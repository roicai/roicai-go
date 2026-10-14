"""Applies the SDK's changes to the code Fern generated.

    python scripts/patch_generated.py build/generated

Each change replaces text Fern is known to write and fails if it is missing,
so a Fern upgrade that changes it stops generation rather than silently
dropping the change:

- List methods request the next page with the `page` token of
  `next_page_url` (internal/page_token.go), not the whole URL.
- Requests send a User-Agent of roicai-go/<version> (internal/version.go).
  Fern sends none when it generates locally.
- Errors lead with the message, code and request ID of the API's error body
  (core/error_message.go), rather than the raw JSON.
"""

from __future__ import annotations

import sys
from pathlib import Path


def main(root: Path) -> None:
    paged = [p for p in root.rglob("client.go") if "next := response.GetNextPageURL()" in p.read_text()]
    if not paged:
        sys.exit("patch_generated: no list method reads next_page_url; did the pagination settings change?")
    for path in paged:
        patch(
            path,
            ("next := response.GetNextPageURL()", "next := internal.PageToken(response.GetNextPageURL())"),
            count=None,
        )
    patch(
        root / "internal/caller.go",
        (
            "\treq.Header.Set(contentTypeHeader, reqContentType)\n",
            "\treq.Header.Set(contentTypeHeader, reqContentType)\n\treq.Header.Set(\"User-Agent\", userAgent)\n",
        ),
    )
    patch(
        root / "core/api_error.go",
        (
            'return fmt.Sprintf("%d: %s", a.StatusCode, a.err.Error())',
            'return fmt.Sprintf("%d: %s", a.StatusCode, errorMessage(a.err))',
        ),
    )


def patch(path: Path, *replacements: tuple[str, str], count: int | None = 1) -> None:
    text = path.read_text()
    for old, new in replacements:
        found = text.count(old)
        if found == 0 or (count is not None and found != count):
            sys.exit(f"patch_generated: expected {count or 'some'} of {old!r} in {path}, found {found}")
        text = text.replace(old, new)
    path.write_text(text)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit(__doc__)
    main(Path(sys.argv[1]))
