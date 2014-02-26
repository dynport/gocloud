package generated

import "time"

type AccountAttributeSetItem struct {
	AttributeName   string                          `xml:"attributeName,omitempty"`
	AttributeValues []*AccountAttributeValueSetItem `xml:"attributeValueSet>item,omitempty"`
}

type AccountAttributeValueSetItem struct {
	AttributeValue string `xml:"attributeValue,omitempty"`
}

type AssignPrivateIpAddressesSetItemRequest struct {
	PrivateIpAddress string `xml:"privateIpAddress,omitempty"`
}

type AttachmentSetItemResponse struct {
	VolumeId            string    `xml:"volumeId,omitempty"`
	InstanceId          string    `xml:"instanceId,omitempty"`
	Device              string    `xml:"device,omitempty"`
	Status              string    `xml:"status,omitempty"`
	AttachTime          time.Time `xml:"attachTime,omitempty"`
	DeleteOnTermination bool      `xml:"deleteOnTermination,omitempty"`
}

type Attachment struct {
	VpcId string `xml:"vpcId,omitempty"`
	State string `xml:"state,omitempty"`
}

type AvailabilityZoneItem struct {
	ZoneName   string                   `xml:"zoneName,omitempty"`
	ZoneState  string                   `xml:"zoneState,omitempty"`
	RegionName string                   `xml:"regionName,omitempty"`
	MessageSet *AvailabilityZoneMessage `xml:"messageSet,omitempty"`
}

type AvailabilityZoneMessage struct {
	Message string `xml:"message,omitempty"`
}

type BlockDeviceMappingItem struct {
	DeviceName  string          `xml:"deviceName,omitempty"`
	VirtualName string          `xml:"virtualName,omitempty"`
	Ebs         *EbsBlockDevice `xml:"ebs,omitempty"`
	NoDevice    struct{}        `xml:"noDevice,omitempty"`
}

type BundleInstanceS3Storage struct {
	AwsAccessKeyId        string `xml:"awsAccessKeyId,omitempty"`
	Bucket                string `xml:"bucket,omitempty"`
	Prefix                string `xml:"prefix,omitempty"`
	UploadPolicy          string `xml:"uploadPolicy,omitempty"`
	UploadPolicySignature string `xml:"uploadPolicySignature,omitempty"`
}

type BundleInstanceTaskError struct {
	Code    string `xml:"code,omitempty"`
	Message string `xml:"message,omitempty"`
}

type BundleInstanceTaskStorage struct {
	S3 *BundleInstanceS3Storage `xml:"S3,omitempty"`
}

type BundleInstanceTask struct {
	InstanceId string                     `xml:"instanceId,omitempty"`
	BundleId   string                     `xml:"bundleId,omitempty"`
	State      string                     `xml:"state,omitempty"`
	StartTime  time.Time                  `xml:"startTime,omitempty"`
	UpdateTime time.Time                  `xml:"updateTime,omitempty"`
	Storage    *BundleInstanceTaskStorage `xml:"storage,omitempty"`
	Progress   string                     `xml:"progress,omitempty"`
	Error      *BundleInstanceTaskError   `xml:"error,omitempty"`
}

type CancelSpotInstanceRequestsResponseSetItem struct {
	SpotInstanceRequestId string `xml:"spotInstanceRequestId,omitempty"`
	State                 string `xml:"state,omitempty"`
}

type ConversionTask struct {
	ConversionTaskId string                     `xml:"conversionTaskId,omitempty"`
	ExpirationTime   string                     `xml:"expirationTime,omitempty"`
	ImportVolume     *ImportVolumeTaskDetails   `xml:"importVolume,omitempty"`
	ImportInstance   *ImportInstanceTaskDetails `xml:"importInstance,omitempty"`
	State            string                     `xml:"state,omitempty"`
	StatusMessage    string                     `xml:"statusMessage,omitempty"`
}

type CreateVolumePermissionItem struct {
	UserId string `xml:"userId,omitempty"`
	Group  string `xml:"group,omitempty"`
}

type CustomerGateway struct {
	CustomerGatewayId string                `xml:"customerGatewayId,omitempty"`
	State             string                `xml:"state,omitempty"`
	Type              string                `xml:"type,omitempty"`
	IpAddress         string                `xml:"ipAddress,omitempty"`
	BgpAsn            int                   `xml:"bgpAsn,omitempty"`
	Tags              []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
}

type DescribeAddressesResponseItem struct {
	PublicIp                string `xml:"publicIp,omitempty"`
	AllocationId            string `xml:"allocationId,omitempty"`
	Domain                  string `xml:"domain,omitempty"`
	InstanceId              string `xml:"instanceId,omitempty"`
	AssociationId           string `xml:"associationId,omitempty"`
	NetworkInterfaceId      string `xml:"networkInterfaceId,omitempty"`
	NetworkInterfaceOwnerId string `xml:"networkInterfaceOwnerId,omitempty"`
}

type DescribeImagesResponseItem struct {
	ImageId             string                    `xml:"imageId,omitempty"`
	ImageLocation       string                    `xml:"imageLocation,omitempty"`
	ImageState          string                    `xml:"imageState,omitempty"`
	ImageOwnerId        string                    `xml:"imageOwnerId,omitempty"`
	IsPublic            bool                      `xml:"isPublic,omitempty"`
	ProductCodes        []*ProductCodesSetItem    `xml:"productCodes>item,omitempty"`
	Architecture        string                    `xml:"architecture,omitempty"`
	ImageType           string                    `xml:"imageType,omitempty"`
	KernelId            string                    `xml:"kernelId,omitempty"`
	RamdiskId           string                    `xml:"ramdiskId,omitempty"`
	Platform            string                    `xml:"platform,omitempty"`
	SriovNetSupport     string                    `xml:"sriovNetSupport,omitempty"`
	StateReason         *StateReason              `xml:"stateReason,omitempty"`
	ImageOwnerAlias     string                    `xml:"imageOwnerAlias,omitempty"`
	Name                string                    `xml:"name,omitempty"`
	Description         string                    `xml:"description,omitempty"`
	RootDeviceType      string                    `xml:"rootDeviceType,omitempty"`
	RootDeviceName      string                    `xml:"rootDeviceName,omitempty"`
	BlockDeviceMappings []*BlockDeviceMappingItem `xml:"blockDeviceMapping>item,omitempty"`
	VirtualizationType  string                    `xml:"virtualizationType,omitempty"`
	Tags                []*ResourceTagSetItem     `xml:"tagSet>item,omitempty"`
	Hypervisor          string                    `xml:"hypervisor,omitempty"`
}

