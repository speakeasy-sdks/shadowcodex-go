# Apps

### Available Operations

* [AppsEventAuthorizationsList](#appseventauthorizationslist) - Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to.
* [AppsPermissionsInfo](#appspermissionsinfo) - Returns list of permissions this app has on a team.
* [AppsPermissionsRequest](#appspermissionsrequest) - Allows an app to request additional scopes
* [AppsPermissionsResourcesList](#appspermissionsresourceslist) - Returns list of resource grants this app has on a team.
* [AppsPermissionsScopesList](#appspermissionsscopeslist) - Returns list of scopes this app has on a team.
* [AppsUninstall](#appsuninstall) - Uninstalls your app from a workspace.

## AppsEventAuthorizationsList

Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to.

API method documentation
<https://api.slack.com/methods/apps.event.authorizations.list>

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
    res, err := s.Apps.AppsEventAuthorizationsList(ctx, operations.AppsEventAuthorizationsListRequest{
        Cursor: slackspec.String("dolorem"),
        EventContext: "sapiente",
        Limit: slackspec.Int64(518201),
        Token: "nihil",
    }, operations.AppsEventAuthorizationsListSecurity{
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

## AppsPermissionsInfo

Returns list of permissions this app has on a team.

API method documentation
<https://api.slack.com/methods/apps.permissions.info>

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
    res, err := s.Apps.AppsPermissionsInfo(ctx, operations.AppsPermissionsInfoRequest{
        Token: slackspec.String("sit"),
    }, operations.AppsPermissionsInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppsPermissionsInfoSchema != nil {
        // handle response
    }
}
```

## AppsPermissionsRequest

Allows an app to request additional scopes

API method documentation
<https://api.slack.com/methods/apps.permissions.request>

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
    res, err := s.Apps.AppsPermissionsRequest(ctx, operations.AppsPermissionsRequestRequest{
        Scopes: "expedita",
        Token: "neque",
        TriggerID: "sed",
    }, operations.AppsPermissionsRequestSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppsPermissionsRequestSchema != nil {
        // handle response
    }
}
```

## AppsPermissionsResourcesList

Returns list of resource grants this app has on a team.

API method documentation
<https://api.slack.com/methods/apps.permissions.resources.list>

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
    res, err := s.Apps.AppsPermissionsResourcesList(ctx, operations.AppsPermissionsResourcesListRequest{
        Cursor: slackspec.String("vel"),
        Limit: slackspec.Int64(730442),
        Token: "voluptas",
    }, operations.AppsPermissionsResourcesListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppsPermissionsResourcesListSuccessSchema != nil {
        // handle response
    }
}
```

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
    res, err := s.Apps.AppsPermissionsScopesList(ctx, operations.AppsPermissionsScopesListRequest{
        Token: "deserunt",
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

## AppsUninstall

Uninstalls your app from a workspace.

API method documentation
<https://api.slack.com/methods/apps.uninstall>

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
    res, err := s.Apps.AppsUninstall(ctx, operations.AppsUninstallRequest{
        ClientID: slackspec.String("quam"),
        ClientSecret: slackspec.String("ipsum"),
        Token: slackspec.String("incidunt"),
    }, operations.AppsUninstallSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppsUninstallSchema != nil {
        // handle response
    }
}
```
