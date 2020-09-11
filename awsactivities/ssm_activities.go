package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SSMActivities struct {
	client ssmiface.SSMAPI
}

func NewSSMActivities(session *session.Session, config ...*aws.Config) *SSMActivities {
	client := ssm.New(session, config...)
	return &SSMActivities{client: client}
}

func (a *SSMActivities) AddTagsToResource(ctx context.Context, input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *SSMActivities) CancelCommand(ctx context.Context, input *ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error) {
	return a.client.CancelCommandWithContext(ctx, input)
}

func (a *SSMActivities) CancelMaintenanceWindowExecution(ctx context.Context, input *ssm.CancelMaintenanceWindowExecutionInput) (*ssm.CancelMaintenanceWindowExecutionOutput, error) {
	return a.client.CancelMaintenanceWindowExecutionWithContext(ctx, input)
}

func (a *SSMActivities) CreateActivation(ctx context.Context, input *ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error) {
	return a.client.CreateActivationWithContext(ctx, input)
}

func (a *SSMActivities) CreateAssociation(ctx context.Context, input *ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error) {
	return a.client.CreateAssociationWithContext(ctx, input)
}

func (a *SSMActivities) CreateAssociationBatch(ctx context.Context, input *ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error) {
	return a.client.CreateAssociationBatchWithContext(ctx, input)
}

func (a *SSMActivities) CreateDocument(ctx context.Context, input *ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error) {
	return a.client.CreateDocumentWithContext(ctx, input)
}

func (a *SSMActivities) CreateMaintenanceWindow(ctx context.Context, input *ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) CreateOpsItem(ctx context.Context, input *ssm.CreateOpsItemInput) (*ssm.CreateOpsItemOutput, error) {
	return a.client.CreateOpsItemWithContext(ctx, input)
}

func (a *SSMActivities) CreatePatchBaseline(ctx context.Context, input *ssm.CreatePatchBaselineInput) (*ssm.CreatePatchBaselineOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreatePatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) CreateResourceDataSync(ctx context.Context, input *ssm.CreateResourceDataSyncInput) (*ssm.CreateResourceDataSyncOutput, error) {
	return a.client.CreateResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) DeleteActivation(ctx context.Context, input *ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error) {
	return a.client.DeleteActivationWithContext(ctx, input)
}

func (a *SSMActivities) DeleteAssociation(ctx context.Context, input *ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error) {
	return a.client.DeleteAssociationWithContext(ctx, input)
}

func (a *SSMActivities) DeleteDocument(ctx context.Context, input *ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error) {
	return a.client.DeleteDocumentWithContext(ctx, input)
}

func (a *SSMActivities) DeleteInventory(ctx context.Context, input *ssm.DeleteInventoryInput) (*ssm.DeleteInventoryOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.DeleteInventoryWithContext(ctx, input)
}

func (a *SSMActivities) DeleteMaintenanceWindow(ctx context.Context, input *ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error) {
	return a.client.DeleteMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) DeleteParameter(ctx context.Context, input *ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error) {
	return a.client.DeleteParameterWithContext(ctx, input)
}

func (a *SSMActivities) DeleteParameters(ctx context.Context, input *ssm.DeleteParametersInput) (*ssm.DeleteParametersOutput, error) {
	return a.client.DeleteParametersWithContext(ctx, input)
}

