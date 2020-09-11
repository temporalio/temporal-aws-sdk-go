
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CloudFormationActivities struct {
    client cloudformationiface.CloudFormationAPI
}

func NewCloudFormationActivities(session *session.Session, config ...*aws.Config) *CloudFormationActivities {
    client := cloudformation.New(session, config...)
    return &CloudFormationActivities{client: client}
}

func (a *CloudFormationActivities) CancelUpdateStack(ctx context.Context, input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
    return a.client.CancelUpdateStackWithContext(ctx, input)
}

func (a *CloudFormationActivities) ContinueUpdateRollback(ctx context.Context, input *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error) {
    return a.client.ContinueUpdateRollbackWithContext(ctx, input)
}

func (a *CloudFormationActivities) CreateChangeSet(ctx context.Context, input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateChangeSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) CreateStack(ctx context.Context, input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
    return a.client.CreateStackWithContext(ctx, input)
}

func (a *CloudFormationActivities) CreateStackInstances(ctx context.Context, input *cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error) {
    return a.client.CreateStackInstancesWithContext(ctx, input)
}

func (a *CloudFormationActivities) CreateStackSet(ctx context.Context, input *cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error) {
    return a.client.CreateStackSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) DeleteChangeSet(ctx context.Context, input *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error) {
    return a.client.DeleteChangeSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) DeleteStack(ctx context.Context, input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
    return a.client.DeleteStackWithContext(ctx, input)
}

func (a *CloudFormationActivities) DeleteStackInstances(ctx context.Context, input *cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error) {
    return a.client.DeleteStackInstancesWithContext(ctx, input)
}

