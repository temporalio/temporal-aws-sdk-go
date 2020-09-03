package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	"go.temporal.io/sdk/workflow"
)

type ECRClient interface {
    BatchCheckLayerAvailability(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error)
    BatchCheckLayerAvailabilityAsync(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) *EcrBatchCheckLayerAvailabilityResult

    BatchDeleteImage(ctx workflow.Context, input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error)
    BatchDeleteImageAsync(ctx workflow.Context, input *ecr.BatchDeleteImageInput) *EcrBatchDeleteImageResult

    BatchGetImage(ctx workflow.Context, input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error)
    BatchGetImageAsync(ctx workflow.Context, input *ecr.BatchGetImageInput) *EcrBatchGetImageResult

    CompleteLayerUpload(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error)
    CompleteLayerUploadAsync(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) *EcrCompleteLayerUploadResult

    CreateRepository(ctx workflow.Context, input *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error)
    CreateRepositoryAsync(ctx workflow.Context, input *ecr.CreateRepositoryInput) *EcrCreateRepositoryResult

    DeleteLifecyclePolicy(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error)
    DeleteLifecyclePolicyAsync(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) *EcrDeleteLifecyclePolicyResult

    DeleteRepository(ctx workflow.Context, input *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error)
    DeleteRepositoryAsync(ctx workflow.Context, input *ecr.DeleteRepositoryInput) *EcrDeleteRepositoryResult

    DeleteRepositoryPolicy(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error)
    DeleteRepositoryPolicyAsync(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) *EcrDeleteRepositoryPolicyResult

    DescribeImageScanFindings(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error)
    DescribeImageScanFindingsAsync(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) *EcrDescribeImageScanFindingsResult

    DescribeImages(ctx workflow.Context, input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error)
    DescribeImagesAsync(ctx workflow.Context, input *ecr.DescribeImagesInput) *EcrDescribeImagesResult

    DescribeRepositories(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error)
    DescribeRepositoriesAsync(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) *EcrDescribeRepositoriesResult

    GetAuthorizationToken(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)
    GetAuthorizationTokenAsync(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) *EcrGetAuthorizationTokenResult

    GetDownloadUrlForLayer(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error)
    GetDownloadUrlForLayerAsync(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) *EcrGetDownloadUrlForLayerResult

    GetLifecyclePolicy(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error)
    GetLifecyclePolicyAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) *EcrGetLifecyclePolicyResult

    GetLifecyclePolicyPreview(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error)
    GetLifecyclePolicyPreviewAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) *EcrGetLifecyclePolicyPreviewResult

    GetRepositoryPolicy(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error)
    GetRepositoryPolicyAsync(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) *EcrGetRepositoryPolicyResult

    InitiateLayerUpload(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error)
    InitiateLayerUploadAsync(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) *EcrInitiateLayerUploadResult

    ListImages(ctx workflow.Context, input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error)
    ListImagesAsync(ctx workflow.Context, input *ecr.ListImagesInput) *EcrListImagesResult

    ListTagsForResource(ctx workflow.Context, input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *ecr.ListTagsForResourceInput) *EcrListTagsForResourceResult

    PutImage(ctx workflow.Context, input *ecr.PutImageInput) (*ecr.PutImageOutput, error)
    PutImageAsync(ctx workflow.Context, input *ecr.PutImageInput) *EcrPutImageResult

    PutImageScanningConfiguration(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error)
    PutImageScanningConfigurationAsync(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) *EcrPutImageScanningConfigurationResult

    PutImageTagMutability(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error)
    PutImageTagMutabilityAsync(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) *EcrPutImageTagMutabilityResult

    PutLifecyclePolicy(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error)
    PutLifecyclePolicyAsync(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) *EcrPutLifecyclePolicyResult

    SetRepositoryPolicy(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error)
    SetRepositoryPolicyAsync(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) *EcrSetRepositoryPolicyResult

    StartImageScan(ctx workflow.Context, input *ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error)
    StartImageScanAsync(ctx workflow.Context, input *ecr.StartImageScanInput) *EcrStartImageScanResult

    StartLifecyclePolicyPreview(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error)
    StartLifecyclePolicyPreviewAsync(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) *EcrStartLifecyclePolicyPreviewResult

    TagResource(ctx workflow.Context, input *ecr.TagResourceInput) (*ecr.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *ecr.TagResourceInput) *EcrTagResourceResult

    UntagResource(ctx workflow.Context, input *ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *ecr.UntagResourceInput) *EcrUntagResourceResult

    UploadLayerPart(ctx workflow.Context, input *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error)
    UploadLayerPartAsync(ctx workflow.Context, input *ecr.UploadLayerPartInput) *EcrUploadLayerPartResult

    WaitUntilImageScanComplete(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) error
    WaitUntilLifecyclePolicyPreviewComplete(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) error}
type EcrBatchCheckLayerAvailabilityResult struct {
	Result workflow.Future
}

func (r *EcrBatchCheckLayerAvailabilityResult) Get(ctx workflow.Context) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
    var output ecr.BatchCheckLayerAvailabilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrBatchDeleteImageResult struct {
	Result workflow.Future
}

