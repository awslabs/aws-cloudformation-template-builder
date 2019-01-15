package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"codecommit/builders/cfn-spec-go/spec"
	"fmt"
	"strings"
)

func main() {
	config := make(map[string]string)

	for resourceType, _ := range spec.Cfn.ResourceTypes {
		name := "My" + strings.Replace(resourceType, "::", "", -1)
		config[name] = resourceType
	}

	t := builder.Cfn.Template(config)
	fmt.Println(format.Yaml(t))

	p := builder.Iam.Policy()
	fmt.Println(format.Yaml(p))
}
