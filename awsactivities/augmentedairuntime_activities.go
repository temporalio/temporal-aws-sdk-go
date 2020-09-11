package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime/augmentedairuntimeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AugmentedAIRuntimeActivities struct {
	client augmentedairuntimeiface.AugmentedAIRuntimeAPI
}

func NewAugmentedAIRuntimeActivities(session *session.Session, config ...*aws.Config) *AugmentedAIRuntimeActivities {
	client := augmentedairuntime.New(session, config...)
	return &AugmentedAIRuntimeActivities{client: client}
}

func (a *AugmentedAIRuntimeActivities) DeleteHumanLoop(ctx context.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
	return a.client.DeleteHumanLoopWithContext(ctx, input)
}

func (a *AugmentedAIRuntimeActivities) DescribeHumanLoop(ctx context.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	return a.client.DescribeHumanLoopWithContext(ctx, input)
}

func (a *AugmentedAIRuntimeActivities) ListHumanLoops(ctx context.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	return a.client.ListHumanLoopsWithContext(ctx, input)
}

func (a *AugmentedAIRuntimeActivities) StartHumanLoop(ctx context.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error) {
	return a.client.StartHumanLoopWithContext(ctx, input)
}

func (a *AugmentedAIRuntimeActivities) StopHumanLoop(ctx context.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error) {
	return a.client.StopHumanLoopWithContext(ctx, input)
}
