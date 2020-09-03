package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/imagebuilder"
	"go.temporal.io/sdk/workflow"
)

type ImagebuilderClient interface {
    CancelImageCreation(ctx workflow.Context, input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error)
    CancelImageCreationAsync(ctx workflow.Context, input *imagebuilder.CancelImageCreationInput) *ImagebuilderCancelImageCreationResult

    CreateComponent(ctx workflow.Context, input *imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error)
    CreateComponentAsync(ctx workflow.Context, input *imagebuilder.CreateComponentInput) *ImagebuilderCreateComponentResult

    CreateDistributionConfiguration(ctx workflow.Context, input *imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error)
    CreateDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.CreateDistributionConfigurationInput) *ImagebuilderCreateDistributionConfigurationResult

    CreateImage(ctx workflow.Context, input *imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error)
    CreateImageAsync(ctx workflow.Context, input *imagebuilder.CreateImageInput) *ImagebuilderCreateImageResult

    CreateImagePipeline(ctx workflow.Context, input *imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error)
    CreateImagePipelineAsync(ctx workflow.Context, input *imagebuilder.CreateImagePipelineInput) *ImagebuilderCreateImagePipelineResult

    CreateImageRecipe(ctx workflow.Context, input *imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error)
    CreateImageRecipeAsync(ctx workflow.Context, input *imagebuilder.CreateImageRecipeInput) *ImagebuilderCreateImageRecipeResult

    CreateInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error)
    CreateInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) *ImagebuilderCreateInfrastructureConfigurationResult

    DeleteComponent(ctx workflow.Context, input *imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error)
    DeleteComponentAsync(ctx workflow.Context, input *imagebuilder.DeleteComponentInput) *ImagebuilderDeleteComponentResult

    DeleteDistributionConfiguration(ctx workflow.Context, input *imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error)
    DeleteDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.DeleteDistributionConfigurationInput) *ImagebuilderDeleteDistributionConfigurationResult

    DeleteImage(ctx workflow.Context, input *imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error)
    DeleteImageAsync(ctx workflow.Context, input *imagebuilder.DeleteImageInput) *ImagebuilderDeleteImageResult

    DeleteImagePipeline(ctx workflow.Context, input *imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error)
    DeleteImagePipelineAsync(ctx workflow.Context, input *imagebuilder.DeleteImagePipelineInput) *ImagebuilderDeleteImagePipelineResult

    DeleteImageRecipe(ctx workflow.Context, input *imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error)
    DeleteImageRecipeAsync(ctx workflow.Context, input *imagebuilder.DeleteImageRecipeInput) *ImagebuilderDeleteImageRecipeResult

    DeleteInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error)
    DeleteInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) *ImagebuilderDeleteInfrastructureConfigurationResult

    GetComponent(ctx workflow.Context, input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error)
    GetComponentAsync(ctx workflow.Context, input *imagebuilder.GetComponentInput) *ImagebuilderGetComponentResult

    GetComponentPolicy(ctx workflow.Context, input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error)
    GetComponentPolicyAsync(ctx workflow.Context, input *imagebuilder.GetComponentPolicyInput) *ImagebuilderGetComponentPolicyResult

    GetDistributionConfiguration(ctx workflow.Context, input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error)
    GetDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.GetDistributionConfigurationInput) *ImagebuilderGetDistributionConfigurationResult

    GetImage(ctx workflow.Context, input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error)
    GetImageAsync(ctx workflow.Context, input *imagebuilder.GetImageInput) *ImagebuilderGetImageResult

    GetImagePipeline(ctx workflow.Context, input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error)
    GetImagePipelineAsync(ctx workflow.Context, input *imagebuilder.GetImagePipelineInput) *ImagebuilderGetImagePipelineResult

    GetImagePolicy(ctx workflow.Context, input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error)
    GetImagePolicyAsync(ctx workflow.Context, input *imagebuilder.GetImagePolicyInput) *ImagebuilderGetImagePolicyResult

    GetImageRecipe(ctx workflow.Context, input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error)
    GetImageRecipeAsync(ctx workflow.Context, input *imagebuilder.GetImageRecipeInput) *ImagebuilderGetImageRecipeResult

    GetImageRecipePolicy(ctx workflow.Context, input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error)
    GetImageRecipePolicyAsync(ctx workflow.Context, input *imagebuilder.GetImageRecipePolicyInput) *ImagebuilderGetImageRecipePolicyResult

    GetInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error)
    GetInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.GetInfrastructureConfigurationInput) *ImagebuilderGetInfrastructureConfigurationResult

    ImportComponent(ctx workflow.Context, input *imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error)
    ImportComponentAsync(ctx workflow.Context, input *imagebuilder.ImportComponentInput) *ImagebuilderImportComponentResult

    ListComponentBuildVersions(ctx workflow.Context, input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error)
    ListComponentBuildVersionsAsync(ctx workflow.Context, input *imagebuilder.ListComponentBuildVersionsInput) *ImagebuilderListComponentBuildVersionsResult

    ListComponents(ctx workflow.Context, input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error)
    ListComponentsAsync(ctx workflow.Context, input *imagebuilder.ListComponentsInput) *ImagebuilderListComponentsResult

    ListDistributionConfigurations(ctx workflow.Context, input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error)
    ListDistributionConfigurationsAsync(ctx workflow.Context, input *imagebuilder.ListDistributionConfigurationsInput) *ImagebuilderListDistributionConfigurationsResult

    ListImageBuildVersions(ctx workflow.Context, input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error)
    ListImageBuildVersionsAsync(ctx workflow.Context, input *imagebuilder.ListImageBuildVersionsInput) *ImagebuilderListImageBuildVersionsResult

    ListImagePipelineImages(ctx workflow.Context, input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error)
    ListImagePipelineImagesAsync(ctx workflow.Context, input *imagebuilder.ListImagePipelineImagesInput) *ImagebuilderListImagePipelineImagesResult

    ListImagePipelines(ctx workflow.Context, input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error)
    ListImagePipelinesAsync(ctx workflow.Context, input *imagebuilder.ListImagePipelinesInput) *ImagebuilderListImagePipelinesResult

    ListImageRecipes(ctx workflow.Context, input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error)
    ListImageRecipesAsync(ctx workflow.Context, input *imagebuilder.ListImageRecipesInput) *ImagebuilderListImageRecipesResult

    ListImages(ctx workflow.Context, input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error)
    ListImagesAsync(ctx workflow.Context, input *imagebuilder.ListImagesInput) *ImagebuilderListImagesResult

    ListInfrastructureConfigurations(ctx workflow.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error)
    ListInfrastructureConfigurationsAsync(ctx workflow.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) *ImagebuilderListInfrastructureConfigurationsResult

    ListTagsForResource(ctx workflow.Context, input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *imagebuilder.ListTagsForResourceInput) *ImagebuilderListTagsForResourceResult

    PutComponentPolicy(ctx workflow.Context, input *imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error)
    PutComponentPolicyAsync(ctx workflow.Context, input *imagebuilder.PutComponentPolicyInput) *ImagebuilderPutComponentPolicyResult

    PutImagePolicy(ctx workflow.Context, input *imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error)
    PutImagePolicyAsync(ctx workflow.Context, input *imagebuilder.PutImagePolicyInput) *ImagebuilderPutImagePolicyResult

    PutImageRecipePolicy(ctx workflow.Context, input *imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error)
    PutImageRecipePolicyAsync(ctx workflow.Context, input *imagebuilder.PutImageRecipePolicyInput) *ImagebuilderPutImageRecipePolicyResult

    StartImagePipelineExecution(ctx workflow.Context, input *imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error)
    StartImagePipelineExecutionAsync(ctx workflow.Context, input *imagebuilder.StartImagePipelineExecutionInput) *ImagebuilderStartImagePipelineExecutionResult

    TagResource(ctx workflow.Context, input *imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *imagebuilder.TagResourceInput) *ImagebuilderTagResourceResult

    UntagResource(ctx workflow.Context, input *imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *imagebuilder.UntagResourceInput) *ImagebuilderUntagResourceResult

    UpdateDistributionConfiguration(ctx workflow.Context, input *imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error)
    UpdateDistributionConfigurationAsync(ctx workflow.Context, input *imagebuilder.UpdateDistributionConfigurationInput) *ImagebuilderUpdateDistributionConfigurationResult

    UpdateImagePipeline(ctx workflow.Context, input *imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error)
    UpdateImagePipelineAsync(ctx workflow.Context, input *imagebuilder.UpdateImagePipelineInput) *ImagebuilderUpdateImagePipelineResult

    UpdateInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error)
    UpdateInfrastructureConfigurationAsync(ctx workflow.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) *ImagebuilderUpdateInfrastructureConfigurationResult
}
type ImagebuilderCancelImageCreationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCancelImageCreationResult) Get(ctx workflow.Context) (*imagebuilder.CancelImageCreationOutput, error) {
    var output imagebuilder.CancelImageCreationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderCreateComponentResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCreateComponentResult) Get(ctx workflow.Context) (*imagebuilder.CreateComponentOutput, error) {
    var output imagebuilder.CreateComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderCreateDistributionConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCreateDistributionConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.CreateDistributionConfigurationOutput, error) {
    var output imagebuilder.CreateDistributionConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderCreateImageResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCreateImageResult) Get(ctx workflow.Context) (*imagebuilder.CreateImageOutput, error) {
    var output imagebuilder.CreateImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderCreateImagePipelineResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCreateImagePipelineResult) Get(ctx workflow.Context) (*imagebuilder.CreateImagePipelineOutput, error) {
    var output imagebuilder.CreateImagePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderCreateImageRecipeResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCreateImageRecipeResult) Get(ctx workflow.Context) (*imagebuilder.CreateImageRecipeOutput, error) {
    var output imagebuilder.CreateImageRecipeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderCreateInfrastructureConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderCreateInfrastructureConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.CreateInfrastructureConfigurationOutput, error) {
    var output imagebuilder.CreateInfrastructureConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderDeleteComponentResult struct {
	Result workflow.Future
}

func (r *ImagebuilderDeleteComponentResult) Get(ctx workflow.Context) (*imagebuilder.DeleteComponentOutput, error) {
    var output imagebuilder.DeleteComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderDeleteDistributionConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderDeleteDistributionConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.DeleteDistributionConfigurationOutput, error) {
    var output imagebuilder.DeleteDistributionConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderDeleteImageResult struct {
	Result workflow.Future
}

func (r *ImagebuilderDeleteImageResult) Get(ctx workflow.Context) (*imagebuilder.DeleteImageOutput, error) {
    var output imagebuilder.DeleteImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderDeleteImagePipelineResult struct {
	Result workflow.Future
}

func (r *ImagebuilderDeleteImagePipelineResult) Get(ctx workflow.Context) (*imagebuilder.DeleteImagePipelineOutput, error) {
    var output imagebuilder.DeleteImagePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderDeleteImageRecipeResult struct {
	Result workflow.Future
}

func (r *ImagebuilderDeleteImageRecipeResult) Get(ctx workflow.Context) (*imagebuilder.DeleteImageRecipeOutput, error) {
    var output imagebuilder.DeleteImageRecipeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderDeleteInfrastructureConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderDeleteInfrastructureConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error) {
    var output imagebuilder.DeleteInfrastructureConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetComponentResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetComponentResult) Get(ctx workflow.Context) (*imagebuilder.GetComponentOutput, error) {
    var output imagebuilder.GetComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetComponentPolicyResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetComponentPolicyResult) Get(ctx workflow.Context) (*imagebuilder.GetComponentPolicyOutput, error) {
    var output imagebuilder.GetComponentPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetDistributionConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetDistributionConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.GetDistributionConfigurationOutput, error) {
    var output imagebuilder.GetDistributionConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetImageResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetImageResult) Get(ctx workflow.Context) (*imagebuilder.GetImageOutput, error) {
    var output imagebuilder.GetImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetImagePipelineResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetImagePipelineResult) Get(ctx workflow.Context) (*imagebuilder.GetImagePipelineOutput, error) {
    var output imagebuilder.GetImagePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetImagePolicyResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetImagePolicyResult) Get(ctx workflow.Context) (*imagebuilder.GetImagePolicyOutput, error) {
    var output imagebuilder.GetImagePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetImageRecipeResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetImageRecipeResult) Get(ctx workflow.Context) (*imagebuilder.GetImageRecipeOutput, error) {
    var output imagebuilder.GetImageRecipeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetImageRecipePolicyResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetImageRecipePolicyResult) Get(ctx workflow.Context) (*imagebuilder.GetImageRecipePolicyOutput, error) {
    var output imagebuilder.GetImageRecipePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderGetInfrastructureConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderGetInfrastructureConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
    var output imagebuilder.GetInfrastructureConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderImportComponentResult struct {
	Result workflow.Future
}

func (r *ImagebuilderImportComponentResult) Get(ctx workflow.Context) (*imagebuilder.ImportComponentOutput, error) {
    var output imagebuilder.ImportComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListComponentBuildVersionsResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListComponentBuildVersionsResult) Get(ctx workflow.Context) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
    var output imagebuilder.ListComponentBuildVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListComponentsResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListComponentsResult) Get(ctx workflow.Context) (*imagebuilder.ListComponentsOutput, error) {
    var output imagebuilder.ListComponentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListDistributionConfigurationsResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListDistributionConfigurationsResult) Get(ctx workflow.Context) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
    var output imagebuilder.ListDistributionConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListImageBuildVersionsResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListImageBuildVersionsResult) Get(ctx workflow.Context) (*imagebuilder.ListImageBuildVersionsOutput, error) {
    var output imagebuilder.ListImageBuildVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListImagePipelineImagesResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListImagePipelineImagesResult) Get(ctx workflow.Context) (*imagebuilder.ListImagePipelineImagesOutput, error) {
    var output imagebuilder.ListImagePipelineImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListImagePipelinesResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListImagePipelinesResult) Get(ctx workflow.Context) (*imagebuilder.ListImagePipelinesOutput, error) {
    var output imagebuilder.ListImagePipelinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListImageRecipesResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListImageRecipesResult) Get(ctx workflow.Context) (*imagebuilder.ListImageRecipesOutput, error) {
    var output imagebuilder.ListImageRecipesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListImagesResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListImagesResult) Get(ctx workflow.Context) (*imagebuilder.ListImagesOutput, error) {
    var output imagebuilder.ListImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListInfrastructureConfigurationsResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListInfrastructureConfigurationsResult) Get(ctx workflow.Context) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
    var output imagebuilder.ListInfrastructureConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ImagebuilderListTagsForResourceResult) Get(ctx workflow.Context) (*imagebuilder.ListTagsForResourceOutput, error) {
    var output imagebuilder.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderPutComponentPolicyResult struct {
	Result workflow.Future
}

func (r *ImagebuilderPutComponentPolicyResult) Get(ctx workflow.Context) (*imagebuilder.PutComponentPolicyOutput, error) {
    var output imagebuilder.PutComponentPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderPutImagePolicyResult struct {
	Result workflow.Future
}

func (r *ImagebuilderPutImagePolicyResult) Get(ctx workflow.Context) (*imagebuilder.PutImagePolicyOutput, error) {
    var output imagebuilder.PutImagePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderPutImageRecipePolicyResult struct {
	Result workflow.Future
}

func (r *ImagebuilderPutImageRecipePolicyResult) Get(ctx workflow.Context) (*imagebuilder.PutImageRecipePolicyOutput, error) {
    var output imagebuilder.PutImageRecipePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderStartImagePipelineExecutionResult struct {
	Result workflow.Future
}

func (r *ImagebuilderStartImagePipelineExecutionResult) Get(ctx workflow.Context) (*imagebuilder.StartImagePipelineExecutionOutput, error) {
    var output imagebuilder.StartImagePipelineExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderTagResourceResult struct {
	Result workflow.Future
}

func (r *ImagebuilderTagResourceResult) Get(ctx workflow.Context) (*imagebuilder.TagResourceOutput, error) {
    var output imagebuilder.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderUntagResourceResult struct {
	Result workflow.Future
}

func (r *ImagebuilderUntagResourceResult) Get(ctx workflow.Context) (*imagebuilder.UntagResourceOutput, error) {
    var output imagebuilder.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderUpdateDistributionConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderUpdateDistributionConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.UpdateDistributionConfigurationOutput, error) {
    var output imagebuilder.UpdateDistributionConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderUpdateImagePipelineResult struct {
	Result workflow.Future
}

func (r *ImagebuilderUpdateImagePipelineResult) Get(ctx workflow.Context) (*imagebuilder.UpdateImagePipelineOutput, error) {
    var output imagebuilder.UpdateImagePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ImagebuilderUpdateInfrastructureConfigurationResult struct {
	Result workflow.Future
}

func (r *ImagebuilderUpdateInfrastructureConfigurationResult) Get(ctx workflow.Context) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error) {
    var output imagebuilder.UpdateInfrastructureConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ImagebuilderStub struct {
    activities ImagebuilderClient
}

func NewImagebuilderStub() ImagebuilderClient {
    return &ImagebuilderStub{}
}

func (a *ImagebuilderStub) CancelImageCreation(ctx workflow.Context, input *imagebuilder.CancelImageCreationInput) (*imagebuilder.CancelImageCreationOutput, error) {
    var output imagebuilder.CancelImageCreationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelImageCreation, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) CreateComponent(ctx workflow.Context, input *imagebuilder.CreateComponentInput) (*imagebuilder.CreateComponentOutput, error) {
    var output imagebuilder.CreateComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) CreateDistributionConfiguration(ctx workflow.Context, input *imagebuilder.CreateDistributionConfigurationInput) (*imagebuilder.CreateDistributionConfigurationOutput, error) {
    var output imagebuilder.CreateDistributionConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDistributionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) CreateImage(ctx workflow.Context, input *imagebuilder.CreateImageInput) (*imagebuilder.CreateImageOutput, error) {
    var output imagebuilder.CreateImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImage, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) CreateImagePipeline(ctx workflow.Context, input *imagebuilder.CreateImagePipelineInput) (*imagebuilder.CreateImagePipelineOutput, error) {
    var output imagebuilder.CreateImagePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImagePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) CreateImageRecipe(ctx workflow.Context, input *imagebuilder.CreateImageRecipeInput) (*imagebuilder.CreateImageRecipeOutput, error) {
    var output imagebuilder.CreateImageRecipeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateImageRecipe, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) CreateInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.CreateInfrastructureConfigurationInput) (*imagebuilder.CreateInfrastructureConfigurationOutput, error) {
    var output imagebuilder.CreateInfrastructureConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInfrastructureConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) DeleteComponent(ctx workflow.Context, input *imagebuilder.DeleteComponentInput) (*imagebuilder.DeleteComponentOutput, error) {
    var output imagebuilder.DeleteComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) DeleteDistributionConfiguration(ctx workflow.Context, input *imagebuilder.DeleteDistributionConfigurationInput) (*imagebuilder.DeleteDistributionConfigurationOutput, error) {
    var output imagebuilder.DeleteDistributionConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDistributionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) DeleteImage(ctx workflow.Context, input *imagebuilder.DeleteImageInput) (*imagebuilder.DeleteImageOutput, error) {
    var output imagebuilder.DeleteImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteImage, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) DeleteImagePipeline(ctx workflow.Context, input *imagebuilder.DeleteImagePipelineInput) (*imagebuilder.DeleteImagePipelineOutput, error) {
    var output imagebuilder.DeleteImagePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteImagePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) DeleteImageRecipe(ctx workflow.Context, input *imagebuilder.DeleteImageRecipeInput) (*imagebuilder.DeleteImageRecipeOutput, error) {
    var output imagebuilder.DeleteImageRecipeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteImageRecipe, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) DeleteInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.DeleteInfrastructureConfigurationInput) (*imagebuilder.DeleteInfrastructureConfigurationOutput, error) {
    var output imagebuilder.DeleteInfrastructureConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInfrastructureConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetComponent(ctx workflow.Context, input *imagebuilder.GetComponentInput) (*imagebuilder.GetComponentOutput, error) {
    var output imagebuilder.GetComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetComponentPolicy(ctx workflow.Context, input *imagebuilder.GetComponentPolicyInput) (*imagebuilder.GetComponentPolicyOutput, error) {
    var output imagebuilder.GetComponentPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetComponentPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetDistributionConfiguration(ctx workflow.Context, input *imagebuilder.GetDistributionConfigurationInput) (*imagebuilder.GetDistributionConfigurationOutput, error) {
    var output imagebuilder.GetDistributionConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDistributionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetImage(ctx workflow.Context, input *imagebuilder.GetImageInput) (*imagebuilder.GetImageOutput, error) {
    var output imagebuilder.GetImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImage, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetImagePipeline(ctx workflow.Context, input *imagebuilder.GetImagePipelineInput) (*imagebuilder.GetImagePipelineOutput, error) {
    var output imagebuilder.GetImagePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImagePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetImagePolicy(ctx workflow.Context, input *imagebuilder.GetImagePolicyInput) (*imagebuilder.GetImagePolicyOutput, error) {
    var output imagebuilder.GetImagePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImagePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetImageRecipe(ctx workflow.Context, input *imagebuilder.GetImageRecipeInput) (*imagebuilder.GetImageRecipeOutput, error) {
    var output imagebuilder.GetImageRecipeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImageRecipe, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetImageRecipePolicy(ctx workflow.Context, input *imagebuilder.GetImageRecipePolicyInput) (*imagebuilder.GetImageRecipePolicyOutput, error) {
    var output imagebuilder.GetImageRecipePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetImageRecipePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) GetInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.GetInfrastructureConfigurationInput) (*imagebuilder.GetInfrastructureConfigurationOutput, error) {
    var output imagebuilder.GetInfrastructureConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInfrastructureConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ImportComponent(ctx workflow.Context, input *imagebuilder.ImportComponentInput) (*imagebuilder.ImportComponentOutput, error) {
    var output imagebuilder.ImportComponentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportComponent, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListComponentBuildVersions(ctx workflow.Context, input *imagebuilder.ListComponentBuildVersionsInput) (*imagebuilder.ListComponentBuildVersionsOutput, error) {
    var output imagebuilder.ListComponentBuildVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListComponentBuildVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListComponents(ctx workflow.Context, input *imagebuilder.ListComponentsInput) (*imagebuilder.ListComponentsOutput, error) {
    var output imagebuilder.ListComponentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListComponents, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListDistributionConfigurations(ctx workflow.Context, input *imagebuilder.ListDistributionConfigurationsInput) (*imagebuilder.ListDistributionConfigurationsOutput, error) {
    var output imagebuilder.ListDistributionConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDistributionConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListImageBuildVersions(ctx workflow.Context, input *imagebuilder.ListImageBuildVersionsInput) (*imagebuilder.ListImageBuildVersionsOutput, error) {
    var output imagebuilder.ListImageBuildVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImageBuildVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListImagePipelineImages(ctx workflow.Context, input *imagebuilder.ListImagePipelineImagesInput) (*imagebuilder.ListImagePipelineImagesOutput, error) {
    var output imagebuilder.ListImagePipelineImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImagePipelineImages, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListImagePipelines(ctx workflow.Context, input *imagebuilder.ListImagePipelinesInput) (*imagebuilder.ListImagePipelinesOutput, error) {
    var output imagebuilder.ListImagePipelinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImagePipelines, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListImageRecipes(ctx workflow.Context, input *imagebuilder.ListImageRecipesInput) (*imagebuilder.ListImageRecipesOutput, error) {
    var output imagebuilder.ListImageRecipesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImageRecipes, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListImages(ctx workflow.Context, input *imagebuilder.ListImagesInput) (*imagebuilder.ListImagesOutput, error) {
    var output imagebuilder.ListImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImages, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListInfrastructureConfigurations(ctx workflow.Context, input *imagebuilder.ListInfrastructureConfigurationsInput) (*imagebuilder.ListInfrastructureConfigurationsOutput, error) {
    var output imagebuilder.ListInfrastructureConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListInfrastructureConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) ListTagsForResource(ctx workflow.Context, input *imagebuilder.ListTagsForResourceInput) (*imagebuilder.ListTagsForResourceOutput, error) {
    var output imagebuilder.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) PutComponentPolicy(ctx workflow.Context, input *imagebuilder.PutComponentPolicyInput) (*imagebuilder.PutComponentPolicyOutput, error) {
    var output imagebuilder.PutComponentPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutComponentPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) PutImagePolicy(ctx workflow.Context, input *imagebuilder.PutImagePolicyInput) (*imagebuilder.PutImagePolicyOutput, error) {
    var output imagebuilder.PutImagePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutImagePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) PutImageRecipePolicy(ctx workflow.Context, input *imagebuilder.PutImageRecipePolicyInput) (*imagebuilder.PutImageRecipePolicyOutput, error) {
    var output imagebuilder.PutImageRecipePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutImageRecipePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) StartImagePipelineExecution(ctx workflow.Context, input *imagebuilder.StartImagePipelineExecutionInput) (*imagebuilder.StartImagePipelineExecutionOutput, error) {
    var output imagebuilder.StartImagePipelineExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartImagePipelineExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) TagResource(ctx workflow.Context, input *imagebuilder.TagResourceInput) (*imagebuilder.TagResourceOutput, error) {
    var output imagebuilder.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) UntagResource(ctx workflow.Context, input *imagebuilder.UntagResourceInput) (*imagebuilder.UntagResourceOutput, error) {
    var output imagebuilder.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) UpdateDistributionConfiguration(ctx workflow.Context, input *imagebuilder.UpdateDistributionConfigurationInput) (*imagebuilder.UpdateDistributionConfigurationOutput, error) {
    var output imagebuilder.UpdateDistributionConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDistributionConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) UpdateImagePipeline(ctx workflow.Context, input *imagebuilder.UpdateImagePipelineInput) (*imagebuilder.UpdateImagePipelineOutput, error) {
    var output imagebuilder.UpdateImagePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateImagePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *ImagebuilderStub) UpdateInfrastructureConfiguration(ctx workflow.Context, input *imagebuilder.UpdateInfrastructureConfigurationInput) (*imagebuilder.UpdateInfrastructureConfigurationOutput, error) {
    var output imagebuilder.UpdateInfrastructureConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateInfrastructureConfiguration, input).Get(ctx, &output)
    return &output, err
}
