package builder

import (
	"codecommit/builders/cfn-skeleton/spec"
)

type Builder struct {
	Spec                      spec.Spec
	IncludeOptionalProperties bool
	BuildIamPolicies          bool
}

var iamBuilder IamBuilder

func init() {
	iamBuilder = NewIamBuilder()
}

func (b Builder) newResource(resourceType string) (map[string]interface{}, map[interface{}]interface{}) {
	rSpec, ok := b.Spec.ResourceTypes[resourceType]
	if !ok {
		panic("No such resource type: " + resourceType)
	}

	// Generate properties
	properties := make(map[string]interface{})
	comments := make(map[interface{}]interface{})
	for name, pSpec := range rSpec.Properties {
		if b.IncludeOptionalProperties || pSpec.Required {
			if b.BuildIamPolicies && (name == "PolicyDocument" || name == "AssumeRolePolicyDocument") {
				properties[name], comments[name] = iamBuilder.Policy()
			} else {
				properties[name], comments[name] = b.newProperty(resourceType, pSpec)
			}
		}
	}

	return map[string]interface{}{
			"Type":       resourceType,
			"Properties": properties,
		}, map[interface{}]interface{}{
			"Properties": comments,
		}
}

func (b Builder) newProperty(resourceType string, pSpec spec.Property) (interface{}, map[interface{}]interface{}) {
	// Correct badly-formed entries
	if pSpec.PrimitiveType == "Map" {
		pSpec.PrimitiveType = ""
		pSpec.Type = "Map"
	}

	comments := make(map[interface{}]interface{})
	if !pSpec.Required {
		comments[""] = "Optional"
	}

	// Primitive types
	if pSpec.PrimitiveType != "" {
		return b.newPrimitive(pSpec.PrimitiveType), comments
	}

	if pSpec.Type == "List" || pSpec.Type == "Map" {
		var value interface{}
		var subComments map[interface{}]interface{}

		if pSpec.PrimitiveItemType != "" {
			value = b.newPrimitive(pSpec.PrimitiveItemType)
		} else if pSpec.ItemType != "" {
			value, subComments = b.newPropertyType(resourceType, pSpec.ItemType)
		} else {
			value = "CHANGEME"
		}

		if pSpec.Type == "List" {
			if subComments != nil {
				comments[0] = subComments
			}

			return []interface{}{value}, comments
		}

		if subComments != nil {
			comments["CHANGEME"] = subComments
		}

		return map[string]interface{}{"CHANGEME": value}, comments
	}

	// Fall through to property types
	return b.newPropertyType(resourceType, pSpec.Type)
}

func (b Builder) newPrimitive(primitiveType string) interface{} {
	switch primitiveType {
	case "String":
		return "CHANGEME"
	case "Integer":
		return 0
	case "Double":
		return 0.0
	case "Long":
		return 0.0
	case "Boolean":
		return false
	case "Timestamp":
		return "1970-01-01 00:00:00"
	case "Json":
		return "{\"JSON\": \"CHANGEME\"}"
	default:
		panic("PRIMITIVE NOT IMPLEMENTED: " + primitiveType)
	}
}

func (b Builder) newPropertyType(resourceType, propertyType string) (interface{}, map[interface{}]interface{}) {
	var ptSpec spec.PropertyType
	var ok bool

	ptSpec, ok = b.Spec.PropertyTypes[propertyType]
	if !ok {
		ptSpec, ok = b.Spec.PropertyTypes[resourceType+"."+propertyType]
	}
	if !ok {
		panic("PTYPE NOT IMPLEMENTED: " + resourceType + "." + propertyType)
	}

	comments := make(map[interface{}]interface{})

	// Generate properties
	properties := make(map[string]interface{})
	for name, pSpec := range ptSpec.Properties {
		if !pSpec.Required {
			comments[name] = "Optional"
		}

		if b.BuildIamPolicies && (name == "PolicyDocument" || name == "AssumeRolePolicyDocument") {
			properties[name], comments[name] = iamBuilder.Policy()
		} else if pSpec.Type == propertyType || pSpec.ItemType == propertyType {
			properties[name] = make(map[string]interface{})
		} else {
			properties[name], _ = b.newProperty(resourceType, pSpec)
		}
	}

	return properties, comments
}
