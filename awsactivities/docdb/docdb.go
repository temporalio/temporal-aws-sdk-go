// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package docdb

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/docdb/docdbiface"
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
	client docdbiface.DocDBAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := docdb.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (docdbiface.DocDBAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return docdb.New(sess), nil
}

func (a *Activities) AddTagsToResource(ctx context.Context, input *docdb.AddTagsToResourceInput) (*docdb.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *Activities) ApplyPendingMaintenanceAction(ctx context.Context, input *docdb.ApplyPendingMaintenanceActionInput) (*docdb.ApplyPendingMaintenanceActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *Activities) CopyDBClusterParameterGroup(ctx context.Context, input *docdb.CopyDBClusterParameterGroupInput) (*docdb.CopyDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *Activities) CopyDBClusterSnapshot(ctx context.Context, input *docdb.CopyDBClusterSnapshotInput) (*docdb.CopyDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CopyDBClusterSnapshotWithContext(ctx, input)
}

func (a *Activities) CreateDBCluster(ctx context.Context, input *docdb.CreateDBClusterInput) (*docdb.CreateDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBClusterWithContext(ctx, input)
}

func (a *Activities) CreateDBClusterParameterGroup(ctx context.Context, input *docdb.CreateDBClusterParameterGroupInput) (*docdb.CreateDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBClusterParameterGroupWithContext(ctx, input)
}

func (a *Activities) CreateDBClusterSnapshot(ctx context.Context, input *docdb.CreateDBClusterSnapshotInput) (*docdb.CreateDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBClusterSnapshotWithContext(ctx, input)
}

func (a *Activities) CreateDBInstance(ctx context.Context, input *docdb.CreateDBInstanceInput) (*docdb.CreateDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBInstanceWithContext(ctx, input)
}

func (a *Activities) CreateDBSubnetGroup(ctx context.Context, input *docdb.CreateDBSubnetGroupInput) (*docdb.CreateDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDBSubnetGroupWithContext(ctx, input)
}

func (a *Activities) DeleteDBCluster(ctx context.Context, input *docdb.DeleteDBClusterInput) (*docdb.DeleteDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBClusterWithContext(ctx, input)
}

func (a *Activities) DeleteDBClusterParameterGroup(ctx context.Context, input *docdb.DeleteDBClusterParameterGroupInput) (*docdb.DeleteDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBClusterParameterGroupWithContext(ctx, input)
}

func (a *Activities) DeleteDBClusterSnapshot(ctx context.Context, input *docdb.DeleteDBClusterSnapshotInput) (*docdb.DeleteDBClusterSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBClusterSnapshotWithContext(ctx, input)
}

func (a *Activities) DeleteDBInstance(ctx context.Context, input *docdb.DeleteDBInstanceInput) (*docdb.DeleteDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBInstanceWithContext(ctx, input)
}

func (a *Activities) DeleteDBSubnetGroup(ctx context.Context, input *docdb.DeleteDBSubnetGroupInput) (*docdb.DeleteDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDBSubnetGroupWithContext(ctx, input)
}

func (a *Activities) DescribeCertificates(ctx context.Context, input *docdb.DescribeCertificatesInput) (*docdb.DescribeCertificatesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCertificatesWithContext(ctx, input)
}

func (a *Activities) DescribeDBClusterParameterGroups(ctx context.Context, input *docdb.DescribeDBClusterParameterGroupsInput) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterParameterGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeDBClusterParameters(ctx context.Context, input *docdb.DescribeDBClusterParametersInput) (*docdb.DescribeDBClusterParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterParametersWithContext(ctx, input)
}

func (a *Activities) DescribeDBClusterSnapshotAttributes(ctx context.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterSnapshotAttributesWithContext(ctx, input)
}

func (a *Activities) DescribeDBClusterSnapshots(ctx context.Context, input *docdb.DescribeDBClusterSnapshotsInput) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClusterSnapshotsWithContext(ctx, input)
}

func (a *Activities) DescribeDBClusters(ctx context.Context, input *docdb.DescribeDBClustersInput) (*docdb.DescribeDBClustersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBClustersWithContext(ctx, input)
}

func (a *Activities) DescribeDBEngineVersions(ctx context.Context, input *docdb.DescribeDBEngineVersionsInput) (*docdb.DescribeDBEngineVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBEngineVersionsWithContext(ctx, input)
}

