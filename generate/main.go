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

var exampleResource = "AWS::S3::Bucket"

func main() {
	// Make file of example resource
	// Needs to be fleshed out to iterate over all resources and properties in the spec.
	// Each resource and property needs to be registered with it's respective map
	// either resourceFuncs or propertyFuncs.
	err := makeFile(exampleResource)
	if err != nil {
		panic(err)
	}

}

// makeFile takes a qualified resource name
// (eg AWS::S3::Bucket) and
// generates a file containin ResourceFunc and
// PropertyFunc
func makeFile(name string) error {
	// Look for the resource

	resource, ok := spec.Cfn.ResourceTypes[name]
	if !ok {
		return errors.New("Cannot resolve resource name: " + name)
	}
	cleanedName := cleanName(name)
	out, err := os.Create(cleanedName + ".go")
	defer out.Close()
	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles("template.txt"))
	err = t.ExecuteTemplate(out, "template", struct {
		Name         string
		RTDefinition interface{}
		PTDefinition interface{}
	}{
		Name:         cleanedName,
		RTDefinition: fmt.Sprintf("%#v", resource),
		PTDefinition: fmt.Sprintf("\"Not yet implemented\""),
	})

	if err != nil {
		return err
	}

	return nil

}

// cleanName takes a qualified name such as
// AWS::S3::Bucket and removes all :: delimiters
func cleanName(name string) string {
	return strings.Replace(name, "::", "", -1)
}
