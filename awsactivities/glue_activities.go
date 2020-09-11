
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/glue/glueiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type GlueActivities struct {
    client glueiface.GlueAPI
}

func NewGlueActivities(session *session.Session, config ...*aws.Config) *GlueActivities {
    client := glue.New(session, config...)
    return &GlueActivities{client: client}
}

func (a *GlueActivities) BatchCreatePartition(ctx context.Context, input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error) {
    return a.client.BatchCreatePartitionWithContext(ctx, input)
}

func (a *GlueActivities) BatchDeleteConnection(ctx context.Context, input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error) {
    return a.client.BatchDeleteConnectionWithContext(ctx, input)
}

func (a *GlueActivities) BatchDeletePartition(ctx context.Context, input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error) {
    return a.client.BatchDeletePartitionWithContext(ctx, input)
}

func (a *GlueActivities) BatchDeleteTable(ctx context.Context, input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error) {
    return a.client.BatchDeleteTableWithContext(ctx, input)
}

func (a *GlueActivities) BatchDeleteTableVersion(ctx context.Context, input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error) {
    return a.client.BatchDeleteTableVersionWithContext(ctx, input)
}

func (a *GlueActivities) BatchGetCrawlers(ctx context.Context, input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error) {
    return a.client.BatchGetCrawlersWithContext(ctx, input)
}

func (a *GlueActivities) BatchGetDevEndpoints(ctx context.Context, input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error) {
    return a.client.BatchGetDevEndpointsWithContext(ctx, input)
}

func (a *GlueActivities) BatchGetJobs(ctx context.Context, input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error) {
    return a.client.BatchGetJobsWithContext(ctx, input)
}

func (a *GlueActivities) BatchGetPartition(ctx context.Context, input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error) {
    return a.client.BatchGetPartitionWithContext(ctx, input)
}

func (a *GlueActivities) BatchGetTriggers(ctx context.Context, input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error) {
    return a.client.BatchGetTriggersWithContext(ctx, input)
}

func (a *GlueActivities) BatchGetWorkflows(ctx context.Context, input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error) {
    return a.client.BatchGetWorkflowsWithContext(ctx, input)
}

func (a *GlueActivities) BatchStopJobRun(ctx context.Context, input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error) {
    return a.client.BatchStopJobRunWithContext(ctx, input)
}

func (a *GlueActivities) CancelMLTaskRun(ctx context.Context, input *glue.CancelMLTaskRunInput) (*glue.CancelMLTaskRunOutput, error) {
    return a.client.CancelMLTaskRunWithContext(ctx, input)
}

func (a *GlueActivities) CreateClassifier(ctx context.Context, input *glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error) {
    return a.client.CreateClassifierWithContext(ctx, input)
}

func (a *GlueActivities) CreateConnection(ctx context.Context, input *glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error) {
    return a.client.CreateConnectionWithContext(ctx, input)
}

func (a *GlueActivities) CreateCrawler(ctx context.Context, input *glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error) {
    return a.client.CreateCrawlerWithContext(ctx, input)
}

func (a *GlueActivities) CreateDatabase(ctx context.Context, input *glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error) {
    return a.client.CreateDatabaseWithContext(ctx, input)
}

func (a *GlueActivities) CreateDevEndpoint(ctx context.Context, input *glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error) {
    return a.client.CreateDevEndpointWithContext(ctx, input)
}

func (a *GlueActivities) CreateJob(ctx context.Context, input *glue.CreateJobInput) (*glue.CreateJobOutput, error) {
    return a.client.CreateJobWithContext(ctx, input)
}

func (a *GlueActivities) CreateMLTransform(ctx context.Context, input *glue.CreateMLTransformInput) (*glue.CreateMLTransformOutput, error) {
    return a.client.CreateMLTransformWithContext(ctx, input)
}

func (a *GlueActivities) CreatePartition(ctx context.Context, input *glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error) {
    return a.client.CreatePartitionWithContext(ctx, input)
}

func (a *GlueActivities) CreateScript(ctx context.Context, input *glue.CreateScriptInput) (*glue.CreateScriptOutput, error) {
    return a.client.CreateScriptWithContext(ctx, input)
}

