package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CloudTrailClient interface {
    AddTags(ctx workflow.Context, input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error)
    AddTagsAsync(ctx workflow.Context, input *cloudtrail.AddTagsInput) *CloudtrailAddTagsResult

    CreateTrail(ctx workflow.Context, input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)
    CreateTrailAsync(ctx workflow.Context, input *cloudtrail.CreateTrailInput) *CloudtrailCreateTrailResult

    DeleteTrail(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)
    DeleteTrailAsync(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) *CloudtrailDeleteTrailResult

    DescribeTrails(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)
    DescribeTrailsAsync(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) *CloudtrailDescribeTrailsResult

    GetEventSelectors(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error)
    GetEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) *CloudtrailGetEventSelectorsResult

    GetInsightSelectors(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error)
    GetInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) *CloudtrailGetInsightSelectorsResult

    GetTrail(ctx workflow.Context, input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error)
    GetTrailAsync(ctx workflow.Context, input *cloudtrail.GetTrailInput) *CloudtrailGetTrailResult

    GetTrailStatus(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)
    GetTrailStatusAsync(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) *CloudtrailGetTrailStatusResult

    ListPublicKeys(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error)
    ListPublicKeysAsync(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) *CloudtrailListPublicKeysResult

    ListTags(ctx workflow.Context, input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *cloudtrail.ListTagsInput) *CloudtrailListTagsResult

    ListTrails(ctx workflow.Context, input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error)
    ListTrailsAsync(ctx workflow.Context, input *cloudtrail.ListTrailsInput) *CloudtrailListTrailsResult

    LookupEvents(ctx workflow.Context, input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)
    LookupEventsAsync(ctx workflow.Context, input *cloudtrail.LookupEventsInput) *CloudtrailLookupEventsResult

    PutEventSelectors(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error)
    PutEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) *CloudtrailPutEventSelectorsResult

    PutInsightSelectors(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error)
    PutInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) *CloudtrailPutInsightSelectorsResult

    RemoveTags(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error)
    RemoveTagsAsync(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) *CloudtrailRemoveTagsResult

    StartLogging(ctx workflow.Context, input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)
    StartLoggingAsync(ctx workflow.Context, input *cloudtrail.StartLoggingInput) *CloudtrailStartLoggingResult

    StopLogging(ctx workflow.Context, input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)
    StopLoggingAsync(ctx workflow.Context, input *cloudtrail.StopLoggingInput) *CloudtrailStopLoggingResult

    UpdateTrail(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
    UpdateTrailAsync(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) *CloudtrailUpdateTrailResult
}
type CloudtrailAddTagsResult struct {
	Result workflow.Future
}

