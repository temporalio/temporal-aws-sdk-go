package ec2demo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsclients"
	"time"
)

func KeepInstance(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	ec2Client := awsclients.NewEC2Stub()
	for {
		instanceId, err2 := launchInstance(ctx, ec2Client)
		if err2 != nil {
			return err2
		}

		err := ec2Client.WaitUntilInstanceTerminated(ctx, &ec2.DescribeInstancesInput{
			InstanceIds: []*string{instanceId},
		})
		if err != nil {
			return err
		}

		logger.Info("EC2 instance terminated.", "instanceId", *instanceId)
	}
	return nil
}

func launchInstance(ctx workflow.Context, ec2Client awsclients.EC2Client) (*string, error) {
	logger := workflow.GetLogger(ctx)
	reservation, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
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

	err = ec2Client.WaitUntilInstanceRunning(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []*string{instanceId},
	})
	if err != nil {
		return nil, err
	}
	logger.Info("EC2 instance running", "instanceId", *instanceId)

	return instanceId, nil
}
