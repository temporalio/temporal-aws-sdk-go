
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type WorkMailMessageFlowActivities struct {
    client workmailmessageflowiface.WorkMailMessageFlowAPI
}

func NewWorkMailMessageFlowActivities(session *session.Session, config ...*aws.Config) *WorkMailMessageFlowActivities {
    client := workmailmessageflow.New(session, config...)
    return &WorkMailMessageFlowActivities{client: client}
}

func (a *WorkMailMessageFlowActivities) GetRawMessageContent(ctx context.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
    return a.client.GetRawMessageContentWithContext(ctx, input)
}
