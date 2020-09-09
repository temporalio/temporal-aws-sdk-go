
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice/transcribestreamingserviceiface"
)

type TranscribeStreamingServiceActivities struct {
    client transcribestreamingserviceiface.TranscribeStreamingServiceAPI
}

func NewTranscribeStreamingServiceActivities(session *session.Session, config ...*aws.Config) *TranscribeStreamingServiceActivities {
    client := transcribestreamingservice.New(session, config...)
    return &TranscribeStreamingServiceActivities{client: client}
}

func (a *TranscribeStreamingServiceActivities) StartStreamTranscription(input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
    return a.client.StartStreamTranscription(input)
}