func (a *SSMActivities) DeletePatchBaseline(ctx context.Context, input *ssm.DeletePatchBaselineInput) (*ssm.DeletePatchBaselineOutput, error) {
	return a.client.DeletePatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) DeleteResourceDataSync(ctx context.Context, input *ssm.DeleteResourceDataSyncInput) (*ssm.DeleteResourceDataSyncOutput, error) {
	return a.client.DeleteResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterManagedInstance(ctx context.Context, input *ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error) {
	return a.client.DeregisterManagedInstanceWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterPatchBaselineForPatchGroup(ctx context.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error) {
	return a.client.DeregisterPatchBaselineForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterTargetFromMaintenanceWindow(ctx context.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error) {
	return a.client.DeregisterTargetFromMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterTaskFromMaintenanceWindow(ctx context.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error) {
	return a.client.DeregisterTaskFromMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) DescribeActivations(ctx context.Context, input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error) {
	return a.client.DescribeActivationsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAssociation(ctx context.Context, input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error) {
	return a.client.DescribeAssociationWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAssociationExecutionTargets(ctx context.Context, input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
	return a.client.DescribeAssociationExecutionTargetsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAssociationExecutions(ctx context.Context, input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error) {
	return a.client.DescribeAssociationExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAutomationExecutions(ctx context.Context, input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error) {
	return a.client.DescribeAutomationExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAutomationStepExecutions(ctx context.Context, input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
	return a.client.DescribeAutomationStepExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAvailablePatches(ctx context.Context, input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error) {
	return a.client.DescribeAvailablePatchesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeDocument(ctx context.Context, input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error) {
	return a.client.DescribeDocumentWithContext(ctx, input)
}

func (a *SSMActivities) DescribeDocumentPermission(ctx context.Context, input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error) {
	return a.client.DescribeDocumentPermissionWithContext(ctx, input)
}

func (a *SSMActivities) DescribeEffectiveInstanceAssociations(ctx context.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
	return a.client.DescribeEffectiveInstanceAssociationsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeEffectivePatchesForPatchBaseline(ctx context.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
	return a.client.DescribeEffectivePatchesForPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstanceAssociationsStatus(ctx context.Context, input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
	return a.client.DescribeInstanceAssociationsStatusWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstanceInformation(ctx context.Context, input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
	return a.client.DescribeInstanceInformationWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstancePatchStates(ctx context.Context, input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error) {
	return a.client.DescribeInstancePatchStatesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstancePatchStatesForPatchGroup(ctx context.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
	return a.client.DescribeInstancePatchStatesForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstancePatches(ctx context.Context, input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error) {
	return a.client.DescribeInstancePatchesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInventoryDeletions(ctx context.Context, input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error) {
	return a.client.DescribeInventoryDeletionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowExecutionTaskInvocations(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	return a.client.DescribeMaintenanceWindowExecutionTaskInvocationsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowExecutionTasks(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
	return a.client.DescribeMaintenanceWindowExecutionTasksWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowExecutions(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
	return a.client.DescribeMaintenanceWindowExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowSchedule(ctx context.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
	return a.client.DescribeMaintenanceWindowScheduleWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowTargets(ctx context.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
	return a.client.DescribeMaintenanceWindowTargetsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowTasks(ctx context.Context, input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
	return a.client.DescribeMaintenanceWindowTasksWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindows(ctx context.Context, input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error) {
	return a.client.DescribeMaintenanceWindowsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowsForTarget(ctx context.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
	return a.client.DescribeMaintenanceWindowsForTargetWithContext(ctx, input)
}

func (a *SSMActivities) DescribeOpsItems(ctx context.Context, input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error) {
	return a.client.DescribeOpsItemsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeParameters(ctx context.Context, input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error) {
	return a.client.DescribeParametersWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchBaselines(ctx context.Context, input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error) {
	return a.client.DescribePatchBaselinesWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchGroupState(ctx context.Context, input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error) {
	return a.client.DescribePatchGroupStateWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchGroups(ctx context.Context, input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error) {
	return a.client.DescribePatchGroupsWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchProperties(ctx context.Context, input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error) {
	return a.client.DescribePatchPropertiesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeSessions(ctx context.Context, input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error) {
	return a.client.DescribeSessionsWithContext(ctx, input)
}

func (a *SSMActivities) GetAutomationExecution(ctx context.Context, input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error) {
	return a.client.GetAutomationExecutionWithContext(ctx, input)
}

func (a *SSMActivities) GetCalendarState(ctx context.Context, input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error) {
	return a.client.GetCalendarStateWithContext(ctx, input)
}

func (a *SSMActivities) GetCommandInvocation(ctx context.Context, input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error) {
	return a.client.GetCommandInvocationWithContext(ctx, input)
}

func (a *SSMActivities) GetConnectionStatus(ctx context.Context, input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error) {
	return a.client.GetConnectionStatusWithContext(ctx, input)
}

func (a *SSMActivities) GetDefaultPatchBaseline(ctx context.Context, input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error) {
	return a.client.GetDefaultPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) GetDeployablePatchSnapshotForInstance(ctx context.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
	return a.client.GetDeployablePatchSnapshotForInstanceWithContext(ctx, input)
}

func (a *SSMActivities) GetDocument(ctx context.Context, input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error) {
	return a.client.GetDocumentWithContext(ctx, input)
}

func (a *SSMActivities) GetInventory(ctx context.Context, input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error) {
	return a.client.GetInventoryWithContext(ctx, input)
}

func (a *SSMActivities) GetInventorySchema(ctx context.Context, input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error) {
	return a.client.GetInventorySchemaWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindow(ctx context.Context, input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error) {
	return a.client.GetMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowExecution(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
	return a.client.GetMaintenanceWindowExecutionWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowExecutionTask(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
	return a.client.GetMaintenanceWindowExecutionTaskWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	return a.client.GetMaintenanceWindowExecutionTaskInvocationWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowTask(ctx context.Context, input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error) {
	return a.client.GetMaintenanceWindowTaskWithContext(ctx, input)
}

func (a *SSMActivities) GetOpsItem(ctx context.Context, input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error) {
	return a.client.GetOpsItemWithContext(ctx, input)
}

func (a *SSMActivities) GetOpsSummary(ctx context.Context, input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error) {
	return a.client.GetOpsSummaryWithContext(ctx, input)
}

func (a *SSMActivities) GetParameter(ctx context.Context, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	return a.client.GetParameterWithContext(ctx, input)
}

func (a *SSMActivities) GetParameterHistory(ctx context.Context, input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error) {
	return a.client.GetParameterHistoryWithContext(ctx, input)
}

func (a *SSMActivities) GetParameters(ctx context.Context, input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error) {
	return a.client.GetParametersWithContext(ctx, input)
}

func (a *SSMActivities) GetParametersByPath(ctx context.Context, input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
	return a.client.GetParametersByPathWithContext(ctx, input)
}

func (a *SSMActivities) GetPatchBaseline(ctx context.Context, input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error) {
	return a.client.GetPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) GetPatchBaselineForPatchGroup(ctx context.Context, input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
	return a.client.GetPatchBaselineForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) GetServiceSetting(ctx context.Context, input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error) {
	return a.client.GetServiceSettingWithContext(ctx, input)
}

func (a *SSMActivities) LabelParameterVersion(ctx context.Context, input *ssm.LabelParameterVersionInput) (*ssm.LabelParameterVersionOutput, error) {
	return a.client.LabelParameterVersionWithContext(ctx, input)
}

func (a *SSMActivities) ListAssociationVersions(ctx context.Context, input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error) {
	return a.client.ListAssociationVersionsWithContext(ctx, input)
}

func (a *SSMActivities) ListAssociations(ctx context.Context, input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error) {
	return a.client.ListAssociationsWithContext(ctx, input)
}

func (a *SSMActivities) ListCommandInvocations(ctx context.Context, input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error) {
	return a.client.ListCommandInvocationsWithContext(ctx, input)
}

func (a *SSMActivities) ListCommands(ctx context.Context, input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error) {
	return a.client.ListCommandsWithContext(ctx, input)
}

func (a *SSMActivities) ListComplianceItems(ctx context.Context, input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error) {
	return a.client.ListComplianceItemsWithContext(ctx, input)
}

func (a *SSMActivities) ListComplianceSummaries(ctx context.Context, input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error) {
	return a.client.ListComplianceSummariesWithContext(ctx, input)
}

func (a *SSMActivities) ListDocumentVersions(ctx context.Context, input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error) {
	return a.client.ListDocumentVersionsWithContext(ctx, input)
}

func (a *SSMActivities) ListDocuments(ctx context.Context, input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error) {
	return a.client.ListDocumentsWithContext(ctx, input)
}

func (a *SSMActivities) ListInventoryEntries(ctx context.Context, input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error) {
	return a.client.ListInventoryEntriesWithContext(ctx, input)
}

func (a *SSMActivities) ListResourceComplianceSummaries(ctx context.Context, input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error) {
	return a.client.ListResourceComplianceSummariesWithContext(ctx, input)
}

func (a *SSMActivities) ListResourceDataSync(ctx context.Context, input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error) {
	return a.client.ListResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) ListTagsForResource(ctx context.Context, input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SSMActivities) ModifyDocumentPermission(ctx context.Context, input *ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error) {
	return a.client.ModifyDocumentPermissionWithContext(ctx, input)
}

func (a *SSMActivities) PutComplianceItems(ctx context.Context, input *ssm.PutComplianceItemsInput) (*ssm.PutComplianceItemsOutput, error) {
	return a.client.PutComplianceItemsWithContext(ctx, input)
}

func (a *SSMActivities) PutInventory(ctx context.Context, input *ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error) {
	return a.client.PutInventoryWithContext(ctx, input)
}

func (a *SSMActivities) PutParameter(ctx context.Context, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	return a.client.PutParameterWithContext(ctx, input)
}

func (a *SSMActivities) RegisterDefaultPatchBaseline(ctx context.Context, input *ssm.RegisterDefaultPatchBaselineInput) (*ssm.RegisterDefaultPatchBaselineOutput, error) {
	return a.client.RegisterDefaultPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) RegisterPatchBaselineForPatchGroup(ctx context.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error) {
	return a.client.RegisterPatchBaselineForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) RegisterTargetWithMaintenanceWindow(ctx context.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.RegisterTargetWithMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) RegisterTaskWithMaintenanceWindow(ctx context.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.RegisterTaskWithMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) RemoveTagsFromResource(ctx context.Context, input *ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *SSMActivities) ResetServiceSetting(ctx context.Context, input *ssm.ResetServiceSettingInput) (*ssm.ResetServiceSettingOutput, error) {
	return a.client.ResetServiceSettingWithContext(ctx, input)
}

func (a *SSMActivities) ResumeSession(ctx context.Context, input *ssm.ResumeSessionInput) (*ssm.ResumeSessionOutput, error) {
	return a.client.ResumeSessionWithContext(ctx, input)
}

func (a *SSMActivities) SendAutomationSignal(ctx context.Context, input *ssm.SendAutomationSignalInput) (*ssm.SendAutomationSignalOutput, error) {
	return a.client.SendAutomationSignalWithContext(ctx, input)
}

func (a *SSMActivities) SendCommand(ctx context.Context, input *ssm.SendCommandInput) (*ssm.SendCommandOutput, error) {
	return a.client.SendCommandWithContext(ctx, input)
}

func (a *SSMActivities) StartAssociationsOnce(ctx context.Context, input *ssm.StartAssociationsOnceInput) (*ssm.StartAssociationsOnceOutput, error) {
	return a.client.StartAssociationsOnceWithContext(ctx, input)
}

func (a *SSMActivities) StartAutomationExecution(ctx context.Context, input *ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.StartAutomationExecutionWithContext(ctx, input)
}

func (a *SSMActivities) StartSession(ctx context.Context, input *ssm.StartSessionInput) (*ssm.StartSessionOutput, error) {
	return a.client.StartSessionWithContext(ctx, input)
}

func (a *SSMActivities) StopAutomationExecution(ctx context.Context, input *ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error) {
	return a.client.StopAutomationExecutionWithContext(ctx, input)
}

func (a *SSMActivities) TerminateSession(ctx context.Context, input *ssm.TerminateSessionInput) (*ssm.TerminateSessionOutput, error) {
	return a.client.TerminateSessionWithContext(ctx, input)
}

func (a *SSMActivities) UpdateAssociation(ctx context.Context, input *ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error) {
	return a.client.UpdateAssociationWithContext(ctx, input)
}

func (a *SSMActivities) UpdateAssociationStatus(ctx context.Context, input *ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error) {
	return a.client.UpdateAssociationStatusWithContext(ctx, input)
}

func (a *SSMActivities) UpdateDocument(ctx context.Context, input *ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error) {
	return a.client.UpdateDocumentWithContext(ctx, input)
}

func (a *SSMActivities) UpdateDocumentDefaultVersion(ctx context.Context, input *ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error) {
	return a.client.UpdateDocumentDefaultVersionWithContext(ctx, input)
}

func (a *SSMActivities) UpdateMaintenanceWindow(ctx context.Context, input *ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error) {
	return a.client.UpdateMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) UpdateMaintenanceWindowTarget(ctx context.Context, input *ssm.UpdateMaintenanceWindowTargetInput) (*ssm.UpdateMaintenanceWindowTargetOutput, error) {
	return a.client.UpdateMaintenanceWindowTargetWithContext(ctx, input)
}

func (a *SSMActivities) UpdateMaintenanceWindowTask(ctx context.Context, input *ssm.UpdateMaintenanceWindowTaskInput) (*ssm.UpdateMaintenanceWindowTaskOutput, error) {
	return a.client.UpdateMaintenanceWindowTaskWithContext(ctx, input)
}

func (a *SSMActivities) UpdateManagedInstanceRole(ctx context.Context, input *ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error) {
	return a.client.UpdateManagedInstanceRoleWithContext(ctx, input)
}

func (a *SSMActivities) UpdateOpsItem(ctx context.Context, input *ssm.UpdateOpsItemInput) (*ssm.UpdateOpsItemOutput, error) {
	return a.client.UpdateOpsItemWithContext(ctx, input)
}

func (a *SSMActivities) UpdatePatchBaseline(ctx context.Context, input *ssm.UpdatePatchBaselineInput) (*ssm.UpdatePatchBaselineOutput, error) {
	return a.client.UpdatePatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) UpdateResourceDataSync(ctx context.Context, input *ssm.UpdateResourceDataSyncInput) (*ssm.UpdateResourceDataSyncOutput, error) {
	return a.client.UpdateResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) UpdateServiceSetting(ctx context.Context, input *ssm.UpdateServiceSettingInput) (*ssm.UpdateServiceSettingOutput, error) {
	return a.client.UpdateServiceSettingWithContext(ctx, input)
}

func (a *SSMActivities) WaitUntilCommandExecuted(ctx context.Context, input *ssm.GetCommandInvocationInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilCommandExecutedWithContext(ctx, input, options...)
	})
}
