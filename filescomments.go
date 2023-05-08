// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package slackspec

import (
	"context"
	"fmt"
	"net/http"
	"slack-spec/pkg/models/operations"
	"slack-spec/pkg/utils"
	"strings"
)

type filesComments struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newFilesComments(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *filesComments {
	return &filesComments{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// FilesCommentsDelete - Deletes an existing comment on a file.
// https://api.slack.com/methods/files.comments.delete - API method documentation

func (s *filesComments) FilesCommentsDelete(ctx context.Context, request operations.FilesCommentsDeleteRequest, security operations.FilesCommentsDeleteSecurity) (*operations.FilesCommentsDeleteResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/files.comments.delete"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := utils.ConfigureSecurityClient(s.defaultClient, security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.FilesCommentsDeleteResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.FilesCommentsDeleteSchema = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.FilesCommentsDeleteErrorSchema = out
		}
	}

	return res, nil
}
