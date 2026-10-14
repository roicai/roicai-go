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
  (core/error_message.go), rather than the raw JSON. A body that isn't the
  API's JSON error, such as a proxy's HTML page, still gives the typed error
  for its status, with a nil Body.
- Fetching pages of one list at the same time doesn't share a query map
  (WithPageToken in internal/page_token.go).
- A request with an empty path parameter fails before it is sent, rather
  than reaching another endpoint (internal/request_checks.go). Trailing
  slashes are trimmed from the base URL.
- A response that asks for a longer wait than the longest backoff, such as a
  429 once the daily quota runs out, isn't retried. X-RateLimit-Reset only
  sets the wait of a 429, since the API sends it on every response.
- ROICAI_API_KEY is only read when neither an API key nor an API key
  function is given.

Fern's own tests of its runtime expect the behaviour these change, so their
expectations change with them.
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
            'return fmt.Sprintf("%d: %s", a.StatusCode, errorMessage(a.StatusCode, a.err))',
        ),
        (
            'return fmt.Sprintf("%d", a.StatusCode)',
            'return fmt.Sprintf("%d: %s", a.StatusCode, errorMessage(a.StatusCode, nil))',
        ),
    )
    patch(
        root / "internal/error_decoder.go",
        (
            "\t\t\t// If we fail to decode the error, we return the\n"
            "\t\t\t// API error as-is.\n"
            "\t\t\treturn apiError\n",
            "\t\t\t// The body isn't the documented error, such as a\n"
            "\t\t\t// proxy's HTML page: keep the typed error, with no Body.\n"
            "\t\t\treturn customError\n",
        ),
    )
    for path in paged:
        patch(
            path,
            (
                "\t\tif pageRequest.Cursor != nil {\n"
                '\t\t\tqueryParams.Set("page", *pageRequest.Cursor)\n'
                "\t\t}\n",
                "\t\tqueryParams := internal.WithPageToken(queryParams, pageRequest.Cursor)\n",
            ),
            count=None,
        )
    patch(
        root / "internal/caller.go",
        (
            "func (c *Caller) Call(ctx context.Context, params *CallParams) (*CallResponse, error) {\n",
            "func (c *Caller) Call(ctx context.Context, params *CallParams) (*CallResponse, error) {\n"
            "\tif err := checkPath(params.Method, params.URL); err != nil {\n"
            "\t\treturn nil, err\n"
            "\t}\n",
        ),
    )
    patch(
        root / "internal/http.go",
        (
            "\t\tif value != \"\" {\n\t\t\treturn value\n",
            "\t\tif value != \"\" {\n\t\t\treturn trimBaseURL(value)\n",
        ),
    )
    patch(
        root / "internal/retrier.go",
        (
            "\t\tdelay, err := r.retryDelay(response, retryAttempt)\n",
            "\t\tif waitsLongerThan(response, maxRetryDelay) {\n"
            "\t\t\treturn nil, retryError\n"
            "\t\t}\n\n"
            "\t\tdelay, err := r.retryDelay(response, retryAttempt)\n",
        ),
        (
            'if rateLimitReset := response.Header.Get("X-RateLimit-Reset"); rateLimitReset != "" {',
            'if rateLimitReset := response.Header.Get("X-RateLimit-Reset"); rateLimitReset != "" && response.StatusCode == http.StatusTooManyRequests {',
        ),
    )
    patch(
        root / "internal/error_decoder_test.go",
        (
            'wantError:      core.NewAPIError(http.StatusNotFound, http.Header{}, errors.New("Resource not found")),',
            'wantError:      &InternalTestNotFoundError{APIError: core.NewAPIError(http.StatusNotFound, http.Header{}, errors.New("Resource not found"))},',
        ),
        (
            "wantError:      core.NewAPIError(http.StatusForbidden, http.Header{}, errors.New(`<html><body>Forbidden</body></html>`)),",
            "wantError:      &InternalTestClientError{APIError: core.NewAPIError(http.StatusForbidden, http.Header{}, errors.New(`<html><body>Forbidden</body></html>`))},",
        ),
    )
    patch(
        root / "internal/retrier_test.go",
        ('require.EqualError(t, err, "500: retry failed")', 'require.EqualError(t, err, "500: Internal Server Error: retry failed")'),
    )
    env_key = [p for p in root.rglob("*.go") if 'options.APIKey = os.Getenv("ROICAI_API_KEY")' in p.read_text()]
    if not env_key:
        sys.exit("patch_generated: no client reads ROICAI_API_KEY; did the auth settings change?")
    for path in env_key:
        patch(
            path,
            (
                'if options.APIKey == "" {\n\t\toptions.APIKey = os.Getenv("ROICAI_API_KEY")',
                'if options.APIKey == "" && options.APIKeyFunc == nil {\n\t\toptions.APIKey = os.Getenv("ROICAI_API_KEY")',
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
