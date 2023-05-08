# AdminUsers

### Available Operations

* [AdminUsersAssign](#adminusersassign) - Add an Enterprise user to a workspace.
* [AdminUsersInvite](#adminusersinvite) - Invite a user to a workspace.
* [AdminUsersList](#adminuserslist) - List users on a workspace
* [AdminUsersRemove](#adminusersremove) - Remove a user from a workspace.
* [AdminUsersSetAdmin](#adminuserssetadmin) - Set an existing guest, regular user, or owner to be an admin user.
* [AdminUsersSetExpiration](#adminuserssetexpiration) - Set an expiration for a guest user
* [AdminUsersSetOwner](#adminuserssetowner) - Set an existing guest, regular user, or admin user to be a workspace owner.
* [AdminUsersSetRegular](#adminuserssetregular) - Set an existing guest user, admin user, or owner to be a regular user.

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
    res, err := s.AdminUsers.AdminUsersAssign(ctx, operations.AdminUsersAssignRequest{
        RequestBody: operations.AdminUsersAssignApplicationJSON{
            ChannelIds: slackspec.String("asperiores"),
            IsRestricted: slackspec.Bool(false),
            IsUltraRestricted: slackspec.Bool(false),
            TeamID: "nihil",
            UserID: "ipsum",
        },
        Token: "voluptate",
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
    res, err := s.AdminUsers.AdminUsersInvite(ctx, operations.AdminUsersInviteRequest{
        RequestBody: operations.AdminUsersInviteApplicationJSON{
            ChannelIds: "id",
            CustomMessage: slackspec.String("saepe"),
            Email: "Brigitte75@gmail.com",
            GuestExpirationTs: slackspec.String("accusamus"),
            IsRestricted: slackspec.Bool(false),
            IsUltraRestricted: slackspec.Bool(false),
            RealName: slackspec.String("ad"),
            Resend: slackspec.Bool(false),
            TeamID: "saepe",
        },
        Token: "suscipit",
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
    res, err := s.AdminUsers.AdminUsersList(ctx, operations.AdminUsersListRequest{
        Cursor: slackspec.String("deserunt"),
        Limit: slackspec.Int64(588317),
        TeamID: "minima",
        Token: "repellendus",
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
    res, err := s.AdminUsers.AdminUsersRemove(ctx, operations.AdminUsersRemoveRequest{
        RequestBody: operations.AdminUsersRemoveApplicationJSON{
            TeamID: "totam",
            UserID: "similique",
        },
        Token: "alias",
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
    res, err := s.AdminUsers.AdminUsersSetAdmin(ctx, operations.AdminUsersSetAdminRequest{
        RequestBody: operations.AdminUsersSetAdminApplicationJSON{
            TeamID: "at",
            UserID: "quaerat",
        },
        Token: "tempora",
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
    res, err := s.AdminUsers.AdminUsersSetExpiration(ctx, operations.AdminUsersSetExpirationRequest{
        RequestBody: operations.AdminUsersSetExpirationApplicationJSON{
            ExpirationTs: 425451,
            TeamID: "quod",
            UserID: "officiis",
        },
        Token: "qui",
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
    res, err := s.AdminUsers.AdminUsersSetOwner(ctx, operations.AdminUsersSetOwnerRequest{
        RequestBody: operations.AdminUsersSetOwnerApplicationJSON{
            TeamID: "dolorum",
            UserID: "a",
        },
        Token: "esse",
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
    res, err := s.AdminUsers.AdminUsersSetRegular(ctx, operations.AdminUsersSetRegularRequest{
        RequestBody: operations.AdminUsersSetRegularApplicationJSON{
            TeamID: "harum",
            UserID: "iusto",
        },
        Token: "ipsum",
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
