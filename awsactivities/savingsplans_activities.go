
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"github.com/aws/aws-sdk-go/service/savingsplans/savingsplansiface"
)

type SavingsPlansActivities struct {
	client savingsplansiface.SavingsPlansAPI
}

func NewSavingsPlansActivities(client savingsplansiface.SavingsPlansAPI) *SavingsPlansActivities {
    return &SavingsPlansActivities{client: client}
}

func (a *SavingsPlansActivities) CreateSavingsPlan(input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error) {
    return a.client.CreateSavingsPlan(input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlanRates(input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
    return a.client.DescribeSavingsPlanRates(input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlans(input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error) {
    return a.client.DescribeSavingsPlans(input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlansOfferingRates(input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
    return a.client.DescribeSavingsPlansOfferingRates(input)
}

func (a *SavingsPlansActivities) DescribeSavingsPlansOfferings(input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
    return a.client.DescribeSavingsPlansOfferings(input)
}

func (a *SavingsPlansActivities) ListTagsForResource(input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResource(input)
}

func (a *SavingsPlansActivities) TagResource(input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error) {
    return a.client.TagResource(input)
}

func (a *SavingsPlansActivities) UntagResource(input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error) {
    return a.client.UntagResource(input)
}
