// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package daxstub

import (
	"github.com/aws/aws-sdk-go/service/dax"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type DAXCreateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXCreateClusterFuture) Get(ctx workflow.Context) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXCreateParameterGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXCreateParameterGroupFuture) Get(ctx workflow.Context) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXCreateSubnetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXCreateSubnetGroupFuture) Get(ctx workflow.Context) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDecreaseReplicationFactorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDecreaseReplicationFactorFuture) Get(ctx workflow.Context) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDeleteClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDeleteClusterFuture) Get(ctx workflow.Context) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDeleteParameterGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDeleteParameterGroupFuture) Get(ctx workflow.Context) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDeleteSubnetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDeleteSubnetGroupFuture) Get(ctx workflow.Context) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeClustersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDescribeClustersFuture) Get(ctx workflow.Context) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeDefaultParametersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDescribeDefaultParametersFuture) Get(ctx workflow.Context) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDescribeEventsFuture) Get(ctx workflow.Context) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeParameterGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDescribeParameterGroupsFuture) Get(ctx workflow.Context) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeParametersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDescribeParametersFuture) Get(ctx workflow.Context) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeSubnetGroupsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXDescribeSubnetGroupsFuture) Get(ctx workflow.Context) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXIncreaseReplicationFactorFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXIncreaseReplicationFactorFuture) Get(ctx workflow.Context) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXListTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXListTagsFuture) Get(ctx workflow.Context) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXRebootNodeFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXRebootNodeFuture) Get(ctx workflow.Context) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXTagResourceFuture) Get(ctx workflow.Context) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXUntagResourceFuture) Get(ctx workflow.Context) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUpdateClusterFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXUpdateClusterFuture) Get(ctx workflow.Context) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUpdateParameterGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXUpdateParameterGroupFuture) Get(ctx workflow.Context) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUpdateSubnetGroupFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DAXUpdateSubnetGroupFuture) Get(ctx workflow.Context) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *DAXCreateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateCluster", input)
	return &DAXCreateClusterFuture{Future: future}
}

func (a *stub) CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *DAXCreateParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateParameterGroup", input)
	return &DAXCreateParameterGroupFuture{Future: future}
}

func (a *stub) CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *DAXCreateSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateSubnetGroup", input)
	return &DAXCreateSubnetGroupFuture{Future: future}
}

func (a *stub) DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DecreaseReplicationFactor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DAXDecreaseReplicationFactorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DecreaseReplicationFactor", input)
	return &DAXDecreaseReplicationFactorFuture{Future: future}
}

func (a *stub) DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DAXDeleteClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteCluster", input)
	return &DAXDeleteClusterFuture{Future: future}
}

func (a *stub) DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DAXDeleteParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteParameterGroup", input)
	return &DAXDeleteParameterGroupFuture{Future: future}
}

func (a *stub) DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DAXDeleteSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteSubnetGroup", input)
	return &DAXDeleteSubnetGroupFuture{Future: future}
}

func (a *stub) DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DAXDescribeClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeClusters", input)
	return &DAXDescribeClustersFuture{Future: future}
}

func (a *stub) DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeDefaultParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DAXDescribeDefaultParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeDefaultParameters", input)
	return &DAXDescribeDefaultParametersFuture{Future: future}
}

func (a *stub) DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DAXDescribeEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeEvents", input)
	return &DAXDescribeEventsFuture{Future: future}
}

func (a *stub) DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameterGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DAXDescribeParameterGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameterGroups", input)
	return &DAXDescribeParameterGroupsFuture{Future: future}
}

func (a *stub) DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DAXDescribeParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameters", input)
	return &DAXDescribeParametersFuture{Future: future}
}

func (a *stub) DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeSubnetGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DAXDescribeSubnetGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeSubnetGroups", input)
	return &DAXDescribeSubnetGroupsFuture{Future: future}
}

func (a *stub) IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.IncreaseReplicationFactor", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *DAXIncreaseReplicationFactorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.IncreaseReplicationFactor", input)
	return &DAXIncreaseReplicationFactorFuture{Future: future}
}

func (a *stub) ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *DAXListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.ListTags", input)
	return &DAXListTagsFuture{Future: future}
}

func (a *stub) RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.RebootNode", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *DAXRebootNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.RebootNode", input)
	return &DAXRebootNodeFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *DAXTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.TagResource", input)
	return &DAXTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *DAXUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UntagResource", input)
	return &DAXUntagResourceFuture{Future: future}
}

func (a *stub) UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *DAXUpdateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateCluster", input)
	return &DAXUpdateClusterFuture{Future: future}
}

func (a *stub) UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *DAXUpdateParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateParameterGroup", input)
	return &DAXUpdateParameterGroupFuture{Future: future}
}

func (a *stub) UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *DAXUpdateSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateSubnetGroup", input)
	return &DAXUpdateSubnetGroupFuture{Future: future}
}
