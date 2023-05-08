# UsersProfile

### Available Operations

* [UsersProfileGet](#usersprofileget) - Retrieves a user's profile information.
* [UsersProfileSet](#usersprofileset) - Set the profile information for a user.

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
    res, err := s.UsersProfile.UsersProfileGet(ctx, operations.UsersProfileGetRequest{
        IncludeLabels: slackspec.Bool(false),
        Token: "ullam",
        User: slackspec.String("quasi"),
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
    res, err := s.UsersProfile.UsersProfileSet(ctx, operations.UsersProfileSetRequest{
        RequestBody: &operations.UsersProfileSetApplicationJSON{
            Name: slackspec.String("Gordon O'Hara"),
            Profile: slackspec.String("animi"),
            User: slackspec.String("ex"),
            Value: slackspec.String("aliquid"),
        },
        Token: "accusantium",
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
