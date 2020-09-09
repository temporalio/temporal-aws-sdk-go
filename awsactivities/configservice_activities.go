package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
)

type ConfigServiceActivities struct {
	client configserviceiface.ConfigServiceAPI
}

func NewConfigServiceActivities(session *session.Session, config ...*aws.Config) *ConfigServiceActivities {
	client := configservice.New(session, config...)
	return &ConfigServiceActivities{client: client}
}

func (a *ConfigServiceActivities) BatchGetAggregateResourceConfig(input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
	return a.client.BatchGetAggregateResourceConfig(input)
}

func (a *ConfigServiceActivities) BatchGetResourceConfig(input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error) {
	return a.client.BatchGetResourceConfig(input)
}

func (a *ConfigServiceActivities) DeleteAggregationAuthorization(input *configservice.DeleteAggregationAuthorizationInput) (*configservice.DeleteAggregationAuthorizationOutput, error) {
	return a.client.DeleteAggregationAuthorization(input)
}

func (a *ConfigServiceActivities) DeleteConfigRule(input *configservice.DeleteConfigRuleInput) (*configservice.DeleteConfigRuleOutput, error) {
	return a.client.DeleteConfigRule(input)
}

func (a *ConfigServiceActivities) DeleteConfigurationAggregator(input *configservice.DeleteConfigurationAggregatorInput) (*configservice.DeleteConfigurationAggregatorOutput, error) {
	return a.client.DeleteConfigurationAggregator(input)
}

func (a *ConfigServiceActivities) DeleteConfigurationRecorder(input *configservice.DeleteConfigurationRecorderInput) (*configservice.DeleteConfigurationRecorderOutput, error) {
	return a.client.DeleteConfigurationRecorder(input)
}

func (a *ConfigServiceActivities) DeleteConformancePack(input *configservice.DeleteConformancePackInput) (*configservice.DeleteConformancePackOutput, error) {
	return a.client.DeleteConformancePack(input)
}

func (a *ConfigServiceActivities) DeleteDeliveryChannel(input *configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error) {
	return a.client.DeleteDeliveryChannel(input)
}

func (a *ConfigServiceActivities) DeleteEvaluationResults(input *configservice.DeleteEvaluationResultsInput) (*configservice.DeleteEvaluationResultsOutput, error) {
	return a.client.DeleteEvaluationResults(input)
}

func (a *ConfigServiceActivities) DeleteOrganizationConfigRule(input *configservice.DeleteOrganizationConfigRuleInput) (*configservice.DeleteOrganizationConfigRuleOutput, error) {
	return a.client.DeleteOrganizationConfigRule(input)
}

func (a *ConfigServiceActivities) DeleteOrganizationConformancePack(input *configservice.DeleteOrganizationConformancePackInput) (*configservice.DeleteOrganizationConformancePackOutput, error) {
	return a.client.DeleteOrganizationConformancePack(input)
}

func (a *ConfigServiceActivities) DeleteRemediationConfiguration(input *configservice.DeleteRemediationConfigurationInput) (*configservice.DeleteRemediationConfigurationOutput, error) {
	return a.client.DeleteRemediationConfiguration(input)
}

func (a *ConfigServiceActivities) DeleteRemediationExceptions(input *configservice.DeleteRemediationExceptionsInput) (*configservice.DeleteRemediationExceptionsOutput, error) {
	return a.client.DeleteRemediationExceptions(input)
}

func (a *ConfigServiceActivities) DeleteResourceConfig(input *configservice.DeleteResourceConfigInput) (*configservice.DeleteResourceConfigOutput, error) {
	return a.client.DeleteResourceConfig(input)
}

func (a *ConfigServiceActivities) DeleteRetentionConfiguration(input *configservice.DeleteRetentionConfigurationInput) (*configservice.DeleteRetentionConfigurationOutput, error) {
	return a.client.DeleteRetentionConfiguration(input)
}

func (a *ConfigServiceActivities) DeliverConfigSnapshot(input *configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error) {
	return a.client.DeliverConfigSnapshot(input)
}

func (a *ConfigServiceActivities) DescribeAggregateComplianceByConfigRules(input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
	return a.client.DescribeAggregateComplianceByConfigRules(input)
}

func (a *ConfigServiceActivities) DescribeAggregationAuthorizations(input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
	return a.client.DescribeAggregationAuthorizations(input)
}

func (a *ConfigServiceActivities) DescribeComplianceByConfigRule(input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
	return a.client.DescribeComplianceByConfigRule(input)
}

func (a *ConfigServiceActivities) DescribeComplianceByResource(input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error) {
	return a.client.DescribeComplianceByResource(input)
}

func (a *ConfigServiceActivities) DescribeConfigRuleEvaluationStatus(input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
	return a.client.DescribeConfigRuleEvaluationStatus(input)
}

