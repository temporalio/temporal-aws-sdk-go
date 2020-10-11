// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediastore

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client mediastoreiface.MediaStoreAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := mediastore.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (mediastoreiface.MediaStoreAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return mediastore.New(sess), nil
}

func (a *Activities) CreateContainer(ctx context.Context, input *mediastore.CreateContainerInput) (*mediastore.CreateContainerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateContainerWithContext(ctx, input)
}

func (a *Activities) DeleteContainer(ctx context.Context, input *mediastore.DeleteContainerInput) (*mediastore.DeleteContainerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteContainerWithContext(ctx, input)
}

func (a *Activities) DeleteContainerPolicy(ctx context.Context, input *mediastore.DeleteContainerPolicyInput) (*mediastore.DeleteContainerPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteContainerPolicyWithContext(ctx, input)
}

func (a *Activities) DeleteCorsPolicy(ctx context.Context, input *mediastore.DeleteCorsPolicyInput) (*mediastore.DeleteCorsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteCorsPolicyWithContext(ctx, input)
}

func (a *Activities) DeleteLifecyclePolicy(ctx context.Context, input *mediastore.DeleteLifecyclePolicyInput) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLifecyclePolicyWithContext(ctx, input)
}

func (a *Activities) DeleteMetricPolicy(ctx context.Context, input *mediastore.DeleteMetricPolicyInput) (*mediastore.DeleteMetricPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteMetricPolicyWithContext(ctx, input)
}

func (a *Activities) DescribeContainer(ctx context.Context, input *mediastore.DescribeContainerInput) (*mediastore.DescribeContainerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeContainerWithContext(ctx, input)
}

func (a *Activities) GetContainerPolicy(ctx context.Context, input *mediastore.GetContainerPolicyInput) (*mediastore.GetContainerPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetContainerPolicyWithContext(ctx, input)
}

func (a *Activities) GetCorsPolicy(ctx context.Context, input *mediastore.GetCorsPolicyInput) (*mediastore.GetCorsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetCorsPolicyWithContext(ctx, input)
}

func (a *Activities) GetLifecyclePolicy(ctx context.Context, input *mediastore.GetLifecyclePolicyInput) (*mediastore.GetLifecyclePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLifecyclePolicyWithContext(ctx, input)
}

func (a *Activities) GetMetricPolicy(ctx context.Context, input *mediastore.GetMetricPolicyInput) (*mediastore.GetMetricPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetMetricPolicyWithContext(ctx, input)
}

func (a *Activities) ListContainers(ctx context.Context, input *mediastore.ListContainersInput) (*mediastore.ListContainersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListContainersWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *mediastore.ListTagsForResourceInput) (*mediastore.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) PutContainerPolicy(ctx context.Context, input *mediastore.PutContainerPolicyInput) (*mediastore.PutContainerPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutContainerPolicyWithContext(ctx, input)
}

func (a *Activities) PutCorsPolicy(ctx context.Context, input *mediastore.PutCorsPolicyInput) (*mediastore.PutCorsPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutCorsPolicyWithContext(ctx, input)
}

func (a *Activities) PutLifecyclePolicy(ctx context.Context, input *mediastore.PutLifecyclePolicyInput) (*mediastore.PutLifecyclePolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutLifecyclePolicyWithContext(ctx, input)
}

func (a *Activities) PutMetricPolicy(ctx context.Context, input *mediastore.PutMetricPolicyInput) (*mediastore.PutMetricPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutMetricPolicyWithContext(ctx, input)
}

func (a *Activities) StartAccessLogging(ctx context.Context, input *mediastore.StartAccessLoggingInput) (*mediastore.StartAccessLoggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartAccessLoggingWithContext(ctx, input)
}

func (a *Activities) StopAccessLogging(ctx context.Context, input *mediastore.StopAccessLoggingInput) (*mediastore.StopAccessLoggingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopAccessLoggingWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *mediastore.TagResourceInput) (*mediastore.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *mediastore.UntagResourceInput) (*mediastore.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}
