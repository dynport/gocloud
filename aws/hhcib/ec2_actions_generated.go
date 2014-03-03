package main

import "time"

type AllocateAddress struct {
	Domain string `aws:"Domain"`
}

type AssignPrivateIpAddresses struct {
	NetworkInterfaceId             string                                    `aws:"NetworkInterfaceId"`
	PrivateIpAddress               []*AssignPrivateIpAddressesSetItemRequest `aws:"PrivateIpAddress"`
	SecondaryPrivateIpAddressCount *IntValue                                 `aws:"SecondaryPrivateIpAddressCount"`
	AllowReassignment              *BoolValue                                `aws:"AllowReassignment"`
}

type AssociateAddress struct {
	PublicIp           string     `aws:"PublicIp"`
	InstanceId         string     `aws:"InstanceId"`
	AllocationId       string     `aws:"AllocationId"`
	NetworkInterfaceId string     `aws:"NetworkInterfaceId"`
	PrivateIpAddress   string     `aws:"PrivateIpAddress"`
	AllowReassociation *BoolValue `aws:"AllowReassociation"`
}

type AssociateDhcpOptions struct {
	DhcpOptionsId string `aws:"DhcpOptionsId"`
	VpcId         string `aws:"VpcId"`
}

type AssociateRouteTable struct {
	RouteTableId string `aws:"RouteTableId"`
	SubnetId     string `aws:"SubnetId"`
}

type AttachInternetGateway struct {
	InternetGatewayId string `aws:"InternetGatewayId"`
	VpcId             string `aws:"VpcId"`
}

type AttachNetworkInterface struct {
	NetworkInterfaceId string    `aws:"NetworkInterfaceId"`
	InstanceId         string    `aws:"InstanceId"`
	DeviceIndex        *IntValue `aws:"DeviceIndex"`
}

type AttachVolume struct {
	VolumeId   string `aws:"VolumeId"`
	InstanceId string `aws:"InstanceId"`
	Device     string `aws:"Device"`
}

type AttachVpnGateway struct {
	VpnGatewayId string `aws:"VpnGatewayId"`
	VpcId        string `aws:"VpcId"`
}

type AuthorizeSecurityGroupEgress struct {
	GroupId       string           `aws:"GroupId"`
	IpPermissions []*IpPermissions `aws:"IpPermissions"`
}

type AuthorizeSecurityGroupIngress struct {
	GroupId       string           `aws:"GroupId"`
	GroupName     string           `aws:"GroupName"`
	IpPermissions []*IpPermissions `aws:"IpPermissions"`
}

type BundleInstance struct {
	InstanceId                     string `aws:"InstanceId"`
	StorageS3Bucket                string `aws:"Storage.S3.Bucket"`
	StorageS3Prefix                string `aws:"Storage.S3.Prefix"`
	StorageS3AWSAccessKeyId        string `aws:"Storage.S3.AWSAccessKeyId"`
	StorageS3UploadPolicy          string `aws:"Storage.S3.UploadPolicy"`
	StorageS3UploadPolicySignature string `aws:"Storage.S3.UploadPolicySignature"`
}

type CancelBundleTask struct {
	BundleId string `aws:"BundleId"`
}

type CancelConversionTask struct {
	ConversionTaskId string `aws:"ConversionTaskId"`
}

type CancelExportTask struct {
	ExportTaskId string `aws:"ExportTaskId"`
}

type CancelReservedInstancesListing struct {
	ReservedInstancesListingId string `aws:"reservedInstancesListingId"`
}

type CancelSpotInstanceRequests struct {
	SpotInstanceRequestIds []string `aws:"SpotInstanceRequestId"`
}

type ConfirmProductInstance struct {
	ProductCode string `aws:"ProductCode"`
	InstanceId  string `aws:"InstanceId"`
}

type CopyImage struct {
	SourceRegion  string `aws:"SourceRegion"`
	SourceImageId string `aws:"SourceImageId"`
	Name          string `aws:"Name"`
	Description   string `aws:"Description"`
	ClientToken   string `aws:"ClientToken"`
}

type CopySnapshot struct {
	SourceRegion     string `aws:"SourceRegion"`
	SourceSnapshotId string `aws:"SourceSnapshotId"`
	Description      string `aws:"Description"`
}

type CreateCustomerGateway struct {
	Type      string    `aws:"Type"`
	IpAddress string    `aws:"IpAddress"`
	BgpAsn    *IntValue `aws:"BgpAsn"`
}

type CreateDhcpOptions struct {
	DhcpConfigurations []*DhcpConfiguration `aws:"DhcpConfiguration"`
}

type CreateImage struct {
	InstanceId          string                `aws:"InstanceId"`
	Name                string                `aws:"Name"`
	Description         string                `aws:"Description"`
	NoReboot            *BoolValue            `aws:"NoReboot"`
	BlockDeviceMappings []*BlockDeviceMapping `aws:"BlockDeviceMapping"`
}

