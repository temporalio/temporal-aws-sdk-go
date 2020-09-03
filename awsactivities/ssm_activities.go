package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type SSMActivities struct {
	client ssmiface.SSMAPI
}

func NewSSMActivities(client ssmiface.SSMAPI) *SSMActivities {
    return &SSMActivities{client: client}
}


func (a *SSMActivities) AddTagsToResource(input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}



func (a *SSMActivities) CancelCommand(input *ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error) {
    return a.client.CancelCommand(input)
}



func (a *SSMActivities) CancelMaintenanceWindowExecution(input *ssm.CancelMaintenanceWindowExecutionInput) (*ssm.CancelMaintenanceWindowExecutionOutput, error) {
    return a.client.CancelMaintenanceWindowExecution(input)
}



func (a *SSMActivities) CreateActivation(input *ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error) {
    return a.client.CreateActivation(input)
}



func (a *SSMActivities) CreateAssociation(input *ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error) {
    return a.client.CreateAssociation(input)
}



func (a *SSMActivities) CreateAssociationBatch(input *ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error) {
    return a.client.CreateAssociationBatch(input)
}



func (a *SSMActivities) CreateDocument(input *ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error) {
    return a.client.CreateDocument(input)
}



func (a *SSMActivities) CreateMaintenanceWindow(input *ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error) {
    return a.client.CreateMaintenanceWindow(input)
}



func (a *SSMActivities) CreateOpsItem(input *ssm.CreateOpsItemInput) (*ssm.CreateOpsItemOutput, error) {
    return a.client.CreateOpsItem(input)
}



func (a *SSMActivities) CreatePatchBaseline(input *ssm.CreatePatchBaselineInput) (*ssm.CreatePatchBaselineOutput, error) {
    return a.client.CreatePatchBaseline(input)
}



func (a *SSMActivities) CreateResourceDataSync(input *ssm.CreateResourceDataSyncInput) (*ssm.CreateResourceDataSyncOutput, error) {
    return a.client.CreateResourceDataSync(input)
}



func (a *SSMActivities) DeleteActivation(input *ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error) {
    return a.client.DeleteActivation(input)
}



func (a *SSMActivities) DeleteAssociation(input *ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error) {
    return a.client.DeleteAssociation(input)
}



func (a *SSMActivities) DeleteDocument(input *ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error) {
    return a.client.DeleteDocument(input)
}



func (a *SSMActivities) DeleteInventory(input *ssm.DeleteInventoryInput) (*ssm.DeleteInventoryOutput, error) {
    return a.client.DeleteInventory(input)
}



func (a *SSMActivities) DeleteMaintenanceWindow(input *ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error) {
    return a.client.DeleteMaintenanceWindow(input)
}



func (a *SSMActivities) DeleteParameter(input *ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error) {
    return a.client.DeleteParameter(input)
}



func (a *SSMActivities) DeleteParameters(input *ssm.DeleteParametersInput) (*ssm.DeleteParametersOutput, error) {
    return a.client.DeleteParameters(input)
}



func (a *SSMActivities) DeletePatchBaseline(input *ssm.DeletePatchBaselineInput) (*ssm.DeletePatchBaselineOutput, error) {
    return a.client.DeletePatchBaseline(input)
}



func (a *SSMActivities) DeleteResourceDataSync(input *ssm.DeleteResourceDataSyncInput) (*ssm.DeleteResourceDataSyncOutput, error) {
    return a.client.DeleteResourceDataSync(input)
}



func (a *SSMActivities) DeregisterManagedInstance(input *ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error) {
    return a.client.DeregisterManagedInstance(input)
}



func (a *SSMActivities) DeregisterPatchBaselineForPatchGroup(input *ssm.DeregisterPatchBaselineForPatchGroupInput) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error) {
    return a.client.DeregisterPatchBaselineForPatchGroup(input)
}



func (a *SSMActivities) DeregisterTargetFromMaintenanceWindow(input *ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error) {
    return a.client.DeregisterTargetFromMaintenanceWindow(input)
}



func (a *SSMActivities) DeregisterTaskFromMaintenanceWindow(input *ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error) {
    return a.client.DeregisterTaskFromMaintenanceWindow(input)
}



func (a *SSMActivities) DescribeActivations(input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error) {
    return a.client.DescribeActivations(input)
}



func (a *SSMActivities) DescribeAssociation(input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error) {
    return a.client.DescribeAssociation(input)
}



func (a *SSMActivities) DescribeAssociationExecutionTargets(input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
    return a.client.DescribeAssociationExecutionTargets(input)
}



