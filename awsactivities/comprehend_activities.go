
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
)

type ComprehendActivities struct {
    client comprehendiface.ComprehendAPI
}

func NewComprehendActivities(session *session.Session, config ...*aws.Config) *ComprehendActivities {
    client := comprehend.New(session, config...)
    return &ComprehendActivities{client: client}
}

func (a *ComprehendActivities) BatchDetectDominantLanguage(input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
    return a.client.BatchDetectDominantLanguage(input)
}

func (a *ComprehendActivities) BatchDetectEntities(input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error) {
    return a.client.BatchDetectEntities(input)
}

func (a *ComprehendActivities) BatchDetectKeyPhrases(input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
    return a.client.BatchDetectKeyPhrases(input)
}

func (a *ComprehendActivities) BatchDetectSentiment(input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error) {
    return a.client.BatchDetectSentiment(input)
}

func (a *ComprehendActivities) BatchDetectSyntax(input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error) {
    return a.client.BatchDetectSyntax(input)
}

func (a *ComprehendActivities) ClassifyDocument(input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error) {
    return a.client.ClassifyDocument(input)
}

func (a *ComprehendActivities) CreateDocumentClassifier(input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error) {
    return a.client.CreateDocumentClassifier(input)
}

func (a *ComprehendActivities) CreateEndpoint(input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error) {
    return a.client.CreateEndpoint(input)
}

func (a *ComprehendActivities) CreateEntityRecognizer(input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error) {
    return a.client.CreateEntityRecognizer(input)
}

func (a *ComprehendActivities) DeleteDocumentClassifier(input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error) {
    return a.client.DeleteDocumentClassifier(input)
}

func (a *ComprehendActivities) DeleteEndpoint(input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error) {
    return a.client.DeleteEndpoint(input)
}

func (a *ComprehendActivities) DeleteEntityRecognizer(input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error) {
    return a.client.DeleteEntityRecognizer(input)
}

func (a *ComprehendActivities) DescribeDocumentClassificationJob(input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
    return a.client.DescribeDocumentClassificationJob(input)
}

func (a *ComprehendActivities) DescribeDocumentClassifier(input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error) {
    return a.client.DescribeDocumentClassifier(input)
}

func (a *ComprehendActivities) DescribeDominantLanguageDetectionJob(input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
    return a.client.DescribeDominantLanguageDetectionJob(input)
}

func (a *ComprehendActivities) DescribeEndpoint(input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error) {
    return a.client.DescribeEndpoint(input)
}

func (a *ComprehendActivities) DescribeEntitiesDetectionJob(input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
    return a.client.DescribeEntitiesDetectionJob(input)
}

func (a *ComprehendActivities) DescribeEntityRecognizer(input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error) {
    return a.client.DescribeEntityRecognizer(input)
}

func (a *ComprehendActivities) DescribeKeyPhrasesDetectionJob(input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
    return a.client.DescribeKeyPhrasesDetectionJob(input)
}

func (a *ComprehendActivities) DescribeSentimentDetectionJob(input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
    return a.client.DescribeSentimentDetectionJob(input)
}

func (a *ComprehendActivities) DescribeTopicsDetectionJob(input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
    return a.client.DescribeTopicsDetectionJob(input)
}

func (a *ComprehendActivities) DetectDominantLanguage(input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error) {
    return a.client.DetectDominantLanguage(input)
}

func (a *ComprehendActivities) DetectEntities(input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error) {
    return a.client.DetectEntities(input)
}

func (a *ComprehendActivities) DetectKeyPhrases(input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error) {
    return a.client.DetectKeyPhrases(input)
}

func (a *ComprehendActivities) DetectSentiment(input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error) {
    return a.client.DetectSentiment(input)
}

func (a *ComprehendActivities) DetectSyntax(input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error) {
    return a.client.DetectSyntax(input)
}

func (a *ComprehendActivities) ListDocumentClassificationJobs(input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error) {
    return a.client.ListDocumentClassificationJobs(input)
}

func (a *ComprehendActivities) ListDocumentClassifiers(input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error) {
    return a.client.ListDocumentClassifiers(input)
}

func (a *ComprehendActivities) ListDominantLanguageDetectionJobs(input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
    return a.client.ListDominantLanguageDetectionJobs(input)
}

func (a *ComprehendActivities) ListEndpoints(input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error) {
    return a.client.ListEndpoints(input)
}

func (a *ComprehendActivities) ListEntitiesDetectionJobs(input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
    return a.client.ListEntitiesDetectionJobs(input)
}

func (a *ComprehendActivities) ListEntityRecognizers(input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error) {
    return a.client.ListEntityRecognizers(input)
}

func (a *ComprehendActivities) ListKeyPhrasesDetectionJobs(input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
    return a.client.ListKeyPhrasesDetectionJobs(input)
}

func (a *ComprehendActivities) ListSentimentDetectionJobs(input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error) {
    return a.client.ListSentimentDetectionJobs(input)
}

func (a *ComprehendActivities) ListTagsForResource(input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ComprehendActivities) ListTopicsDetectionJobs(input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error) {
    return a.client.ListTopicsDetectionJobs(input)
}

func (a *ComprehendActivities) StartDocumentClassificationJob(input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error) {
    return a.client.StartDocumentClassificationJob(input)
}

func (a *ComprehendActivities) StartDominantLanguageDetectionJob(input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error) {
    return a.client.StartDominantLanguageDetectionJob(input)
}

func (a *ComprehendActivities) StartEntitiesDetectionJob(input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error) {
    return a.client.StartEntitiesDetectionJob(input)
}

func (a *ComprehendActivities) StartKeyPhrasesDetectionJob(input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error) {
    return a.client.StartKeyPhrasesDetectionJob(input)
}

func (a *ComprehendActivities) StartSentimentDetectionJob(input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error) {
    return a.client.StartSentimentDetectionJob(input)
}

func (a *ComprehendActivities) StartTopicsDetectionJob(input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error) {
    return a.client.StartTopicsDetectionJob(input)
}

func (a *ComprehendActivities) StopDominantLanguageDetectionJob(input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error) {
    return a.client.StopDominantLanguageDetectionJob(input)
}

func (a *ComprehendActivities) StopEntitiesDetectionJob(input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error) {
    return a.client.StopEntitiesDetectionJob(input)
}

func (a *ComprehendActivities) StopKeyPhrasesDetectionJob(input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error) {
    return a.client.StopKeyPhrasesDetectionJob(input)
}

func (a *ComprehendActivities) StopSentimentDetectionJob(input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error) {
    return a.client.StopSentimentDetectionJob(input)
}

func (a *ComprehendActivities) StopTrainingDocumentClassifier(input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error) {
    return a.client.StopTrainingDocumentClassifier(input)
}

func (a *ComprehendActivities) StopTrainingEntityRecognizer(input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error) {
    return a.client.StopTrainingEntityRecognizer(input)
}

func (a *ComprehendActivities) TagResource(input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ComprehendActivities) UntagResource(input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ComprehendActivities) UpdateEndpoint(input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error) {
    return a.client.UpdateEndpoint(input)
}
