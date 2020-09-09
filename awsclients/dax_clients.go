package awsclients

import (
	"github.com/aws/aws-sdk-go/service/dax"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DAXClient interface {
	CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *DaxCreateClusterResult

	CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error)
	CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *DaxCreateParameterGroupResult

	CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error)
	CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *DaxCreateSubnetGroupResult

	DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error)
	DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DaxDecreaseReplicationFactorResult

	DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DaxDeleteClusterResult

	DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error)
	DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DaxDeleteParameterGroupResult

	DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error)
	DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DaxDeleteSubnetGroupResult

	DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DaxDescribeClustersResult

	DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error)
	DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DaxDescribeDefaultParametersResult

	DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DaxDescribeEventsResult

	DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error)
	DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DaxDescribeParameterGroupsResult

	DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error)
	DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DaxDescribeParametersResult

	DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error)
	DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DaxDescribeSubnetGroupsResult

	IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error)
	IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *DaxIncreaseReplicationFactorResult

	ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *DaxListTagsResult

	RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error)
	RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *DaxRebootNodeResult

	TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *DaxTagResourceResult

	UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *DaxUntagResourceResult

	UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error)
	UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *DaxUpdateClusterResult

	UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error)
	UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *DaxUpdateParameterGroupResult

	UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error)
	UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *DaxUpdateSubnetGroupResult
}

type DaxCreateClusterResult struct {
	Result workflow.Future
}

func (r *DaxCreateClusterResult) Get(ctx workflow.Context) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxCreateParameterGroupResult struct {
	Result workflow.Future
}

func (r *DaxCreateParameterGroupResult) Get(ctx workflow.Context) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxCreateSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DaxCreateSubnetGroupResult) Get(ctx workflow.Context) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDecreaseReplicationFactorResult struct {
	Result workflow.Future
}

func (r *DaxDecreaseReplicationFactorResult) Get(ctx workflow.Context) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDeleteClusterResult struct {
	Result workflow.Future
}

func (r *DaxDeleteClusterResult) Get(ctx workflow.Context) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDeleteParameterGroupResult struct {
	Result workflow.Future
}

func (r *DaxDeleteParameterGroupResult) Get(ctx workflow.Context) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDeleteSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DaxDeleteSubnetGroupResult) Get(ctx workflow.Context) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDescribeClustersResult struct {
	Result workflow.Future
}

func (r *DaxDescribeClustersResult) Get(ctx workflow.Context) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDescribeDefaultParametersResult struct {
	Result workflow.Future
}

func (r *DaxDescribeDefaultParametersResult) Get(ctx workflow.Context) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDescribeEventsResult struct {
	Result workflow.Future
}

func (r *DaxDescribeEventsResult) Get(ctx workflow.Context) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDescribeParameterGroupsResult struct {
	Result workflow.Future
}

func (r *DaxDescribeParameterGroupsResult) Get(ctx workflow.Context) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDescribeParametersResult struct {
	Result workflow.Future
}

func (r *DaxDescribeParametersResult) Get(ctx workflow.Context) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxDescribeSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *DaxDescribeSubnetGroupsResult) Get(ctx workflow.Context) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxIncreaseReplicationFactorResult struct {
	Result workflow.Future
}

func (r *DaxIncreaseReplicationFactorResult) Get(ctx workflow.Context) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxListTagsResult struct {
	Result workflow.Future
}

func (r *DaxListTagsResult) Get(ctx workflow.Context) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxRebootNodeResult struct {
	Result workflow.Future
}

func (r *DaxRebootNodeResult) Get(ctx workflow.Context) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxTagResourceResult struct {
	Result workflow.Future
}

func (r *DaxTagResourceResult) Get(ctx workflow.Context) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxUntagResourceResult struct {
	Result workflow.Future
}

func (r *DaxUntagResourceResult) Get(ctx workflow.Context) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxUpdateClusterResult struct {
	Result workflow.Future
}

