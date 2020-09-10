package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/workdocs/workdocsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type WorkDocsActivities struct {
	client workdocsiface.WorkDocsAPI
}

func NewWorkDocsActivities(session *session.Session, config ...*aws.Config) *WorkDocsActivities {
	client := workdocs.New(session, config...)
	return &WorkDocsActivities{client: client}
}

func (a *WorkDocsActivities) AbortDocumentVersionUpload(ctx context.Context, input *workdocs.AbortDocumentVersionUploadInput) (*workdocs.AbortDocumentVersionUploadOutput, error) {
	return a.client.AbortDocumentVersionUploadWithContext(ctx, input)
}

func (a *WorkDocsActivities) ActivateUser(ctx context.Context, input *workdocs.ActivateUserInput) (*workdocs.ActivateUserOutput, error) {
	return a.client.ActivateUserWithContext(ctx, input)
}

func (a *WorkDocsActivities) AddResourcePermissions(ctx context.Context, input *workdocs.AddResourcePermissionsInput) (*workdocs.AddResourcePermissionsOutput, error) {
	return a.client.AddResourcePermissionsWithContext(ctx, input)
}

func (a *WorkDocsActivities) CreateComment(ctx context.Context, input *workdocs.CreateCommentInput) (*workdocs.CreateCommentOutput, error) {
	return a.client.CreateCommentWithContext(ctx, input)
}

func (a *WorkDocsActivities) CreateCustomMetadata(ctx context.Context, input *workdocs.CreateCustomMetadataInput) (*workdocs.CreateCustomMetadataOutput, error) {
	return a.client.CreateCustomMetadataWithContext(ctx, input)
}

func (a *WorkDocsActivities) CreateFolder(ctx context.Context, input *workdocs.CreateFolderInput) (*workdocs.CreateFolderOutput, error) {
	return a.client.CreateFolderWithContext(ctx, input)
}

func (a *WorkDocsActivities) CreateLabels(ctx context.Context, input *workdocs.CreateLabelsInput) (*workdocs.CreateLabelsOutput, error) {
	return a.client.CreateLabelsWithContext(ctx, input)
}

func (a *WorkDocsActivities) CreateNotificationSubscription(ctx context.Context, input *workdocs.CreateNotificationSubscriptionInput) (*workdocs.CreateNotificationSubscriptionOutput, error) {
	return a.client.CreateNotificationSubscriptionWithContext(ctx, input)
}

