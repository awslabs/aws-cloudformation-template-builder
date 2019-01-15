package builder

import "codecommit/builders/cfn-spec-go/spec"

type iamBuilder struct {
	Builder
}

var Iam = iamBuilder{}

func init() {
	Iam.Spec = spec.Iam
}

func (b iamBuilder) Policy() interface{} {
	return b.newPropertyType("", "Policy")
}
