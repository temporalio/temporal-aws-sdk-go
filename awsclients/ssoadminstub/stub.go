// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ssoadminstub

import (
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type SSOAdminAttachManagedPolicyToPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminAttachManagedPolicyToPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error) {
	var output ssoadmin.AttachManagedPolicyToPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminCreateAccountAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminCreateAccountAssignmentFuture) Get(ctx workflow.Context) (*ssoadmin.CreateAccountAssignmentOutput, error) {
	var output ssoadmin.CreateAccountAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminCreatePermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminCreatePermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.CreatePermissionSetOutput, error) {
	var output ssoadmin.CreatePermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDeleteAccountAssignmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDeleteAccountAssignmentFuture) Get(ctx workflow.Context) (*ssoadmin.DeleteAccountAssignmentOutput, error) {
	var output ssoadmin.DeleteAccountAssignmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDeleteInlinePolicyFromPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDeleteInlinePolicyFromPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error) {
	var output ssoadmin.DeleteInlinePolicyFromPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDeletePermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDeletePermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.DeletePermissionSetOutput, error) {
	var output ssoadmin.DeletePermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDescribeAccountAssignmentCreationStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDescribeAccountAssignmentCreationStatusFuture) Get(ctx workflow.Context) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
	var output ssoadmin.DescribeAccountAssignmentCreationStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDescribeAccountAssignmentDeletionStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDescribeAccountAssignmentDeletionStatusFuture) Get(ctx workflow.Context) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
	var output ssoadmin.DescribeAccountAssignmentDeletionStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDescribePermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDescribePermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.DescribePermissionSetOutput, error) {
	var output ssoadmin.DescribePermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDescribePermissionSetProvisioningStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDescribePermissionSetProvisioningStatusFuture) Get(ctx workflow.Context) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
	var output ssoadmin.DescribePermissionSetProvisioningStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminDetachManagedPolicyFromPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminDetachManagedPolicyFromPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error) {
	var output ssoadmin.DetachManagedPolicyFromPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminGetInlinePolicyForPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminGetInlinePolicyForPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
	var output ssoadmin.GetInlinePolicyForPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListAccountAssignmentCreationStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListAccountAssignmentCreationStatusFuture) Get(ctx workflow.Context) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
	var output ssoadmin.ListAccountAssignmentCreationStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListAccountAssignmentDeletionStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListAccountAssignmentDeletionStatusFuture) Get(ctx workflow.Context) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
	var output ssoadmin.ListAccountAssignmentDeletionStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListAccountAssignmentsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListAccountAssignmentsFuture) Get(ctx workflow.Context) (*ssoadmin.ListAccountAssignmentsOutput, error) {
	var output ssoadmin.ListAccountAssignmentsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListAccountsForProvisionedPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListAccountsForProvisionedPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
	var output ssoadmin.ListAccountsForProvisionedPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListInstancesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListInstancesFuture) Get(ctx workflow.Context) (*ssoadmin.ListInstancesOutput, error) {
	var output ssoadmin.ListInstancesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListManagedPoliciesInPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListManagedPoliciesInPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
	var output ssoadmin.ListManagedPoliciesInPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListPermissionSetProvisioningStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListPermissionSetProvisioningStatusFuture) Get(ctx workflow.Context) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
	var output ssoadmin.ListPermissionSetProvisioningStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListPermissionSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListPermissionSetsFuture) Get(ctx workflow.Context) (*ssoadmin.ListPermissionSetsOutput, error) {
	var output ssoadmin.ListPermissionSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListPermissionSetsProvisionedToAccountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListPermissionSetsProvisionedToAccountFuture) Get(ctx workflow.Context) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
	var output ssoadmin.ListPermissionSetsProvisionedToAccountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminListTagsForResourceFuture) Get(ctx workflow.Context) (*ssoadmin.ListTagsForResourceOutput, error) {
	var output ssoadmin.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminProvisionPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminProvisionPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.ProvisionPermissionSetOutput, error) {
	var output ssoadmin.ProvisionPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminPutInlinePolicyToPermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminPutInlinePolicyToPermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error) {
	var output ssoadmin.PutInlinePolicyToPermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminTagResourceFuture) Get(ctx workflow.Context) (*ssoadmin.TagResourceOutput, error) {
	var output ssoadmin.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminUntagResourceFuture) Get(ctx workflow.Context) (*ssoadmin.UntagResourceOutput, error) {
	var output ssoadmin.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOAdminUpdatePermissionSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SSOAdminUpdatePermissionSetFuture) Get(ctx workflow.Context) (*ssoadmin.UpdatePermissionSetOutput, error) {
	var output ssoadmin.UpdatePermissionSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AttachManagedPolicyToPermissionSet(ctx workflow.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error) {
	var output ssoadmin.AttachManagedPolicyToPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.AttachManagedPolicyToPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AttachManagedPolicyToPermissionSetAsync(ctx workflow.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) *SSOAdminAttachManagedPolicyToPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.AttachManagedPolicyToPermissionSet", input)
	return &SSOAdminAttachManagedPolicyToPermissionSetFuture{Future: future}
}

func (a *stub) CreateAccountAssignment(ctx workflow.Context, input *ssoadmin.CreateAccountAssignmentInput) (*ssoadmin.CreateAccountAssignmentOutput, error) {
	var output ssoadmin.CreateAccountAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.CreateAccountAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAccountAssignmentAsync(ctx workflow.Context, input *ssoadmin.CreateAccountAssignmentInput) *SSOAdminCreateAccountAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.CreateAccountAssignment", input)
	return &SSOAdminCreateAccountAssignmentFuture{Future: future}
}

func (a *stub) CreatePermissionSet(ctx workflow.Context, input *ssoadmin.CreatePermissionSetInput) (*ssoadmin.CreatePermissionSetOutput, error) {
	var output ssoadmin.CreatePermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.CreatePermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePermissionSetAsync(ctx workflow.Context, input *ssoadmin.CreatePermissionSetInput) *SSOAdminCreatePermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.CreatePermissionSet", input)
	return &SSOAdminCreatePermissionSetFuture{Future: future}
}

func (a *stub) DeleteAccountAssignment(ctx workflow.Context, input *ssoadmin.DeleteAccountAssignmentInput) (*ssoadmin.DeleteAccountAssignmentOutput, error) {
	var output ssoadmin.DeleteAccountAssignmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DeleteAccountAssignment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAccountAssignmentAsync(ctx workflow.Context, input *ssoadmin.DeleteAccountAssignmentInput) *SSOAdminDeleteAccountAssignmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DeleteAccountAssignment", input)
	return &SSOAdminDeleteAccountAssignmentFuture{Future: future}
}

func (a *stub) DeleteInlinePolicyFromPermissionSet(ctx workflow.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error) {
	var output ssoadmin.DeleteInlinePolicyFromPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DeleteInlinePolicyFromPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteInlinePolicyFromPermissionSetAsync(ctx workflow.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) *SSOAdminDeleteInlinePolicyFromPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DeleteInlinePolicyFromPermissionSet", input)
	return &SSOAdminDeleteInlinePolicyFromPermissionSetFuture{Future: future}
}

func (a *stub) DeletePermissionSet(ctx workflow.Context, input *ssoadmin.DeletePermissionSetInput) (*ssoadmin.DeletePermissionSetOutput, error) {
	var output ssoadmin.DeletePermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DeletePermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePermissionSetAsync(ctx workflow.Context, input *ssoadmin.DeletePermissionSetInput) *SSOAdminDeletePermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DeletePermissionSet", input)
	return &SSOAdminDeletePermissionSetFuture{Future: future}
}

func (a *stub) DescribeAccountAssignmentCreationStatus(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
	var output ssoadmin.DescribeAccountAssignmentCreationStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribeAccountAssignmentCreationStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAccountAssignmentCreationStatusAsync(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) *SSOAdminDescribeAccountAssignmentCreationStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribeAccountAssignmentCreationStatus", input)
	return &SSOAdminDescribeAccountAssignmentCreationStatusFuture{Future: future}
}

func (a *stub) DescribeAccountAssignmentDeletionStatus(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
	var output ssoadmin.DescribeAccountAssignmentDeletionStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribeAccountAssignmentDeletionStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAccountAssignmentDeletionStatusAsync(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) *SSOAdminDescribeAccountAssignmentDeletionStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribeAccountAssignmentDeletionStatus", input)
	return &SSOAdminDescribeAccountAssignmentDeletionStatusFuture{Future: future}
}

func (a *stub) DescribePermissionSet(ctx workflow.Context, input *ssoadmin.DescribePermissionSetInput) (*ssoadmin.DescribePermissionSetOutput, error) {
	var output ssoadmin.DescribePermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribePermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePermissionSetAsync(ctx workflow.Context, input *ssoadmin.DescribePermissionSetInput) *SSOAdminDescribePermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribePermissionSet", input)
	return &SSOAdminDescribePermissionSetFuture{Future: future}
}

func (a *stub) DescribePermissionSetProvisioningStatus(ctx workflow.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
	var output ssoadmin.DescribePermissionSetProvisioningStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribePermissionSetProvisioningStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePermissionSetProvisioningStatusAsync(ctx workflow.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) *SSOAdminDescribePermissionSetProvisioningStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DescribePermissionSetProvisioningStatus", input)
	return &SSOAdminDescribePermissionSetProvisioningStatusFuture{Future: future}
}

func (a *stub) DetachManagedPolicyFromPermissionSet(ctx workflow.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error) {
	var output ssoadmin.DetachManagedPolicyFromPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DetachManagedPolicyFromPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DetachManagedPolicyFromPermissionSetAsync(ctx workflow.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) *SSOAdminDetachManagedPolicyFromPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.DetachManagedPolicyFromPermissionSet", input)
	return &SSOAdminDetachManagedPolicyFromPermissionSetFuture{Future: future}
}

func (a *stub) GetInlinePolicyForPermissionSet(ctx workflow.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
	var output ssoadmin.GetInlinePolicyForPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.GetInlinePolicyForPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetInlinePolicyForPermissionSetAsync(ctx workflow.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) *SSOAdminGetInlinePolicyForPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.GetInlinePolicyForPermissionSet", input)
	return &SSOAdminGetInlinePolicyForPermissionSetFuture{Future: future}
}

func (a *stub) ListAccountAssignmentCreationStatus(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
	var output ssoadmin.ListAccountAssignmentCreationStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountAssignmentCreationStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountAssignmentCreationStatusAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) *SSOAdminListAccountAssignmentCreationStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountAssignmentCreationStatus", input)
	return &SSOAdminListAccountAssignmentCreationStatusFuture{Future: future}
}

