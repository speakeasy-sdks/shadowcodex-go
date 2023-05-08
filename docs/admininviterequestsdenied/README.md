# AdminInviteRequestsDenied

### Available Operations

* [AdminInviteRequestsDeniedList](#admininviterequestsdeniedlist) - List all denied workspace invite requests.

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
    res, err := s.AdminInviteRequestsDenied.AdminInviteRequestsDeniedList(ctx, operations.AdminInviteRequestsDeniedListRequest{
        Cursor: slackspec.String("adipisci"),
        Limit: slackspec.Int64(992397),
        TeamID: slackspec.String("earum"),
        Token: "modi",
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
