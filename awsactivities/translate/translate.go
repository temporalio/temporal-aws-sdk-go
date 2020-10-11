// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package translate

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/translate/translateiface"
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
	client translateiface.TranslateAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := translate.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (translateiface.TranslateAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return translate.New(sess), nil
}

func (a *Activities) DeleteTerminology(ctx context.Context, input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTerminologyWithContext(ctx, input)
}

func (a *Activities) DescribeTextTranslationJob(ctx context.Context, input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTextTranslationJobWithContext(ctx, input)
}

func (a *Activities) GetTerminology(ctx context.Context, input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTerminologyWithContext(ctx, input)
}

func (a *Activities) ImportTerminology(ctx context.Context, input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ImportTerminologyWithContext(ctx, input)
}

func (a *Activities) ListTerminologies(ctx context.Context, input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTerminologiesWithContext(ctx, input)
}

func (a *Activities) ListTextTranslationJobs(ctx context.Context, input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTextTranslationJobsWithContext(ctx, input)
}

func (a *Activities) StartTextTranslationJob(ctx context.Context, input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.StartTextTranslationJobWithContext(ctx, input)
}

func (a *Activities) StopTextTranslationJob(ctx context.Context, input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopTextTranslationJobWithContext(ctx, input)
}

func (a *Activities) Text(ctx context.Context, input *translate.TextInput) (*translate.TextOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TextWithContext(ctx, input)
}
