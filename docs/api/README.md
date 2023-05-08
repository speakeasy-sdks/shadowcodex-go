# API

### Available Operations

* [APITest](#apitest) - Checks API calling code.

## APITest

Checks API calling code.

API method documentation
<https://api.slack.com/methods/api.test>

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
    res, err := s.API.APITest(ctx, operations.APITestRequest{
        Error: slackspec.String("numquam"),
        Foo: slackspec.String("enim"),
    }, operations.APITestSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APITestSuccessSchema != nil {
        // handle response
    }
}
```
