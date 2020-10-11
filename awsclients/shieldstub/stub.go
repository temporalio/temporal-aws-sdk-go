// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package shieldstub

import (
	"github.com/aws/aws-sdk-go/service/shield"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type ShieldAssociateDRTLogBucketFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldAssociateDRTLogBucketFuture) Get(ctx workflow.Context) (*shield.AssociateDRTLogBucketOutput, error) {
	var output shield.AssociateDRTLogBucketOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldAssociateDRTRoleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldAssociateDRTRoleFuture) Get(ctx workflow.Context) (*shield.AssociateDRTRoleOutput, error) {
	var output shield.AssociateDRTRoleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldAssociateHealthCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldAssociateHealthCheckFuture) Get(ctx workflow.Context) (*shield.AssociateHealthCheckOutput, error) {
	var output shield.AssociateHealthCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldAssociateProactiveEngagementDetailsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldAssociateProactiveEngagementDetailsFuture) Get(ctx workflow.Context) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
	var output shield.AssociateProactiveEngagementDetailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldCreateProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldCreateProtectionFuture) Get(ctx workflow.Context) (*shield.CreateProtectionOutput, error) {
	var output shield.CreateProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldCreateSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldCreateSubscriptionFuture) Get(ctx workflow.Context) (*shield.CreateSubscriptionOutput, error) {
	var output shield.CreateSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDeleteProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDeleteProtectionFuture) Get(ctx workflow.Context) (*shield.DeleteProtectionOutput, error) {
	var output shield.DeleteProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDeleteSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDeleteSubscriptionFuture) Get(ctx workflow.Context) (*shield.DeleteSubscriptionOutput, error) {
	var output shield.DeleteSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDescribeAttackFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDescribeAttackFuture) Get(ctx workflow.Context) (*shield.DescribeAttackOutput, error) {
	var output shield.DescribeAttackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDescribeDRTAccessFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDescribeDRTAccessFuture) Get(ctx workflow.Context) (*shield.DescribeDRTAccessOutput, error) {
	var output shield.DescribeDRTAccessOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDescribeEmergencyContactSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDescribeEmergencyContactSettingsFuture) Get(ctx workflow.Context) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	var output shield.DescribeEmergencyContactSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDescribeProtectionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDescribeProtectionFuture) Get(ctx workflow.Context) (*shield.DescribeProtectionOutput, error) {
	var output shield.DescribeProtectionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDescribeSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDescribeSubscriptionFuture) Get(ctx workflow.Context) (*shield.DescribeSubscriptionOutput, error) {
	var output shield.DescribeSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDisableProactiveEngagementFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDisableProactiveEngagementFuture) Get(ctx workflow.Context) (*shield.DisableProactiveEngagementOutput, error) {
	var output shield.DisableProactiveEngagementOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDisassociateDRTLogBucketFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDisassociateDRTLogBucketFuture) Get(ctx workflow.Context) (*shield.DisassociateDRTLogBucketOutput, error) {
	var output shield.DisassociateDRTLogBucketOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDisassociateDRTRoleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDisassociateDRTRoleFuture) Get(ctx workflow.Context) (*shield.DisassociateDRTRoleOutput, error) {
	var output shield.DisassociateDRTRoleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldDisassociateHealthCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldDisassociateHealthCheckFuture) Get(ctx workflow.Context) (*shield.DisassociateHealthCheckOutput, error) {
	var output shield.DisassociateHealthCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldEnableProactiveEngagementFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldEnableProactiveEngagementFuture) Get(ctx workflow.Context) (*shield.EnableProactiveEngagementOutput, error) {
	var output shield.EnableProactiveEngagementOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldGetSubscriptionStateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldGetSubscriptionStateFuture) Get(ctx workflow.Context) (*shield.GetSubscriptionStateOutput, error) {
	var output shield.GetSubscriptionStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldListAttacksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldListAttacksFuture) Get(ctx workflow.Context) (*shield.ListAttacksOutput, error) {
	var output shield.ListAttacksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldListProtectionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldListProtectionsFuture) Get(ctx workflow.Context) (*shield.ListProtectionsOutput, error) {
	var output shield.ListProtectionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldUpdateEmergencyContactSettingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldUpdateEmergencyContactSettingsFuture) Get(ctx workflow.Context) (*shield.UpdateEmergencyContactSettingsOutput, error) {
	var output shield.UpdateEmergencyContactSettingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ShieldUpdateSubscriptionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ShieldUpdateSubscriptionFuture) Get(ctx workflow.Context) (*shield.UpdateSubscriptionOutput, error) {
	var output shield.UpdateSubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDRTLogBucket(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) (*shield.AssociateDRTLogBucketOutput, error) {
	var output shield.AssociateDRTLogBucketOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.AssociateDRTLogBucket", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDRTLogBucketAsync(ctx workflow.Context, input *shield.AssociateDRTLogBucketInput) *ShieldAssociateDRTLogBucketFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.AssociateDRTLogBucket", input)
	return &ShieldAssociateDRTLogBucketFuture{Future: future}
}

