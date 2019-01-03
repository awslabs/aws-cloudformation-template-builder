package spec

import (
	"sort"
	"strings"
)

type Resource struct {
	typeName             string
	Attributes           map[string]Attribute
	Documentation        string
	Properties           map[string]Property
	AdditionalProperties bool
}

func (r Resource) Skeleton(cfnSpec Spec) Skeleton {
	parts := make(map[string]Skeleton, len(r.Properties))
	requiredNames := make([]string, 0)
	optionalNames := make([]string, 0)

	for name, prop := range r.Properties {
		prop.typeName = r.typeName

		parts[name] = prop.Skeleton(cfnSpec)

		if prop.Required {
			requiredNames = append(requiredNames, name)
		} else {
			optionalNames = append(optionalNames, name)
		}
	}

	sort.Strings(requiredNames)
	sort.Strings(optionalNames)

	builder := strings.Builder{}

	for _, name := range requiredNames {
		builder.WriteString(name)
		builder.WriteString(":")

		if parts[name].Indent {
			builder.WriteString("\n")
		} else {
			builder.WriteString(" ")
		}

		builder.WriteString(parts[name].String())

		builder.WriteString("\n")
	}

	for _, name := range optionalNames {
		builder.WriteString("# ")
		builder.WriteString(name)
		builder.WriteString(":")

		if parts[name].Indent {
			builder.WriteString("\n")
		} else {
			builder.WriteString(" ")
		}

		builder.WriteString(parts[name].Comment())

		builder.WriteString("\n")
	}

	return Skeleton{
		Value:  builder.String(),
		Indent: true,
	}
}
