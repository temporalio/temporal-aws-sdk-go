package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"github.com/aws/aws-sdk-go/service/codegurureviewer/codegururevieweriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type CodeGuruReviewerActivities struct {
	client codegururevieweriface.CodeGuruReviewerAPI
}

func NewCodeGuruReviewerActivities(session *session.Session, config ...*aws.Config) *CodeGuruReviewerActivities {
	client := codegurureviewer.New(session, config...)
	return &CodeGuruReviewerActivities{client: client}
}

func (a *CodeGuruReviewerActivities) AssociateRepository(ctx context.Context, input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error) {
	return a.client.AssociateRepositoryWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) CreateCodeReview(ctx context.Context, input *codegurureviewer.CreateCodeReviewInput) (*codegurureviewer.CreateCodeReviewOutput, error) {
	return a.client.CreateCodeReviewWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) DescribeCodeReview(ctx context.Context, input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error) {
	return a.client.DescribeCodeReviewWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) DescribeRecommendationFeedback(ctx context.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
	return a.client.DescribeRecommendationFeedbackWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) DescribeRepositoryAssociation(ctx context.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
	return a.client.DescribeRepositoryAssociationWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) DisassociateRepository(ctx context.Context, input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error) {
	return a.client.DisassociateRepositoryWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) ListCodeReviews(ctx context.Context, input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error) {
	return a.client.ListCodeReviewsWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) ListRecommendationFeedback(ctx context.Context, input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
	return a.client.ListRecommendationFeedbackWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) ListRecommendations(ctx context.Context, input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error) {
	return a.client.ListRecommendationsWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) ListRepositoryAssociations(ctx context.Context, input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
	return a.client.ListRepositoryAssociationsWithContext(ctx, input)
}

func (a *CodeGuruReviewerActivities) PutRecommendationFeedback(ctx context.Context, input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error) {
	return a.client.PutRecommendationFeedbackWithContext(ctx, input)
}
