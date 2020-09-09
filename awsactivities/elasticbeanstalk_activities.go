
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
)

type ElasticBeanstalkActivities struct {
    client elasticbeanstalkiface.ElasticBeanstalkAPI
}

func NewElasticBeanstalkActivities(session *session.Session, config ...*aws.Config) *ElasticBeanstalkActivities {
    client := elasticbeanstalk.New(session, config...)
    return &ElasticBeanstalkActivities{client: client}
}

func (a *ElasticBeanstalkActivities) AbortEnvironmentUpdate(input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
    return a.client.AbortEnvironmentUpdate(input)
}

func (a *ElasticBeanstalkActivities) ApplyEnvironmentManagedAction(input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error) {
    return a.client.ApplyEnvironmentManagedAction(input)
}

func (a *ElasticBeanstalkActivities) AssociateEnvironmentOperationsRole(input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error) {
    return a.client.AssociateEnvironmentOperationsRole(input)
}

func (a *ElasticBeanstalkActivities) CheckDNSAvailability(input *elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error) {
    return a.client.CheckDNSAvailability(input)
}

func (a *ElasticBeanstalkActivities) ComposeEnvironments(input *elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    return a.client.ComposeEnvironments(input)
}

func (a *ElasticBeanstalkActivities) CreateApplication(input *elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    return a.client.CreateApplication(input)
}

func (a *ElasticBeanstalkActivities) CreateApplicationVersion(input *elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    return a.client.CreateApplicationVersion(input)
}

func (a *ElasticBeanstalkActivities) CreateConfigurationTemplate(input *elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    return a.client.CreateConfigurationTemplate(input)
}

func (a *ElasticBeanstalkActivities) CreateEnvironment(input *elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    return a.client.CreateEnvironment(input)
}

func (a *ElasticBeanstalkActivities) CreatePlatformVersion(input *elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error) {
    return a.client.CreatePlatformVersion(input)
}

func (a *ElasticBeanstalkActivities) CreateStorageLocation(input *elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error) {
    return a.client.CreateStorageLocation(input)
}

func (a *ElasticBeanstalkActivities) DeleteApplication(input *elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error) {
    return a.client.DeleteApplication(input)
}

func (a *ElasticBeanstalkActivities) DeleteApplicationVersion(input *elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error) {
    return a.client.DeleteApplicationVersion(input)
}

func (a *ElasticBeanstalkActivities) DeleteConfigurationTemplate(input *elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error) {
    return a.client.DeleteConfigurationTemplate(input)
}

func (a *ElasticBeanstalkActivities) DeleteEnvironmentConfiguration(input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error) {
    return a.client.DeleteEnvironmentConfiguration(input)
}

func (a *ElasticBeanstalkActivities) DeletePlatformVersion(input *elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error) {
    return a.client.DeletePlatformVersion(input)
}

func (a *ElasticBeanstalkActivities) DescribeAccountAttributes(input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributes(input)
}

func (a *ElasticBeanstalkActivities) DescribeApplicationVersions(input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
    return a.client.DescribeApplicationVersions(input)
}

func (a *ElasticBeanstalkActivities) DescribeApplications(input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
    return a.client.DescribeApplications(input)
}

func (a *ElasticBeanstalkActivities) DescribeConfigurationOptions(input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
    return a.client.DescribeConfigurationOptions(input)
}

func (a *ElasticBeanstalkActivities) DescribeConfigurationSettings(input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
    return a.client.DescribeConfigurationSettings(input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentHealth(input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
    return a.client.DescribeEnvironmentHealth(input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentManagedActionHistory(input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
    return a.client.DescribeEnvironmentManagedActionHistory(input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentManagedActions(input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
    return a.client.DescribeEnvironmentManagedActions(input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentResources(input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
    return a.client.DescribeEnvironmentResources(input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironments(input *elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    return a.client.DescribeEnvironments(input)
}

func (a *ElasticBeanstalkActivities) DescribeEvents(input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}

func (a *ElasticBeanstalkActivities) DescribeInstancesHealth(input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
    return a.client.DescribeInstancesHealth(input)
}

func (a *ElasticBeanstalkActivities) DescribePlatformVersion(input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
    return a.client.DescribePlatformVersion(input)
}

func (a *ElasticBeanstalkActivities) DisassociateEnvironmentOperationsRole(input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error) {
    return a.client.DisassociateEnvironmentOperationsRole(input)
}

func (a *ElasticBeanstalkActivities) ListAvailableSolutionStacks(input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
    return a.client.ListAvailableSolutionStacks(input)
}

func (a *ElasticBeanstalkActivities) ListPlatformBranches(input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
    return a.client.ListPlatformBranches(input)
}

func (a *ElasticBeanstalkActivities) ListPlatformVersions(input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
    return a.client.ListPlatformVersions(input)
}

func (a *ElasticBeanstalkActivities) ListTagsForResource(input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ElasticBeanstalkActivities) RebuildEnvironment(input *elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error) {
    return a.client.RebuildEnvironment(input)
}

func (a *ElasticBeanstalkActivities) RequestEnvironmentInfo(input *elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error) {
    return a.client.RequestEnvironmentInfo(input)
}

func (a *ElasticBeanstalkActivities) RestartAppServer(input *elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error) {
    return a.client.RestartAppServer(input)
}

func (a *ElasticBeanstalkActivities) RetrieveEnvironmentInfo(input *elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error) {
    return a.client.RetrieveEnvironmentInfo(input)
}

func (a *ElasticBeanstalkActivities) SwapEnvironmentCNAMEs(input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error) {
    return a.client.SwapEnvironmentCNAMEs(input)
}

func (a *ElasticBeanstalkActivities) TerminateEnvironment(input *elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    return a.client.TerminateEnvironment(input)
}

func (a *ElasticBeanstalkActivities) UpdateApplication(input *elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    return a.client.UpdateApplication(input)
}

func (a *ElasticBeanstalkActivities) UpdateApplicationResourceLifecycle(input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error) {
    return a.client.UpdateApplicationResourceLifecycle(input)
}

func (a *ElasticBeanstalkActivities) UpdateApplicationVersion(input *elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    return a.client.UpdateApplicationVersion(input)
}

func (a *ElasticBeanstalkActivities) UpdateConfigurationTemplate(input *elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    return a.client.UpdateConfigurationTemplate(input)
}

func (a *ElasticBeanstalkActivities) UpdateEnvironment(input *elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    return a.client.UpdateEnvironment(input)
}

func (a *ElasticBeanstalkActivities) UpdateTagsForResource(input *elasticbeanstalk.UpdateTagsForResourceInput) (*elasticbeanstalk.UpdateTagsForResourceOutput, error) {
    return a.client.UpdateTagsForResource(input)
}

func (a *ElasticBeanstalkActivities) ValidateConfigurationSettings(input *elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error) {
    return a.client.ValidateConfigurationSettings(input)
}

func (a *ElasticBeanstalkActivities) WaitUntilEnvironmentExists(input *elasticbeanstalk.DescribeEnvironmentsInput) error {
    return a.client.WaitUntilEnvironmentExists(input)
}

func (a *ElasticBeanstalkActivities) WaitUntilEnvironmentTerminated(input *elasticbeanstalk.DescribeEnvironmentsInput) error {
    return a.client.WaitUntilEnvironmentTerminated(input)
}

func (a *ElasticBeanstalkActivities) WaitUntilEnvironmentUpdated(input *elasticbeanstalk.DescribeEnvironmentsInput) error {
    return a.client.WaitUntilEnvironmentUpdated(input)
}
