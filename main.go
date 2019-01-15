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

	config = map[string]string{
		"MyBucket": "AWS::S3::Bucket",
	}

	cfn := builder.NewCfnBuilder(false)
	t := cfn.Template(config)
	fmt.Println(format.Yaml(t))

	/*
		fmt.Println()
		fmt.Println("---")
		fmt.Println()

		iam := builder.NewIamBuilder()
		p := map[string]interface{}{
			"PolicyDocument": iam.Policy(),
		}
		fmt.Println(format.Json(p))
	*/
}
