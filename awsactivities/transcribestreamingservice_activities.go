
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice/transcribestreamingserviceiface"
)

type TranscribeStreamingServiceActivities struct {
	client transcribestreamingserviceiface.TranscribeStreamingServiceAPI
}

func NewTranscribeStreamingServiceActivities(client transcribestreamingserviceiface.TranscribeStreamingServiceAPI) *TranscribeStreamingServiceActivities {
    return &TranscribeStreamingServiceActivities{client: client}
}

func (a *TranscribeStreamingServiceActivities) StartStreamTranscription(input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
    return a.client.StartStreamTranscription(input)
}