type CreateInstanceExportTask struct {
	Description               string `aws:"Description"`
	InstanceId                string `aws:"InstanceId"`
	TargetEnvironment         string `aws:"TargetEnvironment"`
	ExportToS3DiskImageFormat string `aws:"ExportToS3.DiskImageFormat"`
	ExportToS3ContainerFormat string `aws:"ExportToS3.ContainerFormat"`
	ExportToS3S3Bucket        string `aws:"ExportToS3.S3Bucket"`
	ExportToS3S3Prefix        string `aws:"ExportToS3.S3Prefix"`
}

type CreateInternetGateway struct {
}

type CreateKeyPair struct {
	KeyName string `aws:"KeyName"`
}

type CreateNetworkAcl struct {
	VpcId string `aws:"VpcId"`
}

type CreateNetworkAclEntry struct {
	NetworkAclId  string     `aws:"NetworkAclId"`
	RuleNumber    *IntValue  `aws:"RuleNumber"`
	Protocol      *IntValue  `aws:"Protocol"`
	RuleAction    string     `aws:"RuleAction"`
	Egress        *BoolValue `aws:"Egress"`
	CidrBlock     string     `aws:"CidrBlock"`
	IcmpCode      *IntValue  `aws:"Icmp.Code"`
	IcmpType      *IntValue  `aws:"Icmp.Type"`
	PortRangeFrom *IntValue  `aws:"PortRange.From"`
	PortRangeTo   *IntValue  `aws:"PortRange.To"`
}

type CreateNetworkInterface struct {
	SubnetId                       string                    `aws:"SubnetId"`
	PrivateIpAddress               string                    `aws:"PrivateIpAddress"`
	SecondaryPrivateIpAddressCount *IntValue                 `aws:"SecondaryPrivateIpAddressCount"`
	Description                    string                    `aws:"Description"`
	SecurityGroupIds               []*SecurityGroupIdSetItem `aws:"SecurityGroupId"`
	PrivateIpAddresses             []*PrivateIpAddresses     `aws:"PrivateIpAddresses"`
}

type CreatePlacementGroup struct {
	GroupName string `aws:"GroupName"`
	Strategy  string `aws:"Strategy"`
}

type CreateReservedInstancesListing struct {
	ReservedInstancesId string                       `aws:"reservedInstancesId"`
	InstanceCount       *IntValue                    `aws:"instanceCount"`
	PriceSchedules      *PriceScheduleRequestSetItem `aws:"priceSchedules"`
	ClientToken         string                       `aws:"clientToken"`
}

type CreateRoute struct {
	RouteTableId         string `aws:"RouteTableId"`
	DestinationCidrBlock string `aws:"DestinationCidrBlock"`
	GatewayId            string `aws:"GatewayId"`
	InstanceId           string `aws:"InstanceId"`
	NetworkInterfaceId   string `aws:"NetworkInterfaceId"`
}

type CreateRouteTable struct {
	VpcId string `aws:"VpcId"`
}

type CreateSecurityGroup struct {
	GroupName        string `aws:"GroupName"`
	GroupDescription string `aws:"GroupDescription"`
	VpcId            string `aws:"VpcId"`
}

type CreateSnapshot struct {
	VolumeId    string `aws:"VolumeId"`
	Description string `aws:"Description"`
}

type CreateSpotDatafeedSubscription struct {
	Bucket string `aws:"Bucket"`
	Prefix string `aws:"Prefix"`
}

type CreateSubnet struct {
	VpcId            string `aws:"VpcId"`
	CidrBlock        string `aws:"CidrBlock"`
	AvailabilityZone string `aws:"AvailabilityZone"`
}

type CreateTags struct {
	ResourceIds []string `aws:"ResourceId"`
	Tags        []*Tag   `aws:"Tag"`
}

type CreateVolume struct {
	Size             string    `aws:"Size"`
	SnapshotId       string    `aws:"SnapshotId"`
	AvailabilityZone string    `aws:"AvailabilityZone"`
	VolumeType       string    `aws:"VolumeType"`
	Iops             *IntValue `aws:"Iops"`
}

type CreateVpc struct {
	CidrBlock       string `aws:"CidrBlock"`
	InstanceTenancy string `aws:"instanceTenancy"`
}

type CreateVpnConnection struct {
	Type                    string     `aws:"Type"`
	CustomerGatewayId       string     `aws:"CustomerGatewayId"`
	VpnGatewayId            string     `aws:"VpnGatewayId"`
	OptionsStaticRoutesOnly *BoolValue `aws:"Options.StaticRoutesOnly"`
}

type CreateVpnConnectionRoute struct {
	DestinationCidrBlock string `aws:"DestinationCidrBlock"`
	VpnConnectionId      string `aws:"VpnConnectionId"`
}

type CreateVpnGateway struct {
	Type string `aws:"Type"`
}

type DeleteCustomerGateway struct {
	CustomerGatewayId string `aws:"CustomerGatewayId"`
}

type DeleteDhcpOptions struct {
	DhcpOptionsId string `aws:"DhcpOptionsId"`
}

type DeleteInternetGateway struct {
	InternetGatewayId string `aws:"InternetGatewayId"`
}

type DeleteKeyPair struct {
	KeyName string `aws:"KeyName"`
}

type DeleteNetworkAcl struct {
	NetworkAclId string `aws:"NetworkAclId"`
}

