package awsclients

import (
	"github.com/aws/aws-sdk-go/service/snowball"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SnowballClient interface {
	CancelCluster(ctx workflow.Context, input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error)
	CancelClusterAsync(ctx workflow.Context, input *snowball.CancelClusterInput) *SnowballCancelClusterResult

	CancelJob(ctx workflow.Context, input *snowball.CancelJobInput) (*snowball.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *snowball.CancelJobInput) *SnowballCancelJobResult

	CreateAddress(ctx workflow.Context, input *snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error)
	CreateAddressAsync(ctx workflow.Context, input *snowball.CreateAddressInput) *SnowballCreateAddressResult

	CreateCluster(ctx workflow.Context, input *snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *snowball.CreateClusterInput) *SnowballCreateClusterResult

	CreateJob(ctx workflow.Context, input *snowball.CreateJobInput) (*snowball.CreateJobOutput, error)
	CreateJobAsync(ctx workflow.Context, input *snowball.CreateJobInput) *SnowballCreateJobResult

	DescribeAddress(ctx workflow.Context, input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error)
	DescribeAddressAsync(ctx workflow.Context, input *snowball.DescribeAddressInput) *SnowballDescribeAddressResult

	DescribeAddresses(ctx workflow.Context, input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error)
	DescribeAddressesAsync(ctx workflow.Context, input *snowball.DescribeAddressesInput) *SnowballDescribeAddressesResult

	DescribeCluster(ctx workflow.Context, input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error)
	DescribeClusterAsync(ctx workflow.Context, input *snowball.DescribeClusterInput) *SnowballDescribeClusterResult

	DescribeJob(ctx workflow.Context, input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error)
	DescribeJobAsync(ctx workflow.Context, input *snowball.DescribeJobInput) *SnowballDescribeJobResult

	GetJobManifest(ctx workflow.Context, input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error)
	GetJobManifestAsync(ctx workflow.Context, input *snowball.GetJobManifestInput) *SnowballGetJobManifestResult

	GetJobUnlockCode(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error)
	GetJobUnlockCodeAsync(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) *SnowballGetJobUnlockCodeResult

	GetSnowballUsage(ctx workflow.Context, input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error)
	GetSnowballUsageAsync(ctx workflow.Context, input *snowball.GetSnowballUsageInput) *SnowballGetSnowballUsageResult

	GetSoftwareUpdates(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error)
	GetSoftwareUpdatesAsync(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) *SnowballGetSoftwareUpdatesResult

	ListClusterJobs(ctx workflow.Context, input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error)
	ListClusterJobsAsync(ctx workflow.Context, input *snowball.ListClusterJobsInput) *SnowballListClusterJobsResult

	ListClusters(ctx workflow.Context, input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *snowball.ListClustersInput) *SnowballListClustersResult

	ListCompatibleImages(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error)
	ListCompatibleImagesAsync(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) *SnowballListCompatibleImagesResult

	ListJobs(ctx workflow.Context, input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error)
	ListJobsAsync(ctx workflow.Context, input *snowball.ListJobsInput) *SnowballListJobsResult

	UpdateCluster(ctx workflow.Context, input *snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error)
	UpdateClusterAsync(ctx workflow.Context, input *snowball.UpdateClusterInput) *SnowballUpdateClusterResult

	UpdateJob(ctx workflow.Context, input *snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error)
	UpdateJobAsync(ctx workflow.Context, input *snowball.UpdateJobInput) *SnowballUpdateJobResult
}

type SnowballCancelClusterResult struct {
	Result workflow.Future
}

func (r *SnowballCancelClusterResult) Get(ctx workflow.Context) (*snowball.CancelClusterOutput, error) {
	var output snowball.CancelClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballCancelJobResult struct {
	Result workflow.Future
}

func (r *SnowballCancelJobResult) Get(ctx workflow.Context) (*snowball.CancelJobOutput, error) {
	var output snowball.CancelJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballCreateAddressResult struct {
	Result workflow.Future
}

func (r *SnowballCreateAddressResult) Get(ctx workflow.Context) (*snowball.CreateAddressOutput, error) {
	var output snowball.CreateAddressOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballCreateClusterResult struct {
	Result workflow.Future
}

func (r *SnowballCreateClusterResult) Get(ctx workflow.Context) (*snowball.CreateClusterOutput, error) {
	var output snowball.CreateClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballCreateJobResult struct {
	Result workflow.Future
}

func (r *SnowballCreateJobResult) Get(ctx workflow.Context) (*snowball.CreateJobOutput, error) {
	var output snowball.CreateJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeAddressResult struct {
	Result workflow.Future
}

func (r *SnowballDescribeAddressResult) Get(ctx workflow.Context) (*snowball.DescribeAddressOutput, error) {
	var output snowball.DescribeAddressOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeAddressesResult struct {
	Result workflow.Future
}

func (r *SnowballDescribeAddressesResult) Get(ctx workflow.Context) (*snowball.DescribeAddressesOutput, error) {
	var output snowball.DescribeAddressesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeClusterResult struct {
	Result workflow.Future
}

func (r *SnowballDescribeClusterResult) Get(ctx workflow.Context) (*snowball.DescribeClusterOutput, error) {
	var output snowball.DescribeClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballDescribeJobResult struct {
	Result workflow.Future
}

func (r *SnowballDescribeJobResult) Get(ctx workflow.Context) (*snowball.DescribeJobOutput, error) {
	var output snowball.DescribeJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballGetJobManifestResult struct {
	Result workflow.Future
}

func (r *SnowballGetJobManifestResult) Get(ctx workflow.Context) (*snowball.GetJobManifestOutput, error) {
	var output snowball.GetJobManifestOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballGetJobUnlockCodeResult struct {
	Result workflow.Future
}

func (r *SnowballGetJobUnlockCodeResult) Get(ctx workflow.Context) (*snowball.GetJobUnlockCodeOutput, error) {
	var output snowball.GetJobUnlockCodeOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballGetSnowballUsageResult struct {
	Result workflow.Future
}

func (r *SnowballGetSnowballUsageResult) Get(ctx workflow.Context) (*snowball.GetSnowballUsageOutput, error) {
	var output snowball.GetSnowballUsageOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballGetSoftwareUpdatesResult struct {
	Result workflow.Future
}

func (r *SnowballGetSoftwareUpdatesResult) Get(ctx workflow.Context) (*snowball.GetSoftwareUpdatesOutput, error) {
	var output snowball.GetSoftwareUpdatesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballListClusterJobsResult struct {
	Result workflow.Future
}

func (r *SnowballListClusterJobsResult) Get(ctx workflow.Context) (*snowball.ListClusterJobsOutput, error) {
	var output snowball.ListClusterJobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballListClustersResult struct {
	Result workflow.Future
}

func (r *SnowballListClustersResult) Get(ctx workflow.Context) (*snowball.ListClustersOutput, error) {
	var output snowball.ListClustersOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballListCompatibleImagesResult struct {
	Result workflow.Future
}

func (r *SnowballListCompatibleImagesResult) Get(ctx workflow.Context) (*snowball.ListCompatibleImagesOutput, error) {
	var output snowball.ListCompatibleImagesOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballListJobsResult struct {
	Result workflow.Future
}

func (r *SnowballListJobsResult) Get(ctx workflow.Context) (*snowball.ListJobsOutput, error) {
	var output snowball.ListJobsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballUpdateClusterResult struct {
	Result workflow.Future
}

func (r *SnowballUpdateClusterResult) Get(ctx workflow.Context) (*snowball.UpdateClusterOutput, error) {
	var output snowball.UpdateClusterOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballUpdateJobResult struct {
	Result workflow.Future
}

func (r *SnowballUpdateJobResult) Get(ctx workflow.Context) (*snowball.UpdateJobOutput, error) {
	var output snowball.UpdateJobOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type SnowballStub struct {
	activities awsactivities.SnowballActivities
}

func NewSnowballStub() SnowballClient {
	return &SnowballStub{}
}

func (a *SnowballStub) CancelCluster(ctx workflow.Context, input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error) {
	var output snowball.CancelClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CancelCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) CancelClusterAsync(ctx workflow.Context, input *snowball.CancelClusterInput) *SnowballCancelClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CancelCluster, input)
	return &SnowballCancelClusterResult{Result: future}
}

func (a *SnowballStub) CancelJob(ctx workflow.Context, input *snowball.CancelJobInput) (*snowball.CancelJobOutput, error) {
	var output snowball.CancelJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) CancelJobAsync(ctx workflow.Context, input *snowball.CancelJobInput) *SnowballCancelJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input)
	return &SnowballCancelJobResult{Result: future}
}

func (a *SnowballStub) CreateAddress(ctx workflow.Context, input *snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error) {
	var output snowball.CreateAddressOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateAddress, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) CreateAddressAsync(ctx workflow.Context, input *snowball.CreateAddressInput) *SnowballCreateAddressResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateAddress, input)
	return &SnowballCreateAddressResult{Result: future}
}

func (a *SnowballStub) CreateCluster(ctx workflow.Context, input *snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error) {
	var output snowball.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) CreateClusterAsync(ctx workflow.Context, input *snowball.CreateClusterInput) *SnowballCreateClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input)
	return &SnowballCreateClusterResult{Result: future}
}

func (a *SnowballStub) CreateJob(ctx workflow.Context, input *snowball.CreateJobInput) (*snowball.CreateJobOutput, error) {
	var output snowball.CreateJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) CreateJobAsync(ctx workflow.Context, input *snowball.CreateJobInput) *SnowballCreateJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
	return &SnowballCreateJobResult{Result: future}
}

func (a *SnowballStub) DescribeAddress(ctx workflow.Context, input *snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error) {
	var output snowball.DescribeAddressOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAddress, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) DescribeAddressAsync(ctx workflow.Context, input *snowball.DescribeAddressInput) *SnowballDescribeAddressResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAddress, input)
	return &SnowballDescribeAddressResult{Result: future}
}

func (a *SnowballStub) DescribeAddresses(ctx workflow.Context, input *snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error) {
	var output snowball.DescribeAddressesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeAddresses, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) DescribeAddressesAsync(ctx workflow.Context, input *snowball.DescribeAddressesInput) *SnowballDescribeAddressesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeAddresses, input)
	return &SnowballDescribeAddressesResult{Result: future}
}

func (a *SnowballStub) DescribeCluster(ctx workflow.Context, input *snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error) {
	var output snowball.DescribeClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) DescribeClusterAsync(ctx workflow.Context, input *snowball.DescribeClusterInput) *SnowballDescribeClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeCluster, input)
	return &SnowballDescribeClusterResult{Result: future}
}

func (a *SnowballStub) DescribeJob(ctx workflow.Context, input *snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error) {
	var output snowball.DescribeJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) DescribeJobAsync(ctx workflow.Context, input *snowball.DescribeJobInput) *SnowballDescribeJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input)
	return &SnowballDescribeJobResult{Result: future}
}

func (a *SnowballStub) GetJobManifest(ctx workflow.Context, input *snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error) {
	var output snowball.GetJobManifestOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetJobManifest, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) GetJobManifestAsync(ctx workflow.Context, input *snowball.GetJobManifestInput) *SnowballGetJobManifestResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetJobManifest, input)
	return &SnowballGetJobManifestResult{Result: future}
}

func (a *SnowballStub) GetJobUnlockCode(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error) {
	var output snowball.GetJobUnlockCodeOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetJobUnlockCode, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) GetJobUnlockCodeAsync(ctx workflow.Context, input *snowball.GetJobUnlockCodeInput) *SnowballGetJobUnlockCodeResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetJobUnlockCode, input)
	return &SnowballGetJobUnlockCodeResult{Result: future}
}

