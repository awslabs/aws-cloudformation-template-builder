
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSElastiCacheSecurityGroupIngressResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute(nil), Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html", Properties:map[string]cf.Property{"CacheSecurityGroupName":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "EC2SecurityGroupName":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "EC2SecurityGroupOwnerId":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}}, AdditionalProperties:false}
}