type DescribeKeyPairsResponseItem struct {
	KeyName        string `xml:"keyName,omitempty"`
	KeyFingerprint string `xml:"keyFingerprint,omitempty"`
}

type DescribeReservedInstancesListingsResponseSetItem struct {
	ReservedInstancesListingId string                `xml:"reservedInstancesListingId,omitempty"`
	ReservedInstancesId        string                `xml:"reservedInstancesId,omitempty"`
	CreateDate                 time.Time             `xml:"createDate,omitempty"`
	UpdateDate                 time.Time             `xml:"updateDate,omitempty"`
	Status                     string                `xml:"status,omitempty"`
	StatusMessage              string                `xml:"statusMessage,omitempty"`
	InstanceCounts             []*InstanceCountsSet  `xml:"instanceCounts,omitempty"`
	PriceSchedules             []*PriceScheduleSet   `xml:"priceSchedules,omitempty"`
	Tags                       []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
	ClientToken                string                `xml:"clientToken,omitempty"`
}

type DescribeReservedInstancesListingSetItem struct {
	ReservedInstancesListingId string `xml:"reservedInstancesListingId,omitempty"`
}

type DescribeReservedInstancesModificationsResponseSetItem struct {
	ReservedInstancesModificationId string                                   `xml:"reservedInstancesModificationId,omitempty"`
	ClientToken                     string                                   `xml:"clientToken,omitempty"`
	ReservedInstancesId             string                                   `xml:"reservedInstancesId,omitempty"`
	ModificationResults             []*ReservedInstancesConfigurationSetItem `xml:"modificationResults>item,omitempty"`
	CreateDate                      string                                   `xml:"createDate,omitempty"`
	UpdateDate                      string                                   `xml:"updateDate,omitempty"`
	EffectiveDate                   string                                   `xml:"effectiveDate,omitempty"`
	Status                          string                                   `xml:"status,omitempty"`
	StatusMessage                   string                                   `xml:"statusMessage,omitempty"`
}

type DescribeReservedInstancesOfferingsResponseSetItem struct {
	ReservedInstancesOfferingId string                     `xml:"reservedInstancesOfferingId,omitempty"`
	InstanceType                string                     `xml:"instanceType,omitempty"`
	AvailabilityZone            string                     `xml:"availabilityZone,omitempty"`
	Duration                    int                        `xml:"duration,omitempty"`
	FixedPrice                  float64                    `xml:"fixedPrice,omitempty"`
	UsagePrice                  float64                    `xml:"usagePrice,omitempty"`
	ProductDescription          string                     `xml:"productDescription,omitempty"`
	InstanceTenancy             string                     `xml:"instanceTenancy,omitempty"`
	CurrencyCode                string                     `xml:"currencyCode,omitempty"`
	OfferingType                string                     `xml:"offeringType,omitempty"`
	RecurringCharges            []*RecurringChargesSetItem `xml:"recurringCharges>item,omitempty"`
	Marketplace                 bool                       `xml:"marketplace,omitempty"`
	PricingDetails              []*PricingDetailsSetItem   `xml:"pricingDetailsSet>item,omitempty"`
}

type DescribeReservedInstancesOfferingsResponse struct {
	RequestId                  string                                               `xml:"requestId,omitempty"`
	ReservedInstancesOfferings []*DescribeReservedInstancesOfferingsResponseSetItem `xml:"reservedInstancesOfferingsSet>item,omitempty"`
	NextToken                  string                                               `xml:"nextToken,omitempty"`
}

type DescribeReservedInstancesResponseSetItem struct {
	ReservedInstancesId string                     `xml:"reservedInstancesId,omitempty"`
	InstanceType        string                     `xml:"instanceType,omitempty"`
	AvailabilityZone    string                     `xml:"availabilityZone,omitempty"`
	Start               time.Time                  `xml:"start,omitempty"`
	Duration            int                        `xml:"duration,omitempty"`
	End                 time.Time                  `xml:"end,omitempty"`
	FixedPrice          float64                    `xml:"fixedPrice,omitempty"`
	UsagePrice          float64                    `xml:"usagePrice,omitempty"`
	InstanceCount       int                        `xml:"instanceCount,omitempty"`
	ProductDescription  string                     `xml:"productDescription,omitempty"`
	State               string                     `xml:"state,omitempty"`
	Tags                []*ResourceTagSetItem      `xml:"tagSet>item,omitempty"`
	InstanceTenancy     string                     `xml:"instanceTenancy,omitempty"`
	CurrencyCode        string                     `xml:"currencyCode,omitempty"`
	OfferingType        string                     `xml:"offeringType,omitempty"`
	RecurringCharges    []*RecurringChargesSetItem `xml:"recurringCharges>item,omitempty"`
}

type DescribeReservedInstancesSetItem struct {
	ReservedInstancesId string `xml:"reservedInstancesId,omitempty"`
}

type DescribeSnapshotsSetItemResponse struct {
	SnapshotId  string                `xml:"snapshotId,omitempty"`
	VolumeId    string                `xml:"volumeId,omitempty"`
	Status      string                `xml:"status,omitempty"`
	StartTime   time.Time             `xml:"startTime,omitempty"`
	Progress    string                `xml:"progress,omitempty"`
	OwnerId     string                `xml:"ownerId,omitempty"`
	VolumeSize  string                `xml:"volumeSize,omitempty"`
	Description string                `xml:"description,omitempty"`
	OwnerAlias  string                `xml:"ownerAlias,omitempty"`
	Tags        []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
}

