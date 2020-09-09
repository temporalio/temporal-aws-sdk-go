
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
)

type MobileAnalyticsActivities struct {
	client mobileanalyticsiface.MobileAnalyticsAPI
}

func NewMobileAnalyticsActivities(session *session.Session, config... *aws.Config) *MobileAnalyticsActivities {
    client := mobileanalytics.New(session, config...)
    return &MobileAnalyticsActivities{client: client}
}

func (a *MobileAnalyticsActivities) PutEvents(input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error) {
    return a.client.PutEvents(input)
}
