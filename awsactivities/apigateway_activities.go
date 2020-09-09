package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
)

type APIGatewayActivities struct {
	client apigatewayiface.APIGatewayAPI
}

func NewAPIGatewayActivities(session *session.Session, config ...*aws.Config) *APIGatewayActivities {
	client := apigateway.New(session, config...)
	return &APIGatewayActivities{client: client}
}

func (a *APIGatewayActivities) CreateApiKey(input *apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error) {
	return a.client.CreateApiKey(input)
}

func (a *APIGatewayActivities) CreateAuthorizer(input *apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error) {
	return a.client.CreateAuthorizer(input)
}

func (a *APIGatewayActivities) CreateBasePathMapping(input *apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error) {
	return a.client.CreateBasePathMapping(input)
}

func (a *APIGatewayActivities) CreateDeployment(input *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error) {
	return a.client.CreateDeployment(input)
}

func (a *APIGatewayActivities) CreateDocumentationPart(input *apigateway.CreateDocumentationPartInput) (*apigateway.DocumentationPart, error) {
	return a.client.CreateDocumentationPart(input)
}

func (a *APIGatewayActivities) CreateDocumentationVersion(input *apigateway.CreateDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
	return a.client.CreateDocumentationVersion(input)
}

func (a *APIGatewayActivities) CreateDomainName(input *apigateway.CreateDomainNameInput) (*apigateway.DomainName, error) {
	return a.client.CreateDomainName(input)
}

func (a *APIGatewayActivities) CreateModel(input *apigateway.CreateModelInput) (*apigateway.Model, error) {
	return a.client.CreateModel(input)
}

func (a *APIGatewayActivities) CreateRequestValidator(input *apigateway.CreateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	return a.client.CreateRequestValidator(input)
}

func (a *APIGatewayActivities) CreateResource(input *apigateway.CreateResourceInput) (*apigateway.Resource, error) {
	return a.client.CreateResource(input)
}

func (a *APIGatewayActivities) CreateRestApi(input *apigateway.CreateRestApiInput) (*apigateway.RestApi, error) {
	return a.client.CreateRestApi(input)
}

func (a *APIGatewayActivities) CreateStage(input *apigateway.CreateStageInput) (*apigateway.Stage, error) {
	return a.client.CreateStage(input)
}

func (a *APIGatewayActivities) CreateUsagePlan(input *apigateway.CreateUsagePlanInput) (*apigateway.UsagePlan, error) {
	return a.client.CreateUsagePlan(input)
}

func (a *APIGatewayActivities) CreateUsagePlanKey(input *apigateway.CreateUsagePlanKeyInput) (*apigateway.UsagePlanKey, error) {
	return a.client.CreateUsagePlanKey(input)
}

func (a *APIGatewayActivities) CreateVpcLink(input *apigateway.CreateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	return a.client.CreateVpcLink(input)
}

func (a *APIGatewayActivities) DeleteApiKey(input *apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error) {
	return a.client.DeleteApiKey(input)
}

func (a *APIGatewayActivities) DeleteAuthorizer(input *apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error) {
	return a.client.DeleteAuthorizer(input)
}

func (a *APIGatewayActivities) DeleteBasePathMapping(input *apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error) {
	return a.client.DeleteBasePathMapping(input)
}

func (a *APIGatewayActivities) DeleteClientCertificate(input *apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error) {
	return a.client.DeleteClientCertificate(input)
}

func (a *APIGatewayActivities) DeleteDeployment(input *apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error) {
	return a.client.DeleteDeployment(input)
}

func (a *APIGatewayActivities) DeleteDocumentationPart(input *apigateway.DeleteDocumentationPartInput) (*apigateway.DeleteDocumentationPartOutput, error) {
	return a.client.DeleteDocumentationPart(input)
}

func (a *APIGatewayActivities) DeleteDocumentationVersion(input *apigateway.DeleteDocumentationVersionInput) (*apigateway.DeleteDocumentationVersionOutput, error) {
	return a.client.DeleteDocumentationVersion(input)
}

func (a *APIGatewayActivities) DeleteDomainName(input *apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error) {
	return a.client.DeleteDomainName(input)
}

func (a *APIGatewayActivities) DeleteGatewayResponse(input *apigateway.DeleteGatewayResponseInput) (*apigateway.DeleteGatewayResponseOutput, error) {
	return a.client.DeleteGatewayResponse(input)
}

func (a *APIGatewayActivities) DeleteIntegration(input *apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error) {
	return a.client.DeleteIntegration(input)
}

func (a *APIGatewayActivities) DeleteIntegrationResponse(input *apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error) {
	return a.client.DeleteIntegrationResponse(input)
}

