package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type KinesisAnalyticsV2Client interface {
    AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error)
    AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) *Kinesisanalyticsv2AddApplicationCloudWatchLoggingOptionResult

    AddApplicationInput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error)
    AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) *Kinesisanalyticsv2AddApplicationInputResult

    AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error)
    AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) *Kinesisanalyticsv2AddApplicationInputProcessingConfigurationResult

    AddApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error)
    AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) *Kinesisanalyticsv2AddApplicationOutputResult

    AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error)
    AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) *Kinesisanalyticsv2AddApplicationReferenceDataSourceResult

    AddApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error)
    AddApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) *Kinesisanalyticsv2AddApplicationVpcConfigurationResult

    CreateApplication(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error)
    CreateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) *Kinesisanalyticsv2CreateApplicationResult

    CreateApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error)
    CreateApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) *Kinesisanalyticsv2CreateApplicationSnapshotResult

    DeleteApplication(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error)
    DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) *Kinesisanalyticsv2DeleteApplicationResult

    DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error)
    DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) *Kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOptionResult

    DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error)
    DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) *Kinesisanalyticsv2DeleteApplicationInputProcessingConfigurationResult

    DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error)
    DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) *Kinesisanalyticsv2DeleteApplicationOutputResult

    DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error)
    DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) *Kinesisanalyticsv2DeleteApplicationReferenceDataSourceResult

    DeleteApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error)
    DeleteApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) *Kinesisanalyticsv2DeleteApplicationSnapshotResult

    DeleteApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error)
    DeleteApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) *Kinesisanalyticsv2DeleteApplicationVpcConfigurationResult

    DescribeApplication(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error)
    DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) *Kinesisanalyticsv2DescribeApplicationResult

    DescribeApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error)
    DescribeApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) *Kinesisanalyticsv2DescribeApplicationSnapshotResult

    DiscoverInputSchema(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error)
    DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) *Kinesisanalyticsv2DiscoverInputSchemaResult

    ListApplicationSnapshots(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error)
    ListApplicationSnapshotsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) *Kinesisanalyticsv2ListApplicationSnapshotsResult

    ListApplications(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error)
    ListApplicationsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) *Kinesisanalyticsv2ListApplicationsResult

    ListTagsForResource(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) *Kinesisanalyticsv2ListTagsForResourceResult

    StartApplication(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error)
    StartApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) *Kinesisanalyticsv2StartApplicationResult

    StopApplication(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error)
    StopApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) *Kinesisanalyticsv2StopApplicationResult

    TagResource(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) *Kinesisanalyticsv2TagResourceResult

    UntagResource(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) *Kinesisanalyticsv2UntagResourceResult

    UpdateApplication(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error)
    UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) *Kinesisanalyticsv2UpdateApplicationResult
}
type Kinesisanalyticsv2AddApplicationCloudWatchLoggingOptionResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2AddApplicationCloudWatchLoggingOptionResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
    var output kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2AddApplicationInputResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2AddApplicationInputResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationInputOutput, error) {
    var output kinesisanalyticsv2.AddApplicationInputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2AddApplicationInputProcessingConfigurationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2AddApplicationInputProcessingConfigurationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error) {
    var output kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2AddApplicationOutputResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2AddApplicationOutputResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationOutputOutput, error) {
    var output kinesisanalyticsv2.AddApplicationOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2AddApplicationReferenceDataSourceResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2AddApplicationReferenceDataSourceResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error) {
    var output kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2AddApplicationVpcConfigurationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2AddApplicationVpcConfigurationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error) {
    var output kinesisanalyticsv2.AddApplicationVpcConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2CreateApplicationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2CreateApplicationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.CreateApplicationOutput, error) {
    var output kinesisanalyticsv2.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2CreateApplicationSnapshotResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2CreateApplicationSnapshotResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error) {
    var output kinesisanalyticsv2.CreateApplicationSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOptionResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOptionResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationInputProcessingConfigurationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationInputProcessingConfigurationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationOutputResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationOutputResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationReferenceDataSourceResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationReferenceDataSourceResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationSnapshotResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationSnapshotResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DeleteApplicationVpcConfigurationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DeleteApplicationVpcConfigurationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DescribeApplicationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DescribeApplicationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
    var output kinesisanalyticsv2.DescribeApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DescribeApplicationSnapshotResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DescribeApplicationSnapshotResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
    var output kinesisanalyticsv2.DescribeApplicationSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2DiscoverInputSchemaResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2DiscoverInputSchemaResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error) {
    var output kinesisanalyticsv2.DiscoverInputSchemaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2ListApplicationSnapshotsResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2ListApplicationSnapshotsResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
    var output kinesisanalyticsv2.ListApplicationSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2ListApplicationsResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2ListApplicationsResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
    var output kinesisanalyticsv2.ListApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2ListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2ListTagsForResourceResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
    var output kinesisanalyticsv2.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2StartApplicationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2StartApplicationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.StartApplicationOutput, error) {
    var output kinesisanalyticsv2.StartApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2StopApplicationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2StopApplicationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.StopApplicationOutput, error) {
    var output kinesisanalyticsv2.StopApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2TagResourceResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2TagResourceResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.TagResourceOutput, error) {
    var output kinesisanalyticsv2.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2UntagResourceResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2UntagResourceResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.UntagResourceOutput, error) {
    var output kinesisanalyticsv2.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Kinesisanalyticsv2UpdateApplicationResult struct {
	Result workflow.Future
}

func (r *Kinesisanalyticsv2UpdateApplicationResult) Get(ctx workflow.Context) (*kinesisanalyticsv2.UpdateApplicationOutput, error) {
    var output kinesisanalyticsv2.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type KinesisAnalyticsV2Stub struct {
    activities awsactivities.KinesisAnalyticsV2Activities
}

func NewKinesisAnalyticsV2Stub() KinesisAnalyticsV2Client {
    return &KinesisAnalyticsV2Stub{}
}
func (a *KinesisAnalyticsV2Stub) AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
    var output kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationCloudWatchLoggingOption, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) *Kinesisanalyticsv2AddApplicationCloudWatchLoggingOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationCloudWatchLoggingOption, input)
    return &Kinesisanalyticsv2AddApplicationCloudWatchLoggingOptionResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) AddApplicationInput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error) {
    var output kinesisanalyticsv2.AddApplicationInputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInput, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) *Kinesisanalyticsv2AddApplicationInputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInput, input)
    return &Kinesisanalyticsv2AddApplicationInputResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error) {
    var output kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInputProcessingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) *Kinesisanalyticsv2AddApplicationInputProcessingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationInputProcessingConfiguration, input)
    return &Kinesisanalyticsv2AddApplicationInputProcessingConfigurationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) AddApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error) {
    var output kinesisanalyticsv2.AddApplicationOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) *Kinesisanalyticsv2AddApplicationOutputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationOutput, input)
    return &Kinesisanalyticsv2AddApplicationOutputResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error) {
    var output kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationReferenceDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) *Kinesisanalyticsv2AddApplicationReferenceDataSourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationReferenceDataSource, input)
    return &Kinesisanalyticsv2AddApplicationReferenceDataSourceResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) AddApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error) {
    var output kinesisanalyticsv2.AddApplicationVpcConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddApplicationVpcConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) *Kinesisanalyticsv2AddApplicationVpcConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddApplicationVpcConfiguration, input)
    return &Kinesisanalyticsv2AddApplicationVpcConfigurationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) CreateApplication(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error) {
    var output kinesisanalyticsv2.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) CreateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) *Kinesisanalyticsv2CreateApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input)
    return &Kinesisanalyticsv2CreateApplicationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) CreateApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error) {
    var output kinesisanalyticsv2.CreateApplicationSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplicationSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) CreateApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) *Kinesisanalyticsv2CreateApplicationSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApplicationSnapshot, input)
    return &Kinesisanalyticsv2CreateApplicationSnapshotResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplication(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) *Kinesisanalyticsv2DeleteApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input)
    return &Kinesisanalyticsv2DeleteApplicationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationCloudWatchLoggingOption, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) *Kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationCloudWatchLoggingOption, input)
    return &Kinesisanalyticsv2DeleteApplicationCloudWatchLoggingOptionResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationInputProcessingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) *Kinesisanalyticsv2DeleteApplicationInputProcessingConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationInputProcessingConfiguration, input)
    return &Kinesisanalyticsv2DeleteApplicationInputProcessingConfigurationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) *Kinesisanalyticsv2DeleteApplicationOutputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationOutput, input)
    return &Kinesisanalyticsv2DeleteApplicationOutputResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationReferenceDataSource, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) *Kinesisanalyticsv2DeleteApplicationReferenceDataSourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationReferenceDataSource, input)
    return &Kinesisanalyticsv2DeleteApplicationReferenceDataSourceResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) *Kinesisanalyticsv2DeleteApplicationSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationSnapshot, input)
    return &Kinesisanalyticsv2DeleteApplicationSnapshotResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DeleteApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error) {
    var output kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationVpcConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) *Kinesisanalyticsv2DeleteApplicationVpcConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationVpcConfiguration, input)
    return &Kinesisanalyticsv2DeleteApplicationVpcConfigurationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DescribeApplication(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
    var output kinesisanalyticsv2.DescribeApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) *Kinesisanalyticsv2DescribeApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeApplication, input)
    return &Kinesisanalyticsv2DescribeApplicationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DescribeApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
    var output kinesisanalyticsv2.DescribeApplicationSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApplicationSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DescribeApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) *Kinesisanalyticsv2DescribeApplicationSnapshotResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeApplicationSnapshot, input)
    return &Kinesisanalyticsv2DescribeApplicationSnapshotResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) DiscoverInputSchema(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error) {
    var output kinesisanalyticsv2.DiscoverInputSchemaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DiscoverInputSchema, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) *Kinesisanalyticsv2DiscoverInputSchemaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DiscoverInputSchema, input)
    return &Kinesisanalyticsv2DiscoverInputSchemaResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) ListApplicationSnapshots(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
    var output kinesisanalyticsv2.ListApplicationSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplicationSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) ListApplicationSnapshotsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) *Kinesisanalyticsv2ListApplicationSnapshotsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListApplicationSnapshots, input)
    return &Kinesisanalyticsv2ListApplicationSnapshotsResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) ListApplications(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
    var output kinesisanalyticsv2.ListApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) ListApplicationsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) *Kinesisanalyticsv2ListApplicationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input)
    return &Kinesisanalyticsv2ListApplicationsResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) ListTagsForResource(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
    var output kinesisanalyticsv2.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) *Kinesisanalyticsv2ListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Kinesisanalyticsv2ListTagsForResourceResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) StartApplication(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error) {
    var output kinesisanalyticsv2.StartApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) StartApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) *Kinesisanalyticsv2StartApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartApplication, input)
    return &Kinesisanalyticsv2StartApplicationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) StopApplication(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error) {
    var output kinesisanalyticsv2.StopApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) StopApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) *Kinesisanalyticsv2StopApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopApplication, input)
    return &Kinesisanalyticsv2StopApplicationResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) TagResource(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error) {
    var output kinesisanalyticsv2.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) TagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) *Kinesisanalyticsv2TagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Kinesisanalyticsv2TagResourceResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) UntagResource(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error) {
    var output kinesisanalyticsv2.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) UntagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) *Kinesisanalyticsv2UntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Kinesisanalyticsv2UntagResourceResult{Result: future}
}
func (a *KinesisAnalyticsV2Stub) UpdateApplication(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error) {
    var output kinesisanalyticsv2.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *KinesisAnalyticsV2Stub) UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) *Kinesisanalyticsv2UpdateApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input)
    return &Kinesisanalyticsv2UpdateApplicationResult{Result: future}
}
