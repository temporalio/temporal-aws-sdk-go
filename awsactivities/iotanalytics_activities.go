package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotanalytics/iotanalyticsiface"
)

type IoTAnalyticsActivities struct {
	client iotanalyticsiface.IoTAnalyticsAPI
}

func NewIoTAnalyticsActivities(client iotanalyticsiface.IoTAnalyticsAPI) *IoTAnalyticsActivities {
    return &IoTAnalyticsActivities{client: client}
}


func (a *IoTAnalyticsActivities) BatchPutMessage(input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
    return a.client.BatchPutMessage(input)
}



func (a *IoTAnalyticsActivities) CancelPipelineReprocessing(input *iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error) {
    return a.client.CancelPipelineReprocessing(input)
}



func (a *IoTAnalyticsActivities) CreateChannel(input *iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error) {
    return a.client.CreateChannel(input)
}



func (a *IoTAnalyticsActivities) CreateDataset(input *iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error) {
    return a.client.CreateDataset(input)
}



func (a *IoTAnalyticsActivities) CreateDatasetContent(input *iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error) {
    return a.client.CreateDatasetContent(input)
}



func (a *IoTAnalyticsActivities) CreateDatastore(input *iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error) {
    return a.client.CreateDatastore(input)
}



func (a *IoTAnalyticsActivities) CreatePipeline(input *iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error) {
    return a.client.CreatePipeline(input)
}



func (a *IoTAnalyticsActivities) DeleteChannel(input *iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error) {
    return a.client.DeleteChannel(input)
}



func (a *IoTAnalyticsActivities) DeleteDataset(input *iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error) {
    return a.client.DeleteDataset(input)
}



func (a *IoTAnalyticsActivities) DeleteDatasetContent(input *iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error) {
    return a.client.DeleteDatasetContent(input)
}



func (a *IoTAnalyticsActivities) DeleteDatastore(input *iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error) {
    return a.client.DeleteDatastore(input)
}



func (a *IoTAnalyticsActivities) DeletePipeline(input *iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error) {
    return a.client.DeletePipeline(input)
}



func (a *IoTAnalyticsActivities) DescribeChannel(input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error) {
    return a.client.DescribeChannel(input)
}



func (a *IoTAnalyticsActivities) DescribeDataset(input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error) {
    return a.client.DescribeDataset(input)
}



func (a *IoTAnalyticsActivities) DescribeDatastore(input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error) {
    return a.client.DescribeDatastore(input)
}



func (a *IoTAnalyticsActivities) DescribeLoggingOptions(input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
    return a.client.DescribeLoggingOptions(input)
}



func (a *IoTAnalyticsActivities) DescribePipeline(input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error) {
    return a.client.DescribePipeline(input)
}



func (a *IoTAnalyticsActivities) GetDatasetContent(input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error) {
    return a.client.GetDatasetContent(input)
}



func (a *IoTAnalyticsActivities) ListChannels(input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error) {
    return a.client.ListChannels(input)
}



func (a *IoTAnalyticsActivities) ListDatasetContents(input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error) {
    return a.client.ListDatasetContents(input)
}



func (a *IoTAnalyticsActivities) ListDatasets(input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error) {
    return a.client.ListDatasets(input)
}



func (a *IoTAnalyticsActivities) ListDatastores(input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error) {
    return a.client.ListDatastores(input)
}



func (a *IoTAnalyticsActivities) ListPipelines(input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error) {
    return a.client.ListPipelines(input)
}



func (a *IoTAnalyticsActivities) ListTagsForResource(input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}



func (a *IoTAnalyticsActivities) PutLoggingOptions(input *iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error) {
    return a.client.PutLoggingOptions(input)
}



func (a *IoTAnalyticsActivities) RunPipelineActivity(input *iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error) {
    return a.client.RunPipelineActivity(input)
}



func (a *IoTAnalyticsActivities) SampleChannelData(input *iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error) {
    return a.client.SampleChannelData(input)
}



func (a *IoTAnalyticsActivities) StartPipelineReprocessing(input *iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error) {
    return a.client.StartPipelineReprocessing(input)
}



func (a *IoTAnalyticsActivities) TagResource(input *iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error) {
    return a.client.TagResource(input)
}



func (a *IoTAnalyticsActivities) UntagResource(input *iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}



func (a *IoTAnalyticsActivities) UpdateChannel(input *iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error) {
    return a.client.UpdateChannel(input)
}



func (a *IoTAnalyticsActivities) UpdateDataset(input *iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error) {
    return a.client.UpdateDataset(input)
}



func (a *IoTAnalyticsActivities) UpdateDatastore(input *iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error) {
    return a.client.UpdateDatastore(input)
}



func (a *IoTAnalyticsActivities) UpdatePipeline(input *iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error) {
    return a.client.UpdatePipeline(input)
}

