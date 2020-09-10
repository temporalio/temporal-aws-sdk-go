package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pi"
	"github.com/aws/aws-sdk-go/service/pi/piiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type PIActivities struct {
	client piiface.PIAPI
}

func NewPIActivities(session *session.Session, config ...*aws.Config) *PIActivities {
	client := pi.New(session, config...)
	return &PIActivities{client: client}
}

func (a *PIActivities) DescribeDimensionKeys(ctx context.Context, input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
	return a.client.DescribeDimensionKeysWithContext(ctx, input)
}

func (a *PIActivities) GetResourceMetrics(ctx context.Context, input *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error) {
	return a.client.GetResourceMetricsWithContext(ctx, input)
}
