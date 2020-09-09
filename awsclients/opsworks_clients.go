package awsclients

import (
	"github.com/aws/aws-sdk-go/service/opsworks"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type OpsWorksClient interface {
       AssignInstance(ctx workflow.Context, input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)
       AssignInstanceAsync(ctx workflow.Context, input *opsworks.AssignInstanceInput) *OpsworksAssignInstanceResult

       AssignVolume(ctx workflow.Context, input *opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)
       AssignVolumeAsync(ctx workflow.Context, input *opsworks.AssignVolumeInput) *OpsworksAssignVolumeResult

       AssociateElasticIp(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)
       AssociateElasticIpAsync(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) *OpsworksAssociateElasticIpResult

       AttachElasticLoadBalancer(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)
       AttachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) *OpsworksAttachElasticLoadBalancerResult

       CloneStack(ctx workflow.Context, input *opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)
       CloneStackAsync(ctx workflow.Context, input *opsworks.CloneStackInput) *OpsworksCloneStackResult

       CreateApp(ctx workflow.Context, input *opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)
       CreateAppAsync(ctx workflow.Context, input *opsworks.CreateAppInput) *OpsworksCreateAppResult

       CreateDeployment(ctx workflow.Context, input *opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)
       CreateDeploymentAsync(ctx workflow.Context, input *opsworks.CreateDeploymentInput) *OpsworksCreateDeploymentResult

       CreateInstance(ctx workflow.Context, input *opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)
       CreateInstanceAsync(ctx workflow.Context, input *opsworks.CreateInstanceInput) *OpsworksCreateInstanceResult

       CreateLayer(ctx workflow.Context, input *opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)
       CreateLayerAsync(ctx workflow.Context, input *opsworks.CreateLayerInput) *OpsworksCreateLayerResult

       CreateStack(ctx workflow.Context, input *opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)
       CreateStackAsync(ctx workflow.Context, input *opsworks.CreateStackInput) *OpsworksCreateStackResult

       CreateUserProfile(ctx workflow.Context, input *opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)
       CreateUserProfileAsync(ctx workflow.Context, input *opsworks.CreateUserProfileInput) *OpsworksCreateUserProfileResult

       DeleteApp(ctx workflow.Context, input *opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)
       DeleteAppAsync(ctx workflow.Context, input *opsworks.DeleteAppInput) *OpsworksDeleteAppResult

       DeleteInstance(ctx workflow.Context, input *opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)
       DeleteInstanceAsync(ctx workflow.Context, input *opsworks.DeleteInstanceInput) *OpsworksDeleteInstanceResult

       DeleteLayer(ctx workflow.Context, input *opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)
       DeleteLayerAsync(ctx workflow.Context, input *opsworks.DeleteLayerInput) *OpsworksDeleteLayerResult

       DeleteStack(ctx workflow.Context, input *opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)
       DeleteStackAsync(ctx workflow.Context, input *opsworks.DeleteStackInput) *OpsworksDeleteStackResult

       DeleteUserProfile(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)
       DeleteUserProfileAsync(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) *OpsworksDeleteUserProfileResult

       DeregisterEcsCluster(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)
       DeregisterEcsClusterAsync(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) *OpsworksDeregisterEcsClusterResult

       DeregisterElasticIp(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)
       DeregisterElasticIpAsync(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) *OpsworksDeregisterElasticIpResult

       DeregisterInstance(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)
       DeregisterInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) *OpsworksDeregisterInstanceResult

       DeregisterRdsDbInstance(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)
       DeregisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) *OpsworksDeregisterRdsDbInstanceResult

       DeregisterVolume(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)
       DeregisterVolumeAsync(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) *OpsworksDeregisterVolumeResult

       DescribeAgentVersions(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)
       DescribeAgentVersionsAsync(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) *OpsworksDescribeAgentVersionsResult

       DescribeApps(ctx workflow.Context, input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)
       DescribeAppsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) *OpsworksDescribeAppsResult

       DescribeCommands(ctx workflow.Context, input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)
       DescribeCommandsAsync(ctx workflow.Context, input *opsworks.DescribeCommandsInput) *OpsworksDescribeCommandsResult

       DescribeDeployments(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)
       DescribeDeploymentsAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) *OpsworksDescribeDeploymentsResult

       DescribeEcsClusters(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)
       DescribeEcsClustersAsync(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) *OpsworksDescribeEcsClustersResult

       DescribeElasticIps(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)
       DescribeElasticIpsAsync(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) *OpsworksDescribeElasticIpsResult

       DescribeElasticLoadBalancers(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)
       DescribeElasticLoadBalancersAsync(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) *OpsworksDescribeElasticLoadBalancersResult

       DescribeInstances(ctx workflow.Context, input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)
       DescribeInstancesAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *OpsworksDescribeInstancesResult

       DescribeLayers(ctx workflow.Context, input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)
       DescribeLayersAsync(ctx workflow.Context, input *opsworks.DescribeLayersInput) *OpsworksDescribeLayersResult

       DescribeLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)
       DescribeLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) *OpsworksDescribeLoadBasedAutoScalingResult

       DescribeMyUserProfile(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)
       DescribeMyUserProfileAsync(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) *OpsworksDescribeMyUserProfileResult

       DescribeOperatingSystems(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error)
       DescribeOperatingSystemsAsync(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) *OpsworksDescribeOperatingSystemsResult

       DescribePermissions(ctx workflow.Context, input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)
       DescribePermissionsAsync(ctx workflow.Context, input *opsworks.DescribePermissionsInput) *OpsworksDescribePermissionsResult

       DescribeRaidArrays(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)
       DescribeRaidArraysAsync(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) *OpsworksDescribeRaidArraysResult

       DescribeRdsDbInstances(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)
       DescribeRdsDbInstancesAsync(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) *OpsworksDescribeRdsDbInstancesResult

       DescribeServiceErrors(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)
       DescribeServiceErrorsAsync(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) *OpsworksDescribeServiceErrorsResult

       DescribeStackProvisioningParameters(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)
       DescribeStackProvisioningParametersAsync(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) *OpsworksDescribeStackProvisioningParametersResult

       DescribeStackSummary(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)
       DescribeStackSummaryAsync(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) *OpsworksDescribeStackSummaryResult

       DescribeStacks(ctx workflow.Context, input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)
       DescribeStacksAsync(ctx workflow.Context, input *opsworks.DescribeStacksInput) *OpsworksDescribeStacksResult

       DescribeTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)
       DescribeTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) *OpsworksDescribeTimeBasedAutoScalingResult

       DescribeUserProfiles(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)
       DescribeUserProfilesAsync(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) *OpsworksDescribeUserProfilesResult

       DescribeVolumes(ctx workflow.Context, input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)
       DescribeVolumesAsync(ctx workflow.Context, input *opsworks.DescribeVolumesInput) *OpsworksDescribeVolumesResult

       DetachElasticLoadBalancer(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)
       DetachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) *OpsworksDetachElasticLoadBalancerResult

       DisassociateElasticIp(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)
       DisassociateElasticIpAsync(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) *OpsworksDisassociateElasticIpResult

       GetHostnameSuggestion(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)
       GetHostnameSuggestionAsync(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) *OpsworksGetHostnameSuggestionResult

       GrantAccess(ctx workflow.Context, input *opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)
       GrantAccessAsync(ctx workflow.Context, input *opsworks.GrantAccessInput) *OpsworksGrantAccessResult

       ListTags(ctx workflow.Context, input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error)
       ListTagsAsync(ctx workflow.Context, input *opsworks.ListTagsInput) *OpsworksListTagsResult

       RebootInstance(ctx workflow.Context, input *opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)
       RebootInstanceAsync(ctx workflow.Context, input *opsworks.RebootInstanceInput) *OpsworksRebootInstanceResult

       RegisterEcsCluster(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)
       RegisterEcsClusterAsync(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) *OpsworksRegisterEcsClusterResult

       RegisterElasticIp(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)
       RegisterElasticIpAsync(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) *OpsworksRegisterElasticIpResult

       RegisterInstance(ctx workflow.Context, input *opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)
       RegisterInstanceAsync(ctx workflow.Context, input *opsworks.RegisterInstanceInput) *OpsworksRegisterInstanceResult

       RegisterRdsDbInstance(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)
       RegisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) *OpsworksRegisterRdsDbInstanceResult

       RegisterVolume(ctx workflow.Context, input *opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)
       RegisterVolumeAsync(ctx workflow.Context, input *opsworks.RegisterVolumeInput) *OpsworksRegisterVolumeResult

       SetLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)
       SetLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) *OpsworksSetLoadBasedAutoScalingResult

       SetPermission(ctx workflow.Context, input *opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)
       SetPermissionAsync(ctx workflow.Context, input *opsworks.SetPermissionInput) *OpsworksSetPermissionResult

       SetTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)
       SetTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) *OpsworksSetTimeBasedAutoScalingResult

       StartInstance(ctx workflow.Context, input *opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)
       StartInstanceAsync(ctx workflow.Context, input *opsworks.StartInstanceInput) *OpsworksStartInstanceResult

       StartStack(ctx workflow.Context, input *opsworks.StartStackInput) (*opsworks.StartStackOutput, error)
       StartStackAsync(ctx workflow.Context, input *opsworks.StartStackInput) *OpsworksStartStackResult

       StopInstance(ctx workflow.Context, input *opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)
       StopInstanceAsync(ctx workflow.Context, input *opsworks.StopInstanceInput) *OpsworksStopInstanceResult

       StopStack(ctx workflow.Context, input *opsworks.StopStackInput) (*opsworks.StopStackOutput, error)
       StopStackAsync(ctx workflow.Context, input *opsworks.StopStackInput) *OpsworksStopStackResult

       TagResource(ctx workflow.Context, input *opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *opsworks.TagResourceInput) *OpsworksTagResourceResult

       UnassignInstance(ctx workflow.Context, input *opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)
       UnassignInstanceAsync(ctx workflow.Context, input *opsworks.UnassignInstanceInput) *OpsworksUnassignInstanceResult

       UnassignVolume(ctx workflow.Context, input *opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)
       UnassignVolumeAsync(ctx workflow.Context, input *opsworks.UnassignVolumeInput) *OpsworksUnassignVolumeResult

       UntagResource(ctx workflow.Context, input *opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *opsworks.UntagResourceInput) *OpsworksUntagResourceResult

       UpdateApp(ctx workflow.Context, input *opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)
       UpdateAppAsync(ctx workflow.Context, input *opsworks.UpdateAppInput) *OpsworksUpdateAppResult

       UpdateElasticIp(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)
       UpdateElasticIpAsync(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) *OpsworksUpdateElasticIpResult

       UpdateInstance(ctx workflow.Context, input *opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)
       UpdateInstanceAsync(ctx workflow.Context, input *opsworks.UpdateInstanceInput) *OpsworksUpdateInstanceResult

       UpdateLayer(ctx workflow.Context, input *opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)
       UpdateLayerAsync(ctx workflow.Context, input *opsworks.UpdateLayerInput) *OpsworksUpdateLayerResult

       UpdateMyUserProfile(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)
       UpdateMyUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) *OpsworksUpdateMyUserProfileResult

       UpdateRdsDbInstance(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)
       UpdateRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) *OpsworksUpdateRdsDbInstanceResult

       UpdateStack(ctx workflow.Context, input *opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)
       UpdateStackAsync(ctx workflow.Context, input *opsworks.UpdateStackInput) *OpsworksUpdateStackResult

       UpdateUserProfile(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)
       UpdateUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) *OpsworksUpdateUserProfileResult

       UpdateVolume(ctx workflow.Context, input *opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
       UpdateVolumeAsync(ctx workflow.Context, input *opsworks.UpdateVolumeInput) *OpsworksUpdateVolumeResult

       WaitUntilAppExists(ctx workflow.Context, input *opsworks.DescribeAppsInput) error
       WaitUntilDeploymentSuccessful(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) error
       WaitUntilInstanceOnline(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
       WaitUntilInstanceRegistered(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
       WaitUntilInstanceStopped(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error
       WaitUntilInstanceTerminated(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error}

type OpsworksAssignInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksAssignInstanceResult) Get(ctx workflow.Context) (*opsworks.AssignInstanceOutput, error) {
    var output opsworks.AssignInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksAssignVolumeResult struct {
	Result workflow.Future
}

func (r *OpsworksAssignVolumeResult) Get(ctx workflow.Context) (*opsworks.AssignVolumeOutput, error) {
    var output opsworks.AssignVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksAssociateElasticIpResult struct {
	Result workflow.Future
}

func (r *OpsworksAssociateElasticIpResult) Get(ctx workflow.Context) (*opsworks.AssociateElasticIpOutput, error) {
    var output opsworks.AssociateElasticIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksAttachElasticLoadBalancerResult struct {
	Result workflow.Future
}

func (r *OpsworksAttachElasticLoadBalancerResult) Get(ctx workflow.Context) (*opsworks.AttachElasticLoadBalancerOutput, error) {
    var output opsworks.AttachElasticLoadBalancerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCloneStackResult struct {
	Result workflow.Future
}

func (r *OpsworksCloneStackResult) Get(ctx workflow.Context) (*opsworks.CloneStackOutput, error) {
    var output opsworks.CloneStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCreateAppResult struct {
	Result workflow.Future
}

func (r *OpsworksCreateAppResult) Get(ctx workflow.Context) (*opsworks.CreateAppOutput, error) {
    var output opsworks.CreateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCreateDeploymentResult struct {
	Result workflow.Future
}

func (r *OpsworksCreateDeploymentResult) Get(ctx workflow.Context) (*opsworks.CreateDeploymentOutput, error) {
    var output opsworks.CreateDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCreateInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksCreateInstanceResult) Get(ctx workflow.Context) (*opsworks.CreateInstanceOutput, error) {
    var output opsworks.CreateInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCreateLayerResult struct {
	Result workflow.Future
}

func (r *OpsworksCreateLayerResult) Get(ctx workflow.Context) (*opsworks.CreateLayerOutput, error) {
    var output opsworks.CreateLayerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCreateStackResult struct {
	Result workflow.Future
}

func (r *OpsworksCreateStackResult) Get(ctx workflow.Context) (*opsworks.CreateStackOutput, error) {
    var output opsworks.CreateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksCreateUserProfileResult struct {
	Result workflow.Future
}

func (r *OpsworksCreateUserProfileResult) Get(ctx workflow.Context) (*opsworks.CreateUserProfileOutput, error) {
    var output opsworks.CreateUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeleteAppResult struct {
	Result workflow.Future
}

func (r *OpsworksDeleteAppResult) Get(ctx workflow.Context) (*opsworks.DeleteAppOutput, error) {
    var output opsworks.DeleteAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeleteInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksDeleteInstanceResult) Get(ctx workflow.Context) (*opsworks.DeleteInstanceOutput, error) {
    var output opsworks.DeleteInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeleteLayerResult struct {
	Result workflow.Future
}

func (r *OpsworksDeleteLayerResult) Get(ctx workflow.Context) (*opsworks.DeleteLayerOutput, error) {
    var output opsworks.DeleteLayerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeleteStackResult struct {
	Result workflow.Future
}

func (r *OpsworksDeleteStackResult) Get(ctx workflow.Context) (*opsworks.DeleteStackOutput, error) {
    var output opsworks.DeleteStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeleteUserProfileResult struct {
	Result workflow.Future
}

func (r *OpsworksDeleteUserProfileResult) Get(ctx workflow.Context) (*opsworks.DeleteUserProfileOutput, error) {
    var output opsworks.DeleteUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeregisterEcsClusterResult struct {
	Result workflow.Future
}

func (r *OpsworksDeregisterEcsClusterResult) Get(ctx workflow.Context) (*opsworks.DeregisterEcsClusterOutput, error) {
    var output opsworks.DeregisterEcsClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeregisterElasticIpResult struct {
	Result workflow.Future
}

func (r *OpsworksDeregisterElasticIpResult) Get(ctx workflow.Context) (*opsworks.DeregisterElasticIpOutput, error) {
    var output opsworks.DeregisterElasticIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeregisterInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksDeregisterInstanceResult) Get(ctx workflow.Context) (*opsworks.DeregisterInstanceOutput, error) {
    var output opsworks.DeregisterInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeregisterRdsDbInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksDeregisterRdsDbInstanceResult) Get(ctx workflow.Context) (*opsworks.DeregisterRdsDbInstanceOutput, error) {
    var output opsworks.DeregisterRdsDbInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDeregisterVolumeResult struct {
	Result workflow.Future
}

func (r *OpsworksDeregisterVolumeResult) Get(ctx workflow.Context) (*opsworks.DeregisterVolumeOutput, error) {
    var output opsworks.DeregisterVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeAgentVersionsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeAgentVersionsResult) Get(ctx workflow.Context) (*opsworks.DescribeAgentVersionsOutput, error) {
    var output opsworks.DescribeAgentVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeAppsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeAppsResult) Get(ctx workflow.Context) (*opsworks.DescribeAppsOutput, error) {
    var output opsworks.DescribeAppsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeCommandsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeCommandsResult) Get(ctx workflow.Context) (*opsworks.DescribeCommandsOutput, error) {
    var output opsworks.DescribeCommandsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeDeploymentsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeDeploymentsResult) Get(ctx workflow.Context) (*opsworks.DescribeDeploymentsOutput, error) {
    var output opsworks.DescribeDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeEcsClustersResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeEcsClustersResult) Get(ctx workflow.Context) (*opsworks.DescribeEcsClustersOutput, error) {
    var output opsworks.DescribeEcsClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeElasticIpsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeElasticIpsResult) Get(ctx workflow.Context) (*opsworks.DescribeElasticIpsOutput, error) {
    var output opsworks.DescribeElasticIpsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeElasticLoadBalancersResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeElasticLoadBalancersResult) Get(ctx workflow.Context) (*opsworks.DescribeElasticLoadBalancersOutput, error) {
    var output opsworks.DescribeElasticLoadBalancersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeInstancesResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeInstancesResult) Get(ctx workflow.Context) (*opsworks.DescribeInstancesOutput, error) {
    var output opsworks.DescribeInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeLayersResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeLayersResult) Get(ctx workflow.Context) (*opsworks.DescribeLayersOutput, error) {
    var output opsworks.DescribeLayersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeLoadBasedAutoScalingResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeLoadBasedAutoScalingResult) Get(ctx workflow.Context) (*opsworks.DescribeLoadBasedAutoScalingOutput, error) {
    var output opsworks.DescribeLoadBasedAutoScalingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeMyUserProfileResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeMyUserProfileResult) Get(ctx workflow.Context) (*opsworks.DescribeMyUserProfileOutput, error) {
    var output opsworks.DescribeMyUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeOperatingSystemsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeOperatingSystemsResult) Get(ctx workflow.Context) (*opsworks.DescribeOperatingSystemsOutput, error) {
    var output opsworks.DescribeOperatingSystemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribePermissionsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribePermissionsResult) Get(ctx workflow.Context) (*opsworks.DescribePermissionsOutput, error) {
    var output opsworks.DescribePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeRaidArraysResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeRaidArraysResult) Get(ctx workflow.Context) (*opsworks.DescribeRaidArraysOutput, error) {
    var output opsworks.DescribeRaidArraysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeRdsDbInstancesResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeRdsDbInstancesResult) Get(ctx workflow.Context) (*opsworks.DescribeRdsDbInstancesOutput, error) {
    var output opsworks.DescribeRdsDbInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeServiceErrorsResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeServiceErrorsResult) Get(ctx workflow.Context) (*opsworks.DescribeServiceErrorsOutput, error) {
    var output opsworks.DescribeServiceErrorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeStackProvisioningParametersResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeStackProvisioningParametersResult) Get(ctx workflow.Context) (*opsworks.DescribeStackProvisioningParametersOutput, error) {
    var output opsworks.DescribeStackProvisioningParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeStackSummaryResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeStackSummaryResult) Get(ctx workflow.Context) (*opsworks.DescribeStackSummaryOutput, error) {
    var output opsworks.DescribeStackSummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeStacksResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeStacksResult) Get(ctx workflow.Context) (*opsworks.DescribeStacksOutput, error) {
    var output opsworks.DescribeStacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeTimeBasedAutoScalingResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeTimeBasedAutoScalingResult) Get(ctx workflow.Context) (*opsworks.DescribeTimeBasedAutoScalingOutput, error) {
    var output opsworks.DescribeTimeBasedAutoScalingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeUserProfilesResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeUserProfilesResult) Get(ctx workflow.Context) (*opsworks.DescribeUserProfilesOutput, error) {
    var output opsworks.DescribeUserProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDescribeVolumesResult struct {
	Result workflow.Future
}

func (r *OpsworksDescribeVolumesResult) Get(ctx workflow.Context) (*opsworks.DescribeVolumesOutput, error) {
    var output opsworks.DescribeVolumesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDetachElasticLoadBalancerResult struct {
	Result workflow.Future
}

func (r *OpsworksDetachElasticLoadBalancerResult) Get(ctx workflow.Context) (*opsworks.DetachElasticLoadBalancerOutput, error) {
    var output opsworks.DetachElasticLoadBalancerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksDisassociateElasticIpResult struct {
	Result workflow.Future
}

func (r *OpsworksDisassociateElasticIpResult) Get(ctx workflow.Context) (*opsworks.DisassociateElasticIpOutput, error) {
    var output opsworks.DisassociateElasticIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksGetHostnameSuggestionResult struct {
	Result workflow.Future
}

func (r *OpsworksGetHostnameSuggestionResult) Get(ctx workflow.Context) (*opsworks.GetHostnameSuggestionOutput, error) {
    var output opsworks.GetHostnameSuggestionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksGrantAccessResult struct {
	Result workflow.Future
}

func (r *OpsworksGrantAccessResult) Get(ctx workflow.Context) (*opsworks.GrantAccessOutput, error) {
    var output opsworks.GrantAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksListTagsResult struct {
	Result workflow.Future
}

func (r *OpsworksListTagsResult) Get(ctx workflow.Context) (*opsworks.ListTagsOutput, error) {
    var output opsworks.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksRebootInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksRebootInstanceResult) Get(ctx workflow.Context) (*opsworks.RebootInstanceOutput, error) {
    var output opsworks.RebootInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksRegisterEcsClusterResult struct {
	Result workflow.Future
}

func (r *OpsworksRegisterEcsClusterResult) Get(ctx workflow.Context) (*opsworks.RegisterEcsClusterOutput, error) {
    var output opsworks.RegisterEcsClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksRegisterElasticIpResult struct {
	Result workflow.Future
}

func (r *OpsworksRegisterElasticIpResult) Get(ctx workflow.Context) (*opsworks.RegisterElasticIpOutput, error) {
    var output opsworks.RegisterElasticIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksRegisterInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksRegisterInstanceResult) Get(ctx workflow.Context) (*opsworks.RegisterInstanceOutput, error) {
    var output opsworks.RegisterInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksRegisterRdsDbInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksRegisterRdsDbInstanceResult) Get(ctx workflow.Context) (*opsworks.RegisterRdsDbInstanceOutput, error) {
    var output opsworks.RegisterRdsDbInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksRegisterVolumeResult struct {
	Result workflow.Future
}

func (r *OpsworksRegisterVolumeResult) Get(ctx workflow.Context) (*opsworks.RegisterVolumeOutput, error) {
    var output opsworks.RegisterVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksSetLoadBasedAutoScalingResult struct {
	Result workflow.Future
}

func (r *OpsworksSetLoadBasedAutoScalingResult) Get(ctx workflow.Context) (*opsworks.SetLoadBasedAutoScalingOutput, error) {
    var output opsworks.SetLoadBasedAutoScalingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksSetPermissionResult struct {
	Result workflow.Future
}

func (r *OpsworksSetPermissionResult) Get(ctx workflow.Context) (*opsworks.SetPermissionOutput, error) {
    var output opsworks.SetPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksSetTimeBasedAutoScalingResult struct {
	Result workflow.Future
}

func (r *OpsworksSetTimeBasedAutoScalingResult) Get(ctx workflow.Context) (*opsworks.SetTimeBasedAutoScalingOutput, error) {
    var output opsworks.SetTimeBasedAutoScalingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksStartInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksStartInstanceResult) Get(ctx workflow.Context) (*opsworks.StartInstanceOutput, error) {
    var output opsworks.StartInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksStartStackResult struct {
	Result workflow.Future
}

func (r *OpsworksStartStackResult) Get(ctx workflow.Context) (*opsworks.StartStackOutput, error) {
    var output opsworks.StartStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksStopInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksStopInstanceResult) Get(ctx workflow.Context) (*opsworks.StopInstanceOutput, error) {
    var output opsworks.StopInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksStopStackResult struct {
	Result workflow.Future
}

func (r *OpsworksStopStackResult) Get(ctx workflow.Context) (*opsworks.StopStackOutput, error) {
    var output opsworks.StopStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksTagResourceResult struct {
	Result workflow.Future
}

func (r *OpsworksTagResourceResult) Get(ctx workflow.Context) (*opsworks.TagResourceOutput, error) {
    var output opsworks.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUnassignInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksUnassignInstanceResult) Get(ctx workflow.Context) (*opsworks.UnassignInstanceOutput, error) {
    var output opsworks.UnassignInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUnassignVolumeResult struct {
	Result workflow.Future
}

func (r *OpsworksUnassignVolumeResult) Get(ctx workflow.Context) (*opsworks.UnassignVolumeOutput, error) {
    var output opsworks.UnassignVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUntagResourceResult struct {
	Result workflow.Future
}

func (r *OpsworksUntagResourceResult) Get(ctx workflow.Context) (*opsworks.UntagResourceOutput, error) {
    var output opsworks.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateAppResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateAppResult) Get(ctx workflow.Context) (*opsworks.UpdateAppOutput, error) {
    var output opsworks.UpdateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateElasticIpResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateElasticIpResult) Get(ctx workflow.Context) (*opsworks.UpdateElasticIpOutput, error) {
    var output opsworks.UpdateElasticIpOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateInstanceResult) Get(ctx workflow.Context) (*opsworks.UpdateInstanceOutput, error) {
    var output opsworks.UpdateInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateLayerResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateLayerResult) Get(ctx workflow.Context) (*opsworks.UpdateLayerOutput, error) {
    var output opsworks.UpdateLayerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateMyUserProfileResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateMyUserProfileResult) Get(ctx workflow.Context) (*opsworks.UpdateMyUserProfileOutput, error) {
    var output opsworks.UpdateMyUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateRdsDbInstanceResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateRdsDbInstanceResult) Get(ctx workflow.Context) (*opsworks.UpdateRdsDbInstanceOutput, error) {
    var output opsworks.UpdateRdsDbInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateStackResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateStackResult) Get(ctx workflow.Context) (*opsworks.UpdateStackOutput, error) {
    var output opsworks.UpdateStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateUserProfileResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateUserProfileResult) Get(ctx workflow.Context) (*opsworks.UpdateUserProfileOutput, error) {
    var output opsworks.UpdateUserProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsworksUpdateVolumeResult struct {
	Result workflow.Future
}

func (r *OpsworksUpdateVolumeResult) Get(ctx workflow.Context) (*opsworks.UpdateVolumeOutput, error) {
    var output opsworks.UpdateVolumeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OpsWorksStub struct {
    activities awsactivities.OpsWorksActivities
}

func NewOpsWorksStub() OpsWorksClient {
    return &OpsWorksStub{}
}

func (a *OpsWorksStub) AssignInstance(ctx workflow.Context, input *opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error) {
    var output opsworks.AssignInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssignInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) AssignInstanceAsync(ctx workflow.Context, input *opsworks.AssignInstanceInput) *OpsworksAssignInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssignInstance, input)
    return &OpsworksAssignInstanceResult{Result: future}
}

func (a *OpsWorksStub) AssignVolume(ctx workflow.Context, input *opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error) {
    var output opsworks.AssignVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssignVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) AssignVolumeAsync(ctx workflow.Context, input *opsworks.AssignVolumeInput) *OpsworksAssignVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssignVolume, input)
    return &OpsworksAssignVolumeResult{Result: future}
}

func (a *OpsWorksStub) AssociateElasticIp(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error) {
    var output opsworks.AssociateElasticIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateElasticIp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) AssociateElasticIpAsync(ctx workflow.Context, input *opsworks.AssociateElasticIpInput) *OpsworksAssociateElasticIpResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateElasticIp, input)
    return &OpsworksAssociateElasticIpResult{Result: future}
}

func (a *OpsWorksStub) AttachElasticLoadBalancer(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error) {
    var output opsworks.AttachElasticLoadBalancerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachElasticLoadBalancer, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) AttachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.AttachElasticLoadBalancerInput) *OpsworksAttachElasticLoadBalancerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachElasticLoadBalancer, input)
    return &OpsworksAttachElasticLoadBalancerResult{Result: future}
}

func (a *OpsWorksStub) CloneStack(ctx workflow.Context, input *opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error) {
    var output opsworks.CloneStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CloneStack, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CloneStackAsync(ctx workflow.Context, input *opsworks.CloneStackInput) *OpsworksCloneStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CloneStack, input)
    return &OpsworksCloneStackResult{Result: future}
}

func (a *OpsWorksStub) CreateApp(ctx workflow.Context, input *opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error) {
    var output opsworks.CreateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CreateAppAsync(ctx workflow.Context, input *opsworks.CreateAppInput) *OpsworksCreateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input)
    return &OpsworksCreateAppResult{Result: future}
}

func (a *OpsWorksStub) CreateDeployment(ctx workflow.Context, input *opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error) {
    var output opsworks.CreateDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CreateDeploymentAsync(ctx workflow.Context, input *opsworks.CreateDeploymentInput) *OpsworksCreateDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input)
    return &OpsworksCreateDeploymentResult{Result: future}
}

func (a *OpsWorksStub) CreateInstance(ctx workflow.Context, input *opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error) {
    var output opsworks.CreateInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CreateInstanceAsync(ctx workflow.Context, input *opsworks.CreateInstanceInput) *OpsworksCreateInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateInstance, input)
    return &OpsworksCreateInstanceResult{Result: future}
}

func (a *OpsWorksStub) CreateLayer(ctx workflow.Context, input *opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error) {
    var output opsworks.CreateLayerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLayer, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CreateLayerAsync(ctx workflow.Context, input *opsworks.CreateLayerInput) *OpsworksCreateLayerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLayer, input)
    return &OpsworksCreateLayerResult{Result: future}
}

func (a *OpsWorksStub) CreateStack(ctx workflow.Context, input *opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error) {
    var output opsworks.CreateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CreateStackAsync(ctx workflow.Context, input *opsworks.CreateStackInput) *OpsworksCreateStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStack, input)
    return &OpsworksCreateStackResult{Result: future}
}

func (a *OpsWorksStub) CreateUserProfile(ctx workflow.Context, input *opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error) {
    var output opsworks.CreateUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) CreateUserProfileAsync(ctx workflow.Context, input *opsworks.CreateUserProfileInput) *OpsworksCreateUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUserProfile, input)
    return &OpsworksCreateUserProfileResult{Result: future}
}

func (a *OpsWorksStub) DeleteApp(ctx workflow.Context, input *opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error) {
    var output opsworks.DeleteAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeleteAppAsync(ctx workflow.Context, input *opsworks.DeleteAppInput) *OpsworksDeleteAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input)
    return &OpsworksDeleteAppResult{Result: future}
}

func (a *OpsWorksStub) DeleteInstance(ctx workflow.Context, input *opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error) {
    var output opsworks.DeleteInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeleteInstanceAsync(ctx workflow.Context, input *opsworks.DeleteInstanceInput) *OpsworksDeleteInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInstance, input)
    return &OpsworksDeleteInstanceResult{Result: future}
}

func (a *OpsWorksStub) DeleteLayer(ctx workflow.Context, input *opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error) {
    var output opsworks.DeleteLayerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLayer, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeleteLayerAsync(ctx workflow.Context, input *opsworks.DeleteLayerInput) *OpsworksDeleteLayerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLayer, input)
    return &OpsworksDeleteLayerResult{Result: future}
}

func (a *OpsWorksStub) DeleteStack(ctx workflow.Context, input *opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error) {
    var output opsworks.DeleteStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStack, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeleteStackAsync(ctx workflow.Context, input *opsworks.DeleteStackInput) *OpsworksDeleteStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteStack, input)
    return &OpsworksDeleteStackResult{Result: future}
}

func (a *OpsWorksStub) DeleteUserProfile(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error) {
    var output opsworks.DeleteUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeleteUserProfileAsync(ctx workflow.Context, input *opsworks.DeleteUserProfileInput) *OpsworksDeleteUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUserProfile, input)
    return &OpsworksDeleteUserProfileResult{Result: future}
}

func (a *OpsWorksStub) DeregisterEcsCluster(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error) {
    var output opsworks.DeregisterEcsClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterEcsCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeregisterEcsClusterAsync(ctx workflow.Context, input *opsworks.DeregisterEcsClusterInput) *OpsworksDeregisterEcsClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterEcsCluster, input)
    return &OpsworksDeregisterEcsClusterResult{Result: future}
}

func (a *OpsWorksStub) DeregisterElasticIp(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error) {
    var output opsworks.DeregisterElasticIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterElasticIp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeregisterElasticIpAsync(ctx workflow.Context, input *opsworks.DeregisterElasticIpInput) *OpsworksDeregisterElasticIpResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterElasticIp, input)
    return &OpsworksDeregisterElasticIpResult{Result: future}
}

func (a *OpsWorksStub) DeregisterInstance(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error) {
    var output opsworks.DeregisterInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeregisterInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterInstanceInput) *OpsworksDeregisterInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterInstance, input)
    return &OpsworksDeregisterInstanceResult{Result: future}
}

func (a *OpsWorksStub) DeregisterRdsDbInstance(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error) {
    var output opsworks.DeregisterRdsDbInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterRdsDbInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeregisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.DeregisterRdsDbInstanceInput) *OpsworksDeregisterRdsDbInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterRdsDbInstance, input)
    return &OpsworksDeregisterRdsDbInstanceResult{Result: future}
}

func (a *OpsWorksStub) DeregisterVolume(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error) {
    var output opsworks.DeregisterVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DeregisterVolumeAsync(ctx workflow.Context, input *opsworks.DeregisterVolumeInput) *OpsworksDeregisterVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterVolume, input)
    return &OpsworksDeregisterVolumeResult{Result: future}
}

func (a *OpsWorksStub) DescribeAgentVersions(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error) {
    var output opsworks.DescribeAgentVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAgentVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeAgentVersionsAsync(ctx workflow.Context, input *opsworks.DescribeAgentVersionsInput) *OpsworksDescribeAgentVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAgentVersions, input)
    return &OpsworksDescribeAgentVersionsResult{Result: future}
}

func (a *OpsWorksStub) DescribeApps(ctx workflow.Context, input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error) {
    var output opsworks.DescribeAppsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApps, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeAppsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) *OpsworksDescribeAppsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeApps, input)
    return &OpsworksDescribeAppsResult{Result: future}
}

func (a *OpsWorksStub) DescribeCommands(ctx workflow.Context, input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error) {
    var output opsworks.DescribeCommandsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCommands, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeCommandsAsync(ctx workflow.Context, input *opsworks.DescribeCommandsInput) *OpsworksDescribeCommandsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCommands, input)
    return &OpsworksDescribeCommandsResult{Result: future}
}

func (a *OpsWorksStub) DescribeDeployments(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error) {
    var output opsworks.DescribeDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeDeploymentsAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) *OpsworksDescribeDeploymentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDeployments, input)
    return &OpsworksDescribeDeploymentsResult{Result: future}
}

func (a *OpsWorksStub) DescribeEcsClusters(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error) {
    var output opsworks.DescribeEcsClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEcsClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeEcsClustersAsync(ctx workflow.Context, input *opsworks.DescribeEcsClustersInput) *OpsworksDescribeEcsClustersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEcsClusters, input)
    return &OpsworksDescribeEcsClustersResult{Result: future}
}

func (a *OpsWorksStub) DescribeElasticIps(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error) {
    var output opsworks.DescribeElasticIpsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticIps, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeElasticIpsAsync(ctx workflow.Context, input *opsworks.DescribeElasticIpsInput) *OpsworksDescribeElasticIpsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticIps, input)
    return &OpsworksDescribeElasticIpsResult{Result: future}
}

func (a *OpsWorksStub) DescribeElasticLoadBalancers(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error) {
    var output opsworks.DescribeElasticLoadBalancersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticLoadBalancers, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeElasticLoadBalancersAsync(ctx workflow.Context, input *opsworks.DescribeElasticLoadBalancersInput) *OpsworksDescribeElasticLoadBalancersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticLoadBalancers, input)
    return &OpsworksDescribeElasticLoadBalancersResult{Result: future}
}

