# AdminTeamsSettings

### Available Operations

* [AdminTeamsSettingsInfo](#adminteamssettingsinfo) - Fetch information about settings in a workspace
* [AdminTeamsSettingsSetDefaultChannels](#adminteamssettingssetdefaultchannels) - Set the default channels of a workspace.
* [AdminTeamsSettingsSetDescription](#adminteamssettingssetdescription) - Set the description of a given workspace.
* [AdminTeamsSettingsSetDiscoverability](#adminteamssettingssetdiscoverability) - An API method that allows admins to set the discoverability of a given workspace
* [AdminTeamsSettingsSetIcon](#adminteamssettingsseticon) - Sets the icon of a workspace.
* [AdminTeamsSettingsSetName](#adminteamssettingssetname) - Set the name of a given workspace.

## AdminTeamsSettingsInfo

Fetch information about settings in a workspace

API method documentation
<https://api.slack.com/methods/admin.teams.settings.info>

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
    res, err := s.AdminTeamsSettings.AdminTeamsSettingsInfo(ctx, operations.AdminTeamsSettingsInfoRequest{
        TeamID: "hic",
        Token: "excepturi",
    }, operations.AdminTeamsSettingsInfoSecurity{
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

## AdminTeamsSettingsSetDefaultChannels

Set the default channels of a workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setDefaultChannels>

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
    res, err := s.AdminTeamsSettings.AdminTeamsSettingsSetDefaultChannels(ctx, operations.AdminTeamsSettingsSetDefaultChannelsRequestBody{
        ChannelIds: "cum",
        TeamID: "voluptate",
        Token: "dignissimos",
    }, operations.AdminTeamsSettingsSetDefaultChannelsSecurity{
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

## AdminTeamsSettingsSetDescription

Set the description of a given workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setDescription>

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
    res, err := s.AdminTeamsSettings.AdminTeamsSettingsSetDescription(ctx, operations.AdminTeamsSettingsSetDescriptionRequest{
        RequestBody: operations.AdminTeamsSettingsSetDescriptionApplicationJSON{
            Description: "reiciendis",
            TeamID: "amet",
        },
        Token: "dolorum",
    }, operations.AdminTeamsSettingsSetDescriptionSecurity{
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

## AdminTeamsSettingsSetDiscoverability

An API method that allows admins to set the discoverability of a given workspace

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setDiscoverability>

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
    res, err := s.AdminTeamsSettings.AdminTeamsSettingsSetDiscoverability(ctx, operations.AdminTeamsSettingsSetDiscoverabilityRequest{
        RequestBody: operations.AdminTeamsSettingsSetDiscoverabilityApplicationJSON{
            Discoverability: "numquam",
            TeamID: "veritatis",
        },
        Token: "ipsa",
    }, operations.AdminTeamsSettingsSetDiscoverabilitySecurity{
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

## AdminTeamsSettingsSetIcon

Sets the icon of a workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setIcon>

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
    res, err := s.AdminTeamsSettings.AdminTeamsSettingsSetIcon(ctx, operations.AdminTeamsSettingsSetIconRequestBody{
        ImageURL: "ipsa",
        TeamID: "iure",
        Token: "odio",
    }, operations.AdminTeamsSettingsSetIconSecurity{
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

## AdminTeamsSettingsSetName

Set the name of a given workspace.

API method documentation
<https://api.slack.com/methods/admin.teams.settings.setName>

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
    res, err := s.AdminTeamsSettings.AdminTeamsSettingsSetName(ctx, operations.AdminTeamsSettingsSetNameRequest{
        RequestBody: operations.AdminTeamsSettingsSetNameApplicationJSON{
            Name: "Sophia Predovic",
            TeamID: "natus",
        },
        Token: "eos",
    }, operations.AdminTeamsSettingsSetNameSecurity{
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
