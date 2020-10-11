// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MediaTailorActivities struct {
	client mediatailoriface.MediaTailorAPI

	sessionFactory SessionFactory
}

func NewMediaTailorActivities(sess *session.Session, config ...*aws.Config) *MediaTailorActivities {
	client := mediatailor.New(sess, config...)
	return &MediaTailorActivities{client: client}
}

func NewMediaTailorActivitiesWithSessionFactory(sessionFactory SessionFactory) *MediaTailorActivities {
	return &MediaTailorActivities{sessionFactory: sessionFactory}
}

func (a *MediaTailorActivities) getClient(ctx context.Context) (mediatailoriface.MediaTailorAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return mediatailor.New(sess), nil
}

func (a *MediaTailorActivities) DeletePlaybackConfiguration(ctx context.Context, input *mediatailor.DeletePlaybackConfigurationInput) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePlaybackConfigurationWithContext(ctx, input)
}

func (a *MediaTailorActivities) GetPlaybackConfiguration(ctx context.Context, input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPlaybackConfigurationWithContext(ctx, input)
}

func (a *MediaTailorActivities) ListPlaybackConfigurations(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPlaybackConfigurationsWithContext(ctx, input)
}

func (a *MediaTailorActivities) ListTagsForResource(ctx context.Context, input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaTailorActivities) PutPlaybackConfiguration(ctx context.Context, input *mediatailor.PutPlaybackConfigurationInput) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPlaybackConfigurationWithContext(ctx, input)
}

func (a *MediaTailorActivities) TagResource(ctx context.Context, input *mediatailor.TagResourceInput) (*mediatailor.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *MediaTailorActivities) UntagResource(ctx context.Context, input *mediatailor.UntagResourceInput) (*mediatailor.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
