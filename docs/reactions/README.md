# Reactions

### Available Operations

* [ReactionsAdd](#reactionsadd) - Adds a reaction to an item.
* [ReactionsGet](#reactionsget) - Gets reactions for an item.
* [ReactionsList](#reactionslist) - Lists reactions made by a user.
* [ReactionsRemove](#reactionsremove) - Removes a reaction from an item.

## ReactionsAdd

Adds a reaction to an item.

API method documentation
<https://api.slack.com/methods/reactions.add>

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
    res, err := s.Reactions.ReactionsAdd(ctx, operations.ReactionsAddRequest{
        RequestBody: operations.ReactionsAddApplicationJSON{
            Channel: "architecto",
            Name: "Francisco Powlowski",
            Timestamp: "nam",
        },
        Token: "tenetur",
    }, operations.ReactionsAddSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReactionsAddSchema != nil {
        // handle response
    }
}
```

## ReactionsGet

Gets reactions for an item.

API method documentation
<https://api.slack.com/methods/reactions.get>

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
    res, err := s.Reactions.ReactionsGet(ctx, operations.ReactionsGetRequest{
        Channel: slackspec.String("laboriosam"),
        File: slackspec.String("alias"),
        FileComment: slackspec.String("amet"),
        Full: slackspec.Bool(false),
        Timestamp: slackspec.String("deserunt"),
        Token: "voluptate",
    }, operations.ReactionsGetSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReactionsGetSuccessSchema != nil {
        // handle response
    }
}
```

## ReactionsList

Lists reactions made by a user.

API method documentation
<https://api.slack.com/methods/reactions.list>

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
    res, err := s.Reactions.ReactionsList(ctx, operations.ReactionsListRequest{
        Count: slackspec.Int64(600392),
        Cursor: slackspec.String("reiciendis"),
        Full: slackspec.Bool(false),
        Limit: slackspec.Int64(588740),
        Page: slackspec.Int64(833819),
        Token: "delectus",
        User: slackspec.String("voluptates"),
    }, operations.ReactionsListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReactionsListSchema != nil {
        // handle response
    }
}
```

## ReactionsRemove

Removes a reaction from an item.

API method documentation
<https://api.slack.com/methods/reactions.remove>

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
    res, err := s.Reactions.ReactionsRemove(ctx, operations.ReactionsRemoveRequest{
        RequestBody: operations.ReactionsRemoveApplicationJSON{
            Channel: slackspec.String("perferendis"),
            File: slackspec.String("est"),
            FileComment: slackspec.String("quidem"),
            Name: "Chelsea Pfannerstill",
            Timestamp: slackspec.String("veniam"),
        },
        Token: "voluptatem",
    }, operations.ReactionsRemoveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReactionsRemoveSchema != nil {
        // handle response
    }
}
```
