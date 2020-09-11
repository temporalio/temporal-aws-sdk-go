package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling/iotsecuretunnelingiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type IoTSecureTunnelingActivities struct {
	client iotsecuretunnelingiface.IoTSecureTunnelingAPI
}

func NewIoTSecureTunnelingActivities(session *session.Session, config ...*aws.Config) *IoTSecureTunnelingActivities {
	client := iotsecuretunneling.New(session, config...)
	return &IoTSecureTunnelingActivities{client: client}
}

func (a *IoTSecureTunnelingActivities) CloseTunnel(ctx context.Context, input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
	return a.client.CloseTunnelWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) DescribeTunnel(ctx context.Context, input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	return a.client.DescribeTunnelWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) ListTagsForResource(ctx context.Context, input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) ListTunnels(ctx context.Context, input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
	return a.client.ListTunnelsWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) OpenTunnel(ctx context.Context, input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error) {
	return a.client.OpenTunnelWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) TagResource(ctx context.Context, input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) UntagResource(ctx context.Context, input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
