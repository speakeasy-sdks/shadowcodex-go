# Users

### Available Operations

* [UsersConversations](#usersconversations) - List conversations the calling user may access.
* [UsersDeletePhoto](#usersdeletephoto) - Delete the user profile photo
* [UsersGetPresence](#usersgetpresence) - Gets user presence information.
* [UsersIdentity](#usersidentity) - Get a user's identity.
* [UsersInfo](#usersinfo) - Gets information about a user.
* [UsersList](#userslist) - Lists all users in a Slack team.
* [UsersLookupByEmail](#userslookupbyemail) - Find a user with an email address.
* [UsersProfileGet](#usersprofileget) - Retrieves a user's profile information.
* [UsersProfileSet](#usersprofileset) - Set the profile information for a user.
* [UsersSetActive](#userssetactive) - Marked a user as active. Deprecated and non-functional.
* [UsersSetPhoto](#userssetphoto) - Set the user profile photo
* [UsersSetPresence](#userssetpresence) - Manually sets user presence.

## UsersConversations

List conversations the calling user may access.

API method documentation
<https://api.slack.com/methods/users.conversations>

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
    res, err := s.Users.UsersConversations(ctx, operations.UsersConversationsRequest{
        Cursor: slackspec.String("aliquam"),
        ExcludeArchived: slackspec.Bool(false),
        Limit: slackspec.Int64(320565),
        Token: slackspec.String("repellat"),
        Types: slackspec.String("alias"),
        User: slackspec.String("corporis"),
    }, operations.UsersConversationsSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersConversationsSuccessSchema != nil {
        // handle response
    }
}
```

## UsersDeletePhoto

Delete the user profile photo

API method documentation
<https://api.slack.com/methods/users.deletePhoto>

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
    res, err := s.Users.UsersDeletePhoto(ctx, operations.UsersDeletePhotoRequestBody{
        Token: "perspiciatis",
    }, operations.UsersDeletePhotoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersDeletePhotoSchema != nil {
        // handle response
    }
}
```

## UsersGetPresence

Gets user presence information.

API method documentation
<https://api.slack.com/methods/users.getPresence>

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
    res, err := s.Users.UsersGetPresence(ctx, operations.UsersGetPresenceRequest{
        Token: "nihil",
        User: slackspec.String("mollitia"),
    }, operations.UsersGetPresenceSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIMethodUsersGetPresence != nil {
        // handle response
    }
}
```

## UsersIdentity

Get a user's identity.

API method documentation
<https://api.slack.com/methods/users.identity>

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
    res, err := s.Users.UsersIdentity(ctx, operations.UsersIdentityRequest{
        Token: slackspec.String("voluptas"),
    }, operations.UsersIdentitySecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersIdentitySchema != nil {
        // handle response
    }
}
```

## UsersInfo

Gets information about a user.

API method documentation
<https://api.slack.com/methods/users.info>

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
    res, err := s.Users.UsersInfo(ctx, operations.UsersInfoRequest{
        IncludeLocale: slackspec.Bool(false),
        Token: "alias",
        User: slackspec.String("maiores"),
    }, operations.UsersInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersInfoSuccessSchema != nil {
        // handle response
    }
}
```

## UsersList

Lists all users in a Slack team.

API method documentation
<https://api.slack.com/methods/users.list>

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
    res, err := s.Users.UsersList(ctx, operations.UsersListRequest{
        Cursor: slackspec.String("reiciendis"),
        IncludeLocale: slackspec.Bool(false),
        Limit: slackspec.Int64(174658),
        Token: slackspec.String("id"),
    }, operations.UsersListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersListSchema != nil {
        // handle response
    }
}
```

## UsersLookupByEmail

Find a user with an email address.

API method documentation
<https://api.slack.com/methods/users.lookupByEmail>

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
    res, err := s.Users.UsersLookupByEmail(ctx, operations.UsersLookupByEmailRequest{
        Email: "Elda6@gmail.com",
        Token: "recusandae",
    }, operations.UsersLookupByEmailSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersLookupByEmailSuccessSchema != nil {
        // handle response
    }
}
```

## UsersProfileGet

Retrieves a user's profile information.

API method documentation
<https://api.slack.com/methods/users.profile.get>

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
    res, err := s.Users.UsersProfileGet(ctx, operations.UsersProfileGetRequest{
        IncludeLabels: slackspec.Bool(false),
        Token: "omnis",
        User: slackspec.String("quaerat"),
    }, operations.UsersProfileGetSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersProfileGetSchema != nil {
        // handle response
    }
}
```

## UsersProfileSet

Set the profile information for a user.

API method documentation
<https://api.slack.com/methods/users.profile.set>

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
    res, err := s.Users.UsersProfileSet(ctx, operations.UsersProfileSetRequest{
        RequestBody: &operations.UsersProfileSetApplicationJSON{
            Name: slackspec.String("Carla Graham"),
            Profile: slackspec.String("debitis"),
            User: slackspec.String("laudantium"),
            Value: slackspec.String("eum"),
        },
        Token: "nemo",
    }, operations.UsersProfileSetSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersProfileSetSchema != nil {
        // handle response
    }
}
```

## UsersSetActive

Marked a user as active. Deprecated and non-functional.

API method documentation
<https://api.slack.com/methods/users.setActive>

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
    res, err := s.Users.UsersSetActive(ctx, operations.UsersSetActiveRequest{
        Token: "recusandae",
    }, operations.UsersSetActiveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersSetActiveSchema != nil {
        // handle response
    }
}
```

## UsersSetPhoto

Set the user profile photo

API method documentation
<https://api.slack.com/methods/users.setPhoto>

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
    res, err := s.Users.UsersSetPhoto(ctx, operations.UsersSetPhotoRequestBody{
        CropW: slackspec.String("esse"),
        CropX: slackspec.String("provident"),
        CropY: slackspec.String("quis"),
        Image: slackspec.String("eum"),
        Token: "reiciendis",
    }, operations.UsersSetPhotoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersSetPhotoSchema != nil {
        // handle response
    }
}
```

## UsersSetPresence

Manually sets user presence.

API method documentation
<https://api.slack.com/methods/users.setPresence>

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
    res, err := s.Users.UsersSetPresence(ctx, operations.UsersSetPresenceRequest{
        RequestBody: operations.UsersSetPresenceApplicationJSON{
            Presence: "provident",
        },
        Token: "aspernatur",
    }, operations.UsersSetPresenceSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsersSetPresenceSchema != nil {
        // handle response
    }
}
```
