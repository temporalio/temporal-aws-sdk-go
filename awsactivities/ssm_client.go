package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"go.temporal.io/sdk/workflow"
)

type SSMClient interface {
    AddTagsToResource(ctx workflow.Context, input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error)
    AddTagsToResourceAsync(ctx workflow.Context, input *ssm.AddTagsToResourceInput) *SsmAddTagsToResourceResult

    CancelCommand(ctx workflow.Context, input *ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error)
    CancelCommandAsync(ctx workflow.Context, input *ssm.CancelCommandInput) *SsmCancelCommandResult

    CancelMaintenanceWindowExecution(ctx workflow.Context, input *ssm.CancelMaintenanceWindowExecutionInput) (*ssm.CancelMaintenanceWindowExecutionOutput, error)
    CancelMaintenanceWindowExecutionAsync(ctx workflow.Context, input *ssm.CancelMaintenanceWindowExecutionInput) *SsmCancelMaintenanceWindowExecutionResult

    CreateActivation(ctx workflow.Context, input *ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error)
    CreateActivationAsync(ctx workflow.Context, input *ssm.CreateActivationInput) *SsmCreateActivationResult

    CreateAssociation(ctx workflow.Context, input *ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error)
    CreateAssociationAsync(ctx workflow.Context, input *ssm.CreateAssociationInput) *SsmCreateAssociationResult

    CreateAssociationBatch(ctx workflow.Context, input *ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error)
    CreateAssociationBatchAsync(ctx workflow.Context, input *ssm.CreateAssociationBatchInput) *SsmCreateAssociationBatchResult

    CreateDocument(ctx workflow.Context, input *ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error)
    CreateDocumentAsync(ctx workflow.Context, input *ssm.CreateDocumentInput) *SsmCreateDocumentResult

    CreateMaintenanceWindow(ctx workflow.Context, input *ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error)
    CreateMaintenanceWindowAsync(ctx workflow.Context, input *ssm.CreateMaintenanceWindowInput) *SsmCreateMaintenanceWindowResult

    CreateOpsItem(ctx workflow.Context, input *ssm.CreateOpsItemInput) (*ssm.CreateOpsItemOutput, error)
    CreateOpsItemAsync(ctx workflow.Context, input *ssm.CreateOpsItemInput) *SsmCreateOpsItemResult

    CreatePatchBaseline(ctx workflow.Context, input *ssm.CreatePatchBaselineInput) (*ssm.CreatePatchBaselineOutput, error)
    CreatePatchBaselineAsync(ctx workflow.Context, input *ssm.CreatePatchBaselineInput) *SsmCreatePatchBaselineResult

    CreateResourceDataSync(ctx workflow.Context, input *ssm.CreateResourceDataSyncInput) (*ssm.CreateResourceDataSyncOutput, error)
    CreateResourceDataSyncAsync(ctx workflow.Context, input *ssm.CreateResourceDataSyncInput) *SsmCreateResourceDataSyncResult

    DeleteActivation(ctx workflow.Context, input *ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error)
    DeleteActivationAsync(ctx workflow.Context, input *ssm.DeleteActivationInput) *SsmDeleteActivationResult

    DeleteAssociation(ctx workflow.Context, input *ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error)
    DeleteAssociationAsync(ctx workflow.Context, input *ssm.DeleteAssociationInput) *SsmDeleteAssociationResult

    DeleteDocument(ctx workflow.Context, input *ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error)
    DeleteDocumentAsync(ctx workflow.Context, input *ssm.DeleteDocumentInput) *SsmDeleteDocumentResult

    DeleteInventory(ctx workflow.Context, input *ssm.DeleteInventoryInput) (*ssm.DeleteInventoryOutput, error)
    DeleteInventoryAsync(ctx workflow.Context, input *ssm.DeleteInventoryInput) *SsmDeleteInventoryResult

    DeleteMaintenanceWindow(ctx workflow.Context, input *ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error)
    DeleteMaintenanceWindowAsync(ctx workflow.Context, input *ssm.DeleteMaintenanceWindowInput) *SsmDeleteMaintenanceWindowResult

    DeleteParameter(ctx workflow.Context, input *ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error)
    DeleteParameterAsync(ctx workflow.Context, input *ssm.DeleteParameterInput) *SsmDeleteParameterResult

    DeleteParameters(ctx workflow.Context, input *ssm.DeleteParametersInput) (*ssm.DeleteParametersOutput, error)
    DeleteParametersAsync(ctx workflow.Context, input *ssm.DeleteParametersInput) *SsmDeleteParametersResult

    DeletePatchBaseline(ctx workflow.Context, input *ssm.DeletePatchBaselineInput) (*ssm.DeletePatchBaselineOutput, error)
    DeletePatchBaselineAsync(ctx workflow.Context, input *ssm.DeletePatchBaselineInput) *SsmDeletePatchBaselineResult

    DeleteResourceDataSync(ctx workflow.Context, input *ssm.DeleteResourceDataSyncInput) (*ssm.DeleteResourceDataSyncOutput, error)
    DeleteResourceDataSyncAsync(ctx workflow.Context, input *ssm.DeleteResourceDataSyncInput) *SsmDeleteResourceDataSyncResult

    DeregisterManagedInstance(ctx workflow.Context, input *ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error)
    DeregisterManagedInstanceAsync(ctx workflow.Context, input *ssm.DeregisterManagedInstanceInput) *SsmDeregisterManagedInstanceResult

    DeregisterPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error)
    DeregisterPatchBaselineForPatchGroupAsync(ctx workflow.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) *SsmDeregisterPatchBaselineForPatchGroupResult

    DeregisterTargetFromMaintenanceWindow(ctx workflow.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error)
    DeregisterTargetFromMaintenanceWindowAsync(ctx workflow.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) *SsmDeregisterTargetFromMaintenanceWindowResult

    DeregisterTaskFromMaintenanceWindow(ctx workflow.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error)
    DeregisterTaskFromMaintenanceWindowAsync(ctx workflow.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) *SsmDeregisterTaskFromMaintenanceWindowResult

    DescribeActivations(ctx workflow.Context, input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error)
    DescribeActivationsAsync(ctx workflow.Context, input *ssm.DescribeActivationsInput) *SsmDescribeActivationsResult

    DescribeAssociation(ctx workflow.Context, input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error)
    DescribeAssociationAsync(ctx workflow.Context, input *ssm.DescribeAssociationInput) *SsmDescribeAssociationResult

    DescribeAssociationExecutionTargets(ctx workflow.Context, input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error)
    DescribeAssociationExecutionTargetsAsync(ctx workflow.Context, input *ssm.DescribeAssociationExecutionTargetsInput) *SsmDescribeAssociationExecutionTargetsResult

    DescribeAssociationExecutions(ctx workflow.Context, input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error)
    DescribeAssociationExecutionsAsync(ctx workflow.Context, input *ssm.DescribeAssociationExecutionsInput) *SsmDescribeAssociationExecutionsResult

    DescribeAutomationExecutions(ctx workflow.Context, input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error)
    DescribeAutomationExecutionsAsync(ctx workflow.Context, input *ssm.DescribeAutomationExecutionsInput) *SsmDescribeAutomationExecutionsResult

    DescribeAutomationStepExecutions(ctx workflow.Context, input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error)
    DescribeAutomationStepExecutionsAsync(ctx workflow.Context, input *ssm.DescribeAutomationStepExecutionsInput) *SsmDescribeAutomationStepExecutionsResult

    DescribeAvailablePatches(ctx workflow.Context, input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error)
    DescribeAvailablePatchesAsync(ctx workflow.Context, input *ssm.DescribeAvailablePatchesInput) *SsmDescribeAvailablePatchesResult

    DescribeDocument(ctx workflow.Context, input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error)
    DescribeDocumentAsync(ctx workflow.Context, input *ssm.DescribeDocumentInput) *SsmDescribeDocumentResult

    DescribeDocumentPermission(ctx workflow.Context, input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error)
    DescribeDocumentPermissionAsync(ctx workflow.Context, input *ssm.DescribeDocumentPermissionInput) *SsmDescribeDocumentPermissionResult

    DescribeEffectiveInstanceAssociations(ctx workflow.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error)
    DescribeEffectiveInstanceAssociationsAsync(ctx workflow.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) *SsmDescribeEffectiveInstanceAssociationsResult

    DescribeEffectivePatchesForPatchBaseline(ctx workflow.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error)
    DescribeEffectivePatchesForPatchBaselineAsync(ctx workflow.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) *SsmDescribeEffectivePatchesForPatchBaselineResult

    DescribeInstanceAssociationsStatus(ctx workflow.Context, input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error)
    DescribeInstanceAssociationsStatusAsync(ctx workflow.Context, input *ssm.DescribeInstanceAssociationsStatusInput) *SsmDescribeInstanceAssociationsStatusResult

    DescribeInstanceInformation(ctx workflow.Context, input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error)
    DescribeInstanceInformationAsync(ctx workflow.Context, input *ssm.DescribeInstanceInformationInput) *SsmDescribeInstanceInformationResult

    DescribeInstancePatchStates(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error)
    DescribeInstancePatchStatesAsync(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesInput) *SsmDescribeInstancePatchStatesResult

    DescribeInstancePatchStatesForPatchGroup(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error)
    DescribeInstancePatchStatesForPatchGroupAsync(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) *SsmDescribeInstancePatchStatesForPatchGroupResult

    DescribeInstancePatches(ctx workflow.Context, input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error)
    DescribeInstancePatchesAsync(ctx workflow.Context, input *ssm.DescribeInstancePatchesInput) *SsmDescribeInstancePatchesResult

    DescribeInventoryDeletions(ctx workflow.Context, input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error)
    DescribeInventoryDeletionsAsync(ctx workflow.Context, input *ssm.DescribeInventoryDeletionsInput) *SsmDescribeInventoryDeletionsResult

    DescribeMaintenanceWindowExecutionTaskInvocations(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error)
    DescribeMaintenanceWindowExecutionTaskInvocationsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) *SsmDescribeMaintenanceWindowExecutionTaskInvocationsResult

    DescribeMaintenanceWindowExecutionTasks(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error)
    DescribeMaintenanceWindowExecutionTasksAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) *SsmDescribeMaintenanceWindowExecutionTasksResult

    DescribeMaintenanceWindowExecutions(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error)
    DescribeMaintenanceWindowExecutionsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) *SsmDescribeMaintenanceWindowExecutionsResult

    DescribeMaintenanceWindowSchedule(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error)
    DescribeMaintenanceWindowScheduleAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) *SsmDescribeMaintenanceWindowScheduleResult

    DescribeMaintenanceWindowTargets(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error)
    DescribeMaintenanceWindowTargetsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) *SsmDescribeMaintenanceWindowTargetsResult

    DescribeMaintenanceWindowTasks(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error)
    DescribeMaintenanceWindowTasksAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTasksInput) *SsmDescribeMaintenanceWindowTasksResult

    DescribeMaintenanceWindows(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error)
    DescribeMaintenanceWindowsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsInput) *SsmDescribeMaintenanceWindowsResult

    DescribeMaintenanceWindowsForTarget(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error)
    DescribeMaintenanceWindowsForTargetAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) *SsmDescribeMaintenanceWindowsForTargetResult

    DescribeOpsItems(ctx workflow.Context, input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error)
    DescribeOpsItemsAsync(ctx workflow.Context, input *ssm.DescribeOpsItemsInput) *SsmDescribeOpsItemsResult

    DescribeParameters(ctx workflow.Context, input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error)
    DescribeParametersAsync(ctx workflow.Context, input *ssm.DescribeParametersInput) *SsmDescribeParametersResult

    DescribePatchBaselines(ctx workflow.Context, input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error)
    DescribePatchBaselinesAsync(ctx workflow.Context, input *ssm.DescribePatchBaselinesInput) *SsmDescribePatchBaselinesResult

    DescribePatchGroupState(ctx workflow.Context, input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error)
    DescribePatchGroupStateAsync(ctx workflow.Context, input *ssm.DescribePatchGroupStateInput) *SsmDescribePatchGroupStateResult

    DescribePatchGroups(ctx workflow.Context, input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error)
    DescribePatchGroupsAsync(ctx workflow.Context, input *ssm.DescribePatchGroupsInput) *SsmDescribePatchGroupsResult

    DescribePatchProperties(ctx workflow.Context, input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error)
    DescribePatchPropertiesAsync(ctx workflow.Context, input *ssm.DescribePatchPropertiesInput) *SsmDescribePatchPropertiesResult

    DescribeSessions(ctx workflow.Context, input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error)
    DescribeSessionsAsync(ctx workflow.Context, input *ssm.DescribeSessionsInput) *SsmDescribeSessionsResult

    GetAutomationExecution(ctx workflow.Context, input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error)
    GetAutomationExecutionAsync(ctx workflow.Context, input *ssm.GetAutomationExecutionInput) *SsmGetAutomationExecutionResult

    GetCalendarState(ctx workflow.Context, input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error)
    GetCalendarStateAsync(ctx workflow.Context, input *ssm.GetCalendarStateInput) *SsmGetCalendarStateResult

    GetCommandInvocation(ctx workflow.Context, input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error)
    GetCommandInvocationAsync(ctx workflow.Context, input *ssm.GetCommandInvocationInput) *SsmGetCommandInvocationResult

    GetConnectionStatus(ctx workflow.Context, input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error)
    GetConnectionStatusAsync(ctx workflow.Context, input *ssm.GetConnectionStatusInput) *SsmGetConnectionStatusResult

    GetDefaultPatchBaseline(ctx workflow.Context, input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error)
    GetDefaultPatchBaselineAsync(ctx workflow.Context, input *ssm.GetDefaultPatchBaselineInput) *SsmGetDefaultPatchBaselineResult

    GetDeployablePatchSnapshotForInstance(ctx workflow.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error)
    GetDeployablePatchSnapshotForInstanceAsync(ctx workflow.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) *SsmGetDeployablePatchSnapshotForInstanceResult

    GetDocument(ctx workflow.Context, input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error)
    GetDocumentAsync(ctx workflow.Context, input *ssm.GetDocumentInput) *SsmGetDocumentResult

    GetInventory(ctx workflow.Context, input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error)
    GetInventoryAsync(ctx workflow.Context, input *ssm.GetInventoryInput) *SsmGetInventoryResult

    GetInventorySchema(ctx workflow.Context, input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error)
    GetInventorySchemaAsync(ctx workflow.Context, input *ssm.GetInventorySchemaInput) *SsmGetInventorySchemaResult

    GetMaintenanceWindow(ctx workflow.Context, input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error)
    GetMaintenanceWindowAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowInput) *SsmGetMaintenanceWindowResult

    GetMaintenanceWindowExecution(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error)
    GetMaintenanceWindowExecutionAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionInput) *SsmGetMaintenanceWindowExecutionResult

    GetMaintenanceWindowExecutionTask(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error)
    GetMaintenanceWindowExecutionTaskAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) *SsmGetMaintenanceWindowExecutionTaskResult

    GetMaintenanceWindowExecutionTaskInvocation(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error)
    GetMaintenanceWindowExecutionTaskInvocationAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) *SsmGetMaintenanceWindowExecutionTaskInvocationResult

    GetMaintenanceWindowTask(ctx workflow.Context, input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error)
    GetMaintenanceWindowTaskAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowTaskInput) *SsmGetMaintenanceWindowTaskResult

    GetOpsItem(ctx workflow.Context, input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error)
    GetOpsItemAsync(ctx workflow.Context, input *ssm.GetOpsItemInput) *SsmGetOpsItemResult

    GetOpsSummary(ctx workflow.Context, input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error)
    GetOpsSummaryAsync(ctx workflow.Context, input *ssm.GetOpsSummaryInput) *SsmGetOpsSummaryResult

    GetParameter(ctx workflow.Context, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error)
    GetParameterAsync(ctx workflow.Context, input *ssm.GetParameterInput) *SsmGetParameterResult

    GetParameterHistory(ctx workflow.Context, input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error)
    GetParameterHistoryAsync(ctx workflow.Context, input *ssm.GetParameterHistoryInput) *SsmGetParameterHistoryResult

    GetParameters(ctx workflow.Context, input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error)
    GetParametersAsync(ctx workflow.Context, input *ssm.GetParametersInput) *SsmGetParametersResult

    GetParametersByPath(ctx workflow.Context, input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error)
    GetParametersByPathAsync(ctx workflow.Context, input *ssm.GetParametersByPathInput) *SsmGetParametersByPathResult

    GetPatchBaseline(ctx workflow.Context, input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error)
    GetPatchBaselineAsync(ctx workflow.Context, input *ssm.GetPatchBaselineInput) *SsmGetPatchBaselineResult

    GetPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error)
    GetPatchBaselineForPatchGroupAsync(ctx workflow.Context, input *ssm.GetPatchBaselineForPatchGroupInput) *SsmGetPatchBaselineForPatchGroupResult

    GetServiceSetting(ctx workflow.Context, input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error)
    GetServiceSettingAsync(ctx workflow.Context, input *ssm.GetServiceSettingInput) *SsmGetServiceSettingResult

    LabelParameterVersion(ctx workflow.Context, input *ssm.LabelParameterVersionInput) (*ssm.LabelParameterVersionOutput, error)
    LabelParameterVersionAsync(ctx workflow.Context, input *ssm.LabelParameterVersionInput) *SsmLabelParameterVersionResult

    ListAssociationVersions(ctx workflow.Context, input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error)
    ListAssociationVersionsAsync(ctx workflow.Context, input *ssm.ListAssociationVersionsInput) *SsmListAssociationVersionsResult

    ListAssociations(ctx workflow.Context, input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error)
    ListAssociationsAsync(ctx workflow.Context, input *ssm.ListAssociationsInput) *SsmListAssociationsResult

    ListCommandInvocations(ctx workflow.Context, input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error)
    ListCommandInvocationsAsync(ctx workflow.Context, input *ssm.ListCommandInvocationsInput) *SsmListCommandInvocationsResult

    ListCommands(ctx workflow.Context, input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error)
    ListCommandsAsync(ctx workflow.Context, input *ssm.ListCommandsInput) *SsmListCommandsResult

    ListComplianceItems(ctx workflow.Context, input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error)
    ListComplianceItemsAsync(ctx workflow.Context, input *ssm.ListComplianceItemsInput) *SsmListComplianceItemsResult

    ListComplianceSummaries(ctx workflow.Context, input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error)
    ListComplianceSummariesAsync(ctx workflow.Context, input *ssm.ListComplianceSummariesInput) *SsmListComplianceSummariesResult

    ListDocumentVersions(ctx workflow.Context, input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error)
    ListDocumentVersionsAsync(ctx workflow.Context, input *ssm.ListDocumentVersionsInput) *SsmListDocumentVersionsResult

    ListDocuments(ctx workflow.Context, input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error)
    ListDocumentsAsync(ctx workflow.Context, input *ssm.ListDocumentsInput) *SsmListDocumentsResult

    ListInventoryEntries(ctx workflow.Context, input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error)
    ListInventoryEntriesAsync(ctx workflow.Context, input *ssm.ListInventoryEntriesInput) *SsmListInventoryEntriesResult

    ListResourceComplianceSummaries(ctx workflow.Context, input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error)
    ListResourceComplianceSummariesAsync(ctx workflow.Context, input *ssm.ListResourceComplianceSummariesInput) *SsmListResourceComplianceSummariesResult

    ListResourceDataSync(ctx workflow.Context, input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error)
    ListResourceDataSyncAsync(ctx workflow.Context, input *ssm.ListResourceDataSyncInput) *SsmListResourceDataSyncResult

    ListTagsForResource(ctx workflow.Context, input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *ssm.ListTagsForResourceInput) *SsmListTagsForResourceResult

    ModifyDocumentPermission(ctx workflow.Context, input *ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error)
    ModifyDocumentPermissionAsync(ctx workflow.Context, input *ssm.ModifyDocumentPermissionInput) *SsmModifyDocumentPermissionResult

    PutComplianceItems(ctx workflow.Context, input *ssm.PutComplianceItemsInput) (*ssm.PutComplianceItemsOutput, error)
    PutComplianceItemsAsync(ctx workflow.Context, input *ssm.PutComplianceItemsInput) *SsmPutComplianceItemsResult

    PutInventory(ctx workflow.Context, input *ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error)
    PutInventoryAsync(ctx workflow.Context, input *ssm.PutInventoryInput) *SsmPutInventoryResult

    PutParameter(ctx workflow.Context, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error)
    PutParameterAsync(ctx workflow.Context, input *ssm.PutParameterInput) *SsmPutParameterResult

    RegisterDefaultPatchBaseline(ctx workflow.Context, input *ssm.RegisterDefaultPatchBaselineInput) (*ssm.RegisterDefaultPatchBaselineOutput, error)
    RegisterDefaultPatchBaselineAsync(ctx workflow.Context, input *ssm.RegisterDefaultPatchBaselineInput) *SsmRegisterDefaultPatchBaselineResult

    RegisterPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error)
    RegisterPatchBaselineForPatchGroupAsync(ctx workflow.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) *SsmRegisterPatchBaselineForPatchGroupResult

    RegisterTargetWithMaintenanceWindow(ctx workflow.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error)
    RegisterTargetWithMaintenanceWindowAsync(ctx workflow.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) *SsmRegisterTargetWithMaintenanceWindowResult

    RegisterTaskWithMaintenanceWindow(ctx workflow.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error)
    RegisterTaskWithMaintenanceWindowAsync(ctx workflow.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) *SsmRegisterTaskWithMaintenanceWindowResult

    RemoveTagsFromResource(ctx workflow.Context, input *ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error)
    RemoveTagsFromResourceAsync(ctx workflow.Context, input *ssm.RemoveTagsFromResourceInput) *SsmRemoveTagsFromResourceResult

    ResetServiceSetting(ctx workflow.Context, input *ssm.ResetServiceSettingInput) (*ssm.ResetServiceSettingOutput, error)
    ResetServiceSettingAsync(ctx workflow.Context, input *ssm.ResetServiceSettingInput) *SsmResetServiceSettingResult

    ResumeSession(ctx workflow.Context, input *ssm.ResumeSessionInput) (*ssm.ResumeSessionOutput, error)
    ResumeSessionAsync(ctx workflow.Context, input *ssm.ResumeSessionInput) *SsmResumeSessionResult

    SendAutomationSignal(ctx workflow.Context, input *ssm.SendAutomationSignalInput) (*ssm.SendAutomationSignalOutput, error)
    SendAutomationSignalAsync(ctx workflow.Context, input *ssm.SendAutomationSignalInput) *SsmSendAutomationSignalResult

    SendCommand(ctx workflow.Context, input *ssm.SendCommandInput) (*ssm.SendCommandOutput, error)
    SendCommandAsync(ctx workflow.Context, input *ssm.SendCommandInput) *SsmSendCommandResult

    StartAssociationsOnce(ctx workflow.Context, input *ssm.StartAssociationsOnceInput) (*ssm.StartAssociationsOnceOutput, error)
    StartAssociationsOnceAsync(ctx workflow.Context, input *ssm.StartAssociationsOnceInput) *SsmStartAssociationsOnceResult

    StartAutomationExecution(ctx workflow.Context, input *ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error)
    StartAutomationExecutionAsync(ctx workflow.Context, input *ssm.StartAutomationExecutionInput) *SsmStartAutomationExecutionResult

    StartSession(ctx workflow.Context, input *ssm.StartSessionInput) (*ssm.StartSessionOutput, error)
    StartSessionAsync(ctx workflow.Context, input *ssm.StartSessionInput) *SsmStartSessionResult

    StopAutomationExecution(ctx workflow.Context, input *ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error)
    StopAutomationExecutionAsync(ctx workflow.Context, input *ssm.StopAutomationExecutionInput) *SsmStopAutomationExecutionResult

    TerminateSession(ctx workflow.Context, input *ssm.TerminateSessionInput) (*ssm.TerminateSessionOutput, error)
    TerminateSessionAsync(ctx workflow.Context, input *ssm.TerminateSessionInput) *SsmTerminateSessionResult

    UpdateAssociation(ctx workflow.Context, input *ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error)
    UpdateAssociationAsync(ctx workflow.Context, input *ssm.UpdateAssociationInput) *SsmUpdateAssociationResult

    UpdateAssociationStatus(ctx workflow.Context, input *ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error)
    UpdateAssociationStatusAsync(ctx workflow.Context, input *ssm.UpdateAssociationStatusInput) *SsmUpdateAssociationStatusResult

    UpdateDocument(ctx workflow.Context, input *ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error)
    UpdateDocumentAsync(ctx workflow.Context, input *ssm.UpdateDocumentInput) *SsmUpdateDocumentResult

    UpdateDocumentDefaultVersion(ctx workflow.Context, input *ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error)
    UpdateDocumentDefaultVersionAsync(ctx workflow.Context, input *ssm.UpdateDocumentDefaultVersionInput) *SsmUpdateDocumentDefaultVersionResult

    UpdateMaintenanceWindow(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error)
    UpdateMaintenanceWindowAsync(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowInput) *SsmUpdateMaintenanceWindowResult

    UpdateMaintenanceWindowTarget(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTargetInput) (*ssm.UpdateMaintenanceWindowTargetOutput, error)
    UpdateMaintenanceWindowTargetAsync(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTargetInput) *SsmUpdateMaintenanceWindowTargetResult

    UpdateMaintenanceWindowTask(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTaskInput) (*ssm.UpdateMaintenanceWindowTaskOutput, error)
    UpdateMaintenanceWindowTaskAsync(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTaskInput) *SsmUpdateMaintenanceWindowTaskResult

    UpdateManagedInstanceRole(ctx workflow.Context, input *ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error)
    UpdateManagedInstanceRoleAsync(ctx workflow.Context, input *ssm.UpdateManagedInstanceRoleInput) *SsmUpdateManagedInstanceRoleResult

    UpdateOpsItem(ctx workflow.Context, input *ssm.UpdateOpsItemInput) (*ssm.UpdateOpsItemOutput, error)
    UpdateOpsItemAsync(ctx workflow.Context, input *ssm.UpdateOpsItemInput) *SsmUpdateOpsItemResult

    UpdatePatchBaseline(ctx workflow.Context, input *ssm.UpdatePatchBaselineInput) (*ssm.UpdatePatchBaselineOutput, error)
    UpdatePatchBaselineAsync(ctx workflow.Context, input *ssm.UpdatePatchBaselineInput) *SsmUpdatePatchBaselineResult

    UpdateResourceDataSync(ctx workflow.Context, input *ssm.UpdateResourceDataSyncInput) (*ssm.UpdateResourceDataSyncOutput, error)
    UpdateResourceDataSyncAsync(ctx workflow.Context, input *ssm.UpdateResourceDataSyncInput) *SsmUpdateResourceDataSyncResult

    UpdateServiceSetting(ctx workflow.Context, input *ssm.UpdateServiceSettingInput) (*ssm.UpdateServiceSettingOutput, error)
    UpdateServiceSettingAsync(ctx workflow.Context, input *ssm.UpdateServiceSettingInput) *SsmUpdateServiceSettingResult

    WaitUntilCommandExecuted(ctx workflow.Context, input *ssm.GetCommandInvocationInput) error}
type SsmAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *SsmAddTagsToResourceResult) Get(ctx workflow.Context) (*ssm.AddTagsToResourceOutput, error) {
    var output ssm.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCancelCommandResult struct {
	Result workflow.Future
}

func (r *SsmCancelCommandResult) Get(ctx workflow.Context) (*ssm.CancelCommandOutput, error) {
    var output ssm.CancelCommandOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCancelMaintenanceWindowExecutionResult struct {
	Result workflow.Future
}

func (r *SsmCancelMaintenanceWindowExecutionResult) Get(ctx workflow.Context) (*ssm.CancelMaintenanceWindowExecutionOutput, error) {
    var output ssm.CancelMaintenanceWindowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateActivationResult struct {
	Result workflow.Future
}

func (r *SsmCreateActivationResult) Get(ctx workflow.Context) (*ssm.CreateActivationOutput, error) {
    var output ssm.CreateActivationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateAssociationResult struct {
	Result workflow.Future
}

func (r *SsmCreateAssociationResult) Get(ctx workflow.Context) (*ssm.CreateAssociationOutput, error) {
    var output ssm.CreateAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateAssociationBatchResult struct {
	Result workflow.Future
}

func (r *SsmCreateAssociationBatchResult) Get(ctx workflow.Context) (*ssm.CreateAssociationBatchOutput, error) {
    var output ssm.CreateAssociationBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateDocumentResult struct {
	Result workflow.Future
}

func (r *SsmCreateDocumentResult) Get(ctx workflow.Context) (*ssm.CreateDocumentOutput, error) {
    var output ssm.CreateDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmCreateMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.CreateMaintenanceWindowOutput, error) {
    var output ssm.CreateMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateOpsItemResult struct {
	Result workflow.Future
}

func (r *SsmCreateOpsItemResult) Get(ctx workflow.Context) (*ssm.CreateOpsItemOutput, error) {
    var output ssm.CreateOpsItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreatePatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmCreatePatchBaselineResult) Get(ctx workflow.Context) (*ssm.CreatePatchBaselineOutput, error) {
    var output ssm.CreatePatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmCreateResourceDataSyncResult struct {
	Result workflow.Future
}

func (r *SsmCreateResourceDataSyncResult) Get(ctx workflow.Context) (*ssm.CreateResourceDataSyncOutput, error) {
    var output ssm.CreateResourceDataSyncOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteActivationResult struct {
	Result workflow.Future
}

func (r *SsmDeleteActivationResult) Get(ctx workflow.Context) (*ssm.DeleteActivationOutput, error) {
    var output ssm.DeleteActivationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteAssociationResult struct {
	Result workflow.Future
}

func (r *SsmDeleteAssociationResult) Get(ctx workflow.Context) (*ssm.DeleteAssociationOutput, error) {
    var output ssm.DeleteAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteDocumentResult struct {
	Result workflow.Future
}

func (r *SsmDeleteDocumentResult) Get(ctx workflow.Context) (*ssm.DeleteDocumentOutput, error) {
    var output ssm.DeleteDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteInventoryResult struct {
	Result workflow.Future
}

func (r *SsmDeleteInventoryResult) Get(ctx workflow.Context) (*ssm.DeleteInventoryOutput, error) {
    var output ssm.DeleteInventoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmDeleteMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.DeleteMaintenanceWindowOutput, error) {
    var output ssm.DeleteMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteParameterResult struct {
	Result workflow.Future
}

func (r *SsmDeleteParameterResult) Get(ctx workflow.Context) (*ssm.DeleteParameterOutput, error) {
    var output ssm.DeleteParameterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteParametersResult struct {
	Result workflow.Future
}

func (r *SsmDeleteParametersResult) Get(ctx workflow.Context) (*ssm.DeleteParametersOutput, error) {
    var output ssm.DeleteParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeletePatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmDeletePatchBaselineResult) Get(ctx workflow.Context) (*ssm.DeletePatchBaselineOutput, error) {
    var output ssm.DeletePatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeleteResourceDataSyncResult struct {
	Result workflow.Future
}

func (r *SsmDeleteResourceDataSyncResult) Get(ctx workflow.Context) (*ssm.DeleteResourceDataSyncOutput, error) {
    var output ssm.DeleteResourceDataSyncOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeregisterManagedInstanceResult struct {
	Result workflow.Future
}

func (r *SsmDeregisterManagedInstanceResult) Get(ctx workflow.Context) (*ssm.DeregisterManagedInstanceOutput, error) {
    var output ssm.DeregisterManagedInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeregisterPatchBaselineForPatchGroupResult struct {
	Result workflow.Future
}

func (r *SsmDeregisterPatchBaselineForPatchGroupResult) Get(ctx workflow.Context) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error) {
    var output ssm.DeregisterPatchBaselineForPatchGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeregisterTargetFromMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmDeregisterTargetFromMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error) {
    var output ssm.DeregisterTargetFromMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDeregisterTaskFromMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmDeregisterTaskFromMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error) {
    var output ssm.DeregisterTaskFromMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeActivationsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeActivationsResult) Get(ctx workflow.Context) (*ssm.DescribeActivationsOutput, error) {
    var output ssm.DescribeActivationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeAssociationResult struct {
	Result workflow.Future
}

func (r *SsmDescribeAssociationResult) Get(ctx workflow.Context) (*ssm.DescribeAssociationOutput, error) {
    var output ssm.DescribeAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeAssociationExecutionTargetsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeAssociationExecutionTargetsResult) Get(ctx workflow.Context) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
    var output ssm.DescribeAssociationExecutionTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeAssociationExecutionsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeAssociationExecutionsResult) Get(ctx workflow.Context) (*ssm.DescribeAssociationExecutionsOutput, error) {
    var output ssm.DescribeAssociationExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeAutomationExecutionsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeAutomationExecutionsResult) Get(ctx workflow.Context) (*ssm.DescribeAutomationExecutionsOutput, error) {
    var output ssm.DescribeAutomationExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeAutomationStepExecutionsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeAutomationStepExecutionsResult) Get(ctx workflow.Context) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
    var output ssm.DescribeAutomationStepExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeAvailablePatchesResult struct {
	Result workflow.Future
}

func (r *SsmDescribeAvailablePatchesResult) Get(ctx workflow.Context) (*ssm.DescribeAvailablePatchesOutput, error) {
    var output ssm.DescribeAvailablePatchesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeDocumentResult struct {
	Result workflow.Future
}

func (r *SsmDescribeDocumentResult) Get(ctx workflow.Context) (*ssm.DescribeDocumentOutput, error) {
    var output ssm.DescribeDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeDocumentPermissionResult struct {
	Result workflow.Future
}

func (r *SsmDescribeDocumentPermissionResult) Get(ctx workflow.Context) (*ssm.DescribeDocumentPermissionOutput, error) {
    var output ssm.DescribeDocumentPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeEffectiveInstanceAssociationsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeEffectiveInstanceAssociationsResult) Get(ctx workflow.Context) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
    var output ssm.DescribeEffectiveInstanceAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeEffectivePatchesForPatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmDescribeEffectivePatchesForPatchBaselineResult) Get(ctx workflow.Context) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
    var output ssm.DescribeEffectivePatchesForPatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeInstanceAssociationsStatusResult struct {
	Result workflow.Future
}

func (r *SsmDescribeInstanceAssociationsStatusResult) Get(ctx workflow.Context) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
    var output ssm.DescribeInstanceAssociationsStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeInstanceInformationResult struct {
	Result workflow.Future
}

func (r *SsmDescribeInstanceInformationResult) Get(ctx workflow.Context) (*ssm.DescribeInstanceInformationOutput, error) {
    var output ssm.DescribeInstanceInformationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeInstancePatchStatesResult struct {
	Result workflow.Future
}

func (r *SsmDescribeInstancePatchStatesResult) Get(ctx workflow.Context) (*ssm.DescribeInstancePatchStatesOutput, error) {
    var output ssm.DescribeInstancePatchStatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeInstancePatchStatesForPatchGroupResult struct {
	Result workflow.Future
}

func (r *SsmDescribeInstancePatchStatesForPatchGroupResult) Get(ctx workflow.Context) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
    var output ssm.DescribeInstancePatchStatesForPatchGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeInstancePatchesResult struct {
	Result workflow.Future
}

func (r *SsmDescribeInstancePatchesResult) Get(ctx workflow.Context) (*ssm.DescribeInstancePatchesOutput, error) {
    var output ssm.DescribeInstancePatchesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeInventoryDeletionsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeInventoryDeletionsResult) Get(ctx workflow.Context) (*ssm.DescribeInventoryDeletionsOutput, error) {
    var output ssm.DescribeInventoryDeletionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowExecutionTaskInvocationsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowExecutionTaskInvocationsResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
    var output ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowExecutionTasksResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowExecutionTasksResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
    var output ssm.DescribeMaintenanceWindowExecutionTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowExecutionsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowExecutionsResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
    var output ssm.DescribeMaintenanceWindowExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowScheduleResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowScheduleResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
    var output ssm.DescribeMaintenanceWindowScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowTargetsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowTargetsResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
    var output ssm.DescribeMaintenanceWindowTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowTasksResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowTasksResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
    var output ssm.DescribeMaintenanceWindowTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowsResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowsOutput, error) {
    var output ssm.DescribeMaintenanceWindowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeMaintenanceWindowsForTargetResult struct {
	Result workflow.Future
}

func (r *SsmDescribeMaintenanceWindowsForTargetResult) Get(ctx workflow.Context) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
    var output ssm.DescribeMaintenanceWindowsForTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeOpsItemsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeOpsItemsResult) Get(ctx workflow.Context) (*ssm.DescribeOpsItemsOutput, error) {
    var output ssm.DescribeOpsItemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeParametersResult struct {
	Result workflow.Future
}

func (r *SsmDescribeParametersResult) Get(ctx workflow.Context) (*ssm.DescribeParametersOutput, error) {
    var output ssm.DescribeParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribePatchBaselinesResult struct {
	Result workflow.Future
}

func (r *SsmDescribePatchBaselinesResult) Get(ctx workflow.Context) (*ssm.DescribePatchBaselinesOutput, error) {
    var output ssm.DescribePatchBaselinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribePatchGroupStateResult struct {
	Result workflow.Future
}

func (r *SsmDescribePatchGroupStateResult) Get(ctx workflow.Context) (*ssm.DescribePatchGroupStateOutput, error) {
    var output ssm.DescribePatchGroupStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribePatchGroupsResult struct {
	Result workflow.Future
}

func (r *SsmDescribePatchGroupsResult) Get(ctx workflow.Context) (*ssm.DescribePatchGroupsOutput, error) {
    var output ssm.DescribePatchGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribePatchPropertiesResult struct {
	Result workflow.Future
}

func (r *SsmDescribePatchPropertiesResult) Get(ctx workflow.Context) (*ssm.DescribePatchPropertiesOutput, error) {
    var output ssm.DescribePatchPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmDescribeSessionsResult struct {
	Result workflow.Future
}

func (r *SsmDescribeSessionsResult) Get(ctx workflow.Context) (*ssm.DescribeSessionsOutput, error) {
    var output ssm.DescribeSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetAutomationExecutionResult struct {
	Result workflow.Future
}

func (r *SsmGetAutomationExecutionResult) Get(ctx workflow.Context) (*ssm.GetAutomationExecutionOutput, error) {
    var output ssm.GetAutomationExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetCalendarStateResult struct {
	Result workflow.Future
}

func (r *SsmGetCalendarStateResult) Get(ctx workflow.Context) (*ssm.GetCalendarStateOutput, error) {
    var output ssm.GetCalendarStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetCommandInvocationResult struct {
	Result workflow.Future
}

func (r *SsmGetCommandInvocationResult) Get(ctx workflow.Context) (*ssm.GetCommandInvocationOutput, error) {
    var output ssm.GetCommandInvocationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetConnectionStatusResult struct {
	Result workflow.Future
}

func (r *SsmGetConnectionStatusResult) Get(ctx workflow.Context) (*ssm.GetConnectionStatusOutput, error) {
    var output ssm.GetConnectionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetDefaultPatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmGetDefaultPatchBaselineResult) Get(ctx workflow.Context) (*ssm.GetDefaultPatchBaselineOutput, error) {
    var output ssm.GetDefaultPatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetDeployablePatchSnapshotForInstanceResult struct {
	Result workflow.Future
}

func (r *SsmGetDeployablePatchSnapshotForInstanceResult) Get(ctx workflow.Context) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
    var output ssm.GetDeployablePatchSnapshotForInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetDocumentResult struct {
	Result workflow.Future
}

func (r *SsmGetDocumentResult) Get(ctx workflow.Context) (*ssm.GetDocumentOutput, error) {
    var output ssm.GetDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetInventoryResult struct {
	Result workflow.Future
}

func (r *SsmGetInventoryResult) Get(ctx workflow.Context) (*ssm.GetInventoryOutput, error) {
    var output ssm.GetInventoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetInventorySchemaResult struct {
	Result workflow.Future
}

func (r *SsmGetInventorySchemaResult) Get(ctx workflow.Context) (*ssm.GetInventorySchemaOutput, error) {
    var output ssm.GetInventorySchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmGetMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.GetMaintenanceWindowOutput, error) {
    var output ssm.GetMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetMaintenanceWindowExecutionResult struct {
	Result workflow.Future
}

func (r *SsmGetMaintenanceWindowExecutionResult) Get(ctx workflow.Context) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
    var output ssm.GetMaintenanceWindowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetMaintenanceWindowExecutionTaskResult struct {
	Result workflow.Future
}

func (r *SsmGetMaintenanceWindowExecutionTaskResult) Get(ctx workflow.Context) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
    var output ssm.GetMaintenanceWindowExecutionTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetMaintenanceWindowExecutionTaskInvocationResult struct {
	Result workflow.Future
}

func (r *SsmGetMaintenanceWindowExecutionTaskInvocationResult) Get(ctx workflow.Context) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
    var output ssm.GetMaintenanceWindowExecutionTaskInvocationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetMaintenanceWindowTaskResult struct {
	Result workflow.Future
}

func (r *SsmGetMaintenanceWindowTaskResult) Get(ctx workflow.Context) (*ssm.GetMaintenanceWindowTaskOutput, error) {
    var output ssm.GetMaintenanceWindowTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetOpsItemResult struct {
	Result workflow.Future
}

func (r *SsmGetOpsItemResult) Get(ctx workflow.Context) (*ssm.GetOpsItemOutput, error) {
    var output ssm.GetOpsItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetOpsSummaryResult struct {
	Result workflow.Future
}

func (r *SsmGetOpsSummaryResult) Get(ctx workflow.Context) (*ssm.GetOpsSummaryOutput, error) {
    var output ssm.GetOpsSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetParameterResult struct {
	Result workflow.Future
}

func (r *SsmGetParameterResult) Get(ctx workflow.Context) (*ssm.GetParameterOutput, error) {
    var output ssm.GetParameterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetParameterHistoryResult struct {
	Result workflow.Future
}

func (r *SsmGetParameterHistoryResult) Get(ctx workflow.Context) (*ssm.GetParameterHistoryOutput, error) {
    var output ssm.GetParameterHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetParametersResult struct {
	Result workflow.Future
}

func (r *SsmGetParametersResult) Get(ctx workflow.Context) (*ssm.GetParametersOutput, error) {
    var output ssm.GetParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetParametersByPathResult struct {
	Result workflow.Future
}

func (r *SsmGetParametersByPathResult) Get(ctx workflow.Context) (*ssm.GetParametersByPathOutput, error) {
    var output ssm.GetParametersByPathOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetPatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmGetPatchBaselineResult) Get(ctx workflow.Context) (*ssm.GetPatchBaselineOutput, error) {
    var output ssm.GetPatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetPatchBaselineForPatchGroupResult struct {
	Result workflow.Future
}

func (r *SsmGetPatchBaselineForPatchGroupResult) Get(ctx workflow.Context) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
    var output ssm.GetPatchBaselineForPatchGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmGetServiceSettingResult struct {
	Result workflow.Future
}

func (r *SsmGetServiceSettingResult) Get(ctx workflow.Context) (*ssm.GetServiceSettingOutput, error) {
    var output ssm.GetServiceSettingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmLabelParameterVersionResult struct {
	Result workflow.Future
}

func (r *SsmLabelParameterVersionResult) Get(ctx workflow.Context) (*ssm.LabelParameterVersionOutput, error) {
    var output ssm.LabelParameterVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListAssociationVersionsResult struct {
	Result workflow.Future
}

func (r *SsmListAssociationVersionsResult) Get(ctx workflow.Context) (*ssm.ListAssociationVersionsOutput, error) {
    var output ssm.ListAssociationVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListAssociationsResult struct {
	Result workflow.Future
}

func (r *SsmListAssociationsResult) Get(ctx workflow.Context) (*ssm.ListAssociationsOutput, error) {
    var output ssm.ListAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListCommandInvocationsResult struct {
	Result workflow.Future
}

func (r *SsmListCommandInvocationsResult) Get(ctx workflow.Context) (*ssm.ListCommandInvocationsOutput, error) {
    var output ssm.ListCommandInvocationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListCommandsResult struct {
	Result workflow.Future
}

func (r *SsmListCommandsResult) Get(ctx workflow.Context) (*ssm.ListCommandsOutput, error) {
    var output ssm.ListCommandsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListComplianceItemsResult struct {
	Result workflow.Future
}

func (r *SsmListComplianceItemsResult) Get(ctx workflow.Context) (*ssm.ListComplianceItemsOutput, error) {
    var output ssm.ListComplianceItemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListComplianceSummariesResult struct {
	Result workflow.Future
}

func (r *SsmListComplianceSummariesResult) Get(ctx workflow.Context) (*ssm.ListComplianceSummariesOutput, error) {
    var output ssm.ListComplianceSummariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListDocumentVersionsResult struct {
	Result workflow.Future
}

func (r *SsmListDocumentVersionsResult) Get(ctx workflow.Context) (*ssm.ListDocumentVersionsOutput, error) {
    var output ssm.ListDocumentVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListDocumentsResult struct {
	Result workflow.Future
}

func (r *SsmListDocumentsResult) Get(ctx workflow.Context) (*ssm.ListDocumentsOutput, error) {
    var output ssm.ListDocumentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListInventoryEntriesResult struct {
	Result workflow.Future
}

func (r *SsmListInventoryEntriesResult) Get(ctx workflow.Context) (*ssm.ListInventoryEntriesOutput, error) {
    var output ssm.ListInventoryEntriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListResourceComplianceSummariesResult struct {
	Result workflow.Future
}

func (r *SsmListResourceComplianceSummariesResult) Get(ctx workflow.Context) (*ssm.ListResourceComplianceSummariesOutput, error) {
    var output ssm.ListResourceComplianceSummariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListResourceDataSyncResult struct {
	Result workflow.Future
}

func (r *SsmListResourceDataSyncResult) Get(ctx workflow.Context) (*ssm.ListResourceDataSyncOutput, error) {
    var output ssm.ListResourceDataSyncOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SsmListTagsForResourceResult) Get(ctx workflow.Context) (*ssm.ListTagsForResourceOutput, error) {
    var output ssm.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmModifyDocumentPermissionResult struct {
	Result workflow.Future
}

func (r *SsmModifyDocumentPermissionResult) Get(ctx workflow.Context) (*ssm.ModifyDocumentPermissionOutput, error) {
    var output ssm.ModifyDocumentPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmPutComplianceItemsResult struct {
	Result workflow.Future
}

func (r *SsmPutComplianceItemsResult) Get(ctx workflow.Context) (*ssm.PutComplianceItemsOutput, error) {
    var output ssm.PutComplianceItemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmPutInventoryResult struct {
	Result workflow.Future
}

func (r *SsmPutInventoryResult) Get(ctx workflow.Context) (*ssm.PutInventoryOutput, error) {
    var output ssm.PutInventoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmPutParameterResult struct {
	Result workflow.Future
}

func (r *SsmPutParameterResult) Get(ctx workflow.Context) (*ssm.PutParameterOutput, error) {
    var output ssm.PutParameterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmRegisterDefaultPatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmRegisterDefaultPatchBaselineResult) Get(ctx workflow.Context) (*ssm.RegisterDefaultPatchBaselineOutput, error) {
    var output ssm.RegisterDefaultPatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmRegisterPatchBaselineForPatchGroupResult struct {
	Result workflow.Future
}

func (r *SsmRegisterPatchBaselineForPatchGroupResult) Get(ctx workflow.Context) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error) {
    var output ssm.RegisterPatchBaselineForPatchGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmRegisterTargetWithMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmRegisterTargetWithMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error) {
    var output ssm.RegisterTargetWithMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmRegisterTaskWithMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmRegisterTaskWithMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error) {
    var output ssm.RegisterTaskWithMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *SsmRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*ssm.RemoveTagsFromResourceOutput, error) {
    var output ssm.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmResetServiceSettingResult struct {
	Result workflow.Future
}

func (r *SsmResetServiceSettingResult) Get(ctx workflow.Context) (*ssm.ResetServiceSettingOutput, error) {
    var output ssm.ResetServiceSettingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmResumeSessionResult struct {
	Result workflow.Future
}

func (r *SsmResumeSessionResult) Get(ctx workflow.Context) (*ssm.ResumeSessionOutput, error) {
    var output ssm.ResumeSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmSendAutomationSignalResult struct {
	Result workflow.Future
}

func (r *SsmSendAutomationSignalResult) Get(ctx workflow.Context) (*ssm.SendAutomationSignalOutput, error) {
    var output ssm.SendAutomationSignalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmSendCommandResult struct {
	Result workflow.Future
}

func (r *SsmSendCommandResult) Get(ctx workflow.Context) (*ssm.SendCommandOutput, error) {
    var output ssm.SendCommandOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmStartAssociationsOnceResult struct {
	Result workflow.Future
}

func (r *SsmStartAssociationsOnceResult) Get(ctx workflow.Context) (*ssm.StartAssociationsOnceOutput, error) {
    var output ssm.StartAssociationsOnceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmStartAutomationExecutionResult struct {
	Result workflow.Future
}

func (r *SsmStartAutomationExecutionResult) Get(ctx workflow.Context) (*ssm.StartAutomationExecutionOutput, error) {
    var output ssm.StartAutomationExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmStartSessionResult struct {
	Result workflow.Future
}

func (r *SsmStartSessionResult) Get(ctx workflow.Context) (*ssm.StartSessionOutput, error) {
    var output ssm.StartSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmStopAutomationExecutionResult struct {
	Result workflow.Future
}

func (r *SsmStopAutomationExecutionResult) Get(ctx workflow.Context) (*ssm.StopAutomationExecutionOutput, error) {
    var output ssm.StopAutomationExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmTerminateSessionResult struct {
	Result workflow.Future
}

func (r *SsmTerminateSessionResult) Get(ctx workflow.Context) (*ssm.TerminateSessionOutput, error) {
    var output ssm.TerminateSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateAssociationResult struct {
	Result workflow.Future
}

func (r *SsmUpdateAssociationResult) Get(ctx workflow.Context) (*ssm.UpdateAssociationOutput, error) {
    var output ssm.UpdateAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateAssociationStatusResult struct {
	Result workflow.Future
}

func (r *SsmUpdateAssociationStatusResult) Get(ctx workflow.Context) (*ssm.UpdateAssociationStatusOutput, error) {
    var output ssm.UpdateAssociationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateDocumentResult struct {
	Result workflow.Future
}

func (r *SsmUpdateDocumentResult) Get(ctx workflow.Context) (*ssm.UpdateDocumentOutput, error) {
    var output ssm.UpdateDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateDocumentDefaultVersionResult struct {
	Result workflow.Future
}

func (r *SsmUpdateDocumentDefaultVersionResult) Get(ctx workflow.Context) (*ssm.UpdateDocumentDefaultVersionOutput, error) {
    var output ssm.UpdateDocumentDefaultVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateMaintenanceWindowResult struct {
	Result workflow.Future
}

func (r *SsmUpdateMaintenanceWindowResult) Get(ctx workflow.Context) (*ssm.UpdateMaintenanceWindowOutput, error) {
    var output ssm.UpdateMaintenanceWindowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateMaintenanceWindowTargetResult struct {
	Result workflow.Future
}

func (r *SsmUpdateMaintenanceWindowTargetResult) Get(ctx workflow.Context) (*ssm.UpdateMaintenanceWindowTargetOutput, error) {
    var output ssm.UpdateMaintenanceWindowTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateMaintenanceWindowTaskResult struct {
	Result workflow.Future
}

func (r *SsmUpdateMaintenanceWindowTaskResult) Get(ctx workflow.Context) (*ssm.UpdateMaintenanceWindowTaskOutput, error) {
    var output ssm.UpdateMaintenanceWindowTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateManagedInstanceRoleResult struct {
	Result workflow.Future
}

func (r *SsmUpdateManagedInstanceRoleResult) Get(ctx workflow.Context) (*ssm.UpdateManagedInstanceRoleOutput, error) {
    var output ssm.UpdateManagedInstanceRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateOpsItemResult struct {
	Result workflow.Future
}

func (r *SsmUpdateOpsItemResult) Get(ctx workflow.Context) (*ssm.UpdateOpsItemOutput, error) {
    var output ssm.UpdateOpsItemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdatePatchBaselineResult struct {
	Result workflow.Future
}

func (r *SsmUpdatePatchBaselineResult) Get(ctx workflow.Context) (*ssm.UpdatePatchBaselineOutput, error) {
    var output ssm.UpdatePatchBaselineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateResourceDataSyncResult struct {
	Result workflow.Future
}

func (r *SsmUpdateResourceDataSyncResult) Get(ctx workflow.Context) (*ssm.UpdateResourceDataSyncOutput, error) {
    var output ssm.UpdateResourceDataSyncOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SsmUpdateServiceSettingResult struct {
	Result workflow.Future
}

func (r *SsmUpdateServiceSettingResult) Get(ctx workflow.Context) (*ssm.UpdateServiceSettingOutput, error) {
    var output ssm.UpdateServiceSettingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SSMStub struct {
    activities SSMClient
}

func NewSSMStub() SSMClient {
    return &SSMStub{}
}

func (a *SSMStub) AddTagsToResource(ctx workflow.Context, input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error) {
    var output ssm.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CancelCommand(ctx workflow.Context, input *ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error) {
    var output ssm.CancelCommandOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelCommand, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CancelMaintenanceWindowExecution(ctx workflow.Context, input *ssm.CancelMaintenanceWindowExecutionInput) (*ssm.CancelMaintenanceWindowExecutionOutput, error) {
    var output ssm.CancelMaintenanceWindowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelMaintenanceWindowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateActivation(ctx workflow.Context, input *ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error) {
    var output ssm.CreateActivationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateActivation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateAssociation(ctx workflow.Context, input *ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error) {
    var output ssm.CreateAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateAssociationBatch(ctx workflow.Context, input *ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error) {
    var output ssm.CreateAssociationBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAssociationBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateDocument(ctx workflow.Context, input *ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error) {
    var output ssm.CreateDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateMaintenanceWindow(ctx workflow.Context, input *ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error) {
    var output ssm.CreateMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateOpsItem(ctx workflow.Context, input *ssm.CreateOpsItemInput) (*ssm.CreateOpsItemOutput, error) {
    var output ssm.CreateOpsItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOpsItem, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreatePatchBaseline(ctx workflow.Context, input *ssm.CreatePatchBaselineInput) (*ssm.CreatePatchBaselineOutput, error) {
    var output ssm.CreatePatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) CreateResourceDataSync(ctx workflow.Context, input *ssm.CreateResourceDataSyncInput) (*ssm.CreateResourceDataSyncOutput, error) {
    var output ssm.CreateResourceDataSyncOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResourceDataSync, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteActivation(ctx workflow.Context, input *ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error) {
    var output ssm.DeleteActivationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteActivation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteAssociation(ctx workflow.Context, input *ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error) {
    var output ssm.DeleteAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteDocument(ctx workflow.Context, input *ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error) {
    var output ssm.DeleteDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteInventory(ctx workflow.Context, input *ssm.DeleteInventoryInput) (*ssm.DeleteInventoryOutput, error) {
    var output ssm.DeleteInventoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInventory, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteMaintenanceWindow(ctx workflow.Context, input *ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error) {
    var output ssm.DeleteMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteParameter(ctx workflow.Context, input *ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error) {
    var output ssm.DeleteParameterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteParameter, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteParameters(ctx workflow.Context, input *ssm.DeleteParametersInput) (*ssm.DeleteParametersOutput, error) {
    var output ssm.DeleteParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeletePatchBaseline(ctx workflow.Context, input *ssm.DeletePatchBaselineInput) (*ssm.DeletePatchBaselineOutput, error) {
    var output ssm.DeletePatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeleteResourceDataSync(ctx workflow.Context, input *ssm.DeleteResourceDataSyncInput) (*ssm.DeleteResourceDataSyncOutput, error) {
    var output ssm.DeleteResourceDataSyncOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceDataSync, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeregisterManagedInstance(ctx workflow.Context, input *ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error) {
    var output ssm.DeregisterManagedInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterManagedInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeregisterPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error) {
    var output ssm.DeregisterPatchBaselineForPatchGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterPatchBaselineForPatchGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeregisterTargetFromMaintenanceWindow(ctx workflow.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error) {
    var output ssm.DeregisterTargetFromMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterTargetFromMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DeregisterTaskFromMaintenanceWindow(ctx workflow.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error) {
    var output ssm.DeregisterTaskFromMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterTaskFromMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeActivations(ctx workflow.Context, input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error) {
    var output ssm.DescribeActivationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeActivations, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeAssociation(ctx workflow.Context, input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error) {
    var output ssm.DescribeAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeAssociationExecutionTargets(ctx workflow.Context, input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
    var output ssm.DescribeAssociationExecutionTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssociationExecutionTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeAssociationExecutions(ctx workflow.Context, input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error) {
    var output ssm.DescribeAssociationExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssociationExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeAutomationExecutions(ctx workflow.Context, input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error) {
    var output ssm.DescribeAutomationExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAutomationExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeAutomationStepExecutions(ctx workflow.Context, input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
    var output ssm.DescribeAutomationStepExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAutomationStepExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeAvailablePatches(ctx workflow.Context, input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error) {
    var output ssm.DescribeAvailablePatchesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAvailablePatches, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeDocument(ctx workflow.Context, input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error) {
    var output ssm.DescribeDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeDocumentPermission(ctx workflow.Context, input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error) {
    var output ssm.DescribeDocumentPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDocumentPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeEffectiveInstanceAssociations(ctx workflow.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
    var output ssm.DescribeEffectiveInstanceAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEffectiveInstanceAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeEffectivePatchesForPatchBaseline(ctx workflow.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
    var output ssm.DescribeEffectivePatchesForPatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEffectivePatchesForPatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeInstanceAssociationsStatus(ctx workflow.Context, input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
    var output ssm.DescribeInstanceAssociationsStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceAssociationsStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeInstanceInformation(ctx workflow.Context, input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
    var output ssm.DescribeInstanceInformationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstanceInformation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeInstancePatchStates(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error) {
    var output ssm.DescribeInstancePatchStatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstancePatchStates, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeInstancePatchStatesForPatchGroup(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
    var output ssm.DescribeInstancePatchStatesForPatchGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstancePatchStatesForPatchGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeInstancePatches(ctx workflow.Context, input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error) {
    var output ssm.DescribeInstancePatchesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstancePatches, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeInventoryDeletions(ctx workflow.Context, input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error) {
    var output ssm.DescribeInventoryDeletionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInventoryDeletions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowExecutionTaskInvocations(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
    var output ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowExecutionTaskInvocations, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowExecutionTasks(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
    var output ssm.DescribeMaintenanceWindowExecutionTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowExecutionTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowExecutions(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
    var output ssm.DescribeMaintenanceWindowExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowSchedule(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
    var output ssm.DescribeMaintenanceWindowScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowTargets(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
    var output ssm.DescribeMaintenanceWindowTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowTasks(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
    var output ssm.DescribeMaintenanceWindowTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindows(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error) {
    var output ssm.DescribeMaintenanceWindowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindows, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeMaintenanceWindowsForTarget(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
    var output ssm.DescribeMaintenanceWindowsForTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMaintenanceWindowsForTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeOpsItems(ctx workflow.Context, input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error) {
    var output ssm.DescribeOpsItemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOpsItems, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeParameters(ctx workflow.Context, input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error) {
    var output ssm.DescribeParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribePatchBaselines(ctx workflow.Context, input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error) {
    var output ssm.DescribePatchBaselinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePatchBaselines, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribePatchGroupState(ctx workflow.Context, input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error) {
    var output ssm.DescribePatchGroupStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePatchGroupState, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribePatchGroups(ctx workflow.Context, input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error) {
    var output ssm.DescribePatchGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePatchGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribePatchProperties(ctx workflow.Context, input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error) {
    var output ssm.DescribePatchPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePatchProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) DescribeSessions(ctx workflow.Context, input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error) {
    var output ssm.DescribeSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetAutomationExecution(ctx workflow.Context, input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error) {
    var output ssm.GetAutomationExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAutomationExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetCalendarState(ctx workflow.Context, input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error) {
    var output ssm.GetCalendarStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCalendarState, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetCommandInvocation(ctx workflow.Context, input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error) {
    var output ssm.GetCommandInvocationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCommandInvocation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetConnectionStatus(ctx workflow.Context, input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error) {
    var output ssm.GetConnectionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnectionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetDefaultPatchBaseline(ctx workflow.Context, input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error) {
    var output ssm.GetDefaultPatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDefaultPatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetDeployablePatchSnapshotForInstance(ctx workflow.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
    var output ssm.GetDeployablePatchSnapshotForInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeployablePatchSnapshotForInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetDocument(ctx workflow.Context, input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error) {
    var output ssm.GetDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetInventory(ctx workflow.Context, input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error) {
    var output ssm.GetInventoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInventory, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetInventorySchema(ctx workflow.Context, input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error) {
    var output ssm.GetInventorySchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInventorySchema, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetMaintenanceWindow(ctx workflow.Context, input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error) {
    var output ssm.GetMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetMaintenanceWindowExecution(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
    var output ssm.GetMaintenanceWindowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMaintenanceWindowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetMaintenanceWindowExecutionTask(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
    var output ssm.GetMaintenanceWindowExecutionTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMaintenanceWindowExecutionTask, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetMaintenanceWindowExecutionTaskInvocation(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
    var output ssm.GetMaintenanceWindowExecutionTaskInvocationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMaintenanceWindowExecutionTaskInvocation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetMaintenanceWindowTask(ctx workflow.Context, input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error) {
    var output ssm.GetMaintenanceWindowTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMaintenanceWindowTask, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetOpsItem(ctx workflow.Context, input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error) {
    var output ssm.GetOpsItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOpsItem, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetOpsSummary(ctx workflow.Context, input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error) {
    var output ssm.GetOpsSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetOpsSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetParameter(ctx workflow.Context, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
    var output ssm.GetParameterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetParameter, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetParameterHistory(ctx workflow.Context, input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error) {
    var output ssm.GetParameterHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetParameterHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetParameters(ctx workflow.Context, input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error) {
    var output ssm.GetParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetParametersByPath(ctx workflow.Context, input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
    var output ssm.GetParametersByPathOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetParametersByPath, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetPatchBaseline(ctx workflow.Context, input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error) {
    var output ssm.GetPatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
    var output ssm.GetPatchBaselineForPatchGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPatchBaselineForPatchGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) GetServiceSetting(ctx workflow.Context, input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error) {
    var output ssm.GetServiceSettingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceSetting, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) LabelParameterVersion(ctx workflow.Context, input *ssm.LabelParameterVersionInput) (*ssm.LabelParameterVersionOutput, error) {
    var output ssm.LabelParameterVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.LabelParameterVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListAssociationVersions(ctx workflow.Context, input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error) {
    var output ssm.ListAssociationVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssociationVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListAssociations(ctx workflow.Context, input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error) {
    var output ssm.ListAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListCommandInvocations(ctx workflow.Context, input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error) {
    var output ssm.ListCommandInvocationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCommandInvocations, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListCommands(ctx workflow.Context, input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error) {
    var output ssm.ListCommandsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCommands, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListComplianceItems(ctx workflow.Context, input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error) {
    var output ssm.ListComplianceItemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListComplianceItems, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListComplianceSummaries(ctx workflow.Context, input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error) {
    var output ssm.ListComplianceSummariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListComplianceSummaries, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListDocumentVersions(ctx workflow.Context, input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error) {
    var output ssm.ListDocumentVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDocumentVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListDocuments(ctx workflow.Context, input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error) {
    var output ssm.ListDocumentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDocuments, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListInventoryEntries(ctx workflow.Context, input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error) {
    var output ssm.ListInventoryEntriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInventoryEntries, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListResourceComplianceSummaries(ctx workflow.Context, input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error) {
    var output ssm.ListResourceComplianceSummariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceComplianceSummaries, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListResourceDataSync(ctx workflow.Context, input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error) {
    var output ssm.ListResourceDataSyncOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceDataSync, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ListTagsForResource(ctx workflow.Context, input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error) {
    var output ssm.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ModifyDocumentPermission(ctx workflow.Context, input *ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error) {
    var output ssm.ModifyDocumentPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDocumentPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) PutComplianceItems(ctx workflow.Context, input *ssm.PutComplianceItemsInput) (*ssm.PutComplianceItemsOutput, error) {
    var output ssm.PutComplianceItemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutComplianceItems, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) PutInventory(ctx workflow.Context, input *ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error) {
    var output ssm.PutInventoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutInventory, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) PutParameter(ctx workflow.Context, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
    var output ssm.PutParameterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutParameter, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) RegisterDefaultPatchBaseline(ctx workflow.Context, input *ssm.RegisterDefaultPatchBaselineInput) (*ssm.RegisterDefaultPatchBaselineOutput, error) {
    var output ssm.RegisterDefaultPatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterDefaultPatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) RegisterPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error) {
    var output ssm.RegisterPatchBaselineForPatchGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterPatchBaselineForPatchGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) RegisterTargetWithMaintenanceWindow(ctx workflow.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error) {
    var output ssm.RegisterTargetWithMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterTargetWithMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) RegisterTaskWithMaintenanceWindow(ctx workflow.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error) {
    var output ssm.RegisterTaskWithMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterTaskWithMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) RemoveTagsFromResource(ctx workflow.Context, input *ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error) {
    var output ssm.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ResetServiceSetting(ctx workflow.Context, input *ssm.ResetServiceSettingInput) (*ssm.ResetServiceSettingOutput, error) {
    var output ssm.ResetServiceSettingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetServiceSetting, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) ResumeSession(ctx workflow.Context, input *ssm.ResumeSessionInput) (*ssm.ResumeSessionOutput, error) {
    var output ssm.ResumeSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResumeSession, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) SendAutomationSignal(ctx workflow.Context, input *ssm.SendAutomationSignalInput) (*ssm.SendAutomationSignalOutput, error) {
    var output ssm.SendAutomationSignalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendAutomationSignal, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) SendCommand(ctx workflow.Context, input *ssm.SendCommandInput) (*ssm.SendCommandOutput, error) {
    var output ssm.SendCommandOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendCommand, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) StartAssociationsOnce(ctx workflow.Context, input *ssm.StartAssociationsOnceInput) (*ssm.StartAssociationsOnceOutput, error) {
    var output ssm.StartAssociationsOnceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartAssociationsOnce, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) StartAutomationExecution(ctx workflow.Context, input *ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error) {
    var output ssm.StartAutomationExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartAutomationExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) StartSession(ctx workflow.Context, input *ssm.StartSessionInput) (*ssm.StartSessionOutput, error) {
    var output ssm.StartSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSession, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) StopAutomationExecution(ctx workflow.Context, input *ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error) {
    var output ssm.StopAutomationExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopAutomationExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) TerminateSession(ctx workflow.Context, input *ssm.TerminateSessionInput) (*ssm.TerminateSessionOutput, error) {
    var output ssm.TerminateSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateSession, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateAssociation(ctx workflow.Context, input *ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error) {
    var output ssm.UpdateAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateAssociationStatus(ctx workflow.Context, input *ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error) {
    var output ssm.UpdateAssociationStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAssociationStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateDocument(ctx workflow.Context, input *ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error) {
    var output ssm.UpdateDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateDocumentDefaultVersion(ctx workflow.Context, input *ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error) {
    var output ssm.UpdateDocumentDefaultVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDocumentDefaultVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateMaintenanceWindow(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error) {
    var output ssm.UpdateMaintenanceWindowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMaintenanceWindow, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateMaintenanceWindowTarget(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTargetInput) (*ssm.UpdateMaintenanceWindowTargetOutput, error) {
    var output ssm.UpdateMaintenanceWindowTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMaintenanceWindowTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateMaintenanceWindowTask(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTaskInput) (*ssm.UpdateMaintenanceWindowTaskOutput, error) {
    var output ssm.UpdateMaintenanceWindowTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMaintenanceWindowTask, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateManagedInstanceRole(ctx workflow.Context, input *ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error) {
    var output ssm.UpdateManagedInstanceRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateManagedInstanceRole, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateOpsItem(ctx workflow.Context, input *ssm.UpdateOpsItemInput) (*ssm.UpdateOpsItemOutput, error) {
    var output ssm.UpdateOpsItemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOpsItem, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdatePatchBaseline(ctx workflow.Context, input *ssm.UpdatePatchBaselineInput) (*ssm.UpdatePatchBaselineOutput, error) {
    var output ssm.UpdatePatchBaselineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePatchBaseline, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateResourceDataSync(ctx workflow.Context, input *ssm.UpdateResourceDataSyncInput) (*ssm.UpdateResourceDataSyncOutput, error) {
    var output ssm.UpdateResourceDataSyncOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateResourceDataSync, input).Get(ctx, &output)
    return &output, err
}

func (a *SSMStub) UpdateServiceSetting(ctx workflow.Context, input *ssm.UpdateServiceSettingInput) (*ssm.UpdateServiceSettingOutput, error) {
    var output ssm.UpdateServiceSettingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceSetting, input).Get(ctx, &output)
    return &output, err
}


func (a *SSMStub) WaitUntilCommandExecuted(ctx workflow.Context, input *ssm.GetCommandInvocationInput) error {
    return a.client.WaitUntilCommandExecuted(input)
}