type DeleteNetworkAclEntry struct {
	NetworkAclId string     `aws:"NetworkAclId"`
	RuleNumber   *IntValue  `aws:"RuleNumber"`
	Egress       *BoolValue `aws:"Egress"`
}

type DeleteNetworkInterface struct {
	NetworkInterfaceId string `aws:"NetworkInterfaceId"`
}

type DeletePlacementGroup struct {
	GroupName string `aws:"GroupName"`
}

type DeleteRoute struct {
	RouteTableId         string `aws:"RouteTableId"`
	DestinationCidrBlock string `aws:"DestinationCidrBlock"`
}

type DeleteRouteTable struct {
	RouteTableId string `aws:"RouteTableId"`
}

type DeleteSecurityGroup struct {
	GroupName string `aws:"GroupName"`
	GroupId   string `aws:"GroupId"`
}

type DeleteSnapshot struct {
	SnapshotId string `aws:"SnapshotId"`
}

type DeleteSpotDatafeedSubscription struct {
}

type DeleteSubnet struct {
	SubnetId string `aws:"SubnetId"`
}

type DeleteTags struct {
	ResourceIds []string `aws:"ResourceId"`
	Tags        []*Tag   `aws:"Tag"`
}

type DeleteVolume struct {
	VolumeId string `aws:"VolumeId"`
}

type DeleteVpc struct {
	VpcId string `aws:"VpcId"`
}

type DeleteVpnConnection struct {
	VpnConnectionId string `aws:"VpnConnectionId"`
}

type DeleteVpnConnectionRoute struct {
	DestinationCidrBlock string `aws:"DestinationCidrBlock"`
	VpnConnectionId      string `aws:"VpnConnectionId"`
}

type DeleteVpnGateway struct {
	VpnGatewayId string `aws:"VpnGatewayId"`
}

type DeregisterImage struct {
	ImageId string `aws:"ImageId"`
}

type DescribeAccountAttributes struct {
	AttributeNames []string `aws:"AttributeName"`
}

type DescribeAddresses struct {
	PublicIps     []string  `aws:"PublicIp"`
	AllocationIds []string  `aws:"AllocationId"`
	Filters       []*Filter `aws:"Filter"`
}

type DescribeAvailabilityZones struct {
	ZoneNames []string  `aws:"ZoneName"`
	Filters   []*Filter `aws:"Filter"`
}

type DescribeBundleTasks struct {
	BundleIds []string  `aws:"BundleId"`
	Filters   []*Filter `aws:"Filter"`
}

type DescribeConversionTasks struct {
	ConversionTaskIds []string `aws:"ConversionTaskId"`
}

type DescribeCustomerGateways struct {
	CustomerGatewayIds []string  `aws:"CustomerGatewayId"`
	Filters            []*Filter `aws:"Filter"`
}

type DescribeDhcpOptions struct {
	DhcpOptionsIds []string  `aws:"DhcpOptionsId"`
	Filters        []*Filter `aws:"Filter"`
}

type DescribeExportTasks struct {
	ExportTaskIds []string `aws:"ExportTaskId"`
}

type DescribeImageAttribute struct {
	ImageId   string `aws:"ImageId"`
	Attribute string `aws:"Attribute"`
}

type DescribeImages struct {
	ExecutableBys []string  `aws:"ExecutableBy"`
	ImageIds      []string  `aws:"ImageId"`
	Owners        []string  `aws:"Owner"`
	Filters       []*Filter `aws:"Filter"`
}

type DescribeInstanceAttribute struct {
	InstanceId string `aws:"InstanceId"`
	Attribute  string `aws:"Attribute"`
}

type DescribeInstances struct {
	InstanceIds []string  `aws:"InstanceId"`
	MaxResults  *IntValue `aws:"MaxResults"`
	NextToken   string    `aws:"NextToken"`
	Filters     []*Filter `aws:"Filter"`
}

type DescribeInstanceStatus struct {
	InstanceId          string     `aws:"InstanceId"`
	IncludeAllInstances *BoolValue `aws:"IncludeAllInstances"`
	MaxResults          *IntValue  `aws:"MaxResults"`
	NextToken           string     `aws:"NextToken"`
	Filters             []*Filter  `aws:"Filter"`
}

type DescribeInternetGateways struct {
	InternetGatewayIds []string  `aws:"InternetGatewayId"`
	Filters            []*Filter `aws:"Filter"`
}

type DescribeKeyPairs struct {
	KeyNames []string  `aws:"KeyName"`
	Filters  []*Filter `aws:"Filter"`
}

type DescribeNetworkAcls struct {
	NetworkAclIds []string  `aws:"NetworkAclId"`
	Filters       []*Filter `aws:"Filter"`
}

type DescribeNetworkInterfaceAttribute struct {
	NetworkInterfaceId string `aws:"NetworkInterfaceId"`
	Attribute          string `aws:"Attribute"`
}

type DescribeNetworkInterfaces struct {
	NetworkInterfaceIds []string  `aws:"NetworkInterfaceId"`
	Filters             []*Filter `aws:"Filter"`
}

type DescribePlacementGroups struct {
	GroupNames []string  `aws:"GroupName"`
	Filters    []*Filter `aws:"Filter"`
}

