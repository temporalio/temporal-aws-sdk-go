package awsclients

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type LightsailClient interface {
	AllocateStaticIp(ctx workflow.Context, input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error)
	AllocateStaticIpAsync(ctx workflow.Context, input *lightsail.AllocateStaticIpInput) *LightsailAllocateStaticIpResult

	AttachCertificateToDistribution(ctx workflow.Context, input *lightsail.AttachCertificateToDistributionInput) (*lightsail.AttachCertificateToDistributionOutput, error)
	AttachCertificateToDistributionAsync(ctx workflow.Context, input *lightsail.AttachCertificateToDistributionInput) *LightsailAttachCertificateToDistributionResult

	AttachDisk(ctx workflow.Context, input *lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error)
	AttachDiskAsync(ctx workflow.Context, input *lightsail.AttachDiskInput) *LightsailAttachDiskResult

	AttachInstancesToLoadBalancer(ctx workflow.Context, input *lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error)
	AttachInstancesToLoadBalancerAsync(ctx workflow.Context, input *lightsail.AttachInstancesToLoadBalancerInput) *LightsailAttachInstancesToLoadBalancerResult

	AttachLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error)
	AttachLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) *LightsailAttachLoadBalancerTlsCertificateResult

	AttachStaticIp(ctx workflow.Context, input *lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error)
	AttachStaticIpAsync(ctx workflow.Context, input *lightsail.AttachStaticIpInput) *LightsailAttachStaticIpResult

	CloseInstancePublicPorts(ctx workflow.Context, input *lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error)
	CloseInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.CloseInstancePublicPortsInput) *LightsailCloseInstancePublicPortsResult

	CopySnapshot(ctx workflow.Context, input *lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error)
	CopySnapshotAsync(ctx workflow.Context, input *lightsail.CopySnapshotInput) *LightsailCopySnapshotResult

	CreateCertificate(ctx workflow.Context, input *lightsail.CreateCertificateInput) (*lightsail.CreateCertificateOutput, error)
	CreateCertificateAsync(ctx workflow.Context, input *lightsail.CreateCertificateInput) *LightsailCreateCertificateResult

	CreateCloudFormationStack(ctx workflow.Context, input *lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error)
	CreateCloudFormationStackAsync(ctx workflow.Context, input *lightsail.CreateCloudFormationStackInput) *LightsailCreateCloudFormationStackResult

	CreateContactMethod(ctx workflow.Context, input *lightsail.CreateContactMethodInput) (*lightsail.CreateContactMethodOutput, error)
	CreateContactMethodAsync(ctx workflow.Context, input *lightsail.CreateContactMethodInput) *LightsailCreateContactMethodResult

	CreateDisk(ctx workflow.Context, input *lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error)
	CreateDiskAsync(ctx workflow.Context, input *lightsail.CreateDiskInput) *LightsailCreateDiskResult

	CreateDiskFromSnapshot(ctx workflow.Context, input *lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error)
	CreateDiskFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateDiskFromSnapshotInput) *LightsailCreateDiskFromSnapshotResult

	CreateDiskSnapshot(ctx workflow.Context, input *lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error)
	CreateDiskSnapshotAsync(ctx workflow.Context, input *lightsail.CreateDiskSnapshotInput) *LightsailCreateDiskSnapshotResult

	CreateDistribution(ctx workflow.Context, input *lightsail.CreateDistributionInput) (*lightsail.CreateDistributionOutput, error)
	CreateDistributionAsync(ctx workflow.Context, input *lightsail.CreateDistributionInput) *LightsailCreateDistributionResult

	CreateDomain(ctx workflow.Context, input *lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error)
	CreateDomainAsync(ctx workflow.Context, input *lightsail.CreateDomainInput) *LightsailCreateDomainResult

	CreateDomainEntry(ctx workflow.Context, input *lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error)
	CreateDomainEntryAsync(ctx workflow.Context, input *lightsail.CreateDomainEntryInput) *LightsailCreateDomainEntryResult

	CreateInstanceSnapshot(ctx workflow.Context, input *lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error)
	CreateInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.CreateInstanceSnapshotInput) *LightsailCreateInstanceSnapshotResult

	CreateInstances(ctx workflow.Context, input *lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error)
	CreateInstancesAsync(ctx workflow.Context, input *lightsail.CreateInstancesInput) *LightsailCreateInstancesResult

	CreateInstancesFromSnapshot(ctx workflow.Context, input *lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error)
	CreateInstancesFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateInstancesFromSnapshotInput) *LightsailCreateInstancesFromSnapshotResult

	CreateKeyPair(ctx workflow.Context, input *lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error)
	CreateKeyPairAsync(ctx workflow.Context, input *lightsail.CreateKeyPairInput) *LightsailCreateKeyPairResult

	CreateLoadBalancer(ctx workflow.Context, input *lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *lightsail.CreateLoadBalancerInput) *LightsailCreateLoadBalancerResult

	CreateLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error)
	CreateLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) *LightsailCreateLoadBalancerTlsCertificateResult

	CreateRelationalDatabase(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error)
	CreateRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseInput) *LightsailCreateRelationalDatabaseResult

	CreateRelationalDatabaseFromSnapshot(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error)
	CreateRelationalDatabaseFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) *LightsailCreateRelationalDatabaseFromSnapshotResult

	CreateRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error)
	CreateRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) *LightsailCreateRelationalDatabaseSnapshotResult

	DeleteAlarm(ctx workflow.Context, input *lightsail.DeleteAlarmInput) (*lightsail.DeleteAlarmOutput, error)
	DeleteAlarmAsync(ctx workflow.Context, input *lightsail.DeleteAlarmInput) *LightsailDeleteAlarmResult

	DeleteAutoSnapshot(ctx workflow.Context, input *lightsail.DeleteAutoSnapshotInput) (*lightsail.DeleteAutoSnapshotOutput, error)
	DeleteAutoSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteAutoSnapshotInput) *LightsailDeleteAutoSnapshotResult

	DeleteCertificate(ctx workflow.Context, input *lightsail.DeleteCertificateInput) (*lightsail.DeleteCertificateOutput, error)
	DeleteCertificateAsync(ctx workflow.Context, input *lightsail.DeleteCertificateInput) *LightsailDeleteCertificateResult

	DeleteContactMethod(ctx workflow.Context, input *lightsail.DeleteContactMethodInput) (*lightsail.DeleteContactMethodOutput, error)
	DeleteContactMethodAsync(ctx workflow.Context, input *lightsail.DeleteContactMethodInput) *LightsailDeleteContactMethodResult

	DeleteDisk(ctx workflow.Context, input *lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error)
	DeleteDiskAsync(ctx workflow.Context, input *lightsail.DeleteDiskInput) *LightsailDeleteDiskResult

	DeleteDiskSnapshot(ctx workflow.Context, input *lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error)
	DeleteDiskSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteDiskSnapshotInput) *LightsailDeleteDiskSnapshotResult

	DeleteDistribution(ctx workflow.Context, input *lightsail.DeleteDistributionInput) (*lightsail.DeleteDistributionOutput, error)
	DeleteDistributionAsync(ctx workflow.Context, input *lightsail.DeleteDistributionInput) *LightsailDeleteDistributionResult

	DeleteDomain(ctx workflow.Context, input *lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error)
	DeleteDomainAsync(ctx workflow.Context, input *lightsail.DeleteDomainInput) *LightsailDeleteDomainResult

	DeleteDomainEntry(ctx workflow.Context, input *lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error)
	DeleteDomainEntryAsync(ctx workflow.Context, input *lightsail.DeleteDomainEntryInput) *LightsailDeleteDomainEntryResult

	DeleteInstance(ctx workflow.Context, input *lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error)
	DeleteInstanceAsync(ctx workflow.Context, input *lightsail.DeleteInstanceInput) *LightsailDeleteInstanceResult

	DeleteInstanceSnapshot(ctx workflow.Context, input *lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error)
	DeleteInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteInstanceSnapshotInput) *LightsailDeleteInstanceSnapshotResult

	DeleteKeyPair(ctx workflow.Context, input *lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error)
	DeleteKeyPairAsync(ctx workflow.Context, input *lightsail.DeleteKeyPairInput) *LightsailDeleteKeyPairResult

	DeleteKnownHostKeys(ctx workflow.Context, input *lightsail.DeleteKnownHostKeysInput) (*lightsail.DeleteKnownHostKeysOutput, error)
	DeleteKnownHostKeysAsync(ctx workflow.Context, input *lightsail.DeleteKnownHostKeysInput) *LightsailDeleteKnownHostKeysResult

	DeleteLoadBalancer(ctx workflow.Context, input *lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *lightsail.DeleteLoadBalancerInput) *LightsailDeleteLoadBalancerResult

	DeleteLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error)
	DeleteLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) *LightsailDeleteLoadBalancerTlsCertificateResult

	DeleteRelationalDatabase(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error)
	DeleteRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseInput) *LightsailDeleteRelationalDatabaseResult

	DeleteRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error)
	DeleteRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) *LightsailDeleteRelationalDatabaseSnapshotResult

	DetachCertificateFromDistribution(ctx workflow.Context, input *lightsail.DetachCertificateFromDistributionInput) (*lightsail.DetachCertificateFromDistributionOutput, error)
	DetachCertificateFromDistributionAsync(ctx workflow.Context, input *lightsail.DetachCertificateFromDistributionInput) *LightsailDetachCertificateFromDistributionResult

	DetachDisk(ctx workflow.Context, input *lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error)
	DetachDiskAsync(ctx workflow.Context, input *lightsail.DetachDiskInput) *LightsailDetachDiskResult

	DetachInstancesFromLoadBalancer(ctx workflow.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error)
	DetachInstancesFromLoadBalancerAsync(ctx workflow.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) *LightsailDetachInstancesFromLoadBalancerResult

	DetachStaticIp(ctx workflow.Context, input *lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error)
	DetachStaticIpAsync(ctx workflow.Context, input *lightsail.DetachStaticIpInput) *LightsailDetachStaticIpResult

	DisableAddOn(ctx workflow.Context, input *lightsail.DisableAddOnInput) (*lightsail.DisableAddOnOutput, error)
	DisableAddOnAsync(ctx workflow.Context, input *lightsail.DisableAddOnInput) *LightsailDisableAddOnResult

	DownloadDefaultKeyPair(ctx workflow.Context, input *lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error)
	DownloadDefaultKeyPairAsync(ctx workflow.Context, input *lightsail.DownloadDefaultKeyPairInput) *LightsailDownloadDefaultKeyPairResult

	EnableAddOn(ctx workflow.Context, input *lightsail.EnableAddOnInput) (*lightsail.EnableAddOnOutput, error)
	EnableAddOnAsync(ctx workflow.Context, input *lightsail.EnableAddOnInput) *LightsailEnableAddOnResult

	ExportSnapshot(ctx workflow.Context, input *lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error)
	ExportSnapshotAsync(ctx workflow.Context, input *lightsail.ExportSnapshotInput) *LightsailExportSnapshotResult

	GetActiveNames(ctx workflow.Context, input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error)
	GetActiveNamesAsync(ctx workflow.Context, input *lightsail.GetActiveNamesInput) *LightsailGetActiveNamesResult

	GetAlarms(ctx workflow.Context, input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error)
	GetAlarmsAsync(ctx workflow.Context, input *lightsail.GetAlarmsInput) *LightsailGetAlarmsResult

	GetAutoSnapshots(ctx workflow.Context, input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error)
	GetAutoSnapshotsAsync(ctx workflow.Context, input *lightsail.GetAutoSnapshotsInput) *LightsailGetAutoSnapshotsResult

	GetBlueprints(ctx workflow.Context, input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error)
	GetBlueprintsAsync(ctx workflow.Context, input *lightsail.GetBlueprintsInput) *LightsailGetBlueprintsResult

	GetBundles(ctx workflow.Context, input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error)
	GetBundlesAsync(ctx workflow.Context, input *lightsail.GetBundlesInput) *LightsailGetBundlesResult

	GetCertificates(ctx workflow.Context, input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error)
	GetCertificatesAsync(ctx workflow.Context, input *lightsail.GetCertificatesInput) *LightsailGetCertificatesResult

	GetCloudFormationStackRecords(ctx workflow.Context, input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error)
	GetCloudFormationStackRecordsAsync(ctx workflow.Context, input *lightsail.GetCloudFormationStackRecordsInput) *LightsailGetCloudFormationStackRecordsResult

	GetContactMethods(ctx workflow.Context, input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error)
	GetContactMethodsAsync(ctx workflow.Context, input *lightsail.GetContactMethodsInput) *LightsailGetContactMethodsResult

	GetDisk(ctx workflow.Context, input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error)
	GetDiskAsync(ctx workflow.Context, input *lightsail.GetDiskInput) *LightsailGetDiskResult

	GetDiskSnapshot(ctx workflow.Context, input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error)
	GetDiskSnapshotAsync(ctx workflow.Context, input *lightsail.GetDiskSnapshotInput) *LightsailGetDiskSnapshotResult

	GetDiskSnapshots(ctx workflow.Context, input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDiskSnapshotsAsync(ctx workflow.Context, input *lightsail.GetDiskSnapshotsInput) *LightsailGetDiskSnapshotsResult

	GetDisks(ctx workflow.Context, input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error)
	GetDisksAsync(ctx workflow.Context, input *lightsail.GetDisksInput) *LightsailGetDisksResult

	GetDistributionBundles(ctx workflow.Context, input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error)
	GetDistributionBundlesAsync(ctx workflow.Context, input *lightsail.GetDistributionBundlesInput) *LightsailGetDistributionBundlesResult

	GetDistributionLatestCacheReset(ctx workflow.Context, input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error)
	GetDistributionLatestCacheResetAsync(ctx workflow.Context, input *lightsail.GetDistributionLatestCacheResetInput) *LightsailGetDistributionLatestCacheResetResult

	GetDistributionMetricData(ctx workflow.Context, input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error)
	GetDistributionMetricDataAsync(ctx workflow.Context, input *lightsail.GetDistributionMetricDataInput) *LightsailGetDistributionMetricDataResult

	GetDistributions(ctx workflow.Context, input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error)
	GetDistributionsAsync(ctx workflow.Context, input *lightsail.GetDistributionsInput) *LightsailGetDistributionsResult

	GetDomain(ctx workflow.Context, input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error)
	GetDomainAsync(ctx workflow.Context, input *lightsail.GetDomainInput) *LightsailGetDomainResult

	GetDomains(ctx workflow.Context, input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error)
	GetDomainsAsync(ctx workflow.Context, input *lightsail.GetDomainsInput) *LightsailGetDomainsResult

	GetExportSnapshotRecords(ctx workflow.Context, input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error)
	GetExportSnapshotRecordsAsync(ctx workflow.Context, input *lightsail.GetExportSnapshotRecordsInput) *LightsailGetExportSnapshotRecordsResult

	GetInstance(ctx workflow.Context, input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error)
	GetInstanceAsync(ctx workflow.Context, input *lightsail.GetInstanceInput) *LightsailGetInstanceResult

	GetInstanceAccessDetails(ctx workflow.Context, input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceAccessDetailsAsync(ctx workflow.Context, input *lightsail.GetInstanceAccessDetailsInput) *LightsailGetInstanceAccessDetailsResult

	GetInstanceMetricData(ctx workflow.Context, input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstanceMetricDataAsync(ctx workflow.Context, input *lightsail.GetInstanceMetricDataInput) *LightsailGetInstanceMetricDataResult

	GetInstancePortStates(ctx workflow.Context, input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstancePortStatesAsync(ctx workflow.Context, input *lightsail.GetInstancePortStatesInput) *LightsailGetInstancePortStatesResult

	GetInstanceSnapshot(ctx workflow.Context, input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.GetInstanceSnapshotInput) *LightsailGetInstanceSnapshotResult

	GetInstanceSnapshots(ctx workflow.Context, input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceSnapshotsAsync(ctx workflow.Context, input *lightsail.GetInstanceSnapshotsInput) *LightsailGetInstanceSnapshotsResult

	GetInstanceState(ctx workflow.Context, input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error)
	GetInstanceStateAsync(ctx workflow.Context, input *lightsail.GetInstanceStateInput) *LightsailGetInstanceStateResult

	GetInstances(ctx workflow.Context, input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error)
	GetInstancesAsync(ctx workflow.Context, input *lightsail.GetInstancesInput) *LightsailGetInstancesResult

	GetKeyPair(ctx workflow.Context, input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairAsync(ctx workflow.Context, input *lightsail.GetKeyPairInput) *LightsailGetKeyPairResult

	GetKeyPairs(ctx workflow.Context, input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error)
	GetKeyPairsAsync(ctx workflow.Context, input *lightsail.GetKeyPairsInput) *LightsailGetKeyPairsResult

	GetLoadBalancer(ctx workflow.Context, input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error)
	GetLoadBalancerAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerInput) *LightsailGetLoadBalancerResult

	GetLoadBalancerMetricData(ctx workflow.Context, input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error)
	GetLoadBalancerMetricDataAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerMetricDataInput) *LightsailGetLoadBalancerMetricDataResult

	GetLoadBalancerTlsCertificates(ctx workflow.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetLoadBalancerTlsCertificatesAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) *LightsailGetLoadBalancerTlsCertificatesResult

	GetLoadBalancers(ctx workflow.Context, input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error)
	GetLoadBalancersAsync(ctx workflow.Context, input *lightsail.GetLoadBalancersInput) *LightsailGetLoadBalancersResult

	GetOperation(ctx workflow.Context, input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error)
	GetOperationAsync(ctx workflow.Context, input *lightsail.GetOperationInput) *LightsailGetOperationResult

	GetOperations(ctx workflow.Context, input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error)
	GetOperationsAsync(ctx workflow.Context, input *lightsail.GetOperationsInput) *LightsailGetOperationsResult

	GetOperationsForResource(ctx workflow.Context, input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error)
	GetOperationsForResourceAsync(ctx workflow.Context, input *lightsail.GetOperationsForResourceInput) *LightsailGetOperationsForResourceResult

	GetRegions(ctx workflow.Context, input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error)
	GetRegionsAsync(ctx workflow.Context, input *lightsail.GetRegionsInput) *LightsailGetRegionsResult

	GetRelationalDatabase(ctx workflow.Context, input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error)
	GetRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseInput) *LightsailGetRelationalDatabaseResult

	GetRelationalDatabaseBlueprints(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
	GetRelationalDatabaseBlueprintsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) *LightsailGetRelationalDatabaseBlueprintsResult

	GetRelationalDatabaseBundles(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
	GetRelationalDatabaseBundlesAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBundlesInput) *LightsailGetRelationalDatabaseBundlesResult

	GetRelationalDatabaseEvents(ctx workflow.Context, input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseEventsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseEventsInput) *LightsailGetRelationalDatabaseEventsResult

	GetRelationalDatabaseLogEvents(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogEventsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) *LightsailGetRelationalDatabaseLogEventsResult

	GetRelationalDatabaseLogStreams(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseLogStreamsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) *LightsailGetRelationalDatabaseLogStreamsResult

	GetRelationalDatabaseMasterUserPassword(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
	GetRelationalDatabaseMasterUserPasswordAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) *LightsailGetRelationalDatabaseMasterUserPasswordResult

	GetRelationalDatabaseMetricData(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
	GetRelationalDatabaseMetricDataAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) *LightsailGetRelationalDatabaseMetricDataResult

	GetRelationalDatabaseParameters(ctx workflow.Context, input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabaseParametersAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseParametersInput) *LightsailGetRelationalDatabaseParametersResult

	GetRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
	GetRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) *LightsailGetRelationalDatabaseSnapshotResult

	GetRelationalDatabaseSnapshots(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetRelationalDatabaseSnapshotsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) *LightsailGetRelationalDatabaseSnapshotsResult

	GetRelationalDatabases(ctx workflow.Context, input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error)
	GetRelationalDatabasesAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabasesInput) *LightsailGetRelationalDatabasesResult

	GetStaticIp(ctx workflow.Context, input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error)
	GetStaticIpAsync(ctx workflow.Context, input *lightsail.GetStaticIpInput) *LightsailGetStaticIpResult

	GetStaticIps(ctx workflow.Context, input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error)
	GetStaticIpsAsync(ctx workflow.Context, input *lightsail.GetStaticIpsInput) *LightsailGetStaticIpsResult

	ImportKeyPair(ctx workflow.Context, input *lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error)
	ImportKeyPairAsync(ctx workflow.Context, input *lightsail.ImportKeyPairInput) *LightsailImportKeyPairResult

	IsVpcPeered(ctx workflow.Context, input *lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error)
	IsVpcPeeredAsync(ctx workflow.Context, input *lightsail.IsVpcPeeredInput) *LightsailIsVpcPeeredResult

	OpenInstancePublicPorts(ctx workflow.Context, input *lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error)
	OpenInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.OpenInstancePublicPortsInput) *LightsailOpenInstancePublicPortsResult

	PeerVpc(ctx workflow.Context, input *lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error)
	PeerVpcAsync(ctx workflow.Context, input *lightsail.PeerVpcInput) *LightsailPeerVpcResult

	PutAlarm(ctx workflow.Context, input *lightsail.PutAlarmInput) (*lightsail.PutAlarmOutput, error)
	PutAlarmAsync(ctx workflow.Context, input *lightsail.PutAlarmInput) *LightsailPutAlarmResult

	PutInstancePublicPorts(ctx workflow.Context, input *lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error)
	PutInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.PutInstancePublicPortsInput) *LightsailPutInstancePublicPortsResult

	RebootInstance(ctx workflow.Context, input *lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error)
	RebootInstanceAsync(ctx workflow.Context, input *lightsail.RebootInstanceInput) *LightsailRebootInstanceResult

	RebootRelationalDatabase(ctx workflow.Context, input *lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error)
	RebootRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.RebootRelationalDatabaseInput) *LightsailRebootRelationalDatabaseResult

	ReleaseStaticIp(ctx workflow.Context, input *lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error)
	ReleaseStaticIpAsync(ctx workflow.Context, input *lightsail.ReleaseStaticIpInput) *LightsailReleaseStaticIpResult

	ResetDistributionCache(ctx workflow.Context, input *lightsail.ResetDistributionCacheInput) (*lightsail.ResetDistributionCacheOutput, error)
	ResetDistributionCacheAsync(ctx workflow.Context, input *lightsail.ResetDistributionCacheInput) *LightsailResetDistributionCacheResult

	SendContactMethodVerification(ctx workflow.Context, input *lightsail.SendContactMethodVerificationInput) (*lightsail.SendContactMethodVerificationOutput, error)
	SendContactMethodVerificationAsync(ctx workflow.Context, input *lightsail.SendContactMethodVerificationInput) *LightsailSendContactMethodVerificationResult

	StartInstance(ctx workflow.Context, input *lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error)
	StartInstanceAsync(ctx workflow.Context, input *lightsail.StartInstanceInput) *LightsailStartInstanceResult

	StartRelationalDatabase(ctx workflow.Context, input *lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error)
	StartRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.StartRelationalDatabaseInput) *LightsailStartRelationalDatabaseResult

	StopInstance(ctx workflow.Context, input *lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error)
	StopInstanceAsync(ctx workflow.Context, input *lightsail.StopInstanceInput) *LightsailStopInstanceResult

	StopRelationalDatabase(ctx workflow.Context, input *lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error)
	StopRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.StopRelationalDatabaseInput) *LightsailStopRelationalDatabaseResult

	TagResource(ctx workflow.Context, input *lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *lightsail.TagResourceInput) *LightsailTagResourceResult

	TestAlarm(ctx workflow.Context, input *lightsail.TestAlarmInput) (*lightsail.TestAlarmOutput, error)
	TestAlarmAsync(ctx workflow.Context, input *lightsail.TestAlarmInput) *LightsailTestAlarmResult

	UnpeerVpc(ctx workflow.Context, input *lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error)
	UnpeerVpcAsync(ctx workflow.Context, input *lightsail.UnpeerVpcInput) *LightsailUnpeerVpcResult

	UntagResource(ctx workflow.Context, input *lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *lightsail.UntagResourceInput) *LightsailUntagResourceResult

	UpdateDistribution(ctx workflow.Context, input *lightsail.UpdateDistributionInput) (*lightsail.UpdateDistributionOutput, error)
	UpdateDistributionAsync(ctx workflow.Context, input *lightsail.UpdateDistributionInput) *LightsailUpdateDistributionResult

	UpdateDistributionBundle(ctx workflow.Context, input *lightsail.UpdateDistributionBundleInput) (*lightsail.UpdateDistributionBundleOutput, error)
	UpdateDistributionBundleAsync(ctx workflow.Context, input *lightsail.UpdateDistributionBundleInput) *LightsailUpdateDistributionBundleResult

	UpdateDomainEntry(ctx workflow.Context, input *lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error)
	UpdateDomainEntryAsync(ctx workflow.Context, input *lightsail.UpdateDomainEntryInput) *LightsailUpdateDomainEntryResult

	UpdateLoadBalancerAttribute(ctx workflow.Context, input *lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error)
	UpdateLoadBalancerAttributeAsync(ctx workflow.Context, input *lightsail.UpdateLoadBalancerAttributeInput) *LightsailUpdateLoadBalancerAttributeResult

	UpdateRelationalDatabase(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error)
	UpdateRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseInput) *LightsailUpdateRelationalDatabaseResult

	UpdateRelationalDatabaseParameters(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error)
	UpdateRelationalDatabaseParametersAsync(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) *LightsailUpdateRelationalDatabaseParametersResult
}

