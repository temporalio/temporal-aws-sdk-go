package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
)

type WorkMailMessageFlowActivities struct {
	client workmailmessageflowiface.WorkMailMessageFlowAPI
}

func NewWorkMailMessageFlowActivities(client workmailmessageflowiface.WorkMailMessageFlowAPI) *WorkMailMessageFlowActivities {
    return &WorkMailMessageFlowActivities{client: client}
}

func (a *WorkMailMessageFlowActivities) GetRawMessageContent(input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
    return a.client.GetRawMessageContent(input)
}
