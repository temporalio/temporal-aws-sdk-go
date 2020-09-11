package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CognitoSyncActivities struct {
	client cognitosynciface.CognitoSyncAPI
}

func NewCognitoSyncActivities(session *session.Session, config ...*aws.Config) *CognitoSyncActivities {
	client := cognitosync.New(session, config...)
	return &CognitoSyncActivities{client: client}
}

func (a *CognitoSyncActivities) BulkPublish(ctx context.Context, input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
	return a.client.BulkPublishWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DeleteDataset(ctx context.Context, input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error) {
	return a.client.DeleteDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DescribeDataset(ctx context.Context, input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error) {
	return a.client.DescribeDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DescribeIdentityPoolUsage(ctx context.Context, input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	return a.client.DescribeIdentityPoolUsageWithContext(ctx, input)
}

func (a *CognitoSyncActivities) DescribeIdentityUsage(ctx context.Context, input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error) {
	return a.client.DescribeIdentityUsageWithContext(ctx, input)
}

func (a *CognitoSyncActivities) GetBulkPublishDetails(ctx context.Context, input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	return a.client.GetBulkPublishDetailsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) GetCognitoEvents(ctx context.Context, input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error) {
	return a.client.GetCognitoEventsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) GetIdentityPoolConfiguration(ctx context.Context, input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	return a.client.GetIdentityPoolConfigurationWithContext(ctx, input)
}

func (a *CognitoSyncActivities) ListDatasets(ctx context.Context, input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error) {
	return a.client.ListDatasetsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) ListIdentityPoolUsage(ctx context.Context, input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	return a.client.ListIdentityPoolUsageWithContext(ctx, input)
}

func (a *CognitoSyncActivities) ListRecords(ctx context.Context, input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error) {
	return a.client.ListRecordsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) RegisterDevice(ctx context.Context, input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error) {
	return a.client.RegisterDeviceWithContext(ctx, input)
}

func (a *CognitoSyncActivities) SetCognitoEvents(ctx context.Context, input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error) {
	return a.client.SetCognitoEventsWithContext(ctx, input)
}

func (a *CognitoSyncActivities) SetIdentityPoolConfiguration(ctx context.Context, input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
	return a.client.SetIdentityPoolConfigurationWithContext(ctx, input)
}

func (a *CognitoSyncActivities) SubscribeToDataset(ctx context.Context, input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error) {
	return a.client.SubscribeToDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) UnsubscribeFromDataset(ctx context.Context, input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
	return a.client.UnsubscribeFromDatasetWithContext(ctx, input)
}

func (a *CognitoSyncActivities) UpdateRecords(ctx context.Context, input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error) {
	return a.client.UpdateRecordsWithContext(ctx, input)
}
