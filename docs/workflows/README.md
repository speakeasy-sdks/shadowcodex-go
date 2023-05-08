# Workflows

### Available Operations

* [WorkflowsStepCompleted](#workflowsstepcompleted) - Indicate that an app's step in a workflow completed execution.
* [WorkflowsStepFailed](#workflowsstepfailed) - Indicate that an app's step in a workflow failed to execute.
* [WorkflowsUpdateStep](#workflowsupdatestep) - Update the configuration for a workflow extension step.

## WorkflowsStepCompleted

Indicate that an app's step in a workflow completed execution.

API method documentation
<https://api.slack.com/methods/workflows.stepCompleted>

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
    res, err := s.Workflows.WorkflowsStepCompleted(ctx, operations.WorkflowsStepCompletedRequest{
        Outputs: slackspec.String("vitae"),
        Token: "rerum",
        WorkflowStepExecuteID: "tempora",
    }, operations.WorkflowsStepCompletedSecurity{
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

## WorkflowsStepFailed

Indicate that an app's step in a workflow failed to execute.

API method documentation
<https://api.slack.com/methods/workflows.stepFailed>

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
    res, err := s.Workflows.WorkflowsStepFailed(ctx, operations.WorkflowsStepFailedRequest{
        Error: "quis",
        Token: "inventore",
        WorkflowStepExecuteID: "fugit",
    }, operations.WorkflowsStepFailedSecurity{
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

## WorkflowsUpdateStep

Update the configuration for a workflow extension step.

API method documentation
<https://api.slack.com/methods/workflows.updateStep>

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
    res, err := s.Workflows.WorkflowsUpdateStep(ctx, operations.WorkflowsUpdateStepRequest{
        Inputs: slackspec.String("cumque"),
        Outputs: slackspec.String("quae"),
        StepImageURL: slackspec.String("perferendis"),
        StepName: slackspec.String("velit"),
        Token: "aspernatur",
        WorkflowStepEditID: "eum",
    }, operations.WorkflowsUpdateStepSecurity{
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
