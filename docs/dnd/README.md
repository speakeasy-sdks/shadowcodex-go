# Dnd

### Available Operations

* [DndEndDnd](#dndenddnd) - Ends the current user's Do Not Disturb session immediately.
* [DndEndSnooze](#dndendsnooze) - Ends the current user's snooze mode immediately.
* [DndInfo](#dndinfo) - Retrieves a user's current Do Not Disturb status.
* [DndSetSnooze](#dndsetsnooze) - Turns on Do Not Disturb mode for the current user, or changes its duration.
* [DndTeamInfo](#dndteaminfo) - Retrieves the Do Not Disturb status for up to 50 users on a team.

## DndEndDnd

Ends the current user's Do Not Disturb session immediately.

API method documentation
<https://api.slack.com/methods/dnd.endDnd>

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
    res, err := s.Dnd.DndEndDnd(ctx, operations.DndEndDndRequest{
        Token: "commodi",
    }, operations.DndEndDndSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DndEndDndSchema != nil {
        // handle response
    }
}
```

## DndEndSnooze

Ends the current user's snooze mode immediately.

API method documentation
<https://api.slack.com/methods/dnd.endSnooze>

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
    res, err := s.Dnd.DndEndSnooze(ctx, operations.DndEndSnoozeRequest{
        Token: "in",
    }, operations.DndEndSnoozeSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DndEndSnoozeSchema != nil {
        // handle response
    }
}
```

## DndInfo

Retrieves a user's current Do Not Disturb status.

API method documentation
<https://api.slack.com/methods/dnd.info>

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
    res, err := s.Dnd.DndInfo(ctx, operations.DndInfoRequest{
        Token: slackspec.String("corporis"),
        User: slackspec.String("reiciendis"),
    }, operations.DndInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DndInfoSchema != nil {
        // handle response
    }
}
```

## DndSetSnooze

Turns on Do Not Disturb mode for the current user, or changes its duration.

API method documentation
<https://api.slack.com/methods/dnd.setSnooze>

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
    res, err := s.Dnd.DndSetSnooze(ctx, operations.DndSetSnoozeRequestBody{
        NumMinutes: "assumenda",
        Token: "nemo",
    }, operations.DndSetSnoozeSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DndSetSnoozeSchema != nil {
        // handle response
    }
}
```

## DndTeamInfo

Retrieves the Do Not Disturb status for up to 50 users on a team.

API method documentation
<https://api.slack.com/methods/dnd.teamInfo>

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
    res, err := s.Dnd.DndTeamInfo(ctx, operations.DndTeamInfoRequest{
        Token: slackspec.String("recusandae"),
        Users: slackspec.String("aliquid"),
    }, operations.DndTeamInfoSecurity{
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
