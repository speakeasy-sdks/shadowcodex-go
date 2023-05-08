// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminConversationsConvertToPrivateSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminConversationsConvertToPrivateApplicationJSON struct {
	// The channel to convert to private.
	ChannelID string `json:"channel_id"`
}

type AdminConversationsConvertToPrivateRequest struct {
	RequestBody AdminConversationsConvertToPrivateApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.conversations:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminConversationsConvertToPrivateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	AdminConversationsConvertToPrivateErrorSchema map[string]map[string]interface{}
	// Typical success response
	AdminConversationsConvertToPrivateSchema map[string]map[string]interface{}
}
