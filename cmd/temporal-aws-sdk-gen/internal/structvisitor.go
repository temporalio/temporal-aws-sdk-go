package internal

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type AWSStructVisitor struct {
	fileSet       *token.FileSet
	definition    InterfaceDefinition
	currentMethod *MethodDefinition
}

func NewAWSStructVisitor(fileSet *token.FileSet) *AWSStructVisitor {
	return &AWSStructVisitor{fileSet: fileSet}
}

func (g *AWSStructVisitor) Visit(node ast.Node) ast.Visitor {
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

func (g *AWSStructVisitor) visitTypeSpec(n *ast.TypeSpec) {
	if g.definition.Name != "" {
		panic("Multiple interfaces in the interface definition file")
	}
	name := n.Name.Name
	// Remove trailing API
	g.definition.Name = name[:len(name)-3]
}

func (g *AWSStructVisitor) visitField(n *ast.Field) *AWSStructVisitor {
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

func (g *AWSStructVisitor) String() string {
	return fmt.Sprintf("%v", g.definition)
}

func (g *AWSStructVisitor) visitFuncType(n *ast.FuncType) {
	//ast.Print(g.fileSet, n.Params.List[0].Type)
	inputExpr := n.Params.List[0].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
	g.currentMethod.Input = &StructDefinition{Package: inputExpr.X.(*ast.Ident).Name, Name: inputExpr.Sel.Name}
	results := n.Results.List
	if len(results) == 2 {
		resultExpr := results[0].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
		g.currentMethod.Output = &StructDefinition{Package: resultExpr.X.(*ast.Ident).Name, Name: resultExpr.Sel.Name}
	}
}
