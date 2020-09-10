package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/outposts"
	"github.com/aws/aws-sdk-go/service/outposts/outpostsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type OutpostsActivities struct {
	client outpostsiface.OutpostsAPI
}

func NewOutpostsActivities(session *session.Session, config ...*aws.Config) *OutpostsActivities {
	client := outposts.New(session, config...)
	return &OutpostsActivities{client: client}
}

func (a *OutpostsActivities) CreateOutpost(ctx context.Context, input *outposts.CreateOutpostInput) (*outposts.CreateOutpostOutput, error) {
	return a.client.CreateOutpostWithContext(ctx, input)
}

func (a *OutpostsActivities) DeleteOutpost(ctx context.Context, input *outposts.DeleteOutpostInput) (*outposts.DeleteOutpostOutput, error) {
	return a.client.DeleteOutpostWithContext(ctx, input)
}

func (a *OutpostsActivities) DeleteSite(ctx context.Context, input *outposts.DeleteSiteInput) (*outposts.DeleteSiteOutput, error) {
	return a.client.DeleteSiteWithContext(ctx, input)
}

func (a *OutpostsActivities) GetOutpost(ctx context.Context, input *outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error) {
	return a.client.GetOutpostWithContext(ctx, input)
}

func (a *OutpostsActivities) GetOutpostInstanceTypes(ctx context.Context, input *outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error) {
	return a.client.GetOutpostInstanceTypesWithContext(ctx, input)
}

func (a *OutpostsActivities) ListOutposts(ctx context.Context, input *outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error) {
	return a.client.ListOutpostsWithContext(ctx, input)
}

func (a *OutpostsActivities) ListSites(ctx context.Context, input *outposts.ListSitesInput) (*outposts.ListSitesOutput, error) {
	return a.client.ListSitesWithContext(ctx, input)
}
