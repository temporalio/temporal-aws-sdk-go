// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package dlmstub

import (
	"github.com/aws/aws-sdk-go/service/dlm"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type DLMCreateLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMCreateLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.CreateLifecyclePolicyOutput, error) {
	var output dlm.CreateLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMDeleteLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMDeleteLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.DeleteLifecyclePolicyOutput, error) {
	var output dlm.DeleteLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMGetLifecyclePoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMGetLifecyclePoliciesFuture) Get(ctx workflow.Context) (*dlm.GetLifecyclePoliciesOutput, error) {
	var output dlm.GetLifecyclePoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMGetLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMGetLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.GetLifecyclePolicyOutput, error) {
	var output dlm.GetLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMListTagsForResourceFuture) Get(ctx workflow.Context) (*dlm.ListTagsForResourceOutput, error) {
	var output dlm.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMTagResourceFuture) Get(ctx workflow.Context) (*dlm.TagResourceOutput, error) {
	var output dlm.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMUntagResourceFuture) Get(ctx workflow.Context) (*dlm.UntagResourceOutput, error) {
	var output dlm.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DLMUpdateLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DLMUpdateLifecyclePolicyFuture) Get(ctx workflow.Context) (*dlm.UpdateLifecyclePolicyOutput, error) {
	var output dlm.UpdateLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLifecyclePolicy(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
	var output dlm.CreateLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.CreateLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) *DLMCreateLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.CreateLifecyclePolicy", input)
	return &DLMCreateLifecyclePolicyFuture{Future: future}
}

func (a *stub) DeleteLifecyclePolicy(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error) {
	var output dlm.DeleteLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.DeleteLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLifecyclePolicyAsync(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) *DLMDeleteLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.DeleteLifecyclePolicy", input)
	return &DLMDeleteLifecyclePolicyFuture{Future: future}
}

func (a *stub) GetLifecyclePolicies(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
	var output dlm.GetLifecyclePoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.GetLifecyclePolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLifecyclePoliciesAsync(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) *DLMGetLifecyclePoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.GetLifecyclePolicies", input)
	return &DLMGetLifecyclePoliciesFuture{Future: future}
}

func (a *stub) GetLifecyclePolicy(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
	var output dlm.GetLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.GetLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLifecyclePolicyAsync(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) *DLMGetLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.GetLifecyclePolicy", input)
	return &DLMGetLifecyclePolicyFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
	var output dlm.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *dlm.ListTagsForResourceInput) *DLMListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.ListTagsForResource", input)
	return &DLMListTagsForResourceFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error) {
	var output dlm.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *dlm.TagResourceInput) *DLMTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.TagResource", input)
	return &DLMTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error) {
	var output dlm.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *dlm.UntagResourceInput) *DLMUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.UntagResource", input)
	return &DLMUntagResourceFuture{Future: future}
}

func (a *stub) UpdateLifecyclePolicy(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error) {
	var output dlm.UpdateLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.dlm.UpdateLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) *DLMUpdateLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dlm.UpdateLifecyclePolicy", input)
	return &DLMUpdateLifecyclePolicyFuture{Future: future}
}
