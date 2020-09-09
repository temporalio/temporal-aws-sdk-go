package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice/forecastqueryserviceiface"
)

type ForecastQueryServiceActivities struct {
	client forecastqueryserviceiface.ForecastQueryServiceAPI
}

func NewForecastQueryServiceActivities(session *session.Session, config ...*aws.Config) *ForecastQueryServiceActivities {
	client := forecastqueryservice.New(session, config...)
	return &ForecastQueryServiceActivities{client: client}
}

func (a *ForecastQueryServiceActivities) QueryForecast(input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
	return a.client.QueryForecast(input)
}