func (a *APIGatewayActivities) DeleteMethod(input *apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error) {
	return a.client.DeleteMethod(input)
}

func (a *APIGatewayActivities) DeleteMethodResponse(input *apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error) {
	return a.client.DeleteMethodResponse(input)
}

func (a *APIGatewayActivities) DeleteModel(input *apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error) {
	return a.client.DeleteModel(input)
}

func (a *APIGatewayActivities) DeleteRequestValidator(input *apigateway.DeleteRequestValidatorInput) (*apigateway.DeleteRequestValidatorOutput, error) {
	return a.client.DeleteRequestValidator(input)
}

func (a *APIGatewayActivities) DeleteResource(input *apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error) {
	return a.client.DeleteResource(input)
}

func (a *APIGatewayActivities) DeleteRestApi(input *apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error) {
	return a.client.DeleteRestApi(input)
}

func (a *APIGatewayActivities) DeleteStage(input *apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error) {
	return a.client.DeleteStage(input)
}

func (a *APIGatewayActivities) DeleteUsagePlan(input *apigateway.DeleteUsagePlanInput) (*apigateway.DeleteUsagePlanOutput, error) {
	return a.client.DeleteUsagePlan(input)
}

func (a *APIGatewayActivities) DeleteUsagePlanKey(input *apigateway.DeleteUsagePlanKeyInput) (*apigateway.DeleteUsagePlanKeyOutput, error) {
	return a.client.DeleteUsagePlanKey(input)
}

func (a *APIGatewayActivities) DeleteVpcLink(input *apigateway.DeleteVpcLinkInput) (*apigateway.DeleteVpcLinkOutput, error) {
	return a.client.DeleteVpcLink(input)
}

func (a *APIGatewayActivities) FlushStageAuthorizersCache(input *apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error) {
	return a.client.FlushStageAuthorizersCache(input)
}

func (a *APIGatewayActivities) FlushStageCache(input *apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error) {
	return a.client.FlushStageCache(input)
}

func (a *APIGatewayActivities) GenerateClientCertificate(input *apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error) {
	return a.client.GenerateClientCertificate(input)
}

func (a *APIGatewayActivities) GetAccount(input *apigateway.GetAccountInput) (*apigateway.Account, error) {
	return a.client.GetAccount(input)
}

func (a *APIGatewayActivities) GetApiKey(input *apigateway.GetApiKeyInput) (*apigateway.ApiKey, error) {
	return a.client.GetApiKey(input)
}

func (a *APIGatewayActivities) GetApiKeys(input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error) {
	return a.client.GetApiKeys(input)
}

func (a *APIGatewayActivities) GetAuthorizer(input *apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error) {
	return a.client.GetAuthorizer(input)
}

func (a *APIGatewayActivities) GetAuthorizers(input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error) {
	return a.client.GetAuthorizers(input)
}

func (a *APIGatewayActivities) GetBasePathMapping(input *apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error) {
	return a.client.GetBasePathMapping(input)
}

func (a *APIGatewayActivities) GetBasePathMappings(input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error) {
	return a.client.GetBasePathMappings(input)
}

func (a *APIGatewayActivities) GetClientCertificate(input *apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error) {
	return a.client.GetClientCertificate(input)
}

func (a *APIGatewayActivities) GetClientCertificates(input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error) {
	return a.client.GetClientCertificates(input)
}

func (a *APIGatewayActivities) GetDeployment(input *apigateway.GetDeploymentInput) (*apigateway.Deployment, error) {
	return a.client.GetDeployment(input)
}

func (a *APIGatewayActivities) GetDeployments(input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error) {
	return a.client.GetDeployments(input)
}

func (a *APIGatewayActivities) GetDocumentationPart(input *apigateway.GetDocumentationPartInput) (*apigateway.DocumentationPart, error) {
	return a.client.GetDocumentationPart(input)
}

func (a *APIGatewayActivities) GetDocumentationParts(input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error) {
	return a.client.GetDocumentationParts(input)
}

func (a *APIGatewayActivities) GetDocumentationVersion(input *apigateway.GetDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
	return a.client.GetDocumentationVersion(input)
}

func (a *APIGatewayActivities) GetDocumentationVersions(input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error) {
	return a.client.GetDocumentationVersions(input)
}

func (a *APIGatewayActivities) GetDomainName(input *apigateway.GetDomainNameInput) (*apigateway.DomainName, error) {
	return a.client.GetDomainName(input)
}

func (a *APIGatewayActivities) GetDomainNames(input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error) {
	return a.client.GetDomainNames(input)
}

func (a *APIGatewayActivities) GetExport(input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error) {
	return a.client.GetExport(input)
}

func (a *APIGatewayActivities) GetGatewayResponse(input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	return a.client.GetGatewayResponse(input)
}

