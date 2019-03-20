
// Package types contains functions that return a resource or property type when called.
// This code is autogenerated. 
// Do not edit it by hand.
package types
import "github.com/awslabs/aws-cloudformation-template-builder/spec/cf"

 
func AWSDMSReplicationInstanceResource() cf.ResourceType {
    return cf.ResourceType{Attributes:map[string]cf.Attribute{"ReplicationInstancePublicIpAddresses":cf.Attribute{ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Type:"List"}, "ReplicationInstancePrivateIpAddresses":cf.Attribute{ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Type:"List"}}, Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html", Properties:map[string]cf.Property{"KmsKeyId":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-kmskeyid", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Immutable"}, "VpcSecurityGroupIds":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-vpcsecuritygroupids", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"String", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Mutable"}, "AllowMajorVersionUpgrade":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-allowmajorversionupgrade", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "PubliclyAccessible":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-publiclyaccessible", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Immutable"}, "MultiAZ":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-multiaz", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "ReplicationInstanceIdentifier":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-replicationinstanceidentifier", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "AvailabilityZone":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-availabilityzone", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Immutable"}, "EngineVersion":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-engineversion", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "PreferredMaintenanceWindow":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-preferredmaintenancewindow", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Mutable"}, "AutoMinorVersionUpgrade":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-autominorversionupgrade", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Boolean", Required:false, Type:"", UpdateType:"Mutable"}, "ReplicationSubnetGroupIdentifier":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-replicationsubnetgroupidentifier", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:false, Type:"", UpdateType:"Immutable"}, "ReplicationInstanceClass":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-replicationinstanceclass", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"String", Required:true, Type:"", UpdateType:"Mutable"}, "Tags":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-tags", DuplicatesAllowed:false, ItemType:"Tag", PrimitiveItemType:"", PrimitiveType:"", Required:false, Type:"List", UpdateType:"Immutable"}, "AllocatedStorage":cf.Property{Documentation:"http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-allocatedstorage", DuplicatesAllowed:false, ItemType:"", PrimitiveItemType:"", PrimitiveType:"Integer", Required:false, Type:"", UpdateType:"Mutable"}}, AdditionalProperties:false}
}


