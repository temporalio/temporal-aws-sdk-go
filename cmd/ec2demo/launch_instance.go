package ec2demo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsclients"
	"time"
)

func LaunchInstance(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	ec2Service := awsclients.NewEC2Stub()
	runInstances := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-0947d2ba12ee1ff75"),
		InstanceType: aws.String("t2.nano"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}
	reservation, err := ec2Service.RunInstances(ctx, runInstances)
	if err != nil {
		return err
	}
	describe := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{reservation.Instances[0].InstanceId},
	}
	ec2Service.WaitUntilInstanceRunning(ctx, describe)
	return nil
}
