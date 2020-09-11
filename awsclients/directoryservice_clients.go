package awsclients

import (
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"go.temporal.io/sdk/workflow"
)

type DirectoryServiceClient interface {
       AcceptSharedDirectory(ctx workflow.Context, input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error)
       AcceptSharedDirectoryAsync(ctx workflow.Context, input *directoryservice.AcceptSharedDirectoryInput) *DirectoryserviceAcceptSharedDirectoryResult

       AddIpRoutes(ctx workflow.Context, input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error)
       AddIpRoutesAsync(ctx workflow.Context, input *directoryservice.AddIpRoutesInput) *DirectoryserviceAddIpRoutesResult

       AddTagsToResource(ctx workflow.Context, input *directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error)
       AddTagsToResourceAsync(ctx workflow.Context, input *directoryservice.AddTagsToResourceInput) *DirectoryserviceAddTagsToResourceResult

       CancelSchemaExtension(ctx workflow.Context, input *directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error)
       CancelSchemaExtensionAsync(ctx workflow.Context, input *directoryservice.CancelSchemaExtensionInput) *DirectoryserviceCancelSchemaExtensionResult

       ConnectDirectory(ctx workflow.Context, input *directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)
       ConnectDirectoryAsync(ctx workflow.Context, input *directoryservice.ConnectDirectoryInput) *DirectoryserviceConnectDirectoryResult

       CreateAlias(ctx workflow.Context, input *directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)
       CreateAliasAsync(ctx workflow.Context, input *directoryservice.CreateAliasInput) *DirectoryserviceCreateAliasResult

       CreateComputer(ctx workflow.Context, input *directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)
       CreateComputerAsync(ctx workflow.Context, input *directoryservice.CreateComputerInput) *DirectoryserviceCreateComputerResult

       CreateConditionalForwarder(ctx workflow.Context, input *directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error)
       CreateConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.CreateConditionalForwarderInput) *DirectoryserviceCreateConditionalForwarderResult

       CreateDirectory(ctx workflow.Context, input *directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)
       CreateDirectoryAsync(ctx workflow.Context, input *directoryservice.CreateDirectoryInput) *DirectoryserviceCreateDirectoryResult

       CreateLogSubscription(ctx workflow.Context, input *directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error)
       CreateLogSubscriptionAsync(ctx workflow.Context, input *directoryservice.CreateLogSubscriptionInput) *DirectoryserviceCreateLogSubscriptionResult

       CreateMicrosoftAD(ctx workflow.Context, input *directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error)
       CreateMicrosoftADAsync(ctx workflow.Context, input *directoryservice.CreateMicrosoftADInput) *DirectoryserviceCreateMicrosoftADResult

       CreateSnapshot(ctx workflow.Context, input *directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)
       CreateSnapshotAsync(ctx workflow.Context, input *directoryservice.CreateSnapshotInput) *DirectoryserviceCreateSnapshotResult

       CreateTrust(ctx workflow.Context, input *directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error)
       CreateTrustAsync(ctx workflow.Context, input *directoryservice.CreateTrustInput) *DirectoryserviceCreateTrustResult

       DeleteConditionalForwarder(ctx workflow.Context, input *directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error)
       DeleteConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.DeleteConditionalForwarderInput) *DirectoryserviceDeleteConditionalForwarderResult

       DeleteDirectory(ctx workflow.Context, input *directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)
       DeleteDirectoryAsync(ctx workflow.Context, input *directoryservice.DeleteDirectoryInput) *DirectoryserviceDeleteDirectoryResult

       DeleteLogSubscription(ctx workflow.Context, input *directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error)
       DeleteLogSubscriptionAsync(ctx workflow.Context, input *directoryservice.DeleteLogSubscriptionInput) *DirectoryserviceDeleteLogSubscriptionResult

       DeleteSnapshot(ctx workflow.Context, input *directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)
       DeleteSnapshotAsync(ctx workflow.Context, input *directoryservice.DeleteSnapshotInput) *DirectoryserviceDeleteSnapshotResult

       DeleteTrust(ctx workflow.Context, input *directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error)
       DeleteTrustAsync(ctx workflow.Context, input *directoryservice.DeleteTrustInput) *DirectoryserviceDeleteTrustResult

       DeregisterCertificate(ctx workflow.Context, input *directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error)
       DeregisterCertificateAsync(ctx workflow.Context, input *directoryservice.DeregisterCertificateInput) *DirectoryserviceDeregisterCertificateResult

       DeregisterEventTopic(ctx workflow.Context, input *directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error)
       DeregisterEventTopicAsync(ctx workflow.Context, input *directoryservice.DeregisterEventTopicInput) *DirectoryserviceDeregisterEventTopicResult

       DescribeCertificate(ctx workflow.Context, input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error)
       DescribeCertificateAsync(ctx workflow.Context, input *directoryservice.DescribeCertificateInput) *DirectoryserviceDescribeCertificateResult

       DescribeConditionalForwarders(ctx workflow.Context, input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error)
       DescribeConditionalForwardersAsync(ctx workflow.Context, input *directoryservice.DescribeConditionalForwardersInput) *DirectoryserviceDescribeConditionalForwardersResult

       DescribeDirectories(ctx workflow.Context, input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)
       DescribeDirectoriesAsync(ctx workflow.Context, input *directoryservice.DescribeDirectoriesInput) *DirectoryserviceDescribeDirectoriesResult

       DescribeDomainControllers(ctx workflow.Context, input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error)
       DescribeDomainControllersAsync(ctx workflow.Context, input *directoryservice.DescribeDomainControllersInput) *DirectoryserviceDescribeDomainControllersResult

       DescribeEventTopics(ctx workflow.Context, input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error)
       DescribeEventTopicsAsync(ctx workflow.Context, input *directoryservice.DescribeEventTopicsInput) *DirectoryserviceDescribeEventTopicsResult

       DescribeLDAPSSettings(ctx workflow.Context, input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error)
       DescribeLDAPSSettingsAsync(ctx workflow.Context, input *directoryservice.DescribeLDAPSSettingsInput) *DirectoryserviceDescribeLDAPSSettingsResult

       DescribeSharedDirectories(ctx workflow.Context, input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error)
       DescribeSharedDirectoriesAsync(ctx workflow.Context, input *directoryservice.DescribeSharedDirectoriesInput) *DirectoryserviceDescribeSharedDirectoriesResult

       DescribeSnapshots(ctx workflow.Context, input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)
       DescribeSnapshotsAsync(ctx workflow.Context, input *directoryservice.DescribeSnapshotsInput) *DirectoryserviceDescribeSnapshotsResult

       DescribeTrusts(ctx workflow.Context, input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error)
       DescribeTrustsAsync(ctx workflow.Context, input *directoryservice.DescribeTrustsInput) *DirectoryserviceDescribeTrustsResult

       DisableLDAPS(ctx workflow.Context, input *directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error)
       DisableLDAPSAsync(ctx workflow.Context, input *directoryservice.DisableLDAPSInput) *DirectoryserviceDisableLDAPSResult

       DisableRadius(ctx workflow.Context, input *directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)
       DisableRadiusAsync(ctx workflow.Context, input *directoryservice.DisableRadiusInput) *DirectoryserviceDisableRadiusResult

       DisableSso(ctx workflow.Context, input *directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error)
       DisableSsoAsync(ctx workflow.Context, input *directoryservice.DisableSsoInput) *DirectoryserviceDisableSsoResult

       EnableLDAPS(ctx workflow.Context, input *directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error)
       EnableLDAPSAsync(ctx workflow.Context, input *directoryservice.EnableLDAPSInput) *DirectoryserviceEnableLDAPSResult

       EnableRadius(ctx workflow.Context, input *directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)
       EnableRadiusAsync(ctx workflow.Context, input *directoryservice.EnableRadiusInput) *DirectoryserviceEnableRadiusResult

       EnableSso(ctx workflow.Context, input *directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error)
       EnableSsoAsync(ctx workflow.Context, input *directoryservice.EnableSsoInput) *DirectoryserviceEnableSsoResult

       GetDirectoryLimits(ctx workflow.Context, input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)
       GetDirectoryLimitsAsync(ctx workflow.Context, input *directoryservice.GetDirectoryLimitsInput) *DirectoryserviceGetDirectoryLimitsResult

       GetSnapshotLimits(ctx workflow.Context, input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)
       GetSnapshotLimitsAsync(ctx workflow.Context, input *directoryservice.GetSnapshotLimitsInput) *DirectoryserviceGetSnapshotLimitsResult

       ListCertificates(ctx workflow.Context, input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error)
       ListCertificatesAsync(ctx workflow.Context, input *directoryservice.ListCertificatesInput) *DirectoryserviceListCertificatesResult

       ListIpRoutes(ctx workflow.Context, input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error)
       ListIpRoutesAsync(ctx workflow.Context, input *directoryservice.ListIpRoutesInput) *DirectoryserviceListIpRoutesResult

       ListLogSubscriptions(ctx workflow.Context, input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error)
       ListLogSubscriptionsAsync(ctx workflow.Context, input *directoryservice.ListLogSubscriptionsInput) *DirectoryserviceListLogSubscriptionsResult

       ListSchemaExtensions(ctx workflow.Context, input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error)
       ListSchemaExtensionsAsync(ctx workflow.Context, input *directoryservice.ListSchemaExtensionsInput) *DirectoryserviceListSchemaExtensionsResult

       ListTagsForResource(ctx workflow.Context, input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *directoryservice.ListTagsForResourceInput) *DirectoryserviceListTagsForResourceResult

       RegisterCertificate(ctx workflow.Context, input *directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error)
       RegisterCertificateAsync(ctx workflow.Context, input *directoryservice.RegisterCertificateInput) *DirectoryserviceRegisterCertificateResult

       RegisterEventTopic(ctx workflow.Context, input *directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error)
       RegisterEventTopicAsync(ctx workflow.Context, input *directoryservice.RegisterEventTopicInput) *DirectoryserviceRegisterEventTopicResult

       RejectSharedDirectory(ctx workflow.Context, input *directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error)
       RejectSharedDirectoryAsync(ctx workflow.Context, input *directoryservice.RejectSharedDirectoryInput) *DirectoryserviceRejectSharedDirectoryResult

       RemoveIpRoutes(ctx workflow.Context, input *directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error)
       RemoveIpRoutesAsync(ctx workflow.Context, input *directoryservice.RemoveIpRoutesInput) *DirectoryserviceRemoveIpRoutesResult

       RemoveTagsFromResource(ctx workflow.Context, input *directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error)
       RemoveTagsFromResourceAsync(ctx workflow.Context, input *directoryservice.RemoveTagsFromResourceInput) *DirectoryserviceRemoveTagsFromResourceResult

       ResetUserPassword(ctx workflow.Context, input *directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error)
       ResetUserPasswordAsync(ctx workflow.Context, input *directoryservice.ResetUserPasswordInput) *DirectoryserviceResetUserPasswordResult

       RestoreFromSnapshot(ctx workflow.Context, input *directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)
       RestoreFromSnapshotAsync(ctx workflow.Context, input *directoryservice.RestoreFromSnapshotInput) *DirectoryserviceRestoreFromSnapshotResult

       ShareDirectory(ctx workflow.Context, input *directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error)
       ShareDirectoryAsync(ctx workflow.Context, input *directoryservice.ShareDirectoryInput) *DirectoryserviceShareDirectoryResult

       StartSchemaExtension(ctx workflow.Context, input *directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error)
       StartSchemaExtensionAsync(ctx workflow.Context, input *directoryservice.StartSchemaExtensionInput) *DirectoryserviceStartSchemaExtensionResult

       UnshareDirectory(ctx workflow.Context, input *directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error)
       UnshareDirectoryAsync(ctx workflow.Context, input *directoryservice.UnshareDirectoryInput) *DirectoryserviceUnshareDirectoryResult

       UpdateConditionalForwarder(ctx workflow.Context, input *directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error)
       UpdateConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.UpdateConditionalForwarderInput) *DirectoryserviceUpdateConditionalForwarderResult

       UpdateNumberOfDomainControllers(ctx workflow.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error)
       UpdateNumberOfDomainControllersAsync(ctx workflow.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) *DirectoryserviceUpdateNumberOfDomainControllersResult

       UpdateRadius(ctx workflow.Context, input *directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)
       UpdateRadiusAsync(ctx workflow.Context, input *directoryservice.UpdateRadiusInput) *DirectoryserviceUpdateRadiusResult

       UpdateTrust(ctx workflow.Context, input *directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error)
       UpdateTrustAsync(ctx workflow.Context, input *directoryservice.UpdateTrustInput) *DirectoryserviceUpdateTrustResult

       VerifyTrust(ctx workflow.Context, input *directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error)
       VerifyTrustAsync(ctx workflow.Context, input *directoryservice.VerifyTrustInput) *DirectoryserviceVerifyTrustResult
}

