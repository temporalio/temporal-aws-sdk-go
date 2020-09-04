package awsclients

import (
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ServiceQuotasClient interface {
    AssociateServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error)
    AssociateServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) *ServicequotasAssociateServiceQuotaTemplateResult

    DeleteServiceQuotaIncreaseRequestFromTemplate(ctx workflow.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error)
    DeleteServiceQuotaIncreaseRequestFromTemplateAsync(ctx workflow.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) *ServicequotasDeleteServiceQuotaIncreaseRequestFromTemplateResult

    DisassociateServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error)
    DisassociateServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) *ServicequotasDisassociateServiceQuotaTemplateResult

    GetAWSDefaultServiceQuota(ctx workflow.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error)
    GetAWSDefaultServiceQuotaAsync(ctx workflow.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) *ServicequotasGetAWSDefaultServiceQuotaResult

    GetAssociationForServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error)
    GetAssociationForServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) *ServicequotasGetAssociationForServiceQuotaTemplateResult

    GetRequestedServiceQuotaChange(ctx workflow.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error)
    GetRequestedServiceQuotaChangeAsync(ctx workflow.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) *ServicequotasGetRequestedServiceQuotaChangeResult

    GetServiceQuota(ctx workflow.Context, input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error)
    GetServiceQuotaAsync(ctx workflow.Context, input *servicequotas.GetServiceQuotaInput) *ServicequotasGetServiceQuotaResult

    GetServiceQuotaIncreaseRequestFromTemplate(ctx workflow.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error)
    GetServiceQuotaIncreaseRequestFromTemplateAsync(ctx workflow.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) *ServicequotasGetServiceQuotaIncreaseRequestFromTemplateResult

    ListAWSDefaultServiceQuotas(ctx workflow.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error)
    ListAWSDefaultServiceQuotasAsync(ctx workflow.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) *ServicequotasListAWSDefaultServiceQuotasResult

    ListRequestedServiceQuotaChangeHistory(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error)
    ListRequestedServiceQuotaChangeHistoryAsync(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) *ServicequotasListRequestedServiceQuotaChangeHistoryResult

    ListRequestedServiceQuotaChangeHistoryByQuota(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
    ListRequestedServiceQuotaChangeHistoryByQuotaAsync(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) *ServicequotasListRequestedServiceQuotaChangeHistoryByQuotaResult

    ListServiceQuotaIncreaseRequestsInTemplate(ctx workflow.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error)
    ListServiceQuotaIncreaseRequestsInTemplateAsync(ctx workflow.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) *ServicequotasListServiceQuotaIncreaseRequestsInTemplateResult

    ListServiceQuotas(ctx workflow.Context, input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error)
    ListServiceQuotasAsync(ctx workflow.Context, input *servicequotas.ListServiceQuotasInput) *ServicequotasListServiceQuotasResult

    ListServices(ctx workflow.Context, input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error)
    ListServicesAsync(ctx workflow.Context, input *servicequotas.ListServicesInput) *ServicequotasListServicesResult

    PutServiceQuotaIncreaseRequestIntoTemplate(ctx workflow.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error)
    PutServiceQuotaIncreaseRequestIntoTemplateAsync(ctx workflow.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) *ServicequotasPutServiceQuotaIncreaseRequestIntoTemplateResult

    RequestServiceQuotaIncrease(ctx workflow.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error)
    RequestServiceQuotaIncreaseAsync(ctx workflow.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) *ServicequotasRequestServiceQuotaIncreaseResult
}
type ServicequotasAssociateServiceQuotaTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasAssociateServiceQuotaTemplateResult) Get(ctx workflow.Context) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
    var output servicequotas.AssociateServiceQuotaTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasDeleteServiceQuotaIncreaseRequestFromTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasDeleteServiceQuotaIncreaseRequestFromTemplateResult) Get(ctx workflow.Context) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
    var output servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasDisassociateServiceQuotaTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasDisassociateServiceQuotaTemplateResult) Get(ctx workflow.Context) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error) {
    var output servicequotas.DisassociateServiceQuotaTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasGetAWSDefaultServiceQuotaResult struct {
	Result workflow.Future
}

func (r *ServicequotasGetAWSDefaultServiceQuotaResult) Get(ctx workflow.Context) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
    var output servicequotas.GetAWSDefaultServiceQuotaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasGetAssociationForServiceQuotaTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasGetAssociationForServiceQuotaTemplateResult) Get(ctx workflow.Context) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
    var output servicequotas.GetAssociationForServiceQuotaTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasGetRequestedServiceQuotaChangeResult struct {
	Result workflow.Future
}

func (r *ServicequotasGetRequestedServiceQuotaChangeResult) Get(ctx workflow.Context) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
    var output servicequotas.GetRequestedServiceQuotaChangeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasGetServiceQuotaResult struct {
	Result workflow.Future
}

