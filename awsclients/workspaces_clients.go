package awsclients

import (
	"github.com/aws/aws-sdk-go/service/workspaces"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type WorkSpacesClient interface {
       AssociateIpGroups(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error)
       AssociateIpGroupsAsync(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) *WorkspacesAssociateIpGroupsResult

       AuthorizeIpRules(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error)
       AuthorizeIpRulesAsync(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) *WorkspacesAuthorizeIpRulesResult

       CopyWorkspaceImage(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error)
       CopyWorkspaceImageAsync(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) *WorkspacesCopyWorkspaceImageResult

       CreateIpGroup(ctx workflow.Context, input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error)
       CreateIpGroupAsync(ctx workflow.Context, input *workspaces.CreateIpGroupInput) *WorkspacesCreateIpGroupResult

       CreateTags(ctx workflow.Context, input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error)
       CreateTagsAsync(ctx workflow.Context, input *workspaces.CreateTagsInput) *WorkspacesCreateTagsResult

       CreateWorkspaces(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)
       CreateWorkspacesAsync(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) *WorkspacesCreateWorkspacesResult

       DeleteIpGroup(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error)
       DeleteIpGroupAsync(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) *WorkspacesDeleteIpGroupResult

       DeleteTags(ctx workflow.Context, input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error)
       DeleteTagsAsync(ctx workflow.Context, input *workspaces.DeleteTagsInput) *WorkspacesDeleteTagsResult

       DeleteWorkspaceImage(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error)
       DeleteWorkspaceImageAsync(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) *WorkspacesDeleteWorkspaceImageResult

       DeregisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error)
       DeregisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) *WorkspacesDeregisterWorkspaceDirectoryResult

       DescribeAccount(ctx workflow.Context, input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error)
       DescribeAccountAsync(ctx workflow.Context, input *workspaces.DescribeAccountInput) *WorkspacesDescribeAccountResult

       DescribeAccountModifications(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error)
       DescribeAccountModificationsAsync(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) *WorkspacesDescribeAccountModificationsResult

       DescribeClientProperties(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error)
       DescribeClientPropertiesAsync(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) *WorkspacesDescribeClientPropertiesResult

       DescribeIpGroups(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error)
       DescribeIpGroupsAsync(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) *WorkspacesDescribeIpGroupsResult

       DescribeTags(ctx workflow.Context, input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error)
       DescribeTagsAsync(ctx workflow.Context, input *workspaces.DescribeTagsInput) *WorkspacesDescribeTagsResult

       DescribeWorkspaceBundles(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)
       DescribeWorkspaceBundlesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) *WorkspacesDescribeWorkspaceBundlesResult

       DescribeWorkspaceDirectories(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
       DescribeWorkspaceDirectoriesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) *WorkspacesDescribeWorkspaceDirectoriesResult

       DescribeWorkspaceImagePermissions(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error)
       DescribeWorkspaceImagePermissionsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) *WorkspacesDescribeWorkspaceImagePermissionsResult

       DescribeWorkspaceImages(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error)
       DescribeWorkspaceImagesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) *WorkspacesDescribeWorkspaceImagesResult

       DescribeWorkspaceSnapshots(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error)
       DescribeWorkspaceSnapshotsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) *WorkspacesDescribeWorkspaceSnapshotsResult

       DescribeWorkspaces(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)
       DescribeWorkspacesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) *WorkspacesDescribeWorkspacesResult

       DescribeWorkspacesConnectionStatus(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
       DescribeWorkspacesConnectionStatusAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) *WorkspacesDescribeWorkspacesConnectionStatusResult

       DisassociateIpGroups(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error)
       DisassociateIpGroupsAsync(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) *WorkspacesDisassociateIpGroupsResult

       ImportWorkspaceImage(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error)
       ImportWorkspaceImageAsync(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) *WorkspacesImportWorkspaceImageResult

       ListAvailableManagementCidrRanges(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error)
       ListAvailableManagementCidrRangesAsync(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) *WorkspacesListAvailableManagementCidrRangesResult

       MigrateWorkspace(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error)
       MigrateWorkspaceAsync(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) *WorkspacesMigrateWorkspaceResult

       ModifyAccount(ctx workflow.Context, input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error)
       ModifyAccountAsync(ctx workflow.Context, input *workspaces.ModifyAccountInput) *WorkspacesModifyAccountResult

       ModifyClientProperties(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error)
       ModifyClientPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) *WorkspacesModifyClientPropertiesResult

       ModifySelfservicePermissions(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error)
       ModifySelfservicePermissionsAsync(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) *WorkspacesModifySelfservicePermissionsResult

       ModifyWorkspaceAccessProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error)
       ModifyWorkspaceAccessPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) *WorkspacesModifyWorkspaceAccessPropertiesResult

       ModifyWorkspaceCreationProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error)
       ModifyWorkspaceCreationPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) *WorkspacesModifyWorkspaceCreationPropertiesResult

       ModifyWorkspaceProperties(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error)
       ModifyWorkspacePropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) *WorkspacesModifyWorkspacePropertiesResult

       ModifyWorkspaceState(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error)
       ModifyWorkspaceStateAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) *WorkspacesModifyWorkspaceStateResult

       RebootWorkspaces(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)
       RebootWorkspacesAsync(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) *WorkspacesRebootWorkspacesResult

       RebuildWorkspaces(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)
       RebuildWorkspacesAsync(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) *WorkspacesRebuildWorkspacesResult

       RegisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error)
       RegisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) *WorkspacesRegisterWorkspaceDirectoryResult

       RestoreWorkspace(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error)
       RestoreWorkspaceAsync(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) *WorkspacesRestoreWorkspaceResult

       RevokeIpRules(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error)
       RevokeIpRulesAsync(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) *WorkspacesRevokeIpRulesResult

       StartWorkspaces(ctx workflow.Context, input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error)
       StartWorkspacesAsync(ctx workflow.Context, input *workspaces.StartWorkspacesInput) *WorkspacesStartWorkspacesResult

       StopWorkspaces(ctx workflow.Context, input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error)
       StopWorkspacesAsync(ctx workflow.Context, input *workspaces.StopWorkspacesInput) *WorkspacesStopWorkspacesResult

       TerminateWorkspaces(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
       TerminateWorkspacesAsync(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) *WorkspacesTerminateWorkspacesResult

       UpdateRulesOfIpGroup(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error)
       UpdateRulesOfIpGroupAsync(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) *WorkspacesUpdateRulesOfIpGroupResult

       UpdateWorkspaceImagePermission(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error)
       UpdateWorkspaceImagePermissionAsync(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) *WorkspacesUpdateWorkspaceImagePermissionResult
}