func (r *DaxUpdateClusterResult) Get(ctx workflow.Context) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxUpdateParameterGroupResult struct {
	Result workflow.Future
}

func (r *DaxUpdateParameterGroupResult) Get(ctx workflow.Context) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DaxUpdateSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DaxUpdateSubnetGroupResult) Get(ctx workflow.Context) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type DAXStub struct {
	activities awsactivities.DAXActivities
}

func NewDAXStub() DAXClient {
	return &DAXStub{}
}

func (a *DAXStub) CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *DaxCreateClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input)
	return &DaxCreateClusterResult{Result: future}
}

func (a *DAXStub) CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *DaxCreateParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateParameterGroup, input)
	return &DaxCreateParameterGroupResult{Result: future}
}

func (a *DAXStub) CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateSubnetGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *DaxCreateSubnetGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateSubnetGroup, input)
	return &DaxCreateSubnetGroupResult{Result: future}
}

func (a *DAXStub) DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DecreaseReplicationFactor, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DaxDecreaseReplicationFactorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DecreaseReplicationFactor, input)
	return &DaxDecreaseReplicationFactorResult{Result: future}
}

func (a *DAXStub) DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DaxDeleteClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input)
	return &DaxDeleteClusterResult{Result: future}
}

func (a *DAXStub) DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DaxDeleteParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteParameterGroup, input)
	return &DaxDeleteParameterGroupResult{Result: future}
}

func (a *DAXStub) DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteSubnetGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DaxDeleteSubnetGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteSubnetGroup, input)
	return &DaxDeleteSubnetGroupResult{Result: future}
}

func (a *DAXStub) DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DaxDescribeClustersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input)
	return &DaxDescribeClustersResult{Result: future}
}

func (a *DAXStub) DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeDefaultParameters, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DaxDescribeDefaultParametersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeDefaultParameters, input)
	return &DaxDescribeDefaultParametersResult{Result: future}
}

func (a *DAXStub) DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DaxDescribeEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input)
	return &DaxDescribeEventsResult{Result: future}
}

func (a *DAXStub) DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeParameterGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DaxDescribeParameterGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeParameterGroups, input)
	return &DaxDescribeParameterGroupsResult{Result: future}
}

func (a *DAXStub) DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeParameters, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DaxDescribeParametersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeParameters, input)
	return &DaxDescribeParametersResult{Result: future}
}

func (a *DAXStub) DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeSubnetGroups, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DaxDescribeSubnetGroupsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeSubnetGroups, input)
	return &DaxDescribeSubnetGroupsResult{Result: future}
}

func (a *DAXStub) IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, a.activities.IncreaseReplicationFactor, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *DaxIncreaseReplicationFactorResult {
	future := workflow.ExecuteActivity(ctx, a.activities.IncreaseReplicationFactor, input)
	return &DaxIncreaseReplicationFactorResult{Result: future}
}

func (a *DAXStub) ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *DaxListTagsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
	return &DaxListTagsResult{Result: future}
}

func (a *DAXStub) RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RebootNode, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *DaxRebootNodeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RebootNode, input)
	return &DaxRebootNodeResult{Result: future}
}

func (a *DAXStub) TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *DaxTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &DaxTagResourceResult{Result: future}
}

func (a *DAXStub) UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *DaxUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &DaxUntagResourceResult{Result: future}
}

func (a *DAXStub) UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *DaxUpdateClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateCluster, input)
	return &DaxUpdateClusterResult{Result: future}
}

func (a *DAXStub) UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateParameterGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *DaxUpdateParameterGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateParameterGroup, input)
	return &DaxUpdateParameterGroupResult{Result: future}
}

func (a *DAXStub) UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateSubnetGroup, input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *DaxUpdateSubnetGroupResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateSubnetGroup, input)
	return &DaxUpdateSubnetGroupResult{Result: future}
}
