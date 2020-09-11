
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/gamelift/gameliftiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type GameLiftActivities struct {
    client gameliftiface.GameLiftAPI
}

func NewGameLiftActivities(session *session.Session, config ...*aws.Config) *GameLiftActivities {
    client := gamelift.New(session, config...)
    return &GameLiftActivities{client: client}
}

func (a *GameLiftActivities) AcceptMatch(ctx context.Context, input *gamelift.AcceptMatchInput) (*gamelift.AcceptMatchOutput, error) {
    return a.client.AcceptMatchWithContext(ctx, input)
}

func (a *GameLiftActivities) ClaimGameServer(ctx context.Context, input *gamelift.ClaimGameServerInput) (*gamelift.ClaimGameServerOutput, error) {
    return a.client.ClaimGameServerWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateAlias(ctx context.Context, input *gamelift.CreateAliasInput) (*gamelift.CreateAliasOutput, error) {
    return a.client.CreateAliasWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateBuild(ctx context.Context, input *gamelift.CreateBuildInput) (*gamelift.CreateBuildOutput, error) {
    return a.client.CreateBuildWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateFleet(ctx context.Context, input *gamelift.CreateFleetInput) (*gamelift.CreateFleetOutput, error) {
    return a.client.CreateFleetWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateGameServerGroup(ctx context.Context, input *gamelift.CreateGameServerGroupInput) (*gamelift.CreateGameServerGroupOutput, error) {
    return a.client.CreateGameServerGroupWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateGameSession(ctx context.Context, input *gamelift.CreateGameSessionInput) (*gamelift.CreateGameSessionOutput, error) {
    return a.client.CreateGameSessionWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateGameSessionQueue(ctx context.Context, input *gamelift.CreateGameSessionQueueInput) (*gamelift.CreateGameSessionQueueOutput, error) {
    return a.client.CreateGameSessionQueueWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateMatchmakingConfiguration(ctx context.Context, input *gamelift.CreateMatchmakingConfigurationInput) (*gamelift.CreateMatchmakingConfigurationOutput, error) {
    return a.client.CreateMatchmakingConfigurationWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateMatchmakingRuleSet(ctx context.Context, input *gamelift.CreateMatchmakingRuleSetInput) (*gamelift.CreateMatchmakingRuleSetOutput, error) {
    return a.client.CreateMatchmakingRuleSetWithContext(ctx, input)
}

func (a *GameLiftActivities) CreatePlayerSession(ctx context.Context, input *gamelift.CreatePlayerSessionInput) (*gamelift.CreatePlayerSessionOutput, error) {
    return a.client.CreatePlayerSessionWithContext(ctx, input)
}

func (a *GameLiftActivities) CreatePlayerSessions(ctx context.Context, input *gamelift.CreatePlayerSessionsInput) (*gamelift.CreatePlayerSessionsOutput, error) {
    return a.client.CreatePlayerSessionsWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateScript(ctx context.Context, input *gamelift.CreateScriptInput) (*gamelift.CreateScriptOutput, error) {
    return a.client.CreateScriptWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateVpcPeeringAuthorization(ctx context.Context, input *gamelift.CreateVpcPeeringAuthorizationInput) (*gamelift.CreateVpcPeeringAuthorizationOutput, error) {
    return a.client.CreateVpcPeeringAuthorizationWithContext(ctx, input)
}

func (a *GameLiftActivities) CreateVpcPeeringConnection(ctx context.Context, input *gamelift.CreateVpcPeeringConnectionInput) (*gamelift.CreateVpcPeeringConnectionOutput, error) {
    return a.client.CreateVpcPeeringConnectionWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteAlias(ctx context.Context, input *gamelift.DeleteAliasInput) (*gamelift.DeleteAliasOutput, error) {
    return a.client.DeleteAliasWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteBuild(ctx context.Context, input *gamelift.DeleteBuildInput) (*gamelift.DeleteBuildOutput, error) {
    return a.client.DeleteBuildWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteFleet(ctx context.Context, input *gamelift.DeleteFleetInput) (*gamelift.DeleteFleetOutput, error) {
    return a.client.DeleteFleetWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteGameServerGroup(ctx context.Context, input *gamelift.DeleteGameServerGroupInput) (*gamelift.DeleteGameServerGroupOutput, error) {
    return a.client.DeleteGameServerGroupWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteGameSessionQueue(ctx context.Context, input *gamelift.DeleteGameSessionQueueInput) (*gamelift.DeleteGameSessionQueueOutput, error) {
    return a.client.DeleteGameSessionQueueWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteMatchmakingConfiguration(ctx context.Context, input *gamelift.DeleteMatchmakingConfigurationInput) (*gamelift.DeleteMatchmakingConfigurationOutput, error) {
    return a.client.DeleteMatchmakingConfigurationWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteMatchmakingRuleSet(ctx context.Context, input *gamelift.DeleteMatchmakingRuleSetInput) (*gamelift.DeleteMatchmakingRuleSetOutput, error) {
    return a.client.DeleteMatchmakingRuleSetWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteScalingPolicy(ctx context.Context, input *gamelift.DeleteScalingPolicyInput) (*gamelift.DeleteScalingPolicyOutput, error) {
    return a.client.DeleteScalingPolicyWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteScript(ctx context.Context, input *gamelift.DeleteScriptInput) (*gamelift.DeleteScriptOutput, error) {
    return a.client.DeleteScriptWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteVpcPeeringAuthorization(ctx context.Context, input *gamelift.DeleteVpcPeeringAuthorizationInput) (*gamelift.DeleteVpcPeeringAuthorizationOutput, error) {
    return a.client.DeleteVpcPeeringAuthorizationWithContext(ctx, input)
}

func (a *GameLiftActivities) DeleteVpcPeeringConnection(ctx context.Context, input *gamelift.DeleteVpcPeeringConnectionInput) (*gamelift.DeleteVpcPeeringConnectionOutput, error) {
    return a.client.DeleteVpcPeeringConnectionWithContext(ctx, input)
}

func (a *GameLiftActivities) DeregisterGameServer(ctx context.Context, input *gamelift.DeregisterGameServerInput) (*gamelift.DeregisterGameServerOutput, error) {
    return a.client.DeregisterGameServerWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeAlias(ctx context.Context, input *gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error) {
    return a.client.DescribeAliasWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeBuild(ctx context.Context, input *gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error) {
    return a.client.DescribeBuildWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeEC2InstanceLimits(ctx context.Context, input *gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
    return a.client.DescribeEC2InstanceLimitsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeFleetAttributes(ctx context.Context, input *gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error) {
    return a.client.DescribeFleetAttributesWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeFleetCapacity(ctx context.Context, input *gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error) {
    return a.client.DescribeFleetCapacityWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeFleetEvents(ctx context.Context, input *gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error) {
    return a.client.DescribeFleetEventsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeFleetPortSettings(ctx context.Context, input *gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error) {
    return a.client.DescribeFleetPortSettingsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeFleetUtilization(ctx context.Context, input *gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error) {
    return a.client.DescribeFleetUtilizationWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameServer(ctx context.Context, input *gamelift.DescribeGameServerInput) (*gamelift.DescribeGameServerOutput, error) {
    return a.client.DescribeGameServerWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameServerGroup(ctx context.Context, input *gamelift.DescribeGameServerGroupInput) (*gamelift.DescribeGameServerGroupOutput, error) {
    return a.client.DescribeGameServerGroupWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameServerInstances(ctx context.Context, input *gamelift.DescribeGameServerInstancesInput) (*gamelift.DescribeGameServerInstancesOutput, error) {
    return a.client.DescribeGameServerInstancesWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameSessionDetails(ctx context.Context, input *gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error) {
    return a.client.DescribeGameSessionDetailsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameSessionPlacement(ctx context.Context, input *gamelift.DescribeGameSessionPlacementInput) (*gamelift.DescribeGameSessionPlacementOutput, error) {
    return a.client.DescribeGameSessionPlacementWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameSessionQueues(ctx context.Context, input *gamelift.DescribeGameSessionQueuesInput) (*gamelift.DescribeGameSessionQueuesOutput, error) {
    return a.client.DescribeGameSessionQueuesWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeGameSessions(ctx context.Context, input *gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error) {
    return a.client.DescribeGameSessionsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeInstances(ctx context.Context, input *gamelift.DescribeInstancesInput) (*gamelift.DescribeInstancesOutput, error) {
    return a.client.DescribeInstancesWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeMatchmaking(ctx context.Context, input *gamelift.DescribeMatchmakingInput) (*gamelift.DescribeMatchmakingOutput, error) {
    return a.client.DescribeMatchmakingWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeMatchmakingConfigurations(ctx context.Context, input *gamelift.DescribeMatchmakingConfigurationsInput) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
    return a.client.DescribeMatchmakingConfigurationsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeMatchmakingRuleSets(ctx context.Context, input *gamelift.DescribeMatchmakingRuleSetsInput) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
    return a.client.DescribeMatchmakingRuleSetsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribePlayerSessions(ctx context.Context, input *gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error) {
    return a.client.DescribePlayerSessionsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeRuntimeConfiguration(ctx context.Context, input *gamelift.DescribeRuntimeConfigurationInput) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
    return a.client.DescribeRuntimeConfigurationWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeScalingPolicies(ctx context.Context, input *gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error) {
    return a.client.DescribeScalingPoliciesWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeScript(ctx context.Context, input *gamelift.DescribeScriptInput) (*gamelift.DescribeScriptOutput, error) {
    return a.client.DescribeScriptWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeVpcPeeringAuthorizations(ctx context.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
    return a.client.DescribeVpcPeeringAuthorizationsWithContext(ctx, input)
}

func (a *GameLiftActivities) DescribeVpcPeeringConnections(ctx context.Context, input *gamelift.DescribeVpcPeeringConnectionsInput) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
    return a.client.DescribeVpcPeeringConnectionsWithContext(ctx, input)
}

func (a *GameLiftActivities) GetGameSessionLogUrl(ctx context.Context, input *gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error) {
    return a.client.GetGameSessionLogUrlWithContext(ctx, input)
}

func (a *GameLiftActivities) GetInstanceAccess(ctx context.Context, input *gamelift.GetInstanceAccessInput) (*gamelift.GetInstanceAccessOutput, error) {
    return a.client.GetInstanceAccessWithContext(ctx, input)
}

func (a *GameLiftActivities) ListAliases(ctx context.Context, input *gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error) {
    return a.client.ListAliasesWithContext(ctx, input)
}

func (a *GameLiftActivities) ListBuilds(ctx context.Context, input *gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error) {
    return a.client.ListBuildsWithContext(ctx, input)
}

func (a *GameLiftActivities) ListFleets(ctx context.Context, input *gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error) {
    return a.client.ListFleetsWithContext(ctx, input)
}

func (a *GameLiftActivities) ListGameServerGroups(ctx context.Context, input *gamelift.ListGameServerGroupsInput) (*gamelift.ListGameServerGroupsOutput, error) {
    return a.client.ListGameServerGroupsWithContext(ctx, input)
}

func (a *GameLiftActivities) ListGameServers(ctx context.Context, input *gamelift.ListGameServersInput) (*gamelift.ListGameServersOutput, error) {
    return a.client.ListGameServersWithContext(ctx, input)
}

func (a *GameLiftActivities) ListScripts(ctx context.Context, input *gamelift.ListScriptsInput) (*gamelift.ListScriptsOutput, error) {
    return a.client.ListScriptsWithContext(ctx, input)
}

func (a *GameLiftActivities) ListTagsForResource(ctx context.Context, input *gamelift.ListTagsForResourceInput) (*gamelift.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *GameLiftActivities) PutScalingPolicy(ctx context.Context, input *gamelift.PutScalingPolicyInput) (*gamelift.PutScalingPolicyOutput, error) {
    return a.client.PutScalingPolicyWithContext(ctx, input)
}

func (a *GameLiftActivities) RegisterGameServer(ctx context.Context, input *gamelift.RegisterGameServerInput) (*gamelift.RegisterGameServerOutput, error) {
    return a.client.RegisterGameServerWithContext(ctx, input)
}

func (a *GameLiftActivities) RequestUploadCredentials(ctx context.Context, input *gamelift.RequestUploadCredentialsInput) (*gamelift.RequestUploadCredentialsOutput, error) {
    return a.client.RequestUploadCredentialsWithContext(ctx, input)
}

func (a *GameLiftActivities) ResolveAlias(ctx context.Context, input *gamelift.ResolveAliasInput) (*gamelift.ResolveAliasOutput, error) {
    return a.client.ResolveAliasWithContext(ctx, input)
}

func (a *GameLiftActivities) ResumeGameServerGroup(ctx context.Context, input *gamelift.ResumeGameServerGroupInput) (*gamelift.ResumeGameServerGroupOutput, error) {
    return a.client.ResumeGameServerGroupWithContext(ctx, input)
}

func (a *GameLiftActivities) SearchGameSessions(ctx context.Context, input *gamelift.SearchGameSessionsInput) (*gamelift.SearchGameSessionsOutput, error) {
    return a.client.SearchGameSessionsWithContext(ctx, input)
}

func (a *GameLiftActivities) StartFleetActions(ctx context.Context, input *gamelift.StartFleetActionsInput) (*gamelift.StartFleetActionsOutput, error) {
    return a.client.StartFleetActionsWithContext(ctx, input)
}

func (a *GameLiftActivities) StartGameSessionPlacement(ctx context.Context, input *gamelift.StartGameSessionPlacementInput) (*gamelift.StartGameSessionPlacementOutput, error) {
    return a.client.StartGameSessionPlacementWithContext(ctx, input)
}

func (a *GameLiftActivities) StartMatchBackfill(ctx context.Context, input *gamelift.StartMatchBackfillInput) (*gamelift.StartMatchBackfillOutput, error) {
    return a.client.StartMatchBackfillWithContext(ctx, input)
}

func (a *GameLiftActivities) StartMatchmaking(ctx context.Context, input *gamelift.StartMatchmakingInput) (*gamelift.StartMatchmakingOutput, error) {
    return a.client.StartMatchmakingWithContext(ctx, input)
}

func (a *GameLiftActivities) StopFleetActions(ctx context.Context, input *gamelift.StopFleetActionsInput) (*gamelift.StopFleetActionsOutput, error) {
    return a.client.StopFleetActionsWithContext(ctx, input)
}

func (a *GameLiftActivities) StopGameSessionPlacement(ctx context.Context, input *gamelift.StopGameSessionPlacementInput) (*gamelift.StopGameSessionPlacementOutput, error) {
    return a.client.StopGameSessionPlacementWithContext(ctx, input)
}

func (a *GameLiftActivities) StopMatchmaking(ctx context.Context, input *gamelift.StopMatchmakingInput) (*gamelift.StopMatchmakingOutput, error) {
    return a.client.StopMatchmakingWithContext(ctx, input)
}

func (a *GameLiftActivities) SuspendGameServerGroup(ctx context.Context, input *gamelift.SuspendGameServerGroupInput) (*gamelift.SuspendGameServerGroupOutput, error) {
    return a.client.SuspendGameServerGroupWithContext(ctx, input)
}

func (a *GameLiftActivities) TagResource(ctx context.Context, input *gamelift.TagResourceInput) (*gamelift.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *GameLiftActivities) UntagResource(ctx context.Context, input *gamelift.UntagResourceInput) (*gamelift.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateAlias(ctx context.Context, input *gamelift.UpdateAliasInput) (*gamelift.UpdateAliasOutput, error) {
    return a.client.UpdateAliasWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateBuild(ctx context.Context, input *gamelift.UpdateBuildInput) (*gamelift.UpdateBuildOutput, error) {
    return a.client.UpdateBuildWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateFleetAttributes(ctx context.Context, input *gamelift.UpdateFleetAttributesInput) (*gamelift.UpdateFleetAttributesOutput, error) {
    return a.client.UpdateFleetAttributesWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateFleetCapacity(ctx context.Context, input *gamelift.UpdateFleetCapacityInput) (*gamelift.UpdateFleetCapacityOutput, error) {
    return a.client.UpdateFleetCapacityWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateFleetPortSettings(ctx context.Context, input *gamelift.UpdateFleetPortSettingsInput) (*gamelift.UpdateFleetPortSettingsOutput, error) {
    return a.client.UpdateFleetPortSettingsWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateGameServer(ctx context.Context, input *gamelift.UpdateGameServerInput) (*gamelift.UpdateGameServerOutput, error) {
    return a.client.UpdateGameServerWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateGameServerGroup(ctx context.Context, input *gamelift.UpdateGameServerGroupInput) (*gamelift.UpdateGameServerGroupOutput, error) {
    return a.client.UpdateGameServerGroupWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateGameSession(ctx context.Context, input *gamelift.UpdateGameSessionInput) (*gamelift.UpdateGameSessionOutput, error) {
    return a.client.UpdateGameSessionWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateGameSessionQueue(ctx context.Context, input *gamelift.UpdateGameSessionQueueInput) (*gamelift.UpdateGameSessionQueueOutput, error) {
    return a.client.UpdateGameSessionQueueWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateMatchmakingConfiguration(ctx context.Context, input *gamelift.UpdateMatchmakingConfigurationInput) (*gamelift.UpdateMatchmakingConfigurationOutput, error) {
    return a.client.UpdateMatchmakingConfigurationWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateRuntimeConfiguration(ctx context.Context, input *gamelift.UpdateRuntimeConfigurationInput) (*gamelift.UpdateRuntimeConfigurationOutput, error) {
    return a.client.UpdateRuntimeConfigurationWithContext(ctx, input)
}

func (a *GameLiftActivities) UpdateScript(ctx context.Context, input *gamelift.UpdateScriptInput) (*gamelift.UpdateScriptOutput, error) {
    return a.client.UpdateScriptWithContext(ctx, input)
}

func (a *GameLiftActivities) ValidateMatchmakingRuleSet(ctx context.Context, input *gamelift.ValidateMatchmakingRuleSetInput) (*gamelift.ValidateMatchmakingRuleSetOutput, error) {
    return a.client.ValidateMatchmakingRuleSetWithContext(ctx, input)
}
