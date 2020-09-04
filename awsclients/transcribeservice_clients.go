package awsclients

import (
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type TranscribeServiceClient interface {
    CreateLanguageModel(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error)
    CreateLanguageModelAsync(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) *TranscribeserviceCreateLanguageModelResult

    CreateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error)
    CreateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) *TranscribeserviceCreateMedicalVocabularyResult

    CreateVocabulary(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error)
    CreateVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) *TranscribeserviceCreateVocabularyResult

    CreateVocabularyFilter(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error)
    CreateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) *TranscribeserviceCreateVocabularyFilterResult

    DeleteLanguageModel(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error)
    DeleteLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) *TranscribeserviceDeleteLanguageModelResult

    DeleteMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error)
    DeleteMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) *TranscribeserviceDeleteMedicalTranscriptionJobResult

    DeleteMedicalVocabulary(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error)
    DeleteMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) *TranscribeserviceDeleteMedicalVocabularyResult

    DeleteTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error)
    DeleteTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) *TranscribeserviceDeleteTranscriptionJobResult

    DeleteVocabulary(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error)
    DeleteVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) *TranscribeserviceDeleteVocabularyResult

    DeleteVocabularyFilter(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error)
    DeleteVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) *TranscribeserviceDeleteVocabularyFilterResult

    DescribeLanguageModel(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error)
    DescribeLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) *TranscribeserviceDescribeLanguageModelResult

    GetMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error)
    GetMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) *TranscribeserviceGetMedicalTranscriptionJobResult

    GetMedicalVocabulary(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error)
    GetMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) *TranscribeserviceGetMedicalVocabularyResult

    GetTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error)
    GetTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) *TranscribeserviceGetTranscriptionJobResult

    GetVocabulary(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error)
    GetVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) *TranscribeserviceGetVocabularyResult

    GetVocabularyFilter(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error)
    GetVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) *TranscribeserviceGetVocabularyFilterResult

    ListLanguageModels(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error)
    ListLanguageModelsAsync(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) *TranscribeserviceListLanguageModelsResult

    ListMedicalTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error)
    ListMedicalTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) *TranscribeserviceListMedicalTranscriptionJobsResult

    ListMedicalVocabularies(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error)
    ListMedicalVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) *TranscribeserviceListMedicalVocabulariesResult

    ListTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error)
    ListTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) *TranscribeserviceListTranscriptionJobsResult

    ListVocabularies(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error)
    ListVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) *TranscribeserviceListVocabulariesResult

    ListVocabularyFilters(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error)
    ListVocabularyFiltersAsync(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) *TranscribeserviceListVocabularyFiltersResult

    StartMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error)
    StartMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) *TranscribeserviceStartMedicalTranscriptionJobResult

    StartTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error)
    StartTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) *TranscribeserviceStartTranscriptionJobResult

    UpdateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error)
    UpdateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) *TranscribeserviceUpdateMedicalVocabularyResult

    UpdateVocabulary(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error)
    UpdateVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) *TranscribeserviceUpdateVocabularyResult

    UpdateVocabularyFilter(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error)
    UpdateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) *TranscribeserviceUpdateVocabularyFilterResult
}
type TranscribeserviceCreateLanguageModelResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceCreateLanguageModelResult) Get(ctx workflow.Context) (*transcribeservice.CreateLanguageModelOutput, error) {
    var output transcribeservice.CreateLanguageModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceCreateMedicalVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceCreateMedicalVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
    var output transcribeservice.CreateMedicalVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceCreateVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceCreateVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.CreateVocabularyOutput, error) {
    var output transcribeservice.CreateVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceCreateVocabularyFilterResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceCreateVocabularyFilterResult) Get(ctx workflow.Context) (*transcribeservice.CreateVocabularyFilterOutput, error) {
    var output transcribeservice.CreateVocabularyFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDeleteLanguageModelResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDeleteLanguageModelResult) Get(ctx workflow.Context) (*transcribeservice.DeleteLanguageModelOutput, error) {
    var output transcribeservice.DeleteLanguageModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDeleteMedicalTranscriptionJobResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDeleteMedicalTranscriptionJobResult) Get(ctx workflow.Context) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
    var output transcribeservice.DeleteMedicalTranscriptionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDeleteMedicalVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDeleteMedicalVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
    var output transcribeservice.DeleteMedicalVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDeleteTranscriptionJobResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDeleteTranscriptionJobResult) Get(ctx workflow.Context) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
    var output transcribeservice.DeleteTranscriptionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDeleteVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDeleteVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.DeleteVocabularyOutput, error) {
    var output transcribeservice.DeleteVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDeleteVocabularyFilterResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDeleteVocabularyFilterResult) Get(ctx workflow.Context) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
    var output transcribeservice.DeleteVocabularyFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceDescribeLanguageModelResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceDescribeLanguageModelResult) Get(ctx workflow.Context) (*transcribeservice.DescribeLanguageModelOutput, error) {
    var output transcribeservice.DescribeLanguageModelOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceGetMedicalTranscriptionJobResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceGetMedicalTranscriptionJobResult) Get(ctx workflow.Context) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
    var output transcribeservice.GetMedicalTranscriptionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceGetMedicalVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceGetMedicalVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.GetMedicalVocabularyOutput, error) {
    var output transcribeservice.GetMedicalVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceGetTranscriptionJobResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceGetTranscriptionJobResult) Get(ctx workflow.Context) (*transcribeservice.GetTranscriptionJobOutput, error) {
    var output transcribeservice.GetTranscriptionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceGetVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceGetVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.GetVocabularyOutput, error) {
    var output transcribeservice.GetVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceGetVocabularyFilterResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceGetVocabularyFilterResult) Get(ctx workflow.Context) (*transcribeservice.GetVocabularyFilterOutput, error) {
    var output transcribeservice.GetVocabularyFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceListLanguageModelsResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceListLanguageModelsResult) Get(ctx workflow.Context) (*transcribeservice.ListLanguageModelsOutput, error) {
    var output transcribeservice.ListLanguageModelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceListMedicalTranscriptionJobsResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceListMedicalTranscriptionJobsResult) Get(ctx workflow.Context) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
    var output transcribeservice.ListMedicalTranscriptionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceListMedicalVocabulariesResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceListMedicalVocabulariesResult) Get(ctx workflow.Context) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
    var output transcribeservice.ListMedicalVocabulariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceListTranscriptionJobsResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceListTranscriptionJobsResult) Get(ctx workflow.Context) (*transcribeservice.ListTranscriptionJobsOutput, error) {
    var output transcribeservice.ListTranscriptionJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceListVocabulariesResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceListVocabulariesResult) Get(ctx workflow.Context) (*transcribeservice.ListVocabulariesOutput, error) {
    var output transcribeservice.ListVocabulariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceListVocabularyFiltersResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceListVocabularyFiltersResult) Get(ctx workflow.Context) (*transcribeservice.ListVocabularyFiltersOutput, error) {
    var output transcribeservice.ListVocabularyFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceStartMedicalTranscriptionJobResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceStartMedicalTranscriptionJobResult) Get(ctx workflow.Context) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
    var output transcribeservice.StartMedicalTranscriptionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceStartTranscriptionJobResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceStartTranscriptionJobResult) Get(ctx workflow.Context) (*transcribeservice.StartTranscriptionJobOutput, error) {
    var output transcribeservice.StartTranscriptionJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceUpdateMedicalVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceUpdateMedicalVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
    var output transcribeservice.UpdateMedicalVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceUpdateVocabularyResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceUpdateVocabularyResult) Get(ctx workflow.Context) (*transcribeservice.UpdateVocabularyOutput, error) {
    var output transcribeservice.UpdateVocabularyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type TranscribeserviceUpdateVocabularyFilterResult struct {
	Result workflow.Future
}

func (r *TranscribeserviceUpdateVocabularyFilterResult) Get(ctx workflow.Context) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
    var output transcribeservice.UpdateVocabularyFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type TranscribeServiceStub struct {
    activities awsactivities.TranscribeServiceActivities
}

func NewTranscribeServiceStub() TranscribeServiceClient {
    return &TranscribeServiceStub{}
}
func (a *TranscribeServiceStub) CreateLanguageModel(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error) {
    var output transcribeservice.CreateLanguageModelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLanguageModel, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) CreateLanguageModelAsync(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) *TranscribeserviceCreateLanguageModelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLanguageModel, input)
    return &TranscribeserviceCreateLanguageModelResult{Result: future}
}
func (a *TranscribeServiceStub) CreateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
    var output transcribeservice.CreateMedicalVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMedicalVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) CreateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) *TranscribeserviceCreateMedicalVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMedicalVocabulary, input)
    return &TranscribeserviceCreateMedicalVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) CreateVocabulary(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error) {
    var output transcribeservice.CreateVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) CreateVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) *TranscribeserviceCreateVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVocabulary, input)
    return &TranscribeserviceCreateVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) CreateVocabularyFilter(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error) {
    var output transcribeservice.CreateVocabularyFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVocabularyFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) CreateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) *TranscribeserviceCreateVocabularyFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVocabularyFilter, input)
    return &TranscribeserviceCreateVocabularyFilterResult{Result: future}
}
func (a *TranscribeServiceStub) DeleteLanguageModel(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error) {
    var output transcribeservice.DeleteLanguageModelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLanguageModel, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DeleteLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) *TranscribeserviceDeleteLanguageModelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLanguageModel, input)
    return &TranscribeserviceDeleteLanguageModelResult{Result: future}
}
func (a *TranscribeServiceStub) DeleteMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
    var output transcribeservice.DeleteMedicalTranscriptionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMedicalTranscriptionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DeleteMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) *TranscribeserviceDeleteMedicalTranscriptionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMedicalTranscriptionJob, input)
    return &TranscribeserviceDeleteMedicalTranscriptionJobResult{Result: future}
}
func (a *TranscribeServiceStub) DeleteMedicalVocabulary(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
    var output transcribeservice.DeleteMedicalVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMedicalVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DeleteMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) *TranscribeserviceDeleteMedicalVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMedicalVocabulary, input)
    return &TranscribeserviceDeleteMedicalVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) DeleteTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
    var output transcribeservice.DeleteTranscriptionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTranscriptionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DeleteTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) *TranscribeserviceDeleteTranscriptionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTranscriptionJob, input)
    return &TranscribeserviceDeleteTranscriptionJobResult{Result: future}
}
func (a *TranscribeServiceStub) DeleteVocabulary(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error) {
    var output transcribeservice.DeleteVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DeleteVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) *TranscribeserviceDeleteVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVocabulary, input)
    return &TranscribeserviceDeleteVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) DeleteVocabularyFilter(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
    var output transcribeservice.DeleteVocabularyFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVocabularyFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DeleteVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) *TranscribeserviceDeleteVocabularyFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVocabularyFilter, input)
    return &TranscribeserviceDeleteVocabularyFilterResult{Result: future}
}
func (a *TranscribeServiceStub) DescribeLanguageModel(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
    var output transcribeservice.DescribeLanguageModelOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLanguageModel, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) DescribeLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) *TranscribeserviceDescribeLanguageModelResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLanguageModel, input)
    return &TranscribeserviceDescribeLanguageModelResult{Result: future}
}
func (a *TranscribeServiceStub) GetMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
    var output transcribeservice.GetMedicalTranscriptionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMedicalTranscriptionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) GetMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) *TranscribeserviceGetMedicalTranscriptionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMedicalTranscriptionJob, input)
    return &TranscribeserviceGetMedicalTranscriptionJobResult{Result: future}
}
func (a *TranscribeServiceStub) GetMedicalVocabulary(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
    var output transcribeservice.GetMedicalVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMedicalVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) GetMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) *TranscribeserviceGetMedicalVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMedicalVocabulary, input)
    return &TranscribeserviceGetMedicalVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) GetTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
    var output transcribeservice.GetTranscriptionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTranscriptionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) GetTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) *TranscribeserviceGetTranscriptionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTranscriptionJob, input)
    return &TranscribeserviceGetTranscriptionJobResult{Result: future}
}
func (a *TranscribeServiceStub) GetVocabulary(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
    var output transcribeservice.GetVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) GetVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) *TranscribeserviceGetVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVocabulary, input)
    return &TranscribeserviceGetVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) GetVocabularyFilter(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
    var output transcribeservice.GetVocabularyFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVocabularyFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) GetVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) *TranscribeserviceGetVocabularyFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVocabularyFilter, input)
    return &TranscribeserviceGetVocabularyFilterResult{Result: future}
}
func (a *TranscribeServiceStub) ListLanguageModels(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
    var output transcribeservice.ListLanguageModelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLanguageModels, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) ListLanguageModelsAsync(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) *TranscribeserviceListLanguageModelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLanguageModels, input)
    return &TranscribeserviceListLanguageModelsResult{Result: future}
}
func (a *TranscribeServiceStub) ListMedicalTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
    var output transcribeservice.ListMedicalTranscriptionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMedicalTranscriptionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) ListMedicalTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) *TranscribeserviceListMedicalTranscriptionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMedicalTranscriptionJobs, input)
    return &TranscribeserviceListMedicalTranscriptionJobsResult{Result: future}
}
func (a *TranscribeServiceStub) ListMedicalVocabularies(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
    var output transcribeservice.ListMedicalVocabulariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMedicalVocabularies, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) ListMedicalVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) *TranscribeserviceListMedicalVocabulariesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMedicalVocabularies, input)
    return &TranscribeserviceListMedicalVocabulariesResult{Result: future}
}
func (a *TranscribeServiceStub) ListTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
    var output transcribeservice.ListTranscriptionJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTranscriptionJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) ListTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) *TranscribeserviceListTranscriptionJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTranscriptionJobs, input)
    return &TranscribeserviceListTranscriptionJobsResult{Result: future}
}
func (a *TranscribeServiceStub) ListVocabularies(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
    var output transcribeservice.ListVocabulariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVocabularies, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) ListVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) *TranscribeserviceListVocabulariesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVocabularies, input)
    return &TranscribeserviceListVocabulariesResult{Result: future}
}
func (a *TranscribeServiceStub) ListVocabularyFilters(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
    var output transcribeservice.ListVocabularyFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVocabularyFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) ListVocabularyFiltersAsync(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) *TranscribeserviceListVocabularyFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVocabularyFilters, input)
    return &TranscribeserviceListVocabularyFiltersResult{Result: future}
}
func (a *TranscribeServiceStub) StartMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
    var output transcribeservice.StartMedicalTranscriptionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMedicalTranscriptionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) StartMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) *TranscribeserviceStartMedicalTranscriptionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMedicalTranscriptionJob, input)
    return &TranscribeserviceStartMedicalTranscriptionJobResult{Result: future}
}
func (a *TranscribeServiceStub) StartTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error) {
    var output transcribeservice.StartTranscriptionJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartTranscriptionJob, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) StartTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) *TranscribeserviceStartTranscriptionJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartTranscriptionJob, input)
    return &TranscribeserviceStartTranscriptionJobResult{Result: future}
}
func (a *TranscribeServiceStub) UpdateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
    var output transcribeservice.UpdateMedicalVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMedicalVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) UpdateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) *TranscribeserviceUpdateMedicalVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMedicalVocabulary, input)
    return &TranscribeserviceUpdateMedicalVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) UpdateVocabulary(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error) {
    var output transcribeservice.UpdateVocabularyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVocabulary, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) UpdateVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) *TranscribeserviceUpdateVocabularyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVocabulary, input)
    return &TranscribeserviceUpdateVocabularyResult{Result: future}
}
func (a *TranscribeServiceStub) UpdateVocabularyFilter(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
    var output transcribeservice.UpdateVocabularyFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVocabularyFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *TranscribeServiceStub) UpdateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) *TranscribeserviceUpdateVocabularyFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVocabularyFilter, input)
    return &TranscribeserviceUpdateVocabularyFilterResult{Result: future}
}
