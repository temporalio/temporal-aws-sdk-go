package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
	"go.temporal.io/sdk/activity"
)

// ensure that activity import is valid even if not used by the generated code
type _ = activity.Info

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
