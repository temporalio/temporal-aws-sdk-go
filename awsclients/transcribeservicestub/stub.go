// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package transcribeservicestub

import (
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type TranscribeServiceCreateLanguageModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceCreateLanguageModelFuture) Get(ctx workflow.Context) (*transcribeservice.CreateLanguageModelOutput, error) {
	var output transcribeservice.CreateLanguageModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceCreateMedicalVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceCreateMedicalVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
	var output transcribeservice.CreateMedicalVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceCreateVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceCreateVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.CreateVocabularyOutput, error) {
	var output transcribeservice.CreateVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceCreateVocabularyFilterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceCreateVocabularyFilterFuture) Get(ctx workflow.Context) (*transcribeservice.CreateVocabularyFilterOutput, error) {
	var output transcribeservice.CreateVocabularyFilterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDeleteLanguageModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDeleteLanguageModelFuture) Get(ctx workflow.Context) (*transcribeservice.DeleteLanguageModelOutput, error) {
	var output transcribeservice.DeleteLanguageModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDeleteMedicalTranscriptionJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDeleteMedicalTranscriptionJobFuture) Get(ctx workflow.Context) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
	var output transcribeservice.DeleteMedicalTranscriptionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDeleteMedicalVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDeleteMedicalVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
	var output transcribeservice.DeleteMedicalVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDeleteTranscriptionJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDeleteTranscriptionJobFuture) Get(ctx workflow.Context) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
	var output transcribeservice.DeleteTranscriptionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDeleteVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDeleteVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.DeleteVocabularyOutput, error) {
	var output transcribeservice.DeleteVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDeleteVocabularyFilterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDeleteVocabularyFilterFuture) Get(ctx workflow.Context) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
	var output transcribeservice.DeleteVocabularyFilterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceDescribeLanguageModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceDescribeLanguageModelFuture) Get(ctx workflow.Context) (*transcribeservice.DescribeLanguageModelOutput, error) {
	var output transcribeservice.DescribeLanguageModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceGetMedicalTranscriptionJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceGetMedicalTranscriptionJobFuture) Get(ctx workflow.Context) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
	var output transcribeservice.GetMedicalTranscriptionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceGetMedicalVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceGetMedicalVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.GetMedicalVocabularyOutput, error) {
	var output transcribeservice.GetMedicalVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceGetTranscriptionJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceGetTranscriptionJobFuture) Get(ctx workflow.Context) (*transcribeservice.GetTranscriptionJobOutput, error) {
	var output transcribeservice.GetTranscriptionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceGetVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceGetVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.GetVocabularyOutput, error) {
	var output transcribeservice.GetVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceGetVocabularyFilterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceGetVocabularyFilterFuture) Get(ctx workflow.Context) (*transcribeservice.GetVocabularyFilterOutput, error) {
	var output transcribeservice.GetVocabularyFilterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceListLanguageModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceListLanguageModelsFuture) Get(ctx workflow.Context) (*transcribeservice.ListLanguageModelsOutput, error) {
	var output transcribeservice.ListLanguageModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceListMedicalTranscriptionJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceListMedicalTranscriptionJobsFuture) Get(ctx workflow.Context) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
	var output transcribeservice.ListMedicalTranscriptionJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceListMedicalVocabulariesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceListMedicalVocabulariesFuture) Get(ctx workflow.Context) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
	var output transcribeservice.ListMedicalVocabulariesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceListTranscriptionJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceListTranscriptionJobsFuture) Get(ctx workflow.Context) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	var output transcribeservice.ListTranscriptionJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceListVocabulariesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceListVocabulariesFuture) Get(ctx workflow.Context) (*transcribeservice.ListVocabulariesOutput, error) {
	var output transcribeservice.ListVocabulariesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceListVocabularyFiltersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceListVocabularyFiltersFuture) Get(ctx workflow.Context) (*transcribeservice.ListVocabularyFiltersOutput, error) {
	var output transcribeservice.ListVocabularyFiltersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceStartMedicalTranscriptionJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceStartMedicalTranscriptionJobFuture) Get(ctx workflow.Context) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
	var output transcribeservice.StartMedicalTranscriptionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceStartTranscriptionJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceStartTranscriptionJobFuture) Get(ctx workflow.Context) (*transcribeservice.StartTranscriptionJobOutput, error) {
	var output transcribeservice.StartTranscriptionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceUpdateMedicalVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceUpdateMedicalVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
	var output transcribeservice.UpdateMedicalVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceUpdateVocabularyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceUpdateVocabularyFuture) Get(ctx workflow.Context) (*transcribeservice.UpdateVocabularyOutput, error) {
	var output transcribeservice.UpdateVocabularyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TranscribeServiceUpdateVocabularyFilterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TranscribeServiceUpdateVocabularyFilterFuture) Get(ctx workflow.Context) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
	var output transcribeservice.UpdateVocabularyFilterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLanguageModel(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error) {
	var output transcribeservice.CreateLanguageModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateLanguageModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLanguageModelAsync(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) *TranscribeServiceCreateLanguageModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateLanguageModel", input)
	return &TranscribeServiceCreateLanguageModelFuture{Future: future}
}

