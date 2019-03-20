package spec

import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

// Cfn is a representation of the CloudFormation specification
var Cfn cf.Spec

// Iam is a representation fo the Iam specification
var Iam cf.Spec

func init() {
	iamInit()
	cfnInit()
	// Create Cfn
	Cfn.PropertyTypes = cfnProperties
	Cfn.ResourceTypes = cfnResources
	Cfn.ResourceSpecificationVersion = cfnResourceSpecificationVersion

	// Create Iam
	Iam.PropertyTypes = iamProperties
	Iam.ResourceTypes = iamResources
	Iam.ResourceSpecificationVersion = iamResourceSpecificationVersion
}
