package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codestarconnections"
	"go.temporal.io/sdk/workflow"
)

type CodeStarConnectionsClient interface {
       CreateConnection(ctx workflow.Context, input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error)
       CreateConnectionAsync(ctx workflow.Context, input *codestarconnections.CreateConnectionInput) *CodestarconnectionsCreateConnectionResult

       CreateHost(ctx workflow.Context, input *codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error)
       CreateHostAsync(ctx workflow.Context, input *codestarconnections.CreateHostInput) *CodestarconnectionsCreateHostResult

       DeleteConnection(ctx workflow.Context, input *codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error)
       DeleteConnectionAsync(ctx workflow.Context, input *codestarconnections.DeleteConnectionInput) *CodestarconnectionsDeleteConnectionResult

       DeleteHost(ctx workflow.Context, input *codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error)
       DeleteHostAsync(ctx workflow.Context, input *codestarconnections.DeleteHostInput) *CodestarconnectionsDeleteHostResult

       GetConnection(ctx workflow.Context, input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error)
       GetConnectionAsync(ctx workflow.Context, input *codestarconnections.GetConnectionInput) *CodestarconnectionsGetConnectionResult

       GetHost(ctx workflow.Context, input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error)
       GetHostAsync(ctx workflow.Context, input *codestarconnections.GetHostInput) *CodestarconnectionsGetHostResult

       ListConnections(ctx workflow.Context, input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error)
       ListConnectionsAsync(ctx workflow.Context, input *codestarconnections.ListConnectionsInput) *CodestarconnectionsListConnectionsResult

       ListHosts(ctx workflow.Context, input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error)
       ListHostsAsync(ctx workflow.Context, input *codestarconnections.ListHostsInput) *CodestarconnectionsListHostsResult

       ListTagsForResource(ctx workflow.Context, input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *codestarconnections.ListTagsForResourceInput) *CodestarconnectionsListTagsForResourceResult

       TagResource(ctx workflow.Context, input *codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *codestarconnections.TagResourceInput) *CodestarconnectionsTagResourceResult

       UntagResource(ctx workflow.Context, input *codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *codestarconnections.UntagResourceInput) *CodestarconnectionsUntagResourceResult
}

type CodeStarConnectionsStub struct {
}

func NewCodeStarConnectionsStub() CodeStarConnectionsClient {
    return &CodeStarConnectionsStub{}
}

type CodestarconnectionsCreateConnectionResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsCreateConnectionResult) Get(ctx workflow.Context) (*codestarconnections.CreateConnectionOutput, error) {
    var output codestarconnections.CreateConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsCreateHostResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsCreateHostResult) Get(ctx workflow.Context) (*codestarconnections.CreateHostOutput, error) {
    var output codestarconnections.CreateHostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsDeleteConnectionResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsDeleteConnectionResult) Get(ctx workflow.Context) (*codestarconnections.DeleteConnectionOutput, error) {
    var output codestarconnections.DeleteConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsDeleteHostResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsDeleteHostResult) Get(ctx workflow.Context) (*codestarconnections.DeleteHostOutput, error) {
    var output codestarconnections.DeleteHostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsGetConnectionResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsGetConnectionResult) Get(ctx workflow.Context) (*codestarconnections.GetConnectionOutput, error) {
    var output codestarconnections.GetConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsGetHostResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsGetHostResult) Get(ctx workflow.Context) (*codestarconnections.GetHostOutput, error) {
    var output codestarconnections.GetHostOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsListConnectionsResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsListConnectionsResult) Get(ctx workflow.Context) (*codestarconnections.ListConnectionsOutput, error) {
    var output codestarconnections.ListConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsListHostsResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsListHostsResult) Get(ctx workflow.Context) (*codestarconnections.ListHostsOutput, error) {
    var output codestarconnections.ListHostsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsListTagsForResourceResult) Get(ctx workflow.Context) (*codestarconnections.ListTagsForResourceOutput, error) {
    var output codestarconnections.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsTagResourceResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsTagResourceResult) Get(ctx workflow.Context) (*codestarconnections.TagResourceOutput, error) {
    var output codestarconnections.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CodestarconnectionsUntagResourceResult struct {
	Result workflow.Future
}

func (r *CodestarconnectionsUntagResourceResult) Get(ctx workflow.Context) (*codestarconnections.UntagResourceOutput, error) {
    var output codestarconnections.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) CreateConnection(ctx workflow.Context, input *codestarconnections.CreateConnectionInput) (*codestarconnections.CreateConnectionOutput, error) {
    var output codestarconnections.CreateConnectionOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.CreateConnection", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) CreateConnectionAsync(ctx workflow.Context, input *codestarconnections.CreateConnectionInput) *CodestarconnectionsCreateConnectionResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.CreateConnection", input)
    return &CodestarconnectionsCreateConnectionResult{Result: future}
}

