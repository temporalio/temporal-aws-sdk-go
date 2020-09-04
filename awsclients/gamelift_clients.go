package awsclients

import (
	"github.com/aws/aws-sdk-go/service/gamelift"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type GameLiftClient interface {
    AcceptMatch(ctx workflow.Context, input *gamelift.AcceptMatchInput) (*gamelift.AcceptMatchOutput, error)
    AcceptMatchAsync(ctx workflow.Context, input *gamelift.AcceptMatchInput) *GameliftAcceptMatchResult

    ClaimGameServer(ctx workflow.Context, input *gamelift.ClaimGameServerInput) (*gamelift.ClaimGameServerOutput, error)
    ClaimGameServerAsync(ctx workflow.Context, input *gamelift.ClaimGameServerInput) *GameliftClaimGameServerResult

    CreateAlias(ctx workflow.Context, input *gamelift.CreateAliasInput) (*gamelift.CreateAliasOutput, error)
    CreateAliasAsync(ctx workflow.Context, input *gamelift.CreateAliasInput) *GameliftCreateAliasResult

    CreateBuild(ctx workflow.Context, input *gamelift.CreateBuildInput) (*gamelift.CreateBuildOutput, error)
    CreateBuildAsync(ctx workflow.Context, input *gamelift.CreateBuildInput) *GameliftCreateBuildResult

    CreateFleet(ctx workflow.Context, input *gamelift.CreateFleetInput) (*gamelift.CreateFleetOutput, error)
    CreateFleetAsync(ctx workflow.Context, input *gamelift.CreateFleetInput) *GameliftCreateFleetResult

    CreateGameServerGroup(ctx workflow.Context, input *gamelift.CreateGameServerGroupInput) (*gamelift.CreateGameServerGroupOutput, error)
    CreateGameServerGroupAsync(ctx workflow.Context, input *gamelift.CreateGameServerGroupInput) *GameliftCreateGameServerGroupResult

    CreateGameSession(ctx workflow.Context, input *gamelift.CreateGameSessionInput) (*gamelift.CreateGameSessionOutput, error)
    CreateGameSessionAsync(ctx workflow.Context, input *gamelift.CreateGameSessionInput) *GameliftCreateGameSessionResult

    CreateGameSessionQueue(ctx workflow.Context, input *gamelift.CreateGameSessionQueueInput) (*gamelift.CreateGameSessionQueueOutput, error)
    CreateGameSessionQueueAsync(ctx workflow.Context, input *gamelift.CreateGameSessionQueueInput) *GameliftCreateGameSessionQueueResult

    CreateMatchmakingConfiguration(ctx workflow.Context, input *gamelift.CreateMatchmakingConfigurationInput) (*gamelift.CreateMatchmakingConfigurationOutput, error)
    CreateMatchmakingConfigurationAsync(ctx workflow.Context, input *gamelift.CreateMatchmakingConfigurationInput) *GameliftCreateMatchmakingConfigurationResult

    CreateMatchmakingRuleSet(ctx workflow.Context, input *gamelift.CreateMatchmakingRuleSetInput) (*gamelift.CreateMatchmakingRuleSetOutput, error)
    CreateMatchmakingRuleSetAsync(ctx workflow.Context, input *gamelift.CreateMatchmakingRuleSetInput) *GameliftCreateMatchmakingRuleSetResult

    CreatePlayerSession(ctx workflow.Context, input *gamelift.CreatePlayerSessionInput) (*gamelift.CreatePlayerSessionOutput, error)
    CreatePlayerSessionAsync(ctx workflow.Context, input *gamelift.CreatePlayerSessionInput) *GameliftCreatePlayerSessionResult

    CreatePlayerSessions(ctx workflow.Context, input *gamelift.CreatePlayerSessionsInput) (*gamelift.CreatePlayerSessionsOutput, error)
    CreatePlayerSessionsAsync(ctx workflow.Context, input *gamelift.CreatePlayerSessionsInput) *GameliftCreatePlayerSessionsResult

    CreateScript(ctx workflow.Context, input *gamelift.CreateScriptInput) (*gamelift.CreateScriptOutput, error)
    CreateScriptAsync(ctx workflow.Context, input *gamelift.CreateScriptInput) *GameliftCreateScriptResult

    CreateVpcPeeringAuthorization(ctx workflow.Context, input *gamelift.CreateVpcPeeringAuthorizationInput) (*gamelift.CreateVpcPeeringAuthorizationOutput, error)
    CreateVpcPeeringAuthorizationAsync(ctx workflow.Context, input *gamelift.CreateVpcPeeringAuthorizationInput) *GameliftCreateVpcPeeringAuthorizationResult

    CreateVpcPeeringConnection(ctx workflow.Context, input *gamelift.CreateVpcPeeringConnectionInput) (*gamelift.CreateVpcPeeringConnectionOutput, error)
    CreateVpcPeeringConnectionAsync(ctx workflow.Context, input *gamelift.CreateVpcPeeringConnectionInput) *GameliftCreateVpcPeeringConnectionResult

    DeleteAlias(ctx workflow.Context, input *gamelift.DeleteAliasInput) (*gamelift.DeleteAliasOutput, error)
    DeleteAliasAsync(ctx workflow.Context, input *gamelift.DeleteAliasInput) *GameliftDeleteAliasResult

    DeleteBuild(ctx workflow.Context, input *gamelift.DeleteBuildInput) (*gamelift.DeleteBuildOutput, error)
    DeleteBuildAsync(ctx workflow.Context, input *gamelift.DeleteBuildInput) *GameliftDeleteBuildResult

    DeleteFleet(ctx workflow.Context, input *gamelift.DeleteFleetInput) (*gamelift.DeleteFleetOutput, error)
    DeleteFleetAsync(ctx workflow.Context, input *gamelift.DeleteFleetInput) *GameliftDeleteFleetResult

    DeleteGameServerGroup(ctx workflow.Context, input *gamelift.DeleteGameServerGroupInput) (*gamelift.DeleteGameServerGroupOutput, error)
    DeleteGameServerGroupAsync(ctx workflow.Context, input *gamelift.DeleteGameServerGroupInput) *GameliftDeleteGameServerGroupResult

    DeleteGameSessionQueue(ctx workflow.Context, input *gamelift.DeleteGameSessionQueueInput) (*gamelift.DeleteGameSessionQueueOutput, error)
    DeleteGameSessionQueueAsync(ctx workflow.Context, input *gamelift.DeleteGameSessionQueueInput) *GameliftDeleteGameSessionQueueResult

    DeleteMatchmakingConfiguration(ctx workflow.Context, input *gamelift.DeleteMatchmakingConfigurationInput) (*gamelift.DeleteMatchmakingConfigurationOutput, error)
    DeleteMatchmakingConfigurationAsync(ctx workflow.Context, input *gamelift.DeleteMatchmakingConfigurationInput) *GameliftDeleteMatchmakingConfigurationResult

    DeleteMatchmakingRuleSet(ctx workflow.Context, input *gamelift.DeleteMatchmakingRuleSetInput) (*gamelift.DeleteMatchmakingRuleSetOutput, error)
    DeleteMatchmakingRuleSetAsync(ctx workflow.Context, input *gamelift.DeleteMatchmakingRuleSetInput) *GameliftDeleteMatchmakingRuleSetResult

    DeleteScalingPolicy(ctx workflow.Context, input *gamelift.DeleteScalingPolicyInput) (*gamelift.DeleteScalingPolicyOutput, error)
    DeleteScalingPolicyAsync(ctx workflow.Context, input *gamelift.DeleteScalingPolicyInput) *GameliftDeleteScalingPolicyResult

    DeleteScript(ctx workflow.Context, input *gamelift.DeleteScriptInput) (*gamelift.DeleteScriptOutput, error)
    DeleteScriptAsync(ctx workflow.Context, input *gamelift.DeleteScriptInput) *GameliftDeleteScriptResult

    DeleteVpcPeeringAuthorization(ctx workflow.Context, input *gamelift.DeleteVpcPeeringAuthorizationInput) (*gamelift.DeleteVpcPeeringAuthorizationOutput, error)
    DeleteVpcPeeringAuthorizationAsync(ctx workflow.Context, input *gamelift.DeleteVpcPeeringAuthorizationInput) *GameliftDeleteVpcPeeringAuthorizationResult

    DeleteVpcPeeringConnection(ctx workflow.Context, input *gamelift.DeleteVpcPeeringConnectionInput) (*gamelift.DeleteVpcPeeringConnectionOutput, error)
    DeleteVpcPeeringConnectionAsync(ctx workflow.Context, input *gamelift.DeleteVpcPeeringConnectionInput) *GameliftDeleteVpcPeeringConnectionResult

    DeregisterGameServer(ctx workflow.Context, input *gamelift.DeregisterGameServerInput) (*gamelift.DeregisterGameServerOutput, error)
    DeregisterGameServerAsync(ctx workflow.Context, input *gamelift.DeregisterGameServerInput) *GameliftDeregisterGameServerResult

    DescribeAlias(ctx workflow.Context, input *gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error)
    DescribeAliasAsync(ctx workflow.Context, input *gamelift.DescribeAliasInput) *GameliftDescribeAliasResult

    DescribeBuild(ctx workflow.Context, input *gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error)
    DescribeBuildAsync(ctx workflow.Context, input *gamelift.DescribeBuildInput) *GameliftDescribeBuildResult

    DescribeEC2InstanceLimits(ctx workflow.Context, input *gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error)
    DescribeEC2InstanceLimitsAsync(ctx workflow.Context, input *gamelift.DescribeEC2InstanceLimitsInput) *GameliftDescribeEC2InstanceLimitsResult

    DescribeFleetAttributes(ctx workflow.Context, input *gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error)
    DescribeFleetAttributesAsync(ctx workflow.Context, input *gamelift.DescribeFleetAttributesInput) *GameliftDescribeFleetAttributesResult

    DescribeFleetCapacity(ctx workflow.Context, input *gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error)
    DescribeFleetCapacityAsync(ctx workflow.Context, input *gamelift.DescribeFleetCapacityInput) *GameliftDescribeFleetCapacityResult

    DescribeFleetEvents(ctx workflow.Context, input *gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error)
    DescribeFleetEventsAsync(ctx workflow.Context, input *gamelift.DescribeFleetEventsInput) *GameliftDescribeFleetEventsResult

    DescribeFleetPortSettings(ctx workflow.Context, input *gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error)
    DescribeFleetPortSettingsAsync(ctx workflow.Context, input *gamelift.DescribeFleetPortSettingsInput) *GameliftDescribeFleetPortSettingsResult

    DescribeFleetUtilization(ctx workflow.Context, input *gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error)
    DescribeFleetUtilizationAsync(ctx workflow.Context, input *gamelift.DescribeFleetUtilizationInput) *GameliftDescribeFleetUtilizationResult

    DescribeGameServer(ctx workflow.Context, input *gamelift.DescribeGameServerInput) (*gamelift.DescribeGameServerOutput, error)
    DescribeGameServerAsync(ctx workflow.Context, input *gamelift.DescribeGameServerInput) *GameliftDescribeGameServerResult

    DescribeGameServerGroup(ctx workflow.Context, input *gamelift.DescribeGameServerGroupInput) (*gamelift.DescribeGameServerGroupOutput, error)
    DescribeGameServerGroupAsync(ctx workflow.Context, input *gamelift.DescribeGameServerGroupInput) *GameliftDescribeGameServerGroupResult

    DescribeGameServerInstances(ctx workflow.Context, input *gamelift.DescribeGameServerInstancesInput) (*gamelift.DescribeGameServerInstancesOutput, error)
    DescribeGameServerInstancesAsync(ctx workflow.Context, input *gamelift.DescribeGameServerInstancesInput) *GameliftDescribeGameServerInstancesResult

    DescribeGameSessionDetails(ctx workflow.Context, input *gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error)
    DescribeGameSessionDetailsAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionDetailsInput) *GameliftDescribeGameSessionDetailsResult

    DescribeGameSessionPlacement(ctx workflow.Context, input *gamelift.DescribeGameSessionPlacementInput) (*gamelift.DescribeGameSessionPlacementOutput, error)
    DescribeGameSessionPlacementAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionPlacementInput) *GameliftDescribeGameSessionPlacementResult

    DescribeGameSessionQueues(ctx workflow.Context, input *gamelift.DescribeGameSessionQueuesInput) (*gamelift.DescribeGameSessionQueuesOutput, error)
    DescribeGameSessionQueuesAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionQueuesInput) *GameliftDescribeGameSessionQueuesResult

    DescribeGameSessions(ctx workflow.Context, input *gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error)
    DescribeGameSessionsAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionsInput) *GameliftDescribeGameSessionsResult

    DescribeInstances(ctx workflow.Context, input *gamelift.DescribeInstancesInput) (*gamelift.DescribeInstancesOutput, error)
    DescribeInstancesAsync(ctx workflow.Context, input *gamelift.DescribeInstancesInput) *GameliftDescribeInstancesResult

    DescribeMatchmaking(ctx workflow.Context, input *gamelift.DescribeMatchmakingInput) (*gamelift.DescribeMatchmakingOutput, error)
    DescribeMatchmakingAsync(ctx workflow.Context, input *gamelift.DescribeMatchmakingInput) *GameliftDescribeMatchmakingResult

    DescribeMatchmakingConfigurations(ctx workflow.Context, input *gamelift.DescribeMatchmakingConfigurationsInput) (*gamelift.DescribeMatchmakingConfigurationsOutput, error)
    DescribeMatchmakingConfigurationsAsync(ctx workflow.Context, input *gamelift.DescribeMatchmakingConfigurationsInput) *GameliftDescribeMatchmakingConfigurationsResult

    DescribeMatchmakingRuleSets(ctx workflow.Context, input *gamelift.DescribeMatchmakingRuleSetsInput) (*gamelift.DescribeMatchmakingRuleSetsOutput, error)
    DescribeMatchmakingRuleSetsAsync(ctx workflow.Context, input *gamelift.DescribeMatchmakingRuleSetsInput) *GameliftDescribeMatchmakingRuleSetsResult

    DescribePlayerSessions(ctx workflow.Context, input *gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error)
    DescribePlayerSessionsAsync(ctx workflow.Context, input *gamelift.DescribePlayerSessionsInput) *GameliftDescribePlayerSessionsResult

    DescribeRuntimeConfiguration(ctx workflow.Context, input *gamelift.DescribeRuntimeConfigurationInput) (*gamelift.DescribeRuntimeConfigurationOutput, error)
    DescribeRuntimeConfigurationAsync(ctx workflow.Context, input *gamelift.DescribeRuntimeConfigurationInput) *GameliftDescribeRuntimeConfigurationResult

    DescribeScalingPolicies(ctx workflow.Context, input *gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error)
    DescribeScalingPoliciesAsync(ctx workflow.Context, input *gamelift.DescribeScalingPoliciesInput) *GameliftDescribeScalingPoliciesResult

    DescribeScript(ctx workflow.Context, input *gamelift.DescribeScriptInput) (*gamelift.DescribeScriptOutput, error)
    DescribeScriptAsync(ctx workflow.Context, input *gamelift.DescribeScriptInput) *GameliftDescribeScriptResult

    DescribeVpcPeeringAuthorizations(ctx workflow.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error)
    DescribeVpcPeeringAuthorizationsAsync(ctx workflow.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput) *GameliftDescribeVpcPeeringAuthorizationsResult

    DescribeVpcPeeringConnections(ctx workflow.Context, input *gamelift.DescribeVpcPeeringConnectionsInput) (*gamelift.DescribeVpcPeeringConnectionsOutput, error)
    DescribeVpcPeeringConnectionsAsync(ctx workflow.Context, input *gamelift.DescribeVpcPeeringConnectionsInput) *GameliftDescribeVpcPeeringConnectionsResult

    GetGameSessionLogUrl(ctx workflow.Context, input *gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error)
    GetGameSessionLogUrlAsync(ctx workflow.Context, input *gamelift.GetGameSessionLogUrlInput) *GameliftGetGameSessionLogUrlResult

    GetInstanceAccess(ctx workflow.Context, input *gamelift.GetInstanceAccessInput) (*gamelift.GetInstanceAccessOutput, error)
    GetInstanceAccessAsync(ctx workflow.Context, input *gamelift.GetInstanceAccessInput) *GameliftGetInstanceAccessResult

    ListAliases(ctx workflow.Context, input *gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error)
    ListAliasesAsync(ctx workflow.Context, input *gamelift.ListAliasesInput) *GameliftListAliasesResult

    ListBuilds(ctx workflow.Context, input *gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error)
    ListBuildsAsync(ctx workflow.Context, input *gamelift.ListBuildsInput) *GameliftListBuildsResult

    ListFleets(ctx workflow.Context, input *gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error)
    ListFleetsAsync(ctx workflow.Context, input *gamelift.ListFleetsInput) *GameliftListFleetsResult

    ListGameServerGroups(ctx workflow.Context, input *gamelift.ListGameServerGroupsInput) (*gamelift.ListGameServerGroupsOutput, error)
    ListGameServerGroupsAsync(ctx workflow.Context, input *gamelift.ListGameServerGroupsInput) *GameliftListGameServerGroupsResult

    ListGameServers(ctx workflow.Context, input *gamelift.ListGameServersInput) (*gamelift.ListGameServersOutput, error)
    ListGameServersAsync(ctx workflow.Context, input *gamelift.ListGameServersInput) *GameliftListGameServersResult

    ListScripts(ctx workflow.Context, input *gamelift.ListScriptsInput) (*gamelift.ListScriptsOutput, error)
    ListScriptsAsync(ctx workflow.Context, input *gamelift.ListScriptsInput) *GameliftListScriptsResult

    ListTagsForResource(ctx workflow.Context, input *gamelift.ListTagsForResourceInput) (*gamelift.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *gamelift.ListTagsForResourceInput) *GameliftListTagsForResourceResult

    PutScalingPolicy(ctx workflow.Context, input *gamelift.PutScalingPolicyInput) (*gamelift.PutScalingPolicyOutput, error)
    PutScalingPolicyAsync(ctx workflow.Context, input *gamelift.PutScalingPolicyInput) *GameliftPutScalingPolicyResult

    RegisterGameServer(ctx workflow.Context, input *gamelift.RegisterGameServerInput) (*gamelift.RegisterGameServerOutput, error)
    RegisterGameServerAsync(ctx workflow.Context, input *gamelift.RegisterGameServerInput) *GameliftRegisterGameServerResult

    RequestUploadCredentials(ctx workflow.Context, input *gamelift.RequestUploadCredentialsInput) (*gamelift.RequestUploadCredentialsOutput, error)
    RequestUploadCredentialsAsync(ctx workflow.Context, input *gamelift.RequestUploadCredentialsInput) *GameliftRequestUploadCredentialsResult

    ResolveAlias(ctx workflow.Context, input *gamelift.ResolveAliasInput) (*gamelift.ResolveAliasOutput, error)
    ResolveAliasAsync(ctx workflow.Context, input *gamelift.ResolveAliasInput) *GameliftResolveAliasResult

    ResumeGameServerGroup(ctx workflow.Context, input *gamelift.ResumeGameServerGroupInput) (*gamelift.ResumeGameServerGroupOutput, error)
    ResumeGameServerGroupAsync(ctx workflow.Context, input *gamelift.ResumeGameServerGroupInput) *GameliftResumeGameServerGroupResult

    SearchGameSessions(ctx workflow.Context, input *gamelift.SearchGameSessionsInput) (*gamelift.SearchGameSessionsOutput, error)
    SearchGameSessionsAsync(ctx workflow.Context, input *gamelift.SearchGameSessionsInput) *GameliftSearchGameSessionsResult

    StartFleetActions(ctx workflow.Context, input *gamelift.StartFleetActionsInput) (*gamelift.StartFleetActionsOutput, error)
    StartFleetActionsAsync(ctx workflow.Context, input *gamelift.StartFleetActionsInput) *GameliftStartFleetActionsResult

    StartGameSessionPlacement(ctx workflow.Context, input *gamelift.StartGameSessionPlacementInput) (*gamelift.StartGameSessionPlacementOutput, error)
    StartGameSessionPlacementAsync(ctx workflow.Context, input *gamelift.StartGameSessionPlacementInput) *GameliftStartGameSessionPlacementResult

    StartMatchBackfill(ctx workflow.Context, input *gamelift.StartMatchBackfillInput) (*gamelift.StartMatchBackfillOutput, error)
    StartMatchBackfillAsync(ctx workflow.Context, input *gamelift.StartMatchBackfillInput) *GameliftStartMatchBackfillResult

    StartMatchmaking(ctx workflow.Context, input *gamelift.StartMatchmakingInput) (*gamelift.StartMatchmakingOutput, error)
    StartMatchmakingAsync(ctx workflow.Context, input *gamelift.StartMatchmakingInput) *GameliftStartMatchmakingResult

    StopFleetActions(ctx workflow.Context, input *gamelift.StopFleetActionsInput) (*gamelift.StopFleetActionsOutput, error)
    StopFleetActionsAsync(ctx workflow.Context, input *gamelift.StopFleetActionsInput) *GameliftStopFleetActionsResult

    StopGameSessionPlacement(ctx workflow.Context, input *gamelift.StopGameSessionPlacementInput) (*gamelift.StopGameSessionPlacementOutput, error)
    StopGameSessionPlacementAsync(ctx workflow.Context, input *gamelift.StopGameSessionPlacementInput) *GameliftStopGameSessionPlacementResult

    StopMatchmaking(ctx workflow.Context, input *gamelift.StopMatchmakingInput) (*gamelift.StopMatchmakingOutput, error)
    StopMatchmakingAsync(ctx workflow.Context, input *gamelift.StopMatchmakingInput) *GameliftStopMatchmakingResult

    SuspendGameServerGroup(ctx workflow.Context, input *gamelift.SuspendGameServerGroupInput) (*gamelift.SuspendGameServerGroupOutput, error)
    SuspendGameServerGroupAsync(ctx workflow.Context, input *gamelift.SuspendGameServerGroupInput) *GameliftSuspendGameServerGroupResult

    TagResource(ctx workflow.Context, input *gamelift.TagResourceInput) (*gamelift.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *gamelift.TagResourceInput) *GameliftTagResourceResult

    UntagResource(ctx workflow.Context, input *gamelift.UntagResourceInput) (*gamelift.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *gamelift.UntagResourceInput) *GameliftUntagResourceResult

    UpdateAlias(ctx workflow.Context, input *gamelift.UpdateAliasInput) (*gamelift.UpdateAliasOutput, error)
    UpdateAliasAsync(ctx workflow.Context, input *gamelift.UpdateAliasInput) *GameliftUpdateAliasResult

    UpdateBuild(ctx workflow.Context, input *gamelift.UpdateBuildInput) (*gamelift.UpdateBuildOutput, error)
    UpdateBuildAsync(ctx workflow.Context, input *gamelift.UpdateBuildInput) *GameliftUpdateBuildResult

    UpdateFleetAttributes(ctx workflow.Context, input *gamelift.UpdateFleetAttributesInput) (*gamelift.UpdateFleetAttributesOutput, error)
    UpdateFleetAttributesAsync(ctx workflow.Context, input *gamelift.UpdateFleetAttributesInput) *GameliftUpdateFleetAttributesResult

    UpdateFleetCapacity(ctx workflow.Context, input *gamelift.UpdateFleetCapacityInput) (*gamelift.UpdateFleetCapacityOutput, error)
    UpdateFleetCapacityAsync(ctx workflow.Context, input *gamelift.UpdateFleetCapacityInput) *GameliftUpdateFleetCapacityResult

    UpdateFleetPortSettings(ctx workflow.Context, input *gamelift.UpdateFleetPortSettingsInput) (*gamelift.UpdateFleetPortSettingsOutput, error)
    UpdateFleetPortSettingsAsync(ctx workflow.Context, input *gamelift.UpdateFleetPortSettingsInput) *GameliftUpdateFleetPortSettingsResult

    UpdateGameServer(ctx workflow.Context, input *gamelift.UpdateGameServerInput) (*gamelift.UpdateGameServerOutput, error)
    UpdateGameServerAsync(ctx workflow.Context, input *gamelift.UpdateGameServerInput) *GameliftUpdateGameServerResult

    UpdateGameServerGroup(ctx workflow.Context, input *gamelift.UpdateGameServerGroupInput) (*gamelift.UpdateGameServerGroupOutput, error)
    UpdateGameServerGroupAsync(ctx workflow.Context, input *gamelift.UpdateGameServerGroupInput) *GameliftUpdateGameServerGroupResult

    UpdateGameSession(ctx workflow.Context, input *gamelift.UpdateGameSessionInput) (*gamelift.UpdateGameSessionOutput, error)
    UpdateGameSessionAsync(ctx workflow.Context, input *gamelift.UpdateGameSessionInput) *GameliftUpdateGameSessionResult

    UpdateGameSessionQueue(ctx workflow.Context, input *gamelift.UpdateGameSessionQueueInput) (*gamelift.UpdateGameSessionQueueOutput, error)
    UpdateGameSessionQueueAsync(ctx workflow.Context, input *gamelift.UpdateGameSessionQueueInput) *GameliftUpdateGameSessionQueueResult

    UpdateMatchmakingConfiguration(ctx workflow.Context, input *gamelift.UpdateMatchmakingConfigurationInput) (*gamelift.UpdateMatchmakingConfigurationOutput, error)
    UpdateMatchmakingConfigurationAsync(ctx workflow.Context, input *gamelift.UpdateMatchmakingConfigurationInput) *GameliftUpdateMatchmakingConfigurationResult

    UpdateRuntimeConfiguration(ctx workflow.Context, input *gamelift.UpdateRuntimeConfigurationInput) (*gamelift.UpdateRuntimeConfigurationOutput, error)
    UpdateRuntimeConfigurationAsync(ctx workflow.Context, input *gamelift.UpdateRuntimeConfigurationInput) *GameliftUpdateRuntimeConfigurationResult

    UpdateScript(ctx workflow.Context, input *gamelift.UpdateScriptInput) (*gamelift.UpdateScriptOutput, error)
    UpdateScriptAsync(ctx workflow.Context, input *gamelift.UpdateScriptInput) *GameliftUpdateScriptResult

    ValidateMatchmakingRuleSet(ctx workflow.Context, input *gamelift.ValidateMatchmakingRuleSetInput) (*gamelift.ValidateMatchmakingRuleSetOutput, error)
    ValidateMatchmakingRuleSetAsync(ctx workflow.Context, input *gamelift.ValidateMatchmakingRuleSetInput) *GameliftValidateMatchmakingRuleSetResult
}

