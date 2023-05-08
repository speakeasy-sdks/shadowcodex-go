# AdminConversationsRestrictAccess

### Available Operations

* [AdminConversationsRestrictAccessAddGroup](#adminconversationsrestrictaccessaddgroup) - Add an allowlist of IDP groups for accessing a channel
* [AdminConversationsRestrictAccessListGroups](#adminconversationsrestrictaccesslistgroups) - List all IDP Groups linked to a channel
* [AdminConversationsRestrictAccessRemoveGroup](#adminconversationsrestrictaccessremovegroup) - Remove a linked IDP group linked from a private channel

## AdminConversationsRestrictAccessAddGroup

Add an allowlist of IDP groups for accessing a channel

API method documentation
<https://api.slack.com/methods/admin.conversations.restrictAccess.addGroup>

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
    res, err := s.AdminConversationsRestrictAccess.AdminConversationsRestrictAccessAddGroup(ctx, operations.AdminConversationsRestrictAccessAddGroupRequestBody{
        ChannelID: "quo",
        GroupID: "illum",
        TeamID: slackspec.String("pariatur"),
        Token: "maxime",
    }, operations.AdminConversationsRestrictAccessAddGroupSecurity{
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

## AdminConversationsRestrictAccessListGroups

List all IDP Groups linked to a channel

API method documentation
<https://api.slack.com/methods/admin.conversations.restrictAccess.listGroups>

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
    res, err := s.AdminConversationsRestrictAccess.AdminConversationsRestrictAccessListGroups(ctx, operations.AdminConversationsRestrictAccessListGroupsRequest{
        ChannelID: "ea",
        TeamID: slackspec.String("excepturi"),
        Token: "odit",
    }, operations.AdminConversationsRestrictAccessListGroupsSecurity{
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

## AdminConversationsRestrictAccessRemoveGroup

Remove a linked IDP group linked from a private channel

API method documentation
<https://api.slack.com/methods/admin.conversations.restrictAccess.removeGroup>

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
    res, err := s.AdminConversationsRestrictAccess.AdminConversationsRestrictAccessRemoveGroup(ctx, operations.AdminConversationsRestrictAccessRemoveGroupRequestBody{
        ChannelID: "ea",
        GroupID: "accusantium",
        TeamID: "ab",
        Token: "maiores",
    }, operations.AdminConversationsRestrictAccessRemoveGroupSecurity{
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
