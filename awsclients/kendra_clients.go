package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kendra"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KendraClient interface {
    BatchDeleteDocument(ctx workflow.Context, input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error)
    BatchDeleteDocumentAsync(ctx workflow.Context, input *kendra.BatchDeleteDocumentInput) *KendraBatchDeleteDocumentResult

    BatchPutDocument(ctx workflow.Context, input *kendra.BatchPutDocumentInput) (*kendra.BatchPutDocumentOutput, error)
    BatchPutDocumentAsync(ctx workflow.Context, input *kendra.BatchPutDocumentInput) *KendraBatchPutDocumentResult

    CreateDataSource(ctx workflow.Context, input *kendra.CreateDataSourceInput) (*kendra.CreateDataSourceOutput, error)
    CreateDataSourceAsync(ctx workflow.Context, input *kendra.CreateDataSourceInput) *KendraCreateDataSourceResult

    CreateFaq(ctx workflow.Context, input *kendra.CreateFaqInput) (*kendra.CreateFaqOutput, error)
    CreateFaqAsync(ctx workflow.Context, input *kendra.CreateFaqInput) *KendraCreateFaqResult

    CreateIndex(ctx workflow.Context, input *kendra.CreateIndexInput) (*kendra.CreateIndexOutput, error)
    CreateIndexAsync(ctx workflow.Context, input *kendra.CreateIndexInput) *KendraCreateIndexResult

    DeleteDataSource(ctx workflow.Context, input *kendra.DeleteDataSourceInput) (*kendra.DeleteDataSourceOutput, error)
    DeleteDataSourceAsync(ctx workflow.Context, input *kendra.DeleteDataSourceInput) *KendraDeleteDataSourceResult

    DeleteFaq(ctx workflow.Context, input *kendra.DeleteFaqInput) (*kendra.DeleteFaqOutput, error)
    DeleteFaqAsync(ctx workflow.Context, input *kendra.DeleteFaqInput) *KendraDeleteFaqResult

    DeleteIndex(ctx workflow.Context, input *kendra.DeleteIndexInput) (*kendra.DeleteIndexOutput, error)
    DeleteIndexAsync(ctx workflow.Context, input *kendra.DeleteIndexInput) *KendraDeleteIndexResult

    DescribeDataSource(ctx workflow.Context, input *kendra.DescribeDataSourceInput) (*kendra.DescribeDataSourceOutput, error)
    DescribeDataSourceAsync(ctx workflow.Context, input *kendra.DescribeDataSourceInput) *KendraDescribeDataSourceResult

    DescribeFaq(ctx workflow.Context, input *kendra.DescribeFaqInput) (*kendra.DescribeFaqOutput, error)
    DescribeFaqAsync(ctx workflow.Context, input *kendra.DescribeFaqInput) *KendraDescribeFaqResult

    DescribeIndex(ctx workflow.Context, input *kendra.DescribeIndexInput) (*kendra.DescribeIndexOutput, error)
    DescribeIndexAsync(ctx workflow.Context, input *kendra.DescribeIndexInput) *KendraDescribeIndexResult

    ListDataSourceSyncJobs(ctx workflow.Context, input *kendra.ListDataSourceSyncJobsInput) (*kendra.ListDataSourceSyncJobsOutput, error)
    ListDataSourceSyncJobsAsync(ctx workflow.Context, input *kendra.ListDataSourceSyncJobsInput) *KendraListDataSourceSyncJobsResult

    ListDataSources(ctx workflow.Context, input *kendra.ListDataSourcesInput) (*kendra.ListDataSourcesOutput, error)
    ListDataSourcesAsync(ctx workflow.Context, input *kendra.ListDataSourcesInput) *KendraListDataSourcesResult

    ListFaqs(ctx workflow.Context, input *kendra.ListFaqsInput) (*kendra.ListFaqsOutput, error)
    ListFaqsAsync(ctx workflow.Context, input *kendra.ListFaqsInput) *KendraListFaqsResult

    ListIndices(ctx workflow.Context, input *kendra.ListIndicesInput) (*kendra.ListIndicesOutput, error)
    ListIndicesAsync(ctx workflow.Context, input *kendra.ListIndicesInput) *KendraListIndicesResult

    ListTagsForResource(ctx workflow.Context, input *kendra.ListTagsForResourceInput) (*kendra.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *kendra.ListTagsForResourceInput) *KendraListTagsForResourceResult

    Query(ctx workflow.Context, input *kendra.QueryInput) (*kendra.QueryOutput, error)
    QueryAsync(ctx workflow.Context, input *kendra.QueryInput) *KendraQueryResult

    StartDataSourceSyncJob(ctx workflow.Context, input *kendra.StartDataSourceSyncJobInput) (*kendra.StartDataSourceSyncJobOutput, error)
    StartDataSourceSyncJobAsync(ctx workflow.Context, input *kendra.StartDataSourceSyncJobInput) *KendraStartDataSourceSyncJobResult

    StopDataSourceSyncJob(ctx workflow.Context, input *kendra.StopDataSourceSyncJobInput) (*kendra.StopDataSourceSyncJobOutput, error)
    StopDataSourceSyncJobAsync(ctx workflow.Context, input *kendra.StopDataSourceSyncJobInput) *KendraStopDataSourceSyncJobResult

    SubmitFeedback(ctx workflow.Context, input *kendra.SubmitFeedbackInput) (*kendra.SubmitFeedbackOutput, error)
    SubmitFeedbackAsync(ctx workflow.Context, input *kendra.SubmitFeedbackInput) *KendraSubmitFeedbackResult

    TagResource(ctx workflow.Context, input *kendra.TagResourceInput) (*kendra.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *kendra.TagResourceInput) *KendraTagResourceResult

    UntagResource(ctx workflow.Context, input *kendra.UntagResourceInput) (*kendra.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *kendra.UntagResourceInput) *KendraUntagResourceResult

    UpdateDataSource(ctx workflow.Context, input *kendra.UpdateDataSourceInput) (*kendra.UpdateDataSourceOutput, error)
    UpdateDataSourceAsync(ctx workflow.Context, input *kendra.UpdateDataSourceInput) *KendraUpdateDataSourceResult

    UpdateIndex(ctx workflow.Context, input *kendra.UpdateIndexInput) (*kendra.UpdateIndexOutput, error)
    UpdateIndexAsync(ctx workflow.Context, input *kendra.UpdateIndexInput) *KendraUpdateIndexResult
}