func (a *GlueActivities) CreateSecurityConfiguration(ctx context.Context, input *glue.CreateSecurityConfigurationInput) (*glue.CreateSecurityConfigurationOutput, error) {
    return a.client.CreateSecurityConfigurationWithContext(ctx, input)
}

func (a *GlueActivities) CreateTable(ctx context.Context, input *glue.CreateTableInput) (*glue.CreateTableOutput, error) {
    return a.client.CreateTableWithContext(ctx, input)
}

func (a *GlueActivities) CreateTrigger(ctx context.Context, input *glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error) {
    return a.client.CreateTriggerWithContext(ctx, input)
}

func (a *GlueActivities) CreateUserDefinedFunction(ctx context.Context, input *glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error) {
    return a.client.CreateUserDefinedFunctionWithContext(ctx, input)
}

func (a *GlueActivities) CreateWorkflow(ctx context.Context, input *glue.CreateWorkflowInput) (*glue.CreateWorkflowOutput, error) {
    return a.client.CreateWorkflowWithContext(ctx, input)
}

func (a *GlueActivities) DeleteClassifier(ctx context.Context, input *glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error) {
    return a.client.DeleteClassifierWithContext(ctx, input)
}

func (a *GlueActivities) DeleteColumnStatisticsForPartition(ctx context.Context, input *glue.DeleteColumnStatisticsForPartitionInput) (*glue.DeleteColumnStatisticsForPartitionOutput, error) {
    return a.client.DeleteColumnStatisticsForPartitionWithContext(ctx, input)
}

func (a *GlueActivities) DeleteColumnStatisticsForTable(ctx context.Context, input *glue.DeleteColumnStatisticsForTableInput) (*glue.DeleteColumnStatisticsForTableOutput, error) {
    return a.client.DeleteColumnStatisticsForTableWithContext(ctx, input)
}

func (a *GlueActivities) DeleteConnection(ctx context.Context, input *glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error) {
    return a.client.DeleteConnectionWithContext(ctx, input)
}

func (a *GlueActivities) DeleteCrawler(ctx context.Context, input *glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error) {
    return a.client.DeleteCrawlerWithContext(ctx, input)
}

func (a *GlueActivities) DeleteDatabase(ctx context.Context, input *glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error) {
    return a.client.DeleteDatabaseWithContext(ctx, input)
}

func (a *GlueActivities) DeleteDevEndpoint(ctx context.Context, input *glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error) {
    return a.client.DeleteDevEndpointWithContext(ctx, input)
}

func (a *GlueActivities) DeleteJob(ctx context.Context, input *glue.DeleteJobInput) (*glue.DeleteJobOutput, error) {
    return a.client.DeleteJobWithContext(ctx, input)
}

func (a *GlueActivities) DeleteMLTransform(ctx context.Context, input *glue.DeleteMLTransformInput) (*glue.DeleteMLTransformOutput, error) {
    return a.client.DeleteMLTransformWithContext(ctx, input)
}

func (a *GlueActivities) DeletePartition(ctx context.Context, input *glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error) {
    return a.client.DeletePartitionWithContext(ctx, input)
}

func (a *GlueActivities) DeleteResourcePolicy(ctx context.Context, input *glue.DeleteResourcePolicyInput) (*glue.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *GlueActivities) DeleteSecurityConfiguration(ctx context.Context, input *glue.DeleteSecurityConfigurationInput) (*glue.DeleteSecurityConfigurationOutput, error) {
    return a.client.DeleteSecurityConfigurationWithContext(ctx, input)
}

func (a *GlueActivities) DeleteTable(ctx context.Context, input *glue.DeleteTableInput) (*glue.DeleteTableOutput, error) {
    return a.client.DeleteTableWithContext(ctx, input)
}

func (a *GlueActivities) DeleteTableVersion(ctx context.Context, input *glue.DeleteTableVersionInput) (*glue.DeleteTableVersionOutput, error) {
    return a.client.DeleteTableVersionWithContext(ctx, input)
}

func (a *GlueActivities) DeleteTrigger(ctx context.Context, input *glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error) {
    return a.client.DeleteTriggerWithContext(ctx, input)
}

func (a *GlueActivities) DeleteUserDefinedFunction(ctx context.Context, input *glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error) {
    return a.client.DeleteUserDefinedFunctionWithContext(ctx, input)
}

