
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/aws/aws-sdk-go/service/iotsitewise/iotsitewiseiface"
)

type IoTSiteWiseActivities struct {
    client iotsitewiseiface.IoTSiteWiseAPI
}

func NewIoTSiteWiseActivities(session *session.Session, config ...*aws.Config) *IoTSiteWiseActivities {
    client := iotsitewise.New(session, config...)
    return &IoTSiteWiseActivities{client: client}
}

func (a *IoTSiteWiseActivities) AssociateAssets(input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error) {
    return a.client.AssociateAssets(input)
}

func (a *IoTSiteWiseActivities) BatchAssociateProjectAssets(input *iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
    return a.client.BatchAssociateProjectAssets(input)
}

func (a *IoTSiteWiseActivities) BatchDisassociateProjectAssets(input *iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
    return a.client.BatchDisassociateProjectAssets(input)
}

func (a *IoTSiteWiseActivities) BatchPutAssetPropertyValue(input *iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
    return a.client.BatchPutAssetPropertyValue(input)
}

func (a *IoTSiteWiseActivities) CreateAccessPolicy(input *iotsitewise.CreateAccessPolicyInput) (*iotsitewise.CreateAccessPolicyOutput, error) {
    return a.client.CreateAccessPolicy(input)
}

func (a *IoTSiteWiseActivities) CreateAsset(input *iotsitewise.CreateAssetInput) (*iotsitewise.CreateAssetOutput, error) {
    return a.client.CreateAsset(input)
}

func (a *IoTSiteWiseActivities) CreateAssetModel(input *iotsitewise.CreateAssetModelInput) (*iotsitewise.CreateAssetModelOutput, error) {
    return a.client.CreateAssetModel(input)
}

func (a *IoTSiteWiseActivities) CreateDashboard(input *iotsitewise.CreateDashboardInput) (*iotsitewise.CreateDashboardOutput, error) {
    return a.client.CreateDashboard(input)
}

func (a *IoTSiteWiseActivities) CreateGateway(input *iotsitewise.CreateGatewayInput) (*iotsitewise.CreateGatewayOutput, error) {
    return a.client.CreateGateway(input)
}

func (a *IoTSiteWiseActivities) CreatePortal(input *iotsitewise.CreatePortalInput) (*iotsitewise.CreatePortalOutput, error) {
    return a.client.CreatePortal(input)
}

func (a *IoTSiteWiseActivities) CreateProject(input *iotsitewise.CreateProjectInput) (*iotsitewise.CreateProjectOutput, error) {
    return a.client.CreateProject(input)
}

func (a *IoTSiteWiseActivities) DeleteAccessPolicy(input *iotsitewise.DeleteAccessPolicyInput) (*iotsitewise.DeleteAccessPolicyOutput, error) {
    return a.client.DeleteAccessPolicy(input)
}

func (a *IoTSiteWiseActivities) DeleteAsset(input *iotsitewise.DeleteAssetInput) (*iotsitewise.DeleteAssetOutput, error) {
    return a.client.DeleteAsset(input)
}

func (a *IoTSiteWiseActivities) DeleteAssetModel(input *iotsitewise.DeleteAssetModelInput) (*iotsitewise.DeleteAssetModelOutput, error) {
    return a.client.DeleteAssetModel(input)
}

func (a *IoTSiteWiseActivities) DeleteDashboard(input *iotsitewise.DeleteDashboardInput) (*iotsitewise.DeleteDashboardOutput, error) {
    return a.client.DeleteDashboard(input)
}

func (a *IoTSiteWiseActivities) DeleteGateway(input *iotsitewise.DeleteGatewayInput) (*iotsitewise.DeleteGatewayOutput, error) {
    return a.client.DeleteGateway(input)
}

func (a *IoTSiteWiseActivities) DeletePortal(input *iotsitewise.DeletePortalInput) (*iotsitewise.DeletePortalOutput, error) {
    return a.client.DeletePortal(input)
}

func (a *IoTSiteWiseActivities) DeleteProject(input *iotsitewise.DeleteProjectInput) (*iotsitewise.DeleteProjectOutput, error) {
    return a.client.DeleteProject(input)
}

func (a *IoTSiteWiseActivities) DescribeAccessPolicy(input *iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error) {
    return a.client.DescribeAccessPolicy(input)
}

func (a *IoTSiteWiseActivities) DescribeAsset(input *iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error) {
    return a.client.DescribeAsset(input)
}

func (a *IoTSiteWiseActivities) DescribeAssetModel(input *iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error) {
    return a.client.DescribeAssetModel(input)
}

func (a *IoTSiteWiseActivities) DescribeAssetProperty(input *iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error) {
    return a.client.DescribeAssetProperty(input)
}

func (a *IoTSiteWiseActivities) DescribeDashboard(input *iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error) {
    return a.client.DescribeDashboard(input)
}

func (a *IoTSiteWiseActivities) DescribeGateway(input *iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error) {
    return a.client.DescribeGateway(input)
}

func (a *IoTSiteWiseActivities) DescribeGatewayCapabilityConfiguration(input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
    return a.client.DescribeGatewayCapabilityConfiguration(input)
}

func (a *IoTSiteWiseActivities) DescribeLoggingOptions(input *iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
    return a.client.DescribeLoggingOptions(input)
}

func (a *IoTSiteWiseActivities) DescribePortal(input *iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error) {
    return a.client.DescribePortal(input)
}

func (a *IoTSiteWiseActivities) DescribeProject(input *iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error) {
    return a.client.DescribeProject(input)
}

