
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSSQSQueuePolicyResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute(nil), Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html", Properties:map[string]cf.Property{"PolicyDocument":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-policydoc", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Json", Required:true, Type:"", UpdateType:"Mutable"}, "Queues":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html#cfn-sqs-queuepolicy-queues", DuplicatesAllowed:true, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:true, Type:"List", UpdateType:"Mutable"}}, AdditionalProperties:false}
}


