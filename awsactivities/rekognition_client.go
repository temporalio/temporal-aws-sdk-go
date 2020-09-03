package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"go.temporal.io/sdk/workflow"
)

type RekognitionClient interface {
    CompareFaces(ctx workflow.Context, input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error)
    CompareFacesAsync(ctx workflow.Context, input *rekognition.CompareFacesInput) *RekognitionCompareFacesResult

    CreateCollection(ctx workflow.Context, input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error)
    CreateCollectionAsync(ctx workflow.Context, input *rekognition.CreateCollectionInput) *RekognitionCreateCollectionResult

    CreateProject(ctx workflow.Context, input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error)
    CreateProjectAsync(ctx workflow.Context, input *rekognition.CreateProjectInput) *RekognitionCreateProjectResult

    CreateProjectVersion(ctx workflow.Context, input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error)
    CreateProjectVersionAsync(ctx workflow.Context, input *rekognition.CreateProjectVersionInput) *RekognitionCreateProjectVersionResult

    CreateStreamProcessor(ctx workflow.Context, input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error)
    CreateStreamProcessorAsync(ctx workflow.Context, input *rekognition.CreateStreamProcessorInput) *RekognitionCreateStreamProcessorResult

    DeleteCollection(ctx workflow.Context, input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error)
    DeleteCollectionAsync(ctx workflow.Context, input *rekognition.DeleteCollectionInput) *RekognitionDeleteCollectionResult

    DeleteFaces(ctx workflow.Context, input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error)
    DeleteFacesAsync(ctx workflow.Context, input *rekognition.DeleteFacesInput) *RekognitionDeleteFacesResult

    DeleteProject(ctx workflow.Context, input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error)
    DeleteProjectAsync(ctx workflow.Context, input *rekognition.DeleteProjectInput) *RekognitionDeleteProjectResult

    DeleteProjectVersion(ctx workflow.Context, input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error)
    DeleteProjectVersionAsync(ctx workflow.Context, input *rekognition.DeleteProjectVersionInput) *RekognitionDeleteProjectVersionResult

    DeleteStreamProcessor(ctx workflow.Context, input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error)
    DeleteStreamProcessorAsync(ctx workflow.Context, input *rekognition.DeleteStreamProcessorInput) *RekognitionDeleteStreamProcessorResult

    DescribeCollection(ctx workflow.Context, input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error)
    DescribeCollectionAsync(ctx workflow.Context, input *rekognition.DescribeCollectionInput) *RekognitionDescribeCollectionResult

    DescribeProjectVersions(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error)
    DescribeProjectVersionsAsync(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) *RekognitionDescribeProjectVersionsResult

    DescribeProjects(ctx workflow.Context, input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error)
    DescribeProjectsAsync(ctx workflow.Context, input *rekognition.DescribeProjectsInput) *RekognitionDescribeProjectsResult

    DescribeStreamProcessor(ctx workflow.Context, input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error)
    DescribeStreamProcessorAsync(ctx workflow.Context, input *rekognition.DescribeStreamProcessorInput) *RekognitionDescribeStreamProcessorResult

    DetectCustomLabels(ctx workflow.Context, input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error)
    DetectCustomLabelsAsync(ctx workflow.Context, input *rekognition.DetectCustomLabelsInput) *RekognitionDetectCustomLabelsResult

    DetectFaces(ctx workflow.Context, input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error)
    DetectFacesAsync(ctx workflow.Context, input *rekognition.DetectFacesInput) *RekognitionDetectFacesResult

    DetectLabels(ctx workflow.Context, input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error)
    DetectLabelsAsync(ctx workflow.Context, input *rekognition.DetectLabelsInput) *RekognitionDetectLabelsResult

    DetectModerationLabels(ctx workflow.Context, input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error)
    DetectModerationLabelsAsync(ctx workflow.Context, input *rekognition.DetectModerationLabelsInput) *RekognitionDetectModerationLabelsResult

    DetectText(ctx workflow.Context, input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error)
    DetectTextAsync(ctx workflow.Context, input *rekognition.DetectTextInput) *RekognitionDetectTextResult

    GetCelebrityInfo(ctx workflow.Context, input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error)
    GetCelebrityInfoAsync(ctx workflow.Context, input *rekognition.GetCelebrityInfoInput) *RekognitionGetCelebrityInfoResult

    GetCelebrityRecognition(ctx workflow.Context, input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error)
    GetCelebrityRecognitionAsync(ctx workflow.Context, input *rekognition.GetCelebrityRecognitionInput) *RekognitionGetCelebrityRecognitionResult

    GetContentModeration(ctx workflow.Context, input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error)
    GetContentModerationAsync(ctx workflow.Context, input *rekognition.GetContentModerationInput) *RekognitionGetContentModerationResult

    GetFaceDetection(ctx workflow.Context, input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error)
    GetFaceDetectionAsync(ctx workflow.Context, input *rekognition.GetFaceDetectionInput) *RekognitionGetFaceDetectionResult

    GetFaceSearch(ctx workflow.Context, input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error)
    GetFaceSearchAsync(ctx workflow.Context, input *rekognition.GetFaceSearchInput) *RekognitionGetFaceSearchResult

    GetLabelDetection(ctx workflow.Context, input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error)
    GetLabelDetectionAsync(ctx workflow.Context, input *rekognition.GetLabelDetectionInput) *RekognitionGetLabelDetectionResult

    GetPersonTracking(ctx workflow.Context, input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error)
    GetPersonTrackingAsync(ctx workflow.Context, input *rekognition.GetPersonTrackingInput) *RekognitionGetPersonTrackingResult

    GetSegmentDetection(ctx workflow.Context, input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error)
    GetSegmentDetectionAsync(ctx workflow.Context, input *rekognition.GetSegmentDetectionInput) *RekognitionGetSegmentDetectionResult

    GetTextDetection(ctx workflow.Context, input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error)
    GetTextDetectionAsync(ctx workflow.Context, input *rekognition.GetTextDetectionInput) *RekognitionGetTextDetectionResult

    IndexFaces(ctx workflow.Context, input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error)
    IndexFacesAsync(ctx workflow.Context, input *rekognition.IndexFacesInput) *RekognitionIndexFacesResult

    ListCollections(ctx workflow.Context, input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error)
    ListCollectionsAsync(ctx workflow.Context, input *rekognition.ListCollectionsInput) *RekognitionListCollectionsResult

    ListFaces(ctx workflow.Context, input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error)
    ListFacesAsync(ctx workflow.Context, input *rekognition.ListFacesInput) *RekognitionListFacesResult

    ListStreamProcessors(ctx workflow.Context, input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error)
    ListStreamProcessorsAsync(ctx workflow.Context, input *rekognition.ListStreamProcessorsInput) *RekognitionListStreamProcessorsResult

    RecognizeCelebrities(ctx workflow.Context, input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error)
    RecognizeCelebritiesAsync(ctx workflow.Context, input *rekognition.RecognizeCelebritiesInput) *RekognitionRecognizeCelebritiesResult

    SearchFaces(ctx workflow.Context, input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error)
    SearchFacesAsync(ctx workflow.Context, input *rekognition.SearchFacesInput) *RekognitionSearchFacesResult

    SearchFacesByImage(ctx workflow.Context, input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error)
    SearchFacesByImageAsync(ctx workflow.Context, input *rekognition.SearchFacesByImageInput) *RekognitionSearchFacesByImageResult

    StartCelebrityRecognition(ctx workflow.Context, input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error)
    StartCelebrityRecognitionAsync(ctx workflow.Context, input *rekognition.StartCelebrityRecognitionInput) *RekognitionStartCelebrityRecognitionResult

    StartContentModeration(ctx workflow.Context, input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error)
    StartContentModerationAsync(ctx workflow.Context, input *rekognition.StartContentModerationInput) *RekognitionStartContentModerationResult

    StartFaceDetection(ctx workflow.Context, input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error)
    StartFaceDetectionAsync(ctx workflow.Context, input *rekognition.StartFaceDetectionInput) *RekognitionStartFaceDetectionResult

    StartFaceSearch(ctx workflow.Context, input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error)
    StartFaceSearchAsync(ctx workflow.Context, input *rekognition.StartFaceSearchInput) *RekognitionStartFaceSearchResult

    StartLabelDetection(ctx workflow.Context, input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error)
    StartLabelDetectionAsync(ctx workflow.Context, input *rekognition.StartLabelDetectionInput) *RekognitionStartLabelDetectionResult

    StartPersonTracking(ctx workflow.Context, input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error)
    StartPersonTrackingAsync(ctx workflow.Context, input *rekognition.StartPersonTrackingInput) *RekognitionStartPersonTrackingResult

    StartProjectVersion(ctx workflow.Context, input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error)
    StartProjectVersionAsync(ctx workflow.Context, input *rekognition.StartProjectVersionInput) *RekognitionStartProjectVersionResult

    StartSegmentDetection(ctx workflow.Context, input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error)
    StartSegmentDetectionAsync(ctx workflow.Context, input *rekognition.StartSegmentDetectionInput) *RekognitionStartSegmentDetectionResult

    StartStreamProcessor(ctx workflow.Context, input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error)
    StartStreamProcessorAsync(ctx workflow.Context, input *rekognition.StartStreamProcessorInput) *RekognitionStartStreamProcessorResult

    StartTextDetection(ctx workflow.Context, input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error)
    StartTextDetectionAsync(ctx workflow.Context, input *rekognition.StartTextDetectionInput) *RekognitionStartTextDetectionResult

    StopProjectVersion(ctx workflow.Context, input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error)
    StopProjectVersionAsync(ctx workflow.Context, input *rekognition.StopProjectVersionInput) *RekognitionStopProjectVersionResult

    StopStreamProcessor(ctx workflow.Context, input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error)
    StopStreamProcessorAsync(ctx workflow.Context, input *rekognition.StopStreamProcessorInput) *RekognitionStopStreamProcessorResult

    WaitUntilProjectVersionRunning(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) error
    WaitUntilProjectVersionTrainingCompleted(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) error}
