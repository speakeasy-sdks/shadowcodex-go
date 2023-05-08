# Migration

### Available Operations

* [MigrationExchange](#migrationexchange) - For Enterprise Grid workspaces, map local user IDs to global user IDs

## MigrationExchange

For Enterprise Grid workspaces, map local user IDs to global user IDs

API method documentation
<https://api.slack.com/methods/migration.exchange>

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
    res, err := s.Migration.MigrationExchange(ctx, operations.MigrationExchangeRequest{
        TeamID: slackspec.String("voluptatem"),
        ToOld: slackspec.Bool(false),
        Token: "culpa",
        Users: "expedita",
    }, operations.MigrationExchangeSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MigrationExchangeSuccessSchema != nil {
        // handle response
    }
}
```
