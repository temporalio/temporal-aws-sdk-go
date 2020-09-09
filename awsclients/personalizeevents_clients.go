package awsclients

import (
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
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
	activities awsactivities.PersonalizeEventsActivities
}

func NewPersonalizeEventsStub() PersonalizeEventsClient {
	return &PersonalizeEventsStub{}
}

func (a *PersonalizeEventsStub) PutEvents(ctx workflow.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
	var output personalizeevents.PutEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *PersonalizeEventsStub) PutEventsAsync(ctx workflow.Context, input *personalizeevents.PutEventsInput) *PersonalizeeventsPutEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input)
	return &PersonalizeeventsPutEventsResult{Result: future}
}
