package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appflow"
	"go.temporal.io/sdk/workflow"
)

type AppflowClient interface {
    CreateConnectorProfile(ctx workflow.Context, input *appflow.CreateConnectorProfileInput) (*appflow.CreateConnectorProfileOutput, error)
    CreateConnectorProfileAsync(ctx workflow.Context, input *appflow.CreateConnectorProfileInput) *AppflowCreateConnectorProfileResult

    CreateFlow(ctx workflow.Context, input *appflow.CreateFlowInput) (*appflow.CreateFlowOutput, error)
    CreateFlowAsync(ctx workflow.Context, input *appflow.CreateFlowInput) *AppflowCreateFlowResult

    DeleteConnectorProfile(ctx workflow.Context, input *appflow.DeleteConnectorProfileInput) (*appflow.DeleteConnectorProfileOutput, error)
    DeleteConnectorProfileAsync(ctx workflow.Context, input *appflow.DeleteConnectorProfileInput) *AppflowDeleteConnectorProfileResult

    DeleteFlow(ctx workflow.Context, input *appflow.DeleteFlowInput) (*appflow.DeleteFlowOutput, error)
    DeleteFlowAsync(ctx workflow.Context, input *appflow.DeleteFlowInput) *AppflowDeleteFlowResult

    DescribeConnectorEntity(ctx workflow.Context, input *appflow.DescribeConnectorEntityInput) (*appflow.DescribeConnectorEntityOutput, error)
    DescribeConnectorEntityAsync(ctx workflow.Context, input *appflow.DescribeConnectorEntityInput) *AppflowDescribeConnectorEntityResult

    DescribeConnectorProfiles(ctx workflow.Context, input *appflow.DescribeConnectorProfilesInput) (*appflow.DescribeConnectorProfilesOutput, error)
    DescribeConnectorProfilesAsync(ctx workflow.Context, input *appflow.DescribeConnectorProfilesInput) *AppflowDescribeConnectorProfilesResult

    DescribeConnectors(ctx workflow.Context, input *appflow.DescribeConnectorsInput) (*appflow.DescribeConnectorsOutput, error)
    DescribeConnectorsAsync(ctx workflow.Context, input *appflow.DescribeConnectorsInput) *AppflowDescribeConnectorsResult

    DescribeFlow(ctx workflow.Context, input *appflow.DescribeFlowInput) (*appflow.DescribeFlowOutput, error)
    DescribeFlowAsync(ctx workflow.Context, input *appflow.DescribeFlowInput) *AppflowDescribeFlowResult

    DescribeFlowExecutionRecords(ctx workflow.Context, input *appflow.DescribeFlowExecutionRecordsInput) (*appflow.DescribeFlowExecutionRecordsOutput, error)
    DescribeFlowExecutionRecordsAsync(ctx workflow.Context, input *appflow.DescribeFlowExecutionRecordsInput) *AppflowDescribeFlowExecutionRecordsResult

    ListConnectorEntities(ctx workflow.Context, input *appflow.ListConnectorEntitiesInput) (*appflow.ListConnectorEntitiesOutput, error)
    ListConnectorEntitiesAsync(ctx workflow.Context, input *appflow.ListConnectorEntitiesInput) *AppflowListConnectorEntitiesResult

    ListFlows(ctx workflow.Context, input *appflow.ListFlowsInput) (*appflow.ListFlowsOutput, error)
    ListFlowsAsync(ctx workflow.Context, input *appflow.ListFlowsInput) *AppflowListFlowsResult

    ListTagsForResource(ctx workflow.Context, input *appflow.ListTagsForResourceInput) (*appflow.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *appflow.ListTagsForResourceInput) *AppflowListTagsForResourceResult

    StartFlow(ctx workflow.Context, input *appflow.StartFlowInput) (*appflow.StartFlowOutput, error)
    StartFlowAsync(ctx workflow.Context, input *appflow.StartFlowInput) *AppflowStartFlowResult

    StopFlow(ctx workflow.Context, input *appflow.StopFlowInput) (*appflow.StopFlowOutput, error)
    StopFlowAsync(ctx workflow.Context, input *appflow.StopFlowInput) *AppflowStopFlowResult

    TagResource(ctx workflow.Context, input *appflow.TagResourceInput) (*appflow.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *appflow.TagResourceInput) *AppflowTagResourceResult

    UntagResource(ctx workflow.Context, input *appflow.UntagResourceInput) (*appflow.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *appflow.UntagResourceInput) *AppflowUntagResourceResult

    UpdateConnectorProfile(ctx workflow.Context, input *appflow.UpdateConnectorProfileInput) (*appflow.UpdateConnectorProfileOutput, error)
    UpdateConnectorProfileAsync(ctx workflow.Context, input *appflow.UpdateConnectorProfileInput) *AppflowUpdateConnectorProfileResult

    UpdateFlow(ctx workflow.Context, input *appflow.UpdateFlowInput) (*appflow.UpdateFlowOutput, error)
    UpdateFlowAsync(ctx workflow.Context, input *appflow.UpdateFlowInput) *AppflowUpdateFlowResult
}
type AppflowCreateConnectorProfileResult struct {
	Result workflow.Future
}