type GameliftAcceptMatchResult struct {
	Result workflow.Future
}

func (r *GameliftAcceptMatchResult) Get(ctx workflow.Context) (*gamelift.AcceptMatchOutput, error) {
    var output gamelift.AcceptMatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftClaimGameServerResult struct {
	Result workflow.Future
}

func (r *GameliftClaimGameServerResult) Get(ctx workflow.Context) (*gamelift.ClaimGameServerOutput, error) {
    var output gamelift.ClaimGameServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateAliasResult struct {
	Result workflow.Future
}

func (r *GameliftCreateAliasResult) Get(ctx workflow.Context) (*gamelift.CreateAliasOutput, error) {
    var output gamelift.CreateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateBuildResult struct {
	Result workflow.Future
}

func (r *GameliftCreateBuildResult) Get(ctx workflow.Context) (*gamelift.CreateBuildOutput, error) {
    var output gamelift.CreateBuildOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateFleetResult struct {
	Result workflow.Future
}

func (r *GameliftCreateFleetResult) Get(ctx workflow.Context) (*gamelift.CreateFleetOutput, error) {
    var output gamelift.CreateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateGameServerGroupResult struct {
	Result workflow.Future
}

func (r *GameliftCreateGameServerGroupResult) Get(ctx workflow.Context) (*gamelift.CreateGameServerGroupOutput, error) {
    var output gamelift.CreateGameServerGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateGameSessionResult struct {
	Result workflow.Future
}

func (r *GameliftCreateGameSessionResult) Get(ctx workflow.Context) (*gamelift.CreateGameSessionOutput, error) {
    var output gamelift.CreateGameSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateGameSessionQueueResult struct {
	Result workflow.Future
}

func (r *GameliftCreateGameSessionQueueResult) Get(ctx workflow.Context) (*gamelift.CreateGameSessionQueueOutput, error) {
    var output gamelift.CreateGameSessionQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateMatchmakingConfigurationResult struct {
	Result workflow.Future
}

func (r *GameliftCreateMatchmakingConfigurationResult) Get(ctx workflow.Context) (*gamelift.CreateMatchmakingConfigurationOutput, error) {
    var output gamelift.CreateMatchmakingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateMatchmakingRuleSetResult struct {
	Result workflow.Future
}

func (r *GameliftCreateMatchmakingRuleSetResult) Get(ctx workflow.Context) (*gamelift.CreateMatchmakingRuleSetOutput, error) {
    var output gamelift.CreateMatchmakingRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreatePlayerSessionResult struct {
	Result workflow.Future
}

func (r *GameliftCreatePlayerSessionResult) Get(ctx workflow.Context) (*gamelift.CreatePlayerSessionOutput, error) {
    var output gamelift.CreatePlayerSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreatePlayerSessionsResult struct {
	Result workflow.Future
}

func (r *GameliftCreatePlayerSessionsResult) Get(ctx workflow.Context) (*gamelift.CreatePlayerSessionsOutput, error) {
    var output gamelift.CreatePlayerSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateScriptResult struct {
	Result workflow.Future
}

func (r *GameliftCreateScriptResult) Get(ctx workflow.Context) (*gamelift.CreateScriptOutput, error) {
    var output gamelift.CreateScriptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateVpcPeeringAuthorizationResult struct {
	Result workflow.Future
}

func (r *GameliftCreateVpcPeeringAuthorizationResult) Get(ctx workflow.Context) (*gamelift.CreateVpcPeeringAuthorizationOutput, error) {
    var output gamelift.CreateVpcPeeringAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftCreateVpcPeeringConnectionResult struct {
	Result workflow.Future
}

func (r *GameliftCreateVpcPeeringConnectionResult) Get(ctx workflow.Context) (*gamelift.CreateVpcPeeringConnectionOutput, error) {
    var output gamelift.CreateVpcPeeringConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteAliasResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteAliasResult) Get(ctx workflow.Context) (*gamelift.DeleteAliasOutput, error) {
    var output gamelift.DeleteAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteBuildResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteBuildResult) Get(ctx workflow.Context) (*gamelift.DeleteBuildOutput, error) {
    var output gamelift.DeleteBuildOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteFleetResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteFleetResult) Get(ctx workflow.Context) (*gamelift.DeleteFleetOutput, error) {
    var output gamelift.DeleteFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteGameServerGroupResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteGameServerGroupResult) Get(ctx workflow.Context) (*gamelift.DeleteGameServerGroupOutput, error) {
    var output gamelift.DeleteGameServerGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteGameSessionQueueResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteGameSessionQueueResult) Get(ctx workflow.Context) (*gamelift.DeleteGameSessionQueueOutput, error) {
    var output gamelift.DeleteGameSessionQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteMatchmakingConfigurationResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteMatchmakingConfigurationResult) Get(ctx workflow.Context) (*gamelift.DeleteMatchmakingConfigurationOutput, error) {
    var output gamelift.DeleteMatchmakingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteMatchmakingRuleSetResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteMatchmakingRuleSetResult) Get(ctx workflow.Context) (*gamelift.DeleteMatchmakingRuleSetOutput, error) {
    var output gamelift.DeleteMatchmakingRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteScalingPolicyResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteScalingPolicyResult) Get(ctx workflow.Context) (*gamelift.DeleteScalingPolicyOutput, error) {
    var output gamelift.DeleteScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteScriptResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteScriptResult) Get(ctx workflow.Context) (*gamelift.DeleteScriptOutput, error) {
    var output gamelift.DeleteScriptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteVpcPeeringAuthorizationResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteVpcPeeringAuthorizationResult) Get(ctx workflow.Context) (*gamelift.DeleteVpcPeeringAuthorizationOutput, error) {
    var output gamelift.DeleteVpcPeeringAuthorizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeleteVpcPeeringConnectionResult struct {
	Result workflow.Future
}

func (r *GameliftDeleteVpcPeeringConnectionResult) Get(ctx workflow.Context) (*gamelift.DeleteVpcPeeringConnectionOutput, error) {
    var output gamelift.DeleteVpcPeeringConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDeregisterGameServerResult struct {
	Result workflow.Future
}

func (r *GameliftDeregisterGameServerResult) Get(ctx workflow.Context) (*gamelift.DeregisterGameServerOutput, error) {
    var output gamelift.DeregisterGameServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeAliasResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeAliasResult) Get(ctx workflow.Context) (*gamelift.DescribeAliasOutput, error) {
    var output gamelift.DescribeAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeBuildResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeBuildResult) Get(ctx workflow.Context) (*gamelift.DescribeBuildOutput, error) {
    var output gamelift.DescribeBuildOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeEC2InstanceLimitsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeEC2InstanceLimitsResult) Get(ctx workflow.Context) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
    var output gamelift.DescribeEC2InstanceLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeFleetAttributesResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeFleetAttributesResult) Get(ctx workflow.Context) (*gamelift.DescribeFleetAttributesOutput, error) {
    var output gamelift.DescribeFleetAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeFleetCapacityResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeFleetCapacityResult) Get(ctx workflow.Context) (*gamelift.DescribeFleetCapacityOutput, error) {
    var output gamelift.DescribeFleetCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeFleetEventsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeFleetEventsResult) Get(ctx workflow.Context) (*gamelift.DescribeFleetEventsOutput, error) {
    var output gamelift.DescribeFleetEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeFleetPortSettingsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeFleetPortSettingsResult) Get(ctx workflow.Context) (*gamelift.DescribeFleetPortSettingsOutput, error) {
    var output gamelift.DescribeFleetPortSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeFleetUtilizationResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeFleetUtilizationResult) Get(ctx workflow.Context) (*gamelift.DescribeFleetUtilizationOutput, error) {
    var output gamelift.DescribeFleetUtilizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameServerResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameServerResult) Get(ctx workflow.Context) (*gamelift.DescribeGameServerOutput, error) {
    var output gamelift.DescribeGameServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameServerGroupResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameServerGroupResult) Get(ctx workflow.Context) (*gamelift.DescribeGameServerGroupOutput, error) {
    var output gamelift.DescribeGameServerGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameServerInstancesResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameServerInstancesResult) Get(ctx workflow.Context) (*gamelift.DescribeGameServerInstancesOutput, error) {
    var output gamelift.DescribeGameServerInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameSessionDetailsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameSessionDetailsResult) Get(ctx workflow.Context) (*gamelift.DescribeGameSessionDetailsOutput, error) {
    var output gamelift.DescribeGameSessionDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameSessionPlacementResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameSessionPlacementResult) Get(ctx workflow.Context) (*gamelift.DescribeGameSessionPlacementOutput, error) {
    var output gamelift.DescribeGameSessionPlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameSessionQueuesResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameSessionQueuesResult) Get(ctx workflow.Context) (*gamelift.DescribeGameSessionQueuesOutput, error) {
    var output gamelift.DescribeGameSessionQueuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeGameSessionsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeGameSessionsResult) Get(ctx workflow.Context) (*gamelift.DescribeGameSessionsOutput, error) {
    var output gamelift.DescribeGameSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeInstancesResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeInstancesResult) Get(ctx workflow.Context) (*gamelift.DescribeInstancesOutput, error) {
    var output gamelift.DescribeInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeMatchmakingResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeMatchmakingResult) Get(ctx workflow.Context) (*gamelift.DescribeMatchmakingOutput, error) {
    var output gamelift.DescribeMatchmakingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeMatchmakingConfigurationsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeMatchmakingConfigurationsResult) Get(ctx workflow.Context) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
    var output gamelift.DescribeMatchmakingConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeMatchmakingRuleSetsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeMatchmakingRuleSetsResult) Get(ctx workflow.Context) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
    var output gamelift.DescribeMatchmakingRuleSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribePlayerSessionsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribePlayerSessionsResult) Get(ctx workflow.Context) (*gamelift.DescribePlayerSessionsOutput, error) {
    var output gamelift.DescribePlayerSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeRuntimeConfigurationResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeRuntimeConfigurationResult) Get(ctx workflow.Context) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
    var output gamelift.DescribeRuntimeConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeScalingPoliciesResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeScalingPoliciesResult) Get(ctx workflow.Context) (*gamelift.DescribeScalingPoliciesOutput, error) {
    var output gamelift.DescribeScalingPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeScriptResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeScriptResult) Get(ctx workflow.Context) (*gamelift.DescribeScriptOutput, error) {
    var output gamelift.DescribeScriptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeVpcPeeringAuthorizationsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeVpcPeeringAuthorizationsResult) Get(ctx workflow.Context) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
    var output gamelift.DescribeVpcPeeringAuthorizationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftDescribeVpcPeeringConnectionsResult struct {
	Result workflow.Future
}

func (r *GameliftDescribeVpcPeeringConnectionsResult) Get(ctx workflow.Context) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
    var output gamelift.DescribeVpcPeeringConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftGetGameSessionLogUrlResult struct {
	Result workflow.Future
}

func (r *GameliftGetGameSessionLogUrlResult) Get(ctx workflow.Context) (*gamelift.GetGameSessionLogUrlOutput, error) {
    var output gamelift.GetGameSessionLogUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftGetInstanceAccessResult struct {
	Result workflow.Future
}

func (r *GameliftGetInstanceAccessResult) Get(ctx workflow.Context) (*gamelift.GetInstanceAccessOutput, error) {
    var output gamelift.GetInstanceAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListAliasesResult struct {
	Result workflow.Future
}

func (r *GameliftListAliasesResult) Get(ctx workflow.Context) (*gamelift.ListAliasesOutput, error) {
    var output gamelift.ListAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListBuildsResult struct {
	Result workflow.Future
}

func (r *GameliftListBuildsResult) Get(ctx workflow.Context) (*gamelift.ListBuildsOutput, error) {
    var output gamelift.ListBuildsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListFleetsResult struct {
	Result workflow.Future
}

func (r *GameliftListFleetsResult) Get(ctx workflow.Context) (*gamelift.ListFleetsOutput, error) {
    var output gamelift.ListFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListGameServerGroupsResult struct {
	Result workflow.Future
}

func (r *GameliftListGameServerGroupsResult) Get(ctx workflow.Context) (*gamelift.ListGameServerGroupsOutput, error) {
    var output gamelift.ListGameServerGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListGameServersResult struct {
	Result workflow.Future
}

func (r *GameliftListGameServersResult) Get(ctx workflow.Context) (*gamelift.ListGameServersOutput, error) {
    var output gamelift.ListGameServersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListScriptsResult struct {
	Result workflow.Future
}

func (r *GameliftListScriptsResult) Get(ctx workflow.Context) (*gamelift.ListScriptsOutput, error) {
    var output gamelift.ListScriptsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *GameliftListTagsForResourceResult) Get(ctx workflow.Context) (*gamelift.ListTagsForResourceOutput, error) {
    var output gamelift.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftPutScalingPolicyResult struct {
	Result workflow.Future
}

func (r *GameliftPutScalingPolicyResult) Get(ctx workflow.Context) (*gamelift.PutScalingPolicyOutput, error) {
    var output gamelift.PutScalingPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftRegisterGameServerResult struct {
	Result workflow.Future
}

func (r *GameliftRegisterGameServerResult) Get(ctx workflow.Context) (*gamelift.RegisterGameServerOutput, error) {
    var output gamelift.RegisterGameServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftRequestUploadCredentialsResult struct {
	Result workflow.Future
}

func (r *GameliftRequestUploadCredentialsResult) Get(ctx workflow.Context) (*gamelift.RequestUploadCredentialsOutput, error) {
    var output gamelift.RequestUploadCredentialsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftResolveAliasResult struct {
	Result workflow.Future
}

func (r *GameliftResolveAliasResult) Get(ctx workflow.Context) (*gamelift.ResolveAliasOutput, error) {
    var output gamelift.ResolveAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftResumeGameServerGroupResult struct {
	Result workflow.Future
}

func (r *GameliftResumeGameServerGroupResult) Get(ctx workflow.Context) (*gamelift.ResumeGameServerGroupOutput, error) {
    var output gamelift.ResumeGameServerGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftSearchGameSessionsResult struct {
	Result workflow.Future
}

func (r *GameliftSearchGameSessionsResult) Get(ctx workflow.Context) (*gamelift.SearchGameSessionsOutput, error) {
    var output gamelift.SearchGameSessionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStartFleetActionsResult struct {
	Result workflow.Future
}

func (r *GameliftStartFleetActionsResult) Get(ctx workflow.Context) (*gamelift.StartFleetActionsOutput, error) {
    var output gamelift.StartFleetActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStartGameSessionPlacementResult struct {
	Result workflow.Future
}

func (r *GameliftStartGameSessionPlacementResult) Get(ctx workflow.Context) (*gamelift.StartGameSessionPlacementOutput, error) {
    var output gamelift.StartGameSessionPlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStartMatchBackfillResult struct {
	Result workflow.Future
}

func (r *GameliftStartMatchBackfillResult) Get(ctx workflow.Context) (*gamelift.StartMatchBackfillOutput, error) {
    var output gamelift.StartMatchBackfillOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStartMatchmakingResult struct {
	Result workflow.Future
}

func (r *GameliftStartMatchmakingResult) Get(ctx workflow.Context) (*gamelift.StartMatchmakingOutput, error) {
    var output gamelift.StartMatchmakingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStopFleetActionsResult struct {
	Result workflow.Future
}

func (r *GameliftStopFleetActionsResult) Get(ctx workflow.Context) (*gamelift.StopFleetActionsOutput, error) {
    var output gamelift.StopFleetActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStopGameSessionPlacementResult struct {
	Result workflow.Future
}

func (r *GameliftStopGameSessionPlacementResult) Get(ctx workflow.Context) (*gamelift.StopGameSessionPlacementOutput, error) {
    var output gamelift.StopGameSessionPlacementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftStopMatchmakingResult struct {
	Result workflow.Future
}

func (r *GameliftStopMatchmakingResult) Get(ctx workflow.Context) (*gamelift.StopMatchmakingOutput, error) {
    var output gamelift.StopMatchmakingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftSuspendGameServerGroupResult struct {
	Result workflow.Future
}

func (r *GameliftSuspendGameServerGroupResult) Get(ctx workflow.Context) (*gamelift.SuspendGameServerGroupOutput, error) {
    var output gamelift.SuspendGameServerGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftTagResourceResult struct {
	Result workflow.Future
}

func (r *GameliftTagResourceResult) Get(ctx workflow.Context) (*gamelift.TagResourceOutput, error) {
    var output gamelift.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUntagResourceResult struct {
	Result workflow.Future
}

func (r *GameliftUntagResourceResult) Get(ctx workflow.Context) (*gamelift.UntagResourceOutput, error) {
    var output gamelift.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateAliasResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateAliasResult) Get(ctx workflow.Context) (*gamelift.UpdateAliasOutput, error) {
    var output gamelift.UpdateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateBuildResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateBuildResult) Get(ctx workflow.Context) (*gamelift.UpdateBuildOutput, error) {
    var output gamelift.UpdateBuildOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateFleetAttributesResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateFleetAttributesResult) Get(ctx workflow.Context) (*gamelift.UpdateFleetAttributesOutput, error) {
    var output gamelift.UpdateFleetAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateFleetCapacityResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateFleetCapacityResult) Get(ctx workflow.Context) (*gamelift.UpdateFleetCapacityOutput, error) {
    var output gamelift.UpdateFleetCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateFleetPortSettingsResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateFleetPortSettingsResult) Get(ctx workflow.Context) (*gamelift.UpdateFleetPortSettingsOutput, error) {
    var output gamelift.UpdateFleetPortSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateGameServerResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateGameServerResult) Get(ctx workflow.Context) (*gamelift.UpdateGameServerOutput, error) {
    var output gamelift.UpdateGameServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateGameServerGroupResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateGameServerGroupResult) Get(ctx workflow.Context) (*gamelift.UpdateGameServerGroupOutput, error) {
    var output gamelift.UpdateGameServerGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateGameSessionResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateGameSessionResult) Get(ctx workflow.Context) (*gamelift.UpdateGameSessionOutput, error) {
    var output gamelift.UpdateGameSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateGameSessionQueueResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateGameSessionQueueResult) Get(ctx workflow.Context) (*gamelift.UpdateGameSessionQueueOutput, error) {
    var output gamelift.UpdateGameSessionQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateMatchmakingConfigurationResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateMatchmakingConfigurationResult) Get(ctx workflow.Context) (*gamelift.UpdateMatchmakingConfigurationOutput, error) {
    var output gamelift.UpdateMatchmakingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateRuntimeConfigurationResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateRuntimeConfigurationResult) Get(ctx workflow.Context) (*gamelift.UpdateRuntimeConfigurationOutput, error) {
    var output gamelift.UpdateRuntimeConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftUpdateScriptResult struct {
	Result workflow.Future
}

func (r *GameliftUpdateScriptResult) Get(ctx workflow.Context) (*gamelift.UpdateScriptOutput, error) {
    var output gamelift.UpdateScriptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameliftValidateMatchmakingRuleSetResult struct {
	Result workflow.Future
}

func (r *GameliftValidateMatchmakingRuleSetResult) Get(ctx workflow.Context) (*gamelift.ValidateMatchmakingRuleSetOutput, error) {
    var output gamelift.ValidateMatchmakingRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GameLiftStub struct {
    activities awsactivities.GameLiftActivities
}

func NewGameLiftStub() GameLiftClient {
    return &GameLiftStub{}
}

func (a *GameLiftStub) AcceptMatch(ctx workflow.Context, input *gamelift.AcceptMatchInput) (*gamelift.AcceptMatchOutput, error) {
    var output gamelift.AcceptMatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptMatch, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) AcceptMatchAsync(ctx workflow.Context, input *gamelift.AcceptMatchInput) *GameliftAcceptMatchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptMatch, input)
    return &GameliftAcceptMatchResult{Result: future}
}

func (a *GameLiftStub) ClaimGameServer(ctx workflow.Context, input *gamelift.ClaimGameServerInput) (*gamelift.ClaimGameServerOutput, error) {
    var output gamelift.ClaimGameServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ClaimGameServer, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ClaimGameServerAsync(ctx workflow.Context, input *gamelift.ClaimGameServerInput) *GameliftClaimGameServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ClaimGameServer, input)
    return &GameliftClaimGameServerResult{Result: future}
}

func (a *GameLiftStub) CreateAlias(ctx workflow.Context, input *gamelift.CreateAliasInput) (*gamelift.CreateAliasOutput, error) {
    var output gamelift.CreateAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateAliasAsync(ctx workflow.Context, input *gamelift.CreateAliasInput) *GameliftCreateAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAlias, input)
    return &GameliftCreateAliasResult{Result: future}
}

func (a *GameLiftStub) CreateBuild(ctx workflow.Context, input *gamelift.CreateBuildInput) (*gamelift.CreateBuildOutput, error) {
    var output gamelift.CreateBuildOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBuild, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateBuildAsync(ctx workflow.Context, input *gamelift.CreateBuildInput) *GameliftCreateBuildResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBuild, input)
    return &GameliftCreateBuildResult{Result: future}
}

func (a *GameLiftStub) CreateFleet(ctx workflow.Context, input *gamelift.CreateFleetInput) (*gamelift.CreateFleetOutput, error) {
    var output gamelift.CreateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateFleetAsync(ctx workflow.Context, input *gamelift.CreateFleetInput) *GameliftCreateFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input)
    return &GameliftCreateFleetResult{Result: future}
}

func (a *GameLiftStub) CreateGameServerGroup(ctx workflow.Context, input *gamelift.CreateGameServerGroupInput) (*gamelift.CreateGameServerGroupOutput, error) {
    var output gamelift.CreateGameServerGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGameServerGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateGameServerGroupAsync(ctx workflow.Context, input *gamelift.CreateGameServerGroupInput) *GameliftCreateGameServerGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGameServerGroup, input)
    return &GameliftCreateGameServerGroupResult{Result: future}
}

func (a *GameLiftStub) CreateGameSession(ctx workflow.Context, input *gamelift.CreateGameSessionInput) (*gamelift.CreateGameSessionOutput, error) {
    var output gamelift.CreateGameSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGameSession, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateGameSessionAsync(ctx workflow.Context, input *gamelift.CreateGameSessionInput) *GameliftCreateGameSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGameSession, input)
    return &GameliftCreateGameSessionResult{Result: future}
}

func (a *GameLiftStub) CreateGameSessionQueue(ctx workflow.Context, input *gamelift.CreateGameSessionQueueInput) (*gamelift.CreateGameSessionQueueOutput, error) {
    var output gamelift.CreateGameSessionQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGameSessionQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateGameSessionQueueAsync(ctx workflow.Context, input *gamelift.CreateGameSessionQueueInput) *GameliftCreateGameSessionQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGameSessionQueue, input)
    return &GameliftCreateGameSessionQueueResult{Result: future}
}

func (a *GameLiftStub) CreateMatchmakingConfiguration(ctx workflow.Context, input *gamelift.CreateMatchmakingConfigurationInput) (*gamelift.CreateMatchmakingConfigurationOutput, error) {
    var output gamelift.CreateMatchmakingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMatchmakingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateMatchmakingConfigurationAsync(ctx workflow.Context, input *gamelift.CreateMatchmakingConfigurationInput) *GameliftCreateMatchmakingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMatchmakingConfiguration, input)
    return &GameliftCreateMatchmakingConfigurationResult{Result: future}
}

func (a *GameLiftStub) CreateMatchmakingRuleSet(ctx workflow.Context, input *gamelift.CreateMatchmakingRuleSetInput) (*gamelift.CreateMatchmakingRuleSetOutput, error) {
    var output gamelift.CreateMatchmakingRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMatchmakingRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateMatchmakingRuleSetAsync(ctx workflow.Context, input *gamelift.CreateMatchmakingRuleSetInput) *GameliftCreateMatchmakingRuleSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMatchmakingRuleSet, input)
    return &GameliftCreateMatchmakingRuleSetResult{Result: future}
}

func (a *GameLiftStub) CreatePlayerSession(ctx workflow.Context, input *gamelift.CreatePlayerSessionInput) (*gamelift.CreatePlayerSessionOutput, error) {
    var output gamelift.CreatePlayerSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlayerSession, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreatePlayerSessionAsync(ctx workflow.Context, input *gamelift.CreatePlayerSessionInput) *GameliftCreatePlayerSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePlayerSession, input)
    return &GameliftCreatePlayerSessionResult{Result: future}
}

func (a *GameLiftStub) CreatePlayerSessions(ctx workflow.Context, input *gamelift.CreatePlayerSessionsInput) (*gamelift.CreatePlayerSessionsOutput, error) {
    var output gamelift.CreatePlayerSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlayerSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreatePlayerSessionsAsync(ctx workflow.Context, input *gamelift.CreatePlayerSessionsInput) *GameliftCreatePlayerSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePlayerSessions, input)
    return &GameliftCreatePlayerSessionsResult{Result: future}
}

func (a *GameLiftStub) CreateScript(ctx workflow.Context, input *gamelift.CreateScriptInput) (*gamelift.CreateScriptOutput, error) {
    var output gamelift.CreateScriptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateScript, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateScriptAsync(ctx workflow.Context, input *gamelift.CreateScriptInput) *GameliftCreateScriptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateScript, input)
    return &GameliftCreateScriptResult{Result: future}
}

func (a *GameLiftStub) CreateVpcPeeringAuthorization(ctx workflow.Context, input *gamelift.CreateVpcPeeringAuthorizationInput) (*gamelift.CreateVpcPeeringAuthorizationOutput, error) {
    var output gamelift.CreateVpcPeeringAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcPeeringAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateVpcPeeringAuthorizationAsync(ctx workflow.Context, input *gamelift.CreateVpcPeeringAuthorizationInput) *GameliftCreateVpcPeeringAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcPeeringAuthorization, input)
    return &GameliftCreateVpcPeeringAuthorizationResult{Result: future}
}

func (a *GameLiftStub) CreateVpcPeeringConnection(ctx workflow.Context, input *gamelift.CreateVpcPeeringConnectionInput) (*gamelift.CreateVpcPeeringConnectionOutput, error) {
    var output gamelift.CreateVpcPeeringConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVpcPeeringConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) CreateVpcPeeringConnectionAsync(ctx workflow.Context, input *gamelift.CreateVpcPeeringConnectionInput) *GameliftCreateVpcPeeringConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVpcPeeringConnection, input)
    return &GameliftCreateVpcPeeringConnectionResult{Result: future}
}

func (a *GameLiftStub) DeleteAlias(ctx workflow.Context, input *gamelift.DeleteAliasInput) (*gamelift.DeleteAliasOutput, error) {
    var output gamelift.DeleteAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteAliasAsync(ctx workflow.Context, input *gamelift.DeleteAliasInput) *GameliftDeleteAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAlias, input)
    return &GameliftDeleteAliasResult{Result: future}
}

func (a *GameLiftStub) DeleteBuild(ctx workflow.Context, input *gamelift.DeleteBuildInput) (*gamelift.DeleteBuildOutput, error) {
    var output gamelift.DeleteBuildOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBuild, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteBuildAsync(ctx workflow.Context, input *gamelift.DeleteBuildInput) *GameliftDeleteBuildResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBuild, input)
    return &GameliftDeleteBuildResult{Result: future}
}

func (a *GameLiftStub) DeleteFleet(ctx workflow.Context, input *gamelift.DeleteFleetInput) (*gamelift.DeleteFleetOutput, error) {
    var output gamelift.DeleteFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteFleetAsync(ctx workflow.Context, input *gamelift.DeleteFleetInput) *GameliftDeleteFleetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFleet, input)
    return &GameliftDeleteFleetResult{Result: future}
}

func (a *GameLiftStub) DeleteGameServerGroup(ctx workflow.Context, input *gamelift.DeleteGameServerGroupInput) (*gamelift.DeleteGameServerGroupOutput, error) {
    var output gamelift.DeleteGameServerGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGameServerGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteGameServerGroupAsync(ctx workflow.Context, input *gamelift.DeleteGameServerGroupInput) *GameliftDeleteGameServerGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGameServerGroup, input)
    return &GameliftDeleteGameServerGroupResult{Result: future}
}

func (a *GameLiftStub) DeleteGameSessionQueue(ctx workflow.Context, input *gamelift.DeleteGameSessionQueueInput) (*gamelift.DeleteGameSessionQueueOutput, error) {
    var output gamelift.DeleteGameSessionQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGameSessionQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteGameSessionQueueAsync(ctx workflow.Context, input *gamelift.DeleteGameSessionQueueInput) *GameliftDeleteGameSessionQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteGameSessionQueue, input)
    return &GameliftDeleteGameSessionQueueResult{Result: future}
}

func (a *GameLiftStub) DeleteMatchmakingConfiguration(ctx workflow.Context, input *gamelift.DeleteMatchmakingConfigurationInput) (*gamelift.DeleteMatchmakingConfigurationOutput, error) {
    var output gamelift.DeleteMatchmakingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMatchmakingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteMatchmakingConfigurationAsync(ctx workflow.Context, input *gamelift.DeleteMatchmakingConfigurationInput) *GameliftDeleteMatchmakingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMatchmakingConfiguration, input)
    return &GameliftDeleteMatchmakingConfigurationResult{Result: future}
}

func (a *GameLiftStub) DeleteMatchmakingRuleSet(ctx workflow.Context, input *gamelift.DeleteMatchmakingRuleSetInput) (*gamelift.DeleteMatchmakingRuleSetOutput, error) {
    var output gamelift.DeleteMatchmakingRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMatchmakingRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteMatchmakingRuleSetAsync(ctx workflow.Context, input *gamelift.DeleteMatchmakingRuleSetInput) *GameliftDeleteMatchmakingRuleSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMatchmakingRuleSet, input)
    return &GameliftDeleteMatchmakingRuleSetResult{Result: future}
}

func (a *GameLiftStub) DeleteScalingPolicy(ctx workflow.Context, input *gamelift.DeleteScalingPolicyInput) (*gamelift.DeleteScalingPolicyOutput, error) {
    var output gamelift.DeleteScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteScalingPolicyAsync(ctx workflow.Context, input *gamelift.DeleteScalingPolicyInput) *GameliftDeleteScalingPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteScalingPolicy, input)
    return &GameliftDeleteScalingPolicyResult{Result: future}
}

func (a *GameLiftStub) DeleteScript(ctx workflow.Context, input *gamelift.DeleteScriptInput) (*gamelift.DeleteScriptOutput, error) {
    var output gamelift.DeleteScriptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteScript, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteScriptAsync(ctx workflow.Context, input *gamelift.DeleteScriptInput) *GameliftDeleteScriptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteScript, input)
    return &GameliftDeleteScriptResult{Result: future}
}

func (a *GameLiftStub) DeleteVpcPeeringAuthorization(ctx workflow.Context, input *gamelift.DeleteVpcPeeringAuthorizationInput) (*gamelift.DeleteVpcPeeringAuthorizationOutput, error) {
    var output gamelift.DeleteVpcPeeringAuthorizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcPeeringAuthorization, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteVpcPeeringAuthorizationAsync(ctx workflow.Context, input *gamelift.DeleteVpcPeeringAuthorizationInput) *GameliftDeleteVpcPeeringAuthorizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcPeeringAuthorization, input)
    return &GameliftDeleteVpcPeeringAuthorizationResult{Result: future}
}

func (a *GameLiftStub) DeleteVpcPeeringConnection(ctx workflow.Context, input *gamelift.DeleteVpcPeeringConnectionInput) (*gamelift.DeleteVpcPeeringConnectionOutput, error) {
    var output gamelift.DeleteVpcPeeringConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcPeeringConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeleteVpcPeeringConnectionAsync(ctx workflow.Context, input *gamelift.DeleteVpcPeeringConnectionInput) *GameliftDeleteVpcPeeringConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVpcPeeringConnection, input)
    return &GameliftDeleteVpcPeeringConnectionResult{Result: future}
}

func (a *GameLiftStub) DeregisterGameServer(ctx workflow.Context, input *gamelift.DeregisterGameServerInput) (*gamelift.DeregisterGameServerOutput, error) {
    var output gamelift.DeregisterGameServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterGameServer, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DeregisterGameServerAsync(ctx workflow.Context, input *gamelift.DeregisterGameServerInput) *GameliftDeregisterGameServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterGameServer, input)
    return &GameliftDeregisterGameServerResult{Result: future}
}

func (a *GameLiftStub) DescribeAlias(ctx workflow.Context, input *gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error) {
    var output gamelift.DescribeAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeAliasAsync(ctx workflow.Context, input *gamelift.DescribeAliasInput) *GameliftDescribeAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAlias, input)
    return &GameliftDescribeAliasResult{Result: future}
}

func (a *GameLiftStub) DescribeBuild(ctx workflow.Context, input *gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error) {
    var output gamelift.DescribeBuildOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBuild, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeBuildAsync(ctx workflow.Context, input *gamelift.DescribeBuildInput) *GameliftDescribeBuildResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBuild, input)
    return &GameliftDescribeBuildResult{Result: future}
}

func (a *GameLiftStub) DescribeEC2InstanceLimits(ctx workflow.Context, input *gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
    var output gamelift.DescribeEC2InstanceLimitsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEC2InstanceLimits, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeEC2InstanceLimitsAsync(ctx workflow.Context, input *gamelift.DescribeEC2InstanceLimitsInput) *GameliftDescribeEC2InstanceLimitsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEC2InstanceLimits, input)
    return &GameliftDescribeEC2InstanceLimitsResult{Result: future}
}

func (a *GameLiftStub) DescribeFleetAttributes(ctx workflow.Context, input *gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error) {
    var output gamelift.DescribeFleetAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeFleetAttributesAsync(ctx workflow.Context, input *gamelift.DescribeFleetAttributesInput) *GameliftDescribeFleetAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetAttributes, input)
    return &GameliftDescribeFleetAttributesResult{Result: future}
}

func (a *GameLiftStub) DescribeFleetCapacity(ctx workflow.Context, input *gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error) {
    var output gamelift.DescribeFleetCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeFleetCapacityAsync(ctx workflow.Context, input *gamelift.DescribeFleetCapacityInput) *GameliftDescribeFleetCapacityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetCapacity, input)
    return &GameliftDescribeFleetCapacityResult{Result: future}
}

func (a *GameLiftStub) DescribeFleetEvents(ctx workflow.Context, input *gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error) {
    var output gamelift.DescribeFleetEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeFleetEventsAsync(ctx workflow.Context, input *gamelift.DescribeFleetEventsInput) *GameliftDescribeFleetEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetEvents, input)
    return &GameliftDescribeFleetEventsResult{Result: future}
}

func (a *GameLiftStub) DescribeFleetPortSettings(ctx workflow.Context, input *gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error) {
    var output gamelift.DescribeFleetPortSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetPortSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeFleetPortSettingsAsync(ctx workflow.Context, input *gamelift.DescribeFleetPortSettingsInput) *GameliftDescribeFleetPortSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetPortSettings, input)
    return &GameliftDescribeFleetPortSettingsResult{Result: future}
}

func (a *GameLiftStub) DescribeFleetUtilization(ctx workflow.Context, input *gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error) {
    var output gamelift.DescribeFleetUtilizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetUtilization, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeFleetUtilizationAsync(ctx workflow.Context, input *gamelift.DescribeFleetUtilizationInput) *GameliftDescribeFleetUtilizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFleetUtilization, input)
    return &GameliftDescribeFleetUtilizationResult{Result: future}
}

func (a *GameLiftStub) DescribeGameServer(ctx workflow.Context, input *gamelift.DescribeGameServerInput) (*gamelift.DescribeGameServerOutput, error) {
    var output gamelift.DescribeGameServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameServer, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameServerAsync(ctx workflow.Context, input *gamelift.DescribeGameServerInput) *GameliftDescribeGameServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameServer, input)
    return &GameliftDescribeGameServerResult{Result: future}
}

