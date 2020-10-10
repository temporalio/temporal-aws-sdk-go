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
	Fields  map[string]*FieldDefinition
}

func NewStructDefinition(pkg, name string) *StructDefinition {
	return &StructDefinition{Package: pkg, Name: name, Fields: make(map[string]*FieldDefinition)}
}

type MethodDefinition struct {
	// Method name
	Name string
	// Input structure
	Input *StructDefinition
	// Output structure
	Output *StructDefinition
}

func (m *MethodDefinition) String() string {
	return fmt.Sprintf("%v(%v) %v", m.Name, m.Input, m.Output)
}

type InterfaceDefinition struct {
	// Service name (lower-cased, URL compatible service name)
	ID string
	// Service name (friendly name)
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