type WorkspacesAssociateIpGroupsResult struct {
	Result workflow.Future
}

func (r *WorkspacesAssociateIpGroupsResult) Get(ctx workflow.Context) (*workspaces.AssociateIpGroupsOutput, error) {
    var output workspaces.AssociateIpGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesAuthorizeIpRulesResult struct {
	Result workflow.Future
}

func (r *WorkspacesAuthorizeIpRulesResult) Get(ctx workflow.Context) (*workspaces.AuthorizeIpRulesOutput, error) {
    var output workspaces.AuthorizeIpRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesCopyWorkspaceImageResult struct {
	Result workflow.Future
}

func (r *WorkspacesCopyWorkspaceImageResult) Get(ctx workflow.Context) (*workspaces.CopyWorkspaceImageOutput, error) {
    var output workspaces.CopyWorkspaceImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesCreateIpGroupResult struct {
	Result workflow.Future
}

func (r *WorkspacesCreateIpGroupResult) Get(ctx workflow.Context) (*workspaces.CreateIpGroupOutput, error) {
    var output workspaces.CreateIpGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesCreateTagsResult struct {
	Result workflow.Future
}

func (r *WorkspacesCreateTagsResult) Get(ctx workflow.Context) (*workspaces.CreateTagsOutput, error) {
    var output workspaces.CreateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesCreateWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesCreateWorkspacesResult) Get(ctx workflow.Context) (*workspaces.CreateWorkspacesOutput, error) {
    var output workspaces.CreateWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDeleteIpGroupResult struct {
	Result workflow.Future
}

func (r *WorkspacesDeleteIpGroupResult) Get(ctx workflow.Context) (*workspaces.DeleteIpGroupOutput, error) {
    var output workspaces.DeleteIpGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDeleteTagsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDeleteTagsResult) Get(ctx workflow.Context) (*workspaces.DeleteTagsOutput, error) {
    var output workspaces.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDeleteWorkspaceImageResult struct {
	Result workflow.Future
}

func (r *WorkspacesDeleteWorkspaceImageResult) Get(ctx workflow.Context) (*workspaces.DeleteWorkspaceImageOutput, error) {
    var output workspaces.DeleteWorkspaceImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDeregisterWorkspaceDirectoryResult struct {
	Result workflow.Future
}

func (r *WorkspacesDeregisterWorkspaceDirectoryResult) Get(ctx workflow.Context) (*workspaces.DeregisterWorkspaceDirectoryOutput, error) {
    var output workspaces.DeregisterWorkspaceDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeAccountResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeAccountResult) Get(ctx workflow.Context) (*workspaces.DescribeAccountOutput, error) {
    var output workspaces.DescribeAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeAccountModificationsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeAccountModificationsResult) Get(ctx workflow.Context) (*workspaces.DescribeAccountModificationsOutput, error) {
    var output workspaces.DescribeAccountModificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeClientPropertiesResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeClientPropertiesResult) Get(ctx workflow.Context) (*workspaces.DescribeClientPropertiesOutput, error) {
    var output workspaces.DescribeClientPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeIpGroupsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeIpGroupsResult) Get(ctx workflow.Context) (*workspaces.DescribeIpGroupsOutput, error) {
    var output workspaces.DescribeIpGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeTagsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeTagsResult) Get(ctx workflow.Context) (*workspaces.DescribeTagsOutput, error) {
    var output workspaces.DescribeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspaceBundlesResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspaceBundlesResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
    var output workspaces.DescribeWorkspaceBundlesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspaceDirectoriesResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspaceDirectoriesResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
    var output workspaces.DescribeWorkspaceDirectoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspaceImagePermissionsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspaceImagePermissionsResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
    var output workspaces.DescribeWorkspaceImagePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspaceImagesResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspaceImagesResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspaceImagesOutput, error) {
    var output workspaces.DescribeWorkspaceImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspaceSnapshotsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspaceSnapshotsResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
    var output workspaces.DescribeWorkspaceSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspacesResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspacesOutput, error) {
    var output workspaces.DescribeWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDescribeWorkspacesConnectionStatusResult struct {
	Result workflow.Future
}

func (r *WorkspacesDescribeWorkspacesConnectionStatusResult) Get(ctx workflow.Context) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
    var output workspaces.DescribeWorkspacesConnectionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesDisassociateIpGroupsResult struct {
	Result workflow.Future
}

func (r *WorkspacesDisassociateIpGroupsResult) Get(ctx workflow.Context) (*workspaces.DisassociateIpGroupsOutput, error) {
    var output workspaces.DisassociateIpGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesImportWorkspaceImageResult struct {
	Result workflow.Future
}

func (r *WorkspacesImportWorkspaceImageResult) Get(ctx workflow.Context) (*workspaces.ImportWorkspaceImageOutput, error) {
    var output workspaces.ImportWorkspaceImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesListAvailableManagementCidrRangesResult struct {
	Result workflow.Future
}

func (r *WorkspacesListAvailableManagementCidrRangesResult) Get(ctx workflow.Context) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
    var output workspaces.ListAvailableManagementCidrRangesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesMigrateWorkspaceResult struct {
	Result workflow.Future
}

func (r *WorkspacesMigrateWorkspaceResult) Get(ctx workflow.Context) (*workspaces.MigrateWorkspaceOutput, error) {
    var output workspaces.MigrateWorkspaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifyAccountResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifyAccountResult) Get(ctx workflow.Context) (*workspaces.ModifyAccountOutput, error) {
    var output workspaces.ModifyAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifyClientPropertiesResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifyClientPropertiesResult) Get(ctx workflow.Context) (*workspaces.ModifyClientPropertiesOutput, error) {
    var output workspaces.ModifyClientPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifySelfservicePermissionsResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifySelfservicePermissionsResult) Get(ctx workflow.Context) (*workspaces.ModifySelfservicePermissionsOutput, error) {
    var output workspaces.ModifySelfservicePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifyWorkspaceAccessPropertiesResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifyWorkspaceAccessPropertiesResult) Get(ctx workflow.Context) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error) {
    var output workspaces.ModifyWorkspaceAccessPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifyWorkspaceCreationPropertiesResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifyWorkspaceCreationPropertiesResult) Get(ctx workflow.Context) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error) {
    var output workspaces.ModifyWorkspaceCreationPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifyWorkspacePropertiesResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifyWorkspacePropertiesResult) Get(ctx workflow.Context) (*workspaces.ModifyWorkspacePropertiesOutput, error) {
    var output workspaces.ModifyWorkspacePropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesModifyWorkspaceStateResult struct {
	Result workflow.Future
}

func (r *WorkspacesModifyWorkspaceStateResult) Get(ctx workflow.Context) (*workspaces.ModifyWorkspaceStateOutput, error) {
    var output workspaces.ModifyWorkspaceStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesRebootWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesRebootWorkspacesResult) Get(ctx workflow.Context) (*workspaces.RebootWorkspacesOutput, error) {
    var output workspaces.RebootWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesRebuildWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesRebuildWorkspacesResult) Get(ctx workflow.Context) (*workspaces.RebuildWorkspacesOutput, error) {
    var output workspaces.RebuildWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesRegisterWorkspaceDirectoryResult struct {
	Result workflow.Future
}

func (r *WorkspacesRegisterWorkspaceDirectoryResult) Get(ctx workflow.Context) (*workspaces.RegisterWorkspaceDirectoryOutput, error) {
    var output workspaces.RegisterWorkspaceDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesRestoreWorkspaceResult struct {
	Result workflow.Future
}

func (r *WorkspacesRestoreWorkspaceResult) Get(ctx workflow.Context) (*workspaces.RestoreWorkspaceOutput, error) {
    var output workspaces.RestoreWorkspaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesRevokeIpRulesResult struct {
	Result workflow.Future
}

func (r *WorkspacesRevokeIpRulesResult) Get(ctx workflow.Context) (*workspaces.RevokeIpRulesOutput, error) {
    var output workspaces.RevokeIpRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesStartWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesStartWorkspacesResult) Get(ctx workflow.Context) (*workspaces.StartWorkspacesOutput, error) {
    var output workspaces.StartWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesStopWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesStopWorkspacesResult) Get(ctx workflow.Context) (*workspaces.StopWorkspacesOutput, error) {
    var output workspaces.StopWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesTerminateWorkspacesResult struct {
	Result workflow.Future
}

func (r *WorkspacesTerminateWorkspacesResult) Get(ctx workflow.Context) (*workspaces.TerminateWorkspacesOutput, error) {
    var output workspaces.TerminateWorkspacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesUpdateRulesOfIpGroupResult struct {
	Result workflow.Future
}

func (r *WorkspacesUpdateRulesOfIpGroupResult) Get(ctx workflow.Context) (*workspaces.UpdateRulesOfIpGroupOutput, error) {
    var output workspaces.UpdateRulesOfIpGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkspacesUpdateWorkspaceImagePermissionResult struct {
	Result workflow.Future
}

func (r *WorkspacesUpdateWorkspaceImagePermissionResult) Get(ctx workflow.Context) (*workspaces.UpdateWorkspaceImagePermissionOutput, error) {
    var output workspaces.UpdateWorkspaceImagePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WorkSpacesStub struct {
    activities awsactivities.WorkSpacesActivities
}

func NewWorkSpacesStub() WorkSpacesClient {
    return &WorkSpacesStub{}
}

func (a *WorkSpacesStub) AssociateIpGroups(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error) {
    var output workspaces.AssociateIpGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateIpGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) AssociateIpGroupsAsync(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) *WorkspacesAssociateIpGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateIpGroups, input)
    return &WorkspacesAssociateIpGroupsResult{Result: future}
}

func (a *WorkSpacesStub) AuthorizeIpRules(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error) {
    var output workspaces.AuthorizeIpRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeIpRules, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) AuthorizeIpRulesAsync(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) *WorkspacesAuthorizeIpRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AuthorizeIpRules, input)
    return &WorkspacesAuthorizeIpRulesResult{Result: future}
}

func (a *WorkSpacesStub) CopyWorkspaceImage(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error) {
    var output workspaces.CopyWorkspaceImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyWorkspaceImage, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) CopyWorkspaceImageAsync(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) *WorkspacesCopyWorkspaceImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyWorkspaceImage, input)
    return &WorkspacesCopyWorkspaceImageResult{Result: future}
}

func (a *WorkSpacesStub) CreateIpGroup(ctx workflow.Context, input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error) {
    var output workspaces.CreateIpGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIpGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) CreateIpGroupAsync(ctx workflow.Context, input *workspaces.CreateIpGroupInput) *WorkspacesCreateIpGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateIpGroup, input)
    return &WorkspacesCreateIpGroupResult{Result: future}
}

func (a *WorkSpacesStub) CreateTags(ctx workflow.Context, input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error) {
    var output workspaces.CreateTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) CreateTagsAsync(ctx workflow.Context, input *workspaces.CreateTagsInput) *WorkspacesCreateTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input)
    return &WorkspacesCreateTagsResult{Result: future}
}

func (a *WorkSpacesStub) CreateWorkspaces(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error) {
    var output workspaces.CreateWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) CreateWorkspacesAsync(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) *WorkspacesCreateWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateWorkspaces, input)
    return &WorkspacesCreateWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) DeleteIpGroup(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error) {
    var output workspaces.DeleteIpGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIpGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DeleteIpGroupAsync(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) *WorkspacesDeleteIpGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteIpGroup, input)
    return &WorkspacesDeleteIpGroupResult{Result: future}
}

func (a *WorkSpacesStub) DeleteTags(ctx workflow.Context, input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error) {
    var output workspaces.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DeleteTagsAsync(ctx workflow.Context, input *workspaces.DeleteTagsInput) *WorkspacesDeleteTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input)
    return &WorkspacesDeleteTagsResult{Result: future}
}

func (a *WorkSpacesStub) DeleteWorkspaceImage(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error) {
    var output workspaces.DeleteWorkspaceImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkspaceImage, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DeleteWorkspaceImageAsync(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) *WorkspacesDeleteWorkspaceImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkspaceImage, input)
    return &WorkspacesDeleteWorkspaceImageResult{Result: future}
}

func (a *WorkSpacesStub) DeregisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error) {
    var output workspaces.DeregisterWorkspaceDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterWorkspaceDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DeregisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) *WorkspacesDeregisterWorkspaceDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterWorkspaceDirectory, input)
    return &WorkspacesDeregisterWorkspaceDirectoryResult{Result: future}
}

