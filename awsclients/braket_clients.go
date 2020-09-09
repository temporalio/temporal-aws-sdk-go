package awsclients

import (
	"github.com/aws/aws-sdk-go/service/braket"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type BraketClient interface {
	CancelQuantumTask(ctx workflow.Context, input *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error)
	CancelQuantumTaskAsync(ctx workflow.Context, input *braket.CancelQuantumTaskInput) *BraketCancelQuantumTaskResult

	CreateQuantumTask(ctx workflow.Context, input *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error)
	CreateQuantumTaskAsync(ctx workflow.Context, input *braket.CreateQuantumTaskInput) *BraketCreateQuantumTaskResult

	GetDevice(ctx workflow.Context, input *braket.GetDeviceInput) (*braket.GetDeviceOutput, error)
	GetDeviceAsync(ctx workflow.Context, input *braket.GetDeviceInput) *BraketGetDeviceResult

	GetQuantumTask(ctx workflow.Context, input *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error)
	GetQuantumTaskAsync(ctx workflow.Context, input *braket.GetQuantumTaskInput) *BraketGetQuantumTaskResult

	SearchDevices(ctx workflow.Context, input *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error)
	SearchDevicesAsync(ctx workflow.Context, input *braket.SearchDevicesInput) *BraketSearchDevicesResult

	SearchQuantumTasks(ctx workflow.Context, input *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error)
	SearchQuantumTasksAsync(ctx workflow.Context, input *braket.SearchQuantumTasksInput) *BraketSearchQuantumTasksResult
}

type BraketCancelQuantumTaskResult struct {
	Result workflow.Future
}

func (r *BraketCancelQuantumTaskResult) Get(ctx workflow.Context) (*braket.CancelQuantumTaskOutput, error) {
	var output braket.CancelQuantumTaskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type BraketCreateQuantumTaskResult struct {
	Result workflow.Future
}

func (r *BraketCreateQuantumTaskResult) Get(ctx workflow.Context) (*braket.CreateQuantumTaskOutput, error) {
	var output braket.CreateQuantumTaskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type BraketGetDeviceResult struct {
	Result workflow.Future
}

func (r *BraketGetDeviceResult) Get(ctx workflow.Context) (*braket.GetDeviceOutput, error) {
	var output braket.GetDeviceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type BraketGetQuantumTaskResult struct {
	Result workflow.Future
}

func (r *BraketGetQuantumTaskResult) Get(ctx workflow.Context) (*braket.GetQuantumTaskOutput, error) {
	var output braket.GetQuantumTaskOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type BraketSearchDevicesResult struct {
	Result workflow.Future
}

func (r *BraketSearchDevicesResult) Get(ctx workflow.Context) (*braket.SearchDevicesOutput, error) {
	var output braket.SearchDevicesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type BraketSearchQuantumTasksResult struct {
	Result workflow.Future
}

func (r *BraketSearchQuantumTasksResult) Get(ctx workflow.Context) (*braket.SearchQuantumTasksOutput, error) {
	var output braket.SearchQuantumTasksOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type BraketStub struct {
	activities awsactivities.BraketActivities
}

func NewBraketStub() BraketClient {
	return &BraketStub{}
}

func (a *BraketStub) CancelQuantumTask(ctx workflow.Context, input *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error) {
	var output braket.CancelQuantumTaskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CancelQuantumTask, input).Get(ctx, &output)
	return &output, err
}

func (a *BraketStub) CancelQuantumTaskAsync(ctx workflow.Context, input *braket.CancelQuantumTaskInput) *BraketCancelQuantumTaskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CancelQuantumTask, input)
	return &BraketCancelQuantumTaskResult{Result: future}
}

func (a *BraketStub) CreateQuantumTask(ctx workflow.Context, input *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error) {
	var output braket.CreateQuantumTaskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateQuantumTask, input).Get(ctx, &output)
	return &output, err
}

func (a *BraketStub) CreateQuantumTaskAsync(ctx workflow.Context, input *braket.CreateQuantumTaskInput) *BraketCreateQuantumTaskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateQuantumTask, input)
	return &BraketCreateQuantumTaskResult{Result: future}
}

func (a *BraketStub) GetDevice(ctx workflow.Context, input *braket.GetDeviceInput) (*braket.GetDeviceOutput, error) {
	var output braket.GetDeviceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input).Get(ctx, &output)
	return &output, err
}

func (a *BraketStub) GetDeviceAsync(ctx workflow.Context, input *braket.GetDeviceInput) *BraketGetDeviceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetDevice, input)
	return &BraketGetDeviceResult{Result: future}
}

func (a *BraketStub) GetQuantumTask(ctx workflow.Context, input *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error) {
	var output braket.GetQuantumTaskOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetQuantumTask, input).Get(ctx, &output)
	return &output, err
}

func (a *BraketStub) GetQuantumTaskAsync(ctx workflow.Context, input *braket.GetQuantumTaskInput) *BraketGetQuantumTaskResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetQuantumTask, input)
	return &BraketGetQuantumTaskResult{Result: future}
}

func (a *BraketStub) SearchDevices(ctx workflow.Context, input *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error) {
	var output braket.SearchDevicesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchDevices, input).Get(ctx, &output)
	return &output, err
}

func (a *BraketStub) SearchDevicesAsync(ctx workflow.Context, input *braket.SearchDevicesInput) *BraketSearchDevicesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchDevices, input)
	return &BraketSearchDevicesResult{Result: future}
}

func (a *BraketStub) SearchQuantumTasks(ctx workflow.Context, input *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error) {
	var output braket.SearchQuantumTasksOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SearchQuantumTasks, input).Get(ctx, &output)
	return &output, err
}

func (a *BraketStub) SearchQuantumTasksAsync(ctx workflow.Context, input *braket.SearchQuantumTasksInput) *BraketSearchQuantumTasksResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SearchQuantumTasks, input)
	return &BraketSearchQuantumTasksResult{Result: future}
}
