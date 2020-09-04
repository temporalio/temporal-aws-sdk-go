package awsclients

import (
	"github.com/aws/aws-sdk-go/service/personalize"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type PersonalizeClient interface {
    CreateBatchInferenceJob(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error)
    CreateBatchInferenceJobAsync(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) *PersonalizeCreateBatchInferenceJobResult

    CreateCampaign(ctx workflow.Context, input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error)
    CreateCampaignAsync(ctx workflow.Context, input *personalize.CreateCampaignInput) *PersonalizeCreateCampaignResult

    CreateDataset(ctx workflow.Context, input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error)
    CreateDatasetAsync(ctx workflow.Context, input *personalize.CreateDatasetInput) *PersonalizeCreateDatasetResult

    CreateDatasetGroup(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error)
    CreateDatasetGroupAsync(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) *PersonalizeCreateDatasetGroupResult

    CreateDatasetImportJob(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error)
    CreateDatasetImportJobAsync(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) *PersonalizeCreateDatasetImportJobResult

    CreateEventTracker(ctx workflow.Context, input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error)
    CreateEventTrackerAsync(ctx workflow.Context, input *personalize.CreateEventTrackerInput) *PersonalizeCreateEventTrackerResult

    CreateFilter(ctx workflow.Context, input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error)
    CreateFilterAsync(ctx workflow.Context, input *personalize.CreateFilterInput) *PersonalizeCreateFilterResult

    CreateSchema(ctx workflow.Context, input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error)
    CreateSchemaAsync(ctx workflow.Context, input *personalize.CreateSchemaInput) *PersonalizeCreateSchemaResult

    CreateSolution(ctx workflow.Context, input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error)
    CreateSolutionAsync(ctx workflow.Context, input *personalize.CreateSolutionInput) *PersonalizeCreateSolutionResult

    CreateSolutionVersion(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error)
    CreateSolutionVersionAsync(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) *PersonalizeCreateSolutionVersionResult

    DeleteCampaign(ctx workflow.Context, input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error)
    DeleteCampaignAsync(ctx workflow.Context, input *personalize.DeleteCampaignInput) *PersonalizeDeleteCampaignResult

    DeleteDataset(ctx workflow.Context, input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error)
    DeleteDatasetAsync(ctx workflow.Context, input *personalize.DeleteDatasetInput) *PersonalizeDeleteDatasetResult

    DeleteDatasetGroup(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error)
    DeleteDatasetGroupAsync(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) *PersonalizeDeleteDatasetGroupResult

    DeleteEventTracker(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error)
    DeleteEventTrackerAsync(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) *PersonalizeDeleteEventTrackerResult

    DeleteFilter(ctx workflow.Context, input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error)
    DeleteFilterAsync(ctx workflow.Context, input *personalize.DeleteFilterInput) *PersonalizeDeleteFilterResult

    DeleteSchema(ctx workflow.Context, input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error)
    DeleteSchemaAsync(ctx workflow.Context, input *personalize.DeleteSchemaInput) *PersonalizeDeleteSchemaResult

    DeleteSolution(ctx workflow.Context, input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error)
    DeleteSolutionAsync(ctx workflow.Context, input *personalize.DeleteSolutionInput) *PersonalizeDeleteSolutionResult

    DescribeAlgorithm(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error)
    DescribeAlgorithmAsync(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) *PersonalizeDescribeAlgorithmResult

    DescribeBatchInferenceJob(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error)
    DescribeBatchInferenceJobAsync(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) *PersonalizeDescribeBatchInferenceJobResult

    DescribeCampaign(ctx workflow.Context, input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error)
    DescribeCampaignAsync(ctx workflow.Context, input *personalize.DescribeCampaignInput) *PersonalizeDescribeCampaignResult

    DescribeDataset(ctx workflow.Context, input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error)
    DescribeDatasetAsync(ctx workflow.Context, input *personalize.DescribeDatasetInput) *PersonalizeDescribeDatasetResult

    DescribeDatasetGroup(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error)
    DescribeDatasetGroupAsync(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) *PersonalizeDescribeDatasetGroupResult

    DescribeDatasetImportJob(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error)
    DescribeDatasetImportJobAsync(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) *PersonalizeDescribeDatasetImportJobResult

    DescribeEventTracker(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error)
    DescribeEventTrackerAsync(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) *PersonalizeDescribeEventTrackerResult

    DescribeFeatureTransformation(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error)
    DescribeFeatureTransformationAsync(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) *PersonalizeDescribeFeatureTransformationResult

    DescribeFilter(ctx workflow.Context, input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error)
    DescribeFilterAsync(ctx workflow.Context, input *personalize.DescribeFilterInput) *PersonalizeDescribeFilterResult

    DescribeRecipe(ctx workflow.Context, input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error)
    DescribeRecipeAsync(ctx workflow.Context, input *personalize.DescribeRecipeInput) *PersonalizeDescribeRecipeResult

    DescribeSchema(ctx workflow.Context, input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error)
    DescribeSchemaAsync(ctx workflow.Context, input *personalize.DescribeSchemaInput) *PersonalizeDescribeSchemaResult

    DescribeSolution(ctx workflow.Context, input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error)
    DescribeSolutionAsync(ctx workflow.Context, input *personalize.DescribeSolutionInput) *PersonalizeDescribeSolutionResult

    DescribeSolutionVersion(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error)
    DescribeSolutionVersionAsync(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) *PersonalizeDescribeSolutionVersionResult

    GetSolutionMetrics(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error)
    GetSolutionMetricsAsync(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) *PersonalizeGetSolutionMetricsResult

    ListBatchInferenceJobs(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error)
    ListBatchInferenceJobsAsync(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) *PersonalizeListBatchInferenceJobsResult

    ListCampaigns(ctx workflow.Context, input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error)
    ListCampaignsAsync(ctx workflow.Context, input *personalize.ListCampaignsInput) *PersonalizeListCampaignsResult

    ListDatasetGroups(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error)
    ListDatasetGroupsAsync(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) *PersonalizeListDatasetGroupsResult

    ListDatasetImportJobs(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error)
    ListDatasetImportJobsAsync(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) *PersonalizeListDatasetImportJobsResult

    ListDatasets(ctx workflow.Context, input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error)
    ListDatasetsAsync(ctx workflow.Context, input *personalize.ListDatasetsInput) *PersonalizeListDatasetsResult

    ListEventTrackers(ctx workflow.Context, input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error)
    ListEventTrackersAsync(ctx workflow.Context, input *personalize.ListEventTrackersInput) *PersonalizeListEventTrackersResult

    ListFilters(ctx workflow.Context, input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error)
    ListFiltersAsync(ctx workflow.Context, input *personalize.ListFiltersInput) *PersonalizeListFiltersResult

    ListRecipes(ctx workflow.Context, input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error)
    ListRecipesAsync(ctx workflow.Context, input *personalize.ListRecipesInput) *PersonalizeListRecipesResult

    ListSchemas(ctx workflow.Context, input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error)
    ListSchemasAsync(ctx workflow.Context, input *personalize.ListSchemasInput) *PersonalizeListSchemasResult

    ListSolutionVersions(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error)
    ListSolutionVersionsAsync(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) *PersonalizeListSolutionVersionsResult

    ListSolutions(ctx workflow.Context, input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error)
    ListSolutionsAsync(ctx workflow.Context, input *personalize.ListSolutionsInput) *PersonalizeListSolutionsResult

    UpdateCampaign(ctx workflow.Context, input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error)
    UpdateCampaignAsync(ctx workflow.Context, input *personalize.UpdateCampaignInput) *PersonalizeUpdateCampaignResult
}
type PersonalizeCreateBatchInferenceJobResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateBatchInferenceJobResult) Get(ctx workflow.Context) (*personalize.CreateBatchInferenceJobOutput, error) {
    var output personalize.CreateBatchInferenceJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateCampaignResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateCampaignResult) Get(ctx workflow.Context) (*personalize.CreateCampaignOutput, error) {
    var output personalize.CreateCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateDatasetResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateDatasetResult) Get(ctx workflow.Context) (*personalize.CreateDatasetOutput, error) {
    var output personalize.CreateDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateDatasetGroupResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateDatasetGroupResult) Get(ctx workflow.Context) (*personalize.CreateDatasetGroupOutput, error) {
    var output personalize.CreateDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateDatasetImportJobResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateDatasetImportJobResult) Get(ctx workflow.Context) (*personalize.CreateDatasetImportJobOutput, error) {
    var output personalize.CreateDatasetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateEventTrackerResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateEventTrackerResult) Get(ctx workflow.Context) (*personalize.CreateEventTrackerOutput, error) {
    var output personalize.CreateEventTrackerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateFilterResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateFilterResult) Get(ctx workflow.Context) (*personalize.CreateFilterOutput, error) {
    var output personalize.CreateFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateSchemaResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateSchemaResult) Get(ctx workflow.Context) (*personalize.CreateSchemaOutput, error) {
    var output personalize.CreateSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateSolutionResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateSolutionResult) Get(ctx workflow.Context) (*personalize.CreateSolutionOutput, error) {
    var output personalize.CreateSolutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeCreateSolutionVersionResult struct {
	Result workflow.Future
}

func (r *PersonalizeCreateSolutionVersionResult) Get(ctx workflow.Context) (*personalize.CreateSolutionVersionOutput, error) {
    var output personalize.CreateSolutionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteCampaignResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteCampaignResult) Get(ctx workflow.Context) (*personalize.DeleteCampaignOutput, error) {
    var output personalize.DeleteCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteDatasetResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteDatasetResult) Get(ctx workflow.Context) (*personalize.DeleteDatasetOutput, error) {
    var output personalize.DeleteDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteDatasetGroupResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteDatasetGroupResult) Get(ctx workflow.Context) (*personalize.DeleteDatasetGroupOutput, error) {
    var output personalize.DeleteDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteEventTrackerResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteEventTrackerResult) Get(ctx workflow.Context) (*personalize.DeleteEventTrackerOutput, error) {
    var output personalize.DeleteEventTrackerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteFilterResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteFilterResult) Get(ctx workflow.Context) (*personalize.DeleteFilterOutput, error) {
    var output personalize.DeleteFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteSchemaResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteSchemaResult) Get(ctx workflow.Context) (*personalize.DeleteSchemaOutput, error) {
    var output personalize.DeleteSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDeleteSolutionResult struct {
	Result workflow.Future
}

func (r *PersonalizeDeleteSolutionResult) Get(ctx workflow.Context) (*personalize.DeleteSolutionOutput, error) {
    var output personalize.DeleteSolutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeAlgorithmResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeAlgorithmResult) Get(ctx workflow.Context) (*personalize.DescribeAlgorithmOutput, error) {
    var output personalize.DescribeAlgorithmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeBatchInferenceJobResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeBatchInferenceJobResult) Get(ctx workflow.Context) (*personalize.DescribeBatchInferenceJobOutput, error) {
    var output personalize.DescribeBatchInferenceJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeCampaignResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeCampaignResult) Get(ctx workflow.Context) (*personalize.DescribeCampaignOutput, error) {
    var output personalize.DescribeCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeDatasetResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeDatasetResult) Get(ctx workflow.Context) (*personalize.DescribeDatasetOutput, error) {
    var output personalize.DescribeDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeDatasetGroupResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeDatasetGroupResult) Get(ctx workflow.Context) (*personalize.DescribeDatasetGroupOutput, error) {
    var output personalize.DescribeDatasetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeDatasetImportJobResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeDatasetImportJobResult) Get(ctx workflow.Context) (*personalize.DescribeDatasetImportJobOutput, error) {
    var output personalize.DescribeDatasetImportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeEventTrackerResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeEventTrackerResult) Get(ctx workflow.Context) (*personalize.DescribeEventTrackerOutput, error) {
    var output personalize.DescribeEventTrackerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeFeatureTransformationResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeFeatureTransformationResult) Get(ctx workflow.Context) (*personalize.DescribeFeatureTransformationOutput, error) {
    var output personalize.DescribeFeatureTransformationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeFilterResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeFilterResult) Get(ctx workflow.Context) (*personalize.DescribeFilterOutput, error) {
    var output personalize.DescribeFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeRecipeResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeRecipeResult) Get(ctx workflow.Context) (*personalize.DescribeRecipeOutput, error) {
    var output personalize.DescribeRecipeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeSchemaResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeSchemaResult) Get(ctx workflow.Context) (*personalize.DescribeSchemaOutput, error) {
    var output personalize.DescribeSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeSolutionResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeSolutionResult) Get(ctx workflow.Context) (*personalize.DescribeSolutionOutput, error) {
    var output personalize.DescribeSolutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeDescribeSolutionVersionResult struct {
	Result workflow.Future
}

func (r *PersonalizeDescribeSolutionVersionResult) Get(ctx workflow.Context) (*personalize.DescribeSolutionVersionOutput, error) {
    var output personalize.DescribeSolutionVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeGetSolutionMetricsResult struct {
	Result workflow.Future
}

func (r *PersonalizeGetSolutionMetricsResult) Get(ctx workflow.Context) (*personalize.GetSolutionMetricsOutput, error) {
    var output personalize.GetSolutionMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListBatchInferenceJobsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListBatchInferenceJobsResult) Get(ctx workflow.Context) (*personalize.ListBatchInferenceJobsOutput, error) {
    var output personalize.ListBatchInferenceJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListCampaignsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListCampaignsResult) Get(ctx workflow.Context) (*personalize.ListCampaignsOutput, error) {
    var output personalize.ListCampaignsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListDatasetGroupsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListDatasetGroupsResult) Get(ctx workflow.Context) (*personalize.ListDatasetGroupsOutput, error) {
    var output personalize.ListDatasetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListDatasetImportJobsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListDatasetImportJobsResult) Get(ctx workflow.Context) (*personalize.ListDatasetImportJobsOutput, error) {
    var output personalize.ListDatasetImportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListDatasetsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListDatasetsResult) Get(ctx workflow.Context) (*personalize.ListDatasetsOutput, error) {
    var output personalize.ListDatasetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListEventTrackersResult struct {
	Result workflow.Future
}

func (r *PersonalizeListEventTrackersResult) Get(ctx workflow.Context) (*personalize.ListEventTrackersOutput, error) {
    var output personalize.ListEventTrackersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListFiltersResult struct {
	Result workflow.Future
}

func (r *PersonalizeListFiltersResult) Get(ctx workflow.Context) (*personalize.ListFiltersOutput, error) {
    var output personalize.ListFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListRecipesResult struct {
	Result workflow.Future
}

func (r *PersonalizeListRecipesResult) Get(ctx workflow.Context) (*personalize.ListRecipesOutput, error) {
    var output personalize.ListRecipesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListSchemasResult struct {
	Result workflow.Future
}

func (r *PersonalizeListSchemasResult) Get(ctx workflow.Context) (*personalize.ListSchemasOutput, error) {
    var output personalize.ListSchemasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListSolutionVersionsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListSolutionVersionsResult) Get(ctx workflow.Context) (*personalize.ListSolutionVersionsOutput, error) {
    var output personalize.ListSolutionVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeListSolutionsResult struct {
	Result workflow.Future
}

func (r *PersonalizeListSolutionsResult) Get(ctx workflow.Context) (*personalize.ListSolutionsOutput, error) {
    var output personalize.ListSolutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PersonalizeUpdateCampaignResult struct {
	Result workflow.Future
}

func (r *PersonalizeUpdateCampaignResult) Get(ctx workflow.Context) (*personalize.UpdateCampaignOutput, error) {
    var output personalize.UpdateCampaignOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type PersonalizeStub struct {
    activities awsactivities.PersonalizeActivities
}

func NewPersonalizeStub() PersonalizeClient {
    return &PersonalizeStub{}
}
func (a *PersonalizeStub) CreateBatchInferenceJob(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) (*personalize.CreateBatchInferenceJobOutput, error) {
    var output personalize.CreateBatchInferenceJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBatchInferenceJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateBatchInferenceJobAsync(ctx workflow.Context, input *personalize.CreateBatchInferenceJobInput) *PersonalizeCreateBatchInferenceJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBatchInferenceJob, input)
    return &PersonalizeCreateBatchInferenceJobResult{Result: future}
}
func (a *PersonalizeStub) CreateCampaign(ctx workflow.Context, input *personalize.CreateCampaignInput) (*personalize.CreateCampaignOutput, error) {
    var output personalize.CreateCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateCampaignAsync(ctx workflow.Context, input *personalize.CreateCampaignInput) *PersonalizeCreateCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCampaign, input)
    return &PersonalizeCreateCampaignResult{Result: future}
}
func (a *PersonalizeStub) CreateDataset(ctx workflow.Context, input *personalize.CreateDatasetInput) (*personalize.CreateDatasetOutput, error) {
    var output personalize.CreateDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateDatasetAsync(ctx workflow.Context, input *personalize.CreateDatasetInput) *PersonalizeCreateDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataset, input)
    return &PersonalizeCreateDatasetResult{Result: future}
}
func (a *PersonalizeStub) CreateDatasetGroup(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) (*personalize.CreateDatasetGroupOutput, error) {
    var output personalize.CreateDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateDatasetGroupAsync(ctx workflow.Context, input *personalize.CreateDatasetGroupInput) *PersonalizeCreateDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetGroup, input)
    return &PersonalizeCreateDatasetGroupResult{Result: future}
}
func (a *PersonalizeStub) CreateDatasetImportJob(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) (*personalize.CreateDatasetImportJobOutput, error) {
    var output personalize.CreateDatasetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateDatasetImportJobAsync(ctx workflow.Context, input *personalize.CreateDatasetImportJobInput) *PersonalizeCreateDatasetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatasetImportJob, input)
    return &PersonalizeCreateDatasetImportJobResult{Result: future}
}
func (a *PersonalizeStub) CreateEventTracker(ctx workflow.Context, input *personalize.CreateEventTrackerInput) (*personalize.CreateEventTrackerOutput, error) {
    var output personalize.CreateEventTrackerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEventTracker, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateEventTrackerAsync(ctx workflow.Context, input *personalize.CreateEventTrackerInput) *PersonalizeCreateEventTrackerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEventTracker, input)
    return &PersonalizeCreateEventTrackerResult{Result: future}
}
func (a *PersonalizeStub) CreateFilter(ctx workflow.Context, input *personalize.CreateFilterInput) (*personalize.CreateFilterOutput, error) {
    var output personalize.CreateFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateFilterAsync(ctx workflow.Context, input *personalize.CreateFilterInput) *PersonalizeCreateFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFilter, input)
    return &PersonalizeCreateFilterResult{Result: future}
}
func (a *PersonalizeStub) CreateSchema(ctx workflow.Context, input *personalize.CreateSchemaInput) (*personalize.CreateSchemaOutput, error) {
    var output personalize.CreateSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateSchemaAsync(ctx workflow.Context, input *personalize.CreateSchemaInput) *PersonalizeCreateSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSchema, input)
    return &PersonalizeCreateSchemaResult{Result: future}
}
func (a *PersonalizeStub) CreateSolution(ctx workflow.Context, input *personalize.CreateSolutionInput) (*personalize.CreateSolutionOutput, error) {
    var output personalize.CreateSolutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSolution, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateSolutionAsync(ctx workflow.Context, input *personalize.CreateSolutionInput) *PersonalizeCreateSolutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSolution, input)
    return &PersonalizeCreateSolutionResult{Result: future}
}
func (a *PersonalizeStub) CreateSolutionVersion(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) (*personalize.CreateSolutionVersionOutput, error) {
    var output personalize.CreateSolutionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSolutionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) CreateSolutionVersionAsync(ctx workflow.Context, input *personalize.CreateSolutionVersionInput) *PersonalizeCreateSolutionVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSolutionVersion, input)
    return &PersonalizeCreateSolutionVersionResult{Result: future}
}
func (a *PersonalizeStub) DeleteCampaign(ctx workflow.Context, input *personalize.DeleteCampaignInput) (*personalize.DeleteCampaignOutput, error) {
    var output personalize.DeleteCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteCampaignAsync(ctx workflow.Context, input *personalize.DeleteCampaignInput) *PersonalizeDeleteCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCampaign, input)
    return &PersonalizeDeleteCampaignResult{Result: future}
}
func (a *PersonalizeStub) DeleteDataset(ctx workflow.Context, input *personalize.DeleteDatasetInput) (*personalize.DeleteDatasetOutput, error) {
    var output personalize.DeleteDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteDatasetAsync(ctx workflow.Context, input *personalize.DeleteDatasetInput) *PersonalizeDeleteDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input)
    return &PersonalizeDeleteDatasetResult{Result: future}
}
func (a *PersonalizeStub) DeleteDatasetGroup(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) (*personalize.DeleteDatasetGroupOutput, error) {
    var output personalize.DeleteDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteDatasetGroupAsync(ctx workflow.Context, input *personalize.DeleteDatasetGroupInput) *PersonalizeDeleteDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDatasetGroup, input)
    return &PersonalizeDeleteDatasetGroupResult{Result: future}
}
func (a *PersonalizeStub) DeleteEventTracker(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) (*personalize.DeleteEventTrackerOutput, error) {
    var output personalize.DeleteEventTrackerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventTracker, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteEventTrackerAsync(ctx workflow.Context, input *personalize.DeleteEventTrackerInput) *PersonalizeDeleteEventTrackerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEventTracker, input)
    return &PersonalizeDeleteEventTrackerResult{Result: future}
}
func (a *PersonalizeStub) DeleteFilter(ctx workflow.Context, input *personalize.DeleteFilterInput) (*personalize.DeleteFilterOutput, error) {
    var output personalize.DeleteFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteFilterAsync(ctx workflow.Context, input *personalize.DeleteFilterInput) *PersonalizeDeleteFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFilter, input)
    return &PersonalizeDeleteFilterResult{Result: future}
}
func (a *PersonalizeStub) DeleteSchema(ctx workflow.Context, input *personalize.DeleteSchemaInput) (*personalize.DeleteSchemaOutput, error) {
    var output personalize.DeleteSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteSchemaAsync(ctx workflow.Context, input *personalize.DeleteSchemaInput) *PersonalizeDeleteSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSchema, input)
    return &PersonalizeDeleteSchemaResult{Result: future}
}
func (a *PersonalizeStub) DeleteSolution(ctx workflow.Context, input *personalize.DeleteSolutionInput) (*personalize.DeleteSolutionOutput, error) {
    var output personalize.DeleteSolutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSolution, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DeleteSolutionAsync(ctx workflow.Context, input *personalize.DeleteSolutionInput) *PersonalizeDeleteSolutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSolution, input)
    return &PersonalizeDeleteSolutionResult{Result: future}
}
func (a *PersonalizeStub) DescribeAlgorithm(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error) {
    var output personalize.DescribeAlgorithmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAlgorithm, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeAlgorithmAsync(ctx workflow.Context, input *personalize.DescribeAlgorithmInput) *PersonalizeDescribeAlgorithmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAlgorithm, input)
    return &PersonalizeDescribeAlgorithmResult{Result: future}
}
func (a *PersonalizeStub) DescribeBatchInferenceJob(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error) {
    var output personalize.DescribeBatchInferenceJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBatchInferenceJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeBatchInferenceJobAsync(ctx workflow.Context, input *personalize.DescribeBatchInferenceJobInput) *PersonalizeDescribeBatchInferenceJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBatchInferenceJob, input)
    return &PersonalizeDescribeBatchInferenceJobResult{Result: future}
}
func (a *PersonalizeStub) DescribeCampaign(ctx workflow.Context, input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error) {
    var output personalize.DescribeCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeCampaignAsync(ctx workflow.Context, input *personalize.DescribeCampaignInput) *PersonalizeDescribeCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCampaign, input)
    return &PersonalizeDescribeCampaignResult{Result: future}
}
func (a *PersonalizeStub) DescribeDataset(ctx workflow.Context, input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error) {
    var output personalize.DescribeDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeDatasetAsync(ctx workflow.Context, input *personalize.DescribeDatasetInput) *PersonalizeDescribeDatasetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input)
    return &PersonalizeDescribeDatasetResult{Result: future}
}
func (a *PersonalizeStub) DescribeDatasetGroup(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error) {
    var output personalize.DescribeDatasetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeDatasetGroupAsync(ctx workflow.Context, input *personalize.DescribeDatasetGroupInput) *PersonalizeDescribeDatasetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetGroup, input)
    return &PersonalizeDescribeDatasetGroupResult{Result: future}
}
func (a *PersonalizeStub) DescribeDatasetImportJob(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error) {
    var output personalize.DescribeDatasetImportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetImportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeDatasetImportJobAsync(ctx workflow.Context, input *personalize.DescribeDatasetImportJobInput) *PersonalizeDescribeDatasetImportJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDatasetImportJob, input)
    return &PersonalizeDescribeDatasetImportJobResult{Result: future}
}
func (a *PersonalizeStub) DescribeEventTracker(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error) {
    var output personalize.DescribeEventTrackerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventTracker, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeEventTrackerAsync(ctx workflow.Context, input *personalize.DescribeEventTrackerInput) *PersonalizeDescribeEventTrackerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventTracker, input)
    return &PersonalizeDescribeEventTrackerResult{Result: future}
}
func (a *PersonalizeStub) DescribeFeatureTransformation(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error) {
    var output personalize.DescribeFeatureTransformationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFeatureTransformation, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeFeatureTransformationAsync(ctx workflow.Context, input *personalize.DescribeFeatureTransformationInput) *PersonalizeDescribeFeatureTransformationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFeatureTransformation, input)
    return &PersonalizeDescribeFeatureTransformationResult{Result: future}
}
func (a *PersonalizeStub) DescribeFilter(ctx workflow.Context, input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error) {
    var output personalize.DescribeFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeFilterAsync(ctx workflow.Context, input *personalize.DescribeFilterInput) *PersonalizeDescribeFilterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFilter, input)
    return &PersonalizeDescribeFilterResult{Result: future}
}
func (a *PersonalizeStub) DescribeRecipe(ctx workflow.Context, input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error) {
    var output personalize.DescribeRecipeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRecipe, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeRecipeAsync(ctx workflow.Context, input *personalize.DescribeRecipeInput) *PersonalizeDescribeRecipeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRecipe, input)
    return &PersonalizeDescribeRecipeResult{Result: future}
}
func (a *PersonalizeStub) DescribeSchema(ctx workflow.Context, input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error) {
    var output personalize.DescribeSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeSchemaAsync(ctx workflow.Context, input *personalize.DescribeSchemaInput) *PersonalizeDescribeSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSchema, input)
    return &PersonalizeDescribeSchemaResult{Result: future}
}
func (a *PersonalizeStub) DescribeSolution(ctx workflow.Context, input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error) {
    var output personalize.DescribeSolutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSolution, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeSolutionAsync(ctx workflow.Context, input *personalize.DescribeSolutionInput) *PersonalizeDescribeSolutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSolution, input)
    return &PersonalizeDescribeSolutionResult{Result: future}
}
func (a *PersonalizeStub) DescribeSolutionVersion(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error) {
    var output personalize.DescribeSolutionVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSolutionVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) DescribeSolutionVersionAsync(ctx workflow.Context, input *personalize.DescribeSolutionVersionInput) *PersonalizeDescribeSolutionVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSolutionVersion, input)
    return &PersonalizeDescribeSolutionVersionResult{Result: future}
}
func (a *PersonalizeStub) GetSolutionMetrics(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error) {
    var output personalize.GetSolutionMetricsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSolutionMetrics, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) GetSolutionMetricsAsync(ctx workflow.Context, input *personalize.GetSolutionMetricsInput) *PersonalizeGetSolutionMetricsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSolutionMetrics, input)
    return &PersonalizeGetSolutionMetricsResult{Result: future}
}
func (a *PersonalizeStub) ListBatchInferenceJobs(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error) {
    var output personalize.ListBatchInferenceJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBatchInferenceJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListBatchInferenceJobsAsync(ctx workflow.Context, input *personalize.ListBatchInferenceJobsInput) *PersonalizeListBatchInferenceJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBatchInferenceJobs, input)
    return &PersonalizeListBatchInferenceJobsResult{Result: future}
}
func (a *PersonalizeStub) ListCampaigns(ctx workflow.Context, input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error) {
    var output personalize.ListCampaignsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCampaigns, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListCampaignsAsync(ctx workflow.Context, input *personalize.ListCampaignsInput) *PersonalizeListCampaignsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCampaigns, input)
    return &PersonalizeListCampaignsResult{Result: future}
}
func (a *PersonalizeStub) ListDatasetGroups(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error) {
    var output personalize.ListDatasetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListDatasetGroupsAsync(ctx workflow.Context, input *personalize.ListDatasetGroupsInput) *PersonalizeListDatasetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasetGroups, input)
    return &PersonalizeListDatasetGroupsResult{Result: future}
}
func (a *PersonalizeStub) ListDatasetImportJobs(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error) {
    var output personalize.ListDatasetImportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasetImportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListDatasetImportJobsAsync(ctx workflow.Context, input *personalize.ListDatasetImportJobsInput) *PersonalizeListDatasetImportJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasetImportJobs, input)
    return &PersonalizeListDatasetImportJobsResult{Result: future}
}
func (a *PersonalizeStub) ListDatasets(ctx workflow.Context, input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error) {
    var output personalize.ListDatasetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListDatasetsAsync(ctx workflow.Context, input *personalize.ListDatasetsInput) *PersonalizeListDatasetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input)
    return &PersonalizeListDatasetsResult{Result: future}
}
func (a *PersonalizeStub) ListEventTrackers(ctx workflow.Context, input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error) {
    var output personalize.ListEventTrackersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEventTrackers, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListEventTrackersAsync(ctx workflow.Context, input *personalize.ListEventTrackersInput) *PersonalizeListEventTrackersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListEventTrackers, input)
    return &PersonalizeListEventTrackersResult{Result: future}
}
func (a *PersonalizeStub) ListFilters(ctx workflow.Context, input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error) {
    var output personalize.ListFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListFiltersAsync(ctx workflow.Context, input *personalize.ListFiltersInput) *PersonalizeListFiltersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFilters, input)
    return &PersonalizeListFiltersResult{Result: future}
}
func (a *PersonalizeStub) ListRecipes(ctx workflow.Context, input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error) {
    var output personalize.ListRecipesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecipes, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListRecipesAsync(ctx workflow.Context, input *personalize.ListRecipesInput) *PersonalizeListRecipesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRecipes, input)
    return &PersonalizeListRecipesResult{Result: future}
}
func (a *PersonalizeStub) ListSchemas(ctx workflow.Context, input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error) {
    var output personalize.ListSchemasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSchemas, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListSchemasAsync(ctx workflow.Context, input *personalize.ListSchemasInput) *PersonalizeListSchemasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSchemas, input)
    return &PersonalizeListSchemasResult{Result: future}
}
func (a *PersonalizeStub) ListSolutionVersions(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error) {
    var output personalize.ListSolutionVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSolutionVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListSolutionVersionsAsync(ctx workflow.Context, input *personalize.ListSolutionVersionsInput) *PersonalizeListSolutionVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSolutionVersions, input)
    return &PersonalizeListSolutionVersionsResult{Result: future}
}
func (a *PersonalizeStub) ListSolutions(ctx workflow.Context, input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error) {
    var output personalize.ListSolutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSolutions, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) ListSolutionsAsync(ctx workflow.Context, input *personalize.ListSolutionsInput) *PersonalizeListSolutionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSolutions, input)
    return &PersonalizeListSolutionsResult{Result: future}
}
func (a *PersonalizeStub) UpdateCampaign(ctx workflow.Context, input *personalize.UpdateCampaignInput) (*personalize.UpdateCampaignOutput, error) {
    var output personalize.UpdateCampaignOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCampaign, input).Get(ctx, &output)
    return &output, err
}

func (a *PersonalizeStub) UpdateCampaignAsync(ctx workflow.Context, input *personalize.UpdateCampaignInput) *PersonalizeUpdateCampaignResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCampaign, input)
    return &PersonalizeUpdateCampaignResult{Result: future}
}