type LightsailAllocateStaticIpResult struct {
	Result workflow.Future
}

func (r *LightsailAllocateStaticIpResult) Get(ctx workflow.Context) (*lightsail.AllocateStaticIpOutput, error) {
	var output lightsail.AllocateStaticIpOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailAttachCertificateToDistributionResult struct {
	Result workflow.Future
}

func (r *LightsailAttachCertificateToDistributionResult) Get(ctx workflow.Context) (*lightsail.AttachCertificateToDistributionOutput, error) {
	var output lightsail.AttachCertificateToDistributionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailAttachDiskResult struct {
	Result workflow.Future
}

func (r *LightsailAttachDiskResult) Get(ctx workflow.Context) (*lightsail.AttachDiskOutput, error) {
	var output lightsail.AttachDiskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailAttachInstancesToLoadBalancerResult struct {
	Result workflow.Future
}

func (r *LightsailAttachInstancesToLoadBalancerResult) Get(ctx workflow.Context) (*lightsail.AttachInstancesToLoadBalancerOutput, error) {
	var output lightsail.AttachInstancesToLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailAttachLoadBalancerTlsCertificateResult struct {
	Result workflow.Future
}

func (r *LightsailAttachLoadBalancerTlsCertificateResult) Get(ctx workflow.Context) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error) {
	var output lightsail.AttachLoadBalancerTlsCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailAttachStaticIpResult struct {
	Result workflow.Future
}

func (r *LightsailAttachStaticIpResult) Get(ctx workflow.Context) (*lightsail.AttachStaticIpOutput, error) {
	var output lightsail.AttachStaticIpOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCloseInstancePublicPortsResult struct {
	Result workflow.Future
}

func (r *LightsailCloseInstancePublicPortsResult) Get(ctx workflow.Context) (*lightsail.CloseInstancePublicPortsOutput, error) {
	var output lightsail.CloseInstancePublicPortsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCopySnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCopySnapshotResult) Get(ctx workflow.Context) (*lightsail.CopySnapshotOutput, error) {
	var output lightsail.CopySnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateCertificateResult struct {
	Result workflow.Future
}

func (r *LightsailCreateCertificateResult) Get(ctx workflow.Context) (*lightsail.CreateCertificateOutput, error) {
	var output lightsail.CreateCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateCloudFormationStackResult struct {
	Result workflow.Future
}

func (r *LightsailCreateCloudFormationStackResult) Get(ctx workflow.Context) (*lightsail.CreateCloudFormationStackOutput, error) {
	var output lightsail.CreateCloudFormationStackOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateContactMethodResult struct {
	Result workflow.Future
}

func (r *LightsailCreateContactMethodResult) Get(ctx workflow.Context) (*lightsail.CreateContactMethodOutput, error) {
	var output lightsail.CreateContactMethodOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateDiskResult struct {
	Result workflow.Future
}

func (r *LightsailCreateDiskResult) Get(ctx workflow.Context) (*lightsail.CreateDiskOutput, error) {
	var output lightsail.CreateDiskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateDiskFromSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCreateDiskFromSnapshotResult) Get(ctx workflow.Context) (*lightsail.CreateDiskFromSnapshotOutput, error) {
	var output lightsail.CreateDiskFromSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateDiskSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCreateDiskSnapshotResult) Get(ctx workflow.Context) (*lightsail.CreateDiskSnapshotOutput, error) {
	var output lightsail.CreateDiskSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateDistributionResult struct {
	Result workflow.Future
}

func (r *LightsailCreateDistributionResult) Get(ctx workflow.Context) (*lightsail.CreateDistributionOutput, error) {
	var output lightsail.CreateDistributionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateDomainResult struct {
	Result workflow.Future
}

func (r *LightsailCreateDomainResult) Get(ctx workflow.Context) (*lightsail.CreateDomainOutput, error) {
	var output lightsail.CreateDomainOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateDomainEntryResult struct {
	Result workflow.Future
}

func (r *LightsailCreateDomainEntryResult) Get(ctx workflow.Context) (*lightsail.CreateDomainEntryOutput, error) {
	var output lightsail.CreateDomainEntryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateInstanceSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCreateInstanceSnapshotResult) Get(ctx workflow.Context) (*lightsail.CreateInstanceSnapshotOutput, error) {
	var output lightsail.CreateInstanceSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateInstancesResult struct {
	Result workflow.Future
}

func (r *LightsailCreateInstancesResult) Get(ctx workflow.Context) (*lightsail.CreateInstancesOutput, error) {
	var output lightsail.CreateInstancesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateInstancesFromSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCreateInstancesFromSnapshotResult) Get(ctx workflow.Context) (*lightsail.CreateInstancesFromSnapshotOutput, error) {
	var output lightsail.CreateInstancesFromSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateKeyPairResult struct {
	Result workflow.Future
}

func (r *LightsailCreateKeyPairResult) Get(ctx workflow.Context) (*lightsail.CreateKeyPairOutput, error) {
	var output lightsail.CreateKeyPairOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateLoadBalancerResult struct {
	Result workflow.Future
}

func (r *LightsailCreateLoadBalancerResult) Get(ctx workflow.Context) (*lightsail.CreateLoadBalancerOutput, error) {
	var output lightsail.CreateLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateLoadBalancerTlsCertificateResult struct {
	Result workflow.Future
}

func (r *LightsailCreateLoadBalancerTlsCertificateResult) Get(ctx workflow.Context) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error) {
	var output lightsail.CreateLoadBalancerTlsCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailCreateRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.CreateRelationalDatabaseOutput, error) {
	var output lightsail.CreateRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateRelationalDatabaseFromSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCreateRelationalDatabaseFromSnapshotResult) Get(ctx workflow.Context) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error) {
	var output lightsail.CreateRelationalDatabaseFromSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailCreateRelationalDatabaseSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailCreateRelationalDatabaseSnapshotResult) Get(ctx workflow.Context) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error) {
	var output lightsail.CreateRelationalDatabaseSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteAlarmResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteAlarmResult) Get(ctx workflow.Context) (*lightsail.DeleteAlarmOutput, error) {
	var output lightsail.DeleteAlarmOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteAutoSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteAutoSnapshotResult) Get(ctx workflow.Context) (*lightsail.DeleteAutoSnapshotOutput, error) {
	var output lightsail.DeleteAutoSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteCertificateResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteCertificateResult) Get(ctx workflow.Context) (*lightsail.DeleteCertificateOutput, error) {
	var output lightsail.DeleteCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteContactMethodResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteContactMethodResult) Get(ctx workflow.Context) (*lightsail.DeleteContactMethodOutput, error) {
	var output lightsail.DeleteContactMethodOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteDiskResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteDiskResult) Get(ctx workflow.Context) (*lightsail.DeleteDiskOutput, error) {
	var output lightsail.DeleteDiskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteDiskSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteDiskSnapshotResult) Get(ctx workflow.Context) (*lightsail.DeleteDiskSnapshotOutput, error) {
	var output lightsail.DeleteDiskSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteDistributionResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteDistributionResult) Get(ctx workflow.Context) (*lightsail.DeleteDistributionOutput, error) {
	var output lightsail.DeleteDistributionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteDomainResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteDomainResult) Get(ctx workflow.Context) (*lightsail.DeleteDomainOutput, error) {
	var output lightsail.DeleteDomainOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteDomainEntryResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteDomainEntryResult) Get(ctx workflow.Context) (*lightsail.DeleteDomainEntryOutput, error) {
	var output lightsail.DeleteDomainEntryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteInstanceResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteInstanceResult) Get(ctx workflow.Context) (*lightsail.DeleteInstanceOutput, error) {
	var output lightsail.DeleteInstanceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteInstanceSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteInstanceSnapshotResult) Get(ctx workflow.Context) (*lightsail.DeleteInstanceSnapshotOutput, error) {
	var output lightsail.DeleteInstanceSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteKeyPairResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteKeyPairResult) Get(ctx workflow.Context) (*lightsail.DeleteKeyPairOutput, error) {
	var output lightsail.DeleteKeyPairOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteKnownHostKeysResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteKnownHostKeysResult) Get(ctx workflow.Context) (*lightsail.DeleteKnownHostKeysOutput, error) {
	var output lightsail.DeleteKnownHostKeysOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteLoadBalancerResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteLoadBalancerResult) Get(ctx workflow.Context) (*lightsail.DeleteLoadBalancerOutput, error) {
	var output lightsail.DeleteLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteLoadBalancerTlsCertificateResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteLoadBalancerTlsCertificateResult) Get(ctx workflow.Context) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error) {
	var output lightsail.DeleteLoadBalancerTlsCertificateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.DeleteRelationalDatabaseOutput, error) {
	var output lightsail.DeleteRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDeleteRelationalDatabaseSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailDeleteRelationalDatabaseSnapshotResult) Get(ctx workflow.Context) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error) {
	var output lightsail.DeleteRelationalDatabaseSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDetachCertificateFromDistributionResult struct {
	Result workflow.Future
}

func (r *LightsailDetachCertificateFromDistributionResult) Get(ctx workflow.Context) (*lightsail.DetachCertificateFromDistributionOutput, error) {
	var output lightsail.DetachCertificateFromDistributionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDetachDiskResult struct {
	Result workflow.Future
}

func (r *LightsailDetachDiskResult) Get(ctx workflow.Context) (*lightsail.DetachDiskOutput, error) {
	var output lightsail.DetachDiskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDetachInstancesFromLoadBalancerResult struct {
	Result workflow.Future
}

func (r *LightsailDetachInstancesFromLoadBalancerResult) Get(ctx workflow.Context) (*lightsail.DetachInstancesFromLoadBalancerOutput, error) {
	var output lightsail.DetachInstancesFromLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDetachStaticIpResult struct {
	Result workflow.Future
}

func (r *LightsailDetachStaticIpResult) Get(ctx workflow.Context) (*lightsail.DetachStaticIpOutput, error) {
	var output lightsail.DetachStaticIpOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDisableAddOnResult struct {
	Result workflow.Future
}

func (r *LightsailDisableAddOnResult) Get(ctx workflow.Context) (*lightsail.DisableAddOnOutput, error) {
	var output lightsail.DisableAddOnOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailDownloadDefaultKeyPairResult struct {
	Result workflow.Future
}

func (r *LightsailDownloadDefaultKeyPairResult) Get(ctx workflow.Context) (*lightsail.DownloadDefaultKeyPairOutput, error) {
	var output lightsail.DownloadDefaultKeyPairOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailEnableAddOnResult struct {
	Result workflow.Future
}

func (r *LightsailEnableAddOnResult) Get(ctx workflow.Context) (*lightsail.EnableAddOnOutput, error) {
	var output lightsail.EnableAddOnOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailExportSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailExportSnapshotResult) Get(ctx workflow.Context) (*lightsail.ExportSnapshotOutput, error) {
	var output lightsail.ExportSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetActiveNamesResult struct {
	Result workflow.Future
}

func (r *LightsailGetActiveNamesResult) Get(ctx workflow.Context) (*lightsail.GetActiveNamesOutput, error) {
	var output lightsail.GetActiveNamesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetAlarmsResult struct {
	Result workflow.Future
}

func (r *LightsailGetAlarmsResult) Get(ctx workflow.Context) (*lightsail.GetAlarmsOutput, error) {
	var output lightsail.GetAlarmsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetAutoSnapshotsResult struct {
	Result workflow.Future
}

func (r *LightsailGetAutoSnapshotsResult) Get(ctx workflow.Context) (*lightsail.GetAutoSnapshotsOutput, error) {
	var output lightsail.GetAutoSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetBlueprintsResult struct {
	Result workflow.Future
}

func (r *LightsailGetBlueprintsResult) Get(ctx workflow.Context) (*lightsail.GetBlueprintsOutput, error) {
	var output lightsail.GetBlueprintsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetBundlesResult struct {
	Result workflow.Future
}

func (r *LightsailGetBundlesResult) Get(ctx workflow.Context) (*lightsail.GetBundlesOutput, error) {
	var output lightsail.GetBundlesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetCertificatesResult struct {
	Result workflow.Future
}

func (r *LightsailGetCertificatesResult) Get(ctx workflow.Context) (*lightsail.GetCertificatesOutput, error) {
	var output lightsail.GetCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetCloudFormationStackRecordsResult struct {
	Result workflow.Future
}

func (r *LightsailGetCloudFormationStackRecordsResult) Get(ctx workflow.Context) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	var output lightsail.GetCloudFormationStackRecordsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetContactMethodsResult struct {
	Result workflow.Future
}

func (r *LightsailGetContactMethodsResult) Get(ctx workflow.Context) (*lightsail.GetContactMethodsOutput, error) {
	var output lightsail.GetContactMethodsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDiskResult struct {
	Result workflow.Future
}

func (r *LightsailGetDiskResult) Get(ctx workflow.Context) (*lightsail.GetDiskOutput, error) {
	var output lightsail.GetDiskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDiskSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailGetDiskSnapshotResult) Get(ctx workflow.Context) (*lightsail.GetDiskSnapshotOutput, error) {
	var output lightsail.GetDiskSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDiskSnapshotsResult struct {
	Result workflow.Future
}

func (r *LightsailGetDiskSnapshotsResult) Get(ctx workflow.Context) (*lightsail.GetDiskSnapshotsOutput, error) {
	var output lightsail.GetDiskSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDisksResult struct {
	Result workflow.Future
}

func (r *LightsailGetDisksResult) Get(ctx workflow.Context) (*lightsail.GetDisksOutput, error) {
	var output lightsail.GetDisksOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDistributionBundlesResult struct {
	Result workflow.Future
}

func (r *LightsailGetDistributionBundlesResult) Get(ctx workflow.Context) (*lightsail.GetDistributionBundlesOutput, error) {
	var output lightsail.GetDistributionBundlesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDistributionLatestCacheResetResult struct {
	Result workflow.Future
}

func (r *LightsailGetDistributionLatestCacheResetResult) Get(ctx workflow.Context) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	var output lightsail.GetDistributionLatestCacheResetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDistributionMetricDataResult struct {
	Result workflow.Future
}

func (r *LightsailGetDistributionMetricDataResult) Get(ctx workflow.Context) (*lightsail.GetDistributionMetricDataOutput, error) {
	var output lightsail.GetDistributionMetricDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDistributionsResult struct {
	Result workflow.Future
}

func (r *LightsailGetDistributionsResult) Get(ctx workflow.Context) (*lightsail.GetDistributionsOutput, error) {
	var output lightsail.GetDistributionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDomainResult struct {
	Result workflow.Future
}

func (r *LightsailGetDomainResult) Get(ctx workflow.Context) (*lightsail.GetDomainOutput, error) {
	var output lightsail.GetDomainOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetDomainsResult struct {
	Result workflow.Future
}

func (r *LightsailGetDomainsResult) Get(ctx workflow.Context) (*lightsail.GetDomainsOutput, error) {
	var output lightsail.GetDomainsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetExportSnapshotRecordsResult struct {
	Result workflow.Future
}

func (r *LightsailGetExportSnapshotRecordsResult) Get(ctx workflow.Context) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	var output lightsail.GetExportSnapshotRecordsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstanceResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstanceResult) Get(ctx workflow.Context) (*lightsail.GetInstanceOutput, error) {
	var output lightsail.GetInstanceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstanceAccessDetailsResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstanceAccessDetailsResult) Get(ctx workflow.Context) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	var output lightsail.GetInstanceAccessDetailsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstanceMetricDataResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstanceMetricDataResult) Get(ctx workflow.Context) (*lightsail.GetInstanceMetricDataOutput, error) {
	var output lightsail.GetInstanceMetricDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstancePortStatesResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstancePortStatesResult) Get(ctx workflow.Context) (*lightsail.GetInstancePortStatesOutput, error) {
	var output lightsail.GetInstancePortStatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstanceSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstanceSnapshotResult) Get(ctx workflow.Context) (*lightsail.GetInstanceSnapshotOutput, error) {
	var output lightsail.GetInstanceSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstanceSnapshotsResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstanceSnapshotsResult) Get(ctx workflow.Context) (*lightsail.GetInstanceSnapshotsOutput, error) {
	var output lightsail.GetInstanceSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstanceStateResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstanceStateResult) Get(ctx workflow.Context) (*lightsail.GetInstanceStateOutput, error) {
	var output lightsail.GetInstanceStateOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetInstancesResult struct {
	Result workflow.Future
}

func (r *LightsailGetInstancesResult) Get(ctx workflow.Context) (*lightsail.GetInstancesOutput, error) {
	var output lightsail.GetInstancesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetKeyPairResult struct {
	Result workflow.Future
}

func (r *LightsailGetKeyPairResult) Get(ctx workflow.Context) (*lightsail.GetKeyPairOutput, error) {
	var output lightsail.GetKeyPairOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetKeyPairsResult struct {
	Result workflow.Future
}

func (r *LightsailGetKeyPairsResult) Get(ctx workflow.Context) (*lightsail.GetKeyPairsOutput, error) {
	var output lightsail.GetKeyPairsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetLoadBalancerResult struct {
	Result workflow.Future
}

func (r *LightsailGetLoadBalancerResult) Get(ctx workflow.Context) (*lightsail.GetLoadBalancerOutput, error) {
	var output lightsail.GetLoadBalancerOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetLoadBalancerMetricDataResult struct {
	Result workflow.Future
}

func (r *LightsailGetLoadBalancerMetricDataResult) Get(ctx workflow.Context) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	var output lightsail.GetLoadBalancerMetricDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetLoadBalancerTlsCertificatesResult struct {
	Result workflow.Future
}

func (r *LightsailGetLoadBalancerTlsCertificatesResult) Get(ctx workflow.Context) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	var output lightsail.GetLoadBalancerTlsCertificatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetLoadBalancersResult struct {
	Result workflow.Future
}

func (r *LightsailGetLoadBalancersResult) Get(ctx workflow.Context) (*lightsail.GetLoadBalancersOutput, error) {
	var output lightsail.GetLoadBalancersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetOperationResult struct {
	Result workflow.Future
}

func (r *LightsailGetOperationResult) Get(ctx workflow.Context) (*lightsail.GetOperationOutput, error) {
	var output lightsail.GetOperationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetOperationsResult struct {
	Result workflow.Future
}

func (r *LightsailGetOperationsResult) Get(ctx workflow.Context) (*lightsail.GetOperationsOutput, error) {
	var output lightsail.GetOperationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetOperationsForResourceResult struct {
	Result workflow.Future
}

func (r *LightsailGetOperationsForResourceResult) Get(ctx workflow.Context) (*lightsail.GetOperationsForResourceOutput, error) {
	var output lightsail.GetOperationsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRegionsResult struct {
	Result workflow.Future
}

func (r *LightsailGetRegionsResult) Get(ctx workflow.Context) (*lightsail.GetRegionsOutput, error) {
	var output lightsail.GetRegionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseOutput, error) {
	var output lightsail.GetRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseBlueprintsResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseBlueprintsResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	var output lightsail.GetRelationalDatabaseBlueprintsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseBundlesResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseBundlesResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	var output lightsail.GetRelationalDatabaseBundlesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseEventsResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseEventsResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	var output lightsail.GetRelationalDatabaseEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseLogEventsResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseLogEventsResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	var output lightsail.GetRelationalDatabaseLogEventsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseLogStreamsResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseLogStreamsResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	var output lightsail.GetRelationalDatabaseLogStreamsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseMasterUserPasswordResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseMasterUserPasswordResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	var output lightsail.GetRelationalDatabaseMasterUserPasswordOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseMetricDataResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseMetricDataResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	var output lightsail.GetRelationalDatabaseMetricDataOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseParametersResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseParametersResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	var output lightsail.GetRelationalDatabaseParametersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseSnapshotResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseSnapshotResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	var output lightsail.GetRelationalDatabaseSnapshotOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabaseSnapshotsResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabaseSnapshotsResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	var output lightsail.GetRelationalDatabaseSnapshotsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetRelationalDatabasesResult struct {
	Result workflow.Future
}

func (r *LightsailGetRelationalDatabasesResult) Get(ctx workflow.Context) (*lightsail.GetRelationalDatabasesOutput, error) {
	var output lightsail.GetRelationalDatabasesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetStaticIpResult struct {
	Result workflow.Future
}

func (r *LightsailGetStaticIpResult) Get(ctx workflow.Context) (*lightsail.GetStaticIpOutput, error) {
	var output lightsail.GetStaticIpOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailGetStaticIpsResult struct {
	Result workflow.Future
}

func (r *LightsailGetStaticIpsResult) Get(ctx workflow.Context) (*lightsail.GetStaticIpsOutput, error) {
	var output lightsail.GetStaticIpsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailImportKeyPairResult struct {
	Result workflow.Future
}

func (r *LightsailImportKeyPairResult) Get(ctx workflow.Context) (*lightsail.ImportKeyPairOutput, error) {
	var output lightsail.ImportKeyPairOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailIsVpcPeeredResult struct {
	Result workflow.Future
}

func (r *LightsailIsVpcPeeredResult) Get(ctx workflow.Context) (*lightsail.IsVpcPeeredOutput, error) {
	var output lightsail.IsVpcPeeredOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailOpenInstancePublicPortsResult struct {
	Result workflow.Future
}

func (r *LightsailOpenInstancePublicPortsResult) Get(ctx workflow.Context) (*lightsail.OpenInstancePublicPortsOutput, error) {
	var output lightsail.OpenInstancePublicPortsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailPeerVpcResult struct {
	Result workflow.Future
}

func (r *LightsailPeerVpcResult) Get(ctx workflow.Context) (*lightsail.PeerVpcOutput, error) {
	var output lightsail.PeerVpcOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailPutAlarmResult struct {
	Result workflow.Future
}

func (r *LightsailPutAlarmResult) Get(ctx workflow.Context) (*lightsail.PutAlarmOutput, error) {
	var output lightsail.PutAlarmOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailPutInstancePublicPortsResult struct {
	Result workflow.Future
}

func (r *LightsailPutInstancePublicPortsResult) Get(ctx workflow.Context) (*lightsail.PutInstancePublicPortsOutput, error) {
	var output lightsail.PutInstancePublicPortsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailRebootInstanceResult struct {
	Result workflow.Future
}

func (r *LightsailRebootInstanceResult) Get(ctx workflow.Context) (*lightsail.RebootInstanceOutput, error) {
	var output lightsail.RebootInstanceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailRebootRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailRebootRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.RebootRelationalDatabaseOutput, error) {
	var output lightsail.RebootRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailReleaseStaticIpResult struct {
	Result workflow.Future
}

func (r *LightsailReleaseStaticIpResult) Get(ctx workflow.Context) (*lightsail.ReleaseStaticIpOutput, error) {
	var output lightsail.ReleaseStaticIpOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailResetDistributionCacheResult struct {
	Result workflow.Future
}

func (r *LightsailResetDistributionCacheResult) Get(ctx workflow.Context) (*lightsail.ResetDistributionCacheOutput, error) {
	var output lightsail.ResetDistributionCacheOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailSendContactMethodVerificationResult struct {
	Result workflow.Future
}

func (r *LightsailSendContactMethodVerificationResult) Get(ctx workflow.Context) (*lightsail.SendContactMethodVerificationOutput, error) {
	var output lightsail.SendContactMethodVerificationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailStartInstanceResult struct {
	Result workflow.Future
}

func (r *LightsailStartInstanceResult) Get(ctx workflow.Context) (*lightsail.StartInstanceOutput, error) {
	var output lightsail.StartInstanceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailStartRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailStartRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.StartRelationalDatabaseOutput, error) {
	var output lightsail.StartRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailStopInstanceResult struct {
	Result workflow.Future
}

func (r *LightsailStopInstanceResult) Get(ctx workflow.Context) (*lightsail.StopInstanceOutput, error) {
	var output lightsail.StopInstanceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailStopRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailStopRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.StopRelationalDatabaseOutput, error) {
	var output lightsail.StopRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailTagResourceResult struct {
	Result workflow.Future
}

func (r *LightsailTagResourceResult) Get(ctx workflow.Context) (*lightsail.TagResourceOutput, error) {
	var output lightsail.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailTestAlarmResult struct {
	Result workflow.Future
}

func (r *LightsailTestAlarmResult) Get(ctx workflow.Context) (*lightsail.TestAlarmOutput, error) {
	var output lightsail.TestAlarmOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUnpeerVpcResult struct {
	Result workflow.Future
}

func (r *LightsailUnpeerVpcResult) Get(ctx workflow.Context) (*lightsail.UnpeerVpcOutput, error) {
	var output lightsail.UnpeerVpcOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUntagResourceResult struct {
	Result workflow.Future
}

func (r *LightsailUntagResourceResult) Get(ctx workflow.Context) (*lightsail.UntagResourceOutput, error) {
	var output lightsail.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUpdateDistributionResult struct {
	Result workflow.Future
}

func (r *LightsailUpdateDistributionResult) Get(ctx workflow.Context) (*lightsail.UpdateDistributionOutput, error) {
	var output lightsail.UpdateDistributionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUpdateDistributionBundleResult struct {
	Result workflow.Future
}

func (r *LightsailUpdateDistributionBundleResult) Get(ctx workflow.Context) (*lightsail.UpdateDistributionBundleOutput, error) {
	var output lightsail.UpdateDistributionBundleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUpdateDomainEntryResult struct {
	Result workflow.Future
}

func (r *LightsailUpdateDomainEntryResult) Get(ctx workflow.Context) (*lightsail.UpdateDomainEntryOutput, error) {
	var output lightsail.UpdateDomainEntryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUpdateLoadBalancerAttributeResult struct {
	Result workflow.Future
}

func (r *LightsailUpdateLoadBalancerAttributeResult) Get(ctx workflow.Context) (*lightsail.UpdateLoadBalancerAttributeOutput, error) {
	var output lightsail.UpdateLoadBalancerAttributeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUpdateRelationalDatabaseResult struct {
	Result workflow.Future
}

func (r *LightsailUpdateRelationalDatabaseResult) Get(ctx workflow.Context) (*lightsail.UpdateRelationalDatabaseOutput, error) {
	var output lightsail.UpdateRelationalDatabaseOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailUpdateRelationalDatabaseParametersResult struct {
	Result workflow.Future
}

func (r *LightsailUpdateRelationalDatabaseParametersResult) Get(ctx workflow.Context) (*lightsail.UpdateRelationalDatabaseParametersOutput, error) {
	var output lightsail.UpdateRelationalDatabaseParametersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type LightsailStub struct {
	activities awsactivities.LightsailActivities
}

func NewLightsailStub() LightsailClient {
	return &LightsailStub{}
}

func (a *LightsailStub) AllocateStaticIp(ctx workflow.Context, input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error) {
	var output lightsail.AllocateStaticIpOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AllocateStaticIp, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) AllocateStaticIpAsync(ctx workflow.Context, input *lightsail.AllocateStaticIpInput) *LightsailAllocateStaticIpResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AllocateStaticIp, input)
	return &LightsailAllocateStaticIpResult{Result: future}
}

func (a *LightsailStub) AttachCertificateToDistribution(ctx workflow.Context, input *lightsail.AttachCertificateToDistributionInput) (*lightsail.AttachCertificateToDistributionOutput, error) {
	var output lightsail.AttachCertificateToDistributionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AttachCertificateToDistribution, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) AttachCertificateToDistributionAsync(ctx workflow.Context, input *lightsail.AttachCertificateToDistributionInput) *LightsailAttachCertificateToDistributionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AttachCertificateToDistribution, input)
	return &LightsailAttachCertificateToDistributionResult{Result: future}
}

func (a *LightsailStub) AttachDisk(ctx workflow.Context, input *lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error) {
	var output lightsail.AttachDiskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AttachDisk, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) AttachDiskAsync(ctx workflow.Context, input *lightsail.AttachDiskInput) *LightsailAttachDiskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AttachDisk, input)
	return &LightsailAttachDiskResult{Result: future}
}

func (a *LightsailStub) AttachInstancesToLoadBalancer(ctx workflow.Context, input *lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error) {
	var output lightsail.AttachInstancesToLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AttachInstancesToLoadBalancer, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) AttachInstancesToLoadBalancerAsync(ctx workflow.Context, input *lightsail.AttachInstancesToLoadBalancerInput) *LightsailAttachInstancesToLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AttachInstancesToLoadBalancer, input)
	return &LightsailAttachInstancesToLoadBalancerResult{Result: future}
}

func (a *LightsailStub) AttachLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error) {
	var output lightsail.AttachLoadBalancerTlsCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AttachLoadBalancerTlsCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) AttachLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) *LightsailAttachLoadBalancerTlsCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AttachLoadBalancerTlsCertificate, input)
	return &LightsailAttachLoadBalancerTlsCertificateResult{Result: future}
}

func (a *LightsailStub) AttachStaticIp(ctx workflow.Context, input *lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error) {
	var output lightsail.AttachStaticIpOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AttachStaticIp, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) AttachStaticIpAsync(ctx workflow.Context, input *lightsail.AttachStaticIpInput) *LightsailAttachStaticIpResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AttachStaticIp, input)
	return &LightsailAttachStaticIpResult{Result: future}
}

func (a *LightsailStub) CloseInstancePublicPorts(ctx workflow.Context, input *lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error) {
	var output lightsail.CloseInstancePublicPortsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CloseInstancePublicPorts, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CloseInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.CloseInstancePublicPortsInput) *LightsailCloseInstancePublicPortsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CloseInstancePublicPorts, input)
	return &LightsailCloseInstancePublicPortsResult{Result: future}
}

func (a *LightsailStub) CopySnapshot(ctx workflow.Context, input *lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error) {
	var output lightsail.CopySnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CopySnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CopySnapshotAsync(ctx workflow.Context, input *lightsail.CopySnapshotInput) *LightsailCopySnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CopySnapshot, input)
	return &LightsailCopySnapshotResult{Result: future}
}

func (a *LightsailStub) CreateCertificate(ctx workflow.Context, input *lightsail.CreateCertificateInput) (*lightsail.CreateCertificateOutput, error) {
	var output lightsail.CreateCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateCertificateAsync(ctx workflow.Context, input *lightsail.CreateCertificateInput) *LightsailCreateCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateCertificate, input)
	return &LightsailCreateCertificateResult{Result: future}
}

func (a *LightsailStub) CreateCloudFormationStack(ctx workflow.Context, input *lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error) {
	var output lightsail.CreateCloudFormationStackOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateCloudFormationStack, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateCloudFormationStackAsync(ctx workflow.Context, input *lightsail.CreateCloudFormationStackInput) *LightsailCreateCloudFormationStackResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateCloudFormationStack, input)
	return &LightsailCreateCloudFormationStackResult{Result: future}
}

func (a *LightsailStub) CreateContactMethod(ctx workflow.Context, input *lightsail.CreateContactMethodInput) (*lightsail.CreateContactMethodOutput, error) {
	var output lightsail.CreateContactMethodOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateContactMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateContactMethodAsync(ctx workflow.Context, input *lightsail.CreateContactMethodInput) *LightsailCreateContactMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateContactMethod, input)
	return &LightsailCreateContactMethodResult{Result: future}
}

func (a *LightsailStub) CreateDisk(ctx workflow.Context, input *lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error) {
	var output lightsail.CreateDiskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDisk, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateDiskAsync(ctx workflow.Context, input *lightsail.CreateDiskInput) *LightsailCreateDiskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDisk, input)
	return &LightsailCreateDiskResult{Result: future}
}

func (a *LightsailStub) CreateDiskFromSnapshot(ctx workflow.Context, input *lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error) {
	var output lightsail.CreateDiskFromSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDiskFromSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateDiskFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateDiskFromSnapshotInput) *LightsailCreateDiskFromSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDiskFromSnapshot, input)
	return &LightsailCreateDiskFromSnapshotResult{Result: future}
}

func (a *LightsailStub) CreateDiskSnapshot(ctx workflow.Context, input *lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error) {
	var output lightsail.CreateDiskSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDiskSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateDiskSnapshotAsync(ctx workflow.Context, input *lightsail.CreateDiskSnapshotInput) *LightsailCreateDiskSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDiskSnapshot, input)
	return &LightsailCreateDiskSnapshotResult{Result: future}
}

func (a *LightsailStub) CreateDistribution(ctx workflow.Context, input *lightsail.CreateDistributionInput) (*lightsail.CreateDistributionOutput, error) {
	var output lightsail.CreateDistributionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDistribution, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateDistributionAsync(ctx workflow.Context, input *lightsail.CreateDistributionInput) *LightsailCreateDistributionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDistribution, input)
	return &LightsailCreateDistributionResult{Result: future}
}

func (a *LightsailStub) CreateDomain(ctx workflow.Context, input *lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error) {
	var output lightsail.CreateDomainOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDomain, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateDomainAsync(ctx workflow.Context, input *lightsail.CreateDomainInput) *LightsailCreateDomainResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDomain, input)
	return &LightsailCreateDomainResult{Result: future}
}

func (a *LightsailStub) CreateDomainEntry(ctx workflow.Context, input *lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error) {
	var output lightsail.CreateDomainEntryOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDomainEntry, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateDomainEntryAsync(ctx workflow.Context, input *lightsail.CreateDomainEntryInput) *LightsailCreateDomainEntryResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDomainEntry, input)
	return &LightsailCreateDomainEntryResult{Result: future}
}

func (a *LightsailStub) CreateInstanceSnapshot(ctx workflow.Context, input *lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error) {
	var output lightsail.CreateInstanceSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.CreateInstanceSnapshotInput) *LightsailCreateInstanceSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateInstanceSnapshot, input)
	return &LightsailCreateInstanceSnapshotResult{Result: future}
}

func (a *LightsailStub) CreateInstances(ctx workflow.Context, input *lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error) {
	var output lightsail.CreateInstancesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateInstances, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateInstancesAsync(ctx workflow.Context, input *lightsail.CreateInstancesInput) *LightsailCreateInstancesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateInstances, input)
	return &LightsailCreateInstancesResult{Result: future}
}

func (a *LightsailStub) CreateInstancesFromSnapshot(ctx workflow.Context, input *lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error) {
	var output lightsail.CreateInstancesFromSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateInstancesFromSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateInstancesFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateInstancesFromSnapshotInput) *LightsailCreateInstancesFromSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateInstancesFromSnapshot, input)
	return &LightsailCreateInstancesFromSnapshotResult{Result: future}
}

func (a *LightsailStub) CreateKeyPair(ctx workflow.Context, input *lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error) {
	var output lightsail.CreateKeyPairOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateKeyPair, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateKeyPairAsync(ctx workflow.Context, input *lightsail.CreateKeyPairInput) *LightsailCreateKeyPairResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateKeyPair, input)
	return &LightsailCreateKeyPairResult{Result: future}
}

func (a *LightsailStub) CreateLoadBalancer(ctx workflow.Context, input *lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error) {
	var output lightsail.CreateLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateLoadBalancer, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateLoadBalancerAsync(ctx workflow.Context, input *lightsail.CreateLoadBalancerInput) *LightsailCreateLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateLoadBalancer, input)
	return &LightsailCreateLoadBalancerResult{Result: future}
}

func (a *LightsailStub) CreateLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error) {
	var output lightsail.CreateLoadBalancerTlsCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateLoadBalancerTlsCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) *LightsailCreateLoadBalancerTlsCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateLoadBalancerTlsCertificate, input)
	return &LightsailCreateLoadBalancerTlsCertificateResult{Result: future}
}

func (a *LightsailStub) CreateRelationalDatabase(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error) {
	var output lightsail.CreateRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseInput) *LightsailCreateRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRelationalDatabase, input)
	return &LightsailCreateRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) CreateRelationalDatabaseFromSnapshot(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error) {
	var output lightsail.CreateRelationalDatabaseFromSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRelationalDatabaseFromSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateRelationalDatabaseFromSnapshotAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) *LightsailCreateRelationalDatabaseFromSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRelationalDatabaseFromSnapshot, input)
	return &LightsailCreateRelationalDatabaseFromSnapshotResult{Result: future}
}

func (a *LightsailStub) CreateRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error) {
	var output lightsail.CreateRelationalDatabaseSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateRelationalDatabaseSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) CreateRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) *LightsailCreateRelationalDatabaseSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateRelationalDatabaseSnapshot, input)
	return &LightsailCreateRelationalDatabaseSnapshotResult{Result: future}
}

func (a *LightsailStub) DeleteAlarm(ctx workflow.Context, input *lightsail.DeleteAlarmInput) (*lightsail.DeleteAlarmOutput, error) {
	var output lightsail.DeleteAlarmOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAlarm, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteAlarmAsync(ctx workflow.Context, input *lightsail.DeleteAlarmInput) *LightsailDeleteAlarmResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAlarm, input)
	return &LightsailDeleteAlarmResult{Result: future}
}

func (a *LightsailStub) DeleteAutoSnapshot(ctx workflow.Context, input *lightsail.DeleteAutoSnapshotInput) (*lightsail.DeleteAutoSnapshotOutput, error) {
	var output lightsail.DeleteAutoSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAutoSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteAutoSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteAutoSnapshotInput) *LightsailDeleteAutoSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAutoSnapshot, input)
	return &LightsailDeleteAutoSnapshotResult{Result: future}
}

func (a *LightsailStub) DeleteCertificate(ctx workflow.Context, input *lightsail.DeleteCertificateInput) (*lightsail.DeleteCertificateOutput, error) {
	var output lightsail.DeleteCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteCertificateAsync(ctx workflow.Context, input *lightsail.DeleteCertificateInput) *LightsailDeleteCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input)
	return &LightsailDeleteCertificateResult{Result: future}
}

func (a *LightsailStub) DeleteContactMethod(ctx workflow.Context, input *lightsail.DeleteContactMethodInput) (*lightsail.DeleteContactMethodOutput, error) {
	var output lightsail.DeleteContactMethodOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteContactMethod, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteContactMethodAsync(ctx workflow.Context, input *lightsail.DeleteContactMethodInput) *LightsailDeleteContactMethodResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteContactMethod, input)
	return &LightsailDeleteContactMethodResult{Result: future}
}

func (a *LightsailStub) DeleteDisk(ctx workflow.Context, input *lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error) {
	var output lightsail.DeleteDiskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDisk, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteDiskAsync(ctx workflow.Context, input *lightsail.DeleteDiskInput) *LightsailDeleteDiskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDisk, input)
	return &LightsailDeleteDiskResult{Result: future}
}

func (a *LightsailStub) DeleteDiskSnapshot(ctx workflow.Context, input *lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error) {
	var output lightsail.DeleteDiskSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDiskSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteDiskSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteDiskSnapshotInput) *LightsailDeleteDiskSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDiskSnapshot, input)
	return &LightsailDeleteDiskSnapshotResult{Result: future}
}

func (a *LightsailStub) DeleteDistribution(ctx workflow.Context, input *lightsail.DeleteDistributionInput) (*lightsail.DeleteDistributionOutput, error) {
	var output lightsail.DeleteDistributionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDistribution, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteDistributionAsync(ctx workflow.Context, input *lightsail.DeleteDistributionInput) *LightsailDeleteDistributionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDistribution, input)
	return &LightsailDeleteDistributionResult{Result: future}
}

func (a *LightsailStub) DeleteDomain(ctx workflow.Context, input *lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error) {
	var output lightsail.DeleteDomainOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomain, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteDomainAsync(ctx workflow.Context, input *lightsail.DeleteDomainInput) *LightsailDeleteDomainResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDomain, input)
	return &LightsailDeleteDomainResult{Result: future}
}

func (a *LightsailStub) DeleteDomainEntry(ctx workflow.Context, input *lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error) {
	var output lightsail.DeleteDomainEntryOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainEntry, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteDomainEntryAsync(ctx workflow.Context, input *lightsail.DeleteDomainEntryInput) *LightsailDeleteDomainEntryResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainEntry, input)
	return &LightsailDeleteDomainEntryResult{Result: future}
}

func (a *LightsailStub) DeleteInstance(ctx workflow.Context, input *lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error) {
	var output lightsail.DeleteInstanceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteInstance, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteInstanceAsync(ctx workflow.Context, input *lightsail.DeleteInstanceInput) *LightsailDeleteInstanceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteInstance, input)
	return &LightsailDeleteInstanceResult{Result: future}
}

func (a *LightsailStub) DeleteInstanceSnapshot(ctx workflow.Context, input *lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error) {
	var output lightsail.DeleteInstanceSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteInstanceSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteInstanceSnapshotInput) *LightsailDeleteInstanceSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteInstanceSnapshot, input)
	return &LightsailDeleteInstanceSnapshotResult{Result: future}
}

func (a *LightsailStub) DeleteKeyPair(ctx workflow.Context, input *lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error) {
	var output lightsail.DeleteKeyPairOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteKeyPair, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteKeyPairAsync(ctx workflow.Context, input *lightsail.DeleteKeyPairInput) *LightsailDeleteKeyPairResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteKeyPair, input)
	return &LightsailDeleteKeyPairResult{Result: future}
}

func (a *LightsailStub) DeleteKnownHostKeys(ctx workflow.Context, input *lightsail.DeleteKnownHostKeysInput) (*lightsail.DeleteKnownHostKeysOutput, error) {
	var output lightsail.DeleteKnownHostKeysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteKnownHostKeys, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteKnownHostKeysAsync(ctx workflow.Context, input *lightsail.DeleteKnownHostKeysInput) *LightsailDeleteKnownHostKeysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteKnownHostKeys, input)
	return &LightsailDeleteKnownHostKeysResult{Result: future}
}

func (a *LightsailStub) DeleteLoadBalancer(ctx workflow.Context, input *lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error) {
	var output lightsail.DeleteLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoadBalancer, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteLoadBalancerAsync(ctx workflow.Context, input *lightsail.DeleteLoadBalancerInput) *LightsailDeleteLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteLoadBalancer, input)
	return &LightsailDeleteLoadBalancerResult{Result: future}
}

func (a *LightsailStub) DeleteLoadBalancerTlsCertificate(ctx workflow.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error) {
	var output lightsail.DeleteLoadBalancerTlsCertificateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoadBalancerTlsCertificate, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteLoadBalancerTlsCertificateAsync(ctx workflow.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) *LightsailDeleteLoadBalancerTlsCertificateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteLoadBalancerTlsCertificate, input)
	return &LightsailDeleteLoadBalancerTlsCertificateResult{Result: future}
}

func (a *LightsailStub) DeleteRelationalDatabase(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error) {
	var output lightsail.DeleteRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseInput) *LightsailDeleteRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRelationalDatabase, input)
	return &LightsailDeleteRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) DeleteRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error) {
	var output lightsail.DeleteRelationalDatabaseSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteRelationalDatabaseSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DeleteRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) *LightsailDeleteRelationalDatabaseSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteRelationalDatabaseSnapshot, input)
	return &LightsailDeleteRelationalDatabaseSnapshotResult{Result: future}
}

func (a *LightsailStub) DetachCertificateFromDistribution(ctx workflow.Context, input *lightsail.DetachCertificateFromDistributionInput) (*lightsail.DetachCertificateFromDistributionOutput, error) {
	var output lightsail.DetachCertificateFromDistributionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetachCertificateFromDistribution, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DetachCertificateFromDistributionAsync(ctx workflow.Context, input *lightsail.DetachCertificateFromDistributionInput) *LightsailDetachCertificateFromDistributionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetachCertificateFromDistribution, input)
	return &LightsailDetachCertificateFromDistributionResult{Result: future}
}

func (a *LightsailStub) DetachDisk(ctx workflow.Context, input *lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error) {
	var output lightsail.DetachDiskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetachDisk, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DetachDiskAsync(ctx workflow.Context, input *lightsail.DetachDiskInput) *LightsailDetachDiskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetachDisk, input)
	return &LightsailDetachDiskResult{Result: future}
}

func (a *LightsailStub) DetachInstancesFromLoadBalancer(ctx workflow.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error) {
	var output lightsail.DetachInstancesFromLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetachInstancesFromLoadBalancer, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DetachInstancesFromLoadBalancerAsync(ctx workflow.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) *LightsailDetachInstancesFromLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetachInstancesFromLoadBalancer, input)
	return &LightsailDetachInstancesFromLoadBalancerResult{Result: future}
}

