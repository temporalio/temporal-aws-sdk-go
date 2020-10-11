// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

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

	sessionFactory SessionFactory
}

func NewSSMActivities(sess *session.Session, config ...*aws.Config) *SSMActivities {
	client := ssm.New(sess, config...)
	return &SSMActivities{client: client}
}

func NewSSMActivitiesWithSessionFactory(sessionFactory SessionFactory) *SSMActivities {
	return &SSMActivities{sessionFactory: sessionFactory}
}

func (a *SSMActivities) getClient(ctx context.Context) (ssmiface.SSMAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return ssm.New(sess), nil
}

func (a *SSMActivities) AddTagsToResource(ctx context.Context, input *ssm.AddTagsToResourceInput) (*ssm.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *SSMActivities) CancelCommand(ctx context.Context, input *ssm.CancelCommandInput) (*ssm.CancelCommandOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelCommandWithContext(ctx, input)
}

func (a *SSMActivities) CancelMaintenanceWindowExecution(ctx context.Context, input *ssm.CancelMaintenanceWindowExecutionInput) (*ssm.CancelMaintenanceWindowExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelMaintenanceWindowExecutionWithContext(ctx, input)
}

func (a *SSMActivities) CreateActivation(ctx context.Context, input *ssm.CreateActivationInput) (*ssm.CreateActivationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateActivationWithContext(ctx, input)
}

func (a *SSMActivities) CreateAssociation(ctx context.Context, input *ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAssociationWithContext(ctx, input)
}

func (a *SSMActivities) CreateAssociationBatch(ctx context.Context, input *ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAssociationBatchWithContext(ctx, input)
}

func (a *SSMActivities) CreateDocument(ctx context.Context, input *ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDocumentWithContext(ctx, input)
}

func (a *SSMActivities) CreateMaintenanceWindow(ctx context.Context, input *ssm.CreateMaintenanceWindowInput) (*ssm.CreateMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) CreateOpsItem(ctx context.Context, input *ssm.CreateOpsItemInput) (*ssm.CreateOpsItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateOpsItemWithContext(ctx, input)
}

func (a *SSMActivities) CreatePatchBaseline(ctx context.Context, input *ssm.CreatePatchBaselineInput) (*ssm.CreatePatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreatePatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) CreateResourceDataSync(ctx context.Context, input *ssm.CreateResourceDataSyncInput) (*ssm.CreateResourceDataSyncOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) DeleteActivation(ctx context.Context, input *ssm.DeleteActivationInput) (*ssm.DeleteActivationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteActivationWithContext(ctx, input)
}

func (a *SSMActivities) DeleteAssociation(ctx context.Context, input *ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAssociationWithContext(ctx, input)
}

func (a *SSMActivities) DeleteDocument(ctx context.Context, input *ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDocumentWithContext(ctx, input)
}

func (a *SSMActivities) DeleteInventory(ctx context.Context, input *ssm.DeleteInventoryInput) (*ssm.DeleteInventoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DeleteInventoryWithContext(ctx, input)
}

func (a *SSMActivities) DeleteMaintenanceWindow(ctx context.Context, input *ssm.DeleteMaintenanceWindowInput) (*ssm.DeleteMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) DeleteParameter(ctx context.Context, input *ssm.DeleteParameterInput) (*ssm.DeleteParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteParameterWithContext(ctx, input)
}

func (a *SSMActivities) DeleteParameters(ctx context.Context, input *ssm.DeleteParametersInput) (*ssm.DeleteParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteParametersWithContext(ctx, input)
}