type DescribeRegions struct {
	RegionNames []string  `aws:"RegionName"`
	Filters     []*Filter `aws:"Filter"`
}

type DescribeReservedInstances struct {
	ReservedInstancesIds []string  `aws:"ReservedInstancesId"`
	OfferingType         string    `aws:"offeringType"`
	Filters              []*Filter `aws:"Filter"`
}

type DescribeReservedInstancesListings struct {
	ReservedInstancesListingIds []*DescribeReservedInstancesListingSetItem `aws:"ReservedInstancesListingId"`
	ReservedInstancesIds        []*DescribeReservedInstancesSetItem        `aws:"ReservedInstancesId"`
	Filters                     []*Filter                                  `aws:"Filter"`
}

type DescribeReservedInstancesModifications struct {
	ReservedInstancesModificationIds []string  `aws:"reservedInstancesModificationId"`
	NextToken                        string    `aws:"nextToken"`
	Filters                          []*Filter `aws:"Filter"`
}

type DescribeReservedInstancesOfferings struct {
	ReservedInstancesOfferingIds []string   `aws:"ReservedInstancesOfferingId"`
	InstanceType                 string     `aws:"InstanceType"`
	AvailabilityZone             string     `aws:"AvailabilityZone"`
	ProductDescription           string     `aws:"ProductDescription"`
	InstanceTenancy              string     `aws:"InstanceTenancy"`
	OfferingType                 string     `aws:"OfferingType"`
	IncludeMarketplace           *BoolValue `aws:"IncludeMarketplace"`
	MinDuration                  *IntValue  `aws:"MinDuration"`
	MaxDuration                  *IntValue  `aws:"MaxDuration"`
	MaxInstanceCount             *IntValue  `aws:"MaxInstanceCount"`
	NextToken                    string     `aws:"NextToken"`
	MaxResults                   *IntValue  `aws:"MaxResults"`
	Filters                      []*Filter  `aws:"Filter"`
}

type DescribeRouteTables struct {
	RouteTableIds []string  `aws:"RouteTableId"`
	Filters       []*Filter `aws:"Filter"`
}

type DescribeSecurityGroups struct {
	GroupNames []string  `aws:"GroupName"`
	GroupIds   []string  `aws:"GroupId"`
	Filters    []*Filter `aws:"Filter"`
}

type DescribeSnapshotAttribute struct {
	SnapshotId string `aws:"SnapshotId"`
	Attribute  string `aws:"Attribute"`
}

type DescribeSnapshots struct {
	SnapshotIds   []string  `aws:"SnapshotId"`
	Owners        []string  `aws:"Owner"`
	RestorableBys []string  `aws:"RestorableBy"`
	Filters       []*Filter `aws:"Filter"`
}

type DescribeSpotDatafeedSubscription struct {
}

type DescribeSpotInstanceRequests struct {
	SpotInstanceRequestIds []string  `aws:"SpotInstanceRequestId"`
	Filters                []*Filter `aws:"Filter"`
}

type DescribeSpotPriceHistory struct {
	StartTime           time.Time `aws:"StartTime"`
	EndTime             time.Time `aws:"EndTime"`
	InstanceTypes       []string  `aws:"InstanceType"`
	ProductDescriptions []string  `aws:"ProductDescription"`
	AvailabilityZone    string    `aws:"AvailabilityZone"`
	MaxResults          *IntValue `aws:"MaxResults"`
	NextToken           string    `aws:"NextToken"`
	Filters             []*Filter `aws:"Filter"`
}

type DescribeSubnets struct {
	SubnetIds []string  `aws:"SubnetId"`
	Filters   []*Filter `aws:"Filter"`
}

type DescribeTags struct {
	MaxResults *IntValue `aws:"MaxResults"`
	NextToken  string    `aws:"NextToken"`
	Filters    []*Filter `aws:"Filter"`
}

type DescribeVolumeAttribute struct {
	VolumeId  string `aws:"VolumeId"`
	Attribute string `aws:"Attribute"`
}

type DescribeVolumes struct {
	VolumeIds []string  `aws:"VolumeId"`
	Filters   []*Filter `aws:"Filter"`
}

type DescribeVolumeStatus struct {
	VolumeIds  []string  `aws:"VolumeId"`
	MaxResults *IntValue `aws:"MaxResults"`
	NextToken  string    `aws:"NextToken"`
	Filters    []*Filter `aws:"Filter"`
}

type DescribeVpcAttribute struct {
	VpcId     string `aws:"VpcId"`
	Attribute string `aws:"Attribute"`
}

type DescribeVpcs struct {
	VpcIds  []string  `aws:"vpcId"`
	Filters []*Filter `aws:"Filter"`
}

type DescribeVpnConnections struct {
	VpnConnectionIds []string  `aws:"VpnConnectionId"`
	Filters          []*Filter `aws:"Filter"`
}

type DescribeVpnGateways struct {
	VpnGatewayIds []string  `aws:"VpnGatewayId"`
	Filters       []*Filter `aws:"Filter"`
}

type DetachInternetGateway struct {
	InternetGatewayId string `aws:"InternetGatewayId"`
	VpcId             string `aws:"VpcId"`
}

