
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSBatchJobDefinition_ContainerPropertiesProperty() cf.PropertyType {
    return cf.PropertyType{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html", Properties:map[string]cf.Property{"User":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-user", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "Memory":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-memory", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:true, Type:"", UpdateType:"Mutable"}, "Privileged":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-privileged", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "JobRoleArn":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-jobrolearn", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "Image":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-image", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "Volumes":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-volumes", DuplicatesAllowed:false, ItemType:"Volumes", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "Command":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-command", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "Environment":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-environment", DuplicatesAllowed:false, ItemType:"Environment", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "InstanceType":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-instancetype", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "ReadonlyRootFilesystem":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-readonlyrootfilesystem", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "Vcpus":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-vcpus", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:true, Type:"", UpdateType:"Mutable"}, "MountPoints":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-mountpoints", DuplicatesAllowed:false, ItemType:"MountPoints", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "Ulimits":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-ulimits", DuplicatesAllowed:false, ItemType:"Ulimit", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}}}
}


