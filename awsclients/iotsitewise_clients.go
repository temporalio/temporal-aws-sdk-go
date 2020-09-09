package awsclients

import (
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type IoTSiteWiseClient interface {
	AssociateAssets(ctx workflow.Context, input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error)
	AssociateAssetsAsync(ctx workflow.Context, input *iotsitewise.AssociateAssetsInput) *IotsitewiseAssociateAssetsResult

	BatchAssociateProjectAssets(ctx workflow.Context, input *iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error)
	BatchAssociateProjectAssetsAsync(ctx workflow.Context, input *iotsitewise.BatchAssociateProjectAssetsInput) *IotsitewiseBatchAssociateProjectAssetsResult

	BatchDisassociateProjectAssets(ctx workflow.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error)
	BatchDisassociateProjectAssetsAsync(ctx workflow.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput) *IotsitewiseBatchDisassociateProjectAssetsResult

	BatchPutAssetPropertyValue(ctx workflow.Context, input *iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error)
	BatchPutAssetPropertyValueAsync(ctx workflow.Context, input *iotsitewise.BatchPutAssetPropertyValueInput) *IotsitewiseBatchPutAssetPropertyValueResult

	CreateAccessPolicy(ctx workflow.Context, input *iotsitewise.CreateAccessPolicyInput) (*iotsitewise.CreateAccessPolicyOutput, error)
	CreateAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.CreateAccessPolicyInput) *IotsitewiseCreateAccessPolicyResult

	CreateAsset(ctx workflow.Context, input *iotsitewise.CreateAssetInput) (*iotsitewise.CreateAssetOutput, error)
	CreateAssetAsync(ctx workflow.Context, input *iotsitewise.CreateAssetInput) *IotsitewiseCreateAssetResult

	CreateAssetModel(ctx workflow.Context, input *iotsitewise.CreateAssetModelInput) (*iotsitewise.CreateAssetModelOutput, error)
	CreateAssetModelAsync(ctx workflow.Context, input *iotsitewise.CreateAssetModelInput) *IotsitewiseCreateAssetModelResult

	CreateDashboard(ctx workflow.Context, input *iotsitewise.CreateDashboardInput) (*iotsitewise.CreateDashboardOutput, error)
	CreateDashboardAsync(ctx workflow.Context, input *iotsitewise.CreateDashboardInput) *IotsitewiseCreateDashboardResult

	CreateGateway(ctx workflow.Context, input *iotsitewise.CreateGatewayInput) (*iotsitewise.CreateGatewayOutput, error)
	CreateGatewayAsync(ctx workflow.Context, input *iotsitewise.CreateGatewayInput) *IotsitewiseCreateGatewayResult

	CreatePortal(ctx workflow.Context, input *iotsitewise.CreatePortalInput) (*iotsitewise.CreatePortalOutput, error)
	CreatePortalAsync(ctx workflow.Context, input *iotsitewise.CreatePortalInput) *IotsitewiseCreatePortalResult

	CreateProject(ctx workflow.Context, input *iotsitewise.CreateProjectInput) (*iotsitewise.CreateProjectOutput, error)
	CreateProjectAsync(ctx workflow.Context, input *iotsitewise.CreateProjectInput) *IotsitewiseCreateProjectResult

	DeleteAccessPolicy(ctx workflow.Context, input *iotsitewise.DeleteAccessPolicyInput) (*iotsitewise.DeleteAccessPolicyOutput, error)
	DeleteAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.DeleteAccessPolicyInput) *IotsitewiseDeleteAccessPolicyResult

	DeleteAsset(ctx workflow.Context, input *iotsitewise.DeleteAssetInput) (*iotsitewise.DeleteAssetOutput, error)
	DeleteAssetAsync(ctx workflow.Context, input *iotsitewise.DeleteAssetInput) *IotsitewiseDeleteAssetResult

	DeleteAssetModel(ctx workflow.Context, input *iotsitewise.DeleteAssetModelInput) (*iotsitewise.DeleteAssetModelOutput, error)
	DeleteAssetModelAsync(ctx workflow.Context, input *iotsitewise.DeleteAssetModelInput) *IotsitewiseDeleteAssetModelResult

	DeleteDashboard(ctx workflow.Context, input *iotsitewise.DeleteDashboardInput) (*iotsitewise.DeleteDashboardOutput, error)
	DeleteDashboardAsync(ctx workflow.Context, input *iotsitewise.DeleteDashboardInput) *IotsitewiseDeleteDashboardResult

	DeleteGateway(ctx workflow.Context, input *iotsitewise.DeleteGatewayInput) (*iotsitewise.DeleteGatewayOutput, error)
	DeleteGatewayAsync(ctx workflow.Context, input *iotsitewise.DeleteGatewayInput) *IotsitewiseDeleteGatewayResult

	DeletePortal(ctx workflow.Context, input *iotsitewise.DeletePortalInput) (*iotsitewise.DeletePortalOutput, error)
	DeletePortalAsync(ctx workflow.Context, input *iotsitewise.DeletePortalInput) *IotsitewiseDeletePortalResult

	DeleteProject(ctx workflow.Context, input *iotsitewise.DeleteProjectInput) (*iotsitewise.DeleteProjectOutput, error)
	DeleteProjectAsync(ctx workflow.Context, input *iotsitewise.DeleteProjectInput) *IotsitewiseDeleteProjectResult

	DescribeAccessPolicy(ctx workflow.Context, input *iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error)
	DescribeAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.DescribeAccessPolicyInput) *IotsitewiseDescribeAccessPolicyResult

	DescribeAsset(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error)
	DescribeAssetAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) *IotsitewiseDescribeAssetResult

	DescribeAssetModel(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error)
	DescribeAssetModelAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) *IotsitewiseDescribeAssetModelResult

	DescribeAssetProperty(ctx workflow.Context, input *iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error)
	DescribeAssetPropertyAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetPropertyInput) *IotsitewiseDescribeAssetPropertyResult

	DescribeDashboard(ctx workflow.Context, input *iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error)
	DescribeDashboardAsync(ctx workflow.Context, input *iotsitewise.DescribeDashboardInput) *IotsitewiseDescribeDashboardResult

	DescribeGateway(ctx workflow.Context, input *iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error)
	DescribeGatewayAsync(ctx workflow.Context, input *iotsitewise.DescribeGatewayInput) *IotsitewiseDescribeGatewayResult

	DescribeGatewayCapabilityConfiguration(ctx workflow.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error)
	DescribeGatewayCapabilityConfigurationAsync(ctx workflow.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) *IotsitewiseDescribeGatewayCapabilityConfigurationResult

	DescribeLoggingOptions(ctx workflow.Context, input *iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotsitewise.DescribeLoggingOptionsInput) *IotsitewiseDescribeLoggingOptionsResult

	DescribePortal(ctx workflow.Context, input *iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error)
	DescribePortalAsync(ctx workflow.Context, input *iotsitewise.DescribePortalInput) *IotsitewiseDescribePortalResult

	DescribeProject(ctx workflow.Context, input *iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error)
	DescribeProjectAsync(ctx workflow.Context, input *iotsitewise.DescribeProjectInput) *IotsitewiseDescribeProjectResult

	DisassociateAssets(ctx workflow.Context, input *iotsitewise.DisassociateAssetsInput) (*iotsitewise.DisassociateAssetsOutput, error)
	DisassociateAssetsAsync(ctx workflow.Context, input *iotsitewise.DisassociateAssetsInput) *IotsitewiseDisassociateAssetsResult

	GetAssetPropertyAggregates(ctx workflow.Context, input *iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error)
	GetAssetPropertyAggregatesAsync(ctx workflow.Context, input *iotsitewise.GetAssetPropertyAggregatesInput) *IotsitewiseGetAssetPropertyAggregatesResult

	GetAssetPropertyValue(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error)
	GetAssetPropertyValueAsync(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueInput) *IotsitewiseGetAssetPropertyValueResult

	GetAssetPropertyValueHistory(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error)
	GetAssetPropertyValueHistoryAsync(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput) *IotsitewiseGetAssetPropertyValueHistoryResult

	ListAccessPolicies(ctx workflow.Context, input *iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error)
	ListAccessPoliciesAsync(ctx workflow.Context, input *iotsitewise.ListAccessPoliciesInput) *IotsitewiseListAccessPoliciesResult

	ListAssetModels(ctx workflow.Context, input *iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error)
	ListAssetModelsAsync(ctx workflow.Context, input *iotsitewise.ListAssetModelsInput) *IotsitewiseListAssetModelsResult

	ListAssets(ctx workflow.Context, input *iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error)
	ListAssetsAsync(ctx workflow.Context, input *iotsitewise.ListAssetsInput) *IotsitewiseListAssetsResult

	ListAssociatedAssets(ctx workflow.Context, input *iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error)
	ListAssociatedAssetsAsync(ctx workflow.Context, input *iotsitewise.ListAssociatedAssetsInput) *IotsitewiseListAssociatedAssetsResult

	ListDashboards(ctx workflow.Context, input *iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error)
	ListDashboardsAsync(ctx workflow.Context, input *iotsitewise.ListDashboardsInput) *IotsitewiseListDashboardsResult

	ListGateways(ctx workflow.Context, input *iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error)
	ListGatewaysAsync(ctx workflow.Context, input *iotsitewise.ListGatewaysInput) *IotsitewiseListGatewaysResult

	ListPortals(ctx workflow.Context, input *iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error)
	ListPortalsAsync(ctx workflow.Context, input *iotsitewise.ListPortalsInput) *IotsitewiseListPortalsResult

	ListProjectAssets(ctx workflow.Context, input *iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error)
	ListProjectAssetsAsync(ctx workflow.Context, input *iotsitewise.ListProjectAssetsInput) *IotsitewiseListProjectAssetsResult

	ListProjects(ctx workflow.Context, input *iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error)
	ListProjectsAsync(ctx workflow.Context, input *iotsitewise.ListProjectsInput) *IotsitewiseListProjectsResult

	ListTagsForResource(ctx workflow.Context, input *iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *iotsitewise.ListTagsForResourceInput) *IotsitewiseListTagsForResourceResult

	PutLoggingOptions(ctx workflow.Context, input *iotsitewise.PutLoggingOptionsInput) (*iotsitewise.PutLoggingOptionsOutput, error)
	PutLoggingOptionsAsync(ctx workflow.Context, input *iotsitewise.PutLoggingOptionsInput) *IotsitewisePutLoggingOptionsResult

	TagResource(ctx workflow.Context, input *iotsitewise.TagResourceInput) (*iotsitewise.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *iotsitewise.TagResourceInput) *IotsitewiseTagResourceResult

	UntagResource(ctx workflow.Context, input *iotsitewise.UntagResourceInput) (*iotsitewise.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *iotsitewise.UntagResourceInput) *IotsitewiseUntagResourceResult

	UpdateAccessPolicy(ctx workflow.Context, input *iotsitewise.UpdateAccessPolicyInput) (*iotsitewise.UpdateAccessPolicyOutput, error)
	UpdateAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.UpdateAccessPolicyInput) *IotsitewiseUpdateAccessPolicyResult

	UpdateAsset(ctx workflow.Context, input *iotsitewise.UpdateAssetInput) (*iotsitewise.UpdateAssetOutput, error)
	UpdateAssetAsync(ctx workflow.Context, input *iotsitewise.UpdateAssetInput) *IotsitewiseUpdateAssetResult

	UpdateAssetModel(ctx workflow.Context, input *iotsitewise.UpdateAssetModelInput) (*iotsitewise.UpdateAssetModelOutput, error)
	UpdateAssetModelAsync(ctx workflow.Context, input *iotsitewise.UpdateAssetModelInput) *IotsitewiseUpdateAssetModelResult

	UpdateAssetProperty(ctx workflow.Context, input *iotsitewise.UpdateAssetPropertyInput) (*iotsitewise.UpdateAssetPropertyOutput, error)
	UpdateAssetPropertyAsync(ctx workflow.Context, input *iotsitewise.UpdateAssetPropertyInput) *IotsitewiseUpdateAssetPropertyResult

	UpdateDashboard(ctx workflow.Context, input *iotsitewise.UpdateDashboardInput) (*iotsitewise.UpdateDashboardOutput, error)
	UpdateDashboardAsync(ctx workflow.Context, input *iotsitewise.UpdateDashboardInput) *IotsitewiseUpdateDashboardResult

	UpdateGateway(ctx workflow.Context, input *iotsitewise.UpdateGatewayInput) (*iotsitewise.UpdateGatewayOutput, error)
	UpdateGatewayAsync(ctx workflow.Context, input *iotsitewise.UpdateGatewayInput) *IotsitewiseUpdateGatewayResult

	UpdateGatewayCapabilityConfiguration(ctx workflow.Context, input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error)
	UpdateGatewayCapabilityConfigurationAsync(ctx workflow.Context, input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) *IotsitewiseUpdateGatewayCapabilityConfigurationResult

	UpdatePortal(ctx workflow.Context, input *iotsitewise.UpdatePortalInput) (*iotsitewise.UpdatePortalOutput, error)
	UpdatePortalAsync(ctx workflow.Context, input *iotsitewise.UpdatePortalInput) *IotsitewiseUpdatePortalResult

	UpdateProject(ctx workflow.Context, input *iotsitewise.UpdateProjectInput) (*iotsitewise.UpdateProjectOutput, error)
	UpdateProjectAsync(ctx workflow.Context, input *iotsitewise.UpdateProjectInput) *IotsitewiseUpdateProjectResult

	WaitUntilAssetActive(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) error
	WaitUntilAssetModelActive(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) error
	WaitUntilAssetModelNotExists(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) error
	WaitUntilAssetNotExists(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) error
	WaitUntilPortalActive(ctx workflow.Context, input *iotsitewise.DescribePortalInput) error
	WaitUntilPortalNotExists(ctx workflow.Context, input *iotsitewise.DescribePortalInput) error
}

type IotsitewiseAssociateAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseAssociateAssetsResult) Get(ctx workflow.Context) (*iotsitewise.AssociateAssetsOutput, error) {
	var output iotsitewise.AssociateAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseBatchAssociateProjectAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseBatchAssociateProjectAssetsResult) Get(ctx workflow.Context) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
	var output iotsitewise.BatchAssociateProjectAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseBatchDisassociateProjectAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseBatchDisassociateProjectAssetsResult) Get(ctx workflow.Context) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
	var output iotsitewise.BatchDisassociateProjectAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseBatchPutAssetPropertyValueResult struct {
	Result workflow.Future
}

func (r *IotsitewiseBatchPutAssetPropertyValueResult) Get(ctx workflow.Context) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
	var output iotsitewise.BatchPutAssetPropertyValueOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreateAccessPolicyResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreateAccessPolicyResult) Get(ctx workflow.Context) (*iotsitewise.CreateAccessPolicyOutput, error) {
	var output iotsitewise.CreateAccessPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreateAssetResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreateAssetResult) Get(ctx workflow.Context) (*iotsitewise.CreateAssetOutput, error) {
	var output iotsitewise.CreateAssetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreateAssetModelResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreateAssetModelResult) Get(ctx workflow.Context) (*iotsitewise.CreateAssetModelOutput, error) {
	var output iotsitewise.CreateAssetModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreateDashboardResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreateDashboardResult) Get(ctx workflow.Context) (*iotsitewise.CreateDashboardOutput, error) {
	var output iotsitewise.CreateDashboardOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreateGatewayResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreateGatewayResult) Get(ctx workflow.Context) (*iotsitewise.CreateGatewayOutput, error) {
	var output iotsitewise.CreateGatewayOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreatePortalResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreatePortalResult) Get(ctx workflow.Context) (*iotsitewise.CreatePortalOutput, error) {
	var output iotsitewise.CreatePortalOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseCreateProjectResult struct {
	Result workflow.Future
}

func (r *IotsitewiseCreateProjectResult) Get(ctx workflow.Context) (*iotsitewise.CreateProjectOutput, error) {
	var output iotsitewise.CreateProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeleteAccessPolicyResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeleteAccessPolicyResult) Get(ctx workflow.Context) (*iotsitewise.DeleteAccessPolicyOutput, error) {
	var output iotsitewise.DeleteAccessPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeleteAssetResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeleteAssetResult) Get(ctx workflow.Context) (*iotsitewise.DeleteAssetOutput, error) {
	var output iotsitewise.DeleteAssetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeleteAssetModelResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeleteAssetModelResult) Get(ctx workflow.Context) (*iotsitewise.DeleteAssetModelOutput, error) {
	var output iotsitewise.DeleteAssetModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeleteDashboardResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeleteDashboardResult) Get(ctx workflow.Context) (*iotsitewise.DeleteDashboardOutput, error) {
	var output iotsitewise.DeleteDashboardOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeleteGatewayResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeleteGatewayResult) Get(ctx workflow.Context) (*iotsitewise.DeleteGatewayOutput, error) {
	var output iotsitewise.DeleteGatewayOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeletePortalResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeletePortalResult) Get(ctx workflow.Context) (*iotsitewise.DeletePortalOutput, error) {
	var output iotsitewise.DeletePortalOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDeleteProjectResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDeleteProjectResult) Get(ctx workflow.Context) (*iotsitewise.DeleteProjectOutput, error) {
	var output iotsitewise.DeleteProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeAccessPolicyResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeAccessPolicyResult) Get(ctx workflow.Context) (*iotsitewise.DescribeAccessPolicyOutput, error) {
	var output iotsitewise.DescribeAccessPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeAssetResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeAssetResult) Get(ctx workflow.Context) (*iotsitewise.DescribeAssetOutput, error) {
	var output iotsitewise.DescribeAssetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeAssetModelResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeAssetModelResult) Get(ctx workflow.Context) (*iotsitewise.DescribeAssetModelOutput, error) {
	var output iotsitewise.DescribeAssetModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeAssetPropertyResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeAssetPropertyResult) Get(ctx workflow.Context) (*iotsitewise.DescribeAssetPropertyOutput, error) {
	var output iotsitewise.DescribeAssetPropertyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeDashboardResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeDashboardResult) Get(ctx workflow.Context) (*iotsitewise.DescribeDashboardOutput, error) {
	var output iotsitewise.DescribeDashboardOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeGatewayResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeGatewayResult) Get(ctx workflow.Context) (*iotsitewise.DescribeGatewayOutput, error) {
	var output iotsitewise.DescribeGatewayOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeGatewayCapabilityConfigurationResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeGatewayCapabilityConfigurationResult) Get(ctx workflow.Context) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
	var output iotsitewise.DescribeGatewayCapabilityConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeLoggingOptionsResult) Get(ctx workflow.Context) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
	var output iotsitewise.DescribeLoggingOptionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribePortalResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribePortalResult) Get(ctx workflow.Context) (*iotsitewise.DescribePortalOutput, error) {
	var output iotsitewise.DescribePortalOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDescribeProjectResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDescribeProjectResult) Get(ctx workflow.Context) (*iotsitewise.DescribeProjectOutput, error) {
	var output iotsitewise.DescribeProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseDisassociateAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseDisassociateAssetsResult) Get(ctx workflow.Context) (*iotsitewise.DisassociateAssetsOutput, error) {
	var output iotsitewise.DisassociateAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseGetAssetPropertyAggregatesResult struct {
	Result workflow.Future
}

func (r *IotsitewiseGetAssetPropertyAggregatesResult) Get(ctx workflow.Context) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
	var output iotsitewise.GetAssetPropertyAggregatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseGetAssetPropertyValueResult struct {
	Result workflow.Future
}

func (r *IotsitewiseGetAssetPropertyValueResult) Get(ctx workflow.Context) (*iotsitewise.GetAssetPropertyValueOutput, error) {
	var output iotsitewise.GetAssetPropertyValueOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseGetAssetPropertyValueHistoryResult struct {
	Result workflow.Future
}

func (r *IotsitewiseGetAssetPropertyValueHistoryResult) Get(ctx workflow.Context) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
	var output iotsitewise.GetAssetPropertyValueHistoryOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListAccessPoliciesResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListAccessPoliciesResult) Get(ctx workflow.Context) (*iotsitewise.ListAccessPoliciesOutput, error) {
	var output iotsitewise.ListAccessPoliciesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListAssetModelsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListAssetModelsResult) Get(ctx workflow.Context) (*iotsitewise.ListAssetModelsOutput, error) {
	var output iotsitewise.ListAssetModelsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListAssetsResult) Get(ctx workflow.Context) (*iotsitewise.ListAssetsOutput, error) {
	var output iotsitewise.ListAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListAssociatedAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListAssociatedAssetsResult) Get(ctx workflow.Context) (*iotsitewise.ListAssociatedAssetsOutput, error) {
	var output iotsitewise.ListAssociatedAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListDashboardsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListDashboardsResult) Get(ctx workflow.Context) (*iotsitewise.ListDashboardsOutput, error) {
	var output iotsitewise.ListDashboardsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListGatewaysResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListGatewaysResult) Get(ctx workflow.Context) (*iotsitewise.ListGatewaysOutput, error) {
	var output iotsitewise.ListGatewaysOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListPortalsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListPortalsResult) Get(ctx workflow.Context) (*iotsitewise.ListPortalsOutput, error) {
	var output iotsitewise.ListPortalsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListProjectAssetsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListProjectAssetsResult) Get(ctx workflow.Context) (*iotsitewise.ListProjectAssetsOutput, error) {
	var output iotsitewise.ListProjectAssetsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListProjectsResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListProjectsResult) Get(ctx workflow.Context) (*iotsitewise.ListProjectsOutput, error) {
	var output iotsitewise.ListProjectsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *IotsitewiseListTagsForResourceResult) Get(ctx workflow.Context) (*iotsitewise.ListTagsForResourceOutput, error) {
	var output iotsitewise.ListTagsForResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewisePutLoggingOptionsResult struct {
	Result workflow.Future
}

func (r *IotsitewisePutLoggingOptionsResult) Get(ctx workflow.Context) (*iotsitewise.PutLoggingOptionsOutput, error) {
	var output iotsitewise.PutLoggingOptionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseTagResourceResult struct {
	Result workflow.Future
}

func (r *IotsitewiseTagResourceResult) Get(ctx workflow.Context) (*iotsitewise.TagResourceOutput, error) {
	var output iotsitewise.TagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUntagResourceResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUntagResourceResult) Get(ctx workflow.Context) (*iotsitewise.UntagResourceOutput, error) {
	var output iotsitewise.UntagResourceOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateAccessPolicyResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateAccessPolicyResult) Get(ctx workflow.Context) (*iotsitewise.UpdateAccessPolicyOutput, error) {
	var output iotsitewise.UpdateAccessPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateAssetResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateAssetResult) Get(ctx workflow.Context) (*iotsitewise.UpdateAssetOutput, error) {
	var output iotsitewise.UpdateAssetOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateAssetModelResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateAssetModelResult) Get(ctx workflow.Context) (*iotsitewise.UpdateAssetModelOutput, error) {
	var output iotsitewise.UpdateAssetModelOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateAssetPropertyResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateAssetPropertyResult) Get(ctx workflow.Context) (*iotsitewise.UpdateAssetPropertyOutput, error) {
	var output iotsitewise.UpdateAssetPropertyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateDashboardResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateDashboardResult) Get(ctx workflow.Context) (*iotsitewise.UpdateDashboardOutput, error) {
	var output iotsitewise.UpdateDashboardOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateGatewayResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateGatewayResult) Get(ctx workflow.Context) (*iotsitewise.UpdateGatewayOutput, error) {
	var output iotsitewise.UpdateGatewayOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateGatewayCapabilityConfigurationResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateGatewayCapabilityConfigurationResult) Get(ctx workflow.Context) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error) {
	var output iotsitewise.UpdateGatewayCapabilityConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdatePortalResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdatePortalResult) Get(ctx workflow.Context) (*iotsitewise.UpdatePortalOutput, error) {
	var output iotsitewise.UpdatePortalOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IotsitewiseUpdateProjectResult struct {
	Result workflow.Future
}

func (r *IotsitewiseUpdateProjectResult) Get(ctx workflow.Context) (*iotsitewise.UpdateProjectOutput, error) {
	var output iotsitewise.UpdateProjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type IoTSiteWiseStub struct {
	activities awsactivities.IoTSiteWiseActivities
}

func NewIoTSiteWiseStub() IoTSiteWiseClient {
	return &IoTSiteWiseStub{}
}

func (a *IoTSiteWiseStub) AssociateAssets(ctx workflow.Context, input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error) {
	var output iotsitewise.AssociateAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AssociateAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) AssociateAssetsAsync(ctx workflow.Context, input *iotsitewise.AssociateAssetsInput) *IotsitewiseAssociateAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AssociateAssets, input)
	return &IotsitewiseAssociateAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) BatchAssociateProjectAssets(ctx workflow.Context, input *iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
	var output iotsitewise.BatchAssociateProjectAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchAssociateProjectAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) BatchAssociateProjectAssetsAsync(ctx workflow.Context, input *iotsitewise.BatchAssociateProjectAssetsInput) *IotsitewiseBatchAssociateProjectAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchAssociateProjectAssets, input)
	return &IotsitewiseBatchAssociateProjectAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) BatchDisassociateProjectAssets(ctx workflow.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
	var output iotsitewise.BatchDisassociateProjectAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchDisassociateProjectAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) BatchDisassociateProjectAssetsAsync(ctx workflow.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput) *IotsitewiseBatchDisassociateProjectAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchDisassociateProjectAssets, input)
	return &IotsitewiseBatchDisassociateProjectAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) BatchPutAssetPropertyValue(ctx workflow.Context, input *iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
	var output iotsitewise.BatchPutAssetPropertyValueOutput
	err := workflow.ExecuteActivity(ctx, a.activities.BatchPutAssetPropertyValue, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) BatchPutAssetPropertyValueAsync(ctx workflow.Context, input *iotsitewise.BatchPutAssetPropertyValueInput) *IotsitewiseBatchPutAssetPropertyValueResult {
	future := workflow.ExecuteActivity(ctx, a.activities.BatchPutAssetPropertyValue, input)
	return &IotsitewiseBatchPutAssetPropertyValueResult{Result: future}
}

func (a *IoTSiteWiseStub) CreateAccessPolicy(ctx workflow.Context, input *iotsitewise.CreateAccessPolicyInput) (*iotsitewise.CreateAccessPolicyOutput, error) {
	var output iotsitewise.CreateAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAccessPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreateAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.CreateAccessPolicyInput) *IotsitewiseCreateAccessPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAccessPolicy, input)
	return &IotsitewiseCreateAccessPolicyResult{Result: future}
}

