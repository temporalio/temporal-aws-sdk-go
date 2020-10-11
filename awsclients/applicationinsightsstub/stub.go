// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package applicationinsightsstub

import (
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type ApplicationInsightsCreateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsCreateApplicationFuture) Get(ctx workflow.Context) (*applicationinsights.CreateApplicationOutput, error) {
	var output applicationinsights.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsCreateComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsCreateComponentFuture) Get(ctx workflow.Context) (*applicationinsights.CreateComponentOutput, error) {
	var output applicationinsights.CreateComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsCreateLogPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsCreateLogPatternFuture) Get(ctx workflow.Context) (*applicationinsights.CreateLogPatternOutput, error) {
	var output applicationinsights.CreateLogPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDeleteApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDeleteApplicationFuture) Get(ctx workflow.Context) (*applicationinsights.DeleteApplicationOutput, error) {
	var output applicationinsights.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDeleteComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDeleteComponentFuture) Get(ctx workflow.Context) (*applicationinsights.DeleteComponentOutput, error) {
	var output applicationinsights.DeleteComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDeleteLogPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDeleteLogPatternFuture) Get(ctx workflow.Context) (*applicationinsights.DeleteLogPatternOutput, error) {
	var output applicationinsights.DeleteLogPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeApplicationFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeApplicationOutput, error) {
	var output applicationinsights.DescribeApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeComponentFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeComponentOutput, error) {
	var output applicationinsights.DescribeComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeComponentConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeComponentConfigurationFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	var output applicationinsights.DescribeComponentConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeComponentConfigurationRecommendationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeComponentConfigurationRecommendationFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	var output applicationinsights.DescribeComponentConfigurationRecommendationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeLogPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeLogPatternFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeLogPatternOutput, error) {
	var output applicationinsights.DescribeLogPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeObservationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeObservationFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeObservationOutput, error) {
	var output applicationinsights.DescribeObservationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeProblemFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeProblemFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeProblemOutput, error) {
	var output applicationinsights.DescribeProblemOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsDescribeProblemObservationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsDescribeProblemObservationsFuture) Get(ctx workflow.Context) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	var output applicationinsights.DescribeProblemObservationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListApplicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListApplicationsFuture) Get(ctx workflow.Context) (*applicationinsights.ListApplicationsOutput, error) {
	var output applicationinsights.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListComponentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListComponentsFuture) Get(ctx workflow.Context) (*applicationinsights.ListComponentsOutput, error) {
	var output applicationinsights.ListComponentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListConfigurationHistoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListConfigurationHistoryFuture) Get(ctx workflow.Context) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	var output applicationinsights.ListConfigurationHistoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListLogPatternSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListLogPatternSetsFuture) Get(ctx workflow.Context) (*applicationinsights.ListLogPatternSetsOutput, error) {
	var output applicationinsights.ListLogPatternSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListLogPatternsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListLogPatternsFuture) Get(ctx workflow.Context) (*applicationinsights.ListLogPatternsOutput, error) {
	var output applicationinsights.ListLogPatternsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListProblemsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListProblemsFuture) Get(ctx workflow.Context) (*applicationinsights.ListProblemsOutput, error) {
	var output applicationinsights.ListProblemsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsListTagsForResourceFuture) Get(ctx workflow.Context) (*applicationinsights.ListTagsForResourceOutput, error) {
	var output applicationinsights.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsTagResourceFuture) Get(ctx workflow.Context) (*applicationinsights.TagResourceOutput, error) {
	var output applicationinsights.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsUntagResourceFuture) Get(ctx workflow.Context) (*applicationinsights.UntagResourceOutput, error) {
	var output applicationinsights.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsUpdateApplicationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsUpdateApplicationFuture) Get(ctx workflow.Context) (*applicationinsights.UpdateApplicationOutput, error) {
	var output applicationinsights.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsUpdateComponentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsUpdateComponentFuture) Get(ctx workflow.Context) (*applicationinsights.UpdateComponentOutput, error) {
	var output applicationinsights.UpdateComponentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsUpdateComponentConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsUpdateComponentConfigurationFuture) Get(ctx workflow.Context) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
	var output applicationinsights.UpdateComponentConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplicationInsightsUpdateLogPatternFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplicationInsightsUpdateLogPatternFuture) Get(ctx workflow.Context) (*applicationinsights.UpdateLogPatternOutput, error) {
	var output applicationinsights.UpdateLogPatternOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplication(ctx workflow.Context, input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error) {
	var output applicationinsights.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateApplicationAsync(ctx workflow.Context, input *applicationinsights.CreateApplicationInput) *ApplicationInsightsCreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.CreateApplication", input)
	return &ApplicationInsightsCreateApplicationFuture{Future: future}
}

func (a *stub) CreateComponent(ctx workflow.Context, input *applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error) {
	var output applicationinsights.CreateComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.CreateComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateComponentAsync(ctx workflow.Context, input *applicationinsights.CreateComponentInput) *ApplicationInsightsCreateComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.CreateComponent", input)
	return &ApplicationInsightsCreateComponentFuture{Future: future}
}

func (a *stub) CreateLogPattern(ctx workflow.Context, input *applicationinsights.CreateLogPatternInput) (*applicationinsights.CreateLogPatternOutput, error) {
	var output applicationinsights.CreateLogPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.CreateLogPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLogPatternAsync(ctx workflow.Context, input *applicationinsights.CreateLogPatternInput) *ApplicationInsightsCreateLogPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.CreateLogPattern", input)
	return &ApplicationInsightsCreateLogPatternFuture{Future: future}
}

func (a *stub) DeleteApplication(ctx workflow.Context, input *applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error) {
	var output applicationinsights.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteApplicationAsync(ctx workflow.Context, input *applicationinsights.DeleteApplicationInput) *ApplicationInsightsDeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DeleteApplication", input)
	return &ApplicationInsightsDeleteApplicationFuture{Future: future}
}

func (a *stub) DeleteComponent(ctx workflow.Context, input *applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error) {
	var output applicationinsights.DeleteComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DeleteComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteComponentAsync(ctx workflow.Context, input *applicationinsights.DeleteComponentInput) *ApplicationInsightsDeleteComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DeleteComponent", input)
	return &ApplicationInsightsDeleteComponentFuture{Future: future}
}

func (a *stub) DeleteLogPattern(ctx workflow.Context, input *applicationinsights.DeleteLogPatternInput) (*applicationinsights.DeleteLogPatternOutput, error) {
	var output applicationinsights.DeleteLogPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DeleteLogPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLogPatternAsync(ctx workflow.Context, input *applicationinsights.DeleteLogPatternInput) *ApplicationInsightsDeleteLogPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DeleteLogPattern", input)
	return &ApplicationInsightsDeleteLogPatternFuture{Future: future}
}

func (a *stub) DescribeApplication(ctx workflow.Context, input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error) {
	var output applicationinsights.DescribeApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeApplicationAsync(ctx workflow.Context, input *applicationinsights.DescribeApplicationInput) *ApplicationInsightsDescribeApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeApplication", input)
	return &ApplicationInsightsDescribeApplicationFuture{Future: future}
}

func (a *stub) DescribeComponent(ctx workflow.Context, input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error) {
	var output applicationinsights.DescribeComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeComponentAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentInput) *ApplicationInsightsDescribeComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeComponent", input)
	return &ApplicationInsightsDescribeComponentFuture{Future: future}
}

