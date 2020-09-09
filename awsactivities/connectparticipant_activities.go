
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/connectparticipant/connectparticipantiface"
)

type ConnectParticipantActivities struct {
    client connectparticipantiface.ConnectParticipantAPI
}

func NewConnectParticipantActivities(session *session.Session, config ...*aws.Config) *ConnectParticipantActivities {
    client := connectparticipant.New(session, config...)
    return &ConnectParticipantActivities{client: client}
}

func (a *ConnectParticipantActivities) CreateParticipantConnection(input *connectparticipant.CreateParticipantConnectionInput) (*connectparticipant.CreateParticipantConnectionOutput, error) {
    return a.client.CreateParticipantConnection(input)
}

func (a *ConnectParticipantActivities) DisconnectParticipant(input *connectparticipant.DisconnectParticipantInput) (*connectparticipant.DisconnectParticipantOutput, error) {
    return a.client.DisconnectParticipant(input)
}

func (a *ConnectParticipantActivities) GetTranscript(input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error) {
    return a.client.GetTranscript(input)
}

func (a *ConnectParticipantActivities) SendEvent(input *connectparticipant.SendEventInput) (*connectparticipant.SendEventOutput, error) {
    return a.client.SendEvent(input)
}

func (a *ConnectParticipantActivities) SendMessage(input *connectparticipant.SendMessageInput) (*connectparticipant.SendMessageOutput, error) {
    return a.client.SendMessage(input)
}
