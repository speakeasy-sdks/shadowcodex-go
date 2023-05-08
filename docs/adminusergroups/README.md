# AdminUsergroups

### Available Operations

* [AdminUsergroupsAddChannels](#adminusergroupsaddchannels) - Add one or more default channels to an IDP group.
* [AdminUsergroupsAddTeams](#adminusergroupsaddteams) - Associate one or more default workspaces with an organization-wide IDP group.
* [AdminUsergroupsListChannels](#adminusergroupslistchannels) - List the channels linked to an org-level IDP group (user group).
* [AdminUsergroupsRemoveChannels](#adminusergroupsremovechannels) - Remove one or more default channels from an org-level IDP group (user group).

## AdminUsergroupsAddChannels

Add one or more default channels to an IDP group.

API method documentation
<https://api.slack.com/methods/admin.usergroups.addChannels>

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
    res, err := s.AdminUsergroups.AdminUsergroupsAddChannels(ctx, operations.AdminUsergroupsAddChannelsRequest{
        RequestBody: operations.AdminUsergroupsAddChannelsApplicationJSON{
            ChannelIds: "atque",
            TeamID: slackspec.String("sit"),
            UsergroupID: "fugiat",
        },
        Token: "ab",
    }, operations.AdminUsergroupsAddChannelsSecurity{
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

## AdminUsergroupsAddTeams

Associate one or more default workspaces with an organization-wide IDP group.

API method documentation
<https://api.slack.com/methods/admin.usergroups.addTeams>

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
    res, err := s.AdminUsergroups.AdminUsergroupsAddTeams(ctx, operations.AdminUsergroupsAddTeamsRequest{
        RequestBody: operations.AdminUsergroupsAddTeamsApplicationJSON{
            AutoProvision: slackspec.Bool(false),
            TeamIds: "soluta",
            UsergroupID: "dolorum",
        },
        Token: "iusto",
    }, operations.AdminUsergroupsAddTeamsSecurity{
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

## AdminUsergroupsListChannels

List the channels linked to an org-level IDP group (user group).

API method documentation
<https://api.slack.com/methods/admin.usergroups.listChannels>

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
    res, err := s.AdminUsergroups.AdminUsergroupsListChannels(ctx, operations.AdminUsergroupsListChannelsRequest{
        IncludeNumMembers: slackspec.Bool(false),
        TeamID: slackspec.String("voluptate"),
        Token: "dolorum",
        UsergroupID: "deleniti",
    }, operations.AdminUsergroupsListChannelsSecurity{
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

## AdminUsergroupsRemoveChannels

Remove one or more default channels from an org-level IDP group (user group).

API method documentation
<https://api.slack.com/methods/admin.usergroups.removeChannels>

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
    res, err := s.AdminUsergroups.AdminUsergroupsRemoveChannels(ctx, operations.AdminUsergroupsRemoveChannelsRequest{
        RequestBody: operations.AdminUsergroupsRemoveChannelsApplicationJSON{
            ChannelIds: "omnis",
            UsergroupID: "necessitatibus",
        },
        Token: "distinctio",
    }, operations.AdminUsergroupsRemoveChannelsSecurity{
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
