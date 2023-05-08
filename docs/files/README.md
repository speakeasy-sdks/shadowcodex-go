# Files

### Available Operations

* [FilesCommentsDelete](#filescommentsdelete) - Deletes an existing comment on a file.
* [FilesDelete](#filesdelete) - Deletes a file.
* [FilesInfo](#filesinfo) - Gets information about a file.
* [FilesList](#fileslist) - List for a team, in a channel, or from a user with applied filters.
* [FilesRemoteAdd](#filesremoteadd) - Adds a file from a remote service
* [FilesRemoteInfo](#filesremoteinfo) - Retrieve information about a remote file added to Slack
* [FilesRemoteList](#filesremotelist) - Retrieve information about a remote file added to Slack
* [FilesRemoteRemove](#filesremoteremove) - Remove a remote file.
* [FilesRemoteShare](#filesremoteshare) - Share a remote file into a channel.
* [FilesRemoteUpdate](#filesremoteupdate) - Updates an existing remote file.
* [FilesRevokePublicURL](#filesrevokepublicurl) - Revokes public/external sharing access for a file
* [FilesSharedPublicURL](#filessharedpublicurl) - Enables a file for public/external sharing.
* [FilesUpload](#filesupload) - Uploads or creates a file.

## FilesCommentsDelete

Deletes an existing comment on a file.

API method documentation
<https://api.slack.com/methods/files.comments.delete>

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
    res, err := s.Files.FilesCommentsDelete(ctx, operations.FilesCommentsDeleteRequest{
        RequestBody: &operations.FilesCommentsDeleteApplicationJSON{
            File: slackspec.String("cum"),
            ID: slackspec.String("375ed4f6-fbee-441f-b331-7fe35b60eb1e"),
        },
        Token: slackspec.String("similique"),
    }, operations.FilesCommentsDeleteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesCommentsDeleteSchema != nil {
        // handle response
    }
}
```

## FilesDelete

Deletes a file.

API method documentation
<https://api.slack.com/methods/files.delete>

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
    res, err := s.Files.FilesDelete(ctx, operations.FilesDeleteRequest{
        RequestBody: &operations.FilesDeleteApplicationJSON{
            File: slackspec.String("tempora"),
        },
        Token: slackspec.String("aspernatur"),
    }, operations.FilesDeleteSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesDeleteSchema != nil {
        // handle response
    }
}
```

## FilesInfo

Gets information about a file.

API method documentation
<https://api.slack.com/methods/files.info>

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
    res, err := s.Files.FilesInfo(ctx, operations.FilesInfoRequest{
        Count: slackspec.String("voluptas"),
        Cursor: slackspec.String("voluptas"),
        File: slackspec.String("voluptas"),
        Limit: slackspec.Int64(324405),
        Page: slackspec.String("nobis"),
        Token: slackspec.String("dolorum"),
    }, operations.FilesInfoSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesInfoSchema != nil {
        // handle response
    }
}
```

## FilesList

List for a team, in a channel, or from a user with applied filters.

API method documentation
<https://api.slack.com/methods/files.list>

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
    res, err := s.Files.FilesList(ctx, operations.FilesListRequest{
        Channel: slackspec.String("adipisci"),
        Count: slackspec.String("minus"),
        Page: slackspec.String("dolores"),
        ShowFilesHiddenByLimit: slackspec.Bool(false),
        Token: slackspec.String("blanditiis"),
        TsFrom: slackspec.Float64(4492.92),
        TsTo: slackspec.Float64(2962.42),
        Types: slackspec.String("aliquam"),
        User: slackspec.String("officiis"),
    }, operations.FilesListSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesListSchema != nil {
        // handle response
    }
}
```

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
    res, err := s.Files.FilesRemoteAdd(ctx, operations.FilesRemoteAddRequestBody{
        ExternalID: slackspec.String("temporibus"),
        ExternalURL: slackspec.String("ullam"),
        Filetype: slackspec.String("adipisci"),
        IndexableFileContents: slackspec.String("cum"),
        PreviewImage: slackspec.String("blanditiis"),
        Title: slackspec.String("Ms."),
        Token: slackspec.String("hic"),
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
    res, err := s.Files.FilesRemoteInfo(ctx, operations.FilesRemoteInfoRequest{
        ExternalID: slackspec.String("nesciunt"),
        File: slackspec.String("culpa"),
        Token: slackspec.String("corrupti"),
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
    res, err := s.Files.FilesRemoteList(ctx, operations.FilesRemoteListRequest{
        Channel: slackspec.String("pariatur"),
        Cursor: slackspec.String("totam"),
        Limit: slackspec.Int64(940210),
        Token: slackspec.String("exercitationem"),
        TsFrom: slackspec.Float64(7507.65),
        TsTo: slackspec.Float64(246.19),
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
    res, err := s.Files.FilesRemoteRemove(ctx, operations.FilesRemoteRemoveRequestBody{
        ExternalID: slackspec.String("rerum"),
        File: slackspec.String("sed"),
        Token: slackspec.String("reiciendis"),
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
    res, err := s.Files.FilesRemoteShare(ctx, operations.FilesRemoteShareRequest{
        Channels: slackspec.String("explicabo"),
        ExternalID: slackspec.String("asperiores"),
        File: slackspec.String("facilis"),
        Token: slackspec.String("voluptate"),
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
    res, err := s.Files.FilesRemoteUpdate(ctx, operations.FilesRemoteUpdateRequestBody{
        ExternalID: slackspec.String("expedita"),
        ExternalURL: slackspec.String("ab"),
        File: slackspec.String("iste"),
        Filetype: slackspec.String("dolore"),
        IndexableFileContents: slackspec.String("laborum"),
        PreviewImage: slackspec.String("sed"),
        Title: slackspec.String("Ms."),
        Token: slackspec.String("commodi"),
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

## FilesRevokePublicURL

Revokes public/external sharing access for a file

API method documentation
<https://api.slack.com/methods/files.revokePublicURL>

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
    res, err := s.Files.FilesRevokePublicURL(ctx, operations.FilesRevokePublicURLRequest{
        RequestBody: &operations.FilesRevokePublicURLApplicationJSON{
            File: slackspec.String("quidem"),
        },
        Token: slackspec.String("explicabo"),
    }, operations.FilesRevokePublicURLSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesRevokePublicURLSchema != nil {
        // handle response
    }
}
```

## FilesSharedPublicURL

Enables a file for public/external sharing.

API method documentation
<https://api.slack.com/methods/files.sharedPublicURL>

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
    res, err := s.Files.FilesSharedPublicURL(ctx, operations.FilesSharedPublicURLRequest{
        RequestBody: &operations.FilesSharedPublicURLApplicationJSON{
            File: slackspec.String("voluptas"),
        },
        Token: slackspec.String("unde"),
    }, operations.FilesSharedPublicURLSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesSharedPublicURLSchema != nil {
        // handle response
    }
}
```

## FilesUpload

Uploads or creates a file.

API method documentation
<https://api.slack.com/methods/files.upload>

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
    res, err := s.Files.FilesUpload(ctx, operations.FilesUploadRequestBody{
        Channels: slackspec.String("architecto"),
        Content: slackspec.String("suscipit"),
        File: slackspec.String("sapiente"),
        Filename: slackspec.String("debitis"),
        Filetype: slackspec.String("illo"),
        InitialComment: slackspec.String("reiciendis"),
        ThreadTs: slackspec.Float64(193),
        Title: slackspec.String("Ms."),
        Token: slackspec.String("maiores"),
    }, operations.FilesUploadSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesUploadSchema != nil {
        // handle response
    }
}
```
