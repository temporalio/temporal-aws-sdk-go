
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"github.com/aws/aws-sdk-go/service/elasticinference/elasticinferenceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ElasticInferenceActivities struct {
    client elasticinferenceiface.ElasticInferenceAPI
}

func NewElasticInferenceActivities(session *session.Session, config ...*aws.Config) *ElasticInferenceActivities {
    client := elasticinference.New(session, config...)
    return &ElasticInferenceActivities{client: client}
}

func (a *ElasticInferenceActivities) DescribeAcceleratorOfferings(ctx context.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
    return a.client.DescribeAcceleratorOfferingsWithContext(ctx, input)
}

func (a *ElasticInferenceActivities) DescribeAcceleratorTypes(ctx context.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
    return a.client.DescribeAcceleratorTypesWithContext(ctx, input)
}

func (a *ElasticInferenceActivities) DescribeAccelerators(ctx context.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error) {
    return a.client.DescribeAcceleratorsWithContext(ctx, input)
}

func (a *ElasticInferenceActivities) ListTagsForResource(ctx context.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ElasticInferenceActivities) TagResource(ctx context.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *ElasticInferenceActivities) UntagResource(ctx context.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}
