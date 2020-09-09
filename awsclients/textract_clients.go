package awsclients

import (
	"github.com/aws/aws-sdk-go/service/textract"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type TextractClient interface {
	AnalyzeDocument(ctx workflow.Context, input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error)
	AnalyzeDocumentAsync(ctx workflow.Context, input *textract.AnalyzeDocumentInput) *TextractAnalyzeDocumentResult

	DetectDocumentText(ctx workflow.Context, input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error)
	DetectDocumentTextAsync(ctx workflow.Context, input *textract.DetectDocumentTextInput) *TextractDetectDocumentTextResult

	GetDocumentAnalysis(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error)
	GetDocumentAnalysisAsync(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) *TextractGetDocumentAnalysisResult

	GetDocumentTextDetection(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error)
	GetDocumentTextDetectionAsync(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) *TextractGetDocumentTextDetectionResult

	StartDocumentAnalysis(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error)
	StartDocumentAnalysisAsync(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) *TextractStartDocumentAnalysisResult

	StartDocumentTextDetection(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error)
	StartDocumentTextDetectionAsync(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) *TextractStartDocumentTextDetectionResult
}

type TextractAnalyzeDocumentResult struct {
	Result workflow.Future
}

func (r *TextractAnalyzeDocumentResult) Get(ctx workflow.Context) (*textract.AnalyzeDocumentOutput, error) {
	var output textract.AnalyzeDocumentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TextractDetectDocumentTextResult struct {
	Result workflow.Future
}

func (r *TextractDetectDocumentTextResult) Get(ctx workflow.Context) (*textract.DetectDocumentTextOutput, error) {
	var output textract.DetectDocumentTextOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TextractGetDocumentAnalysisResult struct {
	Result workflow.Future
}

func (r *TextractGetDocumentAnalysisResult) Get(ctx workflow.Context) (*textract.GetDocumentAnalysisOutput, error) {
	var output textract.GetDocumentAnalysisOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TextractGetDocumentTextDetectionResult struct {
	Result workflow.Future
}

func (r *TextractGetDocumentTextDetectionResult) Get(ctx workflow.Context) (*textract.GetDocumentTextDetectionOutput, error) {
	var output textract.GetDocumentTextDetectionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TextractStartDocumentAnalysisResult struct {
	Result workflow.Future
}

func (r *TextractStartDocumentAnalysisResult) Get(ctx workflow.Context) (*textract.StartDocumentAnalysisOutput, error) {
	var output textract.StartDocumentAnalysisOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TextractStartDocumentTextDetectionResult struct {
	Result workflow.Future
}

func (r *TextractStartDocumentTextDetectionResult) Get(ctx workflow.Context) (*textract.StartDocumentTextDetectionOutput, error) {
	var output textract.StartDocumentTextDetectionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type TextractStub struct {
	activities awsactivities.TextractActivities
}

func NewTextractStub() TextractClient {
	return &TextractStub{}
}

func (a *TextractStub) AnalyzeDocument(ctx workflow.Context, input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
	var output textract.AnalyzeDocumentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AnalyzeDocument, input).Get(ctx, &output)
	return &output, err
}

func (a *TextractStub) AnalyzeDocumentAsync(ctx workflow.Context, input *textract.AnalyzeDocumentInput) *TextractAnalyzeDocumentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AnalyzeDocument, input)
	return &TextractAnalyzeDocumentResult{Result: future}
}

func (a *TextractStub) DetectDocumentText(ctx workflow.Context, input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error) {
	var output textract.DetectDocumentTextOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DetectDocumentText, input).Get(ctx, &output)
	return &output, err
}

func (a *TextractStub) DetectDocumentTextAsync(ctx workflow.Context, input *textract.DetectDocumentTextInput) *TextractDetectDocumentTextResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DetectDocumentText, input)
	return &TextractDetectDocumentTextResult{Result: future}
}

func (a *TextractStub) GetDocumentAnalysis(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error) {
	var output textract.GetDocumentAnalysisOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDocumentAnalysis, input).Get(ctx, &output)
	return &output, err
}

func (a *TextractStub) GetDocumentAnalysisAsync(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) *TextractGetDocumentAnalysisResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDocumentAnalysis, input)
	return &TextractGetDocumentAnalysisResult{Result: future}
}

func (a *TextractStub) GetDocumentTextDetection(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error) {
	var output textract.GetDocumentTextDetectionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDocumentTextDetection, input).Get(ctx, &output)
	return &output, err
}

func (a *TextractStub) GetDocumentTextDetectionAsync(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) *TextractGetDocumentTextDetectionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDocumentTextDetection, input)
	return &TextractGetDocumentTextDetectionResult{Result: future}
}

func (a *TextractStub) StartDocumentAnalysis(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error) {
	var output textract.StartDocumentAnalysisOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartDocumentAnalysis, input).Get(ctx, &output)
	return &output, err
}

func (a *TextractStub) StartDocumentAnalysisAsync(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) *TextractStartDocumentAnalysisResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartDocumentAnalysis, input)
	return &TextractStartDocumentAnalysisResult{Result: future}
}

func (a *TextractStub) StartDocumentTextDetection(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error) {
	var output textract.StartDocumentTextDetectionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.StartDocumentTextDetection, input).Get(ctx, &output)
	return &output, err
}

func (a *TextractStub) StartDocumentTextDetectionAsync(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) *TextractStartDocumentTextDetectionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.StartDocumentTextDetection, input)
	return &TextractStartDocumentTextDetectionResult{Result: future}
}
