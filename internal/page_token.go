// Hand-written: scripts/patch_generated.py makes each list method use this.

package internal

import "net/url"

// PageToken returns the `page` token of a list response's next_page_url, or
// nil on the last page.
//
// List methods request the next page by sending the same request again with
// this token, rather than by following the URL, because the API accepts a
// token only with the filters it was issued for and the URL can leave some
// out.
func PageToken(nextPageURL *string) *string {
	if nextPageURL == nil || *nextPageURL == "" {
		return nil
	}
	parsed, err := url.Parse(*nextPageURL)
	if err != nil {
		return nil
	}
	token := parsed.Query().Get("page")
	if token == "" {
		return nil
	}
	return &token
}

// WithPageToken returns query with its `page` set to token, or query itself
// when token is nil. It copies query rather than changing it, because pages
// of one list can be fetched at the same time and share the first page's
// query.
func WithPageToken(query url.Values, token *string) url.Values {
	if token == nil {
		return query
	}
	copied := make(url.Values, len(query)+1)
	for name, values := range query {
		copied[name] = values
	}
	copied.Set("page", *token)
	return copied
}
