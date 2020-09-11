
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/backup/backupiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type BackupActivities struct {
    client backupiface.BackupAPI
}

func NewBackupActivities(session *session.Session, config ...*aws.Config) *BackupActivities {
    client := backup.New(session, config...)
    return &BackupActivities{client: client}
}

func (a *BackupActivities) CreateBackupPlan(ctx context.Context, input *backup.CreateBackupPlanInput) (*backup.CreateBackupPlanOutput, error) {
    return a.client.CreateBackupPlanWithContext(ctx, input)
}

func (a *BackupActivities) CreateBackupSelection(ctx context.Context, input *backup.CreateBackupSelectionInput) (*backup.CreateBackupSelectionOutput, error) {
    return a.client.CreateBackupSelectionWithContext(ctx, input)
}

func (a *BackupActivities) CreateBackupVault(ctx context.Context, input *backup.CreateBackupVaultInput) (*backup.CreateBackupVaultOutput, error) {
    return a.client.CreateBackupVaultWithContext(ctx, input)
}

func (a *BackupActivities) DeleteBackupPlan(ctx context.Context, input *backup.DeleteBackupPlanInput) (*backup.DeleteBackupPlanOutput, error) {
    return a.client.DeleteBackupPlanWithContext(ctx, input)
}

func (a *BackupActivities) DeleteBackupSelection(ctx context.Context, input *backup.DeleteBackupSelectionInput) (*backup.DeleteBackupSelectionOutput, error) {
    return a.client.DeleteBackupSelectionWithContext(ctx, input)
}

func (a *BackupActivities) DeleteBackupVault(ctx context.Context, input *backup.DeleteBackupVaultInput) (*backup.DeleteBackupVaultOutput, error) {
    return a.client.DeleteBackupVaultWithContext(ctx, input)
}

func (a *BackupActivities) DeleteBackupVaultAccessPolicy(ctx context.Context, input *backup.DeleteBackupVaultAccessPolicyInput) (*backup.DeleteBackupVaultAccessPolicyOutput, error) {
    return a.client.DeleteBackupVaultAccessPolicyWithContext(ctx, input)
}

func (a *BackupActivities) DeleteBackupVaultNotifications(ctx context.Context, input *backup.DeleteBackupVaultNotificationsInput) (*backup.DeleteBackupVaultNotificationsOutput, error) {
    return a.client.DeleteBackupVaultNotificationsWithContext(ctx, input)
}

func (a *BackupActivities) DeleteRecoveryPoint(ctx context.Context, input *backup.DeleteRecoveryPointInput) (*backup.DeleteRecoveryPointOutput, error) {
    return a.client.DeleteRecoveryPointWithContext(ctx, input)
}

func (a *BackupActivities) DescribeBackupJob(ctx context.Context, input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error) {
    return a.client.DescribeBackupJobWithContext(ctx, input)
}

func (a *BackupActivities) DescribeBackupVault(ctx context.Context, input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error) {
    return a.client.DescribeBackupVaultWithContext(ctx, input)
}

func (a *BackupActivities) DescribeCopyJob(ctx context.Context, input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error) {
    return a.client.DescribeCopyJobWithContext(ctx, input)
}

func (a *BackupActivities) DescribeProtectedResource(ctx context.Context, input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error) {
    return a.client.DescribeProtectedResourceWithContext(ctx, input)
}

func (a *BackupActivities) DescribeRecoveryPoint(ctx context.Context, input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error) {
    return a.client.DescribeRecoveryPointWithContext(ctx, input)
}

func (a *BackupActivities) DescribeRegionSettings(ctx context.Context, input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error) {
    return a.client.DescribeRegionSettingsWithContext(ctx, input)
}

func (a *BackupActivities) DescribeRestoreJob(ctx context.Context, input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error) {
    return a.client.DescribeRestoreJobWithContext(ctx, input)
}

func (a *BackupActivities) ExportBackupPlanTemplate(ctx context.Context, input *backup.ExportBackupPlanTemplateInput) (*backup.ExportBackupPlanTemplateOutput, error) {
    return a.client.ExportBackupPlanTemplateWithContext(ctx, input)
}

func (a *BackupActivities) GetBackupPlan(ctx context.Context, input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error) {
    return a.client.GetBackupPlanWithContext(ctx, input)
}

func (a *BackupActivities) GetBackupPlanFromJSON(ctx context.Context, input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error) {
    return a.client.GetBackupPlanFromJSONWithContext(ctx, input)
}

func (a *BackupActivities) GetBackupPlanFromTemplate(ctx context.Context, input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error) {
    return a.client.GetBackupPlanFromTemplateWithContext(ctx, input)
}

func (a *BackupActivities) GetBackupSelection(ctx context.Context, input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error) {
    return a.client.GetBackupSelectionWithContext(ctx, input)
}

func (a *BackupActivities) GetBackupVaultAccessPolicy(ctx context.Context, input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error) {
    return a.client.GetBackupVaultAccessPolicyWithContext(ctx, input)
}

