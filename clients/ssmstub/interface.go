// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ssmstub

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTagsToResource(ctx workflow.Context, input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *ssm.AddTagsToResourceInput) *SSMAddTagsToResourceFuture

	CancelCommand(ctx workflow.Context, input *ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error)
	CancelCommandAsync(ctx workflow.Context, input *ssm.CancelCommandInput) *SSMCancelCommandFuture

	CancelMaintenanceWindowExecution(ctx workflow.Context, input *ssm.CancelMaintenanceWindowExecutionInput) (*ssm.CancelMaintenanceWindowExecutionOutput, error)
	CancelMaintenanceWindowExecutionAsync(ctx workflow.Context, input *ssm.CancelMaintenanceWindowExecutionInput) *SSMCancelMaintenanceWindowExecutionFuture

	CreateActivation(ctx workflow.Context, input *ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error)
	CreateActivationAsync(ctx workflow.Context, input *ssm.CreateActivationInput) *SSMCreateActivationFuture

	CreateAssociation(ctx workflow.Context, input *ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error)
	CreateAssociationAsync(ctx workflow.Context, input *ssm.CreateAssociationInput) *SSMCreateAssociationFuture

	CreateAssociationBatch(ctx workflow.Context, input *ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error)
	CreateAssociationBatchAsync(ctx workflow.Context, input *ssm.CreateAssociationBatchInput) *SSMCreateAssociationBatchFuture

	CreateDocument(ctx workflow.Context, input *ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error)
	CreateDocumentAsync(ctx workflow.Context, input *ssm.CreateDocumentInput) *SSMCreateDocumentFuture

	CreateMaintenanceWindow(ctx workflow.Context, input *ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error)
	CreateMaintenanceWindowAsync(ctx workflow.Context, input *ssm.CreateMaintenanceWindowInput) *SSMCreateMaintenanceWindowFuture

	CreateOpsItem(ctx workflow.Context, input *ssm.CreateOpsItemInput) (*ssm.CreateOpsItemOutput, error)
	CreateOpsItemAsync(ctx workflow.Context, input *ssm.CreateOpsItemInput) *SSMCreateOpsItemFuture

	CreatePatchBaseline(ctx workflow.Context, input *ssm.CreatePatchBaselineInput) (*ssm.CreatePatchBaselineOutput, error)
	CreatePatchBaselineAsync(ctx workflow.Context, input *ssm.CreatePatchBaselineInput) *SSMCreatePatchBaselineFuture

	CreateResourceDataSync(ctx workflow.Context, input *ssm.CreateResourceDataSyncInput) (*ssm.CreateResourceDataSyncOutput, error)
	CreateResourceDataSyncAsync(ctx workflow.Context, input *ssm.CreateResourceDataSyncInput) *SSMCreateResourceDataSyncFuture

	DeleteActivation(ctx workflow.Context, input *ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error)
	DeleteActivationAsync(ctx workflow.Context, input *ssm.DeleteActivationInput) *SSMDeleteActivationFuture

	DeleteAssociation(ctx workflow.Context, input *ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error)
	DeleteAssociationAsync(ctx workflow.Context, input *ssm.DeleteAssociationInput) *SSMDeleteAssociationFuture

	DeleteDocument(ctx workflow.Context, input *ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error)
	DeleteDocumentAsync(ctx workflow.Context, input *ssm.DeleteDocumentInput) *SSMDeleteDocumentFuture

	DeleteInventory(ctx workflow.Context, input *ssm.DeleteInventoryInput) (*ssm.DeleteInventoryOutput, error)
	DeleteInventoryAsync(ctx workflow.Context, input *ssm.DeleteInventoryInput) *SSMDeleteInventoryFuture

	DeleteMaintenanceWindow(ctx workflow.Context, input *ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error)
	DeleteMaintenanceWindowAsync(ctx workflow.Context, input *ssm.DeleteMaintenanceWindowInput) *SSMDeleteMaintenanceWindowFuture

	DeleteParameter(ctx workflow.Context, input *ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error)
	DeleteParameterAsync(ctx workflow.Context, input *ssm.DeleteParameterInput) *SSMDeleteParameterFuture

	DeleteParameters(ctx workflow.Context, input *ssm.DeleteParametersInput) (*ssm.DeleteParametersOutput, error)
	DeleteParametersAsync(ctx workflow.Context, input *ssm.DeleteParametersInput) *SSMDeleteParametersFuture

	DeletePatchBaseline(ctx workflow.Context, input *ssm.DeletePatchBaselineInput) (*ssm.DeletePatchBaselineOutput, error)
	DeletePatchBaselineAsync(ctx workflow.Context, input *ssm.DeletePatchBaselineInput) *SSMDeletePatchBaselineFuture

	DeleteResourceDataSync(ctx workflow.Context, input *ssm.DeleteResourceDataSyncInput) (*ssm.DeleteResourceDataSyncOutput, error)
	DeleteResourceDataSyncAsync(ctx workflow.Context, input *ssm.DeleteResourceDataSyncInput) *SSMDeleteResourceDataSyncFuture

	DeregisterManagedInstance(ctx workflow.Context, input *ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error)
	DeregisterManagedInstanceAsync(ctx workflow.Context, input *ssm.DeregisterManagedInstanceInput) *SSMDeregisterManagedInstanceFuture

	DeregisterPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error)
	DeregisterPatchBaselineForPatchGroupAsync(ctx workflow.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) *SSMDeregisterPatchBaselineForPatchGroupFuture

	DeregisterTargetFromMaintenanceWindow(ctx workflow.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error)
	DeregisterTargetFromMaintenanceWindowAsync(ctx workflow.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) *SSMDeregisterTargetFromMaintenanceWindowFuture

	DeregisterTaskFromMaintenanceWindow(ctx workflow.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error)
	DeregisterTaskFromMaintenanceWindowAsync(ctx workflow.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) *SSMDeregisterTaskFromMaintenanceWindowFuture

	DescribeActivations(ctx workflow.Context, input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error)
	DescribeActivationsAsync(ctx workflow.Context, input *ssm.DescribeActivationsInput) *SSMDescribeActivationsFuture

	DescribeAssociation(ctx workflow.Context, input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error)
	DescribeAssociationAsync(ctx workflow.Context, input *ssm.DescribeAssociationInput) *SSMDescribeAssociationFuture

	DescribeAssociationExecutionTargets(ctx workflow.Context, input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error)
	DescribeAssociationExecutionTargetsAsync(ctx workflow.Context, input *ssm.DescribeAssociationExecutionTargetsInput) *SSMDescribeAssociationExecutionTargetsFuture

	DescribeAssociationExecutions(ctx workflow.Context, input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error)
	DescribeAssociationExecutionsAsync(ctx workflow.Context, input *ssm.DescribeAssociationExecutionsInput) *SSMDescribeAssociationExecutionsFuture

	DescribeAutomationExecutions(ctx workflow.Context, input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error)
	DescribeAutomationExecutionsAsync(ctx workflow.Context, input *ssm.DescribeAutomationExecutionsInput) *SSMDescribeAutomationExecutionsFuture

	DescribeAutomationStepExecutions(ctx workflow.Context, input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error)
	DescribeAutomationStepExecutionsAsync(ctx workflow.Context, input *ssm.DescribeAutomationStepExecutionsInput) *SSMDescribeAutomationStepExecutionsFuture

	DescribeAvailablePatches(ctx workflow.Context, input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error)
	DescribeAvailablePatchesAsync(ctx workflow.Context, input *ssm.DescribeAvailablePatchesInput) *SSMDescribeAvailablePatchesFuture

	DescribeDocument(ctx workflow.Context, input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error)
	DescribeDocumentAsync(ctx workflow.Context, input *ssm.DescribeDocumentInput) *SSMDescribeDocumentFuture

	DescribeDocumentPermission(ctx workflow.Context, input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error)
	DescribeDocumentPermissionAsync(ctx workflow.Context, input *ssm.DescribeDocumentPermissionInput) *SSMDescribeDocumentPermissionFuture

	DescribeEffectiveInstanceAssociations(ctx workflow.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error)
	DescribeEffectiveInstanceAssociationsAsync(ctx workflow.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) *SSMDescribeEffectiveInstanceAssociationsFuture

	DescribeEffectivePatchesForPatchBaseline(ctx workflow.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error)
	DescribeEffectivePatchesForPatchBaselineAsync(ctx workflow.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) *SSMDescribeEffectivePatchesForPatchBaselineFuture

	DescribeInstanceAssociationsStatus(ctx workflow.Context, input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error)
	DescribeInstanceAssociationsStatusAsync(ctx workflow.Context, input *ssm.DescribeInstanceAssociationsStatusInput) *SSMDescribeInstanceAssociationsStatusFuture

	DescribeInstanceInformation(ctx workflow.Context, input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error)
	DescribeInstanceInformationAsync(ctx workflow.Context, input *ssm.DescribeInstanceInformationInput) *SSMDescribeInstanceInformationFuture

	DescribeInstancePatchStates(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error)
	DescribeInstancePatchStatesAsync(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesInput) *SSMDescribeInstancePatchStatesFuture

	DescribeInstancePatchStatesForPatchGroup(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error)
	DescribeInstancePatchStatesForPatchGroupAsync(ctx workflow.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) *SSMDescribeInstancePatchStatesForPatchGroupFuture

	DescribeInstancePatches(ctx workflow.Context, input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error)
	DescribeInstancePatchesAsync(ctx workflow.Context, input *ssm.DescribeInstancePatchesInput) *SSMDescribeInstancePatchesFuture

	DescribeInventoryDeletions(ctx workflow.Context, input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error)
	DescribeInventoryDeletionsAsync(ctx workflow.Context, input *ssm.DescribeInventoryDeletionsInput) *SSMDescribeInventoryDeletionsFuture

	DescribeMaintenanceWindowExecutionTaskInvocations(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error)
	DescribeMaintenanceWindowExecutionTaskInvocationsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) *SSMDescribeMaintenanceWindowExecutionTaskInvocationsFuture

	DescribeMaintenanceWindowExecutionTasks(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error)
	DescribeMaintenanceWindowExecutionTasksAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) *SSMDescribeMaintenanceWindowExecutionTasksFuture

	DescribeMaintenanceWindowExecutions(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error)
	DescribeMaintenanceWindowExecutionsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) *SSMDescribeMaintenanceWindowExecutionsFuture

	DescribeMaintenanceWindowSchedule(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error)
	DescribeMaintenanceWindowScheduleAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) *SSMDescribeMaintenanceWindowScheduleFuture

	DescribeMaintenanceWindowTargets(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error)
	DescribeMaintenanceWindowTargetsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) *SSMDescribeMaintenanceWindowTargetsFuture

	DescribeMaintenanceWindowTasks(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error)
	DescribeMaintenanceWindowTasksAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowTasksInput) *SSMDescribeMaintenanceWindowTasksFuture

	DescribeMaintenanceWindows(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error)
	DescribeMaintenanceWindowsAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsInput) *SSMDescribeMaintenanceWindowsFuture

	DescribeMaintenanceWindowsForTarget(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error)
	DescribeMaintenanceWindowsForTargetAsync(ctx workflow.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) *SSMDescribeMaintenanceWindowsForTargetFuture

	DescribeOpsItems(ctx workflow.Context, input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error)
	DescribeOpsItemsAsync(ctx workflow.Context, input *ssm.DescribeOpsItemsInput) *SSMDescribeOpsItemsFuture

	DescribeParameters(ctx workflow.Context, input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error)
	DescribeParametersAsync(ctx workflow.Context, input *ssm.DescribeParametersInput) *SSMDescribeParametersFuture

	DescribePatchBaselines(ctx workflow.Context, input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error)
	DescribePatchBaselinesAsync(ctx workflow.Context, input *ssm.DescribePatchBaselinesInput) *SSMDescribePatchBaselinesFuture

	DescribePatchGroupState(ctx workflow.Context, input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error)
	DescribePatchGroupStateAsync(ctx workflow.Context, input *ssm.DescribePatchGroupStateInput) *SSMDescribePatchGroupStateFuture

	DescribePatchGroups(ctx workflow.Context, input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error)
	DescribePatchGroupsAsync(ctx workflow.Context, input *ssm.DescribePatchGroupsInput) *SSMDescribePatchGroupsFuture

	DescribePatchProperties(ctx workflow.Context, input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error)
	DescribePatchPropertiesAsync(ctx workflow.Context, input *ssm.DescribePatchPropertiesInput) *SSMDescribePatchPropertiesFuture

	DescribeSessions(ctx workflow.Context, input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error)
	DescribeSessionsAsync(ctx workflow.Context, input *ssm.DescribeSessionsInput) *SSMDescribeSessionsFuture

	GetAutomationExecution(ctx workflow.Context, input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error)
	GetAutomationExecutionAsync(ctx workflow.Context, input *ssm.GetAutomationExecutionInput) *SSMGetAutomationExecutionFuture

	GetCalendarState(ctx workflow.Context, input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error)
	GetCalendarStateAsync(ctx workflow.Context, input *ssm.GetCalendarStateInput) *SSMGetCalendarStateFuture

	GetCommandInvocation(ctx workflow.Context, input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error)
	GetCommandInvocationAsync(ctx workflow.Context, input *ssm.GetCommandInvocationInput) *SSMGetCommandInvocationFuture

	GetConnectionStatus(ctx workflow.Context, input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error)
	GetConnectionStatusAsync(ctx workflow.Context, input *ssm.GetConnectionStatusInput) *SSMGetConnectionStatusFuture

	GetDefaultPatchBaseline(ctx workflow.Context, input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error)
	GetDefaultPatchBaselineAsync(ctx workflow.Context, input *ssm.GetDefaultPatchBaselineInput) *SSMGetDefaultPatchBaselineFuture

	GetDeployablePatchSnapshotForInstance(ctx workflow.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error)
	GetDeployablePatchSnapshotForInstanceAsync(ctx workflow.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) *SSMGetDeployablePatchSnapshotForInstanceFuture

	GetDocument(ctx workflow.Context, input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error)
	GetDocumentAsync(ctx workflow.Context, input *ssm.GetDocumentInput) *SSMGetDocumentFuture

	GetInventory(ctx workflow.Context, input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error)
	GetInventoryAsync(ctx workflow.Context, input *ssm.GetInventoryInput) *SSMGetInventoryFuture

	GetInventorySchema(ctx workflow.Context, input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error)
	GetInventorySchemaAsync(ctx workflow.Context, input *ssm.GetInventorySchemaInput) *SSMGetInventorySchemaFuture

	GetMaintenanceWindow(ctx workflow.Context, input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error)
	GetMaintenanceWindowAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowInput) *SSMGetMaintenanceWindowFuture

	GetMaintenanceWindowExecution(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error)
	GetMaintenanceWindowExecutionAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionInput) *SSMGetMaintenanceWindowExecutionFuture

	GetMaintenanceWindowExecutionTask(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error)
	GetMaintenanceWindowExecutionTaskAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) *SSMGetMaintenanceWindowExecutionTaskFuture

	GetMaintenanceWindowExecutionTaskInvocation(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error)
	GetMaintenanceWindowExecutionTaskInvocationAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) *SSMGetMaintenanceWindowExecutionTaskInvocationFuture

	GetMaintenanceWindowTask(ctx workflow.Context, input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error)
	GetMaintenanceWindowTaskAsync(ctx workflow.Context, input *ssm.GetMaintenanceWindowTaskInput) *SSMGetMaintenanceWindowTaskFuture

	GetOpsItem(ctx workflow.Context, input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error)
	GetOpsItemAsync(ctx workflow.Context, input *ssm.GetOpsItemInput) *SSMGetOpsItemFuture

	GetOpsSummary(ctx workflow.Context, input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error)
	GetOpsSummaryAsync(ctx workflow.Context, input *ssm.GetOpsSummaryInput) *SSMGetOpsSummaryFuture

	GetParameter(ctx workflow.Context, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error)
	GetParameterAsync(ctx workflow.Context, input *ssm.GetParameterInput) *SSMGetParameterFuture

	GetParameterHistory(ctx workflow.Context, input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error)
	GetParameterHistoryAsync(ctx workflow.Context, input *ssm.GetParameterHistoryInput) *SSMGetParameterHistoryFuture

	GetParameters(ctx workflow.Context, input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error)
	GetParametersAsync(ctx workflow.Context, input *ssm.GetParametersInput) *SSMGetParametersFuture

	GetParametersByPath(ctx workflow.Context, input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error)
	GetParametersByPathAsync(ctx workflow.Context, input *ssm.GetParametersByPathInput) *SSMGetParametersByPathFuture

	GetPatchBaseline(ctx workflow.Context, input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error)
	GetPatchBaselineAsync(ctx workflow.Context, input *ssm.GetPatchBaselineInput) *SSMGetPatchBaselineFuture

	GetPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error)
	GetPatchBaselineForPatchGroupAsync(ctx workflow.Context, input *ssm.GetPatchBaselineForPatchGroupInput) *SSMGetPatchBaselineForPatchGroupFuture

	GetServiceSetting(ctx workflow.Context, input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error)
	GetServiceSettingAsync(ctx workflow.Context, input *ssm.GetServiceSettingInput) *SSMGetServiceSettingFuture

	LabelParameterVersion(ctx workflow.Context, input *ssm.LabelParameterVersionInput) (*ssm.LabelParameterVersionOutput, error)
	LabelParameterVersionAsync(ctx workflow.Context, input *ssm.LabelParameterVersionInput) *SSMLabelParameterVersionFuture

	ListAssociationVersions(ctx workflow.Context, input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error)
	ListAssociationVersionsAsync(ctx workflow.Context, input *ssm.ListAssociationVersionsInput) *SSMListAssociationVersionsFuture

	ListAssociations(ctx workflow.Context, input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error)
	ListAssociationsAsync(ctx workflow.Context, input *ssm.ListAssociationsInput) *SSMListAssociationsFuture

	ListCommandInvocations(ctx workflow.Context, input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error)
	ListCommandInvocationsAsync(ctx workflow.Context, input *ssm.ListCommandInvocationsInput) *SSMListCommandInvocationsFuture

	ListCommands(ctx workflow.Context, input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error)
	ListCommandsAsync(ctx workflow.Context, input *ssm.ListCommandsInput) *SSMListCommandsFuture

	ListComplianceItems(ctx workflow.Context, input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error)
	ListComplianceItemsAsync(ctx workflow.Context, input *ssm.ListComplianceItemsInput) *SSMListComplianceItemsFuture

	ListComplianceSummaries(ctx workflow.Context, input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error)
	ListComplianceSummariesAsync(ctx workflow.Context, input *ssm.ListComplianceSummariesInput) *SSMListComplianceSummariesFuture

	ListDocumentVersions(ctx workflow.Context, input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error)
	ListDocumentVersionsAsync(ctx workflow.Context, input *ssm.ListDocumentVersionsInput) *SSMListDocumentVersionsFuture

	ListDocuments(ctx workflow.Context, input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error)
	ListDocumentsAsync(ctx workflow.Context, input *ssm.ListDocumentsInput) *SSMListDocumentsFuture

	ListInventoryEntries(ctx workflow.Context, input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error)
	ListInventoryEntriesAsync(ctx workflow.Context, input *ssm.ListInventoryEntriesInput) *SSMListInventoryEntriesFuture

	ListResourceComplianceSummaries(ctx workflow.Context, input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error)
	ListResourceComplianceSummariesAsync(ctx workflow.Context, input *ssm.ListResourceComplianceSummariesInput) *SSMListResourceComplianceSummariesFuture

	ListResourceDataSync(ctx workflow.Context, input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error)
	ListResourceDataSyncAsync(ctx workflow.Context, input *ssm.ListResourceDataSyncInput) *SSMListResourceDataSyncFuture

	ListTagsForResource(ctx workflow.Context, input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *ssm.ListTagsForResourceInput) *SSMListTagsForResourceFuture

	ModifyDocumentPermission(ctx workflow.Context, input *ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error)
	ModifyDocumentPermissionAsync(ctx workflow.Context, input *ssm.ModifyDocumentPermissionInput) *SSMModifyDocumentPermissionFuture

	PutComplianceItems(ctx workflow.Context, input *ssm.PutComplianceItemsInput) (*ssm.PutComplianceItemsOutput, error)
	PutComplianceItemsAsync(ctx workflow.Context, input *ssm.PutComplianceItemsInput) *SSMPutComplianceItemsFuture

	PutInventory(ctx workflow.Context, input *ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error)
	PutInventoryAsync(ctx workflow.Context, input *ssm.PutInventoryInput) *SSMPutInventoryFuture

	PutParameter(ctx workflow.Context, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error)
	PutParameterAsync(ctx workflow.Context, input *ssm.PutParameterInput) *SSMPutParameterFuture

	RegisterDefaultPatchBaseline(ctx workflow.Context, input *ssm.RegisterDefaultPatchBaselineInput) (*ssm.RegisterDefaultPatchBaselineOutput, error)
	RegisterDefaultPatchBaselineAsync(ctx workflow.Context, input *ssm.RegisterDefaultPatchBaselineInput) *SSMRegisterDefaultPatchBaselineFuture

	RegisterPatchBaselineForPatchGroup(ctx workflow.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error)
	RegisterPatchBaselineForPatchGroupAsync(ctx workflow.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) *SSMRegisterPatchBaselineForPatchGroupFuture

	RegisterTargetWithMaintenanceWindow(ctx workflow.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error)
	RegisterTargetWithMaintenanceWindowAsync(ctx workflow.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) *SSMRegisterTargetWithMaintenanceWindowFuture

	RegisterTaskWithMaintenanceWindow(ctx workflow.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error)
	RegisterTaskWithMaintenanceWindowAsync(ctx workflow.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) *SSMRegisterTaskWithMaintenanceWindowFuture

	RemoveTagsFromResource(ctx workflow.Context, input *ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *ssm.RemoveTagsFromResourceInput) *SSMRemoveTagsFromResourceFuture

	ResetServiceSetting(ctx workflow.Context, input *ssm.ResetServiceSettingInput) (*ssm.ResetServiceSettingOutput, error)
	ResetServiceSettingAsync(ctx workflow.Context, input *ssm.ResetServiceSettingInput) *SSMResetServiceSettingFuture

	ResumeSession(ctx workflow.Context, input *ssm.ResumeSessionInput) (*ssm.ResumeSessionOutput, error)
	ResumeSessionAsync(ctx workflow.Context, input *ssm.ResumeSessionInput) *SSMResumeSessionFuture

	SendAutomationSignal(ctx workflow.Context, input *ssm.SendAutomationSignalInput) (*ssm.SendAutomationSignalOutput, error)
	SendAutomationSignalAsync(ctx workflow.Context, input *ssm.SendAutomationSignalInput) *SSMSendAutomationSignalFuture

	SendCommand(ctx workflow.Context, input *ssm.SendCommandInput) (*ssm.SendCommandOutput, error)
	SendCommandAsync(ctx workflow.Context, input *ssm.SendCommandInput) *SSMSendCommandFuture

	StartAssociationsOnce(ctx workflow.Context, input *ssm.StartAssociationsOnceInput) (*ssm.StartAssociationsOnceOutput, error)
	StartAssociationsOnceAsync(ctx workflow.Context, input *ssm.StartAssociationsOnceInput) *SSMStartAssociationsOnceFuture

	StartAutomationExecution(ctx workflow.Context, input *ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error)
	StartAutomationExecutionAsync(ctx workflow.Context, input *ssm.StartAutomationExecutionInput) *SSMStartAutomationExecutionFuture

	StartSession(ctx workflow.Context, input *ssm.StartSessionInput) (*ssm.StartSessionOutput, error)
	StartSessionAsync(ctx workflow.Context, input *ssm.StartSessionInput) *SSMStartSessionFuture

	StopAutomationExecution(ctx workflow.Context, input *ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error)
	StopAutomationExecutionAsync(ctx workflow.Context, input *ssm.StopAutomationExecutionInput) *SSMStopAutomationExecutionFuture

	TerminateSession(ctx workflow.Context, input *ssm.TerminateSessionInput) (*ssm.TerminateSessionOutput, error)
	TerminateSessionAsync(ctx workflow.Context, input *ssm.TerminateSessionInput) *SSMTerminateSessionFuture

	UpdateAssociation(ctx workflow.Context, input *ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error)
	UpdateAssociationAsync(ctx workflow.Context, input *ssm.UpdateAssociationInput) *SSMUpdateAssociationFuture

	UpdateAssociationStatus(ctx workflow.Context, input *ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error)
	UpdateAssociationStatusAsync(ctx workflow.Context, input *ssm.UpdateAssociationStatusInput) *SSMUpdateAssociationStatusFuture

	UpdateDocument(ctx workflow.Context, input *ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error)
	UpdateDocumentAsync(ctx workflow.Context, input *ssm.UpdateDocumentInput) *SSMUpdateDocumentFuture

	UpdateDocumentDefaultVersion(ctx workflow.Context, input *ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error)
	UpdateDocumentDefaultVersionAsync(ctx workflow.Context, input *ssm.UpdateDocumentDefaultVersionInput) *SSMUpdateDocumentDefaultVersionFuture

	UpdateMaintenanceWindow(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error)
	UpdateMaintenanceWindowAsync(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowInput) *SSMUpdateMaintenanceWindowFuture

	UpdateMaintenanceWindowTarget(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTargetInput) (*ssm.UpdateMaintenanceWindowTargetOutput, error)
	UpdateMaintenanceWindowTargetAsync(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTargetInput) *SSMUpdateMaintenanceWindowTargetFuture

	UpdateMaintenanceWindowTask(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTaskInput) (*ssm.UpdateMaintenanceWindowTaskOutput, error)
	UpdateMaintenanceWindowTaskAsync(ctx workflow.Context, input *ssm.UpdateMaintenanceWindowTaskInput) *SSMUpdateMaintenanceWindowTaskFuture

	UpdateManagedInstanceRole(ctx workflow.Context, input *ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error)
	UpdateManagedInstanceRoleAsync(ctx workflow.Context, input *ssm.UpdateManagedInstanceRoleInput) *SSMUpdateManagedInstanceRoleFuture

	UpdateOpsItem(ctx workflow.Context, input *ssm.UpdateOpsItemInput) (*ssm.UpdateOpsItemOutput, error)
	UpdateOpsItemAsync(ctx workflow.Context, input *ssm.UpdateOpsItemInput) *SSMUpdateOpsItemFuture

	UpdatePatchBaseline(ctx workflow.Context, input *ssm.UpdatePatchBaselineInput) (*ssm.UpdatePatchBaselineOutput, error)
	UpdatePatchBaselineAsync(ctx workflow.Context, input *ssm.UpdatePatchBaselineInput) *SSMUpdatePatchBaselineFuture

	UpdateResourceDataSync(ctx workflow.Context, input *ssm.UpdateResourceDataSyncInput) (*ssm.UpdateResourceDataSyncOutput, error)
	UpdateResourceDataSyncAsync(ctx workflow.Context, input *ssm.UpdateResourceDataSyncInput) *SSMUpdateResourceDataSyncFuture

	UpdateServiceSetting(ctx workflow.Context, input *ssm.UpdateServiceSettingInput) (*ssm.UpdateServiceSettingOutput, error)
	UpdateServiceSettingAsync(ctx workflow.Context, input *ssm.UpdateServiceSettingInput) *SSMUpdateServiceSettingFuture

	WaitUntilCommandExecuted(ctx workflow.Context, input *ssm.GetCommandInvocationInput) error
	WaitUntilCommandExecutedAsync(ctx workflow.Context, input *ssm.GetCommandInvocationInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