func (a *WorkSpacesStub) DescribeAccount(ctx workflow.Context, input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error) {
    var output workspaces.DescribeAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeAccountAsync(ctx workflow.Context, input *workspaces.DescribeAccountInput) *WorkspacesDescribeAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccount, input)
    return &WorkspacesDescribeAccountResult{Result: future}
}

func (a *WorkSpacesStub) DescribeAccountModifications(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error) {
    var output workspaces.DescribeAccountModificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountModifications, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeAccountModificationsAsync(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) *WorkspacesDescribeAccountModificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountModifications, input)
    return &WorkspacesDescribeAccountModificationsResult{Result: future}
}

func (a *WorkSpacesStub) DescribeClientProperties(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error) {
    var output workspaces.DescribeClientPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClientProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeClientPropertiesAsync(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) *WorkspacesDescribeClientPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClientProperties, input)
    return &WorkspacesDescribeClientPropertiesResult{Result: future}
}

func (a *WorkSpacesStub) DescribeIpGroups(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error) {
    var output workspaces.DescribeIpGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIpGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeIpGroupsAsync(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) *WorkspacesDescribeIpGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIpGroups, input)
    return &WorkspacesDescribeIpGroupsResult{Result: future}
}

func (a *WorkSpacesStub) DescribeTags(ctx workflow.Context, input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error) {
    var output workspaces.DescribeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeTagsAsync(ctx workflow.Context, input *workspaces.DescribeTagsInput) *WorkspacesDescribeTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
    return &WorkspacesDescribeTagsResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspaceBundles(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
    var output workspaces.DescribeWorkspaceBundlesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceBundles, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspaceBundlesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) *WorkspacesDescribeWorkspaceBundlesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceBundles, input)
    return &WorkspacesDescribeWorkspaceBundlesResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspaceDirectories(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
    var output workspaces.DescribeWorkspaceDirectoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceDirectories, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspaceDirectoriesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) *WorkspacesDescribeWorkspaceDirectoriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceDirectories, input)
    return &WorkspacesDescribeWorkspaceDirectoriesResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspaceImagePermissions(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
    var output workspaces.DescribeWorkspaceImagePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceImagePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspaceImagePermissionsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) *WorkspacesDescribeWorkspaceImagePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceImagePermissions, input)
    return &WorkspacesDescribeWorkspaceImagePermissionsResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspaceImages(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error) {
    var output workspaces.DescribeWorkspaceImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceImages, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspaceImagesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) *WorkspacesDescribeWorkspaceImagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceImages, input)
    return &WorkspacesDescribeWorkspaceImagesResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspaceSnapshots(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
    var output workspaces.DescribeWorkspaceSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspaceSnapshotsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) *WorkspacesDescribeWorkspaceSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaceSnapshots, input)
    return &WorkspacesDescribeWorkspaceSnapshotsResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspaces(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error) {
    var output workspaces.DescribeWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspacesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) *WorkspacesDescribeWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspaces, input)
    return &WorkspacesDescribeWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) DescribeWorkspacesConnectionStatus(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
    var output workspaces.DescribeWorkspacesConnectionStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspacesConnectionStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DescribeWorkspacesConnectionStatusAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) *WorkspacesDescribeWorkspacesConnectionStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkspacesConnectionStatus, input)
    return &WorkspacesDescribeWorkspacesConnectionStatusResult{Result: future}
}