func (a *OpsWorksStub) DescribeInstances(ctx workflow.Context, input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error) {
    var output opsworks.DescribeInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeInstancesAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) *OpsworksDescribeInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstances, input)
    return &OpsworksDescribeInstancesResult{Result: future}
}

func (a *OpsWorksStub) DescribeLayers(ctx workflow.Context, input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error) {
    var output opsworks.DescribeLayersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLayers, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeLayersAsync(ctx workflow.Context, input *opsworks.DescribeLayersInput) *OpsworksDescribeLayersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLayers, input)
    return &OpsworksDescribeLayersResult{Result: future}
}

func (a *OpsWorksStub) DescribeLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error) {
    var output opsworks.DescribeLoadBasedAutoScalingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoadBasedAutoScaling, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeLoadBasedAutoScalingInput) *OpsworksDescribeLoadBasedAutoScalingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoadBasedAutoScaling, input)
    return &OpsworksDescribeLoadBasedAutoScalingResult{Result: future}
}

func (a *OpsWorksStub) DescribeMyUserProfile(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error) {
    var output opsworks.DescribeMyUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMyUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeMyUserProfileAsync(ctx workflow.Context, input *opsworks.DescribeMyUserProfileInput) *OpsworksDescribeMyUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMyUserProfile, input)
    return &OpsworksDescribeMyUserProfileResult{Result: future}
}

