package main

import (
	"codecommit/builders/cfn-format/format"
	"codecommit/builders/cfn-spec-go/builder"
	"fmt"
	"os"
)

func main() {
	config := map[string]string{
		"MyThing": os.Args[1],
	}

	t := builder.Cfn.Template(config)
	fmt.Println(format.Yaml(t))
}