type DirectoryServiceStub struct {
}

func NewDirectoryServiceStub() DirectoryServiceClient {
    return &DirectoryServiceStub{}
}

type DirectoryserviceAcceptSharedDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceAcceptSharedDirectoryResult) Get(ctx workflow.Context) (*directoryservice.AcceptSharedDirectoryOutput, error) {
    var output directoryservice.AcceptSharedDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceAddIpRoutesResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceAddIpRoutesResult) Get(ctx workflow.Context) (*directoryservice.AddIpRoutesOutput, error) {
    var output directoryservice.AddIpRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceAddTagsToResourceResult) Get(ctx workflow.Context) (*directoryservice.AddTagsToResourceOutput, error) {
    var output directoryservice.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCancelSchemaExtensionResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCancelSchemaExtensionResult) Get(ctx workflow.Context) (*directoryservice.CancelSchemaExtensionOutput, error) {
    var output directoryservice.CancelSchemaExtensionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceConnectDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceConnectDirectoryResult) Get(ctx workflow.Context) (*directoryservice.ConnectDirectoryOutput, error) {
    var output directoryservice.ConnectDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateAliasResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateAliasResult) Get(ctx workflow.Context) (*directoryservice.CreateAliasOutput, error) {
    var output directoryservice.CreateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateComputerResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateComputerResult) Get(ctx workflow.Context) (*directoryservice.CreateComputerOutput, error) {
    var output directoryservice.CreateComputerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateConditionalForwarderResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateConditionalForwarderResult) Get(ctx workflow.Context) (*directoryservice.CreateConditionalForwarderOutput, error) {
    var output directoryservice.CreateConditionalForwarderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateDirectoryResult) Get(ctx workflow.Context) (*directoryservice.CreateDirectoryOutput, error) {
    var output directoryservice.CreateDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateLogSubscriptionResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateLogSubscriptionResult) Get(ctx workflow.Context) (*directoryservice.CreateLogSubscriptionOutput, error) {
    var output directoryservice.CreateLogSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateMicrosoftADResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateMicrosoftADResult) Get(ctx workflow.Context) (*directoryservice.CreateMicrosoftADOutput, error) {
    var output directoryservice.CreateMicrosoftADOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateSnapshotResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateSnapshotResult) Get(ctx workflow.Context) (*directoryservice.CreateSnapshotOutput, error) {
    var output directoryservice.CreateSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceCreateTrustResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceCreateTrustResult) Get(ctx workflow.Context) (*directoryservice.CreateTrustOutput, error) {
    var output directoryservice.CreateTrustOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeleteConditionalForwarderResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeleteConditionalForwarderResult) Get(ctx workflow.Context) (*directoryservice.DeleteConditionalForwarderOutput, error) {
    var output directoryservice.DeleteConditionalForwarderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeleteDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeleteDirectoryResult) Get(ctx workflow.Context) (*directoryservice.DeleteDirectoryOutput, error) {
    var output directoryservice.DeleteDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeleteLogSubscriptionResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeleteLogSubscriptionResult) Get(ctx workflow.Context) (*directoryservice.DeleteLogSubscriptionOutput, error) {
    var output directoryservice.DeleteLogSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeleteSnapshotResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeleteSnapshotResult) Get(ctx workflow.Context) (*directoryservice.DeleteSnapshotOutput, error) {
    var output directoryservice.DeleteSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeleteTrustResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeleteTrustResult) Get(ctx workflow.Context) (*directoryservice.DeleteTrustOutput, error) {
    var output directoryservice.DeleteTrustOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeregisterCertificateResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeregisterCertificateResult) Get(ctx workflow.Context) (*directoryservice.DeregisterCertificateOutput, error) {
    var output directoryservice.DeregisterCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDeregisterEventTopicResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDeregisterEventTopicResult) Get(ctx workflow.Context) (*directoryservice.DeregisterEventTopicOutput, error) {
    var output directoryservice.DeregisterEventTopicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeCertificateResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeCertificateResult) Get(ctx workflow.Context) (*directoryservice.DescribeCertificateOutput, error) {
    var output directoryservice.DescribeCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeConditionalForwardersResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeConditionalForwardersResult) Get(ctx workflow.Context) (*directoryservice.DescribeConditionalForwardersOutput, error) {
    var output directoryservice.DescribeConditionalForwardersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeDirectoriesResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeDirectoriesResult) Get(ctx workflow.Context) (*directoryservice.DescribeDirectoriesOutput, error) {
    var output directoryservice.DescribeDirectoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeDomainControllersResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeDomainControllersResult) Get(ctx workflow.Context) (*directoryservice.DescribeDomainControllersOutput, error) {
    var output directoryservice.DescribeDomainControllersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeEventTopicsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeEventTopicsResult) Get(ctx workflow.Context) (*directoryservice.DescribeEventTopicsOutput, error) {
    var output directoryservice.DescribeEventTopicsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeLDAPSSettingsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeLDAPSSettingsResult) Get(ctx workflow.Context) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
    var output directoryservice.DescribeLDAPSSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeSharedDirectoriesResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeSharedDirectoriesResult) Get(ctx workflow.Context) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
    var output directoryservice.DescribeSharedDirectoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeSnapshotsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeSnapshotsResult) Get(ctx workflow.Context) (*directoryservice.DescribeSnapshotsOutput, error) {
    var output directoryservice.DescribeSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDescribeTrustsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDescribeTrustsResult) Get(ctx workflow.Context) (*directoryservice.DescribeTrustsOutput, error) {
    var output directoryservice.DescribeTrustsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDisableLDAPSResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDisableLDAPSResult) Get(ctx workflow.Context) (*directoryservice.DisableLDAPSOutput, error) {
    var output directoryservice.DisableLDAPSOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDisableRadiusResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDisableRadiusResult) Get(ctx workflow.Context) (*directoryservice.DisableRadiusOutput, error) {
    var output directoryservice.DisableRadiusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceDisableSsoResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceDisableSsoResult) Get(ctx workflow.Context) (*directoryservice.DisableSsoOutput, error) {
    var output directoryservice.DisableSsoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceEnableLDAPSResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceEnableLDAPSResult) Get(ctx workflow.Context) (*directoryservice.EnableLDAPSOutput, error) {
    var output directoryservice.EnableLDAPSOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceEnableRadiusResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceEnableRadiusResult) Get(ctx workflow.Context) (*directoryservice.EnableRadiusOutput, error) {
    var output directoryservice.EnableRadiusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceEnableSsoResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceEnableSsoResult) Get(ctx workflow.Context) (*directoryservice.EnableSsoOutput, error) {
    var output directoryservice.EnableSsoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceGetDirectoryLimitsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceGetDirectoryLimitsResult) Get(ctx workflow.Context) (*directoryservice.GetDirectoryLimitsOutput, error) {
    var output directoryservice.GetDirectoryLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceGetSnapshotLimitsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceGetSnapshotLimitsResult) Get(ctx workflow.Context) (*directoryservice.GetSnapshotLimitsOutput, error) {
    var output directoryservice.GetSnapshotLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceListCertificatesResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceListCertificatesResult) Get(ctx workflow.Context) (*directoryservice.ListCertificatesOutput, error) {
    var output directoryservice.ListCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceListIpRoutesResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceListIpRoutesResult) Get(ctx workflow.Context) (*directoryservice.ListIpRoutesOutput, error) {
    var output directoryservice.ListIpRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceListLogSubscriptionsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceListLogSubscriptionsResult) Get(ctx workflow.Context) (*directoryservice.ListLogSubscriptionsOutput, error) {
    var output directoryservice.ListLogSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceListSchemaExtensionsResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceListSchemaExtensionsResult) Get(ctx workflow.Context) (*directoryservice.ListSchemaExtensionsOutput, error) {
    var output directoryservice.ListSchemaExtensionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceListTagsForResourceResult) Get(ctx workflow.Context) (*directoryservice.ListTagsForResourceOutput, error) {
    var output directoryservice.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceRegisterCertificateResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceRegisterCertificateResult) Get(ctx workflow.Context) (*directoryservice.RegisterCertificateOutput, error) {
    var output directoryservice.RegisterCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceRegisterEventTopicResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceRegisterEventTopicResult) Get(ctx workflow.Context) (*directoryservice.RegisterEventTopicOutput, error) {
    var output directoryservice.RegisterEventTopicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceRejectSharedDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceRejectSharedDirectoryResult) Get(ctx workflow.Context) (*directoryservice.RejectSharedDirectoryOutput, error) {
    var output directoryservice.RejectSharedDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceRemoveIpRoutesResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceRemoveIpRoutesResult) Get(ctx workflow.Context) (*directoryservice.RemoveIpRoutesOutput, error) {
    var output directoryservice.RemoveIpRoutesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*directoryservice.RemoveTagsFromResourceOutput, error) {
    var output directoryservice.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceResetUserPasswordResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceResetUserPasswordResult) Get(ctx workflow.Context) (*directoryservice.ResetUserPasswordOutput, error) {
    var output directoryservice.ResetUserPasswordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceRestoreFromSnapshotResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceRestoreFromSnapshotResult) Get(ctx workflow.Context) (*directoryservice.RestoreFromSnapshotOutput, error) {
    var output directoryservice.RestoreFromSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceShareDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceShareDirectoryResult) Get(ctx workflow.Context) (*directoryservice.ShareDirectoryOutput, error) {
    var output directoryservice.ShareDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceStartSchemaExtensionResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceStartSchemaExtensionResult) Get(ctx workflow.Context) (*directoryservice.StartSchemaExtensionOutput, error) {
    var output directoryservice.StartSchemaExtensionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceUnshareDirectoryResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceUnshareDirectoryResult) Get(ctx workflow.Context) (*directoryservice.UnshareDirectoryOutput, error) {
    var output directoryservice.UnshareDirectoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceUpdateConditionalForwarderResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceUpdateConditionalForwarderResult) Get(ctx workflow.Context) (*directoryservice.UpdateConditionalForwarderOutput, error) {
    var output directoryservice.UpdateConditionalForwarderOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceUpdateNumberOfDomainControllersResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceUpdateNumberOfDomainControllersResult) Get(ctx workflow.Context) (*directoryservice.UpdateNumberOfDomainControllersOutput, error) {
    var output directoryservice.UpdateNumberOfDomainControllersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceUpdateRadiusResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceUpdateRadiusResult) Get(ctx workflow.Context) (*directoryservice.UpdateRadiusOutput, error) {
    var output directoryservice.UpdateRadiusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceUpdateTrustResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceUpdateTrustResult) Get(ctx workflow.Context) (*directoryservice.UpdateTrustOutput, error) {
    var output directoryservice.UpdateTrustOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type DirectoryserviceVerifyTrustResult struct {
	Result workflow.Future
}

func (r *DirectoryserviceVerifyTrustResult) Get(ctx workflow.Context) (*directoryservice.VerifyTrustOutput, error) {
    var output directoryservice.VerifyTrustOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) AcceptSharedDirectory(ctx workflow.Context, input *directoryservice.AcceptSharedDirectoryInput) (*directoryservice.AcceptSharedDirectoryOutput, error) {
    var output directoryservice.AcceptSharedDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.AcceptSharedDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) AcceptSharedDirectoryAsync(ctx workflow.Context, input *directoryservice.AcceptSharedDirectoryInput) *DirectoryserviceAcceptSharedDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.AcceptSharedDirectory", input)
    return &DirectoryserviceAcceptSharedDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) AddIpRoutes(ctx workflow.Context, input *directoryservice.AddIpRoutesInput) (*directoryservice.AddIpRoutesOutput, error) {
    var output directoryservice.AddIpRoutesOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.AddIpRoutes", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) AddIpRoutesAsync(ctx workflow.Context, input *directoryservice.AddIpRoutesInput) *DirectoryserviceAddIpRoutesResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.AddIpRoutes", input)
    return &DirectoryserviceAddIpRoutesResult{Result: future}
}

