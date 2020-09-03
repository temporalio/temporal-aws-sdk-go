package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/aws/aws-sdk-go/service/quicksight/quicksightiface"
)

type QuickSightActivities struct {
	client quicksightiface.QuickSightAPI
}

func NewQuickSightActivities(client quicksightiface.QuickSightAPI) *QuickSightActivities {
    return &QuickSightActivities{client: client}
}

func (a *QuickSightActivities) CancelIngestion(input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error) {
    return a.client.CancelIngestion(input)
}

func (a *QuickSightActivities) CreateAccountCustomization(input *quicksight.CreateAccountCustomizationInput) (*quicksight.CreateAccountCustomizationOutput, error) {
    return a.client.CreateAccountCustomization(input)
}

func (a *QuickSightActivities) CreateAnalysis(input *quicksight.CreateAnalysisInput) (*quicksight.CreateAnalysisOutput, error) {
    return a.client.CreateAnalysis(input)
}

func (a *QuickSightActivities) CreateDashboard(input *quicksight.CreateDashboardInput) (*quicksight.CreateDashboardOutput, error) {
    return a.client.CreateDashboard(input)
}

func (a *QuickSightActivities) CreateDataSet(input *quicksight.CreateDataSetInput) (*quicksight.CreateDataSetOutput, error) {
    return a.client.CreateDataSet(input)
}

func (a *QuickSightActivities) CreateDataSource(input *quicksight.CreateDataSourceInput) (*quicksight.CreateDataSourceOutput, error) {
    return a.client.CreateDataSource(input)
}

func (a *QuickSightActivities) CreateGroup(input *quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error) {
    return a.client.CreateGroup(input)
}

func (a *QuickSightActivities) CreateGroupMembership(input *quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error) {
    return a.client.CreateGroupMembership(input)
}

func (a *QuickSightActivities) CreateIAMPolicyAssignment(input *quicksight.CreateIAMPolicyAssignmentInput) (*quicksight.CreateIAMPolicyAssignmentOutput, error) {
    return a.client.CreateIAMPolicyAssignment(input)
}

func (a *QuickSightActivities) CreateIngestion(input *quicksight.CreateIngestionInput) (*quicksight.CreateIngestionOutput, error) {
    return a.client.CreateIngestion(input)
}

func (a *QuickSightActivities) CreateNamespace(input *quicksight.CreateNamespaceInput) (*quicksight.CreateNamespaceOutput, error) {
    return a.client.CreateNamespace(input)
}

func (a *QuickSightActivities) CreateTemplate(input *quicksight.CreateTemplateInput) (*quicksight.CreateTemplateOutput, error) {
    return a.client.CreateTemplate(input)
}

func (a *QuickSightActivities) CreateTemplateAlias(input *quicksight.CreateTemplateAliasInput) (*quicksight.CreateTemplateAliasOutput, error) {
    return a.client.CreateTemplateAlias(input)
}

func (a *QuickSightActivities) CreateTheme(input *quicksight.CreateThemeInput) (*quicksight.CreateThemeOutput, error) {
    return a.client.CreateTheme(input)
}

func (a *QuickSightActivities) CreateThemeAlias(input *quicksight.CreateThemeAliasInput) (*quicksight.CreateThemeAliasOutput, error) {
    return a.client.CreateThemeAlias(input)
}

func (a *QuickSightActivities) DeleteAccountCustomization(input *quicksight.DeleteAccountCustomizationInput) (*quicksight.DeleteAccountCustomizationOutput, error) {
    return a.client.DeleteAccountCustomization(input)
}

func (a *QuickSightActivities) DeleteAnalysis(input *quicksight.DeleteAnalysisInput) (*quicksight.DeleteAnalysisOutput, error) {
    return a.client.DeleteAnalysis(input)
}

func (a *QuickSightActivities) DeleteDashboard(input *quicksight.DeleteDashboardInput) (*quicksight.DeleteDashboardOutput, error) {
    return a.client.DeleteDashboard(input)
}

func (a *QuickSightActivities) DeleteDataSet(input *quicksight.DeleteDataSetInput) (*quicksight.DeleteDataSetOutput, error) {
    return a.client.DeleteDataSet(input)
}

func (a *QuickSightActivities) DeleteDataSource(input *quicksight.DeleteDataSourceInput) (*quicksight.DeleteDataSourceOutput, error) {
    return a.client.DeleteDataSource(input)
}

func (a *QuickSightActivities) DeleteGroup(input *quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error) {
    return a.client.DeleteGroup(input)
}

func (a *QuickSightActivities) DeleteGroupMembership(input *quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error) {
    return a.client.DeleteGroupMembership(input)
}

