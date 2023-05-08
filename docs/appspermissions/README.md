# AppsPermissions

### Available Operations

* [AppsPermissionsInfo](#appspermissionsinfo) - Returns list of permissions this app has on a team.
* [AppsPermissionsRequest](#appspermissionsrequest) - Allows an app to request additional scopes

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
    res, err := s.AppsPermissions.AppsPermissionsInfo(ctx, operations.AppsPermissionsInfoRequest{
        Token: slackspec.String("soluta"),
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
    res, err := s.AppsPermissions.AppsPermissionsRequest(ctx, operations.AppsPermissionsRequestRequest{
        Scopes: "dicta",
        Token: "laborum",
        TriggerID: "totam",
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
