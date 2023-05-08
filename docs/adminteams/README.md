# AdminTeams

### Available Operations

* [AdminTeamsCreate](#adminteamscreate) - Create an Enterprise team.
* [AdminTeamsList](#adminteamslist) - List all teams on an Enterprise organization

## AdminTeamsCreate

Create an Enterprise team.

API method documentation
<https://api.slack.com/methods/admin.teams.create>

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
    res, err := s.AdminTeams.AdminTeamsCreate(ctx, operations.AdminTeamsCreateRequest{
        RequestBody: operations.AdminTeamsCreateApplicationJSON{
            TeamDescription: slackspec.String("iste"),
            TeamDiscoverability: slackspec.String("dolorum"),
            TeamDomain: "deleniti",
            TeamName: "pariatur",
        },
        Token: "provident",
    }, operations.AdminTeamsCreateSecurity{
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

## AdminTeamsList

List all teams on an Enterprise organization

API method documentation
<https://api.slack.com/methods/admin.teams.list>

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
    res, err := s.AdminTeams.AdminTeamsList(ctx, operations.AdminTeamsListRequest{
        Cursor: slackspec.String("nobis"),
        Limit: slackspec.Int64(730122),
        Token: "delectus",
    }, operations.AdminTeamsListSecurity{
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
