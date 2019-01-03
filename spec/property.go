package spec

type Property struct {
	typeName          string
	Documentation     string
	DuplicatesAllowed bool
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Required          bool
	Type              string
	UpdateType        string
}

func (p Property) listSkeleton(cfnSpec Spec) Skeleton {
	switch {
	case p.PrimitiveItemType != "":
		return Skeleton{
			Value:  "- " + primitiveSkeleton(p.PrimitiveItemType).String(),
			Indent: true,
		}
	case p.ItemType != "":
		subProp := cfnSpec.GetProperty(p, p.ItemType)

		return Skeleton{
			Value:  "- " + subProp.Skeleton(cfnSpec).String(),
			Indent: true,
		}
	default:
		panic("Unimplemented list type")
	}
}

func (p Property) mapSkeleton(cfnSpec Spec) Skeleton {
	switch {
	case p.PrimitiveItemType != "":
		return Skeleton{
			Value:  "KEY: " + primitiveSkeleton(p.PrimitiveItemType).String(),
			Indent: true,
		}
	case p.ItemType != "":
		return Skeleton{
			Value:  "KEY: " + cfnSpec.PropertyTypes[p.Type].Skeleton(cfnSpec).String(),
			Indent: true,
		}
	default:
		panic("Unimplemented list type")
	}
}

func (p Property) Skeleton(cfnSpec Spec) Skeleton {
	switch {
	case p.PrimitiveType != "":
		return primitiveSkeleton(p.PrimitiveType)
	case p.Type != "":
		switch p.Type {
		case "List":
			return p.listSkeleton(cfnSpec)
		case "Map":
			return p.mapSkeleton(cfnSpec)
		default:
			return cfnSpec.PropertyTypes[p.Type].Skeleton(cfnSpec)
		}
	default:
		panic("Unimplemented")
	}
}
