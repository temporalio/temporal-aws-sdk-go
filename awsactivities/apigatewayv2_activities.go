
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
)

type ApiGatewayV2Activities struct {
	client apigatewayv2iface.ApiGatewayV2API
}

func NewApiGatewayV2Activities(session *session.Session, config... *aws.Config) *ApiGatewayV2Activities {
    client := apigatewayv2.New(session, config...)
    return &ApiGatewayV2Activities{client: client}
}

func (a *ApiGatewayV2Activities) CreateApi(input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error) {
    return a.client.CreateApi(input)
}

func (a *ApiGatewayV2Activities) CreateApiMapping(input *apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error) {
    return a.client.CreateApiMapping(input)
}

func (a *ApiGatewayV2Activities) CreateAuthorizer(input *apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error) {
    return a.client.CreateAuthorizer(input)
}

func (a *ApiGatewayV2Activities) CreateDeployment(input *apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error) {
    return a.client.CreateDeployment(input)
}

func (a *ApiGatewayV2Activities) CreateDomainName(input *apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error) {
    return a.client.CreateDomainName(input)
}

func (a *ApiGatewayV2Activities) CreateIntegration(input *apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error) {
    return a.client.CreateIntegration(input)
}

func (a *ApiGatewayV2Activities) CreateIntegrationResponse(input *apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error) {
    return a.client.CreateIntegrationResponse(input)
}

func (a *ApiGatewayV2Activities) CreateModel(input *apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error) {
    return a.client.CreateModel(input)
}

func (a *ApiGatewayV2Activities) CreateRoute(input *apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error) {
    return a.client.CreateRoute(input)
}

func (a *ApiGatewayV2Activities) CreateRouteResponse(input *apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error) {
    return a.client.CreateRouteResponse(input)
}

func (a *ApiGatewayV2Activities) CreateStage(input *apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error) {
    return a.client.CreateStage(input)
}

func (a *ApiGatewayV2Activities) CreateVpcLink(input *apigatewayv2.CreateVpcLinkInput) (*apigatewayv2.CreateVpcLinkOutput, error) {
    return a.client.CreateVpcLink(input)
}

func (a *ApiGatewayV2Activities) DeleteAccessLogSettings(input *apigatewayv2.DeleteAccessLogSettingsInput) (*apigatewayv2.DeleteAccessLogSettingsOutput, error) {
    return a.client.DeleteAccessLogSettings(input)
}

func (a *ApiGatewayV2Activities) DeleteApi(input *apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error) {
    return a.client.DeleteApi(input)
}

func (a *ApiGatewayV2Activities) DeleteApiMapping(input *apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error) {
    return a.client.DeleteApiMapping(input)
}

func (a *ApiGatewayV2Activities) DeleteAuthorizer(input *apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error) {
    return a.client.DeleteAuthorizer(input)
}

func (a *ApiGatewayV2Activities) DeleteCorsConfiguration(input *apigatewayv2.DeleteCorsConfigurationInput) (*apigatewayv2.DeleteCorsConfigurationOutput, error) {
    return a.client.DeleteCorsConfiguration(input)
}

func (a *ApiGatewayV2Activities) DeleteDeployment(input *apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error) {
    return a.client.DeleteDeployment(input)
}

func (a *ApiGatewayV2Activities) DeleteDomainName(input *apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error) {
    return a.client.DeleteDomainName(input)
}

func (a *ApiGatewayV2Activities) DeleteIntegration(input *apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error) {
    return a.client.DeleteIntegration(input)
}

func (a *ApiGatewayV2Activities) DeleteIntegrationResponse(input *apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error) {
    return a.client.DeleteIntegrationResponse(input)
}

func (a *ApiGatewayV2Activities) DeleteModel(input *apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error) {
    return a.client.DeleteModel(input)
}

func (a *ApiGatewayV2Activities) DeleteRoute(input *apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error) {
    return a.client.DeleteRoute(input)
}

func (a *ApiGatewayV2Activities) DeleteRouteRequestParameter(input *apigatewayv2.DeleteRouteRequestParameterInput) (*apigatewayv2.DeleteRouteRequestParameterOutput, error) {
    return a.client.DeleteRouteRequestParameter(input)
}

func (a *ApiGatewayV2Activities) DeleteRouteResponse(input *apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error) {
    return a.client.DeleteRouteResponse(input)
}

func (a *ApiGatewayV2Activities) DeleteRouteSettings(input *apigatewayv2.DeleteRouteSettingsInput) (*apigatewayv2.DeleteRouteSettingsOutput, error) {
    return a.client.DeleteRouteSettings(input)
}

func (a *ApiGatewayV2Activities) DeleteStage(input *apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error) {
    return a.client.DeleteStage(input)
}

func (a *ApiGatewayV2Activities) DeleteVpcLink(input *apigatewayv2.DeleteVpcLinkInput) (*apigatewayv2.DeleteVpcLinkOutput, error) {
    return a.client.DeleteVpcLink(input)
}

func (a *ApiGatewayV2Activities) ExportApi(input *apigatewayv2.ExportApiInput) (*apigatewayv2.ExportApiOutput, error) {
    return a.client.ExportApi(input)
}

func (a *ApiGatewayV2Activities) GetApi(input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error) {
    return a.client.GetApi(input)
}

