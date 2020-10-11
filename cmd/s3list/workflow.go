package s3list

import (
	s32 "github.com/aws/aws-sdk-go/service/s3"
	"go.temporal.io/aws-sdk/awsclients/s3stub"
	"go.temporal.io/sdk/workflow"
	"time"
)

func Workflow(ctx workflow.Context) (*s32.ListBucketsOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	s3 := s3stub.NewClient()
	return s3.ListBuckets(ctx, &s32.ListBucketsInput{})
}
