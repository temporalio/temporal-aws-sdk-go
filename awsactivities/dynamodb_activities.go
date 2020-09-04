
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type DynamoDBActivities struct {
	client dynamodbiface.DynamoDBAPI
}

func NewDynamoDBActivities(client dynamodbiface.DynamoDBAPI) *DynamoDBActivities {
    return &DynamoDBActivities{client: client}
}

func (a *DynamoDBActivities) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
    return a.client.BatchGetItem(input)
}

func (a *DynamoDBActivities) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
    return a.client.BatchWriteItem(input)
}

func (a *DynamoDBActivities) CreateBackup(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
    return a.client.CreateBackup(input)
}

func (a *DynamoDBActivities) CreateGlobalTable(input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
    return a.client.CreateGlobalTable(input)
}

func (a *DynamoDBActivities) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
    return a.client.CreateTable(input)
}

func (a *DynamoDBActivities) DeleteBackup(input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
    return a.client.DeleteBackup(input)
}

func (a *DynamoDBActivities) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
    return a.client.DeleteItem(input)
}

func (a *DynamoDBActivities) DeleteTable(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
    return a.client.DeleteTable(input)
}

func (a *DynamoDBActivities) DescribeBackup(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
    return a.client.DescribeBackup(input)
}

func (a *DynamoDBActivities) DescribeContinuousBackups(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
    return a.client.DescribeContinuousBackups(input)
}

func (a *DynamoDBActivities) DescribeContributorInsights(input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
    return a.client.DescribeContributorInsights(input)
}

func (a *DynamoDBActivities) DescribeEndpoints(input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
    return a.client.DescribeEndpoints(input)
}

func (a *DynamoDBActivities) DescribeGlobalTable(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
    return a.client.DescribeGlobalTable(input)
}

func (a *DynamoDBActivities) DescribeGlobalTableSettings(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
    return a.client.DescribeGlobalTableSettings(input)
}

func (a *DynamoDBActivities) DescribeLimits(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
    return a.client.DescribeLimits(input)
}

func (a *DynamoDBActivities) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
    return a.client.DescribeTable(input)
}

func (a *DynamoDBActivities) DescribeTableReplicaAutoScaling(input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
    return a.client.DescribeTableReplicaAutoScaling(input)
}

func (a *DynamoDBActivities) DescribeTimeToLive(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
    return a.client.DescribeTimeToLive(input)
}

func (a *DynamoDBActivities) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
    return a.client.GetItem(input)
}

func (a *DynamoDBActivities) ListBackups(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
    return a.client.ListBackups(input)
}

func (a *DynamoDBActivities) ListContributorInsights(input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
    return a.client.ListContributorInsights(input)
}

func (a *DynamoDBActivities) ListGlobalTables(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
    return a.client.ListGlobalTables(input)
}

func (a *DynamoDBActivities) ListTables(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
    return a.client.ListTables(input)
}

func (a *DynamoDBActivities) ListTagsOfResource(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
    return a.client.ListTagsOfResource(input)
}

func (a *DynamoDBActivities) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
    return a.client.PutItem(input)
}

func (a *DynamoDBActivities) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
    return a.client.Query(input)
}

func (a *DynamoDBActivities) RestoreTableFromBackup(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
    return a.client.RestoreTableFromBackup(input)
}

func (a *DynamoDBActivities) RestoreTableToPointInTime(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
    return a.client.RestoreTableToPointInTime(input)
}

func (a *DynamoDBActivities) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
    return a.client.Scan(input)
}

func (a *DynamoDBActivities) TagResource(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *DynamoDBActivities) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
    return a.client.TransactGetItems(input)
}

func (a *DynamoDBActivities) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
    return a.client.TransactWriteItems(input)
}

func (a *DynamoDBActivities) UntagResource(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *DynamoDBActivities) UpdateContinuousBackups(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
    return a.client.UpdateContinuousBackups(input)
}

func (a *DynamoDBActivities) UpdateContributorInsights(input *dynamodb.UpdateContributorInsightsInput) (*dynamodb.UpdateContributorInsightsOutput, error) {
    return a.client.UpdateContributorInsights(input)
}

func (a *DynamoDBActivities) UpdateGlobalTable(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
    return a.client.UpdateGlobalTable(input)
}

func (a *DynamoDBActivities) UpdateGlobalTableSettings(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
    return a.client.UpdateGlobalTableSettings(input)
}

func (a *DynamoDBActivities) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
    return a.client.UpdateItem(input)
}

func (a *DynamoDBActivities) UpdateTable(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
    return a.client.UpdateTable(input)
}

func (a *DynamoDBActivities) UpdateTableReplicaAutoScaling(input *dynamodb.UpdateTableReplicaAutoScalingInput) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
    return a.client.UpdateTableReplicaAutoScaling(input)
}

func (a *DynamoDBActivities) UpdateTimeToLive(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
    return a.client.UpdateTimeToLive(input)
}

func (a *DynamoDBActivities) WaitUntilTableExists(input *dynamodb.DescribeTableInput) error {
    return a.client.WaitUntilTableExists(input)
}

func (a *DynamoDBActivities) WaitUntilTableNotExists(input *dynamodb.DescribeTableInput) error {
    return a.client.WaitUntilTableNotExists(input)
}
