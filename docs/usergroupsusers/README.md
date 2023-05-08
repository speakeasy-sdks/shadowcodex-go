# UsergroupsUsers

### Available Operations

* [UsergroupsUsersList](#usergroupsuserslist) - List all users in a User Group
* [UsergroupsUsersUpdate](#usergroupsusersupdate) - Update the list of users for a User Group

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
    res, err := s.UsergroupsUsers.UsergroupsUsersList(ctx, operations.UsergroupsUsersListRequest{
        IncludeDisabled: slackspec.Bool(false),
        Token: "ratione",
        Usergroup: "laborum",
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
    res, err := s.UsergroupsUsers.UsergroupsUsersUpdate(ctx, operations.UsergroupsUsersUpdateRequest{
        RequestBody: operations.UsergroupsUsersUpdateApplicationJSON{
            IncludeCount: slackspec.Bool(false),
            Usergroup: "distinctio",
            Users: "voluptatum",
        },
        Token: "rem",
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
