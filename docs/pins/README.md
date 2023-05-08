# Pins

### Available Operations

* [PinsAdd](#pinsadd) - Pins an item to a channel.
* [PinsList](#pinslist) - Lists items pinned to a channel.
* [PinsRemove](#pinsremove) - Un-pins an item from a channel.

## PinsAdd

Pins an item to a channel.

API method documentation
<https://api.slack.com/methods/pins.add>

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
    res, err := s.Pins.PinsAdd(ctx, operations.PinsAddRequest{
        RequestBody: operations.PinsAddApplicationJSON{
            Channel: "saepe",
            Timestamp: slackspec.String("error"),
        },
        Token: "consequatur",
    }, operations.PinsAddSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PinsAddSchema != nil {
        // handle response
    }
}
```

## PinsList

Lists items pinned to a channel.

API method documentation
<https://api.slack.com/methods/pins.list>

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
    res, err := s.Pins.PinsList(ctx, operations.PinsListRequest{
        Channel: "incidunt",
        Token: "reiciendis",
    }, operations.PinsListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PinsListSuccessSchema != nil {
        // handle response
    }
}
```

## PinsRemove

Un-pins an item from a channel.

API method documentation
<https://api.slack.com/methods/pins.remove>

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
    res, err := s.Pins.PinsRemove(ctx, operations.PinsRemoveRequest{
        RequestBody: operations.PinsRemoveApplicationJSON{
            Channel: "dolorem",
            Timestamp: slackspec.String("harum"),
        },
        Token: "dicta",
    }, operations.PinsRemoveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PinsRemoveSchema != nil {
        // handle response
    }
}
```
