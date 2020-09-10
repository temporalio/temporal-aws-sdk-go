package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
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
func (p *AWSSDKParser) ParseAwsSdk(serviceName string) ([]*InterfaceDefinition, error) {
	var serviceNames []string
	if serviceName == "" {
		serviceDir := p.SdkDirectory + "/service"
		files, err := ioutil.ReadDir(serviceDir)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			if f.IsDir() {
				serviceNames = append(serviceNames, f.Name())
			}
		}
	} else {
		serviceNames = append(serviceNames, serviceName)
	}
	var definitions []*InterfaceDefinition
	structs := make(map[string]*StructDefinition)
	for _, serviceName := range serviceNames {
		definition, serviceStructs, err := p.parseAwsService(serviceName)
		if err != nil {
			return nil, err
		}
		definitions = append(definitions, definition)
		for _, serviceStruct := range serviceStructs {
			structs[serviceStruct.Package+"."+serviceStruct.Name] = serviceStruct
		}
	}
	for _, definition := range definitions {
		for _, method := range definition.Methods {
			if method.Output != nil {
				method.Output = structs[method.Output.Package+"."+method.Output.Name]
			}
			method.Input = structs[method.Input.Package+"."+method.Input.Name]
		}
	}
	return definitions, nil
}

func (p *AWSSDKParser) parseAwsService(serviceName string) (*InterfaceDefinition, []*StructDefinition, error) {
	serviceDir := p.SdkDirectory + "/service"
	fileName := serviceDir + "/" + serviceName + "/" + serviceName + "iface/interface.go"
	definition, err := p.parseAwsServiceInterface(fileName)
	if err != nil {
		return nil, nil, err
	}
	apiFileName := serviceDir + "/" + serviceName + "/api.go"
	structs, err := p.parseAwsApi(strings.ToLower(serviceName), apiFileName)
	if err != nil {
		return nil, nil, err
	}
	return &definition, structs, nil
}

func (p *AWSSDKParser) parseAwsServiceInterface(fileName string) (iDefinition InterfaceDefinition, err error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return InterfaceDefinition{}, err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, string(body), parser.ParseComments)
	if err != nil {
		return InterfaceDefinition{}, err
	}
	v := NewAWSInterfaceVisitor(fileSet)
	ast.Walk(v, file)
	return v.definition, nil
}

func (p *AWSSDKParser) parseAwsApi(packageName, fileName string) (structs []*StructDefinition, err error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, string(body), parser.ParseComments)
	if err != nil {
		return nil, err
	}
	v := NewAWSStructVisitor(fileSet, packageName)
	ast.Walk(v, file)
	return v.Structs, nil
}
