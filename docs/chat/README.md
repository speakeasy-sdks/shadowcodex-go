# Chat

### Available Operations

* [ChatDelete](#chatdelete) - Deletes a message.
* [ChatDeleteScheduledMessage](#chatdeletescheduledmessage) - Deletes a pending scheduled message from the queue.
* [ChatGetPermalink](#chatgetpermalink) - Retrieve a permalink URL for a specific extant message
* [ChatMeMessage](#chatmemessage) - Share a me message into a channel.
* [ChatPostEphemeral](#chatpostephemeral) - Sends an ephemeral message to a user in a channel.
* [ChatPostMessage](#chatpostmessage) - Sends a message to a channel.
* [ChatScheduleMessage](#chatschedulemessage) - Schedules a message to be sent to a channel.
* [ChatScheduledMessagesList](#chatscheduledmessageslist) - Returns a list of scheduled messages.
* [ChatUnfurl](#chatunfurl) - Provide custom unfurl behavior for user-posted URLs
* [ChatUpdate](#chatupdate) - Updates a message.

## ChatDelete

Deletes a message.

API method documentation
<https://api.slack.com/methods/chat.delete>

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
    res, err := s.Chat.ChatDelete(ctx, operations.ChatDeleteRequest{
        RequestBody: &operations.ChatDeleteApplicationJSON{
            AsUser: slackspec.Bool(false),
            Channel: slackspec.String("aliquid"),
            Ts: slackspec.Float64(934.59),
        },
        Token: slackspec.String("saepe"),
    }, operations.ChatDeleteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatDeleteSuccessSchema != nil {
        // handle response
    }
}
```

## ChatDeleteScheduledMessage

Deletes a pending scheduled message from the queue.

API method documentation
<https://api.slack.com/methods/chat.deleteScheduledMessage>

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
    res, err := s.Chat.ChatDeleteScheduledMessage(ctx, operations.ChatDeleteScheduledMessageRequest{
        RequestBody: operations.ChatDeleteScheduledMessageApplicationJSON{
            AsUser: slackspec.Bool(false),
            Channel: "vel",
            ScheduledMessageID: "harum",
        },
        Token: "molestiae",
    }, operations.ChatDeleteScheduledMessageSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatDeleteScheduledMessageSchema != nil {
        // handle response
    }
}
```

## ChatGetPermalink

Retrieve a permalink URL for a specific extant message

API method documentation
<https://api.slack.com/methods/chat.getPermalink>

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
    res, err := s.Chat.ChatGetPermalink(ctx, operations.ChatGetPermalinkRequest{
        Channel: "rerum",
        MessageTs: "occaecati",
        Token: "minima",
    }, operations.ChatGetPermalinkSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatGetPermalinkSuccessSchema != nil {
        // handle response
    }
}
```

## ChatMeMessage

Share a me message into a channel.

API method documentation
<https://api.slack.com/methods/chat.meMessage>

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
    res, err := s.Chat.ChatMeMessage(ctx, operations.ChatMeMessageRequest{
        RequestBody: &operations.ChatMeMessageApplicationJSON{
            Channel: slackspec.String("distinctio"),
            Text: slackspec.String("eligendi"),
        },
        Token: slackspec.String("sit"),
    }, operations.ChatMeMessageSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatMeMessageSchema != nil {
        // handle response
    }
}
```

## ChatPostEphemeral

Sends an ephemeral message to a user in a channel.

API method documentation
<https://api.slack.com/methods/chat.postEphemeral>

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
    res, err := s.Chat.ChatPostEphemeral(ctx, operations.ChatPostEphemeralRequest{
        RequestBody: operations.ChatPostEphemeralApplicationJSON{
            AsUser: slackspec.Bool(false),
            Attachments: slackspec.String("culpa"),
            Blocks: slackspec.String("tempore"),
            Channel: "adipisci",
            IconEmoji: slackspec.String("cumque"),
            IconURL: slackspec.String("consequuntur"),
            LinkNames: slackspec.Bool(false),
            Parse: slackspec.String("consequatur"),
            Text: slackspec.String("minus"),
            ThreadTs: slackspec.String("quaerat"),
            User: "sapiente",
            Username: slackspec.String("Darlene_Koch"),
        },
        Token: "a",
    }, operations.ChatPostEphemeralSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatPostEphemeralSuccessSchema != nil {
        // handle response
    }
}
```

## ChatPostMessage

Sends a message to a channel.

API method documentation
<https://api.slack.com/methods/chat.postMessage>

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
    res, err := s.Chat.ChatPostMessage(ctx, operations.ChatPostMessageRequest{
        RequestBody: operations.ChatPostMessageApplicationJSON{
            AsUser: slackspec.String("nulla"),
            Attachments: slackspec.String("quas"),
            Blocks: slackspec.String("esse"),
            Channel: "quasi",
            IconEmoji: slackspec.String("a"),
            IconURL: slackspec.String("error"),
            LinkNames: slackspec.Bool(false),
            Mrkdwn: slackspec.Bool(false),
            Parse: slackspec.String("sint"),
            ReplyBroadcast: slackspec.Bool(false),
            Text: slackspec.String("pariatur"),
            ThreadTs: slackspec.String("possimus"),
            UnfurlLinks: slackspec.Bool(false),
            UnfurlMedia: slackspec.Bool(false),
            Username: slackspec.String("Carlotta_Upton8"),
        },
        Token: "consequuntur",
    }, operations.ChatPostMessageSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatPostMessageSuccessSchema != nil {
        // handle response
    }
}
```

## ChatScheduleMessage

Schedules a message to be sent to a channel.

API method documentation
<https://api.slack.com/methods/chat.scheduleMessage>

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
    res, err := s.Chat.ChatScheduleMessage(ctx, operations.ChatScheduleMessageRequest{
        RequestBody: &operations.ChatScheduleMessageApplicationJSON{
            AsUser: slackspec.Bool(false),
            Attachments: slackspec.String("quasi"),
            Blocks: slackspec.String("similique"),
            Channel: slackspec.String("culpa"),
            LinkNames: slackspec.Bool(false),
            Parse: slackspec.String("aliquid"),
            PostAt: slackspec.String("tenetur"),
            ReplyBroadcast: slackspec.Bool(false),
            Text: slackspec.String("quae"),
            ThreadTs: slackspec.Float64(9367.47),
            UnfurlLinks: slackspec.Bool(false),
            UnfurlMedia: slackspec.Bool(false),
        },
        Token: slackspec.String("vel"),
    }, operations.ChatScheduleMessageSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatScheduleMessageSuccessSchema != nil {
        // handle response
    }
}
```

## ChatScheduledMessagesList

Returns a list of scheduled messages.

API method documentation
<https://api.slack.com/methods/chat.scheduledMessages.list>

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
    res, err := s.Chat.ChatScheduledMessagesList(ctx, operations.ChatScheduledMessagesListRequest{
        Channel: slackspec.String("in"),
        Cursor: slackspec.String("eius"),
        Latest: slackspec.Float64(7276.97),
        Limit: slackspec.Int64(849039),
        Oldest: slackspec.Float64(7422.38),
        Token: slackspec.String("accusantium"),
    }, operations.ChatScheduledMessagesListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatScheduledMessagesListSchema != nil {
        // handle response
    }
}
```

## ChatUnfurl

Provide custom unfurl behavior for user-posted URLs

API method documentation
<https://api.slack.com/methods/chat.unfurl>

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
    res, err := s.Chat.ChatUnfurl(ctx, operations.ChatUnfurlRequest{
        RequestBody: operations.ChatUnfurlApplicationJSON{
            Channel: "aliquam",
            Ts: "sapiente",
            Unfurls: slackspec.String("dicta"),
            UserAuthMessage: slackspec.String("ullam"),
            UserAuthRequired: slackspec.Bool(false),
            UserAuthURL: slackspec.String("reprehenderit"),
        },
        Token: "ullam",
    }, operations.ChatUnfurlSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatUnfurlSuccessSchema != nil {
        // handle response
    }
}
```

## ChatUpdate

Updates a message.

API method documentation
<https://api.slack.com/methods/chat.update>

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
    res, err := s.Chat.ChatUpdate(ctx, operations.ChatUpdateRequest{
        RequestBody: operations.ChatUpdateApplicationJSON{
            AsUser: slackspec.String("nisi"),
            Attachments: slackspec.String("aut"),
            Blocks: slackspec.String("voluptatum"),
            Channel: "qui",
            LinkNames: slackspec.String("quibusdam"),
            Parse: slackspec.String("ex"),
            Text: slackspec.String("deleniti"),
            Ts: "itaque",
        },
        Token: "dolorum",
    }, operations.ChatUpdateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChatUpdateSuccessSchema != nil {
        // handle response
    }
}
```
