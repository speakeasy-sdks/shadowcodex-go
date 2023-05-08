// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminConversationsSetConversationPrefsSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminConversationsSetConversationPrefsApplicationJSON struct {
	// The channel to set the prefs for
	ChannelID string `json:"channel_id"`
	// The prefs for this channel in a stringified JSON format.
	Prefs string `json:"prefs"`
}

type AdminConversationsSetConversationPrefsRequest struct {
	RequestBody AdminConversationsSetConversationPrefsApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.conversations:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminConversationsSetConversationPrefsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	AdminConversationsSetConversationPrefsErrorSchema map[string]map[string]interface{}
	// Typical success response
	AdminConversationsSetConversationPrefsSchema map[string]map[string]interface{}
}