func (a *QuickSightActivities) DeleteIAMPolicyAssignment(input *quicksight.DeleteIAMPolicyAssignmentInput) (*quicksight.DeleteIAMPolicyAssignmentOutput, error) {
    return a.client.DeleteIAMPolicyAssignment(input)
}

func (a *QuickSightActivities) DeleteNamespace(input *quicksight.DeleteNamespaceInput) (*quicksight.DeleteNamespaceOutput, error) {
    return a.client.DeleteNamespace(input)
}

func (a *QuickSightActivities) DeleteTemplate(input *quicksight.DeleteTemplateInput) (*quicksight.DeleteTemplateOutput, error) {
    return a.client.DeleteTemplate(input)
}

func (a *QuickSightActivities) DeleteTemplateAlias(input *quicksight.DeleteTemplateAliasInput) (*quicksight.DeleteTemplateAliasOutput, error) {
    return a.client.DeleteTemplateAlias(input)
}

func (a *QuickSightActivities) DeleteTheme(input *quicksight.DeleteThemeInput) (*quicksight.DeleteThemeOutput, error) {
    return a.client.DeleteTheme(input)
}

func (a *QuickSightActivities) DeleteThemeAlias(input *quicksight.DeleteThemeAliasInput) (*quicksight.DeleteThemeAliasOutput, error) {
    return a.client.DeleteThemeAlias(input)
}

func (a *QuickSightActivities) DeleteUser(input *quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error) {
    return a.client.DeleteUser(input)
}

func (a *QuickSightActivities) DeleteUserByPrincipalId(input *quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error) {
    return a.client.DeleteUserByPrincipalId(input)
}

func (a *QuickSightActivities) DescribeAccountCustomization(input *quicksight.DescribeAccountCustomizationInput) (*quicksight.DescribeAccountCustomizationOutput, error) {
    return a.client.DescribeAccountCustomization(input)
}

func (a *QuickSightActivities) DescribeAccountSettings(input *quicksight.DescribeAccountSettingsInput) (*quicksight.DescribeAccountSettingsOutput, error) {
    return a.client.DescribeAccountSettings(input)
}

func (a *QuickSightActivities) DescribeAnalysis(input *quicksight.DescribeAnalysisInput) (*quicksight.DescribeAnalysisOutput, error) {
    return a.client.DescribeAnalysis(input)
}

func (a *QuickSightActivities) DescribeAnalysisPermissions(input *quicksight.DescribeAnalysisPermissionsInput) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
    return a.client.DescribeAnalysisPermissions(input)
}

func (a *QuickSightActivities) DescribeDashboard(input *quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error) {
    return a.client.DescribeDashboard(input)
}

func (a *QuickSightActivities) DescribeDashboardPermissions(input *quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error) {
    return a.client.DescribeDashboardPermissions(input)
}

func (a *QuickSightActivities) DescribeDataSet(input *quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error) {
    return a.client.DescribeDataSet(input)
}

func (a *QuickSightActivities) DescribeDataSetPermissions(input *quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error) {
    return a.client.DescribeDataSetPermissions(input)
}

func (a *QuickSightActivities) DescribeDataSource(input *quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error) {
    return a.client.DescribeDataSource(input)
}

func (a *QuickSightActivities) DescribeDataSourcePermissions(input *quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
    return a.client.DescribeDataSourcePermissions(input)
}

func (a *QuickSightActivities) DescribeGroup(input *quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error) {
    return a.client.DescribeGroup(input)
}

func (a *QuickSightActivities) DescribeIAMPolicyAssignment(input *quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
    return a.client.DescribeIAMPolicyAssignment(input)
}

func (a *QuickSightActivities) DescribeIngestion(input *quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error) {
    return a.client.DescribeIngestion(input)
}

func (a *QuickSightActivities) DescribeNamespace(input *quicksight.DescribeNamespaceInput) (*quicksight.DescribeNamespaceOutput, error) {
    return a.client.DescribeNamespace(input)
}

func (a *QuickSightActivities) DescribeTemplate(input *quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error) {
    return a.client.DescribeTemplate(input)
}

func (a *QuickSightActivities) DescribeTemplateAlias(input *quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error) {
    return a.client.DescribeTemplateAlias(input)
}

func (a *QuickSightActivities) DescribeTemplatePermissions(input *quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error) {
    return a.client.DescribeTemplatePermissions(input)
}

func (a *QuickSightActivities) DescribeTheme(input *quicksight.DescribeThemeInput) (*quicksight.DescribeThemeOutput, error) {
    return a.client.DescribeTheme(input)
}

