package main

import (
	"fmt"
	"os"

	"github.com/awslabs/aws-cloudformation-template-builder/builder"
	"github.com/awslabs/aws-cloudformation-template-formatter/format"
	flag "github.com/spf13/pflag"
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

var bareFlag bool
var iamFlag bool
var jsonFlag bool

func init() {
	flag.BoolVarP(&bareFlag, "bare", "b", false, "Produce a minimal template, omitting all optional resource properties.")
	flag.BoolVarP(&iamFlag, "iam", "i", false, "If any resource includes an IAM policy definition, populate that too.")
	flag.BoolVarP(&jsonFlag, "json", "j", false, "Output the template in JSON format (default: YAML).")
}
func die() {
	fmt.Fprint(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	//Parse the flags
	flag.Parse()
	//Get what ever's left after the flags have been parsed
	resourceTypes := flag.Args()
	if len(resourceTypes) == 0 {
		die()
	}
	resources := resolveResources(resourceTypes)
	b := builder.NewCfnBuilder(bareFlag, iamFlag)
	t, c := b.Template(resources)
	if jsonFlag {
		fmt.Println(format.JsonWithComments(t, c))
	} else {
		fmt.Println(format.YamlWithComments(t, c))
	}
}