func (a *IoTSiteWiseStub) CreateAsset(ctx workflow.Context, input *iotsitewise.CreateAssetInput) (*iotsitewise.CreateAssetOutput, error) {
	var output iotsitewise.CreateAssetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAsset, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreateAssetAsync(ctx workflow.Context, input *iotsitewise.CreateAssetInput) *IotsitewiseCreateAssetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAsset, input)
	return &IotsitewiseCreateAssetResult{Result: future}
}

func (a *IoTSiteWiseStub) CreateAssetModel(ctx workflow.Context, input *iotsitewise.CreateAssetModelInput) (*iotsitewise.CreateAssetModelOutput, error) {
	var output iotsitewise.CreateAssetModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAssetModel, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreateAssetModelAsync(ctx workflow.Context, input *iotsitewise.CreateAssetModelInput) *IotsitewiseCreateAssetModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAssetModel, input)
	return &IotsitewiseCreateAssetModelResult{Result: future}
}

func (a *IoTSiteWiseStub) CreateDashboard(ctx workflow.Context, input *iotsitewise.CreateDashboardInput) (*iotsitewise.CreateDashboardOutput, error) {
	var output iotsitewise.CreateDashboardOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateDashboard, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreateDashboardAsync(ctx workflow.Context, input *iotsitewise.CreateDashboardInput) *IotsitewiseCreateDashboardResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateDashboard, input)
	return &IotsitewiseCreateDashboardResult{Result: future}
}

