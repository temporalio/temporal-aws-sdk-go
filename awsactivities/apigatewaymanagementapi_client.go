package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"go.temporal.io/sdk/workflow"
)

type ApiGatewayManagementApiClient interface {
    DeleteConnection(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error)
    DeleteConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) *ApigatewaymanagementapiDeleteConnectionResult

    GetConnection(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error)
    GetConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) *ApigatewaymanagementapiGetConnectionResult

    PostToConnection(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error)
    PostToConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) *ApigatewaymanagementapiPostToConnectionResult
}
type ApigatewaymanagementapiDeleteConnectionResult struct {
	Result workflow.Future
}

func (r *ApigatewaymanagementapiDeleteConnectionResult) Get(ctx workflow.Context) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
    var output apigatewaymanagementapi.DeleteConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApigatewaymanagementapiGetConnectionResult struct {
	Result workflow.Future
}

func (r *ApigatewaymanagementapiGetConnectionResult) Get(ctx workflow.Context) (*apigatewaymanagementapi.GetConnectionOutput, error) {
    var output apigatewaymanagementapi.GetConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApigatewaymanagementapiPostToConnectionResult struct {
	Result workflow.Future
}

func (r *ApigatewaymanagementapiPostToConnectionResult) Get(ctx workflow.Context) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
    var output apigatewaymanagementapi.PostToConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ApiGatewayManagementApiStub struct {
    activities ApiGatewayManagementApiClient
}

func NewApiGatewayManagementApiStub() ApiGatewayManagementApiClient {
    return &ApiGatewayManagementApiStub{}
}

func (a *ApiGatewayManagementApiStub) DeleteConnection(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
    var output apigatewaymanagementapi.DeleteConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ApiGatewayManagementApiStub) GetConnection(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error) {
    var output apigatewaymanagementapi.GetConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ApiGatewayManagementApiStub) PostToConnection(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
    var output apigatewaymanagementapi.PostToConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PostToConnection, input).Get(ctx, &output)
    return &output, err
}
