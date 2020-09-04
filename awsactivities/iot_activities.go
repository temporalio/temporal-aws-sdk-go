
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot/iotiface"
)

type IoTActivities struct {
	client iotiface.IoTAPI
}

func NewIoTActivities(client iotiface.IoTAPI) *IoTActivities {
    return &IoTActivities{client: client}
}

func (a *IoTActivities) AcceptCertificateTransfer(input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error) {
    return a.client.AcceptCertificateTransfer(input)
}

func (a *IoTActivities) AddThingToBillingGroup(input *iot.AddThingToBillingGroupInput) (*iot.AddThingToBillingGroupOutput, error) {
    return a.client.AddThingToBillingGroup(input)
}

func (a *IoTActivities) AddThingToThingGroup(input *iot.AddThingToThingGroupInput) (*iot.AddThingToThingGroupOutput, error) {
    return a.client.AddThingToThingGroup(input)
}

func (a *IoTActivities) AssociateTargetsWithJob(input *iot.AssociateTargetsWithJobInput) (*iot.AssociateTargetsWithJobOutput, error) {
    return a.client.AssociateTargetsWithJob(input)
}

func (a *IoTActivities) AttachPolicy(input *iot.AttachPolicyInput) (*iot.AttachPolicyOutput, error) {
    return a.client.AttachPolicy(input)
}

func (a *IoTActivities) AttachPrincipalPolicy(input *iot.AttachPrincipalPolicyInput) (*iot.AttachPrincipalPolicyOutput, error) {
    return a.client.AttachPrincipalPolicy(input)
}

func (a *IoTActivities) AttachSecurityProfile(input *iot.AttachSecurityProfileInput) (*iot.AttachSecurityProfileOutput, error) {
    return a.client.AttachSecurityProfile(input)
}

func (a *IoTActivities) AttachThingPrincipal(input *iot.AttachThingPrincipalInput) (*iot.AttachThingPrincipalOutput, error) {
    return a.client.AttachThingPrincipal(input)
}

func (a *IoTActivities) CancelAuditMitigationActionsTask(input *iot.CancelAuditMitigationActionsTaskInput) (*iot.CancelAuditMitigationActionsTaskOutput, error) {
    return a.client.CancelAuditMitigationActionsTask(input)
}

func (a *IoTActivities) CancelAuditTask(input *iot.CancelAuditTaskInput) (*iot.CancelAuditTaskOutput, error) {
    return a.client.CancelAuditTask(input)
}

func (a *IoTActivities) CancelCertificateTransfer(input *iot.CancelCertificateTransferInput) (*iot.CancelCertificateTransferOutput, error) {
    return a.client.CancelCertificateTransfer(input)
}

func (a *IoTActivities) CancelJob(input *iot.CancelJobInput) (*iot.CancelJobOutput, error) {
    return a.client.CancelJob(input)
}

func (a *IoTActivities) CancelJobExecution(input *iot.CancelJobExecutionInput) (*iot.CancelJobExecutionOutput, error) {
    return a.client.CancelJobExecution(input)
}

func (a *IoTActivities) ClearDefaultAuthorizer(input *iot.ClearDefaultAuthorizerInput) (*iot.ClearDefaultAuthorizerOutput, error) {
    return a.client.ClearDefaultAuthorizer(input)
}

func (a *IoTActivities) ConfirmTopicRuleDestination(input *iot.ConfirmTopicRuleDestinationInput) (*iot.ConfirmTopicRuleDestinationOutput, error) {
    return a.client.ConfirmTopicRuleDestination(input)
}

func (a *IoTActivities) CreateAuditSuppression(input *iot.CreateAuditSuppressionInput) (*iot.CreateAuditSuppressionOutput, error) {
    return a.client.CreateAuditSuppression(input)
}

func (a *IoTActivities) CreateAuthorizer(input *iot.CreateAuthorizerInput) (*iot.CreateAuthorizerOutput, error) {
    return a.client.CreateAuthorizer(input)
}

func (a *IoTActivities) CreateBillingGroup(input *iot.CreateBillingGroupInput) (*iot.CreateBillingGroupOutput, error) {
    return a.client.CreateBillingGroup(input)
}

func (a *IoTActivities) CreateCertificateFromCsr(input *iot.CreateCertificateFromCsrInput) (*iot.CreateCertificateFromCsrOutput, error) {
    return a.client.CreateCertificateFromCsr(input)
}

func (a *IoTActivities) CreateDimension(input *iot.CreateDimensionInput) (*iot.CreateDimensionOutput, error) {
    return a.client.CreateDimension(input)
}

func (a *IoTActivities) CreateDomainConfiguration(input *iot.CreateDomainConfigurationInput) (*iot.CreateDomainConfigurationOutput, error) {
    return a.client.CreateDomainConfiguration(input)
}

func (a *IoTActivities) CreateDynamicThingGroup(input *iot.CreateDynamicThingGroupInput) (*iot.CreateDynamicThingGroupOutput, error) {
    return a.client.CreateDynamicThingGroup(input)
}

func (a *IoTActivities) CreateJob(input *iot.CreateJobInput) (*iot.CreateJobOutput, error) {
    return a.client.CreateJob(input)
}

func (a *IoTActivities) CreateKeysAndCertificate(input *iot.CreateKeysAndCertificateInput) (*iot.CreateKeysAndCertificateOutput, error) {
    return a.client.CreateKeysAndCertificate(input)
}

