package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

type AWSStructVisitor struct {
	packageName string
	fileSet     *token.FileSet
	Structs     []*StructDefinition
	Current     *StructDefinition
}

func NewAWSStructVisitor(fileSet *token.FileSet, packageName string) *AWSStructVisitor {
	return &AWSStructVisitor{fileSet: fileSet, packageName: packageName}
}

func (g *AWSStructVisitor) Visit(node ast.Node) ast.Visitor {
	if g == nil {
		return nil
	}
	switch n := node.(type) {
	case *ast.TypeSpec:
		return g.visitTypeSpec(n)

	case *ast.FuncType:
		g.visitFuncType(n)
		return g

	case *ast.Field:
		return g.visitField(n)

	default:
		return g
	}
}

func (g *AWSStructVisitor) visitTypeSpec(n *ast.TypeSpec) *AWSStructVisitor {
	if _, ok := n.Type.(*ast.StructType); !ok {
		return nil
	}
	name := n.Name.Name
	//if !strings.HasSuffix(name, "Input") && !strings.HasSuffix(name, "Output") {
	//	return nil
	//}
	g.Current = NewStructDefinition(g.packageName, name)
	g.Structs = append(g.Structs, g.Current)
	return g
}

func (g *AWSStructVisitor) visitField(n *ast.Field) *AWSStructVisitor {
	defer func() {
		if r := recover(); r != nil {
			ast.Print(g.fileSet, n)
			panic(r)
		}
	}()
	if len(n.Names) > 0 {
		name := n.Names[0].Name
		if len(name) <= 1 || g.Current == nil {
			return nil
		}
		typeName := GetTypeName(n.Type)
		//fmt.Printf("    %v: %v\n", name, typeName)
		g.Current.Fields[name] = &FieldDefinition{Name: name, Type: typeName}
	}
	return g
}

func GetTypeName(n ast.Node) string {
	switch t := n.(type) {
	case *ast.StarExpr:
		return GetTypeName(t.X)
	case *ast.ArrayType:
		return "[]" + GetTypeName(t.Elt)
	case *ast.MapType:
		return "map[" + GetTypeName(t.Key) + "]" + GetTypeName(t.Value)
	case *ast.SelectorExpr:
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name
	case *ast.Ident:
		return t.Name
	case *ast.Ellipsis:
		return "[]" + GetTypeName(t.Elt)
	case *ast.FuncType:
		return "func"
	case *ast.ChanType:
		return "channel"
	}
	panic(fmt.Sprintf("unexpected type: %v", n))
}

func (g *AWSStructVisitor) String() string {
	return fmt.Sprintf("%v", g.Structs)
}

func (g *AWSStructVisitor) visitFuncType(n *ast.FuncType) {
	//ast.Print(g.fileSet, n.Params.List[0].Type)
	//inputExpr := n.Params.List[0].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
	//g.currentMethod.Input = &StructDefinition{Package: inputExpr.X.(*ast.Ident).Name, Name: inputExpr.Sel.Name}
	//results := n.Results.List
	//if len(results) == 2 {
	//	resultExpr := results[0].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
	//	g.currentMethod.Output = &StructDefinition{Package: resultExpr.X.(*ast.Ident).Name, Name: resultExpr.Sel.Name}
	//}
}
