# AdminConversationsEkm

### Available Operations

* [AdminConversationsEkmListOriginalConnectedChannelInfo](#adminconversationsekmlistoriginalconnectedchannelinfo) - List all disconnected channels—i.e., channels that were once connected to other workspaces and then disconnected—and the corresponding original channel IDs for key revocation with EKM.

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
    res, err := s.AdminConversationsEkm.AdminConversationsEkmListOriginalConnectedChannelInfo(ctx, operations.AdminConversationsEkmListOriginalConnectedChannelInfoRequest{
        ChannelIds: slackspec.String("consequuntur"),
        Cursor: slackspec.String("praesentium"),
        Limit: slackspec.Int64(615560),
        TeamIds: slackspec.String("magni"),
        Token: "sunt",
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
