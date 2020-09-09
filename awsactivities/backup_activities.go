
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/backup/backupiface"
)

type BackupActivities struct {
	client backupiface.BackupAPI
}

func NewBackupActivities(session *session.Session, config... *aws.Config) *BackupActivities {
    client := backup.New(session, config...)
    return &BackupActivities{client: client}
}

func (a *BackupActivities) CreateBackupPlan(input *backup.CreateBackupPlanInput) (*backup.CreateBackupPlanOutput, error) {
    return a.client.CreateBackupPlan(input)
}

func (a *BackupActivities) CreateBackupSelection(input *backup.CreateBackupSelectionInput) (*backup.CreateBackupSelectionOutput, error) {
    return a.client.CreateBackupSelection(input)
}

func (a *BackupActivities) CreateBackupVault(input *backup.CreateBackupVaultInput) (*backup.CreateBackupVaultOutput, error) {
    return a.client.CreateBackupVault(input)
}

func (a *BackupActivities) DeleteBackupPlan(input *backup.DeleteBackupPlanInput) (*backup.DeleteBackupPlanOutput, error) {
    return a.client.DeleteBackupPlan(input)
}

func (a *BackupActivities) DeleteBackupSelection(input *backup.DeleteBackupSelectionInput) (*backup.DeleteBackupSelectionOutput, error) {
    return a.client.DeleteBackupSelection(input)
}

func (a *BackupActivities) DeleteBackupVault(input *backup.DeleteBackupVaultInput) (*backup.DeleteBackupVaultOutput, error) {
    return a.client.DeleteBackupVault(input)
}

func (a *BackupActivities) DeleteBackupVaultAccessPolicy(input *backup.DeleteBackupVaultAccessPolicyInput) (*backup.DeleteBackupVaultAccessPolicyOutput, error) {
    return a.client.DeleteBackupVaultAccessPolicy(input)
}

func (a *BackupActivities) DeleteBackupVaultNotifications(input *backup.DeleteBackupVaultNotificationsInput) (*backup.DeleteBackupVaultNotificationsOutput, error) {
    return a.client.DeleteBackupVaultNotifications(input)
}

func (a *BackupActivities) DeleteRecoveryPoint(input *backup.DeleteRecoveryPointInput) (*backup.DeleteRecoveryPointOutput, error) {
    return a.client.DeleteRecoveryPoint(input)
}

func (a *BackupActivities) DescribeBackupJob(input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error) {
    return a.client.DescribeBackupJob(input)
}

func (a *BackupActivities) DescribeBackupVault(input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error) {
    return a.client.DescribeBackupVault(input)
}

func (a *BackupActivities) DescribeCopyJob(input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error) {
    return a.client.DescribeCopyJob(input)
}

func (a *BackupActivities) DescribeProtectedResource(input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error) {
    return a.client.DescribeProtectedResource(input)
}

func (a *BackupActivities) DescribeRecoveryPoint(input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error) {
    return a.client.DescribeRecoveryPoint(input)
}

func (a *BackupActivities) DescribeRegionSettings(input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error) {
    return a.client.DescribeRegionSettings(input)
}

func (a *BackupActivities) DescribeRestoreJob(input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error) {
    return a.client.DescribeRestoreJob(input)
}

func (a *BackupActivities) ExportBackupPlanTemplate(input *backup.ExportBackupPlanTemplateInput) (*backup.ExportBackupPlanTemplateOutput, error) {
    return a.client.ExportBackupPlanTemplate(input)
}

func (a *BackupActivities) GetBackupPlan(input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error) {
    return a.client.GetBackupPlan(input)
}

func (a *BackupActivities) GetBackupPlanFromJSON(input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error) {
    return a.client.GetBackupPlanFromJSON(input)
}

func (a *BackupActivities) GetBackupPlanFromTemplate(input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error) {
    return a.client.GetBackupPlanFromTemplate(input)
}

func (a *BackupActivities) GetBackupSelection(input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error) {
    return a.client.GetBackupSelection(input)
}

func (a *BackupActivities) GetBackupVaultAccessPolicy(input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error) {
    return a.client.GetBackupVaultAccessPolicy(input)
}

