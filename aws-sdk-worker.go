package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/urfave/cli/v2"
	"go.temporal.io/aws-sdk/activities"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"os"
)

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var address, namespace, taskQueue, region string
	var verbose bool
	app := cli.NewApp()
	app.Name = "aws-activities"
	app.Usage = "Temporal activities for each AWS SDK API"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "address",
			Required:    false,
			Usage:       "host:port for Temporal frontend service, default is localhost:7233",
			Value:       "localhost:7233",
			Destination: &address,
			EnvVars:     []string{"TEMPORAL_ADDRESS"},
		},
		&cli.StringFlag{
			Name:        "namespace",
			Required:    false,
			Usage:       "Temporal namespace for worker to connect, default is \"default\"",
			Destination: &namespace,
			EnvVars:     []string{"TEMPORAL_NAMESPACE"},
		},
		&cli.StringFlag{
			Name:        "task-queue",
			Usage:       "task queue to start activity worker on, default is \"aws-sdk\"",
			Required:    false,
			Value:       "aws-sdk",
			Destination: &taskQueue,
		},
		&cli.StringFlag{
			Name:        "region",
			Usage:       "aws-region to send requests to",
			Required:    true,
			Destination: &region,
			EnvVars:     []string{"AWS_REGION"},
		},
		&cli.BoolFlag{
			Name:        "verbose",
			Usage:       "Enables LogDebug level of AWS SDK",
			Destination: &verbose,
		},
	}
	app.Action = func(context *cli.Context) error {
		return runWorker(address, namespace, taskQueue, verbose, region)
	}
	return app.Run(args)
}

func runWorker(address, namespace, taskQueue string, verbose bool, region string) error {
	c, err := client.NewClient(client.Options{
		HostPort:  address,
		Namespace: namespace,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	w := worker.New(c, taskQueue, worker.Options{})

	var logLevel aws.LogLevelType
	if verbose {
		logLevel = aws.LogDebug
	} else {
		logLevel = aws.LogOff
	}
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		LogLevel: &logLevel,
	})
	if err != nil {
		return err
	}
	activities.RegisterAwsActivities(w, sess)
	return w.Run(worker.InterruptCh())
}
