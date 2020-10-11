// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/dataexchange/dataexchangeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DataExchangeActivities struct {
	client dataexchangeiface.DataExchangeAPI

	sessionFactory SessionFactory
}

func NewDataExchangeActivities(sess *session.Session, config ...*aws.Config) *DataExchangeActivities {
	client := dataexchange.New(sess, config...)
	return &DataExchangeActivities{client: client}
}

func NewDataExchangeActivitiesWithSessionFactory(sessionFactory SessionFactory) *DataExchangeActivities {
	return &DataExchangeActivities{sessionFactory: sessionFactory}
}

func (a *DataExchangeActivities) getClient(ctx context.Context) (dataexchangeiface.DataExchangeAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return dataexchange.New(sess), nil
}

func (a *DataExchangeActivities) CancelJob(ctx context.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) CreateDataSet(ctx context.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) CreateJob(ctx context.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) CreateRevision(ctx context.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRevisionWithContext(ctx, input)
}

func (a *DataExchangeActivities) DeleteAsset(ctx context.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAssetWithContext(ctx, input)
}

func (a *DataExchangeActivities) DeleteDataSet(ctx context.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) DeleteRevision(ctx context.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRevisionWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetAsset(ctx context.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssetWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetDataSet(ctx context.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetJob(ctx context.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) GetRevision(ctx context.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRevisionWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListDataSetRevisions(ctx context.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDataSetRevisionsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListDataSets(ctx context.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDataSetsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListJobs(ctx context.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListJobsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListRevisionAssets(ctx context.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRevisionAssetsWithContext(ctx, input)
}

func (a *DataExchangeActivities) ListTagsForResource(ctx context.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DataExchangeActivities) StartJob(ctx context.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartJobWithContext(ctx, input)
}

func (a *DataExchangeActivities) TagResource(ctx context.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *DataExchangeActivities) UntagResource(ctx context.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *DataExchangeActivities) UpdateAsset(ctx context.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAssetWithContext(ctx, input)
}

func (a *DataExchangeActivities) UpdateDataSet(ctx context.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDataSetWithContext(ctx, input)
}

func (a *DataExchangeActivities) UpdateRevision(ctx context.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRevisionWithContext(ctx, input)
}
