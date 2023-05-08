# Dialog

### Available Operations

* [DialogOpen](#dialogopen) - Open a dialog with a user

## DialogOpen

Open a dialog with a user

API method documentation
<https://api.slack.com/methods/dialog.open>

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
    res, err := s.Dialog.DialogOpen(ctx, operations.DialogOpenRequest{
        Dialog: "possimus",
        Token: "facilis",
        TriggerID: "cum",
    }, operations.DialogOpenSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DialogOpenSchema != nil {
        // handle response
    }
}
```
