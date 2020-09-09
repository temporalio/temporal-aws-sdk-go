package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface"
)

type MigrationHubConfigActivities struct {
	client migrationhubconfigiface.MigrationHubConfigAPI
}

func NewMigrationHubConfigActivities(session *session.Session, config ...*aws.Config) *MigrationHubConfigActivities {
	client := migrationhubconfig.New(session, config...)
	return &MigrationHubConfigActivities{client: client}
}

func (a *MigrationHubConfigActivities) CreateHomeRegionControl(input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	return a.client.CreateHomeRegionControl(input)
}

func (a *MigrationHubConfigActivities) DescribeHomeRegionControls(input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	return a.client.DescribeHomeRegionControls(input)
}

func (a *MigrationHubConfigActivities) GetHomeRegion(input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	return a.client.GetHomeRegion(input)
}
