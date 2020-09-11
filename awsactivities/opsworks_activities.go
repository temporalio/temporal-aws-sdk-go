package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworks/opsworksiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type OpsWorksActivities struct {
	client opsworksiface.OpsWorksAPI
}

func NewOpsWorksActivities(session *session.Session, config ...*aws.Config) *OpsWorksActivities {
	client := opsworks.New(session, config...)
	return &OpsWorksActivities{client: client}
}

func (a *OpsWorksActivities) AssignInstance(ctx context.Context, input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
	return a.client.AssignInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) AssignVolume(ctx context.Context, input *opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error) {
	return a.client.AssignVolumeWithContext(ctx, input)
}

func (a *OpsWorksActivities) AssociateElasticIp(ctx context.Context, input *opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error) {
	return a.client.AssociateElasticIpWithContext(ctx, input)
}

func (a *OpsWorksActivities) AttachElasticLoadBalancer(ctx context.Context, input *opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error) {
	return a.client.AttachElasticLoadBalancerWithContext(ctx, input)
}

func (a *OpsWorksActivities) CloneStack(ctx context.Context, input *opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error) {
	return a.client.CloneStackWithContext(ctx, input)
}

func (a *OpsWorksActivities) CreateApp(ctx context.Context, input *opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error) {
	return a.client.CreateAppWithContext(ctx, input)
}

func (a *OpsWorksActivities) CreateDeployment(ctx context.Context, input *opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error) {
	return a.client.CreateDeploymentWithContext(ctx, input)
}

func (a *OpsWorksActivities) CreateInstance(ctx context.Context, input *opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error) {
	return a.client.CreateInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) CreateLayer(ctx context.Context, input *opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error) {
	return a.client.CreateLayerWithContext(ctx, input)
}

func (a *OpsWorksActivities) CreateStack(ctx context.Context, input *opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error) {
	return a.client.CreateStackWithContext(ctx, input)
}

