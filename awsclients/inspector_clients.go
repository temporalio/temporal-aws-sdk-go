package awsclients

import (
	"github.com/aws/aws-sdk-go/service/inspector"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type InspectorClient interface {
       AddAttributesToFindings(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error)
       AddAttributesToFindingsAsync(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) *InspectorAddAttributesToFindingsResult

       CreateAssessmentTarget(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error)
       CreateAssessmentTargetAsync(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) *InspectorCreateAssessmentTargetResult

       CreateAssessmentTemplate(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error)
       CreateAssessmentTemplateAsync(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) *InspectorCreateAssessmentTemplateResult

       CreateExclusionsPreview(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error)
       CreateExclusionsPreviewAsync(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) *InspectorCreateExclusionsPreviewResult

       CreateResourceGroup(ctx workflow.Context, input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error)
       CreateResourceGroupAsync(ctx workflow.Context, input *inspector.CreateResourceGroupInput) *InspectorCreateResourceGroupResult

       DeleteAssessmentRun(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error)
       DeleteAssessmentRunAsync(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) *InspectorDeleteAssessmentRunResult

       DeleteAssessmentTarget(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error)
       DeleteAssessmentTargetAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) *InspectorDeleteAssessmentTargetResult

       DeleteAssessmentTemplate(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error)
       DeleteAssessmentTemplateAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) *InspectorDeleteAssessmentTemplateResult

       DescribeAssessmentRuns(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error)
       DescribeAssessmentRunsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) *InspectorDescribeAssessmentRunsResult

       DescribeAssessmentTargets(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error)
       DescribeAssessmentTargetsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) *InspectorDescribeAssessmentTargetsResult

       DescribeAssessmentTemplates(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error)
       DescribeAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) *InspectorDescribeAssessmentTemplatesResult

       DescribeCrossAccountAccessRole(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error)
       DescribeCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) *InspectorDescribeCrossAccountAccessRoleResult

       DescribeExclusions(ctx workflow.Context, input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error)
       DescribeExclusionsAsync(ctx workflow.Context, input *inspector.DescribeExclusionsInput) *InspectorDescribeExclusionsResult

       DescribeFindings(ctx workflow.Context, input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error)
       DescribeFindingsAsync(ctx workflow.Context, input *inspector.DescribeFindingsInput) *InspectorDescribeFindingsResult

       DescribeResourceGroups(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error)
       DescribeResourceGroupsAsync(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) *InspectorDescribeResourceGroupsResult

       DescribeRulesPackages(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error)
       DescribeRulesPackagesAsync(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) *InspectorDescribeRulesPackagesResult

       GetAssessmentReport(ctx workflow.Context, input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error)
       GetAssessmentReportAsync(ctx workflow.Context, input *inspector.GetAssessmentReportInput) *InspectorGetAssessmentReportResult

       GetExclusionsPreview(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error)
       GetExclusionsPreviewAsync(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) *InspectorGetExclusionsPreviewResult

       GetTelemetryMetadata(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error)
       GetTelemetryMetadataAsync(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) *InspectorGetTelemetryMetadataResult

       ListAssessmentRunAgents(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error)
       ListAssessmentRunAgentsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) *InspectorListAssessmentRunAgentsResult

       ListAssessmentRuns(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error)
       ListAssessmentRunsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) *InspectorListAssessmentRunsResult

       ListAssessmentTargets(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error)
       ListAssessmentTargetsAsync(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) *InspectorListAssessmentTargetsResult

       ListAssessmentTemplates(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error)
       ListAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) *InspectorListAssessmentTemplatesResult

       ListEventSubscriptions(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error)
       ListEventSubscriptionsAsync(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) *InspectorListEventSubscriptionsResult

       ListExclusions(ctx workflow.Context, input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error)
       ListExclusionsAsync(ctx workflow.Context, input *inspector.ListExclusionsInput) *InspectorListExclusionsResult

       ListFindings(ctx workflow.Context, input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error)
       ListFindingsAsync(ctx workflow.Context, input *inspector.ListFindingsInput) *InspectorListFindingsResult

       ListRulesPackages(ctx workflow.Context, input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error)
       ListRulesPackagesAsync(ctx workflow.Context, input *inspector.ListRulesPackagesInput) *InspectorListRulesPackagesResult

       ListTagsForResource(ctx workflow.Context, input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *inspector.ListTagsForResourceInput) *InspectorListTagsForResourceResult

       PreviewAgents(ctx workflow.Context, input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error)
       PreviewAgentsAsync(ctx workflow.Context, input *inspector.PreviewAgentsInput) *InspectorPreviewAgentsResult

       RegisterCrossAccountAccessRole(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error)
       RegisterCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) *InspectorRegisterCrossAccountAccessRoleResult

       RemoveAttributesFromFindings(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error)
       RemoveAttributesFromFindingsAsync(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) *InspectorRemoveAttributesFromFindingsResult

       SetTagsForResource(ctx workflow.Context, input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error)
       SetTagsForResourceAsync(ctx workflow.Context, input *inspector.SetTagsForResourceInput) *InspectorSetTagsForResourceResult

       StartAssessmentRun(ctx workflow.Context, input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error)
       StartAssessmentRunAsync(ctx workflow.Context, input *inspector.StartAssessmentRunInput) *InspectorStartAssessmentRunResult

       StopAssessmentRun(ctx workflow.Context, input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error)
       StopAssessmentRunAsync(ctx workflow.Context, input *inspector.StopAssessmentRunInput) *InspectorStopAssessmentRunResult

       SubscribeToEvent(ctx workflow.Context, input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error)
       SubscribeToEventAsync(ctx workflow.Context, input *inspector.SubscribeToEventInput) *InspectorSubscribeToEventResult

       UnsubscribeFromEvent(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error)
       UnsubscribeFromEventAsync(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) *InspectorUnsubscribeFromEventResult

       UpdateAssessmentTarget(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error)
       UpdateAssessmentTargetAsync(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) *InspectorUpdateAssessmentTargetResult
}