func (a *BackupActivities) GetBackupVaultNotifications(input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error) {
    return a.client.GetBackupVaultNotifications(input)
}

func (a *BackupActivities) GetRecoveryPointRestoreMetadata(input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
    return a.client.GetRecoveryPointRestoreMetadata(input)
}

func (a *BackupActivities) GetSupportedResourceTypes(input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error) {
    return a.client.GetSupportedResourceTypes(input)
}

func (a *BackupActivities) ListBackupJobs(input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error) {
    return a.client.ListBackupJobs(input)
}

func (a *BackupActivities) ListBackupPlanTemplates(input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error) {
    return a.client.ListBackupPlanTemplates(input)
}

func (a *BackupActivities) ListBackupPlanVersions(input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error) {
    return a.client.ListBackupPlanVersions(input)
}

func (a *BackupActivities) ListBackupPlans(input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error) {
    return a.client.ListBackupPlans(input)
}

func (a *BackupActivities) ListBackupSelections(input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error) {
    return a.client.ListBackupSelections(input)
}

func (a *BackupActivities) ListBackupVaults(input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error) {
    return a.client.ListBackupVaults(input)
}

func (a *BackupActivities) ListCopyJobs(input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error) {
    return a.client.ListCopyJobs(input)
}

func (a *BackupActivities) ListProtectedResources(input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error) {
    return a.client.ListProtectedResources(input)
}

func (a *BackupActivities) ListRecoveryPointsByBackupVault(input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
    return a.client.ListRecoveryPointsByBackupVault(input)
}

func (a *BackupActivities) ListRecoveryPointsByResource(input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error) {
    return a.client.ListRecoveryPointsByResource(input)
}

func (a *BackupActivities) ListRestoreJobs(input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error) {
    return a.client.ListRestoreJobs(input)
}

func (a *BackupActivities) ListTags(input *backup.ListTagsInput) (*backup.ListTagsOutput, error) {
    return a.client.ListTags(input)
}

func (a *BackupActivities) PutBackupVaultAccessPolicy(input *backup.PutBackupVaultAccessPolicyInput) (*backup.PutBackupVaultAccessPolicyOutput, error) {
    return a.client.PutBackupVaultAccessPolicy(input)
}

func (a *BackupActivities) PutBackupVaultNotifications(input *backup.PutBackupVaultNotificationsInput) (*backup.PutBackupVaultNotificationsOutput, error) {
    return a.client.PutBackupVaultNotifications(input)
}

func (a *BackupActivities) StartBackupJob(input *backup.StartBackupJobInput) (*backup.StartBackupJobOutput, error) {
    return a.client.StartBackupJob(input)
}

func (a *BackupActivities) StartCopyJob(input *backup.StartCopyJobInput) (*backup.StartCopyJobOutput, error) {
    return a.client.StartCopyJob(input)
}

func (a *BackupActivities) StartRestoreJob(input *backup.StartRestoreJobInput) (*backup.StartRestoreJobOutput, error) {
    return a.client.StartRestoreJob(input)
}

func (a *BackupActivities) StopBackupJob(input *backup.StopBackupJobInput) (*backup.StopBackupJobOutput, error) {
    return a.client.StopBackupJob(input)
}

func (a *BackupActivities) TagResource(input *backup.TagResourceInput) (*backup.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *BackupActivities) UntagResource(input *backup.UntagResourceInput) (*backup.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}

func (a *BackupActivities) UpdateBackupPlan(input *backup.UpdateBackupPlanInput) (*backup.UpdateBackupPlanOutput, error) {
    return a.client.UpdateBackupPlan(input)
}

func (a *BackupActivities) UpdateRecoveryPointLifecycle(input *backup.UpdateRecoveryPointLifecycleInput) (*backup.UpdateRecoveryPointLifecycleOutput, error) {
    return a.client.UpdateRecoveryPointLifecycle(input)
}

func (a *BackupActivities) UpdateRegionSettings(input *backup.UpdateRegionSettingsInput) (*backup.UpdateRegionSettingsOutput, error) {
    return a.client.UpdateRegionSettings(input)
}
