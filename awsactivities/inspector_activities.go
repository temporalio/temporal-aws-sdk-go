
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/inspector/inspectoriface"
)

type InspectorActivities struct {
	client inspectoriface.InspectorAPI
}

func NewInspectorActivities(client inspectoriface.InspectorAPI) *InspectorActivities {
    return &InspectorActivities{client: client}
}

func (a *InspectorActivities) AddAttributesToFindings(input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
    return a.client.AddAttributesToFindings(input)
}

func (a *InspectorActivities) CreateAssessmentTarget(input *inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error) {
    return a.client.CreateAssessmentTarget(input)
}

func (a *InspectorActivities) CreateAssessmentTemplate(input *inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error) {
    return a.client.CreateAssessmentTemplate(input)
}

func (a *InspectorActivities) CreateExclusionsPreview(input *inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error) {
    return a.client.CreateExclusionsPreview(input)
}

func (a *InspectorActivities) CreateResourceGroup(input *inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error) {
    return a.client.CreateResourceGroup(input)
}

func (a *InspectorActivities) DeleteAssessmentRun(input *inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error) {
    return a.client.DeleteAssessmentRun(input)
}

func (a *InspectorActivities) DeleteAssessmentTarget(input *inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error) {
    return a.client.DeleteAssessmentTarget(input)
}

func (a *InspectorActivities) DeleteAssessmentTemplate(input *inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error) {
    return a.client.DeleteAssessmentTemplate(input)
}

func (a *InspectorActivities) DescribeAssessmentRuns(input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error) {
    return a.client.DescribeAssessmentRuns(input)
}

func (a *InspectorActivities) DescribeAssessmentTargets(input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error) {
    return a.client.DescribeAssessmentTargets(input)
}

func (a *InspectorActivities) DescribeAssessmentTemplates(input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error) {
    return a.client.DescribeAssessmentTemplates(input)
}

func (a *InspectorActivities) DescribeCrossAccountAccessRole(input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
    return a.client.DescribeCrossAccountAccessRole(input)
}

func (a *InspectorActivities) DescribeExclusions(input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error) {
    return a.client.DescribeExclusions(input)
}

func (a *InspectorActivities) DescribeFindings(input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error) {
    return a.client.DescribeFindings(input)
}

func (a *InspectorActivities) DescribeResourceGroups(input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error) {
    return a.client.DescribeResourceGroups(input)
}

func (a *InspectorActivities) DescribeRulesPackages(input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error) {
    return a.client.DescribeRulesPackages(input)
}

func (a *InspectorActivities) GetAssessmentReport(input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error) {
    return a.client.GetAssessmentReport(input)
}

func (a *InspectorActivities) GetExclusionsPreview(input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error) {
    return a.client.GetExclusionsPreview(input)
}

func (a *InspectorActivities) GetTelemetryMetadata(input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error) {
    return a.client.GetTelemetryMetadata(input)
}

func (a *InspectorActivities) ListAssessmentRunAgents(input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error) {
    return a.client.ListAssessmentRunAgents(input)
}

func (a *InspectorActivities) ListAssessmentRuns(input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error) {
    return a.client.ListAssessmentRuns(input)
}

func (a *InspectorActivities) ListAssessmentTargets(input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error) {
    return a.client.ListAssessmentTargets(input)
}

func (a *InspectorActivities) ListAssessmentTemplates(input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error) {
    return a.client.ListAssessmentTemplates(input)
}

func (a *InspectorActivities) ListEventSubscriptions(input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error) {
    return a.client.ListEventSubscriptions(input)
}

func (a *InspectorActivities) ListExclusions(input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error) {
    return a.client.ListExclusions(input)
}

func (a *InspectorActivities) ListFindings(input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error) {
    return a.client.ListFindings(input)
}

func (a *InspectorActivities) ListRulesPackages(input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error) {
    return a.client.ListRulesPackages(input)
}

func (a *InspectorActivities) ListTagsForResource(input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *InspectorActivities) PreviewAgents(input *inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error) {
    return a.client.PreviewAgents(input)
}

func (a *InspectorActivities) RegisterCrossAccountAccessRole(input *inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error) {
    return a.client.RegisterCrossAccountAccessRole(input)
}

func (a *InspectorActivities) RemoveAttributesFromFindings(input *inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error) {
    return a.client.RemoveAttributesFromFindings(input)
}

func (a *InspectorActivities) SetTagsForResource(input *inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error) {
    return a.client.SetTagsForResource(input)
}

func (a *InspectorActivities) StartAssessmentRun(input *inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error) {
    return a.client.StartAssessmentRun(input)
}

func (a *InspectorActivities) StopAssessmentRun(input *inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error) {
    return a.client.StopAssessmentRun(input)
}

func (a *InspectorActivities) SubscribeToEvent(input *inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error) {
    return a.client.SubscribeToEvent(input)
}

func (a *InspectorActivities) UnsubscribeFromEvent(input *inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error) {
    return a.client.UnsubscribeFromEvent(input)
}

func (a *InspectorActivities) UpdateAssessmentTarget(input *inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error) {
    return a.client.UpdateAssessmentTarget(input)
}