func (a *SSMActivities) DeletePatchBaseline(ctx context.Context, input *ssm.DeletePatchBaselineInput) (*ssm.DeletePatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) DeleteResourceDataSync(ctx context.Context, input *ssm.DeleteResourceDataSyncInput) (*ssm.DeleteResourceDataSyncOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterManagedInstance(ctx context.Context, input *ssm.DeregisterManagedInstanceInput) (*ssm.DeregisterManagedInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterManagedInstanceWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterPatchBaselineForPatchGroup(ctx context.Context, input *ssm.DeregisterPatchBaselineForPatchGroupInput) (*ssm.DeregisterPatchBaselineForPatchGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterPatchBaselineForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterTargetFromMaintenanceWindow(ctx context.Context, input *ssm.DeregisterTargetFromMaintenanceWindowInput) (*ssm.DeregisterTargetFromMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterTargetFromMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) DeregisterTaskFromMaintenanceWindow(ctx context.Context, input *ssm.DeregisterTaskFromMaintenanceWindowInput) (*ssm.DeregisterTaskFromMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterTaskFromMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) DescribeActivations(ctx context.Context, input *ssm.DescribeActivationsInput) (*ssm.DescribeActivationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeActivationsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAssociation(ctx context.Context, input *ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssociationWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAssociationExecutionTargets(ctx context.Context, input *ssm.DescribeAssociationExecutionTargetsInput) (*ssm.DescribeAssociationExecutionTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssociationExecutionTargetsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAssociationExecutions(ctx context.Context, input *ssm.DescribeAssociationExecutionsInput) (*ssm.DescribeAssociationExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssociationExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAutomationExecutions(ctx context.Context, input *ssm.DescribeAutomationExecutionsInput) (*ssm.DescribeAutomationExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAutomationExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAutomationStepExecutions(ctx context.Context, input *ssm.DescribeAutomationStepExecutionsInput) (*ssm.DescribeAutomationStepExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAutomationStepExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeAvailablePatches(ctx context.Context, input *ssm.DescribeAvailablePatchesInput) (*ssm.DescribeAvailablePatchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAvailablePatchesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeDocument(ctx context.Context, input *ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDocumentWithContext(ctx, input)
}

func (a *SSMActivities) DescribeDocumentPermission(ctx context.Context, input *ssm.DescribeDocumentPermissionInput) (*ssm.DescribeDocumentPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDocumentPermissionWithContext(ctx, input)
}

func (a *SSMActivities) DescribeEffectiveInstanceAssociations(ctx context.Context, input *ssm.DescribeEffectiveInstanceAssociationsInput) (*ssm.DescribeEffectiveInstanceAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEffectiveInstanceAssociationsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeEffectivePatchesForPatchBaseline(ctx context.Context, input *ssm.DescribeEffectivePatchesForPatchBaselineInput) (*ssm.DescribeEffectivePatchesForPatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEffectivePatchesForPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstanceAssociationsStatus(ctx context.Context, input *ssm.DescribeInstanceAssociationsStatusInput) (*ssm.DescribeInstanceAssociationsStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInstanceAssociationsStatusWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstanceInformation(ctx context.Context, input *ssm.DescribeInstanceInformationInput) (*ssm.DescribeInstanceInformationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInstanceInformationWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstancePatchStates(ctx context.Context, input *ssm.DescribeInstancePatchStatesInput) (*ssm.DescribeInstancePatchStatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInstancePatchStatesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstancePatchStatesForPatchGroup(ctx context.Context, input *ssm.DescribeInstancePatchStatesForPatchGroupInput) (*ssm.DescribeInstancePatchStatesForPatchGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInstancePatchStatesForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInstancePatches(ctx context.Context, input *ssm.DescribeInstancePatchesInput) (*ssm.DescribeInstancePatchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInstancePatchesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeInventoryDeletions(ctx context.Context, input *ssm.DescribeInventoryDeletionsInput) (*ssm.DescribeInventoryDeletionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInventoryDeletionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowExecutionTaskInvocations(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput) (*ssm.DescribeMaintenanceWindowExecutionTaskInvocationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowExecutionTaskInvocationsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowExecutionTasks(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionTasksInput) (*ssm.DescribeMaintenanceWindowExecutionTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowExecutionTasksWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowExecutions(ctx context.Context, input *ssm.DescribeMaintenanceWindowExecutionsInput) (*ssm.DescribeMaintenanceWindowExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowExecutionsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowSchedule(ctx context.Context, input *ssm.DescribeMaintenanceWindowScheduleInput) (*ssm.DescribeMaintenanceWindowScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowScheduleWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowTargets(ctx context.Context, input *ssm.DescribeMaintenanceWindowTargetsInput) (*ssm.DescribeMaintenanceWindowTargetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowTargetsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowTasks(ctx context.Context, input *ssm.DescribeMaintenanceWindowTasksInput) (*ssm.DescribeMaintenanceWindowTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowTasksWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindows(ctx context.Context, input *ssm.DescribeMaintenanceWindowsInput) (*ssm.DescribeMaintenanceWindowsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeMaintenanceWindowsForTarget(ctx context.Context, input *ssm.DescribeMaintenanceWindowsForTargetInput) (*ssm.DescribeMaintenanceWindowsForTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceWindowsForTargetWithContext(ctx, input)
}

func (a *SSMActivities) DescribeOpsItems(ctx context.Context, input *ssm.DescribeOpsItemsInput) (*ssm.DescribeOpsItemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeOpsItemsWithContext(ctx, input)
}

func (a *SSMActivities) DescribeParameters(ctx context.Context, input *ssm.DescribeParametersInput) (*ssm.DescribeParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeParametersWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchBaselines(ctx context.Context, input *ssm.DescribePatchBaselinesInput) (*ssm.DescribePatchBaselinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePatchBaselinesWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchGroupState(ctx context.Context, input *ssm.DescribePatchGroupStateInput) (*ssm.DescribePatchGroupStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePatchGroupStateWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchGroups(ctx context.Context, input *ssm.DescribePatchGroupsInput) (*ssm.DescribePatchGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePatchGroupsWithContext(ctx, input)
}

func (a *SSMActivities) DescribePatchProperties(ctx context.Context, input *ssm.DescribePatchPropertiesInput) (*ssm.DescribePatchPropertiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePatchPropertiesWithContext(ctx, input)
}

func (a *SSMActivities) DescribeSessions(ctx context.Context, input *ssm.DescribeSessionsInput) (*ssm.DescribeSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSessionsWithContext(ctx, input)
}

func (a *SSMActivities) GetAutomationExecution(ctx context.Context, input *ssm.GetAutomationExecutionInput) (*ssm.GetAutomationExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAutomationExecutionWithContext(ctx, input)
}

func (a *SSMActivities) GetCalendarState(ctx context.Context, input *ssm.GetCalendarStateInput) (*ssm.GetCalendarStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCalendarStateWithContext(ctx, input)
}

func (a *SSMActivities) GetCommandInvocation(ctx context.Context, input *ssm.GetCommandInvocationInput) (*ssm.GetCommandInvocationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCommandInvocationWithContext(ctx, input)
}

func (a *SSMActivities) GetConnectionStatus(ctx context.Context, input *ssm.GetConnectionStatusInput) (*ssm.GetConnectionStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConnectionStatusWithContext(ctx, input)
}

func (a *SSMActivities) GetDefaultPatchBaseline(ctx context.Context, input *ssm.GetDefaultPatchBaselineInput) (*ssm.GetDefaultPatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDefaultPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) GetDeployablePatchSnapshotForInstance(ctx context.Context, input *ssm.GetDeployablePatchSnapshotForInstanceInput) (*ssm.GetDeployablePatchSnapshotForInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeployablePatchSnapshotForInstanceWithContext(ctx, input)
}

func (a *SSMActivities) GetDocument(ctx context.Context, input *ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDocumentWithContext(ctx, input)
}

func (a *SSMActivities) GetInventory(ctx context.Context, input *ssm.GetInventoryInput) (*ssm.GetInventoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInventoryWithContext(ctx, input)
}

func (a *SSMActivities) GetInventorySchema(ctx context.Context, input *ssm.GetInventorySchemaInput) (*ssm.GetInventorySchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInventorySchemaWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindow(ctx context.Context, input *ssm.GetMaintenanceWindowInput) (*ssm.GetMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowExecution(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionInput) (*ssm.GetMaintenanceWindowExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMaintenanceWindowExecutionWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowExecutionTask(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionTaskInput) (*ssm.GetMaintenanceWindowExecutionTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMaintenanceWindowExecutionTaskWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, input *ssm.GetMaintenanceWindowExecutionTaskInvocationInput) (*ssm.GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMaintenanceWindowExecutionTaskInvocationWithContext(ctx, input)
}

func (a *SSMActivities) GetMaintenanceWindowTask(ctx context.Context, input *ssm.GetMaintenanceWindowTaskInput) (*ssm.GetMaintenanceWindowTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMaintenanceWindowTaskWithContext(ctx, input)
}

func (a *SSMActivities) GetOpsItem(ctx context.Context, input *ssm.GetOpsItemInput) (*ssm.GetOpsItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOpsItemWithContext(ctx, input)
}

func (a *SSMActivities) GetOpsSummary(ctx context.Context, input *ssm.GetOpsSummaryInput) (*ssm.GetOpsSummaryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOpsSummaryWithContext(ctx, input)
}

func (a *SSMActivities) GetParameter(ctx context.Context, input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetParameterWithContext(ctx, input)
}

func (a *SSMActivities) GetParameterHistory(ctx context.Context, input *ssm.GetParameterHistoryInput) (*ssm.GetParameterHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetParameterHistoryWithContext(ctx, input)
}

func (a *SSMActivities) GetParameters(ctx context.Context, input *ssm.GetParametersInput) (*ssm.GetParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetParametersWithContext(ctx, input)
}

func (a *SSMActivities) GetParametersByPath(ctx context.Context, input *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetParametersByPathWithContext(ctx, input)
}

func (a *SSMActivities) GetPatchBaseline(ctx context.Context, input *ssm.GetPatchBaselineInput) (*ssm.GetPatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) GetPatchBaselineForPatchGroup(ctx context.Context, input *ssm.GetPatchBaselineForPatchGroupInput) (*ssm.GetPatchBaselineForPatchGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPatchBaselineForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) GetServiceSetting(ctx context.Context, input *ssm.GetServiceSettingInput) (*ssm.GetServiceSettingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetServiceSettingWithContext(ctx, input)
}

func (a *SSMActivities) LabelParameterVersion(ctx context.Context, input *ssm.LabelParameterVersionInput) (*ssm.LabelParameterVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.LabelParameterVersionWithContext(ctx, input)
}

func (a *SSMActivities) ListAssociationVersions(ctx context.Context, input *ssm.ListAssociationVersionsInput) (*ssm.ListAssociationVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssociationVersionsWithContext(ctx, input)
}

func (a *SSMActivities) ListAssociations(ctx context.Context, input *ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssociationsWithContext(ctx, input)
}

func (a *SSMActivities) ListCommandInvocations(ctx context.Context, input *ssm.ListCommandInvocationsInput) (*ssm.ListCommandInvocationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCommandInvocationsWithContext(ctx, input)
}

func (a *SSMActivities) ListCommands(ctx context.Context, input *ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCommandsWithContext(ctx, input)
}

func (a *SSMActivities) ListComplianceItems(ctx context.Context, input *ssm.ListComplianceItemsInput) (*ssm.ListComplianceItemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListComplianceItemsWithContext(ctx, input)
}

func (a *SSMActivities) ListComplianceSummaries(ctx context.Context, input *ssm.ListComplianceSummariesInput) (*ssm.ListComplianceSummariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListComplianceSummariesWithContext(ctx, input)
}

func (a *SSMActivities) ListDocumentVersions(ctx context.Context, input *ssm.ListDocumentVersionsInput) (*ssm.ListDocumentVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDocumentVersionsWithContext(ctx, input)
}

func (a *SSMActivities) ListDocuments(ctx context.Context, input *ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDocumentsWithContext(ctx, input)
}

func (a *SSMActivities) ListInventoryEntries(ctx context.Context, input *ssm.ListInventoryEntriesInput) (*ssm.ListInventoryEntriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInventoryEntriesWithContext(ctx, input)
}

func (a *SSMActivities) ListResourceComplianceSummaries(ctx context.Context, input *ssm.ListResourceComplianceSummariesInput) (*ssm.ListResourceComplianceSummariesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourceComplianceSummariesWithContext(ctx, input)
}

func (a *SSMActivities) ListResourceDataSync(ctx context.Context, input *ssm.ListResourceDataSyncInput) (*ssm.ListResourceDataSyncOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) ListTagsForResource(ctx context.Context, input *ssm.ListTagsForResourceInput) (*ssm.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SSMActivities) ModifyDocumentPermission(ctx context.Context, input *ssm.ModifyDocumentPermissionInput) (*ssm.ModifyDocumentPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDocumentPermissionWithContext(ctx, input)
}

func (a *SSMActivities) PutComplianceItems(ctx context.Context, input *ssm.PutComplianceItemsInput) (*ssm.PutComplianceItemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutComplianceItemsWithContext(ctx, input)
}

func (a *SSMActivities) PutInventory(ctx context.Context, input *ssm.PutInventoryInput) (*ssm.PutInventoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutInventoryWithContext(ctx, input)
}

func (a *SSMActivities) PutParameter(ctx context.Context, input *ssm.PutParameterInput) (*ssm.PutParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutParameterWithContext(ctx, input)
}

func (a *SSMActivities) RegisterDefaultPatchBaseline(ctx context.Context, input *ssm.RegisterDefaultPatchBaselineInput) (*ssm.RegisterDefaultPatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterDefaultPatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) RegisterPatchBaselineForPatchGroup(ctx context.Context, input *ssm.RegisterPatchBaselineForPatchGroupInput) (*ssm.RegisterPatchBaselineForPatchGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterPatchBaselineForPatchGroupWithContext(ctx, input)
}

func (a *SSMActivities) RegisterTargetWithMaintenanceWindow(ctx context.Context, input *ssm.RegisterTargetWithMaintenanceWindowInput) (*ssm.RegisterTargetWithMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.RegisterTargetWithMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) RegisterTaskWithMaintenanceWindow(ctx context.Context, input *ssm.RegisterTaskWithMaintenanceWindowInput) (*ssm.RegisterTaskWithMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.RegisterTaskWithMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) RemoveTagsFromResource(ctx context.Context, input *ssm.RemoveTagsFromResourceInput) (*ssm.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *SSMActivities) ResetServiceSetting(ctx context.Context, input *ssm.ResetServiceSettingInput) (*ssm.ResetServiceSettingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetServiceSettingWithContext(ctx, input)
}

func (a *SSMActivities) ResumeSession(ctx context.Context, input *ssm.ResumeSessionInput) (*ssm.ResumeSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResumeSessionWithContext(ctx, input)
}

func (a *SSMActivities) SendAutomationSignal(ctx context.Context, input *ssm.SendAutomationSignalInput) (*ssm.SendAutomationSignalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendAutomationSignalWithContext(ctx, input)
}

func (a *SSMActivities) SendCommand(ctx context.Context, input *ssm.SendCommandInput) (*ssm.SendCommandOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendCommandWithContext(ctx, input)
}

func (a *SSMActivities) StartAssociationsOnce(ctx context.Context, input *ssm.StartAssociationsOnceInput) (*ssm.StartAssociationsOnceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartAssociationsOnceWithContext(ctx, input)
}

func (a *SSMActivities) StartAutomationExecution(ctx context.Context, input *ssm.StartAutomationExecutionInput) (*ssm.StartAutomationExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.StartAutomationExecutionWithContext(ctx, input)
}

func (a *SSMActivities) StartSession(ctx context.Context, input *ssm.StartSessionInput) (*ssm.StartSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartSessionWithContext(ctx, input)
}

func (a *SSMActivities) StopAutomationExecution(ctx context.Context, input *ssm.StopAutomationExecutionInput) (*ssm.StopAutomationExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopAutomationExecutionWithContext(ctx, input)
}

func (a *SSMActivities) TerminateSession(ctx context.Context, input *ssm.TerminateSessionInput) (*ssm.TerminateSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TerminateSessionWithContext(ctx, input)
}

func (a *SSMActivities) UpdateAssociation(ctx context.Context, input *ssm.UpdateAssociationInput) (*ssm.UpdateAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAssociationWithContext(ctx, input)
}

func (a *SSMActivities) UpdateAssociationStatus(ctx context.Context, input *ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAssociationStatusWithContext(ctx, input)
}

func (a *SSMActivities) UpdateDocument(ctx context.Context, input *ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDocumentWithContext(ctx, input)
}

func (a *SSMActivities) UpdateDocumentDefaultVersion(ctx context.Context, input *ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDocumentDefaultVersionWithContext(ctx, input)
}

func (a *SSMActivities) UpdateMaintenanceWindow(ctx context.Context, input *ssm.UpdateMaintenanceWindowInput) (*ssm.UpdateMaintenanceWindowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMaintenanceWindowWithContext(ctx, input)
}

func (a *SSMActivities) UpdateMaintenanceWindowTarget(ctx context.Context, input *ssm.UpdateMaintenanceWindowTargetInput) (*ssm.UpdateMaintenanceWindowTargetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMaintenanceWindowTargetWithContext(ctx, input)
}

func (a *SSMActivities) UpdateMaintenanceWindowTask(ctx context.Context, input *ssm.UpdateMaintenanceWindowTaskInput) (*ssm.UpdateMaintenanceWindowTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMaintenanceWindowTaskWithContext(ctx, input)
}

func (a *SSMActivities) UpdateManagedInstanceRole(ctx context.Context, input *ssm.UpdateManagedInstanceRoleInput) (*ssm.UpdateManagedInstanceRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateManagedInstanceRoleWithContext(ctx, input)
}

func (a *SSMActivities) UpdateOpsItem(ctx context.Context, input *ssm.UpdateOpsItemInput) (*ssm.UpdateOpsItemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateOpsItemWithContext(ctx, input)
}

func (a *SSMActivities) UpdatePatchBaseline(ctx context.Context, input *ssm.UpdatePatchBaselineInput) (*ssm.UpdatePatchBaselineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePatchBaselineWithContext(ctx, input)
}

func (a *SSMActivities) UpdateResourceDataSync(ctx context.Context, input *ssm.UpdateResourceDataSyncInput) (*ssm.UpdateResourceDataSyncOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResourceDataSyncWithContext(ctx, input)
}

func (a *SSMActivities) UpdateServiceSetting(ctx context.Context, input *ssm.UpdateServiceSettingInput) (*ssm.UpdateServiceSettingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateServiceSettingWithContext(ctx, input)
}

func (a *SSMActivities) WaitUntilCommandExecuted(ctx context.Context, input *ssm.GetCommandInvocationInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilCommandExecutedWithContext(ctx, input, options...)
	})
}