func (a *IoTSiteWiseStub) CreateGateway(ctx workflow.Context, input *iotsitewise.CreateGatewayInput) (*iotsitewise.CreateGatewayOutput, error) {
	var output iotsitewise.CreateGatewayOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateGateway, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreateGatewayAsync(ctx workflow.Context, input *iotsitewise.CreateGatewayInput) *IotsitewiseCreateGatewayResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateGateway, input)
	return &IotsitewiseCreateGatewayResult{Result: future}
}

func (a *IoTSiteWiseStub) CreatePortal(ctx workflow.Context, input *iotsitewise.CreatePortalInput) (*iotsitewise.CreatePortalOutput, error) {
	var output iotsitewise.CreatePortalOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreatePortal, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreatePortalAsync(ctx workflow.Context, input *iotsitewise.CreatePortalInput) *IotsitewiseCreatePortalResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreatePortal, input)
	return &IotsitewiseCreatePortalResult{Result: future}
}

func (a *IoTSiteWiseStub) CreateProject(ctx workflow.Context, input *iotsitewise.CreateProjectInput) (*iotsitewise.CreateProjectOutput, error) {
	var output iotsitewise.CreateProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) CreateProjectAsync(ctx workflow.Context, input *iotsitewise.CreateProjectInput) *IotsitewiseCreateProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input)
	return &IotsitewiseCreateProjectResult{Result: future}
}