func (a *GameLiftStub) DescribeGameServerGroup(ctx workflow.Context, input *gamelift.DescribeGameServerGroupInput) (*gamelift.DescribeGameServerGroupOutput, error) {
    var output gamelift.DescribeGameServerGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameServerGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameServerGroupAsync(ctx workflow.Context, input *gamelift.DescribeGameServerGroupInput) *GameliftDescribeGameServerGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameServerGroup, input)
    return &GameliftDescribeGameServerGroupResult{Result: future}
}

func (a *GameLiftStub) DescribeGameServerInstances(ctx workflow.Context, input *gamelift.DescribeGameServerInstancesInput) (*gamelift.DescribeGameServerInstancesOutput, error) {
    var output gamelift.DescribeGameServerInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameServerInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameServerInstancesAsync(ctx workflow.Context, input *gamelift.DescribeGameServerInstancesInput) *GameliftDescribeGameServerInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameServerInstances, input)
    return &GameliftDescribeGameServerInstancesResult{Result: future}
}

func (a *GameLiftStub) DescribeGameSessionDetails(ctx workflow.Context, input *gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error) {
    var output gamelift.DescribeGameSessionDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessionDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameSessionDetailsAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionDetailsInput) *GameliftDescribeGameSessionDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessionDetails, input)
    return &GameliftDescribeGameSessionDetailsResult{Result: future}
}

