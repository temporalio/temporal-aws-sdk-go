
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"github.com/aws/aws-sdk-go/service/personalizeruntime/personalizeruntimeiface"
)

type PersonalizeRuntimeActivities struct {
	client personalizeruntimeiface.PersonalizeRuntimeAPI
}

func NewPersonalizeRuntimeActivities(client personalizeruntimeiface.PersonalizeRuntimeAPI) *PersonalizeRuntimeActivities {
    return &PersonalizeRuntimeActivities{client: client}
}

func (a *PersonalizeRuntimeActivities) GetPersonalizedRanking(input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
    return a.client.GetPersonalizedRanking(input)
}

func (a *PersonalizeRuntimeActivities) GetRecommendations(input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
    return a.client.GetRecommendations(input)
}