func (a *DirectoryServiceStub) AddTagsToResource(ctx workflow.Context, input *directoryservice.AddTagsToResourceInput) (*directoryservice.AddTagsToResourceOutput, error) {
    var output directoryservice.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.AddTagsToResource", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) AddTagsToResourceAsync(ctx workflow.Context, input *directoryservice.AddTagsToResourceInput) *DirectoryserviceAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.AddTagsToResource", input)
    return &DirectoryserviceAddTagsToResourceResult{Result: future}
}

func (a *DirectoryServiceStub) CancelSchemaExtension(ctx workflow.Context, input *directoryservice.CancelSchemaExtensionInput) (*directoryservice.CancelSchemaExtensionOutput, error) {
    var output directoryservice.CancelSchemaExtensionOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CancelSchemaExtension", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CancelSchemaExtensionAsync(ctx workflow.Context, input *directoryservice.CancelSchemaExtensionInput) *DirectoryserviceCancelSchemaExtensionResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CancelSchemaExtension", input)
    return &DirectoryserviceCancelSchemaExtensionResult{Result: future}
}

func (a *DirectoryServiceStub) ConnectDirectory(ctx workflow.Context, input *directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error) {
    var output directoryservice.ConnectDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ConnectDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ConnectDirectoryAsync(ctx workflow.Context, input *directoryservice.ConnectDirectoryInput) *DirectoryserviceConnectDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ConnectDirectory", input)
    return &DirectoryserviceConnectDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) CreateAlias(ctx workflow.Context, input *directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error) {
    var output directoryservice.CreateAliasOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateAlias", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateAliasAsync(ctx workflow.Context, input *directoryservice.CreateAliasInput) *DirectoryserviceCreateAliasResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateAlias", input)
    return &DirectoryserviceCreateAliasResult{Result: future}
}