func (a *WorkSpacesStub) DisassociateIpGroups(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error) {
    var output workspaces.DisassociateIpGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateIpGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) DisassociateIpGroupsAsync(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) *WorkspacesDisassociateIpGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateIpGroups, input)
    return &WorkspacesDisassociateIpGroupsResult{Result: future}
}

func (a *WorkSpacesStub) ImportWorkspaceImage(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error) {
    var output workspaces.ImportWorkspaceImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportWorkspaceImage, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ImportWorkspaceImageAsync(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) *WorkspacesImportWorkspaceImageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportWorkspaceImage, input)
    return &WorkspacesImportWorkspaceImageResult{Result: future}
}

func (a *WorkSpacesStub) ListAvailableManagementCidrRanges(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
    var output workspaces.ListAvailableManagementCidrRangesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAvailableManagementCidrRanges, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ListAvailableManagementCidrRangesAsync(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) *WorkspacesListAvailableManagementCidrRangesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAvailableManagementCidrRanges, input)
    return &WorkspacesListAvailableManagementCidrRangesResult{Result: future}
}

func (a *WorkSpacesStub) MigrateWorkspace(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error) {
    var output workspaces.MigrateWorkspaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.MigrateWorkspace, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) MigrateWorkspaceAsync(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) *WorkspacesMigrateWorkspaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.MigrateWorkspace, input)
    return &WorkspacesMigrateWorkspaceResult{Result: future}
}

