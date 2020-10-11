// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type KinesisAnalyticsActivities struct {
	client kinesisanalyticsiface.KinesisAnalyticsAPI

	sessionFactory SessionFactory
}

func NewKinesisAnalyticsActivities(sess *session.Session, config ...*aws.Config) *KinesisAnalyticsActivities {
	client := kinesisanalytics.New(sess, config...)
	return &KinesisAnalyticsActivities{client: client}
}

func NewKinesisAnalyticsActivitiesWithSessionFactory(sessionFactory SessionFactory) *KinesisAnalyticsActivities {
	return &KinesisAnalyticsActivities{sessionFactory: sessionFactory}
}

func (a *KinesisAnalyticsActivities) getClient(ctx context.Context) (kinesisanalyticsiface.KinesisAnalyticsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return kinesisanalytics.New(sess), nil
}

func (a *KinesisAnalyticsActivities) AddApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationInput(ctx context.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationInputWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationOutput(ctx context.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationOutputWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationReferenceDataSource(ctx context.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) CreateApplication(ctx context.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplication(ctx context.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationOutput(ctx context.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationOutputWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationReferenceDataSource(ctx context.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DescribeApplication(ctx context.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DiscoverInputSchema(ctx context.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DiscoverInputSchemaWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) ListApplications(ctx context.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationsWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) ListTagsForResource(ctx context.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) StartApplication(ctx context.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) StopApplication(ctx context.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) TagResource(ctx context.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) UntagResource(ctx context.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) UpdateApplication(ctx context.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApplicationWithContext(ctx, input)
}
