package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"codecommit/builders/cfn-spec-go/spec"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <resource type>\n", os.Args[0])
		os.Exit(1)
	}

	suffix := os.Args[1]

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

	name := "My" + strings.Split(options[0], "::")[0]

	b := builder.NewCfnBuilder(true)
	t := b.Template(map[string]string{
		name: options[0],
	})
	fmt.Println(format.Yaml(t))
}
