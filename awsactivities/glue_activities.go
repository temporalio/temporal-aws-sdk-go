
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/glue/glueiface"
)

type GlueActivities struct {
	client glueiface.GlueAPI
}

func NewGlueActivities(session *session.Session, config... *aws.Config) *GlueActivities {
    client := glue.New(session, config...)
    return &GlueActivities{client: client}
}

func (a *GlueActivities) BatchCreatePartition(input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error) {
    return a.client.BatchCreatePartition(input)
}

func (a *GlueActivities) BatchDeleteConnection(input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error) {
    return a.client.BatchDeleteConnection(input)
}

func (a *GlueActivities) BatchDeletePartition(input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error) {
    return a.client.BatchDeletePartition(input)
}

func (a *GlueActivities) BatchDeleteTable(input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error) {
    return a.client.BatchDeleteTable(input)
}

func (a *GlueActivities) BatchDeleteTableVersion(input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error) {
    return a.client.BatchDeleteTableVersion(input)
}

func (a *GlueActivities) BatchGetCrawlers(input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error) {
    return a.client.BatchGetCrawlers(input)
}

func (a *GlueActivities) BatchGetDevEndpoints(input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error) {
    return a.client.BatchGetDevEndpoints(input)
}

func (a *GlueActivities) BatchGetJobs(input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error) {
    return a.client.BatchGetJobs(input)
}

func (a *GlueActivities) BatchGetPartition(input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error) {
    return a.client.BatchGetPartition(input)
}

func (a *GlueActivities) BatchGetTriggers(input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error) {
    return a.client.BatchGetTriggers(input)
}

func (a *GlueActivities) BatchGetWorkflows(input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error) {
    return a.client.BatchGetWorkflows(input)
}

func (a *GlueActivities) BatchStopJobRun(input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error) {
    return a.client.BatchStopJobRun(input)
}

func (a *GlueActivities) CancelMLTaskRun(input *glue.CancelMLTaskRunInput) (*glue.CancelMLTaskRunOutput, error) {
    return a.client.CancelMLTaskRun(input)
}

func (a *GlueActivities) CreateClassifier(input *glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error) {
    return a.client.CreateClassifier(input)
}

func (a *GlueActivities) CreateConnection(input *glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error) {
    return a.client.CreateConnection(input)
}

func (a *GlueActivities) CreateCrawler(input *glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error) {
    return a.client.CreateCrawler(input)
}

func (a *GlueActivities) CreateDatabase(input *glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error) {
    return a.client.CreateDatabase(input)
}

func (a *GlueActivities) CreateDevEndpoint(input *glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error) {
    return a.client.CreateDevEndpoint(input)
}

func (a *GlueActivities) CreateJob(input *glue.CreateJobInput) (*glue.CreateJobOutput, error) {
    return a.client.CreateJob(input)
}

func (a *GlueActivities) CreateMLTransform(input *glue.CreateMLTransformInput) (*glue.CreateMLTransformOutput, error) {
    return a.client.CreateMLTransform(input)
}

func (a *GlueActivities) CreatePartition(input *glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error) {
    return a.client.CreatePartition(input)
}

func (a *GlueActivities) CreateScript(input *glue.CreateScriptInput) (*glue.CreateScriptOutput, error) {
    return a.client.CreateScript(input)
}

func (a *GlueActivities) CreateSecurityConfiguration(input *glue.CreateSecurityConfigurationInput) (*glue.CreateSecurityConfigurationOutput, error) {
    return a.client.CreateSecurityConfiguration(input)
}

func (a *GlueActivities) CreateTable(input *glue.CreateTableInput) (*glue.CreateTableOutput, error) {
    return a.client.CreateTable(input)
}

func (a *GlueActivities) CreateTrigger(input *glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error) {
    return a.client.CreateTrigger(input)
}

func (a *GlueActivities) CreateUserDefinedFunction(input *glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error) {
    return a.client.CreateUserDefinedFunction(input)
}