func (a *IoTActivities) CreateMitigationAction(input *iot.CreateMitigationActionInput) (*iot.CreateMitigationActionOutput, error) {
    return a.client.CreateMitigationAction(input)
}

func (a *IoTActivities) CreateOTAUpdate(input *iot.CreateOTAUpdateInput) (*iot.CreateOTAUpdateOutput, error) {
    return a.client.CreateOTAUpdate(input)
}

func (a *IoTActivities) CreatePolicy(input *iot.CreatePolicyInput) (*iot.CreatePolicyOutput, error) {
    return a.client.CreatePolicy(input)
}

func (a *IoTActivities) CreatePolicyVersion(input *iot.CreatePolicyVersionInput) (*iot.CreatePolicyVersionOutput, error) {
    return a.client.CreatePolicyVersion(input)
}

func (a *IoTActivities) CreateProvisioningClaim(input *iot.CreateProvisioningClaimInput) (*iot.CreateProvisioningClaimOutput, error) {
    return a.client.CreateProvisioningClaim(input)
}

func (a *IoTActivities) CreateProvisioningTemplate(input *iot.CreateProvisioningTemplateInput) (*iot.CreateProvisioningTemplateOutput, error) {
    return a.client.CreateProvisioningTemplate(input)
}

func (a *IoTActivities) CreateProvisioningTemplateVersion(input *iot.CreateProvisioningTemplateVersionInput) (*iot.CreateProvisioningTemplateVersionOutput, error) {
    return a.client.CreateProvisioningTemplateVersion(input)
}

func (a *IoTActivities) CreateRoleAlias(input *iot.CreateRoleAliasInput) (*iot.CreateRoleAliasOutput, error) {
    return a.client.CreateRoleAlias(input)
}

func (a *IoTActivities) CreateScheduledAudit(input *iot.CreateScheduledAuditInput) (*iot.CreateScheduledAuditOutput, error) {
    return a.client.CreateScheduledAudit(input)
}

func (a *IoTActivities) CreateSecurityProfile(input *iot.CreateSecurityProfileInput) (*iot.CreateSecurityProfileOutput, error) {
    return a.client.CreateSecurityProfile(input)
}

func (a *IoTActivities) CreateStream(input *iot.CreateStreamInput) (*iot.CreateStreamOutput, error) {
    return a.client.CreateStream(input)
}

func (a *IoTActivities) CreateThing(input *iot.CreateThingInput) (*iot.CreateThingOutput, error) {
    return a.client.CreateThing(input)
}

func (a *IoTActivities) CreateThingGroup(input *iot.CreateThingGroupInput) (*iot.CreateThingGroupOutput, error) {
    return a.client.CreateThingGroup(input)
}

func (a *IoTActivities) CreateThingType(input *iot.CreateThingTypeInput) (*iot.CreateThingTypeOutput, error) {
    return a.client.CreateThingType(input)
}

func (a *IoTActivities) CreateTopicRule(input *iot.CreateTopicRuleInput) (*iot.CreateTopicRuleOutput, error) {
    return a.client.CreateTopicRule(input)
}

func (a *IoTActivities) CreateTopicRuleDestination(input *iot.CreateTopicRuleDestinationInput) (*iot.CreateTopicRuleDestinationOutput, error) {
    return a.client.CreateTopicRuleDestination(input)
}

func (a *IoTActivities) DeleteAccountAuditConfiguration(input *iot.DeleteAccountAuditConfigurationInput) (*iot.DeleteAccountAuditConfigurationOutput, error) {
    return a.client.DeleteAccountAuditConfiguration(input)
}

func (a *IoTActivities) DeleteAuditSuppression(input *iot.DeleteAuditSuppressionInput) (*iot.DeleteAuditSuppressionOutput, error) {
    return a.client.DeleteAuditSuppression(input)
}

func (a *IoTActivities) DeleteAuthorizer(input *iot.DeleteAuthorizerInput) (*iot.DeleteAuthorizerOutput, error) {
    return a.client.DeleteAuthorizer(input)
}

func (a *IoTActivities) DeleteBillingGroup(input *iot.DeleteBillingGroupInput) (*iot.DeleteBillingGroupOutput, error) {
    return a.client.DeleteBillingGroup(input)
}

func (a *IoTActivities) DeleteCACertificate(input *iot.DeleteCACertificateInput) (*iot.DeleteCACertificateOutput, error) {
    return a.client.DeleteCACertificate(input)
}

func (a *IoTActivities) DeleteCertificate(input *iot.DeleteCertificateInput) (*iot.DeleteCertificateOutput, error) {
    return a.client.DeleteCertificate(input)
}

func (a *IoTActivities) DeleteDimension(input *iot.DeleteDimensionInput) (*iot.DeleteDimensionOutput, error) {
    return a.client.DeleteDimension(input)
}

func (a *IoTActivities) DeleteDomainConfiguration(input *iot.DeleteDomainConfigurationInput) (*iot.DeleteDomainConfigurationOutput, error) {
    return a.client.DeleteDomainConfiguration(input)
}

func (a *IoTActivities) DeleteDynamicThingGroup(input *iot.DeleteDynamicThingGroupInput) (*iot.DeleteDynamicThingGroupOutput, error) {
    return a.client.DeleteDynamicThingGroup(input)
}

func (a *IoTActivities) DeleteJob(input *iot.DeleteJobInput) (*iot.DeleteJobOutput, error) {
    return a.client.DeleteJob(input)
}

