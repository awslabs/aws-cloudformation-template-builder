package spec

import (
	"strings"
)

func Indent(input, indent string) string {
	return indent + strings.Replace(input, "\n", "\n"+indent, -1)
}

func primitiveSkeleton(t string) Skeleton {
	switch t {
	case "String":
		return Skeleton{"\"CHANGEME\"", false}
	case "Long", "Integer", "Double":
		return Skeleton{"0", false}
	case "Boolean":
		return Skeleton{"true", false}
	case "Timestamp":
		return Skeleton{"1970-01-01 00:00:00", false}
	case "Json":
		return Skeleton{"{}", false}
	default:
		panic("Unimplemented primitive: " + t)
	}
}
