package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"go.temporal.io/sdk/workflow"
)

type CognitoSyncClient interface {
    BulkPublish(ctx workflow.Context, input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)
    BulkPublishAsync(ctx workflow.Context, input *cognitosync.BulkPublishInput) *CognitosyncBulkPublishResult

    DeleteDataset(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)
    DeleteDatasetAsync(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) *CognitosyncDeleteDatasetResult

    DescribeDataset(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)
    DescribeDatasetAsync(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) *CognitosyncDescribeDatasetResult

    DescribeIdentityPoolUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)
    DescribeIdentityPoolUsageAsync(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) *CognitosyncDescribeIdentityPoolUsageResult

    DescribeIdentityUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)
    DescribeIdentityUsageAsync(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) *CognitosyncDescribeIdentityUsageResult

    GetBulkPublishDetails(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)
    GetBulkPublishDetailsAsync(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) *CognitosyncGetBulkPublishDetailsResult

    GetCognitoEvents(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error)
    GetCognitoEventsAsync(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) *CognitosyncGetCognitoEventsResult

    GetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)
    GetIdentityPoolConfigurationAsync(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) *CognitosyncGetIdentityPoolConfigurationResult

    ListDatasets(ctx workflow.Context, input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)
    ListDatasetsAsync(ctx workflow.Context, input *cognitosync.ListDatasetsInput) *CognitosyncListDatasetsResult

    ListIdentityPoolUsage(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)
    ListIdentityPoolUsageAsync(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) *CognitosyncListIdentityPoolUsageResult

    ListRecords(ctx workflow.Context, input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)
    ListRecordsAsync(ctx workflow.Context, input *cognitosync.ListRecordsInput) *CognitosyncListRecordsResult

    RegisterDevice(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)
    RegisterDeviceAsync(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) *CognitosyncRegisterDeviceResult

    SetCognitoEvents(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error)
    SetCognitoEventsAsync(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) *CognitosyncSetCognitoEventsResult

    SetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)
    SetIdentityPoolConfigurationAsync(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) *CognitosyncSetIdentityPoolConfigurationResult

    SubscribeToDataset(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)
    SubscribeToDatasetAsync(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) *CognitosyncSubscribeToDatasetResult

    UnsubscribeFromDataset(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)
    UnsubscribeFromDatasetAsync(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) *CognitosyncUnsubscribeFromDatasetResult

    UpdateRecords(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
    UpdateRecordsAsync(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) *CognitosyncUpdateRecordsResult
}
type CognitosyncBulkPublishResult struct {
	Result workflow.Future
}