func (a *SnowballStub) GetSnowballUsage(ctx workflow.Context, input *snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error) {
	var output snowball.GetSnowballUsageOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSnowballUsage, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) GetSnowballUsageAsync(ctx workflow.Context, input *snowball.GetSnowballUsageInput) *SnowballGetSnowballUsageResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSnowballUsage, input)
	return &SnowballGetSnowballUsageResult{Result: future}
}

func (a *SnowballStub) GetSoftwareUpdates(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) (*snowball.GetSoftwareUpdatesOutput, error) {
	var output snowball.GetSoftwareUpdatesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetSoftwareUpdates, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) GetSoftwareUpdatesAsync(ctx workflow.Context, input *snowball.GetSoftwareUpdatesInput) *SnowballGetSoftwareUpdatesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetSoftwareUpdates, input)
	return &SnowballGetSoftwareUpdatesResult{Result: future}
}

func (a *SnowballStub) ListClusterJobs(ctx workflow.Context, input *snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error) {
	var output snowball.ListClusterJobsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListClusterJobs, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) ListClusterJobsAsync(ctx workflow.Context, input *snowball.ListClusterJobsInput) *SnowballListClusterJobsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListClusterJobs, input)
	return &SnowballListClusterJobsResult{Result: future}
}

func (a *SnowballStub) ListClusters(ctx workflow.Context, input *snowball.ListClustersInput) (*snowball.ListClustersOutput, error) {
	var output snowball.ListClustersOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) ListClustersAsync(ctx workflow.Context, input *snowball.ListClustersInput) *SnowballListClustersResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListClusters, input)
	return &SnowballListClustersResult{Result: future}
}