type KendraBatchDeleteDocumentResult struct {
	Result workflow.Future
}

func (r *KendraBatchDeleteDocumentResult) Get(ctx workflow.Context) (*kendra.BatchDeleteDocumentOutput, error) {
    var output kendra.BatchDeleteDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraBatchPutDocumentResult struct {
	Result workflow.Future
}

func (r *KendraBatchPutDocumentResult) Get(ctx workflow.Context) (*kendra.BatchPutDocumentOutput, error) {
    var output kendra.BatchPutDocumentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraCreateDataSourceResult struct {
	Result workflow.Future
}

func (r *KendraCreateDataSourceResult) Get(ctx workflow.Context) (*kendra.CreateDataSourceOutput, error) {
    var output kendra.CreateDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraCreateFaqResult struct {
	Result workflow.Future
}

func (r *KendraCreateFaqResult) Get(ctx workflow.Context) (*kendra.CreateFaqOutput, error) {
    var output kendra.CreateFaqOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraCreateIndexResult struct {
	Result workflow.Future
}

func (r *KendraCreateIndexResult) Get(ctx workflow.Context) (*kendra.CreateIndexOutput, error) {
    var output kendra.CreateIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraDeleteDataSourceResult struct {
	Result workflow.Future
}

func (r *KendraDeleteDataSourceResult) Get(ctx workflow.Context) (*kendra.DeleteDataSourceOutput, error) {
    var output kendra.DeleteDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraDeleteFaqResult struct {
	Result workflow.Future
}

func (r *KendraDeleteFaqResult) Get(ctx workflow.Context) (*kendra.DeleteFaqOutput, error) {
    var output kendra.DeleteFaqOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraDeleteIndexResult struct {
	Result workflow.Future
}

func (r *KendraDeleteIndexResult) Get(ctx workflow.Context) (*kendra.DeleteIndexOutput, error) {
    var output kendra.DeleteIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraDescribeDataSourceResult struct {
	Result workflow.Future
}

func (r *KendraDescribeDataSourceResult) Get(ctx workflow.Context) (*kendra.DescribeDataSourceOutput, error) {
    var output kendra.DescribeDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraDescribeFaqResult struct {
	Result workflow.Future
}

func (r *KendraDescribeFaqResult) Get(ctx workflow.Context) (*kendra.DescribeFaqOutput, error) {
    var output kendra.DescribeFaqOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraDescribeIndexResult struct {
	Result workflow.Future
}

func (r *KendraDescribeIndexResult) Get(ctx workflow.Context) (*kendra.DescribeIndexOutput, error) {
    var output kendra.DescribeIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraListDataSourceSyncJobsResult struct {
	Result workflow.Future
}

func (r *KendraListDataSourceSyncJobsResult) Get(ctx workflow.Context) (*kendra.ListDataSourceSyncJobsOutput, error) {
    var output kendra.ListDataSourceSyncJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraListDataSourcesResult struct {
	Result workflow.Future
}

func (r *KendraListDataSourcesResult) Get(ctx workflow.Context) (*kendra.ListDataSourcesOutput, error) {
    var output kendra.ListDataSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraListFaqsResult struct {
	Result workflow.Future
}

func (r *KendraListFaqsResult) Get(ctx workflow.Context) (*kendra.ListFaqsOutput, error) {
    var output kendra.ListFaqsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraListIndicesResult struct {
	Result workflow.Future
}

func (r *KendraListIndicesResult) Get(ctx workflow.Context) (*kendra.ListIndicesOutput, error) {
    var output kendra.ListIndicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *KendraListTagsForResourceResult) Get(ctx workflow.Context) (*kendra.ListTagsForResourceOutput, error) {
    var output kendra.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraQueryResult struct {
	Result workflow.Future
}

func (r *KendraQueryResult) Get(ctx workflow.Context) (*kendra.QueryOutput, error) {
    var output kendra.QueryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraStartDataSourceSyncJobResult struct {
	Result workflow.Future
}

func (r *KendraStartDataSourceSyncJobResult) Get(ctx workflow.Context) (*kendra.StartDataSourceSyncJobOutput, error) {
    var output kendra.StartDataSourceSyncJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraStopDataSourceSyncJobResult struct {
	Result workflow.Future
}

func (r *KendraStopDataSourceSyncJobResult) Get(ctx workflow.Context) (*kendra.StopDataSourceSyncJobOutput, error) {
    var output kendra.StopDataSourceSyncJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraSubmitFeedbackResult struct {
	Result workflow.Future
}

func (r *KendraSubmitFeedbackResult) Get(ctx workflow.Context) (*kendra.SubmitFeedbackOutput, error) {
    var output kendra.SubmitFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraTagResourceResult struct {
	Result workflow.Future
}

func (r *KendraTagResourceResult) Get(ctx workflow.Context) (*kendra.TagResourceOutput, error) {
    var output kendra.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraUntagResourceResult struct {
	Result workflow.Future
}

func (r *KendraUntagResourceResult) Get(ctx workflow.Context) (*kendra.UntagResourceOutput, error) {
    var output kendra.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraUpdateDataSourceResult struct {
	Result workflow.Future
}

func (r *KendraUpdateDataSourceResult) Get(ctx workflow.Context) (*kendra.UpdateDataSourceOutput, error) {
    var output kendra.UpdateDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraUpdateIndexResult struct {
	Result workflow.Future
}

func (r *KendraUpdateIndexResult) Get(ctx workflow.Context) (*kendra.UpdateIndexOutput, error) {
    var output kendra.UpdateIndexOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type KendraStub struct {
    activities awsactivities.KendraActivities
}

func NewKendraStub() KendraClient {
    return &KendraStub{}
}

func (a *KendraStub) BatchDeleteDocument(ctx workflow.Context, input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error) {
    var output kendra.BatchDeleteDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) BatchDeleteDocumentAsync(ctx workflow.Context, input *kendra.BatchDeleteDocumentInput) *KendraBatchDeleteDocumentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteDocument, input)
    return &KendraBatchDeleteDocumentResult{Result: future}
}

func (a *KendraStub) BatchPutDocument(ctx workflow.Context, input *kendra.BatchPutDocumentInput) (*kendra.BatchPutDocumentOutput, error) {
    var output kendra.BatchPutDocumentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchPutDocument, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) BatchPutDocumentAsync(ctx workflow.Context, input *kendra.BatchPutDocumentInput) *KendraBatchPutDocumentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchPutDocument, input)
    return &KendraBatchPutDocumentResult{Result: future}
}

func (a *KendraStub) CreateDataSource(ctx workflow.Context, input *kendra.CreateDataSourceInput) (*kendra.CreateDataSourceOutput, error) {
    var output kendra.CreateDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) CreateDataSourceAsync(ctx workflow.Context, input *kendra.CreateDataSourceInput) *KendraCreateDataSourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataSource, input)
    return &KendraCreateDataSourceResult{Result: future}
}

func (a *KendraStub) CreateFaq(ctx workflow.Context, input *kendra.CreateFaqInput) (*kendra.CreateFaqOutput, error) {
    var output kendra.CreateFaqOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFaq, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) CreateFaqAsync(ctx workflow.Context, input *kendra.CreateFaqInput) *KendraCreateFaqResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFaq, input)
    return &KendraCreateFaqResult{Result: future}
}

func (a *KendraStub) CreateIndex(ctx workflow.Context, input *kendra.CreateIndexInput) (*kendra.CreateIndexOutput, error) {
    var output kendra.CreateIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) CreateIndexAsync(ctx workflow.Context, input *kendra.CreateIndexInput) *KendraCreateIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateIndex, input)
    return &KendraCreateIndexResult{Result: future}
}

func (a *KendraStub) DeleteDataSource(ctx workflow.Context, input *kendra.DeleteDataSourceInput) (*kendra.DeleteDataSourceOutput, error) {
    var output kendra.DeleteDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) DeleteDataSourceAsync(ctx workflow.Context, input *kendra.DeleteDataSourceInput) *KendraDeleteDataSourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSource, input)
    return &KendraDeleteDataSourceResult{Result: future}
}

func (a *KendraStub) DeleteFaq(ctx workflow.Context, input *kendra.DeleteFaqInput) (*kendra.DeleteFaqOutput, error) {
    var output kendra.DeleteFaqOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFaq, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) DeleteFaqAsync(ctx workflow.Context, input *kendra.DeleteFaqInput) *KendraDeleteFaqResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFaq, input)
    return &KendraDeleteFaqResult{Result: future}
}

func (a *KendraStub) DeleteIndex(ctx workflow.Context, input *kendra.DeleteIndexInput) (*kendra.DeleteIndexOutput, error) {
    var output kendra.DeleteIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) DeleteIndexAsync(ctx workflow.Context, input *kendra.DeleteIndexInput) *KendraDeleteIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteIndex, input)
    return &KendraDeleteIndexResult{Result: future}
}

func (a *KendraStub) DescribeDataSource(ctx workflow.Context, input *kendra.DescribeDataSourceInput) (*kendra.DescribeDataSourceOutput, error) {
    var output kendra.DescribeDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) DescribeDataSourceAsync(ctx workflow.Context, input *kendra.DescribeDataSourceInput) *KendraDescribeDataSourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSource, input)
    return &KendraDescribeDataSourceResult{Result: future}
}

func (a *KendraStub) DescribeFaq(ctx workflow.Context, input *kendra.DescribeFaqInput) (*kendra.DescribeFaqOutput, error) {
    var output kendra.DescribeFaqOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFaq, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) DescribeFaqAsync(ctx workflow.Context, input *kendra.DescribeFaqInput) *KendraDescribeFaqResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFaq, input)
    return &KendraDescribeFaqResult{Result: future}
}

func (a *KendraStub) DescribeIndex(ctx workflow.Context, input *kendra.DescribeIndexInput) (*kendra.DescribeIndexOutput, error) {
    var output kendra.DescribeIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) DescribeIndexAsync(ctx workflow.Context, input *kendra.DescribeIndexInput) *KendraDescribeIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeIndex, input)
    return &KendraDescribeIndexResult{Result: future}
}

func (a *KendraStub) ListDataSourceSyncJobs(ctx workflow.Context, input *kendra.ListDataSourceSyncJobsInput) (*kendra.ListDataSourceSyncJobsOutput, error) {
    var output kendra.ListDataSourceSyncJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataSourceSyncJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) ListDataSourceSyncJobsAsync(ctx workflow.Context, input *kendra.ListDataSourceSyncJobsInput) *KendraListDataSourceSyncJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDataSourceSyncJobs, input)
    return &KendraListDataSourceSyncJobsResult{Result: future}
}

