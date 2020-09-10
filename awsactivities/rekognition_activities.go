package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type RekognitionActivities struct {
	client rekognitioniface.RekognitionAPI
}

func NewRekognitionActivities(session *session.Session, config ...*aws.Config) *RekognitionActivities {
	client := rekognition.New(session, config...)
	return &RekognitionActivities{client: client}
}

func (a *RekognitionActivities) CompareFaces(ctx context.Context, input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
	return a.client.CompareFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateCollection(ctx context.Context, input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error) {
	return a.client.CreateCollectionWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateProject(ctx context.Context, input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error) {
	return a.client.CreateProjectWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateProjectVersion(ctx context.Context, input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error) {
	return a.client.CreateProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateStreamProcessor(ctx context.Context, input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error) {
	return a.client.CreateStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteCollection(ctx context.Context, input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error) {
	return a.client.DeleteCollectionWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteFaces(ctx context.Context, input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error) {
	return a.client.DeleteFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteProject(ctx context.Context, input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error) {
	return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteProjectVersion(ctx context.Context, input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error) {
	return a.client.DeleteProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteStreamProcessor(ctx context.Context, input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error) {
	return a.client.DeleteStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeCollection(ctx context.Context, input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error) {
	return a.client.DescribeCollectionWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeProjectVersions(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error) {
	return a.client.DescribeProjectVersionsWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeProjects(ctx context.Context, input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error) {
	return a.client.DescribeProjectsWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeStreamProcessor(ctx context.Context, input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error) {
	return a.client.DescribeStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectCustomLabels(ctx context.Context, input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error) {
	return a.client.DetectCustomLabelsWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectFaces(ctx context.Context, input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
	return a.client.DetectFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectLabels(ctx context.Context, input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
	return a.client.DetectLabelsWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectModerationLabels(ctx context.Context, input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error) {
	return a.client.DetectModerationLabelsWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectText(ctx context.Context, input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
	return a.client.DetectTextWithContext(ctx, input)
}

func (a *RekognitionActivities) GetCelebrityInfo(ctx context.Context, input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error) {
	return a.client.GetCelebrityInfoWithContext(ctx, input)
}

func (a *RekognitionActivities) GetCelebrityRecognition(ctx context.Context, input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error) {
	return a.client.GetCelebrityRecognitionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetContentModeration(ctx context.Context, input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error) {
	return a.client.GetContentModerationWithContext(ctx, input)
}

func (a *RekognitionActivities) GetFaceDetection(ctx context.Context, input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error) {
	return a.client.GetFaceDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetFaceSearch(ctx context.Context, input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error) {
	return a.client.GetFaceSearchWithContext(ctx, input)
}

func (a *RekognitionActivities) GetLabelDetection(ctx context.Context, input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error) {
	return a.client.GetLabelDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetPersonTracking(ctx context.Context, input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error) {
	return a.client.GetPersonTrackingWithContext(ctx, input)
}

func (a *RekognitionActivities) GetSegmentDetection(ctx context.Context, input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error) {
	return a.client.GetSegmentDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetTextDetection(ctx context.Context, input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error) {
	return a.client.GetTextDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) IndexFaces(ctx context.Context, input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error) {
	return a.client.IndexFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) ListCollections(ctx context.Context, input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error) {
	return a.client.ListCollectionsWithContext(ctx, input)
}

func (a *RekognitionActivities) ListFaces(ctx context.Context, input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error) {
	return a.client.ListFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) ListStreamProcessors(ctx context.Context, input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error) {
	return a.client.ListStreamProcessorsWithContext(ctx, input)
}

func (a *RekognitionActivities) RecognizeCelebrities(ctx context.Context, input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error) {
	return a.client.RecognizeCelebritiesWithContext(ctx, input)
}

func (a *RekognitionActivities) SearchFaces(ctx context.Context, input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error) {
	return a.client.SearchFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) SearchFacesByImage(ctx context.Context, input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error) {
	return a.client.SearchFacesByImageWithContext(ctx, input)
}

func (a *RekognitionActivities) StartCelebrityRecognition(ctx context.Context, input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error) {
	return a.client.StartCelebrityRecognitionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartContentModeration(ctx context.Context, input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error) {
	return a.client.StartContentModerationWithContext(ctx, input)
}

func (a *RekognitionActivities) StartFaceDetection(ctx context.Context, input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error) {
	return a.client.StartFaceDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartFaceSearch(ctx context.Context, input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error) {
	return a.client.StartFaceSearchWithContext(ctx, input)
}

func (a *RekognitionActivities) StartLabelDetection(ctx context.Context, input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error) {
	return a.client.StartLabelDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartPersonTracking(ctx context.Context, input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error) {
	return a.client.StartPersonTrackingWithContext(ctx, input)
}

func (a *RekognitionActivities) StartProjectVersion(ctx context.Context, input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error) {
	return a.client.StartProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartSegmentDetection(ctx context.Context, input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error) {
	return a.client.StartSegmentDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartStreamProcessor(ctx context.Context, input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error) {
	return a.client.StartStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) StartTextDetection(ctx context.Context, input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error) {
	return a.client.StartTextDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StopProjectVersion(ctx context.Context, input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error) {
	return a.client.StopProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) StopStreamProcessor(ctx context.Context, input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error) {
	return a.client.StopStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) WaitUntilProjectVersionRunning(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) error {
	return a.client.WaitUntilProjectVersionRunningWithContext(ctx, input)

}

func (a *RekognitionActivities) WaitUntilProjectVersionTrainingCompleted(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) error {
	return a.client.WaitUntilProjectVersionTrainingCompletedWithContext(ctx, input)

}
