# AdminInviteRequestsApproved

### Available Operations

* [AdminInviteRequestsApprovedList](#admininviterequestsapprovedlist) - List all approved workspace invite requests.

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
    res, err := s.AdminInviteRequestsApproved.AdminInviteRequestsApprovedList(ctx, operations.AdminInviteRequestsApprovedListRequest{
        Cursor: slackspec.String("error"),
        Limit: slackspec.Int64(50370),
        TeamID: slackspec.String("occaecati"),
        Token: "rerum",
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
