package s3list

import (
	s32 "github.com/aws/aws-sdk-go/service/s3"
	"go.temporal.io/aws-sdk/awsclients"
	"go.temporal.io/sdk/workflow"
	"time"
)

func Workflow(ctx workflow.Context) (*s32.ListBucketsOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	s3 := awsclients.NewS3Stub()
	return s3.ListBuckets(ctx, &s32.ListBucketsInput{})
}
