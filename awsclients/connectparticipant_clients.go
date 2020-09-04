package awsclients

import (
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ConnectParticipantClient interface {
    CreateParticipantConnection(ctx workflow.Context, input *connectparticipant.CreateParticipantConnectionInput) (*connectparticipant.CreateParticipantConnectionOutput, error)
    CreateParticipantConnectionAsync(ctx workflow.Context, input *connectparticipant.CreateParticipantConnectionInput) *ConnectparticipantCreateParticipantConnectionResult

    DisconnectParticipant(ctx workflow.Context, input *connectparticipant.DisconnectParticipantInput) (*connectparticipant.DisconnectParticipantOutput, error)
    DisconnectParticipantAsync(ctx workflow.Context, input *connectparticipant.DisconnectParticipantInput) *ConnectparticipantDisconnectParticipantResult

    GetTranscript(ctx workflow.Context, input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error)
    GetTranscriptAsync(ctx workflow.Context, input *connectparticipant.GetTranscriptInput) *ConnectparticipantGetTranscriptResult

    SendEvent(ctx workflow.Context, input *connectparticipant.SendEventInput) (*connectparticipant.SendEventOutput, error)
    SendEventAsync(ctx workflow.Context, input *connectparticipant.SendEventInput) *ConnectparticipantSendEventResult

    SendMessage(ctx workflow.Context, input *connectparticipant.SendMessageInput) (*connectparticipant.SendMessageOutput, error)
    SendMessageAsync(ctx workflow.Context, input *connectparticipant.SendMessageInput) *ConnectparticipantSendMessageResult
}

type ConnectparticipantCreateParticipantConnectionResult struct {
	Result workflow.Future
}

func (r *ConnectparticipantCreateParticipantConnectionResult) Get(ctx workflow.Context) (*connectparticipant.CreateParticipantConnectionOutput, error) {
    var output connectparticipant.CreateParticipantConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectparticipantDisconnectParticipantResult struct {
	Result workflow.Future
}

func (r *ConnectparticipantDisconnectParticipantResult) Get(ctx workflow.Context) (*connectparticipant.DisconnectParticipantOutput, error) {
    var output connectparticipant.DisconnectParticipantOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectparticipantGetTranscriptResult struct {
	Result workflow.Future
}

func (r *ConnectparticipantGetTranscriptResult) Get(ctx workflow.Context) (*connectparticipant.GetTranscriptOutput, error) {
    var output connectparticipant.GetTranscriptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectparticipantSendEventResult struct {
	Result workflow.Future
}

func (r *ConnectparticipantSendEventResult) Get(ctx workflow.Context) (*connectparticipant.SendEventOutput, error) {
    var output connectparticipant.SendEventOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectparticipantSendMessageResult struct {
	Result workflow.Future
}

func (r *ConnectparticipantSendMessageResult) Get(ctx workflow.Context) (*connectparticipant.SendMessageOutput, error) {
    var output connectparticipant.SendMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectParticipantStub struct {
    activities awsactivities.ConnectParticipantActivities
}

func NewConnectParticipantStub() ConnectParticipantClient {
    return &ConnectParticipantStub{}
}

func (a *ConnectParticipantStub) CreateParticipantConnection(ctx workflow.Context, input *connectparticipant.CreateParticipantConnectionInput) (*connectparticipant.CreateParticipantConnectionOutput, error) {
    var output connectparticipant.CreateParticipantConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateParticipantConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectParticipantStub) CreateParticipantConnectionAsync(ctx workflow.Context, input *connectparticipant.CreateParticipantConnectionInput) *ConnectparticipantCreateParticipantConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateParticipantConnection, input)
    return &ConnectparticipantCreateParticipantConnectionResult{Result: future}
}

func (a *ConnectParticipantStub) DisconnectParticipant(ctx workflow.Context, input *connectparticipant.DisconnectParticipantInput) (*connectparticipant.DisconnectParticipantOutput, error) {
    var output connectparticipant.DisconnectParticipantOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisconnectParticipant, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectParticipantStub) DisconnectParticipantAsync(ctx workflow.Context, input *connectparticipant.DisconnectParticipantInput) *ConnectparticipantDisconnectParticipantResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisconnectParticipant, input)
    return &ConnectparticipantDisconnectParticipantResult{Result: future}
}

func (a *ConnectParticipantStub) GetTranscript(ctx workflow.Context, input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error) {
    var output connectparticipant.GetTranscriptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTranscript, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectParticipantStub) GetTranscriptAsync(ctx workflow.Context, input *connectparticipant.GetTranscriptInput) *ConnectparticipantGetTranscriptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTranscript, input)
    return &ConnectparticipantGetTranscriptResult{Result: future}
}

func (a *ConnectParticipantStub) SendEvent(ctx workflow.Context, input *connectparticipant.SendEventInput) (*connectparticipant.SendEventOutput, error) {
    var output connectparticipant.SendEventOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendEvent, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectParticipantStub) SendEventAsync(ctx workflow.Context, input *connectparticipant.SendEventInput) *ConnectparticipantSendEventResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendEvent, input)
    return &ConnectparticipantSendEventResult{Result: future}
}

func (a *ConnectParticipantStub) SendMessage(ctx workflow.Context, input *connectparticipant.SendMessageInput) (*connectparticipant.SendMessageOutput, error) {
    var output connectparticipant.SendMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectParticipantStub) SendMessageAsync(ctx workflow.Context, input *connectparticipant.SendMessageInput) *ConnectparticipantSendMessageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendMessage, input)
    return &ConnectparticipantSendMessageResult{Result: future}
}
