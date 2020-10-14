// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package apigatewayv2

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client apigatewayv2iface.ApiGatewayV2API

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := apigatewayv2.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (apigatewayv2iface.ApiGatewayV2API, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return apigatewayv2.New(sess), nil
}

func (a *Activities) CreateApi(ctx context.Context, input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApiWithContext(ctx, input)
}

func (a *Activities) CreateApiMapping(ctx context.Context, input *apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApiMappingWithContext(ctx, input)
}

func (a *Activities) CreateAuthorizer(ctx context.Context, input *apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAuthorizerWithContext(ctx, input)
}

func (a *Activities) CreateDeployment(ctx context.Context, input *apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeploymentWithContext(ctx, input)
}

func (a *Activities) CreateDomainName(ctx context.Context, input *apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDomainNameWithContext(ctx, input)
}

func (a *Activities) CreateIntegration(ctx context.Context, input *apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateIntegrationWithContext(ctx, input)
}

func (a *Activities) CreateIntegrationResponse(ctx context.Context, input *apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateIntegrationResponseWithContext(ctx, input)
}

func (a *Activities) CreateModel(ctx context.Context, input *apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateModelWithContext(ctx, input)
}

func (a *Activities) CreateRoute(ctx context.Context, input *apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRouteWithContext(ctx, input)
}

func (a *Activities) CreateRouteResponse(ctx context.Context, input *apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRouteResponseWithContext(ctx, input)
}

func (a *Activities) CreateStage(ctx context.Context, input *apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateStageWithContext(ctx, input)
}

func (a *Activities) CreateVpcLink(ctx context.Context, input *apigatewayv2.CreateVpcLinkInput) (*apigatewayv2.CreateVpcLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateVpcLinkWithContext(ctx, input)
}

func (a *Activities) DeleteAccessLogSettings(ctx context.Context, input *apigatewayv2.DeleteAccessLogSettingsInput) (*apigatewayv2.DeleteAccessLogSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAccessLogSettingsWithContext(ctx, input)
}

func (a *Activities) DeleteApi(ctx context.Context, input *apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApiWithContext(ctx, input)
}

func (a *Activities) DeleteApiMapping(ctx context.Context, input *apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApiMappingWithContext(ctx, input)
}

func (a *Activities) DeleteAuthorizer(ctx context.Context, input *apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAuthorizerWithContext(ctx, input)
}

func (a *Activities) DeleteCorsConfiguration(ctx context.Context, input *apigatewayv2.DeleteCorsConfigurationInput) (*apigatewayv2.DeleteCorsConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCorsConfigurationWithContext(ctx, input)
}

func (a *Activities) DeleteDeployment(ctx context.Context, input *apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDeploymentWithContext(ctx, input)
}

func (a *Activities) DeleteDomainName(ctx context.Context, input *apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDomainNameWithContext(ctx, input)
}

func (a *Activities) DeleteIntegration(ctx context.Context, input *apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIntegrationWithContext(ctx, input)
}

func (a *Activities) DeleteIntegrationResponse(ctx context.Context, input *apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIntegrationResponseWithContext(ctx, input)
}

func (a *Activities) DeleteModel(ctx context.Context, input *apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteModelWithContext(ctx, input)
}

func (a *Activities) DeleteRoute(ctx context.Context, input *apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRouteWithContext(ctx, input)
}

func (a *Activities) DeleteRouteRequestParameter(ctx context.Context, input *apigatewayv2.DeleteRouteRequestParameterInput) (*apigatewayv2.DeleteRouteRequestParameterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRouteRequestParameterWithContext(ctx, input)
}

func (a *Activities) DeleteRouteResponse(ctx context.Context, input *apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRouteResponseWithContext(ctx, input)
}

func (a *Activities) DeleteRouteSettings(ctx context.Context, input *apigatewayv2.DeleteRouteSettingsInput) (*apigatewayv2.DeleteRouteSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRouteSettingsWithContext(ctx, input)
}

func (a *Activities) DeleteStage(ctx context.Context, input *apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteStageWithContext(ctx, input)
}

func (a *Activities) DeleteVpcLink(ctx context.Context, input *apigatewayv2.DeleteVpcLinkInput) (*apigatewayv2.DeleteVpcLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVpcLinkWithContext(ctx, input)
}

func (a *Activities) ExportApi(ctx context.Context, input *apigatewayv2.ExportApiInput) (*apigatewayv2.ExportApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExportApiWithContext(ctx, input)
}

func (a *Activities) GetApi(ctx context.Context, input *apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApiWithContext(ctx, input)
}

func (a *Activities) GetApiMapping(ctx context.Context, input *apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApiMappingWithContext(ctx, input)
}

func (a *Activities) GetApiMappings(ctx context.Context, input *apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApiMappingsWithContext(ctx, input)
}

func (a *Activities) GetApis(ctx context.Context, input *apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApisWithContext(ctx, input)
}