func (a *DirectoryServiceStub) CreateComputer(ctx workflow.Context, input *directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error) {
    var output directoryservice.CreateComputerOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateComputer", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateComputerAsync(ctx workflow.Context, input *directoryservice.CreateComputerInput) *DirectoryserviceCreateComputerResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateComputer", input)
    return &DirectoryserviceCreateComputerResult{Result: future}
}

func (a *DirectoryServiceStub) CreateConditionalForwarder(ctx workflow.Context, input *directoryservice.CreateConditionalForwarderInput) (*directoryservice.CreateConditionalForwarderOutput, error) {
    var output directoryservice.CreateConditionalForwarderOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateConditionalForwarder", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.CreateConditionalForwarderInput) *DirectoryserviceCreateConditionalForwarderResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateConditionalForwarder", input)
    return &DirectoryserviceCreateConditionalForwarderResult{Result: future}
}

func (a *DirectoryServiceStub) CreateDirectory(ctx workflow.Context, input *directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error) {
    var output directoryservice.CreateDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateDirectoryAsync(ctx workflow.Context, input *directoryservice.CreateDirectoryInput) *DirectoryserviceCreateDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateDirectory", input)
    return &DirectoryserviceCreateDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) CreateLogSubscription(ctx workflow.Context, input *directoryservice.CreateLogSubscriptionInput) (*directoryservice.CreateLogSubscriptionOutput, error) {
    var output directoryservice.CreateLogSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateLogSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateLogSubscriptionAsync(ctx workflow.Context, input *directoryservice.CreateLogSubscriptionInput) *DirectoryserviceCreateLogSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateLogSubscription", input)
    return &DirectoryserviceCreateLogSubscriptionResult{Result: future}
}

