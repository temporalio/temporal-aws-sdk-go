
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/service/simpledb/simpledbiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type SimpleDBActivities struct {
    client simpledbiface.SimpleDBAPI
}

func NewSimpleDBActivities(session *session.Session, config ...*aws.Config) *SimpleDBActivities {
    client := simpledb.New(session, config...)
    return &SimpleDBActivities{client: client}
}

func (a *SimpleDBActivities) BatchDeleteAttributes(ctx context.Context, input *simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error) {
    return a.client.BatchDeleteAttributesWithContext(ctx, input)
}

func (a *SimpleDBActivities) BatchPutAttributes(ctx context.Context, input *simpledb.BatchPutAttributesInput) (*simpledb.BatchPutAttributesOutput, error) {
    return a.client.BatchPutAttributesWithContext(ctx, input)
}

func (a *SimpleDBActivities) CreateDomain(ctx context.Context, input *simpledb.CreateDomainInput) (*simpledb.CreateDomainOutput, error) {
    return a.client.CreateDomainWithContext(ctx, input)
}

func (a *SimpleDBActivities) DeleteAttributes(ctx context.Context, input *simpledb.DeleteAttributesInput) (*simpledb.DeleteAttributesOutput, error) {
    return a.client.DeleteAttributesWithContext(ctx, input)
}

func (a *SimpleDBActivities) DeleteDomain(ctx context.Context, input *simpledb.DeleteDomainInput) (*simpledb.DeleteDomainOutput, error) {
    return a.client.DeleteDomainWithContext(ctx, input)
}

func (a *SimpleDBActivities) DomainMetadata(ctx context.Context, input *simpledb.DomainMetadataInput) (*simpledb.DomainMetadataOutput, error) {
    return a.client.DomainMetadataWithContext(ctx, input)
}

func (a *SimpleDBActivities) GetAttributes(ctx context.Context, input *simpledb.GetAttributesInput) (*simpledb.GetAttributesOutput, error) {
    return a.client.GetAttributesWithContext(ctx, input)
}

func (a *SimpleDBActivities) ListDomains(ctx context.Context, input *simpledb.ListDomainsInput) (*simpledb.ListDomainsOutput, error) {
    return a.client.ListDomainsWithContext(ctx, input)
}

func (a *SimpleDBActivities) PutAttributes(ctx context.Context, input *simpledb.PutAttributesInput) (*simpledb.PutAttributesOutput, error) {
    return a.client.PutAttributesWithContext(ctx, input)
}

func (a *SimpleDBActivities) Select(ctx context.Context, input *simpledb.SelectInput) (*simpledb.SelectOutput, error) {
    return a.client.SelectWithContext(ctx, input)
}
