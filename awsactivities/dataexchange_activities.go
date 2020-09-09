package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/dataexchange/dataexchangeiface"
)

type DataExchangeActivities struct {
	client dataexchangeiface.DataExchangeAPI
}

func NewDataExchangeActivities(session *session.Session, config ...*aws.Config) *DataExchangeActivities {
	client := dataexchange.New(session, config...)
	return &DataExchangeActivities{client: client}
}

func (a *DataExchangeActivities) CancelJob(input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
	return a.client.CancelJob(input)
}

func (a *DataExchangeActivities) CreateDataSet(input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
	return a.client.CreateDataSet(input)
}

func (a *DataExchangeActivities) CreateJob(input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
	return a.client.CreateJob(input)
}

func (a *DataExchangeActivities) CreateRevision(input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
	return a.client.CreateRevision(input)
}

func (a *DataExchangeActivities) DeleteAsset(input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
	return a.client.DeleteAsset(input)
}

func (a *DataExchangeActivities) DeleteDataSet(input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
	return a.client.DeleteDataSet(input)
}

func (a *DataExchangeActivities) DeleteRevision(input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
	return a.client.DeleteRevision(input)
}

func (a *DataExchangeActivities) GetAsset(input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
	return a.client.GetAsset(input)
}

func (a *DataExchangeActivities) GetDataSet(input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
	return a.client.GetDataSet(input)
}

func (a *DataExchangeActivities) GetJob(input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
	return a.client.GetJob(input)
}

func (a *DataExchangeActivities) GetRevision(input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
	return a.client.GetRevision(input)
}

func (a *DataExchangeActivities) ListDataSetRevisions(input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
	return a.client.ListDataSetRevisions(input)
}

func (a *DataExchangeActivities) ListDataSets(input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
	return a.client.ListDataSets(input)
}

func (a *DataExchangeActivities) ListJobs(input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
	return a.client.ListJobs(input)
}

func (a *DataExchangeActivities) ListRevisionAssets(input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
	return a.client.ListRevisionAssets(input)
}

func (a *DataExchangeActivities) ListTagsForResource(input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *DataExchangeActivities) StartJob(input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
	return a.client.StartJob(input)
}

func (a *DataExchangeActivities) TagResource(input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *DataExchangeActivities) UntagResource(input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *DataExchangeActivities) UpdateAsset(input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
	return a.client.UpdateAsset(input)
}

func (a *DataExchangeActivities) UpdateDataSet(input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
	return a.client.UpdateDataSet(input)
}

func (a *DataExchangeActivities) UpdateRevision(input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
	return a.client.UpdateRevision(input)
}
