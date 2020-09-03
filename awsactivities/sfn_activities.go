package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
)

type SFNActivities struct {
	client sfniface.SFNAPI
}

func NewSFNActivities(client sfniface.SFNAPI) *SFNActivities {
    return &SFNActivities{client: client}
}

func (a *SFNActivities) CreateActivity(input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
    return a.client.CreateActivity(input)
}

func (a *SFNActivities) CreateStateMachine(input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
    return a.client.CreateStateMachine(input)
}

func (a *SFNActivities) DeleteActivity(input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
    return a.client.DeleteActivity(input)
}

func (a *SFNActivities) DeleteStateMachine(input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
    return a.client.DeleteStateMachine(input)
}

func (a *SFNActivities) DescribeActivity(input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
    return a.client.DescribeActivity(input)
}

func (a *SFNActivities) DescribeExecution(input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
    return a.client.DescribeExecution(input)
}

func (a *SFNActivities) DescribeStateMachine(input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
    return a.client.DescribeStateMachine(input)
}

func (a *SFNActivities) DescribeStateMachineForExecution(input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error) {
    return a.client.DescribeStateMachineForExecution(input)
}

func (a *SFNActivities) GetActivityTask(input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
    return a.client.GetActivityTask(input)
}

func (a *SFNActivities) GetExecutionHistory(input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
    return a.client.GetExecutionHistory(input)
}

func (a *SFNActivities) ListActivities(input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
    return a.client.ListActivities(input)
}

func (a *SFNActivities) ListExecutions(input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
    return a.client.ListExecutions(input)
}

func (a *SFNActivities) ListStateMachines(input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
    return a.client.ListStateMachines(input)
}

func (a *SFNActivities) ListTagsForResource(input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SFNActivities) SendTaskFailure(input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
    return a.client.SendTaskFailure(input)
}

func (a *SFNActivities) SendTaskHeartbeat(input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
    return a.client.SendTaskHeartbeat(input)
}

func (a *SFNActivities) SendTaskSuccess(input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
    return a.client.SendTaskSuccess(input)
}

func (a *SFNActivities) StartExecution(input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
    return a.client.StartExecution(input)
}

func (a *SFNActivities) StopExecution(input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
    return a.client.StopExecution(input)
}

func (a *SFNActivities) TagResource(input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SFNActivities) UntagResource(input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *SFNActivities) UpdateStateMachine(input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error) {
    return a.client.UpdateStateMachine(input)
}