func (a *GameLiftStub) DescribeGameSessionPlacement(ctx workflow.Context, input *gamelift.DescribeGameSessionPlacementInput) (*gamelift.DescribeGameSessionPlacementOutput, error) {
    var output gamelift.DescribeGameSessionPlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessionPlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameSessionPlacementAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionPlacementInput) *GameliftDescribeGameSessionPlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessionPlacement, input)
    return &GameliftDescribeGameSessionPlacementResult{Result: future}
}

func (a *GameLiftStub) DescribeGameSessionQueues(ctx workflow.Context, input *gamelift.DescribeGameSessionQueuesInput) (*gamelift.DescribeGameSessionQueuesOutput, error) {
    var output gamelift.DescribeGameSessionQueuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessionQueues, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameSessionQueuesAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionQueuesInput) *GameliftDescribeGameSessionQueuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessionQueues, input)
    return &GameliftDescribeGameSessionQueuesResult{Result: future}
}

func (a *GameLiftStub) DescribeGameSessions(ctx workflow.Context, input *gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error) {
    var output gamelift.DescribeGameSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeGameSessionsAsync(ctx workflow.Context, input *gamelift.DescribeGameSessionsInput) *GameliftDescribeGameSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeGameSessions, input)
    return &GameliftDescribeGameSessionsResult{Result: future}
}

