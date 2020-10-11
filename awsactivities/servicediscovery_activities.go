// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ServiceDiscoveryActivities struct {
	client servicediscoveryiface.ServiceDiscoveryAPI

	sessionFactory SessionFactory
}

func NewServiceDiscoveryActivities(sess *session.Session, config ...*aws.Config) *ServiceDiscoveryActivities {
	client := servicediscovery.New(sess, config...)
	return &ServiceDiscoveryActivities{client: client}
}

func NewServiceDiscoveryActivitiesWithSessionFactory(sessionFactory SessionFactory) *ServiceDiscoveryActivities {
	return &ServiceDiscoveryActivities{sessionFactory: sessionFactory}
}

func (a *ServiceDiscoveryActivities) getClient(ctx context.Context) (servicediscoveryiface.ServiceDiscoveryAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return servicediscovery.New(sess), nil
}

func (a *ServiceDiscoveryActivities) CreateHttpNamespace(ctx context.Context, input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateHttpNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) CreatePrivateDnsNamespace(ctx context.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePrivateDnsNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) CreatePublicDnsNamespace(ctx context.Context, input *servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePublicDnsNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) CreateService(ctx context.Context, input *servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateServiceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DeleteNamespace(ctx context.Context, input *servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DeleteService(ctx context.Context, input *servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteServiceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DeregisterInstance(ctx context.Context, input *servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterInstanceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) DiscoverInstances(ctx context.Context, input *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DiscoverInstancesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetInstance(ctx context.Context, input *servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstanceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetInstancesHealthStatus(ctx context.Context, input *servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetInstancesHealthStatusWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetNamespace(ctx context.Context, input *servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetNamespaceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetOperation(ctx context.Context, input *servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetOperationWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) GetService(ctx context.Context, input *servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetServiceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListInstances(ctx context.Context, input *servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListInstancesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListNamespaces(ctx context.Context, input *servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListNamespacesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListOperations(ctx context.Context, input *servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListOperationsWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListServices(ctx context.Context, input *servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServicesWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) ListTagsForResource(ctx context.Context, input *servicediscovery.ListTagsForResourceInput) (*servicediscovery.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) RegisterInstance(ctx context.Context, input *servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterInstanceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) TagResource(ctx context.Context, input *servicediscovery.TagResourceInput) (*servicediscovery.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) UntagResource(ctx context.Context, input *servicediscovery.UntagResourceInput) (*servicediscovery.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) UpdateInstanceCustomHealthStatus(ctx context.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateInstanceCustomHealthStatusWithContext(ctx, input)
}

func (a *ServiceDiscoveryActivities) UpdateService(ctx context.Context, input *servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateServiceWithContext(ctx, input)
}