type RekognitionCompareFacesResult struct {
	Result workflow.Future
}

func (r *RekognitionCompareFacesResult) Get(ctx workflow.Context) (*rekognition.CompareFacesOutput, error) {
    var output rekognition.CompareFacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionCreateCollectionResult struct {
	Result workflow.Future
}

func (r *RekognitionCreateCollectionResult) Get(ctx workflow.Context) (*rekognition.CreateCollectionOutput, error) {
    var output rekognition.CreateCollectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionCreateProjectResult struct {
	Result workflow.Future
}

func (r *RekognitionCreateProjectResult) Get(ctx workflow.Context) (*rekognition.CreateProjectOutput, error) {
    var output rekognition.CreateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionCreateProjectVersionResult struct {
	Result workflow.Future
}

func (r *RekognitionCreateProjectVersionResult) Get(ctx workflow.Context) (*rekognition.CreateProjectVersionOutput, error) {
    var output rekognition.CreateProjectVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionCreateStreamProcessorResult struct {
	Result workflow.Future
}

func (r *RekognitionCreateStreamProcessorResult) Get(ctx workflow.Context) (*rekognition.CreateStreamProcessorOutput, error) {
    var output rekognition.CreateStreamProcessorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDeleteCollectionResult struct {
	Result workflow.Future
}

func (r *RekognitionDeleteCollectionResult) Get(ctx workflow.Context) (*rekognition.DeleteCollectionOutput, error) {
    var output rekognition.DeleteCollectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDeleteFacesResult struct {
	Result workflow.Future
}

func (r *RekognitionDeleteFacesResult) Get(ctx workflow.Context) (*rekognition.DeleteFacesOutput, error) {
    var output rekognition.DeleteFacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDeleteProjectResult struct {
	Result workflow.Future
}

func (r *RekognitionDeleteProjectResult) Get(ctx workflow.Context) (*rekognition.DeleteProjectOutput, error) {
    var output rekognition.DeleteProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDeleteProjectVersionResult struct {
	Result workflow.Future
}

func (r *RekognitionDeleteProjectVersionResult) Get(ctx workflow.Context) (*rekognition.DeleteProjectVersionOutput, error) {
    var output rekognition.DeleteProjectVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDeleteStreamProcessorResult struct {
	Result workflow.Future
}

func (r *RekognitionDeleteStreamProcessorResult) Get(ctx workflow.Context) (*rekognition.DeleteStreamProcessorOutput, error) {
    var output rekognition.DeleteStreamProcessorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDescribeCollectionResult struct {
	Result workflow.Future
}

func (r *RekognitionDescribeCollectionResult) Get(ctx workflow.Context) (*rekognition.DescribeCollectionOutput, error) {
    var output rekognition.DescribeCollectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDescribeProjectVersionsResult struct {
	Result workflow.Future
}

func (r *RekognitionDescribeProjectVersionsResult) Get(ctx workflow.Context) (*rekognition.DescribeProjectVersionsOutput, error) {
    var output rekognition.DescribeProjectVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDescribeProjectsResult struct {
	Result workflow.Future
}

func (r *RekognitionDescribeProjectsResult) Get(ctx workflow.Context) (*rekognition.DescribeProjectsOutput, error) {
    var output rekognition.DescribeProjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDescribeStreamProcessorResult struct {
	Result workflow.Future
}

func (r *RekognitionDescribeStreamProcessorResult) Get(ctx workflow.Context) (*rekognition.DescribeStreamProcessorOutput, error) {
    var output rekognition.DescribeStreamProcessorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDetectCustomLabelsResult struct {
	Result workflow.Future
}

func (r *RekognitionDetectCustomLabelsResult) Get(ctx workflow.Context) (*rekognition.DetectCustomLabelsOutput, error) {
    var output rekognition.DetectCustomLabelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDetectFacesResult struct {
	Result workflow.Future
}

func (r *RekognitionDetectFacesResult) Get(ctx workflow.Context) (*rekognition.DetectFacesOutput, error) {
    var output rekognition.DetectFacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDetectLabelsResult struct {
	Result workflow.Future
}

func (r *RekognitionDetectLabelsResult) Get(ctx workflow.Context) (*rekognition.DetectLabelsOutput, error) {
    var output rekognition.DetectLabelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDetectModerationLabelsResult struct {
	Result workflow.Future
}

func (r *RekognitionDetectModerationLabelsResult) Get(ctx workflow.Context) (*rekognition.DetectModerationLabelsOutput, error) {
    var output rekognition.DetectModerationLabelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionDetectTextResult struct {
	Result workflow.Future
}

func (r *RekognitionDetectTextResult) Get(ctx workflow.Context) (*rekognition.DetectTextOutput, error) {
    var output rekognition.DetectTextOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetCelebrityInfoResult struct {
	Result workflow.Future
}

func (r *RekognitionGetCelebrityInfoResult) Get(ctx workflow.Context) (*rekognition.GetCelebrityInfoOutput, error) {
    var output rekognition.GetCelebrityInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetCelebrityRecognitionResult struct {
	Result workflow.Future
}

func (r *RekognitionGetCelebrityRecognitionResult) Get(ctx workflow.Context) (*rekognition.GetCelebrityRecognitionOutput, error) {
    var output rekognition.GetCelebrityRecognitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetContentModerationResult struct {
	Result workflow.Future
}

func (r *RekognitionGetContentModerationResult) Get(ctx workflow.Context) (*rekognition.GetContentModerationOutput, error) {
    var output rekognition.GetContentModerationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetFaceDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionGetFaceDetectionResult) Get(ctx workflow.Context) (*rekognition.GetFaceDetectionOutput, error) {
    var output rekognition.GetFaceDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetFaceSearchResult struct {
	Result workflow.Future
}

func (r *RekognitionGetFaceSearchResult) Get(ctx workflow.Context) (*rekognition.GetFaceSearchOutput, error) {
    var output rekognition.GetFaceSearchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetLabelDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionGetLabelDetectionResult) Get(ctx workflow.Context) (*rekognition.GetLabelDetectionOutput, error) {
    var output rekognition.GetLabelDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetPersonTrackingResult struct {
	Result workflow.Future
}

func (r *RekognitionGetPersonTrackingResult) Get(ctx workflow.Context) (*rekognition.GetPersonTrackingOutput, error) {
    var output rekognition.GetPersonTrackingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetSegmentDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionGetSegmentDetectionResult) Get(ctx workflow.Context) (*rekognition.GetSegmentDetectionOutput, error) {
    var output rekognition.GetSegmentDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionGetTextDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionGetTextDetectionResult) Get(ctx workflow.Context) (*rekognition.GetTextDetectionOutput, error) {
    var output rekognition.GetTextDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionIndexFacesResult struct {
	Result workflow.Future
}

func (r *RekognitionIndexFacesResult) Get(ctx workflow.Context) (*rekognition.IndexFacesOutput, error) {
    var output rekognition.IndexFacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionListCollectionsResult struct {
	Result workflow.Future
}

func (r *RekognitionListCollectionsResult) Get(ctx workflow.Context) (*rekognition.ListCollectionsOutput, error) {
    var output rekognition.ListCollectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionListFacesResult struct {
	Result workflow.Future
}

func (r *RekognitionListFacesResult) Get(ctx workflow.Context) (*rekognition.ListFacesOutput, error) {
    var output rekognition.ListFacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionListStreamProcessorsResult struct {
	Result workflow.Future
}

func (r *RekognitionListStreamProcessorsResult) Get(ctx workflow.Context) (*rekognition.ListStreamProcessorsOutput, error) {
    var output rekognition.ListStreamProcessorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionRecognizeCelebritiesResult struct {
	Result workflow.Future
}

func (r *RekognitionRecognizeCelebritiesResult) Get(ctx workflow.Context) (*rekognition.RecognizeCelebritiesOutput, error) {
    var output rekognition.RecognizeCelebritiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionSearchFacesResult struct {
	Result workflow.Future
}

func (r *RekognitionSearchFacesResult) Get(ctx workflow.Context) (*rekognition.SearchFacesOutput, error) {
    var output rekognition.SearchFacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionSearchFacesByImageResult struct {
	Result workflow.Future
}

func (r *RekognitionSearchFacesByImageResult) Get(ctx workflow.Context) (*rekognition.SearchFacesByImageOutput, error) {
    var output rekognition.SearchFacesByImageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartCelebrityRecognitionResult struct {
	Result workflow.Future
}

func (r *RekognitionStartCelebrityRecognitionResult) Get(ctx workflow.Context) (*rekognition.StartCelebrityRecognitionOutput, error) {
    var output rekognition.StartCelebrityRecognitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartContentModerationResult struct {
	Result workflow.Future
}

func (r *RekognitionStartContentModerationResult) Get(ctx workflow.Context) (*rekognition.StartContentModerationOutput, error) {
    var output rekognition.StartContentModerationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartFaceDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionStartFaceDetectionResult) Get(ctx workflow.Context) (*rekognition.StartFaceDetectionOutput, error) {
    var output rekognition.StartFaceDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartFaceSearchResult struct {
	Result workflow.Future
}

func (r *RekognitionStartFaceSearchResult) Get(ctx workflow.Context) (*rekognition.StartFaceSearchOutput, error) {
    var output rekognition.StartFaceSearchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartLabelDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionStartLabelDetectionResult) Get(ctx workflow.Context) (*rekognition.StartLabelDetectionOutput, error) {
    var output rekognition.StartLabelDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartPersonTrackingResult struct {
	Result workflow.Future
}

func (r *RekognitionStartPersonTrackingResult) Get(ctx workflow.Context) (*rekognition.StartPersonTrackingOutput, error) {
    var output rekognition.StartPersonTrackingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartProjectVersionResult struct {
	Result workflow.Future
}

func (r *RekognitionStartProjectVersionResult) Get(ctx workflow.Context) (*rekognition.StartProjectVersionOutput, error) {
    var output rekognition.StartProjectVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartSegmentDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionStartSegmentDetectionResult) Get(ctx workflow.Context) (*rekognition.StartSegmentDetectionOutput, error) {
    var output rekognition.StartSegmentDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartStreamProcessorResult struct {
	Result workflow.Future
}

func (r *RekognitionStartStreamProcessorResult) Get(ctx workflow.Context) (*rekognition.StartStreamProcessorOutput, error) {
    var output rekognition.StartStreamProcessorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStartTextDetectionResult struct {
	Result workflow.Future
}

func (r *RekognitionStartTextDetectionResult) Get(ctx workflow.Context) (*rekognition.StartTextDetectionOutput, error) {
    var output rekognition.StartTextDetectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStopProjectVersionResult struct {
	Result workflow.Future
}

func (r *RekognitionStopProjectVersionResult) Get(ctx workflow.Context) (*rekognition.StopProjectVersionOutput, error) {
    var output rekognition.StopProjectVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RekognitionStopStreamProcessorResult struct {
	Result workflow.Future
}

func (r *RekognitionStopStreamProcessorResult) Get(ctx workflow.Context) (*rekognition.StopStreamProcessorOutput, error) {
    var output rekognition.StopStreamProcessorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type RekognitionStub struct {
    activities RekognitionClient
}

func NewRekognitionStub() RekognitionClient {
    return &RekognitionStub{}
}

func (a *RekognitionStub) CompareFaces(ctx workflow.Context, input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
    var output rekognition.CompareFacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CompareFaces, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) CreateCollection(ctx workflow.Context, input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error) {
    var output rekognition.CreateCollectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCollection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) CreateProject(ctx workflow.Context, input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error) {
    var output rekognition.CreateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) CreateProjectVersion(ctx workflow.Context, input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error) {
    var output rekognition.CreateProjectVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProjectVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) CreateStreamProcessor(ctx workflow.Context, input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error) {
    var output rekognition.CreateStreamProcessorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStreamProcessor, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DeleteCollection(ctx workflow.Context, input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error) {
    var output rekognition.DeleteCollectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCollection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DeleteFaces(ctx workflow.Context, input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error) {
    var output rekognition.DeleteFacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFaces, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DeleteProject(ctx workflow.Context, input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error) {
    var output rekognition.DeleteProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DeleteProjectVersion(ctx workflow.Context, input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error) {
    var output rekognition.DeleteProjectVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProjectVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DeleteStreamProcessor(ctx workflow.Context, input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error) {
    var output rekognition.DeleteStreamProcessorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteStreamProcessor, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DescribeCollection(ctx workflow.Context, input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error) {
    var output rekognition.DescribeCollectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCollection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DescribeProjectVersions(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error) {
    var output rekognition.DescribeProjectVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProjectVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DescribeProjects(ctx workflow.Context, input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error) {
    var output rekognition.DescribeProjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProjects, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DescribeStreamProcessor(ctx workflow.Context, input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error) {
    var output rekognition.DescribeStreamProcessorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeStreamProcessor, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DetectCustomLabels(ctx workflow.Context, input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error) {
    var output rekognition.DetectCustomLabelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectCustomLabels, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DetectFaces(ctx workflow.Context, input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
    var output rekognition.DetectFacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectFaces, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DetectLabels(ctx workflow.Context, input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
    var output rekognition.DetectLabelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectLabels, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DetectModerationLabels(ctx workflow.Context, input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error) {
    var output rekognition.DetectModerationLabelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectModerationLabels, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) DetectText(ctx workflow.Context, input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
    var output rekognition.DetectTextOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetectText, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetCelebrityInfo(ctx workflow.Context, input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error) {
    var output rekognition.GetCelebrityInfoOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCelebrityInfo, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetCelebrityRecognition(ctx workflow.Context, input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error) {
    var output rekognition.GetCelebrityRecognitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCelebrityRecognition, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetContentModeration(ctx workflow.Context, input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error) {
    var output rekognition.GetContentModerationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetContentModeration, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetFaceDetection(ctx workflow.Context, input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error) {
    var output rekognition.GetFaceDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFaceDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetFaceSearch(ctx workflow.Context, input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error) {
    var output rekognition.GetFaceSearchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFaceSearch, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetLabelDetection(ctx workflow.Context, input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error) {
    var output rekognition.GetLabelDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLabelDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetPersonTracking(ctx workflow.Context, input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error) {
    var output rekognition.GetPersonTrackingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPersonTracking, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetSegmentDetection(ctx workflow.Context, input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error) {
    var output rekognition.GetSegmentDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSegmentDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) GetTextDetection(ctx workflow.Context, input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error) {
    var output rekognition.GetTextDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTextDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) IndexFaces(ctx workflow.Context, input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error) {
    var output rekognition.IndexFacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.IndexFaces, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) ListCollections(ctx workflow.Context, input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error) {
    var output rekognition.ListCollectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCollections, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) ListFaces(ctx workflow.Context, input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error) {
    var output rekognition.ListFacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFaces, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) ListStreamProcessors(ctx workflow.Context, input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error) {
    var output rekognition.ListStreamProcessorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStreamProcessors, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) RecognizeCelebrities(ctx workflow.Context, input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error) {
    var output rekognition.RecognizeCelebritiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RecognizeCelebrities, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) SearchFaces(ctx workflow.Context, input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error) {
    var output rekognition.SearchFacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchFaces, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) SearchFacesByImage(ctx workflow.Context, input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error) {
    var output rekognition.SearchFacesByImageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchFacesByImage, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartCelebrityRecognition(ctx workflow.Context, input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error) {
    var output rekognition.StartCelebrityRecognitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartCelebrityRecognition, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartContentModeration(ctx workflow.Context, input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error) {
    var output rekognition.StartContentModerationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartContentModeration, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartFaceDetection(ctx workflow.Context, input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error) {
    var output rekognition.StartFaceDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartFaceDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartFaceSearch(ctx workflow.Context, input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error) {
    var output rekognition.StartFaceSearchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartFaceSearch, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartLabelDetection(ctx workflow.Context, input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error) {
    var output rekognition.StartLabelDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartLabelDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartPersonTracking(ctx workflow.Context, input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error) {
    var output rekognition.StartPersonTrackingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartPersonTracking, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartProjectVersion(ctx workflow.Context, input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error) {
    var output rekognition.StartProjectVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartProjectVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartSegmentDetection(ctx workflow.Context, input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error) {
    var output rekognition.StartSegmentDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSegmentDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartStreamProcessor(ctx workflow.Context, input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error) {
    var output rekognition.StartStreamProcessorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartStreamProcessor, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StartTextDetection(ctx workflow.Context, input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error) {
    var output rekognition.StartTextDetectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartTextDetection, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StopProjectVersion(ctx workflow.Context, input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error) {
    var output rekognition.StopProjectVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopProjectVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *RekognitionStub) StopStreamProcessor(ctx workflow.Context, input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error) {
    var output rekognition.StopStreamProcessorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopStreamProcessor, input).Get(ctx, &output)
    return &output, err
}


func (a *RekognitionStub) WaitUntilProjectVersionRunning(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) error {
    return a.client.WaitUntilProjectVersionRunning(input)
}


func (a *RekognitionStub) WaitUntilProjectVersionTrainingCompleted(ctx workflow.Context, input *rekognition.DescribeProjectVersionsInput) error {
    return a.client.WaitUntilProjectVersionTrainingCompleted(input)
}
