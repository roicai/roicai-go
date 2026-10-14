#!/usr/bin/env bash
# Regenerates the SDK from the ROIC.ai OpenAPI document.
#
#   scripts/generate.sh                  # fetch https://api.roic.ai/openapi.json
#   scripts/generate.sh path/or/url      # another document
#   scripts/generate.sh --committed      # the committed openapi.json, as CI checks
#
# Needs Python 3.10+, Node.js (for npx), Docker and Go. Fern runs its generator
# in Docker, at the versions pinned in fern/; it needs no Fern account.
set -euo pipefail

cd "$(dirname "$0")/.."
source="${1:-https://api.roic.ai/openapi.json}"
python="${PYTHON:-python3}"

# Go files written by hand; generation keeps them and replaces every other one.
hand_written=(
  version.go
  example_test.go
  sdk_test.go
  live_test.go
  core/error_message.go
  core/page_all.go
  internal/page_token.go
  internal/version.go
)

fetched="$(mktemp)"
trap 'rm -f "$fetched"' EXIT

if [[ "$source" == "--committed" ]]; then
  cp openapi.json "$fetched"
elif [[ "$source" == http* ]]; then
  curl -fsSL "$source" -o "$fetched"
else
  cp "$source" "$fetched"
fi

# An operation without SDK names would get a made-up name, and a document with
# none would replace every method.
"$python" - "$fetched" <<'PY'
import json, sys
doc = json.load(open(sys.argv[1]))
ops = [op for path, item in doc["paths"].items() if path.startswith("/v3") for op in item.values() if isinstance(op, dict)]
unnamed = [op.get("operationId", "?") for op in ops if "x-fern-sdk-group-name" not in op]
if not ops or unnamed:
    sys.exit(f"generate: v3 operations without x-fern-sdk-group-name: {unnamed or 'no operations'}")
json.dump(doc, open(sys.argv[1], "w"), indent=2, ensure_ascii=False)
open(sys.argv[1], "a").write("\n")
PY
cp "$fetched" openapi.json

rm -rf build
mkdir -p build
"$python" scripts/prep_spec.py openapi.json build/openapi.sdk.json
"$python" scripts/name_requests.py build/openapi.sdk.json

fern_version="$("$python" -c 'import json; print(json.load(open("fern/fern.config.json"))["version"])')"
# Fern asks before writing into a directory that has files; build/ is new.
npx --yes "fern-api@$fern_version" generate --local --log-level warn < /dev/null

"$python" scripts/patch_generated.py build/generated
rm -rf build/generated/README.md build/generated/CONTRIBUTING.md build/generated/.fern

# Replace the generated code, keeping the hand-written files. Fern's output is
# the module root, so it sits next to scripts/, fern/ and the docs.
keep="$(mktemp -d)"
for file in "${hand_written[@]}"; do
  if [[ -e "$file" ]]; then
    mkdir -p "$keep/$(dirname "$file")"
    cp "$file" "$keep/$file"
  fi
done
find . -name '*.go' -not -path './.git/*' -not -path './build/*' -delete
find . -type d -empty -not -path './.git/*' -delete
rm -f go.mod go.sum reference.md
cp -a build/generated/. .
cp -a "$keep/." .
rm -rf "$keep" build

go mod tidy
gofmt -l . | (! grep .) || { echo "generate: files above are not gofmt-clean" >&2; exit 1; }
go build ./...