func (a *OpsWorksStub) DescribeOperatingSystems(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error) {
    var output opsworks.DescribeOperatingSystemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOperatingSystems, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeOperatingSystemsAsync(ctx workflow.Context, input *opsworks.DescribeOperatingSystemsInput) *OpsworksDescribeOperatingSystemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOperatingSystems, input)
    return &OpsworksDescribeOperatingSystemsResult{Result: future}
}

func (a *OpsWorksStub) DescribePermissions(ctx workflow.Context, input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error) {
    var output opsworks.DescribePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribePermissionsAsync(ctx workflow.Context, input *opsworks.DescribePermissionsInput) *OpsworksDescribePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePermissions, input)
    return &OpsworksDescribePermissionsResult{Result: future}
}

func (a *OpsWorksStub) DescribeRaidArrays(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error) {
    var output opsworks.DescribeRaidArraysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRaidArrays, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeRaidArraysAsync(ctx workflow.Context, input *opsworks.DescribeRaidArraysInput) *OpsworksDescribeRaidArraysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRaidArrays, input)
    return &OpsworksDescribeRaidArraysResult{Result: future}
}

func (a *OpsWorksStub) DescribeRdsDbInstances(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error) {
    var output opsworks.DescribeRdsDbInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRdsDbInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeRdsDbInstancesAsync(ctx workflow.Context, input *opsworks.DescribeRdsDbInstancesInput) *OpsworksDescribeRdsDbInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRdsDbInstances, input)
    return &OpsworksDescribeRdsDbInstancesResult{Result: future}
}