func (a *stub) DescribeComponentConfiguration(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	var output applicationinsights.DescribeComponentConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeComponentConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeComponentConfigurationAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationInput) *ApplicationInsightsDescribeComponentConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeComponentConfiguration", input)
	return &ApplicationInsightsDescribeComponentConfigurationFuture{Future: future}
}

func (a *stub) DescribeComponentConfigurationRecommendation(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	var output applicationinsights.DescribeComponentConfigurationRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeComponentConfigurationRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeComponentConfigurationRecommendationAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) *ApplicationInsightsDescribeComponentConfigurationRecommendationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeComponentConfigurationRecommendation", input)
	return &ApplicationInsightsDescribeComponentConfigurationRecommendationFuture{Future: future}
}

func (a *stub) DescribeLogPattern(ctx workflow.Context, input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error) {
	var output applicationinsights.DescribeLogPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeLogPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLogPatternAsync(ctx workflow.Context, input *applicationinsights.DescribeLogPatternInput) *ApplicationInsightsDescribeLogPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeLogPattern", input)
	return &ApplicationInsightsDescribeLogPatternFuture{Future: future}
}

func (a *stub) DescribeObservation(ctx workflow.Context, input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error) {
	var output applicationinsights.DescribeObservationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeObservation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeObservationAsync(ctx workflow.Context, input *applicationinsights.DescribeObservationInput) *ApplicationInsightsDescribeObservationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeObservation", input)
	return &ApplicationInsightsDescribeObservationFuture{Future: future}
}

func (a *stub) DescribeProblem(ctx workflow.Context, input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error) {
	var output applicationinsights.DescribeProblemOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeProblem", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProblemAsync(ctx workflow.Context, input *applicationinsights.DescribeProblemInput) *ApplicationInsightsDescribeProblemFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeProblem", input)
	return &ApplicationInsightsDescribeProblemFuture{Future: future}
}

func (a *stub) DescribeProblemObservations(ctx workflow.Context, input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	var output applicationinsights.DescribeProblemObservationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeProblemObservations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProblemObservationsAsync(ctx workflow.Context, input *applicationinsights.DescribeProblemObservationsInput) *ApplicationInsightsDescribeProblemObservationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.DescribeProblemObservations", input)
	return &ApplicationInsightsDescribeProblemObservationsFuture{Future: future}
}

