// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package servicequotas

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/aws/aws-sdk-go/service/servicequotas/servicequotasiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client servicequotasiface.ServiceQuotasAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := servicequotas.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (servicequotasiface.ServiceQuotasAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return servicequotas.New(sess), nil
}

func (a *Activities) AssociateServiceQuotaTemplate(ctx context.Context, input *servicequotas.AssociateServiceQuotaTemplateInput) (*servicequotas.AssociateServiceQuotaTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateServiceQuotaTemplateWithContext(ctx, input)
}

func (a *Activities) DeleteServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, input *servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input)
}

func (a *Activities) DisassociateServiceQuotaTemplate(ctx context.Context, input *servicequotas.DisassociateServiceQuotaTemplateInput) (*servicequotas.DisassociateServiceQuotaTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateServiceQuotaTemplateWithContext(ctx, input)
}

func (a *Activities) GetAWSDefaultServiceQuota(ctx context.Context, input *servicequotas.GetAWSDefaultServiceQuotaInput) (*servicequotas.GetAWSDefaultServiceQuotaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAWSDefaultServiceQuotaWithContext(ctx, input)
}

func (a *Activities) GetAssociationForServiceQuotaTemplate(ctx context.Context, input *servicequotas.GetAssociationForServiceQuotaTemplateInput) (*servicequotas.GetAssociationForServiceQuotaTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAssociationForServiceQuotaTemplateWithContext(ctx, input)
}

func (a *Activities) GetRequestedServiceQuotaChange(ctx context.Context, input *servicequotas.GetRequestedServiceQuotaChangeInput) (*servicequotas.GetRequestedServiceQuotaChangeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRequestedServiceQuotaChangeWithContext(ctx, input)
}

func (a *Activities) GetServiceQuota(ctx context.Context, input *servicequotas.GetServiceQuotaInput) (*servicequotas.GetServiceQuotaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetServiceQuotaWithContext(ctx, input)
}

func (a *Activities) GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, input *servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput) (*servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetServiceQuotaIncreaseRequestFromTemplateWithContext(ctx, input)
}

func (a *Activities) ListAWSDefaultServiceQuotas(ctx context.Context, input *servicequotas.ListAWSDefaultServiceQuotasInput) (*servicequotas.ListAWSDefaultServiceQuotasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAWSDefaultServiceQuotasWithContext(ctx, input)
}

func (a *Activities) ListRequestedServiceQuotaChangeHistory(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRequestedServiceQuotaChangeHistoryWithContext(ctx, input)
}

func (a *Activities) ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, input *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput) (*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRequestedServiceQuotaChangeHistoryByQuotaWithContext(ctx, input)
}

func (a *Activities) ListServiceQuotaIncreaseRequestsInTemplate(ctx context.Context, input *servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput) (*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServiceQuotaIncreaseRequestsInTemplateWithContext(ctx, input)
}

func (a *Activities) ListServiceQuotas(ctx context.Context, input *servicequotas.ListServiceQuotasInput) (*servicequotas.ListServiceQuotasOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServiceQuotasWithContext(ctx, input)
}

func (a *Activities) ListServices(ctx context.Context, input *servicequotas.ListServicesInput) (*servicequotas.ListServicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListServicesWithContext(ctx, input)
}

func (a *Activities) PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, input *servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput) (*servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutServiceQuotaIncreaseRequestIntoTemplateWithContext(ctx, input)
}

func (a *Activities) RequestServiceQuotaIncrease(ctx context.Context, input *servicequotas.RequestServiceQuotaIncreaseInput) (*servicequotas.RequestServiceQuotaIncreaseOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RequestServiceQuotaIncreaseWithContext(ctx, input)
}
