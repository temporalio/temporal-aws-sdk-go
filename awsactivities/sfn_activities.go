package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SFNActivities struct {
	client sfniface.SFNAPI
}

func NewSFNActivities(session *session.Session, config ...*aws.Config) *SFNActivities {
	client := sfn.New(session, config...)
	return &SFNActivities{client: client}
}

func (a *SFNActivities) CreateActivity(ctx context.Context, input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
	return a.client.CreateActivityWithContext(ctx, input)
}

func (a *SFNActivities) CreateStateMachine(ctx context.Context, input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
	return a.client.CreateStateMachineWithContext(ctx, input)
}

func (a *SFNActivities) DeleteActivity(ctx context.Context, input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
	return a.client.DeleteActivityWithContext(ctx, input)
}

func (a *SFNActivities) DeleteStateMachine(ctx context.Context, input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
	return a.client.DeleteStateMachineWithContext(ctx, input)
}

func (a *SFNActivities) DescribeActivity(ctx context.Context, input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
	return a.client.DescribeActivityWithContext(ctx, input)
}

func (a *SFNActivities) DescribeExecution(ctx context.Context, input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
	return a.client.DescribeExecutionWithContext(ctx, input)
}

func (a *SFNActivities) DescribeStateMachine(ctx context.Context, input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	return a.client.DescribeStateMachineWithContext(ctx, input)
}

func (a *SFNActivities) DescribeStateMachineForExecution(ctx context.Context, input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	return a.client.DescribeStateMachineForExecutionWithContext(ctx, input)
}

func (a *SFNActivities) GetActivityTask(ctx context.Context, input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
	return a.client.GetActivityTaskWithContext(ctx, input)
}

func (a *SFNActivities) GetExecutionHistory(ctx context.Context, input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
	return a.client.GetExecutionHistoryWithContext(ctx, input)
}

func (a *SFNActivities) ListActivities(ctx context.Context, input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
	return a.client.ListActivitiesWithContext(ctx, input)
}

func (a *SFNActivities) ListExecutions(ctx context.Context, input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
	return a.client.ListExecutionsWithContext(ctx, input)
}

func (a *SFNActivities) ListStateMachines(ctx context.Context, input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
	return a.client.ListStateMachinesWithContext(ctx, input)
}

func (a *SFNActivities) ListTagsForResource(ctx context.Context, input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SFNActivities) SendTaskFailure(ctx context.Context, input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
	return a.client.SendTaskFailureWithContext(ctx, input)
}

func (a *SFNActivities) SendTaskHeartbeat(ctx context.Context, input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
	return a.client.SendTaskHeartbeatWithContext(ctx, input)
}

func (a *SFNActivities) SendTaskSuccess(ctx context.Context, input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
	return a.client.SendTaskSuccessWithContext(ctx, input)
}

func (a *SFNActivities) StartExecution(ctx context.Context, input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	return a.client.StartExecutionWithContext(ctx, input)
}

func (a *SFNActivities) StopExecution(ctx context.Context, input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
	return a.client.StopExecutionWithContext(ctx, input)
}

func (a *SFNActivities) TagResource(ctx context.Context, input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SFNActivities) UntagResource(ctx context.Context, input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *SFNActivities) UpdateStateMachine(ctx context.Context, input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error) {
	return a.client.UpdateStateMachineWithContext(ctx, input)
}
