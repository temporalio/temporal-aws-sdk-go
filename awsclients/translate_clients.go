package awsclients

import (
	"github.com/aws/aws-sdk-go/service/translate"
	"go.temporal.io/sdk/workflow"
)

type TranslateClient interface {
       DeleteTerminology(ctx workflow.Context, input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error)
       DeleteTerminologyAsync(ctx workflow.Context, input *translate.DeleteTerminologyInput) *TranslateDeleteTerminologyResult

       DescribeTextTranslationJob(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error)
       DescribeTextTranslationJobAsync(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) *TranslateDescribeTextTranslationJobResult

       GetTerminology(ctx workflow.Context, input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error)
       GetTerminologyAsync(ctx workflow.Context, input *translate.GetTerminologyInput) *TranslateGetTerminologyResult

       ImportTerminology(ctx workflow.Context, input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error)
       ImportTerminologyAsync(ctx workflow.Context, input *translate.ImportTerminologyInput) *TranslateImportTerminologyResult

       ListTerminologies(ctx workflow.Context, input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error)
       ListTerminologiesAsync(ctx workflow.Context, input *translate.ListTerminologiesInput) *TranslateListTerminologiesResult

       ListTextTranslationJobs(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error)
       ListTextTranslationJobsAsync(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) *TranslateListTextTranslationJobsResult

       StartTextTranslationJob(ctx workflow.Context, input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error)
       StartTextTranslationJobAsync(ctx workflow.Context, input *translate.StartTextTranslationJobInput) *TranslateStartTextTranslationJobResult

       StopTextTranslationJob(ctx workflow.Context, input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error)
       StopTextTranslationJobAsync(ctx workflow.Context, input *translate.StopTextTranslationJobInput) *TranslateStopTextTranslationJobResult

       Text(ctx workflow.Context, input *translate.TextInput) (*translate.TextOutput, error)
       TextAsync(ctx workflow.Context, input *translate.TextInput) *TranslateTextResult
}

type TranslateStub struct {
}

func NewTranslateStub() TranslateClient {
    return &TranslateStub{}
}

type TranslateDeleteTerminologyResult struct {
	Result workflow.Future
}

