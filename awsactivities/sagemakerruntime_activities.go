package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime/sagemakerruntimeiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type SageMakerRuntimeActivities struct {
	client sagemakerruntimeiface.SageMakerRuntimeAPI
}

func NewSageMakerRuntimeActivities(session *session.Session, config ...*aws.Config) *SageMakerRuntimeActivities {
	client := sagemakerruntime.New(session, config...)
	return &SageMakerRuntimeActivities{client: client}
}

func (a *SageMakerRuntimeActivities) InvokeEndpoint(ctx context.Context, input *sagemakerruntime.InvokeEndpointInput) (*sagemakerruntime.InvokeEndpointOutput, error) {
	return a.client.InvokeEndpointWithContext(ctx, input)
}
