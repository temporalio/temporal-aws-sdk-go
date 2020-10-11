// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmediaiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type KinesisVideoArchivedMediaActivities struct {
	client kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI

	sessionFactory SessionFactory
}

func NewKinesisVideoArchivedMediaActivities(sess *session.Session, config ...*aws.Config) *KinesisVideoArchivedMediaActivities {
	client := kinesisvideoarchivedmedia.New(sess, config...)
	return &KinesisVideoArchivedMediaActivities{client: client}
}

func NewKinesisVideoArchivedMediaActivitiesWithSessionFactory(sessionFactory SessionFactory) *KinesisVideoArchivedMediaActivities {
	return &KinesisVideoArchivedMediaActivities{sessionFactory: sessionFactory}
}

func (a *KinesisVideoArchivedMediaActivities) getClient(ctx context.Context) (kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return kinesisvideoarchivedmedia.New(sess), nil
}

func (a *KinesisVideoArchivedMediaActivities) GetClip(ctx context.Context, input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetClipWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) GetDASHStreamingSessionURL(ctx context.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDASHStreamingSessionURLWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) GetHLSStreamingSessionURL(ctx context.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHLSStreamingSessionURLWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) GetMediaForFragmentList(ctx context.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMediaForFragmentListWithContext(ctx, input)
}

func (a *KinesisVideoArchivedMediaActivities) ListFragments(ctx context.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFragmentsWithContext(ctx, input)
}
