// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminConversationsSetTeamsSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminConversationsSetTeamsApplicationJSON struct {
	// The encoded `channel_id` to add or remove to workspaces.
	ChannelID string `json:"channel_id"`
	// True if channel has to be converted to an org channel
	OrgChannel *bool `json:"org_channel,omitempty"`
	// A comma-separated list of workspaces to which the channel should be shared. Not required if the channel is being shared org-wide.
	TargetTeamIds *string `json:"target_team_ids,omitempty"`
	// The workspace to which the channel belongs. Omit this argument if the channel is a cross-workspace shared channel.
	TeamID *string `json:"team_id,omitempty"`
}

type AdminConversationsSetTeamsRequest struct {
	RequestBody AdminConversationsSetTeamsApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.conversations:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminConversationsSetTeamsResponse struct {
	ContentType string
	// Typical error response
	DefaultErrorTemplate map[string]map[string]interface{}
	// Typical success response
	DefaultSuccessTemplate map[string]map[string]interface{}
	StatusCode             int
	RawResponse            *http.Response
}
