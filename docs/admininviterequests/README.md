# AdminInviteRequests

### Available Operations

* [AdminInviteRequestsApprove](#admininviterequestsapprove) - Approve a workspace invite request.
* [AdminInviteRequestsDeny](#admininviterequestsdeny) - Deny a workspace invite request.
* [AdminInviteRequestsList](#admininviterequestslist) - List all pending workspace invite requests.

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
    res, err := s.AdminInviteRequests.AdminInviteRequestsApprove(ctx, operations.AdminInviteRequestsApproveRequest{
        RequestBody: operations.AdminInviteRequestsApproveApplicationJSON{
            InviteRequestID: "nostrum",
            TeamID: slackspec.String("hic"),
        },
        Token: "recusandae",
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
    res, err := s.AdminInviteRequests.AdminInviteRequestsDeny(ctx, operations.AdminInviteRequestsDenyRequest{
        RequestBody: operations.AdminInviteRequestsDenyApplicationJSON{
            InviteRequestID: "omnis",
            TeamID: slackspec.String("facilis"),
        },
        Token: "perspiciatis",
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
    res, err := s.AdminInviteRequests.AdminInviteRequestsList(ctx, operations.AdminInviteRequestsListRequest{
        Cursor: slackspec.String("voluptatem"),
        Limit: slackspec.Int64(783645),
        TeamID: slackspec.String("consequuntur"),
        Token: "blanditiis",
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