func (a *GlueActivities) DeleteWorkflow(ctx context.Context, input *glue.DeleteWorkflowInput) (*glue.DeleteWorkflowOutput, error) {
    return a.client.DeleteWorkflowWithContext(ctx, input)
}

func (a *GlueActivities) GetCatalogImportStatus(ctx context.Context, input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error) {
    return a.client.GetCatalogImportStatusWithContext(ctx, input)
}

func (a *GlueActivities) GetClassifier(ctx context.Context, input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error) {
    return a.client.GetClassifierWithContext(ctx, input)
}

func (a *GlueActivities) GetClassifiers(ctx context.Context, input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error) {
    return a.client.GetClassifiersWithContext(ctx, input)
}

func (a *GlueActivities) GetColumnStatisticsForPartition(ctx context.Context, input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error) {
    return a.client.GetColumnStatisticsForPartitionWithContext(ctx, input)
}

func (a *GlueActivities) GetColumnStatisticsForTable(ctx context.Context, input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error) {
    return a.client.GetColumnStatisticsForTableWithContext(ctx, input)
}

func (a *GlueActivities) GetConnection(ctx context.Context, input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error) {
    return a.client.GetConnectionWithContext(ctx, input)
}

func (a *GlueActivities) GetConnections(ctx context.Context, input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error) {
    return a.client.GetConnectionsWithContext(ctx, input)
}

func (a *GlueActivities) GetCrawler(ctx context.Context, input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error) {
    return a.client.GetCrawlerWithContext(ctx, input)
}

func (a *GlueActivities) GetCrawlerMetrics(ctx context.Context, input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error) {
    return a.client.GetCrawlerMetricsWithContext(ctx, input)
}

func (a *GlueActivities) GetCrawlers(ctx context.Context, input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error) {
    return a.client.GetCrawlersWithContext(ctx, input)
}

func (a *GlueActivities) GetDataCatalogEncryptionSettings(ctx context.Context, input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
    return a.client.GetDataCatalogEncryptionSettingsWithContext(ctx, input)
}

func (a *GlueActivities) GetDatabase(ctx context.Context, input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error) {
    return a.client.GetDatabaseWithContext(ctx, input)
}

func (a *GlueActivities) GetDatabases(ctx context.Context, input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error) {
    return a.client.GetDatabasesWithContext(ctx, input)
}

func (a *GlueActivities) GetDataflowGraph(ctx context.Context, input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error) {
    return a.client.GetDataflowGraphWithContext(ctx, input)
}

func (a *GlueActivities) GetDevEndpoint(ctx context.Context, input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error) {
    return a.client.GetDevEndpointWithContext(ctx, input)
}

func (a *GlueActivities) GetDevEndpoints(ctx context.Context, input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error) {
    return a.client.GetDevEndpointsWithContext(ctx, input)
}

func (a *GlueActivities) GetJob(ctx context.Context, input *glue.GetJobInput) (*glue.GetJobOutput, error) {
    return a.client.GetJobWithContext(ctx, input)
}

func (a *GlueActivities) GetJobBookmark(ctx context.Context, input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error) {
    return a.client.GetJobBookmarkWithContext(ctx, input)
}

func (a *GlueActivities) GetJobRun(ctx context.Context, input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error) {
    return a.client.GetJobRunWithContext(ctx, input)
}

func (a *GlueActivities) GetJobRuns(ctx context.Context, input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error) {
    return a.client.GetJobRunsWithContext(ctx, input)
}

func (a *GlueActivities) GetJobs(ctx context.Context, input *glue.GetJobsInput) (*glue.GetJobsOutput, error) {
    return a.client.GetJobsWithContext(ctx, input)
}

func (a *GlueActivities) GetMLTaskRun(ctx context.Context, input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error) {
    return a.client.GetMLTaskRunWithContext(ctx, input)
}

func (a *GlueActivities) GetMLTaskRuns(ctx context.Context, input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error) {
    return a.client.GetMLTaskRunsWithContext(ctx, input)
}

func (a *GlueActivities) GetMLTransform(ctx context.Context, input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error) {
    return a.client.GetMLTransformWithContext(ctx, input)
}