func (a *OpsWorksStub) DescribeServiceErrors(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error) {
    var output opsworks.DescribeServiceErrorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceErrors, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeServiceErrorsAsync(ctx workflow.Context, input *opsworks.DescribeServiceErrorsInput) *OpsworksDescribeServiceErrorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceErrors, input)
    return &OpsworksDescribeServiceErrorsResult{Result: future}
}

func (a *OpsWorksStub) DescribeStackProvisioningParameters(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error) {
    var output opsworks.DescribeStackProvisioningParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackProvisioningParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeStackProvisioningParametersAsync(ctx workflow.Context, input *opsworks.DescribeStackProvisioningParametersInput) *OpsworksDescribeStackProvisioningParametersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStackProvisioningParameters, input)
    return &OpsworksDescribeStackProvisioningParametersResult{Result: future}
}

func (a *OpsWorksStub) DescribeStackSummary(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error) {
    var output opsworks.DescribeStackSummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStackSummary, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeStackSummaryAsync(ctx workflow.Context, input *opsworks.DescribeStackSummaryInput) *OpsworksDescribeStackSummaryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStackSummary, input)
    return &OpsworksDescribeStackSummaryResult{Result: future}
}

func (a *OpsWorksStub) DescribeStacks(ctx workflow.Context, input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error) {
    var output opsworks.DescribeStacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStacks, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeStacksAsync(ctx workflow.Context, input *opsworks.DescribeStacksInput) *OpsworksDescribeStacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeStacks, input)
    return &OpsworksDescribeStacksResult{Result: future}
}