func (r *TranslateDeleteTerminologyResult) Get(ctx workflow.Context) (*translate.DeleteTerminologyOutput, error) {
    var output translate.DeleteTerminologyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateDescribeTextTranslationJobResult struct {
	Result workflow.Future
}

func (r *TranslateDescribeTextTranslationJobResult) Get(ctx workflow.Context) (*translate.DescribeTextTranslationJobOutput, error) {
    var output translate.DescribeTextTranslationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateGetTerminologyResult struct {
	Result workflow.Future
}

func (r *TranslateGetTerminologyResult) Get(ctx workflow.Context) (*translate.GetTerminologyOutput, error) {
    var output translate.GetTerminologyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateImportTerminologyResult struct {
	Result workflow.Future
}

func (r *TranslateImportTerminologyResult) Get(ctx workflow.Context) (*translate.ImportTerminologyOutput, error) {
    var output translate.ImportTerminologyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateListTerminologiesResult struct {
	Result workflow.Future
}

func (r *TranslateListTerminologiesResult) Get(ctx workflow.Context) (*translate.ListTerminologiesOutput, error) {
    var output translate.ListTerminologiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateListTextTranslationJobsResult struct {
	Result workflow.Future
}

func (r *TranslateListTextTranslationJobsResult) Get(ctx workflow.Context) (*translate.ListTextTranslationJobsOutput, error) {
    var output translate.ListTextTranslationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateStartTextTranslationJobResult struct {
	Result workflow.Future
}

func (r *TranslateStartTextTranslationJobResult) Get(ctx workflow.Context) (*translate.StartTextTranslationJobOutput, error) {
    var output translate.StartTextTranslationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateStopTextTranslationJobResult struct {
	Result workflow.Future
}

func (r *TranslateStopTextTranslationJobResult) Get(ctx workflow.Context) (*translate.StopTextTranslationJobOutput, error) {
    var output translate.StopTextTranslationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type TranslateTextResult struct {
	Result workflow.Future
}

func (r *TranslateTextResult) Get(ctx workflow.Context) (*translate.TextOutput, error) {
    var output translate.TextOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) DeleteTerminology(ctx workflow.Context, input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
    var output translate.DeleteTerminologyOutput
    err := workflow.ExecuteActivity(ctx, "Translate.DeleteTerminology", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) DeleteTerminologyAsync(ctx workflow.Context, input *translate.DeleteTerminologyInput) *TranslateDeleteTerminologyResult {
    future := workflow.ExecuteActivity(ctx, "Translate.DeleteTerminology", input)
    return &TranslateDeleteTerminologyResult{Result: future}
}

func (a *TranslateStub) DescribeTextTranslationJob(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error) {
    var output translate.DescribeTextTranslationJobOutput
    err := workflow.ExecuteActivity(ctx, "Translate.DescribeTextTranslationJob", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) DescribeTextTranslationJobAsync(ctx workflow.Context, input *translate.DescribeTextTranslationJobInput) *TranslateDescribeTextTranslationJobResult {
    future := workflow.ExecuteActivity(ctx, "Translate.DescribeTextTranslationJob", input)
    return &TranslateDescribeTextTranslationJobResult{Result: future}
}

func (a *TranslateStub) GetTerminology(ctx workflow.Context, input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error) {
    var output translate.GetTerminologyOutput
    err := workflow.ExecuteActivity(ctx, "Translate.GetTerminology", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) GetTerminologyAsync(ctx workflow.Context, input *translate.GetTerminologyInput) *TranslateGetTerminologyResult {
    future := workflow.ExecuteActivity(ctx, "Translate.GetTerminology", input)
    return &TranslateGetTerminologyResult{Result: future}
}

func (a *TranslateStub) ImportTerminology(ctx workflow.Context, input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error) {
    var output translate.ImportTerminologyOutput
    err := workflow.ExecuteActivity(ctx, "Translate.ImportTerminology", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) ImportTerminologyAsync(ctx workflow.Context, input *translate.ImportTerminologyInput) *TranslateImportTerminologyResult {
    future := workflow.ExecuteActivity(ctx, "Translate.ImportTerminology", input)
    return &TranslateImportTerminologyResult{Result: future}
}

func (a *TranslateStub) ListTerminologies(ctx workflow.Context, input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error) {
    var output translate.ListTerminologiesOutput
    err := workflow.ExecuteActivity(ctx, "Translate.ListTerminologies", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) ListTerminologiesAsync(ctx workflow.Context, input *translate.ListTerminologiesInput) *TranslateListTerminologiesResult {
    future := workflow.ExecuteActivity(ctx, "Translate.ListTerminologies", input)
    return &TranslateListTerminologiesResult{Result: future}
}

func (a *TranslateStub) ListTextTranslationJobs(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error) {
    var output translate.ListTextTranslationJobsOutput
    err := workflow.ExecuteActivity(ctx, "Translate.ListTextTranslationJobs", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) ListTextTranslationJobsAsync(ctx workflow.Context, input *translate.ListTextTranslationJobsInput) *TranslateListTextTranslationJobsResult {
    future := workflow.ExecuteActivity(ctx, "Translate.ListTextTranslationJobs", input)
    return &TranslateListTextTranslationJobsResult{Result: future}
}

func (a *TranslateStub) StartTextTranslationJob(ctx workflow.Context, input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error) {
    var output translate.StartTextTranslationJobOutput
    err := workflow.ExecuteActivity(ctx, "Translate.StartTextTranslationJob", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) StartTextTranslationJobAsync(ctx workflow.Context, input *translate.StartTextTranslationJobInput) *TranslateStartTextTranslationJobResult {
    future := workflow.ExecuteActivity(ctx, "Translate.StartTextTranslationJob", input)
    return &TranslateStartTextTranslationJobResult{Result: future}
}

func (a *TranslateStub) StopTextTranslationJob(ctx workflow.Context, input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error) {
    var output translate.StopTextTranslationJobOutput
    err := workflow.ExecuteActivity(ctx, "Translate.StopTextTranslationJob", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) StopTextTranslationJobAsync(ctx workflow.Context, input *translate.StopTextTranslationJobInput) *TranslateStopTextTranslationJobResult {
    future := workflow.ExecuteActivity(ctx, "Translate.StopTextTranslationJob", input)
    return &TranslateStopTextTranslationJobResult{Result: future}
}

func (a *TranslateStub) Text(ctx workflow.Context, input *translate.TextInput) (*translate.TextOutput, error) {
    var output translate.TextOutput
    err := workflow.ExecuteActivity(ctx, "Translate.Text", input).Get(ctx, &output)
    return &output, err
}

func (a *TranslateStub) TextAsync(ctx workflow.Context, input *translate.TextInput) *TranslateTextResult {
    future := workflow.ExecuteActivity(ctx, "Translate.Text", input)
    return &TranslateTextResult{Result: future}
}
