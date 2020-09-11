package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"github.com/aws/aws-sdk-go/service/personalizeruntime/personalizeruntimeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type PersonalizeRuntimeActivities struct {
	client personalizeruntimeiface.PersonalizeRuntimeAPI
}

func NewPersonalizeRuntimeActivities(session *session.Session, config ...*aws.Config) *PersonalizeRuntimeActivities {
	client := personalizeruntime.New(session, config...)
	return &PersonalizeRuntimeActivities{client: client}
}

func (a *PersonalizeRuntimeActivities) GetPersonalizedRanking(ctx context.Context, input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	return a.client.GetPersonalizedRankingWithContext(ctx, input)
}

func (a *PersonalizeRuntimeActivities) GetRecommendations(ctx context.Context, input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
	return a.client.GetRecommendationsWithContext(ctx, input)
}
