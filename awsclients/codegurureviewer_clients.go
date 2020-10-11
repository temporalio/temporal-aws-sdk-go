// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"go.temporal.io/sdk/workflow"
)

type CodeGuruReviewerClient interface {
	AssociateRepository(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error)
	AssociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) *CodeGuruReviewerAssociateRepositoryFuture

	CreateCodeReview(ctx workflow.Context, input *codegurureviewer.CreateCodeReviewInput) (*codegurureviewer.CreateCodeReviewOutput, error)
	CreateCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.CreateCodeReviewInput) *CodeGuruReviewerCreateCodeReviewFuture

	DescribeCodeReview(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error)
	DescribeCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) *CodeGuruReviewerDescribeCodeReviewFuture

	DescribeRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error)
	DescribeRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) *CodeGuruReviewerDescribeRecommendationFeedbackFuture

	DescribeRepositoryAssociation(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DescribeRepositoryAssociationAsync(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) *CodeGuruReviewerDescribeRepositoryAssociationFuture

	DisassociateRepository(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error)
	DisassociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) *CodeGuruReviewerDisassociateRepositoryFuture

	ListCodeReviews(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error)
	ListCodeReviewsAsync(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) *CodeGuruReviewerListCodeReviewsFuture

	ListRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error)
	ListRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) *CodeGuruReviewerListRecommendationFeedbackFuture

	ListRecommendations(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error)
	ListRecommendationsAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) *CodeGuruReviewerListRecommendationsFuture

	ListRepositoryAssociations(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListRepositoryAssociationsAsync(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) *CodeGuruReviewerListRepositoryAssociationsFuture

	PutRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error)
	PutRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) *CodeGuruReviewerPutRecommendationFeedbackFuture
}

type CodeGuruReviewerStub struct{}

func NewCodeGuruReviewerStub() CodeGuruReviewerClient {
	return &CodeGuruReviewerStub{}
}

type CodeGuruReviewerAssociateRepositoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerAssociateRepositoryFuture) Get(ctx workflow.Context) (*codegurureviewer.AssociateRepositoryOutput, error) {
	var output codegurureviewer.AssociateRepositoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerCreateCodeReviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerCreateCodeReviewFuture) Get(ctx workflow.Context) (*codegurureviewer.CreateCodeReviewOutput, error) {
	var output codegurureviewer.CreateCodeReviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerDescribeCodeReviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerDescribeCodeReviewFuture) Get(ctx workflow.Context) (*codegurureviewer.DescribeCodeReviewOutput, error) {
	var output codegurureviewer.DescribeCodeReviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerDescribeRecommendationFeedbackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerDescribeRecommendationFeedbackFuture) Get(ctx workflow.Context) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
	var output codegurureviewer.DescribeRecommendationFeedbackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerDescribeRepositoryAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerDescribeRepositoryAssociationFuture) Get(ctx workflow.Context) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
	var output codegurureviewer.DescribeRepositoryAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerDisassociateRepositoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerDisassociateRepositoryFuture) Get(ctx workflow.Context) (*codegurureviewer.DisassociateRepositoryOutput, error) {
	var output codegurureviewer.DisassociateRepositoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerListCodeReviewsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerListCodeReviewsFuture) Get(ctx workflow.Context) (*codegurureviewer.ListCodeReviewsOutput, error) {
	var output codegurureviewer.ListCodeReviewsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerListRecommendationFeedbackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerListRecommendationFeedbackFuture) Get(ctx workflow.Context) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
	var output codegurureviewer.ListRecommendationFeedbackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerListRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerListRecommendationsFuture) Get(ctx workflow.Context) (*codegurureviewer.ListRecommendationsOutput, error) {
	var output codegurureviewer.ListRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerListRepositoryAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerListRepositoryAssociationsFuture) Get(ctx workflow.Context) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
	var output codegurureviewer.ListRepositoryAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CodeGuruReviewerPutRecommendationFeedbackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CodeGuruReviewerPutRecommendationFeedbackFuture) Get(ctx workflow.Context) (*codegurureviewer.PutRecommendationFeedbackOutput, error) {
	var output codegurureviewer.PutRecommendationFeedbackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) AssociateRepository(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error) {
	var output codegurureviewer.AssociateRepositoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.AssociateRepository", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) AssociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) *CodeGuruReviewerAssociateRepositoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.AssociateRepository", input)
	return &CodeGuruReviewerAssociateRepositoryFuture{Future: future}
}

