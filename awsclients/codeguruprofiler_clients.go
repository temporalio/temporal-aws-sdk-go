package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CodeGuruProfilerClient interface {
    AddNotificationChannels(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error)
    AddNotificationChannelsAsync(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) *CodeguruprofilerAddNotificationChannelsResult

    BatchGetFrameMetricData(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error)
    BatchGetFrameMetricDataAsync(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) *CodeguruprofilerBatchGetFrameMetricDataResult

    ConfigureAgent(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error)
    ConfigureAgentAsync(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) *CodeguruprofilerConfigureAgentResult

    CreateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error)
    CreateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) *CodeguruprofilerCreateProfilingGroupResult

    DeleteProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error)
    DeleteProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) *CodeguruprofilerDeleteProfilingGroupResult

    DescribeProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error)
    DescribeProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) *CodeguruprofilerDescribeProfilingGroupResult

    GetFindingsReportAccountSummary(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error)
    GetFindingsReportAccountSummaryAsync(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) *CodeguruprofilerGetFindingsReportAccountSummaryResult

    GetNotificationConfiguration(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error)
    GetNotificationConfigurationAsync(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) *CodeguruprofilerGetNotificationConfigurationResult

    GetPolicy(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error)
    GetPolicyAsync(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) *CodeguruprofilerGetPolicyResult

    GetProfile(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error)
    GetProfileAsync(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) *CodeguruprofilerGetProfileResult

    GetRecommendations(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error)
    GetRecommendationsAsync(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) *CodeguruprofilerGetRecommendationsResult

    ListFindingsReports(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error)
    ListFindingsReportsAsync(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) *CodeguruprofilerListFindingsReportsResult

    ListProfileTimes(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error)
    ListProfileTimesAsync(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) *CodeguruprofilerListProfileTimesResult

    ListProfilingGroups(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error)
    ListProfilingGroupsAsync(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) *CodeguruprofilerListProfilingGroupsResult

    ListTagsForResource(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) *CodeguruprofilerListTagsForResourceResult

    PostAgentProfile(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error)
    PostAgentProfileAsync(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) *CodeguruprofilerPostAgentProfileResult

    PutPermission(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error)
    PutPermissionAsync(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) *CodeguruprofilerPutPermissionResult

    RemoveNotificationChannel(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error)
    RemoveNotificationChannelAsync(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) *CodeguruprofilerRemoveNotificationChannelResult

    RemovePermission(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error)
    RemovePermissionAsync(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) *CodeguruprofilerRemovePermissionResult

    SubmitFeedback(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error)
    SubmitFeedbackAsync(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) *CodeguruprofilerSubmitFeedbackResult

    TagResource(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) *CodeguruprofilerTagResourceResult

    UntagResource(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) *CodeguruprofilerUntagResourceResult

    UpdateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error)
    UpdateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) *CodeguruprofilerUpdateProfilingGroupResult
}

type CodeguruprofilerAddNotificationChannelsResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerAddNotificationChannelsResult) Get(ctx workflow.Context) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
    var output codeguruprofiler.AddNotificationChannelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerBatchGetFrameMetricDataResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerBatchGetFrameMetricDataResult) Get(ctx workflow.Context) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
    var output codeguruprofiler.BatchGetFrameMetricDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerConfigureAgentResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerConfigureAgentResult) Get(ctx workflow.Context) (*codeguruprofiler.ConfigureAgentOutput, error) {
    var output codeguruprofiler.ConfigureAgentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerCreateProfilingGroupResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerCreateProfilingGroupResult) Get(ctx workflow.Context) (*codeguruprofiler.CreateProfilingGroupOutput, error) {
    var output codeguruprofiler.CreateProfilingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerDeleteProfilingGroupResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerDeleteProfilingGroupResult) Get(ctx workflow.Context) (*codeguruprofiler.DeleteProfilingGroupOutput, error) {
    var output codeguruprofiler.DeleteProfilingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerDescribeProfilingGroupResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerDescribeProfilingGroupResult) Get(ctx workflow.Context) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
    var output codeguruprofiler.DescribeProfilingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerGetFindingsReportAccountSummaryResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerGetFindingsReportAccountSummaryResult) Get(ctx workflow.Context) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
    var output codeguruprofiler.GetFindingsReportAccountSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerGetNotificationConfigurationResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerGetNotificationConfigurationResult) Get(ctx workflow.Context) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
    var output codeguruprofiler.GetNotificationConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerGetPolicyResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerGetPolicyResult) Get(ctx workflow.Context) (*codeguruprofiler.GetPolicyOutput, error) {
    var output codeguruprofiler.GetPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerGetProfileResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerGetProfileResult) Get(ctx workflow.Context) (*codeguruprofiler.GetProfileOutput, error) {
    var output codeguruprofiler.GetProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerGetRecommendationsResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerGetRecommendationsResult) Get(ctx workflow.Context) (*codeguruprofiler.GetRecommendationsOutput, error) {
    var output codeguruprofiler.GetRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerListFindingsReportsResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerListFindingsReportsResult) Get(ctx workflow.Context) (*codeguruprofiler.ListFindingsReportsOutput, error) {
    var output codeguruprofiler.ListFindingsReportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerListProfileTimesResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerListProfileTimesResult) Get(ctx workflow.Context) (*codeguruprofiler.ListProfileTimesOutput, error) {
    var output codeguruprofiler.ListProfileTimesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerListProfilingGroupsResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerListProfilingGroupsResult) Get(ctx workflow.Context) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
    var output codeguruprofiler.ListProfilingGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerListTagsForResourceResult) Get(ctx workflow.Context) (*codeguruprofiler.ListTagsForResourceOutput, error) {
    var output codeguruprofiler.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerPostAgentProfileResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerPostAgentProfileResult) Get(ctx workflow.Context) (*codeguruprofiler.PostAgentProfileOutput, error) {
    var output codeguruprofiler.PostAgentProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerPutPermissionResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerPutPermissionResult) Get(ctx workflow.Context) (*codeguruprofiler.PutPermissionOutput, error) {
    var output codeguruprofiler.PutPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerRemoveNotificationChannelResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerRemoveNotificationChannelResult) Get(ctx workflow.Context) (*codeguruprofiler.RemoveNotificationChannelOutput, error) {
    var output codeguruprofiler.RemoveNotificationChannelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerRemovePermissionResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerRemovePermissionResult) Get(ctx workflow.Context) (*codeguruprofiler.RemovePermissionOutput, error) {
    var output codeguruprofiler.RemovePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerSubmitFeedbackResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerSubmitFeedbackResult) Get(ctx workflow.Context) (*codeguruprofiler.SubmitFeedbackOutput, error) {
    var output codeguruprofiler.SubmitFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerTagResourceResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerTagResourceResult) Get(ctx workflow.Context) (*codeguruprofiler.TagResourceOutput, error) {
    var output codeguruprofiler.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerUntagResourceResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerUntagResourceResult) Get(ctx workflow.Context) (*codeguruprofiler.UntagResourceOutput, error) {
    var output codeguruprofiler.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeguruprofilerUpdateProfilingGroupResult struct {
	Result workflow.Future
}

func (r *CodeguruprofilerUpdateProfilingGroupResult) Get(ctx workflow.Context) (*codeguruprofiler.UpdateProfilingGroupOutput, error) {
    var output codeguruprofiler.UpdateProfilingGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeGuruProfilerStub struct {
    activities awsactivities.CodeGuruProfilerActivities
}

func NewCodeGuruProfilerStub() CodeGuruProfilerClient {
    return &CodeGuruProfilerStub{}
}

func (a *CodeGuruProfilerStub) AddNotificationChannels(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
    var output codeguruprofiler.AddNotificationChannelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddNotificationChannels, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) AddNotificationChannelsAsync(ctx workflow.Context, input *codeguruprofiler.AddNotificationChannelsInput) *CodeguruprofilerAddNotificationChannelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddNotificationChannels, input)
    return &CodeguruprofilerAddNotificationChannelsResult{Result: future}
}