func (a *WorkSpacesStub) ModifyAccount(ctx workflow.Context, input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error) {
    var output workspaces.ModifyAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifyAccountAsync(ctx workflow.Context, input *workspaces.ModifyAccountInput) *WorkspacesModifyAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyAccount, input)
    return &WorkspacesModifyAccountResult{Result: future}
}

func (a *WorkSpacesStub) ModifyClientProperties(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error) {
    var output workspaces.ModifyClientPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyClientProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifyClientPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) *WorkspacesModifyClientPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyClientProperties, input)
    return &WorkspacesModifyClientPropertiesResult{Result: future}
}

func (a *WorkSpacesStub) ModifySelfservicePermissions(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error) {
    var output workspaces.ModifySelfservicePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifySelfservicePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifySelfservicePermissionsAsync(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) *WorkspacesModifySelfservicePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifySelfservicePermissions, input)
    return &WorkspacesModifySelfservicePermissionsResult{Result: future}
}

func (a *WorkSpacesStub) ModifyWorkspaceAccessProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error) {
    var output workspaces.ModifyWorkspaceAccessPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceAccessProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifyWorkspaceAccessPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) *WorkspacesModifyWorkspaceAccessPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceAccessProperties, input)
    return &WorkspacesModifyWorkspaceAccessPropertiesResult{Result: future}
}

