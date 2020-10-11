// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldb"
	"github.com/aws/aws-sdk-go/service/qldb/qldbiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type QLDBActivities struct {
	client qldbiface.QLDBAPI

	sessionFactory SessionFactory
}

func NewQLDBActivities(sess *session.Session, config ...*aws.Config) *QLDBActivities {
	client := qldb.New(sess, config...)
	return &QLDBActivities{client: client}
}

func NewQLDBActivitiesWithSessionFactory(sessionFactory SessionFactory) *QLDBActivities {
	return &QLDBActivities{sessionFactory: sessionFactory}
}

func (a *QLDBActivities) getClient(ctx context.Context) (qldbiface.QLDBAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return qldb.New(sess), nil
}

func (a *QLDBActivities) CancelJournalKinesisStream(ctx context.Context, input *qldb.CancelJournalKinesisStreamInput) (*qldb.CancelJournalKinesisStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelJournalKinesisStreamWithContext(ctx, input)
}

func (a *QLDBActivities) CreateLedger(ctx context.Context, input *qldb.CreateLedgerInput) (*qldb.CreateLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) DeleteLedger(ctx context.Context, input *qldb.DeleteLedgerInput) (*qldb.DeleteLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) DescribeJournalKinesisStream(ctx context.Context, input *qldb.DescribeJournalKinesisStreamInput) (*qldb.DescribeJournalKinesisStreamOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeJournalKinesisStreamWithContext(ctx, input)
}

func (a *QLDBActivities) DescribeJournalS3Export(ctx context.Context, input *qldb.DescribeJournalS3ExportInput) (*qldb.DescribeJournalS3ExportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeJournalS3ExportWithContext(ctx, input)
}

func (a *QLDBActivities) DescribeLedger(ctx context.Context, input *qldb.DescribeLedgerInput) (*qldb.DescribeLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) ExportJournalToS3(ctx context.Context, input *qldb.ExportJournalToS3Input) (*qldb.ExportJournalToS3Output, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ExportJournalToS3WithContext(ctx, input)
}

func (a *QLDBActivities) GetBlock(ctx context.Context, input *qldb.GetBlockInput) (*qldb.GetBlockOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBlockWithContext(ctx, input)
}

func (a *QLDBActivities) GetDigest(ctx context.Context, input *qldb.GetDigestInput) (*qldb.GetDigestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDigestWithContext(ctx, input)
}

func (a *QLDBActivities) GetRevision(ctx context.Context, input *qldb.GetRevisionInput) (*qldb.GetRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRevisionWithContext(ctx, input)
}

func (a *QLDBActivities) ListJournalKinesisStreamsForLedger(ctx context.Context, input *qldb.ListJournalKinesisStreamsForLedgerInput) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListJournalKinesisStreamsForLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) ListJournalS3Exports(ctx context.Context, input *qldb.ListJournalS3ExportsInput) (*qldb.ListJournalS3ExportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListJournalS3ExportsWithContext(ctx, input)
}

func (a *QLDBActivities) ListJournalS3ExportsForLedger(ctx context.Context, input *qldb.ListJournalS3ExportsForLedgerInput) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListJournalS3ExportsForLedgerWithContext(ctx, input)
}

func (a *QLDBActivities) ListLedgers(ctx context.Context, input *qldb.ListLedgersInput) (*qldb.ListLedgersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLedgersWithContext(ctx, input)
}

func (a *QLDBActivities) ListTagsForResource(ctx context.Context, input *qldb.ListTagsForResourceInput) (*qldb.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *QLDBActivities) StreamJournalToKinesis(ctx context.Context, input *qldb.StreamJournalToKinesisInput) (*qldb.StreamJournalToKinesisOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StreamJournalToKinesisWithContext(ctx, input)
}

func (a *QLDBActivities) TagResource(ctx context.Context, input *qldb.TagResourceInput) (*qldb.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *QLDBActivities) UntagResource(ctx context.Context, input *qldb.UntagResourceInput) (*qldb.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *QLDBActivities) UpdateLedger(ctx context.Context, input *qldb.UpdateLedgerInput) (*qldb.UpdateLedgerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateLedgerWithContext(ctx, input)
}
