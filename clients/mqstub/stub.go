// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mqstub

import (
	"github.com/aws/aws-sdk-go/service/mq"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type MQCreateBrokerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQCreateBrokerFuture) Get(ctx workflow.Context) (*mq.CreateBrokerResponse, error) {
	var output mq.CreateBrokerResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQCreateConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQCreateConfigurationFuture) Get(ctx workflow.Context) (*mq.CreateConfigurationResponse, error) {
	var output mq.CreateConfigurationResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQCreateTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQCreateTagsFuture) Get(ctx workflow.Context) (*mq.CreateTagsOutput, error) {
	var output mq.CreateTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQCreateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQCreateUserFuture) Get(ctx workflow.Context) (*mq.CreateUserOutput, error) {
	var output mq.CreateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDeleteBrokerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDeleteBrokerFuture) Get(ctx workflow.Context) (*mq.DeleteBrokerResponse, error) {
	var output mq.DeleteBrokerResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDeleteTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDeleteTagsFuture) Get(ctx workflow.Context) (*mq.DeleteTagsOutput, error) {
	var output mq.DeleteTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDeleteUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDeleteUserFuture) Get(ctx workflow.Context) (*mq.DeleteUserOutput, error) {
	var output mq.DeleteUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDescribeBrokerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDescribeBrokerFuture) Get(ctx workflow.Context) (*mq.DescribeBrokerResponse, error) {
	var output mq.DescribeBrokerResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDescribeBrokerEngineTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDescribeBrokerEngineTypesFuture) Get(ctx workflow.Context) (*mq.DescribeBrokerEngineTypesOutput, error) {
	var output mq.DescribeBrokerEngineTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDescribeBrokerInstanceOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDescribeBrokerInstanceOptionsFuture) Get(ctx workflow.Context) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
	var output mq.DescribeBrokerInstanceOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDescribeConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDescribeConfigurationFuture) Get(ctx workflow.Context) (*mq.DescribeConfigurationOutput, error) {
	var output mq.DescribeConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDescribeConfigurationRevisionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDescribeConfigurationRevisionFuture) Get(ctx workflow.Context) (*mq.DescribeConfigurationRevisionResponse, error) {
	var output mq.DescribeConfigurationRevisionResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQDescribeUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQDescribeUserFuture) Get(ctx workflow.Context) (*mq.DescribeUserResponse, error) {
	var output mq.DescribeUserResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQListBrokersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQListBrokersFuture) Get(ctx workflow.Context) (*mq.ListBrokersResponse, error) {
	var output mq.ListBrokersResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQListConfigurationRevisionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQListConfigurationRevisionsFuture) Get(ctx workflow.Context) (*mq.ListConfigurationRevisionsResponse, error) {
	var output mq.ListConfigurationRevisionsResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQListConfigurationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQListConfigurationsFuture) Get(ctx workflow.Context) (*mq.ListConfigurationsResponse, error) {
	var output mq.ListConfigurationsResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQListTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQListTagsFuture) Get(ctx workflow.Context) (*mq.ListTagsOutput, error) {
	var output mq.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQListUsersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQListUsersFuture) Get(ctx workflow.Context) (*mq.ListUsersResponse, error) {
	var output mq.ListUsersResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQRebootBrokerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQRebootBrokerFuture) Get(ctx workflow.Context) (*mq.RebootBrokerOutput, error) {
	var output mq.RebootBrokerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQUpdateBrokerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQUpdateBrokerFuture) Get(ctx workflow.Context) (*mq.UpdateBrokerResponse, error) {
	var output mq.UpdateBrokerResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQUpdateConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQUpdateConfigurationFuture) Get(ctx workflow.Context) (*mq.UpdateConfigurationResponse, error) {
	var output mq.UpdateConfigurationResponse
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MQUpdateUserFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *MQUpdateUserFuture) Get(ctx workflow.Context) (*mq.UpdateUserOutput, error) {
	var output mq.UpdateUserOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBroker(ctx workflow.Context, input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error) {
	var output mq.CreateBrokerResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.CreateBroker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateBrokerAsync(ctx workflow.Context, input *mq.CreateBrokerRequest) *MQCreateBrokerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.CreateBroker", input)
	return &MQCreateBrokerFuture{Future: future}
}

func (a *stub) CreateConfiguration(ctx workflow.Context, input *mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error) {
	var output mq.CreateConfigurationResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.CreateConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationAsync(ctx workflow.Context, input *mq.CreateConfigurationRequest) *MQCreateConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.CreateConfiguration", input)
	return &MQCreateConfigurationFuture{Future: future}
}

func (a *stub) CreateTags(ctx workflow.Context, input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error) {
	var output mq.CreateTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.CreateTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateTagsAsync(ctx workflow.Context, input *mq.CreateTagsInput) *MQCreateTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.CreateTags", input)
	return &MQCreateTagsFuture{Future: future}
}

func (a *stub) CreateUser(ctx workflow.Context, input *mq.CreateUserRequest) (*mq.CreateUserOutput, error) {
	var output mq.CreateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.CreateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateUserAsync(ctx workflow.Context, input *mq.CreateUserRequest) *MQCreateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.CreateUser", input)
	return &MQCreateUserFuture{Future: future}
}

func (a *stub) DeleteBroker(ctx workflow.Context, input *mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error) {
	var output mq.DeleteBrokerResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.DeleteBroker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteBrokerAsync(ctx workflow.Context, input *mq.DeleteBrokerInput) *MQDeleteBrokerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DeleteBroker", input)
	return &MQDeleteBrokerFuture{Future: future}
}

func (a *stub) DeleteTags(ctx workflow.Context, input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error) {
	var output mq.DeleteTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.DeleteTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteTagsAsync(ctx workflow.Context, input *mq.DeleteTagsInput) *MQDeleteTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DeleteTags", input)
	return &MQDeleteTagsFuture{Future: future}
}

func (a *stub) DeleteUser(ctx workflow.Context, input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error) {
	var output mq.DeleteUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.DeleteUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteUserAsync(ctx workflow.Context, input *mq.DeleteUserInput) *MQDeleteUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DeleteUser", input)
	return &MQDeleteUserFuture{Future: future}
}

func (a *stub) DescribeBroker(ctx workflow.Context, input *mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error) {
	var output mq.DescribeBrokerResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.DescribeBroker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBrokerAsync(ctx workflow.Context, input *mq.DescribeBrokerInput) *MQDescribeBrokerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DescribeBroker", input)
	return &MQDescribeBrokerFuture{Future: future}
}

func (a *stub) DescribeBrokerEngineTypes(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error) {
	var output mq.DescribeBrokerEngineTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.DescribeBrokerEngineTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBrokerEngineTypesAsync(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) *MQDescribeBrokerEngineTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DescribeBrokerEngineTypes", input)
	return &MQDescribeBrokerEngineTypesFuture{Future: future}
}

func (a *stub) DescribeBrokerInstanceOptions(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
	var output mq.DescribeBrokerInstanceOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.DescribeBrokerInstanceOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeBrokerInstanceOptionsAsync(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) *MQDescribeBrokerInstanceOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DescribeBrokerInstanceOptions", input)
	return &MQDescribeBrokerInstanceOptionsFuture{Future: future}
}

func (a *stub) DescribeConfiguration(ctx workflow.Context, input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error) {
	var output mq.DescribeConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.DescribeConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeConfigurationAsync(ctx workflow.Context, input *mq.DescribeConfigurationInput) *MQDescribeConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DescribeConfiguration", input)
	return &MQDescribeConfigurationFuture{Future: future}
}

func (a *stub) DescribeConfigurationRevision(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error) {
	var output mq.DescribeConfigurationRevisionResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.DescribeConfigurationRevision", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeConfigurationRevisionAsync(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) *MQDescribeConfigurationRevisionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DescribeConfigurationRevision", input)
	return &MQDescribeConfigurationRevisionFuture{Future: future}
}

