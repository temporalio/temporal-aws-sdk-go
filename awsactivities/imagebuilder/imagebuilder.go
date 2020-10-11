// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package imagebuilder

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/imagebuilder/imagebuilderiface"
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
	client imagebuilderiface.ImagebuilderAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := imagebuilder.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (imagebuilderiface.ImagebuilderAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return imagebuilder.New(sess), nil
}

func (a *Activities) CancelImageCreation(ctx context.Context, input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CancelImageCreationWithContext(ctx, input)
}

func (a *Activities) CreateComponent(ctx context.Context, input *imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateComponentWithContext(ctx, input)
}

func (a *Activities) CreateDistributionConfiguration(ctx context.Context, input *imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateDistributionConfigurationWithContext(ctx, input)
}

func (a *Activities) CreateImage(ctx context.Context, input *imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateImageWithContext(ctx, input)
}

func (a *Activities) CreateImagePipeline(ctx context.Context, input *imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateImagePipelineWithContext(ctx, input)
}

func (a *Activities) CreateImageRecipe(ctx context.Context, input *imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateImageRecipeWithContext(ctx, input)
}

func (a *Activities) CreateInfrastructureConfiguration(ctx context.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateInfrastructureConfigurationWithContext(ctx, input)
}

func (a *Activities) DeleteComponent(ctx context.Context, input *imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteComponentWithContext(ctx, input)
}

func (a *Activities) DeleteDistributionConfiguration(ctx context.Context, input *imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDistributionConfigurationWithContext(ctx, input)
}

func (a *Activities) DeleteImage(ctx context.Context, input *imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteImageWithContext(ctx, input)
}

func (a *Activities) DeleteImagePipeline(ctx context.Context, input *imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteImagePipelineWithContext(ctx, input)
}

func (a *Activities) DeleteImageRecipe(ctx context.Context, input *imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteImageRecipeWithContext(ctx, input)
}

func (a *Activities) DeleteInfrastructureConfiguration(ctx context.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInfrastructureConfigurationWithContext(ctx, input)
}

func (a *Activities) GetComponent(ctx context.Context, input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetComponentWithContext(ctx, input)
}

func (a *Activities) GetComponentPolicy(ctx context.Context, input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetComponentPolicyWithContext(ctx, input)
}

func (a *Activities) GetDistributionConfiguration(ctx context.Context, input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDistributionConfigurationWithContext(ctx, input)
}

func (a *Activities) GetImage(ctx context.Context, input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImageWithContext(ctx, input)
}

func (a *Activities) GetImagePipeline(ctx context.Context, input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImagePipelineWithContext(ctx, input)
}

func (a *Activities) GetImagePolicy(ctx context.Context, input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImagePolicyWithContext(ctx, input)
}

func (a *Activities) GetImageRecipe(ctx context.Context, input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImageRecipeWithContext(ctx, input)
}

func (a *Activities) GetImageRecipePolicy(ctx context.Context, input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetImageRecipePolicyWithContext(ctx, input)
}

func (a *Activities) GetInfrastructureConfiguration(ctx context.Context, input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInfrastructureConfigurationWithContext(ctx, input)
}

func (a *Activities) ImportComponent(ctx context.Context, input *imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.ImportComponentWithContext(ctx, input)
}

func (a *Activities) ListComponentBuildVersions(ctx context.Context, input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListComponentBuildVersionsWithContext(ctx, input)
}

func (a *Activities) ListComponents(ctx context.Context, input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListComponentsWithContext(ctx, input)
}

func (a *Activities) ListDistributionConfigurations(ctx context.Context, input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDistributionConfigurationsWithContext(ctx, input)
}

func (a *Activities) ListImageBuildVersions(ctx context.Context, input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListImageBuildVersionsWithContext(ctx, input)
}

func (a *Activities) ListImagePipelineImages(ctx context.Context, input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListImagePipelineImagesWithContext(ctx, input)
}

func (a *Activities) ListImagePipelines(ctx context.Context, input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListImagePipelinesWithContext(ctx, input)
}

func (a *Activities) ListImageRecipes(ctx context.Context, input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListImageRecipesWithContext(ctx, input)
}

func (a *Activities) ListImages(ctx context.Context, input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListImagesWithContext(ctx, input)
}

func (a *Activities) ListInfrastructureConfigurations(ctx context.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInfrastructureConfigurationsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) PutComponentPolicy(ctx context.Context, input *imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutComponentPolicyWithContext(ctx, input)
}

func (a *Activities) PutImagePolicy(ctx context.Context, input *imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutImagePolicyWithContext(ctx, input)
}

func (a *Activities) PutImageRecipePolicy(ctx context.Context, input *imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutImageRecipePolicyWithContext(ctx, input)
}

func (a *Activities) StartImagePipelineExecution(ctx context.Context, input *imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.StartImagePipelineExecutionWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateDistributionConfiguration(ctx context.Context, input *imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateDistributionConfigurationWithContext(ctx, input)
}

func (a *Activities) UpdateImagePipeline(ctx context.Context, input *imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateImagePipelineWithContext(ctx, input)
}

func (a *Activities) UpdateInfrastructureConfiguration(ctx context.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateInfrastructureConfigurationWithContext(ctx, input)
}
