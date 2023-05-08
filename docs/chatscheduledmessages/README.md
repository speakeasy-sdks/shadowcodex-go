# ChatScheduledMessages

### Available Operations

* [ChatScheduledMessagesList](#chatscheduledmessageslist) - Returns a list of scheduled messages.

## ChatScheduledMessagesList

Returns a list of scheduled messages.

API method documentation
<https://api.slack.com/methods/chat.scheduledMessages.list>

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
    res, err := s.ChatScheduledMessages.ChatScheduledMessagesList(ctx, operations.ChatScheduledMessagesListRequest{
        Channel: slackspec.String("architecto"),
        Cursor: slackspec.String("omnis"),
        Latest: slackspec.Float64(9453.02),
        Limit: slackspec.Int64(98478),
        Oldest: slackspec.Float64(8694.89),
        Token: slackspec.String("et"),
    }, operations.ChatScheduledMessagesListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatScheduledMessagesListSchema != nil {
        // handle response
    }
}
```
