package awsclients

import (
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"go.temporal.io/sdk/workflow"
)

type ApplicationInsightsClient interface {
       CreateApplication(ctx workflow.Context, input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error)
       CreateApplicationAsync(ctx workflow.Context, input *applicationinsights.CreateApplicationInput) *ApplicationinsightsCreateApplicationResult

       CreateComponent(ctx workflow.Context, input *applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error)
       CreateComponentAsync(ctx workflow.Context, input *applicationinsights.CreateComponentInput) *ApplicationinsightsCreateComponentResult

       CreateLogPattern(ctx workflow.Context, input *applicationinsights.CreateLogPatternInput) (*applicationinsights.CreateLogPatternOutput, error)
       CreateLogPatternAsync(ctx workflow.Context, input *applicationinsights.CreateLogPatternInput) *ApplicationinsightsCreateLogPatternResult

       DeleteApplication(ctx workflow.Context, input *applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error)
       DeleteApplicationAsync(ctx workflow.Context, input *applicationinsights.DeleteApplicationInput) *ApplicationinsightsDeleteApplicationResult

       DeleteComponent(ctx workflow.Context, input *applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error)
       DeleteComponentAsync(ctx workflow.Context, input *applicationinsights.DeleteComponentInput) *ApplicationinsightsDeleteComponentResult

       DeleteLogPattern(ctx workflow.Context, input *applicationinsights.DeleteLogPatternInput) (*applicationinsights.DeleteLogPatternOutput, error)
       DeleteLogPatternAsync(ctx workflow.Context, input *applicationinsights.DeleteLogPatternInput) *ApplicationinsightsDeleteLogPatternResult

       DescribeApplication(ctx workflow.Context, input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error)
       DescribeApplicationAsync(ctx workflow.Context, input *applicationinsights.DescribeApplicationInput) *ApplicationinsightsDescribeApplicationResult

       DescribeComponent(ctx workflow.Context, input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error)
       DescribeComponentAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentInput) *ApplicationinsightsDescribeComponentResult

       DescribeComponentConfiguration(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error)
       DescribeComponentConfigurationAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationInput) *ApplicationinsightsDescribeComponentConfigurationResult

       DescribeComponentConfigurationRecommendation(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error)
       DescribeComponentConfigurationRecommendationAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) *ApplicationinsightsDescribeComponentConfigurationRecommendationResult

       DescribeLogPattern(ctx workflow.Context, input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error)
       DescribeLogPatternAsync(ctx workflow.Context, input *applicationinsights.DescribeLogPatternInput) *ApplicationinsightsDescribeLogPatternResult

       DescribeObservation(ctx workflow.Context, input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error)
       DescribeObservationAsync(ctx workflow.Context, input *applicationinsights.DescribeObservationInput) *ApplicationinsightsDescribeObservationResult

       DescribeProblem(ctx workflow.Context, input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error)
       DescribeProblemAsync(ctx workflow.Context, input *applicationinsights.DescribeProblemInput) *ApplicationinsightsDescribeProblemResult

       DescribeProblemObservations(ctx workflow.Context, input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error)
       DescribeProblemObservationsAsync(ctx workflow.Context, input *applicationinsights.DescribeProblemObservationsInput) *ApplicationinsightsDescribeProblemObservationsResult

       ListApplications(ctx workflow.Context, input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error)
       ListApplicationsAsync(ctx workflow.Context, input *applicationinsights.ListApplicationsInput) *ApplicationinsightsListApplicationsResult

       ListComponents(ctx workflow.Context, input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error)
       ListComponentsAsync(ctx workflow.Context, input *applicationinsights.ListComponentsInput) *ApplicationinsightsListComponentsResult

       ListConfigurationHistory(ctx workflow.Context, input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error)
       ListConfigurationHistoryAsync(ctx workflow.Context, input *applicationinsights.ListConfigurationHistoryInput) *ApplicationinsightsListConfigurationHistoryResult

       ListLogPatternSets(ctx workflow.Context, input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error)
       ListLogPatternSetsAsync(ctx workflow.Context, input *applicationinsights.ListLogPatternSetsInput) *ApplicationinsightsListLogPatternSetsResult

       ListLogPatterns(ctx workflow.Context, input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error)
       ListLogPatternsAsync(ctx workflow.Context, input *applicationinsights.ListLogPatternsInput) *ApplicationinsightsListLogPatternsResult

       ListProblems(ctx workflow.Context, input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error)
       ListProblemsAsync(ctx workflow.Context, input *applicationinsights.ListProblemsInput) *ApplicationinsightsListProblemsResult

       ListTagsForResource(ctx workflow.Context, input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *applicationinsights.ListTagsForResourceInput) *ApplicationinsightsListTagsForResourceResult

       TagResource(ctx workflow.Context, input *applicationinsights.TagResourceInput) (*applicationinsights.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *applicationinsights.TagResourceInput) *ApplicationinsightsTagResourceResult

       UntagResource(ctx workflow.Context, input *applicationinsights.UntagResourceInput) (*applicationinsights.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *applicationinsights.UntagResourceInput) *ApplicationinsightsUntagResourceResult

       UpdateApplication(ctx workflow.Context, input *applicationinsights.UpdateApplicationInput) (*applicationinsights.UpdateApplicationOutput, error)
       UpdateApplicationAsync(ctx workflow.Context, input *applicationinsights.UpdateApplicationInput) *ApplicationinsightsUpdateApplicationResult

       UpdateComponent(ctx workflow.Context, input *applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error)
       UpdateComponentAsync(ctx workflow.Context, input *applicationinsights.UpdateComponentInput) *ApplicationinsightsUpdateComponentResult

       UpdateComponentConfiguration(ctx workflow.Context, input *applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error)
       UpdateComponentConfigurationAsync(ctx workflow.Context, input *applicationinsights.UpdateComponentConfigurationInput) *ApplicationinsightsUpdateComponentConfigurationResult

       UpdateLogPattern(ctx workflow.Context, input *applicationinsights.UpdateLogPatternInput) (*applicationinsights.UpdateLogPatternOutput, error)
       UpdateLogPatternAsync(ctx workflow.Context, input *applicationinsights.UpdateLogPatternInput) *ApplicationinsightsUpdateLogPatternResult
}

