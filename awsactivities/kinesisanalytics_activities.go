
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
)

type KinesisAnalyticsActivities struct {
	client kinesisanalyticsiface.KinesisAnalyticsAPI
}

func NewKinesisAnalyticsActivities(client kinesisanalyticsiface.KinesisAnalyticsAPI) *KinesisAnalyticsActivities {
    return &KinesisAnalyticsActivities{client: client}
}

func (a *KinesisAnalyticsActivities) AddApplicationCloudWatchLoggingOption(input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
    return a.client.AddApplicationCloudWatchLoggingOption(input)
}

func (a *KinesisAnalyticsActivities) AddApplicationInput(input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
    return a.client.AddApplicationInput(input)
}

func (a *KinesisAnalyticsActivities) AddApplicationInputProcessingConfiguration(input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
    return a.client.AddApplicationInputProcessingConfiguration(input)
}

func (a *KinesisAnalyticsActivities) AddApplicationOutput(input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
    return a.client.AddApplicationOutput(input)
}

func (a *KinesisAnalyticsActivities) AddApplicationReferenceDataSource(input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
    return a.client.AddApplicationReferenceDataSource(input)
}

func (a *KinesisAnalyticsActivities) CreateApplication(input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
    return a.client.CreateApplication(input)
}

func (a *KinesisAnalyticsActivities) DeleteApplication(input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
    return a.client.DeleteApplication(input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationCloudWatchLoggingOption(input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
    return a.client.DeleteApplicationCloudWatchLoggingOption(input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationInputProcessingConfiguration(input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
    return a.client.DeleteApplicationInputProcessingConfiguration(input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationOutput(input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
    return a.client.DeleteApplicationOutput(input)
}

func (a *KinesisAnalyticsActivities) DeleteApplicationReferenceDataSource(input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
    return a.client.DeleteApplicationReferenceDataSource(input)
}

func (a *KinesisAnalyticsActivities) DescribeApplication(input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
    return a.client.DescribeApplication(input)
}

func (a *KinesisAnalyticsActivities) DiscoverInputSchema(input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
    return a.client.DiscoverInputSchema(input)
}

func (a *KinesisAnalyticsActivities) ListApplications(input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
    return a.client.ListApplications(input)
}

func (a *KinesisAnalyticsActivities) ListTagsForResource(input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *KinesisAnalyticsActivities) StartApplication(input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
    return a.client.StartApplication(input)
}

func (a *KinesisAnalyticsActivities) StopApplication(input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
    return a.client.StopApplication(input)
}

func (a *KinesisAnalyticsActivities) TagResource(input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *KinesisAnalyticsActivities) UntagResource(input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *KinesisAnalyticsActivities) UpdateApplication(input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
    return a.client.UpdateApplication(input)
}