func (a *CodeGuruReviewerStub) CreateCodeReview(ctx workflow.Context, input *codegurureviewer.CreateCodeReviewInput) (*codegurureviewer.CreateCodeReviewOutput, error) {
	var output codegurureviewer.CreateCodeReviewOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.CreateCodeReview", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) CreateCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.CreateCodeReviewInput) *CodeGuruReviewerCreateCodeReviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.CreateCodeReview", input)
	return &CodeGuruReviewerCreateCodeReviewFuture{Future: future}
}

func (a *CodeGuruReviewerStub) DescribeCodeReview(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error) {
	var output codegurureviewer.DescribeCodeReviewOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DescribeCodeReview", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) DescribeCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) *CodeGuruReviewerDescribeCodeReviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DescribeCodeReview", input)
	return &CodeGuruReviewerDescribeCodeReviewFuture{Future: future}
}

func (a *CodeGuruReviewerStub) DescribeRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
	var output codegurureviewer.DescribeRecommendationFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DescribeRecommendationFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) DescribeRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) *CodeGuruReviewerDescribeRecommendationFeedbackFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DescribeRecommendationFeedback", input)
	return &CodeGuruReviewerDescribeRecommendationFeedbackFuture{Future: future}
}

func (a *CodeGuruReviewerStub) DescribeRepositoryAssociation(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
	var output codegurureviewer.DescribeRepositoryAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DescribeRepositoryAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) DescribeRepositoryAssociationAsync(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) *CodeGuruReviewerDescribeRepositoryAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DescribeRepositoryAssociation", input)
	return &CodeGuruReviewerDescribeRepositoryAssociationFuture{Future: future}
}

func (a *CodeGuruReviewerStub) DisassociateRepository(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error) {
	var output codegurureviewer.DisassociateRepositoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DisassociateRepository", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) DisassociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) *CodeGuruReviewerDisassociateRepositoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.DisassociateRepository", input)
	return &CodeGuruReviewerDisassociateRepositoryFuture{Future: future}
}

func (a *CodeGuruReviewerStub) ListCodeReviews(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error) {
	var output codegurureviewer.ListCodeReviewsOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListCodeReviews", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) ListCodeReviewsAsync(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) *CodeGuruReviewerListCodeReviewsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListCodeReviews", input)
	return &CodeGuruReviewerListCodeReviewsFuture{Future: future}
}

func (a *CodeGuruReviewerStub) ListRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
	var output codegurureviewer.ListRecommendationFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListRecommendationFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) ListRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) *CodeGuruReviewerListRecommendationFeedbackFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListRecommendationFeedback", input)
	return &CodeGuruReviewerListRecommendationFeedbackFuture{Future: future}
}

func (a *CodeGuruReviewerStub) ListRecommendations(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error) {
	var output codegurureviewer.ListRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) ListRecommendationsAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) *CodeGuruReviewerListRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListRecommendations", input)
	return &CodeGuruReviewerListRecommendationsFuture{Future: future}
}

func (a *CodeGuruReviewerStub) ListRepositoryAssociations(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
	var output codegurureviewer.ListRepositoryAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListRepositoryAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) ListRepositoryAssociationsAsync(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) *CodeGuruReviewerListRepositoryAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.ListRepositoryAssociations", input)
	return &CodeGuruReviewerListRepositoryAssociationsFuture{Future: future}
}

func (a *CodeGuruReviewerStub) PutRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error) {
	var output codegurureviewer.PutRecommendationFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.PutRecommendationFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *CodeGuruReviewerStub) PutRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) *CodeGuruReviewerPutRecommendationFeedbackFuture {
	future := workflow.ExecuteActivity(ctx, "aws.codegurureviewer.PutRecommendationFeedback", input)
	return &CodeGuruReviewerPutRecommendationFeedbackFuture{Future: future}
}
