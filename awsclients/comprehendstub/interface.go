// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package comprehendstub

import (
	"github.com/aws/aws-sdk-go/service/comprehend"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	BatchDetectDominantLanguage(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error)
	BatchDetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.BatchDetectDominantLanguageInput) *ComprehendBatchDetectDominantLanguageFuture

	BatchDetectEntities(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error)
	BatchDetectEntitiesAsync(ctx workflow.Context, input *comprehend.BatchDetectEntitiesInput) *ComprehendBatchDetectEntitiesFuture

	BatchDetectKeyPhrases(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error)
	BatchDetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.BatchDetectKeyPhrasesInput) *ComprehendBatchDetectKeyPhrasesFuture

	BatchDetectSentiment(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error)
	BatchDetectSentimentAsync(ctx workflow.Context, input *comprehend.BatchDetectSentimentInput) *ComprehendBatchDetectSentimentFuture

	BatchDetectSyntax(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error)
	BatchDetectSyntaxAsync(ctx workflow.Context, input *comprehend.BatchDetectSyntaxInput) *ComprehendBatchDetectSyntaxFuture

	ClassifyDocument(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error)
	ClassifyDocumentAsync(ctx workflow.Context, input *comprehend.ClassifyDocumentInput) *ComprehendClassifyDocumentFuture

	CreateDocumentClassifier(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error)
	CreateDocumentClassifierAsync(ctx workflow.Context, input *comprehend.CreateDocumentClassifierInput) *ComprehendCreateDocumentClassifierFuture

	CreateEndpoint(ctx workflow.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error)
	CreateEndpointAsync(ctx workflow.Context, input *comprehend.CreateEndpointInput) *ComprehendCreateEndpointFuture

	CreateEntityRecognizer(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error)
	CreateEntityRecognizerAsync(ctx workflow.Context, input *comprehend.CreateEntityRecognizerInput) *ComprehendCreateEntityRecognizerFuture

	DeleteDocumentClassifier(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error)
	DeleteDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DeleteDocumentClassifierInput) *ComprehendDeleteDocumentClassifierFuture

	DeleteEndpoint(ctx workflow.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error)
	DeleteEndpointAsync(ctx workflow.Context, input *comprehend.DeleteEndpointInput) *ComprehendDeleteEndpointFuture

	DeleteEntityRecognizer(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error)
	DeleteEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DeleteEntityRecognizerInput) *ComprehendDeleteEntityRecognizerFuture

	DescribeDocumentClassificationJob(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error)
	DescribeDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassificationJobInput) *ComprehendDescribeDocumentClassificationJobFuture

	DescribeDocumentClassifier(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error)
	DescribeDocumentClassifierAsync(ctx workflow.Context, input *comprehend.DescribeDocumentClassifierInput) *ComprehendDescribeDocumentClassifierFuture

	DescribeDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error)
	DescribeDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) *ComprehendDescribeDominantLanguageDetectionJobFuture

	DescribeEndpoint(ctx workflow.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error)
	DescribeEndpointAsync(ctx workflow.Context, input *comprehend.DescribeEndpointInput) *ComprehendDescribeEndpointFuture

	DescribeEntitiesDetectionJob(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error)
	DescribeEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeEntitiesDetectionJobInput) *ComprehendDescribeEntitiesDetectionJobFuture

	DescribeEntityRecognizer(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error)
	DescribeEntityRecognizerAsync(ctx workflow.Context, input *comprehend.DescribeEntityRecognizerInput) *ComprehendDescribeEntityRecognizerFuture

	DescribeKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error)
	DescribeKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) *ComprehendDescribeKeyPhrasesDetectionJobFuture

	DescribePiiEntitiesDetectionJob(ctx workflow.Context, input *comprehend.DescribePiiEntitiesDetectionJobInput) (*comprehend.DescribePiiEntitiesDetectionJobOutput, error)
	DescribePiiEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribePiiEntitiesDetectionJobInput) *ComprehendDescribePiiEntitiesDetectionJobFuture

	DescribeSentimentDetectionJob(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error)
	DescribeSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeSentimentDetectionJobInput) *ComprehendDescribeSentimentDetectionJobFuture

	DescribeTopicsDetectionJob(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error)
	DescribeTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.DescribeTopicsDetectionJobInput) *ComprehendDescribeTopicsDetectionJobFuture

	DetectDominantLanguage(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error)
	DetectDominantLanguageAsync(ctx workflow.Context, input *comprehend.DetectDominantLanguageInput) *ComprehendDetectDominantLanguageFuture

	DetectEntities(ctx workflow.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error)
	DetectEntitiesAsync(ctx workflow.Context, input *comprehend.DetectEntitiesInput) *ComprehendDetectEntitiesFuture

	DetectKeyPhrases(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error)
	DetectKeyPhrasesAsync(ctx workflow.Context, input *comprehend.DetectKeyPhrasesInput) *ComprehendDetectKeyPhrasesFuture

	DetectPiiEntities(ctx workflow.Context, input *comprehend.DetectPiiEntitiesInput) (*comprehend.DetectPiiEntitiesOutput, error)
	DetectPiiEntitiesAsync(ctx workflow.Context, input *comprehend.DetectPiiEntitiesInput) *ComprehendDetectPiiEntitiesFuture

	DetectSentiment(ctx workflow.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error)
	DetectSentimentAsync(ctx workflow.Context, input *comprehend.DetectSentimentInput) *ComprehendDetectSentimentFuture

	DetectSyntax(ctx workflow.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error)
	DetectSyntaxAsync(ctx workflow.Context, input *comprehend.DetectSyntaxInput) *ComprehendDetectSyntaxFuture

	ListDocumentClassificationJobs(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error)
	ListDocumentClassificationJobsAsync(ctx workflow.Context, input *comprehend.ListDocumentClassificationJobsInput) *ComprehendListDocumentClassificationJobsFuture

	ListDocumentClassifiers(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error)
	ListDocumentClassifiersAsync(ctx workflow.Context, input *comprehend.ListDocumentClassifiersInput) *ComprehendListDocumentClassifiersFuture

	ListDominantLanguageDetectionJobs(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error)
	ListDominantLanguageDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) *ComprehendListDominantLanguageDetectionJobsFuture

	ListEndpoints(ctx workflow.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error)
	ListEndpointsAsync(ctx workflow.Context, input *comprehend.ListEndpointsInput) *ComprehendListEndpointsFuture

	ListEntitiesDetectionJobs(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error)
	ListEntitiesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListEntitiesDetectionJobsInput) *ComprehendListEntitiesDetectionJobsFuture

	ListEntityRecognizers(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error)
	ListEntityRecognizersAsync(ctx workflow.Context, input *comprehend.ListEntityRecognizersInput) *ComprehendListEntityRecognizersFuture

	ListKeyPhrasesDetectionJobs(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error)
	ListKeyPhrasesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) *ComprehendListKeyPhrasesDetectionJobsFuture

	ListPiiEntitiesDetectionJobs(ctx workflow.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput) (*comprehend.ListPiiEntitiesDetectionJobsOutput, error)
	ListPiiEntitiesDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListPiiEntitiesDetectionJobsInput) *ComprehendListPiiEntitiesDetectionJobsFuture

	ListSentimentDetectionJobs(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error)
	ListSentimentDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListSentimentDetectionJobsInput) *ComprehendListSentimentDetectionJobsFuture

	ListTagsForResource(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *comprehend.ListTagsForResourceInput) *ComprehendListTagsForResourceFuture

	ListTopicsDetectionJobs(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error)
	ListTopicsDetectionJobsAsync(ctx workflow.Context, input *comprehend.ListTopicsDetectionJobsInput) *ComprehendListTopicsDetectionJobsFuture

	StartDocumentClassificationJob(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error)
	StartDocumentClassificationJobAsync(ctx workflow.Context, input *comprehend.StartDocumentClassificationJobInput) *ComprehendStartDocumentClassificationJobFuture

	StartDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error)
	StartDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StartDominantLanguageDetectionJobInput) *ComprehendStartDominantLanguageDetectionJobFuture

	StartEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error)
	StartEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartEntitiesDetectionJobInput) *ComprehendStartEntitiesDetectionJobFuture

	StartKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error)
	StartKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) *ComprehendStartKeyPhrasesDetectionJobFuture

	StartPiiEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StartPiiEntitiesDetectionJobInput) (*comprehend.StartPiiEntitiesDetectionJobOutput, error)
	StartPiiEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StartPiiEntitiesDetectionJobInput) *ComprehendStartPiiEntitiesDetectionJobFuture

	StartSentimentDetectionJob(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error)
	StartSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StartSentimentDetectionJobInput) *ComprehendStartSentimentDetectionJobFuture

	StartTopicsDetectionJob(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error)
	StartTopicsDetectionJobAsync(ctx workflow.Context, input *comprehend.StartTopicsDetectionJobInput) *ComprehendStartTopicsDetectionJobFuture

	StopDominantLanguageDetectionJob(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error)
	StopDominantLanguageDetectionJobAsync(ctx workflow.Context, input *comprehend.StopDominantLanguageDetectionJobInput) *ComprehendStopDominantLanguageDetectionJobFuture

	StopEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error)
	StopEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopEntitiesDetectionJobInput) *ComprehendStopEntitiesDetectionJobFuture

	StopKeyPhrasesDetectionJob(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error)
	StopKeyPhrasesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) *ComprehendStopKeyPhrasesDetectionJobFuture

	StopPiiEntitiesDetectionJob(ctx workflow.Context, input *comprehend.StopPiiEntitiesDetectionJobInput) (*comprehend.StopPiiEntitiesDetectionJobOutput, error)
	StopPiiEntitiesDetectionJobAsync(ctx workflow.Context, input *comprehend.StopPiiEntitiesDetectionJobInput) *ComprehendStopPiiEntitiesDetectionJobFuture

	StopSentimentDetectionJob(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error)
	StopSentimentDetectionJobAsync(ctx workflow.Context, input *comprehend.StopSentimentDetectionJobInput) *ComprehendStopSentimentDetectionJobFuture

	StopTrainingDocumentClassifier(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error)
	StopTrainingDocumentClassifierAsync(ctx workflow.Context, input *comprehend.StopTrainingDocumentClassifierInput) *ComprehendStopTrainingDocumentClassifierFuture

	StopTrainingEntityRecognizer(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error)
	StopTrainingEntityRecognizerAsync(ctx workflow.Context, input *comprehend.StopTrainingEntityRecognizerInput) *ComprehendStopTrainingEntityRecognizerFuture

	TagResource(ctx workflow.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *comprehend.TagResourceInput) *ComprehendTagResourceFuture

	UntagResource(ctx workflow.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *comprehend.UntagResourceInput) *ComprehendUntagResourceFuture

	UpdateEndpoint(ctx workflow.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error)
	UpdateEndpointAsync(ctx workflow.Context, input *comprehend.UpdateEndpointInput) *ComprehendUpdateEndpointFuture
}

func NewClient() Client {
	return &stub{}
}
