// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotsecuretunnelingstub

import (
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CloseTunnelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CloseTunnelFuture) Get(ctx workflow.Context) (*iotsecuretunneling.CloseTunnelOutput, error) {
	var output iotsecuretunneling.CloseTunnelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeTunnelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeTunnelFuture) Get(ctx workflow.Context) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	var output iotsecuretunneling.DescribeTunnelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	var output iotsecuretunneling.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTunnelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTunnelsFuture) Get(ctx workflow.Context) (*iotsecuretunneling.ListTunnelsOutput, error) {
	var output iotsecuretunneling.ListTunnelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type OpenTunnelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *OpenTunnelFuture) Get(ctx workflow.Context) (*iotsecuretunneling.OpenTunnelOutput, error) {
	var output iotsecuretunneling.OpenTunnelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*iotsecuretunneling.TagResourceOutput, error) {
	var output iotsecuretunneling.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*iotsecuretunneling.UntagResourceOutput, error) {
	var output iotsecuretunneling.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CloseTunnel(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
	var output iotsecuretunneling.CloseTunnelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.CloseTunnel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CloseTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.CloseTunnelInput) *CloseTunnelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.CloseTunnel", input)
	return &CloseTunnelFuture{Future: future}
}

func (a *stub) DescribeTunnel(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	var output iotsecuretunneling.DescribeTunnelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.DescribeTunnel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.DescribeTunnelInput) *DescribeTunnelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.DescribeTunnel", input)
	return &DescribeTunnelFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	var output iotsecuretunneling.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *iotsecuretunneling.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) ListTunnels(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
	var output iotsecuretunneling.ListTunnelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.ListTunnels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTunnelsAsync(ctx workflow.Context, input *iotsecuretunneling.ListTunnelsInput) *ListTunnelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.ListTunnels", input)
	return &ListTunnelsFuture{Future: future}
}

func (a *stub) OpenTunnel(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error) {
	var output iotsecuretunneling.OpenTunnelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.OpenTunnel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) OpenTunnelAsync(ctx workflow.Context, input *iotsecuretunneling.OpenTunnelInput) *OpenTunnelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.OpenTunnel", input)
	return &OpenTunnelFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error) {
	var output iotsecuretunneling.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *iotsecuretunneling.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error) {
	var output iotsecuretunneling.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *iotsecuretunneling.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotsecuretunneling.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}
