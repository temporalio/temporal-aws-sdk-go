package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type LightsailActivities struct {
	client lightsailiface.LightsailAPI
}

func NewLightsailActivities(session *session.Session, config ...*aws.Config) *LightsailActivities {
	client := lightsail.New(session, config...)
	return &LightsailActivities{client: client}
}

func (a *LightsailActivities) AllocateStaticIp(ctx context.Context, input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error) {
	return a.client.AllocateStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) AttachCertificateToDistribution(ctx context.Context, input *lightsail.AttachCertificateToDistributionInput) (*lightsail.AttachCertificateToDistributionOutput, error) {
	return a.client.AttachCertificateToDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) AttachDisk(ctx context.Context, input *lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error) {
	return a.client.AttachDiskWithContext(ctx, input)
}

func (a *LightsailActivities) AttachInstancesToLoadBalancer(ctx context.Context, input *lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error) {
	return a.client.AttachInstancesToLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) AttachLoadBalancerTlsCertificate(ctx context.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error) {
	return a.client.AttachLoadBalancerTlsCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) AttachStaticIp(ctx context.Context, input *lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error) {
	return a.client.AttachStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) CloseInstancePublicPorts(ctx context.Context, input *lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error) {
	return a.client.CloseInstancePublicPortsWithContext(ctx, input)
}