func (a *stub) ListAccountAssignmentDeletionStatus(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
	var output ssoadmin.ListAccountAssignmentDeletionStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountAssignmentDeletionStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountAssignmentDeletionStatusAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) *SSOAdminListAccountAssignmentDeletionStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountAssignmentDeletionStatus", input)
	return &SSOAdminListAccountAssignmentDeletionStatusFuture{Future: future}
}

func (a *stub) ListAccountAssignments(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentsInput) (*ssoadmin.ListAccountAssignmentsOutput, error) {
	var output ssoadmin.ListAccountAssignmentsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountAssignments", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountAssignmentsAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentsInput) *SSOAdminListAccountAssignmentsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountAssignments", input)
	return &SSOAdminListAccountAssignmentsFuture{Future: future}
}

func (a *stub) ListAccountsForProvisionedPermissionSet(ctx workflow.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
	var output ssoadmin.ListAccountsForProvisionedPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountsForProvisionedPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAccountsForProvisionedPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) *SSOAdminListAccountsForProvisionedPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListAccountsForProvisionedPermissionSet", input)
	return &SSOAdminListAccountsForProvisionedPermissionSetFuture{Future: future}
}

func (a *stub) ListInstances(ctx workflow.Context, input *ssoadmin.ListInstancesInput) (*ssoadmin.ListInstancesOutput, error) {
	var output ssoadmin.ListInstancesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListInstances", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInstancesAsync(ctx workflow.Context, input *ssoadmin.ListInstancesInput) *SSOAdminListInstancesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListInstances", input)
	return &SSOAdminListInstancesFuture{Future: future}
}