func (a *stub) ListApplications(ctx workflow.Context, input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error) {
	var output applicationinsights.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListApplicationsAsync(ctx workflow.Context, input *applicationinsights.ListApplicationsInput) *ApplicationInsightsListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListApplications", input)
	return &ApplicationInsightsListApplicationsFuture{Future: future}
}

func (a *stub) ListComponents(ctx workflow.Context, input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error) {
	var output applicationinsights.ListComponentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListComponents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListComponentsAsync(ctx workflow.Context, input *applicationinsights.ListComponentsInput) *ApplicationInsightsListComponentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListComponents", input)
	return &ApplicationInsightsListComponentsFuture{Future: future}
}

func (a *stub) ListConfigurationHistory(ctx workflow.Context, input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	var output applicationinsights.ListConfigurationHistoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListConfigurationHistory", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationHistoryAsync(ctx workflow.Context, input *applicationinsights.ListConfigurationHistoryInput) *ApplicationInsightsListConfigurationHistoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListConfigurationHistory", input)
	return &ApplicationInsightsListConfigurationHistoryFuture{Future: future}
}

func (a *stub) ListLogPatternSets(ctx workflow.Context, input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error) {
	var output applicationinsights.ListLogPatternSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListLogPatternSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLogPatternSetsAsync(ctx workflow.Context, input *applicationinsights.ListLogPatternSetsInput) *ApplicationInsightsListLogPatternSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListLogPatternSets", input)
	return &ApplicationInsightsListLogPatternSetsFuture{Future: future}
}

func (a *stub) ListLogPatterns(ctx workflow.Context, input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error) {
	var output applicationinsights.ListLogPatternsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListLogPatterns", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListLogPatternsAsync(ctx workflow.Context, input *applicationinsights.ListLogPatternsInput) *ApplicationInsightsListLogPatternsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListLogPatterns", input)
	return &ApplicationInsightsListLogPatternsFuture{Future: future}
}

func (a *stub) ListProblems(ctx workflow.Context, input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error) {
	var output applicationinsights.ListProblemsOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListProblems", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProblemsAsync(ctx workflow.Context, input *applicationinsights.ListProblemsInput) *ApplicationInsightsListProblemsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListProblems", input)
	return &ApplicationInsightsListProblemsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error) {
	var output applicationinsights.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *applicationinsights.ListTagsForResourceInput) *ApplicationInsightsListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.ListTagsForResource", input)
	return &ApplicationInsightsListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *applicationinsights.TagResourceInput) (*applicationinsights.TagResourceOutput, error) {
	var output applicationinsights.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *applicationinsights.TagResourceInput) *ApplicationInsightsTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.TagResource", input)
	return &ApplicationInsightsTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *applicationinsights.UntagResourceInput) (*applicationinsights.UntagResourceOutput, error) {
	var output applicationinsights.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *applicationinsights.UntagResourceInput) *ApplicationInsightsUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UntagResource", input)
	return &ApplicationInsightsUntagResourceFuture{Future: future}
}

func (a *stub) UpdateApplication(ctx workflow.Context, input *applicationinsights.UpdateApplicationInput) (*applicationinsights.UpdateApplicationOutput, error) {
	var output applicationinsights.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateApplicationAsync(ctx workflow.Context, input *applicationinsights.UpdateApplicationInput) *ApplicationInsightsUpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateApplication", input)
	return &ApplicationInsightsUpdateApplicationFuture{Future: future}
}

func (a *stub) UpdateComponent(ctx workflow.Context, input *applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error) {
	var output applicationinsights.UpdateComponentOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateComponent", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateComponentAsync(ctx workflow.Context, input *applicationinsights.UpdateComponentInput) *ApplicationInsightsUpdateComponentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateComponent", input)
	return &ApplicationInsightsUpdateComponentFuture{Future: future}
}

func (a *stub) UpdateComponentConfiguration(ctx workflow.Context, input *applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
	var output applicationinsights.UpdateComponentConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateComponentConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateComponentConfigurationAsync(ctx workflow.Context, input *applicationinsights.UpdateComponentConfigurationInput) *ApplicationInsightsUpdateComponentConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateComponentConfiguration", input)
	return &ApplicationInsightsUpdateComponentConfigurationFuture{Future: future}
}

func (a *stub) UpdateLogPattern(ctx workflow.Context, input *applicationinsights.UpdateLogPatternInput) (*applicationinsights.UpdateLogPatternOutput, error) {
	var output applicationinsights.UpdateLogPatternOutput
	err := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateLogPattern", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateLogPatternAsync(ctx workflow.Context, input *applicationinsights.UpdateLogPatternInput) *ApplicationInsightsUpdateLogPatternFuture {
	future := workflow.ExecuteActivity(ctx, "aws.applicationinsights.UpdateLogPattern", input)
	return &ApplicationInsightsUpdateLogPatternFuture{Future: future}
}