func (a *LightsailStub) DetachStaticIp(ctx workflow.Context, input *lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error) {
	var output lightsail.DetachStaticIpOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetachStaticIp, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DetachStaticIpAsync(ctx workflow.Context, input *lightsail.DetachStaticIpInput) *LightsailDetachStaticIpResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetachStaticIp, input)
	return &LightsailDetachStaticIpResult{Result: future}
}

func (a *LightsailStub) DisableAddOn(ctx workflow.Context, input *lightsail.DisableAddOnInput) (*lightsail.DisableAddOnOutput, error) {
	var output lightsail.DisableAddOnOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisableAddOn, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DisableAddOnAsync(ctx workflow.Context, input *lightsail.DisableAddOnInput) *LightsailDisableAddOnResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisableAddOn, input)
	return &LightsailDisableAddOnResult{Result: future}
}

func (a *LightsailStub) DownloadDefaultKeyPair(ctx workflow.Context, input *lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error) {
	var output lightsail.DownloadDefaultKeyPairOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DownloadDefaultKeyPair, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) DownloadDefaultKeyPairAsync(ctx workflow.Context, input *lightsail.DownloadDefaultKeyPairInput) *LightsailDownloadDefaultKeyPairResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DownloadDefaultKeyPair, input)
	return &LightsailDownloadDefaultKeyPairResult{Result: future}
}

