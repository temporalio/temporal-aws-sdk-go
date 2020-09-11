package awsclients

import (
	"github.com/aws/aws-sdk-go/service/simpledb"
	"go.temporal.io/sdk/workflow"
)

type SimpleDBClient interface {
       BatchDeleteAttributes(ctx workflow.Context, input *simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error)
       BatchDeleteAttributesAsync(ctx workflow.Context, input *simpledb.BatchDeleteAttributesInput) *SimpledbBatchDeleteAttributesResult

       BatchPutAttributes(ctx workflow.Context, input *simpledb.BatchPutAttributesInput) (*simpledb.BatchPutAttributesOutput, error)
       BatchPutAttributesAsync(ctx workflow.Context, input *simpledb.BatchPutAttributesInput) *SimpledbBatchPutAttributesResult

       CreateDomain(ctx workflow.Context, input *simpledb.CreateDomainInput) (*simpledb.CreateDomainOutput, error)
       CreateDomainAsync(ctx workflow.Context, input *simpledb.CreateDomainInput) *SimpledbCreateDomainResult

       DeleteAttributes(ctx workflow.Context, input *simpledb.DeleteAttributesInput) (*simpledb.DeleteAttributesOutput, error)
       DeleteAttributesAsync(ctx workflow.Context, input *simpledb.DeleteAttributesInput) *SimpledbDeleteAttributesResult

       DeleteDomain(ctx workflow.Context, input *simpledb.DeleteDomainInput) (*simpledb.DeleteDomainOutput, error)
       DeleteDomainAsync(ctx workflow.Context, input *simpledb.DeleteDomainInput) *SimpledbDeleteDomainResult

       DomainMetadata(ctx workflow.Context, input *simpledb.DomainMetadataInput) (*simpledb.DomainMetadataOutput, error)
       DomainMetadataAsync(ctx workflow.Context, input *simpledb.DomainMetadataInput) *SimpledbDomainMetadataResult

       GetAttributes(ctx workflow.Context, input *simpledb.GetAttributesInput) (*simpledb.GetAttributesOutput, error)
       GetAttributesAsync(ctx workflow.Context, input *simpledb.GetAttributesInput) *SimpledbGetAttributesResult

       ListDomains(ctx workflow.Context, input *simpledb.ListDomainsInput) (*simpledb.ListDomainsOutput, error)
       ListDomainsAsync(ctx workflow.Context, input *simpledb.ListDomainsInput) *SimpledbListDomainsResult

       PutAttributes(ctx workflow.Context, input *simpledb.PutAttributesInput) (*simpledb.PutAttributesOutput, error)
       PutAttributesAsync(ctx workflow.Context, input *simpledb.PutAttributesInput) *SimpledbPutAttributesResult

       Select(ctx workflow.Context, input *simpledb.SelectInput) (*simpledb.SelectOutput, error)
       SelectAsync(ctx workflow.Context, input *simpledb.SelectInput) *SimpledbSelectResult
}

type SimpleDBStub struct {
}

func NewSimpleDBStub() SimpleDBClient {
    return &SimpleDBStub{}
}

type SimpledbBatchDeleteAttributesResult struct {
	Result workflow.Future
}

