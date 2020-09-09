package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface"
)

type DirectoryServiceActivities struct {
	client directoryserviceiface.DirectoryServiceAPI
}

func NewDirectoryServiceActivities(session *session.Session, config ...*aws.Config) *DirectoryServiceActivities {
	client := directoryservice.New(session, config...)
	return &DirectoryServiceActivities{client: client}
}

func (a *DirectoryServiceActivities) AcceptSharedDirectory(input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error) {
	return a.client.AcceptSharedDirectory(input)
}

func (a *DirectoryServiceActivities) AddIpRoutes(input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error) {
	return a.client.AddIpRoutes(input)
}

func (a *DirectoryServiceActivities) AddTagsToResource(input *directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResource(input)
}

func (a *DirectoryServiceActivities) CancelSchemaExtension(input *directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error) {
	return a.client.CancelSchemaExtension(input)
}

func (a *DirectoryServiceActivities) ConnectDirectory(input *directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error) {
	return a.client.ConnectDirectory(input)
}

func (a *DirectoryServiceActivities) CreateAlias(input *directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error) {
	return a.client.CreateAlias(input)
}

func (a *DirectoryServiceActivities) CreateComputer(input *directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error) {
	return a.client.CreateComputer(input)
}

func (a *DirectoryServiceActivities) CreateConditionalForwarder(input *directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error) {
	return a.client.CreateConditionalForwarder(input)
}

func (a *DirectoryServiceActivities) CreateDirectory(input *directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error) {
	return a.client.CreateDirectory(input)
}

func (a *DirectoryServiceActivities) CreateLogSubscription(input *directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error) {
	return a.client.CreateLogSubscription(input)
}

func (a *DirectoryServiceActivities) CreateMicrosoftAD(input *directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error) {
	return a.client.CreateMicrosoftAD(input)
}

func (a *DirectoryServiceActivities) CreateSnapshot(input *directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error) {
	return a.client.CreateSnapshot(input)
}

func (a *DirectoryServiceActivities) CreateTrust(input *directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error) {
	return a.client.CreateTrust(input)
}

func (a *DirectoryServiceActivities) DeleteConditionalForwarder(input *directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error) {
	return a.client.DeleteConditionalForwarder(input)
}

func (a *DirectoryServiceActivities) DeleteDirectory(input *directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error) {
	return a.client.DeleteDirectory(input)
}

func (a *DirectoryServiceActivities) DeleteLogSubscription(input *directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error) {
	return a.client.DeleteLogSubscription(input)
}

func (a *DirectoryServiceActivities) DeleteSnapshot(input *directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error) {
	return a.client.DeleteSnapshot(input)
}

func (a *DirectoryServiceActivities) DeleteTrust(input *directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error) {
	return a.client.DeleteTrust(input)
}

func (a *DirectoryServiceActivities) DeregisterCertificate(input *directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error) {
	return a.client.DeregisterCertificate(input)
}

func (a *DirectoryServiceActivities) DeregisterEventTopic(input *directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error) {
	return a.client.DeregisterEventTopic(input)
}

func (a *DirectoryServiceActivities) DescribeCertificate(input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error) {
	return a.client.DescribeCertificate(input)
}

func (a *DirectoryServiceActivities) DescribeConditionalForwarders(input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error) {
	return a.client.DescribeConditionalForwarders(input)
}

func (a *DirectoryServiceActivities) DescribeDirectories(input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error) {
	return a.client.DescribeDirectories(input)
}

func (a *DirectoryServiceActivities) DescribeDomainControllers(input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error) {
	return a.client.DescribeDomainControllers(input)
}

func (a *DirectoryServiceActivities) DescribeEventTopics(input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error) {
	return a.client.DescribeEventTopics(input)
}

func (a *DirectoryServiceActivities) DescribeLDAPSSettings(input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
	return a.client.DescribeLDAPSSettings(input)
}

func (a *DirectoryServiceActivities) DescribeSharedDirectories(input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
	return a.client.DescribeSharedDirectories(input)
}