type DescribeVolumesSetItemResponse struct {
	VolumeId         string                     `xml:"volumeId,omitempty"`
	Size             string                     `xml:"size,omitempty"`
	SnapshotId       string                     `xml:"snapshotId,omitempty"`
	AvailabilityZone string                     `xml:"availabilityZone,omitempty"`
	Status           string                     `xml:"status,omitempty"`
	CreateTime       time.Time                  `xml:"createTime,omitempty"`
	AttachmentSet    *AttachmentSetItemResponse `xml:"attachmentSet,omitempty"`
	Tags             []*ResourceTagSetItem      `xml:"tagSet>item,omitempty"`
	VolumeType       string                     `xml:"volumeType,omitempty"`
	Iops             int                        `xml:"iops,omitempty"`
}

type DhcpConfigurationItem struct {
	Key      string     `xml:"key,omitempty"`
	ValueSet *DhcpValue `xml:"valueSet,omitempty"`
}

type DhcpOptions struct {
	DhcpOptionsId      string                   `xml:"dhcpOptionsId,omitempty"`
	DhcpConfigurations []*DhcpConfigurationItem `xml:"dhcpConfigurationSet>item,omitempty"`
	Tags               []*ResourceTagSetItem    `xml:"tagSet>item,omitempty"`
}

type DhcpValue struct {
	Value string `xml:"value,omitempty"`
}

type DiskImageDescription struct {
	Format            string `xml:"format,omitempty"`
	Size              int    `xml:"size,omitempty"`
	ImportManifestUrl string `xml:"importManifestUrl,omitempty"`
	Checksum          string `xml:"checksum,omitempty"`
}

type DiskImageVolumeDescription struct {
	Size int    `xml:"size,omitempty"`
	Id   string `xml:"id,omitempty"`
}

type EbsBlockDevice struct {
	SnapshotId          string `xml:"snapshotId,omitempty"`
	VolumeSize          int    `xml:"volumeSize,omitempty"`
	DeleteOnTermination bool   `xml:"deleteOnTermination,omitempty"`
	VolumeType          string `xml:"volumeType,omitempty"`
	Iops                int    `xml:"iops,omitempty"`
}

type EbsInstanceBlockDeviceMappingResponse struct {
	VolumeId            string    `xml:"volumeId,omitempty"`
	Status              string    `xml:"status,omitempty"`
	AttachTime          time.Time `xml:"attachTime,omitempty"`
	DeleteOnTermination bool      `xml:"deleteOnTermination,omitempty"`
}

type ExportTaskResponse struct {
	ExportTaskId   string                      `xml:"exportTaskId,omitempty"`
	Description    string                      `xml:"description,omitempty"`
	State          string                      `xml:"state,omitempty"`
	StatusMessage  string                      `xml:"statusMessage,omitempty"`
	InstanceExport *InstanceExportTaskResponse `xml:"instanceExport,omitempty"`
	ExportToS3     *ExportToS3TaskResponse     `xml:"exportToS3,omitempty"`
}

type ExportToS3TaskResponse struct {
	DiskImageFormat string `xml:"diskImageFormat,omitempty"`
	ContainerFormat string `xml:"containerFormat,omitempty"`
	S3Bucket        string `xml:"s3Bucket,omitempty"`
	S3Key           string `xml:"s3Key,omitempty"`
}

type GroupItem struct {
	GroupId   string `xml:"groupId,omitempty"`
	GroupName string `xml:"groupName,omitempty"`
}

type IamInstanceProfileRequest struct {
	Arn  string `xml:"arn,omitempty"`
	Name string `xml:"name,omitempty"`
}

type IamInstanceProfileResponse struct {
	Arn string `xml:"arn,omitempty"`
	Id  string `xml:"id,omitempty"`
}

type IcmpTypeCode struct {
	Code int `xml:"code,omitempty"`
	Type int `xml:"type,omitempty"`
}

type ImportInstanceTaskDetails struct {
	Volumes     []*ImportInstanceVolumeDetailItem `xml:"volumes>item,omitempty"`
	InstanceId  string                            `xml:"instanceId,omitempty"`
	Platform    string                            `xml:"platform,omitempty"`
	Description string                            `xml:"description,omitempty"`
}

type ImportInstanceVolumeDetailItem struct {
	BytesConverted   int                         `xml:"bytesConverted,omitempty"`
	AvailabilityZone string                      `xml:"availabilityZone,omitempty"`
	Image            *DiskImageDescription       `xml:"image,omitempty"`
	Description      string                      `xml:"description,omitempty"`
	Volume           *DiskImageVolumeDescription `xml:"volume,omitempty"`
	Status           string                      `xml:"status,omitempty"`
	StatusMessage    string                      `xml:"statusMessage,omitempty"`
}

type ImportVolumeTaskDetails struct {
	BytesConverted   int                         `xml:"bytesConverted,omitempty"`
	AvailabilityZone string                      `xml:"availabilityZone,omitempty"`
	Description      string                      `xml:"description,omitempty"`
	Image            *DiskImageDescription       `xml:"image,omitempty"`
	Volume           *DiskImageVolumeDescription `xml:"volume,omitempty"`
}

type InstanceBlockDeviceMappingItem struct {
	DeviceName  string                  `xml:"deviceName,omitempty"`
	VirtualName string                  `xml:"virtualName,omitempty"`
	Ebs         *InstanceEbsBlockDevice `xml:"ebs,omitempty"`
	NoDevice    struct{}                `xml:"noDevice,omitempty"`
}

type InstanceBlockDeviceMappingResponseItem struct {
	DeviceName string                                 `xml:"deviceName,omitempty"`
	Ebs        *EbsInstanceBlockDeviceMappingResponse `xml:"ebs,omitempty"`
}

type InstanceCountsSetItem struct {
	State         string `xml:"state,omitempty"`
	InstanceCount int    `xml:"instanceCount,omitempty"`
}

