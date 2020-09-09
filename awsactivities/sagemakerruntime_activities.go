
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime/sagemakerruntimeiface"
)

type SageMakerRuntimeActivities struct {
    client sagemakerruntimeiface.SageMakerRuntimeAPI
}

func NewSageMakerRuntimeActivities(session *session.Session, config ...*aws.Config) *SageMakerRuntimeActivities {
    client := sagemakerruntime.New(session, config...)
    return &SageMakerRuntimeActivities{client: client}
}

func (a *SageMakerRuntimeActivities) InvokeEndpoint(input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error) {
    return a.client.InvokeEndpoint(input)
}