type InspectorAddAttributesToFindingsResult struct {
	Result workflow.Future
}

func (r *InspectorAddAttributesToFindingsResult) Get(ctx workflow.Context) (*inspector.AddAttributesToFindingsOutput, error) {
    var output inspector.AddAttributesToFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorCreateAssessmentTargetResult struct {
	Result workflow.Future
}

func (r *InspectorCreateAssessmentTargetResult) Get(ctx workflow.Context) (*inspector.CreateAssessmentTargetOutput, error) {
    var output inspector.CreateAssessmentTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorCreateAssessmentTemplateResult struct {
	Result workflow.Future
}

func (r *InspectorCreateAssessmentTemplateResult) Get(ctx workflow.Context) (*inspector.CreateAssessmentTemplateOutput, error) {
    var output inspector.CreateAssessmentTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorCreateExclusionsPreviewResult struct {
	Result workflow.Future
}

func (r *InspectorCreateExclusionsPreviewResult) Get(ctx workflow.Context) (*inspector.CreateExclusionsPreviewOutput, error) {
    var output inspector.CreateExclusionsPreviewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorCreateResourceGroupResult struct {
	Result workflow.Future
}

func (r *InspectorCreateResourceGroupResult) Get(ctx workflow.Context) (*inspector.CreateResourceGroupOutput, error) {
    var output inspector.CreateResourceGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDeleteAssessmentRunResult struct {
	Result workflow.Future
}

func (r *InspectorDeleteAssessmentRunResult) Get(ctx workflow.Context) (*inspector.DeleteAssessmentRunOutput, error) {
    var output inspector.DeleteAssessmentRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDeleteAssessmentTargetResult struct {
	Result workflow.Future
}

func (r *InspectorDeleteAssessmentTargetResult) Get(ctx workflow.Context) (*inspector.DeleteAssessmentTargetOutput, error) {
    var output inspector.DeleteAssessmentTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDeleteAssessmentTemplateResult struct {
	Result workflow.Future
}

func (r *InspectorDeleteAssessmentTemplateResult) Get(ctx workflow.Context) (*inspector.DeleteAssessmentTemplateOutput, error) {
    var output inspector.DeleteAssessmentTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeAssessmentRunsResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeAssessmentRunsResult) Get(ctx workflow.Context) (*inspector.DescribeAssessmentRunsOutput, error) {
    var output inspector.DescribeAssessmentRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeAssessmentTargetsResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeAssessmentTargetsResult) Get(ctx workflow.Context) (*inspector.DescribeAssessmentTargetsOutput, error) {
    var output inspector.DescribeAssessmentTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeAssessmentTemplatesResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeAssessmentTemplatesResult) Get(ctx workflow.Context) (*inspector.DescribeAssessmentTemplatesOutput, error) {
    var output inspector.DescribeAssessmentTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeCrossAccountAccessRoleResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeCrossAccountAccessRoleResult) Get(ctx workflow.Context) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
    var output inspector.DescribeCrossAccountAccessRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeExclusionsResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeExclusionsResult) Get(ctx workflow.Context) (*inspector.DescribeExclusionsOutput, error) {
    var output inspector.DescribeExclusionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeFindingsResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeFindingsResult) Get(ctx workflow.Context) (*inspector.DescribeFindingsOutput, error) {
    var output inspector.DescribeFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeResourceGroupsResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeResourceGroupsResult) Get(ctx workflow.Context) (*inspector.DescribeResourceGroupsOutput, error) {
    var output inspector.DescribeResourceGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorDescribeRulesPackagesResult struct {
	Result workflow.Future
}

func (r *InspectorDescribeRulesPackagesResult) Get(ctx workflow.Context) (*inspector.DescribeRulesPackagesOutput, error) {
    var output inspector.DescribeRulesPackagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorGetAssessmentReportResult struct {
	Result workflow.Future
}

func (r *InspectorGetAssessmentReportResult) Get(ctx workflow.Context) (*inspector.GetAssessmentReportOutput, error) {
    var output inspector.GetAssessmentReportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorGetExclusionsPreviewResult struct {
	Result workflow.Future
}

func (r *InspectorGetExclusionsPreviewResult) Get(ctx workflow.Context) (*inspector.GetExclusionsPreviewOutput, error) {
    var output inspector.GetExclusionsPreviewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorGetTelemetryMetadataResult struct {
	Result workflow.Future
}

func (r *InspectorGetTelemetryMetadataResult) Get(ctx workflow.Context) (*inspector.GetTelemetryMetadataOutput, error) {
    var output inspector.GetTelemetryMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListAssessmentRunAgentsResult struct {
	Result workflow.Future
}

func (r *InspectorListAssessmentRunAgentsResult) Get(ctx workflow.Context) (*inspector.ListAssessmentRunAgentsOutput, error) {
    var output inspector.ListAssessmentRunAgentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListAssessmentRunsResult struct {
	Result workflow.Future
}

func (r *InspectorListAssessmentRunsResult) Get(ctx workflow.Context) (*inspector.ListAssessmentRunsOutput, error) {
    var output inspector.ListAssessmentRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListAssessmentTargetsResult struct {
	Result workflow.Future
}

func (r *InspectorListAssessmentTargetsResult) Get(ctx workflow.Context) (*inspector.ListAssessmentTargetsOutput, error) {
    var output inspector.ListAssessmentTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListAssessmentTemplatesResult struct {
	Result workflow.Future
}

func (r *InspectorListAssessmentTemplatesResult) Get(ctx workflow.Context) (*inspector.ListAssessmentTemplatesOutput, error) {
    var output inspector.ListAssessmentTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListEventSubscriptionsResult struct {
	Result workflow.Future
}

func (r *InspectorListEventSubscriptionsResult) Get(ctx workflow.Context) (*inspector.ListEventSubscriptionsOutput, error) {
    var output inspector.ListEventSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListExclusionsResult struct {
	Result workflow.Future
}

func (r *InspectorListExclusionsResult) Get(ctx workflow.Context) (*inspector.ListExclusionsOutput, error) {
    var output inspector.ListExclusionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListFindingsResult struct {
	Result workflow.Future
}

func (r *InspectorListFindingsResult) Get(ctx workflow.Context) (*inspector.ListFindingsOutput, error) {
    var output inspector.ListFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListRulesPackagesResult struct {
	Result workflow.Future
}

func (r *InspectorListRulesPackagesResult) Get(ctx workflow.Context) (*inspector.ListRulesPackagesOutput, error) {
    var output inspector.ListRulesPackagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *InspectorListTagsForResourceResult) Get(ctx workflow.Context) (*inspector.ListTagsForResourceOutput, error) {
    var output inspector.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorPreviewAgentsResult struct {
	Result workflow.Future
}

func (r *InspectorPreviewAgentsResult) Get(ctx workflow.Context) (*inspector.PreviewAgentsOutput, error) {
    var output inspector.PreviewAgentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorRegisterCrossAccountAccessRoleResult struct {
	Result workflow.Future
}

func (r *InspectorRegisterCrossAccountAccessRoleResult) Get(ctx workflow.Context) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
    var output inspector.RegisterCrossAccountAccessRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorRemoveAttributesFromFindingsResult struct {
	Result workflow.Future
}

func (r *InspectorRemoveAttributesFromFindingsResult) Get(ctx workflow.Context) (*inspector.RemoveAttributesFromFindingsOutput, error) {
    var output inspector.RemoveAttributesFromFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorSetTagsForResourceResult struct {
	Result workflow.Future
}

func (r *InspectorSetTagsForResourceResult) Get(ctx workflow.Context) (*inspector.SetTagsForResourceOutput, error) {
    var output inspector.SetTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorStartAssessmentRunResult struct {
	Result workflow.Future
}

func (r *InspectorStartAssessmentRunResult) Get(ctx workflow.Context) (*inspector.StartAssessmentRunOutput, error) {
    var output inspector.StartAssessmentRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorStopAssessmentRunResult struct {
	Result workflow.Future
}

func (r *InspectorStopAssessmentRunResult) Get(ctx workflow.Context) (*inspector.StopAssessmentRunOutput, error) {
    var output inspector.StopAssessmentRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorSubscribeToEventResult struct {
	Result workflow.Future
}

func (r *InspectorSubscribeToEventResult) Get(ctx workflow.Context) (*inspector.SubscribeToEventOutput, error) {
    var output inspector.SubscribeToEventOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorUnsubscribeFromEventResult struct {
	Result workflow.Future
}

func (r *InspectorUnsubscribeFromEventResult) Get(ctx workflow.Context) (*inspector.UnsubscribeFromEventOutput, error) {
    var output inspector.UnsubscribeFromEventOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorUpdateAssessmentTargetResult struct {
	Result workflow.Future
}

func (r *InspectorUpdateAssessmentTargetResult) Get(ctx workflow.Context) (*inspector.UpdateAssessmentTargetOutput, error) {
    var output inspector.UpdateAssessmentTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type InspectorStub struct {
    activities awsactivities.InspectorActivities
}

func NewInspectorStub() InspectorClient {
    return &InspectorStub{}
}

func (a *InspectorStub) AddAttributesToFindings(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
    var output inspector.AddAttributesToFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddAttributesToFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) AddAttributesToFindingsAsync(ctx workflow.Context, input *inspector.AddAttributesToFindingsInput) *InspectorAddAttributesToFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddAttributesToFindings, input)
    return &InspectorAddAttributesToFindingsResult{Result: future}
}

func (a *InspectorStub) CreateAssessmentTarget(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error) {
    var output inspector.CreateAssessmentTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAssessmentTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) CreateAssessmentTargetAsync(ctx workflow.Context, input *inspector.CreateAssessmentTargetInput) *InspectorCreateAssessmentTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAssessmentTarget, input)
    return &InspectorCreateAssessmentTargetResult{Result: future}
}

func (a *InspectorStub) CreateAssessmentTemplate(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error) {
    var output inspector.CreateAssessmentTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAssessmentTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) CreateAssessmentTemplateAsync(ctx workflow.Context, input *inspector.CreateAssessmentTemplateInput) *InspectorCreateAssessmentTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAssessmentTemplate, input)
    return &InspectorCreateAssessmentTemplateResult{Result: future}
}

func (a *InspectorStub) CreateExclusionsPreview(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error) {
    var output inspector.CreateExclusionsPreviewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateExclusionsPreview, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) CreateExclusionsPreviewAsync(ctx workflow.Context, input *inspector.CreateExclusionsPreviewInput) *InspectorCreateExclusionsPreviewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateExclusionsPreview, input)
    return &InspectorCreateExclusionsPreviewResult{Result: future}
}

func (a *InspectorStub) CreateResourceGroup(ctx workflow.Context, input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error) {
    var output inspector.CreateResourceGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResourceGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) CreateResourceGroupAsync(ctx workflow.Context, input *inspector.CreateResourceGroupInput) *InspectorCreateResourceGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateResourceGroup, input)
    return &InspectorCreateResourceGroupResult{Result: future}
}

func (a *InspectorStub) DeleteAssessmentRun(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error) {
    var output inspector.DeleteAssessmentRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAssessmentRun, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DeleteAssessmentRunAsync(ctx workflow.Context, input *inspector.DeleteAssessmentRunInput) *InspectorDeleteAssessmentRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAssessmentRun, input)
    return &InspectorDeleteAssessmentRunResult{Result: future}
}

func (a *InspectorStub) DeleteAssessmentTarget(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error) {
    var output inspector.DeleteAssessmentTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAssessmentTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DeleteAssessmentTargetAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTargetInput) *InspectorDeleteAssessmentTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAssessmentTarget, input)
    return &InspectorDeleteAssessmentTargetResult{Result: future}
}

func (a *InspectorStub) DeleteAssessmentTemplate(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error) {
    var output inspector.DeleteAssessmentTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAssessmentTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DeleteAssessmentTemplateAsync(ctx workflow.Context, input *inspector.DeleteAssessmentTemplateInput) *InspectorDeleteAssessmentTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAssessmentTemplate, input)
    return &InspectorDeleteAssessmentTemplateResult{Result: future}
}

func (a *InspectorStub) DescribeAssessmentRuns(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error) {
    var output inspector.DescribeAssessmentRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssessmentRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeAssessmentRunsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentRunsInput) *InspectorDescribeAssessmentRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAssessmentRuns, input)
    return &InspectorDescribeAssessmentRunsResult{Result: future}
}

func (a *InspectorStub) DescribeAssessmentTargets(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error) {
    var output inspector.DescribeAssessmentTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssessmentTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeAssessmentTargetsAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTargetsInput) *InspectorDescribeAssessmentTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAssessmentTargets, input)
    return &InspectorDescribeAssessmentTargetsResult{Result: future}
}

func (a *InspectorStub) DescribeAssessmentTemplates(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error) {
    var output inspector.DescribeAssessmentTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssessmentTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.DescribeAssessmentTemplatesInput) *InspectorDescribeAssessmentTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAssessmentTemplates, input)
    return &InspectorDescribeAssessmentTemplatesResult{Result: future}
}

func (a *InspectorStub) DescribeCrossAccountAccessRole(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
    var output inspector.DescribeCrossAccountAccessRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCrossAccountAccessRole, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.DescribeCrossAccountAccessRoleInput) *InspectorDescribeCrossAccountAccessRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCrossAccountAccessRole, input)
    return &InspectorDescribeCrossAccountAccessRoleResult{Result: future}
}

func (a *InspectorStub) DescribeExclusions(ctx workflow.Context, input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error) {
    var output inspector.DescribeExclusionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExclusions, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeExclusionsAsync(ctx workflow.Context, input *inspector.DescribeExclusionsInput) *InspectorDescribeExclusionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeExclusions, input)
    return &InspectorDescribeExclusionsResult{Result: future}
}

func (a *InspectorStub) DescribeFindings(ctx workflow.Context, input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error) {
    var output inspector.DescribeFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeFindingsAsync(ctx workflow.Context, input *inspector.DescribeFindingsInput) *InspectorDescribeFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFindings, input)
    return &InspectorDescribeFindingsResult{Result: future}
}

func (a *InspectorStub) DescribeResourceGroups(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error) {
    var output inspector.DescribeResourceGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeResourceGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeResourceGroupsAsync(ctx workflow.Context, input *inspector.DescribeResourceGroupsInput) *InspectorDescribeResourceGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeResourceGroups, input)
    return &InspectorDescribeResourceGroupsResult{Result: future}
}

func (a *InspectorStub) DescribeRulesPackages(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error) {
    var output inspector.DescribeRulesPackagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRulesPackages, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) DescribeRulesPackagesAsync(ctx workflow.Context, input *inspector.DescribeRulesPackagesInput) *InspectorDescribeRulesPackagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRulesPackages, input)
    return &InspectorDescribeRulesPackagesResult{Result: future}
}

