// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UsersInfoSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type UsersInfoRequest struct {
	// Set this to `true` to receive the locale for this user. Defaults to `false`
	IncludeLocale *bool `queryParam:"style=form,explode=true,name=include_locale"`
	// Authentication token. Requires scope: `users:read`
	Token string `queryParam:"style=form,explode=true,name=token"`
	// User to get info on
	User *string `queryParam:"style=form,explode=true,name=user"`
}

type UsersInfoResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	UsersInfoErrorSchema map[string]map[string]interface{}
	// Typical success response
	UsersInfoSuccessSchema map[string]map[string]interface{}
}