func (r *CloudtrailAddTagsResult) Get(ctx workflow.Context) (*cloudtrail.AddTagsOutput, error) {
    var output cloudtrail.AddTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailCreateTrailResult struct {
	Result workflow.Future
}

func (r *CloudtrailCreateTrailResult) Get(ctx workflow.Context) (*cloudtrail.CreateTrailOutput, error) {
    var output cloudtrail.CreateTrailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailDeleteTrailResult struct {
	Result workflow.Future
}

func (r *CloudtrailDeleteTrailResult) Get(ctx workflow.Context) (*cloudtrail.DeleteTrailOutput, error) {
    var output cloudtrail.DeleteTrailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailDescribeTrailsResult struct {
	Result workflow.Future
}

func (r *CloudtrailDescribeTrailsResult) Get(ctx workflow.Context) (*cloudtrail.DescribeTrailsOutput, error) {
    var output cloudtrail.DescribeTrailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailGetEventSelectorsResult struct {
	Result workflow.Future
}

func (r *CloudtrailGetEventSelectorsResult) Get(ctx workflow.Context) (*cloudtrail.GetEventSelectorsOutput, error) {
    var output cloudtrail.GetEventSelectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailGetInsightSelectorsResult struct {
	Result workflow.Future
}

func (r *CloudtrailGetInsightSelectorsResult) Get(ctx workflow.Context) (*cloudtrail.GetInsightSelectorsOutput, error) {
    var output cloudtrail.GetInsightSelectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailGetTrailResult struct {
	Result workflow.Future
}

func (r *CloudtrailGetTrailResult) Get(ctx workflow.Context) (*cloudtrail.GetTrailOutput, error) {
    var output cloudtrail.GetTrailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailGetTrailStatusResult struct {
	Result workflow.Future
}

func (r *CloudtrailGetTrailStatusResult) Get(ctx workflow.Context) (*cloudtrail.GetTrailStatusOutput, error) {
    var output cloudtrail.GetTrailStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailListPublicKeysResult struct {
	Result workflow.Future
}

func (r *CloudtrailListPublicKeysResult) Get(ctx workflow.Context) (*cloudtrail.ListPublicKeysOutput, error) {
    var output cloudtrail.ListPublicKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailListTagsResult struct {
	Result workflow.Future
}

func (r *CloudtrailListTagsResult) Get(ctx workflow.Context) (*cloudtrail.ListTagsOutput, error) {
    var output cloudtrail.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailListTrailsResult struct {
	Result workflow.Future
}

func (r *CloudtrailListTrailsResult) Get(ctx workflow.Context) (*cloudtrail.ListTrailsOutput, error) {
    var output cloudtrail.ListTrailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailLookupEventsResult struct {
	Result workflow.Future
}

func (r *CloudtrailLookupEventsResult) Get(ctx workflow.Context) (*cloudtrail.LookupEventsOutput, error) {
    var output cloudtrail.LookupEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailPutEventSelectorsResult struct {
	Result workflow.Future
}

func (r *CloudtrailPutEventSelectorsResult) Get(ctx workflow.Context) (*cloudtrail.PutEventSelectorsOutput, error) {
    var output cloudtrail.PutEventSelectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailPutInsightSelectorsResult struct {
	Result workflow.Future
}

func (r *CloudtrailPutInsightSelectorsResult) Get(ctx workflow.Context) (*cloudtrail.PutInsightSelectorsOutput, error) {
    var output cloudtrail.PutInsightSelectorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailRemoveTagsResult struct {
	Result workflow.Future
}

func (r *CloudtrailRemoveTagsResult) Get(ctx workflow.Context) (*cloudtrail.RemoveTagsOutput, error) {
    var output cloudtrail.RemoveTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailStartLoggingResult struct {
	Result workflow.Future
}

func (r *CloudtrailStartLoggingResult) Get(ctx workflow.Context) (*cloudtrail.StartLoggingOutput, error) {
    var output cloudtrail.StartLoggingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailStopLoggingResult struct {
	Result workflow.Future
}

func (r *CloudtrailStopLoggingResult) Get(ctx workflow.Context) (*cloudtrail.StopLoggingOutput, error) {
    var output cloudtrail.StopLoggingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudtrailUpdateTrailResult struct {
	Result workflow.Future
}

func (r *CloudtrailUpdateTrailResult) Get(ctx workflow.Context) (*cloudtrail.UpdateTrailOutput, error) {
    var output cloudtrail.UpdateTrailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudTrailStub struct {
    activities awsactivities.CloudTrailActivities
}

func NewCloudTrailStub() CloudTrailClient {
    return &CloudTrailStub{}
}
func (a *CloudTrailStub) AddTags(ctx workflow.Context, input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
    var output cloudtrail.AddTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTags, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) AddTagsAsync(ctx workflow.Context, input *cloudtrail.AddTagsInput) *CloudtrailAddTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTags, input)
    return &CloudtrailAddTagsResult{Result: future}
}
func (a *CloudTrailStub) CreateTrail(ctx workflow.Context, input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error) {
    var output cloudtrail.CreateTrailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrail, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) CreateTrailAsync(ctx workflow.Context, input *cloudtrail.CreateTrailInput) *CloudtrailCreateTrailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrail, input)
    return &CloudtrailCreateTrailResult{Result: future}
}
func (a *CloudTrailStub) DeleteTrail(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error) {
    var output cloudtrail.DeleteTrailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrail, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) DeleteTrailAsync(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) *CloudtrailDeleteTrailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrail, input)
    return &CloudtrailDeleteTrailResult{Result: future}
}
func (a *CloudTrailStub) DescribeTrails(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error) {
    var output cloudtrail.DescribeTrailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrails, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) DescribeTrailsAsync(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) *CloudtrailDescribeTrailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrails, input)
    return &CloudtrailDescribeTrailsResult{Result: future}
}
func (a *CloudTrailStub) GetEventSelectors(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error) {
    var output cloudtrail.GetEventSelectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEventSelectors, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) GetEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) *CloudtrailGetEventSelectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetEventSelectors, input)
    return &CloudtrailGetEventSelectorsResult{Result: future}
}
func (a *CloudTrailStub) GetInsightSelectors(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error) {
    var output cloudtrail.GetInsightSelectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetInsightSelectors, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) GetInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) *CloudtrailGetInsightSelectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetInsightSelectors, input)
    return &CloudtrailGetInsightSelectorsResult{Result: future}
}
func (a *CloudTrailStub) GetTrail(ctx workflow.Context, input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error) {
    var output cloudtrail.GetTrailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTrail, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) GetTrailAsync(ctx workflow.Context, input *cloudtrail.GetTrailInput) *CloudtrailGetTrailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTrail, input)
    return &CloudtrailGetTrailResult{Result: future}
}
func (a *CloudTrailStub) GetTrailStatus(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error) {
    var output cloudtrail.GetTrailStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTrailStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) GetTrailStatusAsync(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) *CloudtrailGetTrailStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTrailStatus, input)
    return &CloudtrailGetTrailStatusResult{Result: future}
}
func (a *CloudTrailStub) ListPublicKeys(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error) {
    var output cloudtrail.ListPublicKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPublicKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) ListPublicKeysAsync(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) *CloudtrailListPublicKeysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPublicKeys, input)
    return &CloudtrailListPublicKeysResult{Result: future}
}
func (a *CloudTrailStub) ListTags(ctx workflow.Context, input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error) {
    var output cloudtrail.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) ListTagsAsync(ctx workflow.Context, input *cloudtrail.ListTagsInput) *CloudtrailListTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
    return &CloudtrailListTagsResult{Result: future}
}
func (a *CloudTrailStub) ListTrails(ctx workflow.Context, input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error) {
    var output cloudtrail.ListTrailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTrails, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) ListTrailsAsync(ctx workflow.Context, input *cloudtrail.ListTrailsInput) *CloudtrailListTrailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTrails, input)
    return &CloudtrailListTrailsResult{Result: future}
}
func (a *CloudTrailStub) LookupEvents(ctx workflow.Context, input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error) {
    var output cloudtrail.LookupEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.LookupEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) LookupEventsAsync(ctx workflow.Context, input *cloudtrail.LookupEventsInput) *CloudtrailLookupEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.LookupEvents, input)
    return &CloudtrailLookupEventsResult{Result: future}
}
func (a *CloudTrailStub) PutEventSelectors(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error) {
    var output cloudtrail.PutEventSelectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEventSelectors, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) PutEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) *CloudtrailPutEventSelectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutEventSelectors, input)
    return &CloudtrailPutEventSelectorsResult{Result: future}
}
func (a *CloudTrailStub) PutInsightSelectors(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error) {
    var output cloudtrail.PutInsightSelectorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutInsightSelectors, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) PutInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) *CloudtrailPutInsightSelectorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutInsightSelectors, input)
    return &CloudtrailPutInsightSelectorsResult{Result: future}
}
func (a *CloudTrailStub) RemoveTags(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error) {
    var output cloudtrail.RemoveTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTags, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) RemoveTagsAsync(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) *CloudtrailRemoveTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTags, input)
    return &CloudtrailRemoveTagsResult{Result: future}
}
func (a *CloudTrailStub) StartLogging(ctx workflow.Context, input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error) {
    var output cloudtrail.StartLoggingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartLogging, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) StartLoggingAsync(ctx workflow.Context, input *cloudtrail.StartLoggingInput) *CloudtrailStartLoggingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartLogging, input)
    return &CloudtrailStartLoggingResult{Result: future}
}
func (a *CloudTrailStub) StopLogging(ctx workflow.Context, input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error) {
    var output cloudtrail.StopLoggingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopLogging, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) StopLoggingAsync(ctx workflow.Context, input *cloudtrail.StopLoggingInput) *CloudtrailStopLoggingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopLogging, input)
    return &CloudtrailStopLoggingResult{Result: future}
}
func (a *CloudTrailStub) UpdateTrail(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error) {
    var output cloudtrail.UpdateTrailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTrail, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudTrailStub) UpdateTrailAsync(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) *CloudtrailUpdateTrailResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTrail, input)
    return &CloudtrailUpdateTrailResult{Result: future}
}