func (a *QuickSightActivities) DescribeThemeAlias(input *quicksight.DescribeThemeAliasInput) (*quicksight.DescribeThemeAliasOutput, error) {
    return a.client.DescribeThemeAlias(input)
}

func (a *QuickSightActivities) DescribeThemePermissions(input *quicksight.DescribeThemePermissionsInput) (*quicksight.DescribeThemePermissionsOutput, error) {
    return a.client.DescribeThemePermissions(input)
}

func (a *QuickSightActivities) DescribeUser(input *quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error) {
    return a.client.DescribeUser(input)
}

func (a *QuickSightActivities) GetDashboardEmbedUrl(input *quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error) {
    return a.client.GetDashboardEmbedUrl(input)
}

func (a *QuickSightActivities) GetSessionEmbedUrl(input *quicksight.GetSessionEmbedUrlInput) (*quicksight.GetSessionEmbedUrlOutput, error) {
    return a.client.GetSessionEmbedUrl(input)
}

func (a *QuickSightActivities) ListAnalyses(input *quicksight.ListAnalysesInput) (*quicksight.ListAnalysesOutput, error) {
    return a.client.ListAnalyses(input)
}

func (a *QuickSightActivities) ListDashboardVersions(input *quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error) {
    return a.client.ListDashboardVersions(input)
}

func (a *QuickSightActivities) ListDashboards(input *quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error) {
    return a.client.ListDashboards(input)
}

func (a *QuickSightActivities) ListDataSets(input *quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error) {
    return a.client.ListDataSets(input)
}

func (a *QuickSightActivities) ListDataSources(input *quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error) {
    return a.client.ListDataSources(input)
}

func (a *QuickSightActivities) ListGroupMemberships(input *quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error) {
    return a.client.ListGroupMemberships(input)
}

func (a *QuickSightActivities) ListGroups(input *quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error) {
    return a.client.ListGroups(input)
}

func (a *QuickSightActivities) ListIAMPolicyAssignments(input *quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
    return a.client.ListIAMPolicyAssignments(input)
}

func (a *QuickSightActivities) ListIAMPolicyAssignmentsForUser(input *quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
    return a.client.ListIAMPolicyAssignmentsForUser(input)
}

func (a *QuickSightActivities) ListIngestions(input *quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error) {
    return a.client.ListIngestions(input)
}

func (a *QuickSightActivities) ListNamespaces(input *quicksight.ListNamespacesInput) (*quicksight.ListNamespacesOutput, error) {
    return a.client.ListNamespaces(input)
}

func (a *QuickSightActivities) ListTagsForResource(input *quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *QuickSightActivities) ListTemplateAliases(input *quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error) {
    return a.client.ListTemplateAliases(input)
}

func (a *QuickSightActivities) ListTemplateVersions(input *quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error) {
    return a.client.ListTemplateVersions(input)
}

func (a *QuickSightActivities) ListTemplates(input *quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error) {
    return a.client.ListTemplates(input)
}

func (a *QuickSightActivities) ListThemeAliases(input *quicksight.ListThemeAliasesInput) (*quicksight.ListThemeAliasesOutput, error) {
    return a.client.ListThemeAliases(input)
}

func (a *QuickSightActivities) ListThemeVersions(input *quicksight.ListThemeVersionsInput) (*quicksight.ListThemeVersionsOutput, error) {
    return a.client.ListThemeVersions(input)
}

func (a *QuickSightActivities) ListThemes(input *quicksight.ListThemesInput) (*quicksight.ListThemesOutput, error) {
    return a.client.ListThemes(input)
}

func (a *QuickSightActivities) ListUserGroups(input *quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error) {
    return a.client.ListUserGroups(input)
}

func (a *QuickSightActivities) ListUsers(input *quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error) {
    return a.client.ListUsers(input)
}

func (a *QuickSightActivities) RegisterUser(input *quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error) {
    return a.client.RegisterUser(input)
}

func (a *QuickSightActivities) RestoreAnalysis(input *quicksight.RestoreAnalysisInput) (*quicksight.RestoreAnalysisOutput, error) {
    return a.client.RestoreAnalysis(input)
}

func (a *QuickSightActivities) SearchAnalyses(input *quicksight.SearchAnalysesInput) (*quicksight.SearchAnalysesOutput, error) {
    return a.client.SearchAnalyses(input)
}

func (a *QuickSightActivities) SearchDashboards(input *quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error) {
    return a.client.SearchDashboards(input)
}

