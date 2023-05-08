// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemindersDeleteSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type RemindersDeleteApplicationJSON struct {
	// The ID of the reminder
	Reminder *string `json:"reminder,omitempty"`
}

type RemindersDeleteRequest struct {
	RequestBody *RemindersDeleteApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `reminders:write`
	Token *string `header:"style=simple,explode=false,name=token"`
}

type RemindersDeleteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	RemindersDeleteErrorSchema map[string]map[string]interface{}
	// Typical success response
	RemindersDeleteSchema map[string]map[string]interface{}
}
