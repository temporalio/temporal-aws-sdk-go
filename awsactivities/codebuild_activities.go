// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CodeBuildActivities struct {
	client codebuildiface.CodeBuildAPI

	sessionFactory SessionFactory
}

func NewCodeBuildActivities(sess *session.Session, config ...*aws.Config) *CodeBuildActivities {
	client := codebuild.New(sess, config...)
	return &CodeBuildActivities{client: client}
}

func NewCodeBuildActivitiesWithSessionFactory(sessionFactory SessionFactory) *CodeBuildActivities {
	return &CodeBuildActivities{sessionFactory: sessionFactory}
}

func (a *CodeBuildActivities) getClient(ctx context.Context) (codebuildiface.CodeBuildAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return codebuild.New(sess), nil
}

func (a *CodeBuildActivities) BatchDeleteBuilds(ctx context.Context, input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDeleteBuildsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetBuildBatches(ctx context.Context, input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetBuildBatchesWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetBuilds(ctx context.Context, input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetBuildsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetProjects(ctx context.Context, input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetProjectsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetReportGroups(ctx context.Context, input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetReportGroupsWithContext(ctx, input)
}

func (a *CodeBuildActivities) BatchGetReports(ctx context.Context, input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGetReportsWithContext(ctx, input)
}

func (a *CodeBuildActivities) CreateProject(ctx context.Context, input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) CreateReportGroup(ctx context.Context, input *codebuild.CreateReportGroupInput) (*codebuild.CreateReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) CreateWebhook(ctx context.Context, input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateWebhookWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteBuildBatch(ctx context.Context, input *codebuild.DeleteBuildBatchInput) (*codebuild.DeleteBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteProject(ctx context.Context, input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteReport(ctx context.Context, input *codebuild.DeleteReportInput) (*codebuild.DeleteReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReportWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteReportGroup(ctx context.Context, input *codebuild.DeleteReportGroupInput) (*codebuild.DeleteReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteResourcePolicy(ctx context.Context, input *codebuild.DeleteResourcePolicyInput) (*codebuild.DeleteResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResourcePolicyWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteSourceCredentials(ctx context.Context, input *codebuild.DeleteSourceCredentialsInput) (*codebuild.DeleteSourceCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSourceCredentialsWithContext(ctx, input)
}

func (a *CodeBuildActivities) DeleteWebhook(ctx context.Context, input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteWebhookWithContext(ctx, input)
}

func (a *CodeBuildActivities) DescribeCodeCoverages(ctx context.Context, input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCodeCoveragesWithContext(ctx, input)
}

func (a *CodeBuildActivities) DescribeTestCases(ctx context.Context, input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTestCasesWithContext(ctx, input)
}

func (a *CodeBuildActivities) GetResourcePolicy(ctx context.Context, input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResourcePolicyWithContext(ctx, input)
}

func (a *CodeBuildActivities) ImportSourceCredentials(ctx context.Context, input *codebuild.ImportSourceCredentialsInput) (*codebuild.ImportSourceCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ImportSourceCredentialsWithContext(ctx, input)
}

func (a *CodeBuildActivities) InvalidateProjectCache(ctx context.Context, input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.InvalidateProjectCacheWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuildBatches(ctx context.Context, input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBuildBatchesWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuildBatchesForProject(ctx context.Context, input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBuildBatchesForProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuilds(ctx context.Context, input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBuildsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListBuildsForProject(ctx context.Context, input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListBuildsForProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListCuratedEnvironmentImages(ctx context.Context, input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListCuratedEnvironmentImagesWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListProjects(ctx context.Context, input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProjectsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListReportGroups(ctx context.Context, input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListReportGroupsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListReports(ctx context.Context, input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListReportsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListReportsForReportGroup(ctx context.Context, input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListReportsForReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListSharedProjects(ctx context.Context, input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSharedProjectsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListSharedReportGroups(ctx context.Context, input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSharedReportGroupsWithContext(ctx, input)
}

func (a *CodeBuildActivities) ListSourceCredentials(ctx context.Context, input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSourceCredentialsWithContext(ctx, input)
}

func (a *CodeBuildActivities) PutResourcePolicy(ctx context.Context, input *codebuild.PutResourcePolicyInput) (*codebuild.PutResourcePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutResourcePolicyWithContext(ctx, input)
}

func (a *CodeBuildActivities) RetryBuild(ctx context.Context, input *codebuild.RetryBuildInput) (*codebuild.RetryBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RetryBuildWithContext(ctx, input)
}

func (a *CodeBuildActivities) RetryBuildBatch(ctx context.Context, input *codebuild.RetryBuildBatchInput) (*codebuild.RetryBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RetryBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) StartBuild(ctx context.Context, input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartBuildWithContext(ctx, input)
}

func (a *CodeBuildActivities) StartBuildBatch(ctx context.Context, input *codebuild.StartBuildBatchInput) (*codebuild.StartBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) StopBuild(ctx context.Context, input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopBuildWithContext(ctx, input)
}

func (a *CodeBuildActivities) StopBuildBatch(ctx context.Context, input *codebuild.StopBuildBatchInput) (*codebuild.StopBuildBatchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopBuildBatchWithContext(ctx, input)
}

func (a *CodeBuildActivities) UpdateProject(ctx context.Context, input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProjectWithContext(ctx, input)
}

func (a *CodeBuildActivities) UpdateReportGroup(ctx context.Context, input *codebuild.UpdateReportGroupInput) (*codebuild.UpdateReportGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateReportGroupWithContext(ctx, input)
}

func (a *CodeBuildActivities) UpdateWebhook(ctx context.Context, input *codebuild.UpdateWebhookInput) (*codebuild.UpdateWebhookOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateWebhookWithContext(ctx, input)
}