func (a *WorkDocsActivities) CreateUser(ctx context.Context, input *workdocs.CreateUserInput) (*workdocs.CreateUserOutput, error) {
	return a.client.CreateUserWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeactivateUser(ctx context.Context, input *workdocs.DeactivateUserInput) (*workdocs.DeactivateUserOutput, error) {
	return a.client.DeactivateUserWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteComment(ctx context.Context, input *workdocs.DeleteCommentInput) (*workdocs.DeleteCommentOutput, error) {
	return a.client.DeleteCommentWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteCustomMetadata(ctx context.Context, input *workdocs.DeleteCustomMetadataInput) (*workdocs.DeleteCustomMetadataOutput, error) {
	return a.client.DeleteCustomMetadataWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteDocument(ctx context.Context, input *workdocs.DeleteDocumentInput) (*workdocs.DeleteDocumentOutput, error) {
	return a.client.DeleteDocumentWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteFolder(ctx context.Context, input *workdocs.DeleteFolderInput) (*workdocs.DeleteFolderOutput, error) {
	return a.client.DeleteFolderWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteFolderContents(ctx context.Context, input *workdocs.DeleteFolderContentsInput) (*workdocs.DeleteFolderContentsOutput, error) {
	return a.client.DeleteFolderContentsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteLabels(ctx context.Context, input *workdocs.DeleteLabelsInput) (*workdocs.DeleteLabelsOutput, error) {
	return a.client.DeleteLabelsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteNotificationSubscription(ctx context.Context, input *workdocs.DeleteNotificationSubscriptionInput) (*workdocs.DeleteNotificationSubscriptionOutput, error) {
	return a.client.DeleteNotificationSubscriptionWithContext(ctx, input)
}

func (a *WorkDocsActivities) DeleteUser(ctx context.Context, input *workdocs.DeleteUserInput) (*workdocs.DeleteUserOutput, error) {
	return a.client.DeleteUserWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeActivities(ctx context.Context, input *workdocs.DescribeActivitiesInput) (*workdocs.DescribeActivitiesOutput, error) {
	return a.client.DescribeActivitiesWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeComments(ctx context.Context, input *workdocs.DescribeCommentsInput) (*workdocs.DescribeCommentsOutput, error) {
	return a.client.DescribeCommentsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeDocumentVersions(ctx context.Context, input *workdocs.DescribeDocumentVersionsInput) (*workdocs.DescribeDocumentVersionsOutput, error) {
	return a.client.DescribeDocumentVersionsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeFolderContents(ctx context.Context, input *workdocs.DescribeFolderContentsInput) (*workdocs.DescribeFolderContentsOutput, error) {
	return a.client.DescribeFolderContentsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeGroups(ctx context.Context, input *workdocs.DescribeGroupsInput) (*workdocs.DescribeGroupsOutput, error) {
	return a.client.DescribeGroupsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeNotificationSubscriptions(ctx context.Context, input *workdocs.DescribeNotificationSubscriptionsInput) (*workdocs.DescribeNotificationSubscriptionsOutput, error) {
	return a.client.DescribeNotificationSubscriptionsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeResourcePermissions(ctx context.Context, input *workdocs.DescribeResourcePermissionsInput) (*workdocs.DescribeResourcePermissionsOutput, error) {
	return a.client.DescribeResourcePermissionsWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeRootFolders(ctx context.Context, input *workdocs.DescribeRootFoldersInput) (*workdocs.DescribeRootFoldersOutput, error) {
	return a.client.DescribeRootFoldersWithContext(ctx, input)
}

func (a *WorkDocsActivities) DescribeUsers(ctx context.Context, input *workdocs.DescribeUsersInput) (*workdocs.DescribeUsersOutput, error) {
	return a.client.DescribeUsersWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetCurrentUser(ctx context.Context, input *workdocs.GetCurrentUserInput) (*workdocs.GetCurrentUserOutput, error) {
	return a.client.GetCurrentUserWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetDocument(ctx context.Context, input *workdocs.GetDocumentInput) (*workdocs.GetDocumentOutput, error) {
	return a.client.GetDocumentWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetDocumentPath(ctx context.Context, input *workdocs.GetDocumentPathInput) (*workdocs.GetDocumentPathOutput, error) {
	return a.client.GetDocumentPathWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetDocumentVersion(ctx context.Context, input *workdocs.GetDocumentVersionInput) (*workdocs.GetDocumentVersionOutput, error) {
	return a.client.GetDocumentVersionWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetFolder(ctx context.Context, input *workdocs.GetFolderInput) (*workdocs.GetFolderOutput, error) {
	return a.client.GetFolderWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetFolderPath(ctx context.Context, input *workdocs.GetFolderPathInput) (*workdocs.GetFolderPathOutput, error) {
	return a.client.GetFolderPathWithContext(ctx, input)
}

func (a *WorkDocsActivities) GetResources(ctx context.Context, input *workdocs.GetResourcesInput) (*workdocs.GetResourcesOutput, error) {
	return a.client.GetResourcesWithContext(ctx, input)
}

func (a *WorkDocsActivities) InitiateDocumentVersionUpload(ctx context.Context, input *workdocs.InitiateDocumentVersionUploadInput) (*workdocs.InitiateDocumentVersionUploadOutput, error) {
	return a.client.InitiateDocumentVersionUploadWithContext(ctx, input)
}

func (a *WorkDocsActivities) RemoveAllResourcePermissions(ctx context.Context, input *workdocs.RemoveAllResourcePermissionsInput) (*workdocs.RemoveAllResourcePermissionsOutput, error) {
	return a.client.RemoveAllResourcePermissionsWithContext(ctx, input)
}

func (a *WorkDocsActivities) RemoveResourcePermission(ctx context.Context, input *workdocs.RemoveResourcePermissionInput) (*workdocs.RemoveResourcePermissionOutput, error) {
	return a.client.RemoveResourcePermissionWithContext(ctx, input)
}

func (a *WorkDocsActivities) UpdateDocument(ctx context.Context, input *workdocs.UpdateDocumentInput) (*workdocs.UpdateDocumentOutput, error) {
	return a.client.UpdateDocumentWithContext(ctx, input)
}

func (a *WorkDocsActivities) UpdateDocumentVersion(ctx context.Context, input *workdocs.UpdateDocumentVersionInput) (*workdocs.UpdateDocumentVersionOutput, error) {
	return a.client.UpdateDocumentVersionWithContext(ctx, input)
}

func (a *WorkDocsActivities) UpdateFolder(ctx context.Context, input *workdocs.UpdateFolderInput) (*workdocs.UpdateFolderOutput, error) {
	return a.client.UpdateFolderWithContext(ctx, input)
}

func (a *WorkDocsActivities) UpdateUser(ctx context.Context, input *workdocs.UpdateUserInput) (*workdocs.UpdateUserOutput, error) {
	return a.client.UpdateUserWithContext(ctx, input)
}
