# AdminEmoji

### Available Operations

* [AdminEmojiAdd](#adminemojiadd) - Add an emoji.
* [AdminEmojiAddAlias](#adminemojiaddalias) - Add an emoji alias.
* [AdminEmojiList](#adminemojilist) - List emoji for an Enterprise Grid organization.
* [AdminEmojiRemove](#adminemojiremove) - Remove an emoji across an Enterprise Grid organization
* [AdminEmojiRename](#adminemojirename) - Rename an emoji.

## AdminEmojiAdd

Add an emoji.

API method documentation
<https://api.slack.com/methods/admin.emoji.add>

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
    res, err := s.AdminEmoji.AdminEmojiAdd(ctx, operations.AdminEmojiAddRequestBody{
        Name: "Clyde Kling",
        Token: "eaque",
        URL: "pariatur",
    }, operations.AdminEmojiAddSecurity{
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

## AdminEmojiAddAlias

Add an emoji alias.

API method documentation
<https://api.slack.com/methods/admin.emoji.addAlias>

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
    res, err := s.AdminEmoji.AdminEmojiAddAlias(ctx, operations.AdminEmojiAddAliasRequestBody{
        AliasFor: "nemo",
        Name: "Joseph Steuber DDS",
        Token: "corporis",
    }, operations.AdminEmojiAddAliasSecurity{
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

## AdminEmojiList

List emoji for an Enterprise Grid organization.

API method documentation
<https://api.slack.com/methods/admin.emoji.list>

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
    res, err := s.AdminEmoji.AdminEmojiList(ctx, operations.AdminEmojiListRequest{
        Cursor: slackspec.String("hic"),
        Limit: slackspec.Int64(729991),
        Token: "nobis",
    }, operations.AdminEmojiListSecurity{
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

## AdminEmojiRemove

Remove an emoji across an Enterprise Grid organization

API method documentation
<https://api.slack.com/methods/admin.emoji.remove>

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
    res, err := s.AdminEmoji.AdminEmojiRemove(ctx, operations.AdminEmojiRemoveRequestBody{
        Name: "Beatrice Lebsack II",
        Token: "nesciunt",
    }, operations.AdminEmojiRemoveSecurity{
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

## AdminEmojiRename

Rename an emoji.

API method documentation
<https://api.slack.com/methods/admin.emoji.rename>

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
    res, err := s.AdminEmoji.AdminEmojiRename(ctx, operations.AdminEmojiRenameRequestBody{
        Name: "Dorothy Dach",
        NewName: "dolor",
        Token: "vero",
    }, operations.AdminEmojiRenameSecurity{
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
