package awsclients

import (
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ResourceGroupsTaggingAPIClient interface {
       DescribeReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error)
       DescribeReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) *ResourcegroupstaggingapiDescribeReportCreationResult

       GetComplianceSummary(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error)
       GetComplianceSummaryAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) *ResourcegroupstaggingapiGetComplianceSummaryResult

       GetResources(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error)
       GetResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) *ResourcegroupstaggingapiGetResourcesResult

       GetTagKeys(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
       GetTagKeysAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) *ResourcegroupstaggingapiGetTagKeysResult

       GetTagValues(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
       GetTagValuesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) *ResourcegroupstaggingapiGetTagValuesResult

       StartReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error)
       StartReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) *ResourcegroupstaggingapiStartReportCreationResult

       TagResources(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error)
       TagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) *ResourcegroupstaggingapiTagResourcesResult

       UntagResources(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error)
       UntagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) *ResourcegroupstaggingapiUntagResourcesResult
}

type ResourcegroupstaggingapiDescribeReportCreationResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiDescribeReportCreationResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
    var output resourcegroupstaggingapi.DescribeReportCreationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiGetComplianceSummaryResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiGetComplianceSummaryResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
    var output resourcegroupstaggingapi.GetComplianceSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiGetResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiGetResourcesResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
    var output resourcegroupstaggingapi.GetResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiGetTagKeysResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiGetTagKeysResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
    var output resourcegroupstaggingapi.GetTagKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiGetTagValuesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiGetTagValuesResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
    var output resourcegroupstaggingapi.GetTagValuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiStartReportCreationResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiStartReportCreationResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
    var output resourcegroupstaggingapi.StartReportCreationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiTagResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiTagResourcesResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
    var output resourcegroupstaggingapi.TagResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourcegroupstaggingapiUntagResourcesResult struct {
	Result workflow.Future
}

func (r *ResourcegroupstaggingapiUntagResourcesResult) Get(ctx workflow.Context) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
    var output resourcegroupstaggingapi.UntagResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ResourceGroupsTaggingAPIStub struct {
    activities awsactivities.ResourceGroupsTaggingAPIActivities
}

func NewResourceGroupsTaggingAPIStub() ResourceGroupsTaggingAPIClient {
    return &ResourceGroupsTaggingAPIStub{}
}

func (a *ResourceGroupsTaggingAPIStub) DescribeReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
    var output resourcegroupstaggingapi.DescribeReportCreationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReportCreation, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) DescribeReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput) *ResourcegroupstaggingapiDescribeReportCreationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReportCreation, input)
    return &ResourcegroupstaggingapiDescribeReportCreationResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) GetComplianceSummary(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
    var output resourcegroupstaggingapi.GetComplianceSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComplianceSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) GetComplianceSummaryAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput) *ResourcegroupstaggingapiGetComplianceSummaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetComplianceSummary, input)
    return &ResourcegroupstaggingapiGetComplianceSummaryResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) GetResources(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
    var output resourcegroupstaggingapi.GetResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) GetResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetResourcesInput) *ResourcegroupstaggingapiGetResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResources, input)
    return &ResourcegroupstaggingapiGetResourcesResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) GetTagKeys(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
    var output resourcegroupstaggingapi.GetTagKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTagKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) GetTagKeysAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagKeysInput) *ResourcegroupstaggingapiGetTagKeysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTagKeys, input)
    return &ResourcegroupstaggingapiGetTagKeysResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) GetTagValues(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
    var output resourcegroupstaggingapi.GetTagValuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTagValues, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) GetTagValuesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.GetTagValuesInput) *ResourcegroupstaggingapiGetTagValuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTagValues, input)
    return &ResourcegroupstaggingapiGetTagValuesResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) StartReportCreation(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
    var output resourcegroupstaggingapi.StartReportCreationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartReportCreation, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) StartReportCreationAsync(ctx workflow.Context, input *resourcegroupstaggingapi.StartReportCreationInput) *ResourcegroupstaggingapiStartReportCreationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartReportCreation, input)
    return &ResourcegroupstaggingapiStartReportCreationResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) TagResources(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
    var output resourcegroupstaggingapi.TagResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) TagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.TagResourcesInput) *ResourcegroupstaggingapiTagResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResources, input)
    return &ResourcegroupstaggingapiTagResourcesResult{Result: future}
}

func (a *ResourceGroupsTaggingAPIStub) UntagResources(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
    var output resourcegroupstaggingapi.UntagResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ResourceGroupsTaggingAPIStub) UntagResourcesAsync(ctx workflow.Context, input *resourcegroupstaggingapi.UntagResourcesInput) *ResourcegroupstaggingapiUntagResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResources, input)
    return &ResourcegroupstaggingapiUntagResourcesResult{Result: future}
}
