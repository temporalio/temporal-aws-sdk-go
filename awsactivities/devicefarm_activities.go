// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DeviceFarmActivities struct {
	client devicefarmiface.DeviceFarmAPI

	sessionFactory SessionFactory
}

func NewDeviceFarmActivities(sess *session.Session, config ...*aws.Config) *DeviceFarmActivities {
	client := devicefarm.New(sess, config...)
	return &DeviceFarmActivities{client: client}
}

func NewDeviceFarmActivitiesWithSessionFactory(sessionFactory SessionFactory) *DeviceFarmActivities {
	return &DeviceFarmActivities{sessionFactory: sessionFactory}
}

func (a *DeviceFarmActivities) getClient(ctx context.Context) (devicefarmiface.DeviceFarmAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return devicefarm.New(sess), nil
}

func (a *DeviceFarmActivities) CreateDevicePool(ctx context.Context, input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateInstanceProfile(ctx context.Context, input *devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateNetworkProfile(ctx context.Context, input *devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateProject(ctx context.Context, input *devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateRemoteAccessSession(ctx context.Context, input *devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateTestGridProject(ctx context.Context, input *devicefarm.CreateTestGridProjectInput) (*devicefarm.CreateTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateTestGridUrl(ctx context.Context, input *devicefarm.CreateTestGridUrlInput) (*devicefarm.CreateTestGridUrlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTestGridUrlWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateUpload(ctx context.Context, input *devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) CreateVPCEConfiguration(ctx context.Context, input *devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVPCEConfigurationWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteDevicePool(ctx context.Context, input *devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteInstanceProfile(ctx context.Context, input *devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteNetworkProfile(ctx context.Context, input *devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteProject(ctx context.Context, input *devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteRemoteAccessSession(ctx context.Context, input *devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteRun(ctx context.Context, input *devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteTestGridProject(ctx context.Context, input *devicefarm.DeleteTestGridProjectInput) (*devicefarm.DeleteTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteUpload(ctx context.Context, input *devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) DeleteVPCEConfiguration(ctx context.Context, input *devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVPCEConfigurationWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetAccountSettings(ctx context.Context, input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAccountSettingsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDevice(ctx context.Context, input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeviceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDeviceInstance(ctx context.Context, input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeviceInstanceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDevicePool(ctx context.Context, input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetDevicePoolCompatibility(ctx context.Context, input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDevicePoolCompatibilityWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetInstanceProfile(ctx context.Context, input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetJob(ctx context.Context, input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetJobWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetNetworkProfile(ctx context.Context, input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetOfferingStatus(ctx context.Context, input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOfferingStatusWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetProject(ctx context.Context, input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetRemoteAccessSession(ctx context.Context, input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetRun(ctx context.Context, input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetSuite(ctx context.Context, input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSuiteWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetTest(ctx context.Context, input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTestWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetTestGridProject(ctx context.Context, input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetTestGridSession(ctx context.Context, input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTestGridSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetUpload(ctx context.Context, input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) GetVPCEConfiguration(ctx context.Context, input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVPCEConfigurationWithContext(ctx, input)
}

func (a *DeviceFarmActivities) InstallToRemoteAccessSession(ctx context.Context, input *devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InstallToRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListArtifacts(ctx context.Context, input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListArtifactsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListDeviceInstances(ctx context.Context, input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeviceInstancesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListDevicePools(ctx context.Context, input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDevicePoolsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListDevices(ctx context.Context, input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDevicesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListInstanceProfiles(ctx context.Context, input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInstanceProfilesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListJobs(ctx context.Context, input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListJobsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListNetworkProfiles(ctx context.Context, input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNetworkProfilesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListOfferingPromotions(ctx context.Context, input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOfferingPromotionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListOfferingTransactions(ctx context.Context, input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOfferingTransactionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListOfferings(ctx context.Context, input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOfferingsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListProjects(ctx context.Context, input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProjectsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListRemoteAccessSessions(ctx context.Context, input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRemoteAccessSessionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListRuns(ctx context.Context, input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRunsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListSamples(ctx context.Context, input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSamplesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListSuites(ctx context.Context, input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSuitesWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTagsForResource(ctx context.Context, input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridProjects(ctx context.Context, input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTestGridProjectsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridSessionActions(ctx context.Context, input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTestGridSessionActionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridSessionArtifacts(ctx context.Context, input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTestGridSessionArtifactsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTestGridSessions(ctx context.Context, input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTestGridSessionsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListTests(ctx context.Context, input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTestsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListUniqueProblems(ctx context.Context, input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUniqueProblemsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListUploads(ctx context.Context, input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUploadsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ListVPCEConfigurations(ctx context.Context, input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVPCEConfigurationsWithContext(ctx, input)
}

func (a *DeviceFarmActivities) PurchaseOffering(ctx context.Context, input *devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PurchaseOfferingWithContext(ctx, input)
}

func (a *DeviceFarmActivities) RenewOffering(ctx context.Context, input *devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RenewOfferingWithContext(ctx, input)
}

func (a *DeviceFarmActivities) ScheduleRun(ctx context.Context, input *devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ScheduleRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) StopJob(ctx context.Context, input *devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopJobWithContext(ctx, input)
}

func (a *DeviceFarmActivities) StopRemoteAccessSession(ctx context.Context, input *devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopRemoteAccessSessionWithContext(ctx, input)
}

func (a *DeviceFarmActivities) StopRun(ctx context.Context, input *devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopRunWithContext(ctx, input)
}

func (a *DeviceFarmActivities) TagResource(ctx context.Context, input *devicefarm.TagResourceInput) (*devicefarm.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UntagResource(ctx context.Context, input *devicefarm.UntagResourceInput) (*devicefarm.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateDeviceInstance(ctx context.Context, input *devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDeviceInstanceWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateDevicePool(ctx context.Context, input *devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDevicePoolWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateInstanceProfile(ctx context.Context, input *devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateInstanceProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateNetworkProfile(ctx context.Context, input *devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNetworkProfileWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateProject(ctx context.Context, input *devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateTestGridProject(ctx context.Context, input *devicefarm.UpdateTestGridProjectInput) (*devicefarm.UpdateTestGridProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTestGridProjectWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateUpload(ctx context.Context, input *devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUploadWithContext(ctx, input)
}

func (a *DeviceFarmActivities) UpdateVPCEConfiguration(ctx context.Context, input *devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVPCEConfigurationWithContext(ctx, input)
}
