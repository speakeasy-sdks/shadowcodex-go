# Rtm

### Available Operations

* [RtmConnect](#rtmconnect) - Starts a Real Time Messaging session.

## RtmConnect

Starts a Real Time Messaging session.

API method documentation
<https://api.slack.com/methods/rtm.connect>

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
    res, err := s.Rtm.RtmConnect(ctx, operations.RtmConnectRequest{
        BatchPresenceAware: slackspec.Bool(false),
        PresenceSub: slackspec.Bool(false),
        Token: "esse",
    }, operations.RtmConnectSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RtmConnectSchema != nil {
        // handle response
    }
}
```
