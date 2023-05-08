# OauthV2

### Available Operations

* [OauthV2Access](#oauthv2access) - Exchanges a temporary OAuth verifier code for an access token.

## OauthV2Access

Exchanges a temporary OAuth verifier code for an access token.

API method documentation
<https://api.slack.com/methods/oauth.v2.access>

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
    res, err := s.OauthV2.OauthV2Access(ctx, operations.OauthV2AccessRequest{
        ClientID: slackspec.String("sed"),
        ClientSecret: slackspec.String("sit"),
        Code: "vel",
        RedirectURI: slackspec.String("nostrum"),
    }, operations.OauthV2AccessSecurity{
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
