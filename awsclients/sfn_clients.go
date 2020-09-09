package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sfn"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SFNClient interface {
    CreateActivity(ctx workflow.Context, input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error)
    CreateActivityAsync(ctx workflow.Context, input *sfn.CreateActivityInput) *SfnCreateActivityResult

    CreateStateMachine(ctx workflow.Context, input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error)
    CreateStateMachineAsync(ctx workflow.Context, input *sfn.CreateStateMachineInput) *SfnCreateStateMachineResult

    DeleteActivity(ctx workflow.Context, input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error)
    DeleteActivityAsync(ctx workflow.Context, input *sfn.DeleteActivityInput) *SfnDeleteActivityResult

    DeleteStateMachine(ctx workflow.Context, input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error)
    DeleteStateMachineAsync(ctx workflow.Context, input *sfn.DeleteStateMachineInput) *SfnDeleteStateMachineResult

    DescribeActivity(ctx workflow.Context, input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error)
    DescribeActivityAsync(ctx workflow.Context, input *sfn.DescribeActivityInput) *SfnDescribeActivityResult

    DescribeExecution(ctx workflow.Context, input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error)
    DescribeExecutionAsync(ctx workflow.Context, input *sfn.DescribeExecutionInput) *SfnDescribeExecutionResult

    DescribeStateMachine(ctx workflow.Context, input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error)
    DescribeStateMachineAsync(ctx workflow.Context, input *sfn.DescribeStateMachineInput) *SfnDescribeStateMachineResult

    DescribeStateMachineForExecution(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error)
    DescribeStateMachineForExecutionAsync(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) *SfnDescribeStateMachineForExecutionResult

    GetActivityTask(ctx workflow.Context, input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error)
    GetActivityTaskAsync(ctx workflow.Context, input *sfn.GetActivityTaskInput) *SfnGetActivityTaskResult

    GetExecutionHistory(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error)
    GetExecutionHistoryAsync(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) *SfnGetExecutionHistoryResult

    ListActivities(ctx workflow.Context, input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error)
    ListActivitiesAsync(ctx workflow.Context, input *sfn.ListActivitiesInput) *SfnListActivitiesResult

    ListExecutions(ctx workflow.Context, input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error)
    ListExecutionsAsync(ctx workflow.Context, input *sfn.ListExecutionsInput) *SfnListExecutionsResult

    ListStateMachines(ctx workflow.Context, input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error)
    ListStateMachinesAsync(ctx workflow.Context, input *sfn.ListStateMachinesInput) *SfnListStateMachinesResult

    ListTagsForResource(ctx workflow.Context, input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *sfn.ListTagsForResourceInput) *SfnListTagsForResourceResult

    SendTaskFailure(ctx workflow.Context, input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error)
    SendTaskFailureAsync(ctx workflow.Context, input *sfn.SendTaskFailureInput) *SfnSendTaskFailureResult

    SendTaskHeartbeat(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error)
    SendTaskHeartbeatAsync(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) *SfnSendTaskHeartbeatResult

    SendTaskSuccess(ctx workflow.Context, input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error)
    SendTaskSuccessAsync(ctx workflow.Context, input *sfn.SendTaskSuccessInput) *SfnSendTaskSuccessResult

    StartExecution(ctx workflow.Context, input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error)
    StartExecutionAsync(ctx workflow.Context, input *sfn.StartExecutionInput) *SfnStartExecutionResult

    StopExecution(ctx workflow.Context, input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error)
    StopExecutionAsync(ctx workflow.Context, input *sfn.StopExecutionInput) *SfnStopExecutionResult

    TagResource(ctx workflow.Context, input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *sfn.TagResourceInput) *SfnTagResourceResult

    UntagResource(ctx workflow.Context, input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *sfn.UntagResourceInput) *SfnUntagResourceResult

    UpdateStateMachine(ctx workflow.Context, input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error)
    UpdateStateMachineAsync(ctx workflow.Context, input *sfn.UpdateStateMachineInput) *SfnUpdateStateMachineResult
}

type SfnCreateActivityResult struct {
	Result workflow.Future
}

