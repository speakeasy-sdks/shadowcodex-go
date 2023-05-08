# Search

### Available Operations

* [SearchMessages](#searchmessages) - Searches for messages matching a query.

## SearchMessages

Searches for messages matching a query.

API method documentation
<https://api.slack.com/methods/search.messages>

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
    res, err := s.Search.SearchMessages(ctx, operations.SearchMessagesRequest{
        Count: slackspec.Int64(227759),
        Highlight: slackspec.Bool(false),
        Page: slackspec.Int64(826825),
        Query: "ea",
        Sort: slackspec.String("atque"),
        SortDir: slackspec.String("error"),
        Token: "officiis",
    }, operations.SearchMessagesSecurity{
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