func (a *LightsailStub) EnableAddOn(ctx workflow.Context, input *lightsail.EnableAddOnInput) (*lightsail.EnableAddOnOutput, error) {
	var output lightsail.EnableAddOnOutput
	err := workflow.ExecuteActivity(ctx, a.activities.EnableAddOn, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) EnableAddOnAsync(ctx workflow.Context, input *lightsail.EnableAddOnInput) *LightsailEnableAddOnResult {
	future := workflow.ExecuteActivity(ctx, a.activities.EnableAddOn, input)
	return &LightsailEnableAddOnResult{Result: future}
}

func (a *LightsailStub) ExportSnapshot(ctx workflow.Context, input *lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error) {
	var output lightsail.ExportSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ExportSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) ExportSnapshotAsync(ctx workflow.Context, input *lightsail.ExportSnapshotInput) *LightsailExportSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ExportSnapshot, input)
	return &LightsailExportSnapshotResult{Result: future}
}

func (a *LightsailStub) GetActiveNames(ctx workflow.Context, input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error) {
	var output lightsail.GetActiveNamesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetActiveNames, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetActiveNamesAsync(ctx workflow.Context, input *lightsail.GetActiveNamesInput) *LightsailGetActiveNamesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetActiveNames, input)
	return &LightsailGetActiveNamesResult{Result: future}
}

