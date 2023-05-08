# AdminApps

### Available Operations

* [AdminAppsApprove](#adminappsapprove) - Approve an app for installation on a workspace.
* [AdminAppsRestrict](#adminappsrestrict) - Restrict an app for installation on a workspace.

## AdminAppsApprove

Approve an app for installation on a workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.approve>

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
    res, err := s.AdminApps.AdminAppsApprove(ctx, operations.AdminAppsApproveRequest{
        RequestBody: &operations.AdminAppsApproveApplicationJSON{
            AppID: slackspec.String("accusamus"),
            RequestID: slackspec.String("non"),
            TeamID: slackspec.String("occaecati"),
        },
        Token: "enim",
    }, operations.AdminAppsApproveSecurity{
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

## AdminAppsRestrict

Restrict an app for installation on a workspace.

API method documentation
<https://api.slack.com/methods/admin.apps.restrict>

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
    res, err := s.AdminApps.AdminAppsRestrict(ctx, operations.AdminAppsRestrictRequest{
        RequestBody: &operations.AdminAppsRestrictApplicationJSON{
            AppID: slackspec.String("accusamus"),
            RequestID: slackspec.String("delectus"),
            TeamID: slackspec.String("quidem"),
        },
        Token: "provident",
    }, operations.AdminAppsRestrictSecurity{
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
