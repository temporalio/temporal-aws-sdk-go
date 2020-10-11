// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type RekognitionActivities struct {
	client rekognitioniface.RekognitionAPI

	sessionFactory SessionFactory
}

func NewRekognitionActivities(sess *session.Session, config ...*aws.Config) *RekognitionActivities {
	client := rekognition.New(sess, config...)
	return &RekognitionActivities{client: client}
}

func NewRekognitionActivitiesWithSessionFactory(sessionFactory SessionFactory) *RekognitionActivities {
	return &RekognitionActivities{sessionFactory: sessionFactory}
}

func (a *RekognitionActivities) getClient(ctx context.Context) (rekognitioniface.RekognitionAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return rekognition.New(sess), nil
}

func (a *RekognitionActivities) CompareFaces(ctx context.Context, input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CompareFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateCollection(ctx context.Context, input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateCollectionWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateProject(ctx context.Context, input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProjectWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateProjectVersion(ctx context.Context, input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) CreateStreamProcessor(ctx context.Context, input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteCollection(ctx context.Context, input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCollectionWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteFaces(ctx context.Context, input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteProject(ctx context.Context, input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProjectWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteProjectVersion(ctx context.Context, input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) DeleteStreamProcessor(ctx context.Context, input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeCollection(ctx context.Context, input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCollectionWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeProjectVersions(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProjectVersionsWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeProjects(ctx context.Context, input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProjectsWithContext(ctx, input)
}

func (a *RekognitionActivities) DescribeStreamProcessor(ctx context.Context, input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectCustomLabels(ctx context.Context, input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectCustomLabelsWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectFaces(ctx context.Context, input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectLabels(ctx context.Context, input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectLabelsWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectModerationLabels(ctx context.Context, input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectModerationLabelsWithContext(ctx, input)
}

func (a *RekognitionActivities) DetectText(ctx context.Context, input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectTextWithContext(ctx, input)
}

func (a *RekognitionActivities) GetCelebrityInfo(ctx context.Context, input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCelebrityInfoWithContext(ctx, input)
}

func (a *RekognitionActivities) GetCelebrityRecognition(ctx context.Context, input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCelebrityRecognitionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetContentModeration(ctx context.Context, input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetContentModerationWithContext(ctx, input)
}

func (a *RekognitionActivities) GetFaceDetection(ctx context.Context, input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFaceDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetFaceSearch(ctx context.Context, input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFaceSearchWithContext(ctx, input)
}

func (a *RekognitionActivities) GetLabelDetection(ctx context.Context, input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLabelDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetPersonTracking(ctx context.Context, input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPersonTrackingWithContext(ctx, input)
}

func (a *RekognitionActivities) GetSegmentDetection(ctx context.Context, input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSegmentDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) GetTextDetection(ctx context.Context, input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTextDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) IndexFaces(ctx context.Context, input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.IndexFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) ListCollections(ctx context.Context, input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCollectionsWithContext(ctx, input)
}

func (a *RekognitionActivities) ListFaces(ctx context.Context, input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) ListStreamProcessors(ctx context.Context, input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListStreamProcessorsWithContext(ctx, input)
}

func (a *RekognitionActivities) RecognizeCelebrities(ctx context.Context, input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RecognizeCelebritiesWithContext(ctx, input)
}

func (a *RekognitionActivities) SearchFaces(ctx context.Context, input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchFacesWithContext(ctx, input)
}

func (a *RekognitionActivities) SearchFacesByImage(ctx context.Context, input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SearchFacesByImageWithContext(ctx, input)
}

func (a *RekognitionActivities) StartCelebrityRecognition(ctx context.Context, input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartCelebrityRecognitionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartContentModeration(ctx context.Context, input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartContentModerationWithContext(ctx, input)
}

func (a *RekognitionActivities) StartFaceDetection(ctx context.Context, input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartFaceDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartFaceSearch(ctx context.Context, input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartFaceSearchWithContext(ctx, input)
}

func (a *RekognitionActivities) StartLabelDetection(ctx context.Context, input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartLabelDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartPersonTracking(ctx context.Context, input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartPersonTrackingWithContext(ctx, input)
}

func (a *RekognitionActivities) StartProjectVersion(ctx context.Context, input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartSegmentDetection(ctx context.Context, input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartSegmentDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StartStreamProcessor(ctx context.Context, input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) StartTextDetection(ctx context.Context, input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartTextDetectionWithContext(ctx, input)
}

func (a *RekognitionActivities) StopProjectVersion(ctx context.Context, input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopProjectVersionWithContext(ctx, input)
}

func (a *RekognitionActivities) StopStreamProcessor(ctx context.Context, input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopStreamProcessorWithContext(ctx, input)
}

func (a *RekognitionActivities) WaitUntilProjectVersionRunning(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilProjectVersionRunningWithContext(ctx, input, options...)
	})
}

func (a *RekognitionActivities) WaitUntilProjectVersionTrainingCompleted(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilProjectVersionTrainingCompletedWithContext(ctx, input, options...)
	})
}
