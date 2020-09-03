package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
)

type LightsailActivities struct {
	client lightsailiface.LightsailAPI
}

func NewLightsailActivities(client lightsailiface.LightsailAPI) *LightsailActivities {
    return &LightsailActivities{client: client}
}

func (a *LightsailActivities) AllocateStaticIp(input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error) {
    return a.client.AllocateStaticIp(input)
}

func (a *LightsailActivities) AttachCertificateToDistribution(input *lightsail.AttachCertificateToDistributionInput) (*lightsail.AttachCertificateToDistributionOutput, error) {
    return a.client.AttachCertificateToDistribution(input)
}

func (a *LightsailActivities) AttachDisk(input *lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error) {
    return a.client.AttachDisk(input)
}

func (a *LightsailActivities) AttachInstancesToLoadBalancer(input *lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error) {
    return a.client.AttachInstancesToLoadBalancer(input)
}

func (a *LightsailActivities) AttachLoadBalancerTlsCertificate(input *lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error) {
    return a.client.AttachLoadBalancerTlsCertificate(input)
}

func (a *LightsailActivities) AttachStaticIp(input *lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error) {
    return a.client.AttachStaticIp(input)
}

func (a *LightsailActivities) CloseInstancePublicPorts(input *lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error) {
    return a.client.CloseInstancePublicPorts(input)
}

func (a *LightsailActivities) CopySnapshot(input *lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error) {
    return a.client.CopySnapshot(input)
}

func (a *LightsailActivities) CreateCertificate(input *lightsail.CreateCertificateInput) (*lightsail.CreateCertificateOutput, error) {
    return a.client.CreateCertificate(input)
}

func (a *LightsailActivities) CreateCloudFormationStack(input *lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error) {
    return a.client.CreateCloudFormationStack(input)
}

func (a *LightsailActivities) CreateContactMethod(input *lightsail.CreateContactMethodInput) (*lightsail.CreateContactMethodOutput, error) {
    return a.client.CreateContactMethod(input)
}

func (a *LightsailActivities) CreateDisk(input *lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error) {
    return a.client.CreateDisk(input)
}

func (a *LightsailActivities) CreateDiskFromSnapshot(input *lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error) {
    return a.client.CreateDiskFromSnapshot(input)
}

func (a *LightsailActivities) CreateDiskSnapshot(input *lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error) {
    return a.client.CreateDiskSnapshot(input)
}

func (a *LightsailActivities) CreateDistribution(input *lightsail.CreateDistributionInput) (*lightsail.CreateDistributionOutput, error) {
    return a.client.CreateDistribution(input)
}

func (a *LightsailActivities) CreateDomain(input *lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error) {
    return a.client.CreateDomain(input)
}

func (a *LightsailActivities) CreateDomainEntry(input *lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error) {
    return a.client.CreateDomainEntry(input)
}

func (a *LightsailActivities) CreateInstanceSnapshot(input *lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error) {
    return a.client.CreateInstanceSnapshot(input)
}

func (a *LightsailActivities) CreateInstances(input *lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error) {
    return a.client.CreateInstances(input)
}

func (a *LightsailActivities) CreateInstancesFromSnapshot(input *lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error) {
    return a.client.CreateInstancesFromSnapshot(input)
}

func (a *LightsailActivities) CreateKeyPair(input *lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error) {
    return a.client.CreateKeyPair(input)
}

func (a *LightsailActivities) CreateLoadBalancer(input *lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error) {
    return a.client.CreateLoadBalancer(input)
}

func (a *LightsailActivities) CreateLoadBalancerTlsCertificate(input *lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error) {
    return a.client.CreateLoadBalancerTlsCertificate(input)
}

func (a *LightsailActivities) CreateRelationalDatabase(input *lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error) {
    return a.client.CreateRelationalDatabase(input)
}

func (a *LightsailActivities) CreateRelationalDatabaseFromSnapshot(input *lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error) {
    return a.client.CreateRelationalDatabaseFromSnapshot(input)
}

func (a *LightsailActivities) CreateRelationalDatabaseSnapshot(input *lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error) {
    return a.client.CreateRelationalDatabaseSnapshot(input)
}

func (a *LightsailActivities) DeleteAlarm(input *lightsail.DeleteAlarmInput) (*lightsail.DeleteAlarmOutput, error) {
    return a.client.DeleteAlarm(input)
}

func (a *LightsailActivities) DeleteAutoSnapshot(input *lightsail.DeleteAutoSnapshotInput) (*lightsail.DeleteAutoSnapshotOutput, error) {
    return a.client.DeleteAutoSnapshot(input)
}

