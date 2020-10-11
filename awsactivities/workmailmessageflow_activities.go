// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type WorkMailMessageFlowActivities struct {
	client workmailmessageflowiface.WorkMailMessageFlowAPI

	sessionFactory SessionFactory
}

func NewWorkMailMessageFlowActivities(sess *session.Session, config ...*aws.Config) *WorkMailMessageFlowActivities {
	client := workmailmessageflow.New(sess, config...)
	return &WorkMailMessageFlowActivities{client: client}
}

func NewWorkMailMessageFlowActivitiesWithSessionFactory(sessionFactory SessionFactory) *WorkMailMessageFlowActivities {
	return &WorkMailMessageFlowActivities{sessionFactory: sessionFactory}
}

func (a *WorkMailMessageFlowActivities) getClient(ctx context.Context) (workmailmessageflowiface.WorkMailMessageFlowAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return workmailmessageflow.New(sess), nil
}

func (a *WorkMailMessageFlowActivities) GetRawMessageContent(ctx context.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRawMessageContentWithContext(ctx, input)
}
