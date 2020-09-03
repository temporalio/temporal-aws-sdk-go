package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceCatalogClient interface {
    CancelChangeSet(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error)
    CancelChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) *MarketplacecatalogCancelChangeSetResult

    DescribeChangeSet(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error)
    DescribeChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) *MarketplacecatalogDescribeChangeSetResult

    DescribeEntity(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error)
    DescribeEntityAsync(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) *MarketplacecatalogDescribeEntityResult

    ListChangeSets(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error)
    ListChangeSetsAsync(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) *MarketplacecatalogListChangeSetsResult

    ListEntities(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error)
    ListEntitiesAsync(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) *MarketplacecatalogListEntitiesResult

    StartChangeSet(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error)
    StartChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) *MarketplacecatalogStartChangeSetResult
}
type MarketplacecatalogCancelChangeSetResult struct {
	Result workflow.Future
}

func (r *MarketplacecatalogCancelChangeSetResult) Get(ctx workflow.Context) (*marketplacecatalog.CancelChangeSetOutput, error) {
    var output marketplacecatalog.CancelChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacecatalogDescribeChangeSetResult struct {
	Result workflow.Future
}

func (r *MarketplacecatalogDescribeChangeSetResult) Get(ctx workflow.Context) (*marketplacecatalog.DescribeChangeSetOutput, error) {
    var output marketplacecatalog.DescribeChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacecatalogDescribeEntityResult struct {
	Result workflow.Future
}

func (r *MarketplacecatalogDescribeEntityResult) Get(ctx workflow.Context) (*marketplacecatalog.DescribeEntityOutput, error) {
    var output marketplacecatalog.DescribeEntityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacecatalogListChangeSetsResult struct {
	Result workflow.Future
}

func (r *MarketplacecatalogListChangeSetsResult) Get(ctx workflow.Context) (*marketplacecatalog.ListChangeSetsOutput, error) {
    var output marketplacecatalog.ListChangeSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacecatalogListEntitiesResult struct {
	Result workflow.Future
}

func (r *MarketplacecatalogListEntitiesResult) Get(ctx workflow.Context) (*marketplacecatalog.ListEntitiesOutput, error) {
    var output marketplacecatalog.ListEntitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MarketplacecatalogStartChangeSetResult struct {
	Result workflow.Future
}

func (r *MarketplacecatalogStartChangeSetResult) Get(ctx workflow.Context) (*marketplacecatalog.StartChangeSetOutput, error) {
    var output marketplacecatalog.StartChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MarketplaceCatalogStub struct {
    activities MarketplaceCatalogClient
}

func NewMarketplaceCatalogStub() MarketplaceCatalogClient {
    return &MarketplaceCatalogStub{}
}

func (a *MarketplaceCatalogStub) CancelChangeSet(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error) {
    var output marketplacecatalog.CancelChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceCatalogStub) DescribeChangeSet(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error) {
    var output marketplacecatalog.DescribeChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceCatalogStub) DescribeEntity(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error) {
    var output marketplacecatalog.DescribeEntityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEntity, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceCatalogStub) ListChangeSets(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error) {
    var output marketplacecatalog.ListChangeSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChangeSets, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceCatalogStub) ListEntities(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error) {
    var output marketplacecatalog.ListEntitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEntities, input).Get(ctx, &output)
    return &output, err
}

func (a *MarketplaceCatalogStub) StartChangeSet(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error) {
    var output marketplacecatalog.StartChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartChangeSet, input).Get(ctx, &output)
    return &output, err
}
