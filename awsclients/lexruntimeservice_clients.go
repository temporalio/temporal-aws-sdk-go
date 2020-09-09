package awsclients

import (
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type LexRuntimeServiceClient interface {
       DeleteSession(ctx workflow.Context, input *lexruntimeservice.DeleteSessionInput) (*lexruntimeservice.DeleteSessionOutput, error)
       DeleteSessionAsync(ctx workflow.Context, input *lexruntimeservice.DeleteSessionInput) *LexruntimeserviceDeleteSessionResult

       GetSession(ctx workflow.Context, input *lexruntimeservice.GetSessionInput) (*lexruntimeservice.GetSessionOutput, error)
       GetSessionAsync(ctx workflow.Context, input *lexruntimeservice.GetSessionInput) *LexruntimeserviceGetSessionResult

       PostContent(ctx workflow.Context, input *lexruntimeservice.PostContentInput) (*lexruntimeservice.PostContentOutput, error)
       PostContentAsync(ctx workflow.Context, input *lexruntimeservice.PostContentInput) *LexruntimeservicePostContentResult

       PostText(ctx workflow.Context, input *lexruntimeservice.PostTextInput) (*lexruntimeservice.PostTextOutput, error)
       PostTextAsync(ctx workflow.Context, input *lexruntimeservice.PostTextInput) *LexruntimeservicePostTextResult

       PutSession(ctx workflow.Context, input *lexruntimeservice.PutSessionInput) (*lexruntimeservice.PutSessionOutput, error)
       PutSessionAsync(ctx workflow.Context, input *lexruntimeservice.PutSessionInput) *LexruntimeservicePutSessionResult
}

type LexruntimeserviceDeleteSessionResult struct {
	Result workflow.Future
}

func (r *LexruntimeserviceDeleteSessionResult) Get(ctx workflow.Context) (*lexruntimeservice.DeleteSessionOutput, error) {
    var output lexruntimeservice.DeleteSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexruntimeserviceGetSessionResult struct {
	Result workflow.Future
}

func (r *LexruntimeserviceGetSessionResult) Get(ctx workflow.Context) (*lexruntimeservice.GetSessionOutput, error) {
    var output lexruntimeservice.GetSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexruntimeservicePostContentResult struct {
	Result workflow.Future
}

func (r *LexruntimeservicePostContentResult) Get(ctx workflow.Context) (*lexruntimeservice.PostContentOutput, error) {
    var output lexruntimeservice.PostContentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexruntimeservicePostTextResult struct {
	Result workflow.Future
}

func (r *LexruntimeservicePostTextResult) Get(ctx workflow.Context) (*lexruntimeservice.PostTextOutput, error) {
    var output lexruntimeservice.PostTextOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexruntimeservicePutSessionResult struct {
	Result workflow.Future
}

func (r *LexruntimeservicePutSessionResult) Get(ctx workflow.Context) (*lexruntimeservice.PutSessionOutput, error) {
    var output lexruntimeservice.PutSessionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LexRuntimeServiceStub struct {
    activities awsactivities.LexRuntimeServiceActivities
}

func NewLexRuntimeServiceStub() LexRuntimeServiceClient {
    return &LexRuntimeServiceStub{}
}

func (a *LexRuntimeServiceStub) DeleteSession(ctx workflow.Context, input *lexruntimeservice.DeleteSessionInput) (*lexruntimeservice.DeleteSessionOutput, error) {
    var output lexruntimeservice.DeleteSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSession, input).Get(ctx, &output)
    return &output, err
}

func (a *LexRuntimeServiceStub) DeleteSessionAsync(ctx workflow.Context, input *lexruntimeservice.DeleteSessionInput) *LexruntimeserviceDeleteSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSession, input)
    return &LexruntimeserviceDeleteSessionResult{Result: future}
}

func (a *LexRuntimeServiceStub) GetSession(ctx workflow.Context, input *lexruntimeservice.GetSessionInput) (*lexruntimeservice.GetSessionOutput, error) {
    var output lexruntimeservice.GetSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSession, input).Get(ctx, &output)
    return &output, err
}

func (a *LexRuntimeServiceStub) GetSessionAsync(ctx workflow.Context, input *lexruntimeservice.GetSessionInput) *LexruntimeserviceGetSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSession, input)
    return &LexruntimeserviceGetSessionResult{Result: future}
}

func (a *LexRuntimeServiceStub) PostContent(ctx workflow.Context, input *lexruntimeservice.PostContentInput) (*lexruntimeservice.PostContentOutput, error) {
    var output lexruntimeservice.PostContentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PostContent, input).Get(ctx, &output)
    return &output, err
}

func (a *LexRuntimeServiceStub) PostContentAsync(ctx workflow.Context, input *lexruntimeservice.PostContentInput) *LexruntimeservicePostContentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PostContent, input)
    return &LexruntimeservicePostContentResult{Result: future}
}

func (a *LexRuntimeServiceStub) PostText(ctx workflow.Context, input *lexruntimeservice.PostTextInput) (*lexruntimeservice.PostTextOutput, error) {
    var output lexruntimeservice.PostTextOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PostText, input).Get(ctx, &output)
    return &output, err
}

func (a *LexRuntimeServiceStub) PostTextAsync(ctx workflow.Context, input *lexruntimeservice.PostTextInput) *LexruntimeservicePostTextResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PostText, input)
    return &LexruntimeservicePostTextResult{Result: future}
}

func (a *LexRuntimeServiceStub) PutSession(ctx workflow.Context, input *lexruntimeservice.PutSessionInput) (*lexruntimeservice.PutSessionOutput, error) {
    var output lexruntimeservice.PutSessionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSession, input).Get(ctx, &output)
    return &output, err
}

func (a *LexRuntimeServiceStub) PutSessionAsync(ctx workflow.Context, input *lexruntimeservice.PutSessionInput) *LexruntimeservicePutSessionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutSession, input)
    return &LexruntimeservicePutSessionResult{Result: future}
}
