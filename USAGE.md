<!-- Start SDK Example Usage -->
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
    res, err := s.Admin.AdminAppsApprove(ctx, operations.AdminAppsApproveRequest{
        RequestBody: &operations.AdminAppsApproveApplicationJSON{
            AppID: slackspec.String("corrupti"),
            RequestID: slackspec.String("provident"),
            TeamID: slackspec.String("distinctio"),
        },
        Token: "quibusdam",
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
<!-- End SDK Example Usage -->