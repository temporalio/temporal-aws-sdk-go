package main

import (
	"go.temporal.io/aws-sdk/cmd/ec2demo"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

const taskQueue = "ec2demo"

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, taskQueue, worker.Options{})

	w.RegisterWorkflow(ec2demo.KeepInstance)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
