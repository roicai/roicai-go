// Hand-written: scripts/patch_generated.py makes the generated client use
// these.

package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// checkPath returns an error when a path parameter of the request is empty,
// which leaves an empty segment in the URL: GET /v3.0.0/tickers/ would list
// tickers instead of retrieving one.
func checkPath(method string, rawURL string) error {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	if strings.HasSuffix(parsed.Path, "/") || strings.Contains(parsed.Path, "//") {
		return fmt.Errorf("roicai: %s %s: a path parameter, such as the identifier, is empty", method, parsed.Path)
	}
	return nil
}

// trimBaseURL removes trailing slashes from a base URL, since every path the
// client appends starts with one.
func trimBaseURL(baseURL string) string {
	return strings.TrimRight(baseURL, "/")
}

// waitsLongerThan reports whether a response asks the client to wait longer
// than max before retrying, as a 429 does once the daily quota runs out: its
// Retry-After is the time until the quota resets, which can be hours. Retrying
// after max would only fail again.
func waitsLongerThan(response *http.Response, max time.Duration) bool {
	if retryAfter := response.Header.Get("Retry-After"); retryAfter != "" {
		if seconds, err := strconv.Atoi(retryAfter); err == nil {
			return time.Duration(seconds)*time.Second > max
		}
		if at, err := http.ParseTime(retryAfter); err == nil {
			return time.Until(at) > max
		}
	}
	if response.StatusCode == http.StatusTooManyRequests {
		if reset, err := strconv.ParseInt(response.Header.Get("X-RateLimit-Reset"), 10, 64); err == nil {
			return time.Until(time.Unix(reset, 0)) > max
		}
	}
	return false
}
