// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

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

	sessionFactory SessionFactory
}

func NewLightsailActivities(sess *session.Session, config ...*aws.Config) *LightsailActivities {
	client := lightsail.New(sess, config...)
	return &LightsailActivities{client: client}
}

func NewLightsailActivitiesWithSessionFactory(sessionFactory SessionFactory) *LightsailActivities {
	return &LightsailActivities{sessionFactory: sessionFactory}
}

func (a *LightsailActivities) getClient(ctx context.Context) (lightsailiface.LightsailAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return lightsail.New(sess), nil
}

func (a *LightsailActivities) AllocateStaticIp(ctx context.Context, input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AllocateStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) AttachCertificateToDistribution(ctx context.Context, input *lightsail.AttachCertificateToDistributionInput) (*lightsail.AttachCertificateToDistributionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachCertificateToDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) AttachDisk(ctx context.Context, input *lightsail.AttachDiskInput) (*lightsail.AttachDiskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachDiskWithContext(ctx, input)
}

func (a *LightsailActivities) AttachInstancesToLoadBalancer(ctx context.Context, input *lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachInstancesToLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) AttachLoadBalancerTlsCertificate(ctx context.Context, input *lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachLoadBalancerTlsCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) AttachStaticIp(ctx context.Context, input *lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) CloseInstancePublicPorts(ctx context.Context, input *lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CloseInstancePublicPortsWithContext(ctx, input)
}

