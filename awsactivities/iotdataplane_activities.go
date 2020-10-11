// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTDataPlaneActivities struct {
	client iotdataplaneiface.IoTDataPlaneAPI

	sessionFactory SessionFactory
}

func NewIoTDataPlaneActivities(sess *session.Session, config ...*aws.Config) *IoTDataPlaneActivities {
	client := iotdataplane.New(sess, config...)
	return &IoTDataPlaneActivities{client: client}
}

func NewIoTDataPlaneActivitiesWithSessionFactory(sessionFactory SessionFactory) *IoTDataPlaneActivities {
	return &IoTDataPlaneActivities{sessionFactory: sessionFactory}
}

func (a *IoTDataPlaneActivities) getClient(ctx context.Context) (iotdataplaneiface.IoTDataPlaneAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iotdataplane.New(sess), nil
}

func (a *IoTDataPlaneActivities) DeleteThingShadow(ctx context.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteThingShadowWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) GetThingShadow(ctx context.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetThingShadowWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) ListNamedShadowsForThing(ctx context.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNamedShadowsForThingWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) Publish(ctx context.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PublishWithContext(ctx, input)
}

func (a *IoTDataPlaneActivities) UpdateThingShadow(ctx context.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateThingShadowWithContext(ctx, input)
}