func (a *CloudFormationActivities) DeleteStackSet(ctx context.Context, input *cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error) {
    return a.client.DeleteStackSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) DeregisterType(ctx context.Context, input *cloudformation.DeregisterTypeInput) (*cloudformation.DeregisterTypeOutput, error) {
    return a.client.DeregisterTypeWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeAccountLimits(ctx context.Context, input *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error) {
    return a.client.DescribeAccountLimitsWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeChangeSet(ctx context.Context, input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
    return a.client.DescribeChangeSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackDriftDetectionStatus(ctx context.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
    return a.client.DescribeStackDriftDetectionStatusWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackEvents(ctx context.Context, input *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error) {
    return a.client.DescribeStackEventsWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackInstance(ctx context.Context, input *cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error) {
    return a.client.DescribeStackInstanceWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackResource(ctx context.Context, input *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error) {
    return a.client.DescribeStackResourceWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackResourceDrifts(ctx context.Context, input *cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
    return a.client.DescribeStackResourceDriftsWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackResources(ctx context.Context, input *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error) {
    return a.client.DescribeStackResourcesWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackSet(ctx context.Context, input *cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error) {
    return a.client.DescribeStackSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStackSetOperation(ctx context.Context, input *cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error) {
    return a.client.DescribeStackSetOperationWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeStacks(ctx context.Context, input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
    return a.client.DescribeStacksWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeType(ctx context.Context, input *cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error) {
    return a.client.DescribeTypeWithContext(ctx, input)
}

func (a *CloudFormationActivities) DescribeTypeRegistration(ctx context.Context, input *cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error) {
    return a.client.DescribeTypeRegistrationWithContext(ctx, input)
}

func (a *CloudFormationActivities) DetectStackDrift(ctx context.Context, input *cloudformation.DetectStackDriftInput) (*cloudformation.DetectStackDriftOutput, error) {
    return a.client.DetectStackDriftWithContext(ctx, input)
}

func (a *CloudFormationActivities) DetectStackResourceDrift(ctx context.Context, input *cloudformation.DetectStackResourceDriftInput) (*cloudformation.DetectStackResourceDriftOutput, error) {
    return a.client.DetectStackResourceDriftWithContext(ctx, input)
}

func (a *CloudFormationActivities) DetectStackSetDrift(ctx context.Context, input *cloudformation.DetectStackSetDriftInput) (*cloudformation.DetectStackSetDriftOutput, error) {
    return a.client.DetectStackSetDriftWithContext(ctx, input)
}

func (a *CloudFormationActivities) EstimateTemplateCost(ctx context.Context, input *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error) {
    return a.client.EstimateTemplateCostWithContext(ctx, input)
}

func (a *CloudFormationActivities) ExecuteChangeSet(ctx context.Context, input *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error) {
    return a.client.ExecuteChangeSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) GetStackPolicy(ctx context.Context, input *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error) {
    return a.client.GetStackPolicyWithContext(ctx, input)
}

func (a *CloudFormationActivities) GetTemplate(ctx context.Context, input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error) {
    return a.client.GetTemplateWithContext(ctx, input)
}

func (a *CloudFormationActivities) GetTemplateSummary(ctx context.Context, input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error) {
    return a.client.GetTemplateSummaryWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListChangeSets(ctx context.Context, input *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error) {
    return a.client.ListChangeSetsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListExports(ctx context.Context, input *cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error) {
    return a.client.ListExportsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListImports(ctx context.Context, input *cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error) {
    return a.client.ListImportsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListStackInstances(ctx context.Context, input *cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error) {
    return a.client.ListStackInstancesWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListStackResources(ctx context.Context, input *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {
    return a.client.ListStackResourcesWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListStackSetOperationResults(ctx context.Context, input *cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error) {
    return a.client.ListStackSetOperationResultsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListStackSetOperations(ctx context.Context, input *cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error) {
    return a.client.ListStackSetOperationsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListStackSets(ctx context.Context, input *cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error) {
    return a.client.ListStackSetsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListStacks(ctx context.Context, input *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error) {
    return a.client.ListStacksWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListTypeRegistrations(ctx context.Context, input *cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error) {
    return a.client.ListTypeRegistrationsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListTypeVersions(ctx context.Context, input *cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error) {
    return a.client.ListTypeVersionsWithContext(ctx, input)
}

func (a *CloudFormationActivities) ListTypes(ctx context.Context, input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error) {
    return a.client.ListTypesWithContext(ctx, input)
}

func (a *CloudFormationActivities) RecordHandlerProgress(ctx context.Context, input *cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error) {
    return a.client.RecordHandlerProgressWithContext(ctx, input)
}

func (a *CloudFormationActivities) RegisterType(ctx context.Context, input *cloudformation.RegisterTypeInput) (*cloudformation.RegisterTypeOutput, error) {
    return a.client.RegisterTypeWithContext(ctx, input)
}

func (a *CloudFormationActivities) SetStackPolicy(ctx context.Context, input *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error) {
    return a.client.SetStackPolicyWithContext(ctx, input)
}

func (a *CloudFormationActivities) SetTypeDefaultVersion(ctx context.Context, input *cloudformation.SetTypeDefaultVersionInput) (*cloudformation.SetTypeDefaultVersionOutput, error) {
    return a.client.SetTypeDefaultVersionWithContext(ctx, input)
}

func (a *CloudFormationActivities) SignalResource(ctx context.Context, input *cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error) {
    return a.client.SignalResourceWithContext(ctx, input)
}

func (a *CloudFormationActivities) StopStackSetOperation(ctx context.Context, input *cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error) {
    return a.client.StopStackSetOperationWithContext(ctx, input)
}

func (a *CloudFormationActivities) UpdateStack(ctx context.Context, input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error) {
    return a.client.UpdateStackWithContext(ctx, input)
}

func (a *CloudFormationActivities) UpdateStackInstances(ctx context.Context, input *cloudformation.UpdateStackInstancesInput) (*cloudformation.UpdateStackInstancesOutput, error) {
    return a.client.UpdateStackInstancesWithContext(ctx, input)
}

func (a *CloudFormationActivities) UpdateStackSet(ctx context.Context, input *cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error) {
    return a.client.UpdateStackSetWithContext(ctx, input)
}

func (a *CloudFormationActivities) UpdateTerminationProtection(ctx context.Context, input *cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error) {
    return a.client.UpdateTerminationProtectionWithContext(ctx, input)
}

func (a *CloudFormationActivities) ValidateTemplate(ctx context.Context, input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {
    return a.client.ValidateTemplateWithContext(ctx, input)
}

func (a *CloudFormationActivities) WaitUntilChangeSetCreateComplete(ctx context.Context, input *cloudformation.DescribeChangeSetInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilChangeSetCreateCompleteWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilStackCreateComplete(ctx context.Context, input *cloudformation.DescribeStacksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStackCreateCompleteWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilStackDeleteComplete(ctx context.Context, input *cloudformation.DescribeStacksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStackDeleteCompleteWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilStackExists(ctx context.Context, input *cloudformation.DescribeStacksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStackExistsWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilStackImportComplete(ctx context.Context, input *cloudformation.DescribeStacksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStackImportCompleteWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilStackRollbackComplete(ctx context.Context, input *cloudformation.DescribeStacksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStackRollbackCompleteWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilStackUpdateComplete(ctx context.Context, input *cloudformation.DescribeStacksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilStackUpdateCompleteWithContext(ctx, input, options...)
	})
}

func (a *CloudFormationActivities) WaitUntilTypeRegistrationComplete(ctx context.Context, input *cloudformation.DescribeTypeRegistrationInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilTypeRegistrationCompleteWithContext(ctx, input, options...)
	})
}
