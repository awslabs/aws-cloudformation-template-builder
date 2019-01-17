package cmd

import (
	"codecommit/builders/cfn-spec-go/menu"
	"encoding/json"
	"fmt"
)

func Interactive(args []string) {
	out, err := menu.Cfn.Build("AWS::S3::Bucket")
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
