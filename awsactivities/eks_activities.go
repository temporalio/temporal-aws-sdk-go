
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
)

type EKSActivities struct {
    client eksiface.EKSAPI
}

func NewEKSActivities(session *session.Session, config ...*aws.Config) *EKSActivities {
    client := eks.New(session, config...)
    return &EKSActivities{client: client}
}

func (a *EKSActivities) CreateCluster(input *eks.CreateClusterInput) (*eks.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}

func (a *EKSActivities) CreateFargateProfile(input *eks.CreateFargateProfileInput) (*eks.CreateFargateProfileOutput, error) {
    return a.client.CreateFargateProfile(input)
}

func (a *EKSActivities) CreateNodegroup(input *eks.CreateNodegroupInput) (*eks.CreateNodegroupOutput, error) {
    return a.client.CreateNodegroup(input)
}

func (a *EKSActivities) DeleteCluster(input *eks.DeleteClusterInput) (*eks.DeleteClusterOutput, error) {
    return a.client.DeleteCluster(input)
}

func (a *EKSActivities) DeleteFargateProfile(input *eks.DeleteFargateProfileInput) (*eks.DeleteFargateProfileOutput, error) {
    return a.client.DeleteFargateProfile(input)
}

func (a *EKSActivities) DeleteNodegroup(input *eks.DeleteNodegroupInput) (*eks.DeleteNodegroupOutput, error) {
    return a.client.DeleteNodegroup(input)
}

func (a *EKSActivities) DescribeCluster(input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
    return a.client.DescribeCluster(input)
}

func (a *EKSActivities) DescribeFargateProfile(input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error) {
    return a.client.DescribeFargateProfile(input)
}

func (a *EKSActivities) DescribeNodegroup(input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
    return a.client.DescribeNodegroup(input)
}

func (a *EKSActivities) DescribeUpdate(input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {
    return a.client.DescribeUpdate(input)
}

func (a *EKSActivities) ListClusters(input *eks.ListClustersInput) (*eks.ListClustersOutput, error) {
    return a.client.ListClusters(input)
}

func (a *EKSActivities) ListFargateProfiles(input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error) {
    return a.client.ListFargateProfiles(input)
}

func (a *EKSActivities) ListNodegroups(input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error) {
    return a.client.ListNodegroups(input)
}

func (a *EKSActivities) ListTagsForResource(input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *EKSActivities) ListUpdates(input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {
    return a.client.ListUpdates(input)
}

func (a *EKSActivities) TagResource(input *eks.TagResourceInput) (*eks.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *EKSActivities) UntagResource(input *eks.UntagResourceInput) (*eks.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *EKSActivities) UpdateClusterConfig(input *eks.UpdateClusterConfigInput) (*eks.UpdateClusterConfigOutput, error) {
    return a.client.UpdateClusterConfig(input)
}

func (a *EKSActivities) UpdateClusterVersion(input *eks.UpdateClusterVersionInput) (*eks.UpdateClusterVersionOutput, error) {
    return a.client.UpdateClusterVersion(input)
}

func (a *EKSActivities) UpdateNodegroupConfig(input *eks.UpdateNodegroupConfigInput) (*eks.UpdateNodegroupConfigOutput, error) {
    return a.client.UpdateNodegroupConfig(input)
}

func (a *EKSActivities) UpdateNodegroupVersion(input *eks.UpdateNodegroupVersionInput) (*eks.UpdateNodegroupVersionOutput, error) {
    return a.client.UpdateNodegroupVersion(input)
}

func (a *EKSActivities) WaitUntilClusterActive(input *eks.DescribeClusterInput) error {
    return a.client.WaitUntilClusterActive(input)
}

func (a *EKSActivities) WaitUntilClusterDeleted(input *eks.DescribeClusterInput) error {
    return a.client.WaitUntilClusterDeleted(input)
}

func (a *EKSActivities) WaitUntilNodegroupActive(input *eks.DescribeNodegroupInput) error {
    return a.client.WaitUntilNodegroupActive(input)
}

func (a *EKSActivities) WaitUntilNodegroupDeleted(input *eks.DescribeNodegroupInput) error {
    return a.client.WaitUntilNodegroupDeleted(input)
}