func (a *GlueActivities) GetMLTransforms(ctx context.Context, input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error) {
    return a.client.GetMLTransformsWithContext(ctx, input)
}

func (a *GlueActivities) GetMapping(ctx context.Context, input *glue.GetMappingInput) (*glue.GetMappingOutput, error) {
    return a.client.GetMappingWithContext(ctx, input)
}

func (a *GlueActivities) GetPartition(ctx context.Context, input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error) {
    return a.client.GetPartitionWithContext(ctx, input)
}

func (a *GlueActivities) GetPartitionIndexes(ctx context.Context, input *glue.GetPartitionIndexesInput) (*glue.GetPartitionIndexesOutput, error) {
    return a.client.GetPartitionIndexesWithContext(ctx, input)
}

func (a *GlueActivities) GetPartitions(ctx context.Context, input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error) {
    return a.client.GetPartitionsWithContext(ctx, input)
}

func (a *GlueActivities) GetPlan(ctx context.Context, input *glue.GetPlanInput) (*glue.GetPlanOutput, error) {
    return a.client.GetPlanWithContext(ctx, input)
}

func (a *GlueActivities) GetResourcePolicies(ctx context.Context, input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error) {
    return a.client.GetResourcePoliciesWithContext(ctx, input)
}

func (a *GlueActivities) GetResourcePolicy(ctx context.Context, input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error) {
    return a.client.GetResourcePolicyWithContext(ctx, input)
}

func (a *GlueActivities) GetSecurityConfiguration(ctx context.Context, input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error) {
    return a.client.GetSecurityConfigurationWithContext(ctx, input)
}

func (a *GlueActivities) GetSecurityConfigurations(ctx context.Context, input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error) {
    return a.client.GetSecurityConfigurationsWithContext(ctx, input)
}

func (a *GlueActivities) GetTable(ctx context.Context, input *glue.GetTableInput) (*glue.GetTableOutput, error) {
    return a.client.GetTableWithContext(ctx, input)
}

func (a *GlueActivities) GetTableVersion(ctx context.Context, input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error) {
    return a.client.GetTableVersionWithContext(ctx, input)
}

func (a *GlueActivities) GetTableVersions(ctx context.Context, input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error) {
    return a.client.GetTableVersionsWithContext(ctx, input)
}

func (a *GlueActivities) GetTables(ctx context.Context, input *glue.GetTablesInput) (*glue.GetTablesOutput, error) {
    return a.client.GetTablesWithContext(ctx, input)
}

func (a *GlueActivities) GetTags(ctx context.Context, input *glue.GetTagsInput) (*glue.GetTagsOutput, error) {
    return a.client.GetTagsWithContext(ctx, input)
}

func (a *GlueActivities) GetTrigger(ctx context.Context, input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error) {
    return a.client.GetTriggerWithContext(ctx, input)
}

func (a *GlueActivities) GetTriggers(ctx context.Context, input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error) {
    return a.client.GetTriggersWithContext(ctx, input)
}

func (a *GlueActivities) GetUserDefinedFunction(ctx context.Context, input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error) {
    return a.client.GetUserDefinedFunctionWithContext(ctx, input)
}

func (a *GlueActivities) GetUserDefinedFunctions(ctx context.Context, input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error) {
    return a.client.GetUserDefinedFunctionsWithContext(ctx, input)
}

func (a *GlueActivities) GetWorkflow(ctx context.Context, input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error) {
    return a.client.GetWorkflowWithContext(ctx, input)
}

func (a *GlueActivities) GetWorkflowRun(ctx context.Context, input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error) {
    return a.client.GetWorkflowRunWithContext(ctx, input)
}

func (a *GlueActivities) GetWorkflowRunProperties(ctx context.Context, input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error) {
    return a.client.GetWorkflowRunPropertiesWithContext(ctx, input)
}

func (a *GlueActivities) GetWorkflowRuns(ctx context.Context, input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error) {
    return a.client.GetWorkflowRunsWithContext(ctx, input)
}

func (a *GlueActivities) ImportCatalogToGlue(ctx context.Context, input *glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error) {
    return a.client.ImportCatalogToGlueWithContext(ctx, input)
}

