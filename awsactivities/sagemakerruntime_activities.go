
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime/sagemakerruntimeiface"
)

type SageMakerRuntimeActivities struct {
	client sagemakerruntimeiface.SageMakerRuntimeAPI
}

func NewSageMakerRuntimeActivities(client sagemakerruntimeiface.SageMakerRuntimeAPI) *SageMakerRuntimeActivities {
    return &SageMakerRuntimeActivities{client: client}
}

func (a *SageMakerRuntimeActivities) InvokeEndpoint(input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error) {
    return a.client.InvokeEndpoint(input)
}
