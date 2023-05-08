// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FilesCommentsDeleteSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type FilesCommentsDeleteApplicationJSON struct {
	// File to delete a comment from.
	File *string `json:"file,omitempty"`
	// The comment to delete.
	ID *string `json:"id,omitempty"`
}

type FilesCommentsDeleteRequest struct {
	RequestBody *FilesCommentsDeleteApplicationJSON `request:"mediaType=application/json"`
	// Authentication token. Requires scope: `files:write:user`
	Token *string `header:"style=simple,explode=false,name=token"`
}

type FilesCommentsDeleteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Standard failure response when used with an invalid token
	FilesCommentsDeleteErrorSchema map[string]map[string]interface{}
	// Standard success response is very simple
	FilesCommentsDeleteSchema map[string]map[string]interface{}
}
