package s3list

import (
	"github.com/aws/aws-sdk-go/aws"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsclients"
	"github.com/aws/aws-sdk-go/service/ec2"
	"time"
)

func Workflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	ec2Service := awsclients.NewEC2Stub()
	runInstances := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-e7527ed7"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}
	ec2Service.RunInstances(ctx, runInstances)
	return nil
}