func (a *CodeGuruProfilerStub) BatchGetFrameMetricData(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
    var output codeguruprofiler.BatchGetFrameMetricDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetFrameMetricData, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) BatchGetFrameMetricDataAsync(ctx workflow.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput) *CodeguruprofilerBatchGetFrameMetricDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetFrameMetricData, input)
    return &CodeguruprofilerBatchGetFrameMetricDataResult{Result: future}
}

func (a *CodeGuruProfilerStub) ConfigureAgent(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error) {
    var output codeguruprofiler.ConfigureAgentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfigureAgent, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) ConfigureAgentAsync(ctx workflow.Context, input *codeguruprofiler.ConfigureAgentInput) *CodeguruprofilerConfigureAgentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfigureAgent, input)
    return &CodeguruprofilerConfigureAgentResult{Result: future}
}

func (a *CodeGuruProfilerStub) CreateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error) {
    var output codeguruprofiler.CreateProfilingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProfilingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) CreateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.CreateProfilingGroupInput) *CodeguruprofilerCreateProfilingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProfilingGroup, input)
    return &CodeguruprofilerCreateProfilingGroupResult{Result: future}
}

func (a *CodeGuruProfilerStub) DeleteProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error) {
    var output codeguruprofiler.DeleteProfilingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProfilingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) DeleteProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DeleteProfilingGroupInput) *CodeguruprofilerDeleteProfilingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProfilingGroup, input)
    return &CodeguruprofilerDeleteProfilingGroupResult{Result: future}
}

func (a *CodeGuruProfilerStub) DescribeProfilingGroup(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
    var output codeguruprofiler.DescribeProfilingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProfilingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) DescribeProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.DescribeProfilingGroupInput) *CodeguruprofilerDescribeProfilingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProfilingGroup, input)
    return &CodeguruprofilerDescribeProfilingGroupResult{Result: future}
}

func (a *CodeGuruProfilerStub) GetFindingsReportAccountSummary(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
    var output codeguruprofiler.GetFindingsReportAccountSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFindingsReportAccountSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) GetFindingsReportAccountSummaryAsync(ctx workflow.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput) *CodeguruprofilerGetFindingsReportAccountSummaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFindingsReportAccountSummary, input)
    return &CodeguruprofilerGetFindingsReportAccountSummaryResult{Result: future}
}

func (a *CodeGuruProfilerStub) GetNotificationConfiguration(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
    var output codeguruprofiler.GetNotificationConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetNotificationConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) GetNotificationConfigurationAsync(ctx workflow.Context, input *codeguruprofiler.GetNotificationConfigurationInput) *CodeguruprofilerGetNotificationConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetNotificationConfiguration, input)
    return &CodeguruprofilerGetNotificationConfigurationResult{Result: future}
}

func (a *CodeGuruProfilerStub) GetPolicy(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error) {
    var output codeguruprofiler.GetPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) GetPolicyAsync(ctx workflow.Context, input *codeguruprofiler.GetPolicyInput) *CodeguruprofilerGetPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPolicy, input)
    return &CodeguruprofilerGetPolicyResult{Result: future}
}

func (a *CodeGuruProfilerStub) GetProfile(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error) {
    var output codeguruprofiler.GetProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) GetProfileAsync(ctx workflow.Context, input *codeguruprofiler.GetProfileInput) *CodeguruprofilerGetProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetProfile, input)
    return &CodeguruprofilerGetProfileResult{Result: future}
}

func (a *CodeGuruProfilerStub) GetRecommendations(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error) {
    var output codeguruprofiler.GetRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecommendations, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) GetRecommendationsAsync(ctx workflow.Context, input *codeguruprofiler.GetRecommendationsInput) *CodeguruprofilerGetRecommendationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRecommendations, input)
    return &CodeguruprofilerGetRecommendationsResult{Result: future}
}

func (a *CodeGuruProfilerStub) ListFindingsReports(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error) {
    var output codeguruprofiler.ListFindingsReportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFindingsReports, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) ListFindingsReportsAsync(ctx workflow.Context, input *codeguruprofiler.ListFindingsReportsInput) *CodeguruprofilerListFindingsReportsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFindingsReports, input)
    return &CodeguruprofilerListFindingsReportsResult{Result: future}
}

