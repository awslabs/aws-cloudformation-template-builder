
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSS3Bucket_NotificationConfigurationProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html", Properties:map[string]cf.Property{"LambdaConfigurations":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-lambdaconfig", DuplicatesAllowed:false, ItemType:"LambdaConfiguration", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "QueueConfigurations":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-queueconfig", DuplicatesAllowed:false, ItemType:"QueueConfiguration", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "TopicConfigurations":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfig.html#cfn-s3-bucket-notificationconfig-topicconfig", DuplicatesAllowed:false, ItemType:"TopicConfiguration", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}}}
}