func (a *IoTActivities) DeleteJobExecution(input *iot.DeleteJobExecutionInput) (*iot.DeleteJobExecutionOutput, error) {
    return a.client.DeleteJobExecution(input)
}

func (a *IoTActivities) DeleteMitigationAction(input *iot.DeleteMitigationActionInput) (*iot.DeleteMitigationActionOutput, error) {
    return a.client.DeleteMitigationAction(input)
}

func (a *IoTActivities) DeleteOTAUpdate(input *iot.DeleteOTAUpdateInput) (*iot.DeleteOTAUpdateOutput, error) {
    return a.client.DeleteOTAUpdate(input)
}

func (a *IoTActivities) DeletePolicy(input *iot.DeletePolicyInput) (*iot.DeletePolicyOutput, error) {
    return a.client.DeletePolicy(input)
}

func (a *IoTActivities) DeletePolicyVersion(input *iot.DeletePolicyVersionInput) (*iot.DeletePolicyVersionOutput, error) {
    return a.client.DeletePolicyVersion(input)
}

func (a *IoTActivities) DeleteProvisioningTemplate(input *iot.DeleteProvisioningTemplateInput) (*iot.DeleteProvisioningTemplateOutput, error) {
    return a.client.DeleteProvisioningTemplate(input)
}

func (a *IoTActivities) DeleteProvisioningTemplateVersion(input *iot.DeleteProvisioningTemplateVersionInput) (*iot.DeleteProvisioningTemplateVersionOutput, error) {
    return a.client.DeleteProvisioningTemplateVersion(input)
}

func (a *IoTActivities) DeleteRegistrationCode(input *iot.DeleteRegistrationCodeInput) (*iot.DeleteRegistrationCodeOutput, error) {
    return a.client.DeleteRegistrationCode(input)
}

func (a *IoTActivities) DeleteRoleAlias(input *iot.DeleteRoleAliasInput) (*iot.DeleteRoleAliasOutput, error) {
    return a.client.DeleteRoleAlias(input)
}

func (a *IoTActivities) DeleteScheduledAudit(input *iot.DeleteScheduledAuditInput) (*iot.DeleteScheduledAuditOutput, error) {
    return a.client.DeleteScheduledAudit(input)
}

func (a *IoTActivities) DeleteSecurityProfile(input *iot.DeleteSecurityProfileInput) (*iot.DeleteSecurityProfileOutput, error) {
    return a.client.DeleteSecurityProfile(input)
}

func (a *IoTActivities) DeleteStream(input *iot.DeleteStreamInput) (*iot.DeleteStreamOutput, error) {
    return a.client.DeleteStream(input)
}

func (a *IoTActivities) DeleteThing(input *iot.DeleteThingInput) (*iot.DeleteThingOutput, error) {
    return a.client.DeleteThing(input)
}

func (a *IoTActivities) DeleteThingGroup(input *iot.DeleteThingGroupInput) (*iot.DeleteThingGroupOutput, error) {
    return a.client.DeleteThingGroup(input)
}

func (a *IoTActivities) DeleteThingType(input *iot.DeleteThingTypeInput) (*iot.DeleteThingTypeOutput, error) {
    return a.client.DeleteThingType(input)
}

func (a *IoTActivities) DeleteTopicRule(input *iot.DeleteTopicRuleInput) (*iot.DeleteTopicRuleOutput, error) {
    return a.client.DeleteTopicRule(input)
}

func (a *IoTActivities) DeleteTopicRuleDestination(input *iot.DeleteTopicRuleDestinationInput) (*iot.DeleteTopicRuleDestinationOutput, error) {
    return a.client.DeleteTopicRuleDestination(input)
}

func (a *IoTActivities) DeleteV2LoggingLevel(input *iot.DeleteV2LoggingLevelInput) (*iot.DeleteV2LoggingLevelOutput, error) {
    return a.client.DeleteV2LoggingLevel(input)
}

func (a *IoTActivities) DeprecateThingType(input *iot.DeprecateThingTypeInput) (*iot.DeprecateThingTypeOutput, error) {
    return a.client.DeprecateThingType(input)
}

func (a *IoTActivities) DescribeAccountAuditConfiguration(input *iot.DescribeAccountAuditConfigurationInput) (*iot.DescribeAccountAuditConfigurationOutput, error) {
    return a.client.DescribeAccountAuditConfiguration(input)
}

func (a *IoTActivities) DescribeAuditFinding(input *iot.DescribeAuditFindingInput) (*iot.DescribeAuditFindingOutput, error) {
    return a.client.DescribeAuditFinding(input)
}

func (a *IoTActivities) DescribeAuditMitigationActionsTask(input *iot.DescribeAuditMitigationActionsTaskInput) (*iot.DescribeAuditMitigationActionsTaskOutput, error) {
    return a.client.DescribeAuditMitigationActionsTask(input)
}

func (a *IoTActivities) DescribeAuditSuppression(input *iot.DescribeAuditSuppressionInput) (*iot.DescribeAuditSuppressionOutput, error) {
    return a.client.DescribeAuditSuppression(input)
}

func (a *IoTActivities) DescribeAuditTask(input *iot.DescribeAuditTaskInput) (*iot.DescribeAuditTaskOutput, error) {
    return a.client.DescribeAuditTask(input)
}

func (a *IoTActivities) DescribeAuthorizer(input *iot.DescribeAuthorizerInput) (*iot.DescribeAuthorizerOutput, error) {
    return a.client.DescribeAuthorizer(input)
}

