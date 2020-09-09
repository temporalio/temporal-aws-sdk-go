
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/gamelift/gameliftiface"
)

type GameLiftActivities struct {
	client gameliftiface.GameLiftAPI
}

func NewGameLiftActivities(session *session.Session, config... *aws.Config) *GameLiftActivities {
    client := gamelift.New(session, config...)
    return &GameLiftActivities{client: client}
}

func (a *GameLiftActivities) AcceptMatch(input *gamelift.AcceptMatchInput) (*gamelift.AcceptMatchOutput, error) {
    return a.client.AcceptMatch(input)
}

func (a *GameLiftActivities) ClaimGameServer(input *gamelift.ClaimGameServerInput) (*gamelift.ClaimGameServerOutput, error) {
    return a.client.ClaimGameServer(input)
}

func (a *GameLiftActivities) CreateAlias(input *gamelift.CreateAliasInput) (*gamelift.CreateAliasOutput, error) {
    return a.client.CreateAlias(input)
}

func (a *GameLiftActivities) CreateBuild(input *gamelift.CreateBuildInput) (*gamelift.CreateBuildOutput, error) {
    return a.client.CreateBuild(input)
}

func (a *GameLiftActivities) CreateFleet(input *gamelift.CreateFleetInput) (*gamelift.CreateFleetOutput, error) {
    return a.client.CreateFleet(input)
}

func (a *GameLiftActivities) CreateGameServerGroup(input *gamelift.CreateGameServerGroupInput) (*gamelift.CreateGameServerGroupOutput, error) {
    return a.client.CreateGameServerGroup(input)
}

func (a *GameLiftActivities) CreateGameSession(input *gamelift.CreateGameSessionInput) (*gamelift.CreateGameSessionOutput, error) {
    return a.client.CreateGameSession(input)
}

func (a *GameLiftActivities) CreateGameSessionQueue(input *gamelift.CreateGameSessionQueueInput) (*gamelift.CreateGameSessionQueueOutput, error) {
    return a.client.CreateGameSessionQueue(input)
}

func (a *GameLiftActivities) CreateMatchmakingConfiguration(input *gamelift.CreateMatchmakingConfigurationInput) (*gamelift.CreateMatchmakingConfigurationOutput, error) {
    return a.client.CreateMatchmakingConfiguration(input)
}

func (a *GameLiftActivities) CreateMatchmakingRuleSet(input *gamelift.CreateMatchmakingRuleSetInput) (*gamelift.CreateMatchmakingRuleSetOutput, error) {
    return a.client.CreateMatchmakingRuleSet(input)
}

func (a *GameLiftActivities) CreatePlayerSession(input *gamelift.CreatePlayerSessionInput) (*gamelift.CreatePlayerSessionOutput, error) {
    return a.client.CreatePlayerSession(input)
}

func (a *GameLiftActivities) CreatePlayerSessions(input *gamelift.CreatePlayerSessionsInput) (*gamelift.CreatePlayerSessionsOutput, error) {
    return a.client.CreatePlayerSessions(input)
}

func (a *GameLiftActivities) CreateScript(input *gamelift.CreateScriptInput) (*gamelift.CreateScriptOutput, error) {
    return a.client.CreateScript(input)
}

func (a *GameLiftActivities) CreateVpcPeeringAuthorization(input *gamelift.CreateVpcPeeringAuthorizationInput) (*gamelift.CreateVpcPeeringAuthorizationOutput, error) {
    return a.client.CreateVpcPeeringAuthorization(input)
}

func (a *GameLiftActivities) CreateVpcPeeringConnection(input *gamelift.CreateVpcPeeringConnectionInput) (*gamelift.CreateVpcPeeringConnectionOutput, error) {
    return a.client.CreateVpcPeeringConnection(input)
}

func (a *GameLiftActivities) DeleteAlias(input *gamelift.DeleteAliasInput) (*gamelift.DeleteAliasOutput, error) {
    return a.client.DeleteAlias(input)
}

func (a *GameLiftActivities) DeleteBuild(input *gamelift.DeleteBuildInput) (*gamelift.DeleteBuildOutput, error) {
    return a.client.DeleteBuild(input)
}

func (a *GameLiftActivities) DeleteFleet(input *gamelift.DeleteFleetInput) (*gamelift.DeleteFleetOutput, error) {
    return a.client.DeleteFleet(input)
}

func (a *GameLiftActivities) DeleteGameServerGroup(input *gamelift.DeleteGameServerGroupInput) (*gamelift.DeleteGameServerGroupOutput, error) {
    return a.client.DeleteGameServerGroup(input)
}

func (a *GameLiftActivities) DeleteGameSessionQueue(input *gamelift.DeleteGameSessionQueueInput) (*gamelift.DeleteGameSessionQueueOutput, error) {
    return a.client.DeleteGameSessionQueue(input)
}

