// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/datasync"
	"go.temporal.io/sdk/workflow"
)

type DataSyncClient interface {
	CancelTaskExecution(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error)
	CancelTaskExecutionAsync(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) *DatasyncCancelTaskExecutionFuture

	CreateAgent(ctx workflow.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error)
	CreateAgentAsync(ctx workflow.Context, input *datasync.CreateAgentInput) *DatasyncCreateAgentFuture

	CreateLocationEfs(ctx workflow.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error)
	CreateLocationEfsAsync(ctx workflow.Context, input *datasync.CreateLocationEfsInput) *DatasyncCreateLocationEfsFuture

	CreateLocationFsxWindows(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error)
	CreateLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) *DatasyncCreateLocationFsxWindowsFuture

	CreateLocationNfs(ctx workflow.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error)
	CreateLocationNfsAsync(ctx workflow.Context, input *datasync.CreateLocationNfsInput) *DatasyncCreateLocationNfsFuture

	CreateLocationObjectStorage(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error)
	CreateLocationObjectStorageAsync(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) *DatasyncCreateLocationObjectStorageFuture

	CreateLocationS3(ctx workflow.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error)
	CreateLocationS3Async(ctx workflow.Context, input *datasync.CreateLocationS3Input) *DatasyncCreateLocationS3Future

	CreateLocationSmb(ctx workflow.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error)
	CreateLocationSmbAsync(ctx workflow.Context, input *datasync.CreateLocationSmbInput) *DatasyncCreateLocationSmbFuture

	CreateTask(ctx workflow.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error)
	CreateTaskAsync(ctx workflow.Context, input *datasync.CreateTaskInput) *DatasyncCreateTaskFuture

	DeleteAgent(ctx workflow.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error)
	DeleteAgentAsync(ctx workflow.Context, input *datasync.DeleteAgentInput) *DatasyncDeleteAgentFuture

	DeleteLocation(ctx workflow.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error)
	DeleteLocationAsync(ctx workflow.Context, input *datasync.DeleteLocationInput) *DatasyncDeleteLocationFuture

	DeleteTask(ctx workflow.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error)
	DeleteTaskAsync(ctx workflow.Context, input *datasync.DeleteTaskInput) *DatasyncDeleteTaskFuture

	DescribeAgent(ctx workflow.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error)
	DescribeAgentAsync(ctx workflow.Context, input *datasync.DescribeAgentInput) *DatasyncDescribeAgentFuture

	DescribeLocationEfs(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error)
	DescribeLocationEfsAsync(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) *DatasyncDescribeLocationEfsFuture

	DescribeLocationFsxWindows(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error)
	DescribeLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) *DatasyncDescribeLocationFsxWindowsFuture

	DescribeLocationNfs(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error)
	DescribeLocationNfsAsync(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) *DatasyncDescribeLocationNfsFuture

	DescribeLocationObjectStorage(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error)
	DescribeLocationObjectStorageAsync(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) *DatasyncDescribeLocationObjectStorageFuture

	DescribeLocationS3(ctx workflow.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error)
	DescribeLocationS3Async(ctx workflow.Context, input *datasync.DescribeLocationS3Input) *DatasyncDescribeLocationS3Future

	DescribeLocationSmb(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error)
	DescribeLocationSmbAsync(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) *DatasyncDescribeLocationSmbFuture

	DescribeTask(ctx workflow.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error)
	DescribeTaskAsync(ctx workflow.Context, input *datasync.DescribeTaskInput) *DatasyncDescribeTaskFuture

	DescribeTaskExecution(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error)
	DescribeTaskExecutionAsync(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) *DatasyncDescribeTaskExecutionFuture

	ListAgents(ctx workflow.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error)
	ListAgentsAsync(ctx workflow.Context, input *datasync.ListAgentsInput) *DatasyncListAgentsFuture

	ListLocations(ctx workflow.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error)
	ListLocationsAsync(ctx workflow.Context, input *datasync.ListLocationsInput) *DatasyncListLocationsFuture

	ListTagsForResource(ctx workflow.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *datasync.ListTagsForResourceInput) *DatasyncListTagsForResourceFuture

	ListTaskExecutions(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error)
	ListTaskExecutionsAsync(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) *DatasyncListTaskExecutionsFuture

	ListTasks(ctx workflow.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error)
	ListTasksAsync(ctx workflow.Context, input *datasync.ListTasksInput) *DatasyncListTasksFuture

	StartTaskExecution(ctx workflow.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error)
	StartTaskExecutionAsync(ctx workflow.Context, input *datasync.StartTaskExecutionInput) *DatasyncStartTaskExecutionFuture

	TagResource(ctx workflow.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *datasync.TagResourceInput) *DatasyncTagResourceFuture

	UntagResource(ctx workflow.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *datasync.UntagResourceInput) *DatasyncUntagResourceFuture

	UpdateAgent(ctx workflow.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error)
	UpdateAgentAsync(ctx workflow.Context, input *datasync.UpdateAgentInput) *DatasyncUpdateAgentFuture

	UpdateTask(ctx workflow.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error)
	UpdateTaskAsync(ctx workflow.Context, input *datasync.UpdateTaskInput) *DatasyncUpdateTaskFuture
}

