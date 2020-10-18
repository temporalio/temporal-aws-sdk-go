// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package braketstub

import (
	"github.com/aws/aws-sdk-go/service/braket"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelQuantumTask(ctx workflow.Context, input *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error)
	CancelQuantumTaskAsync(ctx workflow.Context, input *braket.CancelQuantumTaskInput) *BraketCancelQuantumTaskFuture

	CreateQuantumTask(ctx workflow.Context, input *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error)
	CreateQuantumTaskAsync(ctx workflow.Context, input *braket.CreateQuantumTaskInput) *BraketCreateQuantumTaskFuture

	GetDevice(ctx workflow.Context, input *braket.GetDeviceInput) (*braket.GetDeviceOutput, error)
	GetDeviceAsync(ctx workflow.Context, input *braket.GetDeviceInput) *BraketGetDeviceFuture

	GetQuantumTask(ctx workflow.Context, input *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error)
	GetQuantumTaskAsync(ctx workflow.Context, input *braket.GetQuantumTaskInput) *BraketGetQuantumTaskFuture

	SearchDevices(ctx workflow.Context, input *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error)
	SearchDevicesAsync(ctx workflow.Context, input *braket.SearchDevicesInput) *BraketSearchDevicesFuture

	SearchQuantumTasks(ctx workflow.Context, input *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error)
	SearchQuantumTasksAsync(ctx workflow.Context, input *braket.SearchQuantumTasksInput) *BraketSearchQuantumTasksFuture
}

func NewClient() Client {
	return &stub{}
}
