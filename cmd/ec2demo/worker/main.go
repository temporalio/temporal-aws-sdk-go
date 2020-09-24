package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	s32 "github.com/aws/aws-sdk-go/service/s3"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"temporal.io/aws-sdk/cmd/s3list"
)

const taskQueue = "s3list"

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, taskQueue, worker.Options{})

	w.RegisterWorkflow(s3list.Workflow)

	err = w.Start()
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
	defer w.Stop()

	options := client.StartWorkflowOptions{
		TaskQueue: taskQueue,
	}
	wRun, err := c.ExecuteWorkflow(context.Background(), options, s3list.Workflow)
	if err != nil {
		log.Fatalln("Failure starting workflow", err)
	}
	var result s32.ListBucketsOutput
	wRun.Get(context.Background(), &result)
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
