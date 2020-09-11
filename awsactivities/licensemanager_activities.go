package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/licensemanager/licensemanageriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type LicenseManagerActivities struct {
	client licensemanageriface.LicenseManagerAPI
}

func NewLicenseManagerActivities(session *session.Session, config ...*aws.Config) *LicenseManagerActivities {
	client := licensemanager.New(session, config...)
	return &LicenseManagerActivities{client: client}
}

func (a *LicenseManagerActivities) CreateLicenseConfiguration(ctx context.Context, input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error) {
	return a.client.CreateLicenseConfigurationWithContext(ctx, input)
}

func (a *LicenseManagerActivities) DeleteLicenseConfiguration(ctx context.Context, input *licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error) {
	return a.client.DeleteLicenseConfigurationWithContext(ctx, input)
}

func (a *LicenseManagerActivities) GetLicenseConfiguration(ctx context.Context, input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error) {
	return a.client.GetLicenseConfigurationWithContext(ctx, input)
}

func (a *LicenseManagerActivities) GetServiceSettings(ctx context.Context, input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error) {
	return a.client.GetServiceSettingsWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListAssociationsForLicenseConfiguration(ctx context.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error) {
	return a.client.ListAssociationsForLicenseConfigurationWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListFailuresForLicenseConfigurationOperations(ctx context.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error) {
	return a.client.ListFailuresForLicenseConfigurationOperationsWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListLicenseConfigurations(ctx context.Context, input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error) {
	return a.client.ListLicenseConfigurationsWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListLicenseSpecificationsForResource(ctx context.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error) {
	return a.client.ListLicenseSpecificationsForResourceWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListResourceInventory(ctx context.Context, input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error) {
	return a.client.ListResourceInventoryWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListTagsForResource(ctx context.Context, input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *LicenseManagerActivities) ListUsageForLicenseConfiguration(ctx context.Context, input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error) {
	return a.client.ListUsageForLicenseConfigurationWithContext(ctx, input)
}

func (a *LicenseManagerActivities) TagResource(ctx context.Context, input *licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *LicenseManagerActivities) UntagResource(ctx context.Context, input *licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *LicenseManagerActivities) UpdateLicenseConfiguration(ctx context.Context, input *licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error) {
	return a.client.UpdateLicenseConfigurationWithContext(ctx, input)
}

func (a *LicenseManagerActivities) UpdateLicenseSpecificationsForResource(ctx context.Context, input *licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error) {
	return a.client.UpdateLicenseSpecificationsForResourceWithContext(ctx, input)
}

func (a *LicenseManagerActivities) UpdateServiceSettings(ctx context.Context, input *licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error) {
	return a.client.UpdateServiceSettingsWithContext(ctx, input)
}