type DataSyncStub struct{}

func NewDataSyncStub() DataSyncClient {
	return &DataSyncStub{}
}

type DatasyncCancelTaskExecutionFuture struct {
	Future workflow.Future
}

func (r *DatasyncCancelTaskExecutionFuture) Get(ctx workflow.Context) (*datasync.CancelTaskExecutionOutput, error) {
	var output datasync.CancelTaskExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateAgentFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateAgentFuture) Get(ctx workflow.Context) (*datasync.CreateAgentOutput, error) {
	var output datasync.CreateAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateLocationEfsFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateLocationEfsFuture) Get(ctx workflow.Context) (*datasync.CreateLocationEfsOutput, error) {
	var output datasync.CreateLocationEfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateLocationFsxWindowsFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateLocationFsxWindowsFuture) Get(ctx workflow.Context) (*datasync.CreateLocationFsxWindowsOutput, error) {
	var output datasync.CreateLocationFsxWindowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateLocationNfsFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateLocationNfsFuture) Get(ctx workflow.Context) (*datasync.CreateLocationNfsOutput, error) {
	var output datasync.CreateLocationNfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateLocationObjectStorageFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateLocationObjectStorageFuture) Get(ctx workflow.Context) (*datasync.CreateLocationObjectStorageOutput, error) {
	var output datasync.CreateLocationObjectStorageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateLocationS3Future struct {
	Future workflow.Future
}

func (r *DatasyncCreateLocationS3Future) Get(ctx workflow.Context) (*datasync.CreateLocationS3Output, error) {
	var output datasync.CreateLocationS3Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateLocationSmbFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateLocationSmbFuture) Get(ctx workflow.Context) (*datasync.CreateLocationSmbOutput, error) {
	var output datasync.CreateLocationSmbOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncCreateTaskFuture struct {
	Future workflow.Future
}

func (r *DatasyncCreateTaskFuture) Get(ctx workflow.Context) (*datasync.CreateTaskOutput, error) {
	var output datasync.CreateTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDeleteAgentFuture struct {
	Future workflow.Future
}

func (r *DatasyncDeleteAgentFuture) Get(ctx workflow.Context) (*datasync.DeleteAgentOutput, error) {
	var output datasync.DeleteAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDeleteLocationFuture struct {
	Future workflow.Future
}

func (r *DatasyncDeleteLocationFuture) Get(ctx workflow.Context) (*datasync.DeleteLocationOutput, error) {
	var output datasync.DeleteLocationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDeleteTaskFuture struct {
	Future workflow.Future
}

func (r *DatasyncDeleteTaskFuture) Get(ctx workflow.Context) (*datasync.DeleteTaskOutput, error) {
	var output datasync.DeleteTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeAgentFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeAgentFuture) Get(ctx workflow.Context) (*datasync.DescribeAgentOutput, error) {
	var output datasync.DescribeAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeLocationEfsFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeLocationEfsFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationEfsOutput, error) {
	var output datasync.DescribeLocationEfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeLocationFsxWindowsFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeLocationFsxWindowsFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	var output datasync.DescribeLocationFsxWindowsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeLocationNfsFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeLocationNfsFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationNfsOutput, error) {
	var output datasync.DescribeLocationNfsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeLocationObjectStorageFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeLocationObjectStorageFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationObjectStorageOutput, error) {
	var output datasync.DescribeLocationObjectStorageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeLocationS3Future struct {
	Future workflow.Future
}

func (r *DatasyncDescribeLocationS3Future) Get(ctx workflow.Context) (*datasync.DescribeLocationS3Output, error) {
	var output datasync.DescribeLocationS3Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeLocationSmbFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeLocationSmbFuture) Get(ctx workflow.Context) (*datasync.DescribeLocationSmbOutput, error) {
	var output datasync.DescribeLocationSmbOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeTaskFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeTaskFuture) Get(ctx workflow.Context) (*datasync.DescribeTaskOutput, error) {
	var output datasync.DescribeTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncDescribeTaskExecutionFuture struct {
	Future workflow.Future
}

func (r *DatasyncDescribeTaskExecutionFuture) Get(ctx workflow.Context) (*datasync.DescribeTaskExecutionOutput, error) {
	var output datasync.DescribeTaskExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncListAgentsFuture struct {
	Future workflow.Future
}

func (r *DatasyncListAgentsFuture) Get(ctx workflow.Context) (*datasync.ListAgentsOutput, error) {
	var output datasync.ListAgentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncListLocationsFuture struct {
	Future workflow.Future
}

func (r *DatasyncListLocationsFuture) Get(ctx workflow.Context) (*datasync.ListLocationsOutput, error) {
	var output datasync.ListLocationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *DatasyncListTagsForResourceFuture) Get(ctx workflow.Context) (*datasync.ListTagsForResourceOutput, error) {
	var output datasync.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncListTaskExecutionsFuture struct {
	Future workflow.Future
}

func (r *DatasyncListTaskExecutionsFuture) Get(ctx workflow.Context) (*datasync.ListTaskExecutionsOutput, error) {
	var output datasync.ListTaskExecutionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncListTasksFuture struct {
	Future workflow.Future
}

func (r *DatasyncListTasksFuture) Get(ctx workflow.Context) (*datasync.ListTasksOutput, error) {
	var output datasync.ListTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncStartTaskExecutionFuture struct {
	Future workflow.Future
}

func (r *DatasyncStartTaskExecutionFuture) Get(ctx workflow.Context) (*datasync.StartTaskExecutionOutput, error) {
	var output datasync.StartTaskExecutionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncTagResourceFuture struct {
	Future workflow.Future
}

func (r *DatasyncTagResourceFuture) Get(ctx workflow.Context) (*datasync.TagResourceOutput, error) {
	var output datasync.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncUntagResourceFuture struct {
	Future workflow.Future
}

func (r *DatasyncUntagResourceFuture) Get(ctx workflow.Context) (*datasync.UntagResourceOutput, error) {
	var output datasync.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncUpdateAgentFuture struct {
	Future workflow.Future
}

func (r *DatasyncUpdateAgentFuture) Get(ctx workflow.Context) (*datasync.UpdateAgentOutput, error) {
	var output datasync.UpdateAgentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DatasyncUpdateTaskFuture struct {
	Future workflow.Future
}

func (r *DatasyncUpdateTaskFuture) Get(ctx workflow.Context) (*datasync.UpdateTaskOutput, error) {
	var output datasync.UpdateTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CancelTaskExecution(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) (*datasync.CancelTaskExecutionOutput, error) {
	var output datasync.CancelTaskExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CancelTaskExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CancelTaskExecutionAsync(ctx workflow.Context, input *datasync.CancelTaskExecutionInput) *DatasyncCancelTaskExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CancelTaskExecution", input)
	return &DatasyncCancelTaskExecutionFuture{Future: future}
}

func (a *DataSyncStub) CreateAgent(ctx workflow.Context, input *datasync.CreateAgentInput) (*datasync.CreateAgentOutput, error) {
	var output datasync.CreateAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateAgentAsync(ctx workflow.Context, input *datasync.CreateAgentInput) *DatasyncCreateAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateAgent", input)
	return &DatasyncCreateAgentFuture{Future: future}
}

func (a *DataSyncStub) CreateLocationEfs(ctx workflow.Context, input *datasync.CreateLocationEfsInput) (*datasync.CreateLocationEfsOutput, error) {
	var output datasync.CreateLocationEfsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationEfs", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateLocationEfsAsync(ctx workflow.Context, input *datasync.CreateLocationEfsInput) *DatasyncCreateLocationEfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationEfs", input)
	return &DatasyncCreateLocationEfsFuture{Future: future}
}

func (a *DataSyncStub) CreateLocationFsxWindows(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) (*datasync.CreateLocationFsxWindowsOutput, error) {
	var output datasync.CreateLocationFsxWindowsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationFsxWindows", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.CreateLocationFsxWindowsInput) *DatasyncCreateLocationFsxWindowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationFsxWindows", input)
	return &DatasyncCreateLocationFsxWindowsFuture{Future: future}
}

func (a *DataSyncStub) CreateLocationNfs(ctx workflow.Context, input *datasync.CreateLocationNfsInput) (*datasync.CreateLocationNfsOutput, error) {
	var output datasync.CreateLocationNfsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationNfs", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateLocationNfsAsync(ctx workflow.Context, input *datasync.CreateLocationNfsInput) *DatasyncCreateLocationNfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationNfs", input)
	return &DatasyncCreateLocationNfsFuture{Future: future}
}

func (a *DataSyncStub) CreateLocationObjectStorage(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) (*datasync.CreateLocationObjectStorageOutput, error) {
	var output datasync.CreateLocationObjectStorageOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationObjectStorage", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateLocationObjectStorageAsync(ctx workflow.Context, input *datasync.CreateLocationObjectStorageInput) *DatasyncCreateLocationObjectStorageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationObjectStorage", input)
	return &DatasyncCreateLocationObjectStorageFuture{Future: future}
}

func (a *DataSyncStub) CreateLocationS3(ctx workflow.Context, input *datasync.CreateLocationS3Input) (*datasync.CreateLocationS3Output, error) {
	var output datasync.CreateLocationS3Output
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationS3", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateLocationS3Async(ctx workflow.Context, input *datasync.CreateLocationS3Input) *DatasyncCreateLocationS3Future {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationS3", input)
	return &DatasyncCreateLocationS3Future{Future: future}
}

func (a *DataSyncStub) CreateLocationSmb(ctx workflow.Context, input *datasync.CreateLocationSmbInput) (*datasync.CreateLocationSmbOutput, error) {
	var output datasync.CreateLocationSmbOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationSmb", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateLocationSmbAsync(ctx workflow.Context, input *datasync.CreateLocationSmbInput) *DatasyncCreateLocationSmbFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateLocationSmb", input)
	return &DatasyncCreateLocationSmbFuture{Future: future}
}

func (a *DataSyncStub) CreateTask(ctx workflow.Context, input *datasync.CreateTaskInput) (*datasync.CreateTaskOutput, error) {
	var output datasync.CreateTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.CreateTask", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) CreateTaskAsync(ctx workflow.Context, input *datasync.CreateTaskInput) *DatasyncCreateTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.CreateTask", input)
	return &DatasyncCreateTaskFuture{Future: future}
}

func (a *DataSyncStub) DeleteAgent(ctx workflow.Context, input *datasync.DeleteAgentInput) (*datasync.DeleteAgentOutput, error) {
	var output datasync.DeleteAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DeleteAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DeleteAgentAsync(ctx workflow.Context, input *datasync.DeleteAgentInput) *DatasyncDeleteAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DeleteAgent", input)
	return &DatasyncDeleteAgentFuture{Future: future}
}

func (a *DataSyncStub) DeleteLocation(ctx workflow.Context, input *datasync.DeleteLocationInput) (*datasync.DeleteLocationOutput, error) {
	var output datasync.DeleteLocationOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DeleteLocation", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DeleteLocationAsync(ctx workflow.Context, input *datasync.DeleteLocationInput) *DatasyncDeleteLocationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DeleteLocation", input)
	return &DatasyncDeleteLocationFuture{Future: future}
}

func (a *DataSyncStub) DeleteTask(ctx workflow.Context, input *datasync.DeleteTaskInput) (*datasync.DeleteTaskOutput, error) {
	var output datasync.DeleteTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DeleteTask", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DeleteTaskAsync(ctx workflow.Context, input *datasync.DeleteTaskInput) *DatasyncDeleteTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DeleteTask", input)
	return &DatasyncDeleteTaskFuture{Future: future}
}

func (a *DataSyncStub) DescribeAgent(ctx workflow.Context, input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error) {
	var output datasync.DescribeAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeAgentAsync(ctx workflow.Context, input *datasync.DescribeAgentInput) *DatasyncDescribeAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeAgent", input)
	return &DatasyncDescribeAgentFuture{Future: future}
}

func (a *DataSyncStub) DescribeLocationEfs(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error) {
	var output datasync.DescribeLocationEfsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationEfs", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeLocationEfsAsync(ctx workflow.Context, input *datasync.DescribeLocationEfsInput) *DatasyncDescribeLocationEfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationEfs", input)
	return &DatasyncDescribeLocationEfsFuture{Future: future}
}

func (a *DataSyncStub) DescribeLocationFsxWindows(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	var output datasync.DescribeLocationFsxWindowsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationFsxWindows", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeLocationFsxWindowsAsync(ctx workflow.Context, input *datasync.DescribeLocationFsxWindowsInput) *DatasyncDescribeLocationFsxWindowsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationFsxWindows", input)
	return &DatasyncDescribeLocationFsxWindowsFuture{Future: future}
}

func (a *DataSyncStub) DescribeLocationNfs(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error) {
	var output datasync.DescribeLocationNfsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationNfs", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeLocationNfsAsync(ctx workflow.Context, input *datasync.DescribeLocationNfsInput) *DatasyncDescribeLocationNfsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationNfs", input)
	return &DatasyncDescribeLocationNfsFuture{Future: future}
}

func (a *DataSyncStub) DescribeLocationObjectStorage(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error) {
	var output datasync.DescribeLocationObjectStorageOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationObjectStorage", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeLocationObjectStorageAsync(ctx workflow.Context, input *datasync.DescribeLocationObjectStorageInput) *DatasyncDescribeLocationObjectStorageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationObjectStorage", input)
	return &DatasyncDescribeLocationObjectStorageFuture{Future: future}
}

func (a *DataSyncStub) DescribeLocationS3(ctx workflow.Context, input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error) {
	var output datasync.DescribeLocationS3Output
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationS3", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeLocationS3Async(ctx workflow.Context, input *datasync.DescribeLocationS3Input) *DatasyncDescribeLocationS3Future {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationS3", input)
	return &DatasyncDescribeLocationS3Future{Future: future}
}

func (a *DataSyncStub) DescribeLocationSmb(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error) {
	var output datasync.DescribeLocationSmbOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationSmb", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeLocationSmbAsync(ctx workflow.Context, input *datasync.DescribeLocationSmbInput) *DatasyncDescribeLocationSmbFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeLocationSmb", input)
	return &DatasyncDescribeLocationSmbFuture{Future: future}
}

func (a *DataSyncStub) DescribeTask(ctx workflow.Context, input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error) {
	var output datasync.DescribeTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeTask", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeTaskAsync(ctx workflow.Context, input *datasync.DescribeTaskInput) *DatasyncDescribeTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeTask", input)
	return &DatasyncDescribeTaskFuture{Future: future}
}

func (a *DataSyncStub) DescribeTaskExecution(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error) {
	var output datasync.DescribeTaskExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeTaskExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) DescribeTaskExecutionAsync(ctx workflow.Context, input *datasync.DescribeTaskExecutionInput) *DatasyncDescribeTaskExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.DescribeTaskExecution", input)
	return &DatasyncDescribeTaskExecutionFuture{Future: future}
}