func (a *WorkSpacesStub) ModifyWorkspaceCreationProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error) {
    var output workspaces.ModifyWorkspaceCreationPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceCreationProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifyWorkspaceCreationPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) *WorkspacesModifyWorkspaceCreationPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceCreationProperties, input)
    return &WorkspacesModifyWorkspaceCreationPropertiesResult{Result: future}
}

func (a *WorkSpacesStub) ModifyWorkspaceProperties(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error) {
    var output workspaces.ModifyWorkspacePropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifyWorkspacePropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) *WorkspacesModifyWorkspacePropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceProperties, input)
    return &WorkspacesModifyWorkspacePropertiesResult{Result: future}
}

func (a *WorkSpacesStub) ModifyWorkspaceState(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error) {
    var output workspaces.ModifyWorkspaceStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceState, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) ModifyWorkspaceStateAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) *WorkspacesModifyWorkspaceStateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyWorkspaceState, input)
    return &WorkspacesModifyWorkspaceStateResult{Result: future}
}

func (a *WorkSpacesStub) RebootWorkspaces(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error) {
    var output workspaces.RebootWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) RebootWorkspacesAsync(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) *WorkspacesRebootWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebootWorkspaces, input)
    return &WorkspacesRebootWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) RebuildWorkspaces(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error) {
    var output workspaces.RebuildWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebuildWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) RebuildWorkspacesAsync(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) *WorkspacesRebuildWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebuildWorkspaces, input)
    return &WorkspacesRebuildWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) RegisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error) {
    var output workspaces.RegisterWorkspaceDirectoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterWorkspaceDirectory, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) RegisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) *WorkspacesRegisterWorkspaceDirectoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterWorkspaceDirectory, input)
    return &WorkspacesRegisterWorkspaceDirectoryResult{Result: future}
}

