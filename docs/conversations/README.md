# Conversations

### Available Operations

* [ConversationsArchive](#conversationsarchive) - Archives a conversation.
* [ConversationsClose](#conversationsclose) - Closes a direct message or multi-person direct message.
* [ConversationsCreate](#conversationscreate) - Initiates a public or private channel-based conversation
* [ConversationsHistory](#conversationshistory) - Fetches a conversation's history of messages and events.
* [ConversationsInfo](#conversationsinfo) - Retrieve information about a conversation.
* [ConversationsInvite](#conversationsinvite) - Invites users to a channel.
* [ConversationsJoin](#conversationsjoin) - Joins an existing conversation.
* [ConversationsKick](#conversationskick) - Removes a user from a conversation.
* [ConversationsLeave](#conversationsleave) - Leaves a conversation.
* [ConversationsList](#conversationslist) - Lists all channels in a Slack team.
* [ConversationsMark](#conversationsmark) - Sets the read cursor in a channel.
* [ConversationsMembers](#conversationsmembers) - Retrieve members of a conversation.
* [ConversationsOpen](#conversationsopen) - Opens or resumes a direct message or multi-person direct message.
* [ConversationsRename](#conversationsrename) - Renames a conversation.
* [ConversationsReplies](#conversationsreplies) - Retrieve a thread of messages posted to a conversation
* [ConversationsSetPurpose](#conversationssetpurpose) - Sets the purpose for a conversation.
* [ConversationsSetTopic](#conversationssettopic) - Sets the topic for a conversation.
* [ConversationsUnarchive](#conversationsunarchive) - Reverses conversation archival.

## ConversationsArchive

Archives a conversation.

API method documentation
<https://api.slack.com/methods/conversations.archive>

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
    res, err := s.Conversations.ConversationsArchive(ctx, operations.ConversationsArchiveRequest{
        RequestBody: &operations.ConversationsArchiveApplicationJSON{
            Channel: slackspec.String("voluptate"),
        },
        Token: slackspec.String("ipsa"),
    }, operations.ConversationsArchiveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsArchiveSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsClose

Closes a direct message or multi-person direct message.

API method documentation
<https://api.slack.com/methods/conversations.close>

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
    res, err := s.Conversations.ConversationsClose(ctx, operations.ConversationsCloseRequest{
        RequestBody: &operations.ConversationsCloseApplicationJSON{
            Channel: slackspec.String("minima"),
        },
        Token: slackspec.String("veritatis"),
    }, operations.ConversationsCloseSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsCloseSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsCreate

Initiates a public or private channel-based conversation

API method documentation
<https://api.slack.com/methods/conversations.create>

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
    res, err := s.Conversations.ConversationsCreate(ctx, operations.ConversationsCreateRequest{
        RequestBody: &operations.ConversationsCreateApplicationJSON{
            IsPrivate: slackspec.Bool(false),
            Name: slackspec.String("Sherry Morar IV"),
        },
        Token: slackspec.String("aut"),
    }, operations.ConversationsCreateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsCreateSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsHistory

Fetches a conversation's history of messages and events.

API method documentation
<https://api.slack.com/methods/conversations.history>

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
    res, err := s.Conversations.ConversationsHistory(ctx, operations.ConversationsHistoryRequest{
        Channel: slackspec.String("laudantium"),
        Cursor: slackspec.String("eum"),
        Inclusive: slackspec.Bool(false),
        Latest: slackspec.Float64(6498.32),
        Limit: slackspec.Int64(68074),
        Oldest: slackspec.Float64(5445.91),
        Token: slackspec.String("non"),
    }, operations.ConversationsHistorySecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsHistorySuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsInfo

Retrieve information about a conversation.

API method documentation
<https://api.slack.com/methods/conversations.info>

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
    res, err := s.Conversations.ConversationsInfo(ctx, operations.ConversationsInfoRequest{
        Channel: slackspec.String("voluptatem"),
        IncludeLocale: slackspec.Bool(false),
        IncludeNumMembers: slackspec.Bool(false),
        Token: slackspec.String("dolor"),
    }, operations.ConversationsInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsInfoSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsInvite

Invites users to a channel.

API method documentation
<https://api.slack.com/methods/conversations.invite>

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
    res, err := s.Conversations.ConversationsInvite(ctx, operations.ConversationsInviteRequest{
        RequestBody: &operations.ConversationsInviteApplicationJSON{
            Channel: slackspec.String("occaecati"),
            Users: slackspec.String("numquam"),
        },
        Token: slackspec.String("impedit"),
    }, operations.ConversationsInviteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsInviteErrorSchema != nil {
        // handle response
    }
}
```

## ConversationsJoin

Joins an existing conversation.

API method documentation
<https://api.slack.com/methods/conversations.join>

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
    res, err := s.Conversations.ConversationsJoin(ctx, operations.ConversationsJoinRequest{
        RequestBody: &operations.ConversationsJoinApplicationJSON{
            Channel: slackspec.String("explicabo"),
        },
        Token: slackspec.String("voluptas"),
    }, operations.ConversationsJoinSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsJoinSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsKick

Removes a user from a conversation.

API method documentation
<https://api.slack.com/methods/conversations.kick>

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
    res, err := s.Conversations.ConversationsKick(ctx, operations.ConversationsKickRequest{
        RequestBody: &operations.ConversationsKickApplicationJSON{
            Channel: slackspec.String("aut"),
            User: slackspec.String("dignissimos"),
        },
        Token: slackspec.String("dicta"),
    }, operations.ConversationsKickSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsKickSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsLeave

Leaves a conversation.

API method documentation
<https://api.slack.com/methods/conversations.leave>

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
    res, err := s.Conversations.ConversationsLeave(ctx, operations.ConversationsLeaveRequest{
        RequestBody: &operations.ConversationsLeaveApplicationJSON{
            Channel: slackspec.String("maiores"),
        },
        Token: slackspec.String("natus"),
    }, operations.ConversationsLeaveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsLeaveSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsList

Lists all channels in a Slack team.

API method documentation
<https://api.slack.com/methods/conversations.list>

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
    res, err := s.Conversations.ConversationsList(ctx, operations.ConversationsListRequest{
        Cursor: slackspec.String("velit"),
        ExcludeArchived: slackspec.Bool(false),
        Limit: slackspec.Int64(974257),
        Token: slackspec.String("voluptas"),
        Types: slackspec.String("asperiores"),
    }, operations.ConversationsListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsListSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsMark

Sets the read cursor in a channel.

API method documentation
<https://api.slack.com/methods/conversations.mark>

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
    res, err := s.Conversations.ConversationsMark(ctx, operations.ConversationsMarkRequest{
        RequestBody: &operations.ConversationsMarkApplicationJSON{
            Channel: slackspec.String("aperiam"),
            Ts: slackspec.Float64(4090.54),
        },
        Token: slackspec.String("quaerat"),
    }, operations.ConversationsMarkSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsMarkSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsMembers

Retrieve members of a conversation.

API method documentation
<https://api.slack.com/methods/conversations.members>

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
    res, err := s.Conversations.ConversationsMembers(ctx, operations.ConversationsMembersRequest{
        Channel: slackspec.String("consequuntur"),
        Cursor: slackspec.String("repellendus"),
        Limit: slackspec.Int64(638762),
        Token: slackspec.String("maxime"),
    }, operations.ConversationsMembersSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsMembersSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsOpen

Opens or resumes a direct message or multi-person direct message.

API method documentation
<https://api.slack.com/methods/conversations.open>

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
    res, err := s.Conversations.ConversationsOpen(ctx, operations.ConversationsOpenRequest{
        RequestBody: &operations.ConversationsOpenApplicationJSON{
            Channel: slackspec.String("dignissimos"),
            ReturnIm: slackspec.Bool(false),
            Users: slackspec.String("officia"),
        },
        Token: slackspec.String("asperiores"),
    }, operations.ConversationsOpenSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsOpenSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsRename

Renames a conversation.

API method documentation
<https://api.slack.com/methods/conversations.rename>

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
    res, err := s.Conversations.ConversationsRename(ctx, operations.ConversationsRenameRequest{
        RequestBody: &operations.ConversationsRenameApplicationJSON{
            Channel: slackspec.String("nemo"),
            Name: slackspec.String("Darlene Sawayn"),
        },
        Token: slackspec.String("ab"),
    }, operations.ConversationsRenameSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsRenameSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsReplies

Retrieve a thread of messages posted to a conversation

API method documentation
<https://api.slack.com/methods/conversations.replies>

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
    res, err := s.Conversations.ConversationsReplies(ctx, operations.ConversationsRepliesRequest{
        Channel: slackspec.String("adipisci"),
        Cursor: slackspec.String("fuga"),
        Inclusive: slackspec.Bool(false),
        Latest: slackspec.Float64(6625.05),
        Limit: slackspec.Int64(380729),
        Oldest: slackspec.Float64(2460.63),
        Token: slackspec.String("culpa"),
        Ts: slackspec.Float64(6658.59),
    }, operations.ConversationsRepliesSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsRepliesSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsSetPurpose

Sets the purpose for a conversation.

API method documentation
<https://api.slack.com/methods/conversations.setPurpose>

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
    res, err := s.Conversations.ConversationsSetPurpose(ctx, operations.ConversationsSetPurposeRequest{
        RequestBody: &operations.ConversationsSetPurposeApplicationJSON{
            Channel: slackspec.String("recusandae"),
            Purpose: slackspec.String("totam"),
        },
        Token: slackspec.String("fugiat"),
    }, operations.ConversationsSetPurposeSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsSetPurposeSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsSetTopic

Sets the topic for a conversation.

API method documentation
<https://api.slack.com/methods/conversations.setTopic>

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
    res, err := s.Conversations.ConversationsSetTopic(ctx, operations.ConversationsSetTopicRequest{
        RequestBody: &operations.ConversationsSetTopicApplicationJSON{
            Channel: slackspec.String("vel"),
            Topic: slackspec.String("ducimus"),
        },
        Token: slackspec.String("quos"),
    }, operations.ConversationsSetTopicSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsSetTopicSuccessSchema != nil {
        // handle response
    }
}
```

## ConversationsUnarchive

Reverses conversation archival.

API method documentation
<https://api.slack.com/methods/conversations.unarchive>

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
    res, err := s.Conversations.ConversationsUnarchive(ctx, operations.ConversationsUnarchiveRequest{
        RequestBody: &operations.ConversationsUnarchiveApplicationJSON{
            Channel: slackspec.String("vel"),
        },
        Token: slackspec.String("labore"),
    }, operations.ConversationsUnarchiveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConversationsUnarchiveSuccessSchema != nil {
        // handle response
    }
}
```
