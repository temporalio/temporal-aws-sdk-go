
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type APIGatewayActivities struct {
    client apigatewayiface.APIGatewayAPI
}

func NewAPIGatewayActivities(session *session.Session, config ...*aws.Config) *APIGatewayActivities {
    client := apigateway.New(session, config...)
    return &APIGatewayActivities{client: client}
}

func (a *APIGatewayActivities) CreateApiKey(ctx context.Context, input *apigateway.CreateApiKeyInput) (*apigateway.ApiKey, error) {
    return a.client.CreateApiKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateAuthorizer(ctx context.Context, input *apigateway.CreateAuthorizerInput) (*apigateway.Authorizer, error) {
    return a.client.CreateAuthorizerWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateBasePathMapping(ctx context.Context, input *apigateway.CreateBasePathMappingInput) (*apigateway.BasePathMapping, error) {
    return a.client.CreateBasePathMappingWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateDeployment(ctx context.Context, input *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error) {
    return a.client.CreateDeploymentWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateDocumentationPart(ctx context.Context, input *apigateway.CreateDocumentationPartInput) (*apigateway.DocumentationPart, error) {
    return a.client.CreateDocumentationPartWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateDocumentationVersion(ctx context.Context, input *apigateway.CreateDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
    return a.client.CreateDocumentationVersionWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateDomainName(ctx context.Context, input *apigateway.CreateDomainNameInput) (*apigateway.DomainName, error) {
    return a.client.CreateDomainNameWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateModel(ctx context.Context, input *apigateway.CreateModelInput) (*apigateway.Model, error) {
    return a.client.CreateModelWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateRequestValidator(ctx context.Context, input *apigateway.CreateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
    return a.client.CreateRequestValidatorWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateResource(ctx context.Context, input *apigateway.CreateResourceInput) (*apigateway.Resource, error) {
    return a.client.CreateResourceWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateRestApi(ctx context.Context, input *apigateway.CreateRestApiInput) (*apigateway.RestApi, error) {
    return a.client.CreateRestApiWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateStage(ctx context.Context, input *apigateway.CreateStageInput) (*apigateway.Stage, error) {
    return a.client.CreateStageWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateUsagePlan(ctx context.Context, input *apigateway.CreateUsagePlanInput) (*apigateway.UsagePlan, error) {
    return a.client.CreateUsagePlanWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateUsagePlanKey(ctx context.Context, input *apigateway.CreateUsagePlanKeyInput) (*apigateway.UsagePlanKey, error) {
    return a.client.CreateUsagePlanKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) CreateVpcLink(ctx context.Context, input *apigateway.CreateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
    return a.client.CreateVpcLinkWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteApiKey(ctx context.Context, input *apigateway.DeleteApiKeyInput) (*apigateway.DeleteApiKeyOutput, error) {
    return a.client.DeleteApiKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteAuthorizer(ctx context.Context, input *apigateway.DeleteAuthorizerInput) (*apigateway.DeleteAuthorizerOutput, error) {
    return a.client.DeleteAuthorizerWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteBasePathMapping(ctx context.Context, input *apigateway.DeleteBasePathMappingInput) (*apigateway.DeleteBasePathMappingOutput, error) {
    return a.client.DeleteBasePathMappingWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteClientCertificate(ctx context.Context, input *apigateway.DeleteClientCertificateInput) (*apigateway.DeleteClientCertificateOutput, error) {
    return a.client.DeleteClientCertificateWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteDeployment(ctx context.Context, input *apigateway.DeleteDeploymentInput) (*apigateway.DeleteDeploymentOutput, error) {
    return a.client.DeleteDeploymentWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteDocumentationPart(ctx context.Context, input *apigateway.DeleteDocumentationPartInput) (*apigateway.DeleteDocumentationPartOutput, error) {
    return a.client.DeleteDocumentationPartWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteDocumentationVersion(ctx context.Context, input *apigateway.DeleteDocumentationVersionInput) (*apigateway.DeleteDocumentationVersionOutput, error) {
    return a.client.DeleteDocumentationVersionWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteDomainName(ctx context.Context, input *apigateway.DeleteDomainNameInput) (*apigateway.DeleteDomainNameOutput, error) {
    return a.client.DeleteDomainNameWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteGatewayResponse(ctx context.Context, input *apigateway.DeleteGatewayResponseInput) (*apigateway.DeleteGatewayResponseOutput, error) {
    return a.client.DeleteGatewayResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteIntegration(ctx context.Context, input *apigateway.DeleteIntegrationInput) (*apigateway.DeleteIntegrationOutput, error) {
    return a.client.DeleteIntegrationWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteIntegrationResponse(ctx context.Context, input *apigateway.DeleteIntegrationResponseInput) (*apigateway.DeleteIntegrationResponseOutput, error) {
    return a.client.DeleteIntegrationResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteMethod(ctx context.Context, input *apigateway.DeleteMethodInput) (*apigateway.DeleteMethodOutput, error) {
    return a.client.DeleteMethodWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteMethodResponse(ctx context.Context, input *apigateway.DeleteMethodResponseInput) (*apigateway.DeleteMethodResponseOutput, error) {
    return a.client.DeleteMethodResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteModel(ctx context.Context, input *apigateway.DeleteModelInput) (*apigateway.DeleteModelOutput, error) {
    return a.client.DeleteModelWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteRequestValidator(ctx context.Context, input *apigateway.DeleteRequestValidatorInput) (*apigateway.DeleteRequestValidatorOutput, error) {
    return a.client.DeleteRequestValidatorWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteResource(ctx context.Context, input *apigateway.DeleteResourceInput) (*apigateway.DeleteResourceOutput, error) {
    return a.client.DeleteResourceWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteRestApi(ctx context.Context, input *apigateway.DeleteRestApiInput) (*apigateway.DeleteRestApiOutput, error) {
    return a.client.DeleteRestApiWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteStage(ctx context.Context, input *apigateway.DeleteStageInput) (*apigateway.DeleteStageOutput, error) {
    return a.client.DeleteStageWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteUsagePlan(ctx context.Context, input *apigateway.DeleteUsagePlanInput) (*apigateway.DeleteUsagePlanOutput, error) {
    return a.client.DeleteUsagePlanWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteUsagePlanKey(ctx context.Context, input *apigateway.DeleteUsagePlanKeyInput) (*apigateway.DeleteUsagePlanKeyOutput, error) {
    return a.client.DeleteUsagePlanKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) DeleteVpcLink(ctx context.Context, input *apigateway.DeleteVpcLinkInput) (*apigateway.DeleteVpcLinkOutput, error) {
    return a.client.DeleteVpcLinkWithContext(ctx, input)
}

func (a *APIGatewayActivities) FlushStageAuthorizersCache(ctx context.Context, input *apigateway.FlushStageAuthorizersCacheInput) (*apigateway.FlushStageAuthorizersCacheOutput, error) {
    return a.client.FlushStageAuthorizersCacheWithContext(ctx, input)
}

func (a *APIGatewayActivities) FlushStageCache(ctx context.Context, input *apigateway.FlushStageCacheInput) (*apigateway.FlushStageCacheOutput, error) {
    return a.client.FlushStageCacheWithContext(ctx, input)
}

func (a *APIGatewayActivities) GenerateClientCertificate(ctx context.Context, input *apigateway.GenerateClientCertificateInput) (*apigateway.ClientCertificate, error) {
    return a.client.GenerateClientCertificateWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetAccount(ctx context.Context, input *apigateway.GetAccountInput) (*apigateway.Account, error) {
    return a.client.GetAccountWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetApiKey(ctx context.Context, input *apigateway.GetApiKeyInput) (*apigateway.ApiKey, error) {
    return a.client.GetApiKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetApiKeys(ctx context.Context, input *apigateway.GetApiKeysInput) (*apigateway.GetApiKeysOutput, error) {
    return a.client.GetApiKeysWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetAuthorizer(ctx context.Context, input *apigateway.GetAuthorizerInput) (*apigateway.Authorizer, error) {
    return a.client.GetAuthorizerWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetAuthorizers(ctx context.Context, input *apigateway.GetAuthorizersInput) (*apigateway.GetAuthorizersOutput, error) {
    return a.client.GetAuthorizersWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetBasePathMapping(ctx context.Context, input *apigateway.GetBasePathMappingInput) (*apigateway.BasePathMapping, error) {
    return a.client.GetBasePathMappingWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetBasePathMappings(ctx context.Context, input *apigateway.GetBasePathMappingsInput) (*apigateway.GetBasePathMappingsOutput, error) {
    return a.client.GetBasePathMappingsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetClientCertificate(ctx context.Context, input *apigateway.GetClientCertificateInput) (*apigateway.ClientCertificate, error) {
    return a.client.GetClientCertificateWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetClientCertificates(ctx context.Context, input *apigateway.GetClientCertificatesInput) (*apigateway.GetClientCertificatesOutput, error) {
    return a.client.GetClientCertificatesWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDeployment(ctx context.Context, input *apigateway.GetDeploymentInput) (*apigateway.Deployment, error) {
    return a.client.GetDeploymentWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDeployments(ctx context.Context, input *apigateway.GetDeploymentsInput) (*apigateway.GetDeploymentsOutput, error) {
    return a.client.GetDeploymentsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDocumentationPart(ctx context.Context, input *apigateway.GetDocumentationPartInput) (*apigateway.DocumentationPart, error) {
    return a.client.GetDocumentationPartWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDocumentationParts(ctx context.Context, input *apigateway.GetDocumentationPartsInput) (*apigateway.GetDocumentationPartsOutput, error) {
    return a.client.GetDocumentationPartsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDocumentationVersion(ctx context.Context, input *apigateway.GetDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
    return a.client.GetDocumentationVersionWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDocumentationVersions(ctx context.Context, input *apigateway.GetDocumentationVersionsInput) (*apigateway.GetDocumentationVersionsOutput, error) {
    return a.client.GetDocumentationVersionsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDomainName(ctx context.Context, input *apigateway.GetDomainNameInput) (*apigateway.DomainName, error) {
    return a.client.GetDomainNameWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetDomainNames(ctx context.Context, input *apigateway.GetDomainNamesInput) (*apigateway.GetDomainNamesOutput, error) {
    return a.client.GetDomainNamesWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetExport(ctx context.Context, input *apigateway.GetExportInput) (*apigateway.GetExportOutput, error) {
    return a.client.GetExportWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetGatewayResponse(ctx context.Context, input *apigateway.GetGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
    return a.client.GetGatewayResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetGatewayResponses(ctx context.Context, input *apigateway.GetGatewayResponsesInput) (*apigateway.GetGatewayResponsesOutput, error) {
    return a.client.GetGatewayResponsesWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetIntegration(ctx context.Context, input *apigateway.GetIntegrationInput) (*apigateway.Integration, error) {
    return a.client.GetIntegrationWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetIntegrationResponse(ctx context.Context, input *apigateway.GetIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
    return a.client.GetIntegrationResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetMethod(ctx context.Context, input *apigateway.GetMethodInput) (*apigateway.Method, error) {
    return a.client.GetMethodWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetMethodResponse(ctx context.Context, input *apigateway.GetMethodResponseInput) (*apigateway.MethodResponse, error) {
    return a.client.GetMethodResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetModel(ctx context.Context, input *apigateway.GetModelInput) (*apigateway.Model, error) {
    return a.client.GetModelWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetModelTemplate(ctx context.Context, input *apigateway.GetModelTemplateInput) (*apigateway.GetModelTemplateOutput, error) {
    return a.client.GetModelTemplateWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetModels(ctx context.Context, input *apigateway.GetModelsInput) (*apigateway.GetModelsOutput, error) {
    return a.client.GetModelsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetRequestValidator(ctx context.Context, input *apigateway.GetRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
    return a.client.GetRequestValidatorWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetRequestValidators(ctx context.Context, input *apigateway.GetRequestValidatorsInput) (*apigateway.GetRequestValidatorsOutput, error) {
    return a.client.GetRequestValidatorsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetResource(ctx context.Context, input *apigateway.GetResourceInput) (*apigateway.Resource, error) {
    return a.client.GetResourceWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetResources(ctx context.Context, input *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error) {
    return a.client.GetResourcesWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetRestApi(ctx context.Context, input *apigateway.GetRestApiInput) (*apigateway.RestApi, error) {
    return a.client.GetRestApiWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetRestApis(ctx context.Context, input *apigateway.GetRestApisInput) (*apigateway.GetRestApisOutput, error) {
    return a.client.GetRestApisWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetSdk(ctx context.Context, input *apigateway.GetSdkInput) (*apigateway.GetSdkOutput, error) {
    return a.client.GetSdkWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetSdkType(ctx context.Context, input *apigateway.GetSdkTypeInput) (*apigateway.SdkType, error) {
    return a.client.GetSdkTypeWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetSdkTypes(ctx context.Context, input *apigateway.GetSdkTypesInput) (*apigateway.GetSdkTypesOutput, error) {
    return a.client.GetSdkTypesWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetStage(ctx context.Context, input *apigateway.GetStageInput) (*apigateway.Stage, error) {
    return a.client.GetStageWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetStages(ctx context.Context, input *apigateway.GetStagesInput) (*apigateway.GetStagesOutput, error) {
    return a.client.GetStagesWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetTags(ctx context.Context, input *apigateway.GetTagsInput) (*apigateway.GetTagsOutput, error) {
    return a.client.GetTagsWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetUsage(ctx context.Context, input *apigateway.GetUsageInput) (*apigateway.Usage, error) {
    return a.client.GetUsageWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetUsagePlan(ctx context.Context, input *apigateway.GetUsagePlanInput) (*apigateway.UsagePlan, error) {
    return a.client.GetUsagePlanWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetUsagePlanKey(ctx context.Context, input *apigateway.GetUsagePlanKeyInput) (*apigateway.UsagePlanKey, error) {
    return a.client.GetUsagePlanKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetUsagePlanKeys(ctx context.Context, input *apigateway.GetUsagePlanKeysInput) (*apigateway.GetUsagePlanKeysOutput, error) {
    return a.client.GetUsagePlanKeysWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetUsagePlans(ctx context.Context, input *apigateway.GetUsagePlansInput) (*apigateway.GetUsagePlansOutput, error) {
    return a.client.GetUsagePlansWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetVpcLink(ctx context.Context, input *apigateway.GetVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
    return a.client.GetVpcLinkWithContext(ctx, input)
}

func (a *APIGatewayActivities) GetVpcLinks(ctx context.Context, input *apigateway.GetVpcLinksInput) (*apigateway.GetVpcLinksOutput, error) {
    return a.client.GetVpcLinksWithContext(ctx, input)
}

func (a *APIGatewayActivities) ImportApiKeys(ctx context.Context, input *apigateway.ImportApiKeysInput) (*apigateway.ImportApiKeysOutput, error) {
    return a.client.ImportApiKeysWithContext(ctx, input)
}

func (a *APIGatewayActivities) ImportDocumentationParts(ctx context.Context, input *apigateway.ImportDocumentationPartsInput) (*apigateway.ImportDocumentationPartsOutput, error) {
    return a.client.ImportDocumentationPartsWithContext(ctx, input)
}

func (a *APIGatewayActivities) ImportRestApi(ctx context.Context, input *apigateway.ImportRestApiInput) (*apigateway.RestApi, error) {
    return a.client.ImportRestApiWithContext(ctx, input)
}

func (a *APIGatewayActivities) PutGatewayResponse(ctx context.Context, input *apigateway.PutGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
    return a.client.PutGatewayResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) PutIntegration(ctx context.Context, input *apigateway.PutIntegrationInput) (*apigateway.Integration, error) {
    return a.client.PutIntegrationWithContext(ctx, input)
}

func (a *APIGatewayActivities) PutIntegrationResponse(ctx context.Context, input *apigateway.PutIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
    return a.client.PutIntegrationResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) PutMethod(ctx context.Context, input *apigateway.PutMethodInput) (*apigateway.Method, error) {
    return a.client.PutMethodWithContext(ctx, input)
}

func (a *APIGatewayActivities) PutMethodResponse(ctx context.Context, input *apigateway.PutMethodResponseInput) (*apigateway.MethodResponse, error) {
    return a.client.PutMethodResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) PutRestApi(ctx context.Context, input *apigateway.PutRestApiInput) (*apigateway.RestApi, error) {
    return a.client.PutRestApiWithContext(ctx, input)
}

func (a *APIGatewayActivities) TagResource(ctx context.Context, input *apigateway.TagResourceInput) (*apigateway.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *APIGatewayActivities) TestInvokeAuthorizer(ctx context.Context, input *apigateway.TestInvokeAuthorizerInput) (*apigateway.TestInvokeAuthorizerOutput, error) {
    return a.client.TestInvokeAuthorizerWithContext(ctx, input)
}

func (a *APIGatewayActivities) TestInvokeMethod(ctx context.Context, input *apigateway.TestInvokeMethodInput) (*apigateway.TestInvokeMethodOutput, error) {
    return a.client.TestInvokeMethodWithContext(ctx, input)
}

func (a *APIGatewayActivities) UntagResource(ctx context.Context, input *apigateway.UntagResourceInput) (*apigateway.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateAccount(ctx context.Context, input *apigateway.UpdateAccountInput) (*apigateway.Account, error) {
    return a.client.UpdateAccountWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateApiKey(ctx context.Context, input *apigateway.UpdateApiKeyInput) (*apigateway.ApiKey, error) {
    return a.client.UpdateApiKeyWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateAuthorizer(ctx context.Context, input *apigateway.UpdateAuthorizerInput) (*apigateway.Authorizer, error) {
    return a.client.UpdateAuthorizerWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateBasePathMapping(ctx context.Context, input *apigateway.UpdateBasePathMappingInput) (*apigateway.BasePathMapping, error) {
    return a.client.UpdateBasePathMappingWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateClientCertificate(ctx context.Context, input *apigateway.UpdateClientCertificateInput) (*apigateway.ClientCertificate, error) {
    return a.client.UpdateClientCertificateWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateDeployment(ctx context.Context, input *apigateway.UpdateDeploymentInput) (*apigateway.Deployment, error) {
    return a.client.UpdateDeploymentWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateDocumentationPart(ctx context.Context, input *apigateway.UpdateDocumentationPartInput) (*apigateway.DocumentationPart, error) {
    return a.client.UpdateDocumentationPartWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateDocumentationVersion(ctx context.Context, input *apigateway.UpdateDocumentationVersionInput) (*apigateway.DocumentationVersion, error) {
    return a.client.UpdateDocumentationVersionWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateDomainName(ctx context.Context, input *apigateway.UpdateDomainNameInput) (*apigateway.DomainName, error) {
    return a.client.UpdateDomainNameWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateGatewayResponse(ctx context.Context, input *apigateway.UpdateGatewayResponseInput) (*apigateway.UpdateGatewayResponseOutput, error) {
    return a.client.UpdateGatewayResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateIntegration(ctx context.Context, input *apigateway.UpdateIntegrationInput) (*apigateway.Integration, error) {
    return a.client.UpdateIntegrationWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateIntegrationResponse(ctx context.Context, input *apigateway.UpdateIntegrationResponseInput) (*apigateway.IntegrationResponse, error) {
    return a.client.UpdateIntegrationResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateMethod(ctx context.Context, input *apigateway.UpdateMethodInput) (*apigateway.Method, error) {
    return a.client.UpdateMethodWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateMethodResponse(ctx context.Context, input *apigateway.UpdateMethodResponseInput) (*apigateway.MethodResponse, error) {
    return a.client.UpdateMethodResponseWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateModel(ctx context.Context, input *apigateway.UpdateModelInput) (*apigateway.Model, error) {
    return a.client.UpdateModelWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateRequestValidator(ctx context.Context, input *apigateway.UpdateRequestValidatorInput) (*apigateway.UpdateRequestValidatorOutput, error) {
    return a.client.UpdateRequestValidatorWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateResource(ctx context.Context, input *apigateway.UpdateResourceInput) (*apigateway.Resource, error) {
    return a.client.UpdateResourceWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateRestApi(ctx context.Context, input *apigateway.UpdateRestApiInput) (*apigateway.RestApi, error) {
    return a.client.UpdateRestApiWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateStage(ctx context.Context, input *apigateway.UpdateStageInput) (*apigateway.Stage, error) {
    return a.client.UpdateStageWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateUsage(ctx context.Context, input *apigateway.UpdateUsageInput) (*apigateway.Usage, error) {
    return a.client.UpdateUsageWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateUsagePlan(ctx context.Context, input *apigateway.UpdateUsagePlanInput) (*apigateway.UsagePlan, error) {
    return a.client.UpdateUsagePlanWithContext(ctx, input)
}

func (a *APIGatewayActivities) UpdateVpcLink(ctx context.Context, input *apigateway.UpdateVpcLinkInput) (*apigateway.UpdateVpcLinkOutput, error) {
    return a.client.UpdateVpcLinkWithContext(ctx, input)
}
