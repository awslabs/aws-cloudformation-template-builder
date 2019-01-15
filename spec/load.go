package spec

import (
	"encoding/json"
	"fmt"
	"os"
)

const cfnSpecFn = "CloudFormationResourceSpecification.json"
const iamSpecFn = "IamSpecification.json"

var Cfn Spec
var Iam Spec

func load(fn string, s *Spec) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err.Error())
	}

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&s)
	if err != nil {
		fmt.Printf("%T\n", err)
		panic(err.Error())
	}
}

func init() {
	load(cfnSpecFn, &Cfn)
	load(iamSpecFn, &Iam)
}
