
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/health/healthiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type HealthActivities struct {
    client healthiface.HealthAPI
}

func NewHealthActivities(session *session.Session, config ...*aws.Config) *HealthActivities {
    client := health.New(session, config...)
    return &HealthActivities{client: client}
}

func (a *HealthActivities) DescribeAffectedAccountsForOrganization(ctx context.Context, input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
    return a.client.DescribeAffectedAccountsForOrganizationWithContext(ctx, input)
}

func (a *HealthActivities) DescribeAffectedEntities(ctx context.Context, input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
    return a.client.DescribeAffectedEntitiesWithContext(ctx, input)
}

func (a *HealthActivities) DescribeAffectedEntitiesForOrganization(ctx context.Context, input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
    return a.client.DescribeAffectedEntitiesForOrganizationWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEntityAggregates(ctx context.Context, input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
    return a.client.DescribeEntityAggregatesWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEventAggregates(ctx context.Context, input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
    return a.client.DescribeEventAggregatesWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEventDetails(ctx context.Context, input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
    return a.client.DescribeEventDetailsWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEventDetailsForOrganization(ctx context.Context, input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error) {
    return a.client.DescribeEventDetailsForOrganizationWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEventTypes(ctx context.Context, input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
    return a.client.DescribeEventTypesWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEvents(ctx context.Context, input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
    return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *HealthActivities) DescribeEventsForOrganization(ctx context.Context, input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error) {
    return a.client.DescribeEventsForOrganizationWithContext(ctx, input)
}

func (a *HealthActivities) DescribeHealthServiceStatusForOrganization(ctx context.Context, input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
    return a.client.DescribeHealthServiceStatusForOrganizationWithContext(ctx, input)
}

func (a *HealthActivities) DisableHealthServiceAccessForOrganization(ctx context.Context, input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
    return a.client.DisableHealthServiceAccessForOrganizationWithContext(ctx, input)
}

func (a *HealthActivities) EnableHealthServiceAccessForOrganization(ctx context.Context, input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
    return a.client.EnableHealthServiceAccessForOrganizationWithContext(ctx, input)
}
