package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworks/opsworksiface"
)

type OpsWorksActivities struct {
	client opsworksiface.OpsWorksAPI
}

func NewOpsWorksActivities(client opsworksiface.OpsWorksAPI) *OpsWorksActivities {
    return &OpsWorksActivities{client: client}
}

func (a *OpsWorksActivities) AssignInstance(input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
    return a.client.AssignInstance(input)
}

func (a *OpsWorksActivities) AssignVolume(input *opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error) {
    return a.client.AssignVolume(input)
}

func (a *OpsWorksActivities) AssociateElasticIp(input *opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error) {
    return a.client.AssociateElasticIp(input)
}

func (a *OpsWorksActivities) AttachElasticLoadBalancer(input *opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error) {
    return a.client.AttachElasticLoadBalancer(input)
}

func (a *OpsWorksActivities) CloneStack(input *opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error) {
    return a.client.CloneStack(input)
}

func (a *OpsWorksActivities) CreateApp(input *opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error) {
    return a.client.CreateApp(input)
}

func (a *OpsWorksActivities) CreateDeployment(input *opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error) {
    return a.client.CreateDeployment(input)
}

func (a *OpsWorksActivities) CreateInstance(input *opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error) {
    return a.client.CreateInstance(input)
}

func (a *OpsWorksActivities) CreateLayer(input *opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error) {
    return a.client.CreateLayer(input)
}

func (a *OpsWorksActivities) CreateStack(input *opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error) {
    return a.client.CreateStack(input)
}

func (a *OpsWorksActivities) CreateUserProfile(input *opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error) {
    return a.client.CreateUserProfile(input)
}

func (a *OpsWorksActivities) DeleteApp(input *opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error) {
    return a.client.DeleteApp(input)
}

func (a *OpsWorksActivities) DeleteInstance(input *opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error) {
    return a.client.DeleteInstance(input)
}

func (a *OpsWorksActivities) DeleteLayer(input *opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error) {
    return a.client.DeleteLayer(input)
}

func (a *OpsWorksActivities) DeleteStack(input *opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error) {
    return a.client.DeleteStack(input)
}

func (a *OpsWorksActivities) DeleteUserProfile(input *opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error) {
    return a.client.DeleteUserProfile(input)
}

func (a *OpsWorksActivities) DeregisterEcsCluster(input *opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error) {
    return a.client.DeregisterEcsCluster(input)
}

func (a *OpsWorksActivities) DeregisterElasticIp(input *opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error) {
    return a.client.DeregisterElasticIp(input)
}

func (a *OpsWorksActivities) DeregisterInstance(input *opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error) {
    return a.client.DeregisterInstance(input)
}

func (a *OpsWorksActivities) DeregisterRdsDbInstance(input *opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error) {
    return a.client.DeregisterRdsDbInstance(input)
}

func (a *OpsWorksActivities) DeregisterVolume(input *opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error) {
    return a.client.DeregisterVolume(input)
}

func (a *OpsWorksActivities) DescribeAgentVersions(input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error) {
    return a.client.DescribeAgentVersions(input)
}

func (a *OpsWorksActivities) DescribeApps(input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error) {
    return a.client.DescribeApps(input)
}

func (a *OpsWorksActivities) DescribeCommands(input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error) {
    return a.client.DescribeCommands(input)
}

func (a *OpsWorksActivities) DescribeDeployments(input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error) {
    return a.client.DescribeDeployments(input)
}

func (a *OpsWorksActivities) DescribeEcsClusters(input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error) {
    return a.client.DescribeEcsClusters(input)
}

func (a *OpsWorksActivities) DescribeElasticIps(input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error) {
    return a.client.DescribeElasticIps(input)
}

func (a *OpsWorksActivities) DescribeElasticLoadBalancers(input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error) {
    return a.client.DescribeElasticLoadBalancers(input)
}

func (a *OpsWorksActivities) DescribeInstances(input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error) {
    return a.client.DescribeInstances(input)
}

func (a *OpsWorksActivities) DescribeLayers(input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error) {
    return a.client.DescribeLayers(input)
}

func (a *OpsWorksActivities) DescribeLoadBasedAutoScaling(input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error) {
    return a.client.DescribeLoadBasedAutoScaling(input)
}

func (a *OpsWorksActivities) DescribeMyUserProfile(input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error) {
    return a.client.DescribeMyUserProfile(input)
}

func (a *OpsWorksActivities) DescribeOperatingSystems(input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error) {
    return a.client.DescribeOperatingSystems(input)
}

func (a *OpsWorksActivities) DescribePermissions(input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error) {
    return a.client.DescribePermissions(input)
}

func (a *OpsWorksActivities) DescribeRaidArrays(input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error) {
    return a.client.DescribeRaidArrays(input)
}

func (a *OpsWorksActivities) DescribeRdsDbInstances(input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error) {
    return a.client.DescribeRdsDbInstances(input)
}

func (a *OpsWorksActivities) DescribeServiceErrors(input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error) {
    return a.client.DescribeServiceErrors(input)
}

func (a *OpsWorksActivities) DescribeStackProvisioningParameters(input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error) {
    return a.client.DescribeStackProvisioningParameters(input)
}

func (a *OpsWorksActivities) DescribeStackSummary(input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error) {
    return a.client.DescribeStackSummary(input)
}

func (a *OpsWorksActivities) DescribeStacks(input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error) {
    return a.client.DescribeStacks(input)
}

func (a *OpsWorksActivities) DescribeTimeBasedAutoScaling(input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error) {
    return a.client.DescribeTimeBasedAutoScaling(input)
}

func (a *OpsWorksActivities) DescribeUserProfiles(input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error) {
    return a.client.DescribeUserProfiles(input)
}

