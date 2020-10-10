package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"temporal.io/aws-sdk/awsactivities"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Failed to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "aws-sdk", worker.Options{})

	// TODO(maxim): command line flags to override session and config options
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Fatalln("Failed to create AWS session", err)
	}
	awsactivities.RegisterAwsActivities(w, sess)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Failed to start worker", err)
	}
}