func (a *LightsailActivities) DeleteCertificate(input *lightsail.DeleteCertificateInput) (*lightsail.DeleteCertificateOutput, error) {
    return a.client.DeleteCertificate(input)
}

func (a *LightsailActivities) DeleteContactMethod(input *lightsail.DeleteContactMethodInput) (*lightsail.DeleteContactMethodOutput, error) {
    return a.client.DeleteContactMethod(input)
}

func (a *LightsailActivities) DeleteDisk(input *lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error) {
    return a.client.DeleteDisk(input)
}

func (a *LightsailActivities) DeleteDiskSnapshot(input *lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error) {
    return a.client.DeleteDiskSnapshot(input)
}

func (a *LightsailActivities) DeleteDistribution(input *lightsail.DeleteDistributionInput) (*lightsail.DeleteDistributionOutput, error) {
    return a.client.DeleteDistribution(input)
}

func (a *LightsailActivities) DeleteDomain(input *lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error) {
    return a.client.DeleteDomain(input)
}

func (a *LightsailActivities) DeleteDomainEntry(input *lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error) {
    return a.client.DeleteDomainEntry(input)
}

func (a *LightsailActivities) DeleteInstance(input *lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error) {
    return a.client.DeleteInstance(input)
}

func (a *LightsailActivities) DeleteInstanceSnapshot(input *lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error) {
    return a.client.DeleteInstanceSnapshot(input)
}

func (a *LightsailActivities) DeleteKeyPair(input *lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error) {
    return a.client.DeleteKeyPair(input)
}

func (a *LightsailActivities) DeleteKnownHostKeys(input *lightsail.DeleteKnownHostKeysInput) (*lightsail.DeleteKnownHostKeysOutput, error) {
    return a.client.DeleteKnownHostKeys(input)
}

func (a *LightsailActivities) DeleteLoadBalancer(input *lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error) {
    return a.client.DeleteLoadBalancer(input)
}

func (a *LightsailActivities) DeleteLoadBalancerTlsCertificate(input *lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error) {
    return a.client.DeleteLoadBalancerTlsCertificate(input)
}

func (a *LightsailActivities) DeleteRelationalDatabase(input *lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error) {
    return a.client.DeleteRelationalDatabase(input)
}

func (a *LightsailActivities) DeleteRelationalDatabaseSnapshot(input *lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error) {
    return a.client.DeleteRelationalDatabaseSnapshot(input)
}

func (a *LightsailActivities) DetachCertificateFromDistribution(input *lightsail.DetachCertificateFromDistributionInput) (*lightsail.DetachCertificateFromDistributionOutput, error) {
    return a.client.DetachCertificateFromDistribution(input)
}

func (a *LightsailActivities) DetachDisk(input *lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error) {
    return a.client.DetachDisk(input)
}

func (a *LightsailActivities) DetachInstancesFromLoadBalancer(input *lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error) {
    return a.client.DetachInstancesFromLoadBalancer(input)
}

func (a *LightsailActivities) DetachStaticIp(input *lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error) {
    return a.client.DetachStaticIp(input)
}

func (a *LightsailActivities) DisableAddOn(input *lightsail.DisableAddOnInput) (*lightsail.DisableAddOnOutput, error) {
    return a.client.DisableAddOn(input)
}

func (a *LightsailActivities) DownloadDefaultKeyPair(input *lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error) {
    return a.client.DownloadDefaultKeyPair(input)
}

func (a *LightsailActivities) EnableAddOn(input *lightsail.EnableAddOnInput) (*lightsail.EnableAddOnOutput, error) {
    return a.client.EnableAddOn(input)
}

func (a *LightsailActivities) ExportSnapshot(input *lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error) {
    return a.client.ExportSnapshot(input)
}

func (a *LightsailActivities) GetActiveNames(input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error) {
    return a.client.GetActiveNames(input)
}

func (a *LightsailActivities) GetAlarms(input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error) {
    return a.client.GetAlarms(input)
}

func (a *LightsailActivities) GetAutoSnapshots(input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error) {
    return a.client.GetAutoSnapshots(input)
}

func (a *LightsailActivities) GetBlueprints(input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error) {
    return a.client.GetBlueprints(input)
}

func (a *LightsailActivities) GetBundles(input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error) {
    return a.client.GetBundles(input)
}

func (a *LightsailActivities) GetCertificates(input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error) {
    return a.client.GetCertificates(input)
}

func (a *LightsailActivities) GetCloudFormationStackRecords(input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
    return a.client.GetCloudFormationStackRecords(input)
}

