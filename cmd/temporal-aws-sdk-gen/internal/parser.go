package internal

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

var methodSuffixesToSkip = []string{"WithContext", "Request", "Pages"}
var blacklistedMethods = map[string]string{
	// deprecated and conflicts with generated code for Invoke
	"Lambda": "InvokeAsync",
}

// ParseAwsSdk calls generator function for each AWS Service iDefinition
func ParseAwsSdk(serviceName string) ([]*InterfaceDefinition, error) {
	var pkgs []*packages.Package

	config := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedTypesInfo,
	}

	if serviceName == "" {
		allPkgs, err := packages.Load(config, "github.com/aws/aws-sdk-go/service/...")
		if err != nil {
			return nil, err
		}

		for _, pkg := range allPkgs {
			if strings.TrimPrefix(pkg.PkgPath, "github.com/aws/aws-sdk-go/service/") == pkg.Name {
				pkgs = append(pkgs, pkg)
			}
		}
	} else {
		servicePkgs, err := packages.Load(config, fmt.Sprintf("github.com/aws/aws-sdk-go/service/%s", serviceName))
		if err != nil {
			return nil, err
		}

		pkgs = append(pkgs, servicePkgs...)
	}

	var definitions []*InterfaceDefinition
	structs := make(map[string]*StructDefinition)

	for _, pkg := range pkgs {
		definition, serviceStructs, err := parseAwsService(pkg)
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

func parseAwsService(pkg *packages.Package) (*InterfaceDefinition, []*StructDefinition, error) {
	var apiFileName string
	for _, file := range pkg.GoFiles {
		if filepath.Base(file) == "api.go" {
			apiFileName = file

			break
		}
	}

	if apiFileName == "" {
		return nil, nil, errors.New("api.go is not found in the service")
	}

	interfaceFileName := filepath.Join(filepath.Dir(apiFileName), fmt.Sprintf("%siface", pkg.Name), "interface.go")
	definition, err := parseAwsServiceInterface(pkg.Name, interfaceFileName)
	if err != nil {
		return nil, nil, err
	}

	structs, err := parseAwsApi(strings.ToLower(pkg.Name), apiFileName)
	if err != nil {
		return nil, nil, err
	}
	return &definition, structs, nil
}

func parseAwsServiceInterface(serviceName string, fileName string) (iDefinition InterfaceDefinition, err error) {
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

	definition := v.definition
	definition.ID = serviceName

	return definition, nil
}

func parseAwsApi(packageName, fileName string) (structs []*StructDefinition, err error) {
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
