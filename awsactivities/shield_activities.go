package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/shield/shieldiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ShieldActivities struct {
	client shieldiface.ShieldAPI
}

func NewShieldActivities(session *session.Session, config ...*aws.Config) *ShieldActivities {
	client := shield.New(session, config...)
	return &ShieldActivities{client: client}
}

func (a *ShieldActivities) AssociateDRTLogBucket(ctx context.Context, input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error) {
	return a.client.AssociateDRTLogBucketWithContext(ctx, input)
}

func (a *ShieldActivities) AssociateDRTRole(ctx context.Context, input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error) {
	return a.client.AssociateDRTRoleWithContext(ctx, input)
}

func (a *ShieldActivities) AssociateHealthCheck(ctx context.Context, input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error) {
	return a.client.AssociateHealthCheckWithContext(ctx, input)
}

func (a *ShieldActivities) AssociateProactiveEngagementDetails(ctx context.Context, input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
	return a.client.AssociateProactiveEngagementDetailsWithContext(ctx, input)
}

func (a *ShieldActivities) CreateProtection(ctx context.Context, input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error) {
	return a.client.CreateProtectionWithContext(ctx, input)
}

func (a *ShieldActivities) CreateSubscription(ctx context.Context, input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error) {
	return a.client.CreateSubscriptionWithContext(ctx, input)
}

func (a *ShieldActivities) DeleteProtection(ctx context.Context, input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error) {
	return a.client.DeleteProtectionWithContext(ctx, input)
}

func (a *ShieldActivities) DeleteSubscription(ctx context.Context, input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error) {
	return a.client.DeleteSubscriptionWithContext(ctx, input)
}

func (a *ShieldActivities) DescribeAttack(ctx context.Context, input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error) {
	return a.client.DescribeAttackWithContext(ctx, input)
}

func (a *ShieldActivities) DescribeDRTAccess(ctx context.Context, input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error) {
	return a.client.DescribeDRTAccessWithContext(ctx, input)
}

func (a *ShieldActivities) DescribeEmergencyContactSettings(ctx context.Context, input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	return a.client.DescribeEmergencyContactSettingsWithContext(ctx, input)
}

func (a *ShieldActivities) DescribeProtection(ctx context.Context, input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error) {
	return a.client.DescribeProtectionWithContext(ctx, input)
}

func (a *ShieldActivities) DescribeSubscription(ctx context.Context, input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error) {
	return a.client.DescribeSubscriptionWithContext(ctx, input)
}

func (a *ShieldActivities) DisableProactiveEngagement(ctx context.Context, input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error) {
	return a.client.DisableProactiveEngagementWithContext(ctx, input)
}

func (a *ShieldActivities) DisassociateDRTLogBucket(ctx context.Context, input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error) {
	return a.client.DisassociateDRTLogBucketWithContext(ctx, input)
}

func (a *ShieldActivities) DisassociateDRTRole(ctx context.Context, input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error) {
	return a.client.DisassociateDRTRoleWithContext(ctx, input)
}

func (a *ShieldActivities) DisassociateHealthCheck(ctx context.Context, input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error) {
	return a.client.DisassociateHealthCheckWithContext(ctx, input)
}

func (a *ShieldActivities) EnableProactiveEngagement(ctx context.Context, input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error) {
	return a.client.EnableProactiveEngagementWithContext(ctx, input)
}

func (a *ShieldActivities) GetSubscriptionState(ctx context.Context, input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error) {
	return a.client.GetSubscriptionStateWithContext(ctx, input)
}

func (a *ShieldActivities) ListAttacks(ctx context.Context, input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error) {
	return a.client.ListAttacksWithContext(ctx, input)
}

func (a *ShieldActivities) ListProtections(ctx context.Context, input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error) {
	return a.client.ListProtectionsWithContext(ctx, input)
}

func (a *ShieldActivities) UpdateEmergencyContactSettings(ctx context.Context, input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error) {
	return a.client.UpdateEmergencyContactSettingsWithContext(ctx, input)
}

func (a *ShieldActivities) UpdateSubscription(ctx context.Context, input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error) {
	return a.client.UpdateSubscriptionWithContext(ctx, input)
}
