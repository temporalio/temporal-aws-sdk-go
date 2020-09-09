package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
)

type DeviceFarmActivities struct {
	client devicefarmiface.DeviceFarmAPI
}

func NewDeviceFarmActivities(session *session.Session, config ...*aws.Config) *DeviceFarmActivities {
	client := devicefarm.New(session, config...)
	return &DeviceFarmActivities{client: client}
}

func (a *DeviceFarmActivities) CreateDevicePool(input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
	return a.client.CreateDevicePool(input)
}

func (a *DeviceFarmActivities) CreateInstanceProfile(input *devicefarm.CreateInstanceProfileInput) (*devicefarm.CreateInstanceProfileOutput, error) {
	return a.client.CreateInstanceProfile(input)
}

func (a *DeviceFarmActivities) CreateNetworkProfile(input *devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error) {
	return a.client.CreateNetworkProfile(input)
}

func (a *DeviceFarmActivities) CreateProject(input *devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error) {
	return a.client.CreateProject(input)
}

func (a *DeviceFarmActivities) CreateRemoteAccessSession(input *devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error) {
	return a.client.CreateRemoteAccessSession(input)
}

func (a *DeviceFarmActivities) CreateTestGridProject(input *devicefarm.CreateTestGridProjectInput) (*devicefarm.CreateTestGridProjectOutput, error) {
	return a.client.CreateTestGridProject(input)
}

func (a *DeviceFarmActivities) CreateTestGridUrl(input *devicefarm.CreateTestGridUrlInput) (*devicefarm.CreateTestGridUrlOutput, error) {
	return a.client.CreateTestGridUrl(input)
}

func (a *DeviceFarmActivities) CreateUpload(input *devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error) {
	return a.client.CreateUpload(input)
}

func (a *DeviceFarmActivities) CreateVPCEConfiguration(input *devicefarm.CreateVPCEConfigurationInput) (*devicefarm.CreateVPCEConfigurationOutput, error) {
	return a.client.CreateVPCEConfiguration(input)
}

func (a *DeviceFarmActivities) DeleteDevicePool(input *devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error) {
	return a.client.DeleteDevicePool(input)
}

func (a *DeviceFarmActivities) DeleteInstanceProfile(input *devicefarm.DeleteInstanceProfileInput) (*devicefarm.DeleteInstanceProfileOutput, error) {
	return a.client.DeleteInstanceProfile(input)
}

func (a *DeviceFarmActivities) DeleteNetworkProfile(input *devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error) {
	return a.client.DeleteNetworkProfile(input)
}

func (a *DeviceFarmActivities) DeleteProject(input *devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error) {
	return a.client.DeleteProject(input)
}

func (a *DeviceFarmActivities) DeleteRemoteAccessSession(input *devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error) {
	return a.client.DeleteRemoteAccessSession(input)
}

func (a *DeviceFarmActivities) DeleteRun(input *devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error) {
	return a.client.DeleteRun(input)
}

func (a *DeviceFarmActivities) DeleteTestGridProject(input *devicefarm.DeleteTestGridProjectInput) (*devicefarm.DeleteTestGridProjectOutput, error) {
	return a.client.DeleteTestGridProject(input)
}

func (a *DeviceFarmActivities) DeleteUpload(input *devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error) {
	return a.client.DeleteUpload(input)
}

func (a *DeviceFarmActivities) DeleteVPCEConfiguration(input *devicefarm.DeleteVPCEConfigurationInput) (*devicefarm.DeleteVPCEConfigurationOutput, error) {
	return a.client.DeleteVPCEConfiguration(input)
}

func (a *DeviceFarmActivities) GetAccountSettings(input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	return a.client.GetAccountSettings(input)
}

func (a *DeviceFarmActivities) GetDevice(input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	return a.client.GetDevice(input)
}

func (a *DeviceFarmActivities) GetDeviceInstance(input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error) {
	return a.client.GetDeviceInstance(input)
}

func (a *DeviceFarmActivities) GetDevicePool(input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	return a.client.GetDevicePool(input)
}

func (a *DeviceFarmActivities) GetDevicePoolCompatibility(input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	return a.client.GetDevicePoolCompatibility(input)
}

func (a *DeviceFarmActivities) GetInstanceProfile(input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error) {
	return a.client.GetInstanceProfile(input)
}

func (a *DeviceFarmActivities) GetJob(input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	return a.client.GetJob(input)
}

func (a *DeviceFarmActivities) GetNetworkProfile(input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
	return a.client.GetNetworkProfile(input)
}

func (a *DeviceFarmActivities) GetOfferingStatus(input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
	return a.client.GetOfferingStatus(input)
}

func (a *DeviceFarmActivities) GetProject(input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	return a.client.GetProject(input)
}

func (a *DeviceFarmActivities) GetRemoteAccessSession(input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	return a.client.GetRemoteAccessSession(input)
}

func (a *DeviceFarmActivities) GetRun(input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
	return a.client.GetRun(input)
}

func (a *DeviceFarmActivities) GetSuite(input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	return a.client.GetSuite(input)
}

func (a *DeviceFarmActivities) GetTest(input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	return a.client.GetTest(input)
}

func (a *DeviceFarmActivities) GetTestGridProject(input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error) {
	return a.client.GetTestGridProject(input)
}

func (a *DeviceFarmActivities) GetTestGridSession(input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error) {
	return a.client.GetTestGridSession(input)
}

func (a *DeviceFarmActivities) GetUpload(input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
	return a.client.GetUpload(input)
}

func (a *DeviceFarmActivities) GetVPCEConfiguration(input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error) {
	return a.client.GetVPCEConfiguration(input)
}

