// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/aws/aws-sdk-go/service/lakeformation/lakeformationiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type LakeFormationActivities struct {
	client lakeformationiface.LakeFormationAPI

	sessionFactory SessionFactory
}

func NewLakeFormationActivities(sess *session.Session, config ...*aws.Config) *LakeFormationActivities {
	client := lakeformation.New(sess, config...)
	return &LakeFormationActivities{client: client}
}

func NewLakeFormationActivitiesWithSessionFactory(sessionFactory SessionFactory) *LakeFormationActivities {
	return &LakeFormationActivities{sessionFactory: sessionFactory}
}

func (a *LakeFormationActivities) getClient(ctx context.Context) (lakeformationiface.LakeFormationAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return lakeformation.New(sess), nil
}

func (a *LakeFormationActivities) BatchGrantPermissions(ctx context.Context, input *lakeformation.BatchGrantPermissionsInput) (*lakeformation.BatchGrantPermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchGrantPermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) BatchRevokePermissions(ctx context.Context, input *lakeformation.BatchRevokePermissionsInput) (*lakeformation.BatchRevokePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchRevokePermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) DeregisterResource(ctx context.Context, input *lakeformation.DeregisterResourceInput) (*lakeformation.DeregisterResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeregisterResourceWithContext(ctx, input)
}

func (a *LakeFormationActivities) DescribeResource(ctx context.Context, input *lakeformation.DescribeResourceInput) (*lakeformation.DescribeResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeResourceWithContext(ctx, input)
}

func (a *LakeFormationActivities) GetDataLakeSettings(ctx context.Context, input *lakeformation.GetDataLakeSettingsInput) (*lakeformation.GetDataLakeSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDataLakeSettingsWithContext(ctx, input)
}

func (a *LakeFormationActivities) GetEffectivePermissionsForPath(ctx context.Context, input *lakeformation.GetEffectivePermissionsForPathInput) (*lakeformation.GetEffectivePermissionsForPathOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEffectivePermissionsForPathWithContext(ctx, input)
}

func (a *LakeFormationActivities) GrantPermissions(ctx context.Context, input *lakeformation.GrantPermissionsInput) (*lakeformation.GrantPermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GrantPermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) ListPermissions(ctx context.Context, input *lakeformation.ListPermissionsInput) (*lakeformation.ListPermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) ListResources(ctx context.Context, input *lakeformation.ListResourcesInput) (*lakeformation.ListResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourcesWithContext(ctx, input)
}

func (a *LakeFormationActivities) PutDataLakeSettings(ctx context.Context, input *lakeformation.PutDataLakeSettingsInput) (*lakeformation.PutDataLakeSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDataLakeSettingsWithContext(ctx, input)
}

func (a *LakeFormationActivities) RegisterResource(ctx context.Context, input *lakeformation.RegisterResourceInput) (*lakeformation.RegisterResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RegisterResourceWithContext(ctx, input)
}

func (a *LakeFormationActivities) RevokePermissions(ctx context.Context, input *lakeformation.RevokePermissionsInput) (*lakeformation.RevokePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RevokePermissionsWithContext(ctx, input)
}

func (a *LakeFormationActivities) UpdateResource(ctx context.Context, input *lakeformation.UpdateResourceInput) (*lakeformation.UpdateResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResourceWithContext(ctx, input)
}