func (a *GameLiftStub) DescribeInstances(ctx workflow.Context, input *gamelift.DescribeInstancesInput) (*gamelift.DescribeInstancesOutput, error) {
    var output gamelift.DescribeInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeInstancesAsync(ctx workflow.Context, input *gamelift.DescribeInstancesInput) *GameliftDescribeInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstances, input)
    return &GameliftDescribeInstancesResult{Result: future}
}

func (a *GameLiftStub) DescribeMatchmaking(ctx workflow.Context, input *gamelift.DescribeMatchmakingInput) (*gamelift.DescribeMatchmakingOutput, error) {
    var output gamelift.DescribeMatchmakingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMatchmaking, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeMatchmakingAsync(ctx workflow.Context, input *gamelift.DescribeMatchmakingInput) *GameliftDescribeMatchmakingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMatchmaking, input)
    return &GameliftDescribeMatchmakingResult{Result: future}
}

func (a *GameLiftStub) DescribeMatchmakingConfigurations(ctx workflow.Context, input *gamelift.DescribeMatchmakingConfigurationsInput) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
    var output gamelift.DescribeMatchmakingConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMatchmakingConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeMatchmakingConfigurationsAsync(ctx workflow.Context, input *gamelift.DescribeMatchmakingConfigurationsInput) *GameliftDescribeMatchmakingConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMatchmakingConfigurations, input)
    return &GameliftDescribeMatchmakingConfigurationsResult{Result: future}
}

