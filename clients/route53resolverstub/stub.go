// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package route53resolverstub

import (
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AssociateResolverEndpointIpAddressFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateResolverEndpointIpAddressFuture) Get(ctx workflow.Context) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
	var output route53resolver.AssociateResolverEndpointIpAddressOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateResolverQueryLogConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateResolverQueryLogConfigFuture) Get(ctx workflow.Context) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
	var output route53resolver.AssociateResolverQueryLogConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AssociateResolverRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateResolverRuleFuture) Get(ctx workflow.Context) (*route53resolver.AssociateResolverRuleOutput, error) {
	var output route53resolver.AssociateResolverRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateResolverEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateResolverEndpointFuture) Get(ctx workflow.Context) (*route53resolver.CreateResolverEndpointOutput, error) {
	var output route53resolver.CreateResolverEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateResolverQueryLogConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateResolverQueryLogConfigFuture) Get(ctx workflow.Context) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
	var output route53resolver.CreateResolverQueryLogConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateResolverRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateResolverRuleFuture) Get(ctx workflow.Context) (*route53resolver.CreateResolverRuleOutput, error) {
	var output route53resolver.CreateResolverRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteResolverEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteResolverEndpointFuture) Get(ctx workflow.Context) (*route53resolver.DeleteResolverEndpointOutput, error) {
	var output route53resolver.DeleteResolverEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteResolverQueryLogConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteResolverQueryLogConfigFuture) Get(ctx workflow.Context) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
	var output route53resolver.DeleteResolverQueryLogConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteResolverRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteResolverRuleFuture) Get(ctx workflow.Context) (*route53resolver.DeleteResolverRuleOutput, error) {
	var output route53resolver.DeleteResolverRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateResolverEndpointIpAddressFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateResolverEndpointIpAddressFuture) Get(ctx workflow.Context) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
	var output route53resolver.DisassociateResolverEndpointIpAddressOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateResolverQueryLogConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateResolverQueryLogConfigFuture) Get(ctx workflow.Context) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
	var output route53resolver.DisassociateResolverQueryLogConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateResolverRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateResolverRuleFuture) Get(ctx workflow.Context) (*route53resolver.DisassociateResolverRuleOutput, error) {
	var output route53resolver.DisassociateResolverRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverEndpointFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverEndpointOutput, error) {
	var output route53resolver.GetResolverEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverQueryLogConfigFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverQueryLogConfigFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
	var output route53resolver.GetResolverQueryLogConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverQueryLogConfigAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverQueryLogConfigAssociationFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
	var output route53resolver.GetResolverQueryLogConfigAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverQueryLogConfigPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverQueryLogConfigPolicyFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
	var output route53resolver.GetResolverQueryLogConfigPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverRuleFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverRuleOutput, error) {
	var output route53resolver.GetResolverRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverRuleAssociationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverRuleAssociationFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	var output route53resolver.GetResolverRuleAssociationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetResolverRulePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetResolverRulePolicyFuture) Get(ctx workflow.Context) (*route53resolver.GetResolverRulePolicyOutput, error) {
	var output route53resolver.GetResolverRulePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResolverEndpointIpAddressesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResolverEndpointIpAddressesFuture) Get(ctx workflow.Context) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	var output route53resolver.ListResolverEndpointIpAddressesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResolverEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResolverEndpointsFuture) Get(ctx workflow.Context) (*route53resolver.ListResolverEndpointsOutput, error) {
	var output route53resolver.ListResolverEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResolverQueryLogConfigAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResolverQueryLogConfigAssociationsFuture) Get(ctx workflow.Context) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
	var output route53resolver.ListResolverQueryLogConfigAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResolverQueryLogConfigsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResolverQueryLogConfigsFuture) Get(ctx workflow.Context) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
	var output route53resolver.ListResolverQueryLogConfigsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResolverRuleAssociationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResolverRuleAssociationsFuture) Get(ctx workflow.Context) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	var output route53resolver.ListResolverRuleAssociationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListResolverRulesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListResolverRulesFuture) Get(ctx workflow.Context) (*route53resolver.ListResolverRulesOutput, error) {
	var output route53resolver.ListResolverRulesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*route53resolver.ListTagsForResourceOutput, error) {
	var output route53resolver.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutResolverQueryLogConfigPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutResolverQueryLogConfigPolicyFuture) Get(ctx workflow.Context) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
	var output route53resolver.PutResolverQueryLogConfigPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutResolverRulePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutResolverRulePolicyFuture) Get(ctx workflow.Context) (*route53resolver.PutResolverRulePolicyOutput, error) {
	var output route53resolver.PutResolverRulePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*route53resolver.TagResourceOutput, error) {
	var output route53resolver.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*route53resolver.UntagResourceOutput, error) {
	var output route53resolver.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateResolverEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateResolverEndpointFuture) Get(ctx workflow.Context) (*route53resolver.UpdateResolverEndpointOutput, error) {
	var output route53resolver.UpdateResolverEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateResolverRuleFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateResolverRuleFuture) Get(ctx workflow.Context) (*route53resolver.UpdateResolverRuleOutput, error) {
	var output route53resolver.UpdateResolverRuleOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateResolverEndpointIpAddress(ctx workflow.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
	var output route53resolver.AssociateResolverEndpointIpAddressOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.AssociateResolverEndpointIpAddress", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateResolverEndpointIpAddressAsync(ctx workflow.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) *AssociateResolverEndpointIpAddressFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.AssociateResolverEndpointIpAddress", input)
	return &AssociateResolverEndpointIpAddressFuture{Future: future}
}