type ApplicationInsightsStub struct {
}

func NewApplicationInsightsStub() ApplicationInsightsClient {
    return &ApplicationInsightsStub{}
}

type ApplicationinsightsCreateApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsCreateApplicationResult) Get(ctx workflow.Context) (*applicationinsights.CreateApplicationOutput, error) {
    var output applicationinsights.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsCreateComponentResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsCreateComponentResult) Get(ctx workflow.Context) (*applicationinsights.CreateComponentOutput, error) {
    var output applicationinsights.CreateComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsCreateLogPatternResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsCreateLogPatternResult) Get(ctx workflow.Context) (*applicationinsights.CreateLogPatternOutput, error) {
    var output applicationinsights.CreateLogPatternOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDeleteApplicationResult) Get(ctx workflow.Context) (*applicationinsights.DeleteApplicationOutput, error) {
    var output applicationinsights.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDeleteComponentResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDeleteComponentResult) Get(ctx workflow.Context) (*applicationinsights.DeleteComponentOutput, error) {
    var output applicationinsights.DeleteComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDeleteLogPatternResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDeleteLogPatternResult) Get(ctx workflow.Context) (*applicationinsights.DeleteLogPatternOutput, error) {
    var output applicationinsights.DeleteLogPatternOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeApplicationResult) Get(ctx workflow.Context) (*applicationinsights.DescribeApplicationOutput, error) {
    var output applicationinsights.DescribeApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeComponentResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeComponentResult) Get(ctx workflow.Context) (*applicationinsights.DescribeComponentOutput, error) {
    var output applicationinsights.DescribeComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeComponentConfigurationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeComponentConfigurationResult) Get(ctx workflow.Context) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
    var output applicationinsights.DescribeComponentConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeComponentConfigurationRecommendationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeComponentConfigurationRecommendationResult) Get(ctx workflow.Context) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
    var output applicationinsights.DescribeComponentConfigurationRecommendationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeLogPatternResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeLogPatternResult) Get(ctx workflow.Context) (*applicationinsights.DescribeLogPatternOutput, error) {
    var output applicationinsights.DescribeLogPatternOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeObservationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeObservationResult) Get(ctx workflow.Context) (*applicationinsights.DescribeObservationOutput, error) {
    var output applicationinsights.DescribeObservationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeProblemResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeProblemResult) Get(ctx workflow.Context) (*applicationinsights.DescribeProblemOutput, error) {
    var output applicationinsights.DescribeProblemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsDescribeProblemObservationsResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsDescribeProblemObservationsResult) Get(ctx workflow.Context) (*applicationinsights.DescribeProblemObservationsOutput, error) {
    var output applicationinsights.DescribeProblemObservationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListApplicationsResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListApplicationsResult) Get(ctx workflow.Context) (*applicationinsights.ListApplicationsOutput, error) {
    var output applicationinsights.ListApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListComponentsResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListComponentsResult) Get(ctx workflow.Context) (*applicationinsights.ListComponentsOutput, error) {
    var output applicationinsights.ListComponentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListConfigurationHistoryResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListConfigurationHistoryResult) Get(ctx workflow.Context) (*applicationinsights.ListConfigurationHistoryOutput, error) {
    var output applicationinsights.ListConfigurationHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListLogPatternSetsResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListLogPatternSetsResult) Get(ctx workflow.Context) (*applicationinsights.ListLogPatternSetsOutput, error) {
    var output applicationinsights.ListLogPatternSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListLogPatternsResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListLogPatternsResult) Get(ctx workflow.Context) (*applicationinsights.ListLogPatternsOutput, error) {
    var output applicationinsights.ListLogPatternsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListProblemsResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListProblemsResult) Get(ctx workflow.Context) (*applicationinsights.ListProblemsOutput, error) {
    var output applicationinsights.ListProblemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsListTagsForResourceResult) Get(ctx workflow.Context) (*applicationinsights.ListTagsForResourceOutput, error) {
    var output applicationinsights.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsTagResourceResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsTagResourceResult) Get(ctx workflow.Context) (*applicationinsights.TagResourceOutput, error) {
    var output applicationinsights.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsUntagResourceResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsUntagResourceResult) Get(ctx workflow.Context) (*applicationinsights.UntagResourceOutput, error) {
    var output applicationinsights.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsUpdateApplicationResult) Get(ctx workflow.Context) (*applicationinsights.UpdateApplicationOutput, error) {
    var output applicationinsights.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsUpdateComponentResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsUpdateComponentResult) Get(ctx workflow.Context) (*applicationinsights.UpdateComponentOutput, error) {
    var output applicationinsights.UpdateComponentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsUpdateComponentConfigurationResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsUpdateComponentConfigurationResult) Get(ctx workflow.Context) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
    var output applicationinsights.UpdateComponentConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type ApplicationinsightsUpdateLogPatternResult struct {
	Result workflow.Future
}

func (r *ApplicationinsightsUpdateLogPatternResult) Get(ctx workflow.Context) (*applicationinsights.UpdateLogPatternOutput, error) {
    var output applicationinsights.UpdateLogPatternOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) CreateApplication(ctx workflow.Context, input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error) {
    var output applicationinsights.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.CreateApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) CreateApplicationAsync(ctx workflow.Context, input *applicationinsights.CreateApplicationInput) *ApplicationinsightsCreateApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.CreateApplication", input)
    return &ApplicationinsightsCreateApplicationResult{Result: future}
}

func (a *ApplicationInsightsStub) CreateComponent(ctx workflow.Context, input *applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error) {
    var output applicationinsights.CreateComponentOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.CreateComponent", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) CreateComponentAsync(ctx workflow.Context, input *applicationinsights.CreateComponentInput) *ApplicationinsightsCreateComponentResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.CreateComponent", input)
    return &ApplicationinsightsCreateComponentResult{Result: future}
}

func (a *ApplicationInsightsStub) CreateLogPattern(ctx workflow.Context, input *applicationinsights.CreateLogPatternInput) (*applicationinsights.CreateLogPatternOutput, error) {
    var output applicationinsights.CreateLogPatternOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.CreateLogPattern", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) CreateLogPatternAsync(ctx workflow.Context, input *applicationinsights.CreateLogPatternInput) *ApplicationinsightsCreateLogPatternResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.CreateLogPattern", input)
    return &ApplicationinsightsCreateLogPatternResult{Result: future}
}

func (a *ApplicationInsightsStub) DeleteApplication(ctx workflow.Context, input *applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error) {
    var output applicationinsights.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DeleteApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DeleteApplicationAsync(ctx workflow.Context, input *applicationinsights.DeleteApplicationInput) *ApplicationinsightsDeleteApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DeleteApplication", input)
    return &ApplicationinsightsDeleteApplicationResult{Result: future}
}

func (a *ApplicationInsightsStub) DeleteComponent(ctx workflow.Context, input *applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error) {
    var output applicationinsights.DeleteComponentOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DeleteComponent", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DeleteComponentAsync(ctx workflow.Context, input *applicationinsights.DeleteComponentInput) *ApplicationinsightsDeleteComponentResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DeleteComponent", input)
    return &ApplicationinsightsDeleteComponentResult{Result: future}
}

func (a *ApplicationInsightsStub) DeleteLogPattern(ctx workflow.Context, input *applicationinsights.DeleteLogPatternInput) (*applicationinsights.DeleteLogPatternOutput, error) {
    var output applicationinsights.DeleteLogPatternOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DeleteLogPattern", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DeleteLogPatternAsync(ctx workflow.Context, input *applicationinsights.DeleteLogPatternInput) *ApplicationinsightsDeleteLogPatternResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DeleteLogPattern", input)
    return &ApplicationinsightsDeleteLogPatternResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeApplication(ctx workflow.Context, input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error) {
    var output applicationinsights.DescribeApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeApplicationAsync(ctx workflow.Context, input *applicationinsights.DescribeApplicationInput) *ApplicationinsightsDescribeApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeApplication", input)
    return &ApplicationinsightsDescribeApplicationResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeComponent(ctx workflow.Context, input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error) {
    var output applicationinsights.DescribeComponentOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeComponent", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeComponentAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentInput) *ApplicationinsightsDescribeComponentResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeComponent", input)
    return &ApplicationinsightsDescribeComponentResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeComponentConfiguration(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
    var output applicationinsights.DescribeComponentConfigurationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeComponentConfiguration", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeComponentConfigurationAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationInput) *ApplicationinsightsDescribeComponentConfigurationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeComponentConfiguration", input)
    return &ApplicationinsightsDescribeComponentConfigurationResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeComponentConfigurationRecommendation(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
    var output applicationinsights.DescribeComponentConfigurationRecommendationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeComponentConfigurationRecommendation", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeComponentConfigurationRecommendationAsync(ctx workflow.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) *ApplicationinsightsDescribeComponentConfigurationRecommendationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeComponentConfigurationRecommendation", input)
    return &ApplicationinsightsDescribeComponentConfigurationRecommendationResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeLogPattern(ctx workflow.Context, input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error) {
    var output applicationinsights.DescribeLogPatternOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeLogPattern", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeLogPatternAsync(ctx workflow.Context, input *applicationinsights.DescribeLogPatternInput) *ApplicationinsightsDescribeLogPatternResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeLogPattern", input)
    return &ApplicationinsightsDescribeLogPatternResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeObservation(ctx workflow.Context, input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error) {
    var output applicationinsights.DescribeObservationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeObservation", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeObservationAsync(ctx workflow.Context, input *applicationinsights.DescribeObservationInput) *ApplicationinsightsDescribeObservationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeObservation", input)
    return &ApplicationinsightsDescribeObservationResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeProblem(ctx workflow.Context, input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error) {
    var output applicationinsights.DescribeProblemOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeProblem", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeProblemAsync(ctx workflow.Context, input *applicationinsights.DescribeProblemInput) *ApplicationinsightsDescribeProblemResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeProblem", input)
    return &ApplicationinsightsDescribeProblemResult{Result: future}
}

func (a *ApplicationInsightsStub) DescribeProblemObservations(ctx workflow.Context, input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error) {
    var output applicationinsights.DescribeProblemObservationsOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeProblemObservations", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) DescribeProblemObservationsAsync(ctx workflow.Context, input *applicationinsights.DescribeProblemObservationsInput) *ApplicationinsightsDescribeProblemObservationsResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.DescribeProblemObservations", input)
    return &ApplicationinsightsDescribeProblemObservationsResult{Result: future}
}

func (a *ApplicationInsightsStub) ListApplications(ctx workflow.Context, input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error) {
    var output applicationinsights.ListApplicationsOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListApplications", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListApplicationsAsync(ctx workflow.Context, input *applicationinsights.ListApplicationsInput) *ApplicationinsightsListApplicationsResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListApplications", input)
    return &ApplicationinsightsListApplicationsResult{Result: future}
}

func (a *ApplicationInsightsStub) ListComponents(ctx workflow.Context, input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error) {
    var output applicationinsights.ListComponentsOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListComponents", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListComponentsAsync(ctx workflow.Context, input *applicationinsights.ListComponentsInput) *ApplicationinsightsListComponentsResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListComponents", input)
    return &ApplicationinsightsListComponentsResult{Result: future}
}

func (a *ApplicationInsightsStub) ListConfigurationHistory(ctx workflow.Context, input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error) {
    var output applicationinsights.ListConfigurationHistoryOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListConfigurationHistory", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListConfigurationHistoryAsync(ctx workflow.Context, input *applicationinsights.ListConfigurationHistoryInput) *ApplicationinsightsListConfigurationHistoryResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListConfigurationHistory", input)
    return &ApplicationinsightsListConfigurationHistoryResult{Result: future}
}

