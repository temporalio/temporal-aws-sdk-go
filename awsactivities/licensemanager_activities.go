package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/licensemanager/licensemanageriface"
)

type LicenseManagerActivities struct {
	client licensemanageriface.LicenseManagerAPI
}

func NewLicenseManagerActivities(session *session.Session, config ...*aws.Config) *LicenseManagerActivities {
	client := licensemanager.New(session, config...)
	return &LicenseManagerActivities{client: client}
}

func (a *LicenseManagerActivities) CreateLicenseConfiguration(input *licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error) {
	return a.client.CreateLicenseConfiguration(input)
}

func (a *LicenseManagerActivities) DeleteLicenseConfiguration(input *licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error) {
	return a.client.DeleteLicenseConfiguration(input)
}

func (a *LicenseManagerActivities) GetLicenseConfiguration(input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error) {
	return a.client.GetLicenseConfiguration(input)
}

func (a *LicenseManagerActivities) GetServiceSettings(input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error) {
	return a.client.GetServiceSettings(input)
}

func (a *LicenseManagerActivities) ListAssociationsForLicenseConfiguration(input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error) {
	return a.client.ListAssociationsForLicenseConfiguration(input)
}

func (a *LicenseManagerActivities) ListFailuresForLicenseConfigurationOperations(input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error) {
	return a.client.ListFailuresForLicenseConfigurationOperations(input)
}

func (a *LicenseManagerActivities) ListLicenseConfigurations(input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error) {
	return a.client.ListLicenseConfigurations(input)
}

func (a *LicenseManagerActivities) ListLicenseSpecificationsForResource(input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error) {
	return a.client.ListLicenseSpecificationsForResource(input)
}

func (a *LicenseManagerActivities) ListResourceInventory(input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error) {
	return a.client.ListResourceInventory(input)
}

func (a *LicenseManagerActivities) ListTagsForResource(input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *LicenseManagerActivities) ListUsageForLicenseConfiguration(input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error) {
	return a.client.ListUsageForLicenseConfiguration(input)
}

func (a *LicenseManagerActivities) TagResource(input *licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *LicenseManagerActivities) UntagResource(input *licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *LicenseManagerActivities) UpdateLicenseConfiguration(input *licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error) {
	return a.client.UpdateLicenseConfiguration(input)
}

func (a *LicenseManagerActivities) UpdateLicenseSpecificationsForResource(input *licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error) {
	return a.client.UpdateLicenseSpecificationsForResource(input)
}

func (a *LicenseManagerActivities) UpdateServiceSettings(input *licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error) {
	return a.client.UpdateServiceSettings(input)
}
