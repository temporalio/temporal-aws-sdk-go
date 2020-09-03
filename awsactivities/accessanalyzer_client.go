package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"go.temporal.io/sdk/workflow"
)

type AccessAnalyzerClient interface {
    CreateAnalyzer(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error)
    CreateAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) *AccessanalyzerCreateAnalyzerResult

    CreateArchiveRule(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error)
    CreateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) *AccessanalyzerCreateArchiveRuleResult

    DeleteAnalyzer(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error)
    DeleteAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) *AccessanalyzerDeleteAnalyzerResult

    DeleteArchiveRule(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error)
    DeleteArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) *AccessanalyzerDeleteArchiveRuleResult

    GetAnalyzedResource(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error)
    GetAnalyzedResourceAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) *AccessanalyzerGetAnalyzedResourceResult

    GetAnalyzer(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error)
    GetAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) *AccessanalyzerGetAnalyzerResult

    GetArchiveRule(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error)
    GetArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) *AccessanalyzerGetArchiveRuleResult

    GetFinding(ctx workflow.Context, input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error)
    GetFindingAsync(ctx workflow.Context, input *accessanalyzer.GetFindingInput) *AccessanalyzerGetFindingResult

    ListAnalyzedResources(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error)
    ListAnalyzedResourcesAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) *AccessanalyzerListAnalyzedResourcesResult

    ListAnalyzers(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error)
    ListAnalyzersAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) *AccessanalyzerListAnalyzersResult

    ListArchiveRules(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error)
    ListArchiveRulesAsync(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) *AccessanalyzerListArchiveRulesResult

    ListFindings(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error)
    ListFindingsAsync(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) *AccessanalyzerListFindingsResult

    ListTagsForResource(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) *AccessanalyzerListTagsForResourceResult

    StartResourceScan(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error)
    StartResourceScanAsync(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) *AccessanalyzerStartResourceScanResult

    TagResource(ctx workflow.Context, input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *accessanalyzer.TagResourceInput) *AccessanalyzerTagResourceResult

    UntagResource(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) *AccessanalyzerUntagResourceResult

    UpdateArchiveRule(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error)
    UpdateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) *AccessanalyzerUpdateArchiveRuleResult

    UpdateFindings(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error)
    UpdateFindingsAsync(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) *AccessanalyzerUpdateFindingsResult
}
type AccessanalyzerCreateAnalyzerResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerCreateAnalyzerResult) Get(ctx workflow.Context) (*accessanalyzer.CreateAnalyzerOutput, error) {
    var output accessanalyzer.CreateAnalyzerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerCreateArchiveRuleResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerCreateArchiveRuleResult) Get(ctx workflow.Context) (*accessanalyzer.CreateArchiveRuleOutput, error) {
    var output accessanalyzer.CreateArchiveRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerDeleteAnalyzerResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerDeleteAnalyzerResult) Get(ctx workflow.Context) (*accessanalyzer.DeleteAnalyzerOutput, error) {
    var output accessanalyzer.DeleteAnalyzerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerDeleteArchiveRuleResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerDeleteArchiveRuleResult) Get(ctx workflow.Context) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
    var output accessanalyzer.DeleteArchiveRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerGetAnalyzedResourceResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerGetAnalyzedResourceResult) Get(ctx workflow.Context) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
    var output accessanalyzer.GetAnalyzedResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerGetAnalyzerResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerGetAnalyzerResult) Get(ctx workflow.Context) (*accessanalyzer.GetAnalyzerOutput, error) {
    var output accessanalyzer.GetAnalyzerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerGetArchiveRuleResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerGetArchiveRuleResult) Get(ctx workflow.Context) (*accessanalyzer.GetArchiveRuleOutput, error) {
    var output accessanalyzer.GetArchiveRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerGetFindingResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerGetFindingResult) Get(ctx workflow.Context) (*accessanalyzer.GetFindingOutput, error) {
    var output accessanalyzer.GetFindingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerListAnalyzedResourcesResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerListAnalyzedResourcesResult) Get(ctx workflow.Context) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
    var output accessanalyzer.ListAnalyzedResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerListAnalyzersResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerListAnalyzersResult) Get(ctx workflow.Context) (*accessanalyzer.ListAnalyzersOutput, error) {
    var output accessanalyzer.ListAnalyzersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerListArchiveRulesResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerListArchiveRulesResult) Get(ctx workflow.Context) (*accessanalyzer.ListArchiveRulesOutput, error) {
    var output accessanalyzer.ListArchiveRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerListFindingsResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerListFindingsResult) Get(ctx workflow.Context) (*accessanalyzer.ListFindingsOutput, error) {
    var output accessanalyzer.ListFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerListTagsForResourceResult) Get(ctx workflow.Context) (*accessanalyzer.ListTagsForResourceOutput, error) {
    var output accessanalyzer.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerStartResourceScanResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerStartResourceScanResult) Get(ctx workflow.Context) (*accessanalyzer.StartResourceScanOutput, error) {
    var output accessanalyzer.StartResourceScanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerTagResourceResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerTagResourceResult) Get(ctx workflow.Context) (*accessanalyzer.TagResourceOutput, error) {
    var output accessanalyzer.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerUntagResourceResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerUntagResourceResult) Get(ctx workflow.Context) (*accessanalyzer.UntagResourceOutput, error) {
    var output accessanalyzer.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerUpdateArchiveRuleResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerUpdateArchiveRuleResult) Get(ctx workflow.Context) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
    var output accessanalyzer.UpdateArchiveRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AccessanalyzerUpdateFindingsResult struct {
	Result workflow.Future
}

