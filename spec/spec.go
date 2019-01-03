package spec

type Spec struct {
	ResourceSpecificationVersion string
	PropertyTypes                map[string]SubProperty
	ResourceTypes                map[string]Resource
}

func (s Spec) GetProperty(p Property, typeName string) SubProperty {
	subTypeName := p.typeName + "." + typeName

	subProp, ok := s.PropertyTypes[subTypeName]

	if !ok {
		panic("No such property type: " + subTypeName)
	}

	subProp.typeName = subTypeName

	return subProp
}

func (s Spec) GetResource(typeName string) Resource {
	res, ok := s.ResourceTypes[typeName]

	if !ok {
		panic("No such resource type: " + typeName)
	}

	res.typeName = typeName

	return res
}
