# TeamProfile

### Available Operations

* [TeamProfileGet](#teamprofileget) - Retrieve a team's profile.

## TeamProfileGet

Retrieve a team's profile.

API method documentation
<https://api.slack.com/methods/team.profile.get>

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
    res, err := s.TeamProfile.TeamProfileGet(ctx, operations.TeamProfileGetRequest{
        Token: "magnam",
        Visibility: slackspec.String("saepe"),
    }, operations.TeamProfileGetSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TeamProfileGetSuccessSchema != nil {
        // handle response
    }
}
```
