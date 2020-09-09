
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog/marketplacecatalogiface"
)

type MarketplaceCatalogActivities struct {
    client marketplacecatalogiface.MarketplaceCatalogAPI
}

func NewMarketplaceCatalogActivities(session *session.Session, config ...*aws.Config) *MarketplaceCatalogActivities {
    client := marketplacecatalog.New(session, config...)
    return &MarketplaceCatalogActivities{client: client}
}

func (a *MarketplaceCatalogActivities) CancelChangeSet(input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error) {
    return a.client.CancelChangeSet(input)
}

func (a *MarketplaceCatalogActivities) DescribeChangeSet(input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error) {
    return a.client.DescribeChangeSet(input)
}

func (a *MarketplaceCatalogActivities) DescribeEntity(input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error) {
    return a.client.DescribeEntity(input)
}

func (a *MarketplaceCatalogActivities) ListChangeSets(input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error) {
    return a.client.ListChangeSets(input)
}

func (a *MarketplaceCatalogActivities) ListEntities(input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error) {
    return a.client.ListEntities(input)
}

func (a *MarketplaceCatalogActivities) StartChangeSet(input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error) {
    return a.client.StartChangeSet(input)
}
