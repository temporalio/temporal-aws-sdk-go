package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
)

type WorkSpacesActivities struct {
	client workspacesiface.WorkSpacesAPI
}

func NewWorkSpacesActivities(client workspacesiface.WorkSpacesAPI) *WorkSpacesActivities {
    return &WorkSpacesActivities{client: client}
}


func (a *WorkSpacesActivities) AssociateIpGroups(input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error) {
    return a.client.AssociateIpGroups(input)
}



func (a *WorkSpacesActivities) AuthorizeIpRules(input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error) {
    return a.client.AuthorizeIpRules(input)
}



func (a *WorkSpacesActivities) CopyWorkspaceImage(input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error) {
    return a.client.CopyWorkspaceImage(input)
}



func (a *WorkSpacesActivities) CreateIpGroup(input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error) {
    return a.client.CreateIpGroup(input)
}



func (a *WorkSpacesActivities) CreateTags(input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error) {
    return a.client.CreateTags(input)
}



func (a *WorkSpacesActivities) CreateWorkspaces(input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error) {
    return a.client.CreateWorkspaces(input)
}



func (a *WorkSpacesActivities) DeleteIpGroup(input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error) {
    return a.client.DeleteIpGroup(input)
}



func (a *WorkSpacesActivities) DeleteTags(input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}



func (a *WorkSpacesActivities) DeleteWorkspaceImage(input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error) {
    return a.client.DeleteWorkspaceImage(input)
}



func (a *WorkSpacesActivities) DeregisterWorkspaceDirectory(input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error) {
    return a.client.DeregisterWorkspaceDirectory(input)
}



func (a *WorkSpacesActivities) DescribeAccount(input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error) {
    return a.client.DescribeAccount(input)
}



func (a *WorkSpacesActivities) DescribeAccountModifications(input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error) {
    return a.client.DescribeAccountModifications(input)
}



func (a *WorkSpacesActivities) DescribeClientProperties(input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error) {
    return a.client.DescribeClientProperties(input)
}



func (a *WorkSpacesActivities) DescribeIpGroups(input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error) {
    return a.client.DescribeIpGroups(input)
}



func (a *WorkSpacesActivities) DescribeTags(input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}



func (a *WorkSpacesActivities) DescribeWorkspaceBundles(input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
    return a.client.DescribeWorkspaceBundles(input)
}



func (a *WorkSpacesActivities) DescribeWorkspaceDirectories(input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
    return a.client.DescribeWorkspaceDirectories(input)
}



func (a *WorkSpacesActivities) DescribeWorkspaceImagePermissions(input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
    return a.client.DescribeWorkspaceImagePermissions(input)
}



func (a *WorkSpacesActivities) DescribeWorkspaceImages(input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error) {
    return a.client.DescribeWorkspaceImages(input)
}



func (a *WorkSpacesActivities) DescribeWorkspaceSnapshots(input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
    return a.client.DescribeWorkspaceSnapshots(input)
}



func (a *WorkSpacesActivities) DescribeWorkspaces(input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error) {
    return a.client.DescribeWorkspaces(input)
}



func (a *WorkSpacesActivities) DescribeWorkspacesConnectionStatus(input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
    return a.client.DescribeWorkspacesConnectionStatus(input)
}



func (a *WorkSpacesActivities) DisassociateIpGroups(input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error) {
    return a.client.DisassociateIpGroups(input)
}



func (a *WorkSpacesActivities) ImportWorkspaceImage(input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error) {
    return a.client.ImportWorkspaceImage(input)
}



func (a *WorkSpacesActivities) ListAvailableManagementCidrRanges(input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
    return a.client.ListAvailableManagementCidrRanges(input)
}



func (a *WorkSpacesActivities) MigrateWorkspace(input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error) {
    return a.client.MigrateWorkspace(input)
}



func (a *WorkSpacesActivities) ModifyAccount(input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error) {
    return a.client.ModifyAccount(input)
}



func (a *WorkSpacesActivities) ModifyClientProperties(input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error) {
    return a.client.ModifyClientProperties(input)
}



func (a *WorkSpacesActivities) ModifySelfservicePermissions(input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error) {
    return a.client.ModifySelfservicePermissions(input)
}



func (a *WorkSpacesActivities) ModifyWorkspaceAccessProperties(input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error) {
    return a.client.ModifyWorkspaceAccessProperties(input)
}



func (a *WorkSpacesActivities) ModifyWorkspaceCreationProperties(input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error) {
    return a.client.ModifyWorkspaceCreationProperties(input)
}



func (a *WorkSpacesActivities) ModifyWorkspaceProperties(input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error) {
    return a.client.ModifyWorkspaceProperties(input)
}



func (a *WorkSpacesActivities) ModifyWorkspaceState(input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error) {
    return a.client.ModifyWorkspaceState(input)
}



func (a *WorkSpacesActivities) RebootWorkspaces(input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error) {
    return a.client.RebootWorkspaces(input)
}



func (a *WorkSpacesActivities) RebuildWorkspaces(input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error) {
    return a.client.RebuildWorkspaces(input)
}



func (a *WorkSpacesActivities) RegisterWorkspaceDirectory(input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error) {
    return a.client.RegisterWorkspaceDirectory(input)
}



func (a *WorkSpacesActivities) RestoreWorkspace(input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error) {
    return a.client.RestoreWorkspace(input)
}



func (a *WorkSpacesActivities) RevokeIpRules(input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error) {
    return a.client.RevokeIpRules(input)
}



func (a *WorkSpacesActivities) StartWorkspaces(input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error) {
    return a.client.StartWorkspaces(input)
}



func (a *WorkSpacesActivities) StopWorkspaces(input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error) {
    return a.client.StopWorkspaces(input)
}



func (a *WorkSpacesActivities) TerminateWorkspaces(input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error) {
    return a.client.TerminateWorkspaces(input)
}



func (a *WorkSpacesActivities) UpdateRulesOfIpGroup(input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error) {
    return a.client.UpdateRulesOfIpGroup(input)
}



func (a *WorkSpacesActivities) UpdateWorkspaceImagePermission(input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error) {
    return a.client.UpdateWorkspaceImagePermission(input)
}

