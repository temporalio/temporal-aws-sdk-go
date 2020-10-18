// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package maciestub

import (
	"github.com/aws/aws-sdk-go/service/macie"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MacieAssociateMemberAccountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieAssociateMemberAccountFuture) Get(ctx workflow.Context) (*macie.AssociateMemberAccountOutput, error) {
	var output macie.AssociateMemberAccountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MacieAssociateS3ResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieAssociateS3ResourcesFuture) Get(ctx workflow.Context) (*macie.AssociateS3ResourcesOutput, error) {
	var output macie.AssociateS3ResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MacieDisassociateMemberAccountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieDisassociateMemberAccountFuture) Get(ctx workflow.Context) (*macie.DisassociateMemberAccountOutput, error) {
	var output macie.DisassociateMemberAccountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MacieDisassociateS3ResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieDisassociateS3ResourcesFuture) Get(ctx workflow.Context) (*macie.DisassociateS3ResourcesOutput, error) {
	var output macie.DisassociateS3ResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MacieListMemberAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieListMemberAccountsFuture) Get(ctx workflow.Context) (*macie.ListMemberAccountsOutput, error) {
	var output macie.ListMemberAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MacieListS3ResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieListS3ResourcesFuture) Get(ctx workflow.Context) (*macie.ListS3ResourcesOutput, error) {
	var output macie.ListS3ResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MacieUpdateS3ResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MacieUpdateS3ResourcesFuture) Get(ctx workflow.Context) (*macie.UpdateS3ResourcesOutput, error) {
	var output macie.UpdateS3ResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateMemberAccount(ctx workflow.Context, input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error) {
	var output macie.AssociateMemberAccountOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.AssociateMemberAccount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateMemberAccountAsync(ctx workflow.Context, input *macie.AssociateMemberAccountInput) *MacieAssociateMemberAccountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.AssociateMemberAccount", input)
	return &MacieAssociateMemberAccountFuture{Future: future}
}

func (a *stub) AssociateS3Resources(ctx workflow.Context, input *macie.AssociateS3ResourcesInput) (*macie.AssociateS3ResourcesOutput, error) {
	var output macie.AssociateS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.AssociateS3Resources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateS3ResourcesAsync(ctx workflow.Context, input *macie.AssociateS3ResourcesInput) *MacieAssociateS3ResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.AssociateS3Resources", input)
	return &MacieAssociateS3ResourcesFuture{Future: future}
}

func (a *stub) DisassociateMemberAccount(ctx workflow.Context, input *macie.DisassociateMemberAccountInput) (*macie.DisassociateMemberAccountOutput, error) {
	var output macie.DisassociateMemberAccountOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.DisassociateMemberAccount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateMemberAccountAsync(ctx workflow.Context, input *macie.DisassociateMemberAccountInput) *MacieDisassociateMemberAccountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.DisassociateMemberAccount", input)
	return &MacieDisassociateMemberAccountFuture{Future: future}
}

func (a *stub) DisassociateS3Resources(ctx workflow.Context, input *macie.DisassociateS3ResourcesInput) (*macie.DisassociateS3ResourcesOutput, error) {
	var output macie.DisassociateS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.DisassociateS3Resources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateS3ResourcesAsync(ctx workflow.Context, input *macie.DisassociateS3ResourcesInput) *MacieDisassociateS3ResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.DisassociateS3Resources", input)
	return &MacieDisassociateS3ResourcesFuture{Future: future}
}

func (a *stub) ListMemberAccounts(ctx workflow.Context, input *macie.ListMemberAccountsInput) (*macie.ListMemberAccountsOutput, error) {
	var output macie.ListMemberAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.ListMemberAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMemberAccountsAsync(ctx workflow.Context, input *macie.ListMemberAccountsInput) *MacieListMemberAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.ListMemberAccounts", input)
	return &MacieListMemberAccountsFuture{Future: future}
}

func (a *stub) ListS3Resources(ctx workflow.Context, input *macie.ListS3ResourcesInput) (*macie.ListS3ResourcesOutput, error) {
	var output macie.ListS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.ListS3Resources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListS3ResourcesAsync(ctx workflow.Context, input *macie.ListS3ResourcesInput) *MacieListS3ResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.ListS3Resources", input)
	return &MacieListS3ResourcesFuture{Future: future}
}

func (a *stub) UpdateS3Resources(ctx workflow.Context, input *macie.UpdateS3ResourcesInput) (*macie.UpdateS3ResourcesOutput, error) {
	var output macie.UpdateS3ResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.macie.UpdateS3Resources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateS3ResourcesAsync(ctx workflow.Context, input *macie.UpdateS3ResourcesInput) *MacieUpdateS3ResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.macie.UpdateS3Resources", input)
	return &MacieUpdateS3ResourcesFuture{Future: future}
}
