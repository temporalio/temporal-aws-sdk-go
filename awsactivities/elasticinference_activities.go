package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"github.com/aws/aws-sdk-go/service/elasticinference/elasticinferenceiface"
)

type ElasticInferenceActivities struct {
	client elasticinferenceiface.ElasticInferenceAPI
}

func NewElasticInferenceActivities(session *session.Session, config ...*aws.Config) *ElasticInferenceActivities {
	client := elasticinference.New(session, config...)
	return &ElasticInferenceActivities{client: client}
}

func (a *ElasticInferenceActivities) DescribeAcceleratorOfferings(input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	return a.client.DescribeAcceleratorOfferings(input)
}

func (a *ElasticInferenceActivities) DescribeAcceleratorTypes(input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	return a.client.DescribeAcceleratorTypes(input)
}

func (a *ElasticInferenceActivities) DescribeAccelerators(input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error) {
	return a.client.DescribeAccelerators(input)
}

func (a *ElasticInferenceActivities) ListTagsForResource(input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *ElasticInferenceActivities) TagResource(input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *ElasticInferenceActivities) UntagResource(input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}