func (a *IoTActivities) DescribeBillingGroup(input *iot.DescribeBillingGroupInput) (*iot.DescribeBillingGroupOutput, error) {
    return a.client.DescribeBillingGroup(input)
}

func (a *IoTActivities) DescribeCACertificate(input *iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error) {
    return a.client.DescribeCACertificate(input)
}

func (a *IoTActivities) DescribeCertificate(input *iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error) {
    return a.client.DescribeCertificate(input)
}

func (a *IoTActivities) DescribeDefaultAuthorizer(input *iot.DescribeDefaultAuthorizerInput) (*iot.DescribeDefaultAuthorizerOutput, error) {
    return a.client.DescribeDefaultAuthorizer(input)
}

func (a *IoTActivities) DescribeDimension(input *iot.DescribeDimensionInput) (*iot.DescribeDimensionOutput, error) {
    return a.client.DescribeDimension(input)
}

func (a *IoTActivities) DescribeDomainConfiguration(input *iot.DescribeDomainConfigurationInput) (*iot.DescribeDomainConfigurationOutput, error) {
    return a.client.DescribeDomainConfiguration(input)
}

func (a *IoTActivities) DescribeEndpoint(input *iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error) {
    return a.client.DescribeEndpoint(input)
}

func (a *IoTActivities) DescribeEventConfigurations(input *iot.DescribeEventConfigurationsInput) (*iot.DescribeEventConfigurationsOutput, error) {
    return a.client.DescribeEventConfigurations(input)
}

func (a *IoTActivities) DescribeIndex(input *iot.DescribeIndexInput) (*iot.DescribeIndexOutput, error) {
    return a.client.DescribeIndex(input)
}

func (a *IoTActivities) DescribeJob(input *iot.DescribeJobInput) (*iot.DescribeJobOutput, error) {
    return a.client.DescribeJob(input)
}

func (a *IoTActivities) DescribeJobExecution(input *iot.DescribeJobExecutionInput) (*iot.DescribeJobExecutionOutput, error) {
    return a.client.DescribeJobExecution(input)
}

func (a *IoTActivities) DescribeMitigationAction(input *iot.DescribeMitigationActionInput) (*iot.DescribeMitigationActionOutput, error) {
    return a.client.DescribeMitigationAction(input)
}

func (a *IoTActivities) DescribeProvisioningTemplate(input *iot.DescribeProvisioningTemplateInput) (*iot.DescribeProvisioningTemplateOutput, error) {
    return a.client.DescribeProvisioningTemplate(input)
}

func (a *IoTActivities) DescribeProvisioningTemplateVersion(input *iot.DescribeProvisioningTemplateVersionInput) (*iot.DescribeProvisioningTemplateVersionOutput, error) {
    return a.client.DescribeProvisioningTemplateVersion(input)
}

func (a *IoTActivities) DescribeRoleAlias(input *iot.DescribeRoleAliasInput) (*iot.DescribeRoleAliasOutput, error) {
    return a.client.DescribeRoleAlias(input)
}

func (a *IoTActivities) DescribeScheduledAudit(input *iot.DescribeScheduledAuditInput) (*iot.DescribeScheduledAuditOutput, error) {
    return a.client.DescribeScheduledAudit(input)
}

func (a *IoTActivities) DescribeSecurityProfile(input *iot.DescribeSecurityProfileInput) (*iot.DescribeSecurityProfileOutput, error) {
    return a.client.DescribeSecurityProfile(input)
}

func (a *IoTActivities) DescribeStream(input *iot.DescribeStreamInput) (*iot.DescribeStreamOutput, error) {
    return a.client.DescribeStream(input)
}

func (a *IoTActivities) DescribeThing(input *iot.DescribeThingInput) (*iot.DescribeThingOutput, error) {
    return a.client.DescribeThing(input)
}

func (a *IoTActivities) DescribeThingGroup(input *iot.DescribeThingGroupInput) (*iot.DescribeThingGroupOutput, error) {
    return a.client.DescribeThingGroup(input)
}

func (a *IoTActivities) DescribeThingRegistrationTask(input *iot.DescribeThingRegistrationTaskInput) (*iot.DescribeThingRegistrationTaskOutput, error) {
    return a.client.DescribeThingRegistrationTask(input)
}

func (a *IoTActivities) DescribeThingType(input *iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error) {
    return a.client.DescribeThingType(input)
}

func (a *IoTActivities) DetachPolicy(input *iot.DetachPolicyInput) (*iot.DetachPolicyOutput, error) {
    return a.client.DetachPolicy(input)
}

func (a *IoTActivities) DetachPrincipalPolicy(input *iot.DetachPrincipalPolicyInput) (*iot.DetachPrincipalPolicyOutput, error) {
    return a.client.DetachPrincipalPolicy(input)
}

func (a *IoTActivities) DetachSecurityProfile(input *iot.DetachSecurityProfileInput) (*iot.DetachSecurityProfileOutput, error) {
    return a.client.DetachSecurityProfile(input)
}

func (a *IoTActivities) DetachThingPrincipal(input *iot.DetachThingPrincipalInput) (*iot.DetachThingPrincipalOutput, error) {
    return a.client.DetachThingPrincipal(input)
}

func (a *IoTActivities) DisableTopicRule(input *iot.DisableTopicRuleInput) (*iot.DisableTopicRuleOutput, error) {
    return a.client.DisableTopicRule(input)
}