func (a *IoTSiteWiseStub) DeleteAccessPolicy(ctx workflow.Context, input *iotsitewise.DeleteAccessPolicyInput) (*iotsitewise.DeleteAccessPolicyOutput, error) {
	var output iotsitewise.DeleteAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeleteAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.DeleteAccessPolicyInput) *IotsitewiseDeleteAccessPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPolicy, input)
	return &IotsitewiseDeleteAccessPolicyResult{Result: future}
}

func (a *IoTSiteWiseStub) DeleteAsset(ctx workflow.Context, input *iotsitewise.DeleteAssetInput) (*iotsitewise.DeleteAssetOutput, error) {
	var output iotsitewise.DeleteAssetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAsset, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeleteAssetAsync(ctx workflow.Context, input *iotsitewise.DeleteAssetInput) *IotsitewiseDeleteAssetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAsset, input)
	return &IotsitewiseDeleteAssetResult{Result: future}
}

func (a *IoTSiteWiseStub) DeleteAssetModel(ctx workflow.Context, input *iotsitewise.DeleteAssetModelInput) (*iotsitewise.DeleteAssetModelOutput, error) {
	var output iotsitewise.DeleteAssetModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteAssetModel, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeleteAssetModelAsync(ctx workflow.Context, input *iotsitewise.DeleteAssetModelInput) *IotsitewiseDeleteAssetModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteAssetModel, input)
	return &IotsitewiseDeleteAssetModelResult{Result: future}
}