func (r *SfnCreateActivityResult) Get(ctx workflow.Context) (*sfn.CreateActivityOutput, error) {
    var output sfn.CreateActivityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnCreateStateMachineResult struct {
	Result workflow.Future
}

func (r *SfnCreateStateMachineResult) Get(ctx workflow.Context) (*sfn.CreateStateMachineOutput, error) {
    var output sfn.CreateStateMachineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnDeleteActivityResult struct {
	Result workflow.Future
}

func (r *SfnDeleteActivityResult) Get(ctx workflow.Context) (*sfn.DeleteActivityOutput, error) {
    var output sfn.DeleteActivityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnDeleteStateMachineResult struct {
	Result workflow.Future
}

func (r *SfnDeleteStateMachineResult) Get(ctx workflow.Context) (*sfn.DeleteStateMachineOutput, error) {
    var output sfn.DeleteStateMachineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnDescribeActivityResult struct {
	Result workflow.Future
}

func (r *SfnDescribeActivityResult) Get(ctx workflow.Context) (*sfn.DescribeActivityOutput, error) {
    var output sfn.DescribeActivityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnDescribeExecutionResult struct {
	Result workflow.Future
}

func (r *SfnDescribeExecutionResult) Get(ctx workflow.Context) (*sfn.DescribeExecutionOutput, error) {
    var output sfn.DescribeExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnDescribeStateMachineResult struct {
	Result workflow.Future
}

func (r *SfnDescribeStateMachineResult) Get(ctx workflow.Context) (*sfn.DescribeStateMachineOutput, error) {
    var output sfn.DescribeStateMachineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnDescribeStateMachineForExecutionResult struct {
	Result workflow.Future
}

func (r *SfnDescribeStateMachineForExecutionResult) Get(ctx workflow.Context) (*sfn.DescribeStateMachineForExecutionOutput, error) {
    var output sfn.DescribeStateMachineForExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnGetActivityTaskResult struct {
	Result workflow.Future
}

func (r *SfnGetActivityTaskResult) Get(ctx workflow.Context) (*sfn.GetActivityTaskOutput, error) {
    var output sfn.GetActivityTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnGetExecutionHistoryResult struct {
	Result workflow.Future
}

func (r *SfnGetExecutionHistoryResult) Get(ctx workflow.Context) (*sfn.GetExecutionHistoryOutput, error) {
    var output sfn.GetExecutionHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnListActivitiesResult struct {
	Result workflow.Future
}

func (r *SfnListActivitiesResult) Get(ctx workflow.Context) (*sfn.ListActivitiesOutput, error) {
    var output sfn.ListActivitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnListExecutionsResult struct {
	Result workflow.Future
}

func (r *SfnListExecutionsResult) Get(ctx workflow.Context) (*sfn.ListExecutionsOutput, error) {
    var output sfn.ListExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnListStateMachinesResult struct {
	Result workflow.Future
}

func (r *SfnListStateMachinesResult) Get(ctx workflow.Context) (*sfn.ListStateMachinesOutput, error) {
    var output sfn.ListStateMachinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SfnListTagsForResourceResult) Get(ctx workflow.Context) (*sfn.ListTagsForResourceOutput, error) {
    var output sfn.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnSendTaskFailureResult struct {
	Result workflow.Future
}

func (r *SfnSendTaskFailureResult) Get(ctx workflow.Context) (*sfn.SendTaskFailureOutput, error) {
    var output sfn.SendTaskFailureOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnSendTaskHeartbeatResult struct {
	Result workflow.Future
}

func (r *SfnSendTaskHeartbeatResult) Get(ctx workflow.Context) (*sfn.SendTaskHeartbeatOutput, error) {
    var output sfn.SendTaskHeartbeatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnSendTaskSuccessResult struct {
	Result workflow.Future
}

func (r *SfnSendTaskSuccessResult) Get(ctx workflow.Context) (*sfn.SendTaskSuccessOutput, error) {
    var output sfn.SendTaskSuccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnStartExecutionResult struct {
	Result workflow.Future
}

func (r *SfnStartExecutionResult) Get(ctx workflow.Context) (*sfn.StartExecutionOutput, error) {
    var output sfn.StartExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnStopExecutionResult struct {
	Result workflow.Future
}

func (r *SfnStopExecutionResult) Get(ctx workflow.Context) (*sfn.StopExecutionOutput, error) {
    var output sfn.StopExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnTagResourceResult struct {
	Result workflow.Future
}

func (r *SfnTagResourceResult) Get(ctx workflow.Context) (*sfn.TagResourceOutput, error) {
    var output sfn.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnUntagResourceResult struct {
	Result workflow.Future
}

func (r *SfnUntagResourceResult) Get(ctx workflow.Context) (*sfn.UntagResourceOutput, error) {
    var output sfn.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SfnUpdateStateMachineResult struct {
	Result workflow.Future
}

func (r *SfnUpdateStateMachineResult) Get(ctx workflow.Context) (*sfn.UpdateStateMachineOutput, error) {
    var output sfn.UpdateStateMachineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SFNStub struct {
    activities awsactivities.SFNActivities
}

func NewSFNStub() SFNClient {
    return &SFNStub{}
}

func (a *SFNStub) CreateActivity(ctx workflow.Context, input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
    var output sfn.CreateActivityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateActivity, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) CreateActivityAsync(ctx workflow.Context, input *sfn.CreateActivityInput) *SfnCreateActivityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateActivity, input)
    return &SfnCreateActivityResult{Result: future}
}

func (a *SFNStub) CreateStateMachine(ctx workflow.Context, input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
    var output sfn.CreateStateMachineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStateMachine, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) CreateStateMachineAsync(ctx workflow.Context, input *sfn.CreateStateMachineInput) *SfnCreateStateMachineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStateMachine, input)
    return &SfnCreateStateMachineResult{Result: future}
}

func (a *SFNStub) DeleteActivity(ctx workflow.Context, input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
    var output sfn.DeleteActivityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteActivity, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) DeleteActivityAsync(ctx workflow.Context, input *sfn.DeleteActivityInput) *SfnDeleteActivityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteActivity, input)
    return &SfnDeleteActivityResult{Result: future}
}

func (a *SFNStub) DeleteStateMachine(ctx workflow.Context, input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
    var output sfn.DeleteStateMachineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStateMachine, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) DeleteStateMachineAsync(ctx workflow.Context, input *sfn.DeleteStateMachineInput) *SfnDeleteStateMachineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteStateMachine, input)
    return &SfnDeleteStateMachineResult{Result: future}
}

func (a *SFNStub) DescribeActivity(ctx workflow.Context, input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
    var output sfn.DescribeActivityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeActivity, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) DescribeActivityAsync(ctx workflow.Context, input *sfn.DescribeActivityInput) *SfnDescribeActivityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeActivity, input)
    return &SfnDescribeActivityResult{Result: future}
}

func (a *SFNStub) DescribeExecution(ctx workflow.Context, input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
    var output sfn.DescribeExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) DescribeExecutionAsync(ctx workflow.Context, input *sfn.DescribeExecutionInput) *SfnDescribeExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeExecution, input)
    return &SfnDescribeExecutionResult{Result: future}
}

func (a *SFNStub) DescribeStateMachine(ctx workflow.Context, input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
    var output sfn.DescribeStateMachineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStateMachine, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) DescribeStateMachineAsync(ctx workflow.Context, input *sfn.DescribeStateMachineInput) *SfnDescribeStateMachineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStateMachine, input)
    return &SfnDescribeStateMachineResult{Result: future}
}

func (a *SFNStub) DescribeStateMachineForExecution(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error) {
    var output sfn.DescribeStateMachineForExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStateMachineForExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) DescribeStateMachineForExecutionAsync(ctx workflow.Context, input *sfn.DescribeStateMachineForExecutionInput) *SfnDescribeStateMachineForExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStateMachineForExecution, input)
    return &SfnDescribeStateMachineForExecutionResult{Result: future}
}

func (a *SFNStub) GetActivityTask(ctx workflow.Context, input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
    var output sfn.GetActivityTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetActivityTask, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) GetActivityTaskAsync(ctx workflow.Context, input *sfn.GetActivityTaskInput) *SfnGetActivityTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetActivityTask, input)
    return &SfnGetActivityTaskResult{Result: future}
}

func (a *SFNStub) GetExecutionHistory(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
    var output sfn.GetExecutionHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetExecutionHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) GetExecutionHistoryAsync(ctx workflow.Context, input *sfn.GetExecutionHistoryInput) *SfnGetExecutionHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetExecutionHistory, input)
    return &SfnGetExecutionHistoryResult{Result: future}
}

func (a *SFNStub) ListActivities(ctx workflow.Context, input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
    var output sfn.ListActivitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListActivities, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) ListActivitiesAsync(ctx workflow.Context, input *sfn.ListActivitiesInput) *SfnListActivitiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListActivities, input)
    return &SfnListActivitiesResult{Result: future}
}