func (a *KendraStub) ListDataSources(ctx workflow.Context, input *kendra.ListDataSourcesInput) (*kendra.ListDataSourcesOutput, error) {
    var output kendra.ListDataSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataSources, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) ListDataSourcesAsync(ctx workflow.Context, input *kendra.ListDataSourcesInput) *KendraListDataSourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDataSources, input)
    return &KendraListDataSourcesResult{Result: future}
}

func (a *KendraStub) ListFaqs(ctx workflow.Context, input *kendra.ListFaqsInput) (*kendra.ListFaqsOutput, error) {
    var output kendra.ListFaqsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFaqs, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) ListFaqsAsync(ctx workflow.Context, input *kendra.ListFaqsInput) *KendraListFaqsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFaqs, input)
    return &KendraListFaqsResult{Result: future}
}

func (a *KendraStub) ListIndices(ctx workflow.Context, input *kendra.ListIndicesInput) (*kendra.ListIndicesOutput, error) {
    var output kendra.ListIndicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIndices, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) ListIndicesAsync(ctx workflow.Context, input *kendra.ListIndicesInput) *KendraListIndicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListIndices, input)
    return &KendraListIndicesResult{Result: future}
}

func (a *KendraStub) ListTagsForResource(ctx workflow.Context, input *kendra.ListTagsForResourceInput) (*kendra.ListTagsForResourceOutput, error) {
    var output kendra.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) ListTagsForResourceAsync(ctx workflow.Context, input *kendra.ListTagsForResourceInput) *KendraListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &KendraListTagsForResourceResult{Result: future}
}

