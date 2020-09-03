package main

import (
	"log"
	"os"
	"strings"
	"temporal.io/temporal-aws-sdk/awsgenerator/internal"
	"text/template"

	"github.com/urfave/cli/v2"
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
			EnvVars:     []string{"TEMPORAL_AWS_SDK_DIR"},
			Destination: &sdkDir,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "output-dir",
			Usage:       "generated code location",
			EnvVars:     []string{"TEMPORAL_OUTPUT_DIR"},
			Destination: &outputDir,
			Required:    true,
		},
	}
	app.Action = func(c *cli.Context) error {
		parser := &internal.AWSSDKParser{SdkDirectory: sdkDir}
		//return parser.ParseAwsService("cloudformation", func(service string, definition internal.InterfaceDefinition) error {
		//	fmt.Println(definition)
		//	return generateCode(templateDir, outputDir, definition)
		//})
		return parser.ParseAwsSdk(func(service string, definition internal.InterfaceDefinition) error {
			return generateCode(templateDir, outputDir, definition)
		})
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func toPrefix(packageName string) string {
	if packageName == "" {
		return packageName
	}
	return strings.ToUpper(packageName[0:1]) + packageName[1:]
}

func generateCode(templateDir string, outputDir string, definition internal.InterfaceDefinition) error {
	funcMap := template.FuncMap{
		"ToUpper":   strings.ToUpper,
		"ToLower":   strings.ToLower,
		"HasPrefix": strings.HasPrefix,
		"ToPrefix":  toPrefix,
	}

	//templateFile := "aws_activity.go.tmpl"
	templateFile := "client.go.tmpl"
	//templates, err := template.ParseGlob(templateDir + "/*")
	// https://stackoverflow.com/a/49043639/1664318 for the New parameter
	templates, err := template.New(templateFile).Funcs(funcMap).ParseFiles(templateDir + "/" + templateFile)
	if err != nil {
		return err
	}
	output := templateFile[0 : len(templateFile)-5]
	outputFile := outputDir + "/" + strings.ToLower(definition.Name) + "_" + output
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	return templates.Execute(f, definition)
}
