package spec

type Spec struct {
	ResourceSpecificationVersion string
	PropertyTypes                map[string]PropertyType
	ResourceTypes                map[string]ResourceType
}

type PropertyType struct {
	Documentation string
	Properties    map[string]Property
}

type ResourceType struct {
	Attributes           map[string]Attribute
	Documentation        string
	Properties           map[string]Property
	AdditionalProperties bool
}

type Property struct {
	Documentation     string
	DuplicatesAllowed bool
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Required          bool
	Type              string
	UpdateType        string
}

type Attribute struct {
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Type              string
}

func (p Property) TypeName() string {
	if p.PrimitiveType != "" {
		if p.PrimitiveType == "List" || p.PrimitiveType == "Map" {
			if p.PrimitiveItemType != "" {
				return p.PrimitiveType + "/" + p.PrimitiveItemType
			}

			return p.PrimitiveType + "/" + p.ItemType
		}

		return p.PrimitiveType
	}

	return p.Type
}