func (a *stub) ListManagedPoliciesInPermissionSet(ctx workflow.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
	var output ssoadmin.ListManagedPoliciesInPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListManagedPoliciesInPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListManagedPoliciesInPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) *SSOAdminListManagedPoliciesInPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListManagedPoliciesInPermissionSet", input)
	return &SSOAdminListManagedPoliciesInPermissionSetFuture{Future: future}
}

func (a *stub) ListPermissionSetProvisioningStatus(ctx workflow.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
	var output ssoadmin.ListPermissionSetProvisioningStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListPermissionSetProvisioningStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPermissionSetProvisioningStatusAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) *SSOAdminListPermissionSetProvisioningStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListPermissionSetProvisioningStatus", input)
	return &SSOAdminListPermissionSetProvisioningStatusFuture{Future: future}
}

func (a *stub) ListPermissionSets(ctx workflow.Context, input *ssoadmin.ListPermissionSetsInput) (*ssoadmin.ListPermissionSetsOutput, error) {
	var output ssoadmin.ListPermissionSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListPermissionSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPermissionSetsAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetsInput) *SSOAdminListPermissionSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListPermissionSets", input)
	return &SSOAdminListPermissionSetsFuture{Future: future}
}

func (a *stub) ListPermissionSetsProvisionedToAccount(ctx workflow.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
	var output ssoadmin.ListPermissionSetsProvisionedToAccountOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListPermissionSetsProvisionedToAccount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPermissionSetsProvisionedToAccountAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) *SSOAdminListPermissionSetsProvisionedToAccountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListPermissionSetsProvisionedToAccount", input)
	return &SSOAdminListPermissionSetsProvisionedToAccountFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *ssoadmin.ListTagsForResourceInput) (*ssoadmin.ListTagsForResourceOutput, error) {
	var output ssoadmin.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *ssoadmin.ListTagsForResourceInput) *SSOAdminListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ListTagsForResource", input)
	return &SSOAdminListTagsForResourceFuture{Future: future}
}

