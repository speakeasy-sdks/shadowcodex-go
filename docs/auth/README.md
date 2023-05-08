# Auth

### Available Operations

* [AuthRevoke](#authrevoke) - Revokes a token.
* [AuthTest](#authtest) - Checks authentication & identity.

## AuthRevoke

Revokes a token.

API method documentation
<https://api.slack.com/methods/auth.revoke>

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
    res, err := s.Auth.AuthRevoke(ctx, operations.AuthRevokeRequest{
        Test: slackspec.Bool(false),
        Token: "facilis",
    }, operations.AuthRevokeSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthRevokeSchema != nil {
        // handle response
    }
}
```

## AuthTest

Checks authentication & identity.

API method documentation
<https://api.slack.com/methods/auth.test>

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
    res, err := s.Auth.AuthTest(ctx, operations.AuthTestRequest{
        Token: "aliquid",
    }, operations.AuthTestSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthTestSuccessSchema != nil {
        // handle response
    }
}
```
