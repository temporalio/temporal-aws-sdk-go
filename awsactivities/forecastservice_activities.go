package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/forecastservice/forecastserviceiface"
)

type ForecastServiceActivities struct {
	client forecastserviceiface.ForecastServiceAPI
}

func NewForecastServiceActivities(client forecastserviceiface.ForecastServiceAPI) *ForecastServiceActivities {
    return &ForecastServiceActivities{client: client}
}

func (a *ForecastServiceActivities) CreateDataset(input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error) {
    return a.client.CreateDataset(input)
}

func (a *ForecastServiceActivities) CreateDatasetGroup(input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error) {
    return a.client.CreateDatasetGroup(input)
}

func (a *ForecastServiceActivities) CreateDatasetImportJob(input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error) {
    return a.client.CreateDatasetImportJob(input)
}

func (a *ForecastServiceActivities) CreateForecast(input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error) {
    return a.client.CreateForecast(input)
}

func (a *ForecastServiceActivities) CreateForecastExportJob(input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error) {
    return a.client.CreateForecastExportJob(input)
}

func (a *ForecastServiceActivities) CreatePredictor(input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error) {
    return a.client.CreatePredictor(input)
}

func (a *ForecastServiceActivities) DeleteDataset(input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error) {
    return a.client.DeleteDataset(input)
}

func (a *ForecastServiceActivities) DeleteDatasetGroup(input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error) {
    return a.client.DeleteDatasetGroup(input)
}

func (a *ForecastServiceActivities) DeleteDatasetImportJob(input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error) {
    return a.client.DeleteDatasetImportJob(input)
}

func (a *ForecastServiceActivities) DeleteForecast(input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error) {
    return a.client.DeleteForecast(input)
}

func (a *ForecastServiceActivities) DeleteForecastExportJob(input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error) {
    return a.client.DeleteForecastExportJob(input)
}

func (a *ForecastServiceActivities) DeletePredictor(input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error) {
    return a.client.DeletePredictor(input)
}

func (a *ForecastServiceActivities) DescribeDataset(input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error) {
    return a.client.DescribeDataset(input)
}

func (a *ForecastServiceActivities) DescribeDatasetGroup(input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error) {
    return a.client.DescribeDatasetGroup(input)
}

func (a *ForecastServiceActivities) DescribeDatasetImportJob(input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error) {
    return a.client.DescribeDatasetImportJob(input)
}

func (a *ForecastServiceActivities) DescribeForecast(input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error) {
    return a.client.DescribeForecast(input)
}

func (a *ForecastServiceActivities) DescribeForecastExportJob(input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error) {
    return a.client.DescribeForecastExportJob(input)
}

func (a *ForecastServiceActivities) DescribePredictor(input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error) {
    return a.client.DescribePredictor(input)
}

func (a *ForecastServiceActivities) GetAccuracyMetrics(input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error) {
    return a.client.GetAccuracyMetrics(input)
}

func (a *ForecastServiceActivities) ListDatasetGroups(input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error) {
    return a.client.ListDatasetGroups(input)
}

func (a *ForecastServiceActivities) ListDatasetImportJobs(input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error) {
    return a.client.ListDatasetImportJobs(input)
}

func (a *ForecastServiceActivities) ListDatasets(input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error) {
    return a.client.ListDatasets(input)
}

func (a *ForecastServiceActivities) ListForecastExportJobs(input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error) {
    return a.client.ListForecastExportJobs(input)
}

func (a *ForecastServiceActivities) ListForecasts(input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error) {
    return a.client.ListForecasts(input)
}

func (a *ForecastServiceActivities) ListPredictors(input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error) {
    return a.client.ListPredictors(input)
}

func (a *ForecastServiceActivities) ListTagsForResource(input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *ForecastServiceActivities) TagResource(input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *ForecastServiceActivities) UntagResource(input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *ForecastServiceActivities) UpdateDatasetGroup(input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error) {
    return a.client.UpdateDatasetGroup(input)
}