func (a *APIGatewayActivities) GetGatewayResponses(input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error) {
	return a.client.GetGatewayResponses(input)
}

func (a *APIGatewayActivities) GetIntegration(input *apigateway.GetIntegrationInput) (*apigateway.Integration, error) {
	return a.client.GetIntegration(input)
}

func (a *APIGatewayActivities) GetIntegrationResponse(input *apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
	return a.client.GetIntegrationResponse(input)
}

func (a *APIGatewayActivities) GetMethod(input *apigateway.GetMethodInput) (*apigateway.Method, error) {
	return a.client.GetMethod(input)
}

func (a *APIGatewayActivities) GetMethodResponse(input *apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error) {
	return a.client.GetMethodResponse(input)
}

func (a *APIGatewayActivities) GetModel(input *apigateway.GetModelInput) (*apigateway.Model, error) {
	return a.client.GetModel(input)
}

func (a *APIGatewayActivities) GetModelTemplate(input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error) {
	return a.client.GetModelTemplate(input)
}

func (a *APIGatewayActivities) GetModels(input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error) {
	return a.client.GetModels(input)
}

func (a *APIGatewayActivities) GetRequestValidator(input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	return a.client.GetRequestValidator(input)
}

func (a *APIGatewayActivities) GetRequestValidators(input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error) {
	return a.client.GetRequestValidators(input)
}

func (a *APIGatewayActivities) GetResource(input *apigateway.GetResourceInput) (*apigateway.Resource, error) {
	return a.client.GetResource(input)
}

func (a *APIGatewayActivities) GetResources(input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error) {
	return a.client.GetResources(input)
}

func (a *APIGatewayActivities) GetRestApi(input *apigateway.GetRestApiInput) (*apigateway.RestApi, error) {
	return a.client.GetRestApi(input)
}

func (a *APIGatewayActivities) GetRestApis(input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error) {
	return a.client.GetRestApis(input)
}

func (a *APIGatewayActivities) GetSdk(input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error) {
	return a.client.GetSdk(input)
}

func (a *APIGatewayActivities) GetSdkType(input *apigateway.GetSdkTypeInput) (*apigateway.SdkType, error) {
	return a.client.GetSdkType(input)
}

func (a *APIGatewayActivities) GetSdkTypes(input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error) {
	return a.client.GetSdkTypes(input)
}

func (a *APIGatewayActivities) GetStage(input *apigateway.GetStageInput) (*apigateway.Stage, error) {
	return a.client.GetStage(input)
}

func (a *APIGatewayActivities) GetStages(input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error) {
	return a.client.GetStages(input)
}

func (a *APIGatewayActivities) GetTags(input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error) {
	return a.client.GetTags(input)
}

func (a *APIGatewayActivities) GetUsage(input *apigateway.GetUsageInput) (*apigateway.Usage, error) {
	return a.client.GetUsage(input)
}

func (a *APIGatewayActivities) GetUsagePlan(input *apigateway.GetUsagePlanInput) (*apigateway.UsagePlan, error) {
	return a.client.GetUsagePlan(input)
}

func (a *APIGatewayActivities) GetUsagePlanKey(input *apigateway.GetUsagePlanKeyInput) (*apigateway.UsagePlanKey, error) {
	return a.client.GetUsagePlanKey(input)
}

func (a *APIGatewayActivities) GetUsagePlanKeys(input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error) {
	return a.client.GetUsagePlanKeys(input)
}

func (a *APIGatewayActivities) GetUsagePlans(input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error) {
	return a.client.GetUsagePlans(input)
}

func (a *APIGatewayActivities) GetVpcLink(input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	return a.client.GetVpcLink(input)
}

func (a *APIGatewayActivities) GetVpcLinks(input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error) {
	return a.client.GetVpcLinks(input)
}

func (a *APIGatewayActivities) ImportApiKeys(input *apigateway.ImportApiKeysInput) (*apigateway.ImportApiKeysOutput, error) {
	return a.client.ImportApiKeys(input)
}

func (a *APIGatewayActivities) ImportDocumentationParts(input *apigateway.ImportDocumentationPartsInput) (*apigateway.ImportDocumentationPartsOutput, error) {
	return a.client.ImportDocumentationParts(input)
}

func (a *APIGatewayActivities) ImportRestApi(input *apigateway.ImportRestApiInput) (*apigateway.RestApi, error) {
	return a.client.ImportRestApi(input)
}

func (a *APIGatewayActivities) PutGatewayResponse(input *apigateway.PutGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	return a.client.PutGatewayResponse(input)
}

func (a *APIGatewayActivities) PutIntegration(input *apigateway.PutIntegrationInput) (*apigateway.Integration, error) {
	return a.client.PutIntegration(input)
}