func (a *SFNStub) ListExecutions(ctx workflow.Context, input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
    var output sfn.ListExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) ListExecutionsAsync(ctx workflow.Context, input *sfn.ListExecutionsInput) *SfnListExecutionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListExecutions, input)
    return &SfnListExecutionsResult{Result: future}
}

func (a *SFNStub) ListStateMachines(ctx workflow.Context, input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
    var output sfn.ListStateMachinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStateMachines, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) ListStateMachinesAsync(ctx workflow.Context, input *sfn.ListStateMachinesInput) *SfnListStateMachinesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListStateMachines, input)
    return &SfnListStateMachinesResult{Result: future}
}

func (a *SFNStub) ListTagsForResource(ctx workflow.Context, input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error) {
    var output sfn.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) ListTagsForResourceAsync(ctx workflow.Context, input *sfn.ListTagsForResourceInput) *SfnListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &SfnListTagsForResourceResult{Result: future}
}

func (a *SFNStub) SendTaskFailure(ctx workflow.Context, input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
    var output sfn.SendTaskFailureOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendTaskFailure, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) SendTaskFailureAsync(ctx workflow.Context, input *sfn.SendTaskFailureInput) *SfnSendTaskFailureResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendTaskFailure, input)
    return &SfnSendTaskFailureResult{Result: future}
}