func (a *stub) ProvisionPermissionSet(ctx workflow.Context, input *ssoadmin.ProvisionPermissionSetInput) (*ssoadmin.ProvisionPermissionSetOutput, error) {
	var output ssoadmin.ProvisionPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ProvisionPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ProvisionPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ProvisionPermissionSetInput) *SSOAdminProvisionPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.ProvisionPermissionSet", input)
	return &SSOAdminProvisionPermissionSetFuture{Future: future}
}

func (a *stub) PutInlinePolicyToPermissionSet(ctx workflow.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error) {
	var output ssoadmin.PutInlinePolicyToPermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.PutInlinePolicyToPermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutInlinePolicyToPermissionSetAsync(ctx workflow.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) *SSOAdminPutInlinePolicyToPermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.PutInlinePolicyToPermissionSet", input)
	return &SSOAdminPutInlinePolicyToPermissionSetFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *ssoadmin.TagResourceInput) (*ssoadmin.TagResourceOutput, error) {
	var output ssoadmin.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *ssoadmin.TagResourceInput) *SSOAdminTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.TagResource", input)
	return &SSOAdminTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *ssoadmin.UntagResourceInput) (*ssoadmin.UntagResourceOutput, error) {
	var output ssoadmin.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *ssoadmin.UntagResourceInput) *SSOAdminUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.UntagResource", input)
	return &SSOAdminUntagResourceFuture{Future: future}
}

func (a *stub) UpdatePermissionSet(ctx workflow.Context, input *ssoadmin.UpdatePermissionSetInput) (*ssoadmin.UpdatePermissionSetOutput, error) {
	var output ssoadmin.UpdatePermissionSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssoadmin.UpdatePermissionSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdatePermissionSetAsync(ctx workflow.Context, input *ssoadmin.UpdatePermissionSetInput) *SSOAdminUpdatePermissionSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssoadmin.UpdatePermissionSet", input)
	return &SSOAdminUpdatePermissionSetFuture{Future: future}
}
