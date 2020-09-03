package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane/iotjobsdataplaneiface"
)

type IoTJobsDataPlaneActivities struct {
	client iotjobsdataplaneiface.IoTJobsDataPlaneAPI
}

func NewIoTJobsDataPlaneActivities(client iotjobsdataplaneiface.IoTJobsDataPlaneAPI) *IoTJobsDataPlaneActivities {
    return &IoTJobsDataPlaneActivities{client: client}
}


func (a *IoTJobsDataPlaneActivities) DescribeJobExecution(input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
    return a.client.DescribeJobExecution(input)
}



func (a *IoTJobsDataPlaneActivities) GetPendingJobExecutions(input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
    return a.client.GetPendingJobExecutions(input)
}



func (a *IoTJobsDataPlaneActivities) StartNextPendingJobExecution(input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error) {
    return a.client.StartNextPendingJobExecution(input)
}



func (a *IoTJobsDataPlaneActivities) UpdateJobExecution(input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error) {
    return a.client.UpdateJobExecution(input)
}