func (a *IoTActivities) EnableTopicRule(input *iot.EnableTopicRuleInput) (*iot.EnableTopicRuleOutput, error) {
    return a.client.EnableTopicRule(input)
}

func (a *IoTActivities) GetCardinality(input *iot.GetCardinalityInput) (*iot.GetCardinalityOutput, error) {
    return a.client.GetCardinality(input)
}

func (a *IoTActivities) GetEffectivePolicies(input *iot.GetEffectivePoliciesInput) (*iot.GetEffectivePoliciesOutput, error) {
    return a.client.GetEffectivePolicies(input)
}

func (a *IoTActivities) GetIndexingConfiguration(input *iot.GetIndexingConfigurationInput) (*iot.GetIndexingConfigurationOutput, error) {
    return a.client.GetIndexingConfiguration(input)
}

func (a *IoTActivities) GetJobDocument(input *iot.GetJobDocumentInput) (*iot.GetJobDocumentOutput, error) {
    return a.client.GetJobDocument(input)
}

func (a *IoTActivities) GetLoggingOptions(input *iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error) {
    return a.client.GetLoggingOptions(input)
}

func (a *IoTActivities) GetOTAUpdate(input *iot.GetOTAUpdateInput) (*iot.GetOTAUpdateOutput, error) {
    return a.client.GetOTAUpdate(input)
}

func (a *IoTActivities) GetPercentiles(input *iot.GetPercentilesInput) (*iot.GetPercentilesOutput, error) {
    return a.client.GetPercentiles(input)
}

func (a *IoTActivities) GetPolicy(input *iot.GetPolicyInput) (*iot.GetPolicyOutput, error) {
    return a.client.GetPolicy(input)
}

func (a *IoTActivities) GetPolicyVersion(input *iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error) {
    return a.client.GetPolicyVersion(input)
}

func (a *IoTActivities) GetRegistrationCode(input *iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error) {
    return a.client.GetRegistrationCode(input)
}

func (a *IoTActivities) GetStatistics(input *iot.GetStatisticsInput) (*iot.GetStatisticsOutput, error) {
    return a.client.GetStatistics(input)
}

func (a *IoTActivities) GetTopicRule(input *iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error) {
    return a.client.GetTopicRule(input)
}

func (a *IoTActivities) GetTopicRuleDestination(input *iot.GetTopicRuleDestinationInput) (*iot.GetTopicRuleDestinationOutput, error) {
    return a.client.GetTopicRuleDestination(input)
}

func (a *IoTActivities) GetV2LoggingOptions(input *iot.GetV2LoggingOptionsInput) (*iot.GetV2LoggingOptionsOutput, error) {
    return a.client.GetV2LoggingOptions(input)
}

func (a *IoTActivities) ListActiveViolations(input *iot.ListActiveViolationsInput) (*iot.ListActiveViolationsOutput, error) {
    return a.client.ListActiveViolations(input)
}

func (a *IoTActivities) ListAttachedPolicies(input *iot.ListAttachedPoliciesInput) (*iot.ListAttachedPoliciesOutput, error) {
    return a.client.ListAttachedPolicies(input)
}

func (a *IoTActivities) ListAuditFindings(input *iot.ListAuditFindingsInput) (*iot.ListAuditFindingsOutput, error) {
    return a.client.ListAuditFindings(input)
}

func (a *IoTActivities) ListAuditMitigationActionsExecutions(input *iot.ListAuditMitigationActionsExecutionsInput) (*iot.ListAuditMitigationActionsExecutionsOutput, error) {
    return a.client.ListAuditMitigationActionsExecutions(input)
}

func (a *IoTActivities) ListAuditMitigationActionsTasks(input *iot.ListAuditMitigationActionsTasksInput) (*iot.ListAuditMitigationActionsTasksOutput, error) {
    return a.client.ListAuditMitigationActionsTasks(input)
}

func (a *IoTActivities) ListAuditSuppressions(input *iot.ListAuditSuppressionsInput) (*iot.ListAuditSuppressionsOutput, error) {
    return a.client.ListAuditSuppressions(input)
}

func (a *IoTActivities) ListAuditTasks(input *iot.ListAuditTasksInput) (*iot.ListAuditTasksOutput, error) {
    return a.client.ListAuditTasks(input)
}

func (a *IoTActivities) ListAuthorizers(input *iot.ListAuthorizersInput) (*iot.ListAuthorizersOutput, error) {
    return a.client.ListAuthorizers(input)
}

func (a *IoTActivities) ListBillingGroups(input *iot.ListBillingGroupsInput) (*iot.ListBillingGroupsOutput, error) {
    return a.client.ListBillingGroups(input)
}

func (a *IoTActivities) ListCACertificates(input *iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error) {
    return a.client.ListCACertificates(input)
}

func (a *IoTActivities) ListCertificates(input *iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error) {
    return a.client.ListCertificates(input)
}

func (a *IoTActivities) ListCertificatesByCA(input *iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error) {
    return a.client.ListCertificatesByCA(input)
}

func (a *IoTActivities) ListDimensions(input *iot.ListDimensionsInput) (*iot.ListDimensionsOutput, error) {
    return a.client.ListDimensions(input)
}

func (a *IoTActivities) ListDomainConfigurations(input *iot.ListDomainConfigurationsInput) (*iot.ListDomainConfigurationsOutput, error) {
    return a.client.ListDomainConfigurations(input)
}