func (a *LightsailStub) GetAlarms(ctx workflow.Context, input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error) {
	var output lightsail.GetAlarmsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAlarms, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetAlarmsAsync(ctx workflow.Context, input *lightsail.GetAlarmsInput) *LightsailGetAlarmsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAlarms, input)
	return &LightsailGetAlarmsResult{Result: future}
}

func (a *LightsailStub) GetAutoSnapshots(ctx workflow.Context, input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error) {
	var output lightsail.GetAutoSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAutoSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetAutoSnapshotsAsync(ctx workflow.Context, input *lightsail.GetAutoSnapshotsInput) *LightsailGetAutoSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAutoSnapshots, input)
	return &LightsailGetAutoSnapshotsResult{Result: future}
}

func (a *LightsailStub) GetBlueprints(ctx workflow.Context, input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error) {
	var output lightsail.GetBlueprintsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBlueprints, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetBlueprintsAsync(ctx workflow.Context, input *lightsail.GetBlueprintsInput) *LightsailGetBlueprintsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBlueprints, input)
	return &LightsailGetBlueprintsResult{Result: future}
}

func (a *LightsailStub) GetBundles(ctx workflow.Context, input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error) {
	var output lightsail.GetBundlesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBundles, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetBundlesAsync(ctx workflow.Context, input *lightsail.GetBundlesInput) *LightsailGetBundlesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBundles, input)
	return &LightsailGetBundlesResult{Result: future}
}

func (a *LightsailStub) GetCertificates(ctx workflow.Context, input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error) {
	var output lightsail.GetCertificatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetCertificates, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetCertificatesAsync(ctx workflow.Context, input *lightsail.GetCertificatesInput) *LightsailGetCertificatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetCertificates, input)
	return &LightsailGetCertificatesResult{Result: future}
}

func (a *LightsailStub) GetCloudFormationStackRecords(ctx workflow.Context, input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	var output lightsail.GetCloudFormationStackRecordsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetCloudFormationStackRecords, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetCloudFormationStackRecordsAsync(ctx workflow.Context, input *lightsail.GetCloudFormationStackRecordsInput) *LightsailGetCloudFormationStackRecordsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetCloudFormationStackRecords, input)
	return &LightsailGetCloudFormationStackRecordsResult{Result: future}
}

