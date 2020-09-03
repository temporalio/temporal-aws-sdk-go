package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mq"
	"go.temporal.io/sdk/workflow"
)

type MQClient interface {
    CreateBroker(ctx workflow.Context, input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error)
    CreateBrokerAsync(ctx workflow.Context, input *mq.CreateBrokerRequest) *MqCreateBrokerResult

    CreateConfiguration(ctx workflow.Context, input *mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error)
    CreateConfigurationAsync(ctx workflow.Context, input *mq.CreateConfigurationRequest) *MqCreateConfigurationResult

    CreateTags(ctx workflow.Context, input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error)
    CreateTagsAsync(ctx workflow.Context, input *mq.CreateTagsInput) *MqCreateTagsResult

    CreateUser(ctx workflow.Context, input *mq.CreateUserRequest) (*mq.CreateUserOutput, error)
    CreateUserAsync(ctx workflow.Context, input *mq.CreateUserRequest) *MqCreateUserResult

    DeleteBroker(ctx workflow.Context, input *mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error)
    DeleteBrokerAsync(ctx workflow.Context, input *mq.DeleteBrokerInput) *MqDeleteBrokerResult

    DeleteTags(ctx workflow.Context, input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error)
    DeleteTagsAsync(ctx workflow.Context, input *mq.DeleteTagsInput) *MqDeleteTagsResult

    DeleteUser(ctx workflow.Context, input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error)
    DeleteUserAsync(ctx workflow.Context, input *mq.DeleteUserInput) *MqDeleteUserResult

    DescribeBroker(ctx workflow.Context, input *mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error)
    DescribeBrokerAsync(ctx workflow.Context, input *mq.DescribeBrokerInput) *MqDescribeBrokerResult

    DescribeBrokerEngineTypes(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error)
    DescribeBrokerEngineTypesAsync(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) *MqDescribeBrokerEngineTypesResult

    DescribeBrokerInstanceOptions(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error)
    DescribeBrokerInstanceOptionsAsync(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) *MqDescribeBrokerInstanceOptionsResult

    DescribeConfiguration(ctx workflow.Context, input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error)
    DescribeConfigurationAsync(ctx workflow.Context, input *mq.DescribeConfigurationInput) *MqDescribeConfigurationResult

    DescribeConfigurationRevision(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error)
    DescribeConfigurationRevisionAsync(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) *MqDescribeConfigurationRevisionResult

    DescribeUser(ctx workflow.Context, input *mq.DescribeUserInput) (*mq.DescribeUserResponse, error)
    DescribeUserAsync(ctx workflow.Context, input *mq.DescribeUserInput) *MqDescribeUserResult

    ListBrokers(ctx workflow.Context, input *mq.ListBrokersInput) (*mq.ListBrokersResponse, error)
    ListBrokersAsync(ctx workflow.Context, input *mq.ListBrokersInput) *MqListBrokersResult

    ListConfigurationRevisions(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error)
    ListConfigurationRevisionsAsync(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) *MqListConfigurationRevisionsResult

    ListConfigurations(ctx workflow.Context, input *mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error)
    ListConfigurationsAsync(ctx workflow.Context, input *mq.ListConfigurationsInput) *MqListConfigurationsResult

    ListTags(ctx workflow.Context, input *mq.ListTagsInput) (*mq.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *mq.ListTagsInput) *MqListTagsResult

    ListUsers(ctx workflow.Context, input *mq.ListUsersInput) (*mq.ListUsersResponse, error)
    ListUsersAsync(ctx workflow.Context, input *mq.ListUsersInput) *MqListUsersResult

    RebootBroker(ctx workflow.Context, input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error)
    RebootBrokerAsync(ctx workflow.Context, input *mq.RebootBrokerInput) *MqRebootBrokerResult

    UpdateBroker(ctx workflow.Context, input *mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error)
    UpdateBrokerAsync(ctx workflow.Context, input *mq.UpdateBrokerRequest) *MqUpdateBrokerResult

    UpdateConfiguration(ctx workflow.Context, input *mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error)
    UpdateConfigurationAsync(ctx workflow.Context, input *mq.UpdateConfigurationRequest) *MqUpdateConfigurationResult

    UpdateUser(ctx workflow.Context, input *mq.UpdateUserRequest) (*mq.UpdateUserOutput, error)
    UpdateUserAsync(ctx workflow.Context, input *mq.UpdateUserRequest) *MqUpdateUserResult
}
type MqCreateBrokerResult struct {
	Result workflow.Future
}

