
package awsactivities

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
)

type WorkMailMessageFlowActivities struct {
	client workmailmessageflowiface.WorkMailMessageFlowAPI
}

func NewWorkMailMessageFlowActivities(session *session.Session, config... *aws.Config) *WorkMailMessageFlowActivities {
    client := workmailmessageflow.New(session, config...)
    return &WorkMailMessageFlowActivities{client: client}
}

func (a *WorkMailMessageFlowActivities) GetRawMessageContent(input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
    return a.client.GetRawMessageContent(input)
}