func (a *BackupActivities) GetBackupVaultNotifications(ctx context.Context, input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error) {
    return a.client.GetBackupVaultNotificationsWithContext(ctx, input)
}

func (a *BackupActivities) GetRecoveryPointRestoreMetadata(ctx context.Context, input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
    return a.client.GetRecoveryPointRestoreMetadataWithContext(ctx, input)
}

func (a *BackupActivities) GetSupportedResourceTypes(ctx context.Context, input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error) {
    return a.client.GetSupportedResourceTypesWithContext(ctx, input)
}

func (a *BackupActivities) ListBackupJobs(ctx context.Context, input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error) {
    return a.client.ListBackupJobsWithContext(ctx, input)
}

func (a *BackupActivities) ListBackupPlanTemplates(ctx context.Context, input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error) {
    return a.client.ListBackupPlanTemplatesWithContext(ctx, input)
}

func (a *BackupActivities) ListBackupPlanVersions(ctx context.Context, input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error) {
    return a.client.ListBackupPlanVersionsWithContext(ctx, input)
}

func (a *BackupActivities) ListBackupPlans(ctx context.Context, input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error) {
    return a.client.ListBackupPlansWithContext(ctx, input)
}

func (a *BackupActivities) ListBackupSelections(ctx context.Context, input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error) {
    return a.client.ListBackupSelectionsWithContext(ctx, input)
}

func (a *BackupActivities) ListBackupVaults(ctx context.Context, input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error) {
    return a.client.ListBackupVaultsWithContext(ctx, input)
}

func (a *BackupActivities) ListCopyJobs(ctx context.Context, input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error) {
    return a.client.ListCopyJobsWithContext(ctx, input)
}

func (a *BackupActivities) ListProtectedResources(ctx context.Context, input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error) {
    return a.client.ListProtectedResourcesWithContext(ctx, input)
}

func (a *BackupActivities) ListRecoveryPointsByBackupVault(ctx context.Context, input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
    return a.client.ListRecoveryPointsByBackupVaultWithContext(ctx, input)
}

func (a *BackupActivities) ListRecoveryPointsByResource(ctx context.Context, input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error) {
    return a.client.ListRecoveryPointsByResourceWithContext(ctx, input)
}

func (a *BackupActivities) ListRestoreJobs(ctx context.Context, input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error) {
    return a.client.ListRestoreJobsWithContext(ctx, input)
}

func (a *BackupActivities) ListTags(ctx context.Context, input *backup.ListTagsInput) (*backup.ListTagsOutput, error) {
    return a.client.ListTagsWithContext(ctx, input)
}

func (a *BackupActivities) PutBackupVaultAccessPolicy(ctx context.Context, input *backup.PutBackupVaultAccessPolicyInput) (*backup.PutBackupVaultAccessPolicyOutput, error) {
    return a.client.PutBackupVaultAccessPolicyWithContext(ctx, input)
}

func (a *BackupActivities) PutBackupVaultNotifications(ctx context.Context, input *backup.PutBackupVaultNotificationsInput) (*backup.PutBackupVaultNotificationsOutput, error) {
    return a.client.PutBackupVaultNotificationsWithContext(ctx, input)
}

func (a *BackupActivities) StartBackupJob(ctx context.Context, input *backup.StartBackupJobInput) (*backup.StartBackupJobOutput, error) {
    return a.client.StartBackupJobWithContext(ctx, input)
}

func (a *BackupActivities) StartCopyJob(ctx context.Context, input *backup.StartCopyJobInput) (*backup.StartCopyJobOutput, error) {
    return a.client.StartCopyJobWithContext(ctx, input)
}

func (a *BackupActivities) StartRestoreJob(ctx context.Context, input *backup.StartRestoreJobInput) (*backup.StartRestoreJobOutput, error) {
    return a.client.StartRestoreJobWithContext(ctx, input)
}

func (a *BackupActivities) StopBackupJob(ctx context.Context, input *backup.StopBackupJobInput) (*backup.StopBackupJobOutput, error) {
    return a.client.StopBackupJobWithContext(ctx, input)
}

func (a *BackupActivities) TagResource(ctx context.Context, input *backup.TagResourceInput) (*backup.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *BackupActivities) UntagResource(ctx context.Context, input *backup.UntagResourceInput) (*backup.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *BackupActivities) UpdateBackupPlan(ctx context.Context, input *backup.UpdateBackupPlanInput) (*backup.UpdateBackupPlanOutput, error) {
    return a.client.UpdateBackupPlanWithContext(ctx, input)
}

func (a *BackupActivities) UpdateRecoveryPointLifecycle(ctx context.Context, input *backup.UpdateRecoveryPointLifecycleInput) (*backup.UpdateRecoveryPointLifecycleOutput, error) {
    return a.client.UpdateRecoveryPointLifecycleWithContext(ctx, input)
}

func (a *BackupActivities) UpdateRegionSettings(ctx context.Context, input *backup.UpdateRegionSettingsInput) (*backup.UpdateRegionSettingsOutput, error) {
    return a.client.UpdateRegionSettingsWithContext(ctx, input)
}
