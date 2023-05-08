# Views

### Available Operations

* [ViewsOpen](#viewsopen) - Open a view for a user.
* [ViewsPublish](#viewspublish) - Publish a static view for a User.
* [ViewsPush](#viewspush) - Push a view onto the stack of a root view.
* [ViewsUpdate](#viewsupdate) - Update an existing view.

## ViewsOpen

Open a view for a user.

API method documentation
<https://api.slack.com/methods/views.open>

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
    res, err := s.Views.ViewsOpen(ctx, operations.ViewsOpenRequest{
        Token: "repellat",
        TriggerID: "doloribus",
        View: "ullam",
    }, operations.ViewsOpenSecurity{
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

## ViewsPublish

Publish a static view for a User.

API method documentation
<https://api.slack.com/methods/views.publish>

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
    res, err := s.Views.ViewsPublish(ctx, operations.ViewsPublishRequest{
        Hash: slackspec.String("in"),
        Token: "nam",
        UserID: "earum",
        View: "officia",
    }, operations.ViewsPublishSecurity{
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

## ViewsPush

Push a view onto the stack of a root view.

API method documentation
<https://api.slack.com/methods/views.push>

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
    res, err := s.Views.ViewsPush(ctx, operations.ViewsPushRequest{
        Token: "laborum",
        TriggerID: "placeat",
        View: "modi",
    }, operations.ViewsPushSecurity{
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

## ViewsUpdate

Update an existing view.

API method documentation
<https://api.slack.com/methods/views.update>

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
    res, err := s.Views.ViewsUpdate(ctx, operations.ViewsUpdateRequest{
        ExternalID: slackspec.String("voluptatibus"),
        Hash: slackspec.String("molestias"),
        Token: "officiis",
        View: slackspec.String("sapiente"),
        ViewID: slackspec.String("cumque"),
    }, operations.ViewsUpdateSecurity{
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
