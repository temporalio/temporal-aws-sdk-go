package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/aws/aws-sdk-go/service/quicksight/quicksightiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type QuickSightActivities struct {
	client quicksightiface.QuickSightAPI
}

func NewQuickSightActivities(session *session.Session, config ...*aws.Config) *QuickSightActivities {
	client := quicksight.New(session, config...)
	return &QuickSightActivities{client: client}
}

func (a *QuickSightActivities) CancelIngestion(ctx context.Context, input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error) {
	return a.client.CancelIngestionWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateAccountCustomization(ctx context.Context, input *quicksight.CreateAccountCustomizationInput) (*quicksight.CreateAccountCustomizationOutput, error) {
	return a.client.CreateAccountCustomizationWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateAnalysis(ctx context.Context, input *quicksight.CreateAnalysisInput) (*quicksight.CreateAnalysisOutput, error) {
	return a.client.CreateAnalysisWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateDashboard(ctx context.Context, input *quicksight.CreateDashboardInput) (*quicksight.CreateDashboardOutput, error) {
	return a.client.CreateDashboardWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateDataSet(ctx context.Context, input *quicksight.CreateDataSetInput) (*quicksight.CreateDataSetOutput, error) {
	return a.client.CreateDataSetWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateDataSource(ctx context.Context, input *quicksight.CreateDataSourceInput) (*quicksight.CreateDataSourceOutput, error) {
	return a.client.CreateDataSourceWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateGroup(ctx context.Context, input *quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error) {
	return a.client.CreateGroupWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateGroupMembership(ctx context.Context, input *quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error) {
	return a.client.CreateGroupMembershipWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateIAMPolicyAssignment(ctx context.Context, input *quicksight.CreateIAMPolicyAssignmentInput) (*quicksight.CreateIAMPolicyAssignmentOutput, error) {
	return a.client.CreateIAMPolicyAssignmentWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateIngestion(ctx context.Context, input *quicksight.CreateIngestionInput) (*quicksight.CreateIngestionOutput, error) {
	return a.client.CreateIngestionWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateNamespace(ctx context.Context, input *quicksight.CreateNamespaceInput) (*quicksight.CreateNamespaceOutput, error) {
	return a.client.CreateNamespaceWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateTemplate(ctx context.Context, input *quicksight.CreateTemplateInput) (*quicksight.CreateTemplateOutput, error) {
	return a.client.CreateTemplateWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateTemplateAlias(ctx context.Context, input *quicksight.CreateTemplateAliasInput) (*quicksight.CreateTemplateAliasOutput, error) {
	return a.client.CreateTemplateAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateTheme(ctx context.Context, input *quicksight.CreateThemeInput) (*quicksight.CreateThemeOutput, error) {
	return a.client.CreateThemeWithContext(ctx, input)
}

func (a *QuickSightActivities) CreateThemeAlias(ctx context.Context, input *quicksight.CreateThemeAliasInput) (*quicksight.CreateThemeAliasOutput, error) {
	return a.client.CreateThemeAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteAccountCustomization(ctx context.Context, input *quicksight.DeleteAccountCustomizationInput) (*quicksight.DeleteAccountCustomizationOutput, error) {
	return a.client.DeleteAccountCustomizationWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteAnalysis(ctx context.Context, input *quicksight.DeleteAnalysisInput) (*quicksight.DeleteAnalysisOutput, error) {
	return a.client.DeleteAnalysisWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteDashboard(ctx context.Context, input *quicksight.DeleteDashboardInput) (*quicksight.DeleteDashboardOutput, error) {
	return a.client.DeleteDashboardWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteDataSet(ctx context.Context, input *quicksight.DeleteDataSetInput) (*quicksight.DeleteDataSetOutput, error) {
	return a.client.DeleteDataSetWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteDataSource(ctx context.Context, input *quicksight.DeleteDataSourceInput) (*quicksight.DeleteDataSourceOutput, error) {
	return a.client.DeleteDataSourceWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteGroup(ctx context.Context, input *quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error) {
	return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteGroupMembership(ctx context.Context, input *quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error) {
	return a.client.DeleteGroupMembershipWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteIAMPolicyAssignment(ctx context.Context, input *quicksight.DeleteIAMPolicyAssignmentInput) (*quicksight.DeleteIAMPolicyAssignmentOutput, error) {
	return a.client.DeleteIAMPolicyAssignmentWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteNamespace(ctx context.Context, input *quicksight.DeleteNamespaceInput) (*quicksight.DeleteNamespaceOutput, error) {
	return a.client.DeleteNamespaceWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteTemplate(ctx context.Context, input *quicksight.DeleteTemplateInput) (*quicksight.DeleteTemplateOutput, error) {
	return a.client.DeleteTemplateWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteTemplateAlias(ctx context.Context, input *quicksight.DeleteTemplateAliasInput) (*quicksight.DeleteTemplateAliasOutput, error) {
	return a.client.DeleteTemplateAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteTheme(ctx context.Context, input *quicksight.DeleteThemeInput) (*quicksight.DeleteThemeOutput, error) {
	return a.client.DeleteThemeWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteThemeAlias(ctx context.Context, input *quicksight.DeleteThemeAliasInput) (*quicksight.DeleteThemeAliasOutput, error) {
	return a.client.DeleteThemeAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteUser(ctx context.Context, input *quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error) {
	return a.client.DeleteUserWithContext(ctx, input)
}

func (a *QuickSightActivities) DeleteUserByPrincipalId(ctx context.Context, input *quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error) {
	return a.client.DeleteUserByPrincipalIdWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeAccountCustomization(ctx context.Context, input *quicksight.DescribeAccountCustomizationInput) (*quicksight.DescribeAccountCustomizationOutput, error) {
	return a.client.DescribeAccountCustomizationWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeAccountSettings(ctx context.Context, input *quicksight.DescribeAccountSettingsInput) (*quicksight.DescribeAccountSettingsOutput, error) {
	return a.client.DescribeAccountSettingsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeAnalysis(ctx context.Context, input *quicksight.DescribeAnalysisInput) (*quicksight.DescribeAnalysisOutput, error) {
	return a.client.DescribeAnalysisWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeAnalysisPermissions(ctx context.Context, input *quicksight.DescribeAnalysisPermissionsInput) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
	return a.client.DescribeAnalysisPermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeDashboard(ctx context.Context, input *quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error) {
	return a.client.DescribeDashboardWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeDashboardPermissions(ctx context.Context, input *quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error) {
	return a.client.DescribeDashboardPermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeDataSet(ctx context.Context, input *quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error) {
	return a.client.DescribeDataSetWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeDataSetPermissions(ctx context.Context, input *quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error) {
	return a.client.DescribeDataSetPermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeDataSource(ctx context.Context, input *quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error) {
	return a.client.DescribeDataSourceWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeDataSourcePermissions(ctx context.Context, input *quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
	return a.client.DescribeDataSourcePermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeGroup(ctx context.Context, input *quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error) {
	return a.client.DescribeGroupWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeIAMPolicyAssignment(ctx context.Context, input *quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
	return a.client.DescribeIAMPolicyAssignmentWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeIngestion(ctx context.Context, input *quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error) {
	return a.client.DescribeIngestionWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeNamespace(ctx context.Context, input *quicksight.DescribeNamespaceInput) (*quicksight.DescribeNamespaceOutput, error) {
	return a.client.DescribeNamespaceWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeTemplate(ctx context.Context, input *quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error) {
	return a.client.DescribeTemplateWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeTemplateAlias(ctx context.Context, input *quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error) {
	return a.client.DescribeTemplateAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeTemplatePermissions(ctx context.Context, input *quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error) {
	return a.client.DescribeTemplatePermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeTheme(ctx context.Context, input *quicksight.DescribeThemeInput) (*quicksight.DescribeThemeOutput, error) {
	return a.client.DescribeThemeWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeThemeAlias(ctx context.Context, input *quicksight.DescribeThemeAliasInput) (*quicksight.DescribeThemeAliasOutput, error) {
	return a.client.DescribeThemeAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeThemePermissions(ctx context.Context, input *quicksight.DescribeThemePermissionsInput) (*quicksight.DescribeThemePermissionsOutput, error) {
	return a.client.DescribeThemePermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) DescribeUser(ctx context.Context, input *quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error) {
	return a.client.DescribeUserWithContext(ctx, input)
}

func (a *QuickSightActivities) GetDashboardEmbedUrl(ctx context.Context, input *quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error) {
	return a.client.GetDashboardEmbedUrlWithContext(ctx, input)
}

func (a *QuickSightActivities) GetSessionEmbedUrl(ctx context.Context, input *quicksight.GetSessionEmbedUrlInput) (*quicksight.GetSessionEmbedUrlOutput, error) {
	return a.client.GetSessionEmbedUrlWithContext(ctx, input)
}

func (a *QuickSightActivities) ListAnalyses(ctx context.Context, input *quicksight.ListAnalysesInput) (*quicksight.ListAnalysesOutput, error) {
	return a.client.ListAnalysesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListDashboardVersions(ctx context.Context, input *quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error) {
	return a.client.ListDashboardVersionsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListDashboards(ctx context.Context, input *quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error) {
	return a.client.ListDashboardsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListDataSets(ctx context.Context, input *quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error) {
	return a.client.ListDataSetsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListDataSources(ctx context.Context, input *quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error) {
	return a.client.ListDataSourcesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListGroupMemberships(ctx context.Context, input *quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error) {
	return a.client.ListGroupMembershipsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListGroups(ctx context.Context, input *quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error) {
	return a.client.ListGroupsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListIAMPolicyAssignments(ctx context.Context, input *quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
	return a.client.ListIAMPolicyAssignmentsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListIAMPolicyAssignmentsForUser(ctx context.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
	return a.client.ListIAMPolicyAssignmentsForUserWithContext(ctx, input)
}

func (a *QuickSightActivities) ListIngestions(ctx context.Context, input *quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error) {
	return a.client.ListIngestionsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListNamespaces(ctx context.Context, input *quicksight.ListNamespacesInput) (*quicksight.ListNamespacesOutput, error) {
	return a.client.ListNamespacesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListTagsForResource(ctx context.Context, input *quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *QuickSightActivities) ListTemplateAliases(ctx context.Context, input *quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error) {
	return a.client.ListTemplateAliasesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListTemplateVersions(ctx context.Context, input *quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error) {
	return a.client.ListTemplateVersionsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListTemplates(ctx context.Context, input *quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error) {
	return a.client.ListTemplatesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListThemeAliases(ctx context.Context, input *quicksight.ListThemeAliasesInput) (*quicksight.ListThemeAliasesOutput, error) {
	return a.client.ListThemeAliasesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListThemeVersions(ctx context.Context, input *quicksight.ListThemeVersionsInput) (*quicksight.ListThemeVersionsOutput, error) {
	return a.client.ListThemeVersionsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListThemes(ctx context.Context, input *quicksight.ListThemesInput) (*quicksight.ListThemesOutput, error) {
	return a.client.ListThemesWithContext(ctx, input)
}

func (a *QuickSightActivities) ListUserGroups(ctx context.Context, input *quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error) {
	return a.client.ListUserGroupsWithContext(ctx, input)
}

func (a *QuickSightActivities) ListUsers(ctx context.Context, input *quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error) {
	return a.client.ListUsersWithContext(ctx, input)
}

func (a *QuickSightActivities) RegisterUser(ctx context.Context, input *quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error) {
	return a.client.RegisterUserWithContext(ctx, input)
}

func (a *QuickSightActivities) RestoreAnalysis(ctx context.Context, input *quicksight.RestoreAnalysisInput) (*quicksight.RestoreAnalysisOutput, error) {
	return a.client.RestoreAnalysisWithContext(ctx, input)
}

func (a *QuickSightActivities) SearchAnalyses(ctx context.Context, input *quicksight.SearchAnalysesInput) (*quicksight.SearchAnalysesOutput, error) {
	return a.client.SearchAnalysesWithContext(ctx, input)
}

func (a *QuickSightActivities) SearchDashboards(ctx context.Context, input *quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error) {
	return a.client.SearchDashboardsWithContext(ctx, input)
}

func (a *QuickSightActivities) TagResource(ctx context.Context, input *quicksight.TagResourceInput) (*quicksight.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *QuickSightActivities) UntagResource(ctx context.Context, input *quicksight.UntagResourceInput) (*quicksight.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateAccountCustomization(ctx context.Context, input *quicksight.UpdateAccountCustomizationInput) (*quicksight.UpdateAccountCustomizationOutput, error) {
	return a.client.UpdateAccountCustomizationWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateAccountSettings(ctx context.Context, input *quicksight.UpdateAccountSettingsInput) (*quicksight.UpdateAccountSettingsOutput, error) {
	return a.client.UpdateAccountSettingsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateAnalysis(ctx context.Context, input *quicksight.UpdateAnalysisInput) (*quicksight.UpdateAnalysisOutput, error) {
	return a.client.UpdateAnalysisWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateAnalysisPermissions(ctx context.Context, input *quicksight.UpdateAnalysisPermissionsInput) (*quicksight.UpdateAnalysisPermissionsOutput, error) {
	return a.client.UpdateAnalysisPermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDashboard(ctx context.Context, input *quicksight.UpdateDashboardInput) (*quicksight.UpdateDashboardOutput, error) {
	return a.client.UpdateDashboardWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDashboardPermissions(ctx context.Context, input *quicksight.UpdateDashboardPermissionsInput) (*quicksight.UpdateDashboardPermissionsOutput, error) {
	return a.client.UpdateDashboardPermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDashboardPublishedVersion(ctx context.Context, input *quicksight.UpdateDashboardPublishedVersionInput) (*quicksight.UpdateDashboardPublishedVersionOutput, error) {
	return a.client.UpdateDashboardPublishedVersionWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDataSet(ctx context.Context, input *quicksight.UpdateDataSetInput) (*quicksight.UpdateDataSetOutput, error) {
	return a.client.UpdateDataSetWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDataSetPermissions(ctx context.Context, input *quicksight.UpdateDataSetPermissionsInput) (*quicksight.UpdateDataSetPermissionsOutput, error) {
	return a.client.UpdateDataSetPermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDataSource(ctx context.Context, input *quicksight.UpdateDataSourceInput) (*quicksight.UpdateDataSourceOutput, error) {
	return a.client.UpdateDataSourceWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateDataSourcePermissions(ctx context.Context, input *quicksight.UpdateDataSourcePermissionsInput) (*quicksight.UpdateDataSourcePermissionsOutput, error) {
	return a.client.UpdateDataSourcePermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateGroup(ctx context.Context, input *quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error) {
	return a.client.UpdateGroupWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateIAMPolicyAssignment(ctx context.Context, input *quicksight.UpdateIAMPolicyAssignmentInput) (*quicksight.UpdateIAMPolicyAssignmentOutput, error) {
	return a.client.UpdateIAMPolicyAssignmentWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateTemplate(ctx context.Context, input *quicksight.UpdateTemplateInput) (*quicksight.UpdateTemplateOutput, error) {
	return a.client.UpdateTemplateWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateTemplateAlias(ctx context.Context, input *quicksight.UpdateTemplateAliasInput) (*quicksight.UpdateTemplateAliasOutput, error) {
	return a.client.UpdateTemplateAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateTemplatePermissions(ctx context.Context, input *quicksight.UpdateTemplatePermissionsInput) (*quicksight.UpdateTemplatePermissionsOutput, error) {
	return a.client.UpdateTemplatePermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateTheme(ctx context.Context, input *quicksight.UpdateThemeInput) (*quicksight.UpdateThemeOutput, error) {
	return a.client.UpdateThemeWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateThemeAlias(ctx context.Context, input *quicksight.UpdateThemeAliasInput) (*quicksight.UpdateThemeAliasOutput, error) {
	return a.client.UpdateThemeAliasWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateThemePermissions(ctx context.Context, input *quicksight.UpdateThemePermissionsInput) (*quicksight.UpdateThemePermissionsOutput, error) {
	return a.client.UpdateThemePermissionsWithContext(ctx, input)
}

func (a *QuickSightActivities) UpdateUser(ctx context.Context, input *quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error) {
	return a.client.UpdateUserWithContext(ctx, input)
}