func (a *DeviceFarmActivities) InstallToRemoteAccessSession(input *devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error) {
	return a.client.InstallToRemoteAccessSession(input)
}

func (a *DeviceFarmActivities) ListArtifacts(input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
	return a.client.ListArtifacts(input)
}

func (a *DeviceFarmActivities) ListDeviceInstances(input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error) {
	return a.client.ListDeviceInstances(input)
}

func (a *DeviceFarmActivities) ListDevicePools(input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
	return a.client.ListDevicePools(input)
}

func (a *DeviceFarmActivities) ListDevices(input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	return a.client.ListDevices(input)
}

func (a *DeviceFarmActivities) ListInstanceProfiles(input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error) {
	return a.client.ListInstanceProfiles(input)
}

func (a *DeviceFarmActivities) ListJobs(input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	return a.client.ListJobs(input)
}

func (a *DeviceFarmActivities) ListNetworkProfiles(input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
	return a.client.ListNetworkProfiles(input)
}

func (a *DeviceFarmActivities) ListOfferingPromotions(input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
	return a.client.ListOfferingPromotions(input)
}

func (a *DeviceFarmActivities) ListOfferingTransactions(input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
	return a.client.ListOfferingTransactions(input)
}

func (a *DeviceFarmActivities) ListOfferings(input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
	return a.client.ListOfferings(input)
}

func (a *DeviceFarmActivities) ListProjects(input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
	return a.client.ListProjects(input)
}

func (a *DeviceFarmActivities) ListRemoteAccessSessions(input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	return a.client.ListRemoteAccessSessions(input)
}

func (a *DeviceFarmActivities) ListRuns(input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	return a.client.ListRuns(input)
}

func (a *DeviceFarmActivities) ListSamples(input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	return a.client.ListSamples(input)
}

func (a *DeviceFarmActivities) ListSuites(input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	return a.client.ListSuites(input)
}

func (a *DeviceFarmActivities) ListTagsForResource(input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *DeviceFarmActivities) ListTestGridProjects(input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error) {
	return a.client.ListTestGridProjects(input)
}

func (a *DeviceFarmActivities) ListTestGridSessionActions(input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error) {
	return a.client.ListTestGridSessionActions(input)
}

func (a *DeviceFarmActivities) ListTestGridSessionArtifacts(input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
	return a.client.ListTestGridSessionArtifacts(input)
}

func (a *DeviceFarmActivities) ListTestGridSessions(input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error) {
	return a.client.ListTestGridSessions(input)
}

func (a *DeviceFarmActivities) ListTests(input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	return a.client.ListTests(input)
}

func (a *DeviceFarmActivities) ListUniqueProblems(input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	return a.client.ListUniqueProblems(input)
}

func (a *DeviceFarmActivities) ListUploads(input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	return a.client.ListUploads(input)
}

func (a *DeviceFarmActivities) ListVPCEConfigurations(input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error) {
	return a.client.ListVPCEConfigurations(input)
}

func (a *DeviceFarmActivities) PurchaseOffering(input *devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error) {
	return a.client.PurchaseOffering(input)
}

func (a *DeviceFarmActivities) RenewOffering(input *devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error) {
	return a.client.RenewOffering(input)
}

func (a *DeviceFarmActivities) ScheduleRun(input *devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error) {
	return a.client.ScheduleRun(input)
}

func (a *DeviceFarmActivities) StopJob(input *devicefarm.StopJobInput) (*devicefarm.StopJobOutput, error) {
	return a.client.StopJob(input)
}

func (a *DeviceFarmActivities) StopRemoteAccessSession(input *devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error) {
	return a.client.StopRemoteAccessSession(input)
}

func (a *DeviceFarmActivities) StopRun(input *devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error) {
	return a.client.StopRun(input)
}

func (a *DeviceFarmActivities) TagResource(input *devicefarm.TagResourceInput) (*devicefarm.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *DeviceFarmActivities) UntagResource(input *devicefarm.UntagResourceInput) (*devicefarm.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *DeviceFarmActivities) UpdateDeviceInstance(input *devicefarm.UpdateDeviceInstanceInput) (*devicefarm.UpdateDeviceInstanceOutput, error) {
	return a.client.UpdateDeviceInstance(input)
}

func (a *DeviceFarmActivities) UpdateDevicePool(input *devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error) {
	return a.client.UpdateDevicePool(input)
}

func (a *DeviceFarmActivities) UpdateInstanceProfile(input *devicefarm.UpdateInstanceProfileInput) (*devicefarm.UpdateInstanceProfileOutput, error) {
	return a.client.UpdateInstanceProfile(input)
}

func (a *DeviceFarmActivities) UpdateNetworkProfile(input *devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error) {
	return a.client.UpdateNetworkProfile(input)
}

func (a *DeviceFarmActivities) UpdateProject(input *devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error) {
	return a.client.UpdateProject(input)
}

func (a *DeviceFarmActivities) UpdateTestGridProject(input *devicefarm.UpdateTestGridProjectInput) (*devicefarm.UpdateTestGridProjectOutput, error) {
	return a.client.UpdateTestGridProject(input)
}

func (a *DeviceFarmActivities) UpdateUpload(input *devicefarm.UpdateUploadInput) (*devicefarm.UpdateUploadOutput, error) {
	return a.client.UpdateUpload(input)
}

func (a *DeviceFarmActivities) UpdateVPCEConfiguration(input *devicefarm.UpdateVPCEConfigurationInput) (*devicefarm.UpdateVPCEConfigurationOutput, error) {
	return a.client.UpdateVPCEConfiguration(input)
}
