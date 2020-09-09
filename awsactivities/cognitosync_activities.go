package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface"
)

type CognitoSyncActivities struct {
	client cognitosynciface.CognitoSyncAPI
}

func NewCognitoSyncActivities(session *session.Session, config ...*aws.Config) *CognitoSyncActivities {
	client := cognitosync.New(session, config...)
	return &CognitoSyncActivities{client: client}
}

func (a *CognitoSyncActivities) BulkPublish(input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
	return a.client.BulkPublish(input)
}

func (a *CognitoSyncActivities) DeleteDataset(input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error) {
	return a.client.DeleteDataset(input)
}

func (a *CognitoSyncActivities) DescribeDataset(input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error) {
	return a.client.DescribeDataset(input)
}

func (a *CognitoSyncActivities) DescribeIdentityPoolUsage(input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	return a.client.DescribeIdentityPoolUsage(input)
}

func (a *CognitoSyncActivities) DescribeIdentityUsage(input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error) {
	return a.client.DescribeIdentityUsage(input)
}

func (a *CognitoSyncActivities) GetBulkPublishDetails(input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	return a.client.GetBulkPublishDetails(input)
}

func (a *CognitoSyncActivities) GetCognitoEvents(input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error) {
	return a.client.GetCognitoEvents(input)
}

func (a *CognitoSyncActivities) GetIdentityPoolConfiguration(input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	return a.client.GetIdentityPoolConfiguration(input)
}

func (a *CognitoSyncActivities) ListDatasets(input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error) {
	return a.client.ListDatasets(input)
}

func (a *CognitoSyncActivities) ListIdentityPoolUsage(input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	return a.client.ListIdentityPoolUsage(input)
}

func (a *CognitoSyncActivities) ListRecords(input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error) {
	return a.client.ListRecords(input)
}

func (a *CognitoSyncActivities) RegisterDevice(input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error) {
	return a.client.RegisterDevice(input)
}

func (a *CognitoSyncActivities) SetCognitoEvents(input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error) {
	return a.client.SetCognitoEvents(input)
}

func (a *CognitoSyncActivities) SetIdentityPoolConfiguration(input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
	return a.client.SetIdentityPoolConfiguration(input)
}

func (a *CognitoSyncActivities) SubscribeToDataset(input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error) {
	return a.client.SubscribeToDataset(input)
}

func (a *CognitoSyncActivities) UnsubscribeFromDataset(input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
	return a.client.UnsubscribeFromDataset(input)
}

func (a *CognitoSyncActivities) UpdateRecords(input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error) {
	return a.client.UpdateRecords(input)
}