func (a *IoTSiteWiseStub) DeleteDashboard(ctx workflow.Context, input *iotsitewise.DeleteDashboardInput) (*iotsitewise.DeleteDashboardOutput, error) {
	var output iotsitewise.DeleteDashboardOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteDashboard, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeleteDashboardAsync(ctx workflow.Context, input *iotsitewise.DeleteDashboardInput) *IotsitewiseDeleteDashboardResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteDashboard, input)
	return &IotsitewiseDeleteDashboardResult{Result: future}
}

func (a *IoTSiteWiseStub) DeleteGateway(ctx workflow.Context, input *iotsitewise.DeleteGatewayInput) (*iotsitewise.DeleteGatewayOutput, error) {
	var output iotsitewise.DeleteGatewayOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteGateway, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeleteGatewayAsync(ctx workflow.Context, input *iotsitewise.DeleteGatewayInput) *IotsitewiseDeleteGatewayResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteGateway, input)
	return &IotsitewiseDeleteGatewayResult{Result: future}
}

func (a *IoTSiteWiseStub) DeletePortal(ctx workflow.Context, input *iotsitewise.DeletePortalInput) (*iotsitewise.DeletePortalOutput, error) {
	var output iotsitewise.DeletePortalOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeletePortal, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeletePortalAsync(ctx workflow.Context, input *iotsitewise.DeletePortalInput) *IotsitewiseDeletePortalResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeletePortal, input)
	return &IotsitewiseDeletePortalResult{Result: future}
}

func (a *IoTSiteWiseStub) DeleteProject(ctx workflow.Context, input *iotsitewise.DeleteProjectInput) (*iotsitewise.DeleteProjectOutput, error) {
	var output iotsitewise.DeleteProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DeleteProjectAsync(ctx workflow.Context, input *iotsitewise.DeleteProjectInput) *IotsitewiseDeleteProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input)
	return &IotsitewiseDeleteProjectResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeAccessPolicy(ctx workflow.Context, input *iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error) {
	var output iotsitewise.DescribeAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccessPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.DescribeAccessPolicyInput) *IotsitewiseDescribeAccessPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccessPolicy, input)
	return &IotsitewiseDescribeAccessPolicyResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeAsset(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error) {
	var output iotsitewise.DescribeAssetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAsset, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeAssetAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) *IotsitewiseDescribeAssetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAsset, input)
	return &IotsitewiseDescribeAssetResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeAssetModel(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error) {
	var output iotsitewise.DescribeAssetModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssetModel, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeAssetModelAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) *IotsitewiseDescribeAssetModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAssetModel, input)
	return &IotsitewiseDescribeAssetModelResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeAssetProperty(ctx workflow.Context, input *iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error) {
	var output iotsitewise.DescribeAssetPropertyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAssetProperty, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeAssetPropertyAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetPropertyInput) *IotsitewiseDescribeAssetPropertyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAssetProperty, input)
	return &IotsitewiseDescribeAssetPropertyResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeDashboard(ctx workflow.Context, input *iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error) {
	var output iotsitewise.DescribeDashboardOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeDashboard, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeDashboardAsync(ctx workflow.Context, input *iotsitewise.DescribeDashboardInput) *IotsitewiseDescribeDashboardResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeDashboard, input)
	return &IotsitewiseDescribeDashboardResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeGateway(ctx workflow.Context, input *iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error) {
	var output iotsitewise.DescribeGatewayOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeGateway, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeGatewayAsync(ctx workflow.Context, input *iotsitewise.DescribeGatewayInput) *IotsitewiseDescribeGatewayResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeGateway, input)
	return &IotsitewiseDescribeGatewayResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeGatewayCapabilityConfiguration(ctx workflow.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
	var output iotsitewise.DescribeGatewayCapabilityConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeGatewayCapabilityConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeGatewayCapabilityConfigurationAsync(ctx workflow.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) *IotsitewiseDescribeGatewayCapabilityConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeGatewayCapabilityConfiguration, input)
	return &IotsitewiseDescribeGatewayCapabilityConfigurationResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeLoggingOptions(ctx workflow.Context, input *iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
	var output iotsitewise.DescribeLoggingOptionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoggingOptions, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotsitewise.DescribeLoggingOptionsInput) *IotsitewiseDescribeLoggingOptionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoggingOptions, input)
	return &IotsitewiseDescribeLoggingOptionsResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribePortal(ctx workflow.Context, input *iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error) {
	var output iotsitewise.DescribePortalOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribePortal, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribePortalAsync(ctx workflow.Context, input *iotsitewise.DescribePortalInput) *IotsitewiseDescribePortalResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribePortal, input)
	return &IotsitewiseDescribePortalResult{Result: future}
}

func (a *IoTSiteWiseStub) DescribeProject(ctx workflow.Context, input *iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error) {
	var output iotsitewise.DescribeProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DescribeProjectAsync(ctx workflow.Context, input *iotsitewise.DescribeProjectInput) *IotsitewiseDescribeProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input)
	return &IotsitewiseDescribeProjectResult{Result: future}
}

