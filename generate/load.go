package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/awslabs/aws-cloudformation-template-builder/spec/cf"
)

const cfnSpecFn = "CloudFormationResourceSpecification.json"
const iamSpecFn = "IamSpecification.json"

// Cfn is a representation of the CloudFormation specification
var Cfn cf.Spec

// Iam is a representation fo the Iam specification
var Iam cf.Spec

func load(fn string, s *cf.Spec) {
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
