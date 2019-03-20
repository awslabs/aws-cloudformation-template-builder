
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSIoT1ClickDeviceResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute{"DeviceId":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}, "Enabled":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Type:""}, "Arn":cf.Attribute{ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Type:""}}, Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html", Properties:map[string]cf.Property{"Enabled":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-enabled", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:true, Type:"", UpdateType:"Mutable"}, "DeviceId":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-device.html#cfn-iot1click-device-deviceid", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Immutable"}}, AdditionalProperties:false}
}


