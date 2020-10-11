// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package accessanalyzerstub

import (
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type Client interface {
	CreateAnalyzer(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) (*accessanalyzer.CreateAnalyzerOutput, error)
	CreateAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.CreateAnalyzerInput) *AccessAnalyzerCreateAnalyzerFuture

	CreateArchiveRule(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) (*accessanalyzer.CreateArchiveRuleOutput, error)
	CreateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.CreateArchiveRuleInput) *AccessAnalyzerCreateArchiveRuleFuture

	DeleteAnalyzer(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) (*accessanalyzer.DeleteAnalyzerOutput, error)
	DeleteAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.DeleteAnalyzerInput) *AccessAnalyzerDeleteAnalyzerFuture

	DeleteArchiveRule(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) (*accessanalyzer.DeleteArchiveRuleOutput, error)
	DeleteArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.DeleteArchiveRuleInput) *AccessAnalyzerDeleteArchiveRuleFuture

	GetAnalyzedResource(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error)
	GetAnalyzedResourceAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzedResourceInput) *AccessAnalyzerGetAnalyzedResourceFuture

	GetAnalyzer(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error)
	GetAnalyzerAsync(ctx workflow.Context, input *accessanalyzer.GetAnalyzerInput) *AccessAnalyzerGetAnalyzerFuture

	GetArchiveRule(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error)
	GetArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.GetArchiveRuleInput) *AccessAnalyzerGetArchiveRuleFuture

	GetFinding(ctx workflow.Context, input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error)
	GetFindingAsync(ctx workflow.Context, input *accessanalyzer.GetFindingInput) *AccessAnalyzerGetFindingFuture

	ListAnalyzedResources(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error)
	ListAnalyzedResourcesAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzedResourcesInput) *AccessAnalyzerListAnalyzedResourcesFuture

	ListAnalyzers(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error)
	ListAnalyzersAsync(ctx workflow.Context, input *accessanalyzer.ListAnalyzersInput) *AccessAnalyzerListAnalyzersFuture

	ListArchiveRules(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error)
	ListArchiveRulesAsync(ctx workflow.Context, input *accessanalyzer.ListArchiveRulesInput) *AccessAnalyzerListArchiveRulesFuture

	ListFindings(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error)
	ListFindingsAsync(ctx workflow.Context, input *accessanalyzer.ListFindingsInput) *AccessAnalyzerListFindingsFuture

	ListTagsForResource(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *accessanalyzer.ListTagsForResourceInput) *AccessAnalyzerListTagsForResourceFuture

	StartResourceScan(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) (*accessanalyzer.StartResourceScanOutput, error)
	StartResourceScanAsync(ctx workflow.Context, input *accessanalyzer.StartResourceScanInput) *AccessAnalyzerStartResourceScanFuture

	TagResource(ctx workflow.Context, input *accessanalyzer.TagResourceInput) (*accessanalyzer.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *accessanalyzer.TagResourceInput) *AccessAnalyzerTagResourceFuture

	UntagResource(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) (*accessanalyzer.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *accessanalyzer.UntagResourceInput) *AccessAnalyzerUntagResourceFuture

	UpdateArchiveRule(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) (*accessanalyzer.UpdateArchiveRuleOutput, error)
	UpdateArchiveRuleAsync(ctx workflow.Context, input *accessanalyzer.UpdateArchiveRuleInput) *AccessAnalyzerUpdateArchiveRuleFuture

	UpdateFindings(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) (*accessanalyzer.UpdateFindingsOutput, error)
	UpdateFindingsAsync(ctx workflow.Context, input *accessanalyzer.UpdateFindingsInput) *AccessAnalyzerUpdateFindingsFuture
}

func NewClient() Client {
	return &stub{}
}
