
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TranscribeServiceActivities struct {
    client transcribeserviceiface.TranscribeServiceAPI
}

func NewTranscribeServiceActivities(session *session.Session, config ...*aws.Config) *TranscribeServiceActivities {
    client := transcribeservice.New(session, config...)
    return &TranscribeServiceActivities{client: client}
}

func (a *TranscribeServiceActivities) CreateLanguageModel(ctx context.Context, input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error) {
    return a.client.CreateLanguageModelWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) CreateMedicalVocabulary(ctx context.Context, input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
    return a.client.CreateMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) CreateVocabulary(ctx context.Context, input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error) {
    return a.client.CreateVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) CreateVocabularyFilter(ctx context.Context, input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error) {
    return a.client.CreateVocabularyFilterWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteLanguageModel(ctx context.Context, input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error) {
    return a.client.DeleteLanguageModelWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
    return a.client.DeleteMedicalTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteMedicalVocabulary(ctx context.Context, input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
    return a.client.DeleteMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteTranscriptionJob(ctx context.Context, input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
    return a.client.DeleteTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteVocabulary(ctx context.Context, input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error) {
    return a.client.DeleteVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DeleteVocabularyFilter(ctx context.Context, input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
    return a.client.DeleteVocabularyFilterWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) DescribeLanguageModel(ctx context.Context, input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
    return a.client.DescribeLanguageModelWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
    return a.client.GetMedicalTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetMedicalVocabulary(ctx context.Context, input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
    return a.client.GetMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetTranscriptionJob(ctx context.Context, input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
    return a.client.GetTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetVocabulary(ctx context.Context, input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
    return a.client.GetVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) GetVocabularyFilter(ctx context.Context, input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
    return a.client.GetVocabularyFilterWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListLanguageModels(ctx context.Context, input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
    return a.client.ListLanguageModelsWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListMedicalTranscriptionJobs(ctx context.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
    return a.client.ListMedicalTranscriptionJobsWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListMedicalVocabularies(ctx context.Context, input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
    return a.client.ListMedicalVocabulariesWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListTranscriptionJobs(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
    return a.client.ListTranscriptionJobsWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListVocabularies(ctx context.Context, input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
    return a.client.ListVocabulariesWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) ListVocabularyFilters(ctx context.Context, input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
    return a.client.ListVocabularyFiltersWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) StartMedicalTranscriptionJob(ctx context.Context, input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
    return a.client.StartMedicalTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) StartTranscriptionJob(ctx context.Context, input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error) {
    return a.client.StartTranscriptionJobWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) UpdateMedicalVocabulary(ctx context.Context, input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
    return a.client.UpdateMedicalVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) UpdateVocabulary(ctx context.Context, input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error) {
    return a.client.UpdateVocabularyWithContext(ctx, input)
}

func (a *TranscribeServiceActivities) UpdateVocabularyFilter(ctx context.Context, input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
    return a.client.UpdateVocabularyFilterWithContext(ctx, input)
}
