# Bots

### Available Operations

* [BotsInfo](#botsinfo) - Gets information about a bot user.

## BotsInfo

Gets information about a bot user.

API method documentation
<https://api.slack.com/methods/bots.info>

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
    res, err := s.Bots.BotsInfo(ctx, operations.BotsInfoRequest{
        Bot: slackspec.String("quam"),
        Token: "molestias",
    }, operations.BotsInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BotsInfoSchema != nil {
        // handle response
    }
}
```