func (a *KendraStub) Query(ctx workflow.Context, input *kendra.QueryInput) (*kendra.QueryOutput, error) {
    var output kendra.QueryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.Query, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) QueryAsync(ctx workflow.Context, input *kendra.QueryInput) *KendraQueryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.Query, input)
    return &KendraQueryResult{Result: future}
}

func (a *KendraStub) StartDataSourceSyncJob(ctx workflow.Context, input *kendra.StartDataSourceSyncJobInput) (*kendra.StartDataSourceSyncJobOutput, error) {
    var output kendra.StartDataSourceSyncJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDataSourceSyncJob, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) StartDataSourceSyncJobAsync(ctx workflow.Context, input *kendra.StartDataSourceSyncJobInput) *KendraStartDataSourceSyncJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartDataSourceSyncJob, input)
    return &KendraStartDataSourceSyncJobResult{Result: future}
}

func (a *KendraStub) StopDataSourceSyncJob(ctx workflow.Context, input *kendra.StopDataSourceSyncJobInput) (*kendra.StopDataSourceSyncJobOutput, error) {
    var output kendra.StopDataSourceSyncJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDataSourceSyncJob, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) StopDataSourceSyncJobAsync(ctx workflow.Context, input *kendra.StopDataSourceSyncJobInput) *KendraStopDataSourceSyncJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopDataSourceSyncJob, input)
    return &KendraStopDataSourceSyncJobResult{Result: future}
}