func (a *CodeGuruProfilerStub) ListProfileTimes(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error) {
    var output codeguruprofiler.ListProfileTimesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProfileTimes, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) ListProfileTimesAsync(ctx workflow.Context, input *codeguruprofiler.ListProfileTimesInput) *CodeguruprofilerListProfileTimesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProfileTimes, input)
    return &CodeguruprofilerListProfileTimesResult{Result: future}
}

func (a *CodeGuruProfilerStub) ListProfilingGroups(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
    var output codeguruprofiler.ListProfilingGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProfilingGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) ListProfilingGroupsAsync(ctx workflow.Context, input *codeguruprofiler.ListProfilingGroupsInput) *CodeguruprofilerListProfilingGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProfilingGroups, input)
    return &CodeguruprofilerListProfilingGroupsResult{Result: future}
}

func (a *CodeGuruProfilerStub) ListTagsForResource(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error) {
    var output codeguruprofiler.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) ListTagsForResourceAsync(ctx workflow.Context, input *codeguruprofiler.ListTagsForResourceInput) *CodeguruprofilerListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &CodeguruprofilerListTagsForResourceResult{Result: future}
}

func (a *CodeGuruProfilerStub) PostAgentProfile(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error) {
    var output codeguruprofiler.PostAgentProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PostAgentProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) PostAgentProfileAsync(ctx workflow.Context, input *codeguruprofiler.PostAgentProfileInput) *CodeguruprofilerPostAgentProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PostAgentProfile, input)
    return &CodeguruprofilerPostAgentProfileResult{Result: future}
}

func (a *CodeGuruProfilerStub) PutPermission(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error) {
    var output codeguruprofiler.PutPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) PutPermissionAsync(ctx workflow.Context, input *codeguruprofiler.PutPermissionInput) *CodeguruprofilerPutPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutPermission, input)
    return &CodeguruprofilerPutPermissionResult{Result: future}
}

func (a *CodeGuruProfilerStub) RemoveNotificationChannel(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error) {
    var output codeguruprofiler.RemoveNotificationChannelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveNotificationChannel, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) RemoveNotificationChannelAsync(ctx workflow.Context, input *codeguruprofiler.RemoveNotificationChannelInput) *CodeguruprofilerRemoveNotificationChannelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveNotificationChannel, input)
    return &CodeguruprofilerRemoveNotificationChannelResult{Result: future}
}

func (a *CodeGuruProfilerStub) RemovePermission(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error) {
    var output codeguruprofiler.RemovePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) RemovePermissionAsync(ctx workflow.Context, input *codeguruprofiler.RemovePermissionInput) *CodeguruprofilerRemovePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input)
    return &CodeguruprofilerRemovePermissionResult{Result: future}
}

func (a *CodeGuruProfilerStub) SubmitFeedback(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error) {
    var output codeguruprofiler.SubmitFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubmitFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) SubmitFeedbackAsync(ctx workflow.Context, input *codeguruprofiler.SubmitFeedbackInput) *CodeguruprofilerSubmitFeedbackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubmitFeedback, input)
    return &CodeguruprofilerSubmitFeedbackResult{Result: future}
}

func (a *CodeGuruProfilerStub) TagResource(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error) {
    var output codeguruprofiler.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) TagResourceAsync(ctx workflow.Context, input *codeguruprofiler.TagResourceInput) *CodeguruprofilerTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &CodeguruprofilerTagResourceResult{Result: future}
}

func (a *CodeGuruProfilerStub) UntagResource(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error) {
    var output codeguruprofiler.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) UntagResourceAsync(ctx workflow.Context, input *codeguruprofiler.UntagResourceInput) *CodeguruprofilerUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &CodeguruprofilerUntagResourceResult{Result: future}
}

func (a *CodeGuruProfilerStub) UpdateProfilingGroup(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error) {
    var output codeguruprofiler.UpdateProfilingGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProfilingGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruProfilerStub) UpdateProfilingGroupAsync(ctx workflow.Context, input *codeguruprofiler.UpdateProfilingGroupInput) *CodeguruprofilerUpdateProfilingGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProfilingGroup, input)
    return &CodeguruprofilerUpdateProfilingGroupResult{Result: future}
}
