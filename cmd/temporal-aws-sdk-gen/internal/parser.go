package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"sort"
	"strings"
)

var methodSuffixesToSkip = []string{"WithContext", "Request", "Pages"}
var blacklistedMethods = map[string]string{
	// deprecated and conflicts with generated code for Invoke
	"Lambda": "InvokeAsync",
}

type MethodDefinition struct {
	// Method name
	Name string
	// Input package name
	InputPackage string
	// Input structure name
	Input string
	// Output package name
	OutputPackage string
	// Output structure name
	Output string
}
type InterfaceDefinition struct {
	// Service name
	Name string
	// Service methods
	Methods []*MethodDefinition
}

func (i InterfaceDefinition) Imports() []string {
	packages := make(map[string]bool)
	for _, method := range i.Methods {
		packages[method.InputPackage] = true
		packages[method.OutputPackage] = true
	}
	delete(packages, "")
	result := make([]string, 0, len(packages))
	for p := range packages {
		result = append(result, p)
	}
	sort.Strings(result)
	return result
}

func (i InterfaceDefinition) String() string {
	result := i.Name + ": "
	for i, method := range i.Methods {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%v", method.Name)
	}
	return result
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
	v := NewAWSInterfaceVisitor(fileSet)
	ast.Walk(v, file)
	return v.definition, nil
}

type AWSInterfaceVisitor struct {
	fileSet       *token.FileSet
	definition    InterfaceDefinition
	currentMethod *MethodDefinition
}

func NewAWSInterfaceVisitor(fileSet *token.FileSet) *AWSInterfaceVisitor {
	return &AWSInterfaceVisitor{fileSet: fileSet}
}

func (g *AWSInterfaceVisitor) Visit(node ast.Node) ast.Visitor {
	if g == nil {
		return nil
	}
	switch n := node.(type) {
	case *ast.TypeSpec:
		g.visitTypeSpec(n)
		return g

	case *ast.FuncType:
		g.visitFuncType(n)
		return g

	case *ast.Field:
		return g.visitField(n)

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

func (g *AWSInterfaceVisitor) visitField(n *ast.Field) *AWSInterfaceVisitor {
	if len(n.Names) > 0 {
		name := n.Names[0].Name
		for _, postfix := range methodSuffixesToSkip {
			if strings.HasSuffix(name, postfix) {
				return nil
			}
		}
		if blacklistedMethods[g.definition.Name] == name {
			return nil
		}
		g.currentMethod = &MethodDefinition{Name: name}
		g.definition.Methods = append(g.definition.Methods, g.currentMethod)
		//ast.Print(g.fileSet, n)
	}
	return g
}

func (g *AWSInterfaceVisitor) String() string {
	return fmt.Sprintf("%v", g.definition)
}

func (g *AWSInterfaceVisitor) visitFuncType(n *ast.FuncType) {
	//ast.Print(g.fileSet, n.Params.List[0].Type)
	inputExpr := n.Params.List[0].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
	g.currentMethod.InputPackage = inputExpr.X.(*ast.Ident).Name
	g.currentMethod.Input = inputExpr.Sel.Name
	results := n.Results.List
	if len(results) == 2 {
		resultExpr := results[0].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
		g.currentMethod.OutputPackage = resultExpr.X.(*ast.Ident).Name
		g.currentMethod.Output = resultExpr.Sel.Name
	}
}