func (a *IoTActivities) ListIndices(input *iot.ListIndicesInput) (*iot.ListIndicesOutput, error) {
    return a.client.ListIndices(input)
}

func (a *IoTActivities) ListJobExecutionsForJob(input *iot.ListJobExecutionsForJobInput) (*iot.ListJobExecutionsForJobOutput, error) {
    return a.client.ListJobExecutionsForJob(input)
}

func (a *IoTActivities) ListJobExecutionsForThing(input *iot.ListJobExecutionsForThingInput) (*iot.ListJobExecutionsForThingOutput, error) {
    return a.client.ListJobExecutionsForThing(input)
}

func (a *IoTActivities) ListJobs(input *iot.ListJobsInput) (*iot.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}

func (a *IoTActivities) ListMitigationActions(input *iot.ListMitigationActionsInput) (*iot.ListMitigationActionsOutput, error) {
    return a.client.ListMitigationActions(input)
}

func (a *IoTActivities) ListOTAUpdates(input *iot.ListOTAUpdatesInput) (*iot.ListOTAUpdatesOutput, error) {
    return a.client.ListOTAUpdates(input)
}

func (a *IoTActivities) ListOutgoingCertificates(input *iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error) {
    return a.client.ListOutgoingCertificates(input)
}

func (a *IoTActivities) ListPolicies(input *iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error) {
    return a.client.ListPolicies(input)
}

func (a *IoTActivities) ListPolicyPrincipals(input *iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error) {
    return a.client.ListPolicyPrincipals(input)
}

func (a *IoTActivities) ListPolicyVersions(input *iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error) {
    return a.client.ListPolicyVersions(input)
}

func (a *IoTActivities) ListPrincipalPolicies(input *iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error) {
    return a.client.ListPrincipalPolicies(input)
}

func (a *IoTActivities) ListPrincipalThings(input *iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error) {
    return a.client.ListPrincipalThings(input)
}

func (a *IoTActivities) ListProvisioningTemplateVersions(input *iot.ListProvisioningTemplateVersionsInput) (*iot.ListProvisioningTemplateVersionsOutput, error) {
    return a.client.ListProvisioningTemplateVersions(input)
}

func (a *IoTActivities) ListProvisioningTemplates(input *iot.ListProvisioningTemplatesInput) (*iot.ListProvisioningTemplatesOutput, error) {
    return a.client.ListProvisioningTemplates(input)
}

func (a *IoTActivities) ListRoleAliases(input *iot.ListRoleAliasesInput) (*iot.ListRoleAliasesOutput, error) {
    return a.client.ListRoleAliases(input)
}

func (a *IoTActivities) ListScheduledAudits(input *iot.ListScheduledAuditsInput) (*iot.ListScheduledAuditsOutput, error) {
    return a.client.ListScheduledAudits(input)
}

func (a *IoTActivities) ListSecurityProfiles(input *iot.ListSecurityProfilesInput) (*iot.ListSecurityProfilesOutput, error) {
    return a.client.ListSecurityProfiles(input)
}

func (a *IoTActivities) ListSecurityProfilesForTarget(input *iot.ListSecurityProfilesForTargetInput) (*iot.ListSecurityProfilesForTargetOutput, error) {
    return a.client.ListSecurityProfilesForTarget(input)
}

func (a *IoTActivities) ListStreams(input *iot.ListStreamsInput) (*iot.ListStreamsOutput, error) {
    return a.client.ListStreams(input)
}

func (a *IoTActivities) ListTagsForResource(input *iot.ListTagsForResourceInput) (*iot.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *IoTActivities) ListTargetsForPolicy(input *iot.ListTargetsForPolicyInput) (*iot.ListTargetsForPolicyOutput, error) {
    return a.client.ListTargetsForPolicy(input)
}

func (a *IoTActivities) ListTargetsForSecurityProfile(input *iot.ListTargetsForSecurityProfileInput) (*iot.ListTargetsForSecurityProfileOutput, error) {
    return a.client.ListTargetsForSecurityProfile(input)
}

func (a *IoTActivities) ListThingGroups(input *iot.ListThingGroupsInput) (*iot.ListThingGroupsOutput, error) {
    return a.client.ListThingGroups(input)
}

func (a *IoTActivities) ListThingGroupsForThing(input *iot.ListThingGroupsForThingInput) (*iot.ListThingGroupsForThingOutput, error) {
    return a.client.ListThingGroupsForThing(input)
}

func (a *IoTActivities) ListThingPrincipals(input *iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error) {
    return a.client.ListThingPrincipals(input)
}

func (a *IoTActivities) ListThingRegistrationTaskReports(input *iot.ListThingRegistrationTaskReportsInput) (*iot.ListThingRegistrationTaskReportsOutput, error) {
    return a.client.ListThingRegistrationTaskReports(input)
}

func (a *IoTActivities) ListThingRegistrationTasks(input *iot.ListThingRegistrationTasksInput) (*iot.ListThingRegistrationTasksOutput, error) {
    return a.client.ListThingRegistrationTasks(input)
}

func (a *IoTActivities) ListThingTypes(input *iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error) {
    return a.client.ListThingTypes(input)
}

func (a *IoTActivities) ListThings(input *iot.ListThingsInput) (*iot.ListThingsOutput, error) {
    return a.client.ListThings(input)
}

func (a *IoTActivities) ListThingsInBillingGroup(input *iot.ListThingsInBillingGroupInput) (*iot.ListThingsInBillingGroupOutput, error) {
    return a.client.ListThingsInBillingGroup(input)
}

