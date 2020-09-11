package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ConfigServiceActivities struct {
	client configserviceiface.ConfigServiceAPI
}

func NewConfigServiceActivities(session *session.Session, config ...*aws.Config) *ConfigServiceActivities {
	client := configservice.New(session, config...)
	return &ConfigServiceActivities{client: client}
}

func (a *ConfigServiceActivities) BatchGetAggregateResourceConfig(ctx context.Context, input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
	return a.client.BatchGetAggregateResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) BatchGetResourceConfig(ctx context.Context, input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error) {
	return a.client.BatchGetResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteAggregationAuthorization(ctx context.Context, input *configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error) {
	return a.client.DeleteAggregationAuthorizationWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteConfigRule(ctx context.Context, input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
	return a.client.DeleteConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteConfigurationAggregator(ctx context.Context, input *configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error) {
	return a.client.DeleteConfigurationAggregatorWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteConfigurationRecorder(ctx context.Context, input *configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error) {
	return a.client.DeleteConfigurationRecorderWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteConformancePack(ctx context.Context, input *configservice.DeleteConformancePackInput) (*configservice.DeleteConformancePackOutput, error) {
	return a.client.DeleteConformancePackWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteDeliveryChannel(ctx context.Context, input *configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error) {
	return a.client.DeleteDeliveryChannelWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteEvaluationResults(ctx context.Context, input *configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error) {
	return a.client.DeleteEvaluationResultsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteOrganizationConfigRule(ctx context.Context, input *configservice.DeleteOrganizationConfigRuleInput) (*configservice.DeleteOrganizationConfigRuleOutput, error) {
	return a.client.DeleteOrganizationConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteOrganizationConformancePack(ctx context.Context, input *configservice.DeleteOrganizationConformancePackInput) (*configservice.DeleteOrganizationConformancePackOutput, error) {
	return a.client.DeleteOrganizationConformancePackWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteRemediationConfiguration(ctx context.Context, input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error) {
	return a.client.DeleteRemediationConfigurationWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteRemediationExceptions(ctx context.Context, input *configservice.DeleteRemediationExceptionsInput) (*configservice.DeleteRemediationExceptionsOutput, error) {
	return a.client.DeleteRemediationExceptionsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteResourceConfig(ctx context.Context, input *configservice.DeleteResourceConfigInput) (*configservice.DeleteResourceConfigOutput, error) {
	return a.client.DeleteResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeleteRetentionConfiguration(ctx context.Context, input *configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error) {
	return a.client.DeleteRetentionConfigurationWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DeliverConfigSnapshot(ctx context.Context, input *configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error) {
	return a.client.DeliverConfigSnapshotWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeAggregateComplianceByConfigRules(ctx context.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
	return a.client.DescribeAggregateComplianceByConfigRulesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeAggregationAuthorizations(ctx context.Context, input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
	return a.client.DescribeAggregationAuthorizationsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeComplianceByConfigRule(ctx context.Context, input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
	return a.client.DescribeComplianceByConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeComplianceByResource(ctx context.Context, input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error) {
	return a.client.DescribeComplianceByResourceWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConfigRuleEvaluationStatus(ctx context.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
	return a.client.DescribeConfigRuleEvaluationStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConfigRules(ctx context.Context, input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error) {
	return a.client.DescribeConfigRulesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConfigurationAggregatorSourcesStatus(ctx context.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	return a.client.DescribeConfigurationAggregatorSourcesStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConfigurationAggregators(ctx context.Context, input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
	return a.client.DescribeConfigurationAggregatorsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConfigurationRecorderStatus(ctx context.Context, input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
	return a.client.DescribeConfigurationRecorderStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConfigurationRecorders(ctx context.Context, input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error) {
	return a.client.DescribeConfigurationRecordersWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConformancePackCompliance(ctx context.Context, input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error) {
	return a.client.DescribeConformancePackComplianceWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConformancePackStatus(ctx context.Context, input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error) {
	return a.client.DescribeConformancePackStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeConformancePacks(ctx context.Context, input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error) {
	return a.client.DescribeConformancePacksWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeDeliveryChannelStatus(ctx context.Context, input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
	return a.client.DescribeDeliveryChannelStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeDeliveryChannels(ctx context.Context, input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error) {
	return a.client.DescribeDeliveryChannelsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConfigRuleStatuses(ctx context.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
	return a.client.DescribeOrganizationConfigRuleStatusesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConfigRules(ctx context.Context, input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
	return a.client.DescribeOrganizationConfigRulesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConformancePackStatuses(ctx context.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
	return a.client.DescribeOrganizationConformancePackStatusesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConformancePacks(ctx context.Context, input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
	return a.client.DescribeOrganizationConformancePacksWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribePendingAggregationRequests(ctx context.Context, input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error) {
	return a.client.DescribePendingAggregationRequestsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeRemediationConfigurations(ctx context.Context, input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error) {
	return a.client.DescribeRemediationConfigurationsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeRemediationExceptions(ctx context.Context, input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error) {
	return a.client.DescribeRemediationExceptionsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeRemediationExecutionStatus(ctx context.Context, input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
	return a.client.DescribeRemediationExecutionStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) DescribeRetentionConfigurations(ctx context.Context, input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error) {
	return a.client.DescribeRetentionConfigurationsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetAggregateComplianceDetailsByConfigRule(ctx context.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	return a.client.GetAggregateComplianceDetailsByConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetAggregateConfigRuleComplianceSummary(ctx context.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
	return a.client.GetAggregateConfigRuleComplianceSummaryWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetAggregateDiscoveredResourceCounts(ctx context.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
	return a.client.GetAggregateDiscoveredResourceCountsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetAggregateResourceConfig(ctx context.Context, input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error) {
	return a.client.GetAggregateResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetComplianceDetailsByConfigRule(ctx context.Context, input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
	return a.client.GetComplianceDetailsByConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetComplianceDetailsByResource(ctx context.Context, input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error) {
	return a.client.GetComplianceDetailsByResourceWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetComplianceSummaryByConfigRule(ctx context.Context, input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
	return a.client.GetComplianceSummaryByConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetComplianceSummaryByResourceType(ctx context.Context, input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
	return a.client.GetComplianceSummaryByResourceTypeWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetConformancePackComplianceDetails(ctx context.Context, input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
	return a.client.GetConformancePackComplianceDetailsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetConformancePackComplianceSummary(ctx context.Context, input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
	return a.client.GetConformancePackComplianceSummaryWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetDiscoveredResourceCounts(ctx context.Context, input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	return a.client.GetDiscoveredResourceCountsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetOrganizationConfigRuleDetailedStatus(ctx context.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
	return a.client.GetOrganizationConfigRuleDetailedStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetOrganizationConformancePackDetailedStatus(ctx context.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
	return a.client.GetOrganizationConformancePackDetailedStatusWithContext(ctx, input)
}

func (a *ConfigServiceActivities) GetResourceConfigHistory(ctx context.Context, input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error) {
	return a.client.GetResourceConfigHistoryWithContext(ctx, input)
}

func (a *ConfigServiceActivities) ListAggregateDiscoveredResources(ctx context.Context, input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
	return a.client.ListAggregateDiscoveredResourcesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) ListDiscoveredResources(ctx context.Context, input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error) {
	return a.client.ListDiscoveredResourcesWithContext(ctx, input)
}

func (a *ConfigServiceActivities) ListTagsForResource(ctx context.Context, input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutAggregationAuthorization(ctx context.Context, input *configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error) {
	return a.client.PutAggregationAuthorizationWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutConfigRule(ctx context.Context, input *configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error) {
	return a.client.PutConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutConfigurationAggregator(ctx context.Context, input *configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error) {
	return a.client.PutConfigurationAggregatorWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutConfigurationRecorder(ctx context.Context, input *configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error) {
	return a.client.PutConfigurationRecorderWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutConformancePack(ctx context.Context, input *configservice.PutConformancePackInput) (*configservice.PutConformancePackOutput, error) {
	return a.client.PutConformancePackWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutDeliveryChannel(ctx context.Context, input *configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error) {
	return a.client.PutDeliveryChannelWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutEvaluations(ctx context.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error) {
	return a.client.PutEvaluationsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutOrganizationConfigRule(ctx context.Context, input *configservice.PutOrganizationConfigRuleInput) (*configservice.PutOrganizationConfigRuleOutput, error) {
	return a.client.PutOrganizationConfigRuleWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutOrganizationConformancePack(ctx context.Context, input *configservice.PutOrganizationConformancePackInput) (*configservice.PutOrganizationConformancePackOutput, error) {
	return a.client.PutOrganizationConformancePackWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutRemediationConfigurations(ctx context.Context, input *configservice.PutRemediationConfigurationsInput) (*configservice.PutRemediationConfigurationsOutput, error) {
	return a.client.PutRemediationConfigurationsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutRemediationExceptions(ctx context.Context, input *configservice.PutRemediationExceptionsInput) (*configservice.PutRemediationExceptionsOutput, error) {
	return a.client.PutRemediationExceptionsWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutResourceConfig(ctx context.Context, input *configservice.PutResourceConfigInput) (*configservice.PutResourceConfigOutput, error) {
	return a.client.PutResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) PutRetentionConfiguration(ctx context.Context, input *configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error) {
	return a.client.PutRetentionConfigurationWithContext(ctx, input)
}

func (a *ConfigServiceActivities) SelectAggregateResourceConfig(ctx context.Context, input *configservice.SelectAggregateResourceConfigInput) (*configservice.SelectAggregateResourceConfigOutput, error) {
	return a.client.SelectAggregateResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) SelectResourceConfig(ctx context.Context, input *configservice.SelectResourceConfigInput) (*configservice.SelectResourceConfigOutput, error) {
	return a.client.SelectResourceConfigWithContext(ctx, input)
}

func (a *ConfigServiceActivities) StartConfigRulesEvaluation(ctx context.Context, input *configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error) {
	return a.client.StartConfigRulesEvaluationWithContext(ctx, input)
}

func (a *ConfigServiceActivities) StartConfigurationRecorder(ctx context.Context, input *configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error) {
	return a.client.StartConfigurationRecorderWithContext(ctx, input)
}

func (a *ConfigServiceActivities) StartRemediationExecution(ctx context.Context, input *configservice.StartRemediationExecutionInput) (*configservice.StartRemediationExecutionOutput, error) {
	return a.client.StartRemediationExecutionWithContext(ctx, input)
}

func (a *ConfigServiceActivities) StopConfigurationRecorder(ctx context.Context, input *configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error) {
	return a.client.StopConfigurationRecorderWithContext(ctx, input)
}

func (a *ConfigServiceActivities) TagResource(ctx context.Context, input *configservice.TagResourceInput) (*configservice.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ConfigServiceActivities) UntagResource(ctx context.Context, input *configservice.UntagResourceInput) (*configservice.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
