// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package quicksightstub

import (
	"github.com/aws/aws-sdk-go/service/quicksight"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelIngestion(ctx workflow.Context, input *quicksight.CancelIngestionInput) (*quicksight.CancelIngestionOutput, error)
	CancelIngestionAsync(ctx workflow.Context, input *quicksight.CancelIngestionInput) *CancelIngestionFuture

	CreateAccountCustomization(ctx workflow.Context, input *quicksight.CreateAccountCustomizationInput) (*quicksight.CreateAccountCustomizationOutput, error)
	CreateAccountCustomizationAsync(ctx workflow.Context, input *quicksight.CreateAccountCustomizationInput) *CreateAccountCustomizationFuture

	CreateAnalysis(ctx workflow.Context, input *quicksight.CreateAnalysisInput) (*quicksight.CreateAnalysisOutput, error)
	CreateAnalysisAsync(ctx workflow.Context, input *quicksight.CreateAnalysisInput) *CreateAnalysisFuture

	CreateDashboard(ctx workflow.Context, input *quicksight.CreateDashboardInput) (*quicksight.CreateDashboardOutput, error)
	CreateDashboardAsync(ctx workflow.Context, input *quicksight.CreateDashboardInput) *CreateDashboardFuture

	CreateDataSet(ctx workflow.Context, input *quicksight.CreateDataSetInput) (*quicksight.CreateDataSetOutput, error)
	CreateDataSetAsync(ctx workflow.Context, input *quicksight.CreateDataSetInput) *CreateDataSetFuture

	CreateDataSource(ctx workflow.Context, input *quicksight.CreateDataSourceInput) (*quicksight.CreateDataSourceOutput, error)
	CreateDataSourceAsync(ctx workflow.Context, input *quicksight.CreateDataSourceInput) *CreateDataSourceFuture

	CreateGroup(ctx workflow.Context, input *quicksight.CreateGroupInput) (*quicksight.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *quicksight.CreateGroupInput) *CreateGroupFuture

	CreateGroupMembership(ctx workflow.Context, input *quicksight.CreateGroupMembershipInput) (*quicksight.CreateGroupMembershipOutput, error)
	CreateGroupMembershipAsync(ctx workflow.Context, input *quicksight.CreateGroupMembershipInput) *CreateGroupMembershipFuture

	CreateIAMPolicyAssignment(ctx workflow.Context, input *quicksight.CreateIAMPolicyAssignmentInput) (*quicksight.CreateIAMPolicyAssignmentOutput, error)
	CreateIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.CreateIAMPolicyAssignmentInput) *CreateIAMPolicyAssignmentFuture

	CreateIngestion(ctx workflow.Context, input *quicksight.CreateIngestionInput) (*quicksight.CreateIngestionOutput, error)
	CreateIngestionAsync(ctx workflow.Context, input *quicksight.CreateIngestionInput) *CreateIngestionFuture

	CreateNamespace(ctx workflow.Context, input *quicksight.CreateNamespaceInput) (*quicksight.CreateNamespaceOutput, error)
	CreateNamespaceAsync(ctx workflow.Context, input *quicksight.CreateNamespaceInput) *CreateNamespaceFuture

	CreateTemplate(ctx workflow.Context, input *quicksight.CreateTemplateInput) (*quicksight.CreateTemplateOutput, error)
	CreateTemplateAsync(ctx workflow.Context, input *quicksight.CreateTemplateInput) *CreateTemplateFuture

	CreateTemplateAlias(ctx workflow.Context, input *quicksight.CreateTemplateAliasInput) (*quicksight.CreateTemplateAliasOutput, error)
	CreateTemplateAliasAsync(ctx workflow.Context, input *quicksight.CreateTemplateAliasInput) *CreateTemplateAliasFuture

	CreateTheme(ctx workflow.Context, input *quicksight.CreateThemeInput) (*quicksight.CreateThemeOutput, error)
	CreateThemeAsync(ctx workflow.Context, input *quicksight.CreateThemeInput) *CreateThemeFuture

	CreateThemeAlias(ctx workflow.Context, input *quicksight.CreateThemeAliasInput) (*quicksight.CreateThemeAliasOutput, error)
	CreateThemeAliasAsync(ctx workflow.Context, input *quicksight.CreateThemeAliasInput) *CreateThemeAliasFuture

	DeleteAccountCustomization(ctx workflow.Context, input *quicksight.DeleteAccountCustomizationInput) (*quicksight.DeleteAccountCustomizationOutput, error)
	DeleteAccountCustomizationAsync(ctx workflow.Context, input *quicksight.DeleteAccountCustomizationInput) *DeleteAccountCustomizationFuture

	DeleteAnalysis(ctx workflow.Context, input *quicksight.DeleteAnalysisInput) (*quicksight.DeleteAnalysisOutput, error)
	DeleteAnalysisAsync(ctx workflow.Context, input *quicksight.DeleteAnalysisInput) *DeleteAnalysisFuture

	DeleteDashboard(ctx workflow.Context, input *quicksight.DeleteDashboardInput) (*quicksight.DeleteDashboardOutput, error)
	DeleteDashboardAsync(ctx workflow.Context, input *quicksight.DeleteDashboardInput) *DeleteDashboardFuture

	DeleteDataSet(ctx workflow.Context, input *quicksight.DeleteDataSetInput) (*quicksight.DeleteDataSetOutput, error)
	DeleteDataSetAsync(ctx workflow.Context, input *quicksight.DeleteDataSetInput) *DeleteDataSetFuture

	DeleteDataSource(ctx workflow.Context, input *quicksight.DeleteDataSourceInput) (*quicksight.DeleteDataSourceOutput, error)
	DeleteDataSourceAsync(ctx workflow.Context, input *quicksight.DeleteDataSourceInput) *DeleteDataSourceFuture

	DeleteGroup(ctx workflow.Context, input *quicksight.DeleteGroupInput) (*quicksight.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *quicksight.DeleteGroupInput) *DeleteGroupFuture

	DeleteGroupMembership(ctx workflow.Context, input *quicksight.DeleteGroupMembershipInput) (*quicksight.DeleteGroupMembershipOutput, error)
	DeleteGroupMembershipAsync(ctx workflow.Context, input *quicksight.DeleteGroupMembershipInput) *DeleteGroupMembershipFuture

	DeleteIAMPolicyAssignment(ctx workflow.Context, input *quicksight.DeleteIAMPolicyAssignmentInput) (*quicksight.DeleteIAMPolicyAssignmentOutput, error)
	DeleteIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.DeleteIAMPolicyAssignmentInput) *DeleteIAMPolicyAssignmentFuture

	DeleteNamespace(ctx workflow.Context, input *quicksight.DeleteNamespaceInput) (*quicksight.DeleteNamespaceOutput, error)
	DeleteNamespaceAsync(ctx workflow.Context, input *quicksight.DeleteNamespaceInput) *DeleteNamespaceFuture

	DeleteTemplate(ctx workflow.Context, input *quicksight.DeleteTemplateInput) (*quicksight.DeleteTemplateOutput, error)
	DeleteTemplateAsync(ctx workflow.Context, input *quicksight.DeleteTemplateInput) *DeleteTemplateFuture

	DeleteTemplateAlias(ctx workflow.Context, input *quicksight.DeleteTemplateAliasInput) (*quicksight.DeleteTemplateAliasOutput, error)
	DeleteTemplateAliasAsync(ctx workflow.Context, input *quicksight.DeleteTemplateAliasInput) *DeleteTemplateAliasFuture

	DeleteTheme(ctx workflow.Context, input *quicksight.DeleteThemeInput) (*quicksight.DeleteThemeOutput, error)
	DeleteThemeAsync(ctx workflow.Context, input *quicksight.DeleteThemeInput) *DeleteThemeFuture

	DeleteThemeAlias(ctx workflow.Context, input *quicksight.DeleteThemeAliasInput) (*quicksight.DeleteThemeAliasOutput, error)
	DeleteThemeAliasAsync(ctx workflow.Context, input *quicksight.DeleteThemeAliasInput) *DeleteThemeAliasFuture

	DeleteUser(ctx workflow.Context, input *quicksight.DeleteUserInput) (*quicksight.DeleteUserOutput, error)
	DeleteUserAsync(ctx workflow.Context, input *quicksight.DeleteUserInput) *DeleteUserFuture

	DeleteUserByPrincipalId(ctx workflow.Context, input *quicksight.DeleteUserByPrincipalIdInput) (*quicksight.DeleteUserByPrincipalIdOutput, error)
	DeleteUserByPrincipalIdAsync(ctx workflow.Context, input *quicksight.DeleteUserByPrincipalIdInput) *DeleteUserByPrincipalIdFuture

	DescribeAccountCustomization(ctx workflow.Context, input *quicksight.DescribeAccountCustomizationInput) (*quicksight.DescribeAccountCustomizationOutput, error)
	DescribeAccountCustomizationAsync(ctx workflow.Context, input *quicksight.DescribeAccountCustomizationInput) *DescribeAccountCustomizationFuture

	DescribeAccountSettings(ctx workflow.Context, input *quicksight.DescribeAccountSettingsInput) (*quicksight.DescribeAccountSettingsOutput, error)
	DescribeAccountSettingsAsync(ctx workflow.Context, input *quicksight.DescribeAccountSettingsInput) *DescribeAccountSettingsFuture

	DescribeAnalysis(ctx workflow.Context, input *quicksight.DescribeAnalysisInput) (*quicksight.DescribeAnalysisOutput, error)
	DescribeAnalysisAsync(ctx workflow.Context, input *quicksight.DescribeAnalysisInput) *DescribeAnalysisFuture

	DescribeAnalysisPermissions(ctx workflow.Context, input *quicksight.DescribeAnalysisPermissionsInput) (*quicksight.DescribeAnalysisPermissionsOutput, error)
	DescribeAnalysisPermissionsAsync(ctx workflow.Context, input *quicksight.DescribeAnalysisPermissionsInput) *DescribeAnalysisPermissionsFuture

	DescribeDashboard(ctx workflow.Context, input *quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error)
	DescribeDashboardAsync(ctx workflow.Context, input *quicksight.DescribeDashboardInput) *DescribeDashboardFuture

	DescribeDashboardPermissions(ctx workflow.Context, input *quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error)
	DescribeDashboardPermissionsAsync(ctx workflow.Context, input *quicksight.DescribeDashboardPermissionsInput) *DescribeDashboardPermissionsFuture

	DescribeDataSet(ctx workflow.Context, input *quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error)
	DescribeDataSetAsync(ctx workflow.Context, input *quicksight.DescribeDataSetInput) *DescribeDataSetFuture

	DescribeDataSetPermissions(ctx workflow.Context, input *quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error)
	DescribeDataSetPermissionsAsync(ctx workflow.Context, input *quicksight.DescribeDataSetPermissionsInput) *DescribeDataSetPermissionsFuture

	DescribeDataSource(ctx workflow.Context, input *quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error)
	DescribeDataSourceAsync(ctx workflow.Context, input *quicksight.DescribeDataSourceInput) *DescribeDataSourceFuture

	DescribeDataSourcePermissions(ctx workflow.Context, input *quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error)
	DescribeDataSourcePermissionsAsync(ctx workflow.Context, input *quicksight.DescribeDataSourcePermissionsInput) *DescribeDataSourcePermissionsFuture

	DescribeGroup(ctx workflow.Context, input *quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupAsync(ctx workflow.Context, input *quicksight.DescribeGroupInput) *DescribeGroupFuture

	DescribeIAMPolicyAssignment(ctx workflow.Context, input *quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error)
	DescribeIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.DescribeIAMPolicyAssignmentInput) *DescribeIAMPolicyAssignmentFuture

	DescribeIngestion(ctx workflow.Context, input *quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error)
	DescribeIngestionAsync(ctx workflow.Context, input *quicksight.DescribeIngestionInput) *DescribeIngestionFuture

	DescribeNamespace(ctx workflow.Context, input *quicksight.DescribeNamespaceInput) (*quicksight.DescribeNamespaceOutput, error)
	DescribeNamespaceAsync(ctx workflow.Context, input *quicksight.DescribeNamespaceInput) *DescribeNamespaceFuture

	DescribeTemplate(ctx workflow.Context, input *quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error)
	DescribeTemplateAsync(ctx workflow.Context, input *quicksight.DescribeTemplateInput) *DescribeTemplateFuture

	DescribeTemplateAlias(ctx workflow.Context, input *quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error)
	DescribeTemplateAliasAsync(ctx workflow.Context, input *quicksight.DescribeTemplateAliasInput) *DescribeTemplateAliasFuture

	DescribeTemplatePermissions(ctx workflow.Context, input *quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error)
	DescribeTemplatePermissionsAsync(ctx workflow.Context, input *quicksight.DescribeTemplatePermissionsInput) *DescribeTemplatePermissionsFuture

	DescribeTheme(ctx workflow.Context, input *quicksight.DescribeThemeInput) (*quicksight.DescribeThemeOutput, error)
	DescribeThemeAsync(ctx workflow.Context, input *quicksight.DescribeThemeInput) *DescribeThemeFuture

	DescribeThemeAlias(ctx workflow.Context, input *quicksight.DescribeThemeAliasInput) (*quicksight.DescribeThemeAliasOutput, error)
	DescribeThemeAliasAsync(ctx workflow.Context, input *quicksight.DescribeThemeAliasInput) *DescribeThemeAliasFuture

	DescribeThemePermissions(ctx workflow.Context, input *quicksight.DescribeThemePermissionsInput) (*quicksight.DescribeThemePermissionsOutput, error)
	DescribeThemePermissionsAsync(ctx workflow.Context, input *quicksight.DescribeThemePermissionsInput) *DescribeThemePermissionsFuture

	DescribeUser(ctx workflow.Context, input *quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error)
	DescribeUserAsync(ctx workflow.Context, input *quicksight.DescribeUserInput) *DescribeUserFuture

	GetDashboardEmbedUrl(ctx workflow.Context, input *quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetDashboardEmbedUrlAsync(ctx workflow.Context, input *quicksight.GetDashboardEmbedUrlInput) *GetDashboardEmbedUrlFuture

	GetSessionEmbedUrl(ctx workflow.Context, input *quicksight.GetSessionEmbedUrlInput) (*quicksight.GetSessionEmbedUrlOutput, error)
	GetSessionEmbedUrlAsync(ctx workflow.Context, input *quicksight.GetSessionEmbedUrlInput) *GetSessionEmbedUrlFuture

	ListAnalyses(ctx workflow.Context, input *quicksight.ListAnalysesInput) (*quicksight.ListAnalysesOutput, error)
	ListAnalysesAsync(ctx workflow.Context, input *quicksight.ListAnalysesInput) *ListAnalysesFuture

	ListDashboardVersions(ctx workflow.Context, input *quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error)
	ListDashboardVersionsAsync(ctx workflow.Context, input *quicksight.ListDashboardVersionsInput) *ListDashboardVersionsFuture

	ListDashboards(ctx workflow.Context, input *quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error)
	ListDashboardsAsync(ctx workflow.Context, input *quicksight.ListDashboardsInput) *ListDashboardsFuture

	ListDataSets(ctx workflow.Context, input *quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error)
	ListDataSetsAsync(ctx workflow.Context, input *quicksight.ListDataSetsInput) *ListDataSetsFuture

	ListDataSources(ctx workflow.Context, input *quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error)
	ListDataSourcesAsync(ctx workflow.Context, input *quicksight.ListDataSourcesInput) *ListDataSourcesFuture

	ListGroupMemberships(ctx workflow.Context, input *quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroupMembershipsAsync(ctx workflow.Context, input *quicksight.ListGroupMembershipsInput) *ListGroupMembershipsFuture

	ListGroups(ctx workflow.Context, input *quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *quicksight.ListGroupsInput) *ListGroupsFuture

	ListIAMPolicyAssignments(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error)
	ListIAMPolicyAssignmentsAsync(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsInput) *ListIAMPolicyAssignmentsFuture

	ListIAMPolicyAssignmentsForUser(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error)
	ListIAMPolicyAssignmentsForUserAsync(ctx workflow.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput) *ListIAMPolicyAssignmentsForUserFuture

	ListIngestions(ctx workflow.Context, input *quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error)
	ListIngestionsAsync(ctx workflow.Context, input *quicksight.ListIngestionsInput) *ListIngestionsFuture

	ListNamespaces(ctx workflow.Context, input *quicksight.ListNamespacesInput) (*quicksight.ListNamespacesOutput, error)
	ListNamespacesAsync(ctx workflow.Context, input *quicksight.ListNamespacesInput) *ListNamespacesFuture

	ListTagsForResource(ctx workflow.Context, input *quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *quicksight.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTemplateAliases(ctx workflow.Context, input *quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error)
	ListTemplateAliasesAsync(ctx workflow.Context, input *quicksight.ListTemplateAliasesInput) *ListTemplateAliasesFuture

	ListTemplateVersions(ctx workflow.Context, input *quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error)
	ListTemplateVersionsAsync(ctx workflow.Context, input *quicksight.ListTemplateVersionsInput) *ListTemplateVersionsFuture

	ListTemplates(ctx workflow.Context, input *quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error)
	ListTemplatesAsync(ctx workflow.Context, input *quicksight.ListTemplatesInput) *ListTemplatesFuture

	ListThemeAliases(ctx workflow.Context, input *quicksight.ListThemeAliasesInput) (*quicksight.ListThemeAliasesOutput, error)
	ListThemeAliasesAsync(ctx workflow.Context, input *quicksight.ListThemeAliasesInput) *ListThemeAliasesFuture

	ListThemeVersions(ctx workflow.Context, input *quicksight.ListThemeVersionsInput) (*quicksight.ListThemeVersionsOutput, error)
	ListThemeVersionsAsync(ctx workflow.Context, input *quicksight.ListThemeVersionsInput) *ListThemeVersionsFuture

	ListThemes(ctx workflow.Context, input *quicksight.ListThemesInput) (*quicksight.ListThemesOutput, error)
	ListThemesAsync(ctx workflow.Context, input *quicksight.ListThemesInput) *ListThemesFuture

	ListUserGroups(ctx workflow.Context, input *quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error)
	ListUserGroupsAsync(ctx workflow.Context, input *quicksight.ListUserGroupsInput) *ListUserGroupsFuture

	ListUsers(ctx workflow.Context, input *quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error)
	ListUsersAsync(ctx workflow.Context, input *quicksight.ListUsersInput) *ListUsersFuture

	RegisterUser(ctx workflow.Context, input *quicksight.RegisterUserInput) (*quicksight.RegisterUserOutput, error)
	RegisterUserAsync(ctx workflow.Context, input *quicksight.RegisterUserInput) *RegisterUserFuture

	RestoreAnalysis(ctx workflow.Context, input *quicksight.RestoreAnalysisInput) (*quicksight.RestoreAnalysisOutput, error)
	RestoreAnalysisAsync(ctx workflow.Context, input *quicksight.RestoreAnalysisInput) *RestoreAnalysisFuture

	SearchAnalyses(ctx workflow.Context, input *quicksight.SearchAnalysesInput) (*quicksight.SearchAnalysesOutput, error)
	SearchAnalysesAsync(ctx workflow.Context, input *quicksight.SearchAnalysesInput) *SearchAnalysesFuture

	SearchDashboards(ctx workflow.Context, input *quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error)
	SearchDashboardsAsync(ctx workflow.Context, input *quicksight.SearchDashboardsInput) *SearchDashboardsFuture

	TagResource(ctx workflow.Context, input *quicksight.TagResourceInput) (*quicksight.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *quicksight.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *quicksight.UntagResourceInput) (*quicksight.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *quicksight.UntagResourceInput) *UntagResourceFuture

	UpdateAccountCustomization(ctx workflow.Context, input *quicksight.UpdateAccountCustomizationInput) (*quicksight.UpdateAccountCustomizationOutput, error)
	UpdateAccountCustomizationAsync(ctx workflow.Context, input *quicksight.UpdateAccountCustomizationInput) *UpdateAccountCustomizationFuture

	UpdateAccountSettings(ctx workflow.Context, input *quicksight.UpdateAccountSettingsInput) (*quicksight.UpdateAccountSettingsOutput, error)
	UpdateAccountSettingsAsync(ctx workflow.Context, input *quicksight.UpdateAccountSettingsInput) *UpdateAccountSettingsFuture

	UpdateAnalysis(ctx workflow.Context, input *quicksight.UpdateAnalysisInput) (*quicksight.UpdateAnalysisOutput, error)
	UpdateAnalysisAsync(ctx workflow.Context, input *quicksight.UpdateAnalysisInput) *UpdateAnalysisFuture

	UpdateAnalysisPermissions(ctx workflow.Context, input *quicksight.UpdateAnalysisPermissionsInput) (*quicksight.UpdateAnalysisPermissionsOutput, error)
	UpdateAnalysisPermissionsAsync(ctx workflow.Context, input *quicksight.UpdateAnalysisPermissionsInput) *UpdateAnalysisPermissionsFuture

	UpdateDashboard(ctx workflow.Context, input *quicksight.UpdateDashboardInput) (*quicksight.UpdateDashboardOutput, error)
	UpdateDashboardAsync(ctx workflow.Context, input *quicksight.UpdateDashboardInput) *UpdateDashboardFuture

	UpdateDashboardPermissions(ctx workflow.Context, input *quicksight.UpdateDashboardPermissionsInput) (*quicksight.UpdateDashboardPermissionsOutput, error)
	UpdateDashboardPermissionsAsync(ctx workflow.Context, input *quicksight.UpdateDashboardPermissionsInput) *UpdateDashboardPermissionsFuture

	UpdateDashboardPublishedVersion(ctx workflow.Context, input *quicksight.UpdateDashboardPublishedVersionInput) (*quicksight.UpdateDashboardPublishedVersionOutput, error)
	UpdateDashboardPublishedVersionAsync(ctx workflow.Context, input *quicksight.UpdateDashboardPublishedVersionInput) *UpdateDashboardPublishedVersionFuture

	UpdateDataSet(ctx workflow.Context, input *quicksight.UpdateDataSetInput) (*quicksight.UpdateDataSetOutput, error)
	UpdateDataSetAsync(ctx workflow.Context, input *quicksight.UpdateDataSetInput) *UpdateDataSetFuture

	UpdateDataSetPermissions(ctx workflow.Context, input *quicksight.UpdateDataSetPermissionsInput) (*quicksight.UpdateDataSetPermissionsOutput, error)
	UpdateDataSetPermissionsAsync(ctx workflow.Context, input *quicksight.UpdateDataSetPermissionsInput) *UpdateDataSetPermissionsFuture

	UpdateDataSource(ctx workflow.Context, input *quicksight.UpdateDataSourceInput) (*quicksight.UpdateDataSourceOutput, error)
	UpdateDataSourceAsync(ctx workflow.Context, input *quicksight.UpdateDataSourceInput) *UpdateDataSourceFuture

	UpdateDataSourcePermissions(ctx workflow.Context, input *quicksight.UpdateDataSourcePermissionsInput) (*quicksight.UpdateDataSourcePermissionsOutput, error)
	UpdateDataSourcePermissionsAsync(ctx workflow.Context, input *quicksight.UpdateDataSourcePermissionsInput) *UpdateDataSourcePermissionsFuture

	UpdateGroup(ctx workflow.Context, input *quicksight.UpdateGroupInput) (*quicksight.UpdateGroupOutput, error)
	UpdateGroupAsync(ctx workflow.Context, input *quicksight.UpdateGroupInput) *UpdateGroupFuture

	UpdateIAMPolicyAssignment(ctx workflow.Context, input *quicksight.UpdateIAMPolicyAssignmentInput) (*quicksight.UpdateIAMPolicyAssignmentOutput, error)
	UpdateIAMPolicyAssignmentAsync(ctx workflow.Context, input *quicksight.UpdateIAMPolicyAssignmentInput) *UpdateIAMPolicyAssignmentFuture

	UpdateTemplate(ctx workflow.Context, input *quicksight.UpdateTemplateInput) (*quicksight.UpdateTemplateOutput, error)
	UpdateTemplateAsync(ctx workflow.Context, input *quicksight.UpdateTemplateInput) *UpdateTemplateFuture

	UpdateTemplateAlias(ctx workflow.Context, input *quicksight.UpdateTemplateAliasInput) (*quicksight.UpdateTemplateAliasOutput, error)
	UpdateTemplateAliasAsync(ctx workflow.Context, input *quicksight.UpdateTemplateAliasInput) *UpdateTemplateAliasFuture

	UpdateTemplatePermissions(ctx workflow.Context, input *quicksight.UpdateTemplatePermissionsInput) (*quicksight.UpdateTemplatePermissionsOutput, error)
	UpdateTemplatePermissionsAsync(ctx workflow.Context, input *quicksight.UpdateTemplatePermissionsInput) *UpdateTemplatePermissionsFuture

	UpdateTheme(ctx workflow.Context, input *quicksight.UpdateThemeInput) (*quicksight.UpdateThemeOutput, error)
	UpdateThemeAsync(ctx workflow.Context, input *quicksight.UpdateThemeInput) *UpdateThemeFuture

	UpdateThemeAlias(ctx workflow.Context, input *quicksight.UpdateThemeAliasInput) (*quicksight.UpdateThemeAliasOutput, error)
	UpdateThemeAliasAsync(ctx workflow.Context, input *quicksight.UpdateThemeAliasInput) *UpdateThemeAliasFuture

	UpdateThemePermissions(ctx workflow.Context, input *quicksight.UpdateThemePermissionsInput) (*quicksight.UpdateThemePermissionsOutput, error)
	UpdateThemePermissionsAsync(ctx workflow.Context, input *quicksight.UpdateThemePermissionsInput) *UpdateThemePermissionsFuture

	UpdateUser(ctx workflow.Context, input *quicksight.UpdateUserInput) (*quicksight.UpdateUserOutput, error)
	UpdateUserAsync(ctx workflow.Context, input *quicksight.UpdateUserInput) *UpdateUserFuture
}

func NewClient() Client {
	return &stub{}
}