func (r *EcrBatchDeleteImageResult) Get(ctx workflow.Context) (*ecr.BatchDeleteImageOutput, error) {
    var output ecr.BatchDeleteImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrBatchGetImageResult struct {
	Result workflow.Future
}

func (r *EcrBatchGetImageResult) Get(ctx workflow.Context) (*ecr.BatchGetImageOutput, error) {
    var output ecr.BatchGetImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrCompleteLayerUploadResult struct {
	Result workflow.Future
}

func (r *EcrCompleteLayerUploadResult) Get(ctx workflow.Context) (*ecr.CompleteLayerUploadOutput, error) {
    var output ecr.CompleteLayerUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrCreateRepositoryResult struct {
	Result workflow.Future
}

func (r *EcrCreateRepositoryResult) Get(ctx workflow.Context) (*ecr.CreateRepositoryOutput, error) {
    var output ecr.CreateRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrDeleteLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *EcrDeleteLifecyclePolicyResult) Get(ctx workflow.Context) (*ecr.DeleteLifecyclePolicyOutput, error) {
    var output ecr.DeleteLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrDeleteRepositoryResult struct {
	Result workflow.Future
}

func (r *EcrDeleteRepositoryResult) Get(ctx workflow.Context) (*ecr.DeleteRepositoryOutput, error) {
    var output ecr.DeleteRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrDeleteRepositoryPolicyResult struct {
	Result workflow.Future
}

func (r *EcrDeleteRepositoryPolicyResult) Get(ctx workflow.Context) (*ecr.DeleteRepositoryPolicyOutput, error) {
    var output ecr.DeleteRepositoryPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrDescribeImageScanFindingsResult struct {
	Result workflow.Future
}

func (r *EcrDescribeImageScanFindingsResult) Get(ctx workflow.Context) (*ecr.DescribeImageScanFindingsOutput, error) {
    var output ecr.DescribeImageScanFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrDescribeImagesResult struct {
	Result workflow.Future
}

func (r *EcrDescribeImagesResult) Get(ctx workflow.Context) (*ecr.DescribeImagesOutput, error) {
    var output ecr.DescribeImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrDescribeRepositoriesResult struct {
	Result workflow.Future
}

func (r *EcrDescribeRepositoriesResult) Get(ctx workflow.Context) (*ecr.DescribeRepositoriesOutput, error) {
    var output ecr.DescribeRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrGetAuthorizationTokenResult struct {
	Result workflow.Future
}

func (r *EcrGetAuthorizationTokenResult) Get(ctx workflow.Context) (*ecr.GetAuthorizationTokenOutput, error) {
    var output ecr.GetAuthorizationTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrGetDownloadUrlForLayerResult struct {
	Result workflow.Future
}

func (r *EcrGetDownloadUrlForLayerResult) Get(ctx workflow.Context) (*ecr.GetDownloadUrlForLayerOutput, error) {
    var output ecr.GetDownloadUrlForLayerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrGetLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *EcrGetLifecyclePolicyResult) Get(ctx workflow.Context) (*ecr.GetLifecyclePolicyOutput, error) {
    var output ecr.GetLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrGetLifecyclePolicyPreviewResult struct {
	Result workflow.Future
}

func (r *EcrGetLifecyclePolicyPreviewResult) Get(ctx workflow.Context) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
    var output ecr.GetLifecyclePolicyPreviewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrGetRepositoryPolicyResult struct {
	Result workflow.Future
}

func (r *EcrGetRepositoryPolicyResult) Get(ctx workflow.Context) (*ecr.GetRepositoryPolicyOutput, error) {
    var output ecr.GetRepositoryPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrInitiateLayerUploadResult struct {
	Result workflow.Future
}

func (r *EcrInitiateLayerUploadResult) Get(ctx workflow.Context) (*ecr.InitiateLayerUploadOutput, error) {
    var output ecr.InitiateLayerUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrListImagesResult struct {
	Result workflow.Future
}

func (r *EcrListImagesResult) Get(ctx workflow.Context) (*ecr.ListImagesOutput, error) {
    var output ecr.ListImagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *EcrListTagsForResourceResult) Get(ctx workflow.Context) (*ecr.ListTagsForResourceOutput, error) {
    var output ecr.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrPutImageResult struct {
	Result workflow.Future
}

func (r *EcrPutImageResult) Get(ctx workflow.Context) (*ecr.PutImageOutput, error) {
    var output ecr.PutImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrPutImageScanningConfigurationResult struct {
	Result workflow.Future
}

func (r *EcrPutImageScanningConfigurationResult) Get(ctx workflow.Context) (*ecr.PutImageScanningConfigurationOutput, error) {
    var output ecr.PutImageScanningConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrPutImageTagMutabilityResult struct {
	Result workflow.Future
}

func (r *EcrPutImageTagMutabilityResult) Get(ctx workflow.Context) (*ecr.PutImageTagMutabilityOutput, error) {
    var output ecr.PutImageTagMutabilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrPutLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *EcrPutLifecyclePolicyResult) Get(ctx workflow.Context) (*ecr.PutLifecyclePolicyOutput, error) {
    var output ecr.PutLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrSetRepositoryPolicyResult struct {
	Result workflow.Future
}

func (r *EcrSetRepositoryPolicyResult) Get(ctx workflow.Context) (*ecr.SetRepositoryPolicyOutput, error) {
    var output ecr.SetRepositoryPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrStartImageScanResult struct {
	Result workflow.Future
}

func (r *EcrStartImageScanResult) Get(ctx workflow.Context) (*ecr.StartImageScanOutput, error) {
    var output ecr.StartImageScanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrStartLifecyclePolicyPreviewResult struct {
	Result workflow.Future
}

func (r *EcrStartLifecyclePolicyPreviewResult) Get(ctx workflow.Context) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
    var output ecr.StartLifecyclePolicyPreviewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrTagResourceResult struct {
	Result workflow.Future
}

func (r *EcrTagResourceResult) Get(ctx workflow.Context) (*ecr.TagResourceOutput, error) {
    var output ecr.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrUntagResourceResult struct {
	Result workflow.Future
}

func (r *EcrUntagResourceResult) Get(ctx workflow.Context) (*ecr.UntagResourceOutput, error) {
    var output ecr.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EcrUploadLayerPartResult struct {
	Result workflow.Future
}

func (r *EcrUploadLayerPartResult) Get(ctx workflow.Context) (*ecr.UploadLayerPartOutput, error) {
    var output ecr.UploadLayerPartOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ECRStub struct {
    activities ECRClient
}

func NewECRStub() ECRClient {
    return &ECRStub{}
}

func (a *ECRStub) BatchCheckLayerAvailability(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
    var output ecr.BatchCheckLayerAvailabilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchCheckLayerAvailability, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) BatchDeleteImage(ctx workflow.Context, input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
    var output ecr.BatchDeleteImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteImage, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) BatchGetImage(ctx workflow.Context, input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
    var output ecr.BatchGetImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetImage, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) CompleteLayerUpload(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error) {
    var output ecr.CompleteLayerUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CompleteLayerUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) CreateRepository(ctx workflow.Context, input *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error) {
    var output ecr.CreateRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) DeleteLifecyclePolicy(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error) {
    var output ecr.DeleteLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) DeleteRepository(ctx workflow.Context, input *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error) {
    var output ecr.DeleteRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) DeleteRepositoryPolicy(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error) {
    var output ecr.DeleteRepositoryPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRepositoryPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) DescribeImageScanFindings(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error) {
    var output ecr.DescribeImageScanFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImageScanFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) DescribeImages(ctx workflow.Context, input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error) {
    var output ecr.DescribeImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImages, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) DescribeRepositories(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
    var output ecr.DescribeRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRepositories, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) GetAuthorizationToken(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
    var output ecr.GetAuthorizationTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizationToken, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) GetDownloadUrlForLayer(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
    var output ecr.GetDownloadUrlForLayerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDownloadUrlForLayer, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) GetLifecyclePolicy(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error) {
    var output ecr.GetLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) GetLifecyclePolicyPreview(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
    var output ecr.GetLifecyclePolicyPreviewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLifecyclePolicyPreview, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) GetRepositoryPolicy(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
    var output ecr.GetRepositoryPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRepositoryPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) InitiateLayerUpload(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error) {
    var output ecr.InitiateLayerUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitiateLayerUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) ListImages(ctx workflow.Context, input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
    var output ecr.ListImagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListImages, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) ListTagsForResource(ctx workflow.Context, input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error) {
    var output ecr.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) PutImage(ctx workflow.Context, input *ecr.PutImageInput) (*ecr.PutImageOutput, error) {
    var output ecr.PutImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutImage, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) PutImageScanningConfiguration(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error) {
    var output ecr.PutImageScanningConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutImageScanningConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) PutImageTagMutability(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error) {
    var output ecr.PutImageTagMutabilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutImageTagMutability, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) PutLifecyclePolicy(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error) {
    var output ecr.PutLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) SetRepositoryPolicy(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error) {
    var output ecr.SetRepositoryPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetRepositoryPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) StartImageScan(ctx workflow.Context, input *ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error) {
    var output ecr.StartImageScanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartImageScan, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) StartLifecyclePolicyPreview(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
    var output ecr.StartLifecyclePolicyPreviewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartLifecyclePolicyPreview, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) TagResource(ctx workflow.Context, input *ecr.TagResourceInput) (*ecr.TagResourceOutput, error) {
    var output ecr.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) UntagResource(ctx workflow.Context, input *ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error) {
    var output ecr.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ECRStub) UploadLayerPart(ctx workflow.Context, input *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error) {
    var output ecr.UploadLayerPartOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadLayerPart, input).Get(ctx, &output)
    return &output, err
}


func (a *ECRStub) WaitUntilImageScanComplete(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) error {
    return a.client.WaitUntilImageScanComplete(input)
}


func (a *ECRStub) WaitUntilLifecyclePolicyPreviewComplete(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) error {
    return a.client.WaitUntilLifecyclePolicyPreviewComplete(input)
}
