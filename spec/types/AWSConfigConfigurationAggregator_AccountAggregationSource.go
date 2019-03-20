
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSConfigConfigurationAggregator_AccountAggregationSourceProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html", Properties:map[string]cf.Property{"AllAwsRegions":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-allawsregions", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "AwsRegions":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-awsregions", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "AccountIds":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationaggregator-accountaggregationsource.html#cfn-config-configurationaggregator-accountaggregationsource-accountids", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:true, Type:"List", UpdateType:"Mutable"}}}
}


