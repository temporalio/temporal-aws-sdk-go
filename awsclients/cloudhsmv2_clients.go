package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CloudHSMV2Client interface {
    CopyBackupToRegion(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error)
    CopyBackupToRegionAsync(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) *Cloudhsmv2CopyBackupToRegionResult

    CreateCluster(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error)
    CreateClusterAsync(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) *Cloudhsmv2CreateClusterResult

    CreateHsm(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error)
    CreateHsmAsync(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) *Cloudhsmv2CreateHsmResult

    DeleteBackup(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error)
    DeleteBackupAsync(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) *Cloudhsmv2DeleteBackupResult

    DeleteCluster(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error)
    DeleteClusterAsync(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) *Cloudhsmv2DeleteClusterResult

    DeleteHsm(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error)
    DeleteHsmAsync(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) *Cloudhsmv2DeleteHsmResult

    DescribeBackups(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error)
    DescribeBackupsAsync(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) *Cloudhsmv2DescribeBackupsResult

    DescribeClusters(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error)
    DescribeClustersAsync(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) *Cloudhsmv2DescribeClustersResult

    InitializeCluster(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error)
    InitializeClusterAsync(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) *Cloudhsmv2InitializeClusterResult

    ListTags(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) *Cloudhsmv2ListTagsResult

    RestoreBackup(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error)
    RestoreBackupAsync(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) *Cloudhsmv2RestoreBackupResult

    TagResource(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) *Cloudhsmv2TagResourceResult

    UntagResource(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) *Cloudhsmv2UntagResourceResult
}
type Cloudhsmv2CopyBackupToRegionResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2CopyBackupToRegionResult) Get(ctx workflow.Context) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
    var output cloudhsmv2.CopyBackupToRegionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2CreateClusterResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2CreateClusterResult) Get(ctx workflow.Context) (*cloudhsmv2.CreateClusterOutput, error) {
    var output cloudhsmv2.CreateClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2CreateHsmResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2CreateHsmResult) Get(ctx workflow.Context) (*cloudhsmv2.CreateHsmOutput, error) {
    var output cloudhsmv2.CreateHsmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2DeleteBackupResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2DeleteBackupResult) Get(ctx workflow.Context) (*cloudhsmv2.DeleteBackupOutput, error) {
    var output cloudhsmv2.DeleteBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2DeleteClusterResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2DeleteClusterResult) Get(ctx workflow.Context) (*cloudhsmv2.DeleteClusterOutput, error) {
    var output cloudhsmv2.DeleteClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2DeleteHsmResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2DeleteHsmResult) Get(ctx workflow.Context) (*cloudhsmv2.DeleteHsmOutput, error) {
    var output cloudhsmv2.DeleteHsmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2DescribeBackupsResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2DescribeBackupsResult) Get(ctx workflow.Context) (*cloudhsmv2.DescribeBackupsOutput, error) {
    var output cloudhsmv2.DescribeBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2DescribeClustersResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2DescribeClustersResult) Get(ctx workflow.Context) (*cloudhsmv2.DescribeClustersOutput, error) {
    var output cloudhsmv2.DescribeClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2InitializeClusterResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2InitializeClusterResult) Get(ctx workflow.Context) (*cloudhsmv2.InitializeClusterOutput, error) {
    var output cloudhsmv2.InitializeClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2ListTagsResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2ListTagsResult) Get(ctx workflow.Context) (*cloudhsmv2.ListTagsOutput, error) {
    var output cloudhsmv2.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2RestoreBackupResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2RestoreBackupResult) Get(ctx workflow.Context) (*cloudhsmv2.RestoreBackupOutput, error) {
    var output cloudhsmv2.RestoreBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2TagResourceResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2TagResourceResult) Get(ctx workflow.Context) (*cloudhsmv2.TagResourceOutput, error) {
    var output cloudhsmv2.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Cloudhsmv2UntagResourceResult struct {
	Result workflow.Future
}

func (r *Cloudhsmv2UntagResourceResult) Get(ctx workflow.Context) (*cloudhsmv2.UntagResourceOutput, error) {
    var output cloudhsmv2.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CloudHSMV2Stub struct {
    activities awsactivities.CloudHSMV2Activities
}

func NewCloudHSMV2Stub() CloudHSMV2Client {
    return &CloudHSMV2Stub{}
}
func (a *CloudHSMV2Stub) CopyBackupToRegion(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
    var output cloudhsmv2.CopyBackupToRegionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyBackupToRegion, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) CopyBackupToRegionAsync(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) *Cloudhsmv2CopyBackupToRegionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyBackupToRegion, input)
    return &Cloudhsmv2CopyBackupToRegionResult{Result: future}
}
func (a *CloudHSMV2Stub) CreateCluster(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
    var output cloudhsmv2.CreateClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) CreateClusterAsync(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) *Cloudhsmv2CreateClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCluster, input)
    return &Cloudhsmv2CreateClusterResult{Result: future}
}
func (a *CloudHSMV2Stub) CreateHsm(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
    var output cloudhsmv2.CreateHsmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHsm, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) CreateHsmAsync(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) *Cloudhsmv2CreateHsmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHsm, input)
    return &Cloudhsmv2CreateHsmResult{Result: future}
}
func (a *CloudHSMV2Stub) DeleteBackup(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
    var output cloudhsmv2.DeleteBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) DeleteBackupAsync(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) *Cloudhsmv2DeleteBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackup, input)
    return &Cloudhsmv2DeleteBackupResult{Result: future}
}
func (a *CloudHSMV2Stub) DeleteCluster(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
    var output cloudhsmv2.DeleteClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) DeleteClusterAsync(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) *Cloudhsmv2DeleteClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCluster, input)
    return &Cloudhsmv2DeleteClusterResult{Result: future}
}
func (a *CloudHSMV2Stub) DeleteHsm(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
    var output cloudhsmv2.DeleteHsmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHsm, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) *Cloudhsmv2DeleteHsmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHsm, input)
    return &Cloudhsmv2DeleteHsmResult{Result: future}
}
func (a *CloudHSMV2Stub) DescribeBackups(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
    var output cloudhsmv2.DescribeBackupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBackups, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) DescribeBackupsAsync(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) *Cloudhsmv2DescribeBackupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBackups, input)
    return &Cloudhsmv2DescribeBackupsResult{Result: future}
}
func (a *CloudHSMV2Stub) DescribeClusters(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
    var output cloudhsmv2.DescribeClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) DescribeClustersAsync(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) *Cloudhsmv2DescribeClustersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeClusters, input)
    return &Cloudhsmv2DescribeClustersResult{Result: future}
}
func (a *CloudHSMV2Stub) InitializeCluster(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
    var output cloudhsmv2.InitializeClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitializeCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) InitializeClusterAsync(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) *Cloudhsmv2InitializeClusterResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InitializeCluster, input)
    return &Cloudhsmv2InitializeClusterResult{Result: future}
}
func (a *CloudHSMV2Stub) ListTags(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
    var output cloudhsmv2.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) ListTagsAsync(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) *Cloudhsmv2ListTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
    return &Cloudhsmv2ListTagsResult{Result: future}
}
func (a *CloudHSMV2Stub) RestoreBackup(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
    var output cloudhsmv2.RestoreBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) RestoreBackupAsync(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) *Cloudhsmv2RestoreBackupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestoreBackup, input)
    return &Cloudhsmv2RestoreBackupResult{Result: future}
}
func (a *CloudHSMV2Stub) TagResource(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
    var output cloudhsmv2.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) TagResourceAsync(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) *Cloudhsmv2TagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Cloudhsmv2TagResourceResult{Result: future}
}
func (a *CloudHSMV2Stub) UntagResource(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
    var output cloudhsmv2.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMV2Stub) UntagResourceAsync(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) *Cloudhsmv2UntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Cloudhsmv2UntagResourceResult{Result: future}
}
