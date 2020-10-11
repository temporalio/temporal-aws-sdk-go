// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package lambda

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
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
	client lambdaiface.LambdaAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := lambda.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (lambdaiface.LambdaAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return lambda.New(sess), nil
}

func (a *Activities) AddLayerVersionPermission(ctx context.Context, input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddLayerVersionPermissionWithContext(ctx, input)
}

func (a *Activities) AddPermission(ctx context.Context, input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddPermissionWithContext(ctx, input)
}

func (a *Activities) CreateAlias(ctx context.Context, input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAliasWithContext(ctx, input)
}

func (a *Activities) CreateEventSourceMapping(ctx context.Context, input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEventSourceMappingWithContext(ctx, input)
}

func (a *Activities) CreateFunction(ctx context.Context, input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFunctionWithContext(ctx, input)
}

func (a *Activities) DeleteAlias(ctx context.Context, input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAliasWithContext(ctx, input)
}

func (a *Activities) DeleteEventSourceMapping(ctx context.Context, input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEventSourceMappingWithContext(ctx, input)
}

func (a *Activities) DeleteFunction(ctx context.Context, input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFunctionWithContext(ctx, input)
}

func (a *Activities) DeleteFunctionConcurrency(ctx context.Context, input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFunctionConcurrencyWithContext(ctx, input)
}

func (a *Activities) DeleteFunctionEventInvokeConfig(ctx context.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *Activities) DeleteLayerVersion(ctx context.Context, input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLayerVersionWithContext(ctx, input)
}

func (a *Activities) DeleteProvisionedConcurrencyConfig(ctx context.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProvisionedConcurrencyConfigWithContext(ctx, input)
}

func (a *Activities) GetAccountSettings(ctx context.Context, input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAccountSettingsWithContext(ctx, input)
}

func (a *Activities) GetAlias(ctx context.Context, input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAliasWithContext(ctx, input)
}

func (a *Activities) GetEventSourceMapping(ctx context.Context, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEventSourceMappingWithContext(ctx, input)
}

func (a *Activities) GetFunction(ctx context.Context, input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionWithContext(ctx, input)
}

func (a *Activities) GetFunctionConcurrency(ctx context.Context, input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionConcurrencyWithContext(ctx, input)
}

func (a *Activities) GetFunctionConfiguration(ctx context.Context, input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionConfigurationWithContext(ctx, input)
}

func (a *Activities) GetFunctionEventInvokeConfig(ctx context.Context, input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *Activities) GetLayerVersion(ctx context.Context, input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLayerVersionWithContext(ctx, input)
}

func (a *Activities) GetLayerVersionByArn(ctx context.Context, input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLayerVersionByArnWithContext(ctx, input)
}

func (a *Activities) GetLayerVersionPolicy(ctx context.Context, input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLayerVersionPolicyWithContext(ctx, input)
}

func (a *Activities) GetPolicy(ctx context.Context, input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPolicyWithContext(ctx, input)
}

func (a *Activities) GetProvisionedConcurrencyConfig(ctx context.Context, input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetProvisionedConcurrencyConfigWithContext(ctx, input)
}

func (a *Activities) Invoke(ctx context.Context, input *lambda.InvokeInput) (*lambda.InvokeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InvokeWithContext(ctx, input)
}

func (a *Activities) ListAliases(ctx context.Context, input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAliasesWithContext(ctx, input)
}

func (a *Activities) ListEventSourceMappings(ctx context.Context, input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEventSourceMappingsWithContext(ctx, input)
}

func (a *Activities) ListFunctionEventInvokeConfigs(ctx context.Context, input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFunctionEventInvokeConfigsWithContext(ctx, input)
}

func (a *Activities) ListFunctions(ctx context.Context, input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFunctionsWithContext(ctx, input)
}

func (a *Activities) ListLayerVersions(ctx context.Context, input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLayerVersionsWithContext(ctx, input)
}

func (a *Activities) ListLayers(ctx context.Context, input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLayersWithContext(ctx, input)
}

func (a *Activities) ListProvisionedConcurrencyConfigs(ctx context.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProvisionedConcurrencyConfigsWithContext(ctx, input)
}

func (a *Activities) ListTags(ctx context.Context, input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsWithContext(ctx, input)
}

func (a *Activities) ListVersionsByFunction(ctx context.Context, input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVersionsByFunctionWithContext(ctx, input)
}

func (a *Activities) PublishLayerVersion(ctx context.Context, input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PublishLayerVersionWithContext(ctx, input)
}

func (a *Activities) PublishVersion(ctx context.Context, input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PublishVersionWithContext(ctx, input)
}

func (a *Activities) PutFunctionConcurrency(ctx context.Context, input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutFunctionConcurrencyWithContext(ctx, input)
}

func (a *Activities) PutFunctionEventInvokeConfig(ctx context.Context, input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *Activities) PutProvisionedConcurrencyConfig(ctx context.Context, input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutProvisionedConcurrencyConfigWithContext(ctx, input)
}

func (a *Activities) RemoveLayerVersionPermission(ctx context.Context, input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveLayerVersionPermissionWithContext(ctx, input)
}

func (a *Activities) RemovePermission(ctx context.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemovePermissionWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateAlias(ctx context.Context, input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAliasWithContext(ctx, input)
}

func (a *Activities) UpdateEventSourceMapping(ctx context.Context, input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateEventSourceMappingWithContext(ctx, input)
}

func (a *Activities) UpdateFunctionCode(ctx context.Context, input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFunctionCodeWithContext(ctx, input)
}

func (a *Activities) UpdateFunctionConfiguration(ctx context.Context, input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFunctionConfigurationWithContext(ctx, input)
}

func (a *Activities) UpdateFunctionEventInvokeConfig(ctx context.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFunctionEventInvokeConfigWithContext(ctx, input)
}

func (a *Activities) WaitUntilFunctionActive(ctx context.Context, input *lambda.GetFunctionConfigurationInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilFunctionActiveWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilFunctionExists(ctx context.Context, input *lambda.GetFunctionInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilFunctionExistsWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilFunctionUpdated(ctx context.Context, input *lambda.GetFunctionConfigurationInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilFunctionUpdatedWithContext(ctx, input, options...)
	})
}