func (r *AccessanalyzerUpdateFindingsResult) Get(ctx workflow.Context) (*accessanalyzer.UpdateFindingsOutput, error) {
    var output accessanalyzer.UpdateFindingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AccessAnalyzerStub struct {
    activities AccessAnalyzerClient
}

func NewAccessAnalyzerStub() AccessAnalyzerClient {
    return &AccessAnalyzerStub{}
}

func (a *AccessAnalyzerStub) CreateAnalyzer(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error) {
    var output accessanalyzer.CreateAnalyzerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAnalyzer, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) CreateArchiveRule(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error) {
    var output accessanalyzer.CreateArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateArchiveRule, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) DeleteAnalyzer(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error) {
    var output accessanalyzer.DeleteAnalyzerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAnalyzer, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) DeleteArchiveRule(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
    var output accessanalyzer.DeleteArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteArchiveRule, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetAnalyzedResource(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
    var output accessanalyzer.GetAnalyzedResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAnalyzedResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetAnalyzer(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error) {
    var output accessanalyzer.GetAnalyzerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAnalyzer, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetArchiveRule(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error) {
    var output accessanalyzer.GetArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetArchiveRule, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetFinding(ctx workflow.Context, input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error) {
    var output accessanalyzer.GetFindingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFinding, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListAnalyzedResources(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
    var output accessanalyzer.ListAnalyzedResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAnalyzedResources, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListAnalyzers(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error) {
    var output accessanalyzer.ListAnalyzersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAnalyzers, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListArchiveRules(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error) {
    var output accessanalyzer.ListArchiveRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListArchiveRules, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListFindings(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error) {
    var output accessanalyzer.ListFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFindings, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListTagsForResource(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error) {
    var output accessanalyzer.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) StartResourceScan(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error) {
    var output accessanalyzer.StartResourceScanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartResourceScan, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) TagResource(ctx workflow.Context, input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error) {
    var output accessanalyzer.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) UntagResource(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error) {
    var output accessanalyzer.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) UpdateArchiveRule(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
    var output accessanalyzer.UpdateArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateArchiveRule, input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) UpdateFindings(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error) {
    var output accessanalyzer.UpdateFindingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFindings, input).Get(ctx, &output)
    return &output, err
}
