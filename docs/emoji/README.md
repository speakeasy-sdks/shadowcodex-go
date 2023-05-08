# Emoji

### Available Operations

* [EmojiList](#emojilist) - Lists custom emoji for a team.

## EmojiList

Lists custom emoji for a team.

API method documentation
<https://api.slack.com/methods/emoji.list>

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
    res, err := s.Emoji.EmojiList(ctx, operations.EmojiListRequest{
        Token: "aperiam",
    }, operations.EmojiListSecurity{
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
