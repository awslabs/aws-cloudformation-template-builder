
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSCloudFrontStreamingDistributionResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute{"DomainName":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}}, Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html", Properties:map[string]cf.Property{"StreamingDistributionConfig":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:true, Type:"StreamingDistributionConfig", UpdateType:"Mutable"}, "Tags":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-streamingdistribution.html#cfn-cloudfront-streamingdistribution-tags", DuplicatesAllowed:false, ItemType:"Tag", PrimitiveItemType:"", PrimitiveType:"", Required:true, Type:"List", UpdateType:"Mutable"}}, AdditionalProperties:false}
}


