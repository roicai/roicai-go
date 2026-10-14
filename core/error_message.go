// Hand-written: scripts/patch_generated.py makes APIError.Error use this.

package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// maxBodyInMessage is how much of a response body that isn't the API's JSON
// error, such as a proxy's error page, an error message quotes.
const maxBodyInMessage = 200

// errorMessage returns the message of an API error body, followed by its code
// and request ID, such as
//
//	No ticker matches the supplied identifier. (code: resource_missing, request_id: req_123)
//
// For any other body it returns the status text, followed by the start of the
// body unless that is empty or HTML.
func errorMessage(statusCode int, err error) string {
	text := ""
	if err != nil {
		text = strings.TrimSpace(err.Error())
	}
	var body struct {
		Error *struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	if json.Unmarshal([]byte(text), &body) == nil && body.Error != nil && body.Error.Message != "" {
		var details []string
		if body.Error.Code != "" {
			details = append(details, "code: "+body.Error.Code)
		}
		if body.Error.RequestID != "" {
			details = append(details, "request_id: "+body.Error.RequestID)
		}
		if len(details) == 0 {
			return body.Error.Message
		}
		return fmt.Sprintf("%s (%s)", body.Error.Message, strings.Join(details, ", "))
	}
	status := http.StatusText(statusCode)
	if text == "" || strings.HasPrefix(text, "<") {
		return status
	}
	text = strings.Join(strings.Fields(text), " ")
	if len(text) > maxBodyInMessage {
		text = text[:maxBodyInMessage] + "…"
	}
	if status == "" {
		return text
	}
	return status + ": " + text
}
