package awsclients

import (
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type QLDBSessionClient interface {
	SendCommand(ctx workflow.Context, input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error)
	SendCommandAsync(ctx workflow.Context, input *qldbsession.SendCommandInput) *QldbsessionSendCommandResult
}

type QldbsessionSendCommandResult struct {
	Result workflow.Future
}

func (r *QldbsessionSendCommandResult) Get(ctx workflow.Context) (*qldbsession.SendCommandOutput, error) {
	var output qldbsession.SendCommandOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type QLDBSessionStub struct {
	activities awsactivities.QLDBSessionActivities
}

func NewQLDBSessionStub() QLDBSessionClient {
	return &QLDBSessionStub{}
}

func (a *QLDBSessionStub) SendCommand(ctx workflow.Context, input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error) {
	var output qldbsession.SendCommandOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SendCommand, input).Get(ctx, &output)
	return &output, err
}

func (a *QLDBSessionStub) SendCommandAsync(ctx workflow.Context, input *qldbsession.SendCommandInput) *QldbsessionSendCommandResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SendCommand, input)
	return &QldbsessionSendCommandResult{Result: future}
}
