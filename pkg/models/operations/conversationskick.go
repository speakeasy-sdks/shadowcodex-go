// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConversationsKickSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type ConversationsKickApplicationJSON struct {
	// ID of conversation to remove user from.
	Channel *string `json:"channel,omitempty"`
	// User ID to be removed.
	User *string `json:"user,omitempty"`
}

type ConversationsKickRequest struct {
	RequestBody *ConversationsKickApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `conversations:write`
	Token *string `header:"style=simple,explode=false,name=token"`
}

type ConversationsKickResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response when you attempt to kick yourself from a channel
	ConversationsKickErrorSchema map[string]map[string]interface{}
	// Typical success response
	ConversationsKickSuccessSchema map[string]map[string]interface{}
}