func (a *GameLiftActivities) DeleteMatchmakingConfiguration(input *gamelift.DeleteMatchmakingConfigurationInput) (*gamelift.DeleteMatchmakingConfigurationOutput, error) {
    return a.client.DeleteMatchmakingConfiguration(input)
}

func (a *GameLiftActivities) DeleteMatchmakingRuleSet(input *gamelift.DeleteMatchmakingRuleSetInput) (*gamelift.DeleteMatchmakingRuleSetOutput, error) {
    return a.client.DeleteMatchmakingRuleSet(input)
}

func (a *GameLiftActivities) DeleteScalingPolicy(input *gamelift.DeleteScalingPolicyInput) (*gamelift.DeleteScalingPolicyOutput, error) {
    return a.client.DeleteScalingPolicy(input)
}

func (a *GameLiftActivities) DeleteScript(input *gamelift.DeleteScriptInput) (*gamelift.DeleteScriptOutput, error) {
    return a.client.DeleteScript(input)
}

func (a *GameLiftActivities) DeleteVpcPeeringAuthorization(input *gamelift.DeleteVpcPeeringAuthorizationInput) (*gamelift.DeleteVpcPeeringAuthorizationOutput, error) {
    return a.client.DeleteVpcPeeringAuthorization(input)
}

func (a *GameLiftActivities) DeleteVpcPeeringConnection(input *gamelift.DeleteVpcPeeringConnectionInput) (*gamelift.DeleteVpcPeeringConnectionOutput, error) {
    return a.client.DeleteVpcPeeringConnection(input)
}

func (a *GameLiftActivities) DeregisterGameServer(input *gamelift.DeregisterGameServerInput) (*gamelift.DeregisterGameServerOutput, error) {
    return a.client.DeregisterGameServer(input)
}

func (a *GameLiftActivities) DescribeAlias(input *gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error) {
    return a.client.DescribeAlias(input)
}

func (a *GameLiftActivities) DescribeBuild(input *gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error) {
    return a.client.DescribeBuild(input)
}

func (a *GameLiftActivities) DescribeEC2InstanceLimits(input *gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
    return a.client.DescribeEC2InstanceLimits(input)
}

func (a *GameLiftActivities) DescribeFleetAttributes(input *gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error) {
    return a.client.DescribeFleetAttributes(input)
}

func (a *GameLiftActivities) DescribeFleetCapacity(input *gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error) {
    return a.client.DescribeFleetCapacity(input)
}

func (a *GameLiftActivities) DescribeFleetEvents(input *gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error) {
    return a.client.DescribeFleetEvents(input)
}

func (a *GameLiftActivities) DescribeFleetPortSettings(input *gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error) {
    return a.client.DescribeFleetPortSettings(input)
}

func (a *GameLiftActivities) DescribeFleetUtilization(input *gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error) {
    return a.client.DescribeFleetUtilization(input)
}

func (a *GameLiftActivities) DescribeGameServer(input *gamelift.DescribeGameServerInput) (*gamelift.DescribeGameServerOutput, error) {
    return a.client.DescribeGameServer(input)
}

func (a *GameLiftActivities) DescribeGameServerGroup(input *gamelift.DescribeGameServerGroupInput) (*gamelift.DescribeGameServerGroupOutput, error) {
    return a.client.DescribeGameServerGroup(input)
}

func (a *GameLiftActivities) DescribeGameServerInstances(input *gamelift.DescribeGameServerInstancesInput) (*gamelift.DescribeGameServerInstancesOutput, error) {
    return a.client.DescribeGameServerInstances(input)
}

func (a *GameLiftActivities) DescribeGameSessionDetails(input *gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error) {
    return a.client.DescribeGameSessionDetails(input)
}

func (a *GameLiftActivities) DescribeGameSessionPlacement(input *gamelift.DescribeGameSessionPlacementInput) (*gamelift.DescribeGameSessionPlacementOutput, error) {
    return a.client.DescribeGameSessionPlacement(input)
}

func (a *GameLiftActivities) DescribeGameSessionQueues(input *gamelift.DescribeGameSessionQueuesInput) (*gamelift.DescribeGameSessionQueuesOutput, error) {
    return a.client.DescribeGameSessionQueues(input)
}

func (a *GameLiftActivities) DescribeGameSessions(input *gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error) {
    return a.client.DescribeGameSessions(input)
}

func (a *GameLiftActivities) DescribeInstances(input *gamelift.DescribeInstancesInput) (*gamelift.DescribeInstancesOutput, error) {
    return a.client.DescribeInstances(input)
}

