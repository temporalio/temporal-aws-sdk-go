
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/batch/batchiface"
)

type BatchActivities struct {
	client batchiface.BatchAPI
}

func NewBatchActivities(client batchiface.BatchAPI) *BatchActivities {
    return &BatchActivities{client: client}
}

func (a *BatchActivities) CancelJob(input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
    return a.client.CancelJob(input)
}

func (a *BatchActivities) CreateComputeEnvironment(input *batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error) {
    return a.client.CreateComputeEnvironment(input)
}

func (a *BatchActivities) CreateJobQueue(input *batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error) {
    return a.client.CreateJobQueue(input)
}

func (a *BatchActivities) DeleteComputeEnvironment(input *batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error) {
    return a.client.DeleteComputeEnvironment(input)
}

func (a *BatchActivities) DeleteJobQueue(input *batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error) {
    return a.client.DeleteJobQueue(input)
}

func (a *BatchActivities) DeregisterJobDefinition(input *batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error) {
    return a.client.DeregisterJobDefinition(input)
}

func (a *BatchActivities) DescribeComputeEnvironments(input *batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error) {
    return a.client.DescribeComputeEnvironments(input)
}

func (a *BatchActivities) DescribeJobDefinitions(input *batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error) {
    return a.client.DescribeJobDefinitions(input)
}

func (a *BatchActivities) DescribeJobQueues(input *batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error) {
    return a.client.DescribeJobQueues(input)
}

func (a *BatchActivities) DescribeJobs(input *batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error) {
    return a.client.DescribeJobs(input)
}

func (a *BatchActivities) ListJobs(input *batch.ListJobsInput) (*batch.ListJobsOutput, error) {
    return a.client.ListJobs(input)
}

func (a *BatchActivities) RegisterJobDefinition(input *batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error) {
    return a.client.RegisterJobDefinition(input)
}

func (a *BatchActivities) SubmitJob(input *batch.SubmitJobInput) (*batch.SubmitJobOutput, error) {
    return a.client.SubmitJob(input)
}

func (a *BatchActivities) TerminateJob(input *batch.TerminateJobInput) (*batch.TerminateJobOutput, error) {
    return a.client.TerminateJob(input)
}

func (a *BatchActivities) UpdateComputeEnvironment(input *batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error) {
    return a.client.UpdateComputeEnvironment(input)
}

func (a *BatchActivities) UpdateJobQueue(input *batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error) {
    return a.client.UpdateJobQueue(input)
}
