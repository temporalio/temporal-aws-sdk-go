package awsgenerator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

type AWSSDKParser struct {
	SdkDirectory string
}

func (p *AWSSDKParser) ProcessAwsSdk(generator func(service string, interfaces map[string][]string) error) error {
	serviceDir := p.SdkDirectory + "/service"
	files, err := ioutil.ReadDir(serviceDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			serviceName := f.Name()
			fileName := serviceDir + "/" + serviceName + "/" + serviceName + "iface/interface.go"
			err2 := p.ProcessAwsService(generator, fileName, serviceName)
			if err2 != nil {
				return err2
			}
		}
	}
	return nil
}

func (p *AWSSDKParser) ProcessAwsService(generator func(service string, interfaces map[string][]string) error, fileName string, serviceName string) error {
	fmt.Println(fileName)
	interfaces, err := p.parse(fileName)
	if err != nil {
		return fmt.Errorf("failure parsing %v: %w", fileName, err)
	}
	err = generator(serviceName, interfaces)
	if err != nil {
		return fmt.Errorf("failure processing %v: %w", fileName, err)
	}
	return nil
}

func (p *AWSSDKParser) parse(fileName string) (interfaces map[string][]string, err error) {
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, fileName, string(body), parser.ParseComments)
	if err != nil {
		return nil, err
	}
	v := NewAWSInterfaceVisitor()
	ast.Walk(v, file)
	return v.interfaces, nil
}

type AWSInterfaceVisitor struct {
	typeName string
	// key is interface name
	// value is the list of methods
	interfaces map[string][]string
}

func NewAWSInterfaceVisitor() *AWSInterfaceVisitor {
	return &AWSInterfaceVisitor{interfaces: make(map[string][]string)}
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
	g.typeName = n.Name.Name
}

func (g *AWSInterfaceVisitor) visitField(n *ast.Field) {
	if len(n.Names) > 0 {
		name := n.Names[0].Name
		if !strings.HasSuffix(name, "WithContext") && !strings.HasSuffix(name, "Request") {
			methods := g.interfaces[g.typeName]
			methods = append(methods, name)
			g.interfaces[g.typeName] = methods
		}
	}
}

func (g *AWSInterfaceVisitor) String() string {
	return fmt.Sprintf("%v", g.interfaces)
}
