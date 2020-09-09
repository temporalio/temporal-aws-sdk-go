
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface"
)

type MachineLearningActivities struct {
	client machinelearningiface.MachineLearningAPI
}

func NewMachineLearningActivities(session *session.Session, config... *aws.Config) *MachineLearningActivities {
    client := machinelearning.New(session, config...)
    return &MachineLearningActivities{client: client}
}

func (a *MachineLearningActivities) AddTags(input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error) {
    return a.client.AddTags(input)
}

func (a *MachineLearningActivities) CreateBatchPrediction(input *machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error) {
    return a.client.CreateBatchPrediction(input)
}

func (a *MachineLearningActivities) CreateDataSourceFromRDS(input *machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
    return a.client.CreateDataSourceFromRDS(input)
}

func (a *MachineLearningActivities) CreateDataSourceFromRedshift(input *machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
    return a.client.CreateDataSourceFromRedshift(input)
}

func (a *MachineLearningActivities) CreateDataSourceFromS3(input *machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error) {
    return a.client.CreateDataSourceFromS3(input)
}

func (a *MachineLearningActivities) CreateEvaluation(input *machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error) {
    return a.client.CreateEvaluation(input)
}

func (a *MachineLearningActivities) CreateMLModel(input *machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error) {
    return a.client.CreateMLModel(input)
}

func (a *MachineLearningActivities) CreateRealtimeEndpoint(input *machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error) {
    return a.client.CreateRealtimeEndpoint(input)
}

func (a *MachineLearningActivities) DeleteBatchPrediction(input *machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error) {
    return a.client.DeleteBatchPrediction(input)
}

func (a *MachineLearningActivities) DeleteDataSource(input *machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error) {
    return a.client.DeleteDataSource(input)
}

func (a *MachineLearningActivities) DeleteEvaluation(input *machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error) {
    return a.client.DeleteEvaluation(input)
}

func (a *MachineLearningActivities) DeleteMLModel(input *machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error) {
    return a.client.DeleteMLModel(input)
}

func (a *MachineLearningActivities) DeleteRealtimeEndpoint(input *machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
    return a.client.DeleteRealtimeEndpoint(input)
}

func (a *MachineLearningActivities) DeleteTags(input *machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error) {
    return a.client.DeleteTags(input)
}

func (a *MachineLearningActivities) DescribeBatchPredictions(input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error) {
    return a.client.DescribeBatchPredictions(input)
}

func (a *MachineLearningActivities) DescribeDataSources(input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error) {
    return a.client.DescribeDataSources(input)
}

func (a *MachineLearningActivities) DescribeEvaluations(input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error) {
    return a.client.DescribeEvaluations(input)
}

func (a *MachineLearningActivities) DescribeMLModels(input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error) {
    return a.client.DescribeMLModels(input)
}

func (a *MachineLearningActivities) DescribeTags(input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error) {
    return a.client.DescribeTags(input)
}

func (a *MachineLearningActivities) GetBatchPrediction(input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error) {
    return a.client.GetBatchPrediction(input)
}

func (a *MachineLearningActivities) GetDataSource(input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error) {
    return a.client.GetDataSource(input)
}

func (a *MachineLearningActivities) GetEvaluation(input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error) {
    return a.client.GetEvaluation(input)
}

func (a *MachineLearningActivities) GetMLModel(input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error) {
    return a.client.GetMLModel(input)
}

func (a *MachineLearningActivities) Predict(input *machinelearning.PredictInput) (*machinelearning.PredictOutput, error) {
    return a.client.Predict(input)
}

func (a *MachineLearningActivities) UpdateBatchPrediction(input *machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error) {
    return a.client.UpdateBatchPrediction(input)
}

func (a *MachineLearningActivities) UpdateDataSource(input *machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error) {
    return a.client.UpdateDataSource(input)
}

func (a *MachineLearningActivities) UpdateEvaluation(input *machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error) {
    return a.client.UpdateEvaluation(input)
}

func (a *MachineLearningActivities) UpdateMLModel(input *machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error) {
    return a.client.UpdateMLModel(input)
}

func (a *MachineLearningActivities) WaitUntilBatchPredictionAvailable(input *machinelearning.DescribeBatchPredictionsInput) error {
    return a.client.WaitUntilBatchPredictionAvailable(input)
}

func (a *MachineLearningActivities) WaitUntilDataSourceAvailable(input *machinelearning.DescribeDataSourcesInput) error {
    return a.client.WaitUntilDataSourceAvailable(input)
}

func (a *MachineLearningActivities) WaitUntilEvaluationAvailable(input *machinelearning.DescribeEvaluationsInput) error {
    return a.client.WaitUntilEvaluationAvailable(input)
}

func (a *MachineLearningActivities) WaitUntilMLModelAvailable(input *machinelearning.DescribeMLModelsInput) error {
    return a.client.WaitUntilMLModelAvailable(input)
}
