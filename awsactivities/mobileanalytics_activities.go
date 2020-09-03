package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
)

type MobileAnalyticsActivities struct {
	client mobileanalyticsiface.MobileAnalyticsAPI
}

func NewMobileAnalyticsActivities(client mobileanalyticsiface.MobileAnalyticsAPI) *MobileAnalyticsActivities {
    return &MobileAnalyticsActivities{client: client}
}


func (a *MobileAnalyticsActivities) PutEvents(input *mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error) {
    return a.client.PutEvents(input)
}