func (a *IoTSiteWiseStub) DisassociateAssets(ctx workflow.Context, input *iotsitewise.DisassociateAssetsInput) (*iotsitewise.DisassociateAssetsOutput, error) {
	var output iotsitewise.DisassociateAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DisassociateAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) DisassociateAssetsAsync(ctx workflow.Context, input *iotsitewise.DisassociateAssetsInput) *IotsitewiseDisassociateAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DisassociateAssets, input)
	return &IotsitewiseDisassociateAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) GetAssetPropertyAggregates(ctx workflow.Context, input *iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
	var output iotsitewise.GetAssetPropertyAggregatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAssetPropertyAggregates, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) GetAssetPropertyAggregatesAsync(ctx workflow.Context, input *iotsitewise.GetAssetPropertyAggregatesInput) *IotsitewiseGetAssetPropertyAggregatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAssetPropertyAggregates, input)
	return &IotsitewiseGetAssetPropertyAggregatesResult{Result: future}
}

func (a *IoTSiteWiseStub) GetAssetPropertyValue(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error) {
	var output iotsitewise.GetAssetPropertyValueOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAssetPropertyValue, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) GetAssetPropertyValueAsync(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueInput) *IotsitewiseGetAssetPropertyValueResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAssetPropertyValue, input)
	return &IotsitewiseGetAssetPropertyValueResult{Result: future}
}

func (a *IoTSiteWiseStub) GetAssetPropertyValueHistory(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
	var output iotsitewise.GetAssetPropertyValueHistoryOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetAssetPropertyValueHistory, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) GetAssetPropertyValueHistoryAsync(ctx workflow.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput) *IotsitewiseGetAssetPropertyValueHistoryResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetAssetPropertyValueHistory, input)
	return &IotsitewiseGetAssetPropertyValueHistoryResult{Result: future}
}

func (a *IoTSiteWiseStub) ListAccessPolicies(ctx workflow.Context, input *iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error) {
	var output iotsitewise.ListAccessPoliciesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListAccessPolicies, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListAccessPoliciesAsync(ctx workflow.Context, input *iotsitewise.ListAccessPoliciesInput) *IotsitewiseListAccessPoliciesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListAccessPolicies, input)
	return &IotsitewiseListAccessPoliciesResult{Result: future}
}

func (a *IoTSiteWiseStub) ListAssetModels(ctx workflow.Context, input *iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error) {
	var output iotsitewise.ListAssetModelsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListAssetModels, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListAssetModelsAsync(ctx workflow.Context, input *iotsitewise.ListAssetModelsInput) *IotsitewiseListAssetModelsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListAssetModels, input)
	return &IotsitewiseListAssetModelsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListAssets(ctx workflow.Context, input *iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error) {
	var output iotsitewise.ListAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListAssetsAsync(ctx workflow.Context, input *iotsitewise.ListAssetsInput) *IotsitewiseListAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListAssets, input)
	return &IotsitewiseListAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListAssociatedAssets(ctx workflow.Context, input *iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error) {
	var output iotsitewise.ListAssociatedAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListAssociatedAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListAssociatedAssetsAsync(ctx workflow.Context, input *iotsitewise.ListAssociatedAssetsInput) *IotsitewiseListAssociatedAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListAssociatedAssets, input)
	return &IotsitewiseListAssociatedAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListDashboards(ctx workflow.Context, input *iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error) {
	var output iotsitewise.ListDashboardsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListDashboards, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListDashboardsAsync(ctx workflow.Context, input *iotsitewise.ListDashboardsInput) *IotsitewiseListDashboardsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListDashboards, input)
	return &IotsitewiseListDashboardsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListGateways(ctx workflow.Context, input *iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error) {
	var output iotsitewise.ListGatewaysOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListGateways, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListGatewaysAsync(ctx workflow.Context, input *iotsitewise.ListGatewaysInput) *IotsitewiseListGatewaysResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListGateways, input)
	return &IotsitewiseListGatewaysResult{Result: future}
}

func (a *IoTSiteWiseStub) ListPortals(ctx workflow.Context, input *iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error) {
	var output iotsitewise.ListPortalsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListPortals, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListPortalsAsync(ctx workflow.Context, input *iotsitewise.ListPortalsInput) *IotsitewiseListPortalsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListPortals, input)
	return &IotsitewiseListPortalsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListProjectAssets(ctx workflow.Context, input *iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error) {
	var output iotsitewise.ListProjectAssetsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListProjectAssets, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListProjectAssetsAsync(ctx workflow.Context, input *iotsitewise.ListProjectAssetsInput) *IotsitewiseListProjectAssetsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListProjectAssets, input)
	return &IotsitewiseListProjectAssetsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListProjects(ctx workflow.Context, input *iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error) {
	var output iotsitewise.ListProjectsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListProjectsAsync(ctx workflow.Context, input *iotsitewise.ListProjectsInput) *IotsitewiseListProjectsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input)
	return &IotsitewiseListProjectsResult{Result: future}
}

func (a *IoTSiteWiseStub) ListTagsForResource(ctx workflow.Context, input *iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error) {
	var output iotsitewise.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) ListTagsForResourceAsync(ctx workflow.Context, input *iotsitewise.ListTagsForResourceInput) *IotsitewiseListTagsForResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
	return &IotsitewiseListTagsForResourceResult{Result: future}
}

func (a *IoTSiteWiseStub) PutLoggingOptions(ctx workflow.Context, input *iotsitewise.PutLoggingOptionsInput) (*iotsitewise.PutLoggingOptionsOutput, error) {
	var output iotsitewise.PutLoggingOptionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutLoggingOptions, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) PutLoggingOptionsAsync(ctx workflow.Context, input *iotsitewise.PutLoggingOptionsInput) *IotsitewisePutLoggingOptionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutLoggingOptions, input)
	return &IotsitewisePutLoggingOptionsResult{Result: future}
}

func (a *IoTSiteWiseStub) TagResource(ctx workflow.Context, input *iotsitewise.TagResourceInput) (*iotsitewise.TagResourceOutput, error) {
	var output iotsitewise.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) TagResourceAsync(ctx workflow.Context, input *iotsitewise.TagResourceInput) *IotsitewiseTagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
	return &IotsitewiseTagResourceResult{Result: future}
}