func (a *GameLiftActivities) DescribeMatchmaking(input *gamelift.DescribeMatchmakingInput) (*gamelift.DescribeMatchmakingOutput, error) {
    return a.client.DescribeMatchmaking(input)
}

func (a *GameLiftActivities) DescribeMatchmakingConfigurations(input *gamelift.DescribeMatchmakingConfigurationsInput) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
    return a.client.DescribeMatchmakingConfigurations(input)
}

func (a *GameLiftActivities) DescribeMatchmakingRuleSets(input *gamelift.DescribeMatchmakingRuleSetsInput) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
    return a.client.DescribeMatchmakingRuleSets(input)
}

func (a *GameLiftActivities) DescribePlayerSessions(input *gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error) {
    return a.client.DescribePlayerSessions(input)
}

func (a *GameLiftActivities) DescribeRuntimeConfiguration(input *gamelift.DescribeRuntimeConfigurationInput) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
    return a.client.DescribeRuntimeConfiguration(input)
}

func (a *GameLiftActivities) DescribeScalingPolicies(input *gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error) {
    return a.client.DescribeScalingPolicies(input)
}

func (a *GameLiftActivities) DescribeScript(input *gamelift.DescribeScriptInput) (*gamelift.DescribeScriptOutput, error) {
    return a.client.DescribeScript(input)
}

func (a *GameLiftActivities) DescribeVpcPeeringAuthorizations(input *gamelift.DescribeVpcPeeringAuthorizationsInput) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
    return a.client.DescribeVpcPeeringAuthorizations(input)
}

func (a *GameLiftActivities) DescribeVpcPeeringConnections(input *gamelift.DescribeVpcPeeringConnectionsInput) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
    return a.client.DescribeVpcPeeringConnections(input)
}

func (a *GameLiftActivities) GetGameSessionLogUrl(input *gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error) {
    return a.client.GetGameSessionLogUrl(input)
}

func (a *GameLiftActivities) GetInstanceAccess(input *gamelift.GetInstanceAccessInput) (*gamelift.GetInstanceAccessOutput, error) {
    return a.client.GetInstanceAccess(input)
}

func (a *GameLiftActivities) ListAliases(input *gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error) {
    return a.client.ListAliases(input)
}

func (a *GameLiftActivities) ListBuilds(input *gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error) {
    return a.client.ListBuilds(input)
}

func (a *GameLiftActivities) ListFleets(input *gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error) {
    return a.client.ListFleets(input)
}

func (a *GameLiftActivities) ListGameServerGroups(input *gamelift.ListGameServerGroupsInput) (*gamelift.ListGameServerGroupsOutput, error) {
    return a.client.ListGameServerGroups(input)
}

func (a *GameLiftActivities) ListGameServers(input *gamelift.ListGameServersInput) (*gamelift.ListGameServersOutput, error) {
    return a.client.ListGameServers(input)
}

func (a *GameLiftActivities) ListScripts(input *gamelift.ListScriptsInput) (*gamelift.ListScriptsOutput, error) {
    return a.client.ListScripts(input)
}

func (a *GameLiftActivities) ListTagsForResource(input *gamelift.ListTagsForResourceInput) (*gamelift.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *GameLiftActivities) PutScalingPolicy(input *gamelift.PutScalingPolicyInput) (*gamelift.PutScalingPolicyOutput, error) {
    return a.client.PutScalingPolicy(input)
}

func (a *GameLiftActivities) RegisterGameServer(input *gamelift.RegisterGameServerInput) (*gamelift.RegisterGameServerOutput, error) {
    return a.client.RegisterGameServer(input)
}

func (a *GameLiftActivities) RequestUploadCredentials(input *gamelift.RequestUploadCredentialsInput) (*gamelift.RequestUploadCredentialsOutput, error) {
    return a.client.RequestUploadCredentials(input)
}

func (a *GameLiftActivities) ResolveAlias(input *gamelift.ResolveAliasInput) (*gamelift.ResolveAliasOutput, error) {
    return a.client.ResolveAlias(input)
}

func (a *GameLiftActivities) ResumeGameServerGroup(input *gamelift.ResumeGameServerGroupInput) (*gamelift.ResumeGameServerGroupOutput, error) {
    return a.client.ResumeGameServerGroup(input)
}

func (a *GameLiftActivities) SearchGameSessions(input *gamelift.SearchGameSessionsInput) (*gamelift.SearchGameSessionsOutput, error) {
    return a.client.SearchGameSessions(input)
}

func (a *GameLiftActivities) StartFleetActions(input *gamelift.StartFleetActionsInput) (*gamelift.StartFleetActionsOutput, error) {
    return a.client.StartFleetActions(input)
}