func (a *LightsailStub) GetContactMethods(ctx workflow.Context, input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error) {
	var output lightsail.GetContactMethodsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetContactMethods, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetContactMethodsAsync(ctx workflow.Context, input *lightsail.GetContactMethodsInput) *LightsailGetContactMethodsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetContactMethods, input)
	return &LightsailGetContactMethodsResult{Result: future}
}

func (a *LightsailStub) GetDisk(ctx workflow.Context, input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error) {
	var output lightsail.GetDiskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDisk, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDiskAsync(ctx workflow.Context, input *lightsail.GetDiskInput) *LightsailGetDiskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDisk, input)
	return &LightsailGetDiskResult{Result: future}
}

func (a *LightsailStub) GetDiskSnapshot(ctx workflow.Context, input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error) {
	var output lightsail.GetDiskSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDiskSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDiskSnapshotAsync(ctx workflow.Context, input *lightsail.GetDiskSnapshotInput) *LightsailGetDiskSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDiskSnapshot, input)
	return &LightsailGetDiskSnapshotResult{Result: future}
}

func (a *LightsailStub) GetDiskSnapshots(ctx workflow.Context, input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error) {
	var output lightsail.GetDiskSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDiskSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDiskSnapshotsAsync(ctx workflow.Context, input *lightsail.GetDiskSnapshotsInput) *LightsailGetDiskSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDiskSnapshots, input)
	return &LightsailGetDiskSnapshotsResult{Result: future}
}

func (a *LightsailStub) GetDisks(ctx workflow.Context, input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error) {
	var output lightsail.GetDisksOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDisks, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDisksAsync(ctx workflow.Context, input *lightsail.GetDisksInput) *LightsailGetDisksResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDisks, input)
	return &LightsailGetDisksResult{Result: future}
}

func (a *LightsailStub) GetDistributionBundles(ctx workflow.Context, input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error) {
	var output lightsail.GetDistributionBundlesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDistributionBundles, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDistributionBundlesAsync(ctx workflow.Context, input *lightsail.GetDistributionBundlesInput) *LightsailGetDistributionBundlesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDistributionBundles, input)
	return &LightsailGetDistributionBundlesResult{Result: future}
}

func (a *LightsailStub) GetDistributionLatestCacheReset(ctx workflow.Context, input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	var output lightsail.GetDistributionLatestCacheResetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDistributionLatestCacheReset, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDistributionLatestCacheResetAsync(ctx workflow.Context, input *lightsail.GetDistributionLatestCacheResetInput) *LightsailGetDistributionLatestCacheResetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDistributionLatestCacheReset, input)
	return &LightsailGetDistributionLatestCacheResetResult{Result: future}
}

func (a *LightsailStub) GetDistributionMetricData(ctx workflow.Context, input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error) {
	var output lightsail.GetDistributionMetricDataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDistributionMetricData, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDistributionMetricDataAsync(ctx workflow.Context, input *lightsail.GetDistributionMetricDataInput) *LightsailGetDistributionMetricDataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDistributionMetricData, input)
	return &LightsailGetDistributionMetricDataResult{Result: future}
}

func (a *LightsailStub) GetDistributions(ctx workflow.Context, input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error) {
	var output lightsail.GetDistributionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDistributions, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDistributionsAsync(ctx workflow.Context, input *lightsail.GetDistributionsInput) *LightsailGetDistributionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDistributions, input)
	return &LightsailGetDistributionsResult{Result: future}
}

func (a *LightsailStub) GetDomain(ctx workflow.Context, input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error) {
	var output lightsail.GetDomainOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDomain, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDomainAsync(ctx workflow.Context, input *lightsail.GetDomainInput) *LightsailGetDomainResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDomain, input)
	return &LightsailGetDomainResult{Result: future}
}

func (a *LightsailStub) GetDomains(ctx workflow.Context, input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error) {
	var output lightsail.GetDomainsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDomains, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetDomainsAsync(ctx workflow.Context, input *lightsail.GetDomainsInput) *LightsailGetDomainsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDomains, input)
	return &LightsailGetDomainsResult{Result: future}
}

func (a *LightsailStub) GetExportSnapshotRecords(ctx workflow.Context, input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	var output lightsail.GetExportSnapshotRecordsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetExportSnapshotRecords, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetExportSnapshotRecordsAsync(ctx workflow.Context, input *lightsail.GetExportSnapshotRecordsInput) *LightsailGetExportSnapshotRecordsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetExportSnapshotRecords, input)
	return &LightsailGetExportSnapshotRecordsResult{Result: future}
}

func (a *LightsailStub) GetInstance(ctx workflow.Context, input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error) {
	var output lightsail.GetInstanceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstance, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstanceAsync(ctx workflow.Context, input *lightsail.GetInstanceInput) *LightsailGetInstanceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstance, input)
	return &LightsailGetInstanceResult{Result: future}
}

func (a *LightsailStub) GetInstanceAccessDetails(ctx workflow.Context, input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	var output lightsail.GetInstanceAccessDetailsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceAccessDetails, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstanceAccessDetailsAsync(ctx workflow.Context, input *lightsail.GetInstanceAccessDetailsInput) *LightsailGetInstanceAccessDetailsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceAccessDetails, input)
	return &LightsailGetInstanceAccessDetailsResult{Result: future}
}

func (a *LightsailStub) GetInstanceMetricData(ctx workflow.Context, input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error) {
	var output lightsail.GetInstanceMetricDataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceMetricData, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstanceMetricDataAsync(ctx workflow.Context, input *lightsail.GetInstanceMetricDataInput) *LightsailGetInstanceMetricDataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceMetricData, input)
	return &LightsailGetInstanceMetricDataResult{Result: future}
}

func (a *LightsailStub) GetInstancePortStates(ctx workflow.Context, input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error) {
	var output lightsail.GetInstancePortStatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstancePortStates, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstancePortStatesAsync(ctx workflow.Context, input *lightsail.GetInstancePortStatesInput) *LightsailGetInstancePortStatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstancePortStates, input)
	return &LightsailGetInstancePortStatesResult{Result: future}
}

func (a *LightsailStub) GetInstanceSnapshot(ctx workflow.Context, input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error) {
	var output lightsail.GetInstanceSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstanceSnapshotAsync(ctx workflow.Context, input *lightsail.GetInstanceSnapshotInput) *LightsailGetInstanceSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceSnapshot, input)
	return &LightsailGetInstanceSnapshotResult{Result: future}
}

func (a *LightsailStub) GetInstanceSnapshots(ctx workflow.Context, input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error) {
	var output lightsail.GetInstanceSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstanceSnapshotsAsync(ctx workflow.Context, input *lightsail.GetInstanceSnapshotsInput) *LightsailGetInstanceSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceSnapshots, input)
	return &LightsailGetInstanceSnapshotsResult{Result: future}
}

func (a *LightsailStub) GetInstanceState(ctx workflow.Context, input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error) {
	var output lightsail.GetInstanceStateOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceState, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstanceStateAsync(ctx workflow.Context, input *lightsail.GetInstanceStateInput) *LightsailGetInstanceStateResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceState, input)
	return &LightsailGetInstanceStateResult{Result: future}
}

func (a *LightsailStub) GetInstances(ctx workflow.Context, input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error) {
	var output lightsail.GetInstancesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetInstances, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetInstancesAsync(ctx workflow.Context, input *lightsail.GetInstancesInput) *LightsailGetInstancesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetInstances, input)
	return &LightsailGetInstancesResult{Result: future}
}

func (a *LightsailStub) GetKeyPair(ctx workflow.Context, input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error) {
	var output lightsail.GetKeyPairOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetKeyPair, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetKeyPairAsync(ctx workflow.Context, input *lightsail.GetKeyPairInput) *LightsailGetKeyPairResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetKeyPair, input)
	return &LightsailGetKeyPairResult{Result: future}
}

func (a *LightsailStub) GetKeyPairs(ctx workflow.Context, input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error) {
	var output lightsail.GetKeyPairsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetKeyPairs, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetKeyPairsAsync(ctx workflow.Context, input *lightsail.GetKeyPairsInput) *LightsailGetKeyPairsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetKeyPairs, input)
	return &LightsailGetKeyPairsResult{Result: future}
}

func (a *LightsailStub) GetLoadBalancer(ctx workflow.Context, input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error) {
	var output lightsail.GetLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancer, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetLoadBalancerAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerInput) *LightsailGetLoadBalancerResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancer, input)
	return &LightsailGetLoadBalancerResult{Result: future}
}

func (a *LightsailStub) GetLoadBalancerMetricData(ctx workflow.Context, input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	var output lightsail.GetLoadBalancerMetricDataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancerMetricData, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetLoadBalancerMetricDataAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerMetricDataInput) *LightsailGetLoadBalancerMetricDataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancerMetricData, input)
	return &LightsailGetLoadBalancerMetricDataResult{Result: future}
}

func (a *LightsailStub) GetLoadBalancerTlsCertificates(ctx workflow.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	var output lightsail.GetLoadBalancerTlsCertificatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancerTlsCertificates, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetLoadBalancerTlsCertificatesAsync(ctx workflow.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) *LightsailGetLoadBalancerTlsCertificatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancerTlsCertificates, input)
	return &LightsailGetLoadBalancerTlsCertificatesResult{Result: future}
}

func (a *LightsailStub) GetLoadBalancers(ctx workflow.Context, input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error) {
	var output lightsail.GetLoadBalancersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancers, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetLoadBalancersAsync(ctx workflow.Context, input *lightsail.GetLoadBalancersInput) *LightsailGetLoadBalancersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetLoadBalancers, input)
	return &LightsailGetLoadBalancersResult{Result: future}
}

func (a *LightsailStub) GetOperation(ctx workflow.Context, input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error) {
	var output lightsail.GetOperationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetOperation, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetOperationAsync(ctx workflow.Context, input *lightsail.GetOperationInput) *LightsailGetOperationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetOperation, input)
	return &LightsailGetOperationResult{Result: future}
}

func (a *LightsailStub) GetOperations(ctx workflow.Context, input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error) {
	var output lightsail.GetOperationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetOperations, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetOperationsAsync(ctx workflow.Context, input *lightsail.GetOperationsInput) *LightsailGetOperationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetOperations, input)
	return &LightsailGetOperationsResult{Result: future}
}

func (a *LightsailStub) GetOperationsForResource(ctx workflow.Context, input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error) {
	var output lightsail.GetOperationsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetOperationsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetOperationsForResourceAsync(ctx workflow.Context, input *lightsail.GetOperationsForResourceInput) *LightsailGetOperationsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetOperationsForResource, input)
	return &LightsailGetOperationsForResourceResult{Result: future}
}

func (a *LightsailStub) GetRegions(ctx workflow.Context, input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error) {
	var output lightsail.GetRegionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRegions, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRegionsAsync(ctx workflow.Context, input *lightsail.GetRegionsInput) *LightsailGetRegionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRegions, input)
	return &LightsailGetRegionsResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabase(ctx workflow.Context, input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error) {
	var output lightsail.GetRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseInput) *LightsailGetRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabase, input)
	return &LightsailGetRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseBlueprints(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	var output lightsail.GetRelationalDatabaseBlueprintsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseBlueprints, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseBlueprintsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) *LightsailGetRelationalDatabaseBlueprintsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseBlueprints, input)
	return &LightsailGetRelationalDatabaseBlueprintsResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseBundles(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	var output lightsail.GetRelationalDatabaseBundlesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseBundles, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseBundlesAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseBundlesInput) *LightsailGetRelationalDatabaseBundlesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseBundles, input)
	return &LightsailGetRelationalDatabaseBundlesResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseEvents(ctx workflow.Context, input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	var output lightsail.GetRelationalDatabaseEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseEventsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseEventsInput) *LightsailGetRelationalDatabaseEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseEvents, input)
	return &LightsailGetRelationalDatabaseEventsResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseLogEvents(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	var output lightsail.GetRelationalDatabaseLogEventsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseLogEvents, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseLogEventsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) *LightsailGetRelationalDatabaseLogEventsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseLogEvents, input)
	return &LightsailGetRelationalDatabaseLogEventsResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseLogStreams(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	var output lightsail.GetRelationalDatabaseLogStreamsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseLogStreams, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseLogStreamsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) *LightsailGetRelationalDatabaseLogStreamsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseLogStreams, input)
	return &LightsailGetRelationalDatabaseLogStreamsResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseMasterUserPassword(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	var output lightsail.GetRelationalDatabaseMasterUserPasswordOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseMasterUserPassword, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseMasterUserPasswordAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) *LightsailGetRelationalDatabaseMasterUserPasswordResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseMasterUserPassword, input)
	return &LightsailGetRelationalDatabaseMasterUserPasswordResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseMetricData(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	var output lightsail.GetRelationalDatabaseMetricDataOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseMetricData, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseMetricDataAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) *LightsailGetRelationalDatabaseMetricDataResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseMetricData, input)
	return &LightsailGetRelationalDatabaseMetricDataResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseParameters(ctx workflow.Context, input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	var output lightsail.GetRelationalDatabaseParametersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseParameters, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseParametersAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseParametersInput) *LightsailGetRelationalDatabaseParametersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseParameters, input)
	return &LightsailGetRelationalDatabaseParametersResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseSnapshot(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	var output lightsail.GetRelationalDatabaseSnapshotOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseSnapshot, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseSnapshotAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) *LightsailGetRelationalDatabaseSnapshotResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseSnapshot, input)
	return &LightsailGetRelationalDatabaseSnapshotResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabaseSnapshots(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	var output lightsail.GetRelationalDatabaseSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseSnapshots, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabaseSnapshotsAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) *LightsailGetRelationalDatabaseSnapshotsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabaseSnapshots, input)
	return &LightsailGetRelationalDatabaseSnapshotsResult{Result: future}
}

