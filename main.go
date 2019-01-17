package main

import (
	"codecommit/builders/cfn-spec-go/cmd"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || (os.Args[1] == "-i" && len(os.Args) < 3) {
		fmt.Fprintf(os.Stderr, "Usage: %s [-i] <resource type>\n", os.Args[0])
		os.Exit(1)
	}

	if os.Args[1] == "-i" {
		cmd.Interactive(os.Args[2:])
	} else {
		cmd.Build(os.Args[1:])
	}
}
