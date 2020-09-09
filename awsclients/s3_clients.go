package awsclients

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type S3Client interface {
	AbortMultipartUpload(ctx workflow.Context, input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error)
	AbortMultipartUploadAsync(ctx workflow.Context, input *s3.AbortMultipartUploadInput) *S3AbortMultipartUploadResult

	CompleteMultipartUpload(ctx workflow.Context, input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadAsync(ctx workflow.Context, input *s3.CompleteMultipartUploadInput) *S3CompleteMultipartUploadResult

	CopyObject(ctx workflow.Context, input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error)
	CopyObjectAsync(ctx workflow.Context, input *s3.CopyObjectInput) *S3CopyObjectResult

	CreateBucket(ctx workflow.Context, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error)
	CreateBucketAsync(ctx workflow.Context, input *s3.CreateBucketInput) *S3CreateBucketResult

	CreateMultipartUpload(ctx workflow.Context, input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error)
	CreateMultipartUploadAsync(ctx workflow.Context, input *s3.CreateMultipartUploadInput) *S3CreateMultipartUploadResult

	DeleteBucket(ctx workflow.Context, input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error)
	DeleteBucketAsync(ctx workflow.Context, input *s3.DeleteBucketInput) *S3DeleteBucketResult

	DeleteBucketAnalyticsConfiguration(ctx workflow.Context, input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error)
	DeleteBucketAnalyticsConfigurationAsync(ctx workflow.Context, input *s3.DeleteBucketAnalyticsConfigurationInput) *S3DeleteBucketAnalyticsConfigurationResult

	DeleteBucketCors(ctx workflow.Context, input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error)
	DeleteBucketCorsAsync(ctx workflow.Context, input *s3.DeleteBucketCorsInput) *S3DeleteBucketCorsResult

	DeleteBucketEncryption(ctx workflow.Context, input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error)
	DeleteBucketEncryptionAsync(ctx workflow.Context, input *s3.DeleteBucketEncryptionInput) *S3DeleteBucketEncryptionResult

	DeleteBucketInventoryConfiguration(ctx workflow.Context, input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error)
	DeleteBucketInventoryConfigurationAsync(ctx workflow.Context, input *s3.DeleteBucketInventoryConfigurationInput) *S3DeleteBucketInventoryConfigurationResult

	DeleteBucketLifecycle(ctx workflow.Context, input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleAsync(ctx workflow.Context, input *s3.DeleteBucketLifecycleInput) *S3DeleteBucketLifecycleResult

	DeleteBucketMetricsConfiguration(ctx workflow.Context, input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error)
	DeleteBucketMetricsConfigurationAsync(ctx workflow.Context, input *s3.DeleteBucketMetricsConfigurationInput) *S3DeleteBucketMetricsConfigurationResult

	DeleteBucketPolicy(ctx workflow.Context, input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error)
	DeleteBucketPolicyAsync(ctx workflow.Context, input *s3.DeleteBucketPolicyInput) *S3DeleteBucketPolicyResult

	DeleteBucketReplication(ctx workflow.Context, input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error)
	DeleteBucketReplicationAsync(ctx workflow.Context, input *s3.DeleteBucketReplicationInput) *S3DeleteBucketReplicationResult

	DeleteBucketTagging(ctx workflow.Context, input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error)
	DeleteBucketTaggingAsync(ctx workflow.Context, input *s3.DeleteBucketTaggingInput) *S3DeleteBucketTaggingResult

	DeleteBucketWebsite(ctx workflow.Context, input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteAsync(ctx workflow.Context, input *s3.DeleteBucketWebsiteInput) *S3DeleteBucketWebsiteResult

	DeleteObject(ctx workflow.Context, input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	DeleteObjectAsync(ctx workflow.Context, input *s3.DeleteObjectInput) *S3DeleteObjectResult

	DeleteObjectTagging(ctx workflow.Context, input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingAsync(ctx workflow.Context, input *s3.DeleteObjectTaggingInput) *S3DeleteObjectTaggingResult

	DeleteObjects(ctx workflow.Context, input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)
	DeleteObjectsAsync(ctx workflow.Context, input *s3.DeleteObjectsInput) *S3DeleteObjectsResult

	DeletePublicAccessBlock(ctx workflow.Context, input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error)
	DeletePublicAccessBlockAsync(ctx workflow.Context, input *s3.DeletePublicAccessBlockInput) *S3DeletePublicAccessBlockResult

	GetBucketAccelerateConfiguration(ctx workflow.Context, input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAccelerateConfigurationAsync(ctx workflow.Context, input *s3.GetBucketAccelerateConfigurationInput) *S3GetBucketAccelerateConfigurationResult

	GetBucketAcl(ctx workflow.Context, input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error)
	GetBucketAclAsync(ctx workflow.Context, input *s3.GetBucketAclInput) *S3GetBucketAclResult

	GetBucketAnalyticsConfiguration(ctx workflow.Context, input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketAnalyticsConfigurationAsync(ctx workflow.Context, input *s3.GetBucketAnalyticsConfigurationInput) *S3GetBucketAnalyticsConfigurationResult

	GetBucketCors(ctx workflow.Context, input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error)
	GetBucketCorsAsync(ctx workflow.Context, input *s3.GetBucketCorsInput) *S3GetBucketCorsResult

	GetBucketEncryption(ctx workflow.Context, input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error)
	GetBucketEncryptionAsync(ctx workflow.Context, input *s3.GetBucketEncryptionInput) *S3GetBucketEncryptionResult

	GetBucketInventoryConfiguration(ctx workflow.Context, input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error)
	GetBucketInventoryConfigurationAsync(ctx workflow.Context, input *s3.GetBucketInventoryConfigurationInput) *S3GetBucketInventoryConfigurationResult

	GetBucketLifecycle(ctx workflow.Context, input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error)
	GetBucketLifecycleAsync(ctx workflow.Context, input *s3.GetBucketLifecycleInput) *S3GetBucketLifecycleResult

	GetBucketLifecycleConfiguration(ctx workflow.Context, input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3.GetBucketLifecycleConfigurationInput) *S3GetBucketLifecycleConfigurationResult

	GetBucketLocation(ctx workflow.Context, input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error)
	GetBucketLocationAsync(ctx workflow.Context, input *s3.GetBucketLocationInput) *S3GetBucketLocationResult

	GetBucketLogging(ctx workflow.Context, input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error)
	GetBucketLoggingAsync(ctx workflow.Context, input *s3.GetBucketLoggingInput) *S3GetBucketLoggingResult

	GetBucketMetricsConfiguration(ctx workflow.Context, input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error)
	GetBucketMetricsConfigurationAsync(ctx workflow.Context, input *s3.GetBucketMetricsConfigurationInput) *S3GetBucketMetricsConfigurationResult

	GetBucketNotification(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error)
	GetBucketNotificationAsync(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) *S3GetBucketNotificationResult

	GetBucketNotificationConfiguration(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error)
	GetBucketNotificationConfigurationAsync(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) *S3GetBucketNotificationConfigurationResult

	GetBucketPolicy(ctx workflow.Context, input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error)
	GetBucketPolicyAsync(ctx workflow.Context, input *s3.GetBucketPolicyInput) *S3GetBucketPolicyResult

	GetBucketPolicyStatus(ctx workflow.Context, input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error)
	GetBucketPolicyStatusAsync(ctx workflow.Context, input *s3.GetBucketPolicyStatusInput) *S3GetBucketPolicyStatusResult

	GetBucketReplication(ctx workflow.Context, input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error)
	GetBucketReplicationAsync(ctx workflow.Context, input *s3.GetBucketReplicationInput) *S3GetBucketReplicationResult

	GetBucketRequestPayment(ctx workflow.Context, input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error)
	GetBucketRequestPaymentAsync(ctx workflow.Context, input *s3.GetBucketRequestPaymentInput) *S3GetBucketRequestPaymentResult

	GetBucketTagging(ctx workflow.Context, input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error)
	GetBucketTaggingAsync(ctx workflow.Context, input *s3.GetBucketTaggingInput) *S3GetBucketTaggingResult

	GetBucketVersioning(ctx workflow.Context, input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error)
	GetBucketVersioningAsync(ctx workflow.Context, input *s3.GetBucketVersioningInput) *S3GetBucketVersioningResult

	GetBucketWebsite(ctx workflow.Context, input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error)
	GetBucketWebsiteAsync(ctx workflow.Context, input *s3.GetBucketWebsiteInput) *S3GetBucketWebsiteResult

	GetObject(ctx workflow.Context, input *s3.GetObjectInput) (*s3.GetObjectOutput, error)
	GetObjectAsync(ctx workflow.Context, input *s3.GetObjectInput) *S3GetObjectResult

	GetObjectAcl(ctx workflow.Context, input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error)
	GetObjectAclAsync(ctx workflow.Context, input *s3.GetObjectAclInput) *S3GetObjectAclResult

	GetObjectLegalHold(ctx workflow.Context, input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error)
	GetObjectLegalHoldAsync(ctx workflow.Context, input *s3.GetObjectLegalHoldInput) *S3GetObjectLegalHoldResult

	GetObjectLockConfiguration(ctx workflow.Context, input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error)
	GetObjectLockConfigurationAsync(ctx workflow.Context, input *s3.GetObjectLockConfigurationInput) *S3GetObjectLockConfigurationResult

	GetObjectRetention(ctx workflow.Context, input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error)
	GetObjectRetentionAsync(ctx workflow.Context, input *s3.GetObjectRetentionInput) *S3GetObjectRetentionResult

	GetObjectTagging(ctx workflow.Context, input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error)
	GetObjectTaggingAsync(ctx workflow.Context, input *s3.GetObjectTaggingInput) *S3GetObjectTaggingResult

	GetObjectTorrent(ctx workflow.Context, input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error)
	GetObjectTorrentAsync(ctx workflow.Context, input *s3.GetObjectTorrentInput) *S3GetObjectTorrentResult

	GetPublicAccessBlock(ctx workflow.Context, input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error)
	GetPublicAccessBlockAsync(ctx workflow.Context, input *s3.GetPublicAccessBlockInput) *S3GetPublicAccessBlockResult

	HeadBucket(ctx workflow.Context, input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error)
	HeadBucketAsync(ctx workflow.Context, input *s3.HeadBucketInput) *S3HeadBucketResult

	HeadObject(ctx workflow.Context, input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error)
	HeadObjectAsync(ctx workflow.Context, input *s3.HeadObjectInput) *S3HeadObjectResult

	ListBucketAnalyticsConfigurations(ctx workflow.Context, input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketAnalyticsConfigurationsAsync(ctx workflow.Context, input *s3.ListBucketAnalyticsConfigurationsInput) *S3ListBucketAnalyticsConfigurationsResult

	ListBucketInventoryConfigurations(ctx workflow.Context, input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error)
	ListBucketInventoryConfigurationsAsync(ctx workflow.Context, input *s3.ListBucketInventoryConfigurationsInput) *S3ListBucketInventoryConfigurationsResult

	ListBucketMetricsConfigurations(ctx workflow.Context, input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error)
	ListBucketMetricsConfigurationsAsync(ctx workflow.Context, input *s3.ListBucketMetricsConfigurationsInput) *S3ListBucketMetricsConfigurationsResult

	ListBuckets(ctx workflow.Context, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error)
	ListBucketsAsync(ctx workflow.Context, input *s3.ListBucketsInput) *S3ListBucketsResult

	ListMultipartUploads(ctx workflow.Context, input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error)
	ListMultipartUploadsAsync(ctx workflow.Context, input *s3.ListMultipartUploadsInput) *S3ListMultipartUploadsResult

	ListObjectVersions(ctx workflow.Context, input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error)
	ListObjectVersionsAsync(ctx workflow.Context, input *s3.ListObjectVersionsInput) *S3ListObjectVersionsResult

	ListObjects(ctx workflow.Context, input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
	ListObjectsAsync(ctx workflow.Context, input *s3.ListObjectsInput) *S3ListObjectsResult

	ListObjectsV2(ctx workflow.Context, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	ListObjectsV2Async(ctx workflow.Context, input *s3.ListObjectsV2Input) *S3ListObjectsV2Result

	ListParts(ctx workflow.Context, input *s3.ListPartsInput) (*s3.ListPartsOutput, error)
	ListPartsAsync(ctx workflow.Context, input *s3.ListPartsInput) *S3ListPartsResult

	PutBucketAccelerateConfiguration(ctx workflow.Context, input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error)
	PutBucketAccelerateConfigurationAsync(ctx workflow.Context, input *s3.PutBucketAccelerateConfigurationInput) *S3PutBucketAccelerateConfigurationResult

	PutBucketAcl(ctx workflow.Context, input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error)
	PutBucketAclAsync(ctx workflow.Context, input *s3.PutBucketAclInput) *S3PutBucketAclResult

	PutBucketAnalyticsConfiguration(ctx workflow.Context, input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error)
	PutBucketAnalyticsConfigurationAsync(ctx workflow.Context, input *s3.PutBucketAnalyticsConfigurationInput) *S3PutBucketAnalyticsConfigurationResult

	PutBucketCors(ctx workflow.Context, input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error)
	PutBucketCorsAsync(ctx workflow.Context, input *s3.PutBucketCorsInput) *S3PutBucketCorsResult

	PutBucketEncryption(ctx workflow.Context, input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error)
	PutBucketEncryptionAsync(ctx workflow.Context, input *s3.PutBucketEncryptionInput) *S3PutBucketEncryptionResult

	PutBucketInventoryConfiguration(ctx workflow.Context, input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error)
	PutBucketInventoryConfigurationAsync(ctx workflow.Context, input *s3.PutBucketInventoryConfigurationInput) *S3PutBucketInventoryConfigurationResult

	PutBucketLifecycle(ctx workflow.Context, input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error)
	PutBucketLifecycleAsync(ctx workflow.Context, input *s3.PutBucketLifecycleInput) *S3PutBucketLifecycleResult

	PutBucketLifecycleConfiguration(ctx workflow.Context, input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3.PutBucketLifecycleConfigurationInput) *S3PutBucketLifecycleConfigurationResult

	PutBucketLogging(ctx workflow.Context, input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error)
	PutBucketLoggingAsync(ctx workflow.Context, input *s3.PutBucketLoggingInput) *S3PutBucketLoggingResult

	PutBucketMetricsConfiguration(ctx workflow.Context, input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error)
	PutBucketMetricsConfigurationAsync(ctx workflow.Context, input *s3.PutBucketMetricsConfigurationInput) *S3PutBucketMetricsConfigurationResult

	PutBucketNotification(ctx workflow.Context, input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error)
	PutBucketNotificationAsync(ctx workflow.Context, input *s3.PutBucketNotificationInput) *S3PutBucketNotificationResult

	PutBucketNotificationConfiguration(ctx workflow.Context, input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error)
	PutBucketNotificationConfigurationAsync(ctx workflow.Context, input *s3.PutBucketNotificationConfigurationInput) *S3PutBucketNotificationConfigurationResult

	PutBucketPolicy(ctx workflow.Context, input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error)
	PutBucketPolicyAsync(ctx workflow.Context, input *s3.PutBucketPolicyInput) *S3PutBucketPolicyResult

	PutBucketReplication(ctx workflow.Context, input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error)
	PutBucketReplicationAsync(ctx workflow.Context, input *s3.PutBucketReplicationInput) *S3PutBucketReplicationResult

	PutBucketRequestPayment(ctx workflow.Context, input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error)
	PutBucketRequestPaymentAsync(ctx workflow.Context, input *s3.PutBucketRequestPaymentInput) *S3PutBucketRequestPaymentResult

	PutBucketTagging(ctx workflow.Context, input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error)
	PutBucketTaggingAsync(ctx workflow.Context, input *s3.PutBucketTaggingInput) *S3PutBucketTaggingResult

	PutBucketVersioning(ctx workflow.Context, input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error)
	PutBucketVersioningAsync(ctx workflow.Context, input *s3.PutBucketVersioningInput) *S3PutBucketVersioningResult

	PutBucketWebsite(ctx workflow.Context, input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error)
	PutBucketWebsiteAsync(ctx workflow.Context, input *s3.PutBucketWebsiteInput) *S3PutBucketWebsiteResult

	PutObject(ctx workflow.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	PutObjectAsync(ctx workflow.Context, input *s3.PutObjectInput) *S3PutObjectResult

	PutObjectAcl(ctx workflow.Context, input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error)
	PutObjectAclAsync(ctx workflow.Context, input *s3.PutObjectAclInput) *S3PutObjectAclResult

	PutObjectLegalHold(ctx workflow.Context, input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error)
	PutObjectLegalHoldAsync(ctx workflow.Context, input *s3.PutObjectLegalHoldInput) *S3PutObjectLegalHoldResult

	PutObjectLockConfiguration(ctx workflow.Context, input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error)
	PutObjectLockConfigurationAsync(ctx workflow.Context, input *s3.PutObjectLockConfigurationInput) *S3PutObjectLockConfigurationResult

	PutObjectRetention(ctx workflow.Context, input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error)
	PutObjectRetentionAsync(ctx workflow.Context, input *s3.PutObjectRetentionInput) *S3PutObjectRetentionResult

	PutObjectTagging(ctx workflow.Context, input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error)
	PutObjectTaggingAsync(ctx workflow.Context, input *s3.PutObjectTaggingInput) *S3PutObjectTaggingResult

	PutPublicAccessBlock(ctx workflow.Context, input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error)
	PutPublicAccessBlockAsync(ctx workflow.Context, input *s3.PutPublicAccessBlockInput) *S3PutPublicAccessBlockResult

	RestoreObject(ctx workflow.Context, input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error)
	RestoreObjectAsync(ctx workflow.Context, input *s3.RestoreObjectInput) *S3RestoreObjectResult

	SelectObjectContent(ctx workflow.Context, input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error)
	SelectObjectContentAsync(ctx workflow.Context, input *s3.SelectObjectContentInput) *S3SelectObjectContentResult

	UploadPart(ctx workflow.Context, input *s3.UploadPartInput) (*s3.UploadPartOutput, error)
	UploadPartAsync(ctx workflow.Context, input *s3.UploadPartInput) *S3UploadPartResult

	UploadPartCopy(ctx workflow.Context, input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error)
	UploadPartCopyAsync(ctx workflow.Context, input *s3.UploadPartCopyInput) *S3UploadPartCopyResult

	WaitUntilBucketExists(ctx workflow.Context, input *s3.HeadBucketInput) error
	WaitUntilBucketNotExists(ctx workflow.Context, input *s3.HeadBucketInput) error
	WaitUntilObjectExists(ctx workflow.Context, input *s3.HeadObjectInput) error
	WaitUntilObjectNotExists(ctx workflow.Context, input *s3.HeadObjectInput) error
}

type S3AbortMultipartUploadResult struct {
	Result workflow.Future
}

func (r *S3AbortMultipartUploadResult) Get(ctx workflow.Context) (*s3.AbortMultipartUploadOutput, error) {
	var output s3.AbortMultipartUploadOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3CompleteMultipartUploadResult struct {
	Result workflow.Future
}

func (r *S3CompleteMultipartUploadResult) Get(ctx workflow.Context) (*s3.CompleteMultipartUploadOutput, error) {
	var output s3.CompleteMultipartUploadOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3CopyObjectResult struct {
	Result workflow.Future
}

func (r *S3CopyObjectResult) Get(ctx workflow.Context) (*s3.CopyObjectOutput, error) {
	var output s3.CopyObjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3CreateBucketResult struct {
	Result workflow.Future
}

func (r *S3CreateBucketResult) Get(ctx workflow.Context) (*s3.CreateBucketOutput, error) {
	var output s3.CreateBucketOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3CreateMultipartUploadResult struct {
	Result workflow.Future
}

func (r *S3CreateMultipartUploadResult) Get(ctx workflow.Context) (*s3.CreateMultipartUploadOutput, error) {
	var output s3.CreateMultipartUploadOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketResult) Get(ctx workflow.Context) (*s3.DeleteBucketOutput, error) {
	var output s3.DeleteBucketOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketAnalyticsConfigurationResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketAnalyticsConfigurationResult) Get(ctx workflow.Context) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	var output s3.DeleteBucketAnalyticsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketCorsResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketCorsResult) Get(ctx workflow.Context) (*s3.DeleteBucketCorsOutput, error) {
	var output s3.DeleteBucketCorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketEncryptionResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketEncryptionResult) Get(ctx workflow.Context) (*s3.DeleteBucketEncryptionOutput, error) {
	var output s3.DeleteBucketEncryptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketInventoryConfigurationResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketInventoryConfigurationResult) Get(ctx workflow.Context) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	var output s3.DeleteBucketInventoryConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketLifecycleResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketLifecycleResult) Get(ctx workflow.Context) (*s3.DeleteBucketLifecycleOutput, error) {
	var output s3.DeleteBucketLifecycleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketMetricsConfigurationResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketMetricsConfigurationResult) Get(ctx workflow.Context) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	var output s3.DeleteBucketMetricsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketPolicyResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketPolicyResult) Get(ctx workflow.Context) (*s3.DeleteBucketPolicyOutput, error) {
	var output s3.DeleteBucketPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketReplicationResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketReplicationResult) Get(ctx workflow.Context) (*s3.DeleteBucketReplicationOutput, error) {
	var output s3.DeleteBucketReplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketTaggingResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketTaggingResult) Get(ctx workflow.Context) (*s3.DeleteBucketTaggingOutput, error) {
	var output s3.DeleteBucketTaggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteBucketWebsiteResult struct {
	Result workflow.Future
}

func (r *S3DeleteBucketWebsiteResult) Get(ctx workflow.Context) (*s3.DeleteBucketWebsiteOutput, error) {
	var output s3.DeleteBucketWebsiteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteObjectResult struct {
	Result workflow.Future
}

func (r *S3DeleteObjectResult) Get(ctx workflow.Context) (*s3.DeleteObjectOutput, error) {
	var output s3.DeleteObjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteObjectTaggingResult struct {
	Result workflow.Future
}

func (r *S3DeleteObjectTaggingResult) Get(ctx workflow.Context) (*s3.DeleteObjectTaggingOutput, error) {
	var output s3.DeleteObjectTaggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeleteObjectsResult struct {
	Result workflow.Future
}

func (r *S3DeleteObjectsResult) Get(ctx workflow.Context) (*s3.DeleteObjectsOutput, error) {
	var output s3.DeleteObjectsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3DeletePublicAccessBlockResult struct {
	Result workflow.Future
}

func (r *S3DeletePublicAccessBlockResult) Get(ctx workflow.Context) (*s3.DeletePublicAccessBlockOutput, error) {
	var output s3.DeletePublicAccessBlockOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketAccelerateConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketAccelerateConfigurationResult) Get(ctx workflow.Context) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	var output s3.GetBucketAccelerateConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketAclResult struct {
	Result workflow.Future
}

func (r *S3GetBucketAclResult) Get(ctx workflow.Context) (*s3.GetBucketAclOutput, error) {
	var output s3.GetBucketAclOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketAnalyticsConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketAnalyticsConfigurationResult) Get(ctx workflow.Context) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	var output s3.GetBucketAnalyticsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketCorsResult struct {
	Result workflow.Future
}

func (r *S3GetBucketCorsResult) Get(ctx workflow.Context) (*s3.GetBucketCorsOutput, error) {
	var output s3.GetBucketCorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketEncryptionResult struct {
	Result workflow.Future
}

func (r *S3GetBucketEncryptionResult) Get(ctx workflow.Context) (*s3.GetBucketEncryptionOutput, error) {
	var output s3.GetBucketEncryptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketInventoryConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketInventoryConfigurationResult) Get(ctx workflow.Context) (*s3.GetBucketInventoryConfigurationOutput, error) {
	var output s3.GetBucketInventoryConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketLifecycleResult struct {
	Result workflow.Future
}

func (r *S3GetBucketLifecycleResult) Get(ctx workflow.Context) (*s3.GetBucketLifecycleOutput, error) {
	var output s3.GetBucketLifecycleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketLifecycleConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketLifecycleConfigurationResult) Get(ctx workflow.Context) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	var output s3.GetBucketLifecycleConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketLocationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketLocationResult) Get(ctx workflow.Context) (*s3.GetBucketLocationOutput, error) {
	var output s3.GetBucketLocationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketLoggingResult struct {
	Result workflow.Future
}

func (r *S3GetBucketLoggingResult) Get(ctx workflow.Context) (*s3.GetBucketLoggingOutput, error) {
	var output s3.GetBucketLoggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketMetricsConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketMetricsConfigurationResult) Get(ctx workflow.Context) (*s3.GetBucketMetricsConfigurationOutput, error) {
	var output s3.GetBucketMetricsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketNotificationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketNotificationResult) Get(ctx workflow.Context) (*s3.NotificationConfigurationDeprecated, error) {
	var output s3.NotificationConfigurationDeprecated
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketNotificationConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketNotificationConfigurationResult) Get(ctx workflow.Context) (*s3.NotificationConfiguration, error) {
	var output s3.NotificationConfiguration
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketPolicyResult struct {
	Result workflow.Future
}

func (r *S3GetBucketPolicyResult) Get(ctx workflow.Context) (*s3.GetBucketPolicyOutput, error) {
	var output s3.GetBucketPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketPolicyStatusResult struct {
	Result workflow.Future
}

func (r *S3GetBucketPolicyStatusResult) Get(ctx workflow.Context) (*s3.GetBucketPolicyStatusOutput, error) {
	var output s3.GetBucketPolicyStatusOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketReplicationResult struct {
	Result workflow.Future
}

func (r *S3GetBucketReplicationResult) Get(ctx workflow.Context) (*s3.GetBucketReplicationOutput, error) {
	var output s3.GetBucketReplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketRequestPaymentResult struct {
	Result workflow.Future
}

func (r *S3GetBucketRequestPaymentResult) Get(ctx workflow.Context) (*s3.GetBucketRequestPaymentOutput, error) {
	var output s3.GetBucketRequestPaymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketTaggingResult struct {
	Result workflow.Future
}

func (r *S3GetBucketTaggingResult) Get(ctx workflow.Context) (*s3.GetBucketTaggingOutput, error) {
	var output s3.GetBucketTaggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketVersioningResult struct {
	Result workflow.Future
}

func (r *S3GetBucketVersioningResult) Get(ctx workflow.Context) (*s3.GetBucketVersioningOutput, error) {
	var output s3.GetBucketVersioningOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetBucketWebsiteResult struct {
	Result workflow.Future
}

func (r *S3GetBucketWebsiteResult) Get(ctx workflow.Context) (*s3.GetBucketWebsiteOutput, error) {
	var output s3.GetBucketWebsiteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectResult struct {
	Result workflow.Future
}

func (r *S3GetObjectResult) Get(ctx workflow.Context) (*s3.GetObjectOutput, error) {
	var output s3.GetObjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectAclResult struct {
	Result workflow.Future
}

func (r *S3GetObjectAclResult) Get(ctx workflow.Context) (*s3.GetObjectAclOutput, error) {
	var output s3.GetObjectAclOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectLegalHoldResult struct {
	Result workflow.Future
}

func (r *S3GetObjectLegalHoldResult) Get(ctx workflow.Context) (*s3.GetObjectLegalHoldOutput, error) {
	var output s3.GetObjectLegalHoldOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectLockConfigurationResult struct {
	Result workflow.Future
}

func (r *S3GetObjectLockConfigurationResult) Get(ctx workflow.Context) (*s3.GetObjectLockConfigurationOutput, error) {
	var output s3.GetObjectLockConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectRetentionResult struct {
	Result workflow.Future
}

func (r *S3GetObjectRetentionResult) Get(ctx workflow.Context) (*s3.GetObjectRetentionOutput, error) {
	var output s3.GetObjectRetentionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectTaggingResult struct {
	Result workflow.Future
}

func (r *S3GetObjectTaggingResult) Get(ctx workflow.Context) (*s3.GetObjectTaggingOutput, error) {
	var output s3.GetObjectTaggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetObjectTorrentResult struct {
	Result workflow.Future
}

func (r *S3GetObjectTorrentResult) Get(ctx workflow.Context) (*s3.GetObjectTorrentOutput, error) {
	var output s3.GetObjectTorrentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3GetPublicAccessBlockResult struct {
	Result workflow.Future
}

func (r *S3GetPublicAccessBlockResult) Get(ctx workflow.Context) (*s3.GetPublicAccessBlockOutput, error) {
	var output s3.GetPublicAccessBlockOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3HeadBucketResult struct {
	Result workflow.Future
}

func (r *S3HeadBucketResult) Get(ctx workflow.Context) (*s3.HeadBucketOutput, error) {
	var output s3.HeadBucketOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3HeadObjectResult struct {
	Result workflow.Future
}

func (r *S3HeadObjectResult) Get(ctx workflow.Context) (*s3.HeadObjectOutput, error) {
	var output s3.HeadObjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListBucketAnalyticsConfigurationsResult struct {
	Result workflow.Future
}

func (r *S3ListBucketAnalyticsConfigurationsResult) Get(ctx workflow.Context) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	var output s3.ListBucketAnalyticsConfigurationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListBucketInventoryConfigurationsResult struct {
	Result workflow.Future
}

func (r *S3ListBucketInventoryConfigurationsResult) Get(ctx workflow.Context) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	var output s3.ListBucketInventoryConfigurationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListBucketMetricsConfigurationsResult struct {
	Result workflow.Future
}

func (r *S3ListBucketMetricsConfigurationsResult) Get(ctx workflow.Context) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	var output s3.ListBucketMetricsConfigurationsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListBucketsResult struct {
	Result workflow.Future
}

func (r *S3ListBucketsResult) Get(ctx workflow.Context) (*s3.ListBucketsOutput, error) {
	var output s3.ListBucketsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListMultipartUploadsResult struct {
	Result workflow.Future
}

func (r *S3ListMultipartUploadsResult) Get(ctx workflow.Context) (*s3.ListMultipartUploadsOutput, error) {
	var output s3.ListMultipartUploadsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListObjectVersionsResult struct {
	Result workflow.Future
}

func (r *S3ListObjectVersionsResult) Get(ctx workflow.Context) (*s3.ListObjectVersionsOutput, error) {
	var output s3.ListObjectVersionsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListObjectsResult struct {
	Result workflow.Future
}

func (r *S3ListObjectsResult) Get(ctx workflow.Context) (*s3.ListObjectsOutput, error) {
	var output s3.ListObjectsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListObjectsV2Result struct {
	Result workflow.Future
}

func (r *S3ListObjectsV2Result) Get(ctx workflow.Context) (*s3.ListObjectsV2Output, error) {
	var output s3.ListObjectsV2Output
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3ListPartsResult struct {
	Result workflow.Future
}

func (r *S3ListPartsResult) Get(ctx workflow.Context) (*s3.ListPartsOutput, error) {
	var output s3.ListPartsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketAccelerateConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketAccelerateConfigurationResult) Get(ctx workflow.Context) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	var output s3.PutBucketAccelerateConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketAclResult struct {
	Result workflow.Future
}

func (r *S3PutBucketAclResult) Get(ctx workflow.Context) (*s3.PutBucketAclOutput, error) {
	var output s3.PutBucketAclOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketAnalyticsConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketAnalyticsConfigurationResult) Get(ctx workflow.Context) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	var output s3.PutBucketAnalyticsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketCorsResult struct {
	Result workflow.Future
}

func (r *S3PutBucketCorsResult) Get(ctx workflow.Context) (*s3.PutBucketCorsOutput, error) {
	var output s3.PutBucketCorsOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketEncryptionResult struct {
	Result workflow.Future
}

func (r *S3PutBucketEncryptionResult) Get(ctx workflow.Context) (*s3.PutBucketEncryptionOutput, error) {
	var output s3.PutBucketEncryptionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketInventoryConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketInventoryConfigurationResult) Get(ctx workflow.Context) (*s3.PutBucketInventoryConfigurationOutput, error) {
	var output s3.PutBucketInventoryConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketLifecycleResult struct {
	Result workflow.Future
}

func (r *S3PutBucketLifecycleResult) Get(ctx workflow.Context) (*s3.PutBucketLifecycleOutput, error) {
	var output s3.PutBucketLifecycleOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketLifecycleConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketLifecycleConfigurationResult) Get(ctx workflow.Context) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	var output s3.PutBucketLifecycleConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketLoggingResult struct {
	Result workflow.Future
}

func (r *S3PutBucketLoggingResult) Get(ctx workflow.Context) (*s3.PutBucketLoggingOutput, error) {
	var output s3.PutBucketLoggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketMetricsConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketMetricsConfigurationResult) Get(ctx workflow.Context) (*s3.PutBucketMetricsConfigurationOutput, error) {
	var output s3.PutBucketMetricsConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketNotificationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketNotificationResult) Get(ctx workflow.Context) (*s3.PutBucketNotificationOutput, error) {
	var output s3.PutBucketNotificationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketNotificationConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketNotificationConfigurationResult) Get(ctx workflow.Context) (*s3.PutBucketNotificationConfigurationOutput, error) {
	var output s3.PutBucketNotificationConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketPolicyResult struct {
	Result workflow.Future
}

func (r *S3PutBucketPolicyResult) Get(ctx workflow.Context) (*s3.PutBucketPolicyOutput, error) {
	var output s3.PutBucketPolicyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketReplicationResult struct {
	Result workflow.Future
}

func (r *S3PutBucketReplicationResult) Get(ctx workflow.Context) (*s3.PutBucketReplicationOutput, error) {
	var output s3.PutBucketReplicationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketRequestPaymentResult struct {
	Result workflow.Future
}

func (r *S3PutBucketRequestPaymentResult) Get(ctx workflow.Context) (*s3.PutBucketRequestPaymentOutput, error) {
	var output s3.PutBucketRequestPaymentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketTaggingResult struct {
	Result workflow.Future
}

func (r *S3PutBucketTaggingResult) Get(ctx workflow.Context) (*s3.PutBucketTaggingOutput, error) {
	var output s3.PutBucketTaggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketVersioningResult struct {
	Result workflow.Future
}

func (r *S3PutBucketVersioningResult) Get(ctx workflow.Context) (*s3.PutBucketVersioningOutput, error) {
	var output s3.PutBucketVersioningOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutBucketWebsiteResult struct {
	Result workflow.Future
}

func (r *S3PutBucketWebsiteResult) Get(ctx workflow.Context) (*s3.PutBucketWebsiteOutput, error) {
	var output s3.PutBucketWebsiteOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutObjectResult struct {
	Result workflow.Future
}

func (r *S3PutObjectResult) Get(ctx workflow.Context) (*s3.PutObjectOutput, error) {
	var output s3.PutObjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutObjectAclResult struct {
	Result workflow.Future
}

func (r *S3PutObjectAclResult) Get(ctx workflow.Context) (*s3.PutObjectAclOutput, error) {
	var output s3.PutObjectAclOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutObjectLegalHoldResult struct {
	Result workflow.Future
}

func (r *S3PutObjectLegalHoldResult) Get(ctx workflow.Context) (*s3.PutObjectLegalHoldOutput, error) {
	var output s3.PutObjectLegalHoldOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutObjectLockConfigurationResult struct {
	Result workflow.Future
}

func (r *S3PutObjectLockConfigurationResult) Get(ctx workflow.Context) (*s3.PutObjectLockConfigurationOutput, error) {
	var output s3.PutObjectLockConfigurationOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutObjectRetentionResult struct {
	Result workflow.Future
}

func (r *S3PutObjectRetentionResult) Get(ctx workflow.Context) (*s3.PutObjectRetentionOutput, error) {
	var output s3.PutObjectRetentionOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutObjectTaggingResult struct {
	Result workflow.Future
}

func (r *S3PutObjectTaggingResult) Get(ctx workflow.Context) (*s3.PutObjectTaggingOutput, error) {
	var output s3.PutObjectTaggingOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3PutPublicAccessBlockResult struct {
	Result workflow.Future
}

func (r *S3PutPublicAccessBlockResult) Get(ctx workflow.Context) (*s3.PutPublicAccessBlockOutput, error) {
	var output s3.PutPublicAccessBlockOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3RestoreObjectResult struct {
	Result workflow.Future
}

func (r *S3RestoreObjectResult) Get(ctx workflow.Context) (*s3.RestoreObjectOutput, error) {
	var output s3.RestoreObjectOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3SelectObjectContentResult struct {
	Result workflow.Future
}

func (r *S3SelectObjectContentResult) Get(ctx workflow.Context) (*s3.SelectObjectContentOutput, error) {
	var output s3.SelectObjectContentOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3UploadPartResult struct {
	Result workflow.Future
}

func (r *S3UploadPartResult) Get(ctx workflow.Context) (*s3.UploadPartOutput, error) {
	var output s3.UploadPartOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3UploadPartCopyResult struct {
	Result workflow.Future
}

func (r *S3UploadPartCopyResult) Get(ctx workflow.Context) (*s3.UploadPartCopyOutput, error) {
	var output s3.UploadPartCopyOutput
	err := r.Result.Get(ctx, &output)
	return &output, err
}

type S3Stub struct {
	activities awsactivities.S3Activities
}

func NewS3Stub() S3Client {
	return &S3Stub{}
}

func (a *S3Stub) AbortMultipartUpload(ctx workflow.Context, input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	var output s3.AbortMultipartUploadOutput
	err := workflow.ExecuteActivity(ctx, a.activities.AbortMultipartUpload, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) AbortMultipartUploadAsync(ctx workflow.Context, input *s3.AbortMultipartUploadInput) *S3AbortMultipartUploadResult {
	future := workflow.ExecuteActivity(ctx, a.activities.AbortMultipartUpload, input)
	return &S3AbortMultipartUploadResult{Result: future}
}

func (a *S3Stub) CompleteMultipartUpload(ctx workflow.Context, input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	var output s3.CompleteMultipartUploadOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CompleteMultipartUpload, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) CompleteMultipartUploadAsync(ctx workflow.Context, input *s3.CompleteMultipartUploadInput) *S3CompleteMultipartUploadResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CompleteMultipartUpload, input)
	return &S3CompleteMultipartUploadResult{Result: future}
}

func (a *S3Stub) CopyObject(ctx workflow.Context, input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	var output s3.CopyObjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CopyObject, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) CopyObjectAsync(ctx workflow.Context, input *s3.CopyObjectInput) *S3CopyObjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CopyObject, input)
	return &S3CopyObjectResult{Result: future}
}

func (a *S3Stub) CreateBucket(ctx workflow.Context, input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	var output s3.CreateBucketOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateBucket, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) CreateBucketAsync(ctx workflow.Context, input *s3.CreateBucketInput) *S3CreateBucketResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateBucket, input)
	return &S3CreateBucketResult{Result: future}
}

func (a *S3Stub) CreateMultipartUpload(ctx workflow.Context, input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	var output s3.CreateMultipartUploadOutput
	err := workflow.ExecuteActivity(ctx, a.activities.CreateMultipartUpload, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) CreateMultipartUploadAsync(ctx workflow.Context, input *s3.CreateMultipartUploadInput) *S3CreateMultipartUploadResult {
	future := workflow.ExecuteActivity(ctx, a.activities.CreateMultipartUpload, input)
	return &S3CreateMultipartUploadResult{Result: future}
}

func (a *S3Stub) DeleteBucket(ctx workflow.Context, input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	var output s3.DeleteBucketOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucket, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketAsync(ctx workflow.Context, input *s3.DeleteBucketInput) *S3DeleteBucketResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucket, input)
	return &S3DeleteBucketResult{Result: future}
}

func (a *S3Stub) DeleteBucketAnalyticsConfiguration(ctx workflow.Context, input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	var output s3.DeleteBucketAnalyticsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketAnalyticsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketAnalyticsConfigurationAsync(ctx workflow.Context, input *s3.DeleteBucketAnalyticsConfigurationInput) *S3DeleteBucketAnalyticsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketAnalyticsConfiguration, input)
	return &S3DeleteBucketAnalyticsConfigurationResult{Result: future}
}

func (a *S3Stub) DeleteBucketCors(ctx workflow.Context, input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	var output s3.DeleteBucketCorsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketCors, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketCorsAsync(ctx workflow.Context, input *s3.DeleteBucketCorsInput) *S3DeleteBucketCorsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketCors, input)
	return &S3DeleteBucketCorsResult{Result: future}
}

func (a *S3Stub) DeleteBucketEncryption(ctx workflow.Context, input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	var output s3.DeleteBucketEncryptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketEncryption, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketEncryptionAsync(ctx workflow.Context, input *s3.DeleteBucketEncryptionInput) *S3DeleteBucketEncryptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketEncryption, input)
	return &S3DeleteBucketEncryptionResult{Result: future}
}

func (a *S3Stub) DeleteBucketInventoryConfiguration(ctx workflow.Context, input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	var output s3.DeleteBucketInventoryConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketInventoryConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketInventoryConfigurationAsync(ctx workflow.Context, input *s3.DeleteBucketInventoryConfigurationInput) *S3DeleteBucketInventoryConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketInventoryConfiguration, input)
	return &S3DeleteBucketInventoryConfigurationResult{Result: future}
}

func (a *S3Stub) DeleteBucketLifecycle(ctx workflow.Context, input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	var output s3.DeleteBucketLifecycleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketLifecycle, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketLifecycleAsync(ctx workflow.Context, input *s3.DeleteBucketLifecycleInput) *S3DeleteBucketLifecycleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketLifecycle, input)
	return &S3DeleteBucketLifecycleResult{Result: future}
}

func (a *S3Stub) DeleteBucketMetricsConfiguration(ctx workflow.Context, input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	var output s3.DeleteBucketMetricsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketMetricsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketMetricsConfigurationAsync(ctx workflow.Context, input *s3.DeleteBucketMetricsConfigurationInput) *S3DeleteBucketMetricsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketMetricsConfiguration, input)
	return &S3DeleteBucketMetricsConfigurationResult{Result: future}
}

func (a *S3Stub) DeleteBucketPolicy(ctx workflow.Context, input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	var output s3.DeleteBucketPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketPolicyAsync(ctx workflow.Context, input *s3.DeleteBucketPolicyInput) *S3DeleteBucketPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketPolicy, input)
	return &S3DeleteBucketPolicyResult{Result: future}
}

func (a *S3Stub) DeleteBucketReplication(ctx workflow.Context, input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	var output s3.DeleteBucketReplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketReplication, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketReplicationAsync(ctx workflow.Context, input *s3.DeleteBucketReplicationInput) *S3DeleteBucketReplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketReplication, input)
	return &S3DeleteBucketReplicationResult{Result: future}
}

func (a *S3Stub) DeleteBucketTagging(ctx workflow.Context, input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	var output s3.DeleteBucketTaggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketTagging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketTaggingAsync(ctx workflow.Context, input *s3.DeleteBucketTaggingInput) *S3DeleteBucketTaggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketTagging, input)
	return &S3DeleteBucketTaggingResult{Result: future}
}

func (a *S3Stub) DeleteBucketWebsite(ctx workflow.Context, input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	var output s3.DeleteBucketWebsiteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketWebsite, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteBucketWebsiteAsync(ctx workflow.Context, input *s3.DeleteBucketWebsiteInput) *S3DeleteBucketWebsiteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteBucketWebsite, input)
	return &S3DeleteBucketWebsiteResult{Result: future}
}

func (a *S3Stub) DeleteObject(ctx workflow.Context, input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	var output s3.DeleteObjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteObject, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteObjectAsync(ctx workflow.Context, input *s3.DeleteObjectInput) *S3DeleteObjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteObject, input)
	return &S3DeleteObjectResult{Result: future}
}

func (a *S3Stub) DeleteObjectTagging(ctx workflow.Context, input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	var output s3.DeleteObjectTaggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteObjectTagging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteObjectTaggingAsync(ctx workflow.Context, input *s3.DeleteObjectTaggingInput) *S3DeleteObjectTaggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteObjectTagging, input)
	return &S3DeleteObjectTaggingResult{Result: future}
}

func (a *S3Stub) DeleteObjects(ctx workflow.Context, input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	var output s3.DeleteObjectsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeleteObjects, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeleteObjectsAsync(ctx workflow.Context, input *s3.DeleteObjectsInput) *S3DeleteObjectsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeleteObjects, input)
	return &S3DeleteObjectsResult{Result: future}
}

func (a *S3Stub) DeletePublicAccessBlock(ctx workflow.Context, input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	var output s3.DeletePublicAccessBlockOutput
	err := workflow.ExecuteActivity(ctx, a.activities.DeletePublicAccessBlock, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) DeletePublicAccessBlockAsync(ctx workflow.Context, input *s3.DeletePublicAccessBlockInput) *S3DeletePublicAccessBlockResult {
	future := workflow.ExecuteActivity(ctx, a.activities.DeletePublicAccessBlock, input)
	return &S3DeletePublicAccessBlockResult{Result: future}
}

func (a *S3Stub) GetBucketAccelerateConfiguration(ctx workflow.Context, input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	var output s3.GetBucketAccelerateConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketAccelerateConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketAccelerateConfigurationAsync(ctx workflow.Context, input *s3.GetBucketAccelerateConfigurationInput) *S3GetBucketAccelerateConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketAccelerateConfiguration, input)
	return &S3GetBucketAccelerateConfigurationResult{Result: future}
}

func (a *S3Stub) GetBucketAcl(ctx workflow.Context, input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	var output s3.GetBucketAclOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketAcl, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketAclAsync(ctx workflow.Context, input *s3.GetBucketAclInput) *S3GetBucketAclResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketAcl, input)
	return &S3GetBucketAclResult{Result: future}
}

func (a *S3Stub) GetBucketAnalyticsConfiguration(ctx workflow.Context, input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	var output s3.GetBucketAnalyticsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketAnalyticsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketAnalyticsConfigurationAsync(ctx workflow.Context, input *s3.GetBucketAnalyticsConfigurationInput) *S3GetBucketAnalyticsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketAnalyticsConfiguration, input)
	return &S3GetBucketAnalyticsConfigurationResult{Result: future}
}

func (a *S3Stub) GetBucketCors(ctx workflow.Context, input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	var output s3.GetBucketCorsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketCors, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketCorsAsync(ctx workflow.Context, input *s3.GetBucketCorsInput) *S3GetBucketCorsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketCors, input)
	return &S3GetBucketCorsResult{Result: future}
}

func (a *S3Stub) GetBucketEncryption(ctx workflow.Context, input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	var output s3.GetBucketEncryptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketEncryption, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketEncryptionAsync(ctx workflow.Context, input *s3.GetBucketEncryptionInput) *S3GetBucketEncryptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketEncryption, input)
	return &S3GetBucketEncryptionResult{Result: future}
}

func (a *S3Stub) GetBucketInventoryConfiguration(ctx workflow.Context, input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	var output s3.GetBucketInventoryConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketInventoryConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketInventoryConfigurationAsync(ctx workflow.Context, input *s3.GetBucketInventoryConfigurationInput) *S3GetBucketInventoryConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketInventoryConfiguration, input)
	return &S3GetBucketInventoryConfigurationResult{Result: future}
}

func (a *S3Stub) GetBucketLifecycle(ctx workflow.Context, input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	var output s3.GetBucketLifecycleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketLifecycle, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketLifecycleAsync(ctx workflow.Context, input *s3.GetBucketLifecycleInput) *S3GetBucketLifecycleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketLifecycle, input)
	return &S3GetBucketLifecycleResult{Result: future}
}

func (a *S3Stub) GetBucketLifecycleConfiguration(ctx workflow.Context, input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	var output s3.GetBucketLifecycleConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketLifecycleConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3.GetBucketLifecycleConfigurationInput) *S3GetBucketLifecycleConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketLifecycleConfiguration, input)
	return &S3GetBucketLifecycleConfigurationResult{Result: future}
}

func (a *S3Stub) GetBucketLocation(ctx workflow.Context, input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	var output s3.GetBucketLocationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketLocation, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketLocationAsync(ctx workflow.Context, input *s3.GetBucketLocationInput) *S3GetBucketLocationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketLocation, input)
	return &S3GetBucketLocationResult{Result: future}
}

func (a *S3Stub) GetBucketLogging(ctx workflow.Context, input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	var output s3.GetBucketLoggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketLogging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketLoggingAsync(ctx workflow.Context, input *s3.GetBucketLoggingInput) *S3GetBucketLoggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketLogging, input)
	return &S3GetBucketLoggingResult{Result: future}
}

func (a *S3Stub) GetBucketMetricsConfiguration(ctx workflow.Context, input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	var output s3.GetBucketMetricsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketMetricsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketMetricsConfigurationAsync(ctx workflow.Context, input *s3.GetBucketMetricsConfigurationInput) *S3GetBucketMetricsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketMetricsConfiguration, input)
	return &S3GetBucketMetricsConfigurationResult{Result: future}
}

func (a *S3Stub) GetBucketNotification(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	var output s3.NotificationConfigurationDeprecated
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketNotification, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketNotificationAsync(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) *S3GetBucketNotificationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketNotification, input)
	return &S3GetBucketNotificationResult{Result: future}
}

func (a *S3Stub) GetBucketNotificationConfiguration(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	var output s3.NotificationConfiguration
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketNotificationConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketNotificationConfigurationAsync(ctx workflow.Context, input *s3.GetBucketNotificationConfigurationRequest) *S3GetBucketNotificationConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketNotificationConfiguration, input)
	return &S3GetBucketNotificationConfigurationResult{Result: future}
}

func (a *S3Stub) GetBucketPolicy(ctx workflow.Context, input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	var output s3.GetBucketPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketPolicyAsync(ctx workflow.Context, input *s3.GetBucketPolicyInput) *S3GetBucketPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketPolicy, input)
	return &S3GetBucketPolicyResult{Result: future}
}

func (a *S3Stub) GetBucketPolicyStatus(ctx workflow.Context, input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	var output s3.GetBucketPolicyStatusOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketPolicyStatus, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketPolicyStatusAsync(ctx workflow.Context, input *s3.GetBucketPolicyStatusInput) *S3GetBucketPolicyStatusResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketPolicyStatus, input)
	return &S3GetBucketPolicyStatusResult{Result: future}
}

func (a *S3Stub) GetBucketReplication(ctx workflow.Context, input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	var output s3.GetBucketReplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketReplication, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketReplicationAsync(ctx workflow.Context, input *s3.GetBucketReplicationInput) *S3GetBucketReplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketReplication, input)
	return &S3GetBucketReplicationResult{Result: future}
}

func (a *S3Stub) GetBucketRequestPayment(ctx workflow.Context, input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	var output s3.GetBucketRequestPaymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketRequestPayment, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketRequestPaymentAsync(ctx workflow.Context, input *s3.GetBucketRequestPaymentInput) *S3GetBucketRequestPaymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketRequestPayment, input)
	return &S3GetBucketRequestPaymentResult{Result: future}
}

func (a *S3Stub) GetBucketTagging(ctx workflow.Context, input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	var output s3.GetBucketTaggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketTagging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketTaggingAsync(ctx workflow.Context, input *s3.GetBucketTaggingInput) *S3GetBucketTaggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketTagging, input)
	return &S3GetBucketTaggingResult{Result: future}
}

func (a *S3Stub) GetBucketVersioning(ctx workflow.Context, input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	var output s3.GetBucketVersioningOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketVersioning, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketVersioningAsync(ctx workflow.Context, input *s3.GetBucketVersioningInput) *S3GetBucketVersioningResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketVersioning, input)
	return &S3GetBucketVersioningResult{Result: future}
}

func (a *S3Stub) GetBucketWebsite(ctx workflow.Context, input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	var output s3.GetBucketWebsiteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetBucketWebsite, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetBucketWebsiteAsync(ctx workflow.Context, input *s3.GetBucketWebsiteInput) *S3GetBucketWebsiteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetBucketWebsite, input)
	return &S3GetBucketWebsiteResult{Result: future}
}

func (a *S3Stub) GetObject(ctx workflow.Context, input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	var output s3.GetObjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObject, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectAsync(ctx workflow.Context, input *s3.GetObjectInput) *S3GetObjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObject, input)
	return &S3GetObjectResult{Result: future}
}

func (a *S3Stub) GetObjectAcl(ctx workflow.Context, input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	var output s3.GetObjectAclOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObjectAcl, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectAclAsync(ctx workflow.Context, input *s3.GetObjectAclInput) *S3GetObjectAclResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObjectAcl, input)
	return &S3GetObjectAclResult{Result: future}
}

func (a *S3Stub) GetObjectLegalHold(ctx workflow.Context, input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	var output s3.GetObjectLegalHoldOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObjectLegalHold, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectLegalHoldAsync(ctx workflow.Context, input *s3.GetObjectLegalHoldInput) *S3GetObjectLegalHoldResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObjectLegalHold, input)
	return &S3GetObjectLegalHoldResult{Result: future}
}

func (a *S3Stub) GetObjectLockConfiguration(ctx workflow.Context, input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	var output s3.GetObjectLockConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObjectLockConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectLockConfigurationAsync(ctx workflow.Context, input *s3.GetObjectLockConfigurationInput) *S3GetObjectLockConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObjectLockConfiguration, input)
	return &S3GetObjectLockConfigurationResult{Result: future}
}

func (a *S3Stub) GetObjectRetention(ctx workflow.Context, input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	var output s3.GetObjectRetentionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObjectRetention, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectRetentionAsync(ctx workflow.Context, input *s3.GetObjectRetentionInput) *S3GetObjectRetentionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObjectRetention, input)
	return &S3GetObjectRetentionResult{Result: future}
}

func (a *S3Stub) GetObjectTagging(ctx workflow.Context, input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	var output s3.GetObjectTaggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObjectTagging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectTaggingAsync(ctx workflow.Context, input *s3.GetObjectTaggingInput) *S3GetObjectTaggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObjectTagging, input)
	return &S3GetObjectTaggingResult{Result: future}
}

func (a *S3Stub) GetObjectTorrent(ctx workflow.Context, input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	var output s3.GetObjectTorrentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetObjectTorrent, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetObjectTorrentAsync(ctx workflow.Context, input *s3.GetObjectTorrentInput) *S3GetObjectTorrentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetObjectTorrent, input)
	return &S3GetObjectTorrentResult{Result: future}
}

func (a *S3Stub) GetPublicAccessBlock(ctx workflow.Context, input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	var output s3.GetPublicAccessBlockOutput
	err := workflow.ExecuteActivity(ctx, a.activities.GetPublicAccessBlock, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) GetPublicAccessBlockAsync(ctx workflow.Context, input *s3.GetPublicAccessBlockInput) *S3GetPublicAccessBlockResult {
	future := workflow.ExecuteActivity(ctx, a.activities.GetPublicAccessBlock, input)
	return &S3GetPublicAccessBlockResult{Result: future}
}

func (a *S3Stub) HeadBucket(ctx workflow.Context, input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	var output s3.HeadBucketOutput
	err := workflow.ExecuteActivity(ctx, a.activities.HeadBucket, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) HeadBucketAsync(ctx workflow.Context, input *s3.HeadBucketInput) *S3HeadBucketResult {
	future := workflow.ExecuteActivity(ctx, a.activities.HeadBucket, input)
	return &S3HeadBucketResult{Result: future}
}

func (a *S3Stub) HeadObject(ctx workflow.Context, input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	var output s3.HeadObjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.HeadObject, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) HeadObjectAsync(ctx workflow.Context, input *s3.HeadObjectInput) *S3HeadObjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.HeadObject, input)
	return &S3HeadObjectResult{Result: future}
}

func (a *S3Stub) ListBucketAnalyticsConfigurations(ctx workflow.Context, input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	var output s3.ListBucketAnalyticsConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBucketAnalyticsConfigurations, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListBucketAnalyticsConfigurationsAsync(ctx workflow.Context, input *s3.ListBucketAnalyticsConfigurationsInput) *S3ListBucketAnalyticsConfigurationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBucketAnalyticsConfigurations, input)
	return &S3ListBucketAnalyticsConfigurationsResult{Result: future}
}

func (a *S3Stub) ListBucketInventoryConfigurations(ctx workflow.Context, input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	var output s3.ListBucketInventoryConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBucketInventoryConfigurations, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListBucketInventoryConfigurationsAsync(ctx workflow.Context, input *s3.ListBucketInventoryConfigurationsInput) *S3ListBucketInventoryConfigurationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBucketInventoryConfigurations, input)
	return &S3ListBucketInventoryConfigurationsResult{Result: future}
}

func (a *S3Stub) ListBucketMetricsConfigurations(ctx workflow.Context, input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	var output s3.ListBucketMetricsConfigurationsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBucketMetricsConfigurations, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListBucketMetricsConfigurationsAsync(ctx workflow.Context, input *s3.ListBucketMetricsConfigurationsInput) *S3ListBucketMetricsConfigurationsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBucketMetricsConfigurations, input)
	return &S3ListBucketMetricsConfigurationsResult{Result: future}
}

func (a *S3Stub) ListBuckets(ctx workflow.Context, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	var output s3.ListBucketsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListBuckets, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListBucketsAsync(ctx workflow.Context, input *s3.ListBucketsInput) *S3ListBucketsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListBuckets, input)
	return &S3ListBucketsResult{Result: future}
}

func (a *S3Stub) ListMultipartUploads(ctx workflow.Context, input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	var output s3.ListMultipartUploadsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListMultipartUploads, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListMultipartUploadsAsync(ctx workflow.Context, input *s3.ListMultipartUploadsInput) *S3ListMultipartUploadsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListMultipartUploads, input)
	return &S3ListMultipartUploadsResult{Result: future}
}

func (a *S3Stub) ListObjectVersions(ctx workflow.Context, input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	var output s3.ListObjectVersionsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListObjectVersions, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListObjectVersionsAsync(ctx workflow.Context, input *s3.ListObjectVersionsInput) *S3ListObjectVersionsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListObjectVersions, input)
	return &S3ListObjectVersionsResult{Result: future}
}

func (a *S3Stub) ListObjects(ctx workflow.Context, input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	var output s3.ListObjectsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListObjects, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListObjectsAsync(ctx workflow.Context, input *s3.ListObjectsInput) *S3ListObjectsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListObjects, input)
	return &S3ListObjectsResult{Result: future}
}

func (a *S3Stub) ListObjectsV2(ctx workflow.Context, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	var output s3.ListObjectsV2Output
	err := workflow.ExecuteActivity(ctx, a.activities.ListObjectsV2, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListObjectsV2Async(ctx workflow.Context, input *s3.ListObjectsV2Input) *S3ListObjectsV2Result {
	future := workflow.ExecuteActivity(ctx, a.activities.ListObjectsV2, input)
	return &S3ListObjectsV2Result{Result: future}
}

func (a *S3Stub) ListParts(ctx workflow.Context, input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	var output s3.ListPartsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.ListParts, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) ListPartsAsync(ctx workflow.Context, input *s3.ListPartsInput) *S3ListPartsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.ListParts, input)
	return &S3ListPartsResult{Result: future}
}

func (a *S3Stub) PutBucketAccelerateConfiguration(ctx workflow.Context, input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	var output s3.PutBucketAccelerateConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketAccelerateConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketAccelerateConfigurationAsync(ctx workflow.Context, input *s3.PutBucketAccelerateConfigurationInput) *S3PutBucketAccelerateConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketAccelerateConfiguration, input)
	return &S3PutBucketAccelerateConfigurationResult{Result: future}
}

func (a *S3Stub) PutBucketAcl(ctx workflow.Context, input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	var output s3.PutBucketAclOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketAcl, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketAclAsync(ctx workflow.Context, input *s3.PutBucketAclInput) *S3PutBucketAclResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketAcl, input)
	return &S3PutBucketAclResult{Result: future}
}

func (a *S3Stub) PutBucketAnalyticsConfiguration(ctx workflow.Context, input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	var output s3.PutBucketAnalyticsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketAnalyticsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketAnalyticsConfigurationAsync(ctx workflow.Context, input *s3.PutBucketAnalyticsConfigurationInput) *S3PutBucketAnalyticsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketAnalyticsConfiguration, input)
	return &S3PutBucketAnalyticsConfigurationResult{Result: future}
}

func (a *S3Stub) PutBucketCors(ctx workflow.Context, input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	var output s3.PutBucketCorsOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketCors, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketCorsAsync(ctx workflow.Context, input *s3.PutBucketCorsInput) *S3PutBucketCorsResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketCors, input)
	return &S3PutBucketCorsResult{Result: future}
}

func (a *S3Stub) PutBucketEncryption(ctx workflow.Context, input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	var output s3.PutBucketEncryptionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketEncryption, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketEncryptionAsync(ctx workflow.Context, input *s3.PutBucketEncryptionInput) *S3PutBucketEncryptionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketEncryption, input)
	return &S3PutBucketEncryptionResult{Result: future}
}

func (a *S3Stub) PutBucketInventoryConfiguration(ctx workflow.Context, input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	var output s3.PutBucketInventoryConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketInventoryConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketInventoryConfigurationAsync(ctx workflow.Context, input *s3.PutBucketInventoryConfigurationInput) *S3PutBucketInventoryConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketInventoryConfiguration, input)
	return &S3PutBucketInventoryConfigurationResult{Result: future}
}

func (a *S3Stub) PutBucketLifecycle(ctx workflow.Context, input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	var output s3.PutBucketLifecycleOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketLifecycle, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketLifecycleAsync(ctx workflow.Context, input *s3.PutBucketLifecycleInput) *S3PutBucketLifecycleResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketLifecycle, input)
	return &S3PutBucketLifecycleResult{Result: future}
}

func (a *S3Stub) PutBucketLifecycleConfiguration(ctx workflow.Context, input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	var output s3.PutBucketLifecycleConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketLifecycleConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketLifecycleConfigurationAsync(ctx workflow.Context, input *s3.PutBucketLifecycleConfigurationInput) *S3PutBucketLifecycleConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketLifecycleConfiguration, input)
	return &S3PutBucketLifecycleConfigurationResult{Result: future}
}

func (a *S3Stub) PutBucketLogging(ctx workflow.Context, input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	var output s3.PutBucketLoggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketLogging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketLoggingAsync(ctx workflow.Context, input *s3.PutBucketLoggingInput) *S3PutBucketLoggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketLogging, input)
	return &S3PutBucketLoggingResult{Result: future}
}