type DetachNetworkInterface struct {
	AttachmentId string     `aws:"AttachmentId"`
	Force        *BoolValue `aws:"Force"`
}

type DetachVolume struct {
	VolumeId   string     `aws:"VolumeId"`
	InstanceId string     `aws:"InstanceId"`
	Device     string     `aws:"Device"`
	Force      *BoolValue `aws:"Force"`
}

type DetachVpnGateway struct {
	VpnGatewayId string `aws:"VpnGatewayId"`
	VpcId        string `aws:"VpcId"`
}

type DisableVgwRoutePropagation struct {
	RouteTableId string `aws:"RouteTableId"`
	GatewayId    string `aws:"GatewayId"`
}

type DisassociateAddress struct {
	PublicIp      string `aws:"PublicIp"`
	AssociationId string `aws:"AssociationId"`
}

type DisassociateRouteTable struct {
	AssociationId string `aws:"AssociationId"`
}

type EnableVgwRoutePropagation struct {
	RouteTableId string `aws:"RouteTableId"`
	GatewayId    string `aws:"GatewayId"`
}

type EnableVolumeIO struct {
	VolumeId string `aws:"VolumeId"`
}

type GetConsoleOutput struct {
	InstanceId string `aws:"InstanceId"`
}

type GetPasswordData struct {
	InstanceId string `aws:"InstanceId"`
}

type ImportInstance struct {
	Description                                          string       `aws:"Description"`
	LaunchSpecificationArchitecture                      string       `aws:"LaunchSpecification.Architecture"`
	LaunchSpecificationGroupNames                        []string     `aws:"LaunchSpecification.GroupName"`
	LaunchSpecificationUserData                          string       `aws:"LaunchSpecification.UserData"`
	LaunchSpecificationInstanceType                      string       `aws:"LaunchSpecification.InstanceType"`
	LaunchSpecificationPlacementAvailabilityZone         string       `aws:"LaunchSpecification.Placement.AvailabilityZone"`
	LaunchSpecificationMonitoringEnabled                 *BoolValue   `aws:"LaunchSpecification.Monitoring.Enabled"`
	LaunchSpecificationSubnetId                          string       `aws:"LaunchSpecification.SubnetId"`
	LaunchSpecificationInstanceInitiatedShutdownBehavior string       `aws:"LaunchSpecification.InstanceInitiatedShutdownBehavior"`
	LaunchSpecificationPrivateIpAddress                  string       `aws:"LaunchSpecification.PrivateIpAddress"`
	Platform                                             string       `aws:"Platform"`
	DiskImages                                           []*DiskImage `aws:"DiskImage"`
}

type ImportKeyPair struct {
	KeyName           string `aws:"KeyName"`
	PublicKeyMaterial string `aws:"PublicKeyMaterial"`
}

type ImportVolume struct {
	AvailabilityZone       string    `aws:"AvailabilityZone"`
	ImageFormat            string    `aws:"Image.Format"`
	ImageBytes             *IntValue `aws:"Image.Bytes"`
	ImageImportManifestUrl string    `aws:"Image.ImportManifestUrl"`
	Description            string    `aws:"Description"`
	VolumeSize             *IntValue `aws:"Volume.Size"`
}

type ModifyImageAttribute struct {
	ImageId                 string                    `aws:"ImageId"`
	ProductCodes            []string                  `aws:"ProductCode"`
	DescriptionValue        string                    `aws:"Description.Value"`
	LaunchPermissionAdds    []*LaunchPermissionAdd    `aws:"LaunchPermission.Add"`
	LaunchPermissionRemoves []*LaunchPermissionRemove `aws:"LaunchPermission.Remove"`
}

type ModifyInstanceAttribute struct {
	InstanceId                             string                          `aws:"InstanceId"`
	BlockDeviceMappingValue                *InstanceBlockDeviceMappingItem `aws:"BlockDeviceMapping.Value"`
	DisableApiTerminationValue             *BoolValue                      `aws:"DisableApiTermination.Value"`
	EbsOptimized                           *BoolValue                      `aws:"EbsOptimized"`
	GroupIds                               []string                        `aws:"GroupId"`
	InstanceInitiatedShutdownBehaviorValue string                          `aws:"InstanceInitiatedShutdownBehavior.Value"`
	InstanceTypeValue                      string                          `aws:"InstanceType.Value"`
	KernelValue                            string                          `aws:"Kernel.Value"`
	RamdiskValue                           string                          `aws:"Ramdisk.Value"`
	SourceDestCheckValue                   *BoolValue                      `aws:"SourceDestCheck.Value"`
	SriovNetSupportValue                   string                          `aws:"SriovNetSupport.Value"`
	UserDataValue                          string                          `aws:"UserData.Value"`
}

type ModifyNetworkInterfaceAttribute struct {
	NetworkInterfaceId            string     `aws:"NetworkInterfaceId"`
	DescriptionValue              string     `aws:"Description.Value"`
	SecurityGroupIds              []string   `aws:"SecurityGroupId"`
	SourceDestCheckValue          *BoolValue `aws:"SourceDestCheck.Value"`
	AttachmentAttachmentId        string     `aws:"Attachment.AttachmentId"`
	AttachmentDeleteOnTermination *BoolValue `aws:"Attachment.DeleteOnTermination"`
}

