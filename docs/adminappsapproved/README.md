# AdminAppsApproved

### Available Operations

* [AdminAppsApprovedList](#adminappsapprovedlist) - List approved apps for an org or workspace.

## AdminAppsApprovedList

List approved apps for an org or workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.approved.list>

### Example Usage

```go
package main

import(
	"context"
	"log"
	"slack-spec"
	"slack-spec/pkg/models/operations"
)

func main() {
    s := slackspec.New()

    ctx := context.Background()
    res, err := s.AdminAppsApproved.AdminAppsApprovedList(ctx, operations.AdminAppsApprovedListRequest{
        Cursor: slackspec.String("nam"),
        EnterpriseID: slackspec.String("id"),
        Limit: slackspec.Int64(501324),
        TeamID: slackspec.String("deleniti"),
        Token: "sapiente",
    }, operations.AdminAppsApprovedListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DefaultSuccessTemplate != nil {
        // handle response
    }
}
```
