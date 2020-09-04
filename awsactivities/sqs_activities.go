
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type SQSActivities struct {
	client sqsiface.SQSAPI
}

func NewSQSActivities(client sqsiface.SQSAPI) *SQSActivities {
    return &SQSActivities{client: client}
}

func (a *SQSActivities) AddPermission(input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
    return a.client.AddPermission(input)
}

func (a *SQSActivities) ChangeMessageVisibility(input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
    return a.client.ChangeMessageVisibility(input)
}

func (a *SQSActivities) ChangeMessageVisibilityBatch(input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
    return a.client.ChangeMessageVisibilityBatch(input)
}

func (a *SQSActivities) CreateQueue(input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
    return a.client.CreateQueue(input)
}

func (a *SQSActivities) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
    return a.client.DeleteMessage(input)
}

func (a *SQSActivities) DeleteMessageBatch(input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
    return a.client.DeleteMessageBatch(input)
}

func (a *SQSActivities) DeleteQueue(input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
    return a.client.DeleteQueue(input)
}

func (a *SQSActivities) GetQueueAttributes(input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
    return a.client.GetQueueAttributes(input)
}

func (a *SQSActivities) GetQueueUrl(input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
    return a.client.GetQueueUrl(input)
}

func (a *SQSActivities) ListDeadLetterSourceQueues(input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
    return a.client.ListDeadLetterSourceQueues(input)
}

func (a *SQSActivities) ListQueueTags(input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
    return a.client.ListQueueTags(input)
}

func (a *SQSActivities) ListQueues(input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
    return a.client.ListQueues(input)
}

func (a *SQSActivities) PurgeQueue(input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
    return a.client.PurgeQueue(input)
}

func (a *SQSActivities) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
    return a.client.ReceiveMessage(input)
}

func (a *SQSActivities) RemovePermission(input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
    return a.client.RemovePermission(input)
}

func (a *SQSActivities) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
    return a.client.SendMessage(input)
}

func (a *SQSActivities) SendMessageBatch(input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
    return a.client.SendMessageBatch(input)
}

func (a *SQSActivities) SetQueueAttributes(input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
    return a.client.SetQueueAttributes(input)
}

func (a *SQSActivities) TagQueue(input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
    return a.client.TagQueue(input)
}

func (a *SQSActivities) UntagQueue(input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
    return a.client.UntagQueue(input)
}