func (a *ApplicationInsightsStub) ListLogPatternSets(ctx workflow.Context, input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error) {
    var output applicationinsights.ListLogPatternSetsOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListLogPatternSets", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListLogPatternSetsAsync(ctx workflow.Context, input *applicationinsights.ListLogPatternSetsInput) *ApplicationinsightsListLogPatternSetsResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListLogPatternSets", input)
    return &ApplicationinsightsListLogPatternSetsResult{Result: future}
}

func (a *ApplicationInsightsStub) ListLogPatterns(ctx workflow.Context, input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error) {
    var output applicationinsights.ListLogPatternsOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListLogPatterns", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListLogPatternsAsync(ctx workflow.Context, input *applicationinsights.ListLogPatternsInput) *ApplicationinsightsListLogPatternsResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListLogPatterns", input)
    return &ApplicationinsightsListLogPatternsResult{Result: future}
}

func (a *ApplicationInsightsStub) ListProblems(ctx workflow.Context, input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error) {
    var output applicationinsights.ListProblemsOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListProblems", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListProblemsAsync(ctx workflow.Context, input *applicationinsights.ListProblemsInput) *ApplicationinsightsListProblemsResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListProblems", input)
    return &ApplicationinsightsListProblemsResult{Result: future}
}

func (a *ApplicationInsightsStub) ListTagsForResource(ctx workflow.Context, input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error) {
    var output applicationinsights.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) ListTagsForResourceAsync(ctx workflow.Context, input *applicationinsights.ListTagsForResourceInput) *ApplicationinsightsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.ListTagsForResource", input)
    return &ApplicationinsightsListTagsForResourceResult{Result: future}
}

func (a *ApplicationInsightsStub) TagResource(ctx workflow.Context, input *applicationinsights.TagResourceInput) (*applicationinsights.TagResourceOutput, error) {
    var output applicationinsights.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) TagResourceAsync(ctx workflow.Context, input *applicationinsights.TagResourceInput) *ApplicationinsightsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.TagResource", input)
    return &ApplicationinsightsTagResourceResult{Result: future}
}

func (a *ApplicationInsightsStub) UntagResource(ctx workflow.Context, input *applicationinsights.UntagResourceInput) (*applicationinsights.UntagResourceOutput, error) {
    var output applicationinsights.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) UntagResourceAsync(ctx workflow.Context, input *applicationinsights.UntagResourceInput) *ApplicationinsightsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.UntagResource", input)
    return &ApplicationinsightsUntagResourceResult{Result: future}
}

func (a *ApplicationInsightsStub) UpdateApplication(ctx workflow.Context, input *applicationinsights.UpdateApplicationInput) (*applicationinsights.UpdateApplicationOutput, error) {
    var output applicationinsights.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateApplication", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) UpdateApplicationAsync(ctx workflow.Context, input *applicationinsights.UpdateApplicationInput) *ApplicationinsightsUpdateApplicationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateApplication", input)
    return &ApplicationinsightsUpdateApplicationResult{Result: future}
}

func (a *ApplicationInsightsStub) UpdateComponent(ctx workflow.Context, input *applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error) {
    var output applicationinsights.UpdateComponentOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateComponent", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) UpdateComponentAsync(ctx workflow.Context, input *applicationinsights.UpdateComponentInput) *ApplicationinsightsUpdateComponentResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateComponent", input)
    return &ApplicationinsightsUpdateComponentResult{Result: future}
}

func (a *ApplicationInsightsStub) UpdateComponentConfiguration(ctx workflow.Context, input *applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
    var output applicationinsights.UpdateComponentConfigurationOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateComponentConfiguration", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) UpdateComponentConfigurationAsync(ctx workflow.Context, input *applicationinsights.UpdateComponentConfigurationInput) *ApplicationinsightsUpdateComponentConfigurationResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateComponentConfiguration", input)
    return &ApplicationinsightsUpdateComponentConfigurationResult{Result: future}
}

func (a *ApplicationInsightsStub) UpdateLogPattern(ctx workflow.Context, input *applicationinsights.UpdateLogPatternInput) (*applicationinsights.UpdateLogPatternOutput, error) {
    var output applicationinsights.UpdateLogPatternOutput
    err := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateLogPattern", input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationInsightsStub) UpdateLogPatternAsync(ctx workflow.Context, input *applicationinsights.UpdateLogPatternInput) *ApplicationinsightsUpdateLogPatternResult {
    future := workflow.ExecuteActivity(ctx, "ApplicationInsights.UpdateLogPattern", input)
    return &ApplicationinsightsUpdateLogPatternResult{Result: future}
}