type ModifyReservedInstances struct {
	ReservedInstancesID string                                 `aws:"reservedInstancesID"`
	ClientToken         string                                 `aws:"clientToken"`
	TargetConfiguration *ReservedInstancesConfigurationSetItem `aws:"targetConfiguration"`
}

type ModifySnapshotAttribute struct {
	SnapshotId                    string                          `aws:"SnapshotId"`
	CreateVolumePermissionAdds    []*CreateVolumePermissionAdd    `aws:"CreateVolumePermission.Add"`
	CreateVolumePermissionRemoves []*CreateVolumePermissionRemove `aws:"CreateVolumePermission.Remove"`
}

type ModifyVolumeAttribute struct {
	VolumeId          string     `aws:"VolumeId"`
	AutoEnableIOValue *BoolValue `aws:"AutoEnableIO.Value"`
}

type ModifyVpcAttribute struct {
	VpcId              string     `aws:"VpcId"`
	EnableDnsSupport   *BoolValue `aws:"enableDnsSupport"`
	EnableDnsHostnames *BoolValue `aws:"enableDnsHostnames"`
}

type MonitorInstances struct {
	InstanceIds []string `aws:"InstanceId"`
}

type PurchaseReservedInstancesOffering struct {
	ReservedInstancesOfferingId string                      `aws:"reservedInstancesOfferingId"`
	InstanceCount               *IntValue                   `aws:"instanceCount"`
	LimitPrice                  *ReservedInstanceLimitPrice `aws:"limitPrice"`
}

type RebootInstances struct {
	InstanceIds []string `aws:"InstanceId"`
}

type RegisterImage struct {
	ImageLocation       string                `aws:"ImageLocation"`
	Name                string                `aws:"Name"`
	Description         string                `aws:"Description"`
	Architecture        string                `aws:"Architecture"`
	RootDeviceName      string                `aws:"RootDeviceName"`
	VirtualizationType  string                `aws:"VirtualizationType"`
	KernelId            string                `aws:"KernelId"`
	RamdiskId           string                `aws:"RamdiskId"`
	SriovNetSupport     string                `aws:"SriovNetSupport"`
	BlockDeviceMappings []*BlockDeviceMapping `aws:"BlockDeviceMapping"`
}

type ReleaseAddress struct {
	PublicIp     string `aws:"PublicIp"`
	AllocationId string `aws:"AllocationId"`
}

type ReplaceNetworkAclAssociation struct {
	AssociationId string `aws:"AssociationId"`
	NetworkAclId  string `aws:"NetworkAclId"`
}

type ReplaceNetworkAclEntry struct {
	NetworkAclId  string     `aws:"NetworkAclId"`
	RuleNumber    *IntValue  `aws:"RuleNumber"`
	Protocol      *IntValue  `aws:"Protocol"`
	RuleAction    string     `aws:"RuleAction"`
	Egress        *BoolValue `aws:"Egress"`
	CidrBlock     string     `aws:"CidrBlock"`
	IcmpCode      *IntValue  `aws:"Icmp.Code"`
	IcmpType      *IntValue  `aws:"Icmp.Type"`
	PortRangeFrom *IntValue  `aws:"PortRange.From"`
	PortRangeTo   *IntValue  `aws:"PortRange.To"`
}

type ReplaceRoute struct {
	RouteTableId         string `aws:"RouteTableId"`
	DestinationCidrBlock string `aws:"DestinationCidrBlock"`
	GatewayId            string `aws:"GatewayId"`
	InstanceId           string `aws:"InstanceId"`
	NetworkInterfaceId   string `aws:"NetworkInterfaceId"`
}

type ReplaceRouteTableAssociation struct {
	AssociationId string `aws:"AssociationId"`
	RouteTableId  string `aws:"RouteTableId"`
}

type ReportInstanceStatus struct {
	InstanceIds []string  `aws:"InstanceId"`
	Status      string    `aws:"Status"`
	StartTime   time.Time `aws:"StartTime"`
	EndTime     time.Time `aws:"EndTime"`
	ReasonCodes []string  `aws:"ReasonCode"`
	Description string    `aws:"Description"`
}

