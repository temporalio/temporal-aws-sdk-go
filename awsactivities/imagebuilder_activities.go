
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"github.com/aws/aws-sdk-go/service/imagebuilder/imagebuilderiface"
)

type ImagebuilderActivities struct {
    client imagebuilderiface.ImagebuilderAPI
}

func NewImagebuilderActivities(session *session.Session, config ...*aws.Config) *ImagebuilderActivities {
    client := imagebuilder.New(session, config...)
    return &ImagebuilderActivities{client: client}
}

func (a *ImagebuilderActivities) CancelImageCreation(input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error) {
    return a.client.CancelImageCreation(input)
}

func (a *ImagebuilderActivities) CreateComponent(input *imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error) {
    return a.client.CreateComponent(input)
}

func (a *ImagebuilderActivities) CreateDistributionConfiguration(input *imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error) {
    return a.client.CreateDistributionConfiguration(input)
}

func (a *ImagebuilderActivities) CreateImage(input *imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error) {
    return a.client.CreateImage(input)
}

func (a *ImagebuilderActivities) CreateImagePipeline(input *imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error) {
    return a.client.CreateImagePipeline(input)
}

func (a *ImagebuilderActivities) CreateImageRecipe(input *imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error) {
    return a.client.CreateImageRecipe(input)
}

func (a *ImagebuilderActivities) CreateInfrastructureConfiguration(input *imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error) {
    return a.client.CreateInfrastructureConfiguration(input)
}

func (a *ImagebuilderActivities) DeleteComponent(input *imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error) {
    return a.client.DeleteComponent(input)
}

func (a *ImagebuilderActivities) DeleteDistributionConfiguration(input *imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error) {
    return a.client.DeleteDistributionConfiguration(input)
}

func (a *ImagebuilderActivities) DeleteImage(input *imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error) {
    return a.client.DeleteImage(input)
}

func (a *ImagebuilderActivities) DeleteImagePipeline(input *imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error) {
    return a.client.DeleteImagePipeline(input)
}

func (a *ImagebuilderActivities) DeleteImageRecipe(input *imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error) {
    return a.client.DeleteImageRecipe(input)
}

func (a *ImagebuilderActivities) DeleteInfrastructureConfiguration(input *imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error) {
    return a.client.DeleteInfrastructureConfiguration(input)
}

func (a *ImagebuilderActivities) GetComponent(input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error) {
    return a.client.GetComponent(input)
}

func (a *ImagebuilderActivities) GetComponentPolicy(input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error) {
    return a.client.GetComponentPolicy(input)
}

func (a *ImagebuilderActivities) GetDistributionConfiguration(input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error) {
    return a.client.GetDistributionConfiguration(input)
}

func (a *ImagebuilderActivities) GetImage(input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error) {
    return a.client.GetImage(input)
}

func (a *ImagebuilderActivities) GetImagePipeline(input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error) {
    return a.client.GetImagePipeline(input)
}

func (a *ImagebuilderActivities) GetImagePolicy(input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error) {
    return a.client.GetImagePolicy(input)
}

func (a *ImagebuilderActivities) GetImageRecipe(input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error) {
    return a.client.GetImageRecipe(input)
}

func (a *ImagebuilderActivities) GetImageRecipePolicy(input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error) {
    return a.client.GetImageRecipePolicy(input)
}

func (a *ImagebuilderActivities) GetInfrastructureConfiguration(input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
    return a.client.GetInfrastructureConfiguration(input)
}

func (a *ImagebuilderActivities) ImportComponent(input *imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error) {
    return a.client.ImportComponent(input)
}

func (a *ImagebuilderActivities) ListComponentBuildVersions(input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
    return a.client.ListComponentBuildVersions(input)
}

func (a *ImagebuilderActivities) ListComponents(input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error) {
    return a.client.ListComponents(input)
}

func (a *ImagebuilderActivities) ListDistributionConfigurations(input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
    return a.client.ListDistributionConfigurations(input)
}

func (a *ImagebuilderActivities) ListImageBuildVersions(input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error) {
    return a.client.ListImageBuildVersions(input)
}

func (a *ImagebuilderActivities) ListImagePipelineImages(input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error) {
    return a.client.ListImagePipelineImages(input)
}

func (a *ImagebuilderActivities) ListImagePipelines(input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error) {
    return a.client.ListImagePipelines(input)
}

func (a *ImagebuilderActivities) ListImageRecipes(input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error) {
    return a.client.ListImageRecipes(input)
}

func (a *ImagebuilderActivities) ListImages(input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error) {
    return a.client.ListImages(input)
}

func (a *ImagebuilderActivities) ListInfrastructureConfigurations(input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
    return a.client.ListInfrastructureConfigurations(input)
}

func (a *ImagebuilderActivities) ListTagsForResource(input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ImagebuilderActivities) PutComponentPolicy(input *imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error) {
    return a.client.PutComponentPolicy(input)
}

func (a *ImagebuilderActivities) PutImagePolicy(input *imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error) {
    return a.client.PutImagePolicy(input)
}

func (a *ImagebuilderActivities) PutImageRecipePolicy(input *imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error) {
    return a.client.PutImageRecipePolicy(input)
}

func (a *ImagebuilderActivities) StartImagePipelineExecution(input *imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error) {
    return a.client.StartImagePipelineExecution(input)
}

func (a *ImagebuilderActivities) TagResource(input *imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ImagebuilderActivities) UntagResource(input *imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ImagebuilderActivities) UpdateDistributionConfiguration(input *imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error) {
    return a.client.UpdateDistributionConfiguration(input)
}

func (a *ImagebuilderActivities) UpdateImagePipeline(input *imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error) {
    return a.client.UpdateImagePipeline(input)
}

func (a *ImagebuilderActivities) UpdateInfrastructureConfiguration(input *imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error) {
    return a.client.UpdateInfrastructureConfiguration(input)
}