func (a *DirectoryServiceActivities) DescribeSnapshots(input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error) {
	return a.client.DescribeSnapshots(input)
}

func (a *DirectoryServiceActivities) DescribeTrusts(input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error) {
	return a.client.DescribeTrusts(input)
}

func (a *DirectoryServiceActivities) DisableLDAPS(input *directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error) {
	return a.client.DisableLDAPS(input)
}

func (a *DirectoryServiceActivities) DisableRadius(input *directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error) {
	return a.client.DisableRadius(input)
}

func (a *DirectoryServiceActivities) DisableSso(input *directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error) {
	return a.client.DisableSso(input)
}

func (a *DirectoryServiceActivities) EnableLDAPS(input *directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error) {
	return a.client.EnableLDAPS(input)
}

func (a *DirectoryServiceActivities) EnableRadius(input *directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error) {
	return a.client.EnableRadius(input)
}

func (a *DirectoryServiceActivities) EnableSso(input *directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error) {
	return a.client.EnableSso(input)
}

func (a *DirectoryServiceActivities) GetDirectoryLimits(input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error) {
	return a.client.GetDirectoryLimits(input)
}

func (a *DirectoryServiceActivities) GetSnapshotLimits(input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error) {
	return a.client.GetSnapshotLimits(input)
}

func (a *DirectoryServiceActivities) ListCertificates(input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error) {
	return a.client.ListCertificates(input)
}

func (a *DirectoryServiceActivities) ListIpRoutes(input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error) {
	return a.client.ListIpRoutes(input)
}

func (a *DirectoryServiceActivities) ListLogSubscriptions(input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error) {
	return a.client.ListLogSubscriptions(input)
}

func (a *DirectoryServiceActivities) ListSchemaExtensions(input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error) {
	return a.client.ListSchemaExtensions(input)
}

func (a *DirectoryServiceActivities) ListTagsForResource(input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *DirectoryServiceActivities) RegisterCertificate(input *directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error) {
	return a.client.RegisterCertificate(input)
}

func (a *DirectoryServiceActivities) RegisterEventTopic(input *directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error) {
	return a.client.RegisterEventTopic(input)
}

func (a *DirectoryServiceActivities) RejectSharedDirectory(input *directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error) {
	return a.client.RejectSharedDirectory(input)
}

func (a *DirectoryServiceActivities) RemoveIpRoutes(input *directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error) {
	return a.client.RemoveIpRoutes(input)
}

func (a *DirectoryServiceActivities) RemoveTagsFromResource(input *directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResource(input)
}

func (a *DirectoryServiceActivities) ResetUserPassword(input *directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error) {
	return a.client.ResetUserPassword(input)
}

func (a *DirectoryServiceActivities) RestoreFromSnapshot(input *directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error) {
	return a.client.RestoreFromSnapshot(input)
}

func (a *DirectoryServiceActivities) ShareDirectory(input *directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error) {
	return a.client.ShareDirectory(input)
}

func (a *DirectoryServiceActivities) StartSchemaExtension(input *directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error) {
	return a.client.StartSchemaExtension(input)
}

func (a *DirectoryServiceActivities) UnshareDirectory(input *directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error) {
	return a.client.UnshareDirectory(input)
}

func (a *DirectoryServiceActivities) UpdateConditionalForwarder(input *directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error) {
	return a.client.UpdateConditionalForwarder(input)
}

func (a *DirectoryServiceActivities) UpdateNumberOfDomainControllers(input *directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error) {
	return a.client.UpdateNumberOfDomainControllers(input)
}

func (a *DirectoryServiceActivities) UpdateRadius(input *directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error) {
	return a.client.UpdateRadius(input)
}

func (a *DirectoryServiceActivities) UpdateTrust(input *directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error) {
	return a.client.UpdateTrust(input)
}

func (a *DirectoryServiceActivities) VerifyTrust(input *directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error) {
	return a.client.VerifyTrust(input)
}