func (a *stub) AssociateDRTRole(ctx workflow.Context, input *shield.AssociateDRTRoleInput) (*shield.AssociateDRTRoleOutput, error) {
	var output shield.AssociateDRTRoleOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.AssociateDRTRole", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateDRTRoleAsync(ctx workflow.Context, input *shield.AssociateDRTRoleInput) *ShieldAssociateDRTRoleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.AssociateDRTRole", input)
	return &ShieldAssociateDRTRoleFuture{Future: future}
}

func (a *stub) AssociateHealthCheck(ctx workflow.Context, input *shield.AssociateHealthCheckInput) (*shield.AssociateHealthCheckOutput, error) {
	var output shield.AssociateHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.AssociateHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateHealthCheckAsync(ctx workflow.Context, input *shield.AssociateHealthCheckInput) *ShieldAssociateHealthCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.AssociateHealthCheck", input)
	return &ShieldAssociateHealthCheckFuture{Future: future}
}

func (a *stub) AssociateProactiveEngagementDetails(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) (*shield.AssociateProactiveEngagementDetailsOutput, error) {
	var output shield.AssociateProactiveEngagementDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.AssociateProactiveEngagementDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateProactiveEngagementDetailsAsync(ctx workflow.Context, input *shield.AssociateProactiveEngagementDetailsInput) *ShieldAssociateProactiveEngagementDetailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.AssociateProactiveEngagementDetails", input)
	return &ShieldAssociateProactiveEngagementDetailsFuture{Future: future}
}

func (a *stub) CreateProtection(ctx workflow.Context, input *shield.CreateProtectionInput) (*shield.CreateProtectionOutput, error) {
	var output shield.CreateProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.CreateProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateProtectionAsync(ctx workflow.Context, input *shield.CreateProtectionInput) *ShieldCreateProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.CreateProtection", input)
	return &ShieldCreateProtectionFuture{Future: future}
}

func (a *stub) CreateSubscription(ctx workflow.Context, input *shield.CreateSubscriptionInput) (*shield.CreateSubscriptionOutput, error) {
	var output shield.CreateSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.CreateSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateSubscriptionAsync(ctx workflow.Context, input *shield.CreateSubscriptionInput) *ShieldCreateSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.CreateSubscription", input)
	return &ShieldCreateSubscriptionFuture{Future: future}
}

func (a *stub) DeleteProtection(ctx workflow.Context, input *shield.DeleteProtectionInput) (*shield.DeleteProtectionOutput, error) {
	var output shield.DeleteProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DeleteProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProtectionAsync(ctx workflow.Context, input *shield.DeleteProtectionInput) *ShieldDeleteProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DeleteProtection", input)
	return &ShieldDeleteProtectionFuture{Future: future}
}

func (a *stub) DeleteSubscription(ctx workflow.Context, input *shield.DeleteSubscriptionInput) (*shield.DeleteSubscriptionOutput, error) {
	var output shield.DeleteSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DeleteSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteSubscriptionAsync(ctx workflow.Context, input *shield.DeleteSubscriptionInput) *ShieldDeleteSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DeleteSubscription", input)
	return &ShieldDeleteSubscriptionFuture{Future: future}
}

func (a *stub) DescribeAttack(ctx workflow.Context, input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error) {
	var output shield.DescribeAttackOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DescribeAttack", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAttackAsync(ctx workflow.Context, input *shield.DescribeAttackInput) *ShieldDescribeAttackFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DescribeAttack", input)
	return &ShieldDescribeAttackFuture{Future: future}
}

func (a *stub) DescribeDRTAccess(ctx workflow.Context, input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error) {
	var output shield.DescribeDRTAccessOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DescribeDRTAccess", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDRTAccessAsync(ctx workflow.Context, input *shield.DescribeDRTAccessInput) *ShieldDescribeDRTAccessFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DescribeDRTAccess", input)
	return &ShieldDescribeDRTAccessFuture{Future: future}
}

func (a *stub) DescribeEmergencyContactSettings(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	var output shield.DescribeEmergencyContactSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DescribeEmergencyContactSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.DescribeEmergencyContactSettingsInput) *ShieldDescribeEmergencyContactSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DescribeEmergencyContactSettings", input)
	return &ShieldDescribeEmergencyContactSettingsFuture{Future: future}
}

func (a *stub) DescribeProtection(ctx workflow.Context, input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error) {
	var output shield.DescribeProtectionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DescribeProtection", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeProtectionAsync(ctx workflow.Context, input *shield.DescribeProtectionInput) *ShieldDescribeProtectionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DescribeProtection", input)
	return &ShieldDescribeProtectionFuture{Future: future}
}

