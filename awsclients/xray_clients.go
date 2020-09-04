package awsclients

import (
	"github.com/aws/aws-sdk-go/service/xray"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type XRayClient interface {
    BatchGetTraces(ctx workflow.Context, input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error)
    BatchGetTracesAsync(ctx workflow.Context, input *xray.BatchGetTracesInput) *XrayBatchGetTracesResult

    CreateGroup(ctx workflow.Context, input *xray.CreateGroupInput) (*xray.CreateGroupOutput, error)
    CreateGroupAsync(ctx workflow.Context, input *xray.CreateGroupInput) *XrayCreateGroupResult

    CreateSamplingRule(ctx workflow.Context, input *xray.CreateSamplingRuleInput) (*xray.CreateSamplingRuleOutput, error)
    CreateSamplingRuleAsync(ctx workflow.Context, input *xray.CreateSamplingRuleInput) *XrayCreateSamplingRuleResult

    DeleteGroup(ctx workflow.Context, input *xray.DeleteGroupInput) (*xray.DeleteGroupOutput, error)
    DeleteGroupAsync(ctx workflow.Context, input *xray.DeleteGroupInput) *XrayDeleteGroupResult

    DeleteSamplingRule(ctx workflow.Context, input *xray.DeleteSamplingRuleInput) (*xray.DeleteSamplingRuleOutput, error)
    DeleteSamplingRuleAsync(ctx workflow.Context, input *xray.DeleteSamplingRuleInput) *XrayDeleteSamplingRuleResult

    GetEncryptionConfig(ctx workflow.Context, input *xray.GetEncryptionConfigInput) (*xray.GetEncryptionConfigOutput, error)
    GetEncryptionConfigAsync(ctx workflow.Context, input *xray.GetEncryptionConfigInput) *XrayGetEncryptionConfigResult

    GetGroup(ctx workflow.Context, input *xray.GetGroupInput) (*xray.GetGroupOutput, error)
    GetGroupAsync(ctx workflow.Context, input *xray.GetGroupInput) *XrayGetGroupResult

    GetGroups(ctx workflow.Context, input *xray.GetGroupsInput) (*xray.GetGroupsOutput, error)
    GetGroupsAsync(ctx workflow.Context, input *xray.GetGroupsInput) *XrayGetGroupsResult

    GetSamplingRules(ctx workflow.Context, input *xray.GetSamplingRulesInput) (*xray.GetSamplingRulesOutput, error)
    GetSamplingRulesAsync(ctx workflow.Context, input *xray.GetSamplingRulesInput) *XrayGetSamplingRulesResult

    GetSamplingStatisticSummaries(ctx workflow.Context, input *xray.GetSamplingStatisticSummariesInput) (*xray.GetSamplingStatisticSummariesOutput, error)
    GetSamplingStatisticSummariesAsync(ctx workflow.Context, input *xray.GetSamplingStatisticSummariesInput) *XrayGetSamplingStatisticSummariesResult

    GetSamplingTargets(ctx workflow.Context, input *xray.GetSamplingTargetsInput) (*xray.GetSamplingTargetsOutput, error)
    GetSamplingTargetsAsync(ctx workflow.Context, input *xray.GetSamplingTargetsInput) *XrayGetSamplingTargetsResult

    GetServiceGraph(ctx workflow.Context, input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error)
    GetServiceGraphAsync(ctx workflow.Context, input *xray.GetServiceGraphInput) *XrayGetServiceGraphResult

    GetTimeSeriesServiceStatistics(ctx workflow.Context, input *xray.GetTimeSeriesServiceStatisticsInput) (*xray.GetTimeSeriesServiceStatisticsOutput, error)
    GetTimeSeriesServiceStatisticsAsync(ctx workflow.Context, input *xray.GetTimeSeriesServiceStatisticsInput) *XrayGetTimeSeriesServiceStatisticsResult

    GetTraceGraph(ctx workflow.Context, input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error)
    GetTraceGraphAsync(ctx workflow.Context, input *xray.GetTraceGraphInput) *XrayGetTraceGraphResult

    GetTraceSummaries(ctx workflow.Context, input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error)
    GetTraceSummariesAsync(ctx workflow.Context, input *xray.GetTraceSummariesInput) *XrayGetTraceSummariesResult

    ListTagsForResource(ctx workflow.Context, input *xray.ListTagsForResourceInput) (*xray.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *xray.ListTagsForResourceInput) *XrayListTagsForResourceResult

    PutEncryptionConfig(ctx workflow.Context, input *xray.PutEncryptionConfigInput) (*xray.PutEncryptionConfigOutput, error)
    PutEncryptionConfigAsync(ctx workflow.Context, input *xray.PutEncryptionConfigInput) *XrayPutEncryptionConfigResult

    PutTelemetryRecords(ctx workflow.Context, input *xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error)
    PutTelemetryRecordsAsync(ctx workflow.Context, input *xray.PutTelemetryRecordsInput) *XrayPutTelemetryRecordsResult

    PutTraceSegments(ctx workflow.Context, input *xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error)
    PutTraceSegmentsAsync(ctx workflow.Context, input *xray.PutTraceSegmentsInput) *XrayPutTraceSegmentsResult

    TagResource(ctx workflow.Context, input *xray.TagResourceInput) (*xray.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *xray.TagResourceInput) *XrayTagResourceResult

    UntagResource(ctx workflow.Context, input *xray.UntagResourceInput) (*xray.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *xray.UntagResourceInput) *XrayUntagResourceResult

    UpdateGroup(ctx workflow.Context, input *xray.UpdateGroupInput) (*xray.UpdateGroupOutput, error)
    UpdateGroupAsync(ctx workflow.Context, input *xray.UpdateGroupInput) *XrayUpdateGroupResult

    UpdateSamplingRule(ctx workflow.Context, input *xray.UpdateSamplingRuleInput) (*xray.UpdateSamplingRuleOutput, error)
    UpdateSamplingRuleAsync(ctx workflow.Context, input *xray.UpdateSamplingRuleInput) *XrayUpdateSamplingRuleResult
}
type XrayBatchGetTracesResult struct {
	Result workflow.Future
}

func (r *XrayBatchGetTracesResult) Get(ctx workflow.Context) (*xray.BatchGetTracesOutput, error) {
    var output xray.BatchGetTracesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayCreateGroupResult struct {
	Result workflow.Future
}

func (r *XrayCreateGroupResult) Get(ctx workflow.Context) (*xray.CreateGroupOutput, error) {
    var output xray.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayCreateSamplingRuleResult struct {
	Result workflow.Future
}

func (r *XrayCreateSamplingRuleResult) Get(ctx workflow.Context) (*xray.CreateSamplingRuleOutput, error) {
    var output xray.CreateSamplingRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayDeleteGroupResult struct {
	Result workflow.Future
}

func (r *XrayDeleteGroupResult) Get(ctx workflow.Context) (*xray.DeleteGroupOutput, error) {
    var output xray.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayDeleteSamplingRuleResult struct {
	Result workflow.Future
}

func (r *XrayDeleteSamplingRuleResult) Get(ctx workflow.Context) (*xray.DeleteSamplingRuleOutput, error) {
    var output xray.DeleteSamplingRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetEncryptionConfigResult struct {
	Result workflow.Future
}

func (r *XrayGetEncryptionConfigResult) Get(ctx workflow.Context) (*xray.GetEncryptionConfigOutput, error) {
    var output xray.GetEncryptionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetGroupResult struct {
	Result workflow.Future
}

func (r *XrayGetGroupResult) Get(ctx workflow.Context) (*xray.GetGroupOutput, error) {
    var output xray.GetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetGroupsResult struct {
	Result workflow.Future
}

func (r *XrayGetGroupsResult) Get(ctx workflow.Context) (*xray.GetGroupsOutput, error) {
    var output xray.GetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetSamplingRulesResult struct {
	Result workflow.Future
}

func (r *XrayGetSamplingRulesResult) Get(ctx workflow.Context) (*xray.GetSamplingRulesOutput, error) {
    var output xray.GetSamplingRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetSamplingStatisticSummariesResult struct {
	Result workflow.Future
}

func (r *XrayGetSamplingStatisticSummariesResult) Get(ctx workflow.Context) (*xray.GetSamplingStatisticSummariesOutput, error) {
    var output xray.GetSamplingStatisticSummariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetSamplingTargetsResult struct {
	Result workflow.Future
}

func (r *XrayGetSamplingTargetsResult) Get(ctx workflow.Context) (*xray.GetSamplingTargetsOutput, error) {
    var output xray.GetSamplingTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetServiceGraphResult struct {
	Result workflow.Future
}

func (r *XrayGetServiceGraphResult) Get(ctx workflow.Context) (*xray.GetServiceGraphOutput, error) {
    var output xray.GetServiceGraphOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetTimeSeriesServiceStatisticsResult struct {
	Result workflow.Future
}

func (r *XrayGetTimeSeriesServiceStatisticsResult) Get(ctx workflow.Context) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
    var output xray.GetTimeSeriesServiceStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetTraceGraphResult struct {
	Result workflow.Future
}

func (r *XrayGetTraceGraphResult) Get(ctx workflow.Context) (*xray.GetTraceGraphOutput, error) {
    var output xray.GetTraceGraphOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayGetTraceSummariesResult struct {
	Result workflow.Future
}

func (r *XrayGetTraceSummariesResult) Get(ctx workflow.Context) (*xray.GetTraceSummariesOutput, error) {
    var output xray.GetTraceSummariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *XrayListTagsForResourceResult) Get(ctx workflow.Context) (*xray.ListTagsForResourceOutput, error) {
    var output xray.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayPutEncryptionConfigResult struct {
	Result workflow.Future
}

func (r *XrayPutEncryptionConfigResult) Get(ctx workflow.Context) (*xray.PutEncryptionConfigOutput, error) {
    var output xray.PutEncryptionConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayPutTelemetryRecordsResult struct {
	Result workflow.Future
}

func (r *XrayPutTelemetryRecordsResult) Get(ctx workflow.Context) (*xray.PutTelemetryRecordsOutput, error) {
    var output xray.PutTelemetryRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayPutTraceSegmentsResult struct {
	Result workflow.Future
}

func (r *XrayPutTraceSegmentsResult) Get(ctx workflow.Context) (*xray.PutTraceSegmentsOutput, error) {
    var output xray.PutTraceSegmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayTagResourceResult struct {
	Result workflow.Future
}

func (r *XrayTagResourceResult) Get(ctx workflow.Context) (*xray.TagResourceOutput, error) {
    var output xray.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayUntagResourceResult struct {
	Result workflow.Future
}

func (r *XrayUntagResourceResult) Get(ctx workflow.Context) (*xray.UntagResourceOutput, error) {
    var output xray.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayUpdateGroupResult struct {
	Result workflow.Future
}

func (r *XrayUpdateGroupResult) Get(ctx workflow.Context) (*xray.UpdateGroupOutput, error) {
    var output xray.UpdateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type XrayUpdateSamplingRuleResult struct {
	Result workflow.Future
}

func (r *XrayUpdateSamplingRuleResult) Get(ctx workflow.Context) (*xray.UpdateSamplingRuleOutput, error) {
    var output xray.UpdateSamplingRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type XRayStub struct {
    activities awsactivities.XRayActivities
}

func NewXRayStub() XRayClient {
    return &XRayStub{}
}
func (a *XRayStub) BatchGetTraces(ctx workflow.Context, input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
    var output xray.BatchGetTracesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetTraces, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) BatchGetTracesAsync(ctx workflow.Context, input *xray.BatchGetTracesInput) *XrayBatchGetTracesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetTraces, input)
    return &XrayBatchGetTracesResult{Result: future}
}
func (a *XRayStub) CreateGroup(ctx workflow.Context, input *xray.CreateGroupInput) (*xray.CreateGroupOutput, error) {
    var output xray.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) CreateGroupAsync(ctx workflow.Context, input *xray.CreateGroupInput) *XrayCreateGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input)
    return &XrayCreateGroupResult{Result: future}
}
func (a *XRayStub) CreateSamplingRule(ctx workflow.Context, input *xray.CreateSamplingRuleInput) (*xray.CreateSamplingRuleOutput, error) {
    var output xray.CreateSamplingRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSamplingRule, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) CreateSamplingRuleAsync(ctx workflow.Context, input *xray.CreateSamplingRuleInput) *XrayCreateSamplingRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSamplingRule, input)
    return &XrayCreateSamplingRuleResult{Result: future}
}
func (a *XRayStub) DeleteGroup(ctx workflow.Context, input *xray.DeleteGroupInput) (*xray.DeleteGroupOutput, error) {
    var output xray.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) DeleteGroupAsync(ctx workflow.Context, input *xray.DeleteGroupInput) *XrayDeleteGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input)
    return &XrayDeleteGroupResult{Result: future}
}
func (a *XRayStub) DeleteSamplingRule(ctx workflow.Context, input *xray.DeleteSamplingRuleInput) (*xray.DeleteSamplingRuleOutput, error) {
    var output xray.DeleteSamplingRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSamplingRule, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) DeleteSamplingRuleAsync(ctx workflow.Context, input *xray.DeleteSamplingRuleInput) *XrayDeleteSamplingRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSamplingRule, input)
    return &XrayDeleteSamplingRuleResult{Result: future}
}
func (a *XRayStub) GetEncryptionConfig(ctx workflow.Context, input *xray.GetEncryptionConfigInput) (*xray.GetEncryptionConfigOutput, error) {
    var output xray.GetEncryptionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEncryptionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetEncryptionConfigAsync(ctx workflow.Context, input *xray.GetEncryptionConfigInput) *XrayGetEncryptionConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEncryptionConfig, input)
    return &XrayGetEncryptionConfigResult{Result: future}
}
func (a *XRayStub) GetGroup(ctx workflow.Context, input *xray.GetGroupInput) (*xray.GetGroupOutput, error) {
    var output xray.GetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetGroupAsync(ctx workflow.Context, input *xray.GetGroupInput) *XrayGetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGroup, input)
    return &XrayGetGroupResult{Result: future}
}
func (a *XRayStub) GetGroups(ctx workflow.Context, input *xray.GetGroupsInput) (*xray.GetGroupsOutput, error) {
    var output xray.GetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetGroupsAsync(ctx workflow.Context, input *xray.GetGroupsInput) *XrayGetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGroups, input)
    return &XrayGetGroupsResult{Result: future}
}
func (a *XRayStub) GetSamplingRules(ctx workflow.Context, input *xray.GetSamplingRulesInput) (*xray.GetSamplingRulesOutput, error) {
    var output xray.GetSamplingRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSamplingRules, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetSamplingRulesAsync(ctx workflow.Context, input *xray.GetSamplingRulesInput) *XrayGetSamplingRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSamplingRules, input)
    return &XrayGetSamplingRulesResult{Result: future}
}
func (a *XRayStub) GetSamplingStatisticSummaries(ctx workflow.Context, input *xray.GetSamplingStatisticSummariesInput) (*xray.GetSamplingStatisticSummariesOutput, error) {
    var output xray.GetSamplingStatisticSummariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSamplingStatisticSummaries, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetSamplingStatisticSummariesAsync(ctx workflow.Context, input *xray.GetSamplingStatisticSummariesInput) *XrayGetSamplingStatisticSummariesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSamplingStatisticSummaries, input)
    return &XrayGetSamplingStatisticSummariesResult{Result: future}
}
func (a *XRayStub) GetSamplingTargets(ctx workflow.Context, input *xray.GetSamplingTargetsInput) (*xray.GetSamplingTargetsOutput, error) {
    var output xray.GetSamplingTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSamplingTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetSamplingTargetsAsync(ctx workflow.Context, input *xray.GetSamplingTargetsInput) *XrayGetSamplingTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSamplingTargets, input)
    return &XrayGetSamplingTargetsResult{Result: future}
}
func (a *XRayStub) GetServiceGraph(ctx workflow.Context, input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error) {
    var output xray.GetServiceGraphOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceGraph, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetServiceGraphAsync(ctx workflow.Context, input *xray.GetServiceGraphInput) *XrayGetServiceGraphResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetServiceGraph, input)
    return &XrayGetServiceGraphResult{Result: future}
}
func (a *XRayStub) GetTimeSeriesServiceStatistics(ctx workflow.Context, input *xray.GetTimeSeriesServiceStatisticsInput) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
    var output xray.GetTimeSeriesServiceStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTimeSeriesServiceStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetTimeSeriesServiceStatisticsAsync(ctx workflow.Context, input *xray.GetTimeSeriesServiceStatisticsInput) *XrayGetTimeSeriesServiceStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTimeSeriesServiceStatistics, input)
    return &XrayGetTimeSeriesServiceStatisticsResult{Result: future}
}
func (a *XRayStub) GetTraceGraph(ctx workflow.Context, input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error) {
    var output xray.GetTraceGraphOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTraceGraph, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetTraceGraphAsync(ctx workflow.Context, input *xray.GetTraceGraphInput) *XrayGetTraceGraphResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTraceGraph, input)
    return &XrayGetTraceGraphResult{Result: future}
}
func (a *XRayStub) GetTraceSummaries(ctx workflow.Context, input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error) {
    var output xray.GetTraceSummariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTraceSummaries, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) GetTraceSummariesAsync(ctx workflow.Context, input *xray.GetTraceSummariesInput) *XrayGetTraceSummariesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTraceSummaries, input)
    return &XrayGetTraceSummariesResult{Result: future}
}
func (a *XRayStub) ListTagsForResource(ctx workflow.Context, input *xray.ListTagsForResourceInput) (*xray.ListTagsForResourceOutput, error) {
    var output xray.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) ListTagsForResourceAsync(ctx workflow.Context, input *xray.ListTagsForResourceInput) *XrayListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &XrayListTagsForResourceResult{Result: future}
}
func (a *XRayStub) PutEncryptionConfig(ctx workflow.Context, input *xray.PutEncryptionConfigInput) (*xray.PutEncryptionConfigOutput, error) {
    var output xray.PutEncryptionConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEncryptionConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) PutEncryptionConfigAsync(ctx workflow.Context, input *xray.PutEncryptionConfigInput) *XrayPutEncryptionConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEncryptionConfig, input)
    return &XrayPutEncryptionConfigResult{Result: future}
}
func (a *XRayStub) PutTelemetryRecords(ctx workflow.Context, input *xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error) {
    var output xray.PutTelemetryRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutTelemetryRecords, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) PutTelemetryRecordsAsync(ctx workflow.Context, input *xray.PutTelemetryRecordsInput) *XrayPutTelemetryRecordsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutTelemetryRecords, input)
    return &XrayPutTelemetryRecordsResult{Result: future}
}
func (a *XRayStub) PutTraceSegments(ctx workflow.Context, input *xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error) {
    var output xray.PutTraceSegmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutTraceSegments, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) PutTraceSegmentsAsync(ctx workflow.Context, input *xray.PutTraceSegmentsInput) *XrayPutTraceSegmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutTraceSegments, input)
    return &XrayPutTraceSegmentsResult{Result: future}
}
func (a *XRayStub) TagResource(ctx workflow.Context, input *xray.TagResourceInput) (*xray.TagResourceOutput, error) {
    var output xray.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) TagResourceAsync(ctx workflow.Context, input *xray.TagResourceInput) *XrayTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &XrayTagResourceResult{Result: future}
}
func (a *XRayStub) UntagResource(ctx workflow.Context, input *xray.UntagResourceInput) (*xray.UntagResourceOutput, error) {
    var output xray.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) UntagResourceAsync(ctx workflow.Context, input *xray.UntagResourceInput) *XrayUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &XrayUntagResourceResult{Result: future}
}
func (a *XRayStub) UpdateGroup(ctx workflow.Context, input *xray.UpdateGroupInput) (*xray.UpdateGroupOutput, error) {
    var output xray.UpdateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) UpdateGroupAsync(ctx workflow.Context, input *xray.UpdateGroupInput) *XrayUpdateGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input)
    return &XrayUpdateGroupResult{Result: future}
}
func (a *XRayStub) UpdateSamplingRule(ctx workflow.Context, input *xray.UpdateSamplingRuleInput) (*xray.UpdateSamplingRuleOutput, error) {
    var output xray.UpdateSamplingRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSamplingRule, input).Get(ctx, &output)
    return &output, err
}

func (a *XRayStub) UpdateSamplingRuleAsync(ctx workflow.Context, input *xray.UpdateSamplingRuleInput) *XrayUpdateSamplingRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateSamplingRule, input)
    return &XrayUpdateSamplingRuleResult{Result: future}
}
