
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog/marketplacecatalogiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MarketplaceCatalogActivities struct {
    client marketplacecatalogiface.MarketplaceCatalogAPI
}

func NewMarketplaceCatalogActivities(session *session.Session, config ...*aws.Config) *MarketplaceCatalogActivities {
    client := marketplacecatalog.New(session, config...)
    return &MarketplaceCatalogActivities{client: client}
}

func (a *MarketplaceCatalogActivities) CancelChangeSet(ctx context.Context, input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error) {
    return a.client.CancelChangeSetWithContext(ctx, input)
}

func (a *MarketplaceCatalogActivities) DescribeChangeSet(ctx context.Context, input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error) {
    return a.client.DescribeChangeSetWithContext(ctx, input)
}

func (a *MarketplaceCatalogActivities) DescribeEntity(ctx context.Context, input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error) {
    return a.client.DescribeEntityWithContext(ctx, input)
}

func (a *MarketplaceCatalogActivities) ListChangeSets(ctx context.Context, input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error) {
    return a.client.ListChangeSetsWithContext(ctx, input)
}

func (a *MarketplaceCatalogActivities) ListEntities(ctx context.Context, input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error) {
    return a.client.ListEntitiesWithContext(ctx, input)
}

func (a *MarketplaceCatalogActivities) StartChangeSet(ctx context.Context, input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error) {
    return a.client.StartChangeSetWithContext(ctx, input)
}
