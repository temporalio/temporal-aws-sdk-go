package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/servicequotas/servicequotasiface"
)

type ServiceQuotasActivities struct {
	client servicequotasiface.ServiceQuotasAPI
}

func NewServiceQuotasActivities(client servicequotasiface.ServiceQuotasAPI) *ServiceQuotasActivities {
    return &ServiceQuotasActivities{client: client}
}

func (a *ServiceQuotasActivities) AssociateServiceQuotaTemplate(input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
    return a.client.AssociateServiceQuotaTemplate(input)
}

func (a *ServiceQuotasActivities) DeleteServiceQuotaIncreaseRequestFromTemplate(input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
    return a.client.DeleteServiceQuotaIncreaseRequestFromTemplate(input)
}

func (a *ServiceQuotasActivities) DisassociateServiceQuotaTemplate(input *servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error) {
    return a.client.DisassociateServiceQuotaTemplate(input)
}

func (a *ServiceQuotasActivities) GetAWSDefaultServiceQuota(input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
    return a.client.GetAWSDefaultServiceQuota(input)
}

func (a *ServiceQuotasActivities) GetAssociationForServiceQuotaTemplate(input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
    return a.client.GetAssociationForServiceQuotaTemplate(input)
}

func (a *ServiceQuotasActivities) GetRequestedServiceQuotaChange(input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
    return a.client.GetRequestedServiceQuotaChange(input)
}

func (a *ServiceQuotasActivities) GetServiceQuota(input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error) {
    return a.client.GetServiceQuota(input)
}

func (a *ServiceQuotasActivities) GetServiceQuotaIncreaseRequestFromTemplate(input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
    return a.client.GetServiceQuotaIncreaseRequestFromTemplate(input)
}

func (a *ServiceQuotasActivities) ListAWSDefaultServiceQuotas(input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
    return a.client.ListAWSDefaultServiceQuotas(input)
}

func (a *ServiceQuotasActivities) ListRequestedServiceQuotaChangeHistory(input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
    return a.client.ListRequestedServiceQuotaChangeHistory(input)
}

func (a *ServiceQuotasActivities) ListRequestedServiceQuotaChangeHistoryByQuota(input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
    return a.client.ListRequestedServiceQuotaChangeHistoryByQuota(input)
}

func (a *ServiceQuotasActivities) ListServiceQuotaIncreaseRequestsInTemplate(input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
    return a.client.ListServiceQuotaIncreaseRequestsInTemplate(input)
}

func (a *ServiceQuotasActivities) ListServiceQuotas(input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error) {
    return a.client.ListServiceQuotas(input)
}

func (a *ServiceQuotasActivities) ListServices(input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error) {
    return a.client.ListServices(input)
}

func (a *ServiceQuotasActivities) PutServiceQuotaIncreaseRequestIntoTemplate(input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
    return a.client.PutServiceQuotaIncreaseRequestIntoTemplate(input)
}

func (a *ServiceQuotasActivities) RequestServiceQuotaIncrease(input *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
    return a.client.RequestServiceQuotaIncrease(input)
}
