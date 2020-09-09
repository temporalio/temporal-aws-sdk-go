package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
	"temporal.io/aws-sdk/cmd/temporal-aws-sdk-gen/internal"
)

func main() {
	var sdkDir, templateDir, outputDir, service string

	app := cli.NewApp()
	app.Name = "temporal-aws-sdk-gen"
	app.Usage = "Generates Temporal Bindings for AWS SDK"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "template-dir",
			Value:       "./cmd/temporal-aws-sdk-gen/templates",
			Usage:       "location of code generation template directory",
			Destination: &templateDir,
		},
		&cli.StringFlag{
			Name:        "aws-sdk-dir",
			Usage:       "location of AWS Go SDK repository",
			Value:       "./awssdkgo",
			Destination: &sdkDir,
		},
		&cli.StringFlag{
			Name:        "output-dir",
			Value:       ".",
			Usage:       "generated code location",
			Destination: &outputDir,
		},
		&cli.StringFlag{
			Name:        "service",
			Usage:       "service to regenerate, default is all services",
			Destination: &service,
		},
	}
	app.Action = func(c *cli.Context) (err error) {
		parser := &internal.AWSSDKParser{SdkDirectory: sdkDir}
		generator := internal.NewGenerator(templateDir)
		s := strings.ToLower(service)
		var definitions []*internal.InterfaceDefinition
		if s != "" {
			err = parser.ParseAwsService(s, func(service string, definition *internal.InterfaceDefinition, sDefinition *internal.StructDefinition) error {
				definitions = append(definitions, definition)
				return nil
			})
		} else {
			err = parser.ParseAwsSdk(func(service string, definition *internal.InterfaceDefinition, sDefinition *internal.StructDefinition) error {
				definitions = append(definitions, definition)
				return nil
			})
		}
		if err != nil {
			log.Fatal(err)
		}
		return generator.GenerateCode(outputDir, definitions)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
