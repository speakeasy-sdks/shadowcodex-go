# AdminTeamsOwners

### Available Operations

* [AdminTeamsOwnersList](#adminteamsownerslist) - List all of the owners on a given workspace.

## AdminTeamsOwnersList

List all of the owners on a given workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.owners.list>

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
    res, err := s.AdminTeamsOwners.AdminTeamsOwnersList(ctx, operations.AdminTeamsOwnersListRequest{
        Cursor: slackspec.String("dolorem"),
        Limit: slackspec.Int64(222443),
        TeamID: "qui",
        Token: "ipsum",
    }, operations.AdminTeamsOwnersListSecurity{
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
