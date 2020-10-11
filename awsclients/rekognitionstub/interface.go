// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package rekognitionstub

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CompareFaces(ctx workflow.Context, input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error)
	CompareFacesAsync(ctx workflow.Context, input *rekognition.CompareFacesInput) *RekognitionCompareFacesFuture

	CreateCollection(ctx workflow.Context, input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error)
	CreateCollectionAsync(ctx workflow.Context, input *rekognition.CreateCollectionInput) *RekognitionCreateCollectionFuture

	CreateProject(ctx workflow.Context, input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *rekognition.CreateProjectInput) *RekognitionCreateProjectFuture

	CreateProjectVersion(ctx workflow.Context, input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error)
	CreateProjectVersionAsync(ctx workflow.Context, input *rekognition.CreateProjectVersionInput) *RekognitionCreateProjectVersionFuture

	CreateStreamProcessor(ctx workflow.Context, input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error)
	CreateStreamProcessorAsync(ctx workflow.Context, input *rekognition.CreateStreamProcessorInput) *RekognitionCreateStreamProcessorFuture

	DeleteCollection(ctx workflow.Context, input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error)
	DeleteCollectionAsync(ctx workflow.Context, input *rekognition.DeleteCollectionInput) *RekognitionDeleteCollectionFuture

	DeleteFaces(ctx workflow.Context, input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error)
	DeleteFacesAsync(ctx workflow.Context, input *rekognition.DeleteFacesInput) *RekognitionDeleteFacesFuture

	DeleteProject(ctx workflow.Context, input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *rekognition.DeleteProjectInput) *RekognitionDeleteProjectFuture

	DeleteProjectVersion(ctx workflow.Context, input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error)
	DeleteProjectVersionAsync(ctx workflow.Context, input *rekognition.DeleteProjectVersionInput) *RekognitionDeleteProjectVersionFuture

	DeleteStreamProcessor(ctx workflow.Context, input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error)
	DeleteStreamProcessorAsync(ctx workflow.Context, input *rekognition.DeleteStreamProcessorInput) *RekognitionDeleteStreamProcessorFuture

	DescribeCollection(ctx workflow.Context, input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error)
	DescribeCollectionAsync(ctx workflow.Context, input *rekognition.DescribeCollectionInput) *RekognitionDescribeCollectionFuture

	DescribeProjectVersions(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error)
	DescribeProjectVersionsAsync(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) *RekognitionDescribeProjectVersionsFuture

	DescribeProjects(ctx workflow.Context, input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error)
	DescribeProjectsAsync(ctx workflow.Context, input *rekognition.DescribeProjectsInput) *RekognitionDescribeProjectsFuture

	DescribeStreamProcessor(ctx workflow.Context, input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error)
	DescribeStreamProcessorAsync(ctx workflow.Context, input *rekognition.DescribeStreamProcessorInput) *RekognitionDescribeStreamProcessorFuture

	DetectCustomLabels(ctx workflow.Context, input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error)
	DetectCustomLabelsAsync(ctx workflow.Context, input *rekognition.DetectCustomLabelsInput) *RekognitionDetectCustomLabelsFuture

	DetectFaces(ctx workflow.Context, input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error)
	DetectFacesAsync(ctx workflow.Context, input *rekognition.DetectFacesInput) *RekognitionDetectFacesFuture

	DetectLabels(ctx workflow.Context, input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error)
	DetectLabelsAsync(ctx workflow.Context, input *rekognition.DetectLabelsInput) *RekognitionDetectLabelsFuture

	DetectModerationLabels(ctx workflow.Context, input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error)
	DetectModerationLabelsAsync(ctx workflow.Context, input *rekognition.DetectModerationLabelsInput) *RekognitionDetectModerationLabelsFuture

	DetectText(ctx workflow.Context, input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error)
	DetectTextAsync(ctx workflow.Context, input *rekognition.DetectTextInput) *RekognitionDetectTextFuture

	GetCelebrityInfo(ctx workflow.Context, input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error)
	GetCelebrityInfoAsync(ctx workflow.Context, input *rekognition.GetCelebrityInfoInput) *RekognitionGetCelebrityInfoFuture

	GetCelebrityRecognition(ctx workflow.Context, input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error)
	GetCelebrityRecognitionAsync(ctx workflow.Context, input *rekognition.GetCelebrityRecognitionInput) *RekognitionGetCelebrityRecognitionFuture

	GetContentModeration(ctx workflow.Context, input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error)
	GetContentModerationAsync(ctx workflow.Context, input *rekognition.GetContentModerationInput) *RekognitionGetContentModerationFuture

	GetFaceDetection(ctx workflow.Context, input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error)
	GetFaceDetectionAsync(ctx workflow.Context, input *rekognition.GetFaceDetectionInput) *RekognitionGetFaceDetectionFuture

	GetFaceSearch(ctx workflow.Context, input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error)
	GetFaceSearchAsync(ctx workflow.Context, input *rekognition.GetFaceSearchInput) *RekognitionGetFaceSearchFuture

	GetLabelDetection(ctx workflow.Context, input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error)
	GetLabelDetectionAsync(ctx workflow.Context, input *rekognition.GetLabelDetectionInput) *RekognitionGetLabelDetectionFuture

	GetPersonTracking(ctx workflow.Context, input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error)
	GetPersonTrackingAsync(ctx workflow.Context, input *rekognition.GetPersonTrackingInput) *RekognitionGetPersonTrackingFuture

	GetSegmentDetection(ctx workflow.Context, input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error)
	GetSegmentDetectionAsync(ctx workflow.Context, input *rekognition.GetSegmentDetectionInput) *RekognitionGetSegmentDetectionFuture

	GetTextDetection(ctx workflow.Context, input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error)
	GetTextDetectionAsync(ctx workflow.Context, input *rekognition.GetTextDetectionInput) *RekognitionGetTextDetectionFuture

	IndexFaces(ctx workflow.Context, input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error)
	IndexFacesAsync(ctx workflow.Context, input *rekognition.IndexFacesInput) *RekognitionIndexFacesFuture

	ListCollections(ctx workflow.Context, input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error)
	ListCollectionsAsync(ctx workflow.Context, input *rekognition.ListCollectionsInput) *RekognitionListCollectionsFuture

	ListFaces(ctx workflow.Context, input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error)
	ListFacesAsync(ctx workflow.Context, input *rekognition.ListFacesInput) *RekognitionListFacesFuture

	ListStreamProcessors(ctx workflow.Context, input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error)
	ListStreamProcessorsAsync(ctx workflow.Context, input *rekognition.ListStreamProcessorsInput) *RekognitionListStreamProcessorsFuture

	RecognizeCelebrities(ctx workflow.Context, input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error)
	RecognizeCelebritiesAsync(ctx workflow.Context, input *rekognition.RecognizeCelebritiesInput) *RekognitionRecognizeCelebritiesFuture

	SearchFaces(ctx workflow.Context, input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error)
	SearchFacesAsync(ctx workflow.Context, input *rekognition.SearchFacesInput) *RekognitionSearchFacesFuture

	SearchFacesByImage(ctx workflow.Context, input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error)
	SearchFacesByImageAsync(ctx workflow.Context, input *rekognition.SearchFacesByImageInput) *RekognitionSearchFacesByImageFuture

	StartCelebrityRecognition(ctx workflow.Context, input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error)
	StartCelebrityRecognitionAsync(ctx workflow.Context, input *rekognition.StartCelebrityRecognitionInput) *RekognitionStartCelebrityRecognitionFuture

	StartContentModeration(ctx workflow.Context, input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error)
	StartContentModerationAsync(ctx workflow.Context, input *rekognition.StartContentModerationInput) *RekognitionStartContentModerationFuture

	StartFaceDetection(ctx workflow.Context, input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error)
	StartFaceDetectionAsync(ctx workflow.Context, input *rekognition.StartFaceDetectionInput) *RekognitionStartFaceDetectionFuture

	StartFaceSearch(ctx workflow.Context, input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error)
	StartFaceSearchAsync(ctx workflow.Context, input *rekognition.StartFaceSearchInput) *RekognitionStartFaceSearchFuture

	StartLabelDetection(ctx workflow.Context, input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error)
	StartLabelDetectionAsync(ctx workflow.Context, input *rekognition.StartLabelDetectionInput) *RekognitionStartLabelDetectionFuture

	StartPersonTracking(ctx workflow.Context, input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error)
	StartPersonTrackingAsync(ctx workflow.Context, input *rekognition.StartPersonTrackingInput) *RekognitionStartPersonTrackingFuture

	StartProjectVersion(ctx workflow.Context, input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error)
	StartProjectVersionAsync(ctx workflow.Context, input *rekognition.StartProjectVersionInput) *RekognitionStartProjectVersionFuture

	StartSegmentDetection(ctx workflow.Context, input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error)
	StartSegmentDetectionAsync(ctx workflow.Context, input *rekognition.StartSegmentDetectionInput) *RekognitionStartSegmentDetectionFuture

	StartStreamProcessor(ctx workflow.Context, input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error)
	StartStreamProcessorAsync(ctx workflow.Context, input *rekognition.StartStreamProcessorInput) *RekognitionStartStreamProcessorFuture

	StartTextDetection(ctx workflow.Context, input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error)
	StartTextDetectionAsync(ctx workflow.Context, input *rekognition.StartTextDetectionInput) *RekognitionStartTextDetectionFuture

	StopProjectVersion(ctx workflow.Context, input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error)
	StopProjectVersionAsync(ctx workflow.Context, input *rekognition.StopProjectVersionInput) *RekognitionStopProjectVersionFuture

	StopStreamProcessor(ctx workflow.Context, input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error)
	StopStreamProcessorAsync(ctx workflow.Context, input *rekognition.StopStreamProcessorInput) *RekognitionStopStreamProcessorFuture

	WaitUntilProjectVersionRunning(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) error
	WaitUntilProjectVersionRunningAsync(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) *awsclients.VoidFuture

	WaitUntilProjectVersionTrainingCompleted(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) error
	WaitUntilProjectVersionTrainingCompletedAsync(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) *awsclients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
