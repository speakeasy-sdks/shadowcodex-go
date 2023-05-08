# AdminConversations

### Available Operations

* [AdminConversationsArchive](#adminconversationsarchive) - Archive a public or private channel.
* [AdminConversationsConvertToPrivate](#adminconversationsconverttoprivate) - Convert a public channel to a private channel.
* [AdminConversationsCreate](#adminconversationscreate) - Create a public or private channel-based conversation.
* [AdminConversationsDelete](#adminconversationsdelete) - Delete a public or private channel.
* [AdminConversationsDisconnectShared](#adminconversationsdisconnectshared) - Disconnect a connected channel from one or more workspaces.
* [AdminConversationsGetConversationPrefs](#adminconversationsgetconversationprefs) - Get conversation preferences for a public or private channel.
* [AdminConversationsGetTeams](#adminconversationsgetteams) - Get all the workspaces a given public or private channel is connected to within this Enterprise org.
* [AdminConversationsInvite](#adminconversationsinvite) - Invite a user to a public or private channel.
* [AdminConversationsRename](#adminconversationsrename) - Rename a public or private channel.
* [AdminConversationsSearch](#adminconversationssearch) - Search for public or private channels in an Enterprise organization.
* [AdminConversationsSetConversationPrefs](#adminconversationssetconversationprefs) - Set the posting permissions for a public or private channel.
* [AdminConversationsSetTeams](#adminconversationssetteams) - Set the workspaces in an Enterprise grid org that connect to a public or private channel.
* [AdminConversationsUnarchive](#adminconversationsunarchive) - Unarchive a public or private channel.

## AdminConversationsArchive

Archive a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.archive>

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
    res, err := s.AdminConversations.AdminConversationsArchive(ctx, operations.AdminConversationsArchiveRequest{
        RequestBody: operations.AdminConversationsArchiveApplicationJSON{
            ChannelID: "magnam",
        },
        Token: "distinctio",
    }, operations.AdminConversationsArchiveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsArchiveSchema != nil {
        // handle response
    }
}
```

## AdminConversationsConvertToPrivate

Convert a public channel to a private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.convertToPrivate>

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
    res, err := s.AdminConversations.AdminConversationsConvertToPrivate(ctx, operations.AdminConversationsConvertToPrivateRequest{
        RequestBody: operations.AdminConversationsConvertToPrivateApplicationJSON{
            ChannelID: "id",
        },
        Token: "labore",
    }, operations.AdminConversationsConvertToPrivateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsConvertToPrivateSchema != nil {
        // handle response
    }
}
```

## AdminConversationsCreate

Create a public or private channel-based conversation.

API method documentation
<https://api.slack.com/methods/admin.conversations.create>

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
    res, err := s.AdminConversations.AdminConversationsCreate(ctx, operations.AdminConversationsCreateRequest{
        RequestBody: operations.AdminConversationsCreateApplicationJSON{
            Description: slackspec.String("labore"),
            IsPrivate: false,
            Name: "Ada Rohan",
            OrgWide: slackspec.Bool(false),
            TeamID: slackspec.String("aspernatur"),
        },
        Token: "architecto",
    }, operations.AdminConversationsCreateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsCreateSchema != nil {
        // handle response
    }
}
```

## AdminConversationsDelete

Delete a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.delete>

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
    res, err := s.AdminConversations.AdminConversationsDelete(ctx, operations.AdminConversationsDeleteRequest{
        RequestBody: operations.AdminConversationsDeleteApplicationJSON{
            ChannelID: "magnam",
        },
        Token: "et",
    }, operations.AdminConversationsDeleteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsDeleteSchema != nil {
        // handle response
    }
}
```

## AdminConversationsDisconnectShared

Disconnect a connected channel from one or more workspaces.

API method documentation
<https://api.slack.com/methods/admin.conversations.disconnectShared>

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
    res, err := s.AdminConversations.AdminConversationsDisconnectShared(ctx, operations.AdminConversationsDisconnectSharedRequest{
        RequestBody: operations.AdminConversationsDisconnectSharedApplicationJSON{
            ChannelID: "excepturi",
            LeavingTeamIds: slackspec.String("ullam"),
        },
        Token: "provident",
    }, operations.AdminConversationsDisconnectSharedSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsRenameSchema != nil {
        // handle response
    }
}
```

## AdminConversationsGetConversationPrefs

Get conversation preferences for a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.getConversationPrefs>

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
    res, err := s.AdminConversations.AdminConversationsGetConversationPrefs(ctx, operations.AdminConversationsGetConversationPrefsRequest{
        ChannelID: "quos",
        Token: "sint",
    }, operations.AdminConversationsGetConversationPrefsSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsGetConversationPrefsSchema != nil {
        // handle response
    }
}
```

## AdminConversationsGetTeams

Get all the workspaces a given public or private channel is connected to within this Enterprise org.

API method documentation
<https://api.slack.com/methods/admin.conversations.getTeams>

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
    res, err := s.AdminConversations.AdminConversationsGetTeams(ctx, operations.AdminConversationsGetTeamsRequest{
        ChannelID: "accusantium",
        Cursor: slackspec.String("mollitia"),
        Limit: slackspec.Int64(968962),
        Token: "mollitia",
    }, operations.AdminConversationsGetTeamsSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsGetTeamsSchema != nil {
        // handle response
    }
}
```

## AdminConversationsInvite

Invite a user to a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.invite>

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
    res, err := s.AdminConversations.AdminConversationsInvite(ctx, operations.AdminConversationsInviteRequest{
        RequestBody: operations.AdminConversationsInviteApplicationJSON{
            ChannelID: "ad",
            UserIds: "eum",
        },
        Token: "dolor",
    }, operations.AdminConversationsInviteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsInviteSchema != nil {
        // handle response
    }
}
```

## AdminConversationsRename

Rename a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.rename>

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
    res, err := s.AdminConversations.AdminConversationsRename(ctx, operations.AdminConversationsRenameRequest{
        RequestBody: operations.AdminConversationsRenameApplicationJSON{
            ChannelID: "necessitatibus",
            Name: "Vivian Boyle",
        },
        Token: "debitis",
    }, operations.AdminConversationsRenameSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsRenameSchema != nil {
        // handle response
    }
}
```

## AdminConversationsSearch

Search for public or private channels in an Enterprise organization.

API method documentation
<https://api.slack.com/methods/admin.conversations.search>

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
    res, err := s.AdminConversations.AdminConversationsSearch(ctx, operations.AdminConversationsSearchRequest{
        Cursor: slackspec.String("eius"),
        Limit: slackspec.Int64(806194),
        Query: slackspec.String("deleniti"),
        SearchChannelTypes: slackspec.String("facilis"),
        Sort: slackspec.String("in"),
        SortDir: slackspec.String("architecto"),
        TeamIds: slackspec.String("architecto"),
        Token: "repudiandae",
    }, operations.AdminConversationsSearchSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsSearchSchema != nil {
        // handle response
    }
}
```

## AdminConversationsSetConversationPrefs

Set the posting permissions for a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.setConversationPrefs>

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
    res, err := s.AdminConversations.AdminConversationsSetConversationPrefs(ctx, operations.AdminConversationsSetConversationPrefsRequest{
        RequestBody: operations.AdminConversationsSetConversationPrefsApplicationJSON{
            ChannelID: "ullam",
            Prefs: "expedita",
        },
        Token: "nihil",
    }, operations.AdminConversationsSetConversationPrefsSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsSetConversationPrefsSchema != nil {
        // handle response
    }
}
```

## AdminConversationsSetTeams

Set the workspaces in an Enterprise grid org that connect to a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.setTeams>

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
    res, err := s.AdminConversations.AdminConversationsSetTeams(ctx, operations.AdminConversationsSetTeamsRequest{
        RequestBody: operations.AdminConversationsSetTeamsApplicationJSON{
            ChannelID: "repellat",
            OrgChannel: slackspec.Bool(false),
            TargetTeamIds: slackspec.String("quibusdam"),
            TeamID: slackspec.String("sed"),
        },
        Token: "saepe",
    }, operations.AdminConversationsSetTeamsSecurity{
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

## AdminConversationsUnarchive

Unarchive a public or private channel.

API method documentation
<https://api.slack.com/methods/admin.conversations.unarchive>

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
    res, err := s.AdminConversations.AdminConversationsUnarchive(ctx, operations.AdminConversationsUnarchiveRequest{
        RequestBody: operations.AdminConversationsUnarchiveApplicationJSON{
            ChannelID: "pariatur",
        },
        Token: "accusantium",
    }, operations.AdminConversationsUnarchiveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AdminConversationsUnarchiveSchema != nil {
        // handle response
    }
}
```
