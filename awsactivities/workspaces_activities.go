
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type WorkSpacesActivities struct {
    client workspacesiface.WorkSpacesAPI
}

func NewWorkSpacesActivities(session *session.Session, config ...*aws.Config) *WorkSpacesActivities {
    client := workspaces.New(session, config...)
    return &WorkSpacesActivities{client: client}
}

func (a *WorkSpacesActivities) AssociateIpGroups(ctx context.Context, input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error) {
    return a.client.AssociateIpGroupsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) AuthorizeIpRules(ctx context.Context, input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error) {
    return a.client.AuthorizeIpRulesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) CopyWorkspaceImage(ctx context.Context, input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error) {
    return a.client.CopyWorkspaceImageWithContext(ctx, input)
}

func (a *WorkSpacesActivities) CreateIpGroup(ctx context.Context, input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error) {
    return a.client.CreateIpGroupWithContext(ctx, input)
}

func (a *WorkSpacesActivities) CreateTags(ctx context.Context, input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error) {
    return a.client.CreateTagsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) CreateWorkspaces(ctx context.Context, input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error) {
    return a.client.CreateWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DeleteIpGroup(ctx context.Context, input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error) {
    return a.client.DeleteIpGroupWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DeleteTags(ctx context.Context, input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error) {
    return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DeleteWorkspaceImage(ctx context.Context, input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error) {
    return a.client.DeleteWorkspaceImageWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DeregisterWorkspaceDirectory(ctx context.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error) {
    return a.client.DeregisterWorkspaceDirectoryWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeAccount(ctx context.Context, input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error) {
    return a.client.DescribeAccountWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeAccountModifications(ctx context.Context, input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error) {
    return a.client.DescribeAccountModificationsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeClientProperties(ctx context.Context, input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error) {
    return a.client.DescribeClientPropertiesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeIpGroups(ctx context.Context, input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error) {
    return a.client.DescribeIpGroupsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeTags(ctx context.Context, input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error) {
    return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspaceBundles(ctx context.Context, input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
    return a.client.DescribeWorkspaceBundlesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspaceDirectories(ctx context.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
    return a.client.DescribeWorkspaceDirectoriesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspaceImagePermissions(ctx context.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
    return a.client.DescribeWorkspaceImagePermissionsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspaceImages(ctx context.Context, input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error) {
    return a.client.DescribeWorkspaceImagesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspaceSnapshots(ctx context.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
    return a.client.DescribeWorkspaceSnapshotsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspaces(ctx context.Context, input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error) {
    return a.client.DescribeWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DescribeWorkspacesConnectionStatus(ctx context.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
    return a.client.DescribeWorkspacesConnectionStatusWithContext(ctx, input)
}

func (a *WorkSpacesActivities) DisassociateIpGroups(ctx context.Context, input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error) {
    return a.client.DisassociateIpGroupsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ImportWorkspaceImage(ctx context.Context, input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error) {
    return a.client.ImportWorkspaceImageWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ListAvailableManagementCidrRanges(ctx context.Context, input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
    return a.client.ListAvailableManagementCidrRangesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) MigrateWorkspace(ctx context.Context, input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error) {
    return a.client.MigrateWorkspaceWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifyAccount(ctx context.Context, input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error) {
    return a.client.ModifyAccountWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifyClientProperties(ctx context.Context, input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error) {
    return a.client.ModifyClientPropertiesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifySelfservicePermissions(ctx context.Context, input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error) {
    return a.client.ModifySelfservicePermissionsWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifyWorkspaceAccessProperties(ctx context.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error) {
    return a.client.ModifyWorkspaceAccessPropertiesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifyWorkspaceCreationProperties(ctx context.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error) {
    return a.client.ModifyWorkspaceCreationPropertiesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifyWorkspaceProperties(ctx context.Context, input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error) {
    return a.client.ModifyWorkspacePropertiesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) ModifyWorkspaceState(ctx context.Context, input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error) {
    return a.client.ModifyWorkspaceStateWithContext(ctx, input)
}

func (a *WorkSpacesActivities) RebootWorkspaces(ctx context.Context, input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error) {
    return a.client.RebootWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) RebuildWorkspaces(ctx context.Context, input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error) {
    return a.client.RebuildWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) RegisterWorkspaceDirectory(ctx context.Context, input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error) {
    return a.client.RegisterWorkspaceDirectoryWithContext(ctx, input)
}

func (a *WorkSpacesActivities) RestoreWorkspace(ctx context.Context, input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error) {
    return a.client.RestoreWorkspaceWithContext(ctx, input)
}

func (a *WorkSpacesActivities) RevokeIpRules(ctx context.Context, input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error) {
    return a.client.RevokeIpRulesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) StartWorkspaces(ctx context.Context, input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error) {
    return a.client.StartWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) StopWorkspaces(ctx context.Context, input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error) {
    return a.client.StopWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) TerminateWorkspaces(ctx context.Context, input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error) {
    return a.client.TerminateWorkspacesWithContext(ctx, input)
}

func (a *WorkSpacesActivities) UpdateRulesOfIpGroup(ctx context.Context, input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error) {
    return a.client.UpdateRulesOfIpGroupWithContext(ctx, input)
}

func (a *WorkSpacesActivities) UpdateWorkspaceImagePermission(ctx context.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error) {
    return a.client.UpdateWorkspaceImagePermissionWithContext(ctx, input)
}