func (a *DataSyncStub) ListAgents(ctx workflow.Context, input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error) {
	var output datasync.ListAgentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.ListAgents", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) ListAgentsAsync(ctx workflow.Context, input *datasync.ListAgentsInput) *DatasyncListAgentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.ListAgents", input)
	return &DatasyncListAgentsFuture{Future: future}
}

func (a *DataSyncStub) ListLocations(ctx workflow.Context, input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error) {
	var output datasync.ListLocationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.ListLocations", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) ListLocationsAsync(ctx workflow.Context, input *datasync.ListLocationsInput) *DatasyncListLocationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.ListLocations", input)
	return &DatasyncListLocationsFuture{Future: future}
}

func (a *DataSyncStub) ListTagsForResource(ctx workflow.Context, input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error) {
	var output datasync.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) ListTagsForResourceAsync(ctx workflow.Context, input *datasync.ListTagsForResourceInput) *DatasyncListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.ListTagsForResource", input)
	return &DatasyncListTagsForResourceFuture{Future: future}
}

func (a *DataSyncStub) ListTaskExecutions(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error) {
	var output datasync.ListTaskExecutionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.ListTaskExecutions", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) ListTaskExecutionsAsync(ctx workflow.Context, input *datasync.ListTaskExecutionsInput) *DatasyncListTaskExecutionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.ListTaskExecutions", input)
	return &DatasyncListTaskExecutionsFuture{Future: future}
}

