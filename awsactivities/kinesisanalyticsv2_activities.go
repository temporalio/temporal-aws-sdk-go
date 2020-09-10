package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2/kinesisanalyticsv2iface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type KinesisAnalyticsV2Activities struct {
	client kinesisanalyticsv2iface.KinesisAnalyticsV2API
}

func NewKinesisAnalyticsV2Activities(session *session.Session, config ...*aws.Config) *KinesisAnalyticsV2Activities {
	client := kinesisanalyticsv2.New(session, config...)
	return &KinesisAnalyticsV2Activities{client: client}
}

func (a *KinesisAnalyticsV2Activities) AddApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
	return a.client.AddApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationInput(ctx context.Context, input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error) {
	return a.client.AddApplicationInputWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error) {
	return a.client.AddApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationOutput(ctx context.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error) {
	return a.client.AddApplicationOutputWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationReferenceDataSource(ctx context.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error) {
	return a.client.AddApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationVpcConfiguration(ctx context.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error) {
	return a.client.AddApplicationVpcConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) CreateApplication(ctx context.Context, input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error) {
	return a.client.CreateApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) CreateApplicationSnapshot(ctx context.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error) {
	return a.client.CreateApplicationSnapshotWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplication(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error) {
	return a.client.DeleteApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationCloudWatchLoggingOption(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	return a.client.DeleteApplicationCloudWatchLoggingOptionWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationInputProcessingConfiguration(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error) {
	return a.client.DeleteApplicationInputProcessingConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationOutput(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error) {
	return a.client.DeleteApplicationOutputWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationReferenceDataSource(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error) {
	return a.client.DeleteApplicationReferenceDataSourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationSnapshot(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error) {
	return a.client.DeleteApplicationSnapshotWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationVpcConfiguration(ctx context.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error) {
	return a.client.DeleteApplicationVpcConfigurationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DescribeApplication(ctx context.Context, input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
	return a.client.DescribeApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DescribeApplicationSnapshot(ctx context.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
	return a.client.DescribeApplicationSnapshotWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) DiscoverInputSchema(ctx context.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error) {
	return a.client.DiscoverInputSchemaWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) ListApplicationSnapshots(ctx context.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
	return a.client.ListApplicationSnapshotsWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) ListApplications(ctx context.Context, input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
	return a.client.ListApplicationsWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) ListTagsForResource(ctx context.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) StartApplication(ctx context.Context, input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error) {
	return a.client.StartApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) StopApplication(ctx context.Context, input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error) {
	return a.client.StopApplicationWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) TagResource(ctx context.Context, input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) UntagResource(ctx context.Context, input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *KinesisAnalyticsV2Activities) UpdateApplication(ctx context.Context, input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error) {
	return a.client.UpdateApplicationWithContext(ctx, input)
}