func (a *OpsWorksActivities) CreateUserProfile(ctx context.Context, input *opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error) {
	return a.client.CreateUserProfileWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeleteApp(ctx context.Context, input *opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error) {
	return a.client.DeleteAppWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeleteInstance(ctx context.Context, input *opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error) {
	return a.client.DeleteInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeleteLayer(ctx context.Context, input *opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error) {
	return a.client.DeleteLayerWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeleteStack(ctx context.Context, input *opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error) {
	return a.client.DeleteStackWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeleteUserProfile(ctx context.Context, input *opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error) {
	return a.client.DeleteUserProfileWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeregisterEcsCluster(ctx context.Context, input *opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error) {
	return a.client.DeregisterEcsClusterWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeregisterElasticIp(ctx context.Context, input *opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error) {
	return a.client.DeregisterElasticIpWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeregisterInstance(ctx context.Context, input *opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error) {
	return a.client.DeregisterInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeregisterRdsDbInstance(ctx context.Context, input *opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error) {
	return a.client.DeregisterRdsDbInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) DeregisterVolume(ctx context.Context, input *opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error) {
	return a.client.DeregisterVolumeWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeAgentVersions(ctx context.Context, input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error) {
	return a.client.DescribeAgentVersionsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeApps(ctx context.Context, input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error) {
	return a.client.DescribeAppsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeCommands(ctx context.Context, input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error) {
	return a.client.DescribeCommandsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeDeployments(ctx context.Context, input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error) {
	return a.client.DescribeDeploymentsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeEcsClusters(ctx context.Context, input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error) {
	return a.client.DescribeEcsClustersWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeElasticIps(ctx context.Context, input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error) {
	return a.client.DescribeElasticIpsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeElasticLoadBalancers(ctx context.Context, input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error) {
	return a.client.DescribeElasticLoadBalancersWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeInstances(ctx context.Context, input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error) {
	return a.client.DescribeInstancesWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeLayers(ctx context.Context, input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error) {
	return a.client.DescribeLayersWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeLoadBasedAutoScaling(ctx context.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error) {
	return a.client.DescribeLoadBasedAutoScalingWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeMyUserProfile(ctx context.Context, input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error) {
	return a.client.DescribeMyUserProfileWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeOperatingSystems(ctx context.Context, input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error) {
	return a.client.DescribeOperatingSystemsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribePermissions(ctx context.Context, input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error) {
	return a.client.DescribePermissionsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeRaidArrays(ctx context.Context, input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error) {
	return a.client.DescribeRaidArraysWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeRdsDbInstances(ctx context.Context, input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error) {
	return a.client.DescribeRdsDbInstancesWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeServiceErrors(ctx context.Context, input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error) {
	return a.client.DescribeServiceErrorsWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeStackProvisioningParameters(ctx context.Context, input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error) {
	return a.client.DescribeStackProvisioningParametersWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeStackSummary(ctx context.Context, input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error) {
	return a.client.DescribeStackSummaryWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeStacks(ctx context.Context, input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error) {
	return a.client.DescribeStacksWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeTimeBasedAutoScaling(ctx context.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error) {
	return a.client.DescribeTimeBasedAutoScalingWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeUserProfiles(ctx context.Context, input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error) {
	return a.client.DescribeUserProfilesWithContext(ctx, input)
}

func (a *OpsWorksActivities) DescribeVolumes(ctx context.Context, input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error) {
	return a.client.DescribeVolumesWithContext(ctx, input)
}

func (a *OpsWorksActivities) DetachElasticLoadBalancer(ctx context.Context, input *opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error) {
	return a.client.DetachElasticLoadBalancerWithContext(ctx, input)
}

func (a *OpsWorksActivities) DisassociateElasticIp(ctx context.Context, input *opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error) {
	return a.client.DisassociateElasticIpWithContext(ctx, input)
}

func (a *OpsWorksActivities) GetHostnameSuggestion(ctx context.Context, input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error) {
	return a.client.GetHostnameSuggestionWithContext(ctx, input)
}

func (a *OpsWorksActivities) GrantAccess(ctx context.Context, input *opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error) {
	return a.client.GrantAccessWithContext(ctx, input)
}

func (a *OpsWorksActivities) ListTags(ctx context.Context, input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *OpsWorksActivities) RebootInstance(ctx context.Context, input *opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error) {
	return a.client.RebootInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) RegisterEcsCluster(ctx context.Context, input *opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error) {
	return a.client.RegisterEcsClusterWithContext(ctx, input)
}

func (a *OpsWorksActivities) RegisterElasticIp(ctx context.Context, input *opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error) {
	return a.client.RegisterElasticIpWithContext(ctx, input)
}

func (a *OpsWorksActivities) RegisterInstance(ctx context.Context, input *opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error) {
	return a.client.RegisterInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) RegisterRdsDbInstance(ctx context.Context, input *opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error) {
	return a.client.RegisterRdsDbInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) RegisterVolume(ctx context.Context, input *opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error) {
	return a.client.RegisterVolumeWithContext(ctx, input)
}

func (a *OpsWorksActivities) SetLoadBasedAutoScaling(ctx context.Context, input *opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error) {
	return a.client.SetLoadBasedAutoScalingWithContext(ctx, input)
}

func (a *OpsWorksActivities) SetPermission(ctx context.Context, input *opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error) {
	return a.client.SetPermissionWithContext(ctx, input)
}

func (a *OpsWorksActivities) SetTimeBasedAutoScaling(ctx context.Context, input *opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error) {
	return a.client.SetTimeBasedAutoScalingWithContext(ctx, input)
}

func (a *OpsWorksActivities) StartInstance(ctx context.Context, input *opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error) {
	return a.client.StartInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) StartStack(ctx context.Context, input *opsworks.StartStackInput) (*opsworks.StartStackOutput, error) {
	return a.client.StartStackWithContext(ctx, input)
}

func (a *OpsWorksActivities) StopInstance(ctx context.Context, input *opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error) {
	return a.client.StopInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) StopStack(ctx context.Context, input *opsworks.StopStackInput) (*opsworks.StopStackOutput, error) {
	return a.client.StopStackWithContext(ctx, input)
}

func (a *OpsWorksActivities) TagResource(ctx context.Context, input *opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *OpsWorksActivities) UnassignInstance(ctx context.Context, input *opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error) {
	return a.client.UnassignInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) UnassignVolume(ctx context.Context, input *opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error) {
	return a.client.UnassignVolumeWithContext(ctx, input)
}

func (a *OpsWorksActivities) UntagResource(ctx context.Context, input *opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateApp(ctx context.Context, input *opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error) {
	return a.client.UpdateAppWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateElasticIp(ctx context.Context, input *opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error) {
	return a.client.UpdateElasticIpWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateInstance(ctx context.Context, input *opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error) {
	return a.client.UpdateInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateLayer(ctx context.Context, input *opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error) {
	return a.client.UpdateLayerWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateMyUserProfile(ctx context.Context, input *opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error) {
	return a.client.UpdateMyUserProfileWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateRdsDbInstance(ctx context.Context, input *opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error) {
	return a.client.UpdateRdsDbInstanceWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateStack(ctx context.Context, input *opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error) {
	return a.client.UpdateStackWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateUserProfile(ctx context.Context, input *opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error) {
	return a.client.UpdateUserProfileWithContext(ctx, input)
}

func (a *OpsWorksActivities) UpdateVolume(ctx context.Context, input *opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error) {
	return a.client.UpdateVolumeWithContext(ctx, input)
}

func (a *OpsWorksActivities) WaitUntilAppExists(ctx context.Context, input *opsworks.DescribeAppsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilAppExistsWithContext(ctx, input, options...)
	})
}

func (a *OpsWorksActivities) WaitUntilDeploymentSuccessful(ctx context.Context, input *opsworks.DescribeDeploymentsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDeploymentSuccessfulWithContext(ctx, input, options...)
	})
}

func (a *OpsWorksActivities) WaitUntilInstanceOnline(ctx context.Context, input *opsworks.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceOnlineWithContext(ctx, input, options...)
	})
}

func (a *OpsWorksActivities) WaitUntilInstanceRegistered(ctx context.Context, input *opsworks.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceRegisteredWithContext(ctx, input, options...)
	})
}

func (a *OpsWorksActivities) WaitUntilInstanceStopped(ctx context.Context, input *opsworks.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceStoppedWithContext(ctx, input, options...)
	})
}

func (a *OpsWorksActivities) WaitUntilInstanceTerminated(ctx context.Context, input *opsworks.DescribeInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceTerminatedWithContext(ctx, input, options...)
	})
}
