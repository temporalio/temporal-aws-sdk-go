// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package textractstub

import (
	"github.com/aws/aws-sdk-go/service/textract"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type TextractAnalyzeDocumentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextractAnalyzeDocumentFuture) Get(ctx workflow.Context) (*textract.AnalyzeDocumentOutput, error) {
	var output textract.AnalyzeDocumentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TextractDetectDocumentTextFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextractDetectDocumentTextFuture) Get(ctx workflow.Context) (*textract.DetectDocumentTextOutput, error) {
	var output textract.DetectDocumentTextOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TextractGetDocumentAnalysisFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextractGetDocumentAnalysisFuture) Get(ctx workflow.Context) (*textract.GetDocumentAnalysisOutput, error) {
	var output textract.GetDocumentAnalysisOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TextractGetDocumentTextDetectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextractGetDocumentTextDetectionFuture) Get(ctx workflow.Context) (*textract.GetDocumentTextDetectionOutput, error) {
	var output textract.GetDocumentTextDetectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TextractStartDocumentAnalysisFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextractStartDocumentAnalysisFuture) Get(ctx workflow.Context) (*textract.StartDocumentAnalysisOutput, error) {
	var output textract.StartDocumentAnalysisOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TextractStartDocumentTextDetectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TextractStartDocumentTextDetectionFuture) Get(ctx workflow.Context) (*textract.StartDocumentTextDetectionOutput, error) {
	var output textract.StartDocumentTextDetectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AnalyzeDocument(ctx workflow.Context, input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
	var output textract.AnalyzeDocumentOutput
	err := workflow.ExecuteActivity(ctx, "aws.textract.AnalyzeDocument", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AnalyzeDocumentAsync(ctx workflow.Context, input *textract.AnalyzeDocumentInput) *TextractAnalyzeDocumentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.textract.AnalyzeDocument", input)
	return &TextractAnalyzeDocumentFuture{Future: future}
}

func (a *stub) DetectDocumentText(ctx workflow.Context, input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error) {
	var output textract.DetectDocumentTextOutput
	err := workflow.ExecuteActivity(ctx, "aws.textract.DetectDocumentText", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DetectDocumentTextAsync(ctx workflow.Context, input *textract.DetectDocumentTextInput) *TextractDetectDocumentTextFuture {
	future := workflow.ExecuteActivity(ctx, "aws.textract.DetectDocumentText", input)
	return &TextractDetectDocumentTextFuture{Future: future}
}

func (a *stub) GetDocumentAnalysis(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error) {
	var output textract.GetDocumentAnalysisOutput
	err := workflow.ExecuteActivity(ctx, "aws.textract.GetDocumentAnalysis", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDocumentAnalysisAsync(ctx workflow.Context, input *textract.GetDocumentAnalysisInput) *TextractGetDocumentAnalysisFuture {
	future := workflow.ExecuteActivity(ctx, "aws.textract.GetDocumentAnalysis", input)
	return &TextractGetDocumentAnalysisFuture{Future: future}
}

func (a *stub) GetDocumentTextDetection(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error) {
	var output textract.GetDocumentTextDetectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.textract.GetDocumentTextDetection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDocumentTextDetectionAsync(ctx workflow.Context, input *textract.GetDocumentTextDetectionInput) *TextractGetDocumentTextDetectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.textract.GetDocumentTextDetection", input)
	return &TextractGetDocumentTextDetectionFuture{Future: future}
}

func (a *stub) StartDocumentAnalysis(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error) {
	var output textract.StartDocumentAnalysisOutput
	err := workflow.ExecuteActivity(ctx, "aws.textract.StartDocumentAnalysis", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDocumentAnalysisAsync(ctx workflow.Context, input *textract.StartDocumentAnalysisInput) *TextractStartDocumentAnalysisFuture {
	future := workflow.ExecuteActivity(ctx, "aws.textract.StartDocumentAnalysis", input)
	return &TextractStartDocumentAnalysisFuture{Future: future}
}

func (a *stub) StartDocumentTextDetection(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error) {
	var output textract.StartDocumentTextDetectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.textract.StartDocumentTextDetection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartDocumentTextDetectionAsync(ctx workflow.Context, input *textract.StartDocumentTextDetectionInput) *TextractStartDocumentTextDetectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.textract.StartDocumentTextDetection", input)
	return &TextractStartDocumentTextDetectionFuture{Future: future}
}