func (r *AppflowCreateConnectorProfileResult) Get(ctx workflow.Context) (*appflow.CreateConnectorProfileOutput, error) {
    var output appflow.CreateConnectorProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowCreateFlowResult struct {
	Result workflow.Future
}

func (r *AppflowCreateFlowResult) Get(ctx workflow.Context) (*appflow.CreateFlowOutput, error) {
    var output appflow.CreateFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDeleteConnectorProfileResult struct {
	Result workflow.Future
}

func (r *AppflowDeleteConnectorProfileResult) Get(ctx workflow.Context) (*appflow.DeleteConnectorProfileOutput, error) {
    var output appflow.DeleteConnectorProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDeleteFlowResult struct {
	Result workflow.Future
}

func (r *AppflowDeleteFlowResult) Get(ctx workflow.Context) (*appflow.DeleteFlowOutput, error) {
    var output appflow.DeleteFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDescribeConnectorEntityResult struct {
	Result workflow.Future
}

func (r *AppflowDescribeConnectorEntityResult) Get(ctx workflow.Context) (*appflow.DescribeConnectorEntityOutput, error) {
    var output appflow.DescribeConnectorEntityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDescribeConnectorProfilesResult struct {
	Result workflow.Future
}

func (r *AppflowDescribeConnectorProfilesResult) Get(ctx workflow.Context) (*appflow.DescribeConnectorProfilesOutput, error) {
    var output appflow.DescribeConnectorProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDescribeConnectorsResult struct {
	Result workflow.Future
}

func (r *AppflowDescribeConnectorsResult) Get(ctx workflow.Context) (*appflow.DescribeConnectorsOutput, error) {
    var output appflow.DescribeConnectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDescribeFlowResult struct {
	Result workflow.Future
}

func (r *AppflowDescribeFlowResult) Get(ctx workflow.Context) (*appflow.DescribeFlowOutput, error) {
    var output appflow.DescribeFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowDescribeFlowExecutionRecordsResult struct {
	Result workflow.Future
}

func (r *AppflowDescribeFlowExecutionRecordsResult) Get(ctx workflow.Context) (*appflow.DescribeFlowExecutionRecordsOutput, error) {
    var output appflow.DescribeFlowExecutionRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowListConnectorEntitiesResult struct {
	Result workflow.Future
}

func (r *AppflowListConnectorEntitiesResult) Get(ctx workflow.Context) (*appflow.ListConnectorEntitiesOutput, error) {
    var output appflow.ListConnectorEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowListFlowsResult struct {
	Result workflow.Future
}

func (r *AppflowListFlowsResult) Get(ctx workflow.Context) (*appflow.ListFlowsOutput, error) {
    var output appflow.ListFlowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AppflowListTagsForResourceResult) Get(ctx workflow.Context) (*appflow.ListTagsForResourceOutput, error) {
    var output appflow.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowStartFlowResult struct {
	Result workflow.Future
}

func (r *AppflowStartFlowResult) Get(ctx workflow.Context) (*appflow.StartFlowOutput, error) {
    var output appflow.StartFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowStopFlowResult struct {
	Result workflow.Future
}

func (r *AppflowStopFlowResult) Get(ctx workflow.Context) (*appflow.StopFlowOutput, error) {
    var output appflow.StopFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowTagResourceResult struct {
	Result workflow.Future
}

func (r *AppflowTagResourceResult) Get(ctx workflow.Context) (*appflow.TagResourceOutput, error) {
    var output appflow.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowUntagResourceResult struct {
	Result workflow.Future
}

func (r *AppflowUntagResourceResult) Get(ctx workflow.Context) (*appflow.UntagResourceOutput, error) {
    var output appflow.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowUpdateConnectorProfileResult struct {
	Result workflow.Future
}

func (r *AppflowUpdateConnectorProfileResult) Get(ctx workflow.Context) (*appflow.UpdateConnectorProfileOutput, error) {
    var output appflow.UpdateConnectorProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppflowUpdateFlowResult struct {
	Result workflow.Future
}

func (r *AppflowUpdateFlowResult) Get(ctx workflow.Context) (*appflow.UpdateFlowOutput, error) {
    var output appflow.UpdateFlowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AppflowStub struct {
    activities AppflowClient
}

func NewAppflowStub() AppflowClient {
    return &AppflowStub{}
}

func (a *AppflowStub) CreateConnectorProfile(ctx workflow.Context, input *appflow.CreateConnectorProfileInput) (*appflow.CreateConnectorProfileOutput, error) {
    var output appflow.CreateConnectorProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConnectorProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) CreateFlow(ctx workflow.Context, input *appflow.CreateFlowInput) (*appflow.CreateFlowOutput, error) {
    var output appflow.CreateFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DeleteConnectorProfile(ctx workflow.Context, input *appflow.DeleteConnectorProfileInput) (*appflow.DeleteConnectorProfileOutput, error) {
    var output appflow.DeleteConnectorProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConnectorProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DeleteFlow(ctx workflow.Context, input *appflow.DeleteFlowInput) (*appflow.DeleteFlowOutput, error) {
    var output appflow.DeleteFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DescribeConnectorEntity(ctx workflow.Context, input *appflow.DescribeConnectorEntityInput) (*appflow.DescribeConnectorEntityOutput, error) {
    var output appflow.DescribeConnectorEntityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectorEntity, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DescribeConnectorProfiles(ctx workflow.Context, input *appflow.DescribeConnectorProfilesInput) (*appflow.DescribeConnectorProfilesOutput, error) {
    var output appflow.DescribeConnectorProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectorProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DescribeConnectors(ctx workflow.Context, input *appflow.DescribeConnectorsInput) (*appflow.DescribeConnectorsOutput, error) {
    var output appflow.DescribeConnectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectors, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DescribeFlow(ctx workflow.Context, input *appflow.DescribeFlowInput) (*appflow.DescribeFlowOutput, error) {
    var output appflow.DescribeFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) DescribeFlowExecutionRecords(ctx workflow.Context, input *appflow.DescribeFlowExecutionRecordsInput) (*appflow.DescribeFlowExecutionRecordsOutput, error) {
    var output appflow.DescribeFlowExecutionRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFlowExecutionRecords, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) ListConnectorEntities(ctx workflow.Context, input *appflow.ListConnectorEntitiesInput) (*appflow.ListConnectorEntitiesOutput, error) {
    var output appflow.ListConnectorEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConnectorEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) ListFlows(ctx workflow.Context, input *appflow.ListFlowsInput) (*appflow.ListFlowsOutput, error) {
    var output appflow.ListFlowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFlows, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) ListTagsForResource(ctx workflow.Context, input *appflow.ListTagsForResourceInput) (*appflow.ListTagsForResourceOutput, error) {
    var output appflow.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) StartFlow(ctx workflow.Context, input *appflow.StartFlowInput) (*appflow.StartFlowOutput, error) {
    var output appflow.StartFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) StopFlow(ctx workflow.Context, input *appflow.StopFlowInput) (*appflow.StopFlowOutput, error) {
    var output appflow.StopFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopFlow, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) TagResource(ctx workflow.Context, input *appflow.TagResourceInput) (*appflow.TagResourceOutput, error) {
    var output appflow.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) UntagResource(ctx workflow.Context, input *appflow.UntagResourceInput) (*appflow.UntagResourceOutput, error) {
    var output appflow.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) UpdateConnectorProfile(ctx workflow.Context, input *appflow.UpdateConnectorProfileInput) (*appflow.UpdateConnectorProfileOutput, error) {
    var output appflow.UpdateConnectorProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConnectorProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppflowStub) UpdateFlow(ctx workflow.Context, input *appflow.UpdateFlowInput) (*appflow.UpdateFlowOutput, error) {
    var output appflow.UpdateFlowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFlow, input).Get(ctx, &output)
    return &output, err
}