func (a *stub) AssociateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
	var output route53resolver.AssociateResolverQueryLogConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.AssociateResolverQueryLogConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) *AssociateResolverQueryLogConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.AssociateResolverQueryLogConfig", input)
	return &AssociateResolverQueryLogConfigFuture{Future: future}
}

func (a *stub) AssociateResolverRule(ctx workflow.Context, input *route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error) {
	var output route53resolver.AssociateResolverRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.AssociateResolverRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateResolverRuleAsync(ctx workflow.Context, input *route53resolver.AssociateResolverRuleInput) *AssociateResolverRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.AssociateResolverRule", input)
	return &AssociateResolverRuleFuture{Future: future}
}

func (a *stub) CreateResolverEndpoint(ctx workflow.Context, input *route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error) {
	var output route53resolver.CreateResolverEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.CreateResolverEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateResolverEndpointAsync(ctx workflow.Context, input *route53resolver.CreateResolverEndpointInput) *CreateResolverEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.CreateResolverEndpoint", input)
	return &CreateResolverEndpointFuture{Future: future}
}

func (a *stub) CreateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
	var output route53resolver.CreateResolverQueryLogConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.CreateResolverQueryLogConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.CreateResolverQueryLogConfigInput) *CreateResolverQueryLogConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.CreateResolverQueryLogConfig", input)
	return &CreateResolverQueryLogConfigFuture{Future: future}
}

func (a *stub) CreateResolverRule(ctx workflow.Context, input *route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error) {
	var output route53resolver.CreateResolverRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.CreateResolverRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateResolverRuleAsync(ctx workflow.Context, input *route53resolver.CreateResolverRuleInput) *CreateResolverRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.CreateResolverRule", input)
	return &CreateResolverRuleFuture{Future: future}
}

func (a *stub) DeleteResolverEndpoint(ctx workflow.Context, input *route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error) {
	var output route53resolver.DeleteResolverEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.DeleteResolverEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteResolverEndpointAsync(ctx workflow.Context, input *route53resolver.DeleteResolverEndpointInput) *DeleteResolverEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.DeleteResolverEndpoint", input)
	return &DeleteResolverEndpointFuture{Future: future}
}

func (a *stub) DeleteResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
	var output route53resolver.DeleteResolverQueryLogConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.DeleteResolverQueryLogConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) *DeleteResolverQueryLogConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.DeleteResolverQueryLogConfig", input)
	return &DeleteResolverQueryLogConfigFuture{Future: future}
}

