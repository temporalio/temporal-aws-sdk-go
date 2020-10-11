// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package opsworksstub

import (
	"github.com/aws/aws-sdk-go/service/opsworks"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	AssignInstance(ctx workflow.Context, input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)
	AssignInstanceAsync(ctx workflow.Context, input *opsworks.AssignInstanceInput) *OpsWorksAssignInstanceFuture

	AssignVolume(ctx workflow.Context, input *opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)
	AssignVolumeAsync(ctx workflow.Context, input *opsworks.AssignVolumeInput) *OpsWorksAssignVolumeFuture

	AssociateElasticIp(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)
	AssociateElasticIpAsync(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) *OpsWorksAssociateElasticIpFuture

	AttachElasticLoadBalancer(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)
	AttachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) *OpsWorksAttachElasticLoadBalancerFuture

	CloneStack(ctx workflow.Context, input *opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)
	CloneStackAsync(ctx workflow.Context, input *opsworks.CloneStackInput) *OpsWorksCloneStackFuture

	CreateApp(ctx workflow.Context, input *opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)
	CreateAppAsync(ctx workflow.Context, input *opsworks.CreateAppInput) *OpsWorksCreateAppFuture

	CreateDeployment(ctx workflow.Context, input *opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *opsworks.CreateDeploymentInput) *OpsWorksCreateDeploymentFuture

	CreateInstance(ctx workflow.Context, input *opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)
	CreateInstanceAsync(ctx workflow.Context, input *opsworks.CreateInstanceInput) *OpsWorksCreateInstanceFuture

	CreateLayer(ctx workflow.Context, input *opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)
	CreateLayerAsync(ctx workflow.Context, input *opsworks.CreateLayerInput) *OpsWorksCreateLayerFuture

	CreateStack(ctx workflow.Context, input *opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)
	CreateStackAsync(ctx workflow.Context, input *opsworks.CreateStackInput) *OpsWorksCreateStackFuture

	CreateUserProfile(ctx workflow.Context, input *opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)
	CreateUserProfileAsync(ctx workflow.Context, input *opsworks.CreateUserProfileInput) *OpsWorksCreateUserProfileFuture

	DeleteApp(ctx workflow.Context, input *opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)
	DeleteAppAsync(ctx workflow.Context, input *opsworks.DeleteAppInput) *OpsWorksDeleteAppFuture

	DeleteInstance(ctx workflow.Context, input *opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)
	DeleteInstanceAsync(ctx workflow.Context, input *opsworks.DeleteInstanceInput) *OpsWorksDeleteInstanceFuture

	DeleteLayer(ctx workflow.Context, input *opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)
	DeleteLayerAsync(ctx workflow.Context, input *opsworks.DeleteLayerInput) *OpsWorksDeleteLayerFuture

	DeleteStack(ctx workflow.Context, input *opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)
	DeleteStackAsync(ctx workflow.Context, input *opsworks.DeleteStackInput) *OpsWorksDeleteStackFuture

	DeleteUserProfile(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)
	DeleteUserProfileAsync(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) *OpsWorksDeleteUserProfileFuture

	DeregisterEcsCluster(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)
	DeregisterEcsClusterAsync(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) *OpsWorksDeregisterEcsClusterFuture

	DeregisterElasticIp(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)
	DeregisterElasticIpAsync(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) *OpsWorksDeregisterElasticIpFuture

	DeregisterInstance(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)
	DeregisterInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) *OpsWorksDeregisterInstanceFuture

	DeregisterRdsDbInstance(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)
	DeregisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) *OpsWorksDeregisterRdsDbInstanceFuture

	DeregisterVolume(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)
	DeregisterVolumeAsync(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) *OpsWorksDeregisterVolumeFuture

	DescribeAgentVersions(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)
	DescribeAgentVersionsAsync(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) *OpsWorksDescribeAgentVersionsFuture

	DescribeApps(ctx workflow.Context, input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)
	DescribeAppsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) *OpsWorksDescribeAppsFuture

	DescribeCommands(ctx workflow.Context, input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)
	DescribeCommandsAsync(ctx workflow.Context, input *opsworks.DescribeCommandsInput) *OpsWorksDescribeCommandsFuture

	DescribeDeployments(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)
	DescribeDeploymentsAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) *OpsWorksDescribeDeploymentsFuture

	DescribeEcsClusters(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)
	DescribeEcsClustersAsync(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) *OpsWorksDescribeEcsClustersFuture

	DescribeElasticIps(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)
	DescribeElasticIpsAsync(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) *OpsWorksDescribeElasticIpsFuture

	DescribeElasticLoadBalancers(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)
	DescribeElasticLoadBalancersAsync(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) *OpsWorksDescribeElasticLoadBalancersFuture

	DescribeInstances(ctx workflow.Context, input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)
	DescribeInstancesAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *OpsWorksDescribeInstancesFuture

	DescribeLayers(ctx workflow.Context, input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)
	DescribeLayersAsync(ctx workflow.Context, input *opsworks.DescribeLayersInput) *OpsWorksDescribeLayersFuture

	DescribeLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
	DescribeLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) *OpsWorksDescribeLoadBasedAutoScalingFuture

	DescribeMyUserProfile(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)
	DescribeMyUserProfileAsync(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) *OpsWorksDescribeMyUserProfileFuture

	DescribeOperatingSystems(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error)
	DescribeOperatingSystemsAsync(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) *OpsWorksDescribeOperatingSystemsFuture

	DescribePermissions(ctx workflow.Context, input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)
	DescribePermissionsAsync(ctx workflow.Context, input *opsworks.DescribePermissionsInput) *OpsWorksDescribePermissionsFuture

	DescribeRaidArrays(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)
	DescribeRaidArraysAsync(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) *OpsWorksDescribeRaidArraysFuture

	DescribeRdsDbInstances(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)
	DescribeRdsDbInstancesAsync(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) *OpsWorksDescribeRdsDbInstancesFuture

	DescribeServiceErrors(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)
	DescribeServiceErrorsAsync(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) *OpsWorksDescribeServiceErrorsFuture

	DescribeStackProvisioningParameters(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)
	DescribeStackProvisioningParametersAsync(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) *OpsWorksDescribeStackProvisioningParametersFuture

	DescribeStackSummary(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)
	DescribeStackSummaryAsync(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) *OpsWorksDescribeStackSummaryFuture

	DescribeStacks(ctx workflow.Context, input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)
	DescribeStacksAsync(ctx workflow.Context, input *opsworks.DescribeStacksInput) *OpsWorksDescribeStacksFuture

	DescribeTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
	DescribeTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) *OpsWorksDescribeTimeBasedAutoScalingFuture

	DescribeUserProfiles(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)
	DescribeUserProfilesAsync(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) *OpsWorksDescribeUserProfilesFuture

	DescribeVolumes(ctx workflow.Context, input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)
	DescribeVolumesAsync(ctx workflow.Context, input *opsworks.DescribeVolumesInput) *OpsWorksDescribeVolumesFuture

	DetachElasticLoadBalancer(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)
	DetachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) *OpsWorksDetachElasticLoadBalancerFuture

	DisassociateElasticIp(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)
	DisassociateElasticIpAsync(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) *OpsWorksDisassociateElasticIpFuture

	GetHostnameSuggestion(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)
	GetHostnameSuggestionAsync(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) *OpsWorksGetHostnameSuggestionFuture

	GrantAccess(ctx workflow.Context, input *opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)
	GrantAccessAsync(ctx workflow.Context, input *opsworks.GrantAccessInput) *OpsWorksGrantAccessFuture

	ListTags(ctx workflow.Context, input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *opsworks.ListTagsInput) *OpsWorksListTagsFuture

	RebootInstance(ctx workflow.Context, input *opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)
	RebootInstanceAsync(ctx workflow.Context, input *opsworks.RebootInstanceInput) *OpsWorksRebootInstanceFuture

	RegisterEcsCluster(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)
	RegisterEcsClusterAsync(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) *OpsWorksRegisterEcsClusterFuture

	RegisterElasticIp(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)
	RegisterElasticIpAsync(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) *OpsWorksRegisterElasticIpFuture

	RegisterInstance(ctx workflow.Context, input *opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)
	RegisterInstanceAsync(ctx workflow.Context, input *opsworks.RegisterInstanceInput) *OpsWorksRegisterInstanceFuture

	RegisterRdsDbInstance(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)
	RegisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) *OpsWorksRegisterRdsDbInstanceFuture

	RegisterVolume(ctx workflow.Context, input *opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)
	RegisterVolumeAsync(ctx workflow.Context, input *opsworks.RegisterVolumeInput) *OpsWorksRegisterVolumeFuture

	SetLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)
	SetLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) *OpsWorksSetLoadBasedAutoScalingFuture

	SetPermission(ctx workflow.Context, input *opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)
	SetPermissionAsync(ctx workflow.Context, input *opsworks.SetPermissionInput) *OpsWorksSetPermissionFuture

	SetTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)
	SetTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) *OpsWorksSetTimeBasedAutoScalingFuture

	StartInstance(ctx workflow.Context, input *opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)
	StartInstanceAsync(ctx workflow.Context, input *opsworks.StartInstanceInput) *OpsWorksStartInstanceFuture

	StartStack(ctx workflow.Context, input *opsworks.StartStackInput) (*opsworks.StartStackOutput, error)
	StartStackAsync(ctx workflow.Context, input *opsworks.StartStackInput) *OpsWorksStartStackFuture

	StopInstance(ctx workflow.Context, input *opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)
	StopInstanceAsync(ctx workflow.Context, input *opsworks.StopInstanceInput) *OpsWorksStopInstanceFuture

	StopStack(ctx workflow.Context, input *opsworks.StopStackInput) (*opsworks.StopStackOutput, error)
	StopStackAsync(ctx workflow.Context, input *opsworks.StopStackInput) *OpsWorksStopStackFuture

	TagResource(ctx workflow.Context, input *opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *opsworks.TagResourceInput) *OpsWorksTagResourceFuture

	UnassignInstance(ctx workflow.Context, input *opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)
	UnassignInstanceAsync(ctx workflow.Context, input *opsworks.UnassignInstanceInput) *OpsWorksUnassignInstanceFuture

	UnassignVolume(ctx workflow.Context, input *opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)
	UnassignVolumeAsync(ctx workflow.Context, input *opsworks.UnassignVolumeInput) *OpsWorksUnassignVolumeFuture

	UntagResource(ctx workflow.Context, input *opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *opsworks.UntagResourceInput) *OpsWorksUntagResourceFuture

	UpdateApp(ctx workflow.Context, input *opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)
	UpdateAppAsync(ctx workflow.Context, input *opsworks.UpdateAppInput) *OpsWorksUpdateAppFuture

	UpdateElasticIp(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)
	UpdateElasticIpAsync(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) *OpsWorksUpdateElasticIpFuture

	UpdateInstance(ctx workflow.Context, input *opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)
	UpdateInstanceAsync(ctx workflow.Context, input *opsworks.UpdateInstanceInput) *OpsWorksUpdateInstanceFuture

	UpdateLayer(ctx workflow.Context, input *opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)
	UpdateLayerAsync(ctx workflow.Context, input *opsworks.UpdateLayerInput) *OpsWorksUpdateLayerFuture

	UpdateMyUserProfile(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)
	UpdateMyUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) *OpsWorksUpdateMyUserProfileFuture

	UpdateRdsDbInstance(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)
	UpdateRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) *OpsWorksUpdateRdsDbInstanceFuture

	UpdateStack(ctx workflow.Context, input *opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)
	UpdateStackAsync(ctx workflow.Context, input *opsworks.UpdateStackInput) *OpsWorksUpdateStackFuture

	UpdateUserProfile(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)
	UpdateUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) *OpsWorksUpdateUserProfileFuture

	UpdateVolume(ctx workflow.Context, input *opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
	UpdateVolumeAsync(ctx workflow.Context, input *opsworks.UpdateVolumeInput) *OpsWorksUpdateVolumeFuture

	WaitUntilAppExists(ctx workflow.Context, input *opsworks.DescribeAppsInput) error
	WaitUntilAppExistsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) *awsclients.VoidFuture

	WaitUntilDeploymentSuccessful(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) error
	WaitUntilDeploymentSuccessfulAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) *awsclients.VoidFuture

	WaitUntilInstanceOnline(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceOnlineAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *awsclients.VoidFuture

	WaitUntilInstanceRegistered(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceRegisteredAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *awsclients.VoidFuture

	WaitUntilInstanceStopped(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceStoppedAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *awsclients.VoidFuture

	WaitUntilInstanceTerminated(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
	WaitUntilInstanceTerminatedAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