func (a *InspectorStub) GetAssessmentReport(ctx workflow.Context, input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error) {
    var output inspector.GetAssessmentReportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAssessmentReport, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) GetAssessmentReportAsync(ctx workflow.Context, input *inspector.GetAssessmentReportInput) *InspectorGetAssessmentReportResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAssessmentReport, input)
    return &InspectorGetAssessmentReportResult{Result: future}
}

func (a *InspectorStub) GetExclusionsPreview(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error) {
    var output inspector.GetExclusionsPreviewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetExclusionsPreview, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) GetExclusionsPreviewAsync(ctx workflow.Context, input *inspector.GetExclusionsPreviewInput) *InspectorGetExclusionsPreviewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetExclusionsPreview, input)
    return &InspectorGetExclusionsPreviewResult{Result: future}
}

func (a *InspectorStub) GetTelemetryMetadata(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error) {
    var output inspector.GetTelemetryMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTelemetryMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) GetTelemetryMetadataAsync(ctx workflow.Context, input *inspector.GetTelemetryMetadataInput) *InspectorGetTelemetryMetadataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTelemetryMetadata, input)
    return &InspectorGetTelemetryMetadataResult{Result: future}
}

func (a *InspectorStub) ListAssessmentRunAgents(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error) {
    var output inspector.ListAssessmentRunAgentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentRunAgents, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListAssessmentRunAgentsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunAgentsInput) *InspectorListAssessmentRunAgentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentRunAgents, input)
    return &InspectorListAssessmentRunAgentsResult{Result: future}
}

