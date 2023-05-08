# AppsPermissionsScopes

### Available Operations

* [AppsPermissionsScopesList](#appspermissionsscopeslist) - Returns list of scopes this app has on a team.

## AppsPermissionsScopesList

Returns list of scopes this app has on a team.

API method documentation
<https://api.slack.com/methods/apps.permissions.scopes.list>

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
    res, err := s.AppsPermissionsScopes.AppsPermissionsScopesList(ctx, operations.AppsPermissionsScopesListRequest{
        Token: "distinctio",
    }, operations.AppsPermissionsScopesListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIPermissionsScopesListSuccessSchema != nil {
        // handle response
    }
}
```
