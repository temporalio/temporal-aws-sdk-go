package internal

import (
	"fmt"
	"sort"
)

type FieldDefinition struct {
	Name string
	Type string
}

type StructDefinition struct {
	Package string
	Name    string
	Fields  []*FieldDefinition
}

type MethodDefinition struct {
	// Method name
	Name string
	// Input structure
	Input *StructDefinition
	// Output structure
	Output *StructDefinition
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
		packages[method.Input.Package] = true
		if method.Output != nil {
			packages[method.Output.Package] = true
		}
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
