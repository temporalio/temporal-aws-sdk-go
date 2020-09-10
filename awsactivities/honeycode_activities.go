package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/honeycode"
	"github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type HoneycodeActivities struct {
	client honeycodeiface.HoneycodeAPI
}

func NewHoneycodeActivities(session *session.Session, config ...*aws.Config) *HoneycodeActivities {
	client := honeycode.New(session, config...)
	return &HoneycodeActivities{client: client}
}

func (a *HoneycodeActivities) GetScreenData(ctx context.Context, input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
	return a.client.GetScreenDataWithContext(ctx, input)
}

func (a *HoneycodeActivities) InvokeScreenAutomation(ctx context.Context, input *honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error) {
	return a.client.InvokeScreenAutomationWithContext(ctx, input)
}
