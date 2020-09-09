
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/shield/shieldiface"
)

type ShieldActivities struct {
	client shieldiface.ShieldAPI
}

func NewShieldActivities(session *session.Session, config... *aws.Config) *ShieldActivities {
    client := shield.New(session, config...)
    return &ShieldActivities{client: client}
}

func (a *ShieldActivities) AssociateDRTLogBucket(input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error) {
    return a.client.AssociateDRTLogBucket(input)
}

func (a *ShieldActivities) AssociateDRTRole(input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error) {
    return a.client.AssociateDRTRole(input)
}

func (a *ShieldActivities) AssociateHealthCheck(input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error) {
    return a.client.AssociateHealthCheck(input)
}

func (a *ShieldActivities) AssociateProactiveEngagementDetails(input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
    return a.client.AssociateProactiveEngagementDetails(input)
}

func (a *ShieldActivities) CreateProtection(input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error) {
    return a.client.CreateProtection(input)
}

func (a *ShieldActivities) CreateSubscription(input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error) {
    return a.client.CreateSubscription(input)
}

func (a *ShieldActivities) DeleteProtection(input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error) {
    return a.client.DeleteProtection(input)
}

func (a *ShieldActivities) DeleteSubscription(input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error) {
    return a.client.DeleteSubscription(input)
}

func (a *ShieldActivities) DescribeAttack(input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error) {
    return a.client.DescribeAttack(input)
}

func (a *ShieldActivities) DescribeDRTAccess(input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error) {
    return a.client.DescribeDRTAccess(input)
}

func (a *ShieldActivities) DescribeEmergencyContactSettings(input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error) {
    return a.client.DescribeEmergencyContactSettings(input)
}

func (a *ShieldActivities) DescribeProtection(input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error) {
    return a.client.DescribeProtection(input)
}

func (a *ShieldActivities) DescribeSubscription(input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error) {
    return a.client.DescribeSubscription(input)
}

func (a *ShieldActivities) DisableProactiveEngagement(input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error) {
    return a.client.DisableProactiveEngagement(input)
}

func (a *ShieldActivities) DisassociateDRTLogBucket(input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error) {
    return a.client.DisassociateDRTLogBucket(input)
}

func (a *ShieldActivities) DisassociateDRTRole(input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error) {
    return a.client.DisassociateDRTRole(input)
}

func (a *ShieldActivities) DisassociateHealthCheck(input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error) {
    return a.client.DisassociateHealthCheck(input)
}

func (a *ShieldActivities) EnableProactiveEngagement(input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error) {
    return a.client.EnableProactiveEngagement(input)
}

func (a *ShieldActivities) GetSubscriptionState(input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error) {
    return a.client.GetSubscriptionState(input)
}

func (a *ShieldActivities) ListAttacks(input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error) {
    return a.client.ListAttacks(input)
}

func (a *ShieldActivities) ListProtections(input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error) {
    return a.client.ListProtections(input)
}

func (a *ShieldActivities) UpdateEmergencyContactSettings(input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error) {
    return a.client.UpdateEmergencyContactSettings(input)
}

func (a *ShieldActivities) UpdateSubscription(input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error) {
    return a.client.UpdateSubscription(input)
}
