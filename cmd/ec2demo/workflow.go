package ec2demo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsclients"
	"time"
)

func KeepInstance(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	ec2Client := awsclients.NewEC2Stub()
	for i := 0; i < 1000; i++ {
		instanceId, err2 := launchInstance(ctx, ec2Client)
		if err2 != nil {
			return err2
		}
		ctx1 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskQueue:           "aws-sdk",
			StartToCloseTimeout: 24 * 364 * time.Hour,
			HeartbeatTimeout:    time.Minute,
			RetryPolicy: &temporal.RetryPolicy{
				NonRetryableErrorTypes: []string{"foo"},

			},
		})
		err := ec2Client.WaitUntilInstanceTerminated(ctx1, &ec2.DescribeInstancesInput{
			InstanceIds: []*string{instanceId},
		})
		if err != nil {
			return err
		}

		logger.Info("EC2 instance terminated.", "instanceId", *instanceId)
	}
	return workflow.NewContinueAsNewError(ctx, KeepInstance)
}

func launchInstance(ctx workflow.Context, ec2Client awsclients.EC2Client) (*string, error) {
	logger := workflow.GetLogger(ctx)
	ctx1 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	reservation, err := ec2Client.RunInstances(ctx1, &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0947d2ba12ee1ff75"),
		InstanceType: aws.String("t2.nano"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})
	if err != nil {
		return nil, err
	}
	instanceId := reservation.Instances[0].InstanceId
	logger.Info("Starting EC2 instance", "instanceId", *instanceId)

	ctx2 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Minute,
		HeartbeatTimeout:    time.Minute,
	})
	err = ec2Client.WaitUntilInstanceRunning(ctx2, &ec2.DescribeInstancesInput{
		InstanceIds: []*string{instanceId},
	})
	if err != nil {
		return nil, err
	}
	logger.Info("EC2 instance running", "instanceId", *instanceId)

	return instanceId, nil
}
