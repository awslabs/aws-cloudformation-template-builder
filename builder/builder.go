package builder

import (
	"codecommit/builders/cfn-spec-go/spec"
)

type Builder struct {
	Spec                      spec.Spec
	IncludeOptionalProperties bool
}

func (b Builder) newResource(resourceType string) map[string]interface{} {
	rSpec, ok := b.Spec.ResourceTypes[resourceType]
	if !ok {
		panic("No such resource type: " + resourceType)
	}

	// Generate properties
	properties := make(map[string]interface{})
	for name, pSpec := range rSpec.Properties {
		if b.IncludeOptionalProperties || pSpec.Required {
			properties[name] = b.newProperty(resourceType, pSpec)
		}
	}

	return map[string]interface{}{
		"Type":       resourceType,
		"Properties": properties,
	}
}

func (b Builder) newProperty(resourceType string, pSpec spec.Property) interface{} {
	// Correctly badly-formed entries
	if pSpec.PrimitiveType == "Map" {
		pSpec.PrimitiveType = ""
		pSpec.Type = "Map"
	}

	// Primitive types
	if pSpec.PrimitiveType != "" {
		return b.newPrimitive(pSpec.PrimitiveType)
	}

	if pSpec.Type == "List" || pSpec.Type == "Map" {
		var value interface{}

		if pSpec.PrimitiveItemType != "" {
			value = b.newPrimitive(pSpec.PrimitiveItemType)
		} else if pSpec.ItemType != "" {
			value = b.newPropertyType(resourceType, pSpec.ItemType)
		} else {
			value = "CHANGEME"
		}

		if pSpec.Type == "List" {
			return []interface{}{value}
		}

		return map[string]interface{}{"CHANGEME": value}
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

func (b Builder) newPropertyType(resourceType, propertyType string) interface{} {
	var ptSpec spec.PropertyType
	var ok bool

	ptSpec, ok = b.Spec.PropertyTypes[propertyType]
	if !ok {
		ptSpec, ok = b.Spec.PropertyTypes[resourceType+"."+propertyType]
	}
	if !ok {
		panic("PTYPE NOT IMPLEMENTED: " + resourceType + "." + propertyType)
	}

	// Generate properties
	properties := make(map[string]interface{})
	for name, pSpec := range ptSpec.Properties {
		if pSpec.Type == propertyType || pSpec.ItemType == propertyType {
			properties[name] = make(map[string]interface{})
		} else {
			properties[name] = b.newProperty(resourceType, pSpec)
		}
	}

	return properties
}