func (a *InspectorStub) ListAssessmentRuns(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error) {
    var output inspector.ListAssessmentRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListAssessmentRunsAsync(ctx workflow.Context, input *inspector.ListAssessmentRunsInput) *InspectorListAssessmentRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentRuns, input)
    return &InspectorListAssessmentRunsResult{Result: future}
}

func (a *InspectorStub) ListAssessmentTargets(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error) {
    var output inspector.ListAssessmentTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListAssessmentTargetsAsync(ctx workflow.Context, input *inspector.ListAssessmentTargetsInput) *InspectorListAssessmentTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentTargets, input)
    return &InspectorListAssessmentTargetsResult{Result: future}
}

func (a *InspectorStub) ListAssessmentTemplates(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error) {
    var output inspector.ListAssessmentTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListAssessmentTemplatesAsync(ctx workflow.Context, input *inspector.ListAssessmentTemplatesInput) *InspectorListAssessmentTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssessmentTemplates, input)
    return &InspectorListAssessmentTemplatesResult{Result: future}
}

func (a *InspectorStub) ListEventSubscriptions(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error) {
    var output inspector.ListEventSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEventSubscriptions, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListEventSubscriptionsAsync(ctx workflow.Context, input *inspector.ListEventSubscriptionsInput) *InspectorListEventSubscriptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEventSubscriptions, input)
    return &InspectorListEventSubscriptionsResult{Result: future}
}