func (a *S3Stub) PutBucketMetricsConfiguration(ctx workflow.Context, input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	var output s3.PutBucketMetricsConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketMetricsConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketMetricsConfigurationAsync(ctx workflow.Context, input *s3.PutBucketMetricsConfigurationInput) *S3PutBucketMetricsConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketMetricsConfiguration, input)
	return &S3PutBucketMetricsConfigurationResult{Result: future}
}

func (a *S3Stub) PutBucketNotification(ctx workflow.Context, input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	var output s3.PutBucketNotificationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketNotification, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketNotificationAsync(ctx workflow.Context, input *s3.PutBucketNotificationInput) *S3PutBucketNotificationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketNotification, input)
	return &S3PutBucketNotificationResult{Result: future}
}

func (a *S3Stub) PutBucketNotificationConfiguration(ctx workflow.Context, input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	var output s3.PutBucketNotificationConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketNotificationConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketNotificationConfigurationAsync(ctx workflow.Context, input *s3.PutBucketNotificationConfigurationInput) *S3PutBucketNotificationConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketNotificationConfiguration, input)
	return &S3PutBucketNotificationConfigurationResult{Result: future}
}

func (a *S3Stub) PutBucketPolicy(ctx workflow.Context, input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	var output s3.PutBucketPolicyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketPolicy, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketPolicyAsync(ctx workflow.Context, input *s3.PutBucketPolicyInput) *S3PutBucketPolicyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketPolicy, input)
	return &S3PutBucketPolicyResult{Result: future}
}