func (a *IoTActivities) ListThingsInThingGroup(input *iot.ListThingsInThingGroupInput) (*iot.ListThingsInThingGroupOutput, error) {
    return a.client.ListThingsInThingGroup(input)
}

func (a *IoTActivities) ListTopicRuleDestinations(input *iot.ListTopicRuleDestinationsInput) (*iot.ListTopicRuleDestinationsOutput, error) {
    return a.client.ListTopicRuleDestinations(input)
}

func (a *IoTActivities) ListTopicRules(input *iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error) {
    return a.client.ListTopicRules(input)
}

func (a *IoTActivities) ListV2LoggingLevels(input *iot.ListV2LoggingLevelsInput) (*iot.ListV2LoggingLevelsOutput, error) {
    return a.client.ListV2LoggingLevels(input)
}

func (a *IoTActivities) ListViolationEvents(input *iot.ListViolationEventsInput) (*iot.ListViolationEventsOutput, error) {
    return a.client.ListViolationEvents(input)
}

func (a *IoTActivities) RegisterCACertificate(input *iot.RegisterCACertificateInput) (*iot.RegisterCACertificateOutput, error) {
    return a.client.RegisterCACertificate(input)
}

func (a *IoTActivities) RegisterCertificate(input *iot.RegisterCertificateInput) (*iot.RegisterCertificateOutput, error) {
    return a.client.RegisterCertificate(input)
}

func (a *IoTActivities) RegisterCertificateWithoutCA(input *iot.RegisterCertificateWithoutCAInput) (*iot.RegisterCertificateWithoutCAOutput, error) {
    return a.client.RegisterCertificateWithoutCA(input)
}

func (a *IoTActivities) RegisterThing(input *iot.RegisterThingInput) (*iot.RegisterThingOutput, error) {
    return a.client.RegisterThing(input)
}

func (a *IoTActivities) RejectCertificateTransfer(input *iot.RejectCertificateTransferInput) (*iot.RejectCertificateTransferOutput, error) {
    return a.client.RejectCertificateTransfer(input)
}

func (a *IoTActivities) RemoveThingFromBillingGroup(input *iot.RemoveThingFromBillingGroupInput) (*iot.RemoveThingFromBillingGroupOutput, error) {
    return a.client.RemoveThingFromBillingGroup(input)
}

func (a *IoTActivities) RemoveThingFromThingGroup(input *iot.RemoveThingFromThingGroupInput) (*iot.RemoveThingFromThingGroupOutput, error) {
    return a.client.RemoveThingFromThingGroup(input)
}

func (a *IoTActivities) ReplaceTopicRule(input *iot.ReplaceTopicRuleInput) (*iot.ReplaceTopicRuleOutput, error) {
    return a.client.ReplaceTopicRule(input)
}

func (a *IoTActivities) SearchIndex(input *iot.SearchIndexInput) (*iot.SearchIndexOutput, error) {
    return a.client.SearchIndex(input)
}

func (a *IoTActivities) SetDefaultAuthorizer(input *iot.SetDefaultAuthorizerInput) (*iot.SetDefaultAuthorizerOutput, error) {
    return a.client.SetDefaultAuthorizer(input)
}

func (a *IoTActivities) SetDefaultPolicyVersion(input *iot.SetDefaultPolicyVersionInput) (*iot.SetDefaultPolicyVersionOutput, error) {
    return a.client.SetDefaultPolicyVersion(input)
}

func (a *IoTActivities) SetLoggingOptions(input *iot.SetLoggingOptionsInput) (*iot.SetLoggingOptionsOutput, error) {
    return a.client.SetLoggingOptions(input)
}

func (a *IoTActivities) SetV2LoggingLevel(input *iot.SetV2LoggingLevelInput) (*iot.SetV2LoggingLevelOutput, error) {
    return a.client.SetV2LoggingLevel(input)
}

func (a *IoTActivities) SetV2LoggingOptions(input *iot.SetV2LoggingOptionsInput) (*iot.SetV2LoggingOptionsOutput, error) {
    return a.client.SetV2LoggingOptions(input)
}

func (a *IoTActivities) StartAuditMitigationActionsTask(input *iot.StartAuditMitigationActionsTaskInput) (*iot.StartAuditMitigationActionsTaskOutput, error) {
    return a.client.StartAuditMitigationActionsTask(input)
}

func (a *IoTActivities) StartOnDemandAuditTask(input *iot.StartOnDemandAuditTaskInput) (*iot.StartOnDemandAuditTaskOutput, error) {
    return a.client.StartOnDemandAuditTask(input)
}

func (a *IoTActivities) StartThingRegistrationTask(input *iot.StartThingRegistrationTaskInput) (*iot.StartThingRegistrationTaskOutput, error) {
    return a.client.StartThingRegistrationTask(input)
}

func (a *IoTActivities) StopThingRegistrationTask(input *iot.StopThingRegistrationTaskInput) (*iot.StopThingRegistrationTaskOutput, error) {
    return a.client.StopThingRegistrationTask(input)
}

func (a *IoTActivities) TagResource(input *iot.TagResourceInput) (*iot.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *IoTActivities) TestAuthorization(input *iot.TestAuthorizationInput) (*iot.TestAuthorizationOutput, error) {
    return a.client.TestAuthorization(input)
}

func (a *IoTActivities) TestInvokeAuthorizer(input *iot.TestInvokeAuthorizerInput) (*iot.TestInvokeAuthorizerOutput, error) {
    return a.client.TestInvokeAuthorizer(input)
}

