package awsclients

import (
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type PersonalizeRuntimeClient interface {
       GetPersonalizedRanking(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error)
       GetPersonalizedRankingAsync(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) *PersonalizeruntimeGetPersonalizedRankingResult

       GetRecommendations(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error)
       GetRecommendationsAsync(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) *PersonalizeruntimeGetRecommendationsResult
}

type PersonalizeruntimeGetPersonalizedRankingResult struct {
	Result workflow.Future
}

func (r *PersonalizeruntimeGetPersonalizedRankingResult) Get(ctx workflow.Context) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
    var output personalizeruntime.GetPersonalizedRankingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeruntimeGetRecommendationsResult struct {
	Result workflow.Future
}

func (r *PersonalizeruntimeGetRecommendationsResult) Get(ctx workflow.Context) (*personalizeruntime.GetRecommendationsOutput, error) {
    var output personalizeruntime.GetRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeRuntimeStub struct {
    activities awsactivities.PersonalizeRuntimeActivities
}

func NewPersonalizeRuntimeStub() PersonalizeRuntimeClient {
    return &PersonalizeRuntimeStub{}
}

func (a *PersonalizeRuntimeStub) GetPersonalizedRanking(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
    var output personalizeruntime.GetPersonalizedRankingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPersonalizedRanking, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeRuntimeStub) GetPersonalizedRankingAsync(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) *PersonalizeruntimeGetPersonalizedRankingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPersonalizedRanking, input)
    return &PersonalizeruntimeGetPersonalizedRankingResult{Result: future}
}

func (a *PersonalizeRuntimeStub) GetRecommendations(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
    var output personalizeruntime.GetRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecommendations, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeRuntimeStub) GetRecommendationsAsync(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) *PersonalizeruntimeGetRecommendationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRecommendations, input)
    return &PersonalizeruntimeGetRecommendationsResult{Result: future}
}