func (a *GlueActivities) CreateWorkflow(input *glue.CreateWorkflowInput) (*glue.CreateWorkflowOutput, error) {
    return a.client.CreateWorkflow(input)
}

func (a *GlueActivities) DeleteClassifier(input *glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error) {
    return a.client.DeleteClassifier(input)
}

func (a *GlueActivities) DeleteColumnStatisticsForPartition(input *glue.DeleteColumnStatisticsForPartitionInput) (*glue.DeleteColumnStatisticsForPartitionOutput, error) {
    return a.client.DeleteColumnStatisticsForPartition(input)
}

func (a *GlueActivities) DeleteColumnStatisticsForTable(input *glue.DeleteColumnStatisticsForTableInput) (*glue.DeleteColumnStatisticsForTableOutput, error) {
    return a.client.DeleteColumnStatisticsForTable(input)
}

func (a *GlueActivities) DeleteConnection(input *glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error) {
    return a.client.DeleteConnection(input)
}

func (a *GlueActivities) DeleteCrawler(input *glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error) {
    return a.client.DeleteCrawler(input)
}

func (a *GlueActivities) DeleteDatabase(input *glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error) {
    return a.client.DeleteDatabase(input)
}

func (a *GlueActivities) DeleteDevEndpoint(input *glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error) {
    return a.client.DeleteDevEndpoint(input)
}

func (a *GlueActivities) DeleteJob(input *glue.DeleteJobInput) (*glue.DeleteJobOutput, error) {
    return a.client.DeleteJob(input)
}

func (a *GlueActivities) DeleteMLTransform(input *glue.DeleteMLTransformInput) (*glue.DeleteMLTransformOutput, error) {
    return a.client.DeleteMLTransform(input)
}

func (a *GlueActivities) DeletePartition(input *glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error) {
    return a.client.DeletePartition(input)
}

func (a *GlueActivities) DeleteResourcePolicy(input *glue.DeleteResourcePolicyInput) (*glue.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicy(input)
}

func (a *GlueActivities) DeleteSecurityConfiguration(input *glue.DeleteSecurityConfigurationInput) (*glue.DeleteSecurityConfigurationOutput, error) {
    return a.client.DeleteSecurityConfiguration(input)
}

func (a *GlueActivities) DeleteTable(input *glue.DeleteTableInput) (*glue.DeleteTableOutput, error) {
    return a.client.DeleteTable(input)
}

func (a *GlueActivities) DeleteTableVersion(input *glue.DeleteTableVersionInput) (*glue.DeleteTableVersionOutput, error) {
    return a.client.DeleteTableVersion(input)
}

func (a *GlueActivities) DeleteTrigger(input *glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error) {
    return a.client.DeleteTrigger(input)
}

func (a *GlueActivities) DeleteUserDefinedFunction(input *glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error) {
    return a.client.DeleteUserDefinedFunction(input)
}

func (a *GlueActivities) DeleteWorkflow(input *glue.DeleteWorkflowInput) (*glue.DeleteWorkflowOutput, error) {
    return a.client.DeleteWorkflow(input)
}

func (a *GlueActivities) GetCatalogImportStatus(input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error) {
    return a.client.GetCatalogImportStatus(input)
}

func (a *GlueActivities) GetClassifier(input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error) {
    return a.client.GetClassifier(input)
}

func (a *GlueActivities) GetClassifiers(input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error) {
    return a.client.GetClassifiers(input)
}

func (a *GlueActivities) GetColumnStatisticsForPartition(input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error) {
    return a.client.GetColumnStatisticsForPartition(input)
}

func (a *GlueActivities) GetColumnStatisticsForTable(input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error) {
    return a.client.GetColumnStatisticsForTable(input)
}

func (a *GlueActivities) GetConnection(input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error) {
    return a.client.GetConnection(input)
}

func (a *GlueActivities) GetConnections(input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error) {
    return a.client.GetConnections(input)
}