func (a *LightsailActivities) CopySnapshot(ctx context.Context, input *lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error) {
	return a.client.CopySnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateCertificate(ctx context.Context, input *lightsail.CreateCertificateInput) (*lightsail.CreateCertificateOutput, error) {
	return a.client.CreateCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) CreateCloudFormationStack(ctx context.Context, input *lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error) {
	return a.client.CreateCloudFormationStackWithContext(ctx, input)
}

func (a *LightsailActivities) CreateContactMethod(ctx context.Context, input *lightsail.CreateContactMethodInput) (*lightsail.CreateContactMethodOutput, error) {
	return a.client.CreateContactMethodWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDisk(ctx context.Context, input *lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error) {
	return a.client.CreateDiskWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDiskFromSnapshot(ctx context.Context, input *lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error) {
	return a.client.CreateDiskFromSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDiskSnapshot(ctx context.Context, input *lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error) {
	return a.client.CreateDiskSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDistribution(ctx context.Context, input *lightsail.CreateDistributionInput) (*lightsail.CreateDistributionOutput, error) {
	return a.client.CreateDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDomain(ctx context.Context, input *lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error) {
	return a.client.CreateDomainWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDomainEntry(ctx context.Context, input *lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error) {
	return a.client.CreateDomainEntryWithContext(ctx, input)
}

func (a *LightsailActivities) CreateInstanceSnapshot(ctx context.Context, input *lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error) {
	return a.client.CreateInstanceSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateInstances(ctx context.Context, input *lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error) {
	return a.client.CreateInstancesWithContext(ctx, input)
}

func (a *LightsailActivities) CreateInstancesFromSnapshot(ctx context.Context, input *lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error) {
	return a.client.CreateInstancesFromSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateKeyPair(ctx context.Context, input *lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error) {
	return a.client.CreateKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) CreateLoadBalancer(ctx context.Context, input *lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error) {
	return a.client.CreateLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) CreateLoadBalancerTlsCertificate(ctx context.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error) {
	return a.client.CreateLoadBalancerTlsCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) CreateRelationalDatabase(ctx context.Context, input *lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error) {
	return a.client.CreateRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) CreateRelationalDatabaseFromSnapshot(ctx context.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error) {
	return a.client.CreateRelationalDatabaseFromSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateRelationalDatabaseSnapshot(ctx context.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error) {
	return a.client.CreateRelationalDatabaseSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteAlarm(ctx context.Context, input *lightsail.DeleteAlarmInput) (*lightsail.DeleteAlarmOutput, error) {
	return a.client.DeleteAlarmWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteAutoSnapshot(ctx context.Context, input *lightsail.DeleteAutoSnapshotInput) (*lightsail.DeleteAutoSnapshotOutput, error) {
	return a.client.DeleteAutoSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteCertificate(ctx context.Context, input *lightsail.DeleteCertificateInput) (*lightsail.DeleteCertificateOutput, error) {
	return a.client.DeleteCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteContactMethod(ctx context.Context, input *lightsail.DeleteContactMethodInput) (*lightsail.DeleteContactMethodOutput, error) {
	return a.client.DeleteContactMethodWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDisk(ctx context.Context, input *lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error) {
	return a.client.DeleteDiskWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDiskSnapshot(ctx context.Context, input *lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error) {
	return a.client.DeleteDiskSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDistribution(ctx context.Context, input *lightsail.DeleteDistributionInput) (*lightsail.DeleteDistributionOutput, error) {
	return a.client.DeleteDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDomain(ctx context.Context, input *lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error) {
	return a.client.DeleteDomainWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDomainEntry(ctx context.Context, input *lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error) {
	return a.client.DeleteDomainEntryWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteInstance(ctx context.Context, input *lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error) {
	return a.client.DeleteInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteInstanceSnapshot(ctx context.Context, input *lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error) {
	return a.client.DeleteInstanceSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteKeyPair(ctx context.Context, input *lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error) {
	return a.client.DeleteKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteKnownHostKeys(ctx context.Context, input *lightsail.DeleteKnownHostKeysInput) (*lightsail.DeleteKnownHostKeysOutput, error) {
	return a.client.DeleteKnownHostKeysWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteLoadBalancer(ctx context.Context, input *lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error) {
	return a.client.DeleteLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteLoadBalancerTlsCertificate(ctx context.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error) {
	return a.client.DeleteLoadBalancerTlsCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteRelationalDatabase(ctx context.Context, input *lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error) {
	return a.client.DeleteRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteRelationalDatabaseSnapshot(ctx context.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error) {
	return a.client.DeleteRelationalDatabaseSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DetachCertificateFromDistribution(ctx context.Context, input *lightsail.DetachCertificateFromDistributionInput) (*lightsail.DetachCertificateFromDistributionOutput, error) {
	return a.client.DetachCertificateFromDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) DetachDisk(ctx context.Context, input *lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error) {
	return a.client.DetachDiskWithContext(ctx, input)
}

func (a *LightsailActivities) DetachInstancesFromLoadBalancer(ctx context.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error) {
	return a.client.DetachInstancesFromLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) DetachStaticIp(ctx context.Context, input *lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error) {
	return a.client.DetachStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) DisableAddOn(ctx context.Context, input *lightsail.DisableAddOnInput) (*lightsail.DisableAddOnOutput, error) {
	return a.client.DisableAddOnWithContext(ctx, input)
}

func (a *LightsailActivities) DownloadDefaultKeyPair(ctx context.Context, input *lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error) {
	return a.client.DownloadDefaultKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) EnableAddOn(ctx context.Context, input *lightsail.EnableAddOnInput) (*lightsail.EnableAddOnOutput, error) {
	return a.client.EnableAddOnWithContext(ctx, input)
}

func (a *LightsailActivities) ExportSnapshot(ctx context.Context, input *lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error) {
	return a.client.ExportSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetActiveNames(ctx context.Context, input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error) {
	return a.client.GetActiveNamesWithContext(ctx, input)
}

func (a *LightsailActivities) GetAlarms(ctx context.Context, input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error) {
	return a.client.GetAlarmsWithContext(ctx, input)
}

func (a *LightsailActivities) GetAutoSnapshots(ctx context.Context, input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error) {
	return a.client.GetAutoSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetBlueprints(ctx context.Context, input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error) {
	return a.client.GetBlueprintsWithContext(ctx, input)
}

func (a *LightsailActivities) GetBundles(ctx context.Context, input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error) {
	return a.client.GetBundlesWithContext(ctx, input)
}

func (a *LightsailActivities) GetCertificates(ctx context.Context, input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error) {
	return a.client.GetCertificatesWithContext(ctx, input)
}

func (a *LightsailActivities) GetCloudFormationStackRecords(ctx context.Context, input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	return a.client.GetCloudFormationStackRecordsWithContext(ctx, input)
}

func (a *LightsailActivities) GetContactMethods(ctx context.Context, input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error) {
	return a.client.GetContactMethodsWithContext(ctx, input)
}

func (a *LightsailActivities) GetDisk(ctx context.Context, input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error) {
	return a.client.GetDiskWithContext(ctx, input)
}

func (a *LightsailActivities) GetDiskSnapshot(ctx context.Context, input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error) {
	return a.client.GetDiskSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetDiskSnapshots(ctx context.Context, input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error) {
	return a.client.GetDiskSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetDisks(ctx context.Context, input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error) {
	return a.client.GetDisksWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributionBundles(ctx context.Context, input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error) {
	return a.client.GetDistributionBundlesWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributionLatestCacheReset(ctx context.Context, input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	return a.client.GetDistributionLatestCacheResetWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributionMetricData(ctx context.Context, input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error) {
	return a.client.GetDistributionMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributions(ctx context.Context, input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error) {
	return a.client.GetDistributionsWithContext(ctx, input)
}

func (a *LightsailActivities) GetDomain(ctx context.Context, input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error) {
	return a.client.GetDomainWithContext(ctx, input)
}

func (a *LightsailActivities) GetDomains(ctx context.Context, input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error) {
	return a.client.GetDomainsWithContext(ctx, input)
}

func (a *LightsailActivities) GetExportSnapshotRecords(ctx context.Context, input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	return a.client.GetExportSnapshotRecordsWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstance(ctx context.Context, input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error) {
	return a.client.GetInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceAccessDetails(ctx context.Context, input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	return a.client.GetInstanceAccessDetailsWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceMetricData(ctx context.Context, input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error) {
	return a.client.GetInstanceMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstancePortStates(ctx context.Context, input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error) {
	return a.client.GetInstancePortStatesWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceSnapshot(ctx context.Context, input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error) {
	return a.client.GetInstanceSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceSnapshots(ctx context.Context, input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error) {
	return a.client.GetInstanceSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceState(ctx context.Context, input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error) {
	return a.client.GetInstanceStateWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstances(ctx context.Context, input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error) {
	return a.client.GetInstancesWithContext(ctx, input)
}

func (a *LightsailActivities) GetKeyPair(ctx context.Context, input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error) {
	return a.client.GetKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) GetKeyPairs(ctx context.Context, input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error) {
	return a.client.GetKeyPairsWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancer(ctx context.Context, input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error) {
	return a.client.GetLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancerMetricData(ctx context.Context, input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	return a.client.GetLoadBalancerMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancerTlsCertificates(ctx context.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	return a.client.GetLoadBalancerTlsCertificatesWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancers(ctx context.Context, input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error) {
	return a.client.GetLoadBalancersWithContext(ctx, input)
}

func (a *LightsailActivities) GetOperation(ctx context.Context, input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error) {
	return a.client.GetOperationWithContext(ctx, input)
}

func (a *LightsailActivities) GetOperations(ctx context.Context, input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error) {
	return a.client.GetOperationsWithContext(ctx, input)
}

func (a *LightsailActivities) GetOperationsForResource(ctx context.Context, input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error) {
	return a.client.GetOperationsForResourceWithContext(ctx, input)
}

func (a *LightsailActivities) GetRegions(ctx context.Context, input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error) {
	return a.client.GetRegionsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabase(ctx context.Context, input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error) {
	return a.client.GetRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseBlueprints(ctx context.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	return a.client.GetRelationalDatabaseBlueprintsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseBundles(ctx context.Context, input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	return a.client.GetRelationalDatabaseBundlesWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseEvents(ctx context.Context, input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	return a.client.GetRelationalDatabaseEventsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseLogEvents(ctx context.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	return a.client.GetRelationalDatabaseLogEventsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseLogStreams(ctx context.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	return a.client.GetRelationalDatabaseLogStreamsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseMasterUserPassword(ctx context.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	return a.client.GetRelationalDatabaseMasterUserPasswordWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseMetricData(ctx context.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	return a.client.GetRelationalDatabaseMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseParameters(ctx context.Context, input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	return a.client.GetRelationalDatabaseParametersWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseSnapshot(ctx context.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	return a.client.GetRelationalDatabaseSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseSnapshots(ctx context.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	return a.client.GetRelationalDatabaseSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabases(ctx context.Context, input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error) {
	return a.client.GetRelationalDatabasesWithContext(ctx, input)
}

func (a *LightsailActivities) GetStaticIp(ctx context.Context, input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error) {
	return a.client.GetStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) GetStaticIps(ctx context.Context, input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error) {
	return a.client.GetStaticIpsWithContext(ctx, input)
}

func (a *LightsailActivities) ImportKeyPair(ctx context.Context, input *lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error) {
	return a.client.ImportKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) IsVpcPeered(ctx context.Context, input *lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error) {
	return a.client.IsVpcPeeredWithContext(ctx, input)
}

func (a *LightsailActivities) OpenInstancePublicPorts(ctx context.Context, input *lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error) {
	return a.client.OpenInstancePublicPortsWithContext(ctx, input)
}

func (a *LightsailActivities) PeerVpc(ctx context.Context, input *lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error) {
	return a.client.PeerVpcWithContext(ctx, input)
}

func (a *LightsailActivities) PutAlarm(ctx context.Context, input *lightsail.PutAlarmInput) (*lightsail.PutAlarmOutput, error) {
	return a.client.PutAlarmWithContext(ctx, input)
}

func (a *LightsailActivities) PutInstancePublicPorts(ctx context.Context, input *lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error) {
	return a.client.PutInstancePublicPortsWithContext(ctx, input)
}

func (a *LightsailActivities) RebootInstance(ctx context.Context, input *lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error) {
	return a.client.RebootInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) RebootRelationalDatabase(ctx context.Context, input *lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error) {
	return a.client.RebootRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) ReleaseStaticIp(ctx context.Context, input *lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error) {
	return a.client.ReleaseStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) ResetDistributionCache(ctx context.Context, input *lightsail.ResetDistributionCacheInput) (*lightsail.ResetDistributionCacheOutput, error) {
	return a.client.ResetDistributionCacheWithContext(ctx, input)
}

func (a *LightsailActivities) SendContactMethodVerification(ctx context.Context, input *lightsail.SendContactMethodVerificationInput) (*lightsail.SendContactMethodVerificationOutput, error) {
	return a.client.SendContactMethodVerificationWithContext(ctx, input)
}

func (a *LightsailActivities) StartInstance(ctx context.Context, input *lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error) {
	return a.client.StartInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) StartRelationalDatabase(ctx context.Context, input *lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error) {
	return a.client.StartRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) StopInstance(ctx context.Context, input *lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error) {
	return a.client.StopInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) StopRelationalDatabase(ctx context.Context, input *lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error) {
	return a.client.StopRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) TagResource(ctx context.Context, input *lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *LightsailActivities) TestAlarm(ctx context.Context, input *lightsail.TestAlarmInput) (*lightsail.TestAlarmOutput, error) {
	return a.client.TestAlarmWithContext(ctx, input)
}

func (a *LightsailActivities) UnpeerVpc(ctx context.Context, input *lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error) {
	return a.client.UnpeerVpcWithContext(ctx, input)
}

func (a *LightsailActivities) UntagResource(ctx context.Context, input *lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateDistribution(ctx context.Context, input *lightsail.UpdateDistributionInput) (*lightsail.UpdateDistributionOutput, error) {
	return a.client.UpdateDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateDistributionBundle(ctx context.Context, input *lightsail.UpdateDistributionBundleInput) (*lightsail.UpdateDistributionBundleOutput, error) {
	return a.client.UpdateDistributionBundleWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateDomainEntry(ctx context.Context, input *lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error) {
	return a.client.UpdateDomainEntryWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateLoadBalancerAttribute(ctx context.Context, input *lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error) {
	return a.client.UpdateLoadBalancerAttributeWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateRelationalDatabase(ctx context.Context, input *lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error) {
	return a.client.UpdateRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateRelationalDatabaseParameters(ctx context.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error) {
	return a.client.UpdateRelationalDatabaseParametersWithContext(ctx, input)
}
