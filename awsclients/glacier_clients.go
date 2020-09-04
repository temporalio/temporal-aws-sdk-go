package awsclients

import (
	"github.com/aws/aws-sdk-go/service/glacier"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type GlacierClient interface {
    AbortMultipartUpload(ctx workflow.Context, input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error)
    AbortMultipartUploadAsync(ctx workflow.Context, input *glacier.AbortMultipartUploadInput) *GlacierAbortMultipartUploadResult

    AbortVaultLock(ctx workflow.Context, input *glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error)
    AbortVaultLockAsync(ctx workflow.Context, input *glacier.AbortVaultLockInput) *GlacierAbortVaultLockResult

    AddTagsToVault(ctx workflow.Context, input *glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error)
    AddTagsToVaultAsync(ctx workflow.Context, input *glacier.AddTagsToVaultInput) *GlacierAddTagsToVaultResult

    CompleteMultipartUpload(ctx workflow.Context, input *glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error)
    CompleteMultipartUploadAsync(ctx workflow.Context, input *glacier.CompleteMultipartUploadInput) *GlacierCompleteMultipartUploadResult

    CompleteVaultLock(ctx workflow.Context, input *glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error)
    CompleteVaultLockAsync(ctx workflow.Context, input *glacier.CompleteVaultLockInput) *GlacierCompleteVaultLockResult

    CreateVault(ctx workflow.Context, input *glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error)
    CreateVaultAsync(ctx workflow.Context, input *glacier.CreateVaultInput) *GlacierCreateVaultResult

    DeleteArchive(ctx workflow.Context, input *glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error)
    DeleteArchiveAsync(ctx workflow.Context, input *glacier.DeleteArchiveInput) *GlacierDeleteArchiveResult

    DeleteVault(ctx workflow.Context, input *glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error)
    DeleteVaultAsync(ctx workflow.Context, input *glacier.DeleteVaultInput) *GlacierDeleteVaultResult

    DeleteVaultAccessPolicy(ctx workflow.Context, input *glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error)
    DeleteVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.DeleteVaultAccessPolicyInput) *GlacierDeleteVaultAccessPolicyResult

    DeleteVaultNotifications(ctx workflow.Context, input *glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error)
    DeleteVaultNotificationsAsync(ctx workflow.Context, input *glacier.DeleteVaultNotificationsInput) *GlacierDeleteVaultNotificationsResult

    DescribeJob(ctx workflow.Context, input *glacier.DescribeJobInput) (*glacier.JobDescription, error)
    DescribeJobAsync(ctx workflow.Context, input *glacier.DescribeJobInput) *GlacierDescribeJobResult

    DescribeVault(ctx workflow.Context, input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error)
    DescribeVaultAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) *GlacierDescribeVaultResult

    GetDataRetrievalPolicy(ctx workflow.Context, input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error)
    GetDataRetrievalPolicyAsync(ctx workflow.Context, input *glacier.GetDataRetrievalPolicyInput) *GlacierGetDataRetrievalPolicyResult

    GetJobOutput(ctx workflow.Context, input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error)
    GetJobOutputAsync(ctx workflow.Context, input *glacier.GetJobOutputInput) *GlacierGetJobOutputResult

    GetVaultAccessPolicy(ctx workflow.Context, input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error)
    GetVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.GetVaultAccessPolicyInput) *GlacierGetVaultAccessPolicyResult

    GetVaultLock(ctx workflow.Context, input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error)
    GetVaultLockAsync(ctx workflow.Context, input *glacier.GetVaultLockInput) *GlacierGetVaultLockResult

    GetVaultNotifications(ctx workflow.Context, input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error)
    GetVaultNotificationsAsync(ctx workflow.Context, input *glacier.GetVaultNotificationsInput) *GlacierGetVaultNotificationsResult

    InitiateJob(ctx workflow.Context, input *glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error)
    InitiateJobAsync(ctx workflow.Context, input *glacier.InitiateJobInput) *GlacierInitiateJobResult

    InitiateMultipartUpload(ctx workflow.Context, input *glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error)
    InitiateMultipartUploadAsync(ctx workflow.Context, input *glacier.InitiateMultipartUploadInput) *GlacierInitiateMultipartUploadResult

    InitiateVaultLock(ctx workflow.Context, input *glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error)
    InitiateVaultLockAsync(ctx workflow.Context, input *glacier.InitiateVaultLockInput) *GlacierInitiateVaultLockResult

    ListJobs(ctx workflow.Context, input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *glacier.ListJobsInput) *GlacierListJobsResult

    ListMultipartUploads(ctx workflow.Context, input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error)
    ListMultipartUploadsAsync(ctx workflow.Context, input *glacier.ListMultipartUploadsInput) *GlacierListMultipartUploadsResult

    ListParts(ctx workflow.Context, input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error)
    ListPartsAsync(ctx workflow.Context, input *glacier.ListPartsInput) *GlacierListPartsResult

    ListProvisionedCapacity(ctx workflow.Context, input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error)
    ListProvisionedCapacityAsync(ctx workflow.Context, input *glacier.ListProvisionedCapacityInput) *GlacierListProvisionedCapacityResult

    ListTagsForVault(ctx workflow.Context, input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error)
    ListTagsForVaultAsync(ctx workflow.Context, input *glacier.ListTagsForVaultInput) *GlacierListTagsForVaultResult

    ListVaults(ctx workflow.Context, input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error)
    ListVaultsAsync(ctx workflow.Context, input *glacier.ListVaultsInput) *GlacierListVaultsResult

    PurchaseProvisionedCapacity(ctx workflow.Context, input *glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error)
    PurchaseProvisionedCapacityAsync(ctx workflow.Context, input *glacier.PurchaseProvisionedCapacityInput) *GlacierPurchaseProvisionedCapacityResult

    RemoveTagsFromVault(ctx workflow.Context, input *glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error)
    RemoveTagsFromVaultAsync(ctx workflow.Context, input *glacier.RemoveTagsFromVaultInput) *GlacierRemoveTagsFromVaultResult

    SetDataRetrievalPolicy(ctx workflow.Context, input *glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error)
    SetDataRetrievalPolicyAsync(ctx workflow.Context, input *glacier.SetDataRetrievalPolicyInput) *GlacierSetDataRetrievalPolicyResult

    SetVaultAccessPolicy(ctx workflow.Context, input *glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error)
    SetVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.SetVaultAccessPolicyInput) *GlacierSetVaultAccessPolicyResult

    SetVaultNotifications(ctx workflow.Context, input *glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error)
    SetVaultNotificationsAsync(ctx workflow.Context, input *glacier.SetVaultNotificationsInput) *GlacierSetVaultNotificationsResult

    UploadArchive(ctx workflow.Context, input *glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error)
    UploadArchiveAsync(ctx workflow.Context, input *glacier.UploadArchiveInput) *GlacierUploadArchiveResult

    UploadMultipartPart(ctx workflow.Context, input *glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error)
    UploadMultipartPartAsync(ctx workflow.Context, input *glacier.UploadMultipartPartInput) *GlacierUploadMultipartPartResult

    WaitUntilVaultExists(ctx workflow.Context, input *glacier.DescribeVaultInput) error
    WaitUntilVaultNotExists(ctx workflow.Context, input *glacier.DescribeVaultInput) error}