func (a *stub) DeleteResolverRule(ctx workflow.Context, input *route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error) {
	var output route53resolver.DeleteResolverRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.DeleteResolverRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteResolverRuleAsync(ctx workflow.Context, input *route53resolver.DeleteResolverRuleInput) *DeleteResolverRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.DeleteResolverRule", input)
	return &DeleteResolverRuleFuture{Future: future}
}

func (a *stub) DisassociateResolverEndpointIpAddress(ctx workflow.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
	var output route53resolver.DisassociateResolverEndpointIpAddressOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.DisassociateResolverEndpointIpAddress", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateResolverEndpointIpAddressAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) *DisassociateResolverEndpointIpAddressFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.DisassociateResolverEndpointIpAddress", input)
	return &DisassociateResolverEndpointIpAddressFuture{Future: future}
}

func (a *stub) DisassociateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
	var output route53resolver.DisassociateResolverQueryLogConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.DisassociateResolverQueryLogConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) *DisassociateResolverQueryLogConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.DisassociateResolverQueryLogConfig", input)
	return &DisassociateResolverQueryLogConfigFuture{Future: future}
}

func (a *stub) DisassociateResolverRule(ctx workflow.Context, input *route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error) {
	var output route53resolver.DisassociateResolverRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.DisassociateResolverRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateResolverRuleAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverRuleInput) *DisassociateResolverRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.DisassociateResolverRule", input)
	return &DisassociateResolverRuleFuture{Future: future}
}

func (a *stub) GetResolverEndpoint(ctx workflow.Context, input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error) {
	var output route53resolver.GetResolverEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverEndpointAsync(ctx workflow.Context, input *route53resolver.GetResolverEndpointInput) *GetResolverEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverEndpoint", input)
	return &GetResolverEndpointFuture{Future: future}
}

func (a *stub) GetResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
	var output route53resolver.GetResolverQueryLogConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverQueryLogConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigInput) *GetResolverQueryLogConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverQueryLogConfig", input)
	return &GetResolverQueryLogConfigFuture{Future: future}
}

func (a *stub) GetResolverQueryLogConfigAssociation(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
	var output route53resolver.GetResolverQueryLogConfigAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverQueryLogConfigAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverQueryLogConfigAssociationAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) *GetResolverQueryLogConfigAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverQueryLogConfigAssociation", input)
	return &GetResolverQueryLogConfigAssociationFuture{Future: future}
}

func (a *stub) GetResolverQueryLogConfigPolicy(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
	var output route53resolver.GetResolverQueryLogConfigPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverQueryLogConfigPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverQueryLogConfigPolicyAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) *GetResolverQueryLogConfigPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverQueryLogConfigPolicy", input)
	return &GetResolverQueryLogConfigPolicyFuture{Future: future}
}

func (a *stub) GetResolverRule(ctx workflow.Context, input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error) {
	var output route53resolver.GetResolverRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverRuleAsync(ctx workflow.Context, input *route53resolver.GetResolverRuleInput) *GetResolverRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverRule", input)
	return &GetResolverRuleFuture{Future: future}
}

func (a *stub) GetResolverRuleAssociation(ctx workflow.Context, input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	var output route53resolver.GetResolverRuleAssociationOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverRuleAssociation", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverRuleAssociationAsync(ctx workflow.Context, input *route53resolver.GetResolverRuleAssociationInput) *GetResolverRuleAssociationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverRuleAssociation", input)
	return &GetResolverRuleAssociationFuture{Future: future}
}

func (a *stub) GetResolverRulePolicy(ctx workflow.Context, input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error) {
	var output route53resolver.GetResolverRulePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverRulePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetResolverRulePolicyAsync(ctx workflow.Context, input *route53resolver.GetResolverRulePolicyInput) *GetResolverRulePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.GetResolverRulePolicy", input)
	return &GetResolverRulePolicyFuture{Future: future}
}

