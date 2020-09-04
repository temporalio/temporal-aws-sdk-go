package awsclients

import (
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type LicenseManagerClient interface {
    CreateLicenseConfiguration(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error)
    CreateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) *LicensemanagerCreateLicenseConfigurationResult

    DeleteLicenseConfiguration(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error)
    DeleteLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) *LicensemanagerDeleteLicenseConfigurationResult

    GetLicenseConfiguration(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error)
    GetLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) *LicensemanagerGetLicenseConfigurationResult

    GetServiceSettings(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error)
    GetServiceSettingsAsync(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) *LicensemanagerGetServiceSettingsResult

    ListAssociationsForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
    ListAssociationsForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) *LicensemanagerListAssociationsForLicenseConfigurationResult

    ListFailuresForLicenseConfigurationOperations(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
    ListFailuresForLicenseConfigurationOperationsAsync(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) *LicensemanagerListFailuresForLicenseConfigurationOperationsResult

    ListLicenseConfigurations(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error)
    ListLicenseConfigurationsAsync(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) *LicensemanagerListLicenseConfigurationsResult

    ListLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
    ListLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) *LicensemanagerListLicenseSpecificationsForResourceResult

    ListResourceInventory(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error)
    ListResourceInventoryAsync(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) *LicensemanagerListResourceInventoryResult

    ListTagsForResource(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) *LicensemanagerListTagsForResourceResult

    ListUsageForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
    ListUsageForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) *LicensemanagerListUsageForLicenseConfigurationResult

    TagResource(ctx workflow.Context, input *licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *licensemanager.TagResourceInput) *LicensemanagerTagResourceResult

    UntagResource(ctx workflow.Context, input *licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *licensemanager.UntagResourceInput) *LicensemanagerUntagResourceResult

    UpdateLicenseConfiguration(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error)
    UpdateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) *LicensemanagerUpdateLicenseConfigurationResult

    UpdateLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
    UpdateLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) *LicensemanagerUpdateLicenseSpecificationsForResourceResult

    UpdateServiceSettings(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error)
    UpdateServiceSettingsAsync(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) *LicensemanagerUpdateServiceSettingsResult
}
type LicensemanagerCreateLicenseConfigurationResult struct {
	Result workflow.Future
}