func (a *QuickSightActivities) TagResource(input *quicksight.TagResourceInput) (*quicksight.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *QuickSightActivities) UntagResource(input *quicksight.UntagResourceInput) (*quicksight.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *QuickSightActivities) UpdateAccountCustomization(input *quicksight.UpdateAccountCustomizationInput) (*quicksight.UpdateAccountCustomizationOutput, error) {
    return a.client.UpdateAccountCustomization(input)
}

func (a *QuickSightActivities) UpdateAccountSettings(input *quicksight.UpdateAccountSettingsInput) (*quicksight.UpdateAccountSettingsOutput, error) {
    return a.client.UpdateAccountSettings(input)
}

func (a *QuickSightActivities) UpdateAnalysis(input *quicksight.UpdateAnalysisInput) (*quicksight.UpdateAnalysisOutput, error) {
    return a.client.UpdateAnalysis(input)
}

func (a *QuickSightActivities) UpdateAnalysisPermissions(input *quicksight.UpdateAnalysisPermissionsInput) (*quicksight.UpdateAnalysisPermissionsOutput, error) {
    return a.client.UpdateAnalysisPermissions(input)
}

func (a *QuickSightActivities) UpdateDashboard(input *quicksight.UpdateDashboardInput) (*quicksight.UpdateDashboardOutput, error) {
    return a.client.UpdateDashboard(input)
}

func (a *QuickSightActivities) UpdateDashboardPermissions(input *quicksight.UpdateDashboardPermissionsInput) (*quicksight.UpdateDashboardPermissionsOutput, error) {
    return a.client.UpdateDashboardPermissions(input)
}

func (a *QuickSightActivities) UpdateDashboardPublishedVersion(input *quicksight.UpdateDashboardPublishedVersionInput) (*quicksight.UpdateDashboardPublishedVersionOutput, error) {
    return a.client.UpdateDashboardPublishedVersion(input)
}

func (a *QuickSightActivities) UpdateDataSet(input *quicksight.UpdateDataSetInput) (*quicksight.UpdateDataSetOutput, error) {
    return a.client.UpdateDataSet(input)
}

func (a *QuickSightActivities) UpdateDataSetPermissions(input *quicksight.UpdateDataSetPermissionsInput) (*quicksight.UpdateDataSetPermissionsOutput, error) {
    return a.client.UpdateDataSetPermissions(input)
}

func (a *QuickSightActivities) UpdateDataSource(input *quicksight.UpdateDataSourceInput) (*quicksight.UpdateDataSourceOutput, error) {
    return a.client.UpdateDataSource(input)
}

func (a *QuickSightActivities) UpdateDataSourcePermissions(input *quicksight.UpdateDataSourcePermissionsInput) (*quicksight.UpdateDataSourcePermissionsOutput, error) {
    return a.client.UpdateDataSourcePermissions(input)
}

func (a *QuickSightActivities) UpdateGroup(input *quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error) {
    return a.client.UpdateGroup(input)
}

func (a *QuickSightActivities) UpdateIAMPolicyAssignment(input *quicksight.UpdateIAMPolicyAssignmentInput) (*quicksight.UpdateIAMPolicyAssignmentOutput, error) {
    return a.client.UpdateIAMPolicyAssignment(input)
}

func (a *QuickSightActivities) UpdateTemplate(input *quicksight.UpdateTemplateInput) (*quicksight.UpdateTemplateOutput, error) {
    return a.client.UpdateTemplate(input)
}

func (a *QuickSightActivities) UpdateTemplateAlias(input *quicksight.UpdateTemplateAliasInput) (*quicksight.UpdateTemplateAliasOutput, error) {
    return a.client.UpdateTemplateAlias(input)
}

func (a *QuickSightActivities) UpdateTemplatePermissions(input *quicksight.UpdateTemplatePermissionsInput) (*quicksight.UpdateTemplatePermissionsOutput, error) {
    return a.client.UpdateTemplatePermissions(input)
}

func (a *QuickSightActivities) UpdateTheme(input *quicksight.UpdateThemeInput) (*quicksight.UpdateThemeOutput, error) {
    return a.client.UpdateTheme(input)
}

func (a *QuickSightActivities) UpdateThemeAlias(input *quicksight.UpdateThemeAliasInput) (*quicksight.UpdateThemeAliasOutput, error) {
    return a.client.UpdateThemeAlias(input)
}

func (a *QuickSightActivities) UpdateThemePermissions(input *quicksight.UpdateThemePermissionsInput) (*quicksight.UpdateThemePermissionsOutput, error) {
    return a.client.UpdateThemePermissions(input)
}

func (a *QuickSightActivities) UpdateUser(input *quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error) {
    return a.client.UpdateUser(input)
}
