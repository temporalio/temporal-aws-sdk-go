package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

type CloudFormationActivities struct {
	client cloudformationiface.CloudFormationAPI
}

func NewCloudFormationActivities(client cloudformationiface.CloudFormationAPI) *CloudFormationActivities {
    return &CloudFormationActivities{client: client}
}


func (a *CloudFormationActivities) CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
    return a.client.CancelUpdateStack(input)
}



func (a *CloudFormationActivities) ContinueUpdateRollback(input *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error) {
    return a.client.ContinueUpdateRollback(input)
}



func (a *CloudFormationActivities) CreateChangeSet(input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
    return a.client.CreateChangeSet(input)
}



func (a *CloudFormationActivities) CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
    return a.client.CreateStack(input)
}



func (a *CloudFormationActivities) CreateStackInstances(input *cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error) {
    return a.client.CreateStackInstances(input)
}



func (a *CloudFormationActivities) CreateStackSet(input *cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error) {
    return a.client.CreateStackSet(input)
}



func (a *CloudFormationActivities) DeleteChangeSet(input *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error) {
    return a.client.DeleteChangeSet(input)
}



func (a *CloudFormationActivities) DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
    return a.client.DeleteStack(input)
}



func (a *CloudFormationActivities) DeleteStackInstances(input *cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error) {
    return a.client.DeleteStackInstances(input)
}



func (a *CloudFormationActivities) DeleteStackSet(input *cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error) {
    return a.client.DeleteStackSet(input)
}



func (a *CloudFormationActivities) DeregisterType(input *cloudformation.DeregisterTypeInput) (*cloudformation.DeregisterTypeOutput, error) {
    return a.client.DeregisterType(input)
}



func (a *CloudFormationActivities) DescribeAccountLimits(input *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error) {
    return a.client.DescribeAccountLimits(input)
}



func (a *CloudFormationActivities) DescribeChangeSet(input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
    return a.client.DescribeChangeSet(input)
}



func (a *CloudFormationActivities) DescribeStackDriftDetectionStatus(input *cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
    return a.client.DescribeStackDriftDetectionStatus(input)
}



func (a *CloudFormationActivities) DescribeStackEvents(input *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error) {
    return a.client.DescribeStackEvents(input)
}



func (a *CloudFormationActivities) DescribeStackInstance(input *cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error) {
    return a.client.DescribeStackInstance(input)
}



func (a *CloudFormationActivities) DescribeStackResource(input *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error) {
    return a.client.DescribeStackResource(input)
}



func (a *CloudFormationActivities) DescribeStackResourceDrifts(input *cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
    return a.client.DescribeStackResourceDrifts(input)
}



func (a *CloudFormationActivities) DescribeStackResources(input *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error) {
    return a.client.DescribeStackResources(input)
}



func (a *CloudFormationActivities) DescribeStackSet(input *cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error) {
    return a.client.DescribeStackSet(input)
}



func (a *CloudFormationActivities) DescribeStackSetOperation(input *cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error) {
    return a.client.DescribeStackSetOperation(input)
}



func (a *CloudFormationActivities) DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
    return a.client.DescribeStacks(input)
}



func (a *CloudFormationActivities) DescribeType(input *cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error) {
    return a.client.DescribeType(input)
}



func (a *CloudFormationActivities) DescribeTypeRegistration(input *cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error) {
    return a.client.DescribeTypeRegistration(input)
}



func (a *CloudFormationActivities) DetectStackDrift(input *cloudformation.DetectStackDriftInput) (*cloudformation.DetectStackDriftOutput, error) {
    return a.client.DetectStackDrift(input)
}



func (a *CloudFormationActivities) DetectStackResourceDrift(input *cloudformation.DetectStackResourceDriftInput) (*cloudformation.DetectStackResourceDriftOutput, error) {
    return a.client.DetectStackResourceDrift(input)
}



func (a *CloudFormationActivities) DetectStackSetDrift(input *cloudformation.DetectStackSetDriftInput) (*cloudformation.DetectStackSetDriftOutput, error) {
    return a.client.DetectStackSetDrift(input)
}



func (a *CloudFormationActivities) EstimateTemplateCost(input *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error) {
    return a.client.EstimateTemplateCost(input)
}



func (a *CloudFormationActivities) ExecuteChangeSet(input *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error) {
    return a.client.ExecuteChangeSet(input)
}



func (a *CloudFormationActivities) GetStackPolicy(input *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error) {
    return a.client.GetStackPolicy(input)
}



