package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
)

type CodeBuildActivities struct {
	client codebuildiface.CodeBuildAPI
}

func NewCodeBuildActivities(client codebuildiface.CodeBuildAPI) *CodeBuildActivities {
    return &CodeBuildActivities{client: client}
}


func (a *CodeBuildActivities) BatchDeleteBuilds(input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
    return a.client.BatchDeleteBuilds(input)
}



func (a *CodeBuildActivities) BatchGetBuildBatches(input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error) {
    return a.client.BatchGetBuildBatches(input)
}



func (a *CodeBuildActivities) BatchGetBuilds(input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
    return a.client.BatchGetBuilds(input)
}



func (a *CodeBuildActivities) BatchGetProjects(input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
    return a.client.BatchGetProjects(input)
}



func (a *CodeBuildActivities) BatchGetReportGroups(input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error) {
    return a.client.BatchGetReportGroups(input)
}



func (a *CodeBuildActivities) BatchGetReports(input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error) {
    return a.client.BatchGetReports(input)
}



func (a *CodeBuildActivities) CreateProject(input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error) {
    return a.client.CreateProject(input)
}



func (a *CodeBuildActivities) CreateReportGroup(input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error) {
    return a.client.CreateReportGroup(input)
}



func (a *CodeBuildActivities) CreateWebhook(input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error) {
    return a.client.CreateWebhook(input)
}



func (a *CodeBuildActivities) DeleteBuildBatch(input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error) {
    return a.client.DeleteBuildBatch(input)
}



func (a *CodeBuildActivities) DeleteProject(input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error) {
    return a.client.DeleteProject(input)
}



func (a *CodeBuildActivities) DeleteReport(input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error) {
    return a.client.DeleteReport(input)
}



func (a *CodeBuildActivities) DeleteReportGroup(input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error) {
    return a.client.DeleteReportGroup(input)
}



func (a *CodeBuildActivities) DeleteResourcePolicy(input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error) {
    return a.client.DeleteResourcePolicy(input)
}



func (a *CodeBuildActivities) DeleteSourceCredentials(input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error) {
    return a.client.DeleteSourceCredentials(input)
}



func (a *CodeBuildActivities) DeleteWebhook(input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error) {
    return a.client.DeleteWebhook(input)
}



func (a *CodeBuildActivities) DescribeCodeCoverages(input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error) {
    return a.client.DescribeCodeCoverages(input)
}



func (a *CodeBuildActivities) DescribeTestCases(input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error) {
    return a.client.DescribeTestCases(input)
}



func (a *CodeBuildActivities) GetResourcePolicy(input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error) {
    return a.client.GetResourcePolicy(input)
}



func (a *CodeBuildActivities) ImportSourceCredentials(input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error) {
    return a.client.ImportSourceCredentials(input)
}



func (a *CodeBuildActivities) InvalidateProjectCache(input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error) {
    return a.client.InvalidateProjectCache(input)
}



func (a *CodeBuildActivities) ListBuildBatches(input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error) {
    return a.client.ListBuildBatches(input)
}



func (a *CodeBuildActivities) ListBuildBatchesForProject(input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error) {
    return a.client.ListBuildBatchesForProject(input)
}



func (a *CodeBuildActivities) ListBuilds(input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
    return a.client.ListBuilds(input)
}



func (a *CodeBuildActivities) ListBuildsForProject(input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
    return a.client.ListBuildsForProject(input)
}



func (a *CodeBuildActivities) ListCuratedEnvironmentImages(input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
    return a.client.ListCuratedEnvironmentImages(input)
}



func (a *CodeBuildActivities) ListProjects(input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
    return a.client.ListProjects(input)
}



func (a *CodeBuildActivities) ListReportGroups(input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error) {
    return a.client.ListReportGroups(input)
}



func (a *CodeBuildActivities) ListReports(input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error) {
    return a.client.ListReports(input)
}



func (a *CodeBuildActivities) ListReportsForReportGroup(input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error) {
    return a.client.ListReportsForReportGroup(input)
}



func (a *CodeBuildActivities) ListSharedProjects(input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error) {
    return a.client.ListSharedProjects(input)
}



func (a *CodeBuildActivities) ListSharedReportGroups(input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error) {
    return a.client.ListSharedReportGroups(input)
}



func (a *CodeBuildActivities) ListSourceCredentials(input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error) {
    return a.client.ListSourceCredentials(input)
}



func (a *CodeBuildActivities) PutResourcePolicy(input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error) {
    return a.client.PutResourcePolicy(input)
}



func (a *CodeBuildActivities) RetryBuild(input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error) {
    return a.client.RetryBuild(input)
}



func (a *CodeBuildActivities) RetryBuildBatch(input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error) {
    return a.client.RetryBuildBatch(input)
}



func (a *CodeBuildActivities) StartBuild(input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error) {
    return a.client.StartBuild(input)
}



func (a *CodeBuildActivities) StartBuildBatch(input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error) {
    return a.client.StartBuildBatch(input)
}



func (a *CodeBuildActivities) StopBuild(input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error) {
    return a.client.StopBuild(input)
}



func (a *CodeBuildActivities) StopBuildBatch(input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error) {
    return a.client.StopBuildBatch(input)
}



func (a *CodeBuildActivities) UpdateProject(input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error) {
    return a.client.UpdateProject(input)
}



func (a *CodeBuildActivities) UpdateReportGroup(input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error) {
    return a.client.UpdateReportGroup(input)
}



func (a *CodeBuildActivities) UpdateWebhook(input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error) {
    return a.client.UpdateWebhook(input)
}

