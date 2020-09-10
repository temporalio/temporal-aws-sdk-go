package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/dax/daxiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type DAXActivities struct {
	client daxiface.DAXAPI
}

func NewDAXActivities(session *session.Session, config ...*aws.Config) *DAXActivities {
	client := dax.New(session, config...)
	return &DAXActivities{client: client}
}

func (a *DAXActivities) CreateCluster(ctx context.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
	return a.client.CreateClusterWithContext(ctx, input)
}

func (a *DAXActivities) CreateParameterGroup(ctx context.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
	return a.client.CreateParameterGroupWithContext(ctx, input)
}

func (a *DAXActivities) CreateSubnetGroup(ctx context.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
	return a.client.CreateSubnetGroupWithContext(ctx, input)
}

func (a *DAXActivities) DecreaseReplicationFactor(ctx context.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
	return a.client.DecreaseReplicationFactorWithContext(ctx, input)
}

func (a *DAXActivities) DeleteCluster(ctx context.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
	return a.client.DeleteClusterWithContext(ctx, input)
}

func (a *DAXActivities) DeleteParameterGroup(ctx context.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
	return a.client.DeleteParameterGroupWithContext(ctx, input)
}

func (a *DAXActivities) DeleteSubnetGroup(ctx context.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
	return a.client.DeleteSubnetGroupWithContext(ctx, input)
}

func (a *DAXActivities) DescribeClusters(ctx context.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
	return a.client.DescribeClustersWithContext(ctx, input)
}

func (a *DAXActivities) DescribeDefaultParameters(ctx context.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
	return a.client.DescribeDefaultParametersWithContext(ctx, input)
}

func (a *DAXActivities) DescribeEvents(ctx context.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *DAXActivities) DescribeParameterGroups(ctx context.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
	return a.client.DescribeParameterGroupsWithContext(ctx, input)
}

func (a *DAXActivities) DescribeParameters(ctx context.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
	return a.client.DescribeParametersWithContext(ctx, input)
}

func (a *DAXActivities) DescribeSubnetGroups(ctx context.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
	return a.client.DescribeSubnetGroupsWithContext(ctx, input)
}

func (a *DAXActivities) IncreaseReplicationFactor(ctx context.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
	return a.client.IncreaseReplicationFactorWithContext(ctx, input)
}

func (a *DAXActivities) ListTags(ctx context.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *DAXActivities) RebootNode(ctx context.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
	return a.client.RebootNodeWithContext(ctx, input)
}

func (a *DAXActivities) TagResource(ctx context.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *DAXActivities) UntagResource(ctx context.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DAXActivities) UpdateCluster(ctx context.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
	return a.client.UpdateClusterWithContext(ctx, input)
}

func (a *DAXActivities) UpdateParameterGroup(ctx context.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
	return a.client.UpdateParameterGroupWithContext(ctx, input)
}

func (a *DAXActivities) UpdateSubnetGroup(ctx context.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
	return a.client.UpdateSubnetGroupWithContext(ctx, input)
}
