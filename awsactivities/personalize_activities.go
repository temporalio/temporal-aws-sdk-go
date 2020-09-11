
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/personalize"
	"github.com/aws/aws-sdk-go/service/personalize/personalizeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type PersonalizeActivities struct {
    client personalizeiface.PersonalizeAPI
}

func NewPersonalizeActivities(session *session.Session, config ...*aws.Config) *PersonalizeActivities {
    client := personalize.New(session, config...)
    return &PersonalizeActivities{client: client}
}

func (a *PersonalizeActivities) CreateBatchInferenceJob(ctx context.Context, input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error) {
    return a.client.CreateBatchInferenceJobWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateCampaign(ctx context.Context, input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error) {
    return a.client.CreateCampaignWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateDataset(ctx context.Context, input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error) {
    return a.client.CreateDatasetWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateDatasetGroup(ctx context.Context, input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error) {
    return a.client.CreateDatasetGroupWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateDatasetImportJob(ctx context.Context, input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error) {
    return a.client.CreateDatasetImportJobWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateEventTracker(ctx context.Context, input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error) {
    return a.client.CreateEventTrackerWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateFilter(ctx context.Context, input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error) {
    return a.client.CreateFilterWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateSchema(ctx context.Context, input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error) {
    return a.client.CreateSchemaWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateSolution(ctx context.Context, input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error) {
    return a.client.CreateSolutionWithContext(ctx, input)
}

func (a *PersonalizeActivities) CreateSolutionVersion(ctx context.Context, input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error) {
    return a.client.CreateSolutionVersionWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteCampaign(ctx context.Context, input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error) {
    return a.client.DeleteCampaignWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteDataset(ctx context.Context, input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error) {
    return a.client.DeleteDatasetWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteDatasetGroup(ctx context.Context, input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error) {
    return a.client.DeleteDatasetGroupWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteEventTracker(ctx context.Context, input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error) {
    return a.client.DeleteEventTrackerWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteFilter(ctx context.Context, input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error) {
    return a.client.DeleteFilterWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteSchema(ctx context.Context, input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error) {
    return a.client.DeleteSchemaWithContext(ctx, input)
}

func (a *PersonalizeActivities) DeleteSolution(ctx context.Context, input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error) {
    return a.client.DeleteSolutionWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeAlgorithm(ctx context.Context, input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error) {
    return a.client.DescribeAlgorithmWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeBatchInferenceJob(ctx context.Context, input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error) {
    return a.client.DescribeBatchInferenceJobWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeCampaign(ctx context.Context, input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error) {
    return a.client.DescribeCampaignWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeDataset(ctx context.Context, input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error) {
    return a.client.DescribeDatasetWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeDatasetGroup(ctx context.Context, input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error) {
    return a.client.DescribeDatasetGroupWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeDatasetImportJob(ctx context.Context, input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error) {
    return a.client.DescribeDatasetImportJobWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeEventTracker(ctx context.Context, input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error) {
    return a.client.DescribeEventTrackerWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeFeatureTransformation(ctx context.Context, input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error) {
    return a.client.DescribeFeatureTransformationWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeFilter(ctx context.Context, input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error) {
    return a.client.DescribeFilterWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeRecipe(ctx context.Context, input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error) {
    return a.client.DescribeRecipeWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeSchema(ctx context.Context, input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error) {
    return a.client.DescribeSchemaWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeSolution(ctx context.Context, input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error) {
    return a.client.DescribeSolutionWithContext(ctx, input)
}

func (a *PersonalizeActivities) DescribeSolutionVersion(ctx context.Context, input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error) {
    return a.client.DescribeSolutionVersionWithContext(ctx, input)
}

func (a *PersonalizeActivities) GetSolutionMetrics(ctx context.Context, input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error) {
    return a.client.GetSolutionMetricsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListBatchInferenceJobs(ctx context.Context, input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error) {
    return a.client.ListBatchInferenceJobsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListCampaigns(ctx context.Context, input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error) {
    return a.client.ListCampaignsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListDatasetGroups(ctx context.Context, input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error) {
    return a.client.ListDatasetGroupsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListDatasetImportJobs(ctx context.Context, input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error) {
    return a.client.ListDatasetImportJobsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListDatasets(ctx context.Context, input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error) {
    return a.client.ListDatasetsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListEventTrackers(ctx context.Context, input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error) {
    return a.client.ListEventTrackersWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListFilters(ctx context.Context, input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error) {
    return a.client.ListFiltersWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListRecipes(ctx context.Context, input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error) {
    return a.client.ListRecipesWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListSchemas(ctx context.Context, input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error) {
    return a.client.ListSchemasWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListSolutionVersions(ctx context.Context, input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error) {
    return a.client.ListSolutionVersionsWithContext(ctx, input)
}

func (a *PersonalizeActivities) ListSolutions(ctx context.Context, input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error) {
    return a.client.ListSolutionsWithContext(ctx, input)
}

func (a *PersonalizeActivities) UpdateCampaign(ctx context.Context, input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error) {
    return a.client.UpdateCampaignWithContext(ctx, input)
}
