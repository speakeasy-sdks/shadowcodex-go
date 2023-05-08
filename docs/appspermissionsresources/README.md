# AppsPermissionsResources

### Available Operations

* [AppsPermissionsResourcesList](#appspermissionsresourceslist) - Returns list of resource grants this app has on a team.

## AppsPermissionsResourcesList

Returns list of resource grants this app has on a team.

API method documentation
<https://api.slack.com/methods/apps.permissions.resources.list>

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
    res, err := s.AppsPermissionsResources.AppsPermissionsResourcesList(ctx, operations.AppsPermissionsResourcesListRequest{
        Cursor: slackspec.String("incidunt"),
        Limit: slackspec.Int64(132068),
        Token: "dolores",
    }, operations.AppsPermissionsResourcesListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppsPermissionsResourcesListSuccessSchema != nil {
        // handle response
    }
}
```