func (a *GameLiftStub) DescribeMatchmakingRuleSets(ctx workflow.Context, input *gamelift.DescribeMatchmakingRuleSetsInput) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
    var output gamelift.DescribeMatchmakingRuleSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMatchmakingRuleSets, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeMatchmakingRuleSetsAsync(ctx workflow.Context, input *gamelift.DescribeMatchmakingRuleSetsInput) *GameliftDescribeMatchmakingRuleSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMatchmakingRuleSets, input)
    return &GameliftDescribeMatchmakingRuleSetsResult{Result: future}
}

func (a *GameLiftStub) DescribePlayerSessions(ctx workflow.Context, input *gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error) {
    var output gamelift.DescribePlayerSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePlayerSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribePlayerSessionsAsync(ctx workflow.Context, input *gamelift.DescribePlayerSessionsInput) *GameliftDescribePlayerSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePlayerSessions, input)
    return &GameliftDescribePlayerSessionsResult{Result: future}
}

func (a *GameLiftStub) DescribeRuntimeConfiguration(ctx workflow.Context, input *gamelift.DescribeRuntimeConfigurationInput) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
    var output gamelift.DescribeRuntimeConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRuntimeConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeRuntimeConfigurationAsync(ctx workflow.Context, input *gamelift.DescribeRuntimeConfigurationInput) *GameliftDescribeRuntimeConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRuntimeConfiguration, input)
    return &GameliftDescribeRuntimeConfigurationResult{Result: future}
}

