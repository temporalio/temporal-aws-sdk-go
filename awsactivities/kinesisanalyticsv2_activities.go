
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2/kinesisanalyticsv2iface"
)

type KinesisAnalyticsV2Activities struct {
	client kinesisanalyticsv2iface.KinesisAnalyticsV2API
}

func NewKinesisAnalyticsV2Activities(session *session.Session, config... *aws.Config) *KinesisAnalyticsV2Activities {
    client := kinesisanalyticsv2.New(session, config...)
    return &KinesisAnalyticsV2Activities{client: client}
}

func (a *KinesisAnalyticsV2Activities) AddApplicationCloudWatchLoggingOption(input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
    return a.client.AddApplicationCloudWatchLoggingOption(input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationInput(input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error) {
    return a.client.AddApplicationInput(input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationInputProcessingConfiguration(input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error) {
    return a.client.AddApplicationInputProcessingConfiguration(input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationOutput(input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error) {
    return a.client.AddApplicationOutput(input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationReferenceDataSource(input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error) {
    return a.client.AddApplicationReferenceDataSource(input)
}

func (a *KinesisAnalyticsV2Activities) AddApplicationVpcConfiguration(input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error) {
    return a.client.AddApplicationVpcConfiguration(input)
}

func (a *KinesisAnalyticsV2Activities) CreateApplication(input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error) {
    return a.client.CreateApplication(input)
}

func (a *KinesisAnalyticsV2Activities) CreateApplicationSnapshot(input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error) {
    return a.client.CreateApplicationSnapshot(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplication(input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error) {
    return a.client.DeleteApplication(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationCloudWatchLoggingOption(input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
    return a.client.DeleteApplicationCloudWatchLoggingOption(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationInputProcessingConfiguration(input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error) {
    return a.client.DeleteApplicationInputProcessingConfiguration(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationOutput(input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error) {
    return a.client.DeleteApplicationOutput(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationReferenceDataSource(input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error) {
    return a.client.DeleteApplicationReferenceDataSource(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationSnapshot(input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error) {
    return a.client.DeleteApplicationSnapshot(input)
}

func (a *KinesisAnalyticsV2Activities) DeleteApplicationVpcConfiguration(input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error) {
    return a.client.DeleteApplicationVpcConfiguration(input)
}

func (a *KinesisAnalyticsV2Activities) DescribeApplication(input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
    return a.client.DescribeApplication(input)
}

func (a *KinesisAnalyticsV2Activities) DescribeApplicationSnapshot(input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
    return a.client.DescribeApplicationSnapshot(input)
}

func (a *KinesisAnalyticsV2Activities) DiscoverInputSchema(input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error) {
    return a.client.DiscoverInputSchema(input)
}

func (a *KinesisAnalyticsV2Activities) ListApplicationSnapshots(input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
    return a.client.ListApplicationSnapshots(input)
}

func (a *KinesisAnalyticsV2Activities) ListApplications(input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
    return a.client.ListApplications(input)
}

func (a *KinesisAnalyticsV2Activities) ListTagsForResource(input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *KinesisAnalyticsV2Activities) StartApplication(input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error) {
    return a.client.StartApplication(input)
}

func (a *KinesisAnalyticsV2Activities) StopApplication(input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error) {
    return a.client.StopApplication(input)
}

func (a *KinesisAnalyticsV2Activities) TagResource(input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *KinesisAnalyticsV2Activities) UntagResource(input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *KinesisAnalyticsV2Activities) UpdateApplication(input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error) {
    return a.client.UpdateApplication(input)
}
