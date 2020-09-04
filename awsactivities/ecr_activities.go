
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
)

type ECRActivities struct {
	client ecriface.ECRAPI
}

func NewECRActivities(client ecriface.ECRAPI) *ECRActivities {
    return &ECRActivities{client: client}
}

func (a *ECRActivities) BatchCheckLayerAvailability(input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
    return a.client.BatchCheckLayerAvailability(input)
}

func (a *ECRActivities) BatchDeleteImage(input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
    return a.client.BatchDeleteImage(input)
}

func (a *ECRActivities) BatchGetImage(input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
    return a.client.BatchGetImage(input)
}

func (a *ECRActivities) CompleteLayerUpload(input *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error) {
    return a.client.CompleteLayerUpload(input)
}

func (a *ECRActivities) CreateRepository(input *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error) {
    return a.client.CreateRepository(input)
}

func (a *ECRActivities) DeleteLifecyclePolicy(input *ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error) {
    return a.client.DeleteLifecyclePolicy(input)
}

func (a *ECRActivities) DeleteRepository(input *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error) {
    return a.client.DeleteRepository(input)
}

func (a *ECRActivities) DeleteRepositoryPolicy(input *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error) {
    return a.client.DeleteRepositoryPolicy(input)
}

func (a *ECRActivities) DescribeImageScanFindings(input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error) {
    return a.client.DescribeImageScanFindings(input)
}

func (a *ECRActivities) DescribeImages(input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error) {
    return a.client.DescribeImages(input)
}

func (a *ECRActivities) DescribeRepositories(input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
    return a.client.DescribeRepositories(input)
}

func (a *ECRActivities) GetAuthorizationToken(input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
    return a.client.GetAuthorizationToken(input)
}

func (a *ECRActivities) GetDownloadUrlForLayer(input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
    return a.client.GetDownloadUrlForLayer(input)
}

func (a *ECRActivities) GetLifecyclePolicy(input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error) {
    return a.client.GetLifecyclePolicy(input)
}

func (a *ECRActivities) GetLifecyclePolicyPreview(input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
    return a.client.GetLifecyclePolicyPreview(input)
}

func (a *ECRActivities) GetRepositoryPolicy(input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
    return a.client.GetRepositoryPolicy(input)
}

func (a *ECRActivities) InitiateLayerUpload(input *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error) {
    return a.client.InitiateLayerUpload(input)
}

func (a *ECRActivities) ListImages(input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
    return a.client.ListImages(input)
}

func (a *ECRActivities) ListTagsForResource(input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ECRActivities) PutImage(input *ecr.PutImageInput) (*ecr.PutImageOutput, error) {
    return a.client.PutImage(input)
}

func (a *ECRActivities) PutImageScanningConfiguration(input *ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error) {
    return a.client.PutImageScanningConfiguration(input)
}

func (a *ECRActivities) PutImageTagMutability(input *ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error) {
    return a.client.PutImageTagMutability(input)
}

func (a *ECRActivities) PutLifecyclePolicy(input *ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error) {
    return a.client.PutLifecyclePolicy(input)
}

func (a *ECRActivities) SetRepositoryPolicy(input *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error) {
    return a.client.SetRepositoryPolicy(input)
}

func (a *ECRActivities) StartImageScan(input *ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error) {
    return a.client.StartImageScan(input)
}

func (a *ECRActivities) StartLifecyclePolicyPreview(input *ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
    return a.client.StartLifecyclePolicyPreview(input)
}

func (a *ECRActivities) TagResource(input *ecr.TagResourceInput) (*ecr.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ECRActivities) UntagResource(input *ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ECRActivities) UploadLayerPart(input *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error) {
    return a.client.UploadLayerPart(input)
}

func (a *ECRActivities) WaitUntilImageScanComplete(input *ecr.DescribeImageScanFindingsInput) error {
    return a.client.WaitUntilImageScanComplete(input)
}

func (a *ECRActivities) WaitUntilLifecyclePolicyPreviewComplete(input *ecr.GetLifecyclePolicyPreviewInput) error {
    return a.client.WaitUntilLifecyclePolicyPreviewComplete(input)
}