func (a *S3Stub) PutBucketReplication(ctx workflow.Context, input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	var output s3.PutBucketReplicationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketReplication, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketReplicationAsync(ctx workflow.Context, input *s3.PutBucketReplicationInput) *S3PutBucketReplicationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketReplication, input)
	return &S3PutBucketReplicationResult{Result: future}
}

func (a *S3Stub) PutBucketRequestPayment(ctx workflow.Context, input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	var output s3.PutBucketRequestPaymentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketRequestPayment, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketRequestPaymentAsync(ctx workflow.Context, input *s3.PutBucketRequestPaymentInput) *S3PutBucketRequestPaymentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketRequestPayment, input)
	return &S3PutBucketRequestPaymentResult{Result: future}
}

func (a *S3Stub) PutBucketTagging(ctx workflow.Context, input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	var output s3.PutBucketTaggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketTagging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketTaggingAsync(ctx workflow.Context, input *s3.PutBucketTaggingInput) *S3PutBucketTaggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketTagging, input)
	return &S3PutBucketTaggingResult{Result: future}
}

func (a *S3Stub) PutBucketVersioning(ctx workflow.Context, input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	var output s3.PutBucketVersioningOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketVersioning, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketVersioningAsync(ctx workflow.Context, input *s3.PutBucketVersioningInput) *S3PutBucketVersioningResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketVersioning, input)
	return &S3PutBucketVersioningResult{Result: future}
}