func (a *GlueActivities) ListCrawlers(ctx context.Context, input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error) {
    return a.client.ListCrawlersWithContext(ctx, input)
}

func (a *GlueActivities) ListDevEndpoints(ctx context.Context, input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error) {
    return a.client.ListDevEndpointsWithContext(ctx, input)
}

func (a *GlueActivities) ListJobs(ctx context.Context, input *glue.ListJobsInput) (*glue.ListJobsOutput, error) {
    return a.client.ListJobsWithContext(ctx, input)
}

func (a *GlueActivities) ListMLTransforms(ctx context.Context, input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error) {
    return a.client.ListMLTransformsWithContext(ctx, input)
}

func (a *GlueActivities) ListTriggers(ctx context.Context, input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error) {
    return a.client.ListTriggersWithContext(ctx, input)
}

func (a *GlueActivities) ListWorkflows(ctx context.Context, input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error) {
    return a.client.ListWorkflowsWithContext(ctx, input)
}

func (a *GlueActivities) PutDataCatalogEncryptionSettings(ctx context.Context, input *glue.PutDataCatalogEncryptionSettingsInput) (*glue.PutDataCatalogEncryptionSettingsOutput, error) {
    return a.client.PutDataCatalogEncryptionSettingsWithContext(ctx, input)
}

func (a *GlueActivities) PutResourcePolicy(ctx context.Context, input *glue.PutResourcePolicyInput) (*glue.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicyWithContext(ctx, input)
}

func (a *GlueActivities) PutWorkflowRunProperties(ctx context.Context, input *glue.PutWorkflowRunPropertiesInput) (*glue.PutWorkflowRunPropertiesOutput, error) {
    return a.client.PutWorkflowRunPropertiesWithContext(ctx, input)
}

func (a *GlueActivities) ResetJobBookmark(ctx context.Context, input *glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error) {
    return a.client.ResetJobBookmarkWithContext(ctx, input)
}

func (a *GlueActivities) ResumeWorkflowRun(ctx context.Context, input *glue.ResumeWorkflowRunInput) (*glue.ResumeWorkflowRunOutput, error) {
    return a.client.ResumeWorkflowRunWithContext(ctx, input)
}

func (a *GlueActivities) SearchTables(ctx context.Context, input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error) {
    return a.client.SearchTablesWithContext(ctx, input)
}

func (a *GlueActivities) StartCrawler(ctx context.Context, input *glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error) {
    return a.client.StartCrawlerWithContext(ctx, input)
}

func (a *GlueActivities) StartCrawlerSchedule(ctx context.Context, input *glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error) {
    return a.client.StartCrawlerScheduleWithContext(ctx, input)
}

func (a *GlueActivities) StartExportLabelsTaskRun(ctx context.Context, input *glue.StartExportLabelsTaskRunInput) (*glue.StartExportLabelsTaskRunOutput, error) {
    return a.client.StartExportLabelsTaskRunWithContext(ctx, input)
}

func (a *GlueActivities) StartImportLabelsTaskRun(ctx context.Context, input *glue.StartImportLabelsTaskRunInput) (*glue.StartImportLabelsTaskRunOutput, error) {
    return a.client.StartImportLabelsTaskRunWithContext(ctx, input)
}

func (a *GlueActivities) StartJobRun(ctx context.Context, input *glue.StartJobRunInput) (*glue.StartJobRunOutput, error) {
    return a.client.StartJobRunWithContext(ctx, input)
}

func (a *GlueActivities) StartMLEvaluationTaskRun(ctx context.Context, input *glue.StartMLEvaluationTaskRunInput) (*glue.StartMLEvaluationTaskRunOutput, error) {
    return a.client.StartMLEvaluationTaskRunWithContext(ctx, input)
}

func (a *GlueActivities) StartMLLabelingSetGenerationTaskRun(ctx context.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error) {
    return a.client.StartMLLabelingSetGenerationTaskRunWithContext(ctx, input)
}

func (a *GlueActivities) StartTrigger(ctx context.Context, input *glue.StartTriggerInput) (*glue.StartTriggerOutput, error) {
    return a.client.StartTriggerWithContext(ctx, input)
}