type InstanceCountsSet struct {
	Items []*InstanceCountsSetItem `xml:"item>item,omitempty"`
}

type InstanceEbsBlockDevice struct {
	DeleteOnTermination bool   `xml:"deleteOnTermination,omitempty"`
	VolumeId            string `xml:"volumeId,omitempty"`
}

type InstanceExportTaskResponse struct {
	InstanceId        string `xml:"instanceId,omitempty"`
	TargetEnvironment string `xml:"targetEnvironment,omitempty"`
}

type InstanceMonitoringState struct {
	State string `xml:"state,omitempty"`
}

type InstanceNetworkInterfaceAssociation struct {
	PublicIp      string `xml:"publicIp,omitempty"`
	PublicDnsName string `xml:"publicDnsName,omitempty"`
	IpOwnerId     string `xml:"ipOwnerId,omitempty"`
}

type InstanceNetworkInterfaceAttachment struct {
	AttachmentID        string    `xml:"attachmentID,omitempty"`
	DeviceIndex         int       `xml:"deviceIndex,omitempty"`
	Status              string    `xml:"status,omitempty"`
	AttachTime          time.Time `xml:"attachTime,omitempty"`
	DeleteOnTermination bool      `xml:"deleteOnTermination,omitempty"`
}

type InstanceNetworkInterfaceSetItemRequest struct {
	NetworkInterfaceId             string                            `xml:"networkInterfaceId,omitempty"`
	DeviceIndex                    int                               `xml:"deviceIndex,omitempty"`
	SubnetId                       string                            `xml:"subnetId,omitempty"`
	Description                    string                            `xml:"description,omitempty"`
	PrivateIpAddress               string                            `xml:"privateIpAddress,omitempty"`
	Groups                         []*SecurityGroupIdSetItem         `xml:"groupSet>item,omitempty"`
	DeleteOnTermination            bool                              `xml:"deleteOnTermination,omitempty"`
	PrivateIpAddressesSet          *PrivateIpAddressesSetItemRequest `xml:"privateIpAddressesSet ,omitempty"`
	SecondaryPrivateIpAddressCount int                               `xml:"secondaryPrivateIpAddressCount,omitempty"`
}

type InstanceNetworkInterfaceSetItem struct {
	NetworkInterfaceId string                               `xml:"networkInterfaceId,omitempty"`
	SubnetId           string                               `xml:"subnetId,omitempty"`
	VpcId              string                               `xml:"vpcId,omitempty"`
	Description        string                               `xml:"description,omitempty"`
	OwnerId            string                               `xml:"ownerId,omitempty"`
	Status             string                               `xml:"status,omitempty"`
	MacAddress         string                               `xml:"macAddress,omitempty"`
	PrivateIpAddress   string                               `xml:"privateIpAddress,omitempty"`
	PrivateDnsName     string                               `xml:"privateDnsName,omitempty"`
	SourceDestCheck    bool                                 `xml:"sourceDestCheck,omitempty"`
	Groups             []*GroupItem                         `xml:"groupSet>item,omitempty"`
	Attachment         *InstanceNetworkInterfaceAttachment  `xml:"attachment,omitempty"`
	Association        *InstanceNetworkInterfaceAssociation `xml:"association,omitempty"`
	PrivateIpAddresses []*InstancePrivateIpAddressesSetItem `xml:"privateIpAddressesSet>item,omitempty"`
}

type InstancePrivateIpAddressesSetItem struct {
	PrivateIpAddress string                               `xml:"privateIpAddress,omitempty"`
	PrivateDnsName   string                               `xml:"privateDnsName,omitempty"`
	Primary          bool                                 `xml:"primary,omitempty"`
	Association      *InstanceNetworkInterfaceAssociation `xml:"association,omitempty"`
}

type InstanceStateChange struct {
	InstanceId    string         `xml:"instanceId,omitempty"`
	CurrentState  *InstanceState `xml:"currentState,omitempty"`
	PreviousState *InstanceState `xml:"previousState,omitempty"`
}

type InstanceState struct {
	Code int    `xml:"code,omitempty"`
	Name string `xml:"name,omitempty"`
}

type InstanceStatusDetailsSet struct {
	Name          string    `xml:"name,omitempty"`
	Status        string    `xml:"status,omitempty"`
	ImpairedSince time.Time `xml:"impairedSince,omitempty"`
}

type InstanceStatusEventsSet struct {
	Item *InstanceStatusEvent `xml:"item,omitempty"`
}

type InstanceStatusEvent struct {
	Code        string    `xml:"code,omitempty"`
	Description string    `xml:"description,omitempty"`
	NotBefore   time.Time `xml:"notBefore,omitempty"`
	NotAfter    time.Time `xml:"notAfter,omitempty"`
}

type InstanceStatusItem struct {
	InstanceId       string                     `xml:"instanceId,omitempty"`
	AvailabilityZone string                     `xml:"availabilityZone,omitempty"`
	Events           []*InstanceStatusEventsSet `xml:"eventsSet,omitempty"`
	InstanceState    *InstanceState             `xml:"instanceState,omitempty"`
	SystemStatus     *InstanceStatus            `xml:"systemStatus,omitempty"`
	InstanceStatus   *InstanceStatus            `xml:"instanceStatus,omitempty"`
}

type InstanceStatusSet struct {
	Items []*InstanceStatusItem `xml:"item>item,omitempty"`
}

type InstanceStatus struct {
	Status  string                      `xml:"status,omitempty"`
	Details []*InstanceStatusDetailsSet `xml:"details,omitempty"`
}

type InternetGatewayAttachment struct {
	VpcId string `xml:"vpcId,omitempty"`
	State string `xml:"state,omitempty"`
}

type InternetGateway struct {
	InternetGatewayId string                     `xml:"internetGatewayId,omitempty"`
	AttachmentSet     *InternetGatewayAttachment `xml:"attachmentSet,omitempty"`
	Tags              []*ResourceTagSetItem      `xml:"tagSet>item,omitempty"`
}

