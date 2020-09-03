package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"go.temporal.io/sdk/workflow"
)

type PersonalizeEventsClient interface {
    PutEvents(ctx workflow.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error)
    PutEventsAsync(ctx workflow.Context, input *personalizeevents.PutEventsInput) *PersonalizeeventsPutEventsResult
}
type PersonalizeeventsPutEventsResult struct {
	Result workflow.Future
}

func (r *PersonalizeeventsPutEventsResult) Get(ctx workflow.Context) (*personalizeevents.PutEventsOutput, error) {
    var output personalizeevents.PutEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type PersonalizeEventsStub struct {
    activities PersonalizeEventsClient
}

func NewPersonalizeEventsStub() PersonalizeEventsClient {
    return &PersonalizeEventsStub{}
}

func (a *PersonalizeEventsStub) PutEvents(ctx workflow.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
    var output personalizeevents.PutEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input).Get(ctx, &output)
    return &output, err
}
