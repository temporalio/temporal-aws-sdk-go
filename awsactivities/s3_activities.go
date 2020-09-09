
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type S3Activities struct {
    client s3iface.S3API
}

func NewS3Activities(session *session.Session, config ...*aws.Config) *S3Activities {
    client := s3.New(session, config...)
    return &S3Activities{client: client}
}

func (a *S3Activities) AbortMultipartUpload(input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
    return a.client.AbortMultipartUpload(input)
}

func (a *S3Activities) CompleteMultipartUpload(input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
    return a.client.CompleteMultipartUpload(input)
}

func (a *S3Activities) CopyObject(input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
    return a.client.CopyObject(input)
}

func (a *S3Activities) CreateBucket(input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
    return a.client.CreateBucket(input)
}

func (a *S3Activities) CreateMultipartUpload(input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
    return a.client.CreateMultipartUpload(input)
}

func (a *S3Activities) DeleteBucket(input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
    return a.client.DeleteBucket(input)
}

func (a *S3Activities) DeleteBucketAnalyticsConfiguration(input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
    return a.client.DeleteBucketAnalyticsConfiguration(input)
}

func (a *S3Activities) DeleteBucketCors(input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
    return a.client.DeleteBucketCors(input)
}

func (a *S3Activities) DeleteBucketEncryption(input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
    return a.client.DeleteBucketEncryption(input)
}

func (a *S3Activities) DeleteBucketInventoryConfiguration(input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
    return a.client.DeleteBucketInventoryConfiguration(input)
}

func (a *S3Activities) DeleteBucketLifecycle(input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
    return a.client.DeleteBucketLifecycle(input)
}

func (a *S3Activities) DeleteBucketMetricsConfiguration(input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
    return a.client.DeleteBucketMetricsConfiguration(input)
}

func (a *S3Activities) DeleteBucketPolicy(input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
    return a.client.DeleteBucketPolicy(input)
}

func (a *S3Activities) DeleteBucketReplication(input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
    return a.client.DeleteBucketReplication(input)
}

func (a *S3Activities) DeleteBucketTagging(input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
    return a.client.DeleteBucketTagging(input)
}

func (a *S3Activities) DeleteBucketWebsite(input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
    return a.client.DeleteBucketWebsite(input)
}

func (a *S3Activities) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
    return a.client.DeleteObject(input)
}

func (a *S3Activities) DeleteObjectTagging(input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
    return a.client.DeleteObjectTagging(input)
}

func (a *S3Activities) DeleteObjects(input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
    return a.client.DeleteObjects(input)
}

func (a *S3Activities) DeletePublicAccessBlock(input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
    return a.client.DeletePublicAccessBlock(input)
}

func (a *S3Activities) GetBucketAccelerateConfiguration(input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
    return a.client.GetBucketAccelerateConfiguration(input)
}

func (a *S3Activities) GetBucketAcl(input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
    return a.client.GetBucketAcl(input)
}

func (a *S3Activities) GetBucketAnalyticsConfiguration(input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
    return a.client.GetBucketAnalyticsConfiguration(input)
}

func (a *S3Activities) GetBucketCors(input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
    return a.client.GetBucketCors(input)
}

func (a *S3Activities) GetBucketEncryption(input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
    return a.client.GetBucketEncryption(input)
}

func (a *S3Activities) GetBucketInventoryConfiguration(input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
    return a.client.GetBucketInventoryConfiguration(input)
}

func (a *S3Activities) GetBucketLifecycle(input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
    return a.client.GetBucketLifecycle(input)
}

func (a *S3Activities) GetBucketLifecycleConfiguration(input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
    return a.client.GetBucketLifecycleConfiguration(input)
}

func (a *S3Activities) GetBucketLocation(input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
    return a.client.GetBucketLocation(input)
}

func (a *S3Activities) GetBucketLogging(input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
    return a.client.GetBucketLogging(input)
}

func (a *S3Activities) GetBucketMetricsConfiguration(input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
    return a.client.GetBucketMetricsConfiguration(input)
}

func (a *S3Activities) GetBucketNotification(input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
    return a.client.GetBucketNotification(input)
}

func (a *S3Activities) GetBucketNotificationConfiguration(input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
    return a.client.GetBucketNotificationConfiguration(input)
}

func (a *S3Activities) GetBucketPolicy(input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
    return a.client.GetBucketPolicy(input)
}

func (a *S3Activities) GetBucketPolicyStatus(input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
    return a.client.GetBucketPolicyStatus(input)
}

func (a *S3Activities) GetBucketReplication(input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
    return a.client.GetBucketReplication(input)
}

func (a *S3Activities) GetBucketRequestPayment(input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
    return a.client.GetBucketRequestPayment(input)
}

func (a *S3Activities) GetBucketTagging(input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
    return a.client.GetBucketTagging(input)
}

func (a *S3Activities) GetBucketVersioning(input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
    return a.client.GetBucketVersioning(input)
}

func (a *S3Activities) GetBucketWebsite(input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
    return a.client.GetBucketWebsite(input)
}

func (a *S3Activities) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
    return a.client.GetObject(input)
}

func (a *S3Activities) GetObjectAcl(input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
    return a.client.GetObjectAcl(input)
}

func (a *S3Activities) GetObjectLegalHold(input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
    return a.client.GetObjectLegalHold(input)
}

func (a *S3Activities) GetObjectLockConfiguration(input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
    return a.client.GetObjectLockConfiguration(input)
}