func (a *DirectoryServiceStub) CreateMicrosoftAD(ctx workflow.Context, input *directoryservice.CreateMicrosoftADInput) (*directoryservice.CreateMicrosoftADOutput, error) {
    var output directoryservice.CreateMicrosoftADOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateMicrosoftAD", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateMicrosoftADAsync(ctx workflow.Context, input *directoryservice.CreateMicrosoftADInput) *DirectoryserviceCreateMicrosoftADResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateMicrosoftAD", input)
    return &DirectoryserviceCreateMicrosoftADResult{Result: future}
}

func (a *DirectoryServiceStub) CreateSnapshot(ctx workflow.Context, input *directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error) {
    var output directoryservice.CreateSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateSnapshotAsync(ctx workflow.Context, input *directoryservice.CreateSnapshotInput) *DirectoryserviceCreateSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateSnapshot", input)
    return &DirectoryserviceCreateSnapshotResult{Result: future}
}

func (a *DirectoryServiceStub) CreateTrust(ctx workflow.Context, input *directoryservice.CreateTrustInput) (*directoryservice.CreateTrustOutput, error) {
    var output directoryservice.CreateTrustOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.CreateTrust", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) CreateTrustAsync(ctx workflow.Context, input *directoryservice.CreateTrustInput) *DirectoryserviceCreateTrustResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.CreateTrust", input)
    return &DirectoryserviceCreateTrustResult{Result: future}
}

func (a *DirectoryServiceStub) DeleteConditionalForwarder(ctx workflow.Context, input *directoryservice.DeleteConditionalForwarderInput) (*directoryservice.DeleteConditionalForwarderOutput, error) {
    var output directoryservice.DeleteConditionalForwarderOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteConditionalForwarder", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeleteConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.DeleteConditionalForwarderInput) *DirectoryserviceDeleteConditionalForwarderResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteConditionalForwarder", input)
    return &DirectoryserviceDeleteConditionalForwarderResult{Result: future}
}

func (a *DirectoryServiceStub) DeleteDirectory(ctx workflow.Context, input *directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error) {
    var output directoryservice.DeleteDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeleteDirectoryAsync(ctx workflow.Context, input *directoryservice.DeleteDirectoryInput) *DirectoryserviceDeleteDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteDirectory", input)
    return &DirectoryserviceDeleteDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) DeleteLogSubscription(ctx workflow.Context, input *directoryservice.DeleteLogSubscriptionInput) (*directoryservice.DeleteLogSubscriptionOutput, error) {
    var output directoryservice.DeleteLogSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteLogSubscription", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeleteLogSubscriptionAsync(ctx workflow.Context, input *directoryservice.DeleteLogSubscriptionInput) *DirectoryserviceDeleteLogSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteLogSubscription", input)
    return &DirectoryserviceDeleteLogSubscriptionResult{Result: future}
}

