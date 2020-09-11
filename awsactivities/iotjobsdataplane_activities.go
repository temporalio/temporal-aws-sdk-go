
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane/iotjobsdataplaneiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type IoTJobsDataPlaneActivities struct {
    client iotjobsdataplaneiface.IoTJobsDataPlaneAPI
}

func NewIoTJobsDataPlaneActivities(session *session.Session, config ...*aws.Config) *IoTJobsDataPlaneActivities {
    client := iotjobsdataplane.New(session, config...)
    return &IoTJobsDataPlaneActivities{client: client}
}

func (a *IoTJobsDataPlaneActivities) DescribeJobExecution(ctx context.Context, input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
    return a.client.DescribeJobExecutionWithContext(ctx, input)
}

func (a *IoTJobsDataPlaneActivities) GetPendingJobExecutions(ctx context.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
    return a.client.GetPendingJobExecutionsWithContext(ctx, input)
}

func (a *IoTJobsDataPlaneActivities) StartNextPendingJobExecution(ctx context.Context, input *iotjobsdataplane.StartNextPendingJobExecutionInput) (*iotjobsdataplane.StartNextPendingJobExecutionOutput, error) {
    return a.client.StartNextPendingJobExecutionWithContext(ctx, input)
}

func (a *IoTJobsDataPlaneActivities) UpdateJobExecution(ctx context.Context, input *iotjobsdataplane.UpdateJobExecutionInput) (*iotjobsdataplane.UpdateJobExecutionOutput, error) {
    return a.client.UpdateJobExecutionWithContext(ctx, input)
}
