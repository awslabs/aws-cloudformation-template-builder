
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSEC2SpotFleet_LoadBalancersConfigProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-loadbalancersconfig.html", Properties:map[string]cf.Property{"ClassicLoadBalancersConfig":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-loadbalancersconfig.html#cfn-ec2-spotfleet-loadbalancersconfig-classicloadbalancersconfig", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"ClassicLoadBalancersConfig", UpdateType:"Mutable"}, "TargetGroupsConfig":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-loadbalancersconfig.html#cfn-ec2-spotfleet-loadbalancersconfig-targetgroupsconfig", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"TargetGroupsConfig", UpdateType:"Mutable"}}}
}


