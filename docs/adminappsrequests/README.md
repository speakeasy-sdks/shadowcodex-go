# AdminAppsRequests

### Available Operations

* [AdminAppsRequestsList](#adminappsrequestslist) - List app requests for a team/workspace.

## AdminAppsRequestsList

List app requests for a team/workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.requests.list>

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
    res, err := s.AdminAppsRequests.AdminAppsRequestsList(ctx, operations.AdminAppsRequestsListRequest{
        Cursor: slackspec.String("amet"),
        Limit: slackspec.Int64(643990),
        TeamID: slackspec.String("nisi"),
        Token: "vel",
    }, operations.AdminAppsRequestsListSecurity{
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
