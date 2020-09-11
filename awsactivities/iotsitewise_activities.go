
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/aws/aws-sdk-go/service/iotsitewise/iotsitewiseiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTSiteWiseActivities struct {
    client iotsitewiseiface.IoTSiteWiseAPI
}

func NewIoTSiteWiseActivities(session *session.Session, config ...*aws.Config) *IoTSiteWiseActivities {
    client := iotsitewise.New(session, config...)
    return &IoTSiteWiseActivities{client: client}
}

func (a *IoTSiteWiseActivities) AssociateAssets(ctx context.Context, input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.AssociateAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) BatchAssociateProjectAssets(ctx context.Context, input *iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.BatchAssociateProjectAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) BatchDisassociateProjectAssets(ctx context.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.BatchDisassociateProjectAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) BatchPutAssetPropertyValue(ctx context.Context, input *iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
    return a.client.BatchPutAssetPropertyValueWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateAccessPolicy(ctx context.Context, input *iotsitewise.CreateAccessPolicyInput) (*iotsitewise.CreateAccessPolicyOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateAsset(ctx context.Context, input *iotsitewise.CreateAssetInput) (*iotsitewise.CreateAssetOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateAssetModel(ctx context.Context, input *iotsitewise.CreateAssetModelInput) (*iotsitewise.CreateAssetModelOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateDashboard(ctx context.Context, input *iotsitewise.CreateDashboardInput) (*iotsitewise.CreateDashboardOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateGateway(ctx context.Context, input *iotsitewise.CreateGatewayInput) (*iotsitewise.CreateGatewayOutput, error) {
    return a.client.CreateGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreatePortal(ctx context.Context, input *iotsitewise.CreatePortalInput) (*iotsitewise.CreatePortalOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreatePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) CreateProject(ctx context.Context, input *iotsitewise.CreateProjectInput) (*iotsitewise.CreateProjectOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.CreateProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteAccessPolicy(ctx context.Context, input *iotsitewise.DeleteAccessPolicyInput) (*iotsitewise.DeleteAccessPolicyOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeleteAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteAsset(ctx context.Context, input *iotsitewise.DeleteAssetInput) (*iotsitewise.DeleteAssetOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeleteAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteAssetModel(ctx context.Context, input *iotsitewise.DeleteAssetModelInput) (*iotsitewise.DeleteAssetModelOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeleteAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteDashboard(ctx context.Context, input *iotsitewise.DeleteDashboardInput) (*iotsitewise.DeleteDashboardOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeleteDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteGateway(ctx context.Context, input *iotsitewise.DeleteGatewayInput) (*iotsitewise.DeleteGatewayOutput, error) {
    return a.client.DeleteGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeletePortal(ctx context.Context, input *iotsitewise.DeletePortalInput) (*iotsitewise.DeletePortalOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeletePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DeleteProject(ctx context.Context, input *iotsitewise.DeleteProjectInput) (*iotsitewise.DeleteProjectOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAccessPolicy(ctx context.Context, input *iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error) {
    return a.client.DescribeAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAsset(ctx context.Context, input *iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error) {
    return a.client.DescribeAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAssetModel(ctx context.Context, input *iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error) {
    return a.client.DescribeAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeAssetProperty(ctx context.Context, input *iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error) {
    return a.client.DescribeAssetPropertyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeDashboard(ctx context.Context, input *iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error) {
    return a.client.DescribeDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeGateway(ctx context.Context, input *iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error) {
    return a.client.DescribeGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeGatewayCapabilityConfiguration(ctx context.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
    return a.client.DescribeGatewayCapabilityConfigurationWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeLoggingOptions(ctx context.Context, input *iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
    return a.client.DescribeLoggingOptionsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribePortal(ctx context.Context, input *iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error) {
    return a.client.DescribePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DescribeProject(ctx context.Context, input *iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error) {
    return a.client.DescribeProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) DisassociateAssets(ctx context.Context, input *iotsitewise.DisassociateAssetsInput) (*iotsitewise.DisassociateAssetsOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.DisassociateAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyAggregates(ctx context.Context, input *iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
    return a.client.GetAssetPropertyAggregatesWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyValue(ctx context.Context, input *iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error) {
    return a.client.GetAssetPropertyValueWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyValueHistory(ctx context.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
    return a.client.GetAssetPropertyValueHistoryWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAccessPolicies(ctx context.Context, input *iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error) {
    return a.client.ListAccessPoliciesWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAssetModels(ctx context.Context, input *iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error) {
    return a.client.ListAssetModelsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAssets(ctx context.Context, input *iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error) {
    return a.client.ListAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListAssociatedAssets(ctx context.Context, input *iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error) {
    return a.client.ListAssociatedAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListDashboards(ctx context.Context, input *iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error) {
    return a.client.ListDashboardsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListGateways(ctx context.Context, input *iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error) {
    return a.client.ListGatewaysWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListPortals(ctx context.Context, input *iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error) {
    return a.client.ListPortalsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListProjectAssets(ctx context.Context, input *iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error) {
    return a.client.ListProjectAssetsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListProjects(ctx context.Context, input *iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error) {
    return a.client.ListProjectsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) ListTagsForResource(ctx context.Context, input *iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) PutLoggingOptions(ctx context.Context, input *iotsitewise.PutLoggingOptionsInput) (*iotsitewise.PutLoggingOptionsOutput, error) {
    return a.client.PutLoggingOptionsWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) TagResource(ctx context.Context, input *iotsitewise.TagResourceInput) (*iotsitewise.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UntagResource(ctx context.Context, input *iotsitewise.UntagResourceInput) (*iotsitewise.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAccessPolicy(ctx context.Context, input *iotsitewise.UpdateAccessPolicyInput) (*iotsitewise.UpdateAccessPolicyOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateAccessPolicyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAsset(ctx context.Context, input *iotsitewise.UpdateAssetInput) (*iotsitewise.UpdateAssetOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateAssetWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAssetModel(ctx context.Context, input *iotsitewise.UpdateAssetModelInput) (*iotsitewise.UpdateAssetModelOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateAssetModelWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateAssetProperty(ctx context.Context, input *iotsitewise.UpdateAssetPropertyInput) (*iotsitewise.UpdateAssetPropertyOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateAssetPropertyWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateDashboard(ctx context.Context, input *iotsitewise.UpdateDashboardInput) (*iotsitewise.UpdateDashboardOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateDashboardWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateGateway(ctx context.Context, input *iotsitewise.UpdateGatewayInput) (*iotsitewise.UpdateGatewayOutput, error) {
    return a.client.UpdateGatewayWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateGatewayCapabilityConfiguration(ctx context.Context, input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error) {
    return a.client.UpdateGatewayCapabilityConfigurationWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdatePortal(ctx context.Context, input *iotsitewise.UpdatePortalInput) (*iotsitewise.UpdatePortalOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdatePortalWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) UpdateProject(ctx context.Context, input *iotsitewise.UpdateProjectInput) (*iotsitewise.UpdateProjectOutput, error) {
    internal.SetClientToken(ctx, &input.ClientToken)
    return a.client.UpdateProjectWithContext(ctx, input)
}

func (a *IoTSiteWiseActivities) WaitUntilAssetActive(ctx context.Context, input *iotsitewise.DescribeAssetInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilAssetActiveWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilAssetModelActive(ctx context.Context, input *iotsitewise.DescribeAssetModelInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilAssetModelActiveWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilAssetModelNotExists(ctx context.Context, input *iotsitewise.DescribeAssetModelInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilAssetModelNotExistsWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilAssetNotExists(ctx context.Context, input *iotsitewise.DescribeAssetInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilAssetNotExistsWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilPortalActive(ctx context.Context, input *iotsitewise.DescribePortalInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilPortalActiveWithContext(ctx, input, options...)
	})
}

func (a *IoTSiteWiseActivities) WaitUntilPortalNotExists(ctx context.Context, input *iotsitewise.DescribePortalInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilPortalNotExistsWithContext(ctx, input, options...)
	})
}