type IpPermission struct {
	IpProtocol string           `xml:"ipProtocol,omitempty"`
	FromPort   int              `xml:"fromPort,omitempty"`
	ToPort     int              `xml:"toPort,omitempty"`
	Groups     *UserIdGroupPair `xml:"groups,omitempty"`
	IpRanges   []*IpRangeItem   `xml:"ipRanges>item,omitempty"`
}

type IpRangeItem struct {
	CidrIp string `xml:"cidrIp,omitempty"`
}

type LaunchPermissionItem struct {
	Group  string `xml:"group,omitempty"`
	UserId string `xml:"userId,omitempty"`
}

type LaunchSpecificationRequest struct {
	ImageId             string                                  `xml:"imageId,omitempty"`
	KeyName             string                                  `xml:"keyName,omitempty"`
	Groups              []*GroupItem                            `xml:"groupSet>item,omitempty"`
	UserData            *UserData                               `xml:"userData,omitempty"`
	InstanceType        string                                  `xml:"instanceType,omitempty"`
	Placement           *PlacementRequest                       `xml:"placement,omitempty"`
	KernelId            string                                  `xml:"kernelId,omitempty"`
	RamdiskId           string                                  `xml:"ramdiskId,omitempty"`
	BlockDeviceMappings []*BlockDeviceMappingItem               `xml:"blockDeviceMapping>item,omitempty"`
	Monitoring          *MonitoringInstance                     `xml:"monitoring,omitempty"`
	SubnetId            string                                  `xml:"subnetId,omitempty"`
	NetworkInterfaceSet *InstanceNetworkInterfaceSetItemRequest `xml:"networkInterfaceSet,omitempty"`
	IamInstanceProfile  *IamInstanceProfileRequest              `xml:"iamInstanceProfile,omitempty"`
	EbsOptimized        bool                                    `xml:"ebsOptimized,omitempty"`
}

type LaunchSpecificationResponse struct {
	ImageId             string                                  `xml:"imageId,omitempty"`
	KeyName             string                                  `xml:"keyName,omitempty"`
	Groups              []*GroupItem                            `xml:"groupSet>item,omitempty"`
	InstanceType        string                                  `xml:"instanceType,omitempty"`
	Placement           *PlacementRequest                       `xml:"placement,omitempty"`
	KernelId            string                                  `xml:"kernelId,omitempty"`
	RamdiskId           string                                  `xml:"ramdiskId,omitempty"`
	BlockDeviceMappings []*BlockDeviceMappingItem               `xml:"blockDeviceMapping>item,omitempty"`
	Monitoring          *MonitoringInstance                     `xml:"monitoring,omitempty"`
	SubnetId            string                                  `xml:"subnetId,omitempty"`
	NetworkInterfaceSet *InstanceNetworkInterfaceSetItemRequest `xml:"networkInterfaceSet,omitempty"`
	IamInstanceProfile  *IamInstanceProfileRequest              `xml:"iamInstanceProfile,omitempty"`
	EbsOptimized        bool                                    `xml:"ebsOptimized,omitempty"`
}

type MonitoringInstance struct {
	Enabled bool `xml:"enabled,omitempty"`
}

type MonitorInstancesResponseSetItem struct {
	InstanceId string                   `xml:"instanceId,omitempty"`
	Monitoring *InstanceMonitoringState `xml:"monitoring,omitempty"`
}

type NetworkAclAssociation struct {
	NetworkAclAssociationId string `xml:"networkAclAssociationId,omitempty"`
	NetworkAclId            string `xml:"networkAclId,omitempty"`
	SubnetId                string `xml:"subnetId,omitempty"`
}

type NetworkAclEntry struct {
	RuleNumber   int           `xml:"ruleNumber,omitempty"`
	Protocol     int           `xml:"protocol,omitempty"`
	RuleAction   string        `xml:"ruleAction,omitempty"`
	Egress       bool          `xml:"egress,omitempty"`
	CidrBlock    string        `xml:"cidrBlock,omitempty"`
	IcmpTypeCode *IcmpTypeCode `xml:"icmpTypeCode,omitempty"`
	PortRange    *PortRange    `xml:"portRange,omitempty"`
}

type NetworkAcl struct {
	NetworkAclId   string                 `xml:"networkAclId,omitempty"`
	VpcId          string                 `xml:"vpcId,omitempty"`
	Default        bool                   `xml:"default,omitempty"`
	EntrySet       *NetworkAclEntry       `xml:"entrySet,omitempty"`
	AssociationSet *NetworkAclAssociation `xml:"associationSet,omitempty"`
	Tags           []*ResourceTagSetItem  `xml:"tagSet>item,omitempty"`
}

type NetworkInterfaceAssociation struct {
	PublicIp      string `xml:"publicIp,omitempty"`
	PublicDnsName string `xml:"publicDnsName,omitempty"`
	IpOwnerId     string `xml:"ipOwnerId,omitempty"`
	AllocationId  string `xml:"allocationId,omitempty"`
	AssociationId string `xml:"associationId,omitempty"`
}

type NetworkInterfaceAttachment struct {
	AttachmentId        string    `xml:"attachmentId,omitempty"`
	InstanceId          string    `xml:"instanceId,omitempty"`
	InstanceOwnerId     string    `xml:"instanceOwnerId,omitempty"`
	DeviceIndex         int       `xml:"deviceIndex,omitempty"`
	Status              string    `xml:"status,omitempty"`
	AttachTime          time.Time `xml:"attachTime,omitempty"`
	DeleteOnTermination bool      `xml:"deleteOnTermination,omitempty"`
}

type NetworkInterfacePrivateIpAddressesSetItem struct {
	PrivateIpAddress string                       `xml:"privateIpAddress,omitempty"`
	PrivateDnsName   string                       `xml:"privateDnsName,omitempty"`
	Primary          bool                         `xml:"primary,omitempty"`
	Association      *NetworkInterfaceAssociation `xml:"association,omitempty"`
}