func (r *MqCreateBrokerResult) Get(ctx workflow.Context) (*mq.CreateBrokerResponse, error) {
    var output mq.CreateBrokerResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqCreateConfigurationResult struct {
	Result workflow.Future
}

func (r *MqCreateConfigurationResult) Get(ctx workflow.Context) (*mq.CreateConfigurationResponse, error) {
    var output mq.CreateConfigurationResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqCreateTagsResult struct {
	Result workflow.Future
}

func (r *MqCreateTagsResult) Get(ctx workflow.Context) (*mq.CreateTagsOutput, error) {
    var output mq.CreateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqCreateUserResult struct {
	Result workflow.Future
}

func (r *MqCreateUserResult) Get(ctx workflow.Context) (*mq.CreateUserOutput, error) {
    var output mq.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDeleteBrokerResult struct {
	Result workflow.Future
}

func (r *MqDeleteBrokerResult) Get(ctx workflow.Context) (*mq.DeleteBrokerResponse, error) {
    var output mq.DeleteBrokerResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDeleteTagsResult struct {
	Result workflow.Future
}

func (r *MqDeleteTagsResult) Get(ctx workflow.Context) (*mq.DeleteTagsOutput, error) {
    var output mq.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDeleteUserResult struct {
	Result workflow.Future
}

func (r *MqDeleteUserResult) Get(ctx workflow.Context) (*mq.DeleteUserOutput, error) {
    var output mq.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDescribeBrokerResult struct {
	Result workflow.Future
}

func (r *MqDescribeBrokerResult) Get(ctx workflow.Context) (*mq.DescribeBrokerResponse, error) {
    var output mq.DescribeBrokerResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDescribeBrokerEngineTypesResult struct {
	Result workflow.Future
}

func (r *MqDescribeBrokerEngineTypesResult) Get(ctx workflow.Context) (*mq.DescribeBrokerEngineTypesOutput, error) {
    var output mq.DescribeBrokerEngineTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDescribeBrokerInstanceOptionsResult struct {
	Result workflow.Future
}

func (r *MqDescribeBrokerInstanceOptionsResult) Get(ctx workflow.Context) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
    var output mq.DescribeBrokerInstanceOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDescribeConfigurationResult struct {
	Result workflow.Future
}

func (r *MqDescribeConfigurationResult) Get(ctx workflow.Context) (*mq.DescribeConfigurationOutput, error) {
    var output mq.DescribeConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDescribeConfigurationRevisionResult struct {
	Result workflow.Future
}

func (r *MqDescribeConfigurationRevisionResult) Get(ctx workflow.Context) (*mq.DescribeConfigurationRevisionResponse, error) {
    var output mq.DescribeConfigurationRevisionResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqDescribeUserResult struct {
	Result workflow.Future
}

func (r *MqDescribeUserResult) Get(ctx workflow.Context) (*mq.DescribeUserResponse, error) {
    var output mq.DescribeUserResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqListBrokersResult struct {
	Result workflow.Future
}

func (r *MqListBrokersResult) Get(ctx workflow.Context) (*mq.ListBrokersResponse, error) {
    var output mq.ListBrokersResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqListConfigurationRevisionsResult struct {
	Result workflow.Future
}

func (r *MqListConfigurationRevisionsResult) Get(ctx workflow.Context) (*mq.ListConfigurationRevisionsResponse, error) {
    var output mq.ListConfigurationRevisionsResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqListConfigurationsResult struct {
	Result workflow.Future
}

func (r *MqListConfigurationsResult) Get(ctx workflow.Context) (*mq.ListConfigurationsResponse, error) {
    var output mq.ListConfigurationsResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqListTagsResult struct {
	Result workflow.Future
}

func (r *MqListTagsResult) Get(ctx workflow.Context) (*mq.ListTagsOutput, error) {
    var output mq.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqListUsersResult struct {
	Result workflow.Future
}

func (r *MqListUsersResult) Get(ctx workflow.Context) (*mq.ListUsersResponse, error) {
    var output mq.ListUsersResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqRebootBrokerResult struct {
	Result workflow.Future
}

func (r *MqRebootBrokerResult) Get(ctx workflow.Context) (*mq.RebootBrokerOutput, error) {
    var output mq.RebootBrokerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqUpdateBrokerResult struct {
	Result workflow.Future
}

func (r *MqUpdateBrokerResult) Get(ctx workflow.Context) (*mq.UpdateBrokerResponse, error) {
    var output mq.UpdateBrokerResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqUpdateConfigurationResult struct {
	Result workflow.Future
}

func (r *MqUpdateConfigurationResult) Get(ctx workflow.Context) (*mq.UpdateConfigurationResponse, error) {
    var output mq.UpdateConfigurationResponse
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MqUpdateUserResult struct {
	Result workflow.Future
}

func (r *MqUpdateUserResult) Get(ctx workflow.Context) (*mq.UpdateUserOutput, error) {
    var output mq.UpdateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MQStub struct {
    activities MQClient
}

func NewMQStub() MQClient {
    return &MQStub{}
}

func (a *MQStub) CreateBroker(ctx workflow.Context, input *mq.CreateBrokerRequest) (*mq.CreateBrokerResponse, error) {
    var output mq.CreateBrokerResponse
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBroker, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) CreateConfiguration(ctx workflow.Context, input *mq.CreateConfigurationRequest) (*mq.CreateConfigurationResponse, error) {
    var output mq.CreateConfigurationResponse
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) CreateTags(ctx workflow.Context, input *mq.CreateTagsInput) (*mq.CreateTagsOutput, error) {
    var output mq.CreateTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) CreateUser(ctx workflow.Context, input *mq.CreateUserRequest) (*mq.CreateUserOutput, error) {
    var output mq.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DeleteBroker(ctx workflow.Context, input *mq.DeleteBrokerInput) (*mq.DeleteBrokerResponse, error) {
    var output mq.DeleteBrokerResponse
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBroker, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DeleteTags(ctx workflow.Context, input *mq.DeleteTagsInput) (*mq.DeleteTagsOutput, error) {
    var output mq.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DeleteUser(ctx workflow.Context, input *mq.DeleteUserInput) (*mq.DeleteUserOutput, error) {
    var output mq.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DescribeBroker(ctx workflow.Context, input *mq.DescribeBrokerInput) (*mq.DescribeBrokerResponse, error) {
    var output mq.DescribeBrokerResponse
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBroker, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DescribeBrokerEngineTypes(ctx workflow.Context, input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error) {
    var output mq.DescribeBrokerEngineTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBrokerEngineTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DescribeBrokerInstanceOptions(ctx workflow.Context, input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
    var output mq.DescribeBrokerInstanceOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBrokerInstanceOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DescribeConfiguration(ctx workflow.Context, input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error) {
    var output mq.DescribeConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DescribeConfigurationRevision(ctx workflow.Context, input *mq.DescribeConfigurationRevisionInput) (*mq.DescribeConfigurationRevisionResponse, error) {
    var output mq.DescribeConfigurationRevisionResponse
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) DescribeUser(ctx workflow.Context, input *mq.DescribeUserInput) (*mq.DescribeUserResponse, error) {
    var output mq.DescribeUserResponse
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) ListBrokers(ctx workflow.Context, input *mq.ListBrokersInput) (*mq.ListBrokersResponse, error) {
    var output mq.ListBrokersResponse
    err := workflow.ExecuteActivity(ctx, a.activities.ListBrokers, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) ListConfigurationRevisions(ctx workflow.Context, input *mq.ListConfigurationRevisionsInput) (*mq.ListConfigurationRevisionsResponse, error) {
    var output mq.ListConfigurationRevisionsResponse
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) ListConfigurations(ctx workflow.Context, input *mq.ListConfigurationsInput) (*mq.ListConfigurationsResponse, error) {
    var output mq.ListConfigurationsResponse
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) ListTags(ctx workflow.Context, input *mq.ListTagsInput) (*mq.ListTagsOutput, error) {
    var output mq.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) ListUsers(ctx workflow.Context, input *mq.ListUsersInput) (*mq.ListUsersResponse, error) {
    var output mq.ListUsersResponse
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) RebootBroker(ctx workflow.Context, input *mq.RebootBrokerInput) (*mq.RebootBrokerOutput, error) {
    var output mq.RebootBrokerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootBroker, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) UpdateBroker(ctx workflow.Context, input *mq.UpdateBrokerRequest) (*mq.UpdateBrokerResponse, error) {
    var output mq.UpdateBrokerResponse
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBroker, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) UpdateConfiguration(ctx workflow.Context, input *mq.UpdateConfigurationRequest) (*mq.UpdateConfigurationResponse, error) {
    var output mq.UpdateConfigurationResponse
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *MQStub) UpdateUser(ctx workflow.Context, input *mq.UpdateUserRequest) (*mq.UpdateUserOutput, error) {
    var output mq.UpdateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUser, input).Get(ctx, &output)
    return &output, err
}