func (a *DirectoryServiceStub) DeleteSnapshot(ctx workflow.Context, input *directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error) {
    var output directoryservice.DeleteSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeleteSnapshotAsync(ctx workflow.Context, input *directoryservice.DeleteSnapshotInput) *DirectoryserviceDeleteSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteSnapshot", input)
    return &DirectoryserviceDeleteSnapshotResult{Result: future}
}

func (a *DirectoryServiceStub) DeleteTrust(ctx workflow.Context, input *directoryservice.DeleteTrustInput) (*directoryservice.DeleteTrustOutput, error) {
    var output directoryservice.DeleteTrustOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteTrust", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeleteTrustAsync(ctx workflow.Context, input *directoryservice.DeleteTrustInput) *DirectoryserviceDeleteTrustResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeleteTrust", input)
    return &DirectoryserviceDeleteTrustResult{Result: future}
}

func (a *DirectoryServiceStub) DeregisterCertificate(ctx workflow.Context, input *directoryservice.DeregisterCertificateInput) (*directoryservice.DeregisterCertificateOutput, error) {
    var output directoryservice.DeregisterCertificateOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeregisterCertificate", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeregisterCertificateAsync(ctx workflow.Context, input *directoryservice.DeregisterCertificateInput) *DirectoryserviceDeregisterCertificateResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeregisterCertificate", input)
    return &DirectoryserviceDeregisterCertificateResult{Result: future}
}

func (a *DirectoryServiceStub) DeregisterEventTopic(ctx workflow.Context, input *directoryservice.DeregisterEventTopicInput) (*directoryservice.DeregisterEventTopicOutput, error) {
    var output directoryservice.DeregisterEventTopicOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DeregisterEventTopic", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DeregisterEventTopicAsync(ctx workflow.Context, input *directoryservice.DeregisterEventTopicInput) *DirectoryserviceDeregisterEventTopicResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DeregisterEventTopic", input)
    return &DirectoryserviceDeregisterEventTopicResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeCertificate(ctx workflow.Context, input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error) {
    var output directoryservice.DescribeCertificateOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeCertificate", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeCertificateAsync(ctx workflow.Context, input *directoryservice.DescribeCertificateInput) *DirectoryserviceDescribeCertificateResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeCertificate", input)
    return &DirectoryserviceDescribeCertificateResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeConditionalForwarders(ctx workflow.Context, input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error) {
    var output directoryservice.DescribeConditionalForwardersOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeConditionalForwarders", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeConditionalForwardersAsync(ctx workflow.Context, input *directoryservice.DescribeConditionalForwardersInput) *DirectoryserviceDescribeConditionalForwardersResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeConditionalForwarders", input)
    return &DirectoryserviceDescribeConditionalForwardersResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeDirectories(ctx workflow.Context, input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error) {
    var output directoryservice.DescribeDirectoriesOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeDirectories", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeDirectoriesAsync(ctx workflow.Context, input *directoryservice.DescribeDirectoriesInput) *DirectoryserviceDescribeDirectoriesResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeDirectories", input)
    return &DirectoryserviceDescribeDirectoriesResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeDomainControllers(ctx workflow.Context, input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error) {
    var output directoryservice.DescribeDomainControllersOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeDomainControllers", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeDomainControllersAsync(ctx workflow.Context, input *directoryservice.DescribeDomainControllersInput) *DirectoryserviceDescribeDomainControllersResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeDomainControllers", input)
    return &DirectoryserviceDescribeDomainControllersResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeEventTopics(ctx workflow.Context, input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error) {
    var output directoryservice.DescribeEventTopicsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeEventTopics", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeEventTopicsAsync(ctx workflow.Context, input *directoryservice.DescribeEventTopicsInput) *DirectoryserviceDescribeEventTopicsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeEventTopics", input)
    return &DirectoryserviceDescribeEventTopicsResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeLDAPSSettings(ctx workflow.Context, input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
    var output directoryservice.DescribeLDAPSSettingsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeLDAPSSettings", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeLDAPSSettingsAsync(ctx workflow.Context, input *directoryservice.DescribeLDAPSSettingsInput) *DirectoryserviceDescribeLDAPSSettingsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeLDAPSSettings", input)
    return &DirectoryserviceDescribeLDAPSSettingsResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeSharedDirectories(ctx workflow.Context, input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
    var output directoryservice.DescribeSharedDirectoriesOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeSharedDirectories", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeSharedDirectoriesAsync(ctx workflow.Context, input *directoryservice.DescribeSharedDirectoriesInput) *DirectoryserviceDescribeSharedDirectoriesResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeSharedDirectories", input)
    return &DirectoryserviceDescribeSharedDirectoriesResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeSnapshots(ctx workflow.Context, input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error) {
    var output directoryservice.DescribeSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeSnapshots", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeSnapshotsAsync(ctx workflow.Context, input *directoryservice.DescribeSnapshotsInput) *DirectoryserviceDescribeSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeSnapshots", input)
    return &DirectoryserviceDescribeSnapshotsResult{Result: future}
}

func (a *DirectoryServiceStub) DescribeTrusts(ctx workflow.Context, input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error) {
    var output directoryservice.DescribeTrustsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeTrusts", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DescribeTrustsAsync(ctx workflow.Context, input *directoryservice.DescribeTrustsInput) *DirectoryserviceDescribeTrustsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DescribeTrusts", input)
    return &DirectoryserviceDescribeTrustsResult{Result: future}
}

func (a *DirectoryServiceStub) DisableLDAPS(ctx workflow.Context, input *directoryservice.DisableLDAPSInput) (*directoryservice.DisableLDAPSOutput, error) {
    var output directoryservice.DisableLDAPSOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DisableLDAPS", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DisableLDAPSAsync(ctx workflow.Context, input *directoryservice.DisableLDAPSInput) *DirectoryserviceDisableLDAPSResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DisableLDAPS", input)
    return &DirectoryserviceDisableLDAPSResult{Result: future}
}

func (a *DirectoryServiceStub) DisableRadius(ctx workflow.Context, input *directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error) {
    var output directoryservice.DisableRadiusOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DisableRadius", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DisableRadiusAsync(ctx workflow.Context, input *directoryservice.DisableRadiusInput) *DirectoryserviceDisableRadiusResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DisableRadius", input)
    return &DirectoryserviceDisableRadiusResult{Result: future}
}

