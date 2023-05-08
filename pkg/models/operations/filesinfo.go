// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FilesInfoSecurity struct {
	SlackAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type FilesInfoRequest struct {
	Count *string `queryParam:"style=form,explode=true,name=count"`
	// Parameter for pagination. File comments are paginated for a single file. Set `cursor` equal to the `next_cursor` attribute returned by the previous request's `response_metadata`. This parameter is optional, but pagination is mandatory: the default value simply fetches the first "page" of the collection of comments. See [pagination](/docs/pagination) for more details.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Specify a file by providing its ID.
	File *string `queryParam:"style=form,explode=true,name=file"`
	// The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn't been reached.
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Page  *string `queryParam:"style=form,explode=true,name=page"`
	// Authentication token. Requires scope: `files:read`
	Token *string `queryParam:"style=form,explode=true,name=token"`
}

type FilesInfoResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Typical error response
	FilesInfoErrorSchema map[string]map[string]interface{}
	// Typical success response
	FilesInfoSchema map[string]map[string]interface{}
}