func (a *KendraStub) SubmitFeedback(ctx workflow.Context, input *kendra.SubmitFeedbackInput) (*kendra.SubmitFeedbackOutput, error) {
    var output kendra.SubmitFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubmitFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) SubmitFeedbackAsync(ctx workflow.Context, input *kendra.SubmitFeedbackInput) *KendraSubmitFeedbackResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SubmitFeedback, input)
    return &KendraSubmitFeedbackResult{Result: future}
}

func (a *KendraStub) TagResource(ctx workflow.Context, input *kendra.TagResourceInput) (*kendra.TagResourceOutput, error) {
    var output kendra.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) TagResourceAsync(ctx workflow.Context, input *kendra.TagResourceInput) *KendraTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &KendraTagResourceResult{Result: future}
}

func (a *KendraStub) UntagResource(ctx workflow.Context, input *kendra.UntagResourceInput) (*kendra.UntagResourceOutput, error) {
    var output kendra.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) UntagResourceAsync(ctx workflow.Context, input *kendra.UntagResourceInput) *KendraUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &KendraUntagResourceResult{Result: future}
}

func (a *KendraStub) UpdateDataSource(ctx workflow.Context, input *kendra.UpdateDataSourceInput) (*kendra.UpdateDataSourceOutput, error) {
    var output kendra.UpdateDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) UpdateDataSourceAsync(ctx workflow.Context, input *kendra.UpdateDataSourceInput) *KendraUpdateDataSourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSource, input)
    return &KendraUpdateDataSourceResult{Result: future}
}

func (a *KendraStub) UpdateIndex(ctx workflow.Context, input *kendra.UpdateIndexInput) (*kendra.UpdateIndexOutput, error) {
    var output kendra.UpdateIndexOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIndex, input).Get(ctx, &output)
    return &output, err
}

func (a *KendraStub) UpdateIndexAsync(ctx workflow.Context, input *kendra.UpdateIndexInput) *KendraUpdateIndexResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateIndex, input)
    return &KendraUpdateIndexResult{Result: future}
}
