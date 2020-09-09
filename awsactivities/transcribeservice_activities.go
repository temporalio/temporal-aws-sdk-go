
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
)

type TranscribeServiceActivities struct {
    client transcribeserviceiface.TranscribeServiceAPI
}

func NewTranscribeServiceActivities(session *session.Session, config ...*aws.Config) *TranscribeServiceActivities {
    client := transcribeservice.New(session, config...)
    return &TranscribeServiceActivities{client: client}
}

func (a *TranscribeServiceActivities) CreateLanguageModel(input *transcribeservice.CreateLanguageModelInput) (*transcribeservice.CreateLanguageModelOutput, error) {
    return a.client.CreateLanguageModel(input)
}

func (a *TranscribeServiceActivities) CreateMedicalVocabulary(input *transcribeservice.CreateMedicalVocabularyInput) (*transcribeservice.CreateMedicalVocabularyOutput, error) {
    return a.client.CreateMedicalVocabulary(input)
}

func (a *TranscribeServiceActivities) CreateVocabulary(input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error) {
    return a.client.CreateVocabulary(input)
}

func (a *TranscribeServiceActivities) CreateVocabularyFilter(input *transcribeservice.CreateVocabularyFilterInput) (*transcribeservice.CreateVocabularyFilterOutput, error) {
    return a.client.CreateVocabularyFilter(input)
}

func (a *TranscribeServiceActivities) DeleteLanguageModel(input *transcribeservice.DeleteLanguageModelInput) (*transcribeservice.DeleteLanguageModelOutput, error) {
    return a.client.DeleteLanguageModel(input)
}

func (a *TranscribeServiceActivities) DeleteMedicalTranscriptionJob(input *transcribeservice.DeleteMedicalTranscriptionJobInput) (*transcribeservice.DeleteMedicalTranscriptionJobOutput, error) {
    return a.client.DeleteMedicalTranscriptionJob(input)
}

func (a *TranscribeServiceActivities) DeleteMedicalVocabulary(input *transcribeservice.DeleteMedicalVocabularyInput) (*transcribeservice.DeleteMedicalVocabularyOutput, error) {
    return a.client.DeleteMedicalVocabulary(input)
}

func (a *TranscribeServiceActivities) DeleteTranscriptionJob(input *transcribeservice.DeleteTranscriptionJobInput) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
    return a.client.DeleteTranscriptionJob(input)
}

func (a *TranscribeServiceActivities) DeleteVocabulary(input *transcribeservice.DeleteVocabularyInput) (*transcribeservice.DeleteVocabularyOutput, error) {
    return a.client.DeleteVocabulary(input)
}

func (a *TranscribeServiceActivities) DeleteVocabularyFilter(input *transcribeservice.DeleteVocabularyFilterInput) (*transcribeservice.DeleteVocabularyFilterOutput, error) {
    return a.client.DeleteVocabularyFilter(input)
}

func (a *TranscribeServiceActivities) DescribeLanguageModel(input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
    return a.client.DescribeLanguageModel(input)
}

func (a *TranscribeServiceActivities) GetMedicalTranscriptionJob(input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
    return a.client.GetMedicalTranscriptionJob(input)
}

func (a *TranscribeServiceActivities) GetMedicalVocabulary(input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
    return a.client.GetMedicalVocabulary(input)
}

func (a *TranscribeServiceActivities) GetTranscriptionJob(input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
    return a.client.GetTranscriptionJob(input)
}

func (a *TranscribeServiceActivities) GetVocabulary(input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
    return a.client.GetVocabulary(input)
}

func (a *TranscribeServiceActivities) GetVocabularyFilter(input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
    return a.client.GetVocabularyFilter(input)
}

func (a *TranscribeServiceActivities) ListLanguageModels(input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
    return a.client.ListLanguageModels(input)
}

func (a *TranscribeServiceActivities) ListMedicalTranscriptionJobs(input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
    return a.client.ListMedicalTranscriptionJobs(input)
}

func (a *TranscribeServiceActivities) ListMedicalVocabularies(input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
    return a.client.ListMedicalVocabularies(input)
}

func (a *TranscribeServiceActivities) ListTranscriptionJobs(input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
    return a.client.ListTranscriptionJobs(input)
}

func (a *TranscribeServiceActivities) ListVocabularies(input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
    return a.client.ListVocabularies(input)
}

func (a *TranscribeServiceActivities) ListVocabularyFilters(input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
    return a.client.ListVocabularyFilters(input)
}

func (a *TranscribeServiceActivities) StartMedicalTranscriptionJob(input *transcribeservice.StartMedicalTranscriptionJobInput) (*transcribeservice.StartMedicalTranscriptionJobOutput, error) {
    return a.client.StartMedicalTranscriptionJob(input)
}

func (a *TranscribeServiceActivities) StartTranscriptionJob(input *transcribeservice.StartTranscriptionJobInput) (*transcribeservice.StartTranscriptionJobOutput, error) {
    return a.client.StartTranscriptionJob(input)
}

func (a *TranscribeServiceActivities) UpdateMedicalVocabulary(input *transcribeservice.UpdateMedicalVocabularyInput) (*transcribeservice.UpdateMedicalVocabularyOutput, error) {
    return a.client.UpdateMedicalVocabulary(input)
}

func (a *TranscribeServiceActivities) UpdateVocabulary(input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error) {
    return a.client.UpdateVocabulary(input)
}

func (a *TranscribeServiceActivities) UpdateVocabularyFilter(input *transcribeservice.UpdateVocabularyFilterInput) (*transcribeservice.UpdateVocabularyFilterOutput, error) {
    return a.client.UpdateVocabularyFilter(input)
}
