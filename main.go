package main

import (
	"codecommit/builders/cfn-spec-go/spec"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const specFn = "CloudFormationResourceSpecification.json"

var cfnSpec spec.Spec

func init() {
	f, err := os.Open(specFn)
	if err != nil {
		panic(err.Error())
	}

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&cfnSpec)
	if err != nil {
		fmt.Printf("%T\n", err)
		panic(err.Error())
	}
}

func template(types []string) string {
	typeCount := make(map[string]int)

	builder := strings.Builder{}

	builder.WriteString("Resources:\n")

	for _, typeName := range types {
		parts := strings.Split(typeName, "::")
		name := parts[len(parts)-1]

		typeCount[name]++

		res := cfnSpec.GetResource(typeName).Skeleton(cfnSpec).Value

		builder.WriteString(fmt.Sprintf("  %s%d:\n", name, typeCount[name]))
		builder.WriteString("    Type: " + typeName + "\n")
		builder.WriteString("    Properties:\n")
		builder.WriteString(spec.Indent(res, "      "))
		builder.WriteString("\n")
	}

	return builder.String()
}

func main() {
	types := []string{
		"AWS::CodeCommit::Repository",
		"AWS::CodeCommit::Repository",
	}

	fmt.Println(template(types))
}
