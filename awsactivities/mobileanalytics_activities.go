package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type MobileAnalyticsActivities struct {
	client mobileanalyticsiface.MobileAnalyticsAPI
}

func NewMobileAnalyticsActivities(session *session.Session, config ...*aws.Config) *MobileAnalyticsActivities {
	client := mobileanalytics.New(session, config...)
	return &MobileAnalyticsActivities{client: client}
}

func (a *MobileAnalyticsActivities) PutEvents(ctx context.Context, input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error) {
	return a.client.PutEventsWithContext(ctx, input)
}