func (a *CloudFormationActivities) GetTemplate(input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error) {
    return a.client.GetTemplate(input)
}



func (a *CloudFormationActivities) GetTemplateSummary(input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error) {
    return a.client.GetTemplateSummary(input)
}



func (a *CloudFormationActivities) ListChangeSets(input *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error) {
    return a.client.ListChangeSets(input)
}



func (a *CloudFormationActivities) ListExports(input *cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error) {
    return a.client.ListExports(input)
}



func (a *CloudFormationActivities) ListImports(input *cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error) {
    return a.client.ListImports(input)
}



func (a *CloudFormationActivities) ListStackInstances(input *cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error) {
    return a.client.ListStackInstances(input)
}



func (a *CloudFormationActivities) ListStackResources(input *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {
    return a.client.ListStackResources(input)
}



func (a *CloudFormationActivities) ListStackSetOperationResults(input *cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error) {
    return a.client.ListStackSetOperationResults(input)
}



func (a *CloudFormationActivities) ListStackSetOperations(input *cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error) {
    return a.client.ListStackSetOperations(input)
}



func (a *CloudFormationActivities) ListStackSets(input *cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error) {
    return a.client.ListStackSets(input)
}



func (a *CloudFormationActivities) ListStacks(input *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error) {
    return a.client.ListStacks(input)
}



func (a *CloudFormationActivities) ListTypeRegistrations(input *cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error) {
    return a.client.ListTypeRegistrations(input)
}



func (a *CloudFormationActivities) ListTypeVersions(input *cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error) {
    return a.client.ListTypeVersions(input)
}



func (a *CloudFormationActivities) ListTypes(input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error) {
    return a.client.ListTypes(input)
}



func (a *CloudFormationActivities) RecordHandlerProgress(input *cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error) {
    return a.client.RecordHandlerProgress(input)
}



func (a *CloudFormationActivities) RegisterType(input *cloudformation.RegisterTypeInput) (*cloudformation.RegisterTypeOutput, error) {
    return a.client.RegisterType(input)
}



func (a *CloudFormationActivities) SetStackPolicy(input *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error) {
    return a.client.SetStackPolicy(input)
}



func (a *CloudFormationActivities) SetTypeDefaultVersion(input *cloudformation.SetTypeDefaultVersionInput) (*cloudformation.SetTypeDefaultVersionOutput, error) {
    return a.client.SetTypeDefaultVersion(input)
}



func (a *CloudFormationActivities) SignalResource(input *cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error) {
    return a.client.SignalResource(input)
}



func (a *CloudFormationActivities) StopStackSetOperation(input *cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error) {
    return a.client.StopStackSetOperation(input)
}



func (a *CloudFormationActivities) UpdateStack(input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error) {
    return a.client.UpdateStack(input)
}



func (a *CloudFormationActivities) UpdateStackInstances(input *cloudformation.UpdateStackInstancesInput) (*cloudformation.UpdateStackInstancesOutput, error) {
    return a.client.UpdateStackInstances(input)
}



func (a *CloudFormationActivities) UpdateStackSet(input *cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error) {
    return a.client.UpdateStackSet(input)
}



func (a *CloudFormationActivities) UpdateTerminationProtection(input *cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error) {
    return a.client.UpdateTerminationProtection(input)
}



func (a *CloudFormationActivities) ValidateTemplate(input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {
    return a.client.ValidateTemplate(input)
}



func (a *CloudFormationActivities) WaitUntilChangeSetCreateComplete(input *cloudformation.DescribeChangeSetInput) error {
    return a.client.WaitUntilChangeSetCreateComplete(input)
}



func (a *CloudFormationActivities) WaitUntilStackCreateComplete(input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackCreateComplete(input)
}



func (a *CloudFormationActivities) WaitUntilStackDeleteComplete(input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackDeleteComplete(input)
}



func (a *CloudFormationActivities) WaitUntilStackExists(input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackExists(input)
}



func (a *CloudFormationActivities) WaitUntilStackImportComplete(input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackImportComplete(input)
}



func (a *CloudFormationActivities) WaitUntilStackRollbackComplete(input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackRollbackComplete(input)
}



func (a *CloudFormationActivities) WaitUntilStackUpdateComplete(input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackUpdateComplete(input)
}



func (a *CloudFormationActivities) WaitUntilTypeRegistrationComplete(input *cloudformation.DescribeTypeRegistrationInput) error {
    return a.client.WaitUntilTypeRegistrationComplete(input)
}

