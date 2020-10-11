// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/textract/textractiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type TextractActivities struct {
	client textractiface.TextractAPI

	sessionFactory SessionFactory
}

func NewTextractActivities(sess *session.Session, config ...*aws.Config) *TextractActivities {
	client := textract.New(sess, config...)
	return &TextractActivities{client: client}
}

func NewTextractActivitiesWithSessionFactory(sessionFactory SessionFactory) *TextractActivities {
	return &TextractActivities{sessionFactory: sessionFactory}
}

func (a *TextractActivities) getClient(ctx context.Context) (textractiface.TextractAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return textract.New(sess), nil
}

func (a *TextractActivities) AnalyzeDocument(ctx context.Context, input *textract.AnalyzeDocumentInput) (*textract.AnalyzeDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AnalyzeDocumentWithContext(ctx, input)
}

func (a *TextractActivities) DetectDocumentText(ctx context.Context, input *textract.DetectDocumentTextInput) (*textract.DetectDocumentTextOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectDocumentTextWithContext(ctx, input)
}

func (a *TextractActivities) GetDocumentAnalysis(ctx context.Context, input *textract.GetDocumentAnalysisInput) (*textract.GetDocumentAnalysisOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDocumentAnalysisWithContext(ctx, input)
}

func (a *TextractActivities) GetDocumentTextDetection(ctx context.Context, input *textract.GetDocumentTextDetectionInput) (*textract.GetDocumentTextDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDocumentTextDetectionWithContext(ctx, input)
}

func (a *TextractActivities) StartDocumentAnalysis(ctx context.Context, input *textract.StartDocumentAnalysisInput) (*textract.StartDocumentAnalysisOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDocumentAnalysisWithContext(ctx, input)
}

func (a *TextractActivities) StartDocumentTextDetection(ctx context.Context, input *textract.StartDocumentTextDetectionInput) (*textract.StartDocumentTextDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDocumentTextDetectionWithContext(ctx, input)
}
