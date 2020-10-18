// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/fsx/fsxiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client fsxiface.FSxAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := fsx.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (fsxiface.FSxAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return fsx.New(sess), nil
}

func (a *Activities) CancelDataRepositoryTask(ctx context.Context, input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelDataRepositoryTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateBackup(ctx context.Context, input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDataRepositoryTask(ctx context.Context, input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDataRepositoryTaskWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFileSystem(ctx context.Context, input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFileSystemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFileSystemFromBackup(ctx context.Context, input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFileSystemFromBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteBackup(ctx context.Context, input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteBackupWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFileSystem(ctx context.Context, input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFileSystemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeBackups(ctx context.Context, input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeBackupsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDataRepositoryTasks(ctx context.Context, input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDataRepositoryTasksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFileSystems(ctx context.Context, input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFileSystemsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFileSystem(ctx context.Context, input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFileSystemWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