func (a *InspectorStub) ListExclusions(ctx workflow.Context, input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error) {
    var output inspector.ListExclusionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListExclusions, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListExclusionsAsync(ctx workflow.Context, input *inspector.ListExclusionsInput) *InspectorListExclusionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListExclusions, input)
    return &InspectorListExclusionsResult{Result: future}
}

func (a *InspectorStub) ListFindings(ctx workflow.Context, input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error) {
    var output inspector.ListFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListFindingsAsync(ctx workflow.Context, input *inspector.ListFindingsInput) *InspectorListFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input)
    return &InspectorListFindingsResult{Result: future}
}

func (a *InspectorStub) ListRulesPackages(ctx workflow.Context, input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error) {
    var output inspector.ListRulesPackagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRulesPackages, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListRulesPackagesAsync(ctx workflow.Context, input *inspector.ListRulesPackagesInput) *InspectorListRulesPackagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRulesPackages, input)
    return &InspectorListRulesPackagesResult{Result: future}
}

func (a *InspectorStub) ListTagsForResource(ctx workflow.Context, input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error) {
    var output inspector.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) ListTagsForResourceAsync(ctx workflow.Context, input *inspector.ListTagsForResourceInput) *InspectorListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &InspectorListTagsForResourceResult{Result: future}
}