func (a *APIGatewayActivities) PutIntegrationResponse(input *apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
	return a.client.PutIntegrationResponse(input)
}

func (a *APIGatewayActivities) PutMethod(input *apigateway.PutMethodInput) (*apigateway.Method, error) {
	return a.client.PutMethod(input)
}

func (a *APIGatewayActivities) PutMethodResponse(input *apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error) {
	return a.client.PutMethodResponse(input)
}

func (a *APIGatewayActivities) PutRestApi(input *apigateway.PutRestApiInput) (*apigateway.RestApi, error) {
	return a.client.PutRestApi(input)
}

func (a *APIGatewayActivities) TagResource(input *apigateway.TagResourceInput) (*apigateway.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *APIGatewayActivities) TestInvokeAuthorizer(input *apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error) {
	return a.client.TestInvokeAuthorizer(input)
}

func (a *APIGatewayActivities) TestInvokeMethod(input *apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error) {
	return a.client.TestInvokeMethod(input)
}

func (a *APIGatewayActivities) UntagResource(input *apigateway.UntagResourceInput) (*apigateway.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *APIGatewayActivities) UpdateAccount(input *apigateway.UpdateAccountInput) (*apigateway.Account, error) {
	return a.client.UpdateAccount(input)
}

func (a *APIGatewayActivities) UpdateApiKey(input *apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error) {
	return a.client.UpdateApiKey(input)
}

func (a *APIGatewayActivities) UpdateAuthorizer(input *apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error) {
	return a.client.UpdateAuthorizer(input)
}

func (a *APIGatewayActivities) UpdateBasePathMapping(input *apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error) {
	return a.client.UpdateBasePathMapping(input)
}

func (a *APIGatewayActivities) UpdateClientCertificate(input *apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error) {
	return a.client.UpdateClientCertificate(input)
}

func (a *APIGatewayActivities) UpdateDeployment(input *apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error) {
	return a.client.UpdateDeployment(input)
}

func (a *APIGatewayActivities) UpdateDocumentationPart(input *apigateway.UpdateDocumentationPartInput) (*apigateway.DocumentationPart, error) {
	return a.client.UpdateDocumentationPart(input)
}

func (a *APIGatewayActivities) UpdateDocumentationVersion(input *apigateway.UpdateDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
	return a.client.UpdateDocumentationVersion(input)
}

func (a *APIGatewayActivities) UpdateDomainName(input *apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error) {
	return a.client.UpdateDomainName(input)
}

func (a *APIGatewayActivities) UpdateGatewayResponse(input *apigateway.UpdateGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
	return a.client.UpdateGatewayResponse(input)
}

func (a *APIGatewayActivities) UpdateIntegration(input *apigateway.UpdateIntegrationInput) (*apigateway.Integration, error) {
	return a.client.UpdateIntegration(input)
}

func (a *APIGatewayActivities) UpdateIntegrationResponse(input *apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
	return a.client.UpdateIntegrationResponse(input)
}

func (a *APIGatewayActivities) UpdateMethod(input *apigateway.UpdateMethodInput) (*apigateway.Method, error) {
	return a.client.UpdateMethod(input)
}

func (a *APIGatewayActivities) UpdateMethodResponse(input *apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error) {
	return a.client.UpdateMethodResponse(input)
}

func (a *APIGatewayActivities) UpdateModel(input *apigateway.UpdateModelInput) (*apigateway.Model, error) {
	return a.client.UpdateModel(input)
}

func (a *APIGatewayActivities) UpdateRequestValidator(input *apigateway.UpdateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
	return a.client.UpdateRequestValidator(input)
}

func (a *APIGatewayActivities) UpdateResource(input *apigateway.UpdateResourceInput) (*apigateway.Resource, error) {
	return a.client.UpdateResource(input)
}

func (a *APIGatewayActivities) UpdateRestApi(input *apigateway.UpdateRestApiInput) (*apigateway.RestApi, error) {
	return a.client.UpdateRestApi(input)
}

func (a *APIGatewayActivities) UpdateStage(input *apigateway.UpdateStageInput) (*apigateway.Stage, error) {
	return a.client.UpdateStage(input)
}

func (a *APIGatewayActivities) UpdateUsage(input *apigateway.UpdateUsageInput) (*apigateway.Usage, error) {
	return a.client.UpdateUsage(input)
}

func (a *APIGatewayActivities) UpdateUsagePlan(input *apigateway.UpdateUsagePlanInput) (*apigateway.UsagePlan, error) {
	return a.client.UpdateUsagePlan(input)
}

func (a *APIGatewayActivities) UpdateVpcLink(input *apigateway.UpdateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
	return a.client.UpdateVpcLink(input)
}