func (a *SSMActivities) DescribeAssociationExecutions(input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error) {
    return a.client.DescribeAssociationExecutions(input)
}



func (a *SSMActivities) DescribeAutomationExecutions(input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error) {
    return a.client.DescribeAutomationExecutions(input)
}



func (a *SSMActivities) DescribeAutomationStepExecutions(input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
    return a.client.DescribeAutomationStepExecutions(input)
}



func (a *SSMActivities) DescribeAvailablePatches(input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error) {
    return a.client.DescribeAvailablePatches(input)
}



func (a *SSMActivities) DescribeDocument(input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error) {
    return a.client.DescribeDocument(input)
}



func (a *SSMActivities) DescribeDocumentPermission(input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error) {
    return a.client.DescribeDocumentPermission(input)
}



func (a *SSMActivities) DescribeEffectiveInstanceAssociations(input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
    return a.client.DescribeEffectiveInstanceAssociations(input)
}



func (a *SSMActivities) DescribeEffectivePatchesForPatchBaseline(input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
    return a.client.DescribeEffectivePatchesForPatchBaseline(input)
}



func (a *SSMActivities) DescribeInstanceAssociationsStatus(input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
    return a.client.DescribeInstanceAssociationsStatus(input)
}



func (a *SSMActivities) DescribeInstanceInformation(input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
    return a.client.DescribeInstanceInformation(input)
}



func (a *SSMActivities) DescribeInstancePatchStates(input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error) {
    return a.client.DescribeInstancePatchStates(input)
}



func (a *SSMActivities) DescribeInstancePatchStatesForPatchGroup(input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
    return a.client.DescribeInstancePatchStatesForPatchGroup(input)
}



func (a *SSMActivities) DescribeInstancePatches(input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error) {
    return a.client.DescribeInstancePatches(input)
}