func (a *IoTSiteWiseActivities) DisassociateAssets(input *iotsitewise.DisassociateAssetsInput) (*iotsitewise.DisassociateAssetsOutput, error) {
    return a.client.DisassociateAssets(input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyAggregates(input *iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
    return a.client.GetAssetPropertyAggregates(input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyValue(input *iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error) {
    return a.client.GetAssetPropertyValue(input)
}

func (a *IoTSiteWiseActivities) GetAssetPropertyValueHistory(input *iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
    return a.client.GetAssetPropertyValueHistory(input)
}

func (a *IoTSiteWiseActivities) ListAccessPolicies(input *iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error) {
    return a.client.ListAccessPolicies(input)
}

func (a *IoTSiteWiseActivities) ListAssetModels(input *iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error) {
    return a.client.ListAssetModels(input)
}

func (a *IoTSiteWiseActivities) ListAssets(input *iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error) {
    return a.client.ListAssets(input)
}

func (a *IoTSiteWiseActivities) ListAssociatedAssets(input *iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error) {
    return a.client.ListAssociatedAssets(input)
}

func (a *IoTSiteWiseActivities) ListDashboards(input *iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error) {
    return a.client.ListDashboards(input)
}

func (a *IoTSiteWiseActivities) ListGateways(input *iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error) {
    return a.client.ListGateways(input)
}

func (a *IoTSiteWiseActivities) ListPortals(input *iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error) {
    return a.client.ListPortals(input)
}

func (a *IoTSiteWiseActivities) ListProjectAssets(input *iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error) {
    return a.client.ListProjectAssets(input)
}

func (a *IoTSiteWiseActivities) ListProjects(input *iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error) {
    return a.client.ListProjects(input)
}

func (a *IoTSiteWiseActivities) ListTagsForResource(input *iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *IoTSiteWiseActivities) PutLoggingOptions(input *iotsitewise.PutLoggingOptionsInput) (*iotsitewise.PutLoggingOptionsOutput, error) {
    return a.client.PutLoggingOptions(input)
}

func (a *IoTSiteWiseActivities) TagResource(input *iotsitewise.TagResourceInput) (*iotsitewise.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *IoTSiteWiseActivities) UntagResource(input *iotsitewise.UntagResourceInput) (*iotsitewise.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *IoTSiteWiseActivities) UpdateAccessPolicy(input *iotsitewise.UpdateAccessPolicyInput) (*iotsitewise.UpdateAccessPolicyOutput, error) {
    return a.client.UpdateAccessPolicy(input)
}

func (a *IoTSiteWiseActivities) UpdateAsset(input *iotsitewise.UpdateAssetInput) (*iotsitewise.UpdateAssetOutput, error) {
    return a.client.UpdateAsset(input)
}

func (a *IoTSiteWiseActivities) UpdateAssetModel(input *iotsitewise.UpdateAssetModelInput) (*iotsitewise.UpdateAssetModelOutput, error) {
    return a.client.UpdateAssetModel(input)
}

func (a *IoTSiteWiseActivities) UpdateAssetProperty(input *iotsitewise.UpdateAssetPropertyInput) (*iotsitewise.UpdateAssetPropertyOutput, error) {
    return a.client.UpdateAssetProperty(input)
}

func (a *IoTSiteWiseActivities) UpdateDashboard(input *iotsitewise.UpdateDashboardInput) (*iotsitewise.UpdateDashboardOutput, error) {
    return a.client.UpdateDashboard(input)
}

func (a *IoTSiteWiseActivities) UpdateGateway(input *iotsitewise.UpdateGatewayInput) (*iotsitewise.UpdateGatewayOutput, error) {
    return a.client.UpdateGateway(input)
}

func (a *IoTSiteWiseActivities) UpdateGatewayCapabilityConfiguration(input *iotsitewise.UpdateGatewayCapabilityConfigurationInput) (*iotsitewise.UpdateGatewayCapabilityConfigurationOutput, error) {
    return a.client.UpdateGatewayCapabilityConfiguration(input)
}

func (a *IoTSiteWiseActivities) UpdatePortal(input *iotsitewise.UpdatePortalInput) (*iotsitewise.UpdatePortalOutput, error) {
    return a.client.UpdatePortal(input)
}

func (a *IoTSiteWiseActivities) UpdateProject(input *iotsitewise.UpdateProjectInput) (*iotsitewise.UpdateProjectOutput, error) {
    return a.client.UpdateProject(input)
}

func (a *IoTSiteWiseActivities) WaitUntilAssetActive(input *iotsitewise.DescribeAssetInput) error {
    return a.client.WaitUntilAssetActive(input)
}

func (a *IoTSiteWiseActivities) WaitUntilAssetModelActive(input *iotsitewise.DescribeAssetModelInput) error {
    return a.client.WaitUntilAssetModelActive(input)
}

func (a *IoTSiteWiseActivities) WaitUntilAssetModelNotExists(input *iotsitewise.DescribeAssetModelInput) error {
    return a.client.WaitUntilAssetModelNotExists(input)
}

func (a *IoTSiteWiseActivities) WaitUntilAssetNotExists(input *iotsitewise.DescribeAssetInput) error {
    return a.client.WaitUntilAssetNotExists(input)
}

func (a *IoTSiteWiseActivities) WaitUntilPortalActive(input *iotsitewise.DescribePortalInput) error {
    return a.client.WaitUntilPortalActive(input)
}

func (a *IoTSiteWiseActivities) WaitUntilPortalNotExists(input *iotsitewise.DescribePortalInput) error {
    return a.client.WaitUntilPortalNotExists(input)
}
