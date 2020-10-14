// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package pinpointemail

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointemail/pinpointemailiface"
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
	client pinpointemailiface.PinpointEmailAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := pinpointemail.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (pinpointemailiface.PinpointEmailAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return pinpointemail.New(sess), nil
}

func (a *Activities) CreateConfigurationSet(ctx context.Context, input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationSetWithContext(ctx, input)
}

func (a *Activities) CreateConfigurationSetEventDestination(ctx context.Context, input *pinpointemail.CreateConfigurationSetEventDestinationInput) (*pinpointemail.CreateConfigurationSetEventDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *Activities) CreateDedicatedIpPool(ctx context.Context, input *pinpointemail.CreateDedicatedIpPoolInput) (*pinpointemail.CreateDedicatedIpPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDedicatedIpPoolWithContext(ctx, input)
}

func (a *Activities) CreateDeliverabilityTestReport(ctx context.Context, input *pinpointemail.CreateDeliverabilityTestReportInput) (*pinpointemail.CreateDeliverabilityTestReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDeliverabilityTestReportWithContext(ctx, input)
}

func (a *Activities) CreateEmailIdentity(ctx context.Context, input *pinpointemail.CreateEmailIdentityInput) (*pinpointemail.CreateEmailIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEmailIdentityWithContext(ctx, input)
}

func (a *Activities) DeleteConfigurationSet(ctx context.Context, input *pinpointemail.DeleteConfigurationSetInput) (*pinpointemail.DeleteConfigurationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConfigurationSetWithContext(ctx, input)
}

func (a *Activities) DeleteConfigurationSetEventDestination(ctx context.Context, input *pinpointemail.DeleteConfigurationSetEventDestinationInput) (*pinpointemail.DeleteConfigurationSetEventDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConfigurationSetEventDestinationWithContext(ctx, input)
}

func (a *Activities) DeleteDedicatedIpPool(ctx context.Context, input *pinpointemail.DeleteDedicatedIpPoolInput) (*pinpointemail.DeleteDedicatedIpPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDedicatedIpPoolWithContext(ctx, input)
}

func (a *Activities) DeleteEmailIdentity(ctx context.Context, input *pinpointemail.DeleteEmailIdentityInput) (*pinpointemail.DeleteEmailIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEmailIdentityWithContext(ctx, input)
}

func (a *Activities) GetAccount(ctx context.Context, input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetAccountWithContext(ctx, input)
}

func (a *Activities) GetBlacklistReports(ctx context.Context, input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetBlacklistReportsWithContext(ctx, input)
}

func (a *Activities) GetConfigurationSet(ctx context.Context, input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConfigurationSetWithContext(ctx, input)
}

func (a *Activities) GetConfigurationSetEventDestinations(ctx context.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetConfigurationSetEventDestinationsWithContext(ctx, input)
}

func (a *Activities) GetDedicatedIp(ctx context.Context, input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDedicatedIpWithContext(ctx, input)
}

func (a *Activities) GetDedicatedIps(ctx context.Context, input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDedicatedIpsWithContext(ctx, input)
}

func (a *Activities) GetDeliverabilityDashboardOptions(ctx context.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeliverabilityDashboardOptionsWithContext(ctx, input)
}

func (a *Activities) GetDeliverabilityTestReport(ctx context.Context, input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDeliverabilityTestReportWithContext(ctx, input)
}

func (a *Activities) GetDomainDeliverabilityCampaign(ctx context.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainDeliverabilityCampaignWithContext(ctx, input)
}

func (a *Activities) GetDomainStatisticsReport(ctx context.Context, input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDomainStatisticsReportWithContext(ctx, input)
}

func (a *Activities) GetEmailIdentity(ctx context.Context, input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetEmailIdentityWithContext(ctx, input)
}

func (a *Activities) ListConfigurationSets(ctx context.Context, input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationSetsWithContext(ctx, input)
}

func (a *Activities) ListDedicatedIpPools(ctx context.Context, input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDedicatedIpPoolsWithContext(ctx, input)
}