func (a *GlueActivities) GetCrawler(input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error) {
    return a.client.GetCrawler(input)
}

func (a *GlueActivities) GetCrawlerMetrics(input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error) {
    return a.client.GetCrawlerMetrics(input)
}

func (a *GlueActivities) GetCrawlers(input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error) {
    return a.client.GetCrawlers(input)
}

func (a *GlueActivities) GetDataCatalogEncryptionSettings(input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
    return a.client.GetDataCatalogEncryptionSettings(input)
}

func (a *GlueActivities) GetDatabase(input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error) {
    return a.client.GetDatabase(input)
}

func (a *GlueActivities) GetDatabases(input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error) {
    return a.client.GetDatabases(input)
}

func (a *GlueActivities) GetDataflowGraph(input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error) {
    return a.client.GetDataflowGraph(input)
}

func (a *GlueActivities) GetDevEndpoint(input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error) {
    return a.client.GetDevEndpoint(input)
}

func (a *GlueActivities) GetDevEndpoints(input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error) {
    return a.client.GetDevEndpoints(input)
}

func (a *GlueActivities) GetJob(input *glue.GetJobInput) (*glue.GetJobOutput, error) {
    return a.client.GetJob(input)
}

func (a *GlueActivities) GetJobBookmark(input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error) {
    return a.client.GetJobBookmark(input)
}

func (a *GlueActivities) GetJobRun(input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error) {
    return a.client.GetJobRun(input)
}

func (a *GlueActivities) GetJobRuns(input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error) {
    return a.client.GetJobRuns(input)
}

func (a *GlueActivities) GetJobs(input *glue.GetJobsInput) (*glue.GetJobsOutput, error) {
    return a.client.GetJobs(input)
}

func (a *GlueActivities) GetMLTaskRun(input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error) {
    return a.client.GetMLTaskRun(input)
}

func (a *GlueActivities) GetMLTaskRuns(input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error) {
    return a.client.GetMLTaskRuns(input)
}

func (a *GlueActivities) GetMLTransform(input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error) {
    return a.client.GetMLTransform(input)
}

func (a *GlueActivities) GetMLTransforms(input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error) {
    return a.client.GetMLTransforms(input)
}

func (a *GlueActivities) GetMapping(input *glue.GetMappingInput) (*glue.GetMappingOutput, error) {
    return a.client.GetMapping(input)
}

func (a *GlueActivities) GetPartition(input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error) {
    return a.client.GetPartition(input)
}

func (a *GlueActivities) GetPartitions(input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error) {
    return a.client.GetPartitions(input)
}

func (a *GlueActivities) GetPlan(input *glue.GetPlanInput) (*glue.GetPlanOutput, error) {
    return a.client.GetPlan(input)
}

func (a *GlueActivities) GetResourcePolicies(input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error) {
    return a.client.GetResourcePolicies(input)
}

func (a *GlueActivities) GetResourcePolicy(input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error) {
    return a.client.GetResourcePolicy(input)
}

func (a *GlueActivities) GetSecurityConfiguration(input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error) {
    return a.client.GetSecurityConfiguration(input)
}

func (a *GlueActivities) GetSecurityConfigurations(input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error) {
    return a.client.GetSecurityConfigurations(input)
}

func (a *GlueActivities) GetTable(input *glue.GetTableInput) (*glue.GetTableOutput, error) {
    return a.client.GetTable(input)
}

func (a *GlueActivities) GetTableVersion(input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error) {
    return a.client.GetTableVersion(input)
}

func (a *GlueActivities) GetTableVersions(input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error) {
    return a.client.GetTableVersions(input)
}

func (a *GlueActivities) GetTables(input *glue.GetTablesInput) (*glue.GetTablesOutput, error) {
    return a.client.GetTables(input)
}

func (a *GlueActivities) GetTags(input *glue.GetTagsInput) (*glue.GetTagsOutput, error) {
    return a.client.GetTags(input)
}