func (a *CodeStarConnectionsStub) CreateHost(ctx workflow.Context, input *codestarconnections.CreateHostInput) (*codestarconnections.CreateHostOutput, error) {
    var output codestarconnections.CreateHostOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.CreateHost", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) CreateHostAsync(ctx workflow.Context, input *codestarconnections.CreateHostInput) *CodestarconnectionsCreateHostResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.CreateHost", input)
    return &CodestarconnectionsCreateHostResult{Result: future}
}

func (a *CodeStarConnectionsStub) DeleteConnection(ctx workflow.Context, input *codestarconnections.DeleteConnectionInput) (*codestarconnections.DeleteConnectionOutput, error) {
    var output codestarconnections.DeleteConnectionOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.DeleteConnection", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) DeleteConnectionAsync(ctx workflow.Context, input *codestarconnections.DeleteConnectionInput) *CodestarconnectionsDeleteConnectionResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.DeleteConnection", input)
    return &CodestarconnectionsDeleteConnectionResult{Result: future}
}

func (a *CodeStarConnectionsStub) DeleteHost(ctx workflow.Context, input *codestarconnections.DeleteHostInput) (*codestarconnections.DeleteHostOutput, error) {
    var output codestarconnections.DeleteHostOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.DeleteHost", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) DeleteHostAsync(ctx workflow.Context, input *codestarconnections.DeleteHostInput) *CodestarconnectionsDeleteHostResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.DeleteHost", input)
    return &CodestarconnectionsDeleteHostResult{Result: future}
}

func (a *CodeStarConnectionsStub) GetConnection(ctx workflow.Context, input *codestarconnections.GetConnectionInput) (*codestarconnections.GetConnectionOutput, error) {
    var output codestarconnections.GetConnectionOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.GetConnection", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) GetConnectionAsync(ctx workflow.Context, input *codestarconnections.GetConnectionInput) *CodestarconnectionsGetConnectionResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.GetConnection", input)
    return &CodestarconnectionsGetConnectionResult{Result: future}
}

func (a *CodeStarConnectionsStub) GetHost(ctx workflow.Context, input *codestarconnections.GetHostInput) (*codestarconnections.GetHostOutput, error) {
    var output codestarconnections.GetHostOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.GetHost", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) GetHostAsync(ctx workflow.Context, input *codestarconnections.GetHostInput) *CodestarconnectionsGetHostResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.GetHost", input)
    return &CodestarconnectionsGetHostResult{Result: future}
}

func (a *CodeStarConnectionsStub) ListConnections(ctx workflow.Context, input *codestarconnections.ListConnectionsInput) (*codestarconnections.ListConnectionsOutput, error) {
    var output codestarconnections.ListConnectionsOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.ListConnections", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) ListConnectionsAsync(ctx workflow.Context, input *codestarconnections.ListConnectionsInput) *CodestarconnectionsListConnectionsResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.ListConnections", input)
    return &CodestarconnectionsListConnectionsResult{Result: future}
}

func (a *CodeStarConnectionsStub) ListHosts(ctx workflow.Context, input *codestarconnections.ListHostsInput) (*codestarconnections.ListHostsOutput, error) {
    var output codestarconnections.ListHostsOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.ListHosts", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) ListHostsAsync(ctx workflow.Context, input *codestarconnections.ListHostsInput) *CodestarconnectionsListHostsResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.ListHosts", input)
    return &CodestarconnectionsListHostsResult{Result: future}
}

func (a *CodeStarConnectionsStub) ListTagsForResource(ctx workflow.Context, input *codestarconnections.ListTagsForResourceInput) (*codestarconnections.ListTagsForResourceOutput, error) {
    var output codestarconnections.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) ListTagsForResourceAsync(ctx workflow.Context, input *codestarconnections.ListTagsForResourceInput) *CodestarconnectionsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.ListTagsForResource", input)
    return &CodestarconnectionsListTagsForResourceResult{Result: future}
}

func (a *CodeStarConnectionsStub) TagResource(ctx workflow.Context, input *codestarconnections.TagResourceInput) (*codestarconnections.TagResourceOutput, error) {
    var output codestarconnections.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) TagResourceAsync(ctx workflow.Context, input *codestarconnections.TagResourceInput) *CodestarconnectionsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.TagResource", input)
    return &CodestarconnectionsTagResourceResult{Result: future}
}

func (a *CodeStarConnectionsStub) UntagResource(ctx workflow.Context, input *codestarconnections.UntagResourceInput) (*codestarconnections.UntagResourceOutput, error) {
    var output codestarconnections.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "CodeStarConnections.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *CodeStarConnectionsStub) UntagResourceAsync(ctx workflow.Context, input *codestarconnections.UntagResourceInput) *CodestarconnectionsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "CodeStarConnections.UntagResource", input)
    return &CodestarconnectionsUntagResourceResult{Result: future}
}
