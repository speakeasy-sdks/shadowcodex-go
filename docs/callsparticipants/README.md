# CallsParticipants

### Available Operations

* [CallsParticipantsAdd](#callsparticipantsadd) - Registers new participants added to a Call.
* [CallsParticipantsRemove](#callsparticipantsremove) - Registers participants removed from a Call.

## CallsParticipantsAdd

Registers new participants added to a Call.

API method documentation
<https://api.slack.com/methods/calls.participants.add>

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
    res, err := s.CallsParticipants.CallsParticipantsAdd(ctx, operations.CallsParticipantsAddRequest{
        RequestBody: operations.CallsParticipantsAddApplicationJSON{
            ID: "cdca4251-904e-4523-87e0-bc7178e4796f",
            Users: "dolores",
        },
        Token: "deserunt",
    }, operations.CallsParticipantsAddSecurity{
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

## CallsParticipantsRemove

Registers participants removed from a Call.

API method documentation
<https://api.slack.com/methods/calls.participants.remove>

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
    res, err := s.CallsParticipants.CallsParticipantsRemove(ctx, operations.CallsParticipantsRemoveRequest{
        RequestBody: operations.CallsParticipantsRemoveApplicationJSON{
            ID: "70c68828-2aa4-4825-a2f2-22e9817ee17c",
            Users: "nam",
        },
        Token: "vero",
    }, operations.CallsParticipantsRemoveSecurity{
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
