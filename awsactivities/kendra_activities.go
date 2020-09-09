
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kendra"
	"github.com/aws/aws-sdk-go/service/kendra/kendraiface"
)

type KendraActivities struct {
    client kendraiface.KendraAPI
}

func NewKendraActivities(session *session.Session, config ...*aws.Config) *KendraActivities {
    client := kendra.New(session, config...)
    return &KendraActivities{client: client}
}

func (a *KendraActivities) BatchDeleteDocument(input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error) {
    return a.client.BatchDeleteDocument(input)
}

func (a *KendraActivities) BatchPutDocument(input *kendra.BatchPutDocumentInput) (*kendra.BatchPutDocumentOutput, error) {
    return a.client.BatchPutDocument(input)
}

func (a *KendraActivities) CreateDataSource(input *kendra.CreateDataSourceInput) (*kendra.CreateDataSourceOutput, error) {
    return a.client.CreateDataSource(input)
}

func (a *KendraActivities) CreateFaq(input *kendra.CreateFaqInput) (*kendra.CreateFaqOutput, error) {
    return a.client.CreateFaq(input)
}

func (a *KendraActivities) CreateIndex(input *kendra.CreateIndexInput) (*kendra.CreateIndexOutput, error) {
    return a.client.CreateIndex(input)
}

func (a *KendraActivities) DeleteDataSource(input *kendra.DeleteDataSourceInput) (*kendra.DeleteDataSourceOutput, error) {
    return a.client.DeleteDataSource(input)
}

func (a *KendraActivities) DeleteFaq(input *kendra.DeleteFaqInput) (*kendra.DeleteFaqOutput, error) {
    return a.client.DeleteFaq(input)
}

func (a *KendraActivities) DeleteIndex(input *kendra.DeleteIndexInput) (*kendra.DeleteIndexOutput, error) {
    return a.client.DeleteIndex(input)
}

func (a *KendraActivities) DescribeDataSource(input *kendra.DescribeDataSourceInput) (*kendra.DescribeDataSourceOutput, error) {
    return a.client.DescribeDataSource(input)
}

func (a *KendraActivities) DescribeFaq(input *kendra.DescribeFaqInput) (*kendra.DescribeFaqOutput, error) {
    return a.client.DescribeFaq(input)
}

func (a *KendraActivities) DescribeIndex(input *kendra.DescribeIndexInput) (*kendra.DescribeIndexOutput, error) {
    return a.client.DescribeIndex(input)
}

func (a *KendraActivities) ListDataSourceSyncJobs(input *kendra.ListDataSourceSyncJobsInput) (*kendra.ListDataSourceSyncJobsOutput, error) {
    return a.client.ListDataSourceSyncJobs(input)
}

func (a *KendraActivities) ListDataSources(input *kendra.ListDataSourcesInput) (*kendra.ListDataSourcesOutput, error) {
    return a.client.ListDataSources(input)
}

func (a *KendraActivities) ListFaqs(input *kendra.ListFaqsInput) (*kendra.ListFaqsOutput, error) {
    return a.client.ListFaqs(input)
}

func (a *KendraActivities) ListIndices(input *kendra.ListIndicesInput) (*kendra.ListIndicesOutput, error) {
    return a.client.ListIndices(input)
}

func (a *KendraActivities) ListTagsForResource(input *kendra.ListTagsForResourceInput) (*kendra.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *KendraActivities) Query(input *kendra.QueryInput) (*kendra.QueryOutput, error) {
    return a.client.Query(input)
}

func (a *KendraActivities) StartDataSourceSyncJob(input *kendra.StartDataSourceSyncJobInput) (*kendra.StartDataSourceSyncJobOutput, error) {
    return a.client.StartDataSourceSyncJob(input)
}

func (a *KendraActivities) StopDataSourceSyncJob(input *kendra.StopDataSourceSyncJobInput) (*kendra.StopDataSourceSyncJobOutput, error) {
    return a.client.StopDataSourceSyncJob(input)
}

func (a *KendraActivities) SubmitFeedback(input *kendra.SubmitFeedbackInput) (*kendra.SubmitFeedbackOutput, error) {
    return a.client.SubmitFeedback(input)
}

func (a *KendraActivities) TagResource(input *kendra.TagResourceInput) (*kendra.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *KendraActivities) UntagResource(input *kendra.UntagResourceInput) (*kendra.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *KendraActivities) UpdateDataSource(input *kendra.UpdateDataSourceInput) (*kendra.UpdateDataSourceOutput, error) {
    return a.client.UpdateDataSource(input)
}

func (a *KendraActivities) UpdateIndex(input *kendra.UpdateIndexInput) (*kendra.UpdateIndexOutput, error) {
    return a.client.UpdateIndex(input)
}
