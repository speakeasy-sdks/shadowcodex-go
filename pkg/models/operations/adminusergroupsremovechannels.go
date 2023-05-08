// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminUsergroupsRemoveChannelsSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminUsergroupsRemoveChannelsApplicationJSON struct {
	// Comma-separated string of channel IDs
	ChannelIds string `json:"channel_ids"`
	// ID of the IDP Group
	UsergroupID string `json:"usergroup_id"`
}

type AdminUsergroupsRemoveChannelsRequest struct {
	RequestBody AdminUsergroupsRemoveChannelsApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.usergroups:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminUsergroupsRemoveChannelsResponse struct {
	ContentType string
	// Typical error response if the token provided is not associated with an Org Admin or Owner
	DefaultErrorTemplate map[string]map[string]interface{}
	// Typical success response
	DefaultSuccessTemplate map[string]map[string]interface{}
	StatusCode             int
	RawResponse            *http.Response
}
