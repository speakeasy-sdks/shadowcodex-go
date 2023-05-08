// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminConversationsArchiveSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminConversationsArchiveApplicationJSON struct {
	// The channel to archive.
	ChannelID string `json:"channel_id"`
}

type AdminConversationsArchiveRequest struct {
	RequestBody AdminConversationsArchiveApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.conversations:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminConversationsArchiveResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	AdminConversationsArchiveErrorSchema map[string]map[string]interface{}
	// Typical success response
	AdminConversationsArchiveSchema map[string]map[string]interface{}
}
