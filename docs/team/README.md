# Team

### Available Operations

* [TeamAccessLogs](#teamaccesslogs) - Gets the access logs for the current team.
* [TeamBillableInfo](#teambillableinfo) - Gets billable users information for the current team.
* [TeamInfo](#teaminfo) - Gets information about the current team.
* [TeamIntegrationLogs](#teamintegrationlogs) - Gets the integration logs for the current team.
* [TeamProfileGet](#teamprofileget) - Retrieve a team's profile.

## TeamAccessLogs

Gets the access logs for the current team.

API method documentation
<https://api.slack.com/methods/team.accessLogs>

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
    res, err := s.Team.TeamAccessLogs(ctx, operations.TeamAccessLogsRequest{
        Before: slackspec.String("sunt"),
        Count: slackspec.String("recusandae"),
        Page: slackspec.String("dolorum"),
        Token: "repellendus",
    }, operations.TeamAccessLogsSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TeamAccessLogsSchema != nil {
        // handle response
    }
}
```

## TeamBillableInfo

Gets billable users information for the current team.

API method documentation
<https://api.slack.com/methods/team.billableInfo>

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
    res, err := s.Team.TeamBillableInfo(ctx, operations.TeamBillableInfoRequest{
        Token: "labore",
        User: slackspec.String("reiciendis"),
    }, operations.TeamBillableInfoSecurity{
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

## TeamInfo

Gets information about the current team.

API method documentation
<https://api.slack.com/methods/team.info>

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
    res, err := s.Team.TeamInfo(ctx, operations.TeamInfoRequest{
        Team: slackspec.String("doloremque"),
        Token: "repudiandae",
    }, operations.TeamInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TeamInfoSchema != nil {
        // handle response
    }
}
```

## TeamIntegrationLogs

Gets the integration logs for the current team.

API method documentation
<https://api.slack.com/methods/team.integrationLogs>

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
    res, err := s.Team.TeamIntegrationLogs(ctx, operations.TeamIntegrationLogsRequest{
        AppID: slackspec.String("dicta"),
        ChangeType: slackspec.String("accusantium"),
        Count: slackspec.String("beatae"),
        Page: slackspec.String("dolores"),
        ServiceID: slackspec.String("enim"),
        Token: "laboriosam",
        User: slackspec.String("velit"),
    }, operations.TeamIntegrationLogsSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TeamIntegrationLogsSchema != nil {
        // handle response
    }
}
```

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
    res, err := s.Team.TeamProfileGet(ctx, operations.TeamProfileGetRequest{
        Token: "a",
        Visibility: slackspec.String("molestias"),
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
