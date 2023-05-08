# FilesRemote

### Available Operations

* [FilesRemoteAdd](#filesremoteadd) - Adds a file from a remote service
* [FilesRemoteInfo](#filesremoteinfo) - Retrieve information about a remote file added to Slack
* [FilesRemoteList](#filesremotelist) - Retrieve information about a remote file added to Slack
* [FilesRemoteRemove](#filesremoteremove) - Remove a remote file.
* [FilesRemoteShare](#filesremoteshare) - Share a remote file into a channel.
* [FilesRemoteUpdate](#filesremoteupdate) - Updates an existing remote file.

## FilesRemoteAdd

Adds a file from a remote service

API method documentation
<https://api.slack.com/methods/files.remote.add>

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
    res, err := s.FilesRemote.FilesRemoteAdd(ctx, operations.FilesRemoteAddRequestBody{
        ExternalID: slackspec.String("magni"),
        ExternalURL: slackspec.String("aperiam"),
        Filetype: slackspec.String("saepe"),
        IndexableFileContents: slackspec.String("numquam"),
        PreviewImage: slackspec.String("veniam"),
        Title: slackspec.String("Ms."),
        Token: slackspec.String("officiis"),
    }, operations.FilesRemoteAddSecurity{
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

## FilesRemoteInfo

Retrieve information about a remote file added to Slack

API method documentation
<https://api.slack.com/methods/files.remote.info>

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
    res, err := s.FilesRemote.FilesRemoteInfo(ctx, operations.FilesRemoteInfoRequest{
        ExternalID: slackspec.String("beatae"),
        File: slackspec.String("laudantium"),
        Token: slackspec.String("exercitationem"),
    }, operations.FilesRemoteInfoSecurity{
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

## FilesRemoteList

Retrieve information about a remote file added to Slack

API method documentation
<https://api.slack.com/methods/files.remote.list>

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
    res, err := s.FilesRemote.FilesRemoteList(ctx, operations.FilesRemoteListRequest{
        Channel: slackspec.String("praesentium"),
        Cursor: slackspec.String("cum"),
        Limit: slackspec.Int64(386827),
        Token: slackspec.String("dolorum"),
        TsFrom: slackspec.Float64(5300.89),
        TsTo: slackspec.Float64(6223.85),
    }, operations.FilesRemoteListSecurity{
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

## FilesRemoteRemove

Remove a remote file.

API method documentation
<https://api.slack.com/methods/files.remote.remove>

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
    res, err := s.FilesRemote.FilesRemoteRemove(ctx, operations.FilesRemoteRemoveRequestBody{
        ExternalID: slackspec.String("hic"),
        File: slackspec.String("expedita"),
        Token: slackspec.String("debitis"),
    }, operations.FilesRemoteRemoveSecurity{
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

## FilesRemoteShare

Share a remote file into a channel.

API method documentation
<https://api.slack.com/methods/files.remote.share>

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
    res, err := s.FilesRemote.FilesRemoteShare(ctx, operations.FilesRemoteShareRequest{
        Channels: slackspec.String("neque"),
        ExternalID: slackspec.String("dolorum"),
        File: slackspec.String("nostrum"),
        Token: slackspec.String("officia"),
    }, operations.FilesRemoteShareSecurity{
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

## FilesRemoteUpdate

Updates an existing remote file.

API method documentation
<https://api.slack.com/methods/files.remote.update>

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
    res, err := s.FilesRemote.FilesRemoteUpdate(ctx, operations.FilesRemoteUpdateRequestBody{
        ExternalID: slackspec.String("dolorum"),
        ExternalURL: slackspec.String("corrupti"),
        File: slackspec.String("accusamus"),
        Filetype: slackspec.String("tempora"),
        IndexableFileContents: slackspec.String("atque"),
        PreviewImage: slackspec.String("fugit"),
        Title: slackspec.String("Mrs."),
        Token: slackspec.String("fugiat"),
    }, operations.FilesRemoteUpdateSecurity{
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