func (a *OpsWorksStub) DescribeTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error) {
    var output opsworks.DescribeTimeBasedAutoScalingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTimeBasedAutoScaling, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.DescribeTimeBasedAutoScalingInput) *OpsworksDescribeTimeBasedAutoScalingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTimeBasedAutoScaling, input)
    return &OpsworksDescribeTimeBasedAutoScalingResult{Result: future}
}

func (a *OpsWorksStub) DescribeUserProfiles(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error) {
    var output opsworks.DescribeUserProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeUserProfilesAsync(ctx workflow.Context, input *opsworks.DescribeUserProfilesInput) *OpsworksDescribeUserProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUserProfiles, input)
    return &OpsworksDescribeUserProfilesResult{Result: future}
}

func (a *OpsWorksStub) DescribeVolumes(ctx workflow.Context, input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error) {
    var output opsworks.DescribeVolumesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumes, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DescribeVolumesAsync(ctx workflow.Context, input *opsworks.DescribeVolumesInput) *OpsworksDescribeVolumesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVolumes, input)
    return &OpsworksDescribeVolumesResult{Result: future}
}

func (a *OpsWorksStub) DetachElasticLoadBalancer(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error) {
    var output opsworks.DetachElasticLoadBalancerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachElasticLoadBalancer, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DetachElasticLoadBalancerAsync(ctx workflow.Context, input *opsworks.DetachElasticLoadBalancerInput) *OpsworksDetachElasticLoadBalancerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachElasticLoadBalancer, input)
    return &OpsworksDetachElasticLoadBalancerResult{Result: future}
}

