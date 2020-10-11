// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package applicationinsights

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"github.com/aws/aws-sdk-go/service/applicationinsights/applicationinsightsiface"
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
	client applicationinsightsiface.ApplicationInsightsAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := applicationinsights.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (applicationinsightsiface.ApplicationInsightsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return applicationinsights.New(sess), nil
}

func (a *Activities) CreateApplication(ctx context.Context, input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApplicationWithContext(ctx, input)
}

func (a *Activities) CreateComponent(ctx context.Context, input *applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateComponentWithContext(ctx, input)
}

func (a *Activities) CreateLogPattern(ctx context.Context, input *applicationinsights.CreateLogPatternInput) (*applicationinsights.CreateLogPatternOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLogPatternWithContext(ctx, input)
}

func (a *Activities) DeleteApplication(ctx context.Context, input *applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApplicationWithContext(ctx, input)
}

func (a *Activities) DeleteComponent(ctx context.Context, input *applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteComponentWithContext(ctx, input)
}

func (a *Activities) DeleteLogPattern(ctx context.Context, input *applicationinsights.DeleteLogPatternInput) (*applicationinsights.DeleteLogPatternOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLogPatternWithContext(ctx, input)
}

func (a *Activities) DescribeApplication(ctx context.Context, input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeApplicationWithContext(ctx, input)
}

func (a *Activities) DescribeComponent(ctx context.Context, input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeComponentWithContext(ctx, input)
}

func (a *Activities) DescribeComponentConfiguration(ctx context.Context, input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeComponentConfigurationWithContext(ctx, input)
}

func (a *Activities) DescribeComponentConfigurationRecommendation(ctx context.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeComponentConfigurationRecommendationWithContext(ctx, input)
}

func (a *Activities) DescribeLogPattern(ctx context.Context, input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLogPatternWithContext(ctx, input)
}

func (a *Activities) DescribeObservation(ctx context.Context, input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeObservationWithContext(ctx, input)
}

func (a *Activities) DescribeProblem(ctx context.Context, input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProblemWithContext(ctx, input)
}

func (a *Activities) DescribeProblemObservations(ctx context.Context, input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProblemObservationsWithContext(ctx, input)
}

func (a *Activities) ListApplications(ctx context.Context, input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApplicationsWithContext(ctx, input)
}

func (a *Activities) ListComponents(ctx context.Context, input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListComponentsWithContext(ctx, input)
}

func (a *Activities) ListConfigurationHistory(ctx context.Context, input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListConfigurationHistoryWithContext(ctx, input)
}

func (a *Activities) ListLogPatternSets(ctx context.Context, input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLogPatternSetsWithContext(ctx, input)
}

func (a *Activities) ListLogPatterns(ctx context.Context, input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLogPatternsWithContext(ctx, input)
}

func (a *Activities) ListProblems(ctx context.Context, input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProblemsWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) TagResource(ctx context.Context, input *applicationinsights.TagResourceInput) (*applicationinsights.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *Activities) UntagResource(ctx context.Context, input *applicationinsights.UntagResourceInput) (*applicationinsights.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *Activities) UpdateApplication(ctx context.Context, input *applicationinsights.UpdateApplicationInput) (*applicationinsights.UpdateApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApplicationWithContext(ctx, input)
}

func (a *Activities) UpdateComponent(ctx context.Context, input *applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateComponentWithContext(ctx, input)
}

func (a *Activities) UpdateComponentConfiguration(ctx context.Context, input *applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateComponentConfigurationWithContext(ctx, input)
}

func (a *Activities) UpdateLogPattern(ctx context.Context, input *applicationinsights.UpdateLogPatternInput) (*applicationinsights.UpdateLogPatternOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateLogPatternWithContext(ctx, input)
}