func (a *OpsWorksActivities) DescribeVolumes(input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error) {
    return a.client.DescribeVolumes(input)
}

func (a *OpsWorksActivities) DetachElasticLoadBalancer(input *opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error) {
    return a.client.DetachElasticLoadBalancer(input)
}

func (a *OpsWorksActivities) DisassociateElasticIp(input *opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error) {
    return a.client.DisassociateElasticIp(input)
}

func (a *OpsWorksActivities) GetHostnameSuggestion(input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error) {
    return a.client.GetHostnameSuggestion(input)
}

func (a *OpsWorksActivities) GrantAccess(input *opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error) {
    return a.client.GrantAccess(input)
}

func (a *OpsWorksActivities) ListTags(input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *OpsWorksActivities) RebootInstance(input *opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error) {
    return a.client.RebootInstance(input)
}

func (a *OpsWorksActivities) RegisterEcsCluster(input *opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error) {
    return a.client.RegisterEcsCluster(input)
}

func (a *OpsWorksActivities) RegisterElasticIp(input *opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error) {
    return a.client.RegisterElasticIp(input)
}

func (a *OpsWorksActivities) RegisterInstance(input *opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error) {
    return a.client.RegisterInstance(input)
}

func (a *OpsWorksActivities) RegisterRdsDbInstance(input *opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error) {
    return a.client.RegisterRdsDbInstance(input)
}

func (a *OpsWorksActivities) RegisterVolume(input *opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error) {
    return a.client.RegisterVolume(input)
}

func (a *OpsWorksActivities) SetLoadBasedAutoScaling(input *opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error) {
    return a.client.SetLoadBasedAutoScaling(input)
}

func (a *OpsWorksActivities) SetPermission(input *opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error) {
    return a.client.SetPermission(input)
}

func (a *OpsWorksActivities) SetTimeBasedAutoScaling(input *opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error) {
    return a.client.SetTimeBasedAutoScaling(input)
}

func (a *OpsWorksActivities) StartInstance(input *opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error) {
    return a.client.StartInstance(input)
}

func (a *OpsWorksActivities) StartStack(input *opsworks.StartStackInput) (*opsworks.StartStackOutput, error) {
    return a.client.StartStack(input)
}

func (a *OpsWorksActivities) StopInstance(input *opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error) {
    return a.client.StopInstance(input)
}

func (a *OpsWorksActivities) StopStack(input *opsworks.StopStackInput) (*opsworks.StopStackOutput, error) {
    return a.client.StopStack(input)
}

func (a *OpsWorksActivities) TagResource(input *opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *OpsWorksActivities) UnassignInstance(input *opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error) {
    return a.client.UnassignInstance(input)
}

func (a *OpsWorksActivities) UnassignVolume(input *opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error) {
    return a.client.UnassignVolume(input)
}

func (a *OpsWorksActivities) UntagResource(input *opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *OpsWorksActivities) UpdateApp(input *opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error) {
    return a.client.UpdateApp(input)
}

func (a *OpsWorksActivities) UpdateElasticIp(input *opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error) {
    return a.client.UpdateElasticIp(input)
}

func (a *OpsWorksActivities) UpdateInstance(input *opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error) {
    return a.client.UpdateInstance(input)
}

func (a *OpsWorksActivities) UpdateLayer(input *opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error) {
    return a.client.UpdateLayer(input)
}

func (a *OpsWorksActivities) UpdateMyUserProfile(input *opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error) {
    return a.client.UpdateMyUserProfile(input)
}

func (a *OpsWorksActivities) UpdateRdsDbInstance(input *opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error) {
    return a.client.UpdateRdsDbInstance(input)
}

func (a *OpsWorksActivities) UpdateStack(input *opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error) {
    return a.client.UpdateStack(input)
}

func (a *OpsWorksActivities) UpdateUserProfile(input *opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error) {
    return a.client.UpdateUserProfile(input)
}

func (a *OpsWorksActivities) UpdateVolume(input *opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error) {
    return a.client.UpdateVolume(input)
}

func (a *OpsWorksActivities) WaitUntilAppExists(input *opsworks.WaitUntilAppExistsInput) (*opsworks.WaitUntilAppExistsOutput, error) {
    return a.client.WaitUntilAppExists(input)
}

func (a *OpsWorksActivities) WaitUntilDeploymentSuccessful(input *opsworks.WaitUntilDeploymentSuccessfulInput) (*opsworks.WaitUntilDeploymentSuccessfulOutput, error) {
    return a.client.WaitUntilDeploymentSuccessful(input)
}

func (a *OpsWorksActivities) WaitUntilInstanceOnline(input *opsworks.WaitUntilInstanceOnlineInput) (*opsworks.WaitUntilInstanceOnlineOutput, error) {
    return a.client.WaitUntilInstanceOnline(input)
}

func (a *OpsWorksActivities) WaitUntilInstanceRegistered(input *opsworks.WaitUntilInstanceRegisteredInput) (*opsworks.WaitUntilInstanceRegisteredOutput, error) {
    return a.client.WaitUntilInstanceRegistered(input)
}

func (a *OpsWorksActivities) WaitUntilInstanceStopped(input *opsworks.WaitUntilInstanceStoppedInput) (*opsworks.WaitUntilInstanceStoppedOutput, error) {
    return a.client.WaitUntilInstanceStopped(input)
}

func (a *OpsWorksActivities) WaitUntilInstanceTerminated(input *opsworks.WaitUntilInstanceTerminatedInput) (*opsworks.WaitUntilInstanceTerminatedOutput, error) {
    return a.client.WaitUntilInstanceTerminated(input)
}
