package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ComprehendActivities struct {
	client comprehendiface.ComprehendAPI
}

func NewComprehendActivities(session *session.Session, config ...*aws.Config) *ComprehendActivities {
	client := comprehend.New(session, config...)
	return &ComprehendActivities{client: client}
}

func (a *ComprehendActivities) BatchDetectDominantLanguage(ctx context.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
	return a.client.BatchDetectDominantLanguageWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectEntities(ctx context.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error) {
	return a.client.BatchDetectEntitiesWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectKeyPhrases(ctx context.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
	return a.client.BatchDetectKeyPhrasesWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectSentiment(ctx context.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error) {
	return a.client.BatchDetectSentimentWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectSyntax(ctx context.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error) {
	return a.client.BatchDetectSyntaxWithContext(ctx, input)
}

func (a *ComprehendActivities) ClassifyDocument(ctx context.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error) {
	return a.client.ClassifyDocumentWithContext(ctx, input)
}

func (a *ComprehendActivities) CreateDocumentClassifier(ctx context.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error) {
	return a.client.CreateDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) CreateEndpoint(ctx context.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error) {
	return a.client.CreateEndpointWithContext(ctx, input)
}

func (a *ComprehendActivities) CreateEntityRecognizer(ctx context.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error) {
	return a.client.CreateEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) DeleteDocumentClassifier(ctx context.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error) {
	return a.client.DeleteDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) DeleteEndpoint(ctx context.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error) {
	return a.client.DeleteEndpointWithContext(ctx, input)
}

func (a *ComprehendActivities) DeleteEntityRecognizer(ctx context.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error) {
	return a.client.DeleteEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeDocumentClassificationJob(ctx context.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
	return a.client.DescribeDocumentClassificationJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeDocumentClassifier(ctx context.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error) {
	return a.client.DescribeDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeDominantLanguageDetectionJob(ctx context.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
	return a.client.DescribeDominantLanguageDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeEndpoint(ctx context.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error) {
	return a.client.DescribeEndpointWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeEntitiesDetectionJob(ctx context.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
	return a.client.DescribeEntitiesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeEntityRecognizer(ctx context.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error) {
	return a.client.DescribeEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
	return a.client.DescribeKeyPhrasesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeSentimentDetectionJob(ctx context.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
	return a.client.DescribeSentimentDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeTopicsDetectionJob(ctx context.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
	return a.client.DescribeTopicsDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectDominantLanguage(ctx context.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error) {
	return a.client.DetectDominantLanguageWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectEntities(ctx context.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error) {
	return a.client.DetectEntitiesWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectKeyPhrases(ctx context.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error) {
	return a.client.DetectKeyPhrasesWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectSentiment(ctx context.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error) {
	return a.client.DetectSentimentWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectSyntax(ctx context.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error) {
	return a.client.DetectSyntaxWithContext(ctx, input)
}

func (a *ComprehendActivities) ListDocumentClassificationJobs(ctx context.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error) {
	return a.client.ListDocumentClassificationJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListDocumentClassifiers(ctx context.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error) {
	return a.client.ListDocumentClassifiersWithContext(ctx, input)
}

func (a *ComprehendActivities) ListDominantLanguageDetectionJobs(ctx context.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
	return a.client.ListDominantLanguageDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListEndpoints(ctx context.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error) {
	return a.client.ListEndpointsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListEntitiesDetectionJobs(ctx context.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
	return a.client.ListEntitiesDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListEntityRecognizers(ctx context.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error) {
	return a.client.ListEntityRecognizersWithContext(ctx, input)
}

func (a *ComprehendActivities) ListKeyPhrasesDetectionJobs(ctx context.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
	return a.client.ListKeyPhrasesDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListSentimentDetectionJobs(ctx context.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error) {
	return a.client.ListSentimentDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListTagsForResource(ctx context.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ComprehendActivities) ListTopicsDetectionJobs(ctx context.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error) {
	return a.client.ListTopicsDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) StartDocumentClassificationJob(ctx context.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error) {
	return a.client.StartDocumentClassificationJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartDominantLanguageDetectionJob(ctx context.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error) {
	return a.client.StartDominantLanguageDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartEntitiesDetectionJob(ctx context.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error) {
	return a.client.StartEntitiesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error) {
	return a.client.StartKeyPhrasesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartSentimentDetectionJob(ctx context.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error) {
	return a.client.StartSentimentDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartTopicsDetectionJob(ctx context.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error) {
	return a.client.StartTopicsDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopDominantLanguageDetectionJob(ctx context.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error) {
	return a.client.StopDominantLanguageDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopEntitiesDetectionJob(ctx context.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error) {
	return a.client.StopEntitiesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error) {
	return a.client.StopKeyPhrasesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopSentimentDetectionJob(ctx context.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error) {
	return a.client.StopSentimentDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopTrainingDocumentClassifier(ctx context.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error) {
	return a.client.StopTrainingDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) StopTrainingEntityRecognizer(ctx context.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error) {
	return a.client.StopTrainingEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) TagResource(ctx context.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ComprehendActivities) UntagResource(ctx context.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ComprehendActivities) UpdateEndpoint(ctx context.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error) {
	return a.client.UpdateEndpointWithContext(ctx, input)
}
