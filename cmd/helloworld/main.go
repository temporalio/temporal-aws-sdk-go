package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"temporal.io/aws-sdk/cmd/helloworld/workflow"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "hello-s3", worker.Options{})

	w.RegisterWorkflow(workflow.Workflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

	//sess, err := session.NewSession(&aws.Config{
	//	Region: aws.String("us-west-2")},
	//)
	//
	//svc := s3.New(sess)
	//
	//result, err := svc.ListBuckets(nil)
	//if err != nil {
	//	panic(fmt.Errorf("Unable to list buckets, %v", err))
	//}
	//
	//fmt.Println("Buckets:")
	//
	//for _, b := range result.Buckets {
	//	fmt.Printf("* %s created on %s\n",
	//		aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	//}
}