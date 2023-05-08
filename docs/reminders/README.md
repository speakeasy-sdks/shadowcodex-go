# Reminders

### Available Operations

* [RemindersAdd](#remindersadd) - Creates a reminder.
* [RemindersComplete](#reminderscomplete) - Marks a reminder as complete.
* [RemindersDelete](#remindersdelete) - Deletes a reminder.
* [RemindersInfo](#remindersinfo) - Gets information about a reminder.
* [RemindersList](#reminderslist) - Lists all reminders created by or for a given user.

## RemindersAdd

Creates a reminder.

API method documentation
<https://api.slack.com/methods/reminders.add>

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
    res, err := s.Reminders.RemindersAdd(ctx, operations.RemindersAddRequest{
        RequestBody: operations.RemindersAddApplicationJSON{
            Text: "quisquam",
            Time: "repudiandae",
            User: slackspec.String("quasi"),
        },
        Token: "atque",
    }, operations.RemindersAddSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemindersAddSchema != nil {
        // handle response
    }
}
```

## RemindersComplete

Marks a reminder as complete.

API method documentation
<https://api.slack.com/methods/reminders.complete>

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
    res, err := s.Reminders.RemindersComplete(ctx, operations.RemindersCompleteRequest{
        RequestBody: &operations.RemindersCompleteApplicationJSON{
            Reminder: slackspec.String("reprehenderit"),
        },
        Token: slackspec.String("asperiores"),
    }, operations.RemindersCompleteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemindersCompleteSchema != nil {
        // handle response
    }
}
```

## RemindersDelete

Deletes a reminder.

API method documentation
<https://api.slack.com/methods/reminders.delete>

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
    res, err := s.Reminders.RemindersDelete(ctx, operations.RemindersDeleteRequest{
        RequestBody: &operations.RemindersDeleteApplicationJSON{
            Reminder: slackspec.String("totam"),
        },
        Token: slackspec.String("suscipit"),
    }, operations.RemindersDeleteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemindersDeleteSchema != nil {
        // handle response
    }
}
```

## RemindersInfo

Gets information about a reminder.

API method documentation
<https://api.slack.com/methods/reminders.info>

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
    res, err := s.Reminders.RemindersInfo(ctx, operations.RemindersInfoRequest{
        Reminder: slackspec.String("quidem"),
        Token: slackspec.String("maxime"),
    }, operations.RemindersInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemindersInfoSchema != nil {
        // handle response
    }
}
```

## RemindersList

Lists all reminders created by or for a given user.

API method documentation
<https://api.slack.com/methods/reminders.list>

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
    res, err := s.Reminders.RemindersList(ctx, operations.RemindersListRequest{
        Token: slackspec.String("et"),
    }, operations.RemindersListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemindersListSchema != nil {
        // handle response
    }
}
```
