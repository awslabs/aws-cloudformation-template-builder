package builder

import "codecommit/builders/cfn-spec-go/spec"

type IamBuilder struct {
	Builder
}

func NewIamBuilder() IamBuilder {
	var b IamBuilder
	b.Spec = spec.Iam

	return b
}

func (b IamBuilder) Policy() interface{} {
	return b.newPropertyType("", "Policy")
}
