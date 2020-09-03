package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/workdocs/workdocsiface"
)

type WorkDocsActivities struct {
	client workdocsiface.WorkDocsAPI
}

func NewWorkDocsActivities(client workdocsiface.WorkDocsAPI) *WorkDocsActivities {
    return &WorkDocsActivities{client: client}
}


func (a *WorkDocsActivities) AbortDocumentVersionUpload(input *workdocs.AbortDocumentVersionUploadInput) (*workdocs.AbortDocumentVersionUploadOutput, error) {
    return a.client.AbortDocumentVersionUpload(input)
}



func (a *WorkDocsActivities) ActivateUser(input *workdocs.ActivateUserInput) (*workdocs.ActivateUserOutput, error) {
    return a.client.ActivateUser(input)
}



func (a *WorkDocsActivities) AddResourcePermissions(input *workdocs.AddResourcePermissionsInput) (*workdocs.AddResourcePermissionsOutput, error) {
    return a.client.AddResourcePermissions(input)
}



func (a *WorkDocsActivities) CreateComment(input *workdocs.CreateCommentInput) (*workdocs.CreateCommentOutput, error) {
    return a.client.CreateComment(input)
}



func (a *WorkDocsActivities) CreateCustomMetadata(input *workdocs.CreateCustomMetadataInput) (*workdocs.CreateCustomMetadataOutput, error) {
    return a.client.CreateCustomMetadata(input)
}



func (a *WorkDocsActivities) CreateFolder(input *workdocs.CreateFolderInput) (*workdocs.CreateFolderOutput, error) {
    return a.client.CreateFolder(input)
}



func (a *WorkDocsActivities) CreateLabels(input *workdocs.CreateLabelsInput) (*workdocs.CreateLabelsOutput, error) {
    return a.client.CreateLabels(input)
}



func (a *WorkDocsActivities) CreateNotificationSubscription(input *workdocs.CreateNotificationSubscriptionInput) (*workdocs.CreateNotificationSubscriptionOutput, error) {
    return a.client.CreateNotificationSubscription(input)
}



func (a *WorkDocsActivities) CreateUser(input *workdocs.CreateUserInput) (*workdocs.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}



func (a *WorkDocsActivities) DeactivateUser(input *workdocs.DeactivateUserInput) (*workdocs.DeactivateUserOutput, error) {
    return a.client.DeactivateUser(input)
}



func (a *WorkDocsActivities) DeleteComment(input *workdocs.DeleteCommentInput) (*workdocs.DeleteCommentOutput, error) {
    return a.client.DeleteComment(input)
}



func (a *WorkDocsActivities) DeleteCustomMetadata(input *workdocs.DeleteCustomMetadataInput) (*workdocs.DeleteCustomMetadataOutput, error) {
    return a.client.DeleteCustomMetadata(input)
}



func (a *WorkDocsActivities) DeleteDocument(input *workdocs.DeleteDocumentInput) (*workdocs.DeleteDocumentOutput, error) {
    return a.client.DeleteDocument(input)
}



func (a *WorkDocsActivities) DeleteFolder(input *workdocs.DeleteFolderInput) (*workdocs.DeleteFolderOutput, error) {
    return a.client.DeleteFolder(input)
}



func (a *WorkDocsActivities) DeleteFolderContents(input *workdocs.DeleteFolderContentsInput) (*workdocs.DeleteFolderContentsOutput, error) {
    return a.client.DeleteFolderContents(input)
}



func (a *WorkDocsActivities) DeleteLabels(input *workdocs.DeleteLabelsInput) (*workdocs.DeleteLabelsOutput, error) {
    return a.client.DeleteLabels(input)
}



func (a *WorkDocsActivities) DeleteNotificationSubscription(input *workdocs.DeleteNotificationSubscriptionInput) (*workdocs.DeleteNotificationSubscriptionOutput, error) {
    return a.client.DeleteNotificationSubscription(input)
}



func (a *WorkDocsActivities) DeleteUser(input *workdocs.DeleteUserInput) (*workdocs.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}



func (a *WorkDocsActivities) DescribeActivities(input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error) {
    return a.client.DescribeActivities(input)
}



func (a *WorkDocsActivities) DescribeComments(input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error) {
    return a.client.DescribeComments(input)
}



func (a *WorkDocsActivities) DescribeDocumentVersions(input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error) {
    return a.client.DescribeDocumentVersions(input)
}



func (a *WorkDocsActivities) DescribeFolderContents(input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error) {
    return a.client.DescribeFolderContents(input)
}



func (a *WorkDocsActivities) DescribeGroups(input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error) {
    return a.client.DescribeGroups(input)
}



func (a *WorkDocsActivities) DescribeNotificationSubscriptions(input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
    return a.client.DescribeNotificationSubscriptions(input)
}



func (a *WorkDocsActivities) DescribeResourcePermissions(input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error) {
    return a.client.DescribeResourcePermissions(input)
}



func (a *WorkDocsActivities) DescribeRootFolders(input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error) {
    return a.client.DescribeRootFolders(input)
}



func (a *WorkDocsActivities) DescribeUsers(input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error) {
    return a.client.DescribeUsers(input)
}



func (a *WorkDocsActivities) GetCurrentUser(input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error) {
    return a.client.GetCurrentUser(input)
}



func (a *WorkDocsActivities) GetDocument(input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error) {
    return a.client.GetDocument(input)
}



func (a *WorkDocsActivities) GetDocumentPath(input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error) {
    return a.client.GetDocumentPath(input)
}



func (a *WorkDocsActivities) GetDocumentVersion(input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error) {
    return a.client.GetDocumentVersion(input)
}



func (a *WorkDocsActivities) GetFolder(input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error) {
    return a.client.GetFolder(input)
}



func (a *WorkDocsActivities) GetFolderPath(input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error) {
    return a.client.GetFolderPath(input)
}



func (a *WorkDocsActivities) GetResources(input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error) {
    return a.client.GetResources(input)
}



func (a *WorkDocsActivities) InitiateDocumentVersionUpload(input *workdocs.InitiateDocumentVersionUploadInput) (*workdocs.InitiateDocumentVersionUploadOutput, error) {
    return a.client.InitiateDocumentVersionUpload(input)
}



func (a *WorkDocsActivities) RemoveAllResourcePermissions(input *workdocs.RemoveAllResourcePermissionsInput) (*workdocs.RemoveAllResourcePermissionsOutput, error) {
    return a.client.RemoveAllResourcePermissions(input)
}



func (a *WorkDocsActivities) RemoveResourcePermission(input *workdocs.RemoveResourcePermissionInput) (*workdocs.RemoveResourcePermissionOutput, error) {
    return a.client.RemoveResourcePermission(input)
}



func (a *WorkDocsActivities) UpdateDocument(input *workdocs.UpdateDocumentInput) (*workdocs.UpdateDocumentOutput, error) {
    return a.client.UpdateDocument(input)
}



func (a *WorkDocsActivities) UpdateDocumentVersion(input *workdocs.UpdateDocumentVersionInput) (*workdocs.UpdateDocumentVersionOutput, error) {
    return a.client.UpdateDocumentVersion(input)
}



func (a *WorkDocsActivities) UpdateFolder(input *workdocs.UpdateFolderInput) (*workdocs.UpdateFolderOutput, error) {
    return a.client.UpdateFolder(input)
}



func (a *WorkDocsActivities) UpdateUser(input *workdocs.UpdateUserInput) (*workdocs.UpdateUserOutput, error) {
    return a.client.UpdateUser(input)
}

