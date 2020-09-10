package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ECRActivities struct {
	client ecriface.ECRAPI
}

func NewECRActivities(session *session.Session, config ...*aws.Config) *ECRActivities {
	client := ecr.New(session, config...)
	return &ECRActivities{client: client}
}

func (a *ECRActivities) BatchCheckLayerAvailability(ctx context.Context, input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	return a.client.BatchCheckLayerAvailabilityWithContext(ctx, input)
}

func (a *ECRActivities) BatchDeleteImage(ctx context.Context, input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
	return a.client.BatchDeleteImageWithContext(ctx, input)
}

func (a *ECRActivities) BatchGetImage(ctx context.Context, input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
	return a.client.BatchGetImageWithContext(ctx, input)
}

func (a *ECRActivities) CompleteLayerUpload(ctx context.Context, input *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error) {
	return a.client.CompleteLayerUploadWithContext(ctx, input)
}

func (a *ECRActivities) CreateRepository(ctx context.Context, input *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error) {
	return a.client.CreateRepositoryWithContext(ctx, input)
}

func (a *ECRActivities) DeleteLifecyclePolicy(ctx context.Context, input *ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error) {
	return a.client.DeleteLifecyclePolicyWithContext(ctx, input)
}

func (a *ECRActivities) DeleteRepository(ctx context.Context, input *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error) {
	return a.client.DeleteRepositoryWithContext(ctx, input)
}

func (a *ECRActivities) DeleteRepositoryPolicy(ctx context.Context, input *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error) {
	return a.client.DeleteRepositoryPolicyWithContext(ctx, input)
}

func (a *ECRActivities) DescribeImageScanFindings(ctx context.Context, input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error) {
	return a.client.DescribeImageScanFindingsWithContext(ctx, input)
}

func (a *ECRActivities) DescribeImages(ctx context.Context, input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error) {
	return a.client.DescribeImagesWithContext(ctx, input)
}

func (a *ECRActivities) DescribeRepositories(ctx context.Context, input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
	return a.client.DescribeRepositoriesWithContext(ctx, input)
}

func (a *ECRActivities) GetAuthorizationToken(ctx context.Context, input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	return a.client.GetAuthorizationTokenWithContext(ctx, input)
}

func (a *ECRActivities) GetDownloadUrlForLayer(ctx context.Context, input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
	return a.client.GetDownloadUrlForLayerWithContext(ctx, input)
}

func (a *ECRActivities) GetLifecyclePolicy(ctx context.Context, input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error) {
	return a.client.GetLifecyclePolicyWithContext(ctx, input)
}

func (a *ECRActivities) GetLifecyclePolicyPreview(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	return a.client.GetLifecyclePolicyPreviewWithContext(ctx, input)
}

func (a *ECRActivities) GetRepositoryPolicy(ctx context.Context, input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
	return a.client.GetRepositoryPolicyWithContext(ctx, input)
}

func (a *ECRActivities) InitiateLayerUpload(ctx context.Context, input *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error) {
	return a.client.InitiateLayerUploadWithContext(ctx, input)
}

func (a *ECRActivities) ListImages(ctx context.Context, input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
	return a.client.ListImagesWithContext(ctx, input)
}

func (a *ECRActivities) ListTagsForResource(ctx context.Context, input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ECRActivities) PutImage(ctx context.Context, input *ecr.PutImageInput) (*ecr.PutImageOutput, error) {
	return a.client.PutImageWithContext(ctx, input)
}

func (a *ECRActivities) PutImageScanningConfiguration(ctx context.Context, input *ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error) {
	return a.client.PutImageScanningConfigurationWithContext(ctx, input)
}

func (a *ECRActivities) PutImageTagMutability(ctx context.Context, input *ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error) {
	return a.client.PutImageTagMutabilityWithContext(ctx, input)
}

func (a *ECRActivities) PutLifecyclePolicy(ctx context.Context, input *ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error) {
	return a.client.PutLifecyclePolicyWithContext(ctx, input)
}

func (a *ECRActivities) SetRepositoryPolicy(ctx context.Context, input *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error) {
	return a.client.SetRepositoryPolicyWithContext(ctx, input)
}

func (a *ECRActivities) StartImageScan(ctx context.Context, input *ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error) {
	return a.client.StartImageScanWithContext(ctx, input)
}

func (a *ECRActivities) StartLifecyclePolicyPreview(ctx context.Context, input *ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
	return a.client.StartLifecyclePolicyPreviewWithContext(ctx, input)
}

func (a *ECRActivities) TagResource(ctx context.Context, input *ecr.TagResourceInput) (*ecr.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ECRActivities) UntagResource(ctx context.Context, input *ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ECRActivities) UploadLayerPart(ctx context.Context, input *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error) {
	return a.client.UploadLayerPartWithContext(ctx, input)
}

func (a *ECRActivities) WaitUntilImageScanComplete(ctx context.Context, input *ecr.DescribeImageScanFindingsInput) error {
	return a.client.WaitUntilImageScanCompleteWithContext(ctx, input)

}

func (a *ECRActivities) WaitUntilLifecyclePolicyPreviewComplete(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput) error {
	return a.client.WaitUntilLifecyclePolicyPreviewCompleteWithContext(ctx, input)

}
