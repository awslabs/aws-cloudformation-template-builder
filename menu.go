package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"fmt"
)

var s3Menu builder.Resource

func init() {
	s3Menu = builder.Resource{
		TypeName: "AWS::S3::Bucket",
		Output:   map[string]interface{}{},
		Menu: []builder.MenuItem{
			{
				Question: "Bucket name:",
				Output: map[string]interface{}{
					"BucketName": builder.PlaceHolder,
				},
			},
			{
				Question: "Is this a website?",
				Options: []builder.Option{
					{
						Name: "Yes",
						Output: map[string]interface{}{
							"WebsiteConfiguration": map[string]interface{}{
								"IndexDocument": "index.html",
							},
						},
					},
					{
						Name: "No",
					},
				},
			},
		},
	}
}

func main() {
	fmt.Println(s3Menu)

	output := map[string]interface{}{
		"Resources": map[string]interface{}{
			"MyBucket": s3Menu.Build(),
		},
	}

	// Check it hasn't changed
	fmt.Println(s3Menu)

	fmt.Println()

	fmt.Println(format.Yaml(output))
}
