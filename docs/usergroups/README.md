# Usergroups

### Available Operations

* [UsergroupsCreate](#usergroupscreate) - Create a User Group
* [UsergroupsDisable](#usergroupsdisable) - Disable an existing User Group
* [UsergroupsEnable](#usergroupsenable) - Enable a User Group
* [UsergroupsList](#usergroupslist) - List all User Groups for a team
* [UsergroupsUpdate](#usergroupsupdate) - Update an existing User Group
* [UsergroupsUsersList](#usergroupsuserslist) - List all users in a User Group
* [UsergroupsUsersUpdate](#usergroupsusersupdate) - Update the list of users for a User Group

## UsergroupsCreate

Create a User Group

API method documentation
<https://api.slack.com/methods/usergroups.create>

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
    res, err := s.Usergroups.UsergroupsCreate(ctx, operations.UsergroupsCreateRequest{
        RequestBody: operations.UsergroupsCreateApplicationJSON{
            Channels: slackspec.String("consequuntur"),
            Description: slackspec.String("occaecati"),
            Handle: slackspec.String("officiis"),
            IncludeCount: slackspec.Bool(false),
            Name: "Arnold Ferry",
        },
        Token: "consequuntur",
    }, operations.UsergroupsCreateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsCreateSchema != nil {
        // handle response
    }
}
```

## UsergroupsDisable

Disable an existing User Group

API method documentation
<https://api.slack.com/methods/usergroups.disable>

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
    res, err := s.Usergroups.UsergroupsDisable(ctx, operations.UsergroupsDisableRequest{
        RequestBody: operations.UsergroupsDisableApplicationJSON{
            IncludeCount: slackspec.Bool(false),
            Usergroup: "fugit",
        },
        Token: "id",
    }, operations.UsergroupsDisableSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsDisableSchema != nil {
        // handle response
    }
}
```

## UsergroupsEnable

Enable a User Group

API method documentation
<https://api.slack.com/methods/usergroups.enable>

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
    res, err := s.Usergroups.UsergroupsEnable(ctx, operations.UsergroupsEnableRequest{
        RequestBody: operations.UsergroupsEnableApplicationJSON{
            IncludeCount: slackspec.Bool(false),
            Usergroup: "quis",
        },
        Token: "reprehenderit",
    }, operations.UsergroupsEnableSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsEnableSchema != nil {
        // handle response
    }
}
```

## UsergroupsList

List all User Groups for a team

API method documentation
<https://api.slack.com/methods/usergroups.list>

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
    res, err := s.Usergroups.UsergroupsList(ctx, operations.UsergroupsListRequest{
        IncludeCount: slackspec.Bool(false),
        IncludeDisabled: slackspec.Bool(false),
        IncludeUsers: slackspec.Bool(false),
        Token: "error",
    }, operations.UsergroupsListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsListSchema != nil {
        // handle response
    }
}
```

## UsergroupsUpdate

Update an existing User Group

API method documentation
<https://api.slack.com/methods/usergroups.update>

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
    res, err := s.Usergroups.UsergroupsUpdate(ctx, operations.UsergroupsUpdateRequest{
        RequestBody: operations.UsergroupsUpdateApplicationJSON{
            Channels: slackspec.String("illo"),
            Description: slackspec.String("corporis"),
            Handle: slackspec.String("quidem"),
            IncludeCount: slackspec.Bool(false),
            Name: slackspec.String("Ms. Melvin Thiel IV"),
            Usergroup: "quae",
        },
        Token: "molestiae",
    }, operations.UsergroupsUpdateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsUpdateSchema != nil {
        // handle response
    }
}
```

## UsergroupsUsersList

List all users in a User Group

API method documentation
<https://api.slack.com/methods/usergroups.users.list>

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
    res, err := s.Usergroups.UsergroupsUsersList(ctx, operations.UsergroupsUsersListRequest{
        IncludeDisabled: slackspec.Bool(false),
        Token: "eveniet",
        Usergroup: "qui",
    }, operations.UsergroupsUsersListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsUsersListSchema != nil {
        // handle response
    }
}
```

## UsergroupsUsersUpdate

Update the list of users for a User Group

API method documentation
<https://api.slack.com/methods/usergroups.users.update>

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
    res, err := s.Usergroups.UsergroupsUsersUpdate(ctx, operations.UsergroupsUsersUpdateRequest{
        RequestBody: operations.UsergroupsUsersUpdateApplicationJSON{
            IncludeCount: slackspec.Bool(false),
            Usergroup: "cum",
            Users: "iure",
        },
        Token: "necessitatibus",
    }, operations.UsergroupsUsersUpdateSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UsergroupsUsersUpdateSchema != nil {
        // handle response
    }
}
```