func (a *SSMActivities) DescribeInventoryDeletions(input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error) {
    return a.client.DescribeInventoryDeletions(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowExecutionTaskInvocations(input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
    return a.client.DescribeMaintenanceWindowExecutionTaskInvocations(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowExecutionTasks(input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
    return a.client.DescribeMaintenanceWindowExecutionTasks(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowExecutions(input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
    return a.client.DescribeMaintenanceWindowExecutions(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowSchedule(input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
    return a.client.DescribeMaintenanceWindowSchedule(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowTargets(input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
    return a.client.DescribeMaintenanceWindowTargets(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowTasks(input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
    return a.client.DescribeMaintenanceWindowTasks(input)
}



func (a *SSMActivities) DescribeMaintenanceWindows(input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error) {
    return a.client.DescribeMaintenanceWindows(input)
}



func (a *SSMActivities) DescribeMaintenanceWindowsForTarget(input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
    return a.client.DescribeMaintenanceWindowsForTarget(input)
}



func (a *SSMActivities) DescribeOpsItems(input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error) {
    return a.client.DescribeOpsItems(input)
}



func (a *SSMActivities) DescribeParameters(input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error) {
    return a.client.DescribeParameters(input)
}



func (a *SSMActivities) DescribePatchBaselines(input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error) {
    return a.client.DescribePatchBaselines(input)
}



func (a *SSMActivities) DescribePatchGroupState(input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error) {
    return a.client.DescribePatchGroupState(input)
}



func (a *SSMActivities) DescribePatchGroups(input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error) {
    return a.client.DescribePatchGroups(input)
}



func (a *SSMActivities) DescribePatchProperties(input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error) {
    return a.client.DescribePatchProperties(input)
}



func (a *SSMActivities) DescribeSessions(input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error) {
    return a.client.DescribeSessions(input)
}



func (a *SSMActivities) GetAutomationExecution(input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error) {
    return a.client.GetAutomationExecution(input)
}



func (a *SSMActivities) GetCalendarState(input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error) {
    return a.client.GetCalendarState(input)
}



func (a *SSMActivities) GetCommandInvocation(input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error) {
    return a.client.GetCommandInvocation(input)
}



func (a *SSMActivities) GetConnectionStatus(input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error) {
    return a.client.GetConnectionStatus(input)
}



func (a *SSMActivities) GetDefaultPatchBaseline(input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error) {
    return a.client.GetDefaultPatchBaseline(input)
}



func (a *SSMActivities) GetDeployablePatchSnapshotForInstance(input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
    return a.client.GetDeployablePatchSnapshotForInstance(input)
}



func (a *SSMActivities) GetDocument(input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error) {
    return a.client.GetDocument(input)
}



func (a *SSMActivities) GetInventory(input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error) {
    return a.client.GetInventory(input)
}



func (a *SSMActivities) GetInventorySchema(input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error) {
    return a.client.GetInventorySchema(input)
}



func (a *SSMActivities) GetMaintenanceWindow(input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error) {
    return a.client.GetMaintenanceWindow(input)
}



func (a *SSMActivities) GetMaintenanceWindowExecution(input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
    return a.client.GetMaintenanceWindowExecution(input)
}



func (a *SSMActivities) GetMaintenanceWindowExecutionTask(input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
    return a.client.GetMaintenanceWindowExecutionTask(input)
}



func (a *SSMActivities) GetMaintenanceWindowExecutionTaskInvocation(input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
    return a.client.GetMaintenanceWindowExecutionTaskInvocation(input)
}



func (a *SSMActivities) GetMaintenanceWindowTask(input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error) {
    return a.client.GetMaintenanceWindowTask(input)
}



func (a *SSMActivities) GetOpsItem(input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error) {
    return a.client.GetOpsItem(input)
}



func (a *SSMActivities) GetOpsSummary(input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error) {
    return a.client.GetOpsSummary(input)
}



func (a *SSMActivities) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
    return a.client.GetParameter(input)
}



func (a *SSMActivities) GetParameterHistory(input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error) {
    return a.client.GetParameterHistory(input)
}



func (a *SSMActivities) GetParameters(input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error) {
    return a.client.GetParameters(input)
}



func (a *SSMActivities) GetParametersByPath(input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
    return a.client.GetParametersByPath(input)
}



func (a *SSMActivities) GetPatchBaseline(input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error) {
    return a.client.GetPatchBaseline(input)
}



func (a *SSMActivities) GetPatchBaselineForPatchGroup(input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
    return a.client.GetPatchBaselineForPatchGroup(input)
}



func (a *SSMActivities) GetServiceSetting(input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error) {
    return a.client.GetServiceSetting(input)
}



func (a *SSMActivities) LabelParameterVersion(input *ssm.LabelParameterVersionInput) (*ssm.LabelParameterVersionOutput, error) {
    return a.client.LabelParameterVersion(input)
}



func (a *SSMActivities) ListAssociationVersions(input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error) {
    return a.client.ListAssociationVersions(input)
}



func (a *SSMActivities) ListAssociations(input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error) {
    return a.client.ListAssociations(input)
}



func (a *SSMActivities) ListCommandInvocations(input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error) {
    return a.client.ListCommandInvocations(input)
}



func (a *SSMActivities) ListCommands(input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error) {
    return a.client.ListCommands(input)
}



func (a *SSMActivities) ListComplianceItems(input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error) {
    return a.client.ListComplianceItems(input)
}



func (a *SSMActivities) ListComplianceSummaries(input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error) {
    return a.client.ListComplianceSummaries(input)
}



func (a *SSMActivities) ListDocumentVersions(input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error) {
    return a.client.ListDocumentVersions(input)
}



func (a *SSMActivities) ListDocuments(input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error) {
    return a.client.ListDocuments(input)
}



func (a *SSMActivities) ListInventoryEntries(input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error) {
    return a.client.ListInventoryEntries(input)
}



func (a *SSMActivities) ListResourceComplianceSummaries(input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error) {
    return a.client.ListResourceComplianceSummaries(input)
}



func (a *SSMActivities) ListResourceDataSync(input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error) {
    return a.client.ListResourceDataSync(input)
}



func (a *SSMActivities) ListTagsForResource(input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *SSMActivities) ModifyDocumentPermission(input *ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error) {
    return a.client.ModifyDocumentPermission(input)
}



func (a *SSMActivities) PutComplianceItems(input *ssm.PutComplianceItemsInput) (*ssm.PutComplianceItemsOutput, error) {
    return a.client.PutComplianceItems(input)
}



func (a *SSMActivities) PutInventory(input *ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error) {
    return a.client.PutInventory(input)
}



func (a *SSMActivities) PutParameter(input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
    return a.client.PutParameter(input)
}



func (a *SSMActivities) RegisterDefaultPatchBaseline(input *ssm.RegisterDefaultPatchBaselineInput) (*ssm.RegisterDefaultPatchBaselineOutput, error) {
    return a.client.RegisterDefaultPatchBaseline(input)
}



func (a *SSMActivities) RegisterPatchBaselineForPatchGroup(input *ssm.RegisterPatchBaselineForPatchGroupInput) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error) {
    return a.client.RegisterPatchBaselineForPatchGroup(input)
}



func (a *SSMActivities) RegisterTargetWithMaintenanceWindow(input *ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error) {
    return a.client.RegisterTargetWithMaintenanceWindow(input)
}



func (a *SSMActivities) RegisterTaskWithMaintenanceWindow(input *ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error) {
    return a.client.RegisterTaskWithMaintenanceWindow(input)
}



func (a *SSMActivities) RemoveTagsFromResource(input *ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}



func (a *SSMActivities) ResetServiceSetting(input *ssm.ResetServiceSettingInput) (*ssm.ResetServiceSettingOutput, error) {
    return a.client.ResetServiceSetting(input)
}



func (a *SSMActivities) ResumeSession(input *ssm.ResumeSessionInput) (*ssm.ResumeSessionOutput, error) {
    return a.client.ResumeSession(input)
}



func (a *SSMActivities) SendAutomationSignal(input *ssm.SendAutomationSignalInput) (*ssm.SendAutomationSignalOutput, error) {
    return a.client.SendAutomationSignal(input)
}



func (a *SSMActivities) SendCommand(input *ssm.SendCommandInput) (*ssm.SendCommandOutput, error) {
    return a.client.SendCommand(input)
}



func (a *SSMActivities) StartAssociationsOnce(input *ssm.StartAssociationsOnceInput) (*ssm.StartAssociationsOnceOutput, error) {
    return a.client.StartAssociationsOnce(input)
}



func (a *SSMActivities) StartAutomationExecution(input *ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error) {
    return a.client.StartAutomationExecution(input)
}



func (a *SSMActivities) StartSession(input *ssm.StartSessionInput) (*ssm.StartSessionOutput, error) {
    return a.client.StartSession(input)
}



func (a *SSMActivities) StopAutomationExecution(input *ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error) {
    return a.client.StopAutomationExecution(input)
}



func (a *SSMActivities) TerminateSession(input *ssm.TerminateSessionInput) (*ssm.TerminateSessionOutput, error) {
    return a.client.TerminateSession(input)
}



func (a *SSMActivities) UpdateAssociation(input *ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error) {
    return a.client.UpdateAssociation(input)
}



func (a *SSMActivities) UpdateAssociationStatus(input *ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error) {
    return a.client.UpdateAssociationStatus(input)
}



func (a *SSMActivities) UpdateDocument(input *ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error) {
    return a.client.UpdateDocument(input)
}



func (a *SSMActivities) UpdateDocumentDefaultVersion(input *ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error) {
    return a.client.UpdateDocumentDefaultVersion(input)
}



func (a *SSMActivities) UpdateMaintenanceWindow(input *ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error) {
    return a.client.UpdateMaintenanceWindow(input)
}



func (a *SSMActivities) UpdateMaintenanceWindowTarget(input *ssm.UpdateMaintenanceWindowTargetInput) (*ssm.UpdateMaintenanceWindowTargetOutput, error) {
    return a.client.UpdateMaintenanceWindowTarget(input)
}



func (a *SSMActivities) UpdateMaintenanceWindowTask(input *ssm.UpdateMaintenanceWindowTaskInput) (*ssm.UpdateMaintenanceWindowTaskOutput, error) {
    return a.client.UpdateMaintenanceWindowTask(input)
}



func (a *SSMActivities) UpdateManagedInstanceRole(input *ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error) {
    return a.client.UpdateManagedInstanceRole(input)
}



func (a *SSMActivities) UpdateOpsItem(input *ssm.UpdateOpsItemInput) (*ssm.UpdateOpsItemOutput, error) {
    return a.client.UpdateOpsItem(input)
}



func (a *SSMActivities) UpdatePatchBaseline(input *ssm.UpdatePatchBaselineInput) (*ssm.UpdatePatchBaselineOutput, error) {
    return a.client.UpdatePatchBaseline(input)
}



func (a *SSMActivities) UpdateResourceDataSync(input *ssm.UpdateResourceDataSyncInput) (*ssm.UpdateResourceDataSyncOutput, error) {
    return a.client.UpdateResourceDataSync(input)
}



func (a *SSMActivities) UpdateServiceSetting(input *ssm.UpdateServiceSettingInput) (*ssm.UpdateServiceSettingOutput, error) {
    return a.client.UpdateServiceSetting(input)
}



func (a *SSMActivities) WaitUntilCommandExecuted(input *ssm.GetCommandInvocationInput) error {
    return a.client.WaitUntilCommandExecuted(input)
}

