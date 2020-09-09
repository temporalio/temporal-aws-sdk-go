package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

var methodSuffixesToSkip = []string{"WithContext", "Request", "Pages"}
var blacklistedMethods = map[string]string{
	// deprecated and conflicts with generated code for Invoke
	"Lambda": "InvokeAsync",
}

type AWSSDKParser struct {
	SdkDirectory string
}

// ParseAwsSdk calls generator function for each AWS Service iDefinition
func (p *AWSSDKParser) ParseAwsSdk(generator func(service string, iDefinition *InterfaceDefinition, sDefinition *StructDefinition) error) error {
	serviceDir := p.SdkDirectory + "/service"
	files, err := ioutil.ReadDir(serviceDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			serviceName := f.Name()
			fileName := serviceDir + "/" + serviceName + "/" + serviceName + "iface/interface.go"
			err := p.parseAwsServiceInterface(serviceName, fileName, generator)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *AWSSDKParser) ParseAwsService(serviceName string, generator func(service string, iDefinition *InterfaceDefinition, sDefinition *StructDefinition) error) error {
	serviceDir := p.SdkDirectory + "/service"
	fileName := serviceDir + "/" + serviceName + "/" + serviceName + "iface/interface.go"
	err := p.parseAwsServiceInterface(serviceName, fileName, generator)
	if err != nil {
		return err
	}
	//apiFileName := serviceDir + "/" + serviceName + "/api.go"
	//err = p.parseAwsApi(serviceName, apiFileName, generator)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (p *AWSSDKParser) parseAwsServiceInterface(serviceName string, fileName string, generator func(service string, iDefinition *InterfaceDefinition, sDefinition *StructDefinition) error) error {
	iDefinition, err := p.parseAwsServiceInterfaceFile(fileName)
	if err != nil {
		return fmt.Errorf("failure parsing %v: %w", fileName, err)
	}
	err = generator(serviceName, iDefinition, nil)
	if err != nil {
		return fmt.Errorf("failure processing %v: %w", fileName, err)
	}
	return nil
}

func (p *AWSSDKParser) parseAwsServiceInterfaceFile(fileName string) (iDefinition *InterfaceDefinition, err error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, string(body), parser.ParseComments)
	if err != nil {
		return nil, err
	}
	v := NewAWSInterfaceVisitor(fileSet)
	ast.Walk(v, file)
	return v.definition, nil
}
