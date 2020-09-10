package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/forecastservice"
	"github.com/aws/aws-sdk-go/service/forecastservice/forecastserviceiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ForecastServiceActivities struct {
	client forecastserviceiface.ForecastServiceAPI
}

func NewForecastServiceActivities(session *session.Session, config ...*aws.Config) *ForecastServiceActivities {
	client := forecastservice.New(session, config...)
	return &ForecastServiceActivities{client: client}
}

func (a *ForecastServiceActivities) CreateDataset(ctx context.Context, input *forecastservice.CreateDatasetInput) (*forecastservice.CreateDatasetOutput, error) {
	return a.client.CreateDatasetWithContext(ctx, input)
}

func (a *ForecastServiceActivities) CreateDatasetGroup(ctx context.Context, input *forecastservice.CreateDatasetGroupInput) (*forecastservice.CreateDatasetGroupOutput, error) {
	return a.client.CreateDatasetGroupWithContext(ctx, input)
}

func (a *ForecastServiceActivities) CreateDatasetImportJob(ctx context.Context, input *forecastservice.CreateDatasetImportJobInput) (*forecastservice.CreateDatasetImportJobOutput, error) {
	return a.client.CreateDatasetImportJobWithContext(ctx, input)
}

func (a *ForecastServiceActivities) CreateForecast(ctx context.Context, input *forecastservice.CreateForecastInput) (*forecastservice.CreateForecastOutput, error) {
	return a.client.CreateForecastWithContext(ctx, input)
}

func (a *ForecastServiceActivities) CreateForecastExportJob(ctx context.Context, input *forecastservice.CreateForecastExportJobInput) (*forecastservice.CreateForecastExportJobOutput, error) {
	return a.client.CreateForecastExportJobWithContext(ctx, input)
}

func (a *ForecastServiceActivities) CreatePredictor(ctx context.Context, input *forecastservice.CreatePredictorInput) (*forecastservice.CreatePredictorOutput, error) {
	return a.client.CreatePredictorWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DeleteDataset(ctx context.Context, input *forecastservice.DeleteDatasetInput) (*forecastservice.DeleteDatasetOutput, error) {
	return a.client.DeleteDatasetWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DeleteDatasetGroup(ctx context.Context, input *forecastservice.DeleteDatasetGroupInput) (*forecastservice.DeleteDatasetGroupOutput, error) {
	return a.client.DeleteDatasetGroupWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DeleteDatasetImportJob(ctx context.Context, input *forecastservice.DeleteDatasetImportJobInput) (*forecastservice.DeleteDatasetImportJobOutput, error) {
	return a.client.DeleteDatasetImportJobWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DeleteForecast(ctx context.Context, input *forecastservice.DeleteForecastInput) (*forecastservice.DeleteForecastOutput, error) {
	return a.client.DeleteForecastWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DeleteForecastExportJob(ctx context.Context, input *forecastservice.DeleteForecastExportJobInput) (*forecastservice.DeleteForecastExportJobOutput, error) {
	return a.client.DeleteForecastExportJobWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DeletePredictor(ctx context.Context, input *forecastservice.DeletePredictorInput) (*forecastservice.DeletePredictorOutput, error) {
	return a.client.DeletePredictorWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DescribeDataset(ctx context.Context, input *forecastservice.DescribeDatasetInput) (*forecastservice.DescribeDatasetOutput, error) {
	return a.client.DescribeDatasetWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DescribeDatasetGroup(ctx context.Context, input *forecastservice.DescribeDatasetGroupInput) (*forecastservice.DescribeDatasetGroupOutput, error) {
	return a.client.DescribeDatasetGroupWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DescribeDatasetImportJob(ctx context.Context, input *forecastservice.DescribeDatasetImportJobInput) (*forecastservice.DescribeDatasetImportJobOutput, error) {
	return a.client.DescribeDatasetImportJobWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DescribeForecast(ctx context.Context, input *forecastservice.DescribeForecastInput) (*forecastservice.DescribeForecastOutput, error) {
	return a.client.DescribeForecastWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DescribeForecastExportJob(ctx context.Context, input *forecastservice.DescribeForecastExportJobInput) (*forecastservice.DescribeForecastExportJobOutput, error) {
	return a.client.DescribeForecastExportJobWithContext(ctx, input)
}

func (a *ForecastServiceActivities) DescribePredictor(ctx context.Context, input *forecastservice.DescribePredictorInput) (*forecastservice.DescribePredictorOutput, error) {
	return a.client.DescribePredictorWithContext(ctx, input)
}

func (a *ForecastServiceActivities) GetAccuracyMetrics(ctx context.Context, input *forecastservice.GetAccuracyMetricsInput) (*forecastservice.GetAccuracyMetricsOutput, error) {
	return a.client.GetAccuracyMetricsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListDatasetGroups(ctx context.Context, input *forecastservice.ListDatasetGroupsInput) (*forecastservice.ListDatasetGroupsOutput, error) {
	return a.client.ListDatasetGroupsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListDatasetImportJobs(ctx context.Context, input *forecastservice.ListDatasetImportJobsInput) (*forecastservice.ListDatasetImportJobsOutput, error) {
	return a.client.ListDatasetImportJobsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListDatasets(ctx context.Context, input *forecastservice.ListDatasetsInput) (*forecastservice.ListDatasetsOutput, error) {
	return a.client.ListDatasetsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListForecastExportJobs(ctx context.Context, input *forecastservice.ListForecastExportJobsInput) (*forecastservice.ListForecastExportJobsOutput, error) {
	return a.client.ListForecastExportJobsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListForecasts(ctx context.Context, input *forecastservice.ListForecastsInput) (*forecastservice.ListForecastsOutput, error) {
	return a.client.ListForecastsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListPredictors(ctx context.Context, input *forecastservice.ListPredictorsInput) (*forecastservice.ListPredictorsOutput, error) {
	return a.client.ListPredictorsWithContext(ctx, input)
}

func (a *ForecastServiceActivities) ListTagsForResource(ctx context.Context, input *forecastservice.ListTagsForResourceInput) (*forecastservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ForecastServiceActivities) TagResource(ctx context.Context, input *forecastservice.TagResourceInput) (*forecastservice.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *ForecastServiceActivities) UntagResource(ctx context.Context, input *forecastservice.UntagResourceInput) (*forecastservice.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *ForecastServiceActivities) UpdateDatasetGroup(ctx context.Context, input *forecastservice.UpdateDatasetGroupInput) (*forecastservice.UpdateDatasetGroupOutput, error) {
	return a.client.UpdateDatasetGroupWithContext(ctx, input)
}
