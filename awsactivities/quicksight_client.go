package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/quicksight"
	"go.temporal.io/sdk/workflow"
)

type QuickSightClient interface {
    CancelIngestion(ctx workflow.Context, input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error)
    CancelIngestionAsync(ctx workflow.Context, input *quicksight.CancelIngestionInput) *QuicksightCancelIngestionResult

    CreateAccountCustomization(ctx workflow.Context, input *quicksight.CreateAccountCustomizationInput) (*quicksight.CreateAccountCustomizationOutput, error)
    CreateAccountCustomizationAsync(ctx workflow.Context, input *quicksight.CreateAccountCustomizationInput) *QuicksightCreateAccountCustomizationResult

    CreateAnalysis(ctx workflow.Context, input *quicksight.CreateAnalysisInput) (*quicksight.CreateAnalysisOutput, error)
    CreateAnalysisAsync(ctx workflow.Context, input *quicksight.CreateAnalysisInput) *QuicksightCreateAnalysisResult

    CreateDashboard(ctx workflow.Context, input *quicksight.CreateDashboardInput) (*quicksight.CreateDashboardOutput, error)
    CreateDashboardAsync(ctx workflow.Context, input *quicksight.CreateDashboardInput) *QuicksightCreateDashboardResult

    CreateDataSet(ctx workflow.Context, input *quicksight.CreateDataSetInput) (*quicksight.CreateDataSetOutput, error)
    CreateDataSetAsync(ctx workflow.Context, input *quicksight.CreateDataSetInput) *QuicksightCreateDataSetResult

    CreateDataSource(ctx workflow.Context, input *quicksight.CreateDataSourceInput) (*quicksight.CreateDataSourceOutput, error)
    CreateDataSourceAsync(ctx workflow.Context, input *quicksight.CreateDataSourceInput) *QuicksightCreateDataSourceResult

    CreateGroup(ctx workflow.Context, input *quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error)
    CreateGroupAsync(ctx workflow.Context, input *quicksight.CreateGroupInput) *QuicksightCreateGroupResult

    CreateGroupMembership(ctx workflow.Context, input *quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error)
    CreateGroupMembershipAsync(ctx workflow.Context, input *quicksight.CreateGroupMembershipInput) *QuicksightCreateGroupMembershipResult

    CreateIAMPolicyAssignment(ctx workflow.Context, input *quicksight.CreateIAMPolicyAssignmentInput) (*quicksight.CreateIAMPolicyAssignmentOutput, error)
    CreateIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.CreateIAMPolicyAssignmentInput) *QuicksightCreateIAMPolicyAssignmentResult

    CreateIngestion(ctx workflow.Context, input *quicksight.CreateIngestionInput) (*quicksight.CreateIngestionOutput, error)
    CreateIngestionAsync(ctx workflow.Context, input *quicksight.CreateIngestionInput) *QuicksightCreateIngestionResult

    CreateNamespace(ctx workflow.Context, input *quicksight.CreateNamespaceInput) (*quicksight.CreateNamespaceOutput, error)
    CreateNamespaceAsync(ctx workflow.Context, input *quicksight.CreateNamespaceInput) *QuicksightCreateNamespaceResult

    CreateTemplate(ctx workflow.Context, input *quicksight.CreateTemplateInput) (*quicksight.CreateTemplateOutput, error)
    CreateTemplateAsync(ctx workflow.Context, input *quicksight.CreateTemplateInput) *QuicksightCreateTemplateResult

    CreateTemplateAlias(ctx workflow.Context, input *quicksight.CreateTemplateAliasInput) (*quicksight.CreateTemplateAliasOutput, error)
    CreateTemplateAliasAsync(ctx workflow.Context, input *quicksight.CreateTemplateAliasInput) *QuicksightCreateTemplateAliasResult

    CreateTheme(ctx workflow.Context, input *quicksight.CreateThemeInput) (*quicksight.CreateThemeOutput, error)
    CreateThemeAsync(ctx workflow.Context, input *quicksight.CreateThemeInput) *QuicksightCreateThemeResult

    CreateThemeAlias(ctx workflow.Context, input *quicksight.CreateThemeAliasInput) (*quicksight.CreateThemeAliasOutput, error)
    CreateThemeAliasAsync(ctx workflow.Context, input *quicksight.CreateThemeAliasInput) *QuicksightCreateThemeAliasResult

    DeleteAccountCustomization(ctx workflow.Context, input *quicksight.DeleteAccountCustomizationInput) (*quicksight.DeleteAccountCustomizationOutput, error)
    DeleteAccountCustomizationAsync(ctx workflow.Context, input *quicksight.DeleteAccountCustomizationInput) *QuicksightDeleteAccountCustomizationResult

    DeleteAnalysis(ctx workflow.Context, input *quicksight.DeleteAnalysisInput) (*quicksight.DeleteAnalysisOutput, error)
    DeleteAnalysisAsync(ctx workflow.Context, input *quicksight.DeleteAnalysisInput) *QuicksightDeleteAnalysisResult

    DeleteDashboard(ctx workflow.Context, input *quicksight.DeleteDashboardInput) (*quicksight.DeleteDashboardOutput, error)
    DeleteDashboardAsync(ctx workflow.Context, input *quicksight.DeleteDashboardInput) *QuicksightDeleteDashboardResult

    DeleteDataSet(ctx workflow.Context, input *quicksight.DeleteDataSetInput) (*quicksight.DeleteDataSetOutput, error)
    DeleteDataSetAsync(ctx workflow.Context, input *quicksight.DeleteDataSetInput) *QuicksightDeleteDataSetResult

    DeleteDataSource(ctx workflow.Context, input *quicksight.DeleteDataSourceInput) (*quicksight.DeleteDataSourceOutput, error)
    DeleteDataSourceAsync(ctx workflow.Context, input *quicksight.DeleteDataSourceInput) *QuicksightDeleteDataSourceResult

    DeleteGroup(ctx workflow.Context, input *quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error)
    DeleteGroupAsync(ctx workflow.Context, input *quicksight.DeleteGroupInput) *QuicksightDeleteGroupResult

    DeleteGroupMembership(ctx workflow.Context, input *quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error)
    DeleteGroupMembershipAsync(ctx workflow.Context, input *quicksight.DeleteGroupMembershipInput) *QuicksightDeleteGroupMembershipResult

    DeleteIAMPolicyAssignment(ctx workflow.Context, input *quicksight.DeleteIAMPolicyAssignmentInput) (*quicksight.DeleteIAMPolicyAssignmentOutput, error)
    DeleteIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.DeleteIAMPolicyAssignmentInput) *QuicksightDeleteIAMPolicyAssignmentResult

    DeleteNamespace(ctx workflow.Context, input *quicksight.DeleteNamespaceInput) (*quicksight.DeleteNamespaceOutput, error)
    DeleteNamespaceAsync(ctx workflow.Context, input *quicksight.DeleteNamespaceInput) *QuicksightDeleteNamespaceResult

    DeleteTemplate(ctx workflow.Context, input *quicksight.DeleteTemplateInput) (*quicksight.DeleteTemplateOutput, error)
    DeleteTemplateAsync(ctx workflow.Context, input *quicksight.DeleteTemplateInput) *QuicksightDeleteTemplateResult

    DeleteTemplateAlias(ctx workflow.Context, input *quicksight.DeleteTemplateAliasInput) (*quicksight.DeleteTemplateAliasOutput, error)
    DeleteTemplateAliasAsync(ctx workflow.Context, input *quicksight.DeleteTemplateAliasInput) *QuicksightDeleteTemplateAliasResult

    DeleteTheme(ctx workflow.Context, input *quicksight.DeleteThemeInput) (*quicksight.DeleteThemeOutput, error)
    DeleteThemeAsync(ctx workflow.Context, input *quicksight.DeleteThemeInput) *QuicksightDeleteThemeResult

    DeleteThemeAlias(ctx workflow.Context, input *quicksight.DeleteThemeAliasInput) (*quicksight.DeleteThemeAliasOutput, error)
    DeleteThemeAliasAsync(ctx workflow.Context, input *quicksight.DeleteThemeAliasInput) *QuicksightDeleteThemeAliasResult

    DeleteUser(ctx workflow.Context, input *quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error)
    DeleteUserAsync(ctx workflow.Context, input *quicksight.DeleteUserInput) *QuicksightDeleteUserResult

    DeleteUserByPrincipalId(ctx workflow.Context, input *quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error)
    DeleteUserByPrincipalIdAsync(ctx workflow.Context, input *quicksight.DeleteUserByPrincipalIdInput) *QuicksightDeleteUserByPrincipalIdResult

    DescribeAccountCustomization(ctx workflow.Context, input *quicksight.DescribeAccountCustomizationInput) (*quicksight.DescribeAccountCustomizationOutput, error)
    DescribeAccountCustomizationAsync(ctx workflow.Context, input *quicksight.DescribeAccountCustomizationInput) *QuicksightDescribeAccountCustomizationResult

    DescribeAccountSettings(ctx workflow.Context, input *quicksight.DescribeAccountSettingsInput) (*quicksight.DescribeAccountSettingsOutput, error)
    DescribeAccountSettingsAsync(ctx workflow.Context, input *quicksight.DescribeAccountSettingsInput) *QuicksightDescribeAccountSettingsResult

    DescribeAnalysis(ctx workflow.Context, input *quicksight.DescribeAnalysisInput) (*quicksight.DescribeAnalysisOutput, error)
    DescribeAnalysisAsync(ctx workflow.Context, input *quicksight.DescribeAnalysisInput) *QuicksightDescribeAnalysisResult

    DescribeAnalysisPermissions(ctx workflow.Context, input *quicksight.DescribeAnalysisPermissionsInput) (*quicksight.DescribeAnalysisPermissionsOutput, error)
    DescribeAnalysisPermissionsAsync(ctx workflow.Context, input *quicksight.DescribeAnalysisPermissionsInput) *QuicksightDescribeAnalysisPermissionsResult

    DescribeDashboard(ctx workflow.Context, input *quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error)
    DescribeDashboardAsync(ctx workflow.Context, input *quicksight.DescribeDashboardInput) *QuicksightDescribeDashboardResult

    DescribeDashboardPermissions(ctx workflow.Context, input *quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error)
    DescribeDashboardPermissionsAsync(ctx workflow.Context, input *quicksight.DescribeDashboardPermissionsInput) *QuicksightDescribeDashboardPermissionsResult

    DescribeDataSet(ctx workflow.Context, input *quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error)
    DescribeDataSetAsync(ctx workflow.Context, input *quicksight.DescribeDataSetInput) *QuicksightDescribeDataSetResult

    DescribeDataSetPermissions(ctx workflow.Context, input *quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error)
    DescribeDataSetPermissionsAsync(ctx workflow.Context, input *quicksight.DescribeDataSetPermissionsInput) *QuicksightDescribeDataSetPermissionsResult

    DescribeDataSource(ctx workflow.Context, input *quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error)
    DescribeDataSourceAsync(ctx workflow.Context, input *quicksight.DescribeDataSourceInput) *QuicksightDescribeDataSourceResult

    DescribeDataSourcePermissions(ctx workflow.Context, input *quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error)
    DescribeDataSourcePermissionsAsync(ctx workflow.Context, input *quicksight.DescribeDataSourcePermissionsInput) *QuicksightDescribeDataSourcePermissionsResult

    DescribeGroup(ctx workflow.Context, input *quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error)
    DescribeGroupAsync(ctx workflow.Context, input *quicksight.DescribeGroupInput) *QuicksightDescribeGroupResult

    DescribeIAMPolicyAssignment(ctx workflow.Context, input *quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error)
    DescribeIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.DescribeIAMPolicyAssignmentInput) *QuicksightDescribeIAMPolicyAssignmentResult

    DescribeIngestion(ctx workflow.Context, input *quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error)
    DescribeIngestionAsync(ctx workflow.Context, input *quicksight.DescribeIngestionInput) *QuicksightDescribeIngestionResult

    DescribeNamespace(ctx workflow.Context, input *quicksight.DescribeNamespaceInput) (*quicksight.DescribeNamespaceOutput, error)
    DescribeNamespaceAsync(ctx workflow.Context, input *quicksight.DescribeNamespaceInput) *QuicksightDescribeNamespaceResult

    DescribeTemplate(ctx workflow.Context, input *quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error)
    DescribeTemplateAsync(ctx workflow.Context, input *quicksight.DescribeTemplateInput) *QuicksightDescribeTemplateResult

    DescribeTemplateAlias(ctx workflow.Context, input *quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error)
    DescribeTemplateAliasAsync(ctx workflow.Context, input *quicksight.DescribeTemplateAliasInput) *QuicksightDescribeTemplateAliasResult

    DescribeTemplatePermissions(ctx workflow.Context, input *quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error)
    DescribeTemplatePermissionsAsync(ctx workflow.Context, input *quicksight.DescribeTemplatePermissionsInput) *QuicksightDescribeTemplatePermissionsResult

    DescribeTheme(ctx workflow.Context, input *quicksight.DescribeThemeInput) (*quicksight.DescribeThemeOutput, error)
    DescribeThemeAsync(ctx workflow.Context, input *quicksight.DescribeThemeInput) *QuicksightDescribeThemeResult

    DescribeThemeAlias(ctx workflow.Context, input *quicksight.DescribeThemeAliasInput) (*quicksight.DescribeThemeAliasOutput, error)
    DescribeThemeAliasAsync(ctx workflow.Context, input *quicksight.DescribeThemeAliasInput) *QuicksightDescribeThemeAliasResult

    DescribeThemePermissions(ctx workflow.Context, input *quicksight.DescribeThemePermissionsInput) (*quicksight.DescribeThemePermissionsOutput, error)
    DescribeThemePermissionsAsync(ctx workflow.Context, input *quicksight.DescribeThemePermissionsInput) *QuicksightDescribeThemePermissionsResult

    DescribeUser(ctx workflow.Context, input *quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error)
    DescribeUserAsync(ctx workflow.Context, input *quicksight.DescribeUserInput) *QuicksightDescribeUserResult

    GetDashboardEmbedUrl(ctx workflow.Context, input *quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error)
    GetDashboardEmbedUrlAsync(ctx workflow.Context, input *quicksight.GetDashboardEmbedUrlInput) *QuicksightGetDashboardEmbedUrlResult

    GetSessionEmbedUrl(ctx workflow.Context, input *quicksight.GetSessionEmbedUrlInput) (*quicksight.GetSessionEmbedUrlOutput, error)
    GetSessionEmbedUrlAsync(ctx workflow.Context, input *quicksight.GetSessionEmbedUrlInput) *QuicksightGetSessionEmbedUrlResult

    ListAnalyses(ctx workflow.Context, input *quicksight.ListAnalysesInput) (*quicksight.ListAnalysesOutput, error)
    ListAnalysesAsync(ctx workflow.Context, input *quicksight.ListAnalysesInput) *QuicksightListAnalysesResult

    ListDashboardVersions(ctx workflow.Context, input *quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error)
    ListDashboardVersionsAsync(ctx workflow.Context, input *quicksight.ListDashboardVersionsInput) *QuicksightListDashboardVersionsResult

    ListDashboards(ctx workflow.Context, input *quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error)
    ListDashboardsAsync(ctx workflow.Context, input *quicksight.ListDashboardsInput) *QuicksightListDashboardsResult

    ListDataSets(ctx workflow.Context, input *quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error)
    ListDataSetsAsync(ctx workflow.Context, input *quicksight.ListDataSetsInput) *QuicksightListDataSetsResult

    ListDataSources(ctx workflow.Context, input *quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error)
    ListDataSourcesAsync(ctx workflow.Context, input *quicksight.ListDataSourcesInput) *QuicksightListDataSourcesResult

    ListGroupMemberships(ctx workflow.Context, input *quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error)
    ListGroupMembershipsAsync(ctx workflow.Context, input *quicksight.ListGroupMembershipsInput) *QuicksightListGroupMembershipsResult

    ListGroups(ctx workflow.Context, input *quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error)
    ListGroupsAsync(ctx workflow.Context, input *quicksight.ListGroupsInput) *QuicksightListGroupsResult

    ListIAMPolicyAssignments(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error)
    ListIAMPolicyAssignmentsAsync(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsInput) *QuicksightListIAMPolicyAssignmentsResult

    ListIAMPolicyAssignmentsForUser(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error)
    ListIAMPolicyAssignmentsForUserAsync(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput) *QuicksightListIAMPolicyAssignmentsForUserResult

    ListIngestions(ctx workflow.Context, input *quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error)
    ListIngestionsAsync(ctx workflow.Context, input *quicksight.ListIngestionsInput) *QuicksightListIngestionsResult

    ListNamespaces(ctx workflow.Context, input *quicksight.ListNamespacesInput) (*quicksight.ListNamespacesOutput, error)
    ListNamespacesAsync(ctx workflow.Context, input *quicksight.ListNamespacesInput) *QuicksightListNamespacesResult

    ListTagsForResource(ctx workflow.Context, input *quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *quicksight.ListTagsForResourceInput) *QuicksightListTagsForResourceResult

    ListTemplateAliases(ctx workflow.Context, input *quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error)
    ListTemplateAliasesAsync(ctx workflow.Context, input *quicksight.ListTemplateAliasesInput) *QuicksightListTemplateAliasesResult

    ListTemplateVersions(ctx workflow.Context, input *quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error)
    ListTemplateVersionsAsync(ctx workflow.Context, input *quicksight.ListTemplateVersionsInput) *QuicksightListTemplateVersionsResult

    ListTemplates(ctx workflow.Context, input *quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error)
    ListTemplatesAsync(ctx workflow.Context, input *quicksight.ListTemplatesInput) *QuicksightListTemplatesResult

    ListThemeAliases(ctx workflow.Context, input *quicksight.ListThemeAliasesInput) (*quicksight.ListThemeAliasesOutput, error)
    ListThemeAliasesAsync(ctx workflow.Context, input *quicksight.ListThemeAliasesInput) *QuicksightListThemeAliasesResult

    ListThemeVersions(ctx workflow.Context, input *quicksight.ListThemeVersionsInput) (*quicksight.ListThemeVersionsOutput, error)
    ListThemeVersionsAsync(ctx workflow.Context, input *quicksight.ListThemeVersionsInput) *QuicksightListThemeVersionsResult

    ListThemes(ctx workflow.Context, input *quicksight.ListThemesInput) (*quicksight.ListThemesOutput, error)
    ListThemesAsync(ctx workflow.Context, input *quicksight.ListThemesInput) *QuicksightListThemesResult

    ListUserGroups(ctx workflow.Context, input *quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error)
    ListUserGroupsAsync(ctx workflow.Context, input *quicksight.ListUserGroupsInput) *QuicksightListUserGroupsResult

    ListUsers(ctx workflow.Context, input *quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error)
    ListUsersAsync(ctx workflow.Context, input *quicksight.ListUsersInput) *QuicksightListUsersResult

    RegisterUser(ctx workflow.Context, input *quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error)
    RegisterUserAsync(ctx workflow.Context, input *quicksight.RegisterUserInput) *QuicksightRegisterUserResult

    RestoreAnalysis(ctx workflow.Context, input *quicksight.RestoreAnalysisInput) (*quicksight.RestoreAnalysisOutput, error)
    RestoreAnalysisAsync(ctx workflow.Context, input *quicksight.RestoreAnalysisInput) *QuicksightRestoreAnalysisResult

    SearchAnalyses(ctx workflow.Context, input *quicksight.SearchAnalysesInput) (*quicksight.SearchAnalysesOutput, error)
    SearchAnalysesAsync(ctx workflow.Context, input *quicksight.SearchAnalysesInput) *QuicksightSearchAnalysesResult

    SearchDashboards(ctx workflow.Context, input *quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error)
    SearchDashboardsAsync(ctx workflow.Context, input *quicksight.SearchDashboardsInput) *QuicksightSearchDashboardsResult

    TagResource(ctx workflow.Context, input *quicksight.TagResourceInput) (*quicksight.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *quicksight.TagResourceInput) *QuicksightTagResourceResult

    UntagResource(ctx workflow.Context, input *quicksight.UntagResourceInput) (*quicksight.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *quicksight.UntagResourceInput) *QuicksightUntagResourceResult

    UpdateAccountCustomization(ctx workflow.Context, input *quicksight.UpdateAccountCustomizationInput) (*quicksight.UpdateAccountCustomizationOutput, error)
    UpdateAccountCustomizationAsync(ctx workflow.Context, input *quicksight.UpdateAccountCustomizationInput) *QuicksightUpdateAccountCustomizationResult

    UpdateAccountSettings(ctx workflow.Context, input *quicksight.UpdateAccountSettingsInput) (*quicksight.UpdateAccountSettingsOutput, error)
    UpdateAccountSettingsAsync(ctx workflow.Context, input *quicksight.UpdateAccountSettingsInput) *QuicksightUpdateAccountSettingsResult

    UpdateAnalysis(ctx workflow.Context, input *quicksight.UpdateAnalysisInput) (*quicksight.UpdateAnalysisOutput, error)
    UpdateAnalysisAsync(ctx workflow.Context, input *quicksight.UpdateAnalysisInput) *QuicksightUpdateAnalysisResult

    UpdateAnalysisPermissions(ctx workflow.Context, input *quicksight.UpdateAnalysisPermissionsInput) (*quicksight.UpdateAnalysisPermissionsOutput, error)
    UpdateAnalysisPermissionsAsync(ctx workflow.Context, input *quicksight.UpdateAnalysisPermissionsInput) *QuicksightUpdateAnalysisPermissionsResult

    UpdateDashboard(ctx workflow.Context, input *quicksight.UpdateDashboardInput) (*quicksight.UpdateDashboardOutput, error)
    UpdateDashboardAsync(ctx workflow.Context, input *quicksight.UpdateDashboardInput) *QuicksightUpdateDashboardResult

    UpdateDashboardPermissions(ctx workflow.Context, input *quicksight.UpdateDashboardPermissionsInput) (*quicksight.UpdateDashboardPermissionsOutput, error)
    UpdateDashboardPermissionsAsync(ctx workflow.Context, input *quicksight.UpdateDashboardPermissionsInput) *QuicksightUpdateDashboardPermissionsResult

    UpdateDashboardPublishedVersion(ctx workflow.Context, input *quicksight.UpdateDashboardPublishedVersionInput) (*quicksight.UpdateDashboardPublishedVersionOutput, error)
    UpdateDashboardPublishedVersionAsync(ctx workflow.Context, input *quicksight.UpdateDashboardPublishedVersionInput) *QuicksightUpdateDashboardPublishedVersionResult

    UpdateDataSet(ctx workflow.Context, input *quicksight.UpdateDataSetInput) (*quicksight.UpdateDataSetOutput, error)
    UpdateDataSetAsync(ctx workflow.Context, input *quicksight.UpdateDataSetInput) *QuicksightUpdateDataSetResult

    UpdateDataSetPermissions(ctx workflow.Context, input *quicksight.UpdateDataSetPermissionsInput) (*quicksight.UpdateDataSetPermissionsOutput, error)
    UpdateDataSetPermissionsAsync(ctx workflow.Context, input *quicksight.UpdateDataSetPermissionsInput) *QuicksightUpdateDataSetPermissionsResult

    UpdateDataSource(ctx workflow.Context, input *quicksight.UpdateDataSourceInput) (*quicksight.UpdateDataSourceOutput, error)
    UpdateDataSourceAsync(ctx workflow.Context, input *quicksight.UpdateDataSourceInput) *QuicksightUpdateDataSourceResult

    UpdateDataSourcePermissions(ctx workflow.Context, input *quicksight.UpdateDataSourcePermissionsInput) (*quicksight.UpdateDataSourcePermissionsOutput, error)
    UpdateDataSourcePermissionsAsync(ctx workflow.Context, input *quicksight.UpdateDataSourcePermissionsInput) *QuicksightUpdateDataSourcePermissionsResult

    UpdateGroup(ctx workflow.Context, input *quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error)
    UpdateGroupAsync(ctx workflow.Context, input *quicksight.UpdateGroupInput) *QuicksightUpdateGroupResult

    UpdateIAMPolicyAssignment(ctx workflow.Context, input *quicksight.UpdateIAMPolicyAssignmentInput) (*quicksight.UpdateIAMPolicyAssignmentOutput, error)
    UpdateIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.UpdateIAMPolicyAssignmentInput) *QuicksightUpdateIAMPolicyAssignmentResult

    UpdateTemplate(ctx workflow.Context, input *quicksight.UpdateTemplateInput) (*quicksight.UpdateTemplateOutput, error)
    UpdateTemplateAsync(ctx workflow.Context, input *quicksight.UpdateTemplateInput) *QuicksightUpdateTemplateResult

    UpdateTemplateAlias(ctx workflow.Context, input *quicksight.UpdateTemplateAliasInput) (*quicksight.UpdateTemplateAliasOutput, error)
    UpdateTemplateAliasAsync(ctx workflow.Context, input *quicksight.UpdateTemplateAliasInput) *QuicksightUpdateTemplateAliasResult

    UpdateTemplatePermissions(ctx workflow.Context, input *quicksight.UpdateTemplatePermissionsInput) (*quicksight.UpdateTemplatePermissionsOutput, error)
    UpdateTemplatePermissionsAsync(ctx workflow.Context, input *quicksight.UpdateTemplatePermissionsInput) *QuicksightUpdateTemplatePermissionsResult

    UpdateTheme(ctx workflow.Context, input *quicksight.UpdateThemeInput) (*quicksight.UpdateThemeOutput, error)
    UpdateThemeAsync(ctx workflow.Context, input *quicksight.UpdateThemeInput) *QuicksightUpdateThemeResult

    UpdateThemeAlias(ctx workflow.Context, input *quicksight.UpdateThemeAliasInput) (*quicksight.UpdateThemeAliasOutput, error)
    UpdateThemeAliasAsync(ctx workflow.Context, input *quicksight.UpdateThemeAliasInput) *QuicksightUpdateThemeAliasResult

    UpdateThemePermissions(ctx workflow.Context, input *quicksight.UpdateThemePermissionsInput) (*quicksight.UpdateThemePermissionsOutput, error)
    UpdateThemePermissionsAsync(ctx workflow.Context, input *quicksight.UpdateThemePermissionsInput) *QuicksightUpdateThemePermissionsResult

    UpdateUser(ctx workflow.Context, input *quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error)
    UpdateUserAsync(ctx workflow.Context, input *quicksight.UpdateUserInput) *QuicksightUpdateUserResult
}
type QuicksightCancelIngestionResult struct {
	Result workflow.Future
}

