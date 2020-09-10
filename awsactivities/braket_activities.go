package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/braket"
	"github.com/aws/aws-sdk-go/service/braket/braketiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type BraketActivities struct {
	client braketiface.BraketAPI
}

func NewBraketActivities(session *session.Session, config ...*aws.Config) *BraketActivities {
	client := braket.New(session, config...)
	return &BraketActivities{client: client}
}

func (a *BraketActivities) CancelQuantumTask(ctx context.Context, input *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CancelQuantumTaskWithContext(ctx, input)
}

func (a *BraketActivities) CreateQuantumTask(ctx context.Context, input *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateQuantumTaskWithContext(ctx, input)
}

func (a *BraketActivities) GetDevice(ctx context.Context, input *braket.GetDeviceInput) (*braket.GetDeviceOutput, error) {
	return a.client.GetDeviceWithContext(ctx, input)
}

func (a *BraketActivities) GetQuantumTask(ctx context.Context, input *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error) {
	return a.client.GetQuantumTaskWithContext(ctx, input)
}

func (a *BraketActivities) SearchDevices(ctx context.Context, input *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error) {
	return a.client.SearchDevicesWithContext(ctx, input)
}

func (a *BraketActivities) SearchQuantumTasks(ctx context.Context, input *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error) {
	return a.client.SearchQuantumTasksWithContext(ctx, input)
}