func (a *DirectoryServiceStub) DisableSso(ctx workflow.Context, input *directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error) {
    var output directoryservice.DisableSsoOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.DisableSso", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) DisableSsoAsync(ctx workflow.Context, input *directoryservice.DisableSsoInput) *DirectoryserviceDisableSsoResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.DisableSso", input)
    return &DirectoryserviceDisableSsoResult{Result: future}
}

func (a *DirectoryServiceStub) EnableLDAPS(ctx workflow.Context, input *directoryservice.EnableLDAPSInput) (*directoryservice.EnableLDAPSOutput, error) {
    var output directoryservice.EnableLDAPSOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.EnableLDAPS", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) EnableLDAPSAsync(ctx workflow.Context, input *directoryservice.EnableLDAPSInput) *DirectoryserviceEnableLDAPSResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.EnableLDAPS", input)
    return &DirectoryserviceEnableLDAPSResult{Result: future}
}

func (a *DirectoryServiceStub) EnableRadius(ctx workflow.Context, input *directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error) {
    var output directoryservice.EnableRadiusOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.EnableRadius", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) EnableRadiusAsync(ctx workflow.Context, input *directoryservice.EnableRadiusInput) *DirectoryserviceEnableRadiusResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.EnableRadius", input)
    return &DirectoryserviceEnableRadiusResult{Result: future}
}

func (a *DirectoryServiceStub) EnableSso(ctx workflow.Context, input *directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error) {
    var output directoryservice.EnableSsoOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.EnableSso", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) EnableSsoAsync(ctx workflow.Context, input *directoryservice.EnableSsoInput) *DirectoryserviceEnableSsoResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.EnableSso", input)
    return &DirectoryserviceEnableSsoResult{Result: future}
}

func (a *DirectoryServiceStub) GetDirectoryLimits(ctx workflow.Context, input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error) {
    var output directoryservice.GetDirectoryLimitsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.GetDirectoryLimits", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) GetDirectoryLimitsAsync(ctx workflow.Context, input *directoryservice.GetDirectoryLimitsInput) *DirectoryserviceGetDirectoryLimitsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.GetDirectoryLimits", input)
    return &DirectoryserviceGetDirectoryLimitsResult{Result: future}
}

func (a *DirectoryServiceStub) GetSnapshotLimits(ctx workflow.Context, input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error) {
    var output directoryservice.GetSnapshotLimitsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.GetSnapshotLimits", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) GetSnapshotLimitsAsync(ctx workflow.Context, input *directoryservice.GetSnapshotLimitsInput) *DirectoryserviceGetSnapshotLimitsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.GetSnapshotLimits", input)
    return &DirectoryserviceGetSnapshotLimitsResult{Result: future}
}

func (a *DirectoryServiceStub) ListCertificates(ctx workflow.Context, input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error) {
    var output directoryservice.ListCertificatesOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ListCertificates", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ListCertificatesAsync(ctx workflow.Context, input *directoryservice.ListCertificatesInput) *DirectoryserviceListCertificatesResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ListCertificates", input)
    return &DirectoryserviceListCertificatesResult{Result: future}
}

func (a *DirectoryServiceStub) ListIpRoutes(ctx workflow.Context, input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error) {
    var output directoryservice.ListIpRoutesOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ListIpRoutes", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ListIpRoutesAsync(ctx workflow.Context, input *directoryservice.ListIpRoutesInput) *DirectoryserviceListIpRoutesResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ListIpRoutes", input)
    return &DirectoryserviceListIpRoutesResult{Result: future}
}

func (a *DirectoryServiceStub) ListLogSubscriptions(ctx workflow.Context, input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error) {
    var output directoryservice.ListLogSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ListLogSubscriptions", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ListLogSubscriptionsAsync(ctx workflow.Context, input *directoryservice.ListLogSubscriptionsInput) *DirectoryserviceListLogSubscriptionsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ListLogSubscriptions", input)
    return &DirectoryserviceListLogSubscriptionsResult{Result: future}
}

func (a *DirectoryServiceStub) ListSchemaExtensions(ctx workflow.Context, input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error) {
    var output directoryservice.ListSchemaExtensionsOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ListSchemaExtensions", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ListSchemaExtensionsAsync(ctx workflow.Context, input *directoryservice.ListSchemaExtensionsInput) *DirectoryserviceListSchemaExtensionsResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ListSchemaExtensions", input)
    return &DirectoryserviceListSchemaExtensionsResult{Result: future}
}

func (a *DirectoryServiceStub) ListTagsForResource(ctx workflow.Context, input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error) {
    var output directoryservice.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *directoryservice.ListTagsForResourceInput) *DirectoryserviceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ListTagsForResource", input)
    return &DirectoryserviceListTagsForResourceResult{Result: future}
}

func (a *DirectoryServiceStub) RegisterCertificate(ctx workflow.Context, input *directoryservice.RegisterCertificateInput) (*directoryservice.RegisterCertificateOutput, error) {
    var output directoryservice.RegisterCertificateOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.RegisterCertificate", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) RegisterCertificateAsync(ctx workflow.Context, input *directoryservice.RegisterCertificateInput) *DirectoryserviceRegisterCertificateResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.RegisterCertificate", input)
    return &DirectoryserviceRegisterCertificateResult{Result: future}
}

func (a *DirectoryServiceStub) RegisterEventTopic(ctx workflow.Context, input *directoryservice.RegisterEventTopicInput) (*directoryservice.RegisterEventTopicOutput, error) {
    var output directoryservice.RegisterEventTopicOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.RegisterEventTopic", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) RegisterEventTopicAsync(ctx workflow.Context, input *directoryservice.RegisterEventTopicInput) *DirectoryserviceRegisterEventTopicResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.RegisterEventTopic", input)
    return &DirectoryserviceRegisterEventTopicResult{Result: future}
}

func (a *DirectoryServiceStub) RejectSharedDirectory(ctx workflow.Context, input *directoryservice.RejectSharedDirectoryInput) (*directoryservice.RejectSharedDirectoryOutput, error) {
    var output directoryservice.RejectSharedDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.RejectSharedDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) RejectSharedDirectoryAsync(ctx workflow.Context, input *directoryservice.RejectSharedDirectoryInput) *DirectoryserviceRejectSharedDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.RejectSharedDirectory", input)
    return &DirectoryserviceRejectSharedDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) RemoveIpRoutes(ctx workflow.Context, input *directoryservice.RemoveIpRoutesInput) (*directoryservice.RemoveIpRoutesOutput, error) {
    var output directoryservice.RemoveIpRoutesOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.RemoveIpRoutes", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) RemoveIpRoutesAsync(ctx workflow.Context, input *directoryservice.RemoveIpRoutesInput) *DirectoryserviceRemoveIpRoutesResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.RemoveIpRoutes", input)
    return &DirectoryserviceRemoveIpRoutesResult{Result: future}
}

