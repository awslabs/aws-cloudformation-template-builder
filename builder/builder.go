package builder

import (
	"github.com/awslabs/aws-cloudformation-template-builder/spec"
	"github.com/awslabs/aws-cloudformation-template-builder/spec/cf"
)

const (
	PolicyDocument           = "PolicyDocument"
	AssumeRolePolicyDocument = "AssumeRolePolicyDocument"
	OptionalTag              = "Optional"
	ChangeMeTag              = "CHANGEME"
)

type Builder struct {
	Spec                      cf.Spec
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
			if b.BuildIamPolicies && (name == PolicyDocument || name == AssumeRolePolicyDocument) {
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

func (b Builder) newProperty(resourceType string, pSpec cf.Property) (interface{}, map[interface{}]interface{}) {
	// Correct badly-formed entries
	if pSpec.PrimitiveType == spec.TypeMap {
		pSpec.PrimitiveType = spec.TypeEmpty
		pSpec.Type = spec.TypeMap
	}

	comments := make(map[interface{}]interface{})
	if !pSpec.Required {
		comments[""] = OptionalTag
	}

	// Primitive types
	if pSpec.PrimitiveType != spec.TypeEmpty {
		return b.newPrimitive(pSpec.PrimitiveType), comments
	}

	if pSpec.Type == spec.TypeList || pSpec.Type == spec.TypeMap {
		var value interface{}
		var subComments map[interface{}]interface{}

		if pSpec.PrimitiveItemType != spec.TypeEmpty {
			value = b.newPrimitive(pSpec.PrimitiveItemType)
		} else if pSpec.ItemType != spec.TypeEmpty {
			value, subComments = b.newPropertyType(resourceType, pSpec.ItemType)
		} else {
			value = ChangeMeTag
		}

		if pSpec.Type == spec.TypeList {
			if subComments != nil {
				comments[0] = subComments
			}

			return []interface{}{value}, comments
		}

		if subComments != nil {
			comments[ChangeMeTag] = subComments
		}

		return map[string]interface{}{ChangeMeTag: value}, comments
	}

	// Fall through to property types
	return b.newPropertyType(resourceType, pSpec.Type)
}

func (b Builder) newPrimitive(primitiveType string) interface{} {
	switch primitiveType {
	case "String":
		return ChangeMeTag
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
	var ptSpec cf.PropertyType
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
			comments[name] = OptionalTag
		}

		if b.BuildIamPolicies && (name == PolicyDocument || name == AssumeRolePolicyDocument) {
			properties[name], comments[name] = iamBuilder.Policy()
		} else if pSpec.Type == propertyType || pSpec.ItemType == propertyType {
			properties[name] = make(map[string]interface{})
		} else {
			properties[name], _ = b.newProperty(resourceType, pSpec)
		}
	}

	return properties, comments
}
