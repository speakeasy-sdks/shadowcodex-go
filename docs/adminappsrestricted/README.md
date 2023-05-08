# AdminAppsRestricted

### Available Operations

* [AdminAppsRestrictedList](#adminappsrestrictedlist) - List restricted apps for an org or workspace.

## AdminAppsRestrictedList

List restricted apps for an org or workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.restricted.list>

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
    res, err := s.AdminAppsRestricted.AdminAppsRestrictedList(ctx, operations.AdminAppsRestrictedListRequest{
        Cursor: slackspec.String("natus"),
        EnterpriseID: slackspec.String("omnis"),
        Limit: slackspec.Int64(474867),
        TeamID: slackspec.String("perferendis"),
        Token: "nihil",
    }, operations.AdminAppsRestrictedListSecurity{
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