func (a *GameLiftStub) DescribeScalingPolicies(ctx workflow.Context, input *gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error) {
    var output gamelift.DescribeScalingPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeScalingPoliciesAsync(ctx workflow.Context, input *gamelift.DescribeScalingPoliciesInput) *GameliftDescribeScalingPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScalingPolicies, input)
    return &GameliftDescribeScalingPoliciesResult{Result: future}
}

func (a *GameLiftStub) DescribeScript(ctx workflow.Context, input *gamelift.DescribeScriptInput) (*gamelift.DescribeScriptOutput, error) {
    var output gamelift.DescribeScriptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeScript, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeScriptAsync(ctx workflow.Context, input *gamelift.DescribeScriptInput) *GameliftDescribeScriptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeScript, input)
    return &GameliftDescribeScriptResult{Result: future}
}

func (a *GameLiftStub) DescribeVpcPeeringAuthorizations(ctx workflow.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
    var output gamelift.DescribeVpcPeeringAuthorizationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcPeeringAuthorizations, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeVpcPeeringAuthorizationsAsync(ctx workflow.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput) *GameliftDescribeVpcPeeringAuthorizationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcPeeringAuthorizations, input)
    return &GameliftDescribeVpcPeeringAuthorizationsResult{Result: future}
}

func (a *GameLiftStub) DescribeVpcPeeringConnections(ctx workflow.Context, input *gamelift.DescribeVpcPeeringConnectionsInput) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
    var output gamelift.DescribeVpcPeeringConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcPeeringConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) DescribeVpcPeeringConnectionsAsync(ctx workflow.Context, input *gamelift.DescribeVpcPeeringConnectionsInput) *GameliftDescribeVpcPeeringConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVpcPeeringConnections, input)
    return &GameliftDescribeVpcPeeringConnectionsResult{Result: future}
}

func (a *GameLiftStub) GetGameSessionLogUrl(ctx workflow.Context, input *gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error) {
    var output gamelift.GetGameSessionLogUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGameSessionLogUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) GetGameSessionLogUrlAsync(ctx workflow.Context, input *gamelift.GetGameSessionLogUrlInput) *GameliftGetGameSessionLogUrlResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetGameSessionLogUrl, input)
    return &GameliftGetGameSessionLogUrlResult{Result: future}
}

func (a *GameLiftStub) GetInstanceAccess(ctx workflow.Context, input *gamelift.GetInstanceAccessInput) (*gamelift.GetInstanceAccessOutput, error) {
    var output gamelift.GetInstanceAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInstanceAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) GetInstanceAccessAsync(ctx workflow.Context, input *gamelift.GetInstanceAccessInput) *GameliftGetInstanceAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInstanceAccess, input)
    return &GameliftGetInstanceAccessResult{Result: future}
}

func (a *GameLiftStub) ListAliases(ctx workflow.Context, input *gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error) {
    var output gamelift.ListAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListAliasesAsync(ctx workflow.Context, input *gamelift.ListAliasesInput) *GameliftListAliasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAliases, input)
    return &GameliftListAliasesResult{Result: future}
}

func (a *GameLiftStub) ListBuilds(ctx workflow.Context, input *gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error) {
    var output gamelift.ListBuildsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBuilds, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListBuildsAsync(ctx workflow.Context, input *gamelift.ListBuildsInput) *GameliftListBuildsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBuilds, input)
    return &GameliftListBuildsResult{Result: future}
}

func (a *GameLiftStub) ListFleets(ctx workflow.Context, input *gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error) {
    var output gamelift.ListFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListFleetsAsync(ctx workflow.Context, input *gamelift.ListFleetsInput) *GameliftListFleetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFleets, input)
    return &GameliftListFleetsResult{Result: future}
}

func (a *GameLiftStub) ListGameServerGroups(ctx workflow.Context, input *gamelift.ListGameServerGroupsInput) (*gamelift.ListGameServerGroupsOutput, error) {
    var output gamelift.ListGameServerGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGameServerGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListGameServerGroupsAsync(ctx workflow.Context, input *gamelift.ListGameServerGroupsInput) *GameliftListGameServerGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGameServerGroups, input)
    return &GameliftListGameServerGroupsResult{Result: future}
}

func (a *GameLiftStub) ListGameServers(ctx workflow.Context, input *gamelift.ListGameServersInput) (*gamelift.ListGameServersOutput, error) {
    var output gamelift.ListGameServersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGameServers, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListGameServersAsync(ctx workflow.Context, input *gamelift.ListGameServersInput) *GameliftListGameServersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListGameServers, input)
    return &GameliftListGameServersResult{Result: future}
}

func (a *GameLiftStub) ListScripts(ctx workflow.Context, input *gamelift.ListScriptsInput) (*gamelift.ListScriptsOutput, error) {
    var output gamelift.ListScriptsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListScripts, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListScriptsAsync(ctx workflow.Context, input *gamelift.ListScriptsInput) *GameliftListScriptsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListScripts, input)
    return &GameliftListScriptsResult{Result: future}
}

func (a *GameLiftStub) ListTagsForResource(ctx workflow.Context, input *gamelift.ListTagsForResourceInput) (*gamelift.ListTagsForResourceOutput, error) {
    var output gamelift.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ListTagsForResourceAsync(ctx workflow.Context, input *gamelift.ListTagsForResourceInput) *GameliftListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &GameliftListTagsForResourceResult{Result: future}
}

func (a *GameLiftStub) PutScalingPolicy(ctx workflow.Context, input *gamelift.PutScalingPolicyInput) (*gamelift.PutScalingPolicyOutput, error) {
    var output gamelift.PutScalingPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutScalingPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) PutScalingPolicyAsync(ctx workflow.Context, input *gamelift.PutScalingPolicyInput) *GameliftPutScalingPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutScalingPolicy, input)
    return &GameliftPutScalingPolicyResult{Result: future}
}

func (a *GameLiftStub) RegisterGameServer(ctx workflow.Context, input *gamelift.RegisterGameServerInput) (*gamelift.RegisterGameServerOutput, error) {
    var output gamelift.RegisterGameServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterGameServer, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) RegisterGameServerAsync(ctx workflow.Context, input *gamelift.RegisterGameServerInput) *GameliftRegisterGameServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterGameServer, input)
    return &GameliftRegisterGameServerResult{Result: future}
}

func (a *GameLiftStub) RequestUploadCredentials(ctx workflow.Context, input *gamelift.RequestUploadCredentialsInput) (*gamelift.RequestUploadCredentialsOutput, error) {
    var output gamelift.RequestUploadCredentialsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestUploadCredentials, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) RequestUploadCredentialsAsync(ctx workflow.Context, input *gamelift.RequestUploadCredentialsInput) *GameliftRequestUploadCredentialsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RequestUploadCredentials, input)
    return &GameliftRequestUploadCredentialsResult{Result: future}
}

func (a *GameLiftStub) ResolveAlias(ctx workflow.Context, input *gamelift.ResolveAliasInput) (*gamelift.ResolveAliasOutput, error) {
    var output gamelift.ResolveAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResolveAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ResolveAliasAsync(ctx workflow.Context, input *gamelift.ResolveAliasInput) *GameliftResolveAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResolveAlias, input)
    return &GameliftResolveAliasResult{Result: future}
}

func (a *GameLiftStub) ResumeGameServerGroup(ctx workflow.Context, input *gamelift.ResumeGameServerGroupInput) (*gamelift.ResumeGameServerGroupOutput, error) {
    var output gamelift.ResumeGameServerGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResumeGameServerGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ResumeGameServerGroupAsync(ctx workflow.Context, input *gamelift.ResumeGameServerGroupInput) *GameliftResumeGameServerGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResumeGameServerGroup, input)
    return &GameliftResumeGameServerGroupResult{Result: future}
}

func (a *GameLiftStub) SearchGameSessions(ctx workflow.Context, input *gamelift.SearchGameSessionsInput) (*gamelift.SearchGameSessionsOutput, error) {
    var output gamelift.SearchGameSessionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchGameSessions, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) SearchGameSessionsAsync(ctx workflow.Context, input *gamelift.SearchGameSessionsInput) *GameliftSearchGameSessionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchGameSessions, input)
    return &GameliftSearchGameSessionsResult{Result: future}
}

