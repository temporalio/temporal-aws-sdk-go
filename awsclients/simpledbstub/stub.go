// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package simpledbstub

import (
	"github.com/aws/aws-sdk-go/service/simpledb"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ awsclients.VoidFuture

type stub struct{}

type SimpleDBBatchDeleteAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBBatchDeleteAttributesFuture) Get(ctx workflow.Context) (*simpledb.BatchDeleteAttributesOutput, error) {
	var output simpledb.BatchDeleteAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBBatchPutAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBBatchPutAttributesFuture) Get(ctx workflow.Context) (*simpledb.BatchPutAttributesOutput, error) {
	var output simpledb.BatchPutAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBCreateDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBCreateDomainFuture) Get(ctx workflow.Context) (*simpledb.CreateDomainOutput, error) {
	var output simpledb.CreateDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBDeleteAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBDeleteAttributesFuture) Get(ctx workflow.Context) (*simpledb.DeleteAttributesOutput, error) {
	var output simpledb.DeleteAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBDeleteDomainFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBDeleteDomainFuture) Get(ctx workflow.Context) (*simpledb.DeleteDomainOutput, error) {
	var output simpledb.DeleteDomainOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBDomainMetadataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBDomainMetadataFuture) Get(ctx workflow.Context) (*simpledb.DomainMetadataOutput, error) {
	var output simpledb.DomainMetadataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBGetAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBGetAttributesFuture) Get(ctx workflow.Context) (*simpledb.GetAttributesOutput, error) {
	var output simpledb.GetAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBListDomainsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBListDomainsFuture) Get(ctx workflow.Context) (*simpledb.ListDomainsOutput, error) {
	var output simpledb.ListDomainsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBPutAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBPutAttributesFuture) Get(ctx workflow.Context) (*simpledb.PutAttributesOutput, error) {
	var output simpledb.PutAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SimpleDBSelectFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SimpleDBSelectFuture) Get(ctx workflow.Context) (*simpledb.SelectOutput, error) {
	var output simpledb.SelectOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteAttributes(ctx workflow.Context, input *simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error) {
	var output simpledb.BatchDeleteAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.BatchDeleteAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteAttributesAsync(ctx workflow.Context, input *simpledb.BatchDeleteAttributesInput) *SimpleDBBatchDeleteAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.BatchDeleteAttributes", input)
	return &SimpleDBBatchDeleteAttributesFuture{Future: future}
}

func (a *stub) BatchPutAttributes(ctx workflow.Context, input *simpledb.BatchPutAttributesInput) (*simpledb.BatchPutAttributesOutput, error) {
	var output simpledb.BatchPutAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.BatchPutAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchPutAttributesAsync(ctx workflow.Context, input *simpledb.BatchPutAttributesInput) *SimpleDBBatchPutAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.BatchPutAttributes", input)
	return &SimpleDBBatchPutAttributesFuture{Future: future}
}

func (a *stub) CreateDomain(ctx workflow.Context, input *simpledb.CreateDomainInput) (*simpledb.CreateDomainOutput, error) {
	var output simpledb.CreateDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.CreateDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDomainAsync(ctx workflow.Context, input *simpledb.CreateDomainInput) *SimpleDBCreateDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.CreateDomain", input)
	return &SimpleDBCreateDomainFuture{Future: future}
}

func (a *stub) DeleteAttributes(ctx workflow.Context, input *simpledb.DeleteAttributesInput) (*simpledb.DeleteAttributesOutput, error) {
	var output simpledb.DeleteAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.DeleteAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAttributesAsync(ctx workflow.Context, input *simpledb.DeleteAttributesInput) *SimpleDBDeleteAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.DeleteAttributes", input)
	return &SimpleDBDeleteAttributesFuture{Future: future}
}

func (a *stub) DeleteDomain(ctx workflow.Context, input *simpledb.DeleteDomainInput) (*simpledb.DeleteDomainOutput, error) {
	var output simpledb.DeleteDomainOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.DeleteDomain", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDomainAsync(ctx workflow.Context, input *simpledb.DeleteDomainInput) *SimpleDBDeleteDomainFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.DeleteDomain", input)
	return &SimpleDBDeleteDomainFuture{Future: future}
}

func (a *stub) DomainMetadata(ctx workflow.Context, input *simpledb.DomainMetadataInput) (*simpledb.DomainMetadataOutput, error) {
	var output simpledb.DomainMetadataOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.DomainMetadata", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DomainMetadataAsync(ctx workflow.Context, input *simpledb.DomainMetadataInput) *SimpleDBDomainMetadataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.DomainMetadata", input)
	return &SimpleDBDomainMetadataFuture{Future: future}
}

func (a *stub) GetAttributes(ctx workflow.Context, input *simpledb.GetAttributesInput) (*simpledb.GetAttributesOutput, error) {
	var output simpledb.GetAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.GetAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAttributesAsync(ctx workflow.Context, input *simpledb.GetAttributesInput) *SimpleDBGetAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.GetAttributes", input)
	return &SimpleDBGetAttributesFuture{Future: future}
}

func (a *stub) ListDomains(ctx workflow.Context, input *simpledb.ListDomainsInput) (*simpledb.ListDomainsOutput, error) {
	var output simpledb.ListDomainsOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.ListDomains", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDomainsAsync(ctx workflow.Context, input *simpledb.ListDomainsInput) *SimpleDBListDomainsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.ListDomains", input)
	return &SimpleDBListDomainsFuture{Future: future}
}

func (a *stub) PutAttributes(ctx workflow.Context, input *simpledb.PutAttributesInput) (*simpledb.PutAttributesOutput, error) {
	var output simpledb.PutAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.PutAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAttributesAsync(ctx workflow.Context, input *simpledb.PutAttributesInput) *SimpleDBPutAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.PutAttributes", input)
	return &SimpleDBPutAttributesFuture{Future: future}
}

func (a *stub) Select(ctx workflow.Context, input *simpledb.SelectInput) (*simpledb.SelectOutput, error) {
	var output simpledb.SelectOutput
	err := workflow.ExecuteActivity(ctx, "aws.simpledb.Select", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SelectAsync(ctx workflow.Context, input *simpledb.SelectInput) *SimpleDBSelectFuture {
	future := workflow.ExecuteActivity(ctx, "aws.simpledb.Select", input)
	return &SimpleDBSelectFuture{Future: future}
}
