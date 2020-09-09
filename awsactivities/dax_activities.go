
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/dax/daxiface"
)

type DAXActivities struct {
    client daxiface.DAXAPI
}

func NewDAXActivities(session *session.Session, config ...*aws.Config) *DAXActivities {
    client := dax.New(session, config...)
    return &DAXActivities{client: client}
}

func (a *DAXActivities) CreateCluster(input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
    return a.client.CreateCluster(input)
}

func (a *DAXActivities) CreateParameterGroup(input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
    return a.client.CreateParameterGroup(input)
}

func (a *DAXActivities) CreateSubnetGroup(input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
    return a.client.CreateSubnetGroup(input)
}

func (a *DAXActivities) DecreaseReplicationFactor(input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
    return a.client.DecreaseReplicationFactor(input)
}

func (a *DAXActivities) DeleteCluster(input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
    return a.client.DeleteCluster(input)
}

func (a *DAXActivities) DeleteParameterGroup(input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
    return a.client.DeleteParameterGroup(input)
}

func (a *DAXActivities) DeleteSubnetGroup(input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
    return a.client.DeleteSubnetGroup(input)
}

func (a *DAXActivities) DescribeClusters(input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
    return a.client.DescribeClusters(input)
}

func (a *DAXActivities) DescribeDefaultParameters(input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
    return a.client.DescribeDefaultParameters(input)
}

func (a *DAXActivities) DescribeEvents(input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}

func (a *DAXActivities) DescribeParameterGroups(input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
    return a.client.DescribeParameterGroups(input)
}

func (a *DAXActivities) DescribeParameters(input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
    return a.client.DescribeParameters(input)
}

func (a *DAXActivities) DescribeSubnetGroups(input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
    return a.client.DescribeSubnetGroups(input)
}

func (a *DAXActivities) IncreaseReplicationFactor(input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
    return a.client.IncreaseReplicationFactor(input)
}

func (a *DAXActivities) ListTags(input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *DAXActivities) RebootNode(input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
    return a.client.RebootNode(input)
}

func (a *DAXActivities) TagResource(input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *DAXActivities) UntagResource(input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *DAXActivities) UpdateCluster(input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
    return a.client.UpdateCluster(input)
}

func (a *DAXActivities) UpdateParameterGroup(input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
    return a.client.UpdateParameterGroup(input)
}

func (a *DAXActivities) UpdateSubnetGroup(input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
    return a.client.UpdateSubnetGroup(input)
}
