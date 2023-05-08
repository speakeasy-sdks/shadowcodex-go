# AppsEventAuthorizations

### Available Operations

* [AppsEventAuthorizationsList](#appseventauthorizationslist) - Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to.

## AppsEventAuthorizationsList

Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to.

API method documentation
<https://api.slack.com/methods/apps.event.authorizations.list>

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
    res, err := s.AppsEventAuthorizations.AppsEventAuthorizationsList(ctx, operations.AppsEventAuthorizationsListRequest{
        Cursor: slackspec.String("qui"),
        EventContext: "cupiditate",
        Limit: slackspec.Int64(807581),
        Token: "pariatur",
    }, operations.AppsEventAuthorizationsListSecurity{
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
