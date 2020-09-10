package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CloudSearchDomainActivities struct {
	client cloudsearchdomainiface.CloudSearchDomainAPI
}

func NewCloudSearchDomainActivities(session *session.Session, config ...*aws.Config) *CloudSearchDomainActivities {
	client := cloudsearchdomain.New(session, config...)
	return &CloudSearchDomainActivities{client: client}
}

func (a *CloudSearchDomainActivities) Search(ctx context.Context, input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	return a.client.SearchWithContext(ctx, input)
}

func (a *CloudSearchDomainActivities) Suggest(ctx context.Context, input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
	return a.client.SuggestWithContext(ctx, input)
}

func (a *CloudSearchDomainActivities) UploadDocuments(ctx context.Context, input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	return a.client.UploadDocumentsWithContext(ctx, input)
}
