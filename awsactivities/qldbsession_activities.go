package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/aws/aws-sdk-go/service/qldbsession/qldbsessioniface"
)

type QLDBSessionActivities struct {
	client qldbsessioniface.QLDBSessionAPI
}

func NewQLDBSessionActivities(session *session.Session, config ...*aws.Config) *QLDBSessionActivities {
	client := qldbsession.New(session, config...)
	return &QLDBSessionActivities{client: client}
}

func (a *QLDBSessionActivities) SendCommand(input *qldbsession.SendCommandInput) (*qldbsession.SendCommandOutput, error) {
	return a.client.SendCommand(input)
}