func (a *GlueActivities) GetTrigger(input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error) {
    return a.client.GetTrigger(input)
}

func (a *GlueActivities) GetTriggers(input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error) {
    return a.client.GetTriggers(input)
}

func (a *GlueActivities) GetUserDefinedFunction(input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error) {
    return a.client.GetUserDefinedFunction(input)
}

func (a *GlueActivities) GetUserDefinedFunctions(input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error) {
    return a.client.GetUserDefinedFunctions(input)
}

func (a *GlueActivities) GetWorkflow(input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error) {
    return a.client.GetWorkflow(input)
}

func (a *GlueActivities) GetWorkflowRun(input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error) {
    return a.client.GetWorkflowRun(input)
}

func (a *GlueActivities) GetWorkflowRunProperties(input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error) {
    return a.client.GetWorkflowRunProperties(input)
}

func (a *GlueActivities) GetWorkflowRuns(input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error) {
    return a.client.GetWorkflowRuns(input)
}

func (a *GlueActivities) ImportCatalogToGlue(input *glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error) {
    return a.client.ImportCatalogToGlue(input)
}

func (a *GlueActivities) ListCrawlers(input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error) {
    return a.client.ListCrawlers(input)
}

func (a *GlueActivities) ListDevEndpoints(input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error) {
    return a.client.ListDevEndpoints(input)
}

func (a *GlueActivities) ListJobs(input *glue.ListJobsInput) (*glue.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}

func (a *GlueActivities) ListMLTransforms(input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error) {
    return a.client.ListMLTransforms(input)
}

func (a *GlueActivities) ListTriggers(input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error) {
    return a.client.ListTriggers(input)
}

func (a *GlueActivities) ListWorkflows(input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error) {
    return a.client.ListWorkflows(input)
}

func (a *GlueActivities) PutDataCatalogEncryptionSettings(input *glue.PutDataCatalogEncryptionSettingsInput) (*glue.PutDataCatalogEncryptionSettingsOutput, error) {
    return a.client.PutDataCatalogEncryptionSettings(input)
}

func (a *GlueActivities) PutResourcePolicy(input *glue.PutResourcePolicyInput) (*glue.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicy(input)
}

func (a *GlueActivities) PutWorkflowRunProperties(input *glue.PutWorkflowRunPropertiesInput) (*glue.PutWorkflowRunPropertiesOutput, error) {
    return a.client.PutWorkflowRunProperties(input)
}

func (a *GlueActivities) ResetJobBookmark(input *glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error) {
    return a.client.ResetJobBookmark(input)
}

func (a *GlueActivities) ResumeWorkflowRun(input *glue.ResumeWorkflowRunInput) (*glue.ResumeWorkflowRunOutput, error) {
    return a.client.ResumeWorkflowRun(input)
}

func (a *GlueActivities) SearchTables(input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error) {
    return a.client.SearchTables(input)
}

func (a *GlueActivities) StartCrawler(input *glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error) {
    return a.client.StartCrawler(input)
}

func (a *GlueActivities) StartCrawlerSchedule(input *glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error) {
    return a.client.StartCrawlerSchedule(input)
}

func (a *GlueActivities) StartExportLabelsTaskRun(input *glue.StartExportLabelsTaskRunInput) (*glue.StartExportLabelsTaskRunOutput, error) {
    return a.client.StartExportLabelsTaskRun(input)
}

func (a *GlueActivities) StartImportLabelsTaskRun(input *glue.StartImportLabelsTaskRunInput) (*glue.StartImportLabelsTaskRunOutput, error) {
    return a.client.StartImportLabelsTaskRun(input)
}

func (a *GlueActivities) StartJobRun(input *glue.StartJobRunInput) (*glue.StartJobRunOutput, error) {
    return a.client.StartJobRun(input)
}

func (a *GlueActivities) StartMLEvaluationTaskRun(input *glue.StartMLEvaluationTaskRunInput) (*glue.StartMLEvaluationTaskRunOutput, error) {
    return a.client.StartMLEvaluationTaskRun(input)
}

