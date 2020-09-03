package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"go.temporal.io/sdk/workflow"
)

type MobileAnalyticsClient interface {
    PutEvents(ctx workflow.Context, input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error)
    PutEventsAsync(ctx workflow.Context, input *mobileanalytics.PutEventsInput) *MobileanalyticsPutEventsResult
}
type MobileanalyticsPutEventsResult struct {
	Result workflow.Future
}

func (r *MobileanalyticsPutEventsResult) Get(ctx workflow.Context) (*mobileanalytics.PutEventsOutput, error) {
    var output mobileanalytics.PutEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MobileAnalyticsStub struct {
    activities MobileAnalyticsClient
}

func NewMobileAnalyticsStub() MobileAnalyticsClient {
    return &MobileAnalyticsStub{}
}

func (a *MobileAnalyticsStub) PutEvents(ctx workflow.Context, input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error) {
    var output mobileanalytics.PutEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input).Get(ctx, &output)
    return &output, err
}
