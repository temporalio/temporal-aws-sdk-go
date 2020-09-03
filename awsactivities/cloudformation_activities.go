package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

type CloudFormationActivities struct {
	client cloudformationiface.CloudFormationAPI
}

func NewCloudFormationActivities(client cloudformationiface.CloudFormationAPI) *CloudFormationActivities {
    return &CloudFormationActivities{client: client}
}

func (a *CloudFormationActivities) CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
    return a.client.CancelUpdateStack(input)
}

func (a *CloudFormationActivities) ContinueUpdateRollback(input *cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error) {
    return a.client.ContinueUpdateRollback(input)
}

func (a *CloudFormationActivities) CreateChangeSet(input *cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error) {
    return a.client.CreateChangeSet(input)
}
