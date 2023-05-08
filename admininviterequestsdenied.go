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

type adminInviteRequestsDenied struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAdminInviteRequestsDenied(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *adminInviteRequestsDenied {
	return &adminInviteRequestsDenied{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// AdminInviteRequestsDeniedList - List all denied workspace invite requests.
// https://api.slack.com/methods/admin.inviteRequests.denied.list - API method documentation

func (s *adminInviteRequestsDenied) AdminInviteRequestsDeniedList(ctx context.Context, request operations.AdminInviteRequestsDeniedListRequest, security operations.AdminInviteRequestsDeniedListSecurity) (*operations.AdminInviteRequestsDeniedListResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/admin.inviteRequests.denied.list"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateHeaders(ctx, req, request)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

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

	res := &operations.AdminInviteRequestsDeniedListResponse{
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

			res.DefaultSuccessTemplate = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out map[string]map[string]interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DefaultErrorTemplate = out
		}
	}

	return res, nil
}