func (a *stub) DescribeSubscription(ctx workflow.Context, input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error) {
	var output shield.DescribeSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DescribeSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSubscriptionAsync(ctx workflow.Context, input *shield.DescribeSubscriptionInput) *ShieldDescribeSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DescribeSubscription", input)
	return &ShieldDescribeSubscriptionFuture{Future: future}
}

func (a *stub) DisableProactiveEngagement(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) (*shield.DisableProactiveEngagementOutput, error) {
	var output shield.DisableProactiveEngagementOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DisableProactiveEngagement", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableProactiveEngagementAsync(ctx workflow.Context, input *shield.DisableProactiveEngagementInput) *ShieldDisableProactiveEngagementFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DisableProactiveEngagement", input)
	return &ShieldDisableProactiveEngagementFuture{Future: future}
}

func (a *stub) DisassociateDRTLogBucket(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) (*shield.DisassociateDRTLogBucketOutput, error) {
	var output shield.DisassociateDRTLogBucketOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DisassociateDRTLogBucket", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateDRTLogBucketAsync(ctx workflow.Context, input *shield.DisassociateDRTLogBucketInput) *ShieldDisassociateDRTLogBucketFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DisassociateDRTLogBucket", input)
	return &ShieldDisassociateDRTLogBucketFuture{Future: future}
}

func (a *stub) DisassociateDRTRole(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) (*shield.DisassociateDRTRoleOutput, error) {
	var output shield.DisassociateDRTRoleOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DisassociateDRTRole", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateDRTRoleAsync(ctx workflow.Context, input *shield.DisassociateDRTRoleInput) *ShieldDisassociateDRTRoleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DisassociateDRTRole", input)
	return &ShieldDisassociateDRTRoleFuture{Future: future}
}

func (a *stub) DisassociateHealthCheck(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) (*shield.DisassociateHealthCheckOutput, error) {
	var output shield.DisassociateHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.DisassociateHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateHealthCheckAsync(ctx workflow.Context, input *shield.DisassociateHealthCheckInput) *ShieldDisassociateHealthCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.DisassociateHealthCheck", input)
	return &ShieldDisassociateHealthCheckFuture{Future: future}
}

func (a *stub) EnableProactiveEngagement(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) (*shield.EnableProactiveEngagementOutput, error) {
	var output shield.EnableProactiveEngagementOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.EnableProactiveEngagement", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableProactiveEngagementAsync(ctx workflow.Context, input *shield.EnableProactiveEngagementInput) *ShieldEnableProactiveEngagementFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.EnableProactiveEngagement", input)
	return &ShieldEnableProactiveEngagementFuture{Future: future}
}

func (a *stub) GetSubscriptionState(ctx workflow.Context, input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error) {
	var output shield.GetSubscriptionStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.GetSubscriptionState", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetSubscriptionStateAsync(ctx workflow.Context, input *shield.GetSubscriptionStateInput) *ShieldGetSubscriptionStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.GetSubscriptionState", input)
	return &ShieldGetSubscriptionStateFuture{Future: future}
}

func (a *stub) ListAttacks(ctx workflow.Context, input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error) {
	var output shield.ListAttacksOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.ListAttacks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAttacksAsync(ctx workflow.Context, input *shield.ListAttacksInput) *ShieldListAttacksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.ListAttacks", input)
	return &ShieldListAttacksFuture{Future: future}
}

func (a *stub) ListProtections(ctx workflow.Context, input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error) {
	var output shield.ListProtectionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.ListProtections", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProtectionsAsync(ctx workflow.Context, input *shield.ListProtectionsInput) *ShieldListProtectionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.ListProtections", input)
	return &ShieldListProtectionsFuture{Future: future}
}

func (a *stub) UpdateEmergencyContactSettings(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) (*shield.UpdateEmergencyContactSettingsOutput, error) {
	var output shield.UpdateEmergencyContactSettingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.UpdateEmergencyContactSettings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateEmergencyContactSettingsAsync(ctx workflow.Context, input *shield.UpdateEmergencyContactSettingsInput) *ShieldUpdateEmergencyContactSettingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.UpdateEmergencyContactSettings", input)
	return &ShieldUpdateEmergencyContactSettingsFuture{Future: future}
}

func (a *stub) UpdateSubscription(ctx workflow.Context, input *shield.UpdateSubscriptionInput) (*shield.UpdateSubscriptionOutput, error) {
	var output shield.UpdateSubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.shield.UpdateSubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateSubscriptionAsync(ctx workflow.Context, input *shield.UpdateSubscriptionInput) *ShieldUpdateSubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.shield.UpdateSubscription", input)
	return &ShieldUpdateSubscriptionFuture{Future: future}
}
