
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type StorageGatewayActivities struct {
    client storagegatewayiface.StorageGatewayAPI
}

func NewStorageGatewayActivities(session *session.Session, config ...*aws.Config) *StorageGatewayActivities {
    client := storagegateway.New(session, config...)
    return &StorageGatewayActivities{client: client}
}

func (a *StorageGatewayActivities) ActivateGateway(ctx context.Context, input *storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error) {
    return a.client.ActivateGatewayWithContext(ctx, input)
}

func (a *StorageGatewayActivities) AddCache(ctx context.Context, input *storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error) {
    return a.client.AddCacheWithContext(ctx, input)
}

func (a *StorageGatewayActivities) AddTagsToResource(ctx context.Context, input *storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *StorageGatewayActivities) AddUploadBuffer(ctx context.Context, input *storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error) {
    return a.client.AddUploadBufferWithContext(ctx, input)
}

func (a *StorageGatewayActivities) AddWorkingStorage(ctx context.Context, input *storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error) {
    return a.client.AddWorkingStorageWithContext(ctx, input)
}

func (a *StorageGatewayActivities) AssignTapePool(ctx context.Context, input *storagegateway.AssignTapePoolInput) (*storagegateway.AssignTapePoolOutput, error) {
    return a.client.AssignTapePoolWithContext(ctx, input)
}

