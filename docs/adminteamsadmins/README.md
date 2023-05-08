# AdminTeamsAdmins

### Available Operations

* [AdminTeamsAdminsList](#adminteamsadminslist) - List all of the admins on a given workspace.

## AdminTeamsAdminsList

List all of the admins on a given workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.admins.list>

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
    res, err := s.AdminTeamsAdmins.AdminTeamsAdminsList(ctx, operations.AdminTeamsAdminsListRequest{
        Cursor: slackspec.String("quaerat"),
        Limit: slackspec.Int64(554242),
        TeamID: "aliquid",
        Token: "dolorem",
    }, operations.AdminTeamsAdminsListSecurity{
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