func (a *LightsailActivities) GetContactMethods(input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error) {
    return a.client.GetContactMethods(input)
}

func (a *LightsailActivities) GetDisk(input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error) {
    return a.client.GetDisk(input)
}

func (a *LightsailActivities) GetDiskSnapshot(input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error) {
    return a.client.GetDiskSnapshot(input)
}

func (a *LightsailActivities) GetDiskSnapshots(input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error) {
    return a.client.GetDiskSnapshots(input)
}

func (a *LightsailActivities) GetDisks(input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error) {
    return a.client.GetDisks(input)
}

func (a *LightsailActivities) GetDistributionBundles(input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error) {
    return a.client.GetDistributionBundles(input)
}

func (a *LightsailActivities) GetDistributionLatestCacheReset(input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
    return a.client.GetDistributionLatestCacheReset(input)
}

func (a *LightsailActivities) GetDistributionMetricData(input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error) {
    return a.client.GetDistributionMetricData(input)
}

func (a *LightsailActivities) GetDistributions(input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error) {
    return a.client.GetDistributions(input)
}

func (a *LightsailActivities) GetDomain(input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error) {
    return a.client.GetDomain(input)
}

func (a *LightsailActivities) GetDomains(input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error) {
    return a.client.GetDomains(input)
}

func (a *LightsailActivities) GetExportSnapshotRecords(input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error) {
    return a.client.GetExportSnapshotRecords(input)
}

func (a *LightsailActivities) GetInstance(input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error) {
    return a.client.GetInstance(input)
}

func (a *LightsailActivities) GetInstanceAccessDetails(input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error) {
    return a.client.GetInstanceAccessDetails(input)
}

func (a *LightsailActivities) GetInstanceMetricData(input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error) {
    return a.client.GetInstanceMetricData(input)
}

func (a *LightsailActivities) GetInstancePortStates(input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error) {
    return a.client.GetInstancePortStates(input)
}

func (a *LightsailActivities) GetInstanceSnapshot(input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error) {
    return a.client.GetInstanceSnapshot(input)
}

func (a *LightsailActivities) GetInstanceSnapshots(input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error) {
    return a.client.GetInstanceSnapshots(input)
}

func (a *LightsailActivities) GetInstanceState(input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error) {
    return a.client.GetInstanceState(input)
}

func (a *LightsailActivities) GetInstances(input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error) {
    return a.client.GetInstances(input)
}

func (a *LightsailActivities) GetKeyPair(input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error) {
    return a.client.GetKeyPair(input)
}

func (a *LightsailActivities) GetKeyPairs(input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error) {
    return a.client.GetKeyPairs(input)
}

func (a *LightsailActivities) GetLoadBalancer(input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error) {
    return a.client.GetLoadBalancer(input)
}

func (a *LightsailActivities) GetLoadBalancerMetricData(input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
    return a.client.GetLoadBalancerMetricData(input)
}

func (a *LightsailActivities) GetLoadBalancerTlsCertificates(input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
    return a.client.GetLoadBalancerTlsCertificates(input)
}

func (a *LightsailActivities) GetLoadBalancers(input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error) {
    return a.client.GetLoadBalancers(input)
}

func (a *LightsailActivities) GetOperation(input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error) {
    return a.client.GetOperation(input)
}

func (a *LightsailActivities) GetOperations(input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error) {
    return a.client.GetOperations(input)
}

func (a *LightsailActivities) GetOperationsForResource(input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error) {
    return a.client.GetOperationsForResource(input)
}

func (a *LightsailActivities) GetRegions(input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error) {
    return a.client.GetRegions(input)
}

func (a *LightsailActivities) GetRelationalDatabase(input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error) {
    return a.client.GetRelationalDatabase(input)
}

func (a *LightsailActivities) GetRelationalDatabaseBlueprints(input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
    return a.client.GetRelationalDatabaseBlueprints(input)
}

func (a *LightsailActivities) GetRelationalDatabaseBundles(input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
    return a.client.GetRelationalDatabaseBundles(input)
}

func (a *LightsailActivities) GetRelationalDatabaseEvents(input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
    return a.client.GetRelationalDatabaseEvents(input)
}

func (a *LightsailActivities) GetRelationalDatabaseLogEvents(input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
    return a.client.GetRelationalDatabaseLogEvents(input)
}

func (a *LightsailActivities) GetRelationalDatabaseLogStreams(input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
    return a.client.GetRelationalDatabaseLogStreams(input)
}

func (a *LightsailActivities) GetRelationalDatabaseMasterUserPassword(input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
    return a.client.GetRelationalDatabaseMasterUserPassword(input)
}