func (a *Activities) DescribeDBInstances(ctx context.Context, input *docdb.DescribeDBInstancesInput) (*docdb.DescribeDBInstancesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBInstancesWithContext(ctx, input)
}

func (a *Activities) DescribeDBSubnetGroups(ctx context.Context, input *docdb.DescribeDBSubnetGroupsInput) (*docdb.DescribeDBSubnetGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDBSubnetGroupsWithContext(ctx, input)
}

func (a *Activities) DescribeEngineDefaultClusterParameters(ctx context.Context, input *docdb.DescribeEngineDefaultClusterParametersInput) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEngineDefaultClusterParametersWithContext(ctx, input)
}

func (a *Activities) DescribeEventCategories(ctx context.Context, input *docdb.DescribeEventCategoriesInput) (*docdb.DescribeEventCategoriesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *docdb.DescribeEventsInput) (*docdb.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEventsWithContext(ctx, input)
}

func (a *Activities) DescribeOrderableDBInstanceOptions(ctx context.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)
}

func (a *Activities) DescribePendingMaintenanceActions(ctx context.Context, input *docdb.DescribePendingMaintenanceActionsInput) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *Activities) FailoverDBCluster(ctx context.Context, input *docdb.FailoverDBClusterInput) (*docdb.FailoverDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.FailoverDBClusterWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *docdb.ListTagsForResourceInput) (*docdb.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ModifyDBCluster(ctx context.Context, input *docdb.ModifyDBClusterInput) (*docdb.ModifyDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBClusterWithContext(ctx, input)
}

func (a *Activities) ModifyDBClusterParameterGroup(ctx context.Context, input *docdb.ModifyDBClusterParameterGroupInput) (*docdb.ModifyDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *Activities) ModifyDBClusterSnapshotAttribute(ctx context.Context, input *docdb.ModifyDBClusterSnapshotAttributeInput) (*docdb.ModifyDBClusterSnapshotAttributeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBClusterSnapshotAttributeWithContext(ctx, input)
}

func (a *Activities) ModifyDBInstance(ctx context.Context, input *docdb.ModifyDBInstanceInput) (*docdb.ModifyDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBInstanceWithContext(ctx, input)
}

func (a *Activities) ModifyDBSubnetGroup(ctx context.Context, input *docdb.ModifyDBSubnetGroupInput) (*docdb.ModifyDBSubnetGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ModifyDBSubnetGroupWithContext(ctx, input)
}

func (a *Activities) RebootDBInstance(ctx context.Context, input *docdb.RebootDBInstanceInput) (*docdb.RebootDBInstanceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RebootDBInstanceWithContext(ctx, input)
}

func (a *Activities) RemoveTagsFromResource(ctx context.Context, input *docdb.RemoveTagsFromResourceInput) (*docdb.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *Activities) ResetDBClusterParameterGroup(ctx context.Context, input *docdb.ResetDBClusterParameterGroupInput) (*docdb.ResetDBClusterParameterGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetDBClusterParameterGroupWithContext(ctx, input)
}

func (a *Activities) RestoreDBClusterFromSnapshot(ctx context.Context, input *docdb.RestoreDBClusterFromSnapshotInput) (*docdb.RestoreDBClusterFromSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreDBClusterFromSnapshotWithContext(ctx, input)
}

func (a *Activities) RestoreDBClusterToPointInTime(ctx context.Context, input *docdb.RestoreDBClusterToPointInTimeInput) (*docdb.RestoreDBClusterToPointInTimeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RestoreDBClusterToPointInTimeWithContext(ctx, input)
}

func (a *Activities) StartDBCluster(ctx context.Context, input *docdb.StartDBClusterInput) (*docdb.StartDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDBClusterWithContext(ctx, input)
}

func (a *Activities) StopDBCluster(ctx context.Context, input *docdb.StopDBClusterInput) (*docdb.StopDBClusterOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopDBClusterWithContext(ctx, input)
}

func (a *Activities) WaitUntilDBInstanceAvailable(ctx context.Context, input *docdb.DescribeDBInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilDBInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *Activities) WaitUntilDBInstanceDeleted(ctx context.Context, input *docdb.DescribeDBInstancesInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return err
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return client.WaitUntilDBInstanceDeletedWithContext(ctx, input, options...)
	})
}