func (a *S3Activities) GetObjectRetention(input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
    return a.client.GetObjectRetention(input)
}

func (a *S3Activities) GetObjectTagging(input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
    return a.client.GetObjectTagging(input)
}

func (a *S3Activities) GetObjectTorrent(input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
    return a.client.GetObjectTorrent(input)
}

func (a *S3Activities) GetPublicAccessBlock(input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
    return a.client.GetPublicAccessBlock(input)
}

func (a *S3Activities) HeadBucket(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
    return a.client.HeadBucket(input)
}

func (a *S3Activities) HeadObject(input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
    return a.client.HeadObject(input)
}

func (a *S3Activities) ListBucketAnalyticsConfigurations(input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
    return a.client.ListBucketAnalyticsConfigurations(input)
}

func (a *S3Activities) ListBucketInventoryConfigurations(input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
    return a.client.ListBucketInventoryConfigurations(input)
}

func (a *S3Activities) ListBucketMetricsConfigurations(input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
    return a.client.ListBucketMetricsConfigurations(input)
}

func (a *S3Activities) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
    return a.client.ListBuckets(input)
}

func (a *S3Activities) ListMultipartUploads(input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
    return a.client.ListMultipartUploads(input)
}

func (a *S3Activities) ListObjectVersions(input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
    return a.client.ListObjectVersions(input)
}

func (a *S3Activities) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
    return a.client.ListObjects(input)
}

func (a *S3Activities) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
    return a.client.ListObjectsV2(input)
}

func (a *S3Activities) ListParts(input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
    return a.client.ListParts(input)
}

func (a *S3Activities) PutBucketAccelerateConfiguration(input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
    return a.client.PutBucketAccelerateConfiguration(input)
}

func (a *S3Activities) PutBucketAcl(input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
    return a.client.PutBucketAcl(input)
}

func (a *S3Activities) PutBucketAnalyticsConfiguration(input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
    return a.client.PutBucketAnalyticsConfiguration(input)
}

func (a *S3Activities) PutBucketCors(input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
    return a.client.PutBucketCors(input)
}

func (a *S3Activities) PutBucketEncryption(input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
    return a.client.PutBucketEncryption(input)
}

func (a *S3Activities) PutBucketInventoryConfiguration(input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
    return a.client.PutBucketInventoryConfiguration(input)
}

func (a *S3Activities) PutBucketLifecycle(input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
    return a.client.PutBucketLifecycle(input)
}

func (a *S3Activities) PutBucketLifecycleConfiguration(input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
    return a.client.PutBucketLifecycleConfiguration(input)
}

func (a *S3Activities) PutBucketLogging(input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
    return a.client.PutBucketLogging(input)
}

func (a *S3Activities) PutBucketMetricsConfiguration(input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
    return a.client.PutBucketMetricsConfiguration(input)
}

func (a *S3Activities) PutBucketNotification(input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
    return a.client.PutBucketNotification(input)
}

func (a *S3Activities) PutBucketNotificationConfiguration(input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
    return a.client.PutBucketNotificationConfiguration(input)
}

func (a *S3Activities) PutBucketPolicy(input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
    return a.client.PutBucketPolicy(input)
}

func (a *S3Activities) PutBucketReplication(input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
    return a.client.PutBucketReplication(input)
}

func (a *S3Activities) PutBucketRequestPayment(input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
    return a.client.PutBucketRequestPayment(input)
}

func (a *S3Activities) PutBucketTagging(input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
    return a.client.PutBucketTagging(input)
}

func (a *S3Activities) PutBucketVersioning(input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
    return a.client.PutBucketVersioning(input)
}

func (a *S3Activities) PutBucketWebsite(input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
    return a.client.PutBucketWebsite(input)
}

func (a *S3Activities) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
    return a.client.PutObject(input)
}

func (a *S3Activities) PutObjectAcl(input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
    return a.client.PutObjectAcl(input)
}

func (a *S3Activities) PutObjectLegalHold(input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
    return a.client.PutObjectLegalHold(input)
}

func (a *S3Activities) PutObjectLockConfiguration(input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
    return a.client.PutObjectLockConfiguration(input)
}

func (a *S3Activities) PutObjectRetention(input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
    return a.client.PutObjectRetention(input)
}

func (a *S3Activities) PutObjectTagging(input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
    return a.client.PutObjectTagging(input)
}

func (a *S3Activities) PutPublicAccessBlock(input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
    return a.client.PutPublicAccessBlock(input)
}

func (a *S3Activities) RestoreObject(input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
    return a.client.RestoreObject(input)
}

func (a *S3Activities) SelectObjectContent(input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
    return a.client.SelectObjectContent(input)
}

func (a *S3Activities) UploadPart(input *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
    return a.client.UploadPart(input)
}

func (a *S3Activities) UploadPartCopy(input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
    return a.client.UploadPartCopy(input)
}

func (a *S3Activities) WaitUntilBucketExists(input *s3.HeadBucketInput) error {
    return a.client.WaitUntilBucketExists(input)
}

func (a *S3Activities) WaitUntilBucketNotExists(input *s3.HeadBucketInput) error {
    return a.client.WaitUntilBucketNotExists(input)
}

func (a *S3Activities) WaitUntilObjectExists(input *s3.HeadObjectInput) error {
    return a.client.WaitUntilObjectExists(input)
}

func (a *S3Activities) WaitUntilObjectNotExists(input *s3.HeadObjectInput) error {
    return a.client.WaitUntilObjectNotExists(input)
}
