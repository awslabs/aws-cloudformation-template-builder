package spec

import (
	"sort"
	"strings"
)

type SubProperty struct {
	typeName      string
	Documentation string
	Properties    map[string]Property
}

func (s SubProperty) Skeleton(cfnSpec Spec) Skeleton {
	parts := make(map[string]Skeleton, len(s.Properties))
	names := make([]string, 0)

	for name, prop := range s.Properties {
		prop.typeName = s.typeName

		parts[name] = prop.Skeleton(cfnSpec)

		names = append(names, name)
	}

	sort.Strings(names)

	builder := strings.Builder{}

	for _, name := range names {
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

	return Skeleton{
		Value:  builder.String(),
		Indent: true,
	}
}