func (a *stub) CreateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
	var output transcribeservice.CreateMedicalVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateMedicalVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) *TranscribeServiceCreateMedicalVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateMedicalVocabulary", input)
	return &TranscribeServiceCreateMedicalVocabularyFuture{Future: future}
}

func (a *stub) CreateVocabulary(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error) {
	var output transcribeservice.CreateVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) *TranscribeServiceCreateVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateVocabulary", input)
	return &TranscribeServiceCreateVocabularyFuture{Future: future}
}

func (a *stub) CreateVocabularyFilter(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error) {
	var output transcribeservice.CreateVocabularyFilterOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateVocabularyFilter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) *TranscribeServiceCreateVocabularyFilterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.CreateVocabularyFilter", input)
	return &TranscribeServiceCreateVocabularyFilterFuture{Future: future}
}

func (a *stub) DeleteLanguageModel(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error) {
	var output transcribeservice.DeleteLanguageModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteLanguageModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) *TranscribeServiceDeleteLanguageModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteLanguageModel", input)
	return &TranscribeServiceDeleteLanguageModelFuture{Future: future}
}

func (a *stub) DeleteMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
	var output transcribeservice.DeleteMedicalTranscriptionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteMedicalTranscriptionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) *TranscribeServiceDeleteMedicalTranscriptionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteMedicalTranscriptionJob", input)
	return &TranscribeServiceDeleteMedicalTranscriptionJobFuture{Future: future}
}

func (a *stub) DeleteMedicalVocabulary(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
	var output transcribeservice.DeleteMedicalVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteMedicalVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) *TranscribeServiceDeleteMedicalVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteMedicalVocabulary", input)
	return &TranscribeServiceDeleteMedicalVocabularyFuture{Future: future}
}

func (a *stub) DeleteTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
	var output transcribeservice.DeleteTranscriptionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteTranscriptionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) *TranscribeServiceDeleteTranscriptionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteTranscriptionJob", input)
	return &TranscribeServiceDeleteTranscriptionJobFuture{Future: future}
}

func (a *stub) DeleteVocabulary(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error) {
	var output transcribeservice.DeleteVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) *TranscribeServiceDeleteVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteVocabulary", input)
	return &TranscribeServiceDeleteVocabularyFuture{Future: future}
}

func (a *stub) DeleteVocabularyFilter(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
	var output transcribeservice.DeleteVocabularyFilterOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteVocabularyFilter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) *TranscribeServiceDeleteVocabularyFilterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DeleteVocabularyFilter", input)
	return &TranscribeServiceDeleteVocabularyFilterFuture{Future: future}
}

func (a *stub) DescribeLanguageModel(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
	var output transcribeservice.DescribeLanguageModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DescribeLanguageModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) *TranscribeServiceDescribeLanguageModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.DescribeLanguageModel", input)
	return &TranscribeServiceDescribeLanguageModelFuture{Future: future}
}

func (a *stub) GetMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
	var output transcribeservice.GetMedicalTranscriptionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetMedicalTranscriptionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) *TranscribeServiceGetMedicalTranscriptionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetMedicalTranscriptionJob", input)
	return &TranscribeServiceGetMedicalTranscriptionJobFuture{Future: future}
}

func (a *stub) GetMedicalVocabulary(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
	var output transcribeservice.GetMedicalVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetMedicalVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) *TranscribeServiceGetMedicalVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetMedicalVocabulary", input)
	return &TranscribeServiceGetMedicalVocabularyFuture{Future: future}
}