func (a *WorkSpacesStub) RestoreWorkspace(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error) {
    var output workspaces.RestoreWorkspaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreWorkspace, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) RestoreWorkspaceAsync(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) *WorkspacesRestoreWorkspaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreWorkspace, input)
    return &WorkspacesRestoreWorkspaceResult{Result: future}
}

func (a *WorkSpacesStub) RevokeIpRules(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error) {
    var output workspaces.RevokeIpRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeIpRules, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) RevokeIpRulesAsync(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) *WorkspacesRevokeIpRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RevokeIpRules, input)
    return &WorkspacesRevokeIpRulesResult{Result: future}
}

func (a *WorkSpacesStub) StartWorkspaces(ctx workflow.Context, input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error) {
    var output workspaces.StartWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) StartWorkspacesAsync(ctx workflow.Context, input *workspaces.StartWorkspacesInput) *WorkspacesStartWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartWorkspaces, input)
    return &WorkspacesStartWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) StopWorkspaces(ctx workflow.Context, input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error) {
    var output workspaces.StopWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) StopWorkspacesAsync(ctx workflow.Context, input *workspaces.StopWorkspacesInput) *WorkspacesStopWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopWorkspaces, input)
    return &WorkspacesStopWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) TerminateWorkspaces(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error) {
    var output workspaces.TerminateWorkspacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateWorkspaces, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) TerminateWorkspacesAsync(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) *WorkspacesTerminateWorkspacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateWorkspaces, input)
    return &WorkspacesTerminateWorkspacesResult{Result: future}
}

func (a *WorkSpacesStub) UpdateRulesOfIpGroup(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error) {
    var output workspaces.UpdateRulesOfIpGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRulesOfIpGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) UpdateRulesOfIpGroupAsync(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) *WorkspacesUpdateRulesOfIpGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRulesOfIpGroup, input)
    return &WorkspacesUpdateRulesOfIpGroupResult{Result: future}
}

func (a *WorkSpacesStub) UpdateWorkspaceImagePermission(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error) {
    var output workspaces.UpdateWorkspaceImagePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkspaceImagePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *WorkSpacesStub) UpdateWorkspaceImagePermissionAsync(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) *WorkspacesUpdateWorkspaceImagePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkspaceImagePermission, input)
    return &WorkspacesUpdateWorkspaceImagePermissionResult{Result: future}
}
