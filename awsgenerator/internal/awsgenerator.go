package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

var methodSuffixesToSkip = []string{"WithContext", "Request", "Pages"}

type InterfaceDefinition struct {
	// Service name
	Name string
	// Service method names
	Methods []string
}

type AWSSDKParser struct {
	SdkDirectory string
}

// ParseAwsSdk calls generator function for each AWS Service definition
func (p *AWSSDKParser) ParseAwsSdk(generator func(service string, interfaces InterfaceDefinition) error) error {
	serviceDir := p.SdkDirectory + "/service"
	files, err := ioutil.ReadDir(serviceDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			serviceName := f.Name()
			fileName := serviceDir + "/" + serviceName + "/" + serviceName + "iface/interface.go"
			err2 := p.processAwsService(serviceName, fileName, generator)
			if err2 != nil {
				return err2
			}
		}
	}
	return nil
}

func (p *AWSSDKParser) ParseAwsService(serviceName string, generator func(service string, definition InterfaceDefinition) error) error {
	serviceDir := p.SdkDirectory + "/service"
	fileName := serviceDir + "/" + serviceName + "/" + serviceName + "iface/interface.go"
	return p.processAwsService(serviceName, fileName, generator)
}

func (p *AWSSDKParser) processAwsService(serviceName string, fileName string, generator func(service string, definition InterfaceDefinition) error) error {
	fmt.Println(fileName)
	definition, err := p.parse(fileName)
	if err != nil {
		return fmt.Errorf("failure parsing %v: %w", fileName, err)
	}
	err = generator(serviceName, definition)
	if err != nil {
		return fmt.Errorf("failure processing %v: %w", fileName, err)
	}
	return nil
}

func (p *AWSSDKParser) parse(fileName string) (definition InterfaceDefinition, err error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return InterfaceDefinition{}, err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, string(body), parser.ParseComments)
	if err != nil {
		return InterfaceDefinition{}, err
	}
	v := NewAWSInterfaceVisitor()
	ast.Walk(v, file)
	return v.definition, nil
}

type AWSInterfaceVisitor struct {
	definition InterfaceDefinition
}

func NewAWSInterfaceVisitor() *AWSInterfaceVisitor {
	return &AWSInterfaceVisitor{}
}

func (g *AWSInterfaceVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.TypeSpec:
		g.visitTypeSpec(n)
		return g

	case *ast.FuncType:
		return nil

	case *ast.Field:
		g.visitField(n)
		return nil
	default:
		return g
	}
}

func (g *AWSInterfaceVisitor) visitTypeSpec(n *ast.TypeSpec) {
	if g.definition.Name != "" {
		panic("Multiple interfaces in the interface definition file")
	}
	name := n.Name.Name
	// Remove trailing API
	g.definition.Name = name[:len(name)-3]
}

func (g *AWSInterfaceVisitor) visitField(n *ast.Field) {
	if len(n.Names) > 0 {
		name := n.Names[0].Name
		for _, postfix := range methodSuffixesToSkip {
			if strings.HasSuffix(name, postfix) {
				return
			}
		}
		g.definition.Methods = append(g.definition.Methods, name)
	}
}

func (g *AWSInterfaceVisitor) String() string {
	return fmt.Sprintf("%v", g.definition)
}
