// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AppMeshActivities struct {
	client appmeshiface.AppMeshAPI

	sessionFactory SessionFactory
}

func NewAppMeshActivities(sess *session.Session, config ...*aws.Config) *AppMeshActivities {
	client := appmesh.New(sess, config...)
	return &AppMeshActivities{client: client}
}

func NewAppMeshActivitiesWithSessionFactory(sessionFactory SessionFactory) *AppMeshActivities {
	return &AppMeshActivities{sessionFactory: sessionFactory}
}

func (a *AppMeshActivities) getClient(ctx context.Context) (appmeshiface.AppMeshAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return appmesh.New(sess), nil
}

func (a *AppMeshActivities) CreateGatewayRoute(ctx context.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateMesh(ctx context.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateRoute(ctx context.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualGateway(ctx context.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualNode(ctx context.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualRouter(ctx context.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualService(ctx context.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateVirtualServiceWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteGatewayRoute(ctx context.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteMesh(ctx context.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteRoute(ctx context.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualGateway(ctx context.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualNode(ctx context.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualRouter(ctx context.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualService(ctx context.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVirtualServiceWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeGatewayRoute(ctx context.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeMesh(ctx context.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeRoute(ctx context.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualGateway(ctx context.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualNode(ctx context.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualRouter(ctx context.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualService(ctx context.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVirtualServiceWithContext(ctx, input)
}

func (a *AppMeshActivities) ListGatewayRoutes(ctx context.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGatewayRoutesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListMeshes(ctx context.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListMeshesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListRoutes(ctx context.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRoutesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListTagsForResource(ctx context.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualGateways(ctx context.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVirtualGatewaysWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualNodes(ctx context.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVirtualNodesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualRouters(ctx context.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVirtualRoutersWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualServices(ctx context.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVirtualServicesWithContext(ctx, input)
}

func (a *AppMeshActivities) TagResource(ctx context.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *AppMeshActivities) UntagResource(ctx context.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateGatewayRoute(ctx context.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateMesh(ctx context.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateRoute(ctx context.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualGateway(ctx context.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualNode(ctx context.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualRouter(ctx context.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualService(ctx context.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.UpdateVirtualServiceWithContext(ctx, input)
}
