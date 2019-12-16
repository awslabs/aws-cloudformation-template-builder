package main

//go:generate go run main.go

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/awslabs/aws-cloudformation-template-builder/spec"
)

const (
	cfnSpecUrl = "https://d2senuesg1djtx.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json"
	cfnSpecFn  = "CloudFormationResourceSpecification.json"
	iamSpecFn  = "IamSpecification.json"
)

func load(r io.Reader, s *spec.Spec) {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}
}

func loadUrl(url string) spec.Spec {
	var s spec.Spec

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	load(resp.Body, &s)

	return s
}

func loadFile(fn string) spec.Spec {
	var s spec.Spec

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	load(f, &s)

	return s
}

func main() {
	var cfn, iam spec.Spec

	// Download cfn spec from AWS
	fmt.Println("CloudFormation spec...")
	// Due to a bug in the current published spec, we'll use a local file for now
	cfn = loadUrl(cfnSpecUrl)
	//cfn = loadFile(cfnSpecFn)

	// Load iam spec from local file
	fmt.Println("Parsing IAM spec...")
	iam = loadFile(iamSpecFn)

	// Write out specs file
	err := ioutil.WriteFile("../builder/specs.go", []byte(fmt.Sprintf(`package builder

import "github.com/awslabs/aws-cloudformation-template-builder/spec"

var CfnSpec = %#v;

var IamSpec = %#v;
`, cfn, iam)), 0644)

	if err != nil {
		panic(err)
	}
}
