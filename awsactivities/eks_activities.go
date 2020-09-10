package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type EKSActivities struct {
	client eksiface.EKSAPI
}

func NewEKSActivities(session *session.Session, config ...*aws.Config) *EKSActivities {
	client := eks.New(session, config...)
	return &EKSActivities{client: client}
}

func (a *EKSActivities) CreateCluster(ctx context.Context, input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error) {
	return a.client.CreateClusterWithContext(ctx, input)
}

func (a *EKSActivities) CreateFargateProfile(ctx context.Context, input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error) {
	return a.client.CreateFargateProfileWithContext(ctx, input)
}

func (a *EKSActivities) CreateNodegroup(ctx context.Context, input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error) {
	return a.client.CreateNodegroupWithContext(ctx, input)
}

func (a *EKSActivities) DeleteCluster(ctx context.Context, input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error) {
	return a.client.DeleteClusterWithContext(ctx, input)
}

func (a *EKSActivities) DeleteFargateProfile(ctx context.Context, input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error) {
	return a.client.DeleteFargateProfileWithContext(ctx, input)
}

func (a *EKSActivities) DeleteNodegroup(ctx context.Context, input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error) {
	return a.client.DeleteNodegroupWithContext(ctx, input)
}

func (a *EKSActivities) DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	return a.client.DescribeClusterWithContext(ctx, input)
}

func (a *EKSActivities) DescribeFargateProfile(ctx context.Context, input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error) {
	return a.client.DescribeFargateProfileWithContext(ctx, input)
}

func (a *EKSActivities) DescribeNodegroup(ctx context.Context, input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
	return a.client.DescribeNodegroupWithContext(ctx, input)
}

func (a *EKSActivities) DescribeUpdate(ctx context.Context, input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {
	return a.client.DescribeUpdateWithContext(ctx, input)
}

func (a *EKSActivities) ListClusters(ctx context.Context, input *eks.ListClustersInput) (*eks.ListClustersOutput, error) {
	return a.client.ListClustersWithContext(ctx, input)
}

func (a *EKSActivities) ListFargateProfiles(ctx context.Context, input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error) {
	return a.client.ListFargateProfilesWithContext(ctx, input)
}

func (a *EKSActivities) ListNodegroups(ctx context.Context, input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error) {
	return a.client.ListNodegroupsWithContext(ctx, input)
}

func (a *EKSActivities) ListTagsForResource(ctx context.Context, input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *EKSActivities) ListUpdates(ctx context.Context, input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {
	return a.client.ListUpdatesWithContext(ctx, input)
}

func (a *EKSActivities) TagResource(ctx context.Context, input *eks.TagResourceInput) (*eks.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *EKSActivities) UntagResource(ctx context.Context, input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *EKSActivities) UpdateClusterConfig(ctx context.Context, input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {
	return a.client.UpdateClusterConfigWithContext(ctx, input)
}

func (a *EKSActivities) UpdateClusterVersion(ctx context.Context, input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error) {
	return a.client.UpdateClusterVersionWithContext(ctx, input)
}

func (a *EKSActivities) UpdateNodegroupConfig(ctx context.Context, input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {
	return a.client.UpdateNodegroupConfigWithContext(ctx, input)
}

func (a *EKSActivities) UpdateNodegroupVersion(ctx context.Context, input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error) {
	return a.client.UpdateNodegroupVersionWithContext(ctx, input)
}

func (a *EKSActivities) WaitUntilClusterActive(ctx context.Context, input *eks.DescribeClusterInput) error {
	return a.client.WaitUntilClusterActiveWithContext(ctx, input)

}

func (a *EKSActivities) WaitUntilClusterDeleted(ctx context.Context, input *eks.DescribeClusterInput) error {
	return a.client.WaitUntilClusterDeletedWithContext(ctx, input)

}

func (a *EKSActivities) WaitUntilNodegroupActive(ctx context.Context, input *eks.DescribeNodegroupInput) error {
	return a.client.WaitUntilNodegroupActiveWithContext(ctx, input)

}

func (a *EKSActivities) WaitUntilNodegroupDeleted(ctx context.Context, input *eks.DescribeNodegroupInput) error {
	return a.client.WaitUntilNodegroupDeletedWithContext(ctx, input)

}