func (a *GameLiftStub) StartFleetActions(ctx workflow.Context, input *gamelift.StartFleetActionsInput) (*gamelift.StartFleetActionsOutput, error) {
    var output gamelift.StartFleetActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartFleetActions, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StartFleetActionsAsync(ctx workflow.Context, input *gamelift.StartFleetActionsInput) *GameliftStartFleetActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartFleetActions, input)
    return &GameliftStartFleetActionsResult{Result: future}
}

func (a *GameLiftStub) StartGameSessionPlacement(ctx workflow.Context, input *gamelift.StartGameSessionPlacementInput) (*gamelift.StartGameSessionPlacementOutput, error) {
    var output gamelift.StartGameSessionPlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartGameSessionPlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StartGameSessionPlacementAsync(ctx workflow.Context, input *gamelift.StartGameSessionPlacementInput) *GameliftStartGameSessionPlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartGameSessionPlacement, input)
    return &GameliftStartGameSessionPlacementResult{Result: future}
}

func (a *GameLiftStub) StartMatchBackfill(ctx workflow.Context, input *gamelift.StartMatchBackfillInput) (*gamelift.StartMatchBackfillOutput, error) {
    var output gamelift.StartMatchBackfillOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMatchBackfill, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StartMatchBackfillAsync(ctx workflow.Context, input *gamelift.StartMatchBackfillInput) *GameliftStartMatchBackfillResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMatchBackfill, input)
    return &GameliftStartMatchBackfillResult{Result: future}
}

func (a *GameLiftStub) StartMatchmaking(ctx workflow.Context, input *gamelift.StartMatchmakingInput) (*gamelift.StartMatchmakingOutput, error) {
    var output gamelift.StartMatchmakingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMatchmaking, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StartMatchmakingAsync(ctx workflow.Context, input *gamelift.StartMatchmakingInput) *GameliftStartMatchmakingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMatchmaking, input)
    return &GameliftStartMatchmakingResult{Result: future}
}

func (a *GameLiftStub) StopFleetActions(ctx workflow.Context, input *gamelift.StopFleetActionsInput) (*gamelift.StopFleetActionsOutput, error) {
    var output gamelift.StopFleetActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopFleetActions, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StopFleetActionsAsync(ctx workflow.Context, input *gamelift.StopFleetActionsInput) *GameliftStopFleetActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopFleetActions, input)
    return &GameliftStopFleetActionsResult{Result: future}
}

func (a *GameLiftStub) StopGameSessionPlacement(ctx workflow.Context, input *gamelift.StopGameSessionPlacementInput) (*gamelift.StopGameSessionPlacementOutput, error) {
    var output gamelift.StopGameSessionPlacementOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopGameSessionPlacement, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StopGameSessionPlacementAsync(ctx workflow.Context, input *gamelift.StopGameSessionPlacementInput) *GameliftStopGameSessionPlacementResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopGameSessionPlacement, input)
    return &GameliftStopGameSessionPlacementResult{Result: future}
}

func (a *GameLiftStub) StopMatchmaking(ctx workflow.Context, input *gamelift.StopMatchmakingInput) (*gamelift.StopMatchmakingOutput, error) {
    var output gamelift.StopMatchmakingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopMatchmaking, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) StopMatchmakingAsync(ctx workflow.Context, input *gamelift.StopMatchmakingInput) *GameliftStopMatchmakingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopMatchmaking, input)
    return &GameliftStopMatchmakingResult{Result: future}
}

func (a *GameLiftStub) SuspendGameServerGroup(ctx workflow.Context, input *gamelift.SuspendGameServerGroupInput) (*gamelift.SuspendGameServerGroupOutput, error) {
    var output gamelift.SuspendGameServerGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SuspendGameServerGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) SuspendGameServerGroupAsync(ctx workflow.Context, input *gamelift.SuspendGameServerGroupInput) *GameliftSuspendGameServerGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SuspendGameServerGroup, input)
    return &GameliftSuspendGameServerGroupResult{Result: future}
}

func (a *GameLiftStub) TagResource(ctx workflow.Context, input *gamelift.TagResourceInput) (*gamelift.TagResourceOutput, error) {
    var output gamelift.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) TagResourceAsync(ctx workflow.Context, input *gamelift.TagResourceInput) *GameliftTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &GameliftTagResourceResult{Result: future}
}

func (a *GameLiftStub) UntagResource(ctx workflow.Context, input *gamelift.UntagResourceInput) (*gamelift.UntagResourceOutput, error) {
    var output gamelift.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UntagResourceAsync(ctx workflow.Context, input *gamelift.UntagResourceInput) *GameliftUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &GameliftUntagResourceResult{Result: future}
}

func (a *GameLiftStub) UpdateAlias(ctx workflow.Context, input *gamelift.UpdateAliasInput) (*gamelift.UpdateAliasOutput, error) {
    var output gamelift.UpdateAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateAliasAsync(ctx workflow.Context, input *gamelift.UpdateAliasInput) *GameliftUpdateAliasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAlias, input)
    return &GameliftUpdateAliasResult{Result: future}
}

func (a *GameLiftStub) UpdateBuild(ctx workflow.Context, input *gamelift.UpdateBuildInput) (*gamelift.UpdateBuildOutput, error) {
    var output gamelift.UpdateBuildOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBuild, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateBuildAsync(ctx workflow.Context, input *gamelift.UpdateBuildInput) *GameliftUpdateBuildResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBuild, input)
    return &GameliftUpdateBuildResult{Result: future}
}

func (a *GameLiftStub) UpdateFleetAttributes(ctx workflow.Context, input *gamelift.UpdateFleetAttributesInput) (*gamelift.UpdateFleetAttributesOutput, error) {
    var output gamelift.UpdateFleetAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateFleetAttributesAsync(ctx workflow.Context, input *gamelift.UpdateFleetAttributesInput) *GameliftUpdateFleetAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetAttributes, input)
    return &GameliftUpdateFleetAttributesResult{Result: future}
}

func (a *GameLiftStub) UpdateFleetCapacity(ctx workflow.Context, input *gamelift.UpdateFleetCapacityInput) (*gamelift.UpdateFleetCapacityOutput, error) {
    var output gamelift.UpdateFleetCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateFleetCapacityAsync(ctx workflow.Context, input *gamelift.UpdateFleetCapacityInput) *GameliftUpdateFleetCapacityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetCapacity, input)
    return &GameliftUpdateFleetCapacityResult{Result: future}
}

func (a *GameLiftStub) UpdateFleetPortSettings(ctx workflow.Context, input *gamelift.UpdateFleetPortSettingsInput) (*gamelift.UpdateFleetPortSettingsOutput, error) {
    var output gamelift.UpdateFleetPortSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetPortSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateFleetPortSettingsAsync(ctx workflow.Context, input *gamelift.UpdateFleetPortSettingsInput) *GameliftUpdateFleetPortSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFleetPortSettings, input)
    return &GameliftUpdateFleetPortSettingsResult{Result: future}
}

func (a *GameLiftStub) UpdateGameServer(ctx workflow.Context, input *gamelift.UpdateGameServerInput) (*gamelift.UpdateGameServerOutput, error) {
    var output gamelift.UpdateGameServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGameServer, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateGameServerAsync(ctx workflow.Context, input *gamelift.UpdateGameServerInput) *GameliftUpdateGameServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGameServer, input)
    return &GameliftUpdateGameServerResult{Result: future}
}

func (a *GameLiftStub) UpdateGameServerGroup(ctx workflow.Context, input *gamelift.UpdateGameServerGroupInput) (*gamelift.UpdateGameServerGroupOutput, error) {
    var output gamelift.UpdateGameServerGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGameServerGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateGameServerGroupAsync(ctx workflow.Context, input *gamelift.UpdateGameServerGroupInput) *GameliftUpdateGameServerGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGameServerGroup, input)
    return &GameliftUpdateGameServerGroupResult{Result: future}
}

func (a *GameLiftStub) UpdateGameSession(ctx workflow.Context, input *gamelift.UpdateGameSessionInput) (*gamelift.UpdateGameSessionOutput, error) {
    var output gamelift.UpdateGameSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGameSession, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateGameSessionAsync(ctx workflow.Context, input *gamelift.UpdateGameSessionInput) *GameliftUpdateGameSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGameSession, input)
    return &GameliftUpdateGameSessionResult{Result: future}
}

func (a *GameLiftStub) UpdateGameSessionQueue(ctx workflow.Context, input *gamelift.UpdateGameSessionQueueInput) (*gamelift.UpdateGameSessionQueueOutput, error) {
    var output gamelift.UpdateGameSessionQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGameSessionQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateGameSessionQueueAsync(ctx workflow.Context, input *gamelift.UpdateGameSessionQueueInput) *GameliftUpdateGameSessionQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateGameSessionQueue, input)
    return &GameliftUpdateGameSessionQueueResult{Result: future}
}

func (a *GameLiftStub) UpdateMatchmakingConfiguration(ctx workflow.Context, input *gamelift.UpdateMatchmakingConfigurationInput) (*gamelift.UpdateMatchmakingConfigurationOutput, error) {
    var output gamelift.UpdateMatchmakingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMatchmakingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateMatchmakingConfigurationAsync(ctx workflow.Context, input *gamelift.UpdateMatchmakingConfigurationInput) *GameliftUpdateMatchmakingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMatchmakingConfiguration, input)
    return &GameliftUpdateMatchmakingConfigurationResult{Result: future}
}

func (a *GameLiftStub) UpdateRuntimeConfiguration(ctx workflow.Context, input *gamelift.UpdateRuntimeConfigurationInput) (*gamelift.UpdateRuntimeConfigurationOutput, error) {
    var output gamelift.UpdateRuntimeConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRuntimeConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateRuntimeConfigurationAsync(ctx workflow.Context, input *gamelift.UpdateRuntimeConfigurationInput) *GameliftUpdateRuntimeConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRuntimeConfiguration, input)
    return &GameliftUpdateRuntimeConfigurationResult{Result: future}
}

func (a *GameLiftStub) UpdateScript(ctx workflow.Context, input *gamelift.UpdateScriptInput) (*gamelift.UpdateScriptOutput, error) {
    var output gamelift.UpdateScriptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateScript, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) UpdateScriptAsync(ctx workflow.Context, input *gamelift.UpdateScriptInput) *GameliftUpdateScriptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateScript, input)
    return &GameliftUpdateScriptResult{Result: future}
}

func (a *GameLiftStub) ValidateMatchmakingRuleSet(ctx workflow.Context, input *gamelift.ValidateMatchmakingRuleSetInput) (*gamelift.ValidateMatchmakingRuleSetOutput, error) {
    var output gamelift.ValidateMatchmakingRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidateMatchmakingRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *GameLiftStub) ValidateMatchmakingRuleSetAsync(ctx workflow.Context, input *gamelift.ValidateMatchmakingRuleSetInput) *GameliftValidateMatchmakingRuleSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ValidateMatchmakingRuleSet, input)
    return &GameliftValidateMatchmakingRuleSetResult{Result: future}
}
