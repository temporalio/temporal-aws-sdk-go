package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type KinesisAnalyticsActivities struct {
	client kinesisanalyticsiface.KinesisAnalyticsAPI
}

func NewKinesisAnalyticsActivities(session *session.Session, config ...*aws.Config) *KinesisAnalyticsActivities {
	client := kinesisanalytics.New(session, config...)
	return &KinesisAnalyticsActivities{client: client}
}

func (a *KinesisAnalyticsActivities) AddApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	return a.client.AddApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationInput(ctx context.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	return a.client.AddApplicationInputWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	return a.client.AddApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationOutput(ctx context.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	return a.client.AddApplicationOutputWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) AddApplicationReferenceDataSource(ctx context.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	return a.client.AddApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) CreateApplication(ctx context.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplication(ctx context.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	return a.client.DeleteApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	return a.client.DeleteApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationOutput(ctx context.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	return a.client.DeleteApplicationOutputWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationReferenceDataSource(ctx context.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	return a.client.DeleteApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DescribeApplication(ctx context.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	return a.client.DescribeApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) DiscoverInputSchema(ctx context.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	return a.client.DiscoverInputSchemaWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) ListApplications(ctx context.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	return a.client.ListApplicationsWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) ListTagsForResource(ctx context.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) StartApplication(ctx context.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	return a.client.StartApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) StopApplication(ctx context.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	return a.client.StopApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) TagResource(ctx context.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) UntagResource(ctx context.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsActivities) UpdateApplication(ctx context.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	return a.client.UpdateApplicationWithContext(ctx, input)
}
