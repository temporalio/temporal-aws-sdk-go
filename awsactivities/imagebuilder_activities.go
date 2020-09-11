package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/imagebuilder/imagebuilderiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ImagebuilderActivities struct {
	client imagebuilderiface.ImagebuilderAPI
}

func NewImagebuilderActivities(session *session.Session, config ...*aws.Config) *ImagebuilderActivities {
	client := imagebuilder.New(session, config...)
	return &ImagebuilderActivities{client: client}
}

func (a *ImagebuilderActivities) CancelImageCreation(ctx context.Context, input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CancelImageCreationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) CreateComponent(ctx context.Context, input *imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateComponentWithContext(ctx, input)
}

func (a *ImagebuilderActivities) CreateDistributionConfiguration(ctx context.Context, input *imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateDistributionConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) CreateImage(ctx context.Context, input *imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateImageWithContext(ctx, input)
}

func (a *ImagebuilderActivities) CreateImagePipeline(ctx context.Context, input *imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateImagePipelineWithContext(ctx, input)
}

func (a *ImagebuilderActivities) CreateImageRecipe(ctx context.Context, input *imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateImageRecipeWithContext(ctx, input)
}

func (a *ImagebuilderActivities) CreateInfrastructureConfiguration(ctx context.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateInfrastructureConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) DeleteComponent(ctx context.Context, input *imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error) {
	return a.client.DeleteComponentWithContext(ctx, input)
}

func (a *ImagebuilderActivities) DeleteDistributionConfiguration(ctx context.Context, input *imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error) {
	return a.client.DeleteDistributionConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) DeleteImage(ctx context.Context, input *imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error) {
	return a.client.DeleteImageWithContext(ctx, input)
}

func (a *ImagebuilderActivities) DeleteImagePipeline(ctx context.Context, input *imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error) {
	return a.client.DeleteImagePipelineWithContext(ctx, input)
}

func (a *ImagebuilderActivities) DeleteImageRecipe(ctx context.Context, input *imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error) {
	return a.client.DeleteImageRecipeWithContext(ctx, input)
}

func (a *ImagebuilderActivities) DeleteInfrastructureConfiguration(ctx context.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error) {
	return a.client.DeleteInfrastructureConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetComponent(ctx context.Context, input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error) {
	return a.client.GetComponentWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetComponentPolicy(ctx context.Context, input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error) {
	return a.client.GetComponentPolicyWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetDistributionConfiguration(ctx context.Context, input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error) {
	return a.client.GetDistributionConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetImage(ctx context.Context, input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error) {
	return a.client.GetImageWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetImagePipeline(ctx context.Context, input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error) {
	return a.client.GetImagePipelineWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetImagePolicy(ctx context.Context, input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error) {
	return a.client.GetImagePolicyWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetImageRecipe(ctx context.Context, input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error) {
	return a.client.GetImageRecipeWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetImageRecipePolicy(ctx context.Context, input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error) {
	return a.client.GetImageRecipePolicyWithContext(ctx, input)
}

func (a *ImagebuilderActivities) GetInfrastructureConfiguration(ctx context.Context, input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
	return a.client.GetInfrastructureConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ImportComponent(ctx context.Context, input *imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.ImportComponentWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListComponentBuildVersions(ctx context.Context, input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
	return a.client.ListComponentBuildVersionsWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListComponents(ctx context.Context, input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error) {
	return a.client.ListComponentsWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListDistributionConfigurations(ctx context.Context, input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
	return a.client.ListDistributionConfigurationsWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListImageBuildVersions(ctx context.Context, input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error) {
	return a.client.ListImageBuildVersionsWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListImagePipelineImages(ctx context.Context, input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error) {
	return a.client.ListImagePipelineImagesWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListImagePipelines(ctx context.Context, input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error) {
	return a.client.ListImagePipelinesWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListImageRecipes(ctx context.Context, input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error) {
	return a.client.ListImageRecipesWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListImages(ctx context.Context, input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error) {
	return a.client.ListImagesWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListInfrastructureConfigurations(ctx context.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
	return a.client.ListInfrastructureConfigurationsWithContext(ctx, input)
}

func (a *ImagebuilderActivities) ListTagsForResource(ctx context.Context, input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ImagebuilderActivities) PutComponentPolicy(ctx context.Context, input *imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error) {
	return a.client.PutComponentPolicyWithContext(ctx, input)
}

func (a *ImagebuilderActivities) PutImagePolicy(ctx context.Context, input *imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error) {
	return a.client.PutImagePolicyWithContext(ctx, input)
}

func (a *ImagebuilderActivities) PutImageRecipePolicy(ctx context.Context, input *imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error) {
	return a.client.PutImageRecipePolicyWithContext(ctx, input)
}

func (a *ImagebuilderActivities) StartImagePipelineExecution(ctx context.Context, input *imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.StartImagePipelineExecutionWithContext(ctx, input)
}

func (a *ImagebuilderActivities) TagResource(ctx context.Context, input *imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ImagebuilderActivities) UntagResource(ctx context.Context, input *imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ImagebuilderActivities) UpdateDistributionConfiguration(ctx context.Context, input *imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateDistributionConfigurationWithContext(ctx, input)
}

func (a *ImagebuilderActivities) UpdateImagePipeline(ctx context.Context, input *imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateImagePipelineWithContext(ctx, input)
}

func (a *ImagebuilderActivities) UpdateInfrastructureConfiguration(ctx context.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateInfrastructureConfigurationWithContext(ctx, input)
}
