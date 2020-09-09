
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface"
)

type StorageGatewayActivities struct {
    client storagegatewayiface.StorageGatewayAPI
}

func NewStorageGatewayActivities(session *session.Session, config ...*aws.Config) *StorageGatewayActivities {
    client := storagegateway.New(session, config...)
    return &StorageGatewayActivities{client: client}
}

func (a *StorageGatewayActivities) ActivateGateway(input *storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error) {
    return a.client.ActivateGateway(input)
}

func (a *StorageGatewayActivities) AddCache(input *storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error) {
    return a.client.AddCache(input)
}

func (a *StorageGatewayActivities) AddTagsToResource(input *storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error) {
    return a.client.AddTagsToResource(input)
}

func (a *StorageGatewayActivities) AddUploadBuffer(input *storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error) {
    return a.client.AddUploadBuffer(input)
}

func (a *StorageGatewayActivities) AddWorkingStorage(input *storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error) {
    return a.client.AddWorkingStorage(input)
}

func (a *StorageGatewayActivities) AssignTapePool(input *storagegateway.AssignTapePoolInput) (*storagegateway.AssignTapePoolOutput, error) {
    return a.client.AssignTapePool(input)
}

func (a *StorageGatewayActivities) AttachVolume(input *storagegateway.AttachVolumeInput) (*storagegateway.AttachVolumeOutput, error) {
    return a.client.AttachVolume(input)
}

func (a *StorageGatewayActivities) CancelArchival(input *storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error) {
    return a.client.CancelArchival(input)
}

func (a *StorageGatewayActivities) CancelRetrieval(input *storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error) {
    return a.client.CancelRetrieval(input)
}

func (a *StorageGatewayActivities) CreateCachediSCSIVolume(input *storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error) {
    return a.client.CreateCachediSCSIVolume(input)
}

func (a *StorageGatewayActivities) CreateNFSFileShare(input *storagegateway.CreateNFSFileShareInput) (*storagegateway.CreateNFSFileShareOutput, error) {
    return a.client.CreateNFSFileShare(input)
}

func (a *StorageGatewayActivities) CreateSMBFileShare(input *storagegateway.CreateSMBFileShareInput) (*storagegateway.CreateSMBFileShareOutput, error) {
    return a.client.CreateSMBFileShare(input)
}

func (a *StorageGatewayActivities) CreateSnapshot(input *storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error) {
    return a.client.CreateSnapshot(input)
}

func (a *StorageGatewayActivities) CreateSnapshotFromVolumeRecoveryPoint(input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error) {
    return a.client.CreateSnapshotFromVolumeRecoveryPoint(input)
}

func (a *StorageGatewayActivities) CreateStorediSCSIVolume(input *storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error) {
    return a.client.CreateStorediSCSIVolume(input)
}

func (a *StorageGatewayActivities) CreateTapePool(input *storagegateway.CreateTapePoolInput) (*storagegateway.CreateTapePoolOutput, error) {
    return a.client.CreateTapePool(input)
}

func (a *StorageGatewayActivities) CreateTapeWithBarcode(input *storagegateway.CreateTapeWithBarcodeInput) (*storagegateway.CreateTapeWithBarcodeOutput, error) {
    return a.client.CreateTapeWithBarcode(input)
}

func (a *StorageGatewayActivities) CreateTapes(input *storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error) {
    return a.client.CreateTapes(input)
}

func (a *StorageGatewayActivities) DeleteAutomaticTapeCreationPolicy(input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error) {
    return a.client.DeleteAutomaticTapeCreationPolicy(input)
}

func (a *StorageGatewayActivities) DeleteBandwidthRateLimit(input *storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error) {
    return a.client.DeleteBandwidthRateLimit(input)
}

func (a *StorageGatewayActivities) DeleteChapCredentials(input *storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error) {
    return a.client.DeleteChapCredentials(input)
}

func (a *StorageGatewayActivities) DeleteFileShare(input *storagegateway.DeleteFileShareInput) (*storagegateway.DeleteFileShareOutput, error) {
    return a.client.DeleteFileShare(input)
}

func (a *StorageGatewayActivities) DeleteGateway(input *storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error) {
    return a.client.DeleteGateway(input)
}

func (a *StorageGatewayActivities) DeleteSnapshotSchedule(input *storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error) {
    return a.client.DeleteSnapshotSchedule(input)
}

func (a *StorageGatewayActivities) DeleteTape(input *storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error) {
    return a.client.DeleteTape(input)
}

