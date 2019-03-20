
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSS3Bucket_RuleProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html", Properties:map[string]cf.Property{"ExpirationInDays":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationindays", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}, "NoncurrentVersionExpirationInDays":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversionexpirationindays", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}, "Status":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-status", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "Transition":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transition", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"Transition", UpdateType:"Mutable"}, "NoncurrentVersionTransitions":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransitions", DuplicatesAllowed:false, ItemType:"NoncurrentVersionTransition", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "Prefix":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-prefix", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "TagFilters":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-rule-tagfilters", DuplicatesAllowed:false, ItemType:"TagFilter", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "Transitions":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-transitions", DuplicatesAllowed:false, ItemType:"Transition", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "AbortIncompleteMultipartUpload":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-rule-abortincompletemultipartupload", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"AbortIncompleteMultipartUpload", UpdateType:"Mutable"}, "ExpirationDate":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-expirationdate", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Timestamp", Required:false, Type:"", UpdateType:"Mutable"}, "Id":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-id", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "NoncurrentVersionTransition":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig-rule.html#cfn-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"NoncurrentVersionTransition", UpdateType:"Mutable"}}}
}