func (a *LightsailActivities) CopySnapshot(ctx context.Context, input *lightsail.CopySnapshotInput) (*lightsail.CopySnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopySnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateCertificate(ctx context.Context, input *lightsail.CreateCertificateInput) (*lightsail.CreateCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) CreateCloudFormationStack(ctx context.Context, input *lightsail.CreateCloudFormationStackInput) (*lightsail.CreateCloudFormationStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCloudFormationStackWithContext(ctx, input)
}

func (a *LightsailActivities) CreateContactMethod(ctx context.Context, input *lightsail.CreateContactMethodInput) (*lightsail.CreateContactMethodOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateContactMethodWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDisk(ctx context.Context, input *lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDiskWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDiskFromSnapshot(ctx context.Context, input *lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDiskFromSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDiskSnapshot(ctx context.Context, input *lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDiskSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDistribution(ctx context.Context, input *lightsail.CreateDistributionInput) (*lightsail.CreateDistributionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDomain(ctx context.Context, input *lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDomainWithContext(ctx, input)
}

func (a *LightsailActivities) CreateDomainEntry(ctx context.Context, input *lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDomainEntryWithContext(ctx, input)
}

func (a *LightsailActivities) CreateInstanceSnapshot(ctx context.Context, input *lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateInstanceSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateInstances(ctx context.Context, input *lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateInstancesWithContext(ctx, input)
}

func (a *LightsailActivities) CreateInstancesFromSnapshot(ctx context.Context, input *lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateInstancesFromSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateKeyPair(ctx context.Context, input *lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) CreateLoadBalancer(ctx context.Context, input *lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) CreateLoadBalancerTlsCertificate(ctx context.Context, input *lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLoadBalancerTlsCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) CreateRelationalDatabase(ctx context.Context, input *lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) CreateRelationalDatabaseFromSnapshot(ctx context.Context, input *lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRelationalDatabaseFromSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) CreateRelationalDatabaseSnapshot(ctx context.Context, input *lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRelationalDatabaseSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteAlarm(ctx context.Context, input *lightsail.DeleteAlarmInput) (*lightsail.DeleteAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAlarmWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteAutoSnapshot(ctx context.Context, input *lightsail.DeleteAutoSnapshotInput) (*lightsail.DeleteAutoSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAutoSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteCertificate(ctx context.Context, input *lightsail.DeleteCertificateInput) (*lightsail.DeleteCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteContactMethod(ctx context.Context, input *lightsail.DeleteContactMethodInput) (*lightsail.DeleteContactMethodOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteContactMethodWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDisk(ctx context.Context, input *lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDiskWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDiskSnapshot(ctx context.Context, input *lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDiskSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDistribution(ctx context.Context, input *lightsail.DeleteDistributionInput) (*lightsail.DeleteDistributionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDomain(ctx context.Context, input *lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDomainWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteDomainEntry(ctx context.Context, input *lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDomainEntryWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteInstance(ctx context.Context, input *lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteInstanceSnapshot(ctx context.Context, input *lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInstanceSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteKeyPair(ctx context.Context, input *lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteKnownHostKeys(ctx context.Context, input *lightsail.DeleteKnownHostKeysInput) (*lightsail.DeleteKnownHostKeysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteKnownHostKeysWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteLoadBalancer(ctx context.Context, input *lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteLoadBalancerTlsCertificate(ctx context.Context, input *lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLoadBalancerTlsCertificateWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteRelationalDatabase(ctx context.Context, input *lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) DeleteRelationalDatabaseSnapshot(ctx context.Context, input *lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRelationalDatabaseSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) DetachCertificateFromDistribution(ctx context.Context, input *lightsail.DetachCertificateFromDistributionInput) (*lightsail.DetachCertificateFromDistributionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetachCertificateFromDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) DetachDisk(ctx context.Context, input *lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetachDiskWithContext(ctx, input)
}

func (a *LightsailActivities) DetachInstancesFromLoadBalancer(ctx context.Context, input *lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetachInstancesFromLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) DetachStaticIp(ctx context.Context, input *lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetachStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) DisableAddOn(ctx context.Context, input *lightsail.DisableAddOnInput) (*lightsail.DisableAddOnOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableAddOnWithContext(ctx, input)
}

func (a *LightsailActivities) DownloadDefaultKeyPair(ctx context.Context, input *lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DownloadDefaultKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) EnableAddOn(ctx context.Context, input *lightsail.EnableAddOnInput) (*lightsail.EnableAddOnOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EnableAddOnWithContext(ctx, input)
}

func (a *LightsailActivities) ExportSnapshot(ctx context.Context, input *lightsail.ExportSnapshotInput) (*lightsail.ExportSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExportSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetActiveNames(ctx context.Context, input *lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetActiveNamesWithContext(ctx, input)
}

func (a *LightsailActivities) GetAlarms(ctx context.Context, input *lightsail.GetAlarmsInput) (*lightsail.GetAlarmsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAlarmsWithContext(ctx, input)
}

func (a *LightsailActivities) GetAutoSnapshots(ctx context.Context, input *lightsail.GetAutoSnapshotsInput) (*lightsail.GetAutoSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAutoSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetBlueprints(ctx context.Context, input *lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBlueprintsWithContext(ctx, input)
}

func (a *LightsailActivities) GetBundles(ctx context.Context, input *lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBundlesWithContext(ctx, input)
}

func (a *LightsailActivities) GetCertificates(ctx context.Context, input *lightsail.GetCertificatesInput) (*lightsail.GetCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCertificatesWithContext(ctx, input)
}

func (a *LightsailActivities) GetCloudFormationStackRecords(ctx context.Context, input *lightsail.GetCloudFormationStackRecordsInput) (*lightsail.GetCloudFormationStackRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCloudFormationStackRecordsWithContext(ctx, input)
}

func (a *LightsailActivities) GetContactMethods(ctx context.Context, input *lightsail.GetContactMethodsInput) (*lightsail.GetContactMethodsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetContactMethodsWithContext(ctx, input)
}

func (a *LightsailActivities) GetDisk(ctx context.Context, input *lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDiskWithContext(ctx, input)
}

func (a *LightsailActivities) GetDiskSnapshot(ctx context.Context, input *lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDiskSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetDiskSnapshots(ctx context.Context, input *lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDiskSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetDisks(ctx context.Context, input *lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDisksWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributionBundles(ctx context.Context, input *lightsail.GetDistributionBundlesInput) (*lightsail.GetDistributionBundlesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDistributionBundlesWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributionLatestCacheReset(ctx context.Context, input *lightsail.GetDistributionLatestCacheResetInput) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDistributionLatestCacheResetWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributionMetricData(ctx context.Context, input *lightsail.GetDistributionMetricDataInput) (*lightsail.GetDistributionMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDistributionMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetDistributions(ctx context.Context, input *lightsail.GetDistributionsInput) (*lightsail.GetDistributionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDistributionsWithContext(ctx, input)
}

func (a *LightsailActivities) GetDomain(ctx context.Context, input *lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainWithContext(ctx, input)
}

func (a *LightsailActivities) GetDomains(ctx context.Context, input *lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainsWithContext(ctx, input)
}

func (a *LightsailActivities) GetExportSnapshotRecords(ctx context.Context, input *lightsail.GetExportSnapshotRecordsInput) (*lightsail.GetExportSnapshotRecordsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetExportSnapshotRecordsWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstance(ctx context.Context, input *lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceAccessDetails(ctx context.Context, input *lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceAccessDetailsWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceMetricData(ctx context.Context, input *lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstancePortStates(ctx context.Context, input *lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstancePortStatesWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceSnapshot(ctx context.Context, input *lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceSnapshots(ctx context.Context, input *lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstanceState(ctx context.Context, input *lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceStateWithContext(ctx, input)
}

func (a *LightsailActivities) GetInstances(ctx context.Context, input *lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstancesWithContext(ctx, input)
}

func (a *LightsailActivities) GetKeyPair(ctx context.Context, input *lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) GetKeyPairs(ctx context.Context, input *lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetKeyPairsWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancer(ctx context.Context, input *lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoadBalancerWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancerMetricData(ctx context.Context, input *lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoadBalancerMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancerTlsCertificates(ctx context.Context, input *lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoadBalancerTlsCertificatesWithContext(ctx, input)
}

func (a *LightsailActivities) GetLoadBalancers(ctx context.Context, input *lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoadBalancersWithContext(ctx, input)
}

func (a *LightsailActivities) GetOperation(ctx context.Context, input *lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOperationWithContext(ctx, input)
}

func (a *LightsailActivities) GetOperations(ctx context.Context, input *lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOperationsWithContext(ctx, input)
}

func (a *LightsailActivities) GetOperationsForResource(ctx context.Context, input *lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOperationsForResourceWithContext(ctx, input)
}

func (a *LightsailActivities) GetRegions(ctx context.Context, input *lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRegionsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabase(ctx context.Context, input *lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseBlueprints(ctx context.Context, input *lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseBlueprintsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseBundles(ctx context.Context, input *lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseBundlesWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseEvents(ctx context.Context, input *lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseEventsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseLogEvents(ctx context.Context, input *lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseLogEventsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseLogStreams(ctx context.Context, input *lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseLogStreamsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseMasterUserPassword(ctx context.Context, input *lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseMasterUserPasswordWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseMetricData(ctx context.Context, input *lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseMetricDataWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseParameters(ctx context.Context, input *lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseParametersWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseSnapshot(ctx context.Context, input *lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseSnapshotWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabaseSnapshots(ctx context.Context, input *lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabaseSnapshotsWithContext(ctx, input)
}

func (a *LightsailActivities) GetRelationalDatabases(ctx context.Context, input *lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRelationalDatabasesWithContext(ctx, input)
}

func (a *LightsailActivities) GetStaticIp(ctx context.Context, input *lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) GetStaticIps(ctx context.Context, input *lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetStaticIpsWithContext(ctx, input)
}

func (a *LightsailActivities) ImportKeyPair(ctx context.Context, input *lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ImportKeyPairWithContext(ctx, input)
}

func (a *LightsailActivities) IsVpcPeered(ctx context.Context, input *lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.IsVpcPeeredWithContext(ctx, input)
}

func (a *LightsailActivities) OpenInstancePublicPorts(ctx context.Context, input *lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.OpenInstancePublicPortsWithContext(ctx, input)
}

func (a *LightsailActivities) PeerVpc(ctx context.Context, input *lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PeerVpcWithContext(ctx, input)
}

func (a *LightsailActivities) PutAlarm(ctx context.Context, input *lightsail.PutAlarmInput) (*lightsail.PutAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAlarmWithContext(ctx, input)
}

func (a *LightsailActivities) PutInstancePublicPorts(ctx context.Context, input *lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutInstancePublicPortsWithContext(ctx, input)
}

func (a *LightsailActivities) RebootInstance(ctx context.Context, input *lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) RebootRelationalDatabase(ctx context.Context, input *lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) ReleaseStaticIp(ctx context.Context, input *lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ReleaseStaticIpWithContext(ctx, input)
}

func (a *LightsailActivities) ResetDistributionCache(ctx context.Context, input *lightsail.ResetDistributionCacheInput) (*lightsail.ResetDistributionCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetDistributionCacheWithContext(ctx, input)
}

func (a *LightsailActivities) SendContactMethodVerification(ctx context.Context, input *lightsail.SendContactMethodVerificationInput) (*lightsail.SendContactMethodVerificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendContactMethodVerificationWithContext(ctx, input)
}

func (a *LightsailActivities) StartInstance(ctx context.Context, input *lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) StartRelationalDatabase(ctx context.Context, input *lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) StopInstance(ctx context.Context, input *lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopInstanceWithContext(ctx, input)
}

func (a *LightsailActivities) StopRelationalDatabase(ctx context.Context, input *lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) TagResource(ctx context.Context, input *lightsail.TagResourceInput) (*lightsail.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *LightsailActivities) TestAlarm(ctx context.Context, input *lightsail.TestAlarmInput) (*lightsail.TestAlarmOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TestAlarmWithContext(ctx, input)
}

func (a *LightsailActivities) UnpeerVpc(ctx context.Context, input *lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UnpeerVpcWithContext(ctx, input)
}

func (a *LightsailActivities) UntagResource(ctx context.Context, input *lightsail.UntagResourceInput) (*lightsail.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateDistribution(ctx context.Context, input *lightsail.UpdateDistributionInput) (*lightsail.UpdateDistributionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDistributionWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateDistributionBundle(ctx context.Context, input *lightsail.UpdateDistributionBundleInput) (*lightsail.UpdateDistributionBundleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDistributionBundleWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateDomainEntry(ctx context.Context, input *lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDomainEntryWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateLoadBalancerAttribute(ctx context.Context, input *lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateLoadBalancerAttributeWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateRelationalDatabase(ctx context.Context, input *lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRelationalDatabaseWithContext(ctx, input)
}

func (a *LightsailActivities) UpdateRelationalDatabaseParameters(ctx context.Context, input *lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRelationalDatabaseParametersWithContext(ctx, input)
}