func (a *S3Stub) PutBucketWebsite(ctx workflow.Context, input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	var output s3.PutBucketWebsiteOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutBucketWebsite, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutBucketWebsiteAsync(ctx workflow.Context, input *s3.PutBucketWebsiteInput) *S3PutBucketWebsiteResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutBucketWebsite, input)
	return &S3PutBucketWebsiteResult{Result: future}
}

func (a *S3Stub) PutObject(ctx workflow.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	var output s3.PutObjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutObject, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutObjectAsync(ctx workflow.Context, input *s3.PutObjectInput) *S3PutObjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutObject, input)
	return &S3PutObjectResult{Result: future}
}

func (a *S3Stub) PutObjectAcl(ctx workflow.Context, input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	var output s3.PutObjectAclOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutObjectAcl, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutObjectAclAsync(ctx workflow.Context, input *s3.PutObjectAclInput) *S3PutObjectAclResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutObjectAcl, input)
	return &S3PutObjectAclResult{Result: future}
}

func (a *S3Stub) PutObjectLegalHold(ctx workflow.Context, input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	var output s3.PutObjectLegalHoldOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutObjectLegalHold, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutObjectLegalHoldAsync(ctx workflow.Context, input *s3.PutObjectLegalHoldInput) *S3PutObjectLegalHoldResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutObjectLegalHold, input)
	return &S3PutObjectLegalHoldResult{Result: future}
}