func (a *GameLiftActivities) StartGameSessionPlacement(input *gamelift.StartGameSessionPlacementInput) (*gamelift.StartGameSessionPlacementOutput, error) {
    return a.client.StartGameSessionPlacement(input)
}

func (a *GameLiftActivities) StartMatchBackfill(input *gamelift.StartMatchBackfillInput) (*gamelift.StartMatchBackfillOutput, error) {
    return a.client.StartMatchBackfill(input)
}

func (a *GameLiftActivities) StartMatchmaking(input *gamelift.StartMatchmakingInput) (*gamelift.StartMatchmakingOutput, error) {
    return a.client.StartMatchmaking(input)
}

func (a *GameLiftActivities) StopFleetActions(input *gamelift.StopFleetActionsInput) (*gamelift.StopFleetActionsOutput, error) {
    return a.client.StopFleetActions(input)
}

func (a *GameLiftActivities) StopGameSessionPlacement(input *gamelift.StopGameSessionPlacementInput) (*gamelift.StopGameSessionPlacementOutput, error) {
    return a.client.StopGameSessionPlacement(input)
}

func (a *GameLiftActivities) StopMatchmaking(input *gamelift.StopMatchmakingInput) (*gamelift.StopMatchmakingOutput, error) {
    return a.client.StopMatchmaking(input)
}

func (a *GameLiftActivities) SuspendGameServerGroup(input *gamelift.SuspendGameServerGroupInput) (*gamelift.SuspendGameServerGroupOutput, error) {
    return a.client.SuspendGameServerGroup(input)
}

func (a *GameLiftActivities) TagResource(input *gamelift.TagResourceInput) (*gamelift.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *GameLiftActivities) UntagResource(input *gamelift.UntagResourceInput) (*gamelift.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *GameLiftActivities) UpdateAlias(input *gamelift.UpdateAliasInput) (*gamelift.UpdateAliasOutput, error) {
    return a.client.UpdateAlias(input)
}

func (a *GameLiftActivities) UpdateBuild(input *gamelift.UpdateBuildInput) (*gamelift.UpdateBuildOutput, error) {
    return a.client.UpdateBuild(input)
}

func (a *GameLiftActivities) UpdateFleetAttributes(input *gamelift.UpdateFleetAttributesInput) (*gamelift.UpdateFleetAttributesOutput, error) {
    return a.client.UpdateFleetAttributes(input)
}

func (a *GameLiftActivities) UpdateFleetCapacity(input *gamelift.UpdateFleetCapacityInput) (*gamelift.UpdateFleetCapacityOutput, error) {
    return a.client.UpdateFleetCapacity(input)
}

func (a *GameLiftActivities) UpdateFleetPortSettings(input *gamelift.UpdateFleetPortSettingsInput) (*gamelift.UpdateFleetPortSettingsOutput, error) {
    return a.client.UpdateFleetPortSettings(input)
}

func (a *GameLiftActivities) UpdateGameServer(input *gamelift.UpdateGameServerInput) (*gamelift.UpdateGameServerOutput, error) {
    return a.client.UpdateGameServer(input)
}

func (a *GameLiftActivities) UpdateGameServerGroup(input *gamelift.UpdateGameServerGroupInput) (*gamelift.UpdateGameServerGroupOutput, error) {
    return a.client.UpdateGameServerGroup(input)
}

func (a *GameLiftActivities) UpdateGameSession(input *gamelift.UpdateGameSessionInput) (*gamelift.UpdateGameSessionOutput, error) {
    return a.client.UpdateGameSession(input)
}

func (a *GameLiftActivities) UpdateGameSessionQueue(input *gamelift.UpdateGameSessionQueueInput) (*gamelift.UpdateGameSessionQueueOutput, error) {
    return a.client.UpdateGameSessionQueue(input)
}

func (a *GameLiftActivities) UpdateMatchmakingConfiguration(input *gamelift.UpdateMatchmakingConfigurationInput) (*gamelift.UpdateMatchmakingConfigurationOutput, error) {
    return a.client.UpdateMatchmakingConfiguration(input)
}

func (a *GameLiftActivities) UpdateRuntimeConfiguration(input *gamelift.UpdateRuntimeConfigurationInput) (*gamelift.UpdateRuntimeConfigurationOutput, error) {
    return a.client.UpdateRuntimeConfiguration(input)
}

func (a *GameLiftActivities) UpdateScript(input *gamelift.UpdateScriptInput) (*gamelift.UpdateScriptOutput, error) {
    return a.client.UpdateScript(input)
}

func (a *GameLiftActivities) ValidateMatchmakingRuleSet(input *gamelift.ValidateMatchmakingRuleSetInput) (*gamelift.ValidateMatchmakingRuleSetOutput, error) {
    return a.client.ValidateMatchmakingRuleSet(input)
}