type RequestSpotInstances struct {
	SpotPrice                                           string                                   `aws:"SpotPrice"`
	InstanceCount                                       *IntValue                                `aws:"InstanceCount"`
	Type                                                string                                   `aws:"Type"`
	ValidFrom                                           time.Time                                `aws:"ValidFrom"`
	ValidUntil                                          time.Time                                `aws:"ValidUntil"`
	LaunchGroup                                         string                                   `aws:"LaunchGroup"`
	AvailabilityZoneGroup                               string                                   `aws:"AvailabilityZoneGroup"`
	LaunchSpecificationImageId                          string                                   `aws:"LaunchSpecification.ImageId"`
	LaunchSpecificationKeyName                          string                                   `aws:"LaunchSpecification.KeyName"`
	LaunchSpecificationSecurityGroupIds                 []string                                 `aws:"LaunchSpecification.SecurityGroupId"`
	LaunchSpecificationSecurityGroups                   []string                                 `aws:"LaunchSpecification.SecurityGroup"`
	LaunchSpecificationUserData                         string                                   `aws:"LaunchSpecification.UserData"`
	LaunchSpecificationInstanceType                     string                                   `aws:"LaunchSpecification.InstanceType"`
	LaunchSpecificationPlacementAvailabilityZone        string                                   `aws:"LaunchSpecification.Placement.AvailabilityZone"`
	LaunchSpecificationPlacementGroupName               string                                   `aws:"LaunchSpecification.Placement.GroupName"`
	LaunchSpecificationKernelId                         string                                   `aws:"LaunchSpecification.KernelId"`
	LaunchSpecificationRamdiskId                        string                                   `aws:"LaunchSpecification.RamdiskId"`
	LaunchSpecificationMonitoringEnabled                string                                   `aws:"LaunchSpecification.Monitoring.Enabled"`
	LaunchSpecificationSubnetId                         string                                   `aws:"LaunchSpecification.SubnetId"`
	LaunchSpecificationNetworkInterfaceSecurityGroupIds []string                                 `aws:"LaunchSpecification.NetworkInterface.SecurityGroupId"`
	LaunchSpecificationIamInstanceProfileArn            string                                   `aws:"LaunchSpecification.IamInstanceProfile.Arn"`
	LaunchSpecificationIamInstanceProfileName           string                                   `aws:"LaunchSpecification.IamInstanceProfile.Name"`
	LaunchSpecificationEbsOptimized                     *BoolValue                               `aws:"LaunchSpecification.EbsOptimized"`
	LaunchSpecificationBlockDeviceMappings              []*LaunchSpecificationBlockDeviceMapping `aws:"LaunchSpecification.BlockDeviceMapping"`
	LaunchSpecificationNetworkInterfaces                []*LaunchSpecificationNetworkInterface   `aws:"LaunchSpecification.NetworkInterface"`
}

type ResetImageAttribute struct {
	ImageId   string `aws:"ImageId"`
	Attribute string `aws:"Attribute"`
}

type ResetInstanceAttribute struct {
	InstanceId string `aws:"InstanceId"`
	Attribute  string `aws:"Attribute"`
}

type ResetNetworkInterfaceAttribute struct {
	NetworkInterfaceId string `aws:"NetworkInterfaceId"`
	Attribute          string `aws:"Attribute"`
}

type ResetSnapshotAttribute struct {
	SnapshotId string `aws:"SnapshotId"`
	Attribute  string `aws:"Attribute"`
}

type RevokeSecurityGroupEgress struct {
	GroupId       string           `aws:"GroupId"`
	IpPermissions []*IpPermissions `aws:"IpPermissions"`
}

type RevokeSecurityGroupIngress struct {
	GroupId       string           `aws:"GroupId"`
	GroupName     string           `aws:"GroupName"`
	IpPermissions []*IpPermissions `aws:"IpPermissions"`
}

type RunInstances struct {
	ImageId                           string                     `aws:"ImageId"`
	MinCount                          *IntValue                  `aws:"MinCount"`
	MaxCount                          *IntValue                  `aws:"MaxCount"`
	KeyName                           string                     `aws:"KeyName"`
	SecurityGroupIds                  []string                   `aws:"SecurityGroupId"`
	SecurityGroups                    []string                   `aws:"SecurityGroup"`
	UserData                          string                     `aws:"UserData"`
	InstanceType                      string                     `aws:"InstanceType"`
	PlacementAvailabilityZone         string                     `aws:"Placement.AvailabilityZone"`
	PlacementGroupName                string                     `aws:"Placement.GroupName"`
	PlacementTenancy                  string                     `aws:"Placement.Tenancy"`
	KernelId                          string                     `aws:"KernelId"`
	RamdiskId                         string                     `aws:"RamdiskId"`
	MonitoringEnabled                 *BoolValue                 `aws:"Monitoring.Enabled"`
	SubnetId                          string                     `aws:"SubnetId"`
	DisableApiTermination             *BoolValue                 `aws:"DisableApiTermination"`
	InstanceInitiatedShutdownBehavior string                     `aws:"InstanceInitiatedShutdownBehavior"`
	PrivateIpAddress                  string                     `aws:"PrivateIpAddress"`
	ClientToken                       string                     `aws:"ClientToken"`
	NetworkInterfaceSecurityGroupIds  []string                   `aws:"NetworkInterface.SecurityGroupId"`
	IamInstanceProfileArn             string                     `aws:"IamInstanceProfile.Arn"`
	IamInstanceProfileName            string                     `aws:"IamInstanceProfile.Name"`
	EbsOptimized                      *BoolValue                 `aws:"EbsOptimized"`
	BlockDeviceMappings               []*BlockDeviceMapping      `aws:"BlockDeviceMapping"`
	NetworkInterfaces                 []*RequestNetworkInterface `aws:"NetworkInterface"`
}

type StartInstances struct {
	InstanceIds []string `aws:"InstanceId"`
}

type StopInstances struct {
	InstanceIds []string   `aws:"InstanceId"`
	Force       *BoolValue `aws:"Force"`
}

type TerminateInstances struct {
	InstanceIds []string `aws:"InstanceId"`
}