func (a *Activities) GetAuthorizer(ctx context.Context, input *apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAuthorizerWithContext(ctx, input)
}

func (a *Activities) GetAuthorizers(ctx context.Context, input *apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAuthorizersWithContext(ctx, input)
}

func (a *Activities) GetDeployment(ctx context.Context, input *apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeploymentWithContext(ctx, input)
}

func (a *Activities) GetDeployments(ctx context.Context, input *apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeploymentsWithContext(ctx, input)
}

func (a *Activities) GetDomainName(ctx context.Context, input *apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainNameWithContext(ctx, input)
}

func (a *Activities) GetDomainNames(ctx context.Context, input *apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainNamesWithContext(ctx, input)
}

func (a *Activities) GetIntegration(ctx context.Context, input *apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntegrationWithContext(ctx, input)
}

func (a *Activities) GetIntegrationResponse(ctx context.Context, input *apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntegrationResponseWithContext(ctx, input)
}

func (a *Activities) GetIntegrationResponses(ctx context.Context, input *apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntegrationResponsesWithContext(ctx, input)
}

func (a *Activities) GetIntegrations(ctx context.Context, input *apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntegrationsWithContext(ctx, input)
}

func (a *Activities) GetModel(ctx context.Context, input *apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetModelWithContext(ctx, input)
}

func (a *Activities) GetModelTemplate(ctx context.Context, input *apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetModelTemplateWithContext(ctx, input)
}

func (a *Activities) GetModels(ctx context.Context, input *apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetModelsWithContext(ctx, input)
}

func (a *Activities) GetRoute(ctx context.Context, input *apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRouteWithContext(ctx, input)
}

func (a *Activities) GetRouteResponse(ctx context.Context, input *apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRouteResponseWithContext(ctx, input)
}

func (a *Activities) GetRouteResponses(ctx context.Context, input *apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRouteResponsesWithContext(ctx, input)
}

func (a *Activities) GetRoutes(ctx context.Context, input *apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRoutesWithContext(ctx, input)
}

func (a *Activities) GetStage(ctx context.Context, input *apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetStageWithContext(ctx, input)
}

func (a *Activities) GetStages(ctx context.Context, input *apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetStagesWithContext(ctx, input)
}

func (a *Activities) GetTags(ctx context.Context, input *apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTagsWithContext(ctx, input)
}

func (a *Activities) GetVpcLink(ctx context.Context, input *apigatewayv2.GetVpcLinkInput) (*apigatewayv2.GetVpcLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVpcLinkWithContext(ctx, input)
}

func (a *Activities) GetVpcLinks(ctx context.Context, input *apigatewayv2.GetVpcLinksInput) (*apigatewayv2.GetVpcLinksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetVpcLinksWithContext(ctx, input)
}

func (a *Activities) ImportApi(ctx context.Context, input *apigatewayv2.ImportApiInput) (*apigatewayv2.ImportApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ImportApiWithContext(ctx, input)
}

func (a *Activities) ReimportApi(ctx context.Context, input *apigatewayv2.ReimportApiInput) (*apigatewayv2.ReimportApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ReimportApiWithContext(ctx, input)
}

func (a *Activities) ResetAuthorizersCache(ctx context.Context, input *apigatewayv2.ResetAuthorizersCacheInput) (*apigatewayv2.ResetAuthorizersCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetAuthorizersCacheWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateApi(ctx context.Context, input *apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApiWithContext(ctx, input)
}

func (a *Activities) UpdateApiMapping(ctx context.Context, input *apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApiMappingWithContext(ctx, input)
}

func (a *Activities) UpdateAuthorizer(ctx context.Context, input *apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAuthorizerWithContext(ctx, input)
}

func (a *Activities) UpdateDeployment(ctx context.Context, input *apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDeploymentWithContext(ctx, input)
}

func (a *Activities) UpdateDomainName(ctx context.Context, input *apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDomainNameWithContext(ctx, input)
}

func (a *Activities) UpdateIntegration(ctx context.Context, input *apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIntegrationWithContext(ctx, input)
}

func (a *Activities) UpdateIntegrationResponse(ctx context.Context, input *apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIntegrationResponseWithContext(ctx, input)
}

func (a *Activities) UpdateModel(ctx context.Context, input *apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateModelWithContext(ctx, input)
}

func (a *Activities) UpdateRoute(ctx context.Context, input *apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRouteWithContext(ctx, input)
}

func (a *Activities) UpdateRouteResponse(ctx context.Context, input *apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRouteResponseWithContext(ctx, input)
}

func (a *Activities) UpdateStage(ctx context.Context, input *apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateStageWithContext(ctx, input)
}

func (a *Activities) UpdateVpcLink(ctx context.Context, input *apigatewayv2.UpdateVpcLinkInput) (*apigatewayv2.UpdateVpcLinkOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVpcLinkWithContext(ctx, input)
}