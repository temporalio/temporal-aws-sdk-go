
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appstream/appstreamiface"
)

type AppStreamActivities struct {
	client appstreamiface.AppStreamAPI
}

func NewAppStreamActivities(client appstreamiface.AppStreamAPI) *AppStreamActivities {
    return &AppStreamActivities{client: client}
}

func (a *AppStreamActivities) AssociateFleet(input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
    return a.client.AssociateFleet(input)
}

func (a *AppStreamActivities) BatchAssociateUserStack(input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error) {
    return a.client.BatchAssociateUserStack(input)
}

func (a *AppStreamActivities) BatchDisassociateUserStack(input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error) {
    return a.client.BatchDisassociateUserStack(input)
}

func (a *AppStreamActivities) CopyImage(input *appstream.CopyImageInput) (*appstream.CopyImageOutput, error) {
    return a.client.CopyImage(input)
}

func (a *AppStreamActivities) CreateDirectoryConfig(input *appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error) {
    return a.client.CreateDirectoryConfig(input)
}

func (a *AppStreamActivities) CreateFleet(input *appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error) {
    return a.client.CreateFleet(input)
}

func (a *AppStreamActivities) CreateImageBuilder(input *appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error) {
    return a.client.CreateImageBuilder(input)
}

func (a *AppStreamActivities) CreateImageBuilderStreamingURL(input *appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error) {
    return a.client.CreateImageBuilderStreamingURL(input)
}

func (a *AppStreamActivities) CreateStack(input *appstream.CreateStackInput) (*appstream.CreateStackOutput, error) {
    return a.client.CreateStack(input)
}

func (a *AppStreamActivities) CreateStreamingURL(input *appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error) {
    return a.client.CreateStreamingURL(input)
}

func (a *AppStreamActivities) CreateUsageReportSubscription(input *appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error) {
    return a.client.CreateUsageReportSubscription(input)
}

func (a *AppStreamActivities) CreateUser(input *appstream.CreateUserInput) (*appstream.CreateUserOutput, error) {
    return a.client.CreateUser(input)
}

func (a *AppStreamActivities) DeleteDirectoryConfig(input *appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error) {
    return a.client.DeleteDirectoryConfig(input)
}

func (a *AppStreamActivities) DeleteFleet(input *appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error) {
    return a.client.DeleteFleet(input)
}

func (a *AppStreamActivities) DeleteImage(input *appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error) {
    return a.client.DeleteImage(input)
}

func (a *AppStreamActivities) DeleteImageBuilder(input *appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error) {
    return a.client.DeleteImageBuilder(input)
}

func (a *AppStreamActivities) DeleteImagePermissions(input *appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error) {
    return a.client.DeleteImagePermissions(input)
}

func (a *AppStreamActivities) DeleteStack(input *appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error) {
    return a.client.DeleteStack(input)
}

func (a *AppStreamActivities) DeleteUsageReportSubscription(input *appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error) {
    return a.client.DeleteUsageReportSubscription(input)
}

func (a *AppStreamActivities) DeleteUser(input *appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *AppStreamActivities) DescribeDirectoryConfigs(input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error) {
    return a.client.DescribeDirectoryConfigs(input)
}

func (a *AppStreamActivities) DescribeFleets(input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error) {
    return a.client.DescribeFleets(input)
}

func (a *AppStreamActivities) DescribeImageBuilders(input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error) {
    return a.client.DescribeImageBuilders(input)
}

func (a *AppStreamActivities) DescribeImagePermissions(input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error) {
    return a.client.DescribeImagePermissions(input)
}

func (a *AppStreamActivities) DescribeImages(input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error) {
    return a.client.DescribeImages(input)
}

func (a *AppStreamActivities) DescribeSessions(input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error) {
    return a.client.DescribeSessions(input)
}

func (a *AppStreamActivities) DescribeStacks(input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error) {
    return a.client.DescribeStacks(input)
}

func (a *AppStreamActivities) DescribeUsageReportSubscriptions(input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
    return a.client.DescribeUsageReportSubscriptions(input)
}

func (a *AppStreamActivities) DescribeUserStackAssociations(input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error) {
    return a.client.DescribeUserStackAssociations(input)
}

func (a *AppStreamActivities) DescribeUsers(input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error) {
    return a.client.DescribeUsers(input)
}

func (a *AppStreamActivities) DisableUser(input *appstream.DisableUserInput) (*appstream.DisableUserOutput, error) {
    return a.client.DisableUser(input)
}

func (a *AppStreamActivities) DisassociateFleet(input *appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error) {
    return a.client.DisassociateFleet(input)
}

func (a *AppStreamActivities) EnableUser(input *appstream.EnableUserInput) (*appstream.EnableUserOutput, error) {
    return a.client.EnableUser(input)
}

func (a *AppStreamActivities) ExpireSession(input *appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error) {
    return a.client.ExpireSession(input)
}

func (a *AppStreamActivities) ListAssociatedFleets(input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error) {
    return a.client.ListAssociatedFleets(input)
}

func (a *AppStreamActivities) ListAssociatedStacks(input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error) {
    return a.client.ListAssociatedStacks(input)
}

func (a *AppStreamActivities) ListTagsForResource(input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *AppStreamActivities) StartFleet(input *appstream.StartFleetInput) (*appstream.StartFleetOutput, error) {
    return a.client.StartFleet(input)
}

func (a *AppStreamActivities) StartImageBuilder(input *appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error) {
    return a.client.StartImageBuilder(input)
}

func (a *AppStreamActivities) StopFleet(input *appstream.StopFleetInput) (*appstream.StopFleetOutput, error) {
    return a.client.StopFleet(input)
}

func (a *AppStreamActivities) StopImageBuilder(input *appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error) {
    return a.client.StopImageBuilder(input)
}

func (a *AppStreamActivities) TagResource(input *appstream.TagResourceInput) (*appstream.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AppStreamActivities) UntagResource(input *appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AppStreamActivities) UpdateDirectoryConfig(input *appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error) {
    return a.client.UpdateDirectoryConfig(input)
}

func (a *AppStreamActivities) UpdateFleet(input *appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error) {
    return a.client.UpdateFleet(input)
}

func (a *AppStreamActivities) UpdateImagePermissions(input *appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error) {
    return a.client.UpdateImagePermissions(input)
}

func (a *AppStreamActivities) UpdateStack(input *appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error) {
    return a.client.UpdateStack(input)
}

func (a *AppStreamActivities) WaitUntilFleetStarted(input *appstream.DescribeFleetsInput) error {
    return a.client.WaitUntilFleetStarted(input)
}

func (a *AppStreamActivities) WaitUntilFleetStopped(input *appstream.DescribeFleetsInput) error {
    return a.client.WaitUntilFleetStopped(input)
}