func (a *GlueActivities) StartMLLabelingSetGenerationTaskRun(input *glue.StartMLLabelingSetGenerationTaskRunInput) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error) {
    return a.client.StartMLLabelingSetGenerationTaskRun(input)
}

func (a *GlueActivities) StartTrigger(input *glue.StartTriggerInput) (*glue.StartTriggerOutput, error) {
    return a.client.StartTrigger(input)
}

func (a *GlueActivities) StartWorkflowRun(input *glue.StartWorkflowRunInput) (*glue.StartWorkflowRunOutput, error) {
    return a.client.StartWorkflowRun(input)
}

func (a *GlueActivities) StopCrawler(input *glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error) {
    return a.client.StopCrawler(input)
}

func (a *GlueActivities) StopCrawlerSchedule(input *glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error) {
    return a.client.StopCrawlerSchedule(input)
}

func (a *GlueActivities) StopTrigger(input *glue.StopTriggerInput) (*glue.StopTriggerOutput, error) {
    return a.client.StopTrigger(input)
}

func (a *GlueActivities) StopWorkflowRun(input *glue.StopWorkflowRunInput) (*glue.StopWorkflowRunOutput, error) {
    return a.client.StopWorkflowRun(input)
}

func (a *GlueActivities) TagResource(input *glue.TagResourceInput) (*glue.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *GlueActivities) UntagResource(input *glue.UntagResourceInput) (*glue.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *GlueActivities) UpdateClassifier(input *glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error) {
    return a.client.UpdateClassifier(input)
}

func (a *GlueActivities) UpdateColumnStatisticsForPartition(input *glue.UpdateColumnStatisticsForPartitionInput) (*glue.UpdateColumnStatisticsForPartitionOutput, error) {
    return a.client.UpdateColumnStatisticsForPartition(input)
}

func (a *GlueActivities) UpdateColumnStatisticsForTable(input *glue.UpdateColumnStatisticsForTableInput) (*glue.UpdateColumnStatisticsForTableOutput, error) {
    return a.client.UpdateColumnStatisticsForTable(input)
}

func (a *GlueActivities) UpdateConnection(input *glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error) {
    return a.client.UpdateConnection(input)
}

func (a *GlueActivities) UpdateCrawler(input *glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error) {
    return a.client.UpdateCrawler(input)
}

func (a *GlueActivities) UpdateCrawlerSchedule(input *glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error) {
    return a.client.UpdateCrawlerSchedule(input)
}

func (a *GlueActivities) UpdateDatabase(input *glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error) {
    return a.client.UpdateDatabase(input)
}

func (a *GlueActivities) UpdateDevEndpoint(input *glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error) {
    return a.client.UpdateDevEndpoint(input)
}

func (a *GlueActivities) UpdateJob(input *glue.UpdateJobInput) (*glue.UpdateJobOutput, error) {
    return a.client.UpdateJob(input)
}

func (a *GlueActivities) UpdateMLTransform(input *glue.UpdateMLTransformInput) (*glue.UpdateMLTransformOutput, error) {
    return a.client.UpdateMLTransform(input)
}

func (a *GlueActivities) UpdatePartition(input *glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error) {
    return a.client.UpdatePartition(input)
}

func (a *GlueActivities) UpdateTable(input *glue.UpdateTableInput) (*glue.UpdateTableOutput, error) {
    return a.client.UpdateTable(input)
}

func (a *GlueActivities) UpdateTrigger(input *glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error) {
    return a.client.UpdateTrigger(input)
}

func (a *GlueActivities) UpdateUserDefinedFunction(input *glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error) {
    return a.client.UpdateUserDefinedFunction(input)
}

func (a *GlueActivities) UpdateWorkflow(input *glue.UpdateWorkflowInput) (*glue.UpdateWorkflowOutput, error) {
    return a.client.UpdateWorkflow(input)
}
