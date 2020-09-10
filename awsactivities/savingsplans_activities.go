package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"github.com/aws/aws-sdk-go/service/savingsplans/savingsplansiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type SavingsPlansActivities struct {
	client savingsplansiface.SavingsPlansAPI
}

func NewSavingsPlansActivities(session *session.Session, config ...*aws.Config) *SavingsPlansActivities {
	client := savingsplans.New(session, config...)
	return &SavingsPlansActivities{client: client}
}

func (a *SavingsPlansActivities) CreateSavingsPlan(ctx context.Context, input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.CreateSavingsPlanWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlanRates(ctx context.Context, input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	return a.client.DescribeSavingsPlanRatesWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlans(ctx context.Context, input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error) {
	return a.client.DescribeSavingsPlansWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlansOfferingRates(ctx context.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	return a.client.DescribeSavingsPlansOfferingRatesWithContext(ctx, input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlansOfferings(ctx context.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	return a.client.DescribeSavingsPlansOfferingsWithContext(ctx, input)
}

func (a *SavingsPlansActivities) ListTagsForResource(ctx context.Context, input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *SavingsPlansActivities) TagResource(ctx context.Context, input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *SavingsPlansActivities) UntagResource(ctx context.Context, input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
