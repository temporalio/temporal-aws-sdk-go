// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MigrationHubConfigActivities struct {
	client migrationhubconfigiface.MigrationHubConfigAPI

	sessionFactory SessionFactory
}

func NewMigrationHubConfigActivities(sess *session.Session, config ...*aws.Config) *MigrationHubConfigActivities {
	client := migrationhubconfig.New(sess, config...)
	return &MigrationHubConfigActivities{client: client}
}

func NewMigrationHubConfigActivitiesWithSessionFactory(sessionFactory SessionFactory) *MigrationHubConfigActivities {
	return &MigrationHubConfigActivities{sessionFactory: sessionFactory}
}

func (a *MigrationHubConfigActivities) getClient(ctx context.Context) (migrationhubconfigiface.MigrationHubConfigAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return migrationhubconfig.New(sess), nil
}

func (a *MigrationHubConfigActivities) CreateHomeRegionControl(ctx context.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateHomeRegionControlWithContext(ctx, input)
}

func (a *MigrationHubConfigActivities) DescribeHomeRegionControls(ctx context.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeHomeRegionControlsWithContext(ctx, input)
}

func (a *MigrationHubConfigActivities) GetHomeRegion(ctx context.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetHomeRegionWithContext(ctx, input)
}