type GlacierAbortMultipartUploadResult struct {
	Result workflow.Future
}

func (r *GlacierAbortMultipartUploadResult) Get(ctx workflow.Context) (*glacier.AbortMultipartUploadOutput, error) {
    var output glacier.AbortMultipartUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierAbortVaultLockResult struct {
	Result workflow.Future
}

func (r *GlacierAbortVaultLockResult) Get(ctx workflow.Context) (*glacier.AbortVaultLockOutput, error) {
    var output glacier.AbortVaultLockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierAddTagsToVaultResult struct {
	Result workflow.Future
}

func (r *GlacierAddTagsToVaultResult) Get(ctx workflow.Context) (*glacier.AddTagsToVaultOutput, error) {
    var output glacier.AddTagsToVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierCompleteMultipartUploadResult struct {
	Result workflow.Future
}

func (r *GlacierCompleteMultipartUploadResult) Get(ctx workflow.Context) (*glacier.ArchiveCreationOutput, error) {
    var output glacier.ArchiveCreationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierCompleteVaultLockResult struct {
	Result workflow.Future
}

func (r *GlacierCompleteVaultLockResult) Get(ctx workflow.Context) (*glacier.CompleteVaultLockOutput, error) {
    var output glacier.CompleteVaultLockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierCreateVaultResult struct {
	Result workflow.Future
}

func (r *GlacierCreateVaultResult) Get(ctx workflow.Context) (*glacier.CreateVaultOutput, error) {
    var output glacier.CreateVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierDeleteArchiveResult struct {
	Result workflow.Future
}

func (r *GlacierDeleteArchiveResult) Get(ctx workflow.Context) (*glacier.DeleteArchiveOutput, error) {
    var output glacier.DeleteArchiveOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierDeleteVaultResult struct {
	Result workflow.Future
}

func (r *GlacierDeleteVaultResult) Get(ctx workflow.Context) (*glacier.DeleteVaultOutput, error) {
    var output glacier.DeleteVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierDeleteVaultAccessPolicyResult struct {
	Result workflow.Future
}

func (r *GlacierDeleteVaultAccessPolicyResult) Get(ctx workflow.Context) (*glacier.DeleteVaultAccessPolicyOutput, error) {
    var output glacier.DeleteVaultAccessPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierDeleteVaultNotificationsResult struct {
	Result workflow.Future
}

func (r *GlacierDeleteVaultNotificationsResult) Get(ctx workflow.Context) (*glacier.DeleteVaultNotificationsOutput, error) {
    var output glacier.DeleteVaultNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierDescribeJobResult struct {
	Result workflow.Future
}

func (r *GlacierDescribeJobResult) Get(ctx workflow.Context) (*glacier.JobDescription, error) {
    var output glacier.JobDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierDescribeVaultResult struct {
	Result workflow.Future
}

func (r *GlacierDescribeVaultResult) Get(ctx workflow.Context) (*glacier.DescribeVaultOutput, error) {
    var output glacier.DescribeVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierGetDataRetrievalPolicyResult struct {
	Result workflow.Future
}

func (r *GlacierGetDataRetrievalPolicyResult) Get(ctx workflow.Context) (*glacier.GetDataRetrievalPolicyOutput, error) {
    var output glacier.GetDataRetrievalPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierGetJobOutputResult struct {
	Result workflow.Future
}

func (r *GlacierGetJobOutputResult) Get(ctx workflow.Context) (*glacier.GetJobOutputOutput, error) {
    var output glacier.GetJobOutputOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierGetVaultAccessPolicyResult struct {
	Result workflow.Future
}

func (r *GlacierGetVaultAccessPolicyResult) Get(ctx workflow.Context) (*glacier.GetVaultAccessPolicyOutput, error) {
    var output glacier.GetVaultAccessPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierGetVaultLockResult struct {
	Result workflow.Future
}

func (r *GlacierGetVaultLockResult) Get(ctx workflow.Context) (*glacier.GetVaultLockOutput, error) {
    var output glacier.GetVaultLockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierGetVaultNotificationsResult struct {
	Result workflow.Future
}

func (r *GlacierGetVaultNotificationsResult) Get(ctx workflow.Context) (*glacier.GetVaultNotificationsOutput, error) {
    var output glacier.GetVaultNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierInitiateJobResult struct {
	Result workflow.Future
}

func (r *GlacierInitiateJobResult) Get(ctx workflow.Context) (*glacier.InitiateJobOutput, error) {
    var output glacier.InitiateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierInitiateMultipartUploadResult struct {
	Result workflow.Future
}

func (r *GlacierInitiateMultipartUploadResult) Get(ctx workflow.Context) (*glacier.InitiateMultipartUploadOutput, error) {
    var output glacier.InitiateMultipartUploadOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierInitiateVaultLockResult struct {
	Result workflow.Future
}

func (r *GlacierInitiateVaultLockResult) Get(ctx workflow.Context) (*glacier.InitiateVaultLockOutput, error) {
    var output glacier.InitiateVaultLockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierListJobsResult struct {
	Result workflow.Future
}

func (r *GlacierListJobsResult) Get(ctx workflow.Context) (*glacier.ListJobsOutput, error) {
    var output glacier.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierListMultipartUploadsResult struct {
	Result workflow.Future
}

func (r *GlacierListMultipartUploadsResult) Get(ctx workflow.Context) (*glacier.ListMultipartUploadsOutput, error) {
    var output glacier.ListMultipartUploadsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierListPartsResult struct {
	Result workflow.Future
}

func (r *GlacierListPartsResult) Get(ctx workflow.Context) (*glacier.ListPartsOutput, error) {
    var output glacier.ListPartsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierListProvisionedCapacityResult struct {
	Result workflow.Future
}

func (r *GlacierListProvisionedCapacityResult) Get(ctx workflow.Context) (*glacier.ListProvisionedCapacityOutput, error) {
    var output glacier.ListProvisionedCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierListTagsForVaultResult struct {
	Result workflow.Future
}

func (r *GlacierListTagsForVaultResult) Get(ctx workflow.Context) (*glacier.ListTagsForVaultOutput, error) {
    var output glacier.ListTagsForVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierListVaultsResult struct {
	Result workflow.Future
}

func (r *GlacierListVaultsResult) Get(ctx workflow.Context) (*glacier.ListVaultsOutput, error) {
    var output glacier.ListVaultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierPurchaseProvisionedCapacityResult struct {
	Result workflow.Future
}

func (r *GlacierPurchaseProvisionedCapacityResult) Get(ctx workflow.Context) (*glacier.PurchaseProvisionedCapacityOutput, error) {
    var output glacier.PurchaseProvisionedCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierRemoveTagsFromVaultResult struct {
	Result workflow.Future
}

func (r *GlacierRemoveTagsFromVaultResult) Get(ctx workflow.Context) (*glacier.RemoveTagsFromVaultOutput, error) {
    var output glacier.RemoveTagsFromVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierSetDataRetrievalPolicyResult struct {
	Result workflow.Future
}

func (r *GlacierSetDataRetrievalPolicyResult) Get(ctx workflow.Context) (*glacier.SetDataRetrievalPolicyOutput, error) {
    var output glacier.SetDataRetrievalPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierSetVaultAccessPolicyResult struct {
	Result workflow.Future
}

func (r *GlacierSetVaultAccessPolicyResult) Get(ctx workflow.Context) (*glacier.SetVaultAccessPolicyOutput, error) {
    var output glacier.SetVaultAccessPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierSetVaultNotificationsResult struct {
	Result workflow.Future
}

func (r *GlacierSetVaultNotificationsResult) Get(ctx workflow.Context) (*glacier.SetVaultNotificationsOutput, error) {
    var output glacier.SetVaultNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierUploadArchiveResult struct {
	Result workflow.Future
}

func (r *GlacierUploadArchiveResult) Get(ctx workflow.Context) (*glacier.ArchiveCreationOutput, error) {
    var output glacier.ArchiveCreationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlacierUploadMultipartPartResult struct {
	Result workflow.Future
}

func (r *GlacierUploadMultipartPartResult) Get(ctx workflow.Context) (*glacier.UploadMultipartPartOutput, error) {
    var output glacier.UploadMultipartPartOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type GlacierStub struct {
    activities awsactivities.GlacierActivities
}

func NewGlacierStub() GlacierClient {
    return &GlacierStub{}
}
func (a *GlacierStub) AbortMultipartUpload(ctx workflow.Context, input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error) {
    var output glacier.AbortMultipartUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AbortMultipartUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) AbortMultipartUploadAsync(ctx workflow.Context, input *glacier.AbortMultipartUploadInput) *GlacierAbortMultipartUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AbortMultipartUpload, input)
    return &GlacierAbortMultipartUploadResult{Result: future}
}
func (a *GlacierStub) AbortVaultLock(ctx workflow.Context, input *glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error) {
    var output glacier.AbortVaultLockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AbortVaultLock, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) AbortVaultLockAsync(ctx workflow.Context, input *glacier.AbortVaultLockInput) *GlacierAbortVaultLockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AbortVaultLock, input)
    return &GlacierAbortVaultLockResult{Result: future}
}
func (a *GlacierStub) AddTagsToVault(ctx workflow.Context, input *glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error) {
    var output glacier.AddTagsToVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToVault, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) AddTagsToVaultAsync(ctx workflow.Context, input *glacier.AddTagsToVaultInput) *GlacierAddTagsToVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTagsToVault, input)
    return &GlacierAddTagsToVaultResult{Result: future}
}
func (a *GlacierStub) CompleteMultipartUpload(ctx workflow.Context, input *glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error) {
    var output glacier.ArchiveCreationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CompleteMultipartUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) CompleteMultipartUploadAsync(ctx workflow.Context, input *glacier.CompleteMultipartUploadInput) *GlacierCompleteMultipartUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CompleteMultipartUpload, input)
    return &GlacierCompleteMultipartUploadResult{Result: future}
}
func (a *GlacierStub) CompleteVaultLock(ctx workflow.Context, input *glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error) {
    var output glacier.CompleteVaultLockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CompleteVaultLock, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) CompleteVaultLockAsync(ctx workflow.Context, input *glacier.CompleteVaultLockInput) *GlacierCompleteVaultLockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CompleteVaultLock, input)
    return &GlacierCompleteVaultLockResult{Result: future}
}
func (a *GlacierStub) CreateVault(ctx workflow.Context, input *glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error) {
    var output glacier.CreateVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateVault, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) CreateVaultAsync(ctx workflow.Context, input *glacier.CreateVaultInput) *GlacierCreateVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateVault, input)
    return &GlacierCreateVaultResult{Result: future}
}
func (a *GlacierStub) DeleteArchive(ctx workflow.Context, input *glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error) {
    var output glacier.DeleteArchiveOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteArchive, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) DeleteArchiveAsync(ctx workflow.Context, input *glacier.DeleteArchiveInput) *GlacierDeleteArchiveResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteArchive, input)
    return &GlacierDeleteArchiveResult{Result: future}
}
func (a *GlacierStub) DeleteVault(ctx workflow.Context, input *glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error) {
    var output glacier.DeleteVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVault, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) DeleteVaultAsync(ctx workflow.Context, input *glacier.DeleteVaultInput) *GlacierDeleteVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVault, input)
    return &GlacierDeleteVaultResult{Result: future}
}
func (a *GlacierStub) DeleteVaultAccessPolicy(ctx workflow.Context, input *glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error) {
    var output glacier.DeleteVaultAccessPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVaultAccessPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) DeleteVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.DeleteVaultAccessPolicyInput) *GlacierDeleteVaultAccessPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVaultAccessPolicy, input)
    return &GlacierDeleteVaultAccessPolicyResult{Result: future}
}
func (a *GlacierStub) DeleteVaultNotifications(ctx workflow.Context, input *glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error) {
    var output glacier.DeleteVaultNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVaultNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) DeleteVaultNotificationsAsync(ctx workflow.Context, input *glacier.DeleteVaultNotificationsInput) *GlacierDeleteVaultNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVaultNotifications, input)
    return &GlacierDeleteVaultNotificationsResult{Result: future}
}
func (a *GlacierStub) DescribeJob(ctx workflow.Context, input *glacier.DescribeJobInput) (*glacier.JobDescription, error) {
    var output glacier.JobDescription
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) DescribeJobAsync(ctx workflow.Context, input *glacier.DescribeJobInput) *GlacierDescribeJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeJob, input)
    return &GlacierDescribeJobResult{Result: future}
}
func (a *GlacierStub) DescribeVault(ctx workflow.Context, input *glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error) {
    var output glacier.DescribeVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVault, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) DescribeVaultAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) *GlacierDescribeVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVault, input)
    return &GlacierDescribeVaultResult{Result: future}
}
func (a *GlacierStub) GetDataRetrievalPolicy(ctx workflow.Context, input *glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error) {
    var output glacier.GetDataRetrievalPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDataRetrievalPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) GetDataRetrievalPolicyAsync(ctx workflow.Context, input *glacier.GetDataRetrievalPolicyInput) *GlacierGetDataRetrievalPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDataRetrievalPolicy, input)
    return &GlacierGetDataRetrievalPolicyResult{Result: future}
}
func (a *GlacierStub) GetJobOutput(ctx workflow.Context, input *glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error) {
    var output glacier.GetJobOutputOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobOutput, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) GetJobOutputAsync(ctx workflow.Context, input *glacier.GetJobOutputInput) *GlacierGetJobOutputResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobOutput, input)
    return &GlacierGetJobOutputResult{Result: future}
}
func (a *GlacierStub) GetVaultAccessPolicy(ctx workflow.Context, input *glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error) {
    var output glacier.GetVaultAccessPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVaultAccessPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) GetVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.GetVaultAccessPolicyInput) *GlacierGetVaultAccessPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVaultAccessPolicy, input)
    return &GlacierGetVaultAccessPolicyResult{Result: future}
}
func (a *GlacierStub) GetVaultLock(ctx workflow.Context, input *glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error) {
    var output glacier.GetVaultLockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVaultLock, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) GetVaultLockAsync(ctx workflow.Context, input *glacier.GetVaultLockInput) *GlacierGetVaultLockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVaultLock, input)
    return &GlacierGetVaultLockResult{Result: future}
}
func (a *GlacierStub) GetVaultNotifications(ctx workflow.Context, input *glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error) {
    var output glacier.GetVaultNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetVaultNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) GetVaultNotificationsAsync(ctx workflow.Context, input *glacier.GetVaultNotificationsInput) *GlacierGetVaultNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetVaultNotifications, input)
    return &GlacierGetVaultNotificationsResult{Result: future}
}
func (a *GlacierStub) InitiateJob(ctx workflow.Context, input *glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error) {
    var output glacier.InitiateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitiateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) InitiateJobAsync(ctx workflow.Context, input *glacier.InitiateJobInput) *GlacierInitiateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InitiateJob, input)
    return &GlacierInitiateJobResult{Result: future}
}
func (a *GlacierStub) InitiateMultipartUpload(ctx workflow.Context, input *glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error) {
    var output glacier.InitiateMultipartUploadOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitiateMultipartUpload, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) InitiateMultipartUploadAsync(ctx workflow.Context, input *glacier.InitiateMultipartUploadInput) *GlacierInitiateMultipartUploadResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InitiateMultipartUpload, input)
    return &GlacierInitiateMultipartUploadResult{Result: future}
}
func (a *GlacierStub) InitiateVaultLock(ctx workflow.Context, input *glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error) {
    var output glacier.InitiateVaultLockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InitiateVaultLock, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) InitiateVaultLockAsync(ctx workflow.Context, input *glacier.InitiateVaultLockInput) *GlacierInitiateVaultLockResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InitiateVaultLock, input)
    return &GlacierInitiateVaultLockResult{Result: future}
}
func (a *GlacierStub) ListJobs(ctx workflow.Context, input *glacier.ListJobsInput) (*glacier.ListJobsOutput, error) {
    var output glacier.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) ListJobsAsync(ctx workflow.Context, input *glacier.ListJobsInput) *GlacierListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &GlacierListJobsResult{Result: future}
}
func (a *GlacierStub) ListMultipartUploads(ctx workflow.Context, input *glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error) {
    var output glacier.ListMultipartUploadsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMultipartUploads, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) ListMultipartUploadsAsync(ctx workflow.Context, input *glacier.ListMultipartUploadsInput) *GlacierListMultipartUploadsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMultipartUploads, input)
    return &GlacierListMultipartUploadsResult{Result: future}
}
func (a *GlacierStub) ListParts(ctx workflow.Context, input *glacier.ListPartsInput) (*glacier.ListPartsOutput, error) {
    var output glacier.ListPartsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListParts, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) ListPartsAsync(ctx workflow.Context, input *glacier.ListPartsInput) *GlacierListPartsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListParts, input)
    return &GlacierListPartsResult{Result: future}
}
func (a *GlacierStub) ListProvisionedCapacity(ctx workflow.Context, input *glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error) {
    var output glacier.ListProvisionedCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisionedCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) ListProvisionedCapacityAsync(ctx workflow.Context, input *glacier.ListProvisionedCapacityInput) *GlacierListProvisionedCapacityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisionedCapacity, input)
    return &GlacierListProvisionedCapacityResult{Result: future}
}
func (a *GlacierStub) ListTagsForVault(ctx workflow.Context, input *glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error) {
    var output glacier.ListTagsForVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForVault, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) ListTagsForVaultAsync(ctx workflow.Context, input *glacier.ListTagsForVaultInput) *GlacierListTagsForVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForVault, input)
    return &GlacierListTagsForVaultResult{Result: future}
}
func (a *GlacierStub) ListVaults(ctx workflow.Context, input *glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error) {
    var output glacier.ListVaultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVaults, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) ListVaultsAsync(ctx workflow.Context, input *glacier.ListVaultsInput) *GlacierListVaultsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVaults, input)
    return &GlacierListVaultsResult{Result: future}
}
func (a *GlacierStub) PurchaseProvisionedCapacity(ctx workflow.Context, input *glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error) {
    var output glacier.PurchaseProvisionedCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseProvisionedCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) PurchaseProvisionedCapacityAsync(ctx workflow.Context, input *glacier.PurchaseProvisionedCapacityInput) *GlacierPurchaseProvisionedCapacityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurchaseProvisionedCapacity, input)
    return &GlacierPurchaseProvisionedCapacityResult{Result: future}
}
func (a *GlacierStub) RemoveTagsFromVault(ctx workflow.Context, input *glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error) {
    var output glacier.RemoveTagsFromVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromVault, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) RemoveTagsFromVaultAsync(ctx workflow.Context, input *glacier.RemoveTagsFromVaultInput) *GlacierRemoveTagsFromVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromVault, input)
    return &GlacierRemoveTagsFromVaultResult{Result: future}
}
func (a *GlacierStub) SetDataRetrievalPolicy(ctx workflow.Context, input *glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error) {
    var output glacier.SetDataRetrievalPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetDataRetrievalPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) SetDataRetrievalPolicyAsync(ctx workflow.Context, input *glacier.SetDataRetrievalPolicyInput) *GlacierSetDataRetrievalPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetDataRetrievalPolicy, input)
    return &GlacierSetDataRetrievalPolicyResult{Result: future}
}
func (a *GlacierStub) SetVaultAccessPolicy(ctx workflow.Context, input *glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error) {
    var output glacier.SetVaultAccessPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetVaultAccessPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) SetVaultAccessPolicyAsync(ctx workflow.Context, input *glacier.SetVaultAccessPolicyInput) *GlacierSetVaultAccessPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetVaultAccessPolicy, input)
    return &GlacierSetVaultAccessPolicyResult{Result: future}
}
func (a *GlacierStub) SetVaultNotifications(ctx workflow.Context, input *glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error) {
    var output glacier.SetVaultNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetVaultNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) SetVaultNotificationsAsync(ctx workflow.Context, input *glacier.SetVaultNotificationsInput) *GlacierSetVaultNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetVaultNotifications, input)
    return &GlacierSetVaultNotificationsResult{Result: future}
}
func (a *GlacierStub) UploadArchive(ctx workflow.Context, input *glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error) {
    var output glacier.ArchiveCreationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadArchive, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) UploadArchiveAsync(ctx workflow.Context, input *glacier.UploadArchiveInput) *GlacierUploadArchiveResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UploadArchive, input)
    return &GlacierUploadArchiveResult{Result: future}
}
func (a *GlacierStub) UploadMultipartPart(ctx workflow.Context, input *glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error) {
    var output glacier.UploadMultipartPartOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UploadMultipartPart, input).Get(ctx, &output)
    return &output, err
}

func (a *GlacierStub) UploadMultipartPartAsync(ctx workflow.Context, input *glacier.UploadMultipartPartInput) *GlacierUploadMultipartPartResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UploadMultipartPart, input)
    return &GlacierUploadMultipartPartResult{Result: future}
}

func (a *GlacierStub) WaitUntilVaultExists(ctx workflow.Context, input *glacier.DescribeVaultInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVaultExists, input).Get(ctx, nil)
}

func (a *GlacierStub) WaitUntilVaultExistsAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVaultExists, input)
}

func (a *GlacierStub) WaitUntilVaultNotExists(ctx workflow.Context, input *glacier.DescribeVaultInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVaultNotExists, input).Get(ctx, nil)
}

func (a *GlacierStub) WaitUntilVaultNotExistsAsync(ctx workflow.Context, input *glacier.DescribeVaultInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilVaultNotExists, input)
}
