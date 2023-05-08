# FilesComments

### Available Operations

* [FilesCommentsDelete](#filescommentsdelete) - Deletes an existing comment on a file.

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
    res, err := s.FilesComments.FilesCommentsDelete(ctx, operations.FilesCommentsDeleteRequest{
        RequestBody: &operations.FilesCommentsDeleteApplicationJSON{
            File: slackspec.String("incidunt"),
            ID: slackspec.String("294e3698-f447-4f60-be8b-445e80ca55ef"),
        },
        Token: slackspec.String("nulla"),
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
