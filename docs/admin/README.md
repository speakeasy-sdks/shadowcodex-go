# Admin

### Available Operations

* [AdminAppsApprove](#adminappsapprove) - Approve an app for installation on a workspace.
* [AdminAppsApprovedList](#adminappsapprovedlist) - List approved apps for an org or workspace.
* [AdminAppsRequestsList](#adminappsrequestslist) - List app requests for a team/workspace.
* [AdminAppsRestrict](#adminappsrestrict) - Restrict an app for installation on a workspace.
* [AdminAppsRestrictedList](#adminappsrestrictedlist) - List restricted apps for an org or workspace.
* [AdminConversationsArchive](#adminconversationsarchive) - Archive a public or private channel.
* [AdminConversationsConvertToPrivate](#adminconversationsconverttoprivate) - Convert a public channel to a private channel.
* [AdminConversationsCreate](#adminconversationscreate) - Create a public or private channel-based conversation.
* [AdminConversationsDelete](#adminconversationsdelete) - Delete a public or private channel.
* [AdminConversationsDisconnectShared](#adminconversationsdisconnectshared) - Disconnect a connected channel from one or more workspaces.
* [AdminConversationsEkmListOriginalConnectedChannelInfo](#adminconversationsekmlistoriginalconnectedchannelinfo) - List all disconnected channels—i.e., channels that were once connected to other workspaces and then disconnected—and the corresponding original channel IDs for key revocation with EKM.
* [AdminConversationsGetConversationPrefs](#adminconversationsgetconversationprefs) - Get conversation preferences for a public or private channel.
* [AdminConversationsGetTeams](#adminconversationsgetteams) - Get all the workspaces a given public or private channel is connected to within this Enterprise org.
* [AdminConversationsInvite](#adminconversationsinvite) - Invite a user to a public or private channel.
* [AdminConversationsRename](#adminconversationsrename) - Rename a public or private channel.
* [AdminConversationsRestrictAccessAddGroup](#adminconversationsrestrictaccessaddgroup) - Add an allowlist of IDP groups for accessing a channel
* [AdminConversationsRestrictAccessListGroups](#adminconversationsrestrictaccesslistgroups) - List all IDP Groups linked to a channel
* [AdminConversationsRestrictAccessRemoveGroup](#adminconversationsrestrictaccessremovegroup) - Remove a linked IDP group linked from a private channel
* [AdminConversationsSearch](#adminconversationssearch) - Search for public or private channels in an Enterprise organization.
* [AdminConversationsSetConversationPrefs](#adminconversationssetconversationprefs) - Set the posting permissions for a public or private channel.
* [AdminConversationsSetTeams](#adminconversationssetteams) - Set the workspaces in an Enterprise grid org that connect to a public or private channel.
* [AdminConversationsUnarchive](#adminconversationsunarchive) - Unarchive a public or private channel.
* [AdminEmojiAdd](#adminemojiadd) - Add an emoji.
* [AdminEmojiAddAlias](#adminemojiaddalias) - Add an emoji alias.
* [AdminEmojiList](#adminemojilist) - List emoji for an Enterprise Grid organization.
* [AdminEmojiRemove](#adminemojiremove) - Remove an emoji across an Enterprise Grid organization
* [AdminEmojiRename](#adminemojirename) - Rename an emoji.
* [AdminInviteRequestsApprove](#admininviterequestsapprove) - Approve a workspace invite request.
* [AdminInviteRequestsApprovedList](#admininviterequestsapprovedlist) - List all approved workspace invite requests.
* [AdminInviteRequestsDeniedList](#admininviterequestsdeniedlist) - List all denied workspace invite requests.
* [AdminInviteRequestsDeny](#admininviterequestsdeny) - Deny a workspace invite request.
* [AdminInviteRequestsList](#admininviterequestslist) - List all pending workspace invite requests.
* [AdminTeamsAdminsList](#adminteamsadminslist) - List all of the admins on a given workspace.
* [AdminTeamsCreate](#adminteamscreate) - Create an Enterprise team.
* [AdminTeamsList](#adminteamslist) - List all teams on an Enterprise organization
* [AdminTeamsOwnersList](#adminteamsownerslist) - List all of the owners on a given workspace.
* [AdminTeamsSettingsInfo](#adminteamssettingsinfo) - Fetch information about settings in a workspace
* [AdminTeamsSettingsSetDefaultChannels](#adminteamssettingssetdefaultchannels) - Set the default channels of a workspace.
* [AdminTeamsSettingsSetDescription](#adminteamssettingssetdescription) - Set the description of a given workspace.
* [AdminTeamsSettingsSetDiscoverability](#adminteamssettingssetdiscoverability) - An API method that allows admins to set the discoverability of a given workspace
* [AdminTeamsSettingsSetIcon](#adminteamssettingsseticon) - Sets the icon of a workspace.
* [AdminTeamsSettingsSetName](#adminteamssettingssetname) - Set the name of a given workspace.
* [AdminUsergroupsAddChannels](#adminusergroupsaddchannels) - Add one or more default channels to an IDP group.
* [AdminUsergroupsAddTeams](#adminusergroupsaddteams) - Associate one or more default workspaces with an organization-wide IDP group.
* [AdminUsergroupsListChannels](#adminusergroupslistchannels) - List the channels linked to an org-level IDP group (user group).
* [AdminUsergroupsRemoveChannels](#adminusergroupsremovechannels) - Remove one or more default channels from an org-level IDP group (user group).
* [AdminUsersAssign](#adminusersassign) - Add an Enterprise user to a workspace.
* [AdminUsersInvite](#adminusersinvite) - Invite a user to a workspace.
* [AdminUsersList](#adminuserslist) - List users on a workspace
* [AdminUsersRemove](#adminusersremove) - Remove a user from a workspace.
* [AdminUsersSessionInvalidate](#adminuserssessioninvalidate) - Invalidate a single session for a user by session_id
* [AdminUsersSessionReset](#adminuserssessionreset) - Wipes all valid sessions on all devices for a given user
* [AdminUsersSetAdmin](#adminuserssetadmin) - Set an existing guest, regular user, or owner to be an admin user.
* [AdminUsersSetExpiration](#adminuserssetexpiration) - Set an expiration for a guest user
* [AdminUsersSetOwner](#adminuserssetowner) - Set an existing guest, regular user, or admin user to be a workspace owner.
* [AdminUsersSetRegular](#adminuserssetregular) - Set an existing guest user, admin user, or owner to be a regular user.

## AdminAppsApprove

Approve an app for installation on a workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.approve>

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
    res, err := s.Admin.AdminAppsApprove(ctx, operations.AdminAppsApproveRequest{
        RequestBody: &operations.AdminAppsApproveApplicationJSON{
            AppID: slackspec.String("unde"),
            RequestID: slackspec.String("nulla"),
            TeamID: slackspec.String("corrupti"),
        },
        Token: "illum",
    }, operations.AdminAppsApproveSecurity{
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
    res, err := s.Admin.AdminAppsApprovedList(ctx, operations.AdminAppsApprovedListRequest{
        Cursor: slackspec.String("vel"),
        EnterpriseID: slackspec.String("error"),
        Limit: slackspec.Int64(645894),
        TeamID: slackspec.String("suscipit"),
        Token: "iure",
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
    res, err := s.Admin.AdminAppsRequestsList(ctx, operations.AdminAppsRequestsListRequest{
        Cursor: slackspec.String("magnam"),
        Limit: slackspec.Int64(891773),
        TeamID: slackspec.String("ipsa"),
        Token: "delectus",
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

## AdminAppsRestrict

Restrict an app for installation on a workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.restrict>

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
    res, err := s.Admin.AdminAppsRestrict(ctx, operations.AdminAppsRestrictRequest{
        RequestBody: &operations.AdminAppsRestrictApplicationJSON{
            AppID: slackspec.String("tempora"),
            RequestID: slackspec.String("suscipit"),
            TeamID: slackspec.String("molestiae"),
        },
        Token: "minus",
    }, operations.AdminAppsRestrictSecurity{
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
    res, err := s.Admin.AdminAppsRestrictedList(ctx, operations.AdminAppsRestrictedListRequest{
        Cursor: slackspec.String("placeat"),
        EnterpriseID: slackspec.String("voluptatum"),
        Limit: slackspec.Int64(479977),
        TeamID: slackspec.String("excepturi"),
        Token: "nisi",
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
    res, err := s.Admin.AdminConversationsArchive(ctx, operations.AdminConversationsArchiveRequest{
        RequestBody: operations.AdminConversationsArchiveApplicationJSON{
            ChannelID: "recusandae",
        },
        Token: "temporibus",
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
    res, err := s.Admin.AdminConversationsConvertToPrivate(ctx, operations.AdminConversationsConvertToPrivateRequest{
        RequestBody: operations.AdminConversationsConvertToPrivateApplicationJSON{
            ChannelID: "ab",
        },
        Token: "quis",
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
    res, err := s.Admin.AdminConversationsCreate(ctx, operations.AdminConversationsCreateRequest{
        RequestBody: operations.AdminConversationsCreateApplicationJSON{
            Description: slackspec.String("veritatis"),
            IsPrivate: false,
            Name: "Christopher Hills",
            OrgWide: slackspec.Bool(false),
            TeamID: slackspec.String("quo"),
        },
        Token: "odit",
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
    res, err := s.Admin.AdminConversationsDelete(ctx, operations.AdminConversationsDeleteRequest{
        RequestBody: operations.AdminConversationsDeleteApplicationJSON{
            ChannelID: "at",
        },
        Token: "at",
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
    res, err := s.Admin.AdminConversationsDisconnectShared(ctx, operations.AdminConversationsDisconnectSharedRequest{
        RequestBody: operations.AdminConversationsDisconnectSharedApplicationJSON{
            ChannelID: "maiores",
            LeavingTeamIds: slackspec.String("molestiae"),
        },
        Token: "quod",
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

## AdminConversationsEkmListOriginalConnectedChannelInfo

List all disconnected channels—i.e., channels that were once connected to other workspaces and then disconnected—and the corresponding original channel IDs for key revocation with EKM.

API method documentation
<https://api.slack.com/methods/admin.conversations.ekm.listOriginalConnectedChannelInfo>

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
    res, err := s.Admin.AdminConversationsEkmListOriginalConnectedChannelInfo(ctx, operations.AdminConversationsEkmListOriginalConnectedChannelInfoRequest{
        ChannelIds: slackspec.String("quod"),
        Cursor: slackspec.String("esse"),
        Limit: slackspec.Int64(520478),
        TeamIds: slackspec.String("porro"),
        Token: "dolorum",
    }, operations.AdminConversationsEkmListOriginalConnectedChannelInfoSecurity{
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
    res, err := s.Admin.AdminConversationsGetConversationPrefs(ctx, operations.AdminConversationsGetConversationPrefsRequest{
        ChannelID: "dicta",
        Token: "nam",
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
    res, err := s.Admin.AdminConversationsGetTeams(ctx, operations.AdminConversationsGetTeamsRequest{
        ChannelID: "officia",
        Cursor: slackspec.String("occaecati"),
        Limit: slackspec.Int64(143353),
        Token: "deleniti",
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
    res, err := s.Admin.AdminConversationsInvite(ctx, operations.AdminConversationsInviteRequest{
        RequestBody: operations.AdminConversationsInviteApplicationJSON{
            ChannelID: "hic",
            UserIds: "optio",
        },
        Token: "totam",
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
    res, err := s.Admin.AdminConversationsRename(ctx, operations.AdminConversationsRenameRequest{
        RequestBody: operations.AdminConversationsRenameApplicationJSON{
            ChannelID: "beatae",
            Name: "Tanya Gleason",
        },
        Token: "cum",
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
    res, err := s.Admin.AdminConversationsRestrictAccessAddGroup(ctx, operations.AdminConversationsRestrictAccessAddGroupRequestBody{
        ChannelID: "esse",
        GroupID: "ipsum",
        TeamID: slackspec.String("excepturi"),
        Token: "aspernatur",
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
    res, err := s.Admin.AdminConversationsRestrictAccessListGroups(ctx, operations.AdminConversationsRestrictAccessListGroupsRequest{
        ChannelID: "perferendis",
        TeamID: slackspec.String("ad"),
        Token: "natus",
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
    res, err := s.Admin.AdminConversationsRestrictAccessRemoveGroup(ctx, operations.AdminConversationsRestrictAccessRemoveGroupRequestBody{
        ChannelID: "sed",
        GroupID: "iste",
        TeamID: "dolor",
        Token: "natus",
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
    res, err := s.Admin.AdminConversationsSearch(ctx, operations.AdminConversationsSearchRequest{
        Cursor: slackspec.String("laboriosam"),
        Limit: slackspec.Int64(943749),
        Query: slackspec.String("saepe"),
        SearchChannelTypes: slackspec.String("fuga"),
        Sort: slackspec.String("in"),
        SortDir: slackspec.String("corporis"),
        TeamIds: slackspec.String("iste"),
        Token: "iure",
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
    res, err := s.Admin.AdminConversationsSetConversationPrefs(ctx, operations.AdminConversationsSetConversationPrefsRequest{
        RequestBody: operations.AdminConversationsSetConversationPrefsApplicationJSON{
            ChannelID: "saepe",
            Prefs: "quidem",
        },
        Token: "architecto",
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
    res, err := s.Admin.AdminConversationsSetTeams(ctx, operations.AdminConversationsSetTeamsRequest{
        RequestBody: operations.AdminConversationsSetTeamsApplicationJSON{
            ChannelID: "ipsa",
            OrgChannel: slackspec.Bool(false),
            TargetTeamIds: slackspec.String("reiciendis"),
            TeamID: slackspec.String("est"),
        },
        Token: "mollitia",
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
    res, err := s.Admin.AdminConversationsUnarchive(ctx, operations.AdminConversationsUnarchiveRequest{
        RequestBody: operations.AdminConversationsUnarchiveApplicationJSON{
            ChannelID: "laborum",
        },
        Token: "dolores",
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

## AdminEmojiAdd

Add an emoji.

API method documentation
<https://api.slack.com/methods/admin.emoji.add>

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
    res, err := s.Admin.AdminEmojiAdd(ctx, operations.AdminEmojiAddRequestBody{
        Name: "Stacy Champlin",
        Token: "omnis",
        URL: "nemo",
    }, operations.AdminEmojiAddSecurity{
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

## AdminEmojiAddAlias

Add an emoji alias.

API method documentation
<https://api.slack.com/methods/admin.emoji.addAlias>

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
    res, err := s.Admin.AdminEmojiAddAlias(ctx, operations.AdminEmojiAddAliasRequestBody{
        AliasFor: "minima",
        Name: "Brian Kessler",
        Token: "sapiente",
    }, operations.AdminEmojiAddAliasSecurity{
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

## AdminEmojiList

List emoji for an Enterprise Grid organization.

API method documentation
<https://api.slack.com/methods/admin.emoji.list>

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
    res, err := s.Admin.AdminEmojiList(ctx, operations.AdminEmojiListRequest{
        Cursor: slackspec.String("architecto"),
        Limit: slackspec.Int64(652790),
        Token: "dolorem",
    }, operations.AdminEmojiListSecurity{
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

## AdminEmojiRemove

Remove an emoji across an Enterprise Grid organization

API method documentation
<https://api.slack.com/methods/admin.emoji.remove>

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
    res, err := s.Admin.AdminEmojiRemove(ctx, operations.AdminEmojiRemoveRequestBody{
        Name: "Carlos Ziemann",
        Token: "numquam",
    }, operations.AdminEmojiRemoveSecurity{
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

## AdminEmojiRename

Rename an emoji.

API method documentation
<https://api.slack.com/methods/admin.emoji.rename>

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
    res, err := s.Admin.AdminEmojiRename(ctx, operations.AdminEmojiRenameRequestBody{
        Name: "Claudia Krajcik",
        NewName: "quia",
        Token: "quis",
    }, operations.AdminEmojiRenameSecurity{
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

## AdminInviteRequestsApprove

Approve a workspace invite request.

API method documentation
<https://api.slack.com/methods/admin.inviteRequests.approve>

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
    res, err := s.Admin.AdminInviteRequestsApprove(ctx, operations.AdminInviteRequestsApproveRequest{
        RequestBody: operations.AdminInviteRequestsApproveApplicationJSON{
            InviteRequestID: "vitae",
            TeamID: slackspec.String("laborum"),
        },
        Token: "animi",
    }, operations.AdminInviteRequestsApproveSecurity{
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

## AdminInviteRequestsApprovedList

List all approved workspace invite requests.

API method documentation
<https://api.slack.com/methods/admin.inviteRequests.approved.list>

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
    res, err := s.Admin.AdminInviteRequestsApprovedList(ctx, operations.AdminInviteRequestsApprovedListRequest{
        Cursor: slackspec.String("enim"),
        Limit: slackspec.Int64(138183),
        TeamID: slackspec.String("quo"),
        Token: "sequi",
    }, operations.AdminInviteRequestsApprovedListSecurity{
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

## AdminInviteRequestsDeniedList

List all denied workspace invite requests.

API method documentation
<https://api.slack.com/methods/admin.inviteRequests.denied.list>

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
    res, err := s.Admin.AdminInviteRequestsDeniedList(ctx, operations.AdminInviteRequestsDeniedListRequest{
        Cursor: slackspec.String("tenetur"),
        Limit: slackspec.Int64(368725),
        TeamID: slackspec.String("id"),
        Token: "possimus",
    }, operations.AdminInviteRequestsDeniedListSecurity{
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

## AdminInviteRequestsDeny

Deny a workspace invite request.

API method documentation
<https://api.slack.com/methods/admin.inviteRequests.deny>

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
    res, err := s.Admin.AdminInviteRequestsDeny(ctx, operations.AdminInviteRequestsDenyRequest{
        RequestBody: operations.AdminInviteRequestsDenyApplicationJSON{
            InviteRequestID: "aut",
            TeamID: slackspec.String("quasi"),
        },
        Token: "error",
    }, operations.AdminInviteRequestsDenySecurity{
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

## AdminInviteRequestsList

List all pending workspace invite requests.

API method documentation
<https://api.slack.com/methods/admin.inviteRequests.list>

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
    res, err := s.Admin.AdminInviteRequestsList(ctx, operations.AdminInviteRequestsListRequest{
        Cursor: slackspec.String("temporibus"),
        Limit: slackspec.Int64(673660),
        TeamID: slackspec.String("quasi"),
        Token: "reiciendis",
    }, operations.AdminInviteRequestsListSecurity{
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
    res, err := s.Admin.AdminTeamsAdminsList(ctx, operations.AdminTeamsAdminsListRequest{
        Cursor: slackspec.String("voluptatibus"),
        Limit: slackspec.Int64(878194),
        TeamID: "nihil",
        Token: "praesentium",
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
    res, err := s.Admin.AdminTeamsCreate(ctx, operations.AdminTeamsCreateRequest{
        RequestBody: operations.AdminTeamsCreateApplicationJSON{
            TeamDescription: slackspec.String("voluptatibus"),
            TeamDiscoverability: slackspec.String("ipsa"),
            TeamDomain: "omnis",
            TeamName: "voluptate",
        },
        Token: "cum",
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
    res, err := s.Admin.AdminTeamsList(ctx, operations.AdminTeamsListRequest{
        Cursor: slackspec.String("perferendis"),
        Limit: slackspec.Int64(39187),
        Token: "reprehenderit",
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
    res, err := s.Admin.AdminTeamsOwnersList(ctx, operations.AdminTeamsOwnersListRequest{
        Cursor: slackspec.String("ut"),
        Limit: slackspec.Int64(979587),
        TeamID: "dicta",
        Token: "corporis",
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

## AdminTeamsSettingsInfo

Fetch information about settings in a workspace

API method documentation
<https://api.slack.com/methods/admin.teams.settings.info>

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
    res, err := s.Admin.AdminTeamsSettingsInfo(ctx, operations.AdminTeamsSettingsInfoRequest{
        TeamID: "dolore",
        Token: "iusto",
    }, operations.AdminTeamsSettingsInfoSecurity{
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

## AdminTeamsSettingsSetDefaultChannels

Set the default channels of a workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setDefaultChannels>

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
    res, err := s.Admin.AdminTeamsSettingsSetDefaultChannels(ctx, operations.AdminTeamsSettingsSetDefaultChannelsRequestBody{
        ChannelIds: "dicta",
        TeamID: "harum",
        Token: "enim",
    }, operations.AdminTeamsSettingsSetDefaultChannelsSecurity{
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

## AdminTeamsSettingsSetDescription

Set the description of a given workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setDescription>

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
    res, err := s.Admin.AdminTeamsSettingsSetDescription(ctx, operations.AdminTeamsSettingsSetDescriptionRequest{
        RequestBody: operations.AdminTeamsSettingsSetDescriptionApplicationJSON{
            Description: "accusamus",
            TeamID: "commodi",
        },
        Token: "repudiandae",
    }, operations.AdminTeamsSettingsSetDescriptionSecurity{
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

## AdminTeamsSettingsSetDiscoverability

An API method that allows admins to set the discoverability of a given workspace

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setDiscoverability>

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
    res, err := s.Admin.AdminTeamsSettingsSetDiscoverability(ctx, operations.AdminTeamsSettingsSetDiscoverabilityRequest{
        RequestBody: operations.AdminTeamsSettingsSetDiscoverabilityApplicationJSON{
            Discoverability: "quae",
            TeamID: "ipsum",
        },
        Token: "quidem",
    }, operations.AdminTeamsSettingsSetDiscoverabilitySecurity{
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

## AdminTeamsSettingsSetIcon

Sets the icon of a workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setIcon>

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
    res, err := s.Admin.AdminTeamsSettingsSetIcon(ctx, operations.AdminTeamsSettingsSetIconRequestBody{
        ImageURL: "molestias",
        TeamID: "excepturi",
        Token: "pariatur",
    }, operations.AdminTeamsSettingsSetIconSecurity{
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

## AdminTeamsSettingsSetName

Set the name of a given workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setName>

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
    res, err := s.Admin.AdminTeamsSettingsSetName(ctx, operations.AdminTeamsSettingsSetNameRequest{
        RequestBody: operations.AdminTeamsSettingsSetNameApplicationJSON{
            Name: "Irma Ledner DVM",
            TeamID: "sint",
        },
        Token: "veritatis",
    }, operations.AdminTeamsSettingsSetNameSecurity{
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
    res, err := s.Admin.AdminUsergroupsAddChannels(ctx, operations.AdminUsergroupsAddChannelsRequest{
        RequestBody: operations.AdminUsergroupsAddChannelsApplicationJSON{
            ChannelIds: "itaque",
            TeamID: slackspec.String("incidunt"),
            UsergroupID: "enim",
        },
        Token: "consequatur",
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
    res, err := s.Admin.AdminUsergroupsAddTeams(ctx, operations.AdminUsergroupsAddTeamsRequest{
        RequestBody: operations.AdminUsergroupsAddTeamsApplicationJSON{
            AutoProvision: slackspec.Bool(false),
            TeamIds: "est",
            UsergroupID: "quibusdam",
        },
        Token: "explicabo",
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
    res, err := s.Admin.AdminUsergroupsListChannels(ctx, operations.AdminUsergroupsListChannelsRequest{
        IncludeNumMembers: slackspec.Bool(false),
        TeamID: slackspec.String("deserunt"),
        Token: "distinctio",
        UsergroupID: "quibusdam",
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
    res, err := s.Admin.AdminUsergroupsRemoveChannels(ctx, operations.AdminUsergroupsRemoveChannelsRequest{
        RequestBody: operations.AdminUsergroupsRemoveChannelsApplicationJSON{
            ChannelIds: "labore",
            UsergroupID: "modi",
        },
        Token: "qui",
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

## AdminUsersAssign

Add an Enterprise user to a workspace.

API method documentation
<https://api.slack.com/methods/admin.users.assign>

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
    res, err := s.Admin.AdminUsersAssign(ctx, operations.AdminUsersAssignRequest{
        RequestBody: operations.AdminUsersAssignApplicationJSON{
            ChannelIds: slackspec.String("aliquid"),
            IsRestricted: slackspec.Bool(false),
            IsUltraRestricted: slackspec.Bool(false),
            TeamID: "cupiditate",
            UserID: "quos",
        },
        Token: "perferendis",
    }, operations.AdminUsersAssignSecurity{
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

## AdminUsersInvite

Invite a user to a workspace.

API method documentation
<https://api.slack.com/methods/admin.users.invite>

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
    res, err := s.Admin.AdminUsersInvite(ctx, operations.AdminUsersInviteRequest{
        RequestBody: operations.AdminUsersInviteApplicationJSON{
            ChannelIds: "magni",
            CustomMessage: slackspec.String("assumenda"),
            Email: "Abigale_Corkery27@yahoo.com",
            GuestExpirationTs: slackspec.String("facilis"),
            IsRestricted: slackspec.Bool(false),
            IsUltraRestricted: slackspec.Bool(false),
            RealName: slackspec.String("tempore"),
            Resend: slackspec.Bool(false),
            TeamID: "labore",
        },
        Token: "delectus",
    }, operations.AdminUsersInviteSecurity{
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

## AdminUsersList

List users on a workspace

API method documentation
<https://api.slack.com/methods/admin.users.list>

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
    res, err := s.Admin.AdminUsersList(ctx, operations.AdminUsersListRequest{
        Cursor: slackspec.String("eum"),
        Limit: slackspec.Int64(248753),
        TeamID: "eligendi",
        Token: "sint",
    }, operations.AdminUsersListSecurity{
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

## AdminUsersRemove

Remove a user from a workspace.

API method documentation
<https://api.slack.com/methods/admin.users.remove>

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
    res, err := s.Admin.AdminUsersRemove(ctx, operations.AdminUsersRemoveRequest{
        RequestBody: operations.AdminUsersRemoveApplicationJSON{
            TeamID: "aliquid",
            UserID: "provident",
        },
        Token: "necessitatibus",
    }, operations.AdminUsersRemoveSecurity{
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

## AdminUsersSessionInvalidate

Invalidate a single session for a user by session_id

API method documentation
<https://api.slack.com/methods/admin.users.session.invalidate>

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
    res, err := s.Admin.AdminUsersSessionInvalidate(ctx, operations.AdminUsersSessionInvalidateRequest{
        RequestBody: operations.AdminUsersSessionInvalidateApplicationJSON{
            SessionID: 572252,
            TeamID: "officia",
        },
        Token: "dolor",
    }, operations.AdminUsersSessionInvalidateSecurity{
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

## AdminUsersSessionReset

Wipes all valid sessions on all devices for a given user

API method documentation
<https://api.slack.com/methods/admin.users.session.reset>

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
    res, err := s.Admin.AdminUsersSessionReset(ctx, operations.AdminUsersSessionResetRequest{
        RequestBody: operations.AdminUsersSessionResetApplicationJSON{
            MobileOnly: slackspec.Bool(false),
            UserID: "debitis",
            WebOnly: slackspec.Bool(false),
        },
        Token: "a",
    }, operations.AdminUsersSessionResetSecurity{
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

## AdminUsersSetAdmin

Set an existing guest, regular user, or owner to be an admin user.

API method documentation
<https://api.slack.com/methods/admin.users.setAdmin>

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
    res, err := s.Admin.AdminUsersSetAdmin(ctx, operations.AdminUsersSetAdminRequest{
        RequestBody: operations.AdminUsersSetAdminApplicationJSON{
            TeamID: "dolorum",
            UserID: "in",
        },
        Token: "in",
    }, operations.AdminUsersSetAdminSecurity{
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

## AdminUsersSetExpiration

Set an expiration for a guest user

API method documentation
<https://api.slack.com/methods/admin.users.setExpiration>

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
    res, err := s.Admin.AdminUsersSetExpiration(ctx, operations.AdminUsersSetExpirationRequest{
        RequestBody: operations.AdminUsersSetExpirationApplicationJSON{
            ExpirationTs: 846409,
            TeamID: "maiores",
            UserID: "rerum",
        },
        Token: "dicta",
    }, operations.AdminUsersSetExpirationSecurity{
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

## AdminUsersSetOwner

Set an existing guest, regular user, or admin user to be a workspace owner.

API method documentation
<https://api.slack.com/methods/admin.users.setOwner>

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
    res, err := s.Admin.AdminUsersSetOwner(ctx, operations.AdminUsersSetOwnerRequest{
        RequestBody: operations.AdminUsersSetOwnerApplicationJSON{
            TeamID: "magnam",
            UserID: "cumque",
        },
        Token: "facere",
    }, operations.AdminUsersSetOwnerSecurity{
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

## AdminUsersSetRegular

Set an existing guest user, admin user, or owner to be a regular user.

API method documentation
<https://api.slack.com/methods/admin.users.setRegular>

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
    res, err := s.Admin.AdminUsersSetRegular(ctx, operations.AdminUsersSetRegularRequest{
        RequestBody: operations.AdminUsersSetRegularApplicationJSON{
            TeamID: "ea",
            UserID: "aliquid",
        },
        Token: "laborum",
    }, operations.AdminUsersSetRegularSecurity{
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