type UnassignPrivateIpAddresses struct {
	NetworkInterfaceId string                                    `aws:"NetworkInterfaceId"`
	PrivateIpAddress   []*AssignPrivateIpAddressesSetItemRequest `aws:"PrivateIpAddress"`
}

type UnmonitorInstances struct {
	InstanceIds []string `aws:"InstanceId"`
}

type IpPermissions struct {
	IpProtocol string      `aws:"IpProtocol"`
	FromPort   *IntValue   `aws:"FromPort"`
	ToPort     *IntValue   `aws:"ToPort"`
	Groups     []*Groups   `aws:"Groups"`
	IpRanges   []*IpRanges `aws:"IpRanges"`
}

type IpRanges struct {
	CidrIp string `aws:"CidrIp"`
}

type DhcpConfiguration struct {
	Key    string   `aws:"Key"`
	Values []string `aws:"Values"`
}

type PrivateIpAddresses struct {
	PrivateIpAddress string     `aws:"PrivateIpAddress"`
	Primary          *BoolValue `aws:"Primary"`
}

type CreateVolumePermissionAdd struct {
	UserId string `aws:"UserId"`
	Group  string `aws:"Group"`
}

type CreateVolumePermissionRemove struct {
	UserId string `aws:"UserId"`
	Group  string `aws:"Group"`
}

type DiskImage struct {
	ImageFormat            string    `aws:"Image.Format"`
	ImageBytes             *IntValue `aws:"Image.Bytes"`
	ImageImportManifestUrl string    `aws:"Image.ImportManifestUrl"`
	ImageDescription       string    `aws:"Image.Description"`
	VolumeSize             struct{}  `aws:"Volume.Size"`
}

type LaunchPermissionAdd struct {
	UserId string `aws:"UserId"`
	Group  string `aws:"Group"`
}

type LaunchSpecificationNetworkInterface struct {
	NetworkInterfaceId             string     `aws:"NetworkInterfaceId"`
	DeviceIndex                    *IntValue  `aws:"DeviceIndex"`
	SubnetId                       string     `aws:"SubnetId"`
	Description                    string     `aws:"Description"`
	SecondaryPrivateIpAddressCount *IntValue  `aws:"SecondaryPrivateIpAddressCount\n\t\t\t\t\t\t"`
	AssociatePublicIpAddress       *BoolValue `aws:"AssociatePublicIpAddress"`
	PrivateIpAddress               string     `aws:"PrivateIpAddress"`
	PrivateIpAddresses             string     `aws:"PrivateIpAddresses"`
	DeleteOnTermination            *BoolValue `aws:"DeleteOnTermination"`
}

type Groups struct {
	GroupName string `aws:"GroupName "`
	GroupId   string `aws:"GroupId "`
	UserId    string `aws:"UserId "`
}

type BlockDeviceMapping struct {
	DeviceName             string     `aws:"DeviceName"`
	VirtualName            string     `aws:"VirtualName"`
	EbsSnapshotId          string     `aws:"Ebs.SnapshotId"`
	EbsVolumeSize          *IntValue  `aws:"Ebs.VolumeSize"`
	EbsDeleteOnTermination *BoolValue `aws:"Ebs.DeleteOnTermination"`
	EbsVolumeType          string     `aws:"Ebs.VolumeType"`
	EbsIops                *IntValue  `aws:"Ebs.Iops"`
	NoDevice               struct{}   `aws:"NoDevice"`
}

type Tag struct {
	Key   string `aws:"Key"`
	Value string `aws:"Value"`
}

type Filter struct {
	Name   string   `aws:"Name"`
	Values []string `aws:"Values"`
}

type LaunchPermissionRemove struct {
	UserId string `aws:"UserId"`
	Group  string `aws:"Group"`
}

type LaunchSpecificationBlockDeviceMapping struct {
	DeviceName             string     `aws:"DeviceName"`
	NoDevice               *BoolValue `aws:"NoDevice"`
	VirtualName            string     `aws:"VirtualName"`
	EbsSnapshotId          string     `aws:"Ebs.SnapshotId"`
	EbsVolumeSize          *IntValue  `aws:"Ebs.VolumeSize"`
	EbsDeleteOnTermination *BoolValue `aws:"Ebs.DeleteOnTermination"`
	EbsVolumeType          string     `aws:"Ebs.VolumeType"`
	EbsIops                *IntValue  `aws:"Ebs.Iops"`
}

type RequestNetworkInterface struct {
	Description                    string     `aws:"Description"`
	DeleteOnTermination            *BoolValue `aws:"DeleteOnTermination"`
	SecondaryPrivateIpAddressCount struct{}   `aws:"SecondaryPrivateIpAddressCount"`
	AssociatePublicIpAddress       *BoolValue `aws:"AssociatePublicIpAddress"`
	NetworkInterfaceId             string     `aws:"NetworkInterfaceId"`
	DeviceIndex                    *IntValue  `aws:"DeviceIndex"`
	SubnetId                       string     `aws:"SubnetId"`
	PrivateIpAddress               string     `aws:"PrivateIpAddress"`
	PrivateIpAddresses             *BoolValue `aws:"PrivateIpAddresses"`
}