func (a *stub) DescribeUser(ctx workflow.Context, input *mq.DescribeUserInput) (*mq.DescribeUserResponse, error) {
	var output mq.DescribeUserResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.DescribeUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeUserAsync(ctx workflow.Context, input *mq.DescribeUserInput) *MQDescribeUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.DescribeUser", input)
	return &MQDescribeUserFuture{Future: future}
}

func (a *stub) ListBrokers(ctx workflow.Context, input *mq.ListBrokersInput) (*mq.ListBrokersResponse, error) {
	var output mq.ListBrokersResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.ListBrokers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListBrokersAsync(ctx workflow.Context, input *mq.ListBrokersInput) *MQListBrokersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.ListBrokers", input)
	return &MQListBrokersFuture{Future: future}
}

func (a *stub) ListConfigurationRevisions(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error) {
	var output mq.ListConfigurationRevisionsResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.ListConfigurationRevisions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationRevisionsAsync(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) *MQListConfigurationRevisionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.ListConfigurationRevisions", input)
	return &MQListConfigurationRevisionsFuture{Future: future}
}

func (a *stub) ListConfigurations(ctx workflow.Context, input *mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error) {
	var output mq.ListConfigurationsResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.ListConfigurations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationsAsync(ctx workflow.Context, input *mq.ListConfigurationsInput) *MQListConfigurationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.ListConfigurations", input)
	return &MQListConfigurationsFuture{Future: future}
}