func (a *OpsWorksStub) DisassociateElasticIp(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error) {
    var output opsworks.DisassociateElasticIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateElasticIp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) DisassociateElasticIpAsync(ctx workflow.Context, input *opsworks.DisassociateElasticIpInput) *OpsworksDisassociateElasticIpResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateElasticIp, input)
    return &OpsworksDisassociateElasticIpResult{Result: future}
}

func (a *OpsWorksStub) GetHostnameSuggestion(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error) {
    var output opsworks.GetHostnameSuggestionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHostnameSuggestion, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) GetHostnameSuggestionAsync(ctx workflow.Context, input *opsworks.GetHostnameSuggestionInput) *OpsworksGetHostnameSuggestionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetHostnameSuggestion, input)
    return &OpsworksGetHostnameSuggestionResult{Result: future}
}

func (a *OpsWorksStub) GrantAccess(ctx workflow.Context, input *opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error) {
    var output opsworks.GrantAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GrantAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) GrantAccessAsync(ctx workflow.Context, input *opsworks.GrantAccessInput) *OpsworksGrantAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GrantAccess, input)
    return &OpsworksGrantAccessResult{Result: future}
}

func (a *OpsWorksStub) ListTags(ctx workflow.Context, input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error) {
    var output opsworks.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) ListTagsAsync(ctx workflow.Context, input *opsworks.ListTagsInput) *OpsworksListTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
    return &OpsworksListTagsResult{Result: future}
}