func (a *DataSyncStub) ListTasks(ctx workflow.Context, input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error) {
	var output datasync.ListTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.ListTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) ListTasksAsync(ctx workflow.Context, input *datasync.ListTasksInput) *DatasyncListTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.ListTasks", input)
	return &DatasyncListTasksFuture{Future: future}
}

func (a *DataSyncStub) StartTaskExecution(ctx workflow.Context, input *datasync.StartTaskExecutionInput) (*datasync.StartTaskExecutionOutput, error) {
	var output datasync.StartTaskExecutionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.StartTaskExecution", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) StartTaskExecutionAsync(ctx workflow.Context, input *datasync.StartTaskExecutionInput) *DatasyncStartTaskExecutionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.StartTaskExecution", input)
	return &DatasyncStartTaskExecutionFuture{Future: future}
}

func (a *DataSyncStub) TagResource(ctx workflow.Context, input *datasync.TagResourceInput) (*datasync.TagResourceOutput, error) {
	var output datasync.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) TagResourceAsync(ctx workflow.Context, input *datasync.TagResourceInput) *DatasyncTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.TagResource", input)
	return &DatasyncTagResourceFuture{Future: future}
}

func (a *DataSyncStub) UntagResource(ctx workflow.Context, input *datasync.UntagResourceInput) (*datasync.UntagResourceOutput, error) {
	var output datasync.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) UntagResourceAsync(ctx workflow.Context, input *datasync.UntagResourceInput) *DatasyncUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.UntagResource", input)
	return &DatasyncUntagResourceFuture{Future: future}
}

func (a *DataSyncStub) UpdateAgent(ctx workflow.Context, input *datasync.UpdateAgentInput) (*datasync.UpdateAgentOutput, error) {
	var output datasync.UpdateAgentOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.UpdateAgent", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) UpdateAgentAsync(ctx workflow.Context, input *datasync.UpdateAgentInput) *DatasyncUpdateAgentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.UpdateAgent", input)
	return &DatasyncUpdateAgentFuture{Future: future}
}

func (a *DataSyncStub) UpdateTask(ctx workflow.Context, input *datasync.UpdateTaskInput) (*datasync.UpdateTaskOutput, error) {
	var output datasync.UpdateTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.datasync.UpdateTask", input).Get(ctx, &output)
	return &output, err
}

func (a *DataSyncStub) UpdateTaskAsync(ctx workflow.Context, input *datasync.UpdateTaskInput) *DatasyncUpdateTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datasync.UpdateTask", input)
	return &DatasyncUpdateTaskFuture{Future: future}
}
