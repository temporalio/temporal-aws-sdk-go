package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice/transcribestreamingserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type TranscribeStreamingServiceActivities struct {
	client transcribestreamingserviceiface.TranscribeStreamingServiceAPI
}

func NewTranscribeStreamingServiceActivities(session *session.Session, config ...*aws.Config) *TranscribeStreamingServiceActivities {
	client := transcribestreamingservice.New(session, config...)
	return &TranscribeStreamingServiceActivities{client: client}
}

func (a *TranscribeStreamingServiceActivities) StartStreamTranscription(ctx context.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
	return a.client.StartStreamTranscriptionWithContext(ctx, input)
}
