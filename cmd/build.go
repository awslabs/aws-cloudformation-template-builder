package cmd

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"codecommit/builders/cfn-spec-go/spec"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Build(args []string) {
	resources := make(map[string]string)

	for _, suffix := range args {
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

		resources[name] = options[0]
	}

	b := builder.NewCfnBuilder(true)
	t, c := b.Template(resources)
	fmt.Println(format.YamlWithComments(t, c))
}
