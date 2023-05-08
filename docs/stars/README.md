# Stars

### Available Operations

* [StarsAdd](#starsadd) - Adds a star to an item.
* [StarsList](#starslist) - Lists stars for a user.
* [StarsRemove](#starsremove) - Removes a star from an item.

## StarsAdd

Adds a star to an item.

API method documentation
<https://api.slack.com/methods/stars.add>

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
    res, err := s.Stars.StarsAdd(ctx, operations.StarsAddRequest{
        RequestBody: &operations.StarsAddApplicationJSON{
            Channel: slackspec.String("officiis"),
            File: slackspec.String("accusamus"),
            FileComment: slackspec.String("natus"),
            Timestamp: slackspec.String("minima"),
        },
        Token: "aspernatur",
    }, operations.StarsAddSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StarsAddSchema != nil {
        // handle response
    }
}
```

## StarsList

Lists stars for a user.

API method documentation
<https://api.slack.com/methods/stars.list>

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
    res, err := s.Stars.StarsList(ctx, operations.StarsListRequest{
        Count: slackspec.String("ex"),
        Cursor: slackspec.String("maiores"),
        Limit: slackspec.Int64(544647),
        Page: slackspec.String("at"),
        Token: slackspec.String("error"),
    }, operations.StarsListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StarsListSchema != nil {
        // handle response
    }
}
```

## StarsRemove

Removes a star from an item.

API method documentation
<https://api.slack.com/methods/stars.remove>

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
    res, err := s.Stars.StarsRemove(ctx, operations.StarsRemoveRequest{
        RequestBody: &operations.StarsRemoveApplicationJSON{
            Channel: slackspec.String("blanditiis"),
            File: slackspec.String("suscipit"),
            FileComment: slackspec.String("repudiandae"),
            Timestamp: slackspec.String("atque"),
        },
        Token: "atque",
    }, operations.StarsRemoveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StarsRemoveSchema != nil {
        // handle response
    }
}
```
