// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice/forecastqueryserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ForecastQueryServiceActivities struct {
	client forecastqueryserviceiface.ForecastQueryServiceAPI

	sessionFactory SessionFactory
}

func NewForecastQueryServiceActivities(sess *session.Session, config ...*aws.Config) *ForecastQueryServiceActivities {
	client := forecastqueryservice.New(sess, config...)
	return &ForecastQueryServiceActivities{client: client}
}

func NewForecastQueryServiceActivitiesWithSessionFactory(sessionFactory SessionFactory) *ForecastQueryServiceActivities {
	return &ForecastQueryServiceActivities{sessionFactory: sessionFactory}
}

func (a *ForecastQueryServiceActivities) getClient(ctx context.Context) (forecastqueryserviceiface.ForecastQueryServiceAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return forecastqueryservice.New(sess), nil
}

func (a *ForecastQueryServiceActivities) QueryForecast(ctx context.Context, input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.QueryForecastWithContext(ctx, input)
}
