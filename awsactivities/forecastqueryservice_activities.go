
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/forecastqueryservice"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice/forecastqueryserviceiface"
)

type ForecastQueryServiceActivities struct {
	client forecastqueryserviceiface.ForecastQueryServiceAPI
}

func NewForecastQueryServiceActivities(client forecastqueryserviceiface.ForecastQueryServiceAPI) *ForecastQueryServiceActivities {
    return &ForecastQueryServiceActivities{client: client}
}

func (a *ForecastQueryServiceActivities) QueryForecast(input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
    return a.client.QueryForecast(input)
}
