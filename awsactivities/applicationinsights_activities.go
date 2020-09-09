package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"github.com/aws/aws-sdk-go/service/applicationinsights/applicationinsightsiface"
)

type ApplicationInsightsActivities struct {
	client applicationinsightsiface.ApplicationInsightsAPI
}

func NewApplicationInsightsActivities(session *session.Session, config ...*aws.Config) *ApplicationInsightsActivities {
	client := applicationinsights.New(session, config...)
	return &ApplicationInsightsActivities{client: client}
}

func (a *ApplicationInsightsActivities) CreateApplication(input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error) {
	return a.client.CreateApplication(input)
}

func (a *ApplicationInsightsActivities) CreateComponent(input *applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error) {
	return a.client.CreateComponent(input)
}

func (a *ApplicationInsightsActivities) CreateLogPattern(input *applicationinsights.CreateLogPatternInput) (*applicationinsights.CreateLogPatternOutput, error) {
	return a.client.CreateLogPattern(input)
}

func (a *ApplicationInsightsActivities) DeleteApplication(input *applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error) {
	return a.client.DeleteApplication(input)
}

func (a *ApplicationInsightsActivities) DeleteComponent(input *applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error) {
	return a.client.DeleteComponent(input)
}

func (a *ApplicationInsightsActivities) DeleteLogPattern(input *applicationinsights.DeleteLogPatternInput) (*applicationinsights.DeleteLogPatternOutput, error) {
	return a.client.DeleteLogPattern(input)
}

func (a *ApplicationInsightsActivities) DescribeApplication(input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error) {
	return a.client.DescribeApplication(input)
}

func (a *ApplicationInsightsActivities) DescribeComponent(input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error) {
	return a.client.DescribeComponent(input)
}

func (a *ApplicationInsightsActivities) DescribeComponentConfiguration(input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	return a.client.DescribeComponentConfiguration(input)
}

func (a *ApplicationInsightsActivities) DescribeComponentConfigurationRecommendation(input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	return a.client.DescribeComponentConfigurationRecommendation(input)
}

func (a *ApplicationInsightsActivities) DescribeLogPattern(input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error) {
	return a.client.DescribeLogPattern(input)
}

func (a *ApplicationInsightsActivities) DescribeObservation(input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error) {
	return a.client.DescribeObservation(input)
}

func (a *ApplicationInsightsActivities) DescribeProblem(input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error) {
	return a.client.DescribeProblem(input)
}

func (a *ApplicationInsightsActivities) DescribeProblemObservations(input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	return a.client.DescribeProblemObservations(input)
}

func (a *ApplicationInsightsActivities) ListApplications(input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error) {
	return a.client.ListApplications(input)
}

func (a *ApplicationInsightsActivities) ListComponents(input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error) {
	return a.client.ListComponents(input)
}

func (a *ApplicationInsightsActivities) ListConfigurationHistory(input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	return a.client.ListConfigurationHistory(input)
}

func (a *ApplicationInsightsActivities) ListLogPatternSets(input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error) {
	return a.client.ListLogPatternSets(input)
}

func (a *ApplicationInsightsActivities) ListLogPatterns(input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error) {
	return a.client.ListLogPatterns(input)
}

func (a *ApplicationInsightsActivities) ListProblems(input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error) {
	return a.client.ListProblems(input)
}

func (a *ApplicationInsightsActivities) ListTagsForResource(input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResource(input)
}

func (a *ApplicationInsightsActivities) TagResource(input *applicationinsights.TagResourceInput) (*applicationinsights.TagResourceOutput, error) {
	return a.client.TagResource(input)
}

func (a *ApplicationInsightsActivities) UntagResource(input *applicationinsights.UntagResourceInput) (*applicationinsights.UntagResourceOutput, error) {
	return a.client.UntagResource(input)
}

func (a *ApplicationInsightsActivities) UpdateApplication(input *applicationinsights.UpdateApplicationInput) (*applicationinsights.UpdateApplicationOutput, error) {
	return a.client.UpdateApplication(input)
}

func (a *ApplicationInsightsActivities) UpdateComponent(input *applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error) {
	return a.client.UpdateComponent(input)
}

func (a *ApplicationInsightsActivities) UpdateComponentConfiguration(input *applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
	return a.client.UpdateComponentConfiguration(input)
}

func (a *ApplicationInsightsActivities) UpdateLogPattern(input *applicationinsights.UpdateLogPatternInput) (*applicationinsights.UpdateLogPatternOutput, error) {
	return a.client.UpdateLogPattern(input)
}
