
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DynamoDBActivities struct {
    client dynamodbiface.DynamoDBAPI
}

func NewDynamoDBActivities(session *session.Session, config ...*aws.Config) *DynamoDBActivities {
    client := dynamodb.New(session, config...)
    return &DynamoDBActivities{client: client}
}

func (a *DynamoDBActivities) BatchGetItem(ctx context.Context, input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
    return a.client.BatchGetItemWithContext(ctx, input)
}

func (a *DynamoDBActivities) BatchWriteItem(ctx context.Context, input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
    return a.client.BatchWriteItemWithContext(ctx, input)
}

func (a *DynamoDBActivities) CreateBackup(ctx context.Context, input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
    return a.client.CreateBackupWithContext(ctx, input)
}

func (a *DynamoDBActivities) CreateGlobalTable(ctx context.Context, input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
    return a.client.CreateGlobalTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) CreateTable(ctx context.Context, input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
    return a.client.CreateTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) DeleteBackup(ctx context.Context, input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
    return a.client.DeleteBackupWithContext(ctx, input)
}

func (a *DynamoDBActivities) DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
    return a.client.DeleteItemWithContext(ctx, input)
}

func (a *DynamoDBActivities) DeleteTable(ctx context.Context, input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
    return a.client.DeleteTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeBackup(ctx context.Context, input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
    return a.client.DescribeBackupWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeContinuousBackups(ctx context.Context, input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
    return a.client.DescribeContinuousBackupsWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeContributorInsights(ctx context.Context, input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
    return a.client.DescribeContributorInsightsWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeEndpoints(ctx context.Context, input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
    return a.client.DescribeEndpointsWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeGlobalTable(ctx context.Context, input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
    return a.client.DescribeGlobalTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeGlobalTableSettings(ctx context.Context, input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
    return a.client.DescribeGlobalTableSettingsWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeLimits(ctx context.Context, input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
    return a.client.DescribeLimitsWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeTable(ctx context.Context, input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
    return a.client.DescribeTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeTableReplicaAutoScaling(ctx context.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
    return a.client.DescribeTableReplicaAutoScalingWithContext(ctx, input)
}

func (a *DynamoDBActivities) DescribeTimeToLive(ctx context.Context, input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
    return a.client.DescribeTimeToLiveWithContext(ctx, input)
}

func (a *DynamoDBActivities) GetItem(ctx context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
    return a.client.GetItemWithContext(ctx, input)
}

func (a *DynamoDBActivities) ListBackups(ctx context.Context, input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
    return a.client.ListBackupsWithContext(ctx, input)
}

func (a *DynamoDBActivities) ListContributorInsights(ctx context.Context, input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
    return a.client.ListContributorInsightsWithContext(ctx, input)
}

func (a *DynamoDBActivities) ListGlobalTables(ctx context.Context, input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
    return a.client.ListGlobalTablesWithContext(ctx, input)
}

func (a *DynamoDBActivities) ListTables(ctx context.Context, input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
    return a.client.ListTablesWithContext(ctx, input)
}

func (a *DynamoDBActivities) ListTagsOfResource(ctx context.Context, input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
    return a.client.ListTagsOfResourceWithContext(ctx, input)
}

func (a *DynamoDBActivities) PutItem(ctx context.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
    return a.client.PutItemWithContext(ctx, input)
}

func (a *DynamoDBActivities) Query(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
    return a.client.QueryWithContext(ctx, input)
}

func (a *DynamoDBActivities) RestoreTableFromBackup(ctx context.Context, input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
    return a.client.RestoreTableFromBackupWithContext(ctx, input)
}

func (a *DynamoDBActivities) RestoreTableToPointInTime(ctx context.Context, input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
    return a.client.RestoreTableToPointInTimeWithContext(ctx, input)
}

func (a *DynamoDBActivities) Scan(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
    return a.client.ScanWithContext(ctx, input)
}

func (a *DynamoDBActivities) TagResource(ctx context.Context, input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *DynamoDBActivities) TransactGetItems(ctx context.Context, input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
    return a.client.TransactGetItemsWithContext(ctx, input)
}

func (a *DynamoDBActivities) TransactWriteItems(ctx context.Context, input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
    return a.client.TransactWriteItemsWithContext(ctx, input)
}

func (a *DynamoDBActivities) UntagResource(ctx context.Context, input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateContinuousBackups(ctx context.Context, input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
    return a.client.UpdateContinuousBackupsWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateContributorInsights(ctx context.Context, input *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
    return a.client.UpdateContributorInsightsWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateGlobalTable(ctx context.Context, input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
    return a.client.UpdateGlobalTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateGlobalTableSettings(ctx context.Context, input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
    return a.client.UpdateGlobalTableSettingsWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
    return a.client.UpdateItemWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateTable(ctx context.Context, input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
    return a.client.UpdateTableWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateTableReplicaAutoScaling(ctx context.Context, input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
    return a.client.UpdateTableReplicaAutoScalingWithContext(ctx, input)
}

func (a *DynamoDBActivities) UpdateTimeToLive(ctx context.Context, input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
    return a.client.UpdateTimeToLiveWithContext(ctx, input)
}

func (a *DynamoDBActivities) WaitUntilTableExists(ctx context.Context, input *dynamodb.DescribeTableInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilTableExistsWithContext(ctx, input, options...)
	})
}

func (a *DynamoDBActivities) WaitUntilTableNotExists(ctx context.Context, input *dynamodb.DescribeTableInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilTableNotExistsWithContext(ctx, input, options...)
	})
}
