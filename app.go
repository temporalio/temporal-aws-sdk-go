package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	gen "temporal.io/temporal-aws-sdk/awsgenerator"
)

const (
	// Version is the controlled version string. It should be updated every time
	// before we release a new version.
	Version = "0.0.1"

	// Flags
	FlagTemplateDir = "template-dir"
	FlagAwsSdkDir   = "aws-sdk-dir"
)

func main() {
	app := cli.NewApp()
	app.Name = "aws2temporal"
	app.Usage = "Generates Temporal Bindings for AWS SDK"
	app.Version = Version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:   FlagTemplateDir,
			Value:  ".",
			Usage:  "location of code generation template directory",
			EnvVar: "TEMPORAL_AWS_SDK_TEMPLATE_DIR",
		},
		&cli.StringFlag{
			Name:   FlagAwsSdkDir,
			Value:  ".",
			Usage:  "location of AWS Go SDK repository",
			EnvVar: "TEMPORAL_AWS_SDK_DIR",
		},
	}
	app.Action = func(c *cli.Context) error {
		parser := &gen.AWSSDKParser{SdkDirectory: "/Users/maxim/support/aws-sdk-go"}
		return parser.ProcessAwsSdk(func(service string, interfaces map[string][]string) error {
			fmt.Println(interfaces)
			return nil
		})
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
