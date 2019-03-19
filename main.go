package main

import (
	"fmt"
	"os"

	"github.com/awslabs/aws-cloudformation-template-builder/builder"

	"github.com/awslabs/aws-cloudformation-template-formatter/format"
)

const (
	styleJSON = "json"
	styleYAML = "yaml"
)

const usage = `Usage: cfn-skeleton [OPTIONS] [RESOURCE TYPES...]

  cfn-skeleton is a tool that generates skeleton CloudFormation templates
  containing definitions for the resource types that you specify.

  You can use a short name for a resource type so long as it is unambiguous.
  For example 'Bucket', 'S3::Bucket', and 'AWS::S3::Bucket' refer to the same type.
  But 'Instance' would need disambiguation.

Options:

  -b, --bare  Produce a minimal template, omitting all optional resource properties.
  -i, --iam   If any resource includes an IAM policy definition, populate that too.
  -j, --json  Output the template in JSON format (default: YAML).
  --help      Show this message and exit.
`

func die() {
	fmt.Fprint(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die()
	}

	// TODO: Use a CLI library
	includeOptional := true
	buildIamPolicies := false
	style := styleYAML
	resourceTypes := make([]string, 0)

	// Parse params
	for _, arg := range os.Args[1:] {
		switch {
		case arg == "-b" || arg == "--bare":
			includeOptional = false
		case arg == "-i" || arg == "--iam":
			buildIamPolicies = true
		case arg == "-j" || arg == "--json":
			style = styleJSON
		case arg == "--help":
			die()
		case arg[0] == '-':
			die()
		default:
			resourceTypes = append(resourceTypes, arg)
		}
	}

	// Refuse to build an empty template?
	if len(resourceTypes) == 0 {
		die()
	}

	// Resolve resource types
	resources := resolveResources(resourceTypes)

	// Generate the template
	b := builder.NewCfnBuilder(includeOptional, buildIamPolicies)
	t, c := b.Template(resources)

	// Output the result
	if style == styleJSON {
		fmt.Println(format.JsonWithComments(t, c))
	} else {
		fmt.Println(format.YamlWithComments(t, c))
	}
}