func (a *InspectorStub) PreviewAgents(ctx workflow.Context, input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error) {
    var output inspector.PreviewAgentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PreviewAgents, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) PreviewAgentsAsync(ctx workflow.Context, input *inspector.PreviewAgentsInput) *InspectorPreviewAgentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PreviewAgents, input)
    return &InspectorPreviewAgentsResult{Result: future}
}

func (a *InspectorStub) RegisterCrossAccountAccessRole(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
    var output inspector.RegisterCrossAccountAccessRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterCrossAccountAccessRole, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) RegisterCrossAccountAccessRoleAsync(ctx workflow.Context, input *inspector.RegisterCrossAccountAccessRoleInput) *InspectorRegisterCrossAccountAccessRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterCrossAccountAccessRole, input)
    return &InspectorRegisterCrossAccountAccessRoleResult{Result: future}
}

func (a *InspectorStub) RemoveAttributesFromFindings(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error) {
    var output inspector.RemoveAttributesFromFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveAttributesFromFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) RemoveAttributesFromFindingsAsync(ctx workflow.Context, input *inspector.RemoveAttributesFromFindingsInput) *InspectorRemoveAttributesFromFindingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveAttributesFromFindings, input)
    return &InspectorRemoveAttributesFromFindingsResult{Result: future}
}

