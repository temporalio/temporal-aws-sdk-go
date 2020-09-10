package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"github.com/aws/aws-sdk-go/service/personalizeevents/personalizeeventsiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type PersonalizeEventsActivities struct {
	client personalizeeventsiface.PersonalizeEventsAPI
}

func NewPersonalizeEventsActivities(session *session.Session, config ...*aws.Config) *PersonalizeEventsActivities {
	client := personalizeevents.New(session, config...)
	return &PersonalizeEventsActivities{client: client}
}

func (a *PersonalizeEventsActivities) PutEvents(ctx context.Context, input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
	return a.client.PutEventsWithContext(ctx, input)
}