func (a *GlueActivities) StartWorkflowRun(ctx context.Context, input *glue.StartWorkflowRunInput) (*glue.StartWorkflowRunOutput, error) {
    return a.client.StartWorkflowRunWithContext(ctx, input)
}

func (a *GlueActivities) StopCrawler(ctx context.Context, input *glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error) {
    return a.client.StopCrawlerWithContext(ctx, input)
}

func (a *GlueActivities) StopCrawlerSchedule(ctx context.Context, input *glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error) {
    return a.client.StopCrawlerScheduleWithContext(ctx, input)
}

func (a *GlueActivities) StopTrigger(ctx context.Context, input *glue.StopTriggerInput) (*glue.StopTriggerOutput, error) {
    return a.client.StopTriggerWithContext(ctx, input)
}

func (a *GlueActivities) StopWorkflowRun(ctx context.Context, input *glue.StopWorkflowRunInput) (*glue.StopWorkflowRunOutput, error) {
    return a.client.StopWorkflowRunWithContext(ctx, input)
}

func (a *GlueActivities) TagResource(ctx context.Context, input *glue.TagResourceInput) (*glue.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *GlueActivities) UntagResource(ctx context.Context, input *glue.UntagResourceInput) (*glue.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *GlueActivities) UpdateClassifier(ctx context.Context, input *glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error) {
    return a.client.UpdateClassifierWithContext(ctx, input)
}

func (a *GlueActivities) UpdateColumnStatisticsForPartition(ctx context.Context, input *glue.UpdateColumnStatisticsForPartitionInput) (*glue.UpdateColumnStatisticsForPartitionOutput, error) {
    return a.client.UpdateColumnStatisticsForPartitionWithContext(ctx, input)
}

func (a *GlueActivities) UpdateColumnStatisticsForTable(ctx context.Context, input *glue.UpdateColumnStatisticsForTableInput) (*glue.UpdateColumnStatisticsForTableOutput, error) {
    return a.client.UpdateColumnStatisticsForTableWithContext(ctx, input)
}

func (a *GlueActivities) UpdateConnection(ctx context.Context, input *glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error) {
    return a.client.UpdateConnectionWithContext(ctx, input)
}

func (a *GlueActivities) UpdateCrawler(ctx context.Context, input *glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error) {
    return a.client.UpdateCrawlerWithContext(ctx, input)
}

func (a *GlueActivities) UpdateCrawlerSchedule(ctx context.Context, input *glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error) {
    return a.client.UpdateCrawlerScheduleWithContext(ctx, input)
}

func (a *GlueActivities) UpdateDatabase(ctx context.Context, input *glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error) {
    return a.client.UpdateDatabaseWithContext(ctx, input)
}

func (a *GlueActivities) UpdateDevEndpoint(ctx context.Context, input *glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error) {
    return a.client.UpdateDevEndpointWithContext(ctx, input)
}

func (a *GlueActivities) UpdateJob(ctx context.Context, input *glue.UpdateJobInput) (*glue.UpdateJobOutput, error) {
    return a.client.UpdateJobWithContext(ctx, input)
}

func (a *GlueActivities) UpdateMLTransform(ctx context.Context, input *glue.UpdateMLTransformInput) (*glue.UpdateMLTransformOutput, error) {
    return a.client.UpdateMLTransformWithContext(ctx, input)
}

func (a *GlueActivities) UpdatePartition(ctx context.Context, input *glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error) {
    return a.client.UpdatePartitionWithContext(ctx, input)
}

func (a *GlueActivities) UpdateTable(ctx context.Context, input *glue.UpdateTableInput) (*glue.UpdateTableOutput, error) {
    return a.client.UpdateTableWithContext(ctx, input)
}

func (a *GlueActivities) UpdateTrigger(ctx context.Context, input *glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error) {
    return a.client.UpdateTriggerWithContext(ctx, input)
}

func (a *GlueActivities) UpdateUserDefinedFunction(ctx context.Context, input *glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error) {
    return a.client.UpdateUserDefinedFunctionWithContext(ctx, input)
}

func (a *GlueActivities) UpdateWorkflow(ctx context.Context, input *glue.UpdateWorkflowInput) (*glue.UpdateWorkflowOutput, error) {
    return a.client.UpdateWorkflowWithContext(ctx, input)
}