func (r *QuicksightCancelIngestionResult) Get(ctx workflow.Context) (*quicksight.CancelIngestionOutput, error) {
    var output quicksight.CancelIngestionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateAccountCustomizationResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateAccountCustomizationResult) Get(ctx workflow.Context) (*quicksight.CreateAccountCustomizationOutput, error) {
    var output quicksight.CreateAccountCustomizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateAnalysisResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateAnalysisResult) Get(ctx workflow.Context) (*quicksight.CreateAnalysisOutput, error) {
    var output quicksight.CreateAnalysisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateDashboardResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateDashboardResult) Get(ctx workflow.Context) (*quicksight.CreateDashboardOutput, error) {
    var output quicksight.CreateDashboardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateDataSetResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateDataSetResult) Get(ctx workflow.Context) (*quicksight.CreateDataSetOutput, error) {
    var output quicksight.CreateDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateDataSourceResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateDataSourceResult) Get(ctx workflow.Context) (*quicksight.CreateDataSourceOutput, error) {
    var output quicksight.CreateDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateGroupResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateGroupResult) Get(ctx workflow.Context) (*quicksight.CreateGroupOutput, error) {
    var output quicksight.CreateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateGroupMembershipResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateGroupMembershipResult) Get(ctx workflow.Context) (*quicksight.CreateGroupMembershipOutput, error) {
    var output quicksight.CreateGroupMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateIAMPolicyAssignmentResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateIAMPolicyAssignmentResult) Get(ctx workflow.Context) (*quicksight.CreateIAMPolicyAssignmentOutput, error) {
    var output quicksight.CreateIAMPolicyAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateIngestionResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateIngestionResult) Get(ctx workflow.Context) (*quicksight.CreateIngestionOutput, error) {
    var output quicksight.CreateIngestionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateNamespaceResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateNamespaceResult) Get(ctx workflow.Context) (*quicksight.CreateNamespaceOutput, error) {
    var output quicksight.CreateNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateTemplateResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateTemplateResult) Get(ctx workflow.Context) (*quicksight.CreateTemplateOutput, error) {
    var output quicksight.CreateTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateTemplateAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateTemplateAliasResult) Get(ctx workflow.Context) (*quicksight.CreateTemplateAliasOutput, error) {
    var output quicksight.CreateTemplateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateThemeResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateThemeResult) Get(ctx workflow.Context) (*quicksight.CreateThemeOutput, error) {
    var output quicksight.CreateThemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightCreateThemeAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightCreateThemeAliasResult) Get(ctx workflow.Context) (*quicksight.CreateThemeAliasOutput, error) {
    var output quicksight.CreateThemeAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteAccountCustomizationResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteAccountCustomizationResult) Get(ctx workflow.Context) (*quicksight.DeleteAccountCustomizationOutput, error) {
    var output quicksight.DeleteAccountCustomizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteAnalysisResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteAnalysisResult) Get(ctx workflow.Context) (*quicksight.DeleteAnalysisOutput, error) {
    var output quicksight.DeleteAnalysisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteDashboardResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteDashboardResult) Get(ctx workflow.Context) (*quicksight.DeleteDashboardOutput, error) {
    var output quicksight.DeleteDashboardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteDataSetResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteDataSetResult) Get(ctx workflow.Context) (*quicksight.DeleteDataSetOutput, error) {
    var output quicksight.DeleteDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteDataSourceResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteDataSourceResult) Get(ctx workflow.Context) (*quicksight.DeleteDataSourceOutput, error) {
    var output quicksight.DeleteDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteGroupResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteGroupResult) Get(ctx workflow.Context) (*quicksight.DeleteGroupOutput, error) {
    var output quicksight.DeleteGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteGroupMembershipResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteGroupMembershipResult) Get(ctx workflow.Context) (*quicksight.DeleteGroupMembershipOutput, error) {
    var output quicksight.DeleteGroupMembershipOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteIAMPolicyAssignmentResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteIAMPolicyAssignmentResult) Get(ctx workflow.Context) (*quicksight.DeleteIAMPolicyAssignmentOutput, error) {
    var output quicksight.DeleteIAMPolicyAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteNamespaceResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteNamespaceResult) Get(ctx workflow.Context) (*quicksight.DeleteNamespaceOutput, error) {
    var output quicksight.DeleteNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteTemplateResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteTemplateResult) Get(ctx workflow.Context) (*quicksight.DeleteTemplateOutput, error) {
    var output quicksight.DeleteTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteTemplateAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteTemplateAliasResult) Get(ctx workflow.Context) (*quicksight.DeleteTemplateAliasOutput, error) {
    var output quicksight.DeleteTemplateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteThemeResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteThemeResult) Get(ctx workflow.Context) (*quicksight.DeleteThemeOutput, error) {
    var output quicksight.DeleteThemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteThemeAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteThemeAliasResult) Get(ctx workflow.Context) (*quicksight.DeleteThemeAliasOutput, error) {
    var output quicksight.DeleteThemeAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteUserResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteUserResult) Get(ctx workflow.Context) (*quicksight.DeleteUserOutput, error) {
    var output quicksight.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDeleteUserByPrincipalIdResult struct {
	Result workflow.Future
}

func (r *QuicksightDeleteUserByPrincipalIdResult) Get(ctx workflow.Context) (*quicksight.DeleteUserByPrincipalIdOutput, error) {
    var output quicksight.DeleteUserByPrincipalIdOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeAccountCustomizationResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeAccountCustomizationResult) Get(ctx workflow.Context) (*quicksight.DescribeAccountCustomizationOutput, error) {
    var output quicksight.DescribeAccountCustomizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeAccountSettingsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeAccountSettingsResult) Get(ctx workflow.Context) (*quicksight.DescribeAccountSettingsOutput, error) {
    var output quicksight.DescribeAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeAnalysisResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeAnalysisResult) Get(ctx workflow.Context) (*quicksight.DescribeAnalysisOutput, error) {
    var output quicksight.DescribeAnalysisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeAnalysisPermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeAnalysisPermissionsResult) Get(ctx workflow.Context) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
    var output quicksight.DescribeAnalysisPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeDashboardResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeDashboardResult) Get(ctx workflow.Context) (*quicksight.DescribeDashboardOutput, error) {
    var output quicksight.DescribeDashboardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeDashboardPermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeDashboardPermissionsResult) Get(ctx workflow.Context) (*quicksight.DescribeDashboardPermissionsOutput, error) {
    var output quicksight.DescribeDashboardPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeDataSetResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeDataSetResult) Get(ctx workflow.Context) (*quicksight.DescribeDataSetOutput, error) {
    var output quicksight.DescribeDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeDataSetPermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeDataSetPermissionsResult) Get(ctx workflow.Context) (*quicksight.DescribeDataSetPermissionsOutput, error) {
    var output quicksight.DescribeDataSetPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeDataSourceResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeDataSourceResult) Get(ctx workflow.Context) (*quicksight.DescribeDataSourceOutput, error) {
    var output quicksight.DescribeDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeDataSourcePermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeDataSourcePermissionsResult) Get(ctx workflow.Context) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
    var output quicksight.DescribeDataSourcePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeGroupResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeGroupResult) Get(ctx workflow.Context) (*quicksight.DescribeGroupOutput, error) {
    var output quicksight.DescribeGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeIAMPolicyAssignmentResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeIAMPolicyAssignmentResult) Get(ctx workflow.Context) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
    var output quicksight.DescribeIAMPolicyAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeIngestionResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeIngestionResult) Get(ctx workflow.Context) (*quicksight.DescribeIngestionOutput, error) {
    var output quicksight.DescribeIngestionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeNamespaceResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeNamespaceResult) Get(ctx workflow.Context) (*quicksight.DescribeNamespaceOutput, error) {
    var output quicksight.DescribeNamespaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeTemplateResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeTemplateResult) Get(ctx workflow.Context) (*quicksight.DescribeTemplateOutput, error) {
    var output quicksight.DescribeTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeTemplateAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeTemplateAliasResult) Get(ctx workflow.Context) (*quicksight.DescribeTemplateAliasOutput, error) {
    var output quicksight.DescribeTemplateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeTemplatePermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeTemplatePermissionsResult) Get(ctx workflow.Context) (*quicksight.DescribeTemplatePermissionsOutput, error) {
    var output quicksight.DescribeTemplatePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeThemeResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeThemeResult) Get(ctx workflow.Context) (*quicksight.DescribeThemeOutput, error) {
    var output quicksight.DescribeThemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeThemeAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeThemeAliasResult) Get(ctx workflow.Context) (*quicksight.DescribeThemeAliasOutput, error) {
    var output quicksight.DescribeThemeAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeThemePermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeThemePermissionsResult) Get(ctx workflow.Context) (*quicksight.DescribeThemePermissionsOutput, error) {
    var output quicksight.DescribeThemePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightDescribeUserResult struct {
	Result workflow.Future
}

func (r *QuicksightDescribeUserResult) Get(ctx workflow.Context) (*quicksight.DescribeUserOutput, error) {
    var output quicksight.DescribeUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightGetDashboardEmbedUrlResult struct {
	Result workflow.Future
}

func (r *QuicksightGetDashboardEmbedUrlResult) Get(ctx workflow.Context) (*quicksight.GetDashboardEmbedUrlOutput, error) {
    var output quicksight.GetDashboardEmbedUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightGetSessionEmbedUrlResult struct {
	Result workflow.Future
}

func (r *QuicksightGetSessionEmbedUrlResult) Get(ctx workflow.Context) (*quicksight.GetSessionEmbedUrlOutput, error) {
    var output quicksight.GetSessionEmbedUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListAnalysesResult struct {
	Result workflow.Future
}

func (r *QuicksightListAnalysesResult) Get(ctx workflow.Context) (*quicksight.ListAnalysesOutput, error) {
    var output quicksight.ListAnalysesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListDashboardVersionsResult struct {
	Result workflow.Future
}

func (r *QuicksightListDashboardVersionsResult) Get(ctx workflow.Context) (*quicksight.ListDashboardVersionsOutput, error) {
    var output quicksight.ListDashboardVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListDashboardsResult struct {
	Result workflow.Future
}

func (r *QuicksightListDashboardsResult) Get(ctx workflow.Context) (*quicksight.ListDashboardsOutput, error) {
    var output quicksight.ListDashboardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListDataSetsResult struct {
	Result workflow.Future
}

func (r *QuicksightListDataSetsResult) Get(ctx workflow.Context) (*quicksight.ListDataSetsOutput, error) {
    var output quicksight.ListDataSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListDataSourcesResult struct {
	Result workflow.Future
}

func (r *QuicksightListDataSourcesResult) Get(ctx workflow.Context) (*quicksight.ListDataSourcesOutput, error) {
    var output quicksight.ListDataSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListGroupMembershipsResult struct {
	Result workflow.Future
}

func (r *QuicksightListGroupMembershipsResult) Get(ctx workflow.Context) (*quicksight.ListGroupMembershipsOutput, error) {
    var output quicksight.ListGroupMembershipsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListGroupsResult struct {
	Result workflow.Future
}

func (r *QuicksightListGroupsResult) Get(ctx workflow.Context) (*quicksight.ListGroupsOutput, error) {
    var output quicksight.ListGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListIAMPolicyAssignmentsResult struct {
	Result workflow.Future
}

func (r *QuicksightListIAMPolicyAssignmentsResult) Get(ctx workflow.Context) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
    var output quicksight.ListIAMPolicyAssignmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListIAMPolicyAssignmentsForUserResult struct {
	Result workflow.Future
}

func (r *QuicksightListIAMPolicyAssignmentsForUserResult) Get(ctx workflow.Context) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
    var output quicksight.ListIAMPolicyAssignmentsForUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListIngestionsResult struct {
	Result workflow.Future
}

func (r *QuicksightListIngestionsResult) Get(ctx workflow.Context) (*quicksight.ListIngestionsOutput, error) {
    var output quicksight.ListIngestionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListNamespacesResult struct {
	Result workflow.Future
}

func (r *QuicksightListNamespacesResult) Get(ctx workflow.Context) (*quicksight.ListNamespacesOutput, error) {
    var output quicksight.ListNamespacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *QuicksightListTagsForResourceResult) Get(ctx workflow.Context) (*quicksight.ListTagsForResourceOutput, error) {
    var output quicksight.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListTemplateAliasesResult struct {
	Result workflow.Future
}

func (r *QuicksightListTemplateAliasesResult) Get(ctx workflow.Context) (*quicksight.ListTemplateAliasesOutput, error) {
    var output quicksight.ListTemplateAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListTemplateVersionsResult struct {
	Result workflow.Future
}

func (r *QuicksightListTemplateVersionsResult) Get(ctx workflow.Context) (*quicksight.ListTemplateVersionsOutput, error) {
    var output quicksight.ListTemplateVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListTemplatesResult struct {
	Result workflow.Future
}

func (r *QuicksightListTemplatesResult) Get(ctx workflow.Context) (*quicksight.ListTemplatesOutput, error) {
    var output quicksight.ListTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListThemeAliasesResult struct {
	Result workflow.Future
}

func (r *QuicksightListThemeAliasesResult) Get(ctx workflow.Context) (*quicksight.ListThemeAliasesOutput, error) {
    var output quicksight.ListThemeAliasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListThemeVersionsResult struct {
	Result workflow.Future
}

func (r *QuicksightListThemeVersionsResult) Get(ctx workflow.Context) (*quicksight.ListThemeVersionsOutput, error) {
    var output quicksight.ListThemeVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListThemesResult struct {
	Result workflow.Future
}

func (r *QuicksightListThemesResult) Get(ctx workflow.Context) (*quicksight.ListThemesOutput, error) {
    var output quicksight.ListThemesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListUserGroupsResult struct {
	Result workflow.Future
}

func (r *QuicksightListUserGroupsResult) Get(ctx workflow.Context) (*quicksight.ListUserGroupsOutput, error) {
    var output quicksight.ListUserGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightListUsersResult struct {
	Result workflow.Future
}

func (r *QuicksightListUsersResult) Get(ctx workflow.Context) (*quicksight.ListUsersOutput, error) {
    var output quicksight.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightRegisterUserResult struct {
	Result workflow.Future
}

func (r *QuicksightRegisterUserResult) Get(ctx workflow.Context) (*quicksight.RegisterUserOutput, error) {
    var output quicksight.RegisterUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightRestoreAnalysisResult struct {
	Result workflow.Future
}

func (r *QuicksightRestoreAnalysisResult) Get(ctx workflow.Context) (*quicksight.RestoreAnalysisOutput, error) {
    var output quicksight.RestoreAnalysisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightSearchAnalysesResult struct {
	Result workflow.Future
}

func (r *QuicksightSearchAnalysesResult) Get(ctx workflow.Context) (*quicksight.SearchAnalysesOutput, error) {
    var output quicksight.SearchAnalysesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightSearchDashboardsResult struct {
	Result workflow.Future
}

func (r *QuicksightSearchDashboardsResult) Get(ctx workflow.Context) (*quicksight.SearchDashboardsOutput, error) {
    var output quicksight.SearchDashboardsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightTagResourceResult struct {
	Result workflow.Future
}

func (r *QuicksightTagResourceResult) Get(ctx workflow.Context) (*quicksight.TagResourceOutput, error) {
    var output quicksight.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUntagResourceResult struct {
	Result workflow.Future
}

func (r *QuicksightUntagResourceResult) Get(ctx workflow.Context) (*quicksight.UntagResourceOutput, error) {
    var output quicksight.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateAccountCustomizationResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateAccountCustomizationResult) Get(ctx workflow.Context) (*quicksight.UpdateAccountCustomizationOutput, error) {
    var output quicksight.UpdateAccountCustomizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateAccountSettingsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateAccountSettingsResult) Get(ctx workflow.Context) (*quicksight.UpdateAccountSettingsOutput, error) {
    var output quicksight.UpdateAccountSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateAnalysisResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateAnalysisResult) Get(ctx workflow.Context) (*quicksight.UpdateAnalysisOutput, error) {
    var output quicksight.UpdateAnalysisOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateAnalysisPermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateAnalysisPermissionsResult) Get(ctx workflow.Context) (*quicksight.UpdateAnalysisPermissionsOutput, error) {
    var output quicksight.UpdateAnalysisPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDashboardResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDashboardResult) Get(ctx workflow.Context) (*quicksight.UpdateDashboardOutput, error) {
    var output quicksight.UpdateDashboardOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDashboardPermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDashboardPermissionsResult) Get(ctx workflow.Context) (*quicksight.UpdateDashboardPermissionsOutput, error) {
    var output quicksight.UpdateDashboardPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDashboardPublishedVersionResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDashboardPublishedVersionResult) Get(ctx workflow.Context) (*quicksight.UpdateDashboardPublishedVersionOutput, error) {
    var output quicksight.UpdateDashboardPublishedVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDataSetResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDataSetResult) Get(ctx workflow.Context) (*quicksight.UpdateDataSetOutput, error) {
    var output quicksight.UpdateDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDataSetPermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDataSetPermissionsResult) Get(ctx workflow.Context) (*quicksight.UpdateDataSetPermissionsOutput, error) {
    var output quicksight.UpdateDataSetPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDataSourceResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDataSourceResult) Get(ctx workflow.Context) (*quicksight.UpdateDataSourceOutput, error) {
    var output quicksight.UpdateDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateDataSourcePermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateDataSourcePermissionsResult) Get(ctx workflow.Context) (*quicksight.UpdateDataSourcePermissionsOutput, error) {
    var output quicksight.UpdateDataSourcePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateGroupResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateGroupResult) Get(ctx workflow.Context) (*quicksight.UpdateGroupOutput, error) {
    var output quicksight.UpdateGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateIAMPolicyAssignmentResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateIAMPolicyAssignmentResult) Get(ctx workflow.Context) (*quicksight.UpdateIAMPolicyAssignmentOutput, error) {
    var output quicksight.UpdateIAMPolicyAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateTemplateResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateTemplateResult) Get(ctx workflow.Context) (*quicksight.UpdateTemplateOutput, error) {
    var output quicksight.UpdateTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateTemplateAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateTemplateAliasResult) Get(ctx workflow.Context) (*quicksight.UpdateTemplateAliasOutput, error) {
    var output quicksight.UpdateTemplateAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateTemplatePermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateTemplatePermissionsResult) Get(ctx workflow.Context) (*quicksight.UpdateTemplatePermissionsOutput, error) {
    var output quicksight.UpdateTemplatePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateThemeResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateThemeResult) Get(ctx workflow.Context) (*quicksight.UpdateThemeOutput, error) {
    var output quicksight.UpdateThemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateThemeAliasResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateThemeAliasResult) Get(ctx workflow.Context) (*quicksight.UpdateThemeAliasOutput, error) {
    var output quicksight.UpdateThemeAliasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateThemePermissionsResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateThemePermissionsResult) Get(ctx workflow.Context) (*quicksight.UpdateThemePermissionsOutput, error) {
    var output quicksight.UpdateThemePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type QuicksightUpdateUserResult struct {
	Result workflow.Future
}

func (r *QuicksightUpdateUserResult) Get(ctx workflow.Context) (*quicksight.UpdateUserOutput, error) {
    var output quicksight.UpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type QuickSightStub struct {
    activities QuickSightClient
}

func NewQuickSightStub() QuickSightClient {
    return &QuickSightStub{}
}

func (a *QuickSightStub) CancelIngestion(ctx workflow.Context, input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error) {
    var output quicksight.CancelIngestionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelIngestion, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateAccountCustomization(ctx workflow.Context, input *quicksight.CreateAccountCustomizationInput) (*quicksight.CreateAccountCustomizationOutput, error) {
    var output quicksight.CreateAccountCustomizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccountCustomization, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateAnalysis(ctx workflow.Context, input *quicksight.CreateAnalysisInput) (*quicksight.CreateAnalysisOutput, error) {
    var output quicksight.CreateAnalysisOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAnalysis, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateDashboard(ctx workflow.Context, input *quicksight.CreateDashboardInput) (*quicksight.CreateDashboardOutput, error) {
    var output quicksight.CreateDashboardOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDashboard, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateDataSet(ctx workflow.Context, input *quicksight.CreateDataSetInput) (*quicksight.CreateDataSetOutput, error) {
    var output quicksight.CreateDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateDataSource(ctx workflow.Context, input *quicksight.CreateDataSourceInput) (*quicksight.CreateDataSourceOutput, error) {
    var output quicksight.CreateDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateGroup(ctx workflow.Context, input *quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error) {
    var output quicksight.CreateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateGroupMembership(ctx workflow.Context, input *quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error) {
    var output quicksight.CreateGroupMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGroupMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateIAMPolicyAssignment(ctx workflow.Context, input *quicksight.CreateIAMPolicyAssignmentInput) (*quicksight.CreateIAMPolicyAssignmentOutput, error) {
    var output quicksight.CreateIAMPolicyAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIAMPolicyAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateIngestion(ctx workflow.Context, input *quicksight.CreateIngestionInput) (*quicksight.CreateIngestionOutput, error) {
    var output quicksight.CreateIngestionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIngestion, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateNamespace(ctx workflow.Context, input *quicksight.CreateNamespaceInput) (*quicksight.CreateNamespaceOutput, error) {
    var output quicksight.CreateNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateTemplate(ctx workflow.Context, input *quicksight.CreateTemplateInput) (*quicksight.CreateTemplateOutput, error) {
    var output quicksight.CreateTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateTemplateAlias(ctx workflow.Context, input *quicksight.CreateTemplateAliasInput) (*quicksight.CreateTemplateAliasOutput, error) {
    var output quicksight.CreateTemplateAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTemplateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateTheme(ctx workflow.Context, input *quicksight.CreateThemeInput) (*quicksight.CreateThemeOutput, error) {
    var output quicksight.CreateThemeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTheme, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) CreateThemeAlias(ctx workflow.Context, input *quicksight.CreateThemeAliasInput) (*quicksight.CreateThemeAliasOutput, error) {
    var output quicksight.CreateThemeAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateThemeAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteAccountCustomization(ctx workflow.Context, input *quicksight.DeleteAccountCustomizationInput) (*quicksight.DeleteAccountCustomizationOutput, error) {
    var output quicksight.DeleteAccountCustomizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccountCustomization, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteAnalysis(ctx workflow.Context, input *quicksight.DeleteAnalysisInput) (*quicksight.DeleteAnalysisOutput, error) {
    var output quicksight.DeleteAnalysisOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAnalysis, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteDashboard(ctx workflow.Context, input *quicksight.DeleteDashboardInput) (*quicksight.DeleteDashboardOutput, error) {
    var output quicksight.DeleteDashboardOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDashboard, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteDataSet(ctx workflow.Context, input *quicksight.DeleteDataSetInput) (*quicksight.DeleteDataSetOutput, error) {
    var output quicksight.DeleteDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteDataSource(ctx workflow.Context, input *quicksight.DeleteDataSourceInput) (*quicksight.DeleteDataSourceOutput, error) {
    var output quicksight.DeleteDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteGroup(ctx workflow.Context, input *quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error) {
    var output quicksight.DeleteGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteGroupMembership(ctx workflow.Context, input *quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error) {
    var output quicksight.DeleteGroupMembershipOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGroupMembership, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteIAMPolicyAssignment(ctx workflow.Context, input *quicksight.DeleteIAMPolicyAssignmentInput) (*quicksight.DeleteIAMPolicyAssignmentOutput, error) {
    var output quicksight.DeleteIAMPolicyAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIAMPolicyAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteNamespace(ctx workflow.Context, input *quicksight.DeleteNamespaceInput) (*quicksight.DeleteNamespaceOutput, error) {
    var output quicksight.DeleteNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteTemplate(ctx workflow.Context, input *quicksight.DeleteTemplateInput) (*quicksight.DeleteTemplateOutput, error) {
    var output quicksight.DeleteTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteTemplateAlias(ctx workflow.Context, input *quicksight.DeleteTemplateAliasInput) (*quicksight.DeleteTemplateAliasOutput, error) {
    var output quicksight.DeleteTemplateAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTemplateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteTheme(ctx workflow.Context, input *quicksight.DeleteThemeInput) (*quicksight.DeleteThemeOutput, error) {
    var output quicksight.DeleteThemeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTheme, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteThemeAlias(ctx workflow.Context, input *quicksight.DeleteThemeAliasInput) (*quicksight.DeleteThemeAliasOutput, error) {
    var output quicksight.DeleteThemeAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteThemeAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteUser(ctx workflow.Context, input *quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error) {
    var output quicksight.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DeleteUserByPrincipalId(ctx workflow.Context, input *quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error) {
    var output quicksight.DeleteUserByPrincipalIdOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserByPrincipalId, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeAccountCustomization(ctx workflow.Context, input *quicksight.DescribeAccountCustomizationInput) (*quicksight.DescribeAccountCustomizationOutput, error) {
    var output quicksight.DescribeAccountCustomizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountCustomization, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeAccountSettings(ctx workflow.Context, input *quicksight.DescribeAccountSettingsInput) (*quicksight.DescribeAccountSettingsOutput, error) {
    var output quicksight.DescribeAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeAnalysis(ctx workflow.Context, input *quicksight.DescribeAnalysisInput) (*quicksight.DescribeAnalysisOutput, error) {
    var output quicksight.DescribeAnalysisOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAnalysis, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeAnalysisPermissions(ctx workflow.Context, input *quicksight.DescribeAnalysisPermissionsInput) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
    var output quicksight.DescribeAnalysisPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAnalysisPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeDashboard(ctx workflow.Context, input *quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error) {
    var output quicksight.DescribeDashboardOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDashboard, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeDashboardPermissions(ctx workflow.Context, input *quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error) {
    var output quicksight.DescribeDashboardPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDashboardPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeDataSet(ctx workflow.Context, input *quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error) {
    var output quicksight.DescribeDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeDataSetPermissions(ctx workflow.Context, input *quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error) {
    var output quicksight.DescribeDataSetPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSetPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeDataSource(ctx workflow.Context, input *quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error) {
    var output quicksight.DescribeDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeDataSourcePermissions(ctx workflow.Context, input *quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
    var output quicksight.DescribeDataSourcePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataSourcePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeGroup(ctx workflow.Context, input *quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error) {
    var output quicksight.DescribeGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeIAMPolicyAssignment(ctx workflow.Context, input *quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
    var output quicksight.DescribeIAMPolicyAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIAMPolicyAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeIngestion(ctx workflow.Context, input *quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error) {
    var output quicksight.DescribeIngestionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIngestion, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeNamespace(ctx workflow.Context, input *quicksight.DescribeNamespaceInput) (*quicksight.DescribeNamespaceOutput, error) {
    var output quicksight.DescribeNamespaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeNamespace, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeTemplate(ctx workflow.Context, input *quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error) {
    var output quicksight.DescribeTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeTemplateAlias(ctx workflow.Context, input *quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error) {
    var output quicksight.DescribeTemplateAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTemplateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeTemplatePermissions(ctx workflow.Context, input *quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error) {
    var output quicksight.DescribeTemplatePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTemplatePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeTheme(ctx workflow.Context, input *quicksight.DescribeThemeInput) (*quicksight.DescribeThemeOutput, error) {
    var output quicksight.DescribeThemeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTheme, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeThemeAlias(ctx workflow.Context, input *quicksight.DescribeThemeAliasInput) (*quicksight.DescribeThemeAliasOutput, error) {
    var output quicksight.DescribeThemeAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeThemeAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeThemePermissions(ctx workflow.Context, input *quicksight.DescribeThemePermissionsInput) (*quicksight.DescribeThemePermissionsOutput, error) {
    var output quicksight.DescribeThemePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeThemePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) DescribeUser(ctx workflow.Context, input *quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error) {
    var output quicksight.DescribeUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) GetDashboardEmbedUrl(ctx workflow.Context, input *quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error) {
    var output quicksight.GetDashboardEmbedUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDashboardEmbedUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) GetSessionEmbedUrl(ctx workflow.Context, input *quicksight.GetSessionEmbedUrlInput) (*quicksight.GetSessionEmbedUrlOutput, error) {
    var output quicksight.GetSessionEmbedUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSessionEmbedUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListAnalyses(ctx workflow.Context, input *quicksight.ListAnalysesInput) (*quicksight.ListAnalysesOutput, error) {
    var output quicksight.ListAnalysesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAnalyses, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListDashboardVersions(ctx workflow.Context, input *quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error) {
    var output quicksight.ListDashboardVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDashboardVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListDashboards(ctx workflow.Context, input *quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error) {
    var output quicksight.ListDashboardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDashboards, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListDataSets(ctx workflow.Context, input *quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error) {
    var output quicksight.ListDataSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataSets, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListDataSources(ctx workflow.Context, input *quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error) {
    var output quicksight.ListDataSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataSources, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListGroupMemberships(ctx workflow.Context, input *quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error) {
    var output quicksight.ListGroupMembershipsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroupMemberships, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListGroups(ctx workflow.Context, input *quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error) {
    var output quicksight.ListGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListIAMPolicyAssignments(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
    var output quicksight.ListIAMPolicyAssignmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIAMPolicyAssignments, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListIAMPolicyAssignmentsForUser(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
    var output quicksight.ListIAMPolicyAssignmentsForUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIAMPolicyAssignmentsForUser, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListIngestions(ctx workflow.Context, input *quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error) {
    var output quicksight.ListIngestionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIngestions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListNamespaces(ctx workflow.Context, input *quicksight.ListNamespacesInput) (*quicksight.ListNamespacesOutput, error) {
    var output quicksight.ListNamespacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListNamespaces, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListTagsForResource(ctx workflow.Context, input *quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error) {
    var output quicksight.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListTemplateAliases(ctx workflow.Context, input *quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error) {
    var output quicksight.ListTemplateAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTemplateAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListTemplateVersions(ctx workflow.Context, input *quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error) {
    var output quicksight.ListTemplateVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTemplateVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListTemplates(ctx workflow.Context, input *quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error) {
    var output quicksight.ListTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListThemeAliases(ctx workflow.Context, input *quicksight.ListThemeAliasesInput) (*quicksight.ListThemeAliasesOutput, error) {
    var output quicksight.ListThemeAliasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThemeAliases, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListThemeVersions(ctx workflow.Context, input *quicksight.ListThemeVersionsInput) (*quicksight.ListThemeVersionsOutput, error) {
    var output quicksight.ListThemeVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThemeVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListThemes(ctx workflow.Context, input *quicksight.ListThemesInput) (*quicksight.ListThemesOutput, error) {
    var output quicksight.ListThemesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListThemes, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListUserGroups(ctx workflow.Context, input *quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error) {
    var output quicksight.ListUserGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) ListUsers(ctx workflow.Context, input *quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error) {
    var output quicksight.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) RegisterUser(ctx workflow.Context, input *quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error) {
    var output quicksight.RegisterUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterUser, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) RestoreAnalysis(ctx workflow.Context, input *quicksight.RestoreAnalysisInput) (*quicksight.RestoreAnalysisOutput, error) {
    var output quicksight.RestoreAnalysisOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreAnalysis, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) SearchAnalyses(ctx workflow.Context, input *quicksight.SearchAnalysesInput) (*quicksight.SearchAnalysesOutput, error) {
    var output quicksight.SearchAnalysesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchAnalyses, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) SearchDashboards(ctx workflow.Context, input *quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error) {
    var output quicksight.SearchDashboardsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchDashboards, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) TagResource(ctx workflow.Context, input *quicksight.TagResourceInput) (*quicksight.TagResourceOutput, error) {
    var output quicksight.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UntagResource(ctx workflow.Context, input *quicksight.UntagResourceInput) (*quicksight.UntagResourceOutput, error) {
    var output quicksight.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateAccountCustomization(ctx workflow.Context, input *quicksight.UpdateAccountCustomizationInput) (*quicksight.UpdateAccountCustomizationOutput, error) {
    var output quicksight.UpdateAccountCustomizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountCustomization, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateAccountSettings(ctx workflow.Context, input *quicksight.UpdateAccountSettingsInput) (*quicksight.UpdateAccountSettingsOutput, error) {
    var output quicksight.UpdateAccountSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateAnalysis(ctx workflow.Context, input *quicksight.UpdateAnalysisInput) (*quicksight.UpdateAnalysisOutput, error) {
    var output quicksight.UpdateAnalysisOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAnalysis, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateAnalysisPermissions(ctx workflow.Context, input *quicksight.UpdateAnalysisPermissionsInput) (*quicksight.UpdateAnalysisPermissionsOutput, error) {
    var output quicksight.UpdateAnalysisPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAnalysisPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDashboard(ctx workflow.Context, input *quicksight.UpdateDashboardInput) (*quicksight.UpdateDashboardOutput, error) {
    var output quicksight.UpdateDashboardOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDashboard, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDashboardPermissions(ctx workflow.Context, input *quicksight.UpdateDashboardPermissionsInput) (*quicksight.UpdateDashboardPermissionsOutput, error) {
    var output quicksight.UpdateDashboardPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDashboardPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDashboardPublishedVersion(ctx workflow.Context, input *quicksight.UpdateDashboardPublishedVersionInput) (*quicksight.UpdateDashboardPublishedVersionOutput, error) {
    var output quicksight.UpdateDashboardPublishedVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDashboardPublishedVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDataSet(ctx workflow.Context, input *quicksight.UpdateDataSetInput) (*quicksight.UpdateDataSetOutput, error) {
    var output quicksight.UpdateDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDataSetPermissions(ctx workflow.Context, input *quicksight.UpdateDataSetPermissionsInput) (*quicksight.UpdateDataSetPermissionsOutput, error) {
    var output quicksight.UpdateDataSetPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSetPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDataSource(ctx workflow.Context, input *quicksight.UpdateDataSourceInput) (*quicksight.UpdateDataSourceOutput, error) {
    var output quicksight.UpdateDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateDataSourcePermissions(ctx workflow.Context, input *quicksight.UpdateDataSourcePermissionsInput) (*quicksight.UpdateDataSourcePermissionsOutput, error) {
    var output quicksight.UpdateDataSourcePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSourcePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateGroup(ctx workflow.Context, input *quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error) {
    var output quicksight.UpdateGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateIAMPolicyAssignment(ctx workflow.Context, input *quicksight.UpdateIAMPolicyAssignmentInput) (*quicksight.UpdateIAMPolicyAssignmentOutput, error) {
    var output quicksight.UpdateIAMPolicyAssignmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIAMPolicyAssignment, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateTemplate(ctx workflow.Context, input *quicksight.UpdateTemplateInput) (*quicksight.UpdateTemplateOutput, error) {
    var output quicksight.UpdateTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateTemplateAlias(ctx workflow.Context, input *quicksight.UpdateTemplateAliasInput) (*quicksight.UpdateTemplateAliasOutput, error) {
    var output quicksight.UpdateTemplateAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTemplateAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateTemplatePermissions(ctx workflow.Context, input *quicksight.UpdateTemplatePermissionsInput) (*quicksight.UpdateTemplatePermissionsOutput, error) {
    var output quicksight.UpdateTemplatePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTemplatePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateTheme(ctx workflow.Context, input *quicksight.UpdateThemeInput) (*quicksight.UpdateThemeOutput, error) {
    var output quicksight.UpdateThemeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTheme, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateThemeAlias(ctx workflow.Context, input *quicksight.UpdateThemeAliasInput) (*quicksight.UpdateThemeAliasOutput, error) {
    var output quicksight.UpdateThemeAliasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateThemeAlias, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateThemePermissions(ctx workflow.Context, input *quicksight.UpdateThemePermissionsInput) (*quicksight.UpdateThemePermissionsOutput, error) {
    var output quicksight.UpdateThemePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateThemePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *QuickSightStub) UpdateUser(ctx workflow.Context, input *quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error) {
    var output quicksight.UpdateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUser, input).Get(ctx, &output)
    return &output, err
}
