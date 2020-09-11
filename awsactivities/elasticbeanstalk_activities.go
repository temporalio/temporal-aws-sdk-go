
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ElasticBeanstalkActivities struct {
    client elasticbeanstalkiface.ElasticBeanstalkAPI
}

func NewElasticBeanstalkActivities(session *session.Session, config ...*aws.Config) *ElasticBeanstalkActivities {
    client := elasticbeanstalk.New(session, config...)
    return &ElasticBeanstalkActivities{client: client}
}

func (a *ElasticBeanstalkActivities) AbortEnvironmentUpdate(ctx context.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
    return a.client.AbortEnvironmentUpdateWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ApplyEnvironmentManagedAction(ctx context.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error) {
    return a.client.ApplyEnvironmentManagedActionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) AssociateEnvironmentOperationsRole(ctx context.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error) {
    return a.client.AssociateEnvironmentOperationsRoleWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CheckDNSAvailability(ctx context.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error) {
    return a.client.CheckDNSAvailabilityWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ComposeEnvironments(ctx context.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    return a.client.ComposeEnvironmentsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CreateApplication(ctx context.Context, input *elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CreateApplicationVersion(ctx context.Context, input *elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    return a.client.CreateApplicationVersionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CreateConfigurationTemplate(ctx context.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    return a.client.CreateConfigurationTemplateWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CreateEnvironment(ctx context.Context, input *elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    return a.client.CreateEnvironmentWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CreatePlatformVersion(ctx context.Context, input *elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error) {
    return a.client.CreatePlatformVersionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) CreateStorageLocation(ctx context.Context, input *elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error) {
    return a.client.CreateStorageLocationWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DeleteApplication(ctx context.Context, input *elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error) {
    return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DeleteApplicationVersion(ctx context.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error) {
    return a.client.DeleteApplicationVersionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DeleteConfigurationTemplate(ctx context.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error) {
    return a.client.DeleteConfigurationTemplateWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DeleteEnvironmentConfiguration(ctx context.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error) {
    return a.client.DeleteEnvironmentConfigurationWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DeletePlatformVersion(ctx context.Context, input *elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error) {
    return a.client.DeletePlatformVersionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeAccountAttributes(ctx context.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
    return a.client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeApplicationVersions(ctx context.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
    return a.client.DescribeApplicationVersionsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeApplications(ctx context.Context, input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
    return a.client.DescribeApplicationsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeConfigurationOptions(ctx context.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
    return a.client.DescribeConfigurationOptionsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeConfigurationSettings(ctx context.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
    return a.client.DescribeConfigurationSettingsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentHealth(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
    return a.client.DescribeEnvironmentHealthWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentManagedActionHistory(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
    return a.client.DescribeEnvironmentManagedActionHistoryWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentManagedActions(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
    return a.client.DescribeEnvironmentManagedActionsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironmentResources(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
    return a.client.DescribeEnvironmentResourcesWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeEnvironments(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    return a.client.DescribeEnvironmentsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeEvents(ctx context.Context, input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error) {
    return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribeInstancesHealth(ctx context.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
    return a.client.DescribeInstancesHealthWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DescribePlatformVersion(ctx context.Context, input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
    return a.client.DescribePlatformVersionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) DisassociateEnvironmentOperationsRole(ctx context.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error) {
    return a.client.DisassociateEnvironmentOperationsRoleWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ListAvailableSolutionStacks(ctx context.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
    return a.client.ListAvailableSolutionStacksWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ListPlatformBranches(ctx context.Context, input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
    return a.client.ListPlatformBranchesWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ListPlatformVersions(ctx context.Context, input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
    return a.client.ListPlatformVersionsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ListTagsForResource(ctx context.Context, input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) RebuildEnvironment(ctx context.Context, input *elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error) {
    return a.client.RebuildEnvironmentWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) RequestEnvironmentInfo(ctx context.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error) {
    return a.client.RequestEnvironmentInfoWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) RestartAppServer(ctx context.Context, input *elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error) {
    return a.client.RestartAppServerWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) RetrieveEnvironmentInfo(ctx context.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error) {
    return a.client.RetrieveEnvironmentInfoWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) SwapEnvironmentCNAMEs(ctx context.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error) {
    return a.client.SwapEnvironmentCNAMEsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) TerminateEnvironment(ctx context.Context, input *elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    return a.client.TerminateEnvironmentWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) UpdateApplication(ctx context.Context, input *elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    return a.client.UpdateApplicationWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) UpdateApplicationResourceLifecycle(ctx context.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error) {
    return a.client.UpdateApplicationResourceLifecycleWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) UpdateApplicationVersion(ctx context.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    return a.client.UpdateApplicationVersionWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) UpdateConfigurationTemplate(ctx context.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    return a.client.UpdateConfigurationTemplateWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) UpdateEnvironment(ctx context.Context, input *elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    return a.client.UpdateEnvironmentWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) UpdateTagsForResource(ctx context.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) (*elasticbeanstalk.UpdateTagsForResourceOutput, error) {
    return a.client.UpdateTagsForResourceWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) ValidateConfigurationSettings(ctx context.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error) {
    return a.client.ValidateConfigurationSettingsWithContext(ctx, input)
}

func (a *ElasticBeanstalkActivities) WaitUntilEnvironmentExists(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilEnvironmentExistsWithContext(ctx, input, options...)
	})
}

func (a *ElasticBeanstalkActivities) WaitUntilEnvironmentTerminated(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilEnvironmentTerminatedWithContext(ctx, input, options...)
	})
}

func (a *ElasticBeanstalkActivities) WaitUntilEnvironmentUpdated(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilEnvironmentUpdatedWithContext(ctx, input, options...)
	})
}