func (a *OpsWorksStub) RebootInstance(ctx workflow.Context, input *opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error) {
    var output opsworks.RebootInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) RebootInstanceAsync(ctx workflow.Context, input *opsworks.RebootInstanceInput) *OpsworksRebootInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebootInstance, input)
    return &OpsworksRebootInstanceResult{Result: future}
}

func (a *OpsWorksStub) RegisterEcsCluster(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error) {
    var output opsworks.RegisterEcsClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterEcsCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) RegisterEcsClusterAsync(ctx workflow.Context, input *opsworks.RegisterEcsClusterInput) *OpsworksRegisterEcsClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterEcsCluster, input)
    return &OpsworksRegisterEcsClusterResult{Result: future}
}

func (a *OpsWorksStub) RegisterElasticIp(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error) {
    var output opsworks.RegisterElasticIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterElasticIp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) RegisterElasticIpAsync(ctx workflow.Context, input *opsworks.RegisterElasticIpInput) *OpsworksRegisterElasticIpResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterElasticIp, input)
    return &OpsworksRegisterElasticIpResult{Result: future}
}

func (a *OpsWorksStub) RegisterInstance(ctx workflow.Context, input *opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error) {
    var output opsworks.RegisterInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) RegisterInstanceAsync(ctx workflow.Context, input *opsworks.RegisterInstanceInput) *OpsworksRegisterInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterInstance, input)
    return &OpsworksRegisterInstanceResult{Result: future}
}

func (a *OpsWorksStub) RegisterRdsDbInstance(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error) {
    var output opsworks.RegisterRdsDbInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterRdsDbInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) RegisterRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.RegisterRdsDbInstanceInput) *OpsworksRegisterRdsDbInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterRdsDbInstance, input)
    return &OpsworksRegisterRdsDbInstanceResult{Result: future}
}

func (a *OpsWorksStub) RegisterVolume(ctx workflow.Context, input *opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error) {
    var output opsworks.RegisterVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) RegisterVolumeAsync(ctx workflow.Context, input *opsworks.RegisterVolumeInput) *OpsworksRegisterVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterVolume, input)
    return &OpsworksRegisterVolumeResult{Result: future}
}

func (a *OpsWorksStub) SetLoadBasedAutoScaling(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error) {
    var output opsworks.SetLoadBasedAutoScalingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetLoadBasedAutoScaling, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) SetLoadBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetLoadBasedAutoScalingInput) *OpsworksSetLoadBasedAutoScalingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetLoadBasedAutoScaling, input)
    return &OpsworksSetLoadBasedAutoScalingResult{Result: future}
}

func (a *OpsWorksStub) SetPermission(ctx workflow.Context, input *opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error) {
    var output opsworks.SetPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) SetPermissionAsync(ctx workflow.Context, input *opsworks.SetPermissionInput) *OpsworksSetPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetPermission, input)
    return &OpsworksSetPermissionResult{Result: future}
}

func (a *OpsWorksStub) SetTimeBasedAutoScaling(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error) {
    var output opsworks.SetTimeBasedAutoScalingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetTimeBasedAutoScaling, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) SetTimeBasedAutoScalingAsync(ctx workflow.Context, input *opsworks.SetTimeBasedAutoScalingInput) *OpsworksSetTimeBasedAutoScalingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetTimeBasedAutoScaling, input)
    return &OpsworksSetTimeBasedAutoScalingResult{Result: future}
}

func (a *OpsWorksStub) StartInstance(ctx workflow.Context, input *opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error) {
    var output opsworks.StartInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) StartInstanceAsync(ctx workflow.Context, input *opsworks.StartInstanceInput) *OpsworksStartInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartInstance, input)
    return &OpsworksStartInstanceResult{Result: future}
}

func (a *OpsWorksStub) StartStack(ctx workflow.Context, input *opsworks.StartStackInput) (*opsworks.StartStackOutput, error) {
    var output opsworks.StartStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartStack, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) StartStackAsync(ctx workflow.Context, input *opsworks.StartStackInput) *OpsworksStartStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartStack, input)
    return &OpsworksStartStackResult{Result: future}
}