func (a *LightsailStub) GetRelationalDatabases(ctx workflow.Context, input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error) {
	var output lightsail.GetRelationalDatabasesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabases, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetRelationalDatabasesAsync(ctx workflow.Context, input *lightsail.GetRelationalDatabasesInput) *LightsailGetRelationalDatabasesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetRelationalDatabases, input)
	return &LightsailGetRelationalDatabasesResult{Result: future}
}

func (a *LightsailStub) GetStaticIp(ctx workflow.Context, input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error) {
	var output lightsail.GetStaticIpOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetStaticIp, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetStaticIpAsync(ctx workflow.Context, input *lightsail.GetStaticIpInput) *LightsailGetStaticIpResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetStaticIp, input)
	return &LightsailGetStaticIpResult{Result: future}
}

func (a *LightsailStub) GetStaticIps(ctx workflow.Context, input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error) {
	var output lightsail.GetStaticIpsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetStaticIps, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) GetStaticIpsAsync(ctx workflow.Context, input *lightsail.GetStaticIpsInput) *LightsailGetStaticIpsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetStaticIps, input)
	return &LightsailGetStaticIpsResult{Result: future}
}

func (a *LightsailStub) ImportKeyPair(ctx workflow.Context, input *lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error) {
	var output lightsail.ImportKeyPairOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ImportKeyPair, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) ImportKeyPairAsync(ctx workflow.Context, input *lightsail.ImportKeyPairInput) *LightsailImportKeyPairResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ImportKeyPair, input)
	return &LightsailImportKeyPairResult{Result: future}
}

func (a *LightsailStub) IsVpcPeered(ctx workflow.Context, input *lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error) {
	var output lightsail.IsVpcPeeredOutput
	err := workflow.ExecuteActivity(ctx, a.activities.IsVpcPeered, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) IsVpcPeeredAsync(ctx workflow.Context, input *lightsail.IsVpcPeeredInput) *LightsailIsVpcPeeredResult {
	future := workflow.ExecuteActivity(ctx, a.activities.IsVpcPeered, input)
	return &LightsailIsVpcPeeredResult{Result: future}
}

func (a *LightsailStub) OpenInstancePublicPorts(ctx workflow.Context, input *lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error) {
	var output lightsail.OpenInstancePublicPortsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.OpenInstancePublicPorts, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) OpenInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.OpenInstancePublicPortsInput) *LightsailOpenInstancePublicPortsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.OpenInstancePublicPorts, input)
	return &LightsailOpenInstancePublicPortsResult{Result: future}
}

func (a *LightsailStub) PeerVpc(ctx workflow.Context, input *lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error) {
	var output lightsail.PeerVpcOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PeerVpc, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) PeerVpcAsync(ctx workflow.Context, input *lightsail.PeerVpcInput) *LightsailPeerVpcResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PeerVpc, input)
	return &LightsailPeerVpcResult{Result: future}
}

func (a *LightsailStub) PutAlarm(ctx workflow.Context, input *lightsail.PutAlarmInput) (*lightsail.PutAlarmOutput, error) {
	var output lightsail.PutAlarmOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutAlarm, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) PutAlarmAsync(ctx workflow.Context, input *lightsail.PutAlarmInput) *LightsailPutAlarmResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutAlarm, input)
	return &LightsailPutAlarmResult{Result: future}
}

func (a *LightsailStub) PutInstancePublicPorts(ctx workflow.Context, input *lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error) {
	var output lightsail.PutInstancePublicPortsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutInstancePublicPorts, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) PutInstancePublicPortsAsync(ctx workflow.Context, input *lightsail.PutInstancePublicPortsInput) *LightsailPutInstancePublicPortsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutInstancePublicPorts, input)
	return &LightsailPutInstancePublicPortsResult{Result: future}
}

func (a *LightsailStub) RebootInstance(ctx workflow.Context, input *lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error) {
	var output lightsail.RebootInstanceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RebootInstance, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) RebootInstanceAsync(ctx workflow.Context, input *lightsail.RebootInstanceInput) *LightsailRebootInstanceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RebootInstance, input)
	return &LightsailRebootInstanceResult{Result: future}
}

func (a *LightsailStub) RebootRelationalDatabase(ctx workflow.Context, input *lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error) {
	var output lightsail.RebootRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RebootRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) RebootRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.RebootRelationalDatabaseInput) *LightsailRebootRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RebootRelationalDatabase, input)
	return &LightsailRebootRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) ReleaseStaticIp(ctx workflow.Context, input *lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error) {
	var output lightsail.ReleaseStaticIpOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ReleaseStaticIp, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) ReleaseStaticIpAsync(ctx workflow.Context, input *lightsail.ReleaseStaticIpInput) *LightsailReleaseStaticIpResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ReleaseStaticIp, input)
	return &LightsailReleaseStaticIpResult{Result: future}
}

func (a *LightsailStub) ResetDistributionCache(ctx workflow.Context, input *lightsail.ResetDistributionCacheInput) (*lightsail.ResetDistributionCacheOutput, error) {
	var output lightsail.ResetDistributionCacheOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ResetDistributionCache, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) ResetDistributionCacheAsync(ctx workflow.Context, input *lightsail.ResetDistributionCacheInput) *LightsailResetDistributionCacheResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ResetDistributionCache, input)
	return &LightsailResetDistributionCacheResult{Result: future}
}

func (a *LightsailStub) SendContactMethodVerification(ctx workflow.Context, input *lightsail.SendContactMethodVerificationInput) (*lightsail.SendContactMethodVerificationOutput, error) {
	var output lightsail.SendContactMethodVerificationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SendContactMethodVerification, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) SendContactMethodVerificationAsync(ctx workflow.Context, input *lightsail.SendContactMethodVerificationInput) *LightsailSendContactMethodVerificationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SendContactMethodVerification, input)
	return &LightsailSendContactMethodVerificationResult{Result: future}
}

func (a *LightsailStub) StartInstance(ctx workflow.Context, input *lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error) {
	var output lightsail.StartInstanceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartInstance, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) StartInstanceAsync(ctx workflow.Context, input *lightsail.StartInstanceInput) *LightsailStartInstanceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartInstance, input)
	return &LightsailStartInstanceResult{Result: future}
}

func (a *LightsailStub) StartRelationalDatabase(ctx workflow.Context, input *lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error) {
	var output lightsail.StartRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) StartRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.StartRelationalDatabaseInput) *LightsailStartRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartRelationalDatabase, input)
	return &LightsailStartRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) StopInstance(ctx workflow.Context, input *lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error) {
	var output lightsail.StopInstanceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopInstance, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) StopInstanceAsync(ctx workflow.Context, input *lightsail.StopInstanceInput) *LightsailStopInstanceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopInstance, input)
	return &LightsailStopInstanceResult{Result: future}
}

func (a *LightsailStub) StopRelationalDatabase(ctx workflow.Context, input *lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error) {
	var output lightsail.StopRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StopRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) StopRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.StopRelationalDatabaseInput) *LightsailStopRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StopRelationalDatabase, input)
	return &LightsailStopRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) TagResource(ctx workflow.Context, input *lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error) {
	var output lightsail.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) TagResourceAsync(ctx workflow.Context, input *lightsail.TagResourceInput) *LightsailTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &LightsailTagResourceResult{Result: future}
}

func (a *LightsailStub) TestAlarm(ctx workflow.Context, input *lightsail.TestAlarmInput) (*lightsail.TestAlarmOutput, error) {
	var output lightsail.TestAlarmOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TestAlarm, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) TestAlarmAsync(ctx workflow.Context, input *lightsail.TestAlarmInput) *LightsailTestAlarmResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TestAlarm, input)
	return &LightsailTestAlarmResult{Result: future}
}

func (a *LightsailStub) UnpeerVpc(ctx workflow.Context, input *lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error) {
	var output lightsail.UnpeerVpcOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UnpeerVpc, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UnpeerVpcAsync(ctx workflow.Context, input *lightsail.UnpeerVpcInput) *LightsailUnpeerVpcResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UnpeerVpc, input)
	return &LightsailUnpeerVpcResult{Result: future}
}

func (a *LightsailStub) UntagResource(ctx workflow.Context, input *lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error) {
	var output lightsail.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UntagResourceAsync(ctx workflow.Context, input *lightsail.UntagResourceInput) *LightsailUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &LightsailUntagResourceResult{Result: future}
}

func (a *LightsailStub) UpdateDistribution(ctx workflow.Context, input *lightsail.UpdateDistributionInput) (*lightsail.UpdateDistributionOutput, error) {
	var output lightsail.UpdateDistributionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDistribution, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UpdateDistributionAsync(ctx workflow.Context, input *lightsail.UpdateDistributionInput) *LightsailUpdateDistributionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDistribution, input)
	return &LightsailUpdateDistributionResult{Result: future}
}

func (a *LightsailStub) UpdateDistributionBundle(ctx workflow.Context, input *lightsail.UpdateDistributionBundleInput) (*lightsail.UpdateDistributionBundleOutput, error) {
	var output lightsail.UpdateDistributionBundleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDistributionBundle, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UpdateDistributionBundleAsync(ctx workflow.Context, input *lightsail.UpdateDistributionBundleInput) *LightsailUpdateDistributionBundleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDistributionBundle, input)
	return &LightsailUpdateDistributionBundleResult{Result: future}
}

func (a *LightsailStub) UpdateDomainEntry(ctx workflow.Context, input *lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error) {
	var output lightsail.UpdateDomainEntryOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainEntry, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UpdateDomainEntryAsync(ctx workflow.Context, input *lightsail.UpdateDomainEntryInput) *LightsailUpdateDomainEntryResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainEntry, input)
	return &LightsailUpdateDomainEntryResult{Result: future}
}

func (a *LightsailStub) UpdateLoadBalancerAttribute(ctx workflow.Context, input *lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error) {
	var output lightsail.UpdateLoadBalancerAttributeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateLoadBalancerAttribute, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UpdateLoadBalancerAttributeAsync(ctx workflow.Context, input *lightsail.UpdateLoadBalancerAttributeInput) *LightsailUpdateLoadBalancerAttributeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateLoadBalancerAttribute, input)
	return &LightsailUpdateLoadBalancerAttributeResult{Result: future}
}

func (a *LightsailStub) UpdateRelationalDatabase(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error) {
	var output lightsail.UpdateRelationalDatabaseOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRelationalDatabase, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UpdateRelationalDatabaseAsync(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseInput) *LightsailUpdateRelationalDatabaseResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRelationalDatabase, input)
	return &LightsailUpdateRelationalDatabaseResult{Result: future}
}

func (a *LightsailStub) UpdateRelationalDatabaseParameters(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error) {
	var output lightsail.UpdateRelationalDatabaseParametersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateRelationalDatabaseParameters, input).Get(ctx, &output)
	return &output, err
}

func (a *LightsailStub) UpdateRelationalDatabaseParametersAsync(ctx workflow.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) *LightsailUpdateRelationalDatabaseParametersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateRelationalDatabaseParameters, input)
	return &LightsailUpdateRelationalDatabaseParametersResult{Result: future}
}