func (a *LightsailActivities) GetRelationalDatabaseMetricData(input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
    return a.client.GetRelationalDatabaseMetricData(input)
}

func (a *LightsailActivities) GetRelationalDatabaseParameters(input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
    return a.client.GetRelationalDatabaseParameters(input)
}

func (a *LightsailActivities) GetRelationalDatabaseSnapshot(input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
    return a.client.GetRelationalDatabaseSnapshot(input)
}

func (a *LightsailActivities) GetRelationalDatabaseSnapshots(input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
    return a.client.GetRelationalDatabaseSnapshots(input)
}

func (a *LightsailActivities) GetRelationalDatabases(input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error) {
    return a.client.GetRelationalDatabases(input)
}

func (a *LightsailActivities) GetStaticIp(input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error) {
    return a.client.GetStaticIp(input)
}

func (a *LightsailActivities) GetStaticIps(input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error) {
    return a.client.GetStaticIps(input)
}

func (a *LightsailActivities) ImportKeyPair(input *lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error) {
    return a.client.ImportKeyPair(input)
}

func (a *LightsailActivities) IsVpcPeered(input *lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error) {
    return a.client.IsVpcPeered(input)
}

func (a *LightsailActivities) OpenInstancePublicPorts(input *lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error) {
    return a.client.OpenInstancePublicPorts(input)
}

func (a *LightsailActivities) PeerVpc(input *lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error) {
    return a.client.PeerVpc(input)
}

func (a *LightsailActivities) PutAlarm(input *lightsail.PutAlarmInput) (*lightsail.PutAlarmOutput, error) {
    return a.client.PutAlarm(input)
}

func (a *LightsailActivities) PutInstancePublicPorts(input *lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error) {
    return a.client.PutInstancePublicPorts(input)
}

func (a *LightsailActivities) RebootInstance(input *lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error) {
    return a.client.RebootInstance(input)
}

func (a *LightsailActivities) RebootRelationalDatabase(input *lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error) {
    return a.client.RebootRelationalDatabase(input)
}

func (a *LightsailActivities) ReleaseStaticIp(input *lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error) {
    return a.client.ReleaseStaticIp(input)
}

func (a *LightsailActivities) ResetDistributionCache(input *lightsail.ResetDistributionCacheInput) (*lightsail.ResetDistributionCacheOutput, error) {
    return a.client.ResetDistributionCache(input)
}

func (a *LightsailActivities) SendContactMethodVerification(input *lightsail.SendContactMethodVerificationInput) (*lightsail.SendContactMethodVerificationOutput, error) {
    return a.client.SendContactMethodVerification(input)
}

func (a *LightsailActivities) StartInstance(input *lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error) {
    return a.client.StartInstance(input)
}

func (a *LightsailActivities) StartRelationalDatabase(input *lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error) {
    return a.client.StartRelationalDatabase(input)
}

func (a *LightsailActivities) StopInstance(input *lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error) {
    return a.client.StopInstance(input)
}

func (a *LightsailActivities) StopRelationalDatabase(input *lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error) {
    return a.client.StopRelationalDatabase(input)
}

func (a *LightsailActivities) TagResource(input *lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *LightsailActivities) TestAlarm(input *lightsail.TestAlarmInput) (*lightsail.TestAlarmOutput, error) {
    return a.client.TestAlarm(input)
}

func (a *LightsailActivities) UnpeerVpc(input *lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error) {
    return a.client.UnpeerVpc(input)
}

func (a *LightsailActivities) UntagResource(input *lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *LightsailActivities) UpdateDistribution(input *lightsail.UpdateDistributionInput) (*lightsail.UpdateDistributionOutput, error) {
    return a.client.UpdateDistribution(input)
}

func (a *LightsailActivities) UpdateDistributionBundle(input *lightsail.UpdateDistributionBundleInput) (*lightsail.UpdateDistributionBundleOutput, error) {
    return a.client.UpdateDistributionBundle(input)
}

func (a *LightsailActivities) UpdateDomainEntry(input *lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error) {
    return a.client.UpdateDomainEntry(input)
}

func (a *LightsailActivities) UpdateLoadBalancerAttribute(input *lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error) {
    return a.client.UpdateLoadBalancerAttribute(input)
}

func (a *LightsailActivities) UpdateRelationalDatabase(input *lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error) {
    return a.client.UpdateRelationalDatabase(input)
}

func (a *LightsailActivities) UpdateRelationalDatabaseParameters(input *lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error) {
    return a.client.UpdateRelationalDatabaseParameters(input)
}