func (r *SimpledbBatchDeleteAttributesResult) Get(ctx workflow.Context) (*simpledb.BatchDeleteAttributesOutput, error) {
    var output simpledb.BatchDeleteAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbBatchPutAttributesResult struct {
	Result workflow.Future
}

func (r *SimpledbBatchPutAttributesResult) Get(ctx workflow.Context) (*simpledb.BatchPutAttributesOutput, error) {
    var output simpledb.BatchPutAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbCreateDomainResult struct {
	Result workflow.Future
}

func (r *SimpledbCreateDomainResult) Get(ctx workflow.Context) (*simpledb.CreateDomainOutput, error) {
    var output simpledb.CreateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbDeleteAttributesResult struct {
	Result workflow.Future
}

func (r *SimpledbDeleteAttributesResult) Get(ctx workflow.Context) (*simpledb.DeleteAttributesOutput, error) {
    var output simpledb.DeleteAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbDeleteDomainResult struct {
	Result workflow.Future
}

func (r *SimpledbDeleteDomainResult) Get(ctx workflow.Context) (*simpledb.DeleteDomainOutput, error) {
    var output simpledb.DeleteDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbDomainMetadataResult struct {
	Result workflow.Future
}

func (r *SimpledbDomainMetadataResult) Get(ctx workflow.Context) (*simpledb.DomainMetadataOutput, error) {
    var output simpledb.DomainMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbGetAttributesResult struct {
	Result workflow.Future
}

func (r *SimpledbGetAttributesResult) Get(ctx workflow.Context) (*simpledb.GetAttributesOutput, error) {
    var output simpledb.GetAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbListDomainsResult struct {
	Result workflow.Future
}

func (r *SimpledbListDomainsResult) Get(ctx workflow.Context) (*simpledb.ListDomainsOutput, error) {
    var output simpledb.ListDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbPutAttributesResult struct {
	Result workflow.Future
}

func (r *SimpledbPutAttributesResult) Get(ctx workflow.Context) (*simpledb.PutAttributesOutput, error) {
    var output simpledb.PutAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SimpledbSelectResult struct {
	Result workflow.Future
}

func (r *SimpledbSelectResult) Get(ctx workflow.Context) (*simpledb.SelectOutput, error) {
    var output simpledb.SelectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) BatchDeleteAttributes(ctx workflow.Context, input *simpledb.BatchDeleteAttributesInput) (*simpledb.BatchDeleteAttributesOutput, error) {
    var output simpledb.BatchDeleteAttributesOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.BatchDeleteAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) BatchDeleteAttributesAsync(ctx workflow.Context, input *simpledb.BatchDeleteAttributesInput) *SimpledbBatchDeleteAttributesResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.BatchDeleteAttributes", input)
    return &SimpledbBatchDeleteAttributesResult{Result: future}
}

func (a *SimpleDBStub) BatchPutAttributes(ctx workflow.Context, input *simpledb.BatchPutAttributesInput) (*simpledb.BatchPutAttributesOutput, error) {
    var output simpledb.BatchPutAttributesOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.BatchPutAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) BatchPutAttributesAsync(ctx workflow.Context, input *simpledb.BatchPutAttributesInput) *SimpledbBatchPutAttributesResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.BatchPutAttributes", input)
    return &SimpledbBatchPutAttributesResult{Result: future}
}

func (a *SimpleDBStub) CreateDomain(ctx workflow.Context, input *simpledb.CreateDomainInput) (*simpledb.CreateDomainOutput, error) {
    var output simpledb.CreateDomainOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.CreateDomain", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) CreateDomainAsync(ctx workflow.Context, input *simpledb.CreateDomainInput) *SimpledbCreateDomainResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.CreateDomain", input)
    return &SimpledbCreateDomainResult{Result: future}
}

func (a *SimpleDBStub) DeleteAttributes(ctx workflow.Context, input *simpledb.DeleteAttributesInput) (*simpledb.DeleteAttributesOutput, error) {
    var output simpledb.DeleteAttributesOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.DeleteAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) DeleteAttributesAsync(ctx workflow.Context, input *simpledb.DeleteAttributesInput) *SimpledbDeleteAttributesResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.DeleteAttributes", input)
    return &SimpledbDeleteAttributesResult{Result: future}
}

func (a *SimpleDBStub) DeleteDomain(ctx workflow.Context, input *simpledb.DeleteDomainInput) (*simpledb.DeleteDomainOutput, error) {
    var output simpledb.DeleteDomainOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.DeleteDomain", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) DeleteDomainAsync(ctx workflow.Context, input *simpledb.DeleteDomainInput) *SimpledbDeleteDomainResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.DeleteDomain", input)
    return &SimpledbDeleteDomainResult{Result: future}
}

func (a *SimpleDBStub) DomainMetadata(ctx workflow.Context, input *simpledb.DomainMetadataInput) (*simpledb.DomainMetadataOutput, error) {
    var output simpledb.DomainMetadataOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.DomainMetadata", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) DomainMetadataAsync(ctx workflow.Context, input *simpledb.DomainMetadataInput) *SimpledbDomainMetadataResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.DomainMetadata", input)
    return &SimpledbDomainMetadataResult{Result: future}
}

func (a *SimpleDBStub) GetAttributes(ctx workflow.Context, input *simpledb.GetAttributesInput) (*simpledb.GetAttributesOutput, error) {
    var output simpledb.GetAttributesOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.GetAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) GetAttributesAsync(ctx workflow.Context, input *simpledb.GetAttributesInput) *SimpledbGetAttributesResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.GetAttributes", input)
    return &SimpledbGetAttributesResult{Result: future}
}

func (a *SimpleDBStub) ListDomains(ctx workflow.Context, input *simpledb.ListDomainsInput) (*simpledb.ListDomainsOutput, error) {
    var output simpledb.ListDomainsOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.ListDomains", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) ListDomainsAsync(ctx workflow.Context, input *simpledb.ListDomainsInput) *SimpledbListDomainsResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.ListDomains", input)
    return &SimpledbListDomainsResult{Result: future}
}

func (a *SimpleDBStub) PutAttributes(ctx workflow.Context, input *simpledb.PutAttributesInput) (*simpledb.PutAttributesOutput, error) {
    var output simpledb.PutAttributesOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.PutAttributes", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) PutAttributesAsync(ctx workflow.Context, input *simpledb.PutAttributesInput) *SimpledbPutAttributesResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.PutAttributes", input)
    return &SimpledbPutAttributesResult{Result: future}
}

func (a *SimpleDBStub) Select(ctx workflow.Context, input *simpledb.SelectInput) (*simpledb.SelectOutput, error) {
    var output simpledb.SelectOutput
    err := workflow.ExecuteActivity(ctx, "SimpleDB.Select", input).Get(ctx, &output)
    return &output, err
}

func (a *SimpleDBStub) SelectAsync(ctx workflow.Context, input *simpledb.SelectInput) *SimpledbSelectResult {
    future := workflow.ExecuteActivity(ctx, "SimpleDB.Select", input)
    return &SimpledbSelectResult{Result: future}
}