func (a *DirectoryServiceStub) RemoveTagsFromResource(ctx workflow.Context, input *directoryservice.RemoveTagsFromResourceInput) (*directoryservice.RemoveTagsFromResourceOutput, error) {
    var output directoryservice.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.RemoveTagsFromResource", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *directoryservice.RemoveTagsFromResourceInput) *DirectoryserviceRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.RemoveTagsFromResource", input)
    return &DirectoryserviceRemoveTagsFromResourceResult{Result: future}
}

func (a *DirectoryServiceStub) ResetUserPassword(ctx workflow.Context, input *directoryservice.ResetUserPasswordInput) (*directoryservice.ResetUserPasswordOutput, error) {
    var output directoryservice.ResetUserPasswordOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ResetUserPassword", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ResetUserPasswordAsync(ctx workflow.Context, input *directoryservice.ResetUserPasswordInput) *DirectoryserviceResetUserPasswordResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ResetUserPassword", input)
    return &DirectoryserviceResetUserPasswordResult{Result: future}
}

func (a *DirectoryServiceStub) RestoreFromSnapshot(ctx workflow.Context, input *directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error) {
    var output directoryservice.RestoreFromSnapshotOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.RestoreFromSnapshot", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) RestoreFromSnapshotAsync(ctx workflow.Context, input *directoryservice.RestoreFromSnapshotInput) *DirectoryserviceRestoreFromSnapshotResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.RestoreFromSnapshot", input)
    return &DirectoryserviceRestoreFromSnapshotResult{Result: future}
}

func (a *DirectoryServiceStub) ShareDirectory(ctx workflow.Context, input *directoryservice.ShareDirectoryInput) (*directoryservice.ShareDirectoryOutput, error) {
    var output directoryservice.ShareDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.ShareDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) ShareDirectoryAsync(ctx workflow.Context, input *directoryservice.ShareDirectoryInput) *DirectoryserviceShareDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.ShareDirectory", input)
    return &DirectoryserviceShareDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) StartSchemaExtension(ctx workflow.Context, input *directoryservice.StartSchemaExtensionInput) (*directoryservice.StartSchemaExtensionOutput, error) {
    var output directoryservice.StartSchemaExtensionOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.StartSchemaExtension", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) StartSchemaExtensionAsync(ctx workflow.Context, input *directoryservice.StartSchemaExtensionInput) *DirectoryserviceStartSchemaExtensionResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.StartSchemaExtension", input)
    return &DirectoryserviceStartSchemaExtensionResult{Result: future}
}

func (a *DirectoryServiceStub) UnshareDirectory(ctx workflow.Context, input *directoryservice.UnshareDirectoryInput) (*directoryservice.UnshareDirectoryOutput, error) {
    var output directoryservice.UnshareDirectoryOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.UnshareDirectory", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) UnshareDirectoryAsync(ctx workflow.Context, input *directoryservice.UnshareDirectoryInput) *DirectoryserviceUnshareDirectoryResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.UnshareDirectory", input)
    return &DirectoryserviceUnshareDirectoryResult{Result: future}
}

func (a *DirectoryServiceStub) UpdateConditionalForwarder(ctx workflow.Context, input *directoryservice.UpdateConditionalForwarderInput) (*directoryservice.UpdateConditionalForwarderOutput, error) {
    var output directoryservice.UpdateConditionalForwarderOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateConditionalForwarder", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) UpdateConditionalForwarderAsync(ctx workflow.Context, input *directoryservice.UpdateConditionalForwarderInput) *DirectoryserviceUpdateConditionalForwarderResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateConditionalForwarder", input)
    return &DirectoryserviceUpdateConditionalForwarderResult{Result: future}
}

func (a *DirectoryServiceStub) UpdateNumberOfDomainControllers(ctx workflow.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) (*directoryservice.UpdateNumberOfDomainControllersOutput, error) {
    var output directoryservice.UpdateNumberOfDomainControllersOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateNumberOfDomainControllers", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) UpdateNumberOfDomainControllersAsync(ctx workflow.Context, input *directoryservice.UpdateNumberOfDomainControllersInput) *DirectoryserviceUpdateNumberOfDomainControllersResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateNumberOfDomainControllers", input)
    return &DirectoryserviceUpdateNumberOfDomainControllersResult{Result: future}
}

func (a *DirectoryServiceStub) UpdateRadius(ctx workflow.Context, input *directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error) {
    var output directoryservice.UpdateRadiusOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateRadius", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) UpdateRadiusAsync(ctx workflow.Context, input *directoryservice.UpdateRadiusInput) *DirectoryserviceUpdateRadiusResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateRadius", input)
    return &DirectoryserviceUpdateRadiusResult{Result: future}
}

func (a *DirectoryServiceStub) UpdateTrust(ctx workflow.Context, input *directoryservice.UpdateTrustInput) (*directoryservice.UpdateTrustOutput, error) {
    var output directoryservice.UpdateTrustOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateTrust", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) UpdateTrustAsync(ctx workflow.Context, input *directoryservice.UpdateTrustInput) *DirectoryserviceUpdateTrustResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.UpdateTrust", input)
    return &DirectoryserviceUpdateTrustResult{Result: future}
}

func (a *DirectoryServiceStub) VerifyTrust(ctx workflow.Context, input *directoryservice.VerifyTrustInput) (*directoryservice.VerifyTrustOutput, error) {
    var output directoryservice.VerifyTrustOutput
    err := workflow.ExecuteActivity(ctx, "DirectoryService.VerifyTrust", input).Get(ctx, &output)
    return &output, err
}

func (a *DirectoryServiceStub) VerifyTrustAsync(ctx workflow.Context, input *directoryservice.VerifyTrustInput) *DirectoryserviceVerifyTrustResult {
    future := workflow.ExecuteActivity(ctx, "DirectoryService.VerifyTrust", input)
    return &DirectoryserviceVerifyTrustResult{Result: future}
}
