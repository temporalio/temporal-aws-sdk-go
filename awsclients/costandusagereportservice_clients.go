package awsclients

import (
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CostandUsageReportServiceClient interface {
    DeleteReportDefinition(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error)
    DeleteReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) *CostandusagereportserviceDeleteReportDefinitionResult

    DescribeReportDefinitions(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error)
    DescribeReportDefinitionsAsync(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) *CostandusagereportserviceDescribeReportDefinitionsResult

    ModifyReportDefinition(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error)
    ModifyReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) *CostandusagereportserviceModifyReportDefinitionResult

    PutReportDefinition(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error)
    PutReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) *CostandusagereportservicePutReportDefinitionResult
}

type CostandusagereportserviceDeleteReportDefinitionResult struct {
	Result workflow.Future
}

func (r *CostandusagereportserviceDeleteReportDefinitionResult) Get(ctx workflow.Context) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
    var output costandusagereportservice.DeleteReportDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CostandusagereportserviceDescribeReportDefinitionsResult struct {
	Result workflow.Future
}

func (r *CostandusagereportserviceDescribeReportDefinitionsResult) Get(ctx workflow.Context) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
    var output costandusagereportservice.DescribeReportDefinitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CostandusagereportserviceModifyReportDefinitionResult struct {
	Result workflow.Future
}

func (r *CostandusagereportserviceModifyReportDefinitionResult) Get(ctx workflow.Context) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
    var output costandusagereportservice.ModifyReportDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CostandusagereportservicePutReportDefinitionResult struct {
	Result workflow.Future
}

func (r *CostandusagereportservicePutReportDefinitionResult) Get(ctx workflow.Context) (*costandusagereportservice.PutReportDefinitionOutput, error) {
    var output costandusagereportservice.PutReportDefinitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CostandUsageReportServiceStub struct {
    activities awsactivities.CostandUsageReportServiceActivities
}

func NewCostandUsageReportServiceStub() CostandUsageReportServiceClient {
    return &CostandUsageReportServiceStub{}
}

func (a *CostandUsageReportServiceStub) DeleteReportDefinition(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
    var output costandusagereportservice.DeleteReportDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReportDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *CostandUsageReportServiceStub) DeleteReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) *CostandusagereportserviceDeleteReportDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReportDefinition, input)
    return &CostandusagereportserviceDeleteReportDefinitionResult{Result: future}
}

func (a *CostandUsageReportServiceStub) DescribeReportDefinitions(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
    var output costandusagereportservice.DescribeReportDefinitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReportDefinitions, input).Get(ctx, &output)
    return &output, err
}

func (a *CostandUsageReportServiceStub) DescribeReportDefinitionsAsync(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) *CostandusagereportserviceDescribeReportDefinitionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReportDefinitions, input)
    return &CostandusagereportserviceDescribeReportDefinitionsResult{Result: future}
}

func (a *CostandUsageReportServiceStub) ModifyReportDefinition(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
    var output costandusagereportservice.ModifyReportDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyReportDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *CostandUsageReportServiceStub) ModifyReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) *CostandusagereportserviceModifyReportDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyReportDefinition, input)
    return &CostandusagereportserviceModifyReportDefinitionResult{Result: future}
}

func (a *CostandUsageReportServiceStub) PutReportDefinition(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error) {
    var output costandusagereportservice.PutReportDefinitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutReportDefinition, input).Get(ctx, &output)
    return &output, err
}

func (a *CostandUsageReportServiceStub) PutReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) *CostandusagereportservicePutReportDefinitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutReportDefinition, input)
    return &CostandusagereportservicePutReportDefinitionResult{Result: future}
}