func (a *StorageGatewayActivities) AttachVolume(ctx context.Context, input *storagegateway.AttachVolumeInput) (*storagegateway.AttachVolumeOutput, error) {
    return a.client.AttachVolumeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CancelArchival(ctx context.Context, input *storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error) {
    return a.client.CancelArchivalWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CancelRetrieval(ctx context.Context, input *storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error) {
    return a.client.CancelRetrievalWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateCachediSCSIVolume(ctx context.Context, input *storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateCachediSCSIVolumeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateNFSFileShare(ctx context.Context, input *storagegateway.CreateNFSFileShareInput) (*storagegateway.CreateNFSFileShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateNFSFileShareWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateSMBFileShare(ctx context.Context, input *storagegateway.CreateSMBFileShareInput) (*storagegateway.CreateSMBFileShareOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateSMBFileShareWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateSnapshot(ctx context.Context, input *storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error) {
    return a.client.CreateSnapshotWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateSnapshotFromVolumeRecoveryPoint(ctx context.Context, input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error) {
    return a.client.CreateSnapshotFromVolumeRecoveryPointWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateStorediSCSIVolume(ctx context.Context, input *storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error) {
    return a.client.CreateStorediSCSIVolumeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateTapePool(ctx context.Context, input *storagegateway.CreateTapePoolInput) (*storagegateway.CreateTapePoolOutput, error) {
    return a.client.CreateTapePoolWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateTapeWithBarcode(ctx context.Context, input *storagegateway.CreateTapeWithBarcodeInput) (*storagegateway.CreateTapeWithBarcodeOutput, error) {
    return a.client.CreateTapeWithBarcodeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) CreateTapes(ctx context.Context, input *storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateTapesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteAutomaticTapeCreationPolicy(ctx context.Context, input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error) {
    return a.client.DeleteAutomaticTapeCreationPolicyWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteBandwidthRateLimit(ctx context.Context, input *storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error) {
    return a.client.DeleteBandwidthRateLimitWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteChapCredentials(ctx context.Context, input *storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error) {
    return a.client.DeleteChapCredentialsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteFileShare(ctx context.Context, input *storagegateway.DeleteFileShareInput) (*storagegateway.DeleteFileShareOutput, error) {
    return a.client.DeleteFileShareWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteGateway(ctx context.Context, input *storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error) {
    return a.client.DeleteGatewayWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteSnapshotSchedule(ctx context.Context, input *storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error) {
    return a.client.DeleteSnapshotScheduleWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteTape(ctx context.Context, input *storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error) {
    return a.client.DeleteTapeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteTapeArchive(ctx context.Context, input *storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error) {
    return a.client.DeleteTapeArchiveWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteTapePool(ctx context.Context, input *storagegateway.DeleteTapePoolInput) (*storagegateway.DeleteTapePoolOutput, error) {
    return a.client.DeleteTapePoolWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DeleteVolume(ctx context.Context, input *storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error) {
    return a.client.DeleteVolumeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeAvailabilityMonitorTest(ctx context.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
    return a.client.DescribeAvailabilityMonitorTestWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeBandwidthRateLimit(ctx context.Context, input *storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
    return a.client.DescribeBandwidthRateLimitWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeCache(ctx context.Context, input *storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error) {
    return a.client.DescribeCacheWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeCachediSCSIVolumes(ctx context.Context, input *storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
    return a.client.DescribeCachediSCSIVolumesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeChapCredentials(ctx context.Context, input *storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error) {
    return a.client.DescribeChapCredentialsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeGatewayInformation(ctx context.Context, input *storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error) {
    return a.client.DescribeGatewayInformationWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeMaintenanceStartTime(ctx context.Context, input *storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
    return a.client.DescribeMaintenanceStartTimeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeNFSFileShares(ctx context.Context, input *storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error) {
    return a.client.DescribeNFSFileSharesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeSMBFileShares(ctx context.Context, input *storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error) {
    return a.client.DescribeSMBFileSharesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeSMBSettings(ctx context.Context, input *storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error) {
    return a.client.DescribeSMBSettingsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeSnapshotSchedule(ctx context.Context, input *storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
    return a.client.DescribeSnapshotScheduleWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeStorediSCSIVolumes(ctx context.Context, input *storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
    return a.client.DescribeStorediSCSIVolumesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeTapeArchives(ctx context.Context, input *storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error) {
    return a.client.DescribeTapeArchivesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeTapeRecoveryPoints(ctx context.Context, input *storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
    return a.client.DescribeTapeRecoveryPointsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeTapes(ctx context.Context, input *storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error) {
    return a.client.DescribeTapesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeUploadBuffer(ctx context.Context, input *storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error) {
    return a.client.DescribeUploadBufferWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeVTLDevices(ctx context.Context, input *storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error) {
    return a.client.DescribeVTLDevicesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DescribeWorkingStorage(ctx context.Context, input *storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error) {
    return a.client.DescribeWorkingStorageWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DetachVolume(ctx context.Context, input *storagegateway.DetachVolumeInput) (*storagegateway.DetachVolumeOutput, error) {
    return a.client.DetachVolumeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) DisableGateway(ctx context.Context, input *storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error) {
    return a.client.DisableGatewayWithContext(ctx, input)
}

func (a *StorageGatewayActivities) JoinDomain(ctx context.Context, input *storagegateway.JoinDomainInput) (*storagegateway.JoinDomainOutput, error) {
    return a.client.JoinDomainWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListAutomaticTapeCreationPolicies(ctx context.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
    return a.client.ListAutomaticTapeCreationPoliciesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListFileShares(ctx context.Context, input *storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error) {
    return a.client.ListFileSharesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListGateways(ctx context.Context, input *storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error) {
    return a.client.ListGatewaysWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListLocalDisks(ctx context.Context, input *storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error) {
    return a.client.ListLocalDisksWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListTagsForResource(ctx context.Context, input *storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListTapePools(ctx context.Context, input *storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error) {
    return a.client.ListTapePoolsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListTapes(ctx context.Context, input *storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error) {
    return a.client.ListTapesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListVolumeInitiators(ctx context.Context, input *storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error) {
    return a.client.ListVolumeInitiatorsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListVolumeRecoveryPoints(ctx context.Context, input *storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
    return a.client.ListVolumeRecoveryPointsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ListVolumes(ctx context.Context, input *storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error) {
    return a.client.ListVolumesWithContext(ctx, input)
}

func (a *StorageGatewayActivities) NotifyWhenUploaded(ctx context.Context, input *storagegateway.NotifyWhenUploadedInput) (*storagegateway.NotifyWhenUploadedOutput, error) {
    return a.client.NotifyWhenUploadedWithContext(ctx, input)
}

func (a *StorageGatewayActivities) RefreshCache(ctx context.Context, input *storagegateway.RefreshCacheInput) (*storagegateway.RefreshCacheOutput, error) {
    return a.client.RefreshCacheWithContext(ctx, input)
}

func (a *StorageGatewayActivities) RemoveTagsFromResource(ctx context.Context, input *storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ResetCache(ctx context.Context, input *storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error) {
    return a.client.ResetCacheWithContext(ctx, input)
}

func (a *StorageGatewayActivities) RetrieveTapeArchive(ctx context.Context, input *storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error) {
    return a.client.RetrieveTapeArchiveWithContext(ctx, input)
}

func (a *StorageGatewayActivities) RetrieveTapeRecoveryPoint(ctx context.Context, input *storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error) {
    return a.client.RetrieveTapeRecoveryPointWithContext(ctx, input)
}

func (a *StorageGatewayActivities) SetLocalConsolePassword(ctx context.Context, input *storagegateway.SetLocalConsolePasswordInput) (*storagegateway.SetLocalConsolePasswordOutput, error) {
    return a.client.SetLocalConsolePasswordWithContext(ctx, input)
}

func (a *StorageGatewayActivities) SetSMBGuestPassword(ctx context.Context, input *storagegateway.SetSMBGuestPasswordInput) (*storagegateway.SetSMBGuestPasswordOutput, error) {
    return a.client.SetSMBGuestPasswordWithContext(ctx, input)
}

func (a *StorageGatewayActivities) ShutdownGateway(ctx context.Context, input *storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error) {
    return a.client.ShutdownGatewayWithContext(ctx, input)
}

func (a *StorageGatewayActivities) StartAvailabilityMonitorTest(ctx context.Context, input *storagegateway.StartAvailabilityMonitorTestInput) (*storagegateway.StartAvailabilityMonitorTestOutput, error) {
    return a.client.StartAvailabilityMonitorTestWithContext(ctx, input)
}

func (a *StorageGatewayActivities) StartGateway(ctx context.Context, input *storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error) {
    return a.client.StartGatewayWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateAutomaticTapeCreationPolicy(ctx context.Context, input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error) {
    return a.client.UpdateAutomaticTapeCreationPolicyWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateBandwidthRateLimit(ctx context.Context, input *storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error) {
    return a.client.UpdateBandwidthRateLimitWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateChapCredentials(ctx context.Context, input *storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error) {
    return a.client.UpdateChapCredentialsWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateGatewayInformation(ctx context.Context, input *storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error) {
    return a.client.UpdateGatewayInformationWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateGatewaySoftwareNow(ctx context.Context, input *storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error) {
    return a.client.UpdateGatewaySoftwareNowWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateMaintenanceStartTime(ctx context.Context, input *storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error) {
    return a.client.UpdateMaintenanceStartTimeWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateNFSFileShare(ctx context.Context, input *storagegateway.UpdateNFSFileShareInput) (*storagegateway.UpdateNFSFileShareOutput, error) {
    return a.client.UpdateNFSFileShareWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateSMBFileShare(ctx context.Context, input *storagegateway.UpdateSMBFileShareInput) (*storagegateway.UpdateSMBFileShareOutput, error) {
    return a.client.UpdateSMBFileShareWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateSMBSecurityStrategy(ctx context.Context, input *storagegateway.UpdateSMBSecurityStrategyInput) (*storagegateway.UpdateSMBSecurityStrategyOutput, error) {
    return a.client.UpdateSMBSecurityStrategyWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateSnapshotSchedule(ctx context.Context, input *storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error) {
    return a.client.UpdateSnapshotScheduleWithContext(ctx, input)
}

func (a *StorageGatewayActivities) UpdateVTLDeviceType(ctx context.Context, input *storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error) {
    return a.client.UpdateVTLDeviceTypeWithContext(ctx, input)
}