type NetworkInterface struct {
	NetworkInterfaceId string                                       `xml:"networkInterfaceId,omitempty"`
	SubnetId           string                                       `xml:"subnetId,omitempty"`
	VpcId              string                                       `xml:"vpcId,omitempty"`
	AvailabilityZone   string                                       `xml:"availabilityZone,omitempty"`
	Description        string                                       `xml:"description,omitempty"`
	OwnerId            string                                       `xml:"ownerId,omitempty"`
	RequesterId        string                                       `xml:"requesterId,omitempty"`
	RequesterManaged   string                                       `xml:"requesterManaged,omitempty"`
	Status             string                                       `xml:"status,omitempty"`
	MacAddress         string                                       `xml:"macAddress,omitempty"`
	PrivateIpAddress   string                                       `xml:"privateIpAddress,omitempty"`
	PrivateDnsName     string                                       `xml:"privateDnsName,omitempty"`
	SourceDestCheck    bool                                         `xml:"sourceDestCheck,omitempty"`
	Groups             []*GroupItem                                 `xml:"groupSet>item,omitempty"`
	Attachment         *NetworkInterfaceAttachment                  `xml:"attachment,omitempty"`
	Association        *NetworkInterfaceAssociation                 `xml:"association,omitempty"`
	Tags               []*ResourceTagSetItem                        `xml:"tagSet>item,omitempty"`
	PrivateIpAddresses []*NetworkInterfacePrivateIpAddressesSetItem `xml:"privateIpAddressesSet>item,omitempty"`
}

type PlacementGroupInfo struct {
	GroupName string `xml:"groupName,omitempty"`
	Strategy  string `xml:"strategy,omitempty"`
	State     string `xml:"state,omitempty"`
}

type PlacementRequest struct {
	AvailabilityZone string `xml:"availabilityZone,omitempty"`
	GroupName        string `xml:"groupName,omitempty"`
}

type PlacementResponse struct {
	AvailabilityZone string `xml:"availabilityZone,omitempty"`
	GroupName        string `xml:"groupName,omitempty"`
	Tenancy          string `xml:"tenancy,omitempty"`
}

type PortRange struct {
	From int `xml:"from,omitempty"`
	To   int `xml:"to,omitempty"`
}

type PriceScheduleRequestSetItem struct {
	Term         int     `xml:"term,omitempty"`
	Price        float64 `xml:"price,omitempty"`
	CurrencyCode string  `xml:"currencyCode,omitempty"`
}

type PriceScheduleSetItem struct {
	Term         int     `xml:"term,omitempty"`
	Price        float64 `xml:"price,omitempty"`
	CurrencyCode string  `xml:"currencyCode,omitempty"`
	Active       bool    `xml:"active,omitempty"`
}

type PriceScheduleSet struct {
	Items []*PriceScheduleSetItem `xml:"item>item,omitempty"`
}

type PricingDetailsSetItem struct {
	Price int `xml:"price,omitempty"`
	Count int `xml:"count,omitempty"`
}

type PrivateIpAddressesSetItemRequest struct {
	PrivateIpAddressesSet *AssignPrivateIpAddressesSetItemRequest `xml:"privateIpAddressesSet,omitempty"`
	Primary               bool                                    `xml:"primary,omitempty"`
}

type ProductCodeItem struct {
	ProductCode string `xml:"productCode,omitempty"`
}

type ProductCodesSetItem struct {
	ProductCode string `xml:"productCode,omitempty"`
	Type        string `xml:"type,omitempty"`
}

type ProductDescriptionSetItem struct {
	ProductDescription string `xml:"productDescription,omitempty"`
}

type PropagatingVgw struct {
	GatewayID string `xml:"gatewayID,omitempty"`
}

type RecurringChargesSetItem struct {
	Frequency string  `xml:"frequency,omitempty"`
	Amount    float64 `xml:"amount,omitempty"`
}

type RegionItem struct {
	RegionName     string `xml:"regionName,omitempty"`
	RegionEndpoint string `xml:"regionEndpoint,omitempty"`
}

type ReservationInfo struct {
	ReservationId string                  `xml:"reservationId,omitempty"`
	OwnerId       string                  `xml:"ownerId,omitempty"`
	Groups        []*GroupItem            `xml:"groupSet>item,omitempty"`
	Instances     []*RunningInstancesItem `xml:"instancesSet>item,omitempty"`
	RequesterId   string                  `xml:"requesterId,omitempty"`
}

type ReservedInstanceLimitPrice struct {
	Amount       float64 `xml:"amount,omitempty"`
	CurrencyCode float64 `xml:"currencyCode,omitempty"`
}

type ReservedInstancesConfigurationSetItem struct {
	AvailabilityZone string `xml:"availabilityZone,omitempty"`
	Platform         string `xml:"platform,omitempty"`
	InstanceCount    int    `xml:"instanceCount,omitempty"`
	InstanceType     string `xml:"instanceType,omitempty"`
}

type ReservedInstancesModificationResultSetItem struct {
	ReservedInstancesId  string                                   `xml:"reservedInstancesId,omitempty"`
	TargetConfigurations []*ReservedInstancesConfigurationSetItem `xml:"targetConfiguration>item,omitempty"`
}