func (a *ConfigServiceActivities) DescribeConfigRules(input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error) {
	return a.client.DescribeConfigRules(input)
}

func (a *ConfigServiceActivities) DescribeConfigurationAggregatorSourcesStatus(input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	return a.client.DescribeConfigurationAggregatorSourcesStatus(input)
}

func (a *ConfigServiceActivities) DescribeConfigurationAggregators(input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
	return a.client.DescribeConfigurationAggregators(input)
}

func (a *ConfigServiceActivities) DescribeConfigurationRecorderStatus(input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
	return a.client.DescribeConfigurationRecorderStatus(input)
}

func (a *ConfigServiceActivities) DescribeConfigurationRecorders(input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error) {
	return a.client.DescribeConfigurationRecorders(input)
}

func (a *ConfigServiceActivities) DescribeConformancePackCompliance(input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error) {
	return a.client.DescribeConformancePackCompliance(input)
}

func (a *ConfigServiceActivities) DescribeConformancePackStatus(input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error) {
	return a.client.DescribeConformancePackStatus(input)
}

func (a *ConfigServiceActivities) DescribeConformancePacks(input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error) {
	return a.client.DescribeConformancePacks(input)
}

func (a *ConfigServiceActivities) DescribeDeliveryChannelStatus(input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
	return a.client.DescribeDeliveryChannelStatus(input)
}

