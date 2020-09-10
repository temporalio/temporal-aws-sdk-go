package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/connectparticipant/connectparticipantiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

type ConnectParticipantActivities struct {
	client connectparticipantiface.ConnectParticipantAPI
}

func NewConnectParticipantActivities(session *session.Session, config ...*aws.Config) *ConnectParticipantActivities {
	client := connectparticipant.New(session, config...)
	return &ConnectParticipantActivities{client: client}
}

func (a *ConnectParticipantActivities) CreateParticipantConnection(ctx context.Context, input *connectparticipant.CreateParticipantConnectionInput) (*connectparticipant.CreateParticipantConnectionOutput, error) {
	return a.client.CreateParticipantConnectionWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) DisconnectParticipant(ctx context.Context, input *connectparticipant.DisconnectParticipantInput) (*connectparticipant.DisconnectParticipantOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.DisconnectParticipantWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) GetTranscript(ctx context.Context, input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error) {
	return a.client.GetTranscriptWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) SendEvent(ctx context.Context, input *connectparticipant.SendEventInput) (*connectparticipant.SendEventOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.SendEventWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) SendMessage(ctx context.Context, input *connectparticipant.SendMessageInput) (*connectparticipant.SendMessageOutput, error) {
	// Use the same token during retries
	if input.ClientToken == nil {
		info := activity.GetInfo(ctx)
		token := info.WorkflowExecution.RunID + "-" + info.ActivityID
		input.ClientToken = &token
	}
	return a.client.SendMessageWithContext(ctx, input)
}