func (a *StorageGatewayActivities) DeleteTapeArchive(input *storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error) {
    return a.client.DeleteTapeArchive(input)
}

func (a *StorageGatewayActivities) DeleteTapePool(input *storagegateway.DeleteTapePoolInput) (*storagegateway.DeleteTapePoolOutput, error) {
    return a.client.DeleteTapePool(input)
}

func (a *StorageGatewayActivities) DeleteVolume(input *storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error) {
    return a.client.DeleteVolume(input)
}

func (a *StorageGatewayActivities) DescribeAvailabilityMonitorTest(input *storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
    return a.client.DescribeAvailabilityMonitorTest(input)
}

func (a *StorageGatewayActivities) DescribeBandwidthRateLimit(input *storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
    return a.client.DescribeBandwidthRateLimit(input)
}

func (a *StorageGatewayActivities) DescribeCache(input *storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error) {
    return a.client.DescribeCache(input)
}

func (a *StorageGatewayActivities) DescribeCachediSCSIVolumes(input *storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
    return a.client.DescribeCachediSCSIVolumes(input)
}

func (a *StorageGatewayActivities) DescribeChapCredentials(input *storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error) {
    return a.client.DescribeChapCredentials(input)
}

func (a *StorageGatewayActivities) DescribeGatewayInformation(input *storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error) {
    return a.client.DescribeGatewayInformation(input)
}

func (a *StorageGatewayActivities) DescribeMaintenanceStartTime(input *storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
    return a.client.DescribeMaintenanceStartTime(input)
}

func (a *StorageGatewayActivities) DescribeNFSFileShares(input *storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error) {
    return a.client.DescribeNFSFileShares(input)
}

func (a *StorageGatewayActivities) DescribeSMBFileShares(input *storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error) {
    return a.client.DescribeSMBFileShares(input)
}

func (a *StorageGatewayActivities) DescribeSMBSettings(input *storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error) {
    return a.client.DescribeSMBSettings(input)
}

func (a *StorageGatewayActivities) DescribeSnapshotSchedule(input *storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
    return a.client.DescribeSnapshotSchedule(input)
}

func (a *StorageGatewayActivities) DescribeStorediSCSIVolumes(input *storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
    return a.client.DescribeStorediSCSIVolumes(input)
}

func (a *StorageGatewayActivities) DescribeTapeArchives(input *storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error) {
    return a.client.DescribeTapeArchives(input)
}

func (a *StorageGatewayActivities) DescribeTapeRecoveryPoints(input *storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
    return a.client.DescribeTapeRecoveryPoints(input)
}

func (a *StorageGatewayActivities) DescribeTapes(input *storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error) {
    return a.client.DescribeTapes(input)
}

func (a *StorageGatewayActivities) DescribeUploadBuffer(input *storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error) {
    return a.client.DescribeUploadBuffer(input)
}

func (a *StorageGatewayActivities) DescribeVTLDevices(input *storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error) {
    return a.client.DescribeVTLDevices(input)
}

func (a *StorageGatewayActivities) DescribeWorkingStorage(input *storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error) {
    return a.client.DescribeWorkingStorage(input)
}

func (a *StorageGatewayActivities) DetachVolume(input *storagegateway.DetachVolumeInput) (*storagegateway.DetachVolumeOutput, error) {
    return a.client.DetachVolume(input)
}

func (a *StorageGatewayActivities) DisableGateway(input *storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error) {
    return a.client.DisableGateway(input)
}

func (a *StorageGatewayActivities) JoinDomain(input *storagegateway.JoinDomainInput) (*storagegateway.JoinDomainOutput, error) {
    return a.client.JoinDomain(input)
}

func (a *StorageGatewayActivities) ListAutomaticTapeCreationPolicies(input *storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
    return a.client.ListAutomaticTapeCreationPolicies(input)
}

func (a *StorageGatewayActivities) ListFileShares(input *storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error) {
    return a.client.ListFileShares(input)
}

func (a *StorageGatewayActivities) ListGateways(input *storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error) {
    return a.client.ListGateways(input)
}

func (a *StorageGatewayActivities) ListLocalDisks(input *storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error) {
    return a.client.ListLocalDisks(input)
}

func (a *StorageGatewayActivities) ListTagsForResource(input *storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *StorageGatewayActivities) ListTapePools(input *storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error) {
    return a.client.ListTapePools(input)
}

func (a *StorageGatewayActivities) ListTapes(input *storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error) {
    return a.client.ListTapes(input)
}