func (a *ConfigServiceActivities) DescribeDeliveryChannels(input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error) {
	return a.client.DescribeDeliveryChannels(input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConfigRuleStatuses(input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
	return a.client.DescribeOrganizationConfigRuleStatuses(input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConfigRules(input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
	return a.client.DescribeOrganizationConfigRules(input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConformancePackStatuses(input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
	return a.client.DescribeOrganizationConformancePackStatuses(input)
}

func (a *ConfigServiceActivities) DescribeOrganizationConformancePacks(input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
	return a.client.DescribeOrganizationConformancePacks(input)
}

func (a *ConfigServiceActivities) DescribePendingAggregationRequests(input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error) {
	return a.client.DescribePendingAggregationRequests(input)
}

func (a *ConfigServiceActivities) DescribeRemediationConfigurations(input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error) {
	return a.client.DescribeRemediationConfigurations(input)
}

func (a *ConfigServiceActivities) DescribeRemediationExceptions(input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error) {
	return a.client.DescribeRemediationExceptions(input)
}

func (a *ConfigServiceActivities) DescribeRemediationExecutionStatus(input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
	return a.client.DescribeRemediationExecutionStatus(input)
}

func (a *ConfigServiceActivities) DescribeRetentionConfigurations(input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error) {
	return a.client.DescribeRetentionConfigurations(input)
}

func (a *ConfigServiceActivities) GetAggregateComplianceDetailsByConfigRule(input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	return a.client.GetAggregateComplianceDetailsByConfigRule(input)
}

func (a *ConfigServiceActivities) GetAggregateConfigRuleComplianceSummary(input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
	return a.client.GetAggregateConfigRuleComplianceSummary(input)
}

func (a *ConfigServiceActivities) GetAggregateDiscoveredResourceCounts(input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
	return a.client.GetAggregateDiscoveredResourceCounts(input)
}

func (a *ConfigServiceActivities) GetAggregateResourceConfig(input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error) {
	return a.client.GetAggregateResourceConfig(input)
}

func (a *ConfigServiceActivities) GetComplianceDetailsByConfigRule(input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
	return a.client.GetComplianceDetailsByConfigRule(input)
}

func (a *ConfigServiceActivities) GetComplianceDetailsByResource(input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error) {
	return a.client.GetComplianceDetailsByResource(input)
}

func (a *ConfigServiceActivities) GetComplianceSummaryByConfigRule(input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
	return a.client.GetComplianceSummaryByConfigRule(input)
}

func (a *ConfigServiceActivities) GetComplianceSummaryByResourceType(input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
	return a.client.GetComplianceSummaryByResourceType(input)
}

func (a *ConfigServiceActivities) GetConformancePackComplianceDetails(input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
	return a.client.GetConformancePackComplianceDetails(input)
}

func (a *ConfigServiceActivities) GetConformancePackComplianceSummary(input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
	return a.client.GetConformancePackComplianceSummary(input)
}

func (a *ConfigServiceActivities) GetDiscoveredResourceCounts(input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	return a.client.GetDiscoveredResourceCounts(input)
}

func (a *ConfigServiceActivities) GetOrganizationConfigRuleDetailedStatus(input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
	return a.client.GetOrganizationConfigRuleDetailedStatus(input)
}

func (a *ConfigServiceActivities) GetOrganizationConformancePackDetailedStatus(input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
	return a.client.GetOrganizationConformancePackDetailedStatus(input)
}

func (a *ConfigServiceActivities) GetResourceConfigHistory(input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error) {
	return a.client.GetResourceConfigHistory(input)
}

func (a *ConfigServiceActivities) ListAggregateDiscoveredResources(input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
	return a.client.ListAggregateDiscoveredResources(input)
}

func (a *ConfigServiceActivities) ListDiscoveredResources(input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error) {
	return a.client.ListDiscoveredResources(input)
}

func (a *ConfigServiceActivities) ListTagsForResource(input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *ConfigServiceActivities) PutAggregationAuthorization(input *configservice.PutAggregationAuthorizationInput) (*configservice.PutAggregationAuthorizationOutput, error) {
	return a.client.PutAggregationAuthorization(input)
}

func (a *ConfigServiceActivities) PutConfigRule(input *configservice.PutConfigRuleInput) (*configservice.PutConfigRuleOutput, error) {
	return a.client.PutConfigRule(input)
}

func (a *ConfigServiceActivities) PutConfigurationAggregator(input *configservice.PutConfigurationAggregatorInput) (*configservice.PutConfigurationAggregatorOutput, error) {
	return a.client.PutConfigurationAggregator(input)
}

func (a *ConfigServiceActivities) PutConfigurationRecorder(input *configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error) {
	return a.client.PutConfigurationRecorder(input)
}

func (a *ConfigServiceActivities) PutConformancePack(input *configservice.PutConformancePackInput) (*configservice.PutConformancePackOutput, error) {
	return a.client.PutConformancePack(input)
}

func (a *ConfigServiceActivities) PutDeliveryChannel(input *configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error) {
	return a.client.PutDeliveryChannel(input)
}

func (a *ConfigServiceActivities) PutEvaluations(input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error) {
	return a.client.PutEvaluations(input)
}

func (a *ConfigServiceActivities) PutOrganizationConfigRule(input *configservice.PutOrganizationConfigRuleInput) (*configservice.PutOrganizationConfigRuleOutput, error) {
	return a.client.PutOrganizationConfigRule(input)
}

func (a *ConfigServiceActivities) PutOrganizationConformancePack(input *configservice.PutOrganizationConformancePackInput) (*configservice.PutOrganizationConformancePackOutput, error) {
	return a.client.PutOrganizationConformancePack(input)
}

func (a *ConfigServiceActivities) PutRemediationConfigurations(input *configservice.PutRemediationConfigurationsInput) (*configservice.PutRemediationConfigurationsOutput, error) {
	return a.client.PutRemediationConfigurations(input)
}

func (a *ConfigServiceActivities) PutRemediationExceptions(input *configservice.PutRemediationExceptionsInput) (*configservice.PutRemediationExceptionsOutput, error) {
	return a.client.PutRemediationExceptions(input)
}

func (a *ConfigServiceActivities) PutResourceConfig(input *configservice.PutResourceConfigInput) (*configservice.PutResourceConfigOutput, error) {
	return a.client.PutResourceConfig(input)
}

func (a *ConfigServiceActivities) PutRetentionConfiguration(input *configservice.PutRetentionConfigurationInput) (*configservice.PutRetentionConfigurationOutput, error) {
	return a.client.PutRetentionConfiguration(input)
}

func (a *ConfigServiceActivities) SelectAggregateResourceConfig(input *configservice.SelectAggregateResourceConfigInput) (*configservice.SelectAggregateResourceConfigOutput, error) {
	return a.client.SelectAggregateResourceConfig(input)
}

func (a *ConfigServiceActivities) SelectResourceConfig(input *configservice.SelectResourceConfigInput) (*configservice.SelectResourceConfigOutput, error) {
	return a.client.SelectResourceConfig(input)
}

func (a *ConfigServiceActivities) StartConfigRulesEvaluation(input *configservice.StartConfigRulesEvaluationInput) (*configservice.StartConfigRulesEvaluationOutput, error) {
	return a.client.StartConfigRulesEvaluation(input)
}

func (a *ConfigServiceActivities) StartConfigurationRecorder(input *configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error) {
	return a.client.StartConfigurationRecorder(input)
}

func (a *ConfigServiceActivities) StartRemediationExecution(input *configservice.StartRemediationExecutionInput) (*configservice.StartRemediationExecutionOutput, error) {
	return a.client.StartRemediationExecution(input)
}

func (a *ConfigServiceActivities) StopConfigurationRecorder(input *configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error) {
	return a.client.StopConfigurationRecorder(input)
}

func (a *ConfigServiceActivities) TagResource(input *configservice.TagResourceInput) (*configservice.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *ConfigServiceActivities) UntagResource(input *configservice.UntagResourceInput) (*configservice.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}
