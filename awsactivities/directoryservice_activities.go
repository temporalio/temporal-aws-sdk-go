package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type DirectoryServiceActivities struct {
	client directoryserviceiface.DirectoryServiceAPI
}

func NewDirectoryServiceActivities(session *session.Session, config ...*aws.Config) *DirectoryServiceActivities {
	client := directoryservice.New(session, config...)
	return &DirectoryServiceActivities{client: client}
}

func (a *DirectoryServiceActivities) AcceptSharedDirectory(ctx context.Context, input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error) {
	return a.client.AcceptSharedDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) AddIpRoutes(ctx context.Context, input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error) {
	return a.client.AddIpRoutesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) AddTagsToResource(ctx context.Context, input *directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CancelSchemaExtension(ctx context.Context, input *directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error) {
	return a.client.CancelSchemaExtensionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ConnectDirectory(ctx context.Context, input *directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error) {
	return a.client.ConnectDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateAlias(ctx context.Context, input *directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error) {
	return a.client.CreateAliasWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateComputer(ctx context.Context, input *directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error) {
	return a.client.CreateComputerWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateConditionalForwarder(ctx context.Context, input *directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error) {
	return a.client.CreateConditionalForwarderWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateDirectory(ctx context.Context, input *directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error) {
	return a.client.CreateDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateLogSubscription(ctx context.Context, input *directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error) {
	return a.client.CreateLogSubscriptionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateMicrosoftAD(ctx context.Context, input *directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error) {
	return a.client.CreateMicrosoftADWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateSnapshot(ctx context.Context, input *directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error) {
	return a.client.CreateSnapshotWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) CreateTrust(ctx context.Context, input *directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error) {
	return a.client.CreateTrustWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteConditionalForwarder(ctx context.Context, input *directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error) {
	return a.client.DeleteConditionalForwarderWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteDirectory(ctx context.Context, input *directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error) {
	return a.client.DeleteDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteLogSubscription(ctx context.Context, input *directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error) {
	return a.client.DeleteLogSubscriptionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteSnapshot(ctx context.Context, input *directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error) {
	return a.client.DeleteSnapshotWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeleteTrust(ctx context.Context, input *directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error) {
	return a.client.DeleteTrustWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeregisterCertificate(ctx context.Context, input *directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error) {
	return a.client.DeregisterCertificateWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DeregisterEventTopic(ctx context.Context, input *directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error) {
	return a.client.DeregisterEventTopicWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeCertificate(ctx context.Context, input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error) {
	return a.client.DescribeCertificateWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeConditionalForwarders(ctx context.Context, input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error) {
	return a.client.DescribeConditionalForwardersWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeDirectories(ctx context.Context, input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error) {
	return a.client.DescribeDirectoriesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeDomainControllers(ctx context.Context, input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error) {
	return a.client.DescribeDomainControllersWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeEventTopics(ctx context.Context, input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error) {
	return a.client.DescribeEventTopicsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeLDAPSSettings(ctx context.Context, input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
	return a.client.DescribeLDAPSSettingsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeSharedDirectories(ctx context.Context, input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
	return a.client.DescribeSharedDirectoriesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeSnapshots(ctx context.Context, input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error) {
	return a.client.DescribeSnapshotsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DescribeTrusts(ctx context.Context, input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error) {
	return a.client.DescribeTrustsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DisableLDAPS(ctx context.Context, input *directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error) {
	return a.client.DisableLDAPSWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DisableRadius(ctx context.Context, input *directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error) {
	return a.client.DisableRadiusWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) DisableSso(ctx context.Context, input *directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error) {
	return a.client.DisableSsoWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) EnableLDAPS(ctx context.Context, input *directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error) {
	return a.client.EnableLDAPSWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) EnableRadius(ctx context.Context, input *directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error) {
	return a.client.EnableRadiusWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) EnableSso(ctx context.Context, input *directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error) {
	return a.client.EnableSsoWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) GetDirectoryLimits(ctx context.Context, input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error) {
	return a.client.GetDirectoryLimitsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) GetSnapshotLimits(ctx context.Context, input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error) {
	return a.client.GetSnapshotLimitsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListCertificates(ctx context.Context, input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error) {
	return a.client.ListCertificatesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListIpRoutes(ctx context.Context, input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error) {
	return a.client.ListIpRoutesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListLogSubscriptions(ctx context.Context, input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error) {
	return a.client.ListLogSubscriptionsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListSchemaExtensions(ctx context.Context, input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error) {
	return a.client.ListSchemaExtensionsWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ListTagsForResource(ctx context.Context, input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RegisterCertificate(ctx context.Context, input *directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error) {
	return a.client.RegisterCertificateWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RegisterEventTopic(ctx context.Context, input *directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error) {
	return a.client.RegisterEventTopicWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RejectSharedDirectory(ctx context.Context, input *directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error) {
	return a.client.RejectSharedDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RemoveIpRoutes(ctx context.Context, input *directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error) {
	return a.client.RemoveIpRoutesWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RemoveTagsFromResource(ctx context.Context, input *directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ResetUserPassword(ctx context.Context, input *directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error) {
	return a.client.ResetUserPasswordWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) RestoreFromSnapshot(ctx context.Context, input *directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error) {
	return a.client.RestoreFromSnapshotWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) ShareDirectory(ctx context.Context, input *directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error) {
	return a.client.ShareDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) StartSchemaExtension(ctx context.Context, input *directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error) {
	return a.client.StartSchemaExtensionWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UnshareDirectory(ctx context.Context, input *directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error) {
	return a.client.UnshareDirectoryWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateConditionalForwarder(ctx context.Context, input *directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error) {
	return a.client.UpdateConditionalForwarderWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateNumberOfDomainControllers(ctx context.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error) {
	return a.client.UpdateNumberOfDomainControllersWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateRadius(ctx context.Context, input *directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error) {
	return a.client.UpdateRadiusWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) UpdateTrust(ctx context.Context, input *directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error) {
	return a.client.UpdateTrustWithContext(ctx, input)
}

func (a *DirectoryServiceActivities) VerifyTrust(ctx context.Context, input *directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error) {
	return a.client.VerifyTrustWithContext(ctx, input)
}
