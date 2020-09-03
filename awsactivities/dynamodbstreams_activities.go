package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
)

type DynamoDBStreamsActivities struct {
	client dynamodbstreamsiface.DynamoDBStreamsAPI
}

func NewDynamoDBStreamsActivities(client dynamodbstreamsiface.DynamoDBStreamsAPI) *DynamoDBStreamsActivities {
    return &DynamoDBStreamsActivities{client: client}
}

func (a *DynamoDBStreamsActivities) DescribeStream(input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
    return a.client.DescribeStream(input)
}

func (a *DynamoDBStreamsActivities) GetRecords(input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
    return a.client.GetRecords(input)
}

func (a *DynamoDBStreamsActivities) GetShardIterator(input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
    return a.client.GetShardIterator(input)
}

func (a *DynamoDBStreamsActivities) ListStreams(input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
    return a.client.ListStreams(input)
}
