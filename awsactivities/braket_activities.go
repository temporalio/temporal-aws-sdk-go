
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/braket"
	"github.com/aws/aws-sdk-go/service/braket/braketiface"
)

type BraketActivities struct {
	client braketiface.BraketAPI
}

func NewBraketActivities(session *session.Session, config... *aws.Config) *BraketActivities {
    client := braket.New(session, config...)
    return &BraketActivities{client: client}
}

func (a *BraketActivities) CancelQuantumTask(input *braket.CancelQuantumTaskInput) (*braket.CancelQuantumTaskOutput, error) {
    return a.client.CancelQuantumTask(input)
}

func (a *BraketActivities) CreateQuantumTask(input *braket.CreateQuantumTaskInput) (*braket.CreateQuantumTaskOutput, error) {
    return a.client.CreateQuantumTask(input)
}

func (a *BraketActivities) GetDevice(input *braket.GetDeviceInput) (*braket.GetDeviceOutput, error) {
    return a.client.GetDevice(input)
}

func (a *BraketActivities) GetQuantumTask(input *braket.GetQuantumTaskInput) (*braket.GetQuantumTaskOutput, error) {
    return a.client.GetQuantumTask(input)
}

func (a *BraketActivities) SearchDevices(input *braket.SearchDevicesInput) (*braket.SearchDevicesOutput, error) {
    return a.client.SearchDevices(input)
}

func (a *BraketActivities) SearchQuantumTasks(input *braket.SearchQuantumTasksInput) (*braket.SearchQuantumTasksOutput, error) {
    return a.client.SearchQuantumTasks(input)
}
