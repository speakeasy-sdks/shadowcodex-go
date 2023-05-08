# Calls

### Available Operations

* [CallsAdd](#callsadd) - Registers a new Call.
* [CallsEnd](#callsend) - Ends a Call.
* [CallsInfo](#callsinfo) - Returns information about a Call.
* [CallsParticipantsAdd](#callsparticipantsadd) - Registers new participants added to a Call.
* [CallsParticipantsRemove](#callsparticipantsremove) - Registers participants removed from a Call.
* [CallsUpdate](#callsupdate) - Updates information about a Call.

## CallsAdd

Registers a new Call.

API method documentation
<https://api.slack.com/methods/calls.add>

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
    res, err := s.Calls.CallsAdd(ctx, operations.CallsAddRequest{
        RequestBody: operations.CallsAddApplicationJSON{
            CreatedBy: slackspec.String("temporibus"),
            DateStart: slackspec.Int(183280),
            DesktopAppJoinURL: slackspec.String("neque"),
            ExternalDisplayID: slackspec.String("fugit"),
            ExternalUniqueID: "magni",
            JoinURL: "odio",
            Title: slackspec.String("Mr."),
            Users: slackspec.String("ullam"),
        },
        Token: "nam",
    }, operations.CallsAddSecurity{
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

## CallsEnd

Ends a Call.

API method documentation
<https://api.slack.com/methods/calls.end>

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
    res, err := s.Calls.CallsEnd(ctx, operations.CallsEndRequest{
        RequestBody: operations.CallsEndApplicationJSON{
            Duration: slackspec.Int(940432),
            ID: "0cbb1e31-b8b9-40f3-843a-1108e0adcf4b",
        },
        Token: "cupiditate",
    }, operations.CallsEndSecurity{
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

## CallsInfo

Returns information about a Call.

API method documentation
<https://api.slack.com/methods/calls.info>

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
    res, err := s.Calls.CallsInfo(ctx, operations.CallsInfoRequest{
        ID: "21879fce-953f-473e-b7fb-c7abd74dd39c",
        Token: "aut",
    }, operations.CallsInfoSecurity{
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
    res, err := s.Calls.CallsParticipantsAdd(ctx, operations.CallsParticipantsAddRequest{
        RequestBody: operations.CallsParticipantsAddApplicationJSON{
            ID: "f5d2cff7-c70a-4456-a6d4-36813f16d9f5",
            Users: "sapiente",
        },
        Token: "quisquam",
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
    res, err := s.Calls.CallsParticipantsRemove(ctx, operations.CallsParticipantsRemoveRequest{
        RequestBody: operations.CallsParticipantsRemoveApplicationJSON{
            ID: "e6c55614-6c3e-4250-bb00-8c42e141aac3",
            Users: "eum",
        },
        Token: "autem",
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

## CallsUpdate

Updates information about a Call.

API method documentation
<https://api.slack.com/methods/calls.update>

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
    res, err := s.Calls.CallsUpdate(ctx, operations.CallsUpdateRequest{
        RequestBody: operations.CallsUpdateApplicationJSON{
            DesktopAppJoinURL: slackspec.String("nobis"),
            ID: "8dd6b144-2907-4474-b78a-7bd466d28c10",
            JoinURL: slackspec.String("id"),
            Title: slackspec.String("Miss"),
        },
        Token: "neque",
    }, operations.CallsUpdateSecurity{
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
