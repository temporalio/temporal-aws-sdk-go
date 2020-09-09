
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/personalize"
	"github.com/aws/aws-sdk-go/service/personalize/personalizeiface"
)

type PersonalizeActivities struct {
    client personalizeiface.PersonalizeAPI
}

func NewPersonalizeActivities(session *session.Session, config ...*aws.Config) *PersonalizeActivities {
    client := personalize.New(session, config...)
    return &PersonalizeActivities{client: client}
}

func (a *PersonalizeActivities) CreateBatchInferenceJob(input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error) {
    return a.client.CreateBatchInferenceJob(input)
}

func (a *PersonalizeActivities) CreateCampaign(input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error) {
    return a.client.CreateCampaign(input)
}

func (a *PersonalizeActivities) CreateDataset(input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error) {
    return a.client.CreateDataset(input)
}

func (a *PersonalizeActivities) CreateDatasetGroup(input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error) {
    return a.client.CreateDatasetGroup(input)
}

func (a *PersonalizeActivities) CreateDatasetImportJob(input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error) {
    return a.client.CreateDatasetImportJob(input)
}

func (a *PersonalizeActivities) CreateEventTracker(input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error) {
    return a.client.CreateEventTracker(input)
}

func (a *PersonalizeActivities) CreateFilter(input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error) {
    return a.client.CreateFilter(input)
}

func (a *PersonalizeActivities) CreateSchema(input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error) {
    return a.client.CreateSchema(input)
}

func (a *PersonalizeActivities) CreateSolution(input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error) {
    return a.client.CreateSolution(input)
}

func (a *PersonalizeActivities) CreateSolutionVersion(input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error) {
    return a.client.CreateSolutionVersion(input)
}

func (a *PersonalizeActivities) DeleteCampaign(input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error) {
    return a.client.DeleteCampaign(input)
}

func (a *PersonalizeActivities) DeleteDataset(input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error) {
    return a.client.DeleteDataset(input)
}

func (a *PersonalizeActivities) DeleteDatasetGroup(input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error) {
    return a.client.DeleteDatasetGroup(input)
}

func (a *PersonalizeActivities) DeleteEventTracker(input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error) {
    return a.client.DeleteEventTracker(input)
}

func (a *PersonalizeActivities) DeleteFilter(input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error) {
    return a.client.DeleteFilter(input)
}

func (a *PersonalizeActivities) DeleteSchema(input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error) {
    return a.client.DeleteSchema(input)
}

func (a *PersonalizeActivities) DeleteSolution(input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error) {
    return a.client.DeleteSolution(input)
}

func (a *PersonalizeActivities) DescribeAlgorithm(input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error) {
    return a.client.DescribeAlgorithm(input)
}

func (a *PersonalizeActivities) DescribeBatchInferenceJob(input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error) {
    return a.client.DescribeBatchInferenceJob(input)
}

func (a *PersonalizeActivities) DescribeCampaign(input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error) {
    return a.client.DescribeCampaign(input)
}

func (a *PersonalizeActivities) DescribeDataset(input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error) {
    return a.client.DescribeDataset(input)
}

func (a *PersonalizeActivities) DescribeDatasetGroup(input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error) {
    return a.client.DescribeDatasetGroup(input)
}

func (a *PersonalizeActivities) DescribeDatasetImportJob(input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error) {
    return a.client.DescribeDatasetImportJob(input)
}

func (a *PersonalizeActivities) DescribeEventTracker(input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error) {
    return a.client.DescribeEventTracker(input)
}

func (a *PersonalizeActivities) DescribeFeatureTransformation(input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error) {
    return a.client.DescribeFeatureTransformation(input)
}

func (a *PersonalizeActivities) DescribeFilter(input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error) {
    return a.client.DescribeFilter(input)
}

func (a *PersonalizeActivities) DescribeRecipe(input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error) {
    return a.client.DescribeRecipe(input)
}

func (a *PersonalizeActivities) DescribeSchema(input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error) {
    return a.client.DescribeSchema(input)
}

func (a *PersonalizeActivities) DescribeSolution(input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error) {
    return a.client.DescribeSolution(input)
}

func (a *PersonalizeActivities) DescribeSolutionVersion(input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error) {
    return a.client.DescribeSolutionVersion(input)
}

func (a *PersonalizeActivities) GetSolutionMetrics(input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error) {
    return a.client.GetSolutionMetrics(input)
}

func (a *PersonalizeActivities) ListBatchInferenceJobs(input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error) {
    return a.client.ListBatchInferenceJobs(input)
}

func (a *PersonalizeActivities) ListCampaigns(input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error) {
    return a.client.ListCampaigns(input)
}

func (a *PersonalizeActivities) ListDatasetGroups(input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error) {
    return a.client.ListDatasetGroups(input)
}

func (a *PersonalizeActivities) ListDatasetImportJobs(input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error) {
    return a.client.ListDatasetImportJobs(input)
}

func (a *PersonalizeActivities) ListDatasets(input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error) {
    return a.client.ListDatasets(input)
}

func (a *PersonalizeActivities) ListEventTrackers(input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error) {
    return a.client.ListEventTrackers(input)
}

func (a *PersonalizeActivities) ListFilters(input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error) {
    return a.client.ListFilters(input)
}

func (a *PersonalizeActivities) ListRecipes(input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error) {
    return a.client.ListRecipes(input)
}

func (a *PersonalizeActivities) ListSchemas(input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error) {
    return a.client.ListSchemas(input)
}

func (a *PersonalizeActivities) ListSolutionVersions(input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error) {
    return a.client.ListSolutionVersions(input)
}

func (a *PersonalizeActivities) ListSolutions(input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error) {
    return a.client.ListSolutions(input)
}

func (a *PersonalizeActivities) UpdateCampaign(input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error) {
    return a.client.UpdateCampaign(input)
}
