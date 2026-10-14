// Hand-written: scripts/patch_generated.py makes APIError.Error use this.

package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

// errorMessage returns the message of an API error body, followed by its code
// and request ID, such as
//
//	No ticker matches the supplied identifier. (code: resource_missing, request_id: req_123)
//
// It returns err's own text when that isn't an API error body.
func errorMessage(err error) string {
	var body struct {
		Error *struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	text := err.Error()
	if json.Unmarshal([]byte(text), &body) != nil || body.Error == nil || body.Error.Message == "" {
		return text
	}
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