func (a *IoTActivities) TransferCertificate(input *iot.TransferCertificateInput) (*iot.TransferCertificateOutput, error) {
    return a.client.TransferCertificate(input)
}

func (a *IoTActivities) UntagResource(input *iot.UntagResourceInput) (*iot.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *IoTActivities) UpdateAccountAuditConfiguration(input *iot.UpdateAccountAuditConfigurationInput) (*iot.UpdateAccountAuditConfigurationOutput, error) {
    return a.client.UpdateAccountAuditConfiguration(input)
}

func (a *IoTActivities) UpdateAuditSuppression(input *iot.UpdateAuditSuppressionInput) (*iot.UpdateAuditSuppressionOutput, error) {
    return a.client.UpdateAuditSuppression(input)
}

func (a *IoTActivities) UpdateAuthorizer(input *iot.UpdateAuthorizerInput) (*iot.UpdateAuthorizerOutput, error) {
    return a.client.UpdateAuthorizer(input)
}

func (a *IoTActivities) UpdateBillingGroup(input *iot.UpdateBillingGroupInput) (*iot.UpdateBillingGroupOutput, error) {
    return a.client.UpdateBillingGroup(input)
}

func (a *IoTActivities) UpdateCACertificate(input *iot.UpdateCACertificateInput) (*iot.UpdateCACertificateOutput, error) {
    return a.client.UpdateCACertificate(input)
}

func (a *IoTActivities) UpdateCertificate(input *iot.UpdateCertificateInput) (*iot.UpdateCertificateOutput, error) {
    return a.client.UpdateCertificate(input)
}

func (a *IoTActivities) UpdateDimension(input *iot.UpdateDimensionInput) (*iot.UpdateDimensionOutput, error) {
    return a.client.UpdateDimension(input)
}

func (a *IoTActivities) UpdateDomainConfiguration(input *iot.UpdateDomainConfigurationInput) (*iot.UpdateDomainConfigurationOutput, error) {
    return a.client.UpdateDomainConfiguration(input)
}

func (a *IoTActivities) UpdateDynamicThingGroup(input *iot.UpdateDynamicThingGroupInput) (*iot.UpdateDynamicThingGroupOutput, error) {
    return a.client.UpdateDynamicThingGroup(input)
}

func (a *IoTActivities) UpdateEventConfigurations(input *iot.UpdateEventConfigurationsInput) (*iot.UpdateEventConfigurationsOutput, error) {
    return a.client.UpdateEventConfigurations(input)
}

func (a *IoTActivities) UpdateIndexingConfiguration(input *iot.UpdateIndexingConfigurationInput) (*iot.UpdateIndexingConfigurationOutput, error) {
    return a.client.UpdateIndexingConfiguration(input)
}

func (a *IoTActivities) UpdateJob(input *iot.UpdateJobInput) (*iot.UpdateJobOutput, error) {
    return a.client.UpdateJob(input)
}

func (a *IoTActivities) UpdateMitigationAction(input *iot.UpdateMitigationActionInput) (*iot.UpdateMitigationActionOutput, error) {
    return a.client.UpdateMitigationAction(input)
}

func (a *IoTActivities) UpdateProvisioningTemplate(input *iot.UpdateProvisioningTemplateInput) (*iot.UpdateProvisioningTemplateOutput, error) {
    return a.client.UpdateProvisioningTemplate(input)
}

func (a *IoTActivities) UpdateRoleAlias(input *iot.UpdateRoleAliasInput) (*iot.UpdateRoleAliasOutput, error) {
    return a.client.UpdateRoleAlias(input)
}

func (a *IoTActivities) UpdateScheduledAudit(input *iot.UpdateScheduledAuditInput) (*iot.UpdateScheduledAuditOutput, error) {
    return a.client.UpdateScheduledAudit(input)
}

func (a *IoTActivities) UpdateSecurityProfile(input *iot.UpdateSecurityProfileInput) (*iot.UpdateSecurityProfileOutput, error) {
    return a.client.UpdateSecurityProfile(input)
}

func (a *IoTActivities) UpdateStream(input *iot.UpdateStreamInput) (*iot.UpdateStreamOutput, error) {
    return a.client.UpdateStream(input)
}

func (a *IoTActivities) UpdateThing(input *iot.UpdateThingInput) (*iot.UpdateThingOutput, error) {
    return a.client.UpdateThing(input)
}

func (a *IoTActivities) UpdateThingGroup(input *iot.UpdateThingGroupInput) (*iot.UpdateThingGroupOutput, error) {
    return a.client.UpdateThingGroup(input)
}

func (a *IoTActivities) UpdateThingGroupsForThing(input *iot.UpdateThingGroupsForThingInput) (*iot.UpdateThingGroupsForThingOutput, error) {
    return a.client.UpdateThingGroupsForThing(input)
}

func (a *IoTActivities) UpdateTopicRuleDestination(input *iot.UpdateTopicRuleDestinationInput) (*iot.UpdateTopicRuleDestinationOutput, error) {
    return a.client.UpdateTopicRuleDestination(input)
}

func (a *IoTActivities) ValidateSecurityProfileBehaviors(input *iot.ValidateSecurityProfileBehaviorsInput) (*iot.ValidateSecurityProfileBehaviorsOutput, error) {
    return a.client.ValidateSecurityProfileBehaviors(input)
}
