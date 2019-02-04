package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-skeleton/builder"
	"codecommit/builders/cfn-skeleton/spec"
	"fmt"
	"os"
	"sort"
	"strings"
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

func resolveType(suffix string) string {
	options := spec.Cfn.ResolveResource(suffix)

	if len(options) == 0 {
		fmt.Fprintf(os.Stderr, "No resource type found matching '%s'\n", suffix)
		os.Exit(1)
	} else if len(options) != 1 {
		fmt.Fprintf(os.Stderr, "Ambiguous resource type '%s'; could be any of:\n", suffix)
		sort.Strings(options)
		for _, option := range options {
			fmt.Fprintf(os.Stderr, "  %s\n", option)
		}
		os.Exit(1)
	}

	return options[0]
}

func makeName(resourceType string) string {
	parts := strings.Split(resourceType, "::")
	return "My" + parts[len(parts)-1]
}

func main() {
	if len(os.Args) < 2 {
		die()
	}

	includeOptional := true
	buildIamPolicies := false
	style := styleYAML
	resourceTypes := make([]string, 0)

	// Parse params
	for _, arg := range os.Args[1:] {
		if arg == "-b" || arg == "--bare" {
			includeOptional = false
		} else if arg == "-i" || arg == "--iam" {
			buildIamPolicies = true
		} else if arg == "-j" || arg == "--json" {
			style = styleJSON
		} else if arg == "--help" {
			die()
		} else if arg[0] == '-' {
			die()
		} else {
			resourceTypes = append(resourceTypes, arg)
		}
	}

	// Refuse to build an empty template?
	if len(resourceTypes) == 0 {
		die()
	}

	// Resolve resource types
	resources := make(map[string]string)
	for _, r := range resourceTypes {
		r = resolveType(r)
		resources[makeName(r)] = r
	}

	// Generate the template
	b := builder.NewCfnBuilder(includeOptional, buildIamPolicies)
	t, c := b.Template(resources)

	if style == styleJSON {
		fmt.Println(format.JsonWithComments(t, c))
	} else {
		fmt.Println(format.YamlWithComments(t, c))
	}
}