func (a *OpsWorksStub) StopInstance(ctx workflow.Context, input *opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error) {
    var output opsworks.StopInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) StopInstanceAsync(ctx workflow.Context, input *opsworks.StopInstanceInput) *OpsworksStopInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopInstance, input)
    return &OpsworksStopInstanceResult{Result: future}
}

func (a *OpsWorksStub) StopStack(ctx workflow.Context, input *opsworks.StopStackInput) (*opsworks.StopStackOutput, error) {
    var output opsworks.StopStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopStack, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) StopStackAsync(ctx workflow.Context, input *opsworks.StopStackInput) *OpsworksStopStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopStack, input)
    return &OpsworksStopStackResult{Result: future}
}

func (a *OpsWorksStub) TagResource(ctx workflow.Context, input *opsworks.TagResourceInput) (*opsworks.TagResourceOutput, error) {
    var output opsworks.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) TagResourceAsync(ctx workflow.Context, input *opsworks.TagResourceInput) *OpsworksTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &OpsworksTagResourceResult{Result: future}
}

func (a *OpsWorksStub) UnassignInstance(ctx workflow.Context, input *opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error) {
    var output opsworks.UnassignInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnassignInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UnassignInstanceAsync(ctx workflow.Context, input *opsworks.UnassignInstanceInput) *OpsworksUnassignInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnassignInstance, input)
    return &OpsworksUnassignInstanceResult{Result: future}
}

func (a *OpsWorksStub) UnassignVolume(ctx workflow.Context, input *opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error) {
    var output opsworks.UnassignVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnassignVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UnassignVolumeAsync(ctx workflow.Context, input *opsworks.UnassignVolumeInput) *OpsworksUnassignVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UnassignVolume, input)
    return &OpsworksUnassignVolumeResult{Result: future}
}

func (a *OpsWorksStub) UntagResource(ctx workflow.Context, input *opsworks.UntagResourceInput) (*opsworks.UntagResourceOutput, error) {
    var output opsworks.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UntagResourceAsync(ctx workflow.Context, input *opsworks.UntagResourceInput) *OpsworksUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &OpsworksUntagResourceResult{Result: future}
}

func (a *OpsWorksStub) UpdateApp(ctx workflow.Context, input *opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error) {
    var output opsworks.UpdateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateAppAsync(ctx workflow.Context, input *opsworks.UpdateAppInput) *OpsworksUpdateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApp, input)
    return &OpsworksUpdateAppResult{Result: future}
}

func (a *OpsWorksStub) UpdateElasticIp(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error) {
    var output opsworks.UpdateElasticIpOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateElasticIp, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateElasticIpAsync(ctx workflow.Context, input *opsworks.UpdateElasticIpInput) *OpsworksUpdateElasticIpResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateElasticIp, input)
    return &OpsworksUpdateElasticIpResult{Result: future}
}

func (a *OpsWorksStub) UpdateInstance(ctx workflow.Context, input *opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error) {
    var output opsworks.UpdateInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateInstanceAsync(ctx workflow.Context, input *opsworks.UpdateInstanceInput) *OpsworksUpdateInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateInstance, input)
    return &OpsworksUpdateInstanceResult{Result: future}
}

func (a *OpsWorksStub) UpdateLayer(ctx workflow.Context, input *opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error) {
    var output opsworks.UpdateLayerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLayer, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateLayerAsync(ctx workflow.Context, input *opsworks.UpdateLayerInput) *OpsworksUpdateLayerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateLayer, input)
    return &OpsworksUpdateLayerResult{Result: future}
}

func (a *OpsWorksStub) UpdateMyUserProfile(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error) {
    var output opsworks.UpdateMyUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMyUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateMyUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateMyUserProfileInput) *OpsworksUpdateMyUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMyUserProfile, input)
    return &OpsworksUpdateMyUserProfileResult{Result: future}
}

func (a *OpsWorksStub) UpdateRdsDbInstance(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error) {
    var output opsworks.UpdateRdsDbInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRdsDbInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateRdsDbInstanceAsync(ctx workflow.Context, input *opsworks.UpdateRdsDbInstanceInput) *OpsworksUpdateRdsDbInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRdsDbInstance, input)
    return &OpsworksUpdateRdsDbInstanceResult{Result: future}
}

func (a *OpsWorksStub) UpdateStack(ctx workflow.Context, input *opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error) {
    var output opsworks.UpdateStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateStack, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateStackAsync(ctx workflow.Context, input *opsworks.UpdateStackInput) *OpsworksUpdateStackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateStack, input)
    return &OpsworksUpdateStackResult{Result: future}
}

func (a *OpsWorksStub) UpdateUserProfile(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error) {
    var output opsworks.UpdateUserProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateUserProfileAsync(ctx workflow.Context, input *opsworks.UpdateUserProfileInput) *OpsworksUpdateUserProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserProfile, input)
    return &OpsworksUpdateUserProfileResult{Result: future}
}

func (a *OpsWorksStub) UpdateVolume(ctx workflow.Context, input *opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error) {
    var output opsworks.UpdateVolumeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVolume, input).Get(ctx, &output)
    return &output, err
}

func (a *OpsWorksStub) UpdateVolumeAsync(ctx workflow.Context, input *opsworks.UpdateVolumeInput) *OpsworksUpdateVolumeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVolume, input)
    return &OpsworksUpdateVolumeResult{Result: future}
}

func (a *OpsWorksStub) WaitUntilAppExists(ctx workflow.Context, input *opsworks.DescribeAppsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAppExists, input).Get(ctx, nil)
}

func (a *OpsWorksStub) WaitUntilAppExistsAsync(ctx workflow.Context, input *opsworks.DescribeAppsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAppExists, input)
}


func (a *OpsWorksStub) WaitUntilDeploymentSuccessful(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDeploymentSuccessful, input).Get(ctx, nil)
}

func (a *OpsWorksStub) WaitUntilDeploymentSuccessfulAsync(ctx workflow.Context, input *opsworks.DescribeDeploymentsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilDeploymentSuccessful, input)
}


func (a *OpsWorksStub) WaitUntilInstanceOnline(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceOnline, input).Get(ctx, nil)
}

func (a *OpsWorksStub) WaitUntilInstanceOnlineAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceOnline, input)
}


func (a *OpsWorksStub) WaitUntilInstanceRegistered(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceRegistered, input).Get(ctx, nil)
}

func (a *OpsWorksStub) WaitUntilInstanceRegisteredAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceRegistered, input)
}


func (a *OpsWorksStub) WaitUntilInstanceStopped(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceStopped, input).Get(ctx, nil)
}

func (a *OpsWorksStub) WaitUntilInstanceStoppedAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceStopped, input)
}


func (a *OpsWorksStub) WaitUntilInstanceTerminated(ctx workflow.Context, input *opsworks.DescribeInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceTerminated, input).Get(ctx, nil)
}

func (a *OpsWorksStub) WaitUntilInstanceTerminatedAsync(ctx workflow.Context, input *opsworks.DescribeInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilInstanceTerminated, input)
}

