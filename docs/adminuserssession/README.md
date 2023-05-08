# AdminUsersSession

### Available Operations

* [AdminUsersSessionInvalidate](#adminuserssessioninvalidate) - Invalidate a single session for a user by session_id
* [AdminUsersSessionReset](#adminuserssessionreset) - Wipes all valid sessions on all devices for a given user

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
    res, err := s.AdminUsersSession.AdminUsersSessionInvalidate(ctx, operations.AdminUsersSessionInvalidateRequest{
        RequestBody: operations.AdminUsersSessionInvalidateApplicationJSON{
            SessionID: 788740,
            TeamID: "tenetur",
        },
        Token: "amet",
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
    res, err := s.AdminUsersSession.AdminUsersSessionReset(ctx, operations.AdminUsersSessionResetRequest{
        RequestBody: operations.AdminUsersSessionResetApplicationJSON{
            MobileOnly: slackspec.Bool(false),
            UserID: "tempore",
            WebOnly: slackspec.Bool(false),
        },
        Token: "accusamus",
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