func (a *stub) ListTags(ctx workflow.Context, input *mq.ListTagsInput) (*mq.ListTagsOutput, error) {
	var output mq.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsAsync(ctx workflow.Context, input *mq.ListTagsInput) *MQListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.ListTags", input)
	return &MQListTagsFuture{Future: future}
}

func (a *stub) ListUsers(ctx workflow.Context, input *mq.ListUsersInput) (*mq.ListUsersResponse, error) {
	var output mq.ListUsersResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.ListUsers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListUsersAsync(ctx workflow.Context, input *mq.ListUsersInput) *MQListUsersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.ListUsers", input)
	return &MQListUsersFuture{Future: future}
}

func (a *stub) RebootBroker(ctx workflow.Context, input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error) {
	var output mq.RebootBrokerOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.RebootBroker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RebootBrokerAsync(ctx workflow.Context, input *mq.RebootBrokerInput) *MQRebootBrokerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.RebootBroker", input)
	return &MQRebootBrokerFuture{Future: future}
}

func (a *stub) UpdateBroker(ctx workflow.Context, input *mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error) {
	var output mq.UpdateBrokerResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.UpdateBroker", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateBrokerAsync(ctx workflow.Context, input *mq.UpdateBrokerRequest) *MQUpdateBrokerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.UpdateBroker", input)
	return &MQUpdateBrokerFuture{Future: future}
}

func (a *stub) UpdateConfiguration(ctx workflow.Context, input *mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error) {
	var output mq.UpdateConfigurationResponse
	err := workflow.ExecuteActivity(ctx, "aws.mq.UpdateConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateConfigurationAsync(ctx workflow.Context, input *mq.UpdateConfigurationRequest) *MQUpdateConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.UpdateConfiguration", input)
	return &MQUpdateConfigurationFuture{Future: future}
}

func (a *stub) UpdateUser(ctx workflow.Context, input *mq.UpdateUserRequest) (*mq.UpdateUserOutput, error) {
	var output mq.UpdateUserOutput
	err := workflow.ExecuteActivity(ctx, "aws.mq.UpdateUser", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateUserAsync(ctx workflow.Context, input *mq.UpdateUserRequest) *MQUpdateUserFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mq.UpdateUser", input)
	return &MQUpdateUserFuture{Future: future}
}
