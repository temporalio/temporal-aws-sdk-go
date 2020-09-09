
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/polly/pollyiface"
)

type PollyActivities struct {
    client pollyiface.PollyAPI
}

func NewPollyActivities(session *session.Session, config ...*aws.Config) *PollyActivities {
    client := polly.New(session, config...)
    return &PollyActivities{client: client}
}

func (a *PollyActivities) DeleteLexicon(input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error) {
    return a.client.DeleteLexicon(input)
}

func (a *PollyActivities) DescribeVoices(input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error) {
    return a.client.DescribeVoices(input)
}

func (a *PollyActivities) GetLexicon(input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error) {
    return a.client.GetLexicon(input)
}

func (a *PollyActivities) GetSpeechSynthesisTask(input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error) {
    return a.client.GetSpeechSynthesisTask(input)
}

func (a *PollyActivities) ListLexicons(input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error) {
    return a.client.ListLexicons(input)
}

func (a *PollyActivities) ListSpeechSynthesisTasks(input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error) {
    return a.client.ListSpeechSynthesisTasks(input)
}

func (a *PollyActivities) PutLexicon(input *polly.PutLexiconInput) (*polly.PutLexiconOutput, error) {
    return a.client.PutLexicon(input)
}

func (a *PollyActivities) StartSpeechSynthesisTask(input *polly.StartSpeechSynthesisTaskInput) (*polly.StartSpeechSynthesisTaskOutput, error) {
    return a.client.StartSpeechSynthesisTask(input)
}

func (a *PollyActivities) SynthesizeSpeech(input *polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error) {
    return a.client.SynthesizeSpeech(input)
}