func (a *SnowballStub) ListCompatibleImages(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) (*snowball.ListCompatibleImagesOutput, error) {
	var output snowball.ListCompatibleImagesOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListCompatibleImages, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) ListCompatibleImagesAsync(ctx workflow.Context, input *snowball.ListCompatibleImagesInput) *SnowballListCompatibleImagesResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListCompatibleImages, input)
	return &SnowballListCompatibleImagesResult{Result: future}
}

func (a *SnowballStub) ListJobs(ctx workflow.Context, input *snowball.ListJobsInput) (*snowball.ListJobsOutput, error) {
	var output snowball.ListJobsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) ListJobsAsync(ctx workflow.Context, input *snowball.ListJobsInput) *SnowballListJobsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
	return &SnowballListJobsResult{Result: future}
}

func (a *SnowballStub) UpdateCluster(ctx workflow.Context, input *snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error) {
	var output snowball.UpdateClusterOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateCluster, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) UpdateClusterAsync(ctx workflow.Context, input *snowball.UpdateClusterInput) *SnowballUpdateClusterResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateCluster, input)
	return &SnowballUpdateClusterResult{Result: future}
}

func (a *SnowballStub) UpdateJob(ctx workflow.Context, input *snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error) {
	var output snowball.UpdateJobOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UpdateJob, input).Get(ctx, &output)
	return &output, err
}

func (a *SnowballStub) UpdateJobAsync(ctx workflow.Context, input *snowball.UpdateJobInput) *SnowballUpdateJobResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UpdateJob, input)
	return &SnowballUpdateJobResult{Result: future}
}
