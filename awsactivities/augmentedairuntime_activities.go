
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime/augmentedairuntimeiface"
)

type AugmentedAIRuntimeActivities struct {
    client augmentedairuntimeiface.AugmentedAIRuntimeAPI
}

func NewAugmentedAIRuntimeActivities(session *session.Session, config ...*aws.Config) *AugmentedAIRuntimeActivities {
    client := augmentedairuntime.New(session, config...)
    return &AugmentedAIRuntimeActivities{client: client}
}

func (a *AugmentedAIRuntimeActivities) DeleteHumanLoop(input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
    return a.client.DeleteHumanLoop(input)
}

func (a *AugmentedAIRuntimeActivities) DescribeHumanLoop(input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
    return a.client.DescribeHumanLoop(input)
}

func (a *AugmentedAIRuntimeActivities) ListHumanLoops(input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
    return a.client.ListHumanLoops(input)
}

func (a *AugmentedAIRuntimeActivities) StartHumanLoop(input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error) {
    return a.client.StartHumanLoop(input)
}

func (a *AugmentedAIRuntimeActivities) StopHumanLoop(input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error) {
    return a.client.StopHumanLoop(input)
}
