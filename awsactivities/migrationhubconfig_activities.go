package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type MigrationHubConfigActivities struct {
	client migrationhubconfigiface.MigrationHubConfigAPI
}

func NewMigrationHubConfigActivities(session *session.Session, config ...*aws.Config) *MigrationHubConfigActivities {
	client := migrationhubconfig.New(session, config...)
	return &MigrationHubConfigActivities{client: client}
}

func (a *MigrationHubConfigActivities) CreateHomeRegionControl(ctx context.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	return a.client.CreateHomeRegionControlWithContext(ctx, input)
}

func (a *MigrationHubConfigActivities) DescribeHomeRegionControls(ctx context.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	return a.client.DescribeHomeRegionControlsWithContext(ctx, input)
}

func (a *MigrationHubConfigActivities) GetHomeRegion(ctx context.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	return a.client.GetHomeRegionWithContext(ctx, input)
}
