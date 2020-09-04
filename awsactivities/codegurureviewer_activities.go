
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"github.com/aws/aws-sdk-go/service/codegurureviewer/codegururevieweriface"
)

type CodeGuruReviewerActivities struct {
	client codegururevieweriface.CodeGuruReviewerAPI
}

func NewCodeGuruReviewerActivities(client codegururevieweriface.CodeGuruReviewerAPI) *CodeGuruReviewerActivities {
    return &CodeGuruReviewerActivities{client: client}
}

func (a *CodeGuruReviewerActivities) AssociateRepository(input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error) {
    return a.client.AssociateRepository(input)
}

func (a *CodeGuruReviewerActivities) DescribeCodeReview(input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error) {
    return a.client.DescribeCodeReview(input)
}

func (a *CodeGuruReviewerActivities) DescribeRecommendationFeedback(input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
    return a.client.DescribeRecommendationFeedback(input)
}

func (a *CodeGuruReviewerActivities) DescribeRepositoryAssociation(input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
    return a.client.DescribeRepositoryAssociation(input)
}

func (a *CodeGuruReviewerActivities) DisassociateRepository(input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error) {
    return a.client.DisassociateRepository(input)
}

func (a *CodeGuruReviewerActivities) ListCodeReviews(input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error) {
    return a.client.ListCodeReviews(input)
}

func (a *CodeGuruReviewerActivities) ListRecommendationFeedback(input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
    return a.client.ListRecommendationFeedback(input)
}

func (a *CodeGuruReviewerActivities) ListRecommendations(input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error) {
    return a.client.ListRecommendations(input)
}

func (a *CodeGuruReviewerActivities) ListRepositoryAssociations(input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
    return a.client.ListRepositoryAssociations(input)
}

func (a *CodeGuruReviewerActivities) PutRecommendationFeedback(input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error) {
    return a.client.PutRecommendationFeedback(input)
}
