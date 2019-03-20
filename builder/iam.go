package builder

import "github.com/awslabs/aws-cloudformation-template-builder/spec"

type IamBuilder struct {
	Builder
}

func NewIamBuilder() IamBuilder {
	var b IamBuilder
	b.Spec = spec.Iam

	return b
}

func (b IamBuilder) Policy() (interface{}, map[interface{}]interface{}) {
	return b.newPropertyType("", "Policy")
}
