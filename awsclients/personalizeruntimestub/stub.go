// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package personalizeruntimestub

import (
	"github.com/aws/aws-sdk-go/service/personalizeruntime"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type PersonalizeRuntimeGetPersonalizedRankingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PersonalizeRuntimeGetPersonalizedRankingFuture) Get(ctx workflow.Context) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	var output personalizeruntime.GetPersonalizedRankingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PersonalizeRuntimeGetRecommendationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PersonalizeRuntimeGetRecommendationsFuture) Get(ctx workflow.Context) (*personalizeruntime.GetRecommendationsOutput, error) {
	var output personalizeruntime.GetRecommendationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPersonalizedRanking(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) (*personalizeruntime.GetPersonalizedRankingOutput, error) {
	var output personalizeruntime.GetPersonalizedRankingOutput
	err := workflow.ExecuteActivity(ctx, "aws.personalizeruntime.GetPersonalizedRanking", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPersonalizedRankingAsync(ctx workflow.Context, input *personalizeruntime.GetPersonalizedRankingInput) *PersonalizeRuntimeGetPersonalizedRankingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.personalizeruntime.GetPersonalizedRanking", input)
	return &PersonalizeRuntimeGetPersonalizedRankingFuture{Future: future}
}

func (a *stub) GetRecommendations(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) (*personalizeruntime.GetRecommendationsOutput, error) {
	var output personalizeruntime.GetRecommendationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.personalizeruntime.GetRecommendations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRecommendationsAsync(ctx workflow.Context, input *personalizeruntime.GetRecommendationsInput) *PersonalizeRuntimeGetRecommendationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.personalizeruntime.GetRecommendations", input)
	return &PersonalizeRuntimeGetRecommendationsFuture{Future: future}
}