func (a *SFNStub) SendTaskHeartbeat(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
    var output sfn.SendTaskHeartbeatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendTaskHeartbeat, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) SendTaskHeartbeatAsync(ctx workflow.Context, input *sfn.SendTaskHeartbeatInput) *SfnSendTaskHeartbeatResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendTaskHeartbeat, input)
    return &SfnSendTaskHeartbeatResult{Result: future}
}

func (a *SFNStub) SendTaskSuccess(ctx workflow.Context, input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
    var output sfn.SendTaskSuccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendTaskSuccess, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) SendTaskSuccessAsync(ctx workflow.Context, input *sfn.SendTaskSuccessInput) *SfnSendTaskSuccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendTaskSuccess, input)
    return &SfnSendTaskSuccessResult{Result: future}
}

func (a *SFNStub) StartExecution(ctx workflow.Context, input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
    var output sfn.StartExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) StartExecutionAsync(ctx workflow.Context, input *sfn.StartExecutionInput) *SfnStartExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartExecution, input)
    return &SfnStartExecutionResult{Result: future}
}

func (a *SFNStub) StopExecution(ctx workflow.Context, input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
    var output sfn.StopExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) StopExecutionAsync(ctx workflow.Context, input *sfn.StopExecutionInput) *SfnStopExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopExecution, input)
    return &SfnStopExecutionResult{Result: future}
}

func (a *SFNStub) TagResource(ctx workflow.Context, input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error) {
    var output sfn.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) TagResourceAsync(ctx workflow.Context, input *sfn.TagResourceInput) *SfnTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &SfnTagResourceResult{Result: future}
}

func (a *SFNStub) UntagResource(ctx workflow.Context, input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error) {
    var output sfn.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) UntagResourceAsync(ctx workflow.Context, input *sfn.UntagResourceInput) *SfnUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &SfnUntagResourceResult{Result: future}
}

func (a *SFNStub) UpdateStateMachine(ctx workflow.Context, input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error) {
    var output sfn.UpdateStateMachineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStateMachine, input).Get(ctx, &output)
    return &output, err
}

func (a *SFNStub) UpdateStateMachineAsync(ctx workflow.Context, input *sfn.UpdateStateMachineInput) *SfnUpdateStateMachineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateStateMachine, input)
    return &SfnUpdateStateMachineResult{Result: future}
}