func (r *LicensemanagerCreateLicenseConfigurationResult) Get(ctx workflow.Context) (*licensemanager.CreateLicenseConfigurationOutput, error) {
    var output licensemanager.CreateLicenseConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerDeleteLicenseConfigurationResult struct {
	Result workflow.Future
}

func (r *LicensemanagerDeleteLicenseConfigurationResult) Get(ctx workflow.Context) (*licensemanager.DeleteLicenseConfigurationOutput, error) {
    var output licensemanager.DeleteLicenseConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerGetLicenseConfigurationResult struct {
	Result workflow.Future
}

func (r *LicensemanagerGetLicenseConfigurationResult) Get(ctx workflow.Context) (*licensemanager.GetLicenseConfigurationOutput, error) {
    var output licensemanager.GetLicenseConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerGetServiceSettingsResult struct {
	Result workflow.Future
}

func (r *LicensemanagerGetServiceSettingsResult) Get(ctx workflow.Context) (*licensemanager.GetServiceSettingsOutput, error) {
    var output licensemanager.GetServiceSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListAssociationsForLicenseConfigurationResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListAssociationsForLicenseConfigurationResult) Get(ctx workflow.Context) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error) {
    var output licensemanager.ListAssociationsForLicenseConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListFailuresForLicenseConfigurationOperationsResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListFailuresForLicenseConfigurationOperationsResult) Get(ctx workflow.Context) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error) {
    var output licensemanager.ListFailuresForLicenseConfigurationOperationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListLicenseConfigurationsResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListLicenseConfigurationsResult) Get(ctx workflow.Context) (*licensemanager.ListLicenseConfigurationsOutput, error) {
    var output licensemanager.ListLicenseConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListLicenseSpecificationsForResourceResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListLicenseSpecificationsForResourceResult) Get(ctx workflow.Context) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error) {
    var output licensemanager.ListLicenseSpecificationsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListResourceInventoryResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListResourceInventoryResult) Get(ctx workflow.Context) (*licensemanager.ListResourceInventoryOutput, error) {
    var output licensemanager.ListResourceInventoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListTagsForResourceResult) Get(ctx workflow.Context) (*licensemanager.ListTagsForResourceOutput, error) {
    var output licensemanager.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerListUsageForLicenseConfigurationResult struct {
	Result workflow.Future
}

func (r *LicensemanagerListUsageForLicenseConfigurationResult) Get(ctx workflow.Context) (*licensemanager.ListUsageForLicenseConfigurationOutput, error) {
    var output licensemanager.ListUsageForLicenseConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerTagResourceResult struct {
	Result workflow.Future
}

func (r *LicensemanagerTagResourceResult) Get(ctx workflow.Context) (*licensemanager.TagResourceOutput, error) {
    var output licensemanager.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerUntagResourceResult struct {
	Result workflow.Future
}

func (r *LicensemanagerUntagResourceResult) Get(ctx workflow.Context) (*licensemanager.UntagResourceOutput, error) {
    var output licensemanager.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerUpdateLicenseConfigurationResult struct {
	Result workflow.Future
}

func (r *LicensemanagerUpdateLicenseConfigurationResult) Get(ctx workflow.Context) (*licensemanager.UpdateLicenseConfigurationOutput, error) {
    var output licensemanager.UpdateLicenseConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerUpdateLicenseSpecificationsForResourceResult struct {
	Result workflow.Future
}

func (r *LicensemanagerUpdateLicenseSpecificationsForResourceResult) Get(ctx workflow.Context) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error) {
    var output licensemanager.UpdateLicenseSpecificationsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type LicensemanagerUpdateServiceSettingsResult struct {
	Result workflow.Future
}

func (r *LicensemanagerUpdateServiceSettingsResult) Get(ctx workflow.Context) (*licensemanager.UpdateServiceSettingsOutput, error) {
    var output licensemanager.UpdateServiceSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type LicenseManagerStub struct {
    activities awsactivities.LicenseManagerActivities
}

func NewLicenseManagerStub() LicenseManagerClient {
    return &LicenseManagerStub{}
}
func (a *LicenseManagerStub) CreateLicenseConfiguration(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error) {
    var output licensemanager.CreateLicenseConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLicenseConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) CreateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.CreateLicenseConfigurationInput) *LicensemanagerCreateLicenseConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLicenseConfiguration, input)
    return &LicensemanagerCreateLicenseConfigurationResult{Result: future}
}
func (a *LicenseManagerStub) DeleteLicenseConfiguration(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error) {
    var output licensemanager.DeleteLicenseConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLicenseConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) DeleteLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.DeleteLicenseConfigurationInput) *LicensemanagerDeleteLicenseConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLicenseConfiguration, input)
    return &LicensemanagerDeleteLicenseConfigurationResult{Result: future}
}
func (a *LicenseManagerStub) GetLicenseConfiguration(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error) {
    var output licensemanager.GetLicenseConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLicenseConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) GetLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.GetLicenseConfigurationInput) *LicensemanagerGetLicenseConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLicenseConfiguration, input)
    return &LicensemanagerGetLicenseConfigurationResult{Result: future}
}
func (a *LicenseManagerStub) GetServiceSettings(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error) {
    var output licensemanager.GetServiceSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) GetServiceSettingsAsync(ctx workflow.Context, input *licensemanager.GetServiceSettingsInput) *LicensemanagerGetServiceSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetServiceSettings, input)
    return &LicensemanagerGetServiceSettingsResult{Result: future}
}
func (a *LicenseManagerStub) ListAssociationsForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error) {
    var output licensemanager.ListAssociationsForLicenseConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAssociationsForLicenseConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListAssociationsForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) *LicensemanagerListAssociationsForLicenseConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAssociationsForLicenseConfiguration, input)
    return &LicensemanagerListAssociationsForLicenseConfigurationResult{Result: future}
}
func (a *LicenseManagerStub) ListFailuresForLicenseConfigurationOperations(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error) {
    var output licensemanager.ListFailuresForLicenseConfigurationOperationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFailuresForLicenseConfigurationOperations, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListFailuresForLicenseConfigurationOperationsAsync(ctx workflow.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) *LicensemanagerListFailuresForLicenseConfigurationOperationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListFailuresForLicenseConfigurationOperations, input)
    return &LicensemanagerListFailuresForLicenseConfigurationOperationsResult{Result: future}
}
func (a *LicenseManagerStub) ListLicenseConfigurations(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error) {
    var output licensemanager.ListLicenseConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLicenseConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListLicenseConfigurationsAsync(ctx workflow.Context, input *licensemanager.ListLicenseConfigurationsInput) *LicensemanagerListLicenseConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLicenseConfigurations, input)
    return &LicensemanagerListLicenseConfigurationsResult{Result: future}
}
func (a *LicenseManagerStub) ListLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error) {
    var output licensemanager.ListLicenseSpecificationsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLicenseSpecificationsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) *LicensemanagerListLicenseSpecificationsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLicenseSpecificationsForResource, input)
    return &LicensemanagerListLicenseSpecificationsForResourceResult{Result: future}
}
func (a *LicenseManagerStub) ListResourceInventory(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error) {
    var output licensemanager.ListResourceInventoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceInventory, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListResourceInventoryAsync(ctx workflow.Context, input *licensemanager.ListResourceInventoryInput) *LicensemanagerListResourceInventoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResourceInventory, input)
    return &LicensemanagerListResourceInventoryResult{Result: future}
}
func (a *LicenseManagerStub) ListTagsForResource(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error) {
    var output licensemanager.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListTagsForResourceAsync(ctx workflow.Context, input *licensemanager.ListTagsForResourceInput) *LicensemanagerListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &LicensemanagerListTagsForResourceResult{Result: future}
}
func (a *LicenseManagerStub) ListUsageForLicenseConfiguration(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error) {
    var output licensemanager.ListUsageForLicenseConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsageForLicenseConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) ListUsageForLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) *LicensemanagerListUsageForLicenseConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUsageForLicenseConfiguration, input)
    return &LicensemanagerListUsageForLicenseConfigurationResult{Result: future}
}
func (a *LicenseManagerStub) TagResource(ctx workflow.Context, input *licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error) {
    var output licensemanager.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) TagResourceAsync(ctx workflow.Context, input *licensemanager.TagResourceInput) *LicensemanagerTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &LicensemanagerTagResourceResult{Result: future}
}
func (a *LicenseManagerStub) UntagResource(ctx workflow.Context, input *licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error) {
    var output licensemanager.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) UntagResourceAsync(ctx workflow.Context, input *licensemanager.UntagResourceInput) *LicensemanagerUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &LicensemanagerUntagResourceResult{Result: future}
}
func (a *LicenseManagerStub) UpdateLicenseConfiguration(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error) {
    var output licensemanager.UpdateLicenseConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLicenseConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) UpdateLicenseConfigurationAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseConfigurationInput) *LicensemanagerUpdateLicenseConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateLicenseConfiguration, input)
    return &LicensemanagerUpdateLicenseConfigurationResult{Result: future}
}
func (a *LicenseManagerStub) UpdateLicenseSpecificationsForResource(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error) {
    var output licensemanager.UpdateLicenseSpecificationsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLicenseSpecificationsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) UpdateLicenseSpecificationsForResourceAsync(ctx workflow.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) *LicensemanagerUpdateLicenseSpecificationsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateLicenseSpecificationsForResource, input)
    return &LicensemanagerUpdateLicenseSpecificationsForResourceResult{Result: future}
}
func (a *LicenseManagerStub) UpdateServiceSettings(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error) {
    var output licensemanager.UpdateServiceSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *LicenseManagerStub) UpdateServiceSettingsAsync(ctx workflow.Context, input *licensemanager.UpdateServiceSettingsInput) *LicensemanagerUpdateServiceSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceSettings, input)
    return &LicensemanagerUpdateServiceSettingsResult{Result: future}
}
