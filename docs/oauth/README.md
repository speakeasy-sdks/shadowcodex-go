# Oauth

### Available Operations

* [OauthAccess](#oauthaccess) - Exchanges a temporary OAuth verifier code for an access token.
* [OauthToken](#oauthtoken) - Exchanges a temporary OAuth verifier code for a workspace token.
* [OauthV2Access](#oauthv2access) - Exchanges a temporary OAuth verifier code for an access token.

## OauthAccess

Exchanges a temporary OAuth verifier code for an access token.

API method documentation
<https://api.slack.com/methods/oauth.access>

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
    res, err := s.Oauth.OauthAccess(ctx, operations.OauthAccessRequest{
        ClientID: slackspec.String("magnam"),
        ClientSecret: slackspec.String("consequatur"),
        Code: slackspec.String("esse"),
        RedirectURI: slackspec.String("ipsam"),
        SingleChannel: slackspec.Bool(false),
    }, operations.OauthAccessSecurity{
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

## OauthToken

Exchanges a temporary OAuth verifier code for a workspace token.

API method documentation
<https://api.slack.com/methods/oauth.token>

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
    res, err := s.Oauth.OauthToken(ctx, operations.OauthTokenRequest{
        ClientID: slackspec.String("sit"),
        ClientSecret: slackspec.String("voluptatum"),
        Code: slackspec.String("quas"),
        RedirectURI: slackspec.String("repudiandae"),
        SingleChannel: slackspec.Bool(false),
    }, operations.OauthTokenSecurity{
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
    res, err := s.Oauth.OauthV2Access(ctx, operations.OauthV2AccessRequest{
        ClientID: slackspec.String("corporis"),
        ClientSecret: slackspec.String("et"),
        Code: "blanditiis",
        RedirectURI: slackspec.String("ex"),
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
