// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package augmentedairuntime

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime/augmentedairuntimeiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client augmentedairuntimeiface.AugmentedAIRuntimeAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := augmentedairuntime.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (augmentedairuntimeiface.AugmentedAIRuntimeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return augmentedairuntime.New(sess), nil
}

func (a *Activities) DeleteHumanLoop(ctx context.Context, input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteHumanLoopWithContext(ctx, input)
}

func (a *Activities) DescribeHumanLoop(ctx context.Context, input *augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeHumanLoopWithContext(ctx, input)
}

func (a *Activities) ListHumanLoops(ctx context.Context, input *augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListHumanLoopsWithContext(ctx, input)
}

func (a *Activities) StartHumanLoop(ctx context.Context, input *augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartHumanLoopWithContext(ctx, input)
}

func (a *Activities) StopHumanLoop(ctx context.Context, input *augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopHumanLoopWithContext(ctx, input)
}
