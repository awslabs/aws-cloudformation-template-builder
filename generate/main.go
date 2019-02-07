// Package generate is used to generate resource and property functions from the CloudFormation specification.
// Each function is registered with it's respective map, resourceFuncs or propertyFuncs.
//
package main

import (
	"codecommit/builders/cfn-skeleton/spec"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// PropertyFunc is a function that returns a PropertyType
// of the named function
type PropertyFunc func() spec.PropertyType

// ResourceFunc is a function that returns a ResourceType
// of the named function
type ResourceFunc func() spec.ResourceType

// map resouce names their respective ResourceFunc
// and PropertyFunc
var (
	resourceFuncs map[string]ResourceFunc
	propertyFuncs map[string]PropertyFunc
)

const (
	resourceTemplate = "resource.tmpl"
	propertyTemplate = "property.tmpl"
)

func main() {
	// Make file of example resource
	// Needs to be fleshed out to iterate over all resources and properties in the spec.
	// Each resource and property needs to be registered with it's respective map
	// either resourceFuncs or propertyFuncs.
	resources := spec.Cfn.ResourceTypes
	//properties := spec.Cfn.PropertyTypes

	// Make resourceFuncs

	for name, _ := range resources {
		generateResource(name)

	}

	// Make propertyFuncs
	// This won't work
	// propertyFunc not yet implemented.

	// for name, _ := range properties {
	// 	generateProperty(name)
	// }

}

func generateResource(name string) {

	// get the resource/property type
	resource, err := getResourceType(name)
	if err != nil {
		panic(err)
	}

	// Write the file
	err = writeResourceFile(name, resource)
	if err != nil {
		panic(err)
	}
}

// getResourceType returns a ResourceType for a given name. If it
// cannot find the type, it will return an error and an empty
// spec.ResourceType
func getResourceType(name string) (spec.ResourceType, error) {
	resource, ok := spec.Cfn.ResourceTypes[name]
	if !ok {
		return spec.ResourceType{}, errors.New("Cannot resolve resource name: " + name)
	}
	return resource, nil
}

// writeResourceFile takes a qualified resource name
// (eg AWS::S3::Bucket) and
// generates a file containin ResourceFunc and
// PropertyFunc
func writeResourceFile(name string, rt spec.ResourceType) error {

	cleanedName := nameFromAWSType(name)
	out, err := os.Create(cleanedName + ".go")
	defer out.Close()
	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles(resourceTemplate))
	err = t.ExecuteTemplate(out, "template", struct {
		Name         string
		RTDefinition interface{}
	}{
		Name: cleanedName,

		// Use %#v directive to output the entire struct in parseable format
		RTDefinition: fmt.Sprintf("%#v", rt),
	})

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
