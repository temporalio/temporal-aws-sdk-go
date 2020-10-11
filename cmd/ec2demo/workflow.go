package ec2demo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.temporal.io/aws-sdk/clients/ec2stub"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func KeepInstance(ctx workflow.Context) error {
	ctx1 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	ctx2 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 24 * 364 * time.Hour,
		HeartbeatTimeout:    time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			NonRetryableErrorTypes: []string{"foo"},
		},
	})
	logger := workflow.GetLogger(ctx)
	ec2Client := ec2stub.NewClient()
	for i := 0; i < 1000; i++ {
		reservation, err := ec2Client.RunInstances(ctx1, &ec2.RunInstancesInput{
			ImageId:      aws.String("ami-0947d2ba12ee1ff75"),
			InstanceType: aws.String("t2.nano"),
			MinCount:     aws.Int64(1),
			MaxCount:     aws.Int64(1),
		})
		if err != nil {
			return err
		}
		instanceId := reservation.Instances[0].InstanceId
		logger.Info("Starting EC2 instance", "instanceId", *instanceId)

		err = ec2Client.WaitUntilInstanceTerminated(ctx2, &ec2.DescribeInstancesInput{
			InstanceIds: []*string{instanceId},
		})
		if err != nil {
			return err
		}
		logger.Info("EC2 instance terminated. Relaunching...", "instanceId", *instanceId)
	}
	// Start new run atomically to keep history size bounded
	return workflow.NewContinueAsNewError(ctx, KeepInstance)
}
