// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminTeamsCreateSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminTeamsCreateApplicationJSON struct {
	// Description for the team.
	TeamDescription *string `json:"team_description,omitempty"`
	// Who can join the team. A team's discoverability can be `open`, `closed`, `invite_only`, or `unlisted`.
	TeamDiscoverability *string `json:"team_discoverability,omitempty"`
	// Team domain (for example, slacksoftballteam).
	TeamDomain string `json:"team_domain"`
	// Team name (for example, Slack Softball Team).
	TeamName string `json:"team_name"`
}

type AdminTeamsCreateRequest struct {
	RequestBody AdminTeamsCreateApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.teams:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminTeamsCreateResponse struct {
	ContentType string
	// Typical error response
	DefaultErrorTemplate map[string]map[string]interface{}
	// Typical success response
	DefaultSuccessTemplate map[string]map[string]interface{}
	StatusCode             int
	RawResponse            *http.Response
}
