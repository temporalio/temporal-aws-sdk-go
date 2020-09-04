package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
)

type RekognitionActivities struct {
	client rekognitioniface.RekognitionAPI
}

func NewRekognitionActivities(client rekognitioniface.RekognitionAPI) *RekognitionActivities {
    return &RekognitionActivities{client: client}
}


func (a *RekognitionActivities) CompareFaces(input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
    return a.client.CompareFaces(input)
}



func (a *RekognitionActivities) CreateCollection(input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error) {
    return a.client.CreateCollection(input)
}



func (a *RekognitionActivities) CreateProject(input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error) {
    return a.client.CreateProject(input)
}



func (a *RekognitionActivities) CreateProjectVersion(input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error) {
    return a.client.CreateProjectVersion(input)
}



func (a *RekognitionActivities) CreateStreamProcessor(input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error) {
    return a.client.CreateStreamProcessor(input)
}



func (a *RekognitionActivities) DeleteCollection(input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error) {
    return a.client.DeleteCollection(input)
}



func (a *RekognitionActivities) DeleteFaces(input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error) {
    return a.client.DeleteFaces(input)
}



func (a *RekognitionActivities) DeleteProject(input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error) {
    return a.client.DeleteProject(input)
}



func (a *RekognitionActivities) DeleteProjectVersion(input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error) {
    return a.client.DeleteProjectVersion(input)
}



func (a *RekognitionActivities) DeleteStreamProcessor(input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error) {
    return a.client.DeleteStreamProcessor(input)
}



func (a *RekognitionActivities) DescribeCollection(input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error) {
    return a.client.DescribeCollection(input)
}



func (a *RekognitionActivities) DescribeProjectVersions(input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error) {
    return a.client.DescribeProjectVersions(input)
}



func (a *RekognitionActivities) DescribeProjects(input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error) {
    return a.client.DescribeProjects(input)
}



func (a *RekognitionActivities) DescribeStreamProcessor(input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error) {
    return a.client.DescribeStreamProcessor(input)
}



func (a *RekognitionActivities) DetectCustomLabels(input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error) {
    return a.client.DetectCustomLabels(input)
}



func (a *RekognitionActivities) DetectFaces(input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
    return a.client.DetectFaces(input)
}



func (a *RekognitionActivities) DetectLabels(input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
    return a.client.DetectLabels(input)
}



func (a *RekognitionActivities) DetectModerationLabels(input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error) {
    return a.client.DetectModerationLabels(input)
}



func (a *RekognitionActivities) DetectText(input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
    return a.client.DetectText(input)
}



func (a *RekognitionActivities) GetCelebrityInfo(input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error) {
    return a.client.GetCelebrityInfo(input)
}



func (a *RekognitionActivities) GetCelebrityRecognition(input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error) {
    return a.client.GetCelebrityRecognition(input)
}



func (a *RekognitionActivities) GetContentModeration(input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error) {
    return a.client.GetContentModeration(input)
}



func (a *RekognitionActivities) GetFaceDetection(input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error) {
    return a.client.GetFaceDetection(input)
}



func (a *RekognitionActivities) GetFaceSearch(input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error) {
    return a.client.GetFaceSearch(input)
}



func (a *RekognitionActivities) GetLabelDetection(input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error) {
    return a.client.GetLabelDetection(input)
}



func (a *RekognitionActivities) GetPersonTracking(input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error) {
    return a.client.GetPersonTracking(input)
}



func (a *RekognitionActivities) GetSegmentDetection(input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error) {
    return a.client.GetSegmentDetection(input)
}



func (a *RekognitionActivities) GetTextDetection(input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error) {
    return a.client.GetTextDetection(input)
}



func (a *RekognitionActivities) IndexFaces(input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error) {
    return a.client.IndexFaces(input)
}



func (a *RekognitionActivities) ListCollections(input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error) {
    return a.client.ListCollections(input)
}



func (a *RekognitionActivities) ListFaces(input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error) {
    return a.client.ListFaces(input)
}



func (a *RekognitionActivities) ListStreamProcessors(input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error) {
    return a.client.ListStreamProcessors(input)
}



func (a *RekognitionActivities) RecognizeCelebrities(input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error) {
    return a.client.RecognizeCelebrities(input)
}



func (a *RekognitionActivities) SearchFaces(input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error) {
    return a.client.SearchFaces(input)
}



func (a *RekognitionActivities) SearchFacesByImage(input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error) {
    return a.client.SearchFacesByImage(input)
}



func (a *RekognitionActivities) StartCelebrityRecognition(input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error) {
    return a.client.StartCelebrityRecognition(input)
}



func (a *RekognitionActivities) StartContentModeration(input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error) {
    return a.client.StartContentModeration(input)
}



func (a *RekognitionActivities) StartFaceDetection(input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error) {
    return a.client.StartFaceDetection(input)
}



func (a *RekognitionActivities) StartFaceSearch(input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error) {
    return a.client.StartFaceSearch(input)
}



func (a *RekognitionActivities) StartLabelDetection(input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error) {
    return a.client.StartLabelDetection(input)
}



func (a *RekognitionActivities) StartPersonTracking(input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error) {
    return a.client.StartPersonTracking(input)
}



func (a *RekognitionActivities) StartProjectVersion(input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error) {
    return a.client.StartProjectVersion(input)
}



func (a *RekognitionActivities) StartSegmentDetection(input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error) {
    return a.client.StartSegmentDetection(input)
}



func (a *RekognitionActivities) StartStreamProcessor(input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error) {
    return a.client.StartStreamProcessor(input)
}



func (a *RekognitionActivities) StartTextDetection(input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error) {
    return a.client.StartTextDetection(input)
}



func (a *RekognitionActivities) StopProjectVersion(input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error) {
    return a.client.StopProjectVersion(input)
}



func (a *RekognitionActivities) StopStreamProcessor(input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error) {
    return a.client.StopStreamProcessor(input)
}



func (a *RekognitionActivities) WaitUntilProjectVersionRunning(input *rekognition.DescribeProjectVersionsInput) error {
    return a.client.WaitUntilProjectVersionRunning(input)
}



func (a *RekognitionActivities) WaitUntilProjectVersionTrainingCompleted(input *rekognition.DescribeProjectVersionsInput) error {
    return a.client.WaitUntilProjectVersionTrainingCompleted(input)
}

