package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func toPrefix(packageName string) string {
	if packageName == "" {
		return packageName
	}
	return strings.ToUpper(packageName[0:1]) + packageName[1:]
}

type TemporalAWSGenerator struct {
	TemplateDir   string
	outputStructs map[string]bool
}

func NewGenerator(templateDir string) *TemporalAWSGenerator {
	return &TemporalAWSGenerator{TemplateDir: templateDir, outputStructs: make(map[string]bool)}
}

func (g *TemporalAWSGenerator) GenerateCode(outputDir string, definition InterfaceDefinition) error {
	var templateFiles []string
	files, err := ioutil.ReadDir(g.TemplateDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if !f.IsDir() {
			templateName := f.Name()
			if strings.HasSuffix(templateName, ".tmpl") {
				templateFiles = append(templateFiles, strings.TrimSuffix(templateName, ".tmpl"))
			}
		}
	}
	for _, templateFile := range templateFiles {
		err := g.generateFromSingleTemplate(templateFile+".tmpl", outputDir+"/aws"+templateFile, definition)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *TemporalAWSGenerator) generateFromSingleTemplate(templateFile string, outputDir string, definition InterfaceDefinition) error {
	funcMap := template.FuncMap{
		"ToUpper":   strings.ToUpper,
		"ToLower":   strings.ToLower,
		"HasPrefix": strings.HasPrefix,
		"ToPrefix":  toPrefix,
		"IsDuplicatedOutput": func(outputPackage, output string) bool {
			_, ok := g.outputStructs[outputPackage+output]
			return ok
		},
	}
	templates, err := template.New(templateFile).Funcs(funcMap).ParseFiles(g.TemplateDir + "/" + templateFile)
	if err != nil {
		return err
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.Mkdir(outputDir, 0700)
		if err != nil {
			return err
		}
	}
	output := templateFile[0 : len(templateFile)-5]
	outputFile := outputDir + "/" + strings.ToLower(definition.Name) + "_" + output + ".go"
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	err = templates.Execute(f, definition)
	if err != nil {
		return err
	}
	for _, method := range definition.Methods {
		g.outputStructs[method.OutputPackage+method.Output] = true
	}
	fmt.Println(outputFile)
	return nil
}
