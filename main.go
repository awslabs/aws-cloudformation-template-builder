package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"codecommit/builders/cfn-spec-go/spec"
	"fmt"
	"strings"
)

func main() {
	config := builder.NewTemplateConfig()

	for resourceType, _ := range spec.Cfn.ResourceTypes {
		name := "My" + strings.Replace(resourceType, "::", "", -1)
		config.Resources[name] = resourceType
	}

	t := builder.NewTemplate(config)

	fmt.Println(format.Yaml(t))
}
