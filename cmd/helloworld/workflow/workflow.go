package workflow

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	s32 "github.com/aws/aws-sdk-go/service/s3"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsclients"
	"time"
)

func Workflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "aws-sdk",
		StartToCloseTimeout: 10 * time.Second,
	})
	s3 := awsclients.NewS3Stub()
	result, err := s3.ListBuckets(ctx, &s32.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("Unable to list buckets, %w", err)
	}
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	return nil
}