func (a *stub) GetTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
	var output transcribeservice.GetTranscriptionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetTranscriptionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) *TranscribeServiceGetTranscriptionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetTranscriptionJob", input)
	return &TranscribeServiceGetTranscriptionJobFuture{Future: future}
}

func (a *stub) GetVocabulary(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
	var output transcribeservice.GetVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) *TranscribeServiceGetVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetVocabulary", input)
	return &TranscribeServiceGetVocabularyFuture{Future: future}
}

func (a *stub) GetVocabularyFilter(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
	var output transcribeservice.GetVocabularyFilterOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetVocabularyFilter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) *TranscribeServiceGetVocabularyFilterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.GetVocabularyFilter", input)
	return &TranscribeServiceGetVocabularyFilterFuture{Future: future}
}

func (a *stub) ListLanguageModels(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
	var output transcribeservice.ListLanguageModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListLanguageModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLanguageModelsAsync(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) *TranscribeServiceListLanguageModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListLanguageModels", input)
	return &TranscribeServiceListLanguageModelsFuture{Future: future}
}

func (a *stub) ListMedicalTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
	var output transcribeservice.ListMedicalTranscriptionJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListMedicalTranscriptionJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMedicalTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) *TranscribeServiceListMedicalTranscriptionJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListMedicalTranscriptionJobs", input)
	return &TranscribeServiceListMedicalTranscriptionJobsFuture{Future: future}
}

func (a *stub) ListMedicalVocabularies(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
	var output transcribeservice.ListMedicalVocabulariesOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListMedicalVocabularies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMedicalVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) *TranscribeServiceListMedicalVocabulariesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListMedicalVocabularies", input)
	return &TranscribeServiceListMedicalVocabulariesFuture{Future: future}
}

func (a *stub) ListTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	var output transcribeservice.ListTranscriptionJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListTranscriptionJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) *TranscribeServiceListTranscriptionJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListTranscriptionJobs", input)
	return &TranscribeServiceListTranscriptionJobsFuture{Future: future}
}

func (a *stub) ListVocabularies(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
	var output transcribeservice.ListVocabulariesOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListVocabularies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) *TranscribeServiceListVocabulariesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListVocabularies", input)
	return &TranscribeServiceListVocabulariesFuture{Future: future}
}

func (a *stub) ListVocabularyFilters(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
	var output transcribeservice.ListVocabularyFiltersOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListVocabularyFilters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListVocabularyFiltersAsync(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) *TranscribeServiceListVocabularyFiltersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.ListVocabularyFilters", input)
	return &TranscribeServiceListVocabularyFiltersFuture{Future: future}
}

func (a *stub) StartMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
	var output transcribeservice.StartMedicalTranscriptionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.StartMedicalTranscriptionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) *TranscribeServiceStartMedicalTranscriptionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.StartMedicalTranscriptionJob", input)
	return &TranscribeServiceStartMedicalTranscriptionJobFuture{Future: future}
}

func (a *stub) StartTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error) {
	var output transcribeservice.StartTranscriptionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.StartTranscriptionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) *TranscribeServiceStartTranscriptionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.StartTranscriptionJob", input)
	return &TranscribeServiceStartTranscriptionJobFuture{Future: future}
}

func (a *stub) UpdateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
	var output transcribeservice.UpdateMedicalVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.UpdateMedicalVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) *TranscribeServiceUpdateMedicalVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.UpdateMedicalVocabulary", input)
	return &TranscribeServiceUpdateMedicalVocabularyFuture{Future: future}
}

func (a *stub) UpdateVocabulary(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error) {
	var output transcribeservice.UpdateVocabularyOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.UpdateVocabulary", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) *TranscribeServiceUpdateVocabularyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.UpdateVocabulary", input)
	return &TranscribeServiceUpdateVocabularyFuture{Future: future}
}

func (a *stub) UpdateVocabularyFilter(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
	var output transcribeservice.UpdateVocabularyFilterOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribeservice.UpdateVocabularyFilter", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) *TranscribeServiceUpdateVocabularyFilterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribeservice.UpdateVocabularyFilter", input)
	return &TranscribeServiceUpdateVocabularyFilterFuture{Future: future}
}
