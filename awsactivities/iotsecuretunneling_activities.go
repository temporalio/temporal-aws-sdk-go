// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling/iotsecuretunnelingiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTSecureTunnelingActivities struct {
	client iotsecuretunnelingiface.IoTSecureTunnelingAPI

	sessionFactory SessionFactory
}

func NewIoTSecureTunnelingActivities(sess *session.Session, config ...*aws.Config) *IoTSecureTunnelingActivities {
	client := iotsecuretunneling.New(sess, config...)
	return &IoTSecureTunnelingActivities{client: client}
}

func NewIoTSecureTunnelingActivitiesWithSessionFactory(sessionFactory SessionFactory) *IoTSecureTunnelingActivities {
	return &IoTSecureTunnelingActivities{sessionFactory: sessionFactory}
}

func (a *IoTSecureTunnelingActivities) getClient(ctx context.Context) (iotsecuretunnelingiface.IoTSecureTunnelingAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return iotsecuretunneling.New(sess), nil
}

func (a *IoTSecureTunnelingActivities) CloseTunnel(ctx context.Context, input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CloseTunnelWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) DescribeTunnel(ctx context.Context, input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTunnelWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) ListTagsForResource(ctx context.Context, input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) ListTunnels(ctx context.Context, input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTunnelsWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) OpenTunnel(ctx context.Context, input *iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.OpenTunnelWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) TagResource(ctx context.Context, input *iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *IoTSecureTunnelingActivities) UntagResource(ctx context.Context, input *iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
