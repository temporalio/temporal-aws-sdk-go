package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type MediaStoreDataClient interface {
    DeleteObject(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error)
    DeleteObjectAsync(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) *MediastoredataDeleteObjectResult

    DescribeObject(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error)
    DescribeObjectAsync(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) *MediastoredataDescribeObjectResult

    GetObject(ctx workflow.Context, input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error)
    GetObjectAsync(ctx workflow.Context, input *mediastoredata.GetObjectInput) *MediastoredataGetObjectResult

    ListItems(ctx workflow.Context, input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error)
    ListItemsAsync(ctx workflow.Context, input *mediastoredata.ListItemsInput) *MediastoredataListItemsResult

    PutObject(ctx workflow.Context, input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error)
    PutObjectAsync(ctx workflow.Context, input *mediastoredata.PutObjectInput) *MediastoredataPutObjectResult
}
type MediastoredataDeleteObjectResult struct {
	Result workflow.Future
}

func (r *MediastoredataDeleteObjectResult) Get(ctx workflow.Context) (*mediastoredata.DeleteObjectOutput, error) {
    var output mediastoredata.DeleteObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediastoredataDescribeObjectResult struct {
	Result workflow.Future
}

func (r *MediastoredataDescribeObjectResult) Get(ctx workflow.Context) (*mediastoredata.DescribeObjectOutput, error) {
    var output mediastoredata.DescribeObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediastoredataGetObjectResult struct {
	Result workflow.Future
}

func (r *MediastoredataGetObjectResult) Get(ctx workflow.Context) (*mediastoredata.GetObjectOutput, error) {
    var output mediastoredata.GetObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediastoredataListItemsResult struct {
	Result workflow.Future
}

func (r *MediastoredataListItemsResult) Get(ctx workflow.Context) (*mediastoredata.ListItemsOutput, error) {
    var output mediastoredata.ListItemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MediastoredataPutObjectResult struct {
	Result workflow.Future
}

func (r *MediastoredataPutObjectResult) Get(ctx workflow.Context) (*mediastoredata.PutObjectOutput, error) {
    var output mediastoredata.PutObjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MediaStoreDataStub struct {
    activities awsactivities.MediaStoreDataActivities
}

func NewMediaStoreDataStub() MediaStoreDataClient {
    return &MediaStoreDataStub{}
}
func (a *MediaStoreDataStub) DeleteObject(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
    var output mediastoredata.DeleteObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteObject, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaStoreDataStub) DeleteObjectAsync(ctx workflow.Context, input *mediastoredata.DeleteObjectInput) *MediastoredataDeleteObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteObject, input)
    return &MediastoredataDeleteObjectResult{Result: future}
}
func (a *MediaStoreDataStub) DescribeObject(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error) {
    var output mediastoredata.DescribeObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeObject, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaStoreDataStub) DescribeObjectAsync(ctx workflow.Context, input *mediastoredata.DescribeObjectInput) *MediastoredataDescribeObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeObject, input)
    return &MediastoredataDescribeObjectResult{Result: future}
}
func (a *MediaStoreDataStub) GetObject(ctx workflow.Context, input *mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error) {
    var output mediastoredata.GetObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetObject, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaStoreDataStub) GetObjectAsync(ctx workflow.Context, input *mediastoredata.GetObjectInput) *MediastoredataGetObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetObject, input)
    return &MediastoredataGetObjectResult{Result: future}
}
func (a *MediaStoreDataStub) ListItems(ctx workflow.Context, input *mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error) {
    var output mediastoredata.ListItemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListItems, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaStoreDataStub) ListItemsAsync(ctx workflow.Context, input *mediastoredata.ListItemsInput) *MediastoredataListItemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListItems, input)
    return &MediastoredataListItemsResult{Result: future}
}
func (a *MediaStoreDataStub) PutObject(ctx workflow.Context, input *mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error) {
    var output mediastoredata.PutObjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutObject, input).Get(ctx, &output)
    return &output, err
}

func (a *MediaStoreDataStub) PutObjectAsync(ctx workflow.Context, input *mediastoredata.PutObjectInput) *MediastoredataPutObjectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutObject, input)
    return &MediastoredataPutObjectResult{Result: future}
}
