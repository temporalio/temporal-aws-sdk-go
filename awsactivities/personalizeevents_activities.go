
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"github.com/aws/aws-sdk-go/service/personalizeevents/personalizeeventsiface"
)

type PersonalizeEventsActivities struct {
	client personalizeeventsiface.PersonalizeEventsAPI
}

func NewPersonalizeEventsActivities(client personalizeeventsiface.PersonalizeEventsAPI) *PersonalizeEventsActivities {
    return &PersonalizeEventsActivities{client: client}
}

func (a *PersonalizeEventsActivities) PutEvents(input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
    return a.client.PutEvents(input)
}
