package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"codecommit/builders/cfn-spec-go/spec"
	"fmt"
)

func main() {
	for name, _ := range spec.Cfn.ResourceTypes {
		t := builder.NewTemplate(map[string]string{
			"MyThing": name,
		})

		fmt.Println(format.Yaml(t))
	}
}