func (a *InspectorStub) SetTagsForResource(ctx workflow.Context, input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error) {
    var output inspector.SetTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) SetTagsForResourceAsync(ctx workflow.Context, input *inspector.SetTagsForResourceInput) *InspectorSetTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetTagsForResource, input)
    return &InspectorSetTagsForResourceResult{Result: future}
}

func (a *InspectorStub) StartAssessmentRun(ctx workflow.Context, input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error) {
    var output inspector.StartAssessmentRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartAssessmentRun, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) StartAssessmentRunAsync(ctx workflow.Context, input *inspector.StartAssessmentRunInput) *InspectorStartAssessmentRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartAssessmentRun, input)
    return &InspectorStartAssessmentRunResult{Result: future}
}

func (a *InspectorStub) StopAssessmentRun(ctx workflow.Context, input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error) {
    var output inspector.StopAssessmentRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopAssessmentRun, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) StopAssessmentRunAsync(ctx workflow.Context, input *inspector.StopAssessmentRunInput) *InspectorStopAssessmentRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopAssessmentRun, input)
    return &InspectorStopAssessmentRunResult{Result: future}
}

func (a *InspectorStub) SubscribeToEvent(ctx workflow.Context, input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error) {
    var output inspector.SubscribeToEventOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubscribeToEvent, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) SubscribeToEventAsync(ctx workflow.Context, input *inspector.SubscribeToEventInput) *InspectorSubscribeToEventResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubscribeToEvent, input)
    return &InspectorSubscribeToEventResult{Result: future}
}

func (a *InspectorStub) UnsubscribeFromEvent(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error) {
    var output inspector.UnsubscribeFromEventOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnsubscribeFromEvent, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) UnsubscribeFromEventAsync(ctx workflow.Context, input *inspector.UnsubscribeFromEventInput) *InspectorUnsubscribeFromEventResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnsubscribeFromEvent, input)
    return &InspectorUnsubscribeFromEventResult{Result: future}
}

func (a *InspectorStub) UpdateAssessmentTarget(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error) {
    var output inspector.UpdateAssessmentTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAssessmentTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *InspectorStub) UpdateAssessmentTargetAsync(ctx workflow.Context, input *inspector.UpdateAssessmentTargetInput) *InspectorUpdateAssessmentTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAssessmentTarget, input)
    return &InspectorUpdateAssessmentTargetResult{Result: future}
}