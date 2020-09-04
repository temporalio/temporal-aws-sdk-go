package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"temporal.io/aws-sdk/awsgenerator/internal"
)

func main() {
	var sdkDir, templateDir, outputDir string

	app := cli.NewApp()
	app.Name = "aws2temporal"
	app.Usage = "Generates Temporal Bindings for AWS SDK"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "template-dir",
			Value:       "./awsgenerator/templates",
			Usage:       "location of code generation template directory",
			EnvVars:     []string{"TEMPORAL_AWS_SDK_TEMPLATE_DIR"},
			Destination: &templateDir,
		},
		&cli.StringFlag{
			Name:        "aws-sdk-dir",
			Usage:       "location of AWS Go SDK repository",
			Value:       "./awssdkgo",
			EnvVars:     []string{"TEMPORAL_AWS_SDK_DIR"},
			Destination: &sdkDir,
		},
		&cli.StringFlag{
			Name:        "output-dir",
			Value:       ".",
			Usage:       "generated code location",
			EnvVars:     []string{"TEMPORAL_OUTPUT_DIR"},
			Destination: &outputDir,
		},
	}
	app.Action = func(c *cli.Context) error {
		parser := &internal.AWSSDKParser{SdkDirectory: sdkDir}
		generator := internal.NewGenerator(templateDir)
		//return parser.ParseAwsService("cloudformation", func(service string, definition internal.InterfaceDefinition) error {
		//	fmt.Println(definition)
		//	return generateCode(templateDir, outputDir, definition)
		//})
		return parser.ParseAwsSdk(func(service string, definition internal.InterfaceDefinition) error {
			return generator.GenerateCode(outputDir, definition)
		})
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
