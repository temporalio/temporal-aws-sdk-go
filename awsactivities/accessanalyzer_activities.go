package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"github.com/aws/aws-sdk-go/service/accessanalyzer/accessanalyzeriface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type AccessAnalyzerActivities struct {
	client accessanalyzeriface.AccessAnalyzerAPI
}

func NewAccessAnalyzerActivities(session *session.Session, config ...*aws.Config) *AccessAnalyzerActivities {
	client := accessanalyzer.New(session, config...)
	return &AccessAnalyzerActivities{client: client}
}

func (a *AccessAnalyzerActivities) CreateAnalyzer(ctx context.Context, input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateAnalyzerWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) CreateArchiveRule(ctx context.Context, input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateArchiveRuleWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) DeleteAnalyzer(ctx context.Context, input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.DeleteAnalyzerWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) DeleteArchiveRule(ctx context.Context, input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.DeleteArchiveRuleWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) GetAnalyzedResource(ctx context.Context, input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	return a.client.GetAnalyzedResourceWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) GetAnalyzer(ctx context.Context, input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error) {
	return a.client.GetAnalyzerWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) GetArchiveRule(ctx context.Context, input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error) {
	return a.client.GetArchiveRuleWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) GetFinding(ctx context.Context, input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error) {
	return a.client.GetFindingWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) ListAnalyzedResources(ctx context.Context, input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	return a.client.ListAnalyzedResourcesWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) ListAnalyzers(ctx context.Context, input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error) {
	return a.client.ListAnalyzersWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) ListArchiveRules(ctx context.Context, input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error) {
	return a.client.ListArchiveRulesWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) ListFindings(ctx context.Context, input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error) {
	return a.client.ListFindingsWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) ListTagsForResource(ctx context.Context, input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) StartResourceScan(ctx context.Context, input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error) {
	return a.client.StartResourceScanWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) TagResource(ctx context.Context, input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) UntagResource(ctx context.Context, input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) UpdateArchiveRule(ctx context.Context, input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.UpdateArchiveRuleWithContext(ctx, input)
}

func (a *AccessAnalyzerActivities) UpdateFindings(ctx context.Context, input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.UpdateFindingsWithContext(ctx, input)
}