func (a *IoTSiteWiseStub) UntagResource(ctx workflow.Context, input *iotsitewise.UntagResourceInput) (*iotsitewise.UntagResourceOutput, error) {
	var output iotsitewise.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UntagResourceAsync(ctx workflow.Context, input *iotsitewise.UntagResourceInput) *IotsitewiseUntagResourceResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
	return &IotsitewiseUntagResourceResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateAccessPolicy(ctx workflow.Context, input *iotsitewise.UpdateAccessPolicyInput) (*iotsitewise.UpdateAccessPolicyOutput, error) {
	var output iotsitewise.UpdateAccessPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccessPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateAccessPolicyAsync(ctx workflow.Context, input *iotsitewise.UpdateAccessPolicyInput) *IotsitewiseUpdateAccessPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAccessPolicy, input)
	return &IotsitewiseUpdateAccessPolicyResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateAsset(ctx workflow.Context, input *iotsitewise.UpdateAssetInput) (*iotsitewise.UpdateAssetOutput, error) {
	var output iotsitewise.UpdateAssetOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAsset, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateAssetAsync(ctx workflow.Context, input *iotsitewise.UpdateAssetInput) *IotsitewiseUpdateAssetResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAsset, input)
	return &IotsitewiseUpdateAssetResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateAssetModel(ctx workflow.Context, input *iotsitewise.UpdateAssetModelInput) (*iotsitewise.UpdateAssetModelOutput, error) {
	var output iotsitewise.UpdateAssetModelOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAssetModel, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateAssetModelAsync(ctx workflow.Context, input *iotsitewise.UpdateAssetModelInput) *IotsitewiseUpdateAssetModelResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAssetModel, input)
	return &IotsitewiseUpdateAssetModelResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateAssetProperty(ctx workflow.Context, input *iotsitewise.UpdateAssetPropertyInput) (*iotsitewise.UpdateAssetPropertyOutput, error) {
	var output iotsitewise.UpdateAssetPropertyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateAssetProperty, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateAssetPropertyAsync(ctx workflow.Context, input *iotsitewise.UpdateAssetPropertyInput) *IotsitewiseUpdateAssetPropertyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateAssetProperty, input)
	return &IotsitewiseUpdateAssetPropertyResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateDashboard(ctx workflow.Context, input *iotsitewise.UpdateDashboardInput) (*iotsitewise.UpdateDashboardOutput, error) {
	var output iotsitewise.UpdateDashboardOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateDashboard, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateDashboardAsync(ctx workflow.Context, input *iotsitewise.UpdateDashboardInput) *IotsitewiseUpdateDashboardResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateDashboard, input)
	return &IotsitewiseUpdateDashboardResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateGateway(ctx workflow.Context, input *iotsitewise.UpdateGatewayInput) (*iotsitewise.UpdateGatewayOutput, error) {
	var output iotsitewise.UpdateGatewayOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateGateway, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateGatewayAsync(ctx workflow.Context, input *iotsitewise.UpdateGatewayInput) *IotsitewiseUpdateGatewayResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateGateway, input)
	return &IotsitewiseUpdateGatewayResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateGatewayCapabilityConfiguration(ctx workflow.Context, input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error) {
	var output iotsitewise.UpdateGatewayCapabilityConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayCapabilityConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateGatewayCapabilityConfigurationAsync(ctx workflow.Context, input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) *IotsitewiseUpdateGatewayCapabilityConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateGatewayCapabilityConfiguration, input)
	return &IotsitewiseUpdateGatewayCapabilityConfigurationResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdatePortal(ctx workflow.Context, input *iotsitewise.UpdatePortalInput) (*iotsitewise.UpdatePortalOutput, error) {
	var output iotsitewise.UpdatePortalOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdatePortal, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdatePortalAsync(ctx workflow.Context, input *iotsitewise.UpdatePortalInput) *IotsitewiseUpdatePortalResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdatePortal, input)
	return &IotsitewiseUpdatePortalResult{Result: future}
}

func (a *IoTSiteWiseStub) UpdateProject(ctx workflow.Context, input *iotsitewise.UpdateProjectInput) (*iotsitewise.UpdateProjectOutput, error) {
	var output iotsitewise.UpdateProjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input).Get(ctx, &output)
	return &output, err
}

func (a *IoTSiteWiseStub) UpdateProjectAsync(ctx workflow.Context, input *iotsitewise.UpdateProjectInput) *IotsitewiseUpdateProjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input)
	return &IotsitewiseUpdateProjectResult{Result: future}
}

func (a *IoTSiteWiseStub) WaitUntilAssetActive(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetActive, input).Get(ctx, nil)
}

func (a *IoTSiteWiseStub) WaitUntilAssetActiveAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetActive, input)
}

func (a *IoTSiteWiseStub) WaitUntilAssetModelActive(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetModelActive, input).Get(ctx, nil)
}

func (a *IoTSiteWiseStub) WaitUntilAssetModelActiveAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetModelActive, input)
}

func (a *IoTSiteWiseStub) WaitUntilAssetModelNotExists(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetModelNotExists, input).Get(ctx, nil)
}

func (a *IoTSiteWiseStub) WaitUntilAssetModelNotExistsAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetModelInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetModelNotExists, input)
}

func (a *IoTSiteWiseStub) WaitUntilAssetNotExists(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetNotExists, input).Get(ctx, nil)
}

func (a *IoTSiteWiseStub) WaitUntilAssetNotExistsAsync(ctx workflow.Context, input *iotsitewise.DescribeAssetInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilAssetNotExists, input)
}

func (a *IoTSiteWiseStub) WaitUntilPortalActive(ctx workflow.Context, input *iotsitewise.DescribePortalInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilPortalActive, input).Get(ctx, nil)
}

func (a *IoTSiteWiseStub) WaitUntilPortalActiveAsync(ctx workflow.Context, input *iotsitewise.DescribePortalInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilPortalActive, input)
}

func (a *IoTSiteWiseStub) WaitUntilPortalNotExists(ctx workflow.Context, input *iotsitewise.DescribePortalInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilPortalNotExists, input).Get(ctx, nil)
}

func (a *IoTSiteWiseStub) WaitUntilPortalNotExistsAsync(ctx workflow.Context, input *iotsitewise.DescribePortalInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilPortalNotExists, input)
}
