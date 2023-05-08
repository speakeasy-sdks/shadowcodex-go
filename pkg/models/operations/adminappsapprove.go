// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AdminAppsApproveSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type AdminAppsApproveApplicationJSON struct {
	// The id of the app to approve.
	AppID *string `json:"app_id,omitempty"`
	// The id of the request to approve.
	RequestID *string `json:"request_id,omitempty"`
	TeamID    *string `json:"team_id,omitempty"`
}

type AdminAppsApproveRequest struct {
	RequestBody *AdminAppsApproveApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `admin.apps:write`
	Token string `header:"style=simple,explode=false,name=token"`
}

type AdminAppsApproveResponse struct {
	ContentType string
	// Typical error response
	DefaultErrorTemplate map[string]map[string]interface{}
	// Typical success response
	DefaultSuccessTemplate map[string]map[string]interface{}
	StatusCode             int
	RawResponse            *http.Response
}
