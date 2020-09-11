
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

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
