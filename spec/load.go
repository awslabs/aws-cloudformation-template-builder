package spec

import (
	"encoding/json"
	"fmt"
	"os"
)

const cfnSpecFn = "CloudFormationResourceSpecification.json"

var Cfn Spec

func init() {
	f, err := os.Open(cfnSpecFn)
	if err != nil {
		panic(err.Error())
	}

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&Cfn)
	if err != nil {
		fmt.Printf("%T\n", err)
		panic(err.Error())
	}
}
