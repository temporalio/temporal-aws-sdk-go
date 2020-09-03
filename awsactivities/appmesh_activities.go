package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
)

type AppMeshActivities struct {
	client appmeshiface.AppMeshAPI
}

func NewAppMeshActivities(client appmeshiface.AppMeshAPI) *AppMeshActivities {
    return &AppMeshActivities{client: client}
}

func (a *AppMeshActivities) CreateGatewayRoute(input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
    return a.client.CreateGatewayRoute(input)
}

func (a *AppMeshActivities) CreateMesh(input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
    return a.client.CreateMesh(input)
}

func (a *AppMeshActivities) CreateRoute(input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error) {
    return a.client.CreateRoute(input)
}

func (a *AppMeshActivities) CreateVirtualGateway(input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error) {
    return a.client.CreateVirtualGateway(input)
}

func (a *AppMeshActivities) CreateVirtualNode(input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error) {
    return a.client.CreateVirtualNode(input)
}

func (a *AppMeshActivities) CreateVirtualRouter(input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error) {
    return a.client.CreateVirtualRouter(input)
}

func (a *AppMeshActivities) CreateVirtualService(input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error) {
    return a.client.CreateVirtualService(input)
}

func (a *AppMeshActivities) DeleteGatewayRoute(input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error) {
    return a.client.DeleteGatewayRoute(input)
}

func (a *AppMeshActivities) DeleteMesh(input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error) {
    return a.client.DeleteMesh(input)
}

func (a *AppMeshActivities) DeleteRoute(input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error) {
    return a.client.DeleteRoute(input)
}

func (a *AppMeshActivities) DeleteVirtualGateway(input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error) {
    return a.client.DeleteVirtualGateway(input)
}

func (a *AppMeshActivities) DeleteVirtualNode(input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error) {
    return a.client.DeleteVirtualNode(input)
}

func (a *AppMeshActivities) DeleteVirtualRouter(input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error) {
    return a.client.DeleteVirtualRouter(input)
}

func (a *AppMeshActivities) DeleteVirtualService(input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error) {
    return a.client.DeleteVirtualService(input)
}

func (a *AppMeshActivities) DescribeGatewayRoute(input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
    return a.client.DescribeGatewayRoute(input)
}

func (a *AppMeshActivities) DescribeMesh(input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
    return a.client.DescribeMesh(input)
}

func (a *AppMeshActivities) DescribeRoute(input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
    return a.client.DescribeRoute(input)
}

func (a *AppMeshActivities) DescribeVirtualGateway(input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
    return a.client.DescribeVirtualGateway(input)
}

func (a *AppMeshActivities) DescribeVirtualNode(input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
    return a.client.DescribeVirtualNode(input)
}

func (a *AppMeshActivities) DescribeVirtualRouter(input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
    return a.client.DescribeVirtualRouter(input)
}

func (a *AppMeshActivities) DescribeVirtualService(input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
    return a.client.DescribeVirtualService(input)
}

func (a *AppMeshActivities) ListGatewayRoutes(input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
    return a.client.ListGatewayRoutes(input)
}

func (a *AppMeshActivities) ListMeshes(input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
    return a.client.ListMeshes(input)
}

func (a *AppMeshActivities) ListRoutes(input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
    return a.client.ListRoutes(input)
}

func (a *AppMeshActivities) ListTagsForResource(input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *AppMeshActivities) ListVirtualGateways(input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
    return a.client.ListVirtualGateways(input)
}

func (a *AppMeshActivities) ListVirtualNodes(input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
    return a.client.ListVirtualNodes(input)
}

func (a *AppMeshActivities) ListVirtualRouters(input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
    return a.client.ListVirtualRouters(input)
}

func (a *AppMeshActivities) ListVirtualServices(input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
    return a.client.ListVirtualServices(input)
}

func (a *AppMeshActivities) TagResource(input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *AppMeshActivities) UntagResource(input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *AppMeshActivities) UpdateGatewayRoute(input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error) {
    return a.client.UpdateGatewayRoute(input)
}

func (a *AppMeshActivities) UpdateMesh(input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error) {
    return a.client.UpdateMesh(input)
}

func (a *AppMeshActivities) UpdateRoute(input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error) {
    return a.client.UpdateRoute(input)
}

func (a *AppMeshActivities) UpdateVirtualGateway(input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error) {
    return a.client.UpdateVirtualGateway(input)
}

func (a *AppMeshActivities) UpdateVirtualNode(input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error) {
    return a.client.UpdateVirtualNode(input)
}

func (a *AppMeshActivities) UpdateVirtualRouter(input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error) {
    return a.client.UpdateVirtualRouter(input)
}

func (a *AppMeshActivities) UpdateVirtualService(input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error) {
    return a.client.UpdateVirtualService(input)
}
