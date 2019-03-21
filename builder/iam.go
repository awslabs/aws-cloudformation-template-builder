package builder

type IamBuilder struct {
	Builder
}

func NewIamBuilder() IamBuilder {
	var b IamBuilder
	b.Spec = IamSpec

	return b
}

func (b IamBuilder) Policy() (interface{}, map[interface{}]interface{}) {
	return b.newPropertyType("", "Policy")
}