type ResourceTagSetItem struct {
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

type RouteTableAssociation struct {
	RouteTableAssociationId string `xml:"routeTableAssociationId,omitempty"`
	RouteTableId            string `xml:"routeTableId,omitempty"`
	SubnetId                string `xml:"subnetId,omitempty"`
	Main                    bool   `xml:"main,omitempty"`
}

type RouteTable struct {
	RouteTableId      string                 `xml:"routeTableId,omitempty"`
	VpcId             string                 `xml:"vpcId,omitempty"`
	RouteSet          *Route                 `xml:"routeSet,omitempty"`
	AssociationSet    *RouteTableAssociation `xml:"associationSet,omitempty"`
	PropagatingVgwSet *PropagatingVgw        `xml:"propagatingVgwSet,omitempty"`
	Tags              []*ResourceTagSetItem  `xml:"tagSet>item,omitempty"`
}

type Route struct {
	DestinationCidrBlock string `xml:"destinationCidrBlock,omitempty"`
	GatewayId            string `xml:"gatewayId,omitempty"`
	InstanceId           string `xml:"instanceId,omitempty"`
	InstanceOwnerId      string `xml:"instanceOwnerId,omitempty"`
	NetworkInterfaceId   string `xml:"networkInterfaceId,omitempty"`
	State                string `xml:"state,omitempty"`
	Origin               string `xml:"origin,omitempty"`
}

type RunningInstancesItem struct {
	InstanceId            string                                    `xml:"instanceId,omitempty"`
	ImageId               string                                    `xml:"imageId,omitempty"`
	InstanceState         *InstanceState                            `xml:"instanceState,omitempty"`
	PrivateDnsName        string                                    `xml:"privateDnsName,omitempty"`
	DnsName               string                                    `xml:"dnsName,omitempty"`
	Reason                string                                    `xml:"reason,omitempty"`
	KeyName               string                                    `xml:"keyName,omitempty"`
	AmiLaunchIndex        string                                    `xml:"amiLaunchIndex,omitempty"`
	ProductCodes          []*ProductCodesSetItem                    `xml:"productCodes>item,omitempty"`
	InstanceType          string                                    `xml:"instanceType,omitempty"`
	LaunchTime            time.Time                                 `xml:"launchTime,omitempty"`
	Placement             *PlacementResponse                        `xml:"placement,omitempty"`
	KernelId              string                                    `xml:"kernelId,omitempty"`
	RamdiskId             string                                    `xml:"ramdiskId,omitempty"`
	Platform              string                                    `xml:"platform,omitempty"`
	Monitoring            *InstanceMonitoringState                  `xml:"monitoring,omitempty"`
	SubnetId              string                                    `xml:"subnetId,omitempty"`
	VpcId                 string                                    `xml:"vpcId,omitempty"`
	PrivateIpAddress      string                                    `xml:"privateIpAddress,omitempty"`
	IpAddress             string                                    `xml:"ipAddress,omitempty"`
	SourceDestCheck       bool                                      `xml:"sourceDestCheck,omitempty"`
	Groups                []*GroupItem                              `xml:"groupSet>item,omitempty"`
	StateReason           *StateReason                              `xml:"stateReason,omitempty"`
	Architecture          string                                    `xml:"architecture,omitempty"`
	RootDeviceType        string                                    `xml:"rootDeviceType,omitempty"`
	RootDeviceName        string                                    `xml:"rootDeviceName,omitempty"`
	BlockDeviceMappings   []*InstanceBlockDeviceMappingResponseItem `xml:"blockDeviceMapping>item,omitempty"`
	InstanceLifecycle     string                                    `xml:"instanceLifecycle,omitempty"`
	SpotInstanceRequestId string                                    `xml:"spotInstanceRequestId,omitempty"`
	VirtualizationType    string                                    `xml:"virtualizationType,omitempty"`
	ClientToken           string                                    `xml:"clientToken,omitempty"`
	Tags                  []*ResourceTagSetItem                     `xml:"tagSet>item,omitempty"`
	Hypervisor            string                                    `xml:"hypervisor,omitempty"`
	NetworkInterfaces     []*InstanceNetworkInterfaceSetItem        `xml:"networkInterfaceSet>item,omitempty"`
	IamInstanceProfile    *IamInstanceProfileResponse               `xml:"iamInstanceProfile,omitempty"`
	EbsOptimized          bool                                      `xml:"ebsOptimized,omitempty"`
	SriovNetSupport       string                                    `xml:"sriovNetSupport,omitempty"`
}

type SecurityGroupIdSetItem struct {
	GroupId string `xml:"groupId,omitempty"`
}

type SecurityGroupItem struct {
	OwnerId             string                `xml:"ownerId,omitempty"`
	GroupId             string                `xml:"groupId,omitempty"`
	GroupName           string                `xml:"groupName,omitempty"`
	GroupDescription    string                `xml:"groupDescription,omitempty"`
	VpcId               string                `xml:"vpcId,omitempty"`
	IpPermissions       *IpPermission         `xml:"ipPermissions,omitempty"`
	IpPermissionsEgress *IpPermission         `xml:"ipPermissionsEgress,omitempty"`
	Tags                []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
}

type SpotDatafeedSubscription struct {
	OwnerId string                  `xml:"ownerId,omitempty"`
	Bucket  string                  `xml:"bucket,omitempty"`
	Prefix  string                  `xml:"prefix,omitempty"`
	State   string                  `xml:"state,omitempty"`
	Fault   *SpotInstanceStateFault `xml:"fault,omitempty"`
}

type SpotInstanceRequestSetItem struct {
	SpotInstanceRequestId    string                       `xml:"spotInstanceRequestId,omitempty"`
	SpotPrice                string                       `xml:"spotPrice,omitempty"`
	Type                     string                       `xml:"type,omitempty"`
	State                    string                       `xml:"state,omitempty"`
	Fault                    *SpotInstanceStateFault      `xml:"fault,omitempty"`
	Status                   *SpotInstanceStatusMessage   `xml:"status,omitempty"`
	ValidFrom                time.Time                    `xml:"validFrom,omitempty"`
	ValidUntil               time.Time                    `xml:"validUntil,omitempty"`
	LaunchGroup              string                       `xml:"launchGroup,omitempty"`
	AvailabilityZoneGroup    string                       `xml:"availabilityZoneGroup,omitempty"`
	LaunchedAvailabilityZone string                       `xml:"launchedAvailabilityZone,omitempty"`
	LaunchSpecification      *LaunchSpecificationResponse `xml:"launchSpecification,omitempty"`
	InstanceId               string                       `xml:"instanceId,omitempty"`
	CreateTime               time.Time                    `xml:"createTime,omitempty"`
	ProductDescription       string                       `xml:"productDescription,omitempty"`
	Tags                     []*ResourceTagSetItem        `xml:"tagSet>item,omitempty"`
}

type SpotInstanceStateFault struct {
	Code    string `xml:"code,omitempty"`
	Message string `xml:"message,omitempty"`
}

type SpotInstanceStatusMessage struct {
	Code       string    `xml:"code,omitempty"`
	UpdateTime time.Time `xml:"updateTime,omitempty"`
	Message    string    `xml:"message,omitempty"`
}

type SpotPriceHistorySetItem struct {
	InstanceType       string    `xml:"instanceType,omitempty"`
	ProductDescription string    `xml:"productDescription,omitempty"`
	SpotPrice          string    `xml:"spotPrice,omitempty"`
	Timestamp          time.Time `xml:"timestamp,omitempty"`
	AvailabilityZone   string    `xml:"availabilityZone,omitempty"`
}

type StateReason struct {
	Code    string `xml:"code,omitempty"`
	Message string `xml:"message,omitempty"`
}

type Subnet struct {
	SubnetId                string                `xml:"subnetId,omitempty"`
	State                   string                `xml:"state,omitempty"`
	VpcId                   string                `xml:"vpcId,omitempty"`
	CidrBlock               string                `xml:"cidrBlock,omitempty"`
	AvailableIpAddressCount int                   `xml:"availableIpAddressCount,omitempty"`
	AvailabilityZone        string                `xml:"availabilityZone,omitempty"`
	DefaultForAz            bool                  `xml:"defaultForAz,omitempty"`
	MapPublicIpOnLaunch     bool                  `xml:"mapPublicIpOnLaunch,omitempty"`
	Tags                    []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
}

type TagSetItem struct {
	ResourceId   string `xml:"resourceId,omitempty"`
	ResourceType string `xml:"resourceType,omitempty"`
	Key          string `xml:"key,omitempty"`
	Value        string `xml:"value,omitempty"`
}

type UserData struct {
	Data string `xml:"data,omitempty"`
}

type UserIdGroupPair struct {
	UserId    string `xml:"userId,omitempty"`
	GroupId   string `xml:"groupId,omitempty"`
	GroupName string `xml:"groupName,omitempty"`
}

type VolumeStatusItem struct {
	VolumeId         string                    `xml:"volumeId,omitempty"`
	AvailabilityZone string                    `xml:"availabilityZone,omitempty"`
	VolumeStatus     *VolumeStatusInfo         `xml:"volumeStatus,omitempty"`
	Events           []*VolumeStatusEventItem  `xml:"eventSet>item,omitempty"`
	Actions          []*VolumeStatusActionItem `xml:"actionSet>item,omitempty"`
}

type VolumeStatusInfo struct {
	Status  string                     `xml:"status,omitempty"`
	Details []*VolumeStatusDetailsItem `xml:"details>item,omitempty"`
}

type VolumeStatusDetailsItem struct {
	Name   string `xml:"name,omitempty"`
	Status string `xml:"status,omitempty"`
}

type VolumeStatusEventItem struct {
	EventType   string    `xml:"eventType,omitempty"`
	EventId     string    `xml:"eventId,omitempty"`
	Description string    `xml:"description,omitempty"`
	NotBefore   time.Time `xml:"notBefore,omitempty"`
	NotAfter    time.Time `xml:"notAfter,omitempty"`
}

type VolumeStatusActionItem struct {
	Code        string `xml:"code,omitempty"`
	EventType   string `xml:"eventType,omitempty"`
	EventId     string `xml:"eventId,omitempty"`
	Description string `xml:"description,omitempty"`
}

type Vpc struct {
	VpcId           string                `xml:"vpcId,omitempty"`
	State           string                `xml:"state,omitempty"`
	CidrBlock       string                `xml:"cidrBlock,omitempty"`
	DhcpOptionsId   string                `xml:"dhcpOptionsId,omitempty"`
	Tags            []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
	InstanceTenancy string                `xml:"instanceTenancy,omitempty"`
	IsDefault       bool                  `xml:"isDefault,omitempty"`
}

type VpnConnectionOptionsResponse struct {
	StaticRoutesOnly bool `xml:"staticRoutesOnly,omitempty"`
}

type VpnConnection struct {
	VpnConnectionId              string                        `xml:"vpnConnectionId,omitempty"`
	State                        string                        `xml:"state,omitempty"`
	CustomerGatewayConfiguration string                        `xml:"customerGatewayConfiguration,omitempty"`
	Type                         string                        `xml:"type,omitempty"`
	CustomerGatewayId            string                        `xml:"customerGatewayId,omitempty"`
	VpnGatewayId                 string                        `xml:"vpnGatewayId,omitempty"`
	Tags                         []*ResourceTagSetItem         `xml:"tagSet>item,omitempty"`
	VgwTelemetry                 *VpnTunnelTelemetry           `xml:"vgwTelemetry,omitempty"`
	Options                      *VpnConnectionOptionsResponse `xml:"options,omitempty"`
	Routes                       *VpnStaticRoute               `xml:"routes,omitempty"`
}

type VpnGateway struct {
	VpnGatewayId     string                `xml:"vpnGatewayId,omitempty"`
	State            string                `xml:"state,omitempty"`
	Type             string                `xml:"type,omitempty"`
	AvailabilityZone string                `xml:"availabilityZone,omitempty"`
	Attachments      *Attachment           `xml:"attachments,omitempty"`
	Tags             []*ResourceTagSetItem `xml:"tagSet>item,omitempty"`
}

type VpnStaticRoute struct {
	DestinationCidrBlock string `xml:"destinationCidrBlock,omitempty"`
	Source               string `xml:"source,omitempty"`
	State                string `xml:"state,omitempty"`
}

type VpnTunnelTelemetry struct {
	OutsideIpAddress   string    `xml:"outsideIpAddress,omitempty"`
	Status             string    `xml:"status,omitempty"`
	LastStatusChange   time.Time `xml:"lastStatusChange,omitempty"`
	StatusMessage      string    `xml:"statusMessage,omitempty"`
	AcceptedRouteCount int       `xml:"acceptedRouteCount,omitempty"`
}
