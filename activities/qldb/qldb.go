// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldb/qldbiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client qldbiface.QLDBAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := qldb.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (qldbiface.QLDBAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return qldb.New(sess), nil
}

func (a *Activities) CancelJournalKinesisStream(ctx context.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CancelJournalKinesisStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateLedger(ctx context.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateLedgerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteLedger(ctx context.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteLedgerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeJournalKinesisStream(ctx context.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeJournalKinesisStreamWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeJournalS3Export(ctx context.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeJournalS3ExportWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeLedger(ctx context.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeLedgerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExportJournalToS3(ctx context.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExportJournalToS3WithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetBlock(ctx context.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetBlockWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetDigest(ctx context.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetDigestWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetRevision(ctx context.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetRevisionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJournalKinesisStreamsForLedger(ctx context.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJournalKinesisStreamsForLedgerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJournalS3Exports(ctx context.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJournalS3ExportsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListJournalS3ExportsForLedger(ctx context.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListJournalS3ExportsForLedgerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListLedgers(ctx context.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListLedgersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StreamJournalToKinesis(ctx context.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StreamJournalToKinesisWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateLedger(ctx context.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateLedgerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}
