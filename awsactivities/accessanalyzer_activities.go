package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"github.com/aws/aws-sdk-go/service/accessanalyzer/accessanalyzeriface"
)

type AccessAnalyzerActivities struct {
	client accessanalyzeriface.AccessAnalyzerAPI
}

func NewAccessAnalyzerActivities(session *session.Session, config ...*aws.Config) *AccessAnalyzerActivities {
	client := accessanalyzer.New(session, config...)
	return &AccessAnalyzerActivities{client: client}
}

func (a *AccessAnalyzerActivities) CreateAnalyzer(input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error) {
	return a.client.CreateAnalyzer(input)
}

func (a *AccessAnalyzerActivities) CreateArchiveRule(input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error) {
	return a.client.CreateArchiveRule(input)
}

func (a *AccessAnalyzerActivities) DeleteAnalyzer(input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error) {
	return a.client.DeleteAnalyzer(input)
}

func (a *AccessAnalyzerActivities) DeleteArchiveRule(input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error) {
	return a.client.DeleteArchiveRule(input)
}

func (a *AccessAnalyzerActivities) GetAnalyzedResource(input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	return a.client.GetAnalyzedResource(input)
}

func (a *AccessAnalyzerActivities) GetAnalyzer(input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error) {
	return a.client.GetAnalyzer(input)
}

func (a *AccessAnalyzerActivities) GetArchiveRule(input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error) {
	return a.client.GetArchiveRule(input)
}

func (a *AccessAnalyzerActivities) GetFinding(input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error) {
	return a.client.GetFinding(input)
}

func (a *AccessAnalyzerActivities) ListAnalyzedResources(input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	return a.client.ListAnalyzedResources(input)
}

func (a *AccessAnalyzerActivities) ListAnalyzers(input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error) {
	return a.client.ListAnalyzers(input)
}

func (a *AccessAnalyzerActivities) ListArchiveRules(input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error) {
	return a.client.ListArchiveRules(input)
}

func (a *AccessAnalyzerActivities) ListFindings(input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error) {
	return a.client.ListFindings(input)
}

func (a *AccessAnalyzerActivities) ListTagsForResource(input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *AccessAnalyzerActivities) StartResourceScan(input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error) {
	return a.client.StartResourceScan(input)
}

func (a *AccessAnalyzerActivities) TagResource(input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *AccessAnalyzerActivities) UntagResource(input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *AccessAnalyzerActivities) UpdateArchiveRule(input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error) {
	return a.client.UpdateArchiveRule(input)
}

func (a *AccessAnalyzerActivities) UpdateFindings(input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error) {
	return a.client.UpdateFindings(input)
}
