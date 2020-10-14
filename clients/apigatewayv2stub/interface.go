// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package apigatewayv2stub

import (
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateApi(ctx workflow.Context, input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error)
	CreateApiAsync(ctx workflow.Context, input *apigatewayv2.CreateApiInput) *ApiGatewayV2CreateApiFuture

	CreateApiMapping(ctx workflow.Context, input *apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error)
	CreateApiMappingAsync(ctx workflow.Context, input *apigatewayv2.CreateApiMappingInput) *ApiGatewayV2CreateApiMappingFuture

	CreateAuthorizer(ctx workflow.Context, input *apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error)
	CreateAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.CreateAuthorizerInput) *ApiGatewayV2CreateAuthorizerFuture

	CreateDeployment(ctx workflow.Context, input *apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error)
	CreateDeploymentAsync(ctx workflow.Context, input *apigatewayv2.CreateDeploymentInput) *ApiGatewayV2CreateDeploymentFuture

	CreateDomainName(ctx workflow.Context, input *apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error)
	CreateDomainNameAsync(ctx workflow.Context, input *apigatewayv2.CreateDomainNameInput) *ApiGatewayV2CreateDomainNameFuture

	CreateIntegration(ctx workflow.Context, input *apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error)
	CreateIntegrationAsync(ctx workflow.Context, input *apigatewayv2.CreateIntegrationInput) *ApiGatewayV2CreateIntegrationFuture

	CreateIntegrationResponse(ctx workflow.Context, input *apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error)
	CreateIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.CreateIntegrationResponseInput) *ApiGatewayV2CreateIntegrationResponseFuture

	CreateModel(ctx workflow.Context, input *apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error)
	CreateModelAsync(ctx workflow.Context, input *apigatewayv2.CreateModelInput) *ApiGatewayV2CreateModelFuture

	CreateRoute(ctx workflow.Context, input *apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error)
	CreateRouteAsync(ctx workflow.Context, input *apigatewayv2.CreateRouteInput) *ApiGatewayV2CreateRouteFuture

	CreateRouteResponse(ctx workflow.Context, input *apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error)
	CreateRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.CreateRouteResponseInput) *ApiGatewayV2CreateRouteResponseFuture

	CreateStage(ctx workflow.Context, input *apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error)
	CreateStageAsync(ctx workflow.Context, input *apigatewayv2.CreateStageInput) *ApiGatewayV2CreateStageFuture

	CreateVpcLink(ctx workflow.Context, input *apigatewayv2.CreateVpcLinkInput) (*apigatewayv2.CreateVpcLinkOutput, error)
	CreateVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.CreateVpcLinkInput) *ApiGatewayV2CreateVpcLinkFuture

	DeleteAccessLogSettings(ctx workflow.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) (*apigatewayv2.DeleteAccessLogSettingsOutput, error)
	DeleteAccessLogSettingsAsync(ctx workflow.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) *ApiGatewayV2DeleteAccessLogSettingsFuture

	DeleteApi(ctx workflow.Context, input *apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error)
	DeleteApiAsync(ctx workflow.Context, input *apigatewayv2.DeleteApiInput) *ApiGatewayV2DeleteApiFuture

	DeleteApiMapping(ctx workflow.Context, input *apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error)
	DeleteApiMappingAsync(ctx workflow.Context, input *apigatewayv2.DeleteApiMappingInput) *ApiGatewayV2DeleteApiMappingFuture

	DeleteAuthorizer(ctx workflow.Context, input *apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error)
	DeleteAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.DeleteAuthorizerInput) *ApiGatewayV2DeleteAuthorizerFuture

	DeleteCorsConfiguration(ctx workflow.Context, input *apigatewayv2.DeleteCorsConfigurationInput) (*apigatewayv2.DeleteCorsConfigurationOutput, error)
	DeleteCorsConfigurationAsync(ctx workflow.Context, input *apigatewayv2.DeleteCorsConfigurationInput) *ApiGatewayV2DeleteCorsConfigurationFuture

	DeleteDeployment(ctx workflow.Context, input *apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error)
	DeleteDeploymentAsync(ctx workflow.Context, input *apigatewayv2.DeleteDeploymentInput) *ApiGatewayV2DeleteDeploymentFuture

	DeleteDomainName(ctx workflow.Context, input *apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error)
	DeleteDomainNameAsync(ctx workflow.Context, input *apigatewayv2.DeleteDomainNameInput) *ApiGatewayV2DeleteDomainNameFuture

	DeleteIntegration(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error)
	DeleteIntegrationAsync(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationInput) *ApiGatewayV2DeleteIntegrationFuture

	DeleteIntegrationResponse(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.DeleteIntegrationResponseInput) *ApiGatewayV2DeleteIntegrationResponseFuture

	DeleteModel(ctx workflow.Context, input *apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error)
	DeleteModelAsync(ctx workflow.Context, input *apigatewayv2.DeleteModelInput) *ApiGatewayV2DeleteModelFuture

	DeleteRoute(ctx workflow.Context, input *apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error)
	DeleteRouteAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteInput) *ApiGatewayV2DeleteRouteFuture

	DeleteRouteRequestParameter(ctx workflow.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) (*apigatewayv2.DeleteRouteRequestParameterOutput, error)
	DeleteRouteRequestParameterAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) *ApiGatewayV2DeleteRouteRequestParameterFuture

	DeleteRouteResponse(ctx workflow.Context, input *apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error)
	DeleteRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteResponseInput) *ApiGatewayV2DeleteRouteResponseFuture

	DeleteRouteSettings(ctx workflow.Context, input *apigatewayv2.DeleteRouteSettingsInput) (*apigatewayv2.DeleteRouteSettingsOutput, error)
	DeleteRouteSettingsAsync(ctx workflow.Context, input *apigatewayv2.DeleteRouteSettingsInput) *ApiGatewayV2DeleteRouteSettingsFuture

	DeleteStage(ctx workflow.Context, input *apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error)
	DeleteStageAsync(ctx workflow.Context, input *apigatewayv2.DeleteStageInput) *ApiGatewayV2DeleteStageFuture

	DeleteVpcLink(ctx workflow.Context, input *apigatewayv2.DeleteVpcLinkInput) (*apigatewayv2.DeleteVpcLinkOutput, error)
	DeleteVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.DeleteVpcLinkInput) *ApiGatewayV2DeleteVpcLinkFuture

	ExportApi(ctx workflow.Context, input *apigatewayv2.ExportApiInput) (*apigatewayv2.ExportApiOutput, error)
	ExportApiAsync(ctx workflow.Context, input *apigatewayv2.ExportApiInput) *ApiGatewayV2ExportApiFuture

	GetApi(ctx workflow.Context, input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error)
	GetApiAsync(ctx workflow.Context, input *apigatewayv2.GetApiInput) *ApiGatewayV2GetApiFuture

	GetApiMapping(ctx workflow.Context, input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error)
	GetApiMappingAsync(ctx workflow.Context, input *apigatewayv2.GetApiMappingInput) *ApiGatewayV2GetApiMappingFuture

	GetApiMappings(ctx workflow.Context, input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApiMappingsAsync(ctx workflow.Context, input *apigatewayv2.GetApiMappingsInput) *ApiGatewayV2GetApiMappingsFuture

	GetApis(ctx workflow.Context, input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error)
	GetApisAsync(ctx workflow.Context, input *apigatewayv2.GetApisInput) *ApiGatewayV2GetApisFuture

	GetAuthorizer(ctx workflow.Context, input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error)
	GetAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.GetAuthorizerInput) *ApiGatewayV2GetAuthorizerFuture

	GetAuthorizers(ctx workflow.Context, input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error)
	GetAuthorizersAsync(ctx workflow.Context, input *apigatewayv2.GetAuthorizersInput) *ApiGatewayV2GetAuthorizersFuture

	GetDeployment(ctx workflow.Context, input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error)
	GetDeploymentAsync(ctx workflow.Context, input *apigatewayv2.GetDeploymentInput) *ApiGatewayV2GetDeploymentFuture

	GetDeployments(ctx workflow.Context, input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDeploymentsAsync(ctx workflow.Context, input *apigatewayv2.GetDeploymentsInput) *ApiGatewayV2GetDeploymentsFuture

	GetDomainName(ctx workflow.Context, input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error)
	GetDomainNameAsync(ctx workflow.Context, input *apigatewayv2.GetDomainNameInput) *ApiGatewayV2GetDomainNameFuture

	GetDomainNames(ctx workflow.Context, input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error)
	GetDomainNamesAsync(ctx workflow.Context, input *apigatewayv2.GetDomainNamesInput) *ApiGatewayV2GetDomainNamesFuture

	GetIntegration(ctx workflow.Context, input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error)
	GetIntegrationAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationInput) *ApiGatewayV2GetIntegrationFuture

	GetIntegrationResponse(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error)
	GetIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponseInput) *ApiGatewayV2GetIntegrationResponseFuture

	GetIntegrationResponses(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrationResponsesAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationResponsesInput) *ApiGatewayV2GetIntegrationResponsesFuture

	GetIntegrations(ctx workflow.Context, input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error)
	GetIntegrationsAsync(ctx workflow.Context, input *apigatewayv2.GetIntegrationsInput) *ApiGatewayV2GetIntegrationsFuture

	GetModel(ctx workflow.Context, input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error)
	GetModelAsync(ctx workflow.Context, input *apigatewayv2.GetModelInput) *ApiGatewayV2GetModelFuture

	GetModelTemplate(ctx workflow.Context, input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error)
	GetModelTemplateAsync(ctx workflow.Context, input *apigatewayv2.GetModelTemplateInput) *ApiGatewayV2GetModelTemplateFuture

	GetModels(ctx workflow.Context, input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error)
	GetModelsAsync(ctx workflow.Context, input *apigatewayv2.GetModelsInput) *ApiGatewayV2GetModelsFuture

	GetRoute(ctx workflow.Context, input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error)
	GetRouteAsync(ctx workflow.Context, input *apigatewayv2.GetRouteInput) *ApiGatewayV2GetRouteFuture

	GetRouteResponse(ctx workflow.Context, input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error)
	GetRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.GetRouteResponseInput) *ApiGatewayV2GetRouteResponseFuture

	GetRouteResponses(ctx workflow.Context, input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRouteResponsesAsync(ctx workflow.Context, input *apigatewayv2.GetRouteResponsesInput) *ApiGatewayV2GetRouteResponsesFuture

	GetRoutes(ctx workflow.Context, input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error)
	GetRoutesAsync(ctx workflow.Context, input *apigatewayv2.GetRoutesInput) *ApiGatewayV2GetRoutesFuture

	GetStage(ctx workflow.Context, input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error)
	GetStageAsync(ctx workflow.Context, input *apigatewayv2.GetStageInput) *ApiGatewayV2GetStageFuture

	GetStages(ctx workflow.Context, input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error)
	GetStagesAsync(ctx workflow.Context, input *apigatewayv2.GetStagesInput) *ApiGatewayV2GetStagesFuture

	GetTags(ctx workflow.Context, input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *apigatewayv2.GetTagsInput) *ApiGatewayV2GetTagsFuture

	GetVpcLink(ctx workflow.Context, input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error)
	GetVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.GetVpcLinkInput) *ApiGatewayV2GetVpcLinkFuture

	GetVpcLinks(ctx workflow.Context, input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error)
	GetVpcLinksAsync(ctx workflow.Context, input *apigatewayv2.GetVpcLinksInput) *ApiGatewayV2GetVpcLinksFuture

	ImportApi(ctx workflow.Context, input *apigatewayv2.ImportApiInput) (*apigatewayv2.ImportApiOutput, error)
	ImportApiAsync(ctx workflow.Context, input *apigatewayv2.ImportApiInput) *ApiGatewayV2ImportApiFuture

	ReimportApi(ctx workflow.Context, input *apigatewayv2.ReimportApiInput) (*apigatewayv2.ReimportApiOutput, error)
	ReimportApiAsync(ctx workflow.Context, input *apigatewayv2.ReimportApiInput) *ApiGatewayV2ReimportApiFuture

	ResetAuthorizersCache(ctx workflow.Context, input *apigatewayv2.ResetAuthorizersCacheInput) (*apigatewayv2.ResetAuthorizersCacheOutput, error)
	ResetAuthorizersCacheAsync(ctx workflow.Context, input *apigatewayv2.ResetAuthorizersCacheInput) *ApiGatewayV2ResetAuthorizersCacheFuture

	TagResource(ctx workflow.Context, input *apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *apigatewayv2.TagResourceInput) *ApiGatewayV2TagResourceFuture

	UntagResource(ctx workflow.Context, input *apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *apigatewayv2.UntagResourceInput) *ApiGatewayV2UntagResourceFuture

	UpdateApi(ctx workflow.Context, input *apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error)
	UpdateApiAsync(ctx workflow.Context, input *apigatewayv2.UpdateApiInput) *ApiGatewayV2UpdateApiFuture

	UpdateApiMapping(ctx workflow.Context, input *apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error)
	UpdateApiMappingAsync(ctx workflow.Context, input *apigatewayv2.UpdateApiMappingInput) *ApiGatewayV2UpdateApiMappingFuture

	UpdateAuthorizer(ctx workflow.Context, input *apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error)
	UpdateAuthorizerAsync(ctx workflow.Context, input *apigatewayv2.UpdateAuthorizerInput) *ApiGatewayV2UpdateAuthorizerFuture

	UpdateDeployment(ctx workflow.Context, input *apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error)
	UpdateDeploymentAsync(ctx workflow.Context, input *apigatewayv2.UpdateDeploymentInput) *ApiGatewayV2UpdateDeploymentFuture

	UpdateDomainName(ctx workflow.Context, input *apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error)
	UpdateDomainNameAsync(ctx workflow.Context, input *apigatewayv2.UpdateDomainNameInput) *ApiGatewayV2UpdateDomainNameFuture

	UpdateIntegration(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error)
	UpdateIntegrationAsync(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationInput) *ApiGatewayV2UpdateIntegrationFuture

	UpdateIntegrationResponse(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error)
	UpdateIntegrationResponseAsync(ctx workflow.Context, input *apigatewayv2.UpdateIntegrationResponseInput) *ApiGatewayV2UpdateIntegrationResponseFuture

	UpdateModel(ctx workflow.Context, input *apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error)
	UpdateModelAsync(ctx workflow.Context, input *apigatewayv2.UpdateModelInput) *ApiGatewayV2UpdateModelFuture

	UpdateRoute(ctx workflow.Context, input *apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error)
	UpdateRouteAsync(ctx workflow.Context, input *apigatewayv2.UpdateRouteInput) *ApiGatewayV2UpdateRouteFuture

	UpdateRouteResponse(ctx workflow.Context, input *apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error)
	UpdateRouteResponseAsync(ctx workflow.Context, input *apigatewayv2.UpdateRouteResponseInput) *ApiGatewayV2UpdateRouteResponseFuture

	UpdateStage(ctx workflow.Context, input *apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error)
	UpdateStageAsync(ctx workflow.Context, input *apigatewayv2.UpdateStageInput) *ApiGatewayV2UpdateStageFuture

	UpdateVpcLink(ctx workflow.Context, input *apigatewayv2.UpdateVpcLinkInput) (*apigatewayv2.UpdateVpcLinkOutput, error)
	UpdateVpcLinkAsync(ctx workflow.Context, input *apigatewayv2.UpdateVpcLinkInput) *ApiGatewayV2UpdateVpcLinkFuture
}

func NewClient() Client {
	return &stub{}
}