func (r *ServicequotasGetServiceQuotaResult) Get(ctx workflow.Context) (*servicequotas.GetServiceQuotaOutput, error) {
    var output servicequotas.GetServiceQuotaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasGetServiceQuotaIncreaseRequestFromTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasGetServiceQuotaIncreaseRequestFromTemplateResult) Get(ctx workflow.Context) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
    var output servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasListAWSDefaultServiceQuotasResult struct {
	Result workflow.Future
}

func (r *ServicequotasListAWSDefaultServiceQuotasResult) Get(ctx workflow.Context) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
    var output servicequotas.ListAWSDefaultServiceQuotasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasListRequestedServiceQuotaChangeHistoryResult struct {
	Result workflow.Future
}

func (r *ServicequotasListRequestedServiceQuotaChangeHistoryResult) Get(ctx workflow.Context) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
    var output servicequotas.ListRequestedServiceQuotaChangeHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasListRequestedServiceQuotaChangeHistoryByQuotaResult struct {
	Result workflow.Future
}

func (r *ServicequotasListRequestedServiceQuotaChangeHistoryByQuotaResult) Get(ctx workflow.Context) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
    var output servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasListServiceQuotaIncreaseRequestsInTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasListServiceQuotaIncreaseRequestsInTemplateResult) Get(ctx workflow.Context) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
    var output servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasListServiceQuotasResult struct {
	Result workflow.Future
}