func (a *ApiGatewayV2Activities) GetApiMapping(input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error) {
    return a.client.GetApiMapping(input)
}

func (a *ApiGatewayV2Activities) GetApiMappings(input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error) {
    return a.client.GetApiMappings(input)
}

func (a *ApiGatewayV2Activities) GetApis(input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error) {
    return a.client.GetApis(input)
}

func (a *ApiGatewayV2Activities) GetAuthorizer(input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error) {
    return a.client.GetAuthorizer(input)
}

func (a *ApiGatewayV2Activities) GetAuthorizers(input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error) {
    return a.client.GetAuthorizers(input)
}

func (a *ApiGatewayV2Activities) GetDeployment(input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error) {
    return a.client.GetDeployment(input)
}

func (a *ApiGatewayV2Activities) GetDeployments(input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error) {
    return a.client.GetDeployments(input)
}

func (a *ApiGatewayV2Activities) GetDomainName(input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error) {
    return a.client.GetDomainName(input)
}

func (a *ApiGatewayV2Activities) GetDomainNames(input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error) {
    return a.client.GetDomainNames(input)
}

func (a *ApiGatewayV2Activities) GetIntegration(input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error) {
    return a.client.GetIntegration(input)
}

func (a *ApiGatewayV2Activities) GetIntegrationResponse(input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error) {
    return a.client.GetIntegrationResponse(input)
}

func (a *ApiGatewayV2Activities) GetIntegrationResponses(input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
    return a.client.GetIntegrationResponses(input)
}

func (a *ApiGatewayV2Activities) GetIntegrations(input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error) {
    return a.client.GetIntegrations(input)
}

func (a *ApiGatewayV2Activities) GetModel(input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error) {
    return a.client.GetModel(input)
}

func (a *ApiGatewayV2Activities) GetModelTemplate(input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error) {
    return a.client.GetModelTemplate(input)
}

func (a *ApiGatewayV2Activities) GetModels(input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error) {
    return a.client.GetModels(input)
}

func (a *ApiGatewayV2Activities) GetRoute(input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error) {
    return a.client.GetRoute(input)
}

func (a *ApiGatewayV2Activities) GetRouteResponse(input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error) {
    return a.client.GetRouteResponse(input)
}

func (a *ApiGatewayV2Activities) GetRouteResponses(input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error) {
    return a.client.GetRouteResponses(input)
}

func (a *ApiGatewayV2Activities) GetRoutes(input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error) {
    return a.client.GetRoutes(input)
}

func (a *ApiGatewayV2Activities) GetStage(input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error) {
    return a.client.GetStage(input)
}

func (a *ApiGatewayV2Activities) GetStages(input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error) {
    return a.client.GetStages(input)
}

func (a *ApiGatewayV2Activities) GetTags(input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error) {
    return a.client.GetTags(input)
}

func (a *ApiGatewayV2Activities) GetVpcLink(input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error) {
    return a.client.GetVpcLink(input)
}

func (a *ApiGatewayV2Activities) GetVpcLinks(input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error) {
    return a.client.GetVpcLinks(input)
}

func (a *ApiGatewayV2Activities) ImportApi(input *apigatewayv2.ImportApiInput) (*apigatewayv2.ImportApiOutput, error) {
    return a.client.ImportApi(input)
}

func (a *ApiGatewayV2Activities) ReimportApi(input *apigatewayv2.ReimportApiInput) (*apigatewayv2.ReimportApiOutput, error) {
    return a.client.ReimportApi(input)
}

func (a *ApiGatewayV2Activities) TagResource(input *apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ApiGatewayV2Activities) UntagResource(input *apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ApiGatewayV2Activities) UpdateApi(input *apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error) {
    return a.client.UpdateApi(input)
}

func (a *ApiGatewayV2Activities) UpdateApiMapping(input *apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error) {
    return a.client.UpdateApiMapping(input)
}

func (a *ApiGatewayV2Activities) UpdateAuthorizer(input *apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error) {
    return a.client.UpdateAuthorizer(input)
}

func (a *ApiGatewayV2Activities) UpdateDeployment(input *apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error) {
    return a.client.UpdateDeployment(input)
}

func (a *ApiGatewayV2Activities) UpdateDomainName(input *apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error) {
    return a.client.UpdateDomainName(input)
}

func (a *ApiGatewayV2Activities) UpdateIntegration(input *apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error) {
    return a.client.UpdateIntegration(input)
}

func (a *ApiGatewayV2Activities) UpdateIntegrationResponse(input *apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error) {
    return a.client.UpdateIntegrationResponse(input)
}

func (a *ApiGatewayV2Activities) UpdateModel(input *apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error) {
    return a.client.UpdateModel(input)
}

func (a *ApiGatewayV2Activities) UpdateRoute(input *apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error) {
    return a.client.UpdateRoute(input)
}

func (a *ApiGatewayV2Activities) UpdateRouteResponse(input *apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error) {
    return a.client.UpdateRouteResponse(input)
}

func (a *ApiGatewayV2Activities) UpdateStage(input *apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error) {
    return a.client.UpdateStage(input)
}

func (a *ApiGatewayV2Activities) UpdateVpcLink(input *apigatewayv2.UpdateVpcLinkInput) (*apigatewayv2.UpdateVpcLinkOutput, error) {
    return a.client.UpdateVpcLink(input)
}