func (a *StorageGatewayActivities) ListVolumeInitiators(input *storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error) {
    return a.client.ListVolumeInitiators(input)
}

func (a *StorageGatewayActivities) ListVolumeRecoveryPoints(input *storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
    return a.client.ListVolumeRecoveryPoints(input)
}

func (a *StorageGatewayActivities) ListVolumes(input *storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error) {
    return a.client.ListVolumes(input)
}

func (a *StorageGatewayActivities) NotifyWhenUploaded(input *storagegateway.NotifyWhenUploadedInput) (*storagegateway.NotifyWhenUploadedOutput, error) {
    return a.client.NotifyWhenUploaded(input)
}

func (a *StorageGatewayActivities) RefreshCache(input *storagegateway.RefreshCacheInput) (*storagegateway.RefreshCacheOutput, error) {
    return a.client.RefreshCache(input)
}

func (a *StorageGatewayActivities) RemoveTagsFromResource(input *storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error) {
    return a.client.RemoveTagsFromResource(input)
}

func (a *StorageGatewayActivities) ResetCache(input *storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error) {
    return a.client.ResetCache(input)
}

func (a *StorageGatewayActivities) RetrieveTapeArchive(input *storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error) {
    return a.client.RetrieveTapeArchive(input)
}

func (a *StorageGatewayActivities) RetrieveTapeRecoveryPoint(input *storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error) {
    return a.client.RetrieveTapeRecoveryPoint(input)
}

func (a *StorageGatewayActivities) SetLocalConsolePassword(input *storagegateway.SetLocalConsolePasswordInput) (*storagegateway.SetLocalConsolePasswordOutput, error) {
    return a.client.SetLocalConsolePassword(input)
}

func (a *StorageGatewayActivities) SetSMBGuestPassword(input *storagegateway.SetSMBGuestPasswordInput) (*storagegateway.SetSMBGuestPasswordOutput, error) {
    return a.client.SetSMBGuestPassword(input)
}

func (a *StorageGatewayActivities) ShutdownGateway(input *storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error) {
    return a.client.ShutdownGateway(input)
}

func (a *StorageGatewayActivities) StartAvailabilityMonitorTest(input *storagegateway.StartAvailabilityMonitorTestInput) (*storagegateway.StartAvailabilityMonitorTestOutput, error) {
    return a.client.StartAvailabilityMonitorTest(input)
}

func (a *StorageGatewayActivities) StartGateway(input *storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error) {
    return a.client.StartGateway(input)
}

func (a *StorageGatewayActivities) UpdateAutomaticTapeCreationPolicy(input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error) {
    return a.client.UpdateAutomaticTapeCreationPolicy(input)
}

func (a *StorageGatewayActivities) UpdateBandwidthRateLimit(input *storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error) {
    return a.client.UpdateBandwidthRateLimit(input)
}

func (a *StorageGatewayActivities) UpdateChapCredentials(input *storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error) {
    return a.client.UpdateChapCredentials(input)
}

func (a *StorageGatewayActivities) UpdateGatewayInformation(input *storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error) {
    return a.client.UpdateGatewayInformation(input)
}

func (a *StorageGatewayActivities) UpdateGatewaySoftwareNow(input *storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error) {
    return a.client.UpdateGatewaySoftwareNow(input)
}

func (a *StorageGatewayActivities) UpdateMaintenanceStartTime(input *storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error) {
    return a.client.UpdateMaintenanceStartTime(input)
}

func (a *StorageGatewayActivities) UpdateNFSFileShare(input *storagegateway.UpdateNFSFileShareInput) (*storagegateway.UpdateNFSFileShareOutput, error) {
    return a.client.UpdateNFSFileShare(input)
}

func (a *StorageGatewayActivities) UpdateSMBFileShare(input *storagegateway.UpdateSMBFileShareInput) (*storagegateway.UpdateSMBFileShareOutput, error) {
    return a.client.UpdateSMBFileShare(input)
}

func (a *StorageGatewayActivities) UpdateSMBSecurityStrategy(input *storagegateway.UpdateSMBSecurityStrategyInput) (*storagegateway.UpdateSMBSecurityStrategyOutput, error) {
    return a.client.UpdateSMBSecurityStrategy(input)
}

func (a *StorageGatewayActivities) UpdateSnapshotSchedule(input *storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error) {
    return a.client.UpdateSnapshotSchedule(input)
}

func (a *StorageGatewayActivities) UpdateVTLDeviceType(input *storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error) {
    return a.client.UpdateVTLDeviceType(input)
}