func (a *stub) ListResolverEndpointIpAddresses(ctx workflow.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	var output route53resolver.ListResolverEndpointIpAddressesOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverEndpointIpAddresses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolverEndpointIpAddressesAsync(ctx workflow.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) *ListResolverEndpointIpAddressesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverEndpointIpAddresses", input)
	return &ListResolverEndpointIpAddressesFuture{Future: future}
}

func (a *stub) ListResolverEndpoints(ctx workflow.Context, input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error) {
	var output route53resolver.ListResolverEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolverEndpointsAsync(ctx workflow.Context, input *route53resolver.ListResolverEndpointsInput) *ListResolverEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverEndpoints", input)
	return &ListResolverEndpointsFuture{Future: future}
}

func (a *stub) ListResolverQueryLogConfigAssociations(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
	var output route53resolver.ListResolverQueryLogConfigAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverQueryLogConfigAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolverQueryLogConfigAssociationsAsync(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) *ListResolverQueryLogConfigAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverQueryLogConfigAssociations", input)
	return &ListResolverQueryLogConfigAssociationsFuture{Future: future}
}

func (a *stub) ListResolverQueryLogConfigs(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
	var output route53resolver.ListResolverQueryLogConfigsOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverQueryLogConfigs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolverQueryLogConfigsAsync(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigsInput) *ListResolverQueryLogConfigsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverQueryLogConfigs", input)
	return &ListResolverQueryLogConfigsFuture{Future: future}
}

func (a *stub) ListResolverRuleAssociations(ctx workflow.Context, input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	var output route53resolver.ListResolverRuleAssociationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverRuleAssociations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolverRuleAssociationsAsync(ctx workflow.Context, input *route53resolver.ListResolverRuleAssociationsInput) *ListResolverRuleAssociationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverRuleAssociations", input)
	return &ListResolverRuleAssociationsFuture{Future: future}
}

func (a *stub) ListResolverRules(ctx workflow.Context, input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error) {
	var output route53resolver.ListResolverRulesOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverRules", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListResolverRulesAsync(ctx workflow.Context, input *route53resolver.ListResolverRulesInput) *ListResolverRulesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListResolverRules", input)
	return &ListResolverRulesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error) {
	var output route53resolver.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *route53resolver.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PutResolverQueryLogConfigPolicy(ctx workflow.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
	var output route53resolver.PutResolverQueryLogConfigPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.PutResolverQueryLogConfigPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutResolverQueryLogConfigPolicyAsync(ctx workflow.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) *PutResolverQueryLogConfigPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.PutResolverQueryLogConfigPolicy", input)
	return &PutResolverQueryLogConfigPolicyFuture{Future: future}
}

func (a *stub) PutResolverRulePolicy(ctx workflow.Context, input *route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error) {
	var output route53resolver.PutResolverRulePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.PutResolverRulePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutResolverRulePolicyAsync(ctx workflow.Context, input *route53resolver.PutResolverRulePolicyInput) *PutResolverRulePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.PutResolverRulePolicy", input)
	return &PutResolverRulePolicyFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error) {
	var output route53resolver.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *route53resolver.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error) {
	var output route53resolver.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *route53resolver.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateResolverEndpoint(ctx workflow.Context, input *route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error) {
	var output route53resolver.UpdateResolverEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.UpdateResolverEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateResolverEndpointAsync(ctx workflow.Context, input *route53resolver.UpdateResolverEndpointInput) *UpdateResolverEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.UpdateResolverEndpoint", input)
	return &UpdateResolverEndpointFuture{Future: future}
}

func (a *stub) UpdateResolverRule(ctx workflow.Context, input *route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error) {
	var output route53resolver.UpdateResolverRuleOutput
	err := workflow.ExecuteActivity(ctx, "aws.route53resolver.UpdateResolverRule", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateResolverRuleAsync(ctx workflow.Context, input *route53resolver.UpdateResolverRuleInput) *UpdateResolverRuleFuture {
	future := workflow.ExecuteActivity(ctx, "aws.route53resolver.UpdateResolverRule", input)
	return &UpdateResolverRuleFuture{Future: future}
}
