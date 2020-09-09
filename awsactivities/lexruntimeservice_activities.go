
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface"
)

type LexRuntimeServiceActivities struct {
    client lexruntimeserviceiface.LexRuntimeServiceAPI
}

func NewLexRuntimeServiceActivities(session *session.Session, config ...*aws.Config) *LexRuntimeServiceActivities {
    client := lexruntimeservice.New(session, config...)
    return &LexRuntimeServiceActivities{client: client}
}

func (a *LexRuntimeServiceActivities) DeleteSession(input *lexruntimeservice.DeleteSessionInput) (*lexruntimeservice.DeleteSessionOutput, error) {
    return a.client.DeleteSession(input)
}

func (a *LexRuntimeServiceActivities) GetSession(input *lexruntimeservice.GetSessionInput) (*lexruntimeservice.GetSessionOutput, error) {
    return a.client.GetSession(input)
}

func (a *LexRuntimeServiceActivities) PostContent(input *lexruntimeservice.PostContentInput) (*lexruntimeservice.PostContentOutput, error) {
    return a.client.PostContent(input)
}

func (a *LexRuntimeServiceActivities) PostText(input *lexruntimeservice.PostTextInput) (*lexruntimeservice.PostTextOutput, error) {
    return a.client.PostText(input)
}

func (a *LexRuntimeServiceActivities) PutSession(input *lexruntimeservice.PutSessionInput) (*lexruntimeservice.PutSessionOutput, error) {
    return a.client.PutSession(input)
}
