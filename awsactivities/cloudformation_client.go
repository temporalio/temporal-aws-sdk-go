package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"go.temporal.io/sdk/workflow"
)

type CloudFormationClient interface {
    CancelUpdateStack(ctx workflow.Context, input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)
    CancelUpdateStackAsync(ctx workflow.Context, input *cloudformation.CancelUpdateStackInput) *CloudformationCancelUpdateStackResult

    ContinueUpdateRollback(ctx workflow.Context, input *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error)
    ContinueUpdateRollbackAsync(ctx workflow.Context, input *cloudformation.ContinueUpdateRollbackInput) *CloudformationContinueUpdateRollbackResult

    CreateChangeSet(ctx workflow.Context, input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error)
    CreateChangeSetAsync(ctx workflow.Context, input *cloudformation.CreateChangeSetInput) *CloudformationCreateChangeSetResult

    CreateStack(ctx workflow.Context, input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
    CreateStackAsync(ctx workflow.Context, input *cloudformation.CreateStackInput) *CloudformationCreateStackResult

    CreateStackInstances(ctx workflow.Context, input *cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error)
    CreateStackInstancesAsync(ctx workflow.Context, input *cloudformation.CreateStackInstancesInput) *CloudformationCreateStackInstancesResult

    CreateStackSet(ctx workflow.Context, input *cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error)
    CreateStackSetAsync(ctx workflow.Context, input *cloudformation.CreateStackSetInput) *CloudformationCreateStackSetResult

    DeleteChangeSet(ctx workflow.Context, input *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error)
    DeleteChangeSetAsync(ctx workflow.Context, input *cloudformation.DeleteChangeSetInput) *CloudformationDeleteChangeSetResult

    DeleteStack(ctx workflow.Context, input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)
    DeleteStackAsync(ctx workflow.Context, input *cloudformation.DeleteStackInput) *CloudformationDeleteStackResult

    DeleteStackInstances(ctx workflow.Context, input *cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error)
    DeleteStackInstancesAsync(ctx workflow.Context, input *cloudformation.DeleteStackInstancesInput) *CloudformationDeleteStackInstancesResult

    DeleteStackSet(ctx workflow.Context, input *cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error)
    DeleteStackSetAsync(ctx workflow.Context, input *cloudformation.DeleteStackSetInput) *CloudformationDeleteStackSetResult

    DeregisterType(ctx workflow.Context, input *cloudformation.DeregisterTypeInput) (*cloudformation.DeregisterTypeOutput, error)
    DeregisterTypeAsync(ctx workflow.Context, input *cloudformation.DeregisterTypeInput) *CloudformationDeregisterTypeResult

    DescribeAccountLimits(ctx workflow.Context, input *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error)
    DescribeAccountLimitsAsync(ctx workflow.Context, input *cloudformation.DescribeAccountLimitsInput) *CloudformationDescribeAccountLimitsResult

    DescribeChangeSet(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error)
    DescribeChangeSetAsync(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) *CloudformationDescribeChangeSetResult

    DescribeStackDriftDetectionStatus(ctx workflow.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error)
    DescribeStackDriftDetectionStatusAsync(ctx workflow.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput) *CloudformationDescribeStackDriftDetectionStatusResult

    DescribeStackEvents(ctx workflow.Context, input *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)
    DescribeStackEventsAsync(ctx workflow.Context, input *cloudformation.DescribeStackEventsInput) *CloudformationDescribeStackEventsResult

    DescribeStackInstance(ctx workflow.Context, input *cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error)
    DescribeStackInstanceAsync(ctx workflow.Context, input *cloudformation.DescribeStackInstanceInput) *CloudformationDescribeStackInstanceResult

    DescribeStackResource(ctx workflow.Context, input *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)
    DescribeStackResourceAsync(ctx workflow.Context, input *cloudformation.DescribeStackResourceInput) *CloudformationDescribeStackResourceResult

    DescribeStackResourceDrifts(ctx workflow.Context, input *cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error)
    DescribeStackResourceDriftsAsync(ctx workflow.Context, input *cloudformation.DescribeStackResourceDriftsInput) *CloudformationDescribeStackResourceDriftsResult

    DescribeStackResources(ctx workflow.Context, input *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)
    DescribeStackResourcesAsync(ctx workflow.Context, input *cloudformation.DescribeStackResourcesInput) *CloudformationDescribeStackResourcesResult

    DescribeStackSet(ctx workflow.Context, input *cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error)
    DescribeStackSetAsync(ctx workflow.Context, input *cloudformation.DescribeStackSetInput) *CloudformationDescribeStackSetResult

    DescribeStackSetOperation(ctx workflow.Context, input *cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error)
    DescribeStackSetOperationAsync(ctx workflow.Context, input *cloudformation.DescribeStackSetOperationInput) *CloudformationDescribeStackSetOperationResult

    DescribeStacks(ctx workflow.Context, input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
    DescribeStacksAsync(ctx workflow.Context, input *cloudformation.DescribeStacksInput) *CloudformationDescribeStacksResult

    DescribeType(ctx workflow.Context, input *cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error)
    DescribeTypeAsync(ctx workflow.Context, input *cloudformation.DescribeTypeInput) *CloudformationDescribeTypeResult

    DescribeTypeRegistration(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error)
    DescribeTypeRegistrationAsync(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) *CloudformationDescribeTypeRegistrationResult

    DetectStackDrift(ctx workflow.Context, input *cloudformation.DetectStackDriftInput) (*cloudformation.DetectStackDriftOutput, error)
    DetectStackDriftAsync(ctx workflow.Context, input *cloudformation.DetectStackDriftInput) *CloudformationDetectStackDriftResult

    DetectStackResourceDrift(ctx workflow.Context, input *cloudformation.DetectStackResourceDriftInput) (*cloudformation.DetectStackResourceDriftOutput, error)
    DetectStackResourceDriftAsync(ctx workflow.Context, input *cloudformation.DetectStackResourceDriftInput) *CloudformationDetectStackResourceDriftResult

    DetectStackSetDrift(ctx workflow.Context, input *cloudformation.DetectStackSetDriftInput) (*cloudformation.DetectStackSetDriftOutput, error)
    DetectStackSetDriftAsync(ctx workflow.Context, input *cloudformation.DetectStackSetDriftInput) *CloudformationDetectStackSetDriftResult

    EstimateTemplateCost(ctx workflow.Context, input *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)
    EstimateTemplateCostAsync(ctx workflow.Context, input *cloudformation.EstimateTemplateCostInput) *CloudformationEstimateTemplateCostResult

    ExecuteChangeSet(ctx workflow.Context, input *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error)
    ExecuteChangeSetAsync(ctx workflow.Context, input *cloudformation.ExecuteChangeSetInput) *CloudformationExecuteChangeSetResult

    GetStackPolicy(ctx workflow.Context, input *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)
    GetStackPolicyAsync(ctx workflow.Context, input *cloudformation.GetStackPolicyInput) *CloudformationGetStackPolicyResult

    GetTemplate(ctx workflow.Context, input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)
    GetTemplateAsync(ctx workflow.Context, input *cloudformation.GetTemplateInput) *CloudformationGetTemplateResult

    GetTemplateSummary(ctx workflow.Context, input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)
    GetTemplateSummaryAsync(ctx workflow.Context, input *cloudformation.GetTemplateSummaryInput) *CloudformationGetTemplateSummaryResult

    ListChangeSets(ctx workflow.Context, input *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error)
    ListChangeSetsAsync(ctx workflow.Context, input *cloudformation.ListChangeSetsInput) *CloudformationListChangeSetsResult

    ListExports(ctx workflow.Context, input *cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error)
    ListExportsAsync(ctx workflow.Context, input *cloudformation.ListExportsInput) *CloudformationListExportsResult

    ListImports(ctx workflow.Context, input *cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error)
    ListImportsAsync(ctx workflow.Context, input *cloudformation.ListImportsInput) *CloudformationListImportsResult

    ListStackInstances(ctx workflow.Context, input *cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error)
    ListStackInstancesAsync(ctx workflow.Context, input *cloudformation.ListStackInstancesInput) *CloudformationListStackInstancesResult

    ListStackResources(ctx workflow.Context, input *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)
    ListStackResourcesAsync(ctx workflow.Context, input *cloudformation.ListStackResourcesInput) *CloudformationListStackResourcesResult

    ListStackSetOperationResults(ctx workflow.Context, input *cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error)
    ListStackSetOperationResultsAsync(ctx workflow.Context, input *cloudformation.ListStackSetOperationResultsInput) *CloudformationListStackSetOperationResultsResult

    ListStackSetOperations(ctx workflow.Context, input *cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error)
    ListStackSetOperationsAsync(ctx workflow.Context, input *cloudformation.ListStackSetOperationsInput) *CloudformationListStackSetOperationsResult

    ListStackSets(ctx workflow.Context, input *cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error)
    ListStackSetsAsync(ctx workflow.Context, input *cloudformation.ListStackSetsInput) *CloudformationListStackSetsResult

    ListStacks(ctx workflow.Context, input *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)
    ListStacksAsync(ctx workflow.Context, input *cloudformation.ListStacksInput) *CloudformationListStacksResult

    ListTypeRegistrations(ctx workflow.Context, input *cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error)
    ListTypeRegistrationsAsync(ctx workflow.Context, input *cloudformation.ListTypeRegistrationsInput) *CloudformationListTypeRegistrationsResult

    ListTypeVersions(ctx workflow.Context, input *cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error)
    ListTypeVersionsAsync(ctx workflow.Context, input *cloudformation.ListTypeVersionsInput) *CloudformationListTypeVersionsResult

    ListTypes(ctx workflow.Context, input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error)
    ListTypesAsync(ctx workflow.Context, input *cloudformation.ListTypesInput) *CloudformationListTypesResult

    RecordHandlerProgress(ctx workflow.Context, input *cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error)
    RecordHandlerProgressAsync(ctx workflow.Context, input *cloudformation.RecordHandlerProgressInput) *CloudformationRecordHandlerProgressResult

    RegisterType(ctx workflow.Context, input *cloudformation.RegisterTypeInput) (*cloudformation.RegisterTypeOutput, error)
    RegisterTypeAsync(ctx workflow.Context, input *cloudformation.RegisterTypeInput) *CloudformationRegisterTypeResult

    SetStackPolicy(ctx workflow.Context, input *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)
    SetStackPolicyAsync(ctx workflow.Context, input *cloudformation.SetStackPolicyInput) *CloudformationSetStackPolicyResult

    SetTypeDefaultVersion(ctx workflow.Context, input *cloudformation.SetTypeDefaultVersionInput) (*cloudformation.SetTypeDefaultVersionOutput, error)
    SetTypeDefaultVersionAsync(ctx workflow.Context, input *cloudformation.SetTypeDefaultVersionInput) *CloudformationSetTypeDefaultVersionResult

    SignalResource(ctx workflow.Context, input *cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)
    SignalResourceAsync(ctx workflow.Context, input *cloudformation.SignalResourceInput) *CloudformationSignalResourceResult

    StopStackSetOperation(ctx workflow.Context, input *cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error)
    StopStackSetOperationAsync(ctx workflow.Context, input *cloudformation.StopStackSetOperationInput) *CloudformationStopStackSetOperationResult

    UpdateStack(ctx workflow.Context, input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)
    UpdateStackAsync(ctx workflow.Context, input *cloudformation.UpdateStackInput) *CloudformationUpdateStackResult

    UpdateStackInstances(ctx workflow.Context, input *cloudformation.UpdateStackInstancesInput) (*cloudformation.UpdateStackInstancesOutput, error)
    UpdateStackInstancesAsync(ctx workflow.Context, input *cloudformation.UpdateStackInstancesInput) *CloudformationUpdateStackInstancesResult

    UpdateStackSet(ctx workflow.Context, input *cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error)
    UpdateStackSetAsync(ctx workflow.Context, input *cloudformation.UpdateStackSetInput) *CloudformationUpdateStackSetResult

    UpdateTerminationProtection(ctx workflow.Context, input *cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error)
    UpdateTerminationProtectionAsync(ctx workflow.Context, input *cloudformation.UpdateTerminationProtectionInput) *CloudformationUpdateTerminationProtectionResult

    ValidateTemplate(ctx workflow.Context, input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
    ValidateTemplateAsync(ctx workflow.Context, input *cloudformation.ValidateTemplateInput) *CloudformationValidateTemplateResult

    WaitUntilChangeSetCreateComplete(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) error
    WaitUntilStackCreateComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
    WaitUntilStackDeleteComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
    WaitUntilStackExists(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
    WaitUntilStackImportComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
    WaitUntilStackRollbackComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
    WaitUntilStackUpdateComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error
    WaitUntilTypeRegistrationComplete(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) error}
type CloudformationCancelUpdateStackResult struct {
	Result workflow.Future
}

func (r *CloudformationCancelUpdateStackResult) Get(ctx workflow.Context) (*cloudformation.CancelUpdateStackOutput, error) {
    var output cloudformation.CancelUpdateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationContinueUpdateRollbackResult struct {
	Result workflow.Future
}

func (r *CloudformationContinueUpdateRollbackResult) Get(ctx workflow.Context) (*cloudformation.ContinueUpdateRollbackOutput, error) {
    var output cloudformation.ContinueUpdateRollbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationCreateChangeSetResult struct {
	Result workflow.Future
}

func (r *CloudformationCreateChangeSetResult) Get(ctx workflow.Context) (*cloudformation.CreateChangeSetOutput, error) {
    var output cloudformation.CreateChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationCreateStackResult struct {
	Result workflow.Future
}

func (r *CloudformationCreateStackResult) Get(ctx workflow.Context) (*cloudformation.CreateStackOutput, error) {
    var output cloudformation.CreateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationCreateStackInstancesResult struct {
	Result workflow.Future
}

func (r *CloudformationCreateStackInstancesResult) Get(ctx workflow.Context) (*cloudformation.CreateStackInstancesOutput, error) {
    var output cloudformation.CreateStackInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationCreateStackSetResult struct {
	Result workflow.Future
}

func (r *CloudformationCreateStackSetResult) Get(ctx workflow.Context) (*cloudformation.CreateStackSetOutput, error) {
    var output cloudformation.CreateStackSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDeleteChangeSetResult struct {
	Result workflow.Future
}

func (r *CloudformationDeleteChangeSetResult) Get(ctx workflow.Context) (*cloudformation.DeleteChangeSetOutput, error) {
    var output cloudformation.DeleteChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDeleteStackResult struct {
	Result workflow.Future
}

func (r *CloudformationDeleteStackResult) Get(ctx workflow.Context) (*cloudformation.DeleteStackOutput, error) {
    var output cloudformation.DeleteStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDeleteStackInstancesResult struct {
	Result workflow.Future
}

func (r *CloudformationDeleteStackInstancesResult) Get(ctx workflow.Context) (*cloudformation.DeleteStackInstancesOutput, error) {
    var output cloudformation.DeleteStackInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDeleteStackSetResult struct {
	Result workflow.Future
}

func (r *CloudformationDeleteStackSetResult) Get(ctx workflow.Context) (*cloudformation.DeleteStackSetOutput, error) {
    var output cloudformation.DeleteStackSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDeregisterTypeResult struct {
	Result workflow.Future
}

func (r *CloudformationDeregisterTypeResult) Get(ctx workflow.Context) (*cloudformation.DeregisterTypeOutput, error) {
    var output cloudformation.DeregisterTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeAccountLimitsResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeAccountLimitsResult) Get(ctx workflow.Context) (*cloudformation.DescribeAccountLimitsOutput, error) {
    var output cloudformation.DescribeAccountLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeChangeSetResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeChangeSetResult) Get(ctx workflow.Context) (*cloudformation.DescribeChangeSetOutput, error) {
    var output cloudformation.DescribeChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackDriftDetectionStatusResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackDriftDetectionStatusResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
    var output cloudformation.DescribeStackDriftDetectionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackEventsResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackEventsResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackEventsOutput, error) {
    var output cloudformation.DescribeStackEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackInstanceResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackInstanceResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackInstanceOutput, error) {
    var output cloudformation.DescribeStackInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackResourceResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackResourceResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackResourceOutput, error) {
    var output cloudformation.DescribeStackResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackResourceDriftsResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackResourceDriftsResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
    var output cloudformation.DescribeStackResourceDriftsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackResourcesResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackResourcesResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackResourcesOutput, error) {
    var output cloudformation.DescribeStackResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackSetResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackSetResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackSetOutput, error) {
    var output cloudformation.DescribeStackSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStackSetOperationResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStackSetOperationResult) Get(ctx workflow.Context) (*cloudformation.DescribeStackSetOperationOutput, error) {
    var output cloudformation.DescribeStackSetOperationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeStacksResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeStacksResult) Get(ctx workflow.Context) (*cloudformation.DescribeStacksOutput, error) {
    var output cloudformation.DescribeStacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeTypeResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeTypeResult) Get(ctx workflow.Context) (*cloudformation.DescribeTypeOutput, error) {
    var output cloudformation.DescribeTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDescribeTypeRegistrationResult struct {
	Result workflow.Future
}

func (r *CloudformationDescribeTypeRegistrationResult) Get(ctx workflow.Context) (*cloudformation.DescribeTypeRegistrationOutput, error) {
    var output cloudformation.DescribeTypeRegistrationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDetectStackDriftResult struct {
	Result workflow.Future
}

func (r *CloudformationDetectStackDriftResult) Get(ctx workflow.Context) (*cloudformation.DetectStackDriftOutput, error) {
    var output cloudformation.DetectStackDriftOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDetectStackResourceDriftResult struct {
	Result workflow.Future
}

func (r *CloudformationDetectStackResourceDriftResult) Get(ctx workflow.Context) (*cloudformation.DetectStackResourceDriftOutput, error) {
    var output cloudformation.DetectStackResourceDriftOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationDetectStackSetDriftResult struct {
	Result workflow.Future
}

func (r *CloudformationDetectStackSetDriftResult) Get(ctx workflow.Context) (*cloudformation.DetectStackSetDriftOutput, error) {
    var output cloudformation.DetectStackSetDriftOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationEstimateTemplateCostResult struct {
	Result workflow.Future
}

func (r *CloudformationEstimateTemplateCostResult) Get(ctx workflow.Context) (*cloudformation.EstimateTemplateCostOutput, error) {
    var output cloudformation.EstimateTemplateCostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationExecuteChangeSetResult struct {
	Result workflow.Future
}

func (r *CloudformationExecuteChangeSetResult) Get(ctx workflow.Context) (*cloudformation.ExecuteChangeSetOutput, error) {
    var output cloudformation.ExecuteChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationGetStackPolicyResult struct {
	Result workflow.Future
}

func (r *CloudformationGetStackPolicyResult) Get(ctx workflow.Context) (*cloudformation.GetStackPolicyOutput, error) {
    var output cloudformation.GetStackPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationGetTemplateResult struct {
	Result workflow.Future
}

func (r *CloudformationGetTemplateResult) Get(ctx workflow.Context) (*cloudformation.GetTemplateOutput, error) {
    var output cloudformation.GetTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationGetTemplateSummaryResult struct {
	Result workflow.Future
}

func (r *CloudformationGetTemplateSummaryResult) Get(ctx workflow.Context) (*cloudformation.GetTemplateSummaryOutput, error) {
    var output cloudformation.GetTemplateSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListChangeSetsResult struct {
	Result workflow.Future
}

func (r *CloudformationListChangeSetsResult) Get(ctx workflow.Context) (*cloudformation.ListChangeSetsOutput, error) {
    var output cloudformation.ListChangeSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListExportsResult struct {
	Result workflow.Future
}

func (r *CloudformationListExportsResult) Get(ctx workflow.Context) (*cloudformation.ListExportsOutput, error) {
    var output cloudformation.ListExportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListImportsResult struct {
	Result workflow.Future
}

func (r *CloudformationListImportsResult) Get(ctx workflow.Context) (*cloudformation.ListImportsOutput, error) {
    var output cloudformation.ListImportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListStackInstancesResult struct {
	Result workflow.Future
}

func (r *CloudformationListStackInstancesResult) Get(ctx workflow.Context) (*cloudformation.ListStackInstancesOutput, error) {
    var output cloudformation.ListStackInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListStackResourcesResult struct {
	Result workflow.Future
}

func (r *CloudformationListStackResourcesResult) Get(ctx workflow.Context) (*cloudformation.ListStackResourcesOutput, error) {
    var output cloudformation.ListStackResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListStackSetOperationResultsResult struct {
	Result workflow.Future
}

func (r *CloudformationListStackSetOperationResultsResult) Get(ctx workflow.Context) (*cloudformation.ListStackSetOperationResultsOutput, error) {
    var output cloudformation.ListStackSetOperationResultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListStackSetOperationsResult struct {
	Result workflow.Future
}

func (r *CloudformationListStackSetOperationsResult) Get(ctx workflow.Context) (*cloudformation.ListStackSetOperationsOutput, error) {
    var output cloudformation.ListStackSetOperationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListStackSetsResult struct {
	Result workflow.Future
}

func (r *CloudformationListStackSetsResult) Get(ctx workflow.Context) (*cloudformation.ListStackSetsOutput, error) {
    var output cloudformation.ListStackSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListStacksResult struct {
	Result workflow.Future
}

func (r *CloudformationListStacksResult) Get(ctx workflow.Context) (*cloudformation.ListStacksOutput, error) {
    var output cloudformation.ListStacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListTypeRegistrationsResult struct {
	Result workflow.Future
}

func (r *CloudformationListTypeRegistrationsResult) Get(ctx workflow.Context) (*cloudformation.ListTypeRegistrationsOutput, error) {
    var output cloudformation.ListTypeRegistrationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListTypeVersionsResult struct {
	Result workflow.Future
}

func (r *CloudformationListTypeVersionsResult) Get(ctx workflow.Context) (*cloudformation.ListTypeVersionsOutput, error) {
    var output cloudformation.ListTypeVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationListTypesResult struct {
	Result workflow.Future
}

func (r *CloudformationListTypesResult) Get(ctx workflow.Context) (*cloudformation.ListTypesOutput, error) {
    var output cloudformation.ListTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationRecordHandlerProgressResult struct {
	Result workflow.Future
}

func (r *CloudformationRecordHandlerProgressResult) Get(ctx workflow.Context) (*cloudformation.RecordHandlerProgressOutput, error) {
    var output cloudformation.RecordHandlerProgressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationRegisterTypeResult struct {
	Result workflow.Future
}

func (r *CloudformationRegisterTypeResult) Get(ctx workflow.Context) (*cloudformation.RegisterTypeOutput, error) {
    var output cloudformation.RegisterTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationSetStackPolicyResult struct {
	Result workflow.Future
}

func (r *CloudformationSetStackPolicyResult) Get(ctx workflow.Context) (*cloudformation.SetStackPolicyOutput, error) {
    var output cloudformation.SetStackPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationSetTypeDefaultVersionResult struct {
	Result workflow.Future
}

func (r *CloudformationSetTypeDefaultVersionResult) Get(ctx workflow.Context) (*cloudformation.SetTypeDefaultVersionOutput, error) {
    var output cloudformation.SetTypeDefaultVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationSignalResourceResult struct {
	Result workflow.Future
}

func (r *CloudformationSignalResourceResult) Get(ctx workflow.Context) (*cloudformation.SignalResourceOutput, error) {
    var output cloudformation.SignalResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationStopStackSetOperationResult struct {
	Result workflow.Future
}

func (r *CloudformationStopStackSetOperationResult) Get(ctx workflow.Context) (*cloudformation.StopStackSetOperationOutput, error) {
    var output cloudformation.StopStackSetOperationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationUpdateStackResult struct {
	Result workflow.Future
}

func (r *CloudformationUpdateStackResult) Get(ctx workflow.Context) (*cloudformation.UpdateStackOutput, error) {
    var output cloudformation.UpdateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationUpdateStackInstancesResult struct {
	Result workflow.Future
}

func (r *CloudformationUpdateStackInstancesResult) Get(ctx workflow.Context) (*cloudformation.UpdateStackInstancesOutput, error) {
    var output cloudformation.UpdateStackInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationUpdateStackSetResult struct {
	Result workflow.Future
}

func (r *CloudformationUpdateStackSetResult) Get(ctx workflow.Context) (*cloudformation.UpdateStackSetOutput, error) {
    var output cloudformation.UpdateStackSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationUpdateTerminationProtectionResult struct {
	Result workflow.Future
}

func (r *CloudformationUpdateTerminationProtectionResult) Get(ctx workflow.Context) (*cloudformation.UpdateTerminationProtectionOutput, error) {
    var output cloudformation.UpdateTerminationProtectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudformationValidateTemplateResult struct {
	Result workflow.Future
}

func (r *CloudformationValidateTemplateResult) Get(ctx workflow.Context) (*cloudformation.ValidateTemplateOutput, error) {
    var output cloudformation.ValidateTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudFormationStub struct {
    activities CloudFormationClient
}

func NewCloudFormationStub() CloudFormationClient {
    return &CloudFormationStub{}
}

func (a *CloudFormationStub) CancelUpdateStack(ctx workflow.Context, input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
    var output cloudformation.CancelUpdateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelUpdateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ContinueUpdateRollback(ctx workflow.Context, input *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error) {
    var output cloudformation.ContinueUpdateRollbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ContinueUpdateRollback, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) CreateChangeSet(ctx workflow.Context, input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
    var output cloudformation.CreateChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) CreateStack(ctx workflow.Context, input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
    var output cloudformation.CreateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) CreateStackInstances(ctx workflow.Context, input *cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error) {
    var output cloudformation.CreateStackInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStackInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) CreateStackSet(ctx workflow.Context, input *cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error) {
    var output cloudformation.CreateStackSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStackSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DeleteChangeSet(ctx workflow.Context, input *cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error) {
    var output cloudformation.DeleteChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DeleteStack(ctx workflow.Context, input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
    var output cloudformation.DeleteStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStack, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DeleteStackInstances(ctx workflow.Context, input *cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error) {
    var output cloudformation.DeleteStackInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStackInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DeleteStackSet(ctx workflow.Context, input *cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error) {
    var output cloudformation.DeleteStackSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStackSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DeregisterType(ctx workflow.Context, input *cloudformation.DeregisterTypeInput) (*cloudformation.DeregisterTypeOutput, error) {
    var output cloudformation.DeregisterTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterType, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeAccountLimits(ctx workflow.Context, input *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error) {
    var output cloudformation.DescribeAccountLimitsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountLimits, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeChangeSet(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
    var output cloudformation.DescribeChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackDriftDetectionStatus(ctx workflow.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
    var output cloudformation.DescribeStackDriftDetectionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackDriftDetectionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackEvents(ctx workflow.Context, input *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error) {
    var output cloudformation.DescribeStackEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackInstance(ctx workflow.Context, input *cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error) {
    var output cloudformation.DescribeStackInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackResource(ctx workflow.Context, input *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error) {
    var output cloudformation.DescribeStackResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackResourceDrifts(ctx workflow.Context, input *cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
    var output cloudformation.DescribeStackResourceDriftsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackResourceDrifts, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackResources(ctx workflow.Context, input *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error) {
    var output cloudformation.DescribeStackResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackResources, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackSet(ctx workflow.Context, input *cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error) {
    var output cloudformation.DescribeStackSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStackSetOperation(ctx workflow.Context, input *cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error) {
    var output cloudformation.DescribeStackSetOperationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackSetOperation, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeStacks(ctx workflow.Context, input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
    var output cloudformation.DescribeStacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStacks, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeType(ctx workflow.Context, input *cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error) {
    var output cloudformation.DescribeTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeType, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DescribeTypeRegistration(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error) {
    var output cloudformation.DescribeTypeRegistrationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTypeRegistration, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DetectStackDrift(ctx workflow.Context, input *cloudformation.DetectStackDriftInput) (*cloudformation.DetectStackDriftOutput, error) {
    var output cloudformation.DetectStackDriftOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectStackDrift, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DetectStackResourceDrift(ctx workflow.Context, input *cloudformation.DetectStackResourceDriftInput) (*cloudformation.DetectStackResourceDriftOutput, error) {
    var output cloudformation.DetectStackResourceDriftOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectStackResourceDrift, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) DetectStackSetDrift(ctx workflow.Context, input *cloudformation.DetectStackSetDriftInput) (*cloudformation.DetectStackSetDriftOutput, error) {
    var output cloudformation.DetectStackSetDriftOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectStackSetDrift, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) EstimateTemplateCost(ctx workflow.Context, input *cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error) {
    var output cloudformation.EstimateTemplateCostOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EstimateTemplateCost, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ExecuteChangeSet(ctx workflow.Context, input *cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error) {
    var output cloudformation.ExecuteChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExecuteChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) GetStackPolicy(ctx workflow.Context, input *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error) {
    var output cloudformation.GetStackPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetStackPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) GetTemplate(ctx workflow.Context, input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error) {
    var output cloudformation.GetTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) GetTemplateSummary(ctx workflow.Context, input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error) {
    var output cloudformation.GetTemplateSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTemplateSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListChangeSets(ctx workflow.Context, input *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error) {
    var output cloudformation.ListChangeSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChangeSets, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListExports(ctx workflow.Context, input *cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error) {
    var output cloudformation.ListExportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListExports, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListImports(ctx workflow.Context, input *cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error) {
    var output cloudformation.ListImportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImports, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListStackInstances(ctx workflow.Context, input *cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error) {
    var output cloudformation.ListStackInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStackInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListStackResources(ctx workflow.Context, input *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {
    var output cloudformation.ListStackResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStackResources, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListStackSetOperationResults(ctx workflow.Context, input *cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error) {
    var output cloudformation.ListStackSetOperationResultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStackSetOperationResults, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListStackSetOperations(ctx workflow.Context, input *cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error) {
    var output cloudformation.ListStackSetOperationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStackSetOperations, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListStackSets(ctx workflow.Context, input *cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error) {
    var output cloudformation.ListStackSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStackSets, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListStacks(ctx workflow.Context, input *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error) {
    var output cloudformation.ListStacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStacks, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListTypeRegistrations(ctx workflow.Context, input *cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error) {
    var output cloudformation.ListTypeRegistrationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTypeRegistrations, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListTypeVersions(ctx workflow.Context, input *cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error) {
    var output cloudformation.ListTypeVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTypeVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ListTypes(ctx workflow.Context, input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error) {
    var output cloudformation.ListTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) RecordHandlerProgress(ctx workflow.Context, input *cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error) {
    var output cloudformation.RecordHandlerProgressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RecordHandlerProgress, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) RegisterType(ctx workflow.Context, input *cloudformation.RegisterTypeInput) (*cloudformation.RegisterTypeOutput, error) {
    var output cloudformation.RegisterTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterType, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) SetStackPolicy(ctx workflow.Context, input *cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error) {
    var output cloudformation.SetStackPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetStackPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) SetTypeDefaultVersion(ctx workflow.Context, input *cloudformation.SetTypeDefaultVersionInput) (*cloudformation.SetTypeDefaultVersionOutput, error) {
    var output cloudformation.SetTypeDefaultVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetTypeDefaultVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) SignalResource(ctx workflow.Context, input *cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error) {
    var output cloudformation.SignalResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SignalResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) StopStackSetOperation(ctx workflow.Context, input *cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error) {
    var output cloudformation.StopStackSetOperationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopStackSetOperation, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) UpdateStack(ctx workflow.Context, input *cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error) {
    var output cloudformation.UpdateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) UpdateStackInstances(ctx workflow.Context, input *cloudformation.UpdateStackInstancesInput) (*cloudformation.UpdateStackInstancesOutput, error) {
    var output cloudformation.UpdateStackInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStackInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) UpdateStackSet(ctx workflow.Context, input *cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error) {
    var output cloudformation.UpdateStackSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStackSet, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) UpdateTerminationProtection(ctx workflow.Context, input *cloudformation.UpdateTerminationProtectionInput) (*cloudformation.UpdateTerminationProtectionOutput, error) {
    var output cloudformation.UpdateTerminationProtectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTerminationProtection, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudFormationStub) ValidateTemplate(ctx workflow.Context, input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {
    var output cloudformation.ValidateTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidateTemplate, input).Get(ctx, &output)
    return &output, err
}


func (a *CloudFormationStub) WaitUntilChangeSetCreateComplete(ctx workflow.Context, input *cloudformation.DescribeChangeSetInput) error {
    return a.client.WaitUntilChangeSetCreateComplete(input)
}


func (a *CloudFormationStub) WaitUntilStackCreateComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackCreateComplete(input)
}


func (a *CloudFormationStub) WaitUntilStackDeleteComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackDeleteComplete(input)
}


func (a *CloudFormationStub) WaitUntilStackExists(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackExists(input)
}


func (a *CloudFormationStub) WaitUntilStackImportComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackImportComplete(input)
}


func (a *CloudFormationStub) WaitUntilStackRollbackComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackRollbackComplete(input)
}


func (a *CloudFormationStub) WaitUntilStackUpdateComplete(ctx workflow.Context, input *cloudformation.DescribeStacksInput) error {
    return a.client.WaitUntilStackUpdateComplete(input)
}


func (a *CloudFormationStub) WaitUntilTypeRegistrationComplete(ctx workflow.Context, input *cloudformation.DescribeTypeRegistrationInput) error {
    return a.client.WaitUntilTypeRegistrationComplete(input)
}