func (a *Activities) ListDeliverabilityTestReports(ctx context.Context, input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDeliverabilityTestReportsWithContext(ctx, input)
}

func (a *Activities) ListDomainDeliverabilityCampaigns(ctx context.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDomainDeliverabilityCampaignsWithContext(ctx, input)
}

func (a *Activities) ListEmailIdentities(ctx context.Context, input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEmailIdentitiesWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) PutAccountDedicatedIpWarmupAttributes(ctx context.Context, input *pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) (*pinpointemail.PutAccountDedicatedIpWarmupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAccountDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *Activities) PutAccountSendingAttributes(ctx context.Context, input *pinpointemail.PutAccountSendingAttributesInput) (*pinpointemail.PutAccountSendingAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutAccountSendingAttributesWithContext(ctx, input)
}

func (a *Activities) PutConfigurationSetDeliveryOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetDeliveryOptionsInput) (*pinpointemail.PutConfigurationSetDeliveryOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetDeliveryOptionsWithContext(ctx, input)
}

func (a *Activities) PutConfigurationSetReputationOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetReputationOptionsInput) (*pinpointemail.PutConfigurationSetReputationOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetReputationOptionsWithContext(ctx, input)
}

func (a *Activities) PutConfigurationSetSendingOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetSendingOptionsInput) (*pinpointemail.PutConfigurationSetSendingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetSendingOptionsWithContext(ctx, input)
}

func (a *Activities) PutConfigurationSetTrackingOptions(ctx context.Context, input *pinpointemail.PutConfigurationSetTrackingOptionsInput) (*pinpointemail.PutConfigurationSetTrackingOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutConfigurationSetTrackingOptionsWithContext(ctx, input)
}

func (a *Activities) PutDedicatedIpInPool(ctx context.Context, input *pinpointemail.PutDedicatedIpInPoolInput) (*pinpointemail.PutDedicatedIpInPoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDedicatedIpInPoolWithContext(ctx, input)
}

func (a *Activities) PutDedicatedIpWarmupAttributes(ctx context.Context, input *pinpointemail.PutDedicatedIpWarmupAttributesInput) (*pinpointemail.PutDedicatedIpWarmupAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDedicatedIpWarmupAttributesWithContext(ctx, input)
}

func (a *Activities) PutDeliverabilityDashboardOption(ctx context.Context, input *pinpointemail.PutDeliverabilityDashboardOptionInput) (*pinpointemail.PutDeliverabilityDashboardOptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutDeliverabilityDashboardOptionWithContext(ctx, input)
}

func (a *Activities) PutEmailIdentityDkimAttributes(ctx context.Context, input *pinpointemail.PutEmailIdentityDkimAttributesInput) (*pinpointemail.PutEmailIdentityDkimAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityDkimAttributesWithContext(ctx, input)
}

func (a *Activities) PutEmailIdentityFeedbackAttributes(ctx context.Context, input *pinpointemail.PutEmailIdentityFeedbackAttributesInput) (*pinpointemail.PutEmailIdentityFeedbackAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityFeedbackAttributesWithContext(ctx, input)
}

func (a *Activities) PutEmailIdentityMailFromAttributes(ctx context.Context, input *pinpointemail.PutEmailIdentityMailFromAttributesInput) (*pinpointemail.PutEmailIdentityMailFromAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutEmailIdentityMailFromAttributesWithContext(ctx, input)
}

func (a *Activities) SendEmail(ctx context.Context, input *pinpointemail.SendEmailInput) (*pinpointemail.SendEmailOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SendEmailWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *pinpointemail.TagResourceInput) (*pinpointemail.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *pinpointemail.UntagResourceInput) (*pinpointemail.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateConfigurationSetEventDestination(ctx context.Context, input *pinpointemail.UpdateConfigurationSetEventDestinationInput) (*pinpointemail.UpdateConfigurationSetEventDestinationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateConfigurationSetEventDestinationWithContext(ctx, input)
}