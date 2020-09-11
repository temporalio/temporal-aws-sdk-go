package awsclients

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

type AccessAnalyzerStub struct {
}

func NewAccessAnalyzerStub() AccessAnalyzerClient {
    return &AccessAnalyzerStub{}
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

func (a *AccessAnalyzerStub) CreateAnalyzer(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error) {
    var output accessanalyzer.CreateAnalyzerOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.CreateAnalyzer", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) CreateAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) *AccessanalyzerCreateAnalyzerResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.CreateAnalyzer", input)
    return &AccessanalyzerCreateAnalyzerResult{Result: future}
}

func (a *AccessAnalyzerStub) CreateArchiveRule(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error) {
    var output accessanalyzer.CreateArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.CreateArchiveRule", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) CreateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) *AccessanalyzerCreateArchiveRuleResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.CreateArchiveRule", input)
    return &AccessanalyzerCreateArchiveRuleResult{Result: future}
}

func (a *AccessAnalyzerStub) DeleteAnalyzer(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error) {
    var output accessanalyzer.DeleteAnalyzerOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.DeleteAnalyzer", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) DeleteAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) *AccessanalyzerDeleteAnalyzerResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.DeleteAnalyzer", input)
    return &AccessanalyzerDeleteAnalyzerResult{Result: future}
}

func (a *AccessAnalyzerStub) DeleteArchiveRule(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
    var output accessanalyzer.DeleteArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.DeleteArchiveRule", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) DeleteArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) *AccessanalyzerDeleteArchiveRuleResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.DeleteArchiveRule", input)
    return &AccessanalyzerDeleteArchiveRuleResult{Result: future}
}

func (a *AccessAnalyzerStub) GetAnalyzedResource(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
    var output accessanalyzer.GetAnalyzedResourceOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetAnalyzedResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetAnalyzedResourceAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) *AccessanalyzerGetAnalyzedResourceResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetAnalyzedResource", input)
    return &AccessanalyzerGetAnalyzedResourceResult{Result: future}
}

func (a *AccessAnalyzerStub) GetAnalyzer(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error) {
    var output accessanalyzer.GetAnalyzerOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetAnalyzer", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) *AccessanalyzerGetAnalyzerResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetAnalyzer", input)
    return &AccessanalyzerGetAnalyzerResult{Result: future}
}

func (a *AccessAnalyzerStub) GetArchiveRule(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error) {
    var output accessanalyzer.GetArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetArchiveRule", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) *AccessanalyzerGetArchiveRuleResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetArchiveRule", input)
    return &AccessanalyzerGetArchiveRuleResult{Result: future}
}

func (a *AccessAnalyzerStub) GetFinding(ctx workflow.Context, input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error) {
    var output accessanalyzer.GetFindingOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetFinding", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) GetFindingAsync(ctx workflow.Context, input *accessanalyzer.GetFindingInput) *AccessanalyzerGetFindingResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.GetFinding", input)
    return &AccessanalyzerGetFindingResult{Result: future}
}

func (a *AccessAnalyzerStub) ListAnalyzedResources(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
    var output accessanalyzer.ListAnalyzedResourcesOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListAnalyzedResources", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListAnalyzedResourcesAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) *AccessanalyzerListAnalyzedResourcesResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListAnalyzedResources", input)
    return &AccessanalyzerListAnalyzedResourcesResult{Result: future}
}

func (a *AccessAnalyzerStub) ListAnalyzers(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error) {
    var output accessanalyzer.ListAnalyzersOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListAnalyzers", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListAnalyzersAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) *AccessanalyzerListAnalyzersResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListAnalyzers", input)
    return &AccessanalyzerListAnalyzersResult{Result: future}
}

func (a *AccessAnalyzerStub) ListArchiveRules(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error) {
    var output accessanalyzer.ListArchiveRulesOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListArchiveRules", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListArchiveRulesAsync(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) *AccessanalyzerListArchiveRulesResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListArchiveRules", input)
    return &AccessanalyzerListArchiveRulesResult{Result: future}
}

func (a *AccessAnalyzerStub) ListFindings(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error) {
    var output accessanalyzer.ListFindingsOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListFindings", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListFindingsAsync(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) *AccessanalyzerListFindingsResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListFindings", input)
    return &AccessanalyzerListFindingsResult{Result: future}
}

func (a *AccessAnalyzerStub) ListTagsForResource(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error) {
    var output accessanalyzer.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) ListTagsForResourceAsync(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) *AccessanalyzerListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.ListTagsForResource", input)
    return &AccessanalyzerListTagsForResourceResult{Result: future}
}

func (a *AccessAnalyzerStub) StartResourceScan(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error) {
    var output accessanalyzer.StartResourceScanOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.StartResourceScan", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) StartResourceScanAsync(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) *AccessanalyzerStartResourceScanResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.StartResourceScan", input)
    return &AccessanalyzerStartResourceScanResult{Result: future}
}

func (a *AccessAnalyzerStub) TagResource(ctx workflow.Context, input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error) {
    var output accessanalyzer.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) TagResourceAsync(ctx workflow.Context, input *accessanalyzer.TagResourceInput) *AccessanalyzerTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.TagResource", input)
    return &AccessanalyzerTagResourceResult{Result: future}
}

func (a *AccessAnalyzerStub) UntagResource(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error) {
    var output accessanalyzer.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) UntagResourceAsync(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) *AccessanalyzerUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.UntagResource", input)
    return &AccessanalyzerUntagResourceResult{Result: future}
}

func (a *AccessAnalyzerStub) UpdateArchiveRule(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
    var output accessanalyzer.UpdateArchiveRuleOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.UpdateArchiveRule", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) UpdateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) *AccessanalyzerUpdateArchiveRuleResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.UpdateArchiveRule", input)
    return &AccessanalyzerUpdateArchiveRuleResult{Result: future}
}

func (a *AccessAnalyzerStub) UpdateFindings(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error) {
    var output accessanalyzer.UpdateFindingsOutput
    err := workflow.ExecuteActivity(ctx, "AccessAnalyzer.UpdateFindings", input).Get(ctx, &output)
    return &output, err
}

func (a *AccessAnalyzerStub) UpdateFindingsAsync(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) *AccessanalyzerUpdateFindingsResult {
    future := workflow.ExecuteActivity(ctx, "AccessAnalyzer.UpdateFindings", input)
    return &AccessanalyzerUpdateFindingsResult{Result: future}
}
