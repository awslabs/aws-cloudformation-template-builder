// Package generate is used to generate resource and property functions from the CloudFormation specification.
// Each function is registered with it's respective map, resourceFuncs or propertyFuncs.
//
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/awslabs/aws-cloudformation-template-builder/spec"
	"github.com/awslabs/aws-cloudformation-template-builder/spec/cf"
)

// Filename of the templates used
const (
	resourceTemplate = "generate/resource.tmpl"
	propertyTemplate = "generate/property.tmpl"
	initTemplate     = "generate/init.tmpl"
)

type funcs struct {
	Name string
	Fun  string
}

func main() {

	resources := spec.Cfn.ResourceTypes
	properties := spec.Cfn.PropertyTypes

	resourceFuncs := []funcs{}
	// Make resourceFuncs
	for rt := range resources {
		fun := generateResource(rt)
		resourceFuncs = append(resourceFuncs, funcs{rt, fun})

	}

	propertyFuncs := []funcs{}
	// Make propertyFuncs
	for pt := range properties {
		fun := generateProperty(pt)
		propertyFuncs = append(propertyFuncs, funcs{pt, fun})
	}

	// Make init() func
	generateInit(resourceFuncs, propertyFuncs)
}

// generateInit creates the init() function
// that registers each resource and property
// in a containing map
func generateInit(resources, properties []funcs) {
	var b bytes.Buffer
	t := template.Must(template.ParseFiles(initTemplate))
	err := t.ExecuteTemplate(&b, "template", struct {
		Resources  []funcs // Name of the function to be created
		Properties []funcs // Body of the function
	}{
		Resources:  resources,
		Properties: properties,
	})

	if err != nil {
		panic(err)
	}
	writeFile("init", b.Bytes())
}

// generateResource creates a file containing a
// function that returns a resourceType
func generateResource(resourceType string) string {

	name := nameFromAWSType(resourceType)
	// get the resource/property type
	resource, err := getResourceType(resourceType)
	if err != nil {
		panic(err)
	}

	b, err := build(name, resourceTemplate, resource)
	if err != nil {
		panic(err)
	}

	// Write the file
	err = writeFile("types/"+name, b)
	if err != nil {
		panic(err)
	}

	return name
}

// generateProperty creates a file containing a
// function that returns a propertyType
func generateProperty(propertyType string) string {
	name := nameFromAWSType(propertyType)
	property, err := getPropertyType(propertyType)
	if err != nil {
		panic(err)
	}

	b, err := build(name, propertyTemplate, property)
	if err != nil {
		panic(err)
	}

	err = writeFile("types/"+name, b)
	if err != nil {
		panic(err)
	}

	return name
}

// build takes a name, templateName and a resourceType or propertyType
// it executes the template into a buffer and returns the array of bytes
func build(name string, templateName string, input interface{}) ([]byte, error) {
	var b bytes.Buffer
	t := template.Must(template.ParseFiles(templateName))
	err := t.ExecuteTemplate(&b, "template", struct {
		Name      string      // Name of the function to be created
		ReturnVal interface{} // Body of the function
	}{
		Name: name,

		// Use %#v directive to output the entire struct in parseable format
		ReturnVal: fmt.Sprintf("%#v", input),
	})

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// writeFile writes the provided byte sequence to a file of the provided name
func writeFile(name string, b []byte) error {
	out, err := os.Create("spec/" + name + ".go")
	defer out.Close()
	if err != nil {
		return err
	}

	_, err = out.Write(b)
	if err != nil {
		return err
	}

	return nil

}

// nameFromAWSType takes an CloudFormation type such as
// AWS::S3::Bucket and removes all :: delimiters to
// produce a name used for assignments
func nameFromAWSType(name string) string {
	temp := strings.Replace(name, "::", "", -1)
	temp = strings.Replace(temp, ".", "_", -1)
	return temp
}

// getResourceType returns a ResourceType for a given name. If it
// cannot find the type, it will return an error and an empty
// spec.ResourceType
func getResourceType(name string) (cf.ResourceType, error) {
	resource, ok := spec.Cfn.ResourceTypes[name]
	if !ok {
		return cf.ResourceType{}, errors.New("Cannot resolve resource name: " + name)
	}
	return resource, nil
}

// getPropertyType returns a PropertyType for a given name.
// If it cannot find the type, it will return an error and an empty
// spec.PropertyType
func getPropertyType(name string) (cf.PropertyType, error) {
	property, ok := spec.Cfn.PropertyTypes[name]
	if !ok {
		return cf.PropertyType{}, errors.New("Cannot resolve property name: " + name)
	}
	return property, nil
}
