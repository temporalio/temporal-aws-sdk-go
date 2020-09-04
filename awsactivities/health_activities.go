package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/health/healthiface"
)

type HealthActivities struct {
	client healthiface.HealthAPI
}

func NewHealthActivities(client healthiface.HealthAPI) *HealthActivities {
    return &HealthActivities{client: client}
}


func (a *HealthActivities) DescribeAffectedAccountsForOrganization(input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
    return a.client.DescribeAffectedAccountsForOrganization(input)
}



func (a *HealthActivities) DescribeAffectedEntities(input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
    return a.client.DescribeAffectedEntities(input)
}



func (a *HealthActivities) DescribeAffectedEntitiesForOrganization(input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
    return a.client.DescribeAffectedEntitiesForOrganization(input)
}



func (a *HealthActivities) DescribeEntityAggregates(input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
    return a.client.DescribeEntityAggregates(input)
}



func (a *HealthActivities) DescribeEventAggregates(input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
    return a.client.DescribeEventAggregates(input)
}



func (a *HealthActivities) DescribeEventDetails(input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
    return a.client.DescribeEventDetails(input)
}



func (a *HealthActivities) DescribeEventDetailsForOrganization(input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error) {
    return a.client.DescribeEventDetailsForOrganization(input)
}



func (a *HealthActivities) DescribeEventTypes(input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
    return a.client.DescribeEventTypes(input)
}



func (a *HealthActivities) DescribeEvents(input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
    return a.client.DescribeEvents(input)
}



func (a *HealthActivities) DescribeEventsForOrganization(input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error) {
    return a.client.DescribeEventsForOrganization(input)
}



func (a *HealthActivities) DescribeHealthServiceStatusForOrganization(input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
    return a.client.DescribeHealthServiceStatusForOrganization(input)
}



func (a *HealthActivities) DisableHealthServiceAccessForOrganization(input *health.DisableHealthServiceAccessForOrganizationInput) (*health.DisableHealthServiceAccessForOrganizationOutput, error) {
    return a.client.DisableHealthServiceAccessForOrganization(input)
}



func (a *HealthActivities) EnableHealthServiceAccessForOrganization(input *health.EnableHealthServiceAccessForOrganizationInput) (*health.EnableHealthServiceAccessForOrganizationOutput, error) {
    return a.client.EnableHealthServiceAccessForOrganization(input)
}