func (a *S3Stub) PutObjectLockConfiguration(ctx workflow.Context, input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	var output s3.PutObjectLockConfigurationOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutObjectLockConfiguration, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutObjectLockConfigurationAsync(ctx workflow.Context, input *s3.PutObjectLockConfigurationInput) *S3PutObjectLockConfigurationResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutObjectLockConfiguration, input)
	return &S3PutObjectLockConfigurationResult{Result: future}
}

func (a *S3Stub) PutObjectRetention(ctx workflow.Context, input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	var output s3.PutObjectRetentionOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutObjectRetention, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutObjectRetentionAsync(ctx workflow.Context, input *s3.PutObjectRetentionInput) *S3PutObjectRetentionResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutObjectRetention, input)
	return &S3PutObjectRetentionResult{Result: future}
}

func (a *S3Stub) PutObjectTagging(ctx workflow.Context, input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	var output s3.PutObjectTaggingOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutObjectTagging, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutObjectTaggingAsync(ctx workflow.Context, input *s3.PutObjectTaggingInput) *S3PutObjectTaggingResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutObjectTagging, input)
	return &S3PutObjectTaggingResult{Result: future}
}

func (a *S3Stub) PutPublicAccessBlock(ctx workflow.Context, input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	var output s3.PutPublicAccessBlockOutput
	err := workflow.ExecuteActivity(ctx, a.activities.PutPublicAccessBlock, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) PutPublicAccessBlockAsync(ctx workflow.Context, input *s3.PutPublicAccessBlockInput) *S3PutPublicAccessBlockResult {
	future := workflow.ExecuteActivity(ctx, a.activities.PutPublicAccessBlock, input)
	return &S3PutPublicAccessBlockResult{Result: future}
}

func (a *S3Stub) RestoreObject(ctx workflow.Context, input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	var output s3.RestoreObjectOutput
	err := workflow.ExecuteActivity(ctx, a.activities.RestoreObject, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) RestoreObjectAsync(ctx workflow.Context, input *s3.RestoreObjectInput) *S3RestoreObjectResult {
	future := workflow.ExecuteActivity(ctx, a.activities.RestoreObject, input)
	return &S3RestoreObjectResult{Result: future}
}

func (a *S3Stub) SelectObjectContent(ctx workflow.Context, input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	var output s3.SelectObjectContentOutput
	err := workflow.ExecuteActivity(ctx, a.activities.SelectObjectContent, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) SelectObjectContentAsync(ctx workflow.Context, input *s3.SelectObjectContentInput) *S3SelectObjectContentResult {
	future := workflow.ExecuteActivity(ctx, a.activities.SelectObjectContent, input)
	return &S3SelectObjectContentResult{Result: future}
}

func (a *S3Stub) UploadPart(ctx workflow.Context, input *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	var output s3.UploadPartOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UploadPart, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) UploadPartAsync(ctx workflow.Context, input *s3.UploadPartInput) *S3UploadPartResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UploadPart, input)
	return &S3UploadPartResult{Result: future}
}

func (a *S3Stub) UploadPartCopy(ctx workflow.Context, input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	var output s3.UploadPartCopyOutput
	err := workflow.ExecuteActivity(ctx, a.activities.UploadPartCopy, input).Get(ctx, &output)
	return &output, err
}

func (a *S3Stub) UploadPartCopyAsync(ctx workflow.Context, input *s3.UploadPartCopyInput) *S3UploadPartCopyResult {
	future := workflow.ExecuteActivity(ctx, a.activities.UploadPartCopy, input)
	return &S3UploadPartCopyResult{Result: future}
}

func (a *S3Stub) WaitUntilBucketExists(ctx workflow.Context, input *s3.HeadBucketInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBucketExists, input).Get(ctx, nil)
}

func (a *S3Stub) WaitUntilBucketExistsAsync(ctx workflow.Context, input *s3.HeadBucketInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBucketExists, input)
}

func (a *S3Stub) WaitUntilBucketNotExists(ctx workflow.Context, input *s3.HeadBucketInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBucketNotExists, input).Get(ctx, nil)
}

func (a *S3Stub) WaitUntilBucketNotExistsAsync(ctx workflow.Context, input *s3.HeadBucketInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilBucketNotExists, input)
}

func (a *S3Stub) WaitUntilObjectExists(ctx workflow.Context, input *s3.HeadObjectInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilObjectExists, input).Get(ctx, nil)
}

func (a *S3Stub) WaitUntilObjectExistsAsync(ctx workflow.Context, input *s3.HeadObjectInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilObjectExists, input)
}

func (a *S3Stub) WaitUntilObjectNotExists(ctx workflow.Context, input *s3.HeadObjectInput) error {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilObjectNotExists, input).Get(ctx, nil)
}

func (a *S3Stub) WaitUntilObjectNotExistsAsync(ctx workflow.Context, input *s3.HeadObjectInput) workflow.Future {
	return workflow.ExecuteActivity(ctx, a.activities.WaitUntilObjectNotExists, input)
}
