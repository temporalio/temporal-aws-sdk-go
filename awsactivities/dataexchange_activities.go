
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/dataexchange/dataexchangeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DataExchangeActivities struct {
    client dataexchangeiface.DataExchangeAPI
}

func NewDataExchangeActivities(session *session.Session, config ...*aws.Config) *DataExchangeActivities {
    client := dataexchange.New(session, config...)
    return &DataExchangeActivities{client: client}
}

func (a *DataExchangeActivities) CancelJob(ctx context.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
    return a.client.CancelJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) CreateDataSet(ctx context.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
    return a.client.CreateDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) CreateJob(ctx context.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
    return a.client.CreateJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) CreateRevision(ctx context.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
    return a.client.CreateRevisionWithContext(ctx, input)
}

func (a *DataExchangeActivities) DeleteAsset(ctx context.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
    return a.client.DeleteAssetWithContext(ctx, input)
}

func (a *DataExchangeActivities) DeleteDataSet(ctx context.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
    return a.client.DeleteDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) DeleteRevision(ctx context.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
    return a.client.DeleteRevisionWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetAsset(ctx context.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
    return a.client.GetAssetWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetDataSet(ctx context.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
    return a.client.GetDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetJob(ctx context.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
    return a.client.GetJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetRevision(ctx context.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
    return a.client.GetRevisionWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListDataSetRevisions(ctx context.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
    return a.client.ListDataSetRevisionsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListDataSets(ctx context.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
    return a.client.ListDataSetsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListJobs(ctx context.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
    return a.client.ListJobsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListRevisionAssets(ctx context.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
    return a.client.ListRevisionAssetsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListTagsForResource(ctx context.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DataExchangeActivities) StartJob(ctx context.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
    return a.client.StartJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) TagResource(ctx context.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *DataExchangeActivities) UntagResource(ctx context.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DataExchangeActivities) UpdateAsset(ctx context.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
    return a.client.UpdateAssetWithContext(ctx, input)
}

func (a *DataExchangeActivities) UpdateDataSet(ctx context.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
    return a.client.UpdateDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) UpdateRevision(ctx context.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
    return a.client.UpdateRevisionWithContext(ctx, input)
}
