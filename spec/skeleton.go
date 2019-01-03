package spec

import "strings"

type Skeleton struct {
	Value  string
	Indent bool
}

func (s Skeleton) String() string {
	builder := strings.Builder{}

	if s.Indent {
		builder.WriteString(Indent(s.Value, "  "))
	} else {
		builder.WriteString(s.Value)
	}

	return builder.String()
}

func (s Skeleton) Comment() string {
	if s.Indent {
		return Indent(s.String(), "# ")
	}

	return s.String()
}