func (r *ServicequotasListServiceQuotasResult) Get(ctx workflow.Context) (*servicequotas.ListServiceQuotasOutput, error) {
    var output servicequotas.ListServiceQuotasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasListServicesResult struct {
	Result workflow.Future
}

func (r *ServicequotasListServicesResult) Get(ctx workflow.Context) (*servicequotas.ListServicesOutput, error) {
    var output servicequotas.ListServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasPutServiceQuotaIncreaseRequestIntoTemplateResult struct {
	Result workflow.Future
}

func (r *ServicequotasPutServiceQuotaIncreaseRequestIntoTemplateResult) Get(ctx workflow.Context) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
    var output servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicequotasRequestServiceQuotaIncreaseResult struct {
	Result workflow.Future
}

func (r *ServicequotasRequestServiceQuotaIncreaseResult) Get(ctx workflow.Context) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
    var output servicequotas.RequestServiceQuotaIncreaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ServiceQuotasStub struct {
    activities awsactivities.ServiceQuotasActivities
}

func NewServiceQuotasStub() ServiceQuotasClient {
    return &ServiceQuotasStub{}
}
func (a *ServiceQuotasStub) AssociateServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
    var output servicequotas.AssociateServiceQuotaTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateServiceQuotaTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) AssociateServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) *ServicequotasAssociateServiceQuotaTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateServiceQuotaTemplate, input)
    return &ServicequotasAssociateServiceQuotaTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) DeleteServiceQuotaIncreaseRequestFromTemplate(ctx workflow.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
    var output servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServiceQuotaIncreaseRequestFromTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) DeleteServiceQuotaIncreaseRequestFromTemplateAsync(ctx workflow.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) *ServicequotasDeleteServiceQuotaIncreaseRequestFromTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteServiceQuotaIncreaseRequestFromTemplate, input)
    return &ServicequotasDeleteServiceQuotaIncreaseRequestFromTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) DisassociateServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error) {
    var output servicequotas.DisassociateServiceQuotaTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateServiceQuotaTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) DisassociateServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) *ServicequotasDisassociateServiceQuotaTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateServiceQuotaTemplate, input)
    return &ServicequotasDisassociateServiceQuotaTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) GetAWSDefaultServiceQuota(ctx workflow.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
    var output servicequotas.GetAWSDefaultServiceQuotaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAWSDefaultServiceQuota, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) GetAWSDefaultServiceQuotaAsync(ctx workflow.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) *ServicequotasGetAWSDefaultServiceQuotaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAWSDefaultServiceQuota, input)
    return &ServicequotasGetAWSDefaultServiceQuotaResult{Result: future}
}
func (a *ServiceQuotasStub) GetAssociationForServiceQuotaTemplate(ctx workflow.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
    var output servicequotas.GetAssociationForServiceQuotaTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAssociationForServiceQuotaTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) GetAssociationForServiceQuotaTemplateAsync(ctx workflow.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) *ServicequotasGetAssociationForServiceQuotaTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAssociationForServiceQuotaTemplate, input)
    return &ServicequotasGetAssociationForServiceQuotaTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) GetRequestedServiceQuotaChange(ctx workflow.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
    var output servicequotas.GetRequestedServiceQuotaChangeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRequestedServiceQuotaChange, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) GetRequestedServiceQuotaChangeAsync(ctx workflow.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) *ServicequotasGetRequestedServiceQuotaChangeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRequestedServiceQuotaChange, input)
    return &ServicequotasGetRequestedServiceQuotaChangeResult{Result: future}
}
func (a *ServiceQuotasStub) GetServiceQuota(ctx workflow.Context, input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error) {
    var output servicequotas.GetServiceQuotaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceQuota, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) GetServiceQuotaAsync(ctx workflow.Context, input *servicequotas.GetServiceQuotaInput) *ServicequotasGetServiceQuotaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetServiceQuota, input)
    return &ServicequotasGetServiceQuotaResult{Result: future}
}
func (a *ServiceQuotasStub) GetServiceQuotaIncreaseRequestFromTemplate(ctx workflow.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
    var output servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetServiceQuotaIncreaseRequestFromTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) GetServiceQuotaIncreaseRequestFromTemplateAsync(ctx workflow.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) *ServicequotasGetServiceQuotaIncreaseRequestFromTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetServiceQuotaIncreaseRequestFromTemplate, input)
    return &ServicequotasGetServiceQuotaIncreaseRequestFromTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) ListAWSDefaultServiceQuotas(ctx workflow.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
    var output servicequotas.ListAWSDefaultServiceQuotasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAWSDefaultServiceQuotas, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) ListAWSDefaultServiceQuotasAsync(ctx workflow.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) *ServicequotasListAWSDefaultServiceQuotasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAWSDefaultServiceQuotas, input)
    return &ServicequotasListAWSDefaultServiceQuotasResult{Result: future}
}
func (a *ServiceQuotasStub) ListRequestedServiceQuotaChangeHistory(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
    var output servicequotas.ListRequestedServiceQuotaChangeHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRequestedServiceQuotaChangeHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) ListRequestedServiceQuotaChangeHistoryAsync(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) *ServicequotasListRequestedServiceQuotaChangeHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRequestedServiceQuotaChangeHistory, input)
    return &ServicequotasListRequestedServiceQuotaChangeHistoryResult{Result: future}
}
func (a *ServiceQuotasStub) ListRequestedServiceQuotaChangeHistoryByQuota(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
    var output servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRequestedServiceQuotaChangeHistoryByQuota, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) ListRequestedServiceQuotaChangeHistoryByQuotaAsync(ctx workflow.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) *ServicequotasListRequestedServiceQuotaChangeHistoryByQuotaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRequestedServiceQuotaChangeHistoryByQuota, input)
    return &ServicequotasListRequestedServiceQuotaChangeHistoryByQuotaResult{Result: future}
}
func (a *ServiceQuotasStub) ListServiceQuotaIncreaseRequestsInTemplate(ctx workflow.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
    var output servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServiceQuotaIncreaseRequestsInTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) ListServiceQuotaIncreaseRequestsInTemplateAsync(ctx workflow.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) *ServicequotasListServiceQuotaIncreaseRequestsInTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServiceQuotaIncreaseRequestsInTemplate, input)
    return &ServicequotasListServiceQuotaIncreaseRequestsInTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) ListServiceQuotas(ctx workflow.Context, input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error) {
    var output servicequotas.ListServiceQuotasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServiceQuotas, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) ListServiceQuotasAsync(ctx workflow.Context, input *servicequotas.ListServiceQuotasInput) *ServicequotasListServiceQuotasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServiceQuotas, input)
    return &ServicequotasListServiceQuotasResult{Result: future}
}
func (a *ServiceQuotasStub) ListServices(ctx workflow.Context, input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error) {
    var output servicequotas.ListServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServices, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) ListServicesAsync(ctx workflow.Context, input *servicequotas.ListServicesInput) *ServicequotasListServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServices, input)
    return &ServicequotasListServicesResult{Result: future}
}
func (a *ServiceQuotasStub) PutServiceQuotaIncreaseRequestIntoTemplate(ctx workflow.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
    var output servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutServiceQuotaIncreaseRequestIntoTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) PutServiceQuotaIncreaseRequestIntoTemplateAsync(ctx workflow.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) *ServicequotasPutServiceQuotaIncreaseRequestIntoTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutServiceQuotaIncreaseRequestIntoTemplate, input)
    return &ServicequotasPutServiceQuotaIncreaseRequestIntoTemplateResult{Result: future}
}
func (a *ServiceQuotasStub) RequestServiceQuotaIncrease(ctx workflow.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
    var output servicequotas.RequestServiceQuotaIncreaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestServiceQuotaIncrease, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceQuotasStub) RequestServiceQuotaIncreaseAsync(ctx workflow.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) *ServicequotasRequestServiceQuotaIncreaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RequestServiceQuotaIncrease, input)
    return &ServicequotasRequestServiceQuotaIncreaseResult{Result: future}
}
