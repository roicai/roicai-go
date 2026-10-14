# Contributing

This guide is for maintainers of the ROIC.ai Go library. To report a bug or request a feature, open a [GitHub issue](https://github.com/roicai/roicai-go/issues).

## Development

The library is generated from the API's OpenAPI document with [Fern](https://github.com/fern-api/fern)'s Go generator, which runs locally in Docker and needs no Fern account. Don't edit generated files by hand: CI regenerates the library from `openapi.json` and fails on any difference.

These files are written by hand, and generation keeps them:

| File | What it does |
|---|---|
| `internal/version.go` | The module's version, sent in the `User-Agent` as `roicai-go/<version>`. |
| `version.go` | Exports that version as `roicai.Version`. |
| `internal/page_token.go` | Takes the `page` token out of a list's `next_page_url`, so the next page is the same request with that token, on a copy of the query so pages can be fetched concurrently. |
| `core/page_all.go` | `Page.All`, which returns an iterator over every item of every page. |
| `core/error_message.go` | Errors lead with the message, code and request ID of the API's error body, or the status text for any other body. |
| `internal/request_checks.go` | Rejects a request with an empty path parameter, trims the base URL's trailing slash, and skips a retry when the server asks for a wait longer than a minute. |
| `doc.go`, `client/doc.go`, `core/doc.go`, `option/doc.go` | Package documentation for pkg.go.dev. |
| `sdk_test.go`, `live_test.go`, `example_test.go` | Offline tests against a local server, tests against the live API, and the examples shown on pkg.go.dev. |

`scripts/generate.sh` prepares the document before Fern reads it, and patches the generated code afterwards:

- `scripts/prep_spec.py` keeps the v3 operations, drops the `apikey` query parameter, nests the dotted group names, adds Fern's pagination settings, and gives inline schemas short names. The Python library uses the same script.
- `scripts/name_requests.py` names each request type and its parameters' enums after the operation, such as `RetrieveTickerRequest` and `SearchTickersSearchBy`.
- `scripts/patch_generated.py` makes the generated code use the hand-written files above. Each patch fails if Fern's output changes, so a Fern upgrade can't silently drop one.

Test:

```sh
go vet ./...
go test ./...                     # set ROICAI_API_KEY to also run the live tests
```

## Regenerating the library

Install Go, Python 3.10+, Node.js and Docker, then run the script. Fern's CLI and generator versions are pinned in `fern/`:

```sh
scripts/generate.sh               # from https://api.roic.ai/openapi.json
scripts/generate.sh ./spec.json   # from another document
scripts/generate.sh --committed   # from the committed openapi.json, as CI does
```

The **Codegen** workflow does this after every API deploy and opens a pull request when the library changed. Review the `openapi.json` diff in that pull request: a removed or renamed method, group, type or field is a breaking change (see [Releasing](#releasing)).

To upgrade Fern, change the versions in `fern/fern.config.json` and `fern/generators.yml` in a pull request of their own, and review the diff they make.

## Releasing

1. Open a pull request that sets `Version` in `internal/version.go`. The library follows semver: while it is `v0`, a breaking change needs a new minor version.
2. After it merges, publish a GitHub release whose tag is `v` plus that version, such as `v0.1.0`.

The **Release** workflow checks that the tag matches `internal/version.go`, runs the tests, and asks the Go module proxy for the new version, so `go get` and pkg.go.dev see it within minutes. A Go module has no registry to publish to: its tags are the releases, so never move or delete a tag once it's pushed.
