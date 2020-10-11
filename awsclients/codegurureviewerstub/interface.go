// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package codegurureviewerstub

import (
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
    "go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

type Client interface {
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

func NewClient() Client {
	return &stub{}
}

