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

type Client interface {
	CreateLanguageModel(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error)
	CreateLanguageModelAsync(ctx workflow.Context, input *transcribeservice.CreateLanguageModelInput) *TranscribeServiceCreateLanguageModelFuture

	CreateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error)
	CreateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateMedicalVocabularyInput) *TranscribeServiceCreateMedicalVocabularyFuture

	CreateVocabulary(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error)
	CreateVocabularyAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyInput) *TranscribeServiceCreateVocabularyFuture

	CreateVocabularyFilter(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error)
	CreateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.CreateVocabularyFilterInput) *TranscribeServiceCreateVocabularyFilterFuture

	DeleteLanguageModel(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error)
	DeleteLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DeleteLanguageModelInput) *TranscribeServiceDeleteLanguageModelFuture

	DeleteMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error)
	DeleteMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) *TranscribeServiceDeleteMedicalTranscriptionJobFuture

	DeleteMedicalVocabulary(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error)
	DeleteMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteMedicalVocabularyInput) *TranscribeServiceDeleteMedicalVocabularyFuture

	DeleteTranscriptionJob(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error)
	DeleteTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.DeleteTranscriptionJobInput) *TranscribeServiceDeleteTranscriptionJobFuture

	DeleteVocabulary(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error)
	DeleteVocabularyAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyInput) *TranscribeServiceDeleteVocabularyFuture

	DeleteVocabularyFilter(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error)
	DeleteVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.DeleteVocabularyFilterInput) *TranscribeServiceDeleteVocabularyFilterFuture

	DescribeLanguageModel(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error)
	DescribeLanguageModelAsync(ctx workflow.Context, input *transcribeservice.DescribeLanguageModelInput) *TranscribeServiceDescribeLanguageModelFuture

	GetMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error)
	GetMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) *TranscribeServiceGetMedicalTranscriptionJobFuture

	GetMedicalVocabulary(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error)
	GetMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetMedicalVocabularyInput) *TranscribeServiceGetMedicalVocabularyFuture

	GetTranscriptionJob(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error)
	GetTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.GetTranscriptionJobInput) *TranscribeServiceGetTranscriptionJobFuture

	GetVocabulary(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error)
	GetVocabularyAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyInput) *TranscribeServiceGetVocabularyFuture

	GetVocabularyFilter(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error)
	GetVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.GetVocabularyFilterInput) *TranscribeServiceGetVocabularyFilterFuture

	ListLanguageModels(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error)
	ListLanguageModelsAsync(ctx workflow.Context, input *transcribeservice.ListLanguageModelsInput) *TranscribeServiceListLanguageModelsFuture

	ListMedicalTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error)
	ListMedicalTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) *TranscribeServiceListMedicalTranscriptionJobsFuture

	ListMedicalVocabularies(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error)
	ListMedicalVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListMedicalVocabulariesInput) *TranscribeServiceListMedicalVocabulariesFuture

	ListTranscriptionJobs(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error)
	ListTranscriptionJobsAsync(ctx workflow.Context, input *transcribeservice.ListTranscriptionJobsInput) *TranscribeServiceListTranscriptionJobsFuture

	ListVocabularies(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error)
	ListVocabulariesAsync(ctx workflow.Context, input *transcribeservice.ListVocabulariesInput) *TranscribeServiceListVocabulariesFuture

	ListVocabularyFilters(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error)
	ListVocabularyFiltersAsync(ctx workflow.Context, input *transcribeservice.ListVocabularyFiltersInput) *TranscribeServiceListVocabularyFiltersFuture

	StartMedicalTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error)
	StartMedicalTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) *TranscribeServiceStartMedicalTranscriptionJobFuture

	StartTranscriptionJob(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error)
	StartTranscriptionJobAsync(ctx workflow.Context, input *transcribeservice.StartTranscriptionJobInput) *TranscribeServiceStartTranscriptionJobFuture

	UpdateMedicalVocabulary(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error)
	UpdateMedicalVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateMedicalVocabularyInput) *TranscribeServiceUpdateMedicalVocabularyFuture

	UpdateVocabulary(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error)
	UpdateVocabularyAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyInput) *TranscribeServiceUpdateVocabularyFuture

	UpdateVocabularyFilter(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error)
	UpdateVocabularyFilterAsync(ctx workflow.Context, input *transcribeservice.UpdateVocabularyFilterInput) *TranscribeServiceUpdateVocabularyFilterFuture
}

func NewClient() Client {
	return &stub{}
}
