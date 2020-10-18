// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client sfniface.SFNAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := sfn.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (sfniface.SFNAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return sfn.New(sess), nil
}

func (a *Activities) CreateActivity(ctx context.Context, input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateActivityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStateMachine(ctx context.Context, input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStateMachineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteActivity(ctx context.Context, input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteActivityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteStateMachine(ctx context.Context, input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteStateMachineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeActivity(ctx context.Context, input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeActivityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeExecution(ctx context.Context, input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeExecutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStateMachine(ctx context.Context, input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStateMachineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStateMachineForExecution(ctx context.Context, input *sfn.DescribeStateMachineForExecutionInput) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStateMachineForExecutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetActivityTask(ctx context.Context, input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetActivityTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetExecutionHistory(ctx context.Context, input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetExecutionHistoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListActivities(ctx context.Context, input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListActivitiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListExecutions(ctx context.Context, input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListExecutionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStateMachines(ctx context.Context, input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStateMachinesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *sfn.ListTagsForResourceInput) (*sfn.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendTaskFailure(ctx context.Context, input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendTaskFailureWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendTaskHeartbeat(ctx context.Context, input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendTaskHeartbeatWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SendTaskSuccess(ctx context.Context, input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SendTaskSuccessWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartExecution(ctx context.Context, input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartExecutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopExecution(ctx context.Context, input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopExecutionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *sfn.TagResourceInput) (*sfn.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *sfn.UntagResourceInput) (*sfn.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateStateMachine(ctx context.Context, input *sfn.UpdateStateMachineInput) (*sfn.UpdateStateMachineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateStateMachineWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