func (r *CognitosyncBulkPublishResult) Get(ctx workflow.Context) (*cognitosync.BulkPublishOutput, error) {
    var output cognitosync.BulkPublishOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncDeleteDatasetResult struct {
	Result workflow.Future
}

func (r *CognitosyncDeleteDatasetResult) Get(ctx workflow.Context) (*cognitosync.DeleteDatasetOutput, error) {
    var output cognitosync.DeleteDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncDescribeDatasetResult struct {
	Result workflow.Future
}

func (r *CognitosyncDescribeDatasetResult) Get(ctx workflow.Context) (*cognitosync.DescribeDatasetOutput, error) {
    var output cognitosync.DescribeDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncDescribeIdentityPoolUsageResult struct {
	Result workflow.Future
}

func (r *CognitosyncDescribeIdentityPoolUsageResult) Get(ctx workflow.Context) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
    var output cognitosync.DescribeIdentityPoolUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncDescribeIdentityUsageResult struct {
	Result workflow.Future
}

func (r *CognitosyncDescribeIdentityUsageResult) Get(ctx workflow.Context) (*cognitosync.DescribeIdentityUsageOutput, error) {
    var output cognitosync.DescribeIdentityUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncGetBulkPublishDetailsResult struct {
	Result workflow.Future
}

func (r *CognitosyncGetBulkPublishDetailsResult) Get(ctx workflow.Context) (*cognitosync.GetBulkPublishDetailsOutput, error) {
    var output cognitosync.GetBulkPublishDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncGetCognitoEventsResult struct {
	Result workflow.Future
}

func (r *CognitosyncGetCognitoEventsResult) Get(ctx workflow.Context) (*cognitosync.GetCognitoEventsOutput, error) {
    var output cognitosync.GetCognitoEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncGetIdentityPoolConfigurationResult struct {
	Result workflow.Future
}

func (r *CognitosyncGetIdentityPoolConfigurationResult) Get(ctx workflow.Context) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
    var output cognitosync.GetIdentityPoolConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncListDatasetsResult struct {
	Result workflow.Future
}

func (r *CognitosyncListDatasetsResult) Get(ctx workflow.Context) (*cognitosync.ListDatasetsOutput, error) {
    var output cognitosync.ListDatasetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncListIdentityPoolUsageResult struct {
	Result workflow.Future
}

func (r *CognitosyncListIdentityPoolUsageResult) Get(ctx workflow.Context) (*cognitosync.ListIdentityPoolUsageOutput, error) {
    var output cognitosync.ListIdentityPoolUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncListRecordsResult struct {
	Result workflow.Future
}

func (r *CognitosyncListRecordsResult) Get(ctx workflow.Context) (*cognitosync.ListRecordsOutput, error) {
    var output cognitosync.ListRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncRegisterDeviceResult struct {
	Result workflow.Future
}

func (r *CognitosyncRegisterDeviceResult) Get(ctx workflow.Context) (*cognitosync.RegisterDeviceOutput, error) {
    var output cognitosync.RegisterDeviceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncSetCognitoEventsResult struct {
	Result workflow.Future
}

func (r *CognitosyncSetCognitoEventsResult) Get(ctx workflow.Context) (*cognitosync.SetCognitoEventsOutput, error) {
    var output cognitosync.SetCognitoEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncSetIdentityPoolConfigurationResult struct {
	Result workflow.Future
}

func (r *CognitosyncSetIdentityPoolConfigurationResult) Get(ctx workflow.Context) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
    var output cognitosync.SetIdentityPoolConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncSubscribeToDatasetResult struct {
	Result workflow.Future
}

func (r *CognitosyncSubscribeToDatasetResult) Get(ctx workflow.Context) (*cognitosync.SubscribeToDatasetOutput, error) {
    var output cognitosync.SubscribeToDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncUnsubscribeFromDatasetResult struct {
	Result workflow.Future
}

func (r *CognitosyncUnsubscribeFromDatasetResult) Get(ctx workflow.Context) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
    var output cognitosync.UnsubscribeFromDatasetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CognitosyncUpdateRecordsResult struct {
	Result workflow.Future
}

func (r *CognitosyncUpdateRecordsResult) Get(ctx workflow.Context) (*cognitosync.UpdateRecordsOutput, error) {
    var output cognitosync.UpdateRecordsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CognitoSyncStub struct {
    activities CognitoSyncClient
}

func NewCognitoSyncStub() CognitoSyncClient {
    return &CognitoSyncStub{}
}

func (a *CognitoSyncStub) BulkPublish(ctx workflow.Context, input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
    var output cognitosync.BulkPublishOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BulkPublish, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) DeleteDataset(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error) {
    var output cognitosync.DeleteDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) DescribeDataset(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error) {
    var output cognitosync.DescribeDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) DescribeIdentityPoolUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
    var output cognitosync.DescribeIdentityPoolUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityPoolUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) DescribeIdentityUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error) {
    var output cognitosync.DescribeIdentityUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeIdentityUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) GetBulkPublishDetails(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error) {
    var output cognitosync.GetBulkPublishDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBulkPublishDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) GetCognitoEvents(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error) {
    var output cognitosync.GetCognitoEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCognitoEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) GetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
    var output cognitosync.GetIdentityPoolConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityPoolConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) ListDatasets(ctx workflow.Context, input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error) {
    var output cognitosync.ListDatasetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDatasets, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) ListIdentityPoolUsage(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error) {
    var output cognitosync.ListIdentityPoolUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIdentityPoolUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) ListRecords(ctx workflow.Context, input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error) {
    var output cognitosync.ListRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecords, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) RegisterDevice(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error) {
    var output cognitosync.RegisterDeviceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterDevice, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) SetCognitoEvents(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error) {
    var output cognitosync.SetCognitoEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetCognitoEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) SetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
    var output cognitosync.SetIdentityPoolConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityPoolConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) SubscribeToDataset(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error) {
    var output cognitosync.SubscribeToDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SubscribeToDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) UnsubscribeFromDataset(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
    var output cognitosync.UnsubscribeFromDatasetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnsubscribeFromDataset, input).Get(ctx, &output)
    return &output, err
}

func (a *CognitoSyncStub) UpdateRecords(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error) {
    var output cognitosync.UpdateRecordsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRecords, input).Get(ctx, &output)
    return &output, err
}
