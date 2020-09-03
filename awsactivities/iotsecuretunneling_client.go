package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"go.temporal.io/sdk/workflow"
)

type IoTSecureTunnelingClient interface {
    CloseTunnel(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error)
    CloseTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) *IotsecuretunnelingCloseTunnelResult

    DescribeTunnel(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error)
    DescribeTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) *IotsecuretunnelingDescribeTunnelResult

    ListTagsForResource(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) *IotsecuretunnelingListTagsForResourceResult

    ListTunnels(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error)
    ListTunnelsAsync(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) *IotsecuretunnelingListTunnelsResult

    OpenTunnel(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error)
    OpenTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) *IotsecuretunnelingOpenTunnelResult

    TagResource(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) *IotsecuretunnelingTagResourceResult

    UntagResource(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) *IotsecuretunnelingUntagResourceResult
}
type IotsecuretunnelingCloseTunnelResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingCloseTunnelResult) Get(ctx workflow.Context) (*iotsecuretunneling.CloseTunnelOutput, error) {
    var output iotsecuretunneling.CloseTunnelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotsecuretunnelingDescribeTunnelResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingDescribeTunnelResult) Get(ctx workflow.Context) (*iotsecuretunneling.DescribeTunnelOutput, error) {
    var output iotsecuretunneling.DescribeTunnelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotsecuretunnelingListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingListTagsForResourceResult) Get(ctx workflow.Context) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
    var output iotsecuretunneling.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotsecuretunnelingListTunnelsResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingListTunnelsResult) Get(ctx workflow.Context) (*iotsecuretunneling.ListTunnelsOutput, error) {
    var output iotsecuretunneling.ListTunnelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotsecuretunnelingOpenTunnelResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingOpenTunnelResult) Get(ctx workflow.Context) (*iotsecuretunneling.OpenTunnelOutput, error) {
    var output iotsecuretunneling.OpenTunnelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotsecuretunnelingTagResourceResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingTagResourceResult) Get(ctx workflow.Context) (*iotsecuretunneling.TagResourceOutput, error) {
    var output iotsecuretunneling.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type IotsecuretunnelingUntagResourceResult struct {
	Result workflow.Future
}

func (r *IotsecuretunnelingUntagResourceResult) Get(ctx workflow.Context) (*iotsecuretunneling.UntagResourceOutput, error) {
    var output iotsecuretunneling.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type IoTSecureTunnelingStub struct {
    activities IoTSecureTunnelingClient
}

func NewIoTSecureTunnelingStub() IoTSecureTunnelingClient {
    return &IoTSecureTunnelingStub{}
}

func (a *IoTSecureTunnelingStub) CloseTunnel(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
    var output iotsecuretunneling.CloseTunnelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CloseTunnel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTSecureTunnelingStub) DescribeTunnel(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
    var output iotsecuretunneling.DescribeTunnelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTunnel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTSecureTunnelingStub) ListTagsForResource(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
    var output iotsecuretunneling.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTSecureTunnelingStub) ListTunnels(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
    var output iotsecuretunneling.ListTunnelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTunnels, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTSecureTunnelingStub) OpenTunnel(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error) {
    var output iotsecuretunneling.OpenTunnelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.OpenTunnel, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTSecureTunnelingStub) TagResource(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error) {
    var output iotsecuretunneling.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *IoTSecureTunnelingStub) UntagResource(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error) {
    var output iotsecuretunneling.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}
