package main

type AutoScalingAutoScalingGroup struct {
	AvailabilityZones         interface{} `json:"AvailabilityZones,omitempty"`
	Cooldown                  interface{} `json:"Cooldown,omitempty"`
	DesiredCapacity           interface{} `json:"DesiredCapacity,omitempty"`
	HealthCheckGracePeriod    interface{} `json:"HealthCheckGracePeriod,omitempty"`
	HealthCheckType           interface{} `json:"HealthCheckType,omitempty"`
	InstanceId                interface{} `json:"InstanceId,omitempty"`
	LaunchConfigurationName   interface{} `json:"LaunchConfigurationName,omitempty"`
	LoadBalancerNames         interface{} `json:"LoadBalancerNames,omitempty"`
	MaxSize                   interface{} `json:"MaxSize,omitempty"`
	MinSize                   interface{} `json:"MinSize,omitempty"`
	NotificationConfiguration interface{} `json:"NotificationConfiguration,omitempty"`
	Tags                      interface{} `json:"Tags,omitempty"`
	TerminationPolicies       interface{} `json:"TerminationPolicies,omitempty"`
	VPCZoneIdentifier         interface{} `json:"VPCZoneIdentifier,omitempty"`
}

type AutoScalingLaunchConfiguration struct {
	AssociatePublicIpAddress interface{} `json:"AssociatePublicIpAddress,omitempty"`
	BlockDeviceMappings      interface{} `json:"BlockDeviceMappings,omitempty"`
	EbsOptimized             interface{} `json:"EbsOptimized,omitempty"`
	IamInstanceProfile       interface{} `json:"IamInstanceProfile,omitempty"`
	ImageId                  interface{} `json:"ImageId,omitempty"`
	InstanceId               interface{} `json:"InstanceId,omitempty"`
	InstanceMonitoring       interface{} `json:"InstanceMonitoring,omitempty"`
	InstanceType             interface{} `json:"InstanceType,omitempty"`
	KernelId                 interface{} `json:"KernelId,omitempty"`
	KeyName                  interface{} `json:"KeyName,omitempty"`
	RamDiskId                interface{} `json:"RamDiskId,omitempty"`
	SecurityGroups           interface{} `json:"SecurityGroups,omitempty"`
	SpotPrice                interface{} `json:"SpotPrice,omitempty"`
	UserData                 interface{} `json:"UserData,omitempty"`
}

type AutoScalingScalingPolicy struct {
	AdjustmentType       interface{} `json:"AdjustmentType,omitempty"`
	AutoScalingGroupName interface{} `json:"AutoScalingGroupName,omitempty"`
	Cooldown             interface{} `json:"Cooldown,omitempty"`
	ScalingAdjustment    interface{} `json:"ScalingAdjustment,omitempty"`
}

type AutoScalingTrigger struct {
	AutoScalingGroupName      interface{} `json:"AutoScalingGroupName,omitempty"`
	BreachDuration            interface{} `json:"BreachDuration,omitempty"`
	Dimensions                interface{} `json:"Dimensions,omitempty"`
	LowerBreachScaleIncrement interface{} `json:"LowerBreachScaleIncrement,omitempty"`
	LowerThreshold            interface{} `json:"LowerThreshold,omitempty"`
	MetricName                interface{} `json:"MetricName,omitempty"`
	Namespace                 interface{} `json:"Namespace,omitempty"`
	Period                    interface{} `json:"Period,omitempty"`
	Statistic                 interface{} `json:"Statistic,omitempty"`
	Unit                      interface{} `json:"Unit,omitempty"`
	UpperBreachScaleIncrement interface{} `json:"UpperBreachScaleIncrement,omitempty"`
	UpperThreshold            interface{} `json:"UpperThreshold,omitempty"`
}

type CloudFormationAuthentication struct {
	AccessKeyId interface{} `json:"accessKeyId,omitempty"`
	Buckets     interface{} `json:"buckets,omitempty"`
	Password    interface{} `json:"password,omitempty"`
	SecretKey   interface{} `json:"secretKey,omitempty"`
	Type        interface{} `json:"type,omitempty"`
	Uris        interface{} `json:"uris,omitempty"`
	Username    interface{} `json:"username,omitempty"`
	RoleName    interface{} `json:"roleName,omitempty"`
}

type CloudFormationCustomResource struct {
	ServiceToken interface{} `json:"ServiceToken,omitempty"`
}

type CloudFormationInit struct {
}

type CloudFormationStack struct {
	TemplateURL      interface{} `json:"TemplateURL,omitempty"`
	TimeoutInMinutes interface{} `json:"TimeoutInMinutes,omitempty"`
	Parameters       interface{} `json:"Parameters,omitempty"`
}

type CloudFormationWaitCondition struct {
	Count   interface{} `json:"Count,omitempty"`
	Handle  interface{} `json:"Handle,omitempty"`
	Timeout interface{} `json:"Timeout,omitempty"`
}

type CloudFormationWaitConditionHandle struct {
}

type CloudFrontDistribution struct {
	DistributionConfig interface{} `json:"DistributionConfig,omitempty"`
}

type CloudWatchAlarm struct {
}

type DynamoDBTable struct {
	AttributeDefinitions   interface{} `json:"AttributeDefinitions,omitempty"`
	GlobalSecondaryIndexes interface{} `json:"GlobalSecondaryIndexes,omitempty"`
	KeySchema              interface{} `json:"KeySchema,omitempty"`
	LocalSecondaryIndexes  interface{} `json:"LocalSecondaryIndexes,omitempty"`
	ProvisionedThroughput  interface{} `json:"ProvisionedThroughput,omitempty"`
	TableName              interface{} `json:"TableName,omitempty"`
}

type EC2CustomerGateway struct {
	BgpAsn    interface{} `json:"BgpAsn,omitempty"`
	IpAddress interface{} `json:"IpAddress,omitempty"`
	Tags      interface{} `json:"Tags,omitempty"`
	Type      interface{} `json:"Type,omitempty"`
}

type EC2DHCPOptions struct {
	DomainName         interface{} `json:"DomainName,omitempty"`
	DomainNameServers  interface{} `json:"DomainNameServers,omitempty"`
	NetbiosNameServers interface{} `json:"NetbiosNameServers,omitempty"`
	NetbiosNodeType    interface{} `json:"NetbiosNodeType,omitempty"`
	NtpServers         interface{} `json:"NtpServers,omitempty"`
	Tags               interface{} `json:"Tags,omitempty"`
}

type EC2EIP struct {
	InstanceId interface{} `json:"InstanceId,omitempty"`
	Domain     interface{} `json:"Domain,omitempty"`
}

type EC2EIPAssociation struct {
	AllocationId       interface{} `json:"AllocationId,omitempty"`
	EIP                interface{} `json:"EIP,omitempty"`
	InstanceId         interface{} `json:"InstanceId,omitempty"`
	NetworkInterfaceId interface{} `json:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   interface{} `json:"PrivateIpAddress,omitempty"`
}

type EC2Instance struct {
	AvailabilityZone      interface{} `json:"AvailabilityZone,omitempty"`
	BlockDeviceMappings   interface{} `json:"BlockDeviceMappings,omitempty"`
	DisableApiTermination interface{} `json:"DisableApiTermination,omitempty"`
	EbsOptimized          interface{} `json:"EbsOptimized,omitempty"`
	IamInstanceProfile    interface{} `json:"IamInstanceProfile,omitempty"`
	ImageId               interface{} `json:"ImageId,omitempty"`
	InstanceType          interface{} `json:"InstanceType,omitempty"`
	KernelId              interface{} `json:"KernelId,omitempty"`
	KeyName               interface{} `json:"KeyName,omitempty"`
	Monitoring            interface{} `json:"Monitoring,omitempty"`
	NetworkInterfaces     interface{} `json:"NetworkInterfaces,omitempty"`
	PlacementGroupName    interface{} `json:"PlacementGroupName,omitempty"`
	PrivateIpAddress      interface{} `json:"PrivateIpAddress,omitempty"`
	RamdiskId             interface{} `json:"RamdiskId,omitempty"`
	SecurityGroupIds      interface{} `json:"SecurityGroupIds,omitempty"`
	SecurityGroups        interface{} `json:"SecurityGroups,omitempty"`
	SourceDestCheck       interface{} `json:"SourceDestCheck,omitempty"`
	SubnetId              interface{} `json:"SubnetId,omitempty"`
	Tags                  interface{} `json:"Tags,omitempty"`
	Tenancy               interface{} `json:"Tenancy,omitempty"`
	UserData              interface{} `json:"UserData,omitempty"`
	Volumes               interface{} `json:"Volumes,omitempty"`
}

type EC2InternetGateway struct {
	Tags interface{} `json:"Tags,omitempty"`
}

type EC2NetworkAcl struct {
	Tags  interface{} `json:"Tags,omitempty"`
	VpcId interface{} `json:"VpcId,omitempty"`
}

type EC2NetworkAclEntry struct {
	CidrBlock    interface{} `json:"CidrBlock,omitempty"`
	Egress       interface{} `json:"Egress,omitempty"`
	Icmp         interface{} `json:"Icmp,omitempty"`
	NetworkAclId interface{} `json:"NetworkAclId,omitempty"`
	PortRange    interface{} `json:"PortRange,omitempty"`
	Protocol     interface{} `json:"Protocol,omitempty"`
	RuleAction   interface{} `json:"RuleAction,omitempty"`
	RuleNumber   interface{} `json:"RuleNumber,omitempty"`
}

type EC2NetworkInterface struct {
	Description                    interface{} `json:"Description,omitempty"`
	GroupSet                       interface{} `json:"GroupSet,omitempty"`
	PrivateIpAddress               interface{} `json:"PrivateIpAddress,omitempty"`
	PrivateIpAddresses             interface{} `json:"PrivateIpAddresses,omitempty"`
	SecondaryPrivateIpAddressCount interface{} `json:"SecondaryPrivateIpAddressCount,omitempty"`
	SourceDestCheck                interface{} `json:"SourceDestCheck,omitempty"`
	SubnetId                       interface{} `json:"SubnetId,omitempty"`
	Tags                           interface{} `json:"Tags,omitempty"`
}

type EC2NetworkInterfaceAttachment struct {
	DeleteOnTermination interface{} `json:"DeleteOnTermination,omitempty"`
	DeviceIndex         interface{} `json:"DeviceIndex,omitempty"`
	InstanceId          interface{} `json:"InstanceId,omitempty"`
	NetworkInterfaceId  interface{} `json:"NetworkInterfaceId,omitempty"`
}

type EC2Route struct {
	DestinationCidrBlock interface{} `json:"DestinationCidrBlock,omitempty"`
	GatewayId            interface{} `json:"GatewayId,omitempty"`
	InstanceId           interface{} `json:"InstanceId,omitempty"`
	NetworkInterfaceId   interface{} `json:"NetworkInterfaceId,omitempty"`
	RouteTableId         interface{} `json:"RouteTableId,omitempty"`
}

type EC2RouteTable struct {
	VpcId interface{} `json:"VpcId,omitempty"`
	Tags  interface{} `json:"Tags,omitempty"`
}

type EC2SecurityGroup struct {
	GroupDescription     interface{} `json:"GroupDescription,omitempty"`
	SecurityGroupEgress  interface{} `json:"SecurityGroupEgress,omitempty"`
	SecurityGroupIngress interface{} `json:"SecurityGroupIngress,omitempty"`
	Tags                 interface{} `json:"Tags,omitempty"`
	VpcId                interface{} `json:"VpcId,omitempty"`
}

type EC2SecurityGroupIngress struct {
	GroupName                  interface{} `json:"GroupName,omitempty"`
	GroupId                    interface{} `json:"GroupId,omitempty"`
	IpProtocol                 interface{} `json:"IpProtocol,omitempty"`
	CidrIp                     interface{} `json:"CidrIp,omitempty"`
	SourceSecurityGroupName    interface{} `json:"SourceSecurityGroupName,omitempty"`
	SourceSecurityGroupId      interface{} `json:"SourceSecurityGroupId,omitempty"`
	SourceSecurityGroupOwnerId interface{} `json:"SourceSecurityGroupOwnerId,omitempty"`
	FromPort                   interface{} `json:"FromPort,omitempty"`
	ToPort                     interface{} `json:"ToPort,omitempty"`
}

type EC2SecurityGroupEgress struct {
	CidrIp                     interface{} `json:"CidrIp,omitempty"`
	DestinationSecurityGroupId interface{} `json:"DestinationSecurityGroupId,omitempty"`
	FromPort                   interface{} `json:"FromPort,omitempty"`
	GroupId                    interface{} `json:"GroupId,omitempty"`
	IpProtocol                 interface{} `json:"IpProtocol,omitempty"`
	ToPort                     interface{} `json:"ToPort,omitempty"`
}

type EC2Subnet struct {
	AvailabilityZone interface{} `json:"AvailabilityZone,omitempty"`
	CidrBlock        interface{} `json:"CidrBlock,omitempty"`
	Tags             interface{} `json:"Tags,omitempty"`
	VpcId            interface{} `json:"VpcId,omitempty"`
}

type EC2SubnetNetworkAclAssociation struct {
	SubnetId     interface{} `json:"SubnetId,omitempty"`
	NetworkAclId interface{} `json:"NetworkAclId,omitempty"`
}

type EC2SubnetRouteTableAssociation struct {
	RouteTableId interface{} `json:"RouteTableId,omitempty"`
	SubnetId     interface{} `json:"SubnetId,omitempty"`
}

type EC2Volume struct {
	AvailabilityZone interface{} `json:"AvailabilityZone,omitempty"`
	Iops             interface{} `json:"Iops,omitempty"`
	Size             interface{} `json:"Size,omitempty"`
	SnapshotId       interface{} `json:"SnapshotId,omitempty"`
	Tags             interface{} `json:"Tags,omitempty"`
	VolumeType       interface{} `json:"VolumeType,omitempty"`
}

type EC2VolumeAttachment struct {
	Device     interface{} `json:"Device,omitempty"`
	InstanceId interface{} `json:"InstanceId,omitempty"`
	VolumeId   interface{} `json:"VolumeId,omitempty"`
}

type EC2VPC struct {
	CidrBlock          interface{} `json:"CidrBlock,omitempty"`
	EnableDnsSupport   interface{} `json:"EnableDnsSupport,omitempty"`
	EnableDnsHostnames interface{} `json:"EnableDnsHostnames,omitempty"`
	InstanceTenancy    interface{} `json:"InstanceTenancy,omitempty"`
	Tags               interface{} `json:"Tags,omitempty"`
}

type EC2VPCDHCPOptionsAssociation struct {
	DhcpOptionsId interface{} `json:"DhcpOptionsId,omitempty"`
	VpcId         interface{} `json:"VpcId,omitempty"`
}

type EC2VPCGatewayAttachment struct {
	InternetGatewayId interface{} `json:"InternetGatewayId,omitempty"`
	VpcId             interface{} `json:"VpcId,omitempty"`
	VpnGatewayId      interface{} `json:"VpnGatewayId,omitempty"`
}

type EC2VPNConnection struct {
	Type              interface{} `json:"Type,omitempty"`
	CustomerGatewayId interface{} `json:"CustomerGatewayId,omitempty"`
	StaticRoutesOnly  interface{} `json:"StaticRoutesOnly,omitempty"`
	Tags              interface{} `json:"Tags,omitempty"`
	VpnGatewayId      interface{} `json:"VpnGatewayId,omitempty"`
}

type EC2VPNConnectionRoute struct {
	DestinationCidrBlock interface{} `json:"DestinationCidrBlock,omitempty"`
	VpnConnectionId      interface{} `json:"VpnConnectionId,omitempty"`
}

type EC2VPNGateway struct {
	Type interface{} `json:"Type,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

type EC2VPNGatewayRoutePropagation struct {
	RouteTableIds interface{} `json:"RouteTableIds,omitempty"`
	VpnGatewayId  interface{} `json:"VpnGatewayId,omitempty"`
}

type ElastiCacheCacheCluster struct {
	AutoMinorVersionUpgrade    interface{} `json:"AutoMinorVersionUpgrade,omitempty"`
	CacheNodeType              interface{} `json:"CacheNodeType,omitempty"`
	CacheParameterGroupName    interface{} `json:"CacheParameterGroupName,omitempty"`
	CacheSecurityGroupNames    interface{} `json:"CacheSecurityGroupNames,omitempty"`
	CacheSubnetGroupName       interface{} `json:"CacheSubnetGroupName,omitempty"`
	ClusterName                interface{} `json:"ClusterName,omitempty"`
	Engine                     interface{} `json:"Engine,omitempty"`
	EngineVersion              interface{} `json:"EngineVersion,omitempty"`
	NotificationTopicArn       interface{} `json:"NotificationTopicArn,omitempty"`
	NumCacheNodes              interface{} `json:"NumCacheNodes,omitempty"`
	Port                       interface{} `json:"Port,omitempty"`
	PreferredAvailabilityZone  interface{} `json:"PreferredAvailabilityZone,omitempty"`
	PreferredMaintenanceWindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	SnapshotArns               interface{} `json:"SnapshotArns,omitempty"`
	VpcSecurityGroupIds        interface{} `json:" VpcSecurityGroupIds ,omitempty"`
}

type ElastiCacheParameterGroup struct {
	CacheParameterGroupFamily interface{} `json:"CacheParameterGroupFamily,omitempty"`
	Description               interface{} `json:"Description,omitempty"`
	Properties                interface{} `json:"Properties,omitempty"`
}

type ElastiCacheSecurityGroup struct {
}

type ElastiCacheSecurityGroupIngress struct {
}

type ElastiCacheSubnetGroup struct {
	Description interface{} `json:"Description,omitempty"`
	SubnetIds   interface{} `json:"SubnetIds,omitempty"`
}

type ElasticBeanstalkApplication struct {
	ApplicationName interface{} `json:"ApplicationName,omitempty"`
	Description     interface{} `json:"Description,omitempty"`
}

type ElasticBeanstalkEnvironment struct {
	ApplicationName   interface{} `json:"ApplicationName,omitempty"`
	CNAMEPrefix       interface{} `json:"CNAMEPrefix,omitempty"`
	Description       interface{} `json:"Description,omitempty"`
	EnvironmentName   interface{} `json:"EnvironmentName,omitempty"`
	OptionSettings    interface{} `json:"OptionSettings,omitempty"`
	SolutionStackName interface{} `json:"SolutionStackName,omitempty"`
	TemplateName      interface{} `json:"TemplateName,omitempty"`
	Tier              interface{} `json:"Tier,omitempty"`
	VersionLabel      interface{} `json:"VersionLabel,omitempty"`
}

type ElasticLoadBalancingLoadBalancer struct {
	AppCookieStickinessPolicy interface{} `json:"AppCookieStickinessPolicy,omitempty"`
	AvailabilityZones         interface{} `json:"AvailabilityZones,omitempty"`
	CrossZone                 interface{} `json:"CrossZone,omitempty"`
	HealthCheck               interface{} `json:"HealthCheck,omitempty"`
	Instances                 interface{} `json:"Instances,omitempty"`
	LBCookieStickinessPolicy  interface{} `json:"LBCookieStickinessPolicy,omitempty"`
	LoadBalancerName          interface{} `json:"LoadBalancerName,omitempty"`
	Listeners                 interface{} `json:"Listeners,omitempty"`
	Policies                  interface{} `json:"Policies,omitempty"`
	Scheme                    interface{} `json:"Scheme,omitempty"`
	SecurityGroups            interface{} `json:"SecurityGroups,omitempty"`
	Subnets                   interface{} `json:"Subnets,omitempty"`
}

type IAMAccessKey struct {
	Serial   interface{} `json:"Serial,omitempty"`
	Status   interface{} `json:"Status,omitempty"`
	UserName interface{} `json:"UserName,omitempty"`
}

type IAMGroup struct {
	Path     interface{} `json:"Path,omitempty"`
	Policies interface{} `json:"Policies,omitempty"`
}

type IAMInstanceProfile struct {
	Path  interface{} `json:"Path,omitempty"`
	Roles interface{} `json:"Roles,omitempty"`
}

type IAMPolicy struct {
	Groups         interface{} `json:"Groups,omitempty"`
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`
	PolicyName     interface{} `json:"PolicyName,omitempty"`
	Roles          interface{} `json:"Roles,omitempty"`
	Users          interface{} `json:"Users,omitempty"`
}

type IAMRole struct {
	AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument,omitempty"`
	Path                     interface{} `json:"Path,omitempty"`
	Policies                 interface{} `json:"Policies,omitempty"`
}

type IAMUser struct {
	Path         interface{} `json:"Path,omitempty"`
	Groups       interface{} `json:"Groups,omitempty"`
	LoginProfile interface{} `json:"LoginProfile,omitempty"`
	Policies     interface{} `json:"Policies,omitempty"`
}

type IAMUserToGroupAddition struct {
	GroupName interface{} `json:"GroupName,omitempty"`
	Users     interface{} `json:"Users,omitempty"`
}

type RDSDBInstance struct {
	AllocatedStorage           interface{} `json:"AllocatedStorage,omitempty"`
	AllowMajorVersionUpgrade   interface{} `json:"AllowMajorVersionUpgrade,omitempty"`
	AutoMinorVersionUpgrade    interface{} `json:"AutoMinorVersionUpgrade,omitempty"`
	AvailabilityZone           interface{} `json:"AvailabilityZone,omitempty"`
	BackupRetentionPeriod      interface{} `json:"BackupRetentionPeriod,omitempty"`
	DBInstanceClass            interface{} `json:"DBInstanceClass,omitempty"`
	DBInstanceIdentifier       interface{} `json:"DBInstanceIdentifier,omitempty"`
	DBName                     interface{} `json:"DBName,omitempty"`
	DBParameterGroupName       interface{} `json:"DBParameterGroupName,omitempty"`
	DBSecurityGroups           interface{} `json:"DBSecurityGroups,omitempty"`
	DBSnapshotIdentifier       interface{} `json:"DBSnapshotIdentifier,omitempty"`
	DBSubnetGroupName          interface{} `json:"DBSubnetGroupName,omitempty"`
	Engine                     interface{} `json:"Engine,omitempty"`
	EngineVersion              interface{} `json:"EngineVersion,omitempty"`
	Iops                       interface{} `json:"Iops,omitempty"`
	LicenseModel               interface{} `json:"LicenseModel,omitempty"`
	MasterUsername             interface{} `json:"MasterUsername,omitempty"`
	MasterUserPassword         interface{} `json:"MasterUserPassword,omitempty"`
	MultiAZ                    interface{} `json:"MultiAZ,omitempty"`
	Port                       interface{} `json:"Port,omitempty"`
	PreferredBackupWindow      interface{} `json:"PreferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	SourceDBInstanceIdentifier interface{} `json:"SourceDBInstanceIdentifier,omitempty"`
	Tags                       interface{} `json:"Tags,omitempty"`
	VPCSecurityGroups          interface{} `json:"VPCSecurityGroups,omitempty"`
}

type RDSDBParameterGroup struct {
	Description interface{} `json:"Description,omitempty"`
	Family      interface{} `json:"Family,omitempty"`
	Parameters  interface{} `json:"Parameters,omitempty"`
	Tags        interface{} `json:"Tags,omitempty"`
}

type RDSDBSubnetGroup struct {
	DBSubnetGroupDescription interface{} `json:"DBSubnetGroupDescription,omitempty"`
	SubnetIds                interface{} `json:"SubnetIds,omitempty"`
	Tags                     interface{} `json:"Tags,omitempty"`
}

type RDSDBSecurityGroup struct {
	EC2VpcId               interface{} `json:"EC2VpcId,omitempty"`
	DBSecurityGroupIngress interface{} `json:"DBSecurityGroupIngress,omitempty"`
	GroupDescription       interface{} `json:"GroupDescription,omitempty"`
	Tags                   interface{} `json:"Tags,omitempty"`
}

type RDSDBSecurityGroupIngress struct {
	CIDRIP                  interface{} `json:"CIDRIP,omitempty"`
	DBSecurityGroupName     interface{} `json:"DBSecurityGroupName,omitempty"`
	EC2SecurityGroupId      interface{} `json:"EC2SecurityGroupId,omitempty"`
	EC2SecurityGroupName    interface{} `json:"EC2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerId interface{} `json:"EC2SecurityGroupOwnerId,omitempty"`
}

type Route53RecordSet struct {
	AliasTarget     interface{} `json:"AliasTarget,omitempty"`
	Comment         interface{} `json:"Comment,omitempty"`
	HostedZoneId    interface{} `json:"HostedZoneId,omitempty"`
	HostedZoneName  interface{} `json:"HostedZoneName,omitempty"`
	Name            interface{} `json:"Name,omitempty"`
	Region          interface{} `json:"Region,omitempty"`
	ResourceRecords interface{} `json:"ResourceRecords,omitempty"`
	SetIdentifier   interface{} `json:"SetIdentifier,omitempty"`
	TTL             interface{} `json:"TTL,omitempty"`
	Type            interface{} `json:"Type,omitempty"`
	Weight          interface{} `json:"Weight,omitempty"`
}

type Route53RecordSetGroup struct {
	HostedZoneId   interface{} `json:"HostedZoneId,omitempty"`
	HostedZoneName interface{} `json:"HostedZoneName,omitempty"`
	RecordSets     interface{} `json:"RecordSets,omitempty"`
	Comment        interface{} `json:"Comment,omitempty"`
}

type S3Bucket struct {
	AccessControl        interface{} `json:"AccessControl,omitempty"`
	BucketName           interface{} `json:"BucketName,omitempty"`
	Tags                 interface{} `json:"Tags,omitempty"`
	WebsiteConfiguration interface{} `json:"WebsiteConfiguration,omitempty"`
}

type S3BucketPolicy struct {
	Bucket         interface{} `json:"Bucket,omitempty"`
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`
}

type SDBDomain struct {
}

type SNSTopicPolicy struct {
}

type SNSTopic struct {
	DisplayName  interface{} `json:"DisplayName,omitempty"`
	Subscription interface{} `json:"Subscription,omitempty"`
	TopicName    interface{} `json:"TopicName,omitempty"`
}

type SQSQueue struct {
	DelaySeconds                  interface{} `json:"DelaySeconds,omitempty"`
	MaximumMessageSize            interface{} `json:"MaximumMessageSize,omitempty"`
	MessageRetentionPeriod        interface{} `json:"MessageRetentionPeriod,omitempty"`
	QueueName                     interface{} `json:"QueueName,omitempty"`
	ReceiveMessageWaitTimeSeconds interface{} `json:"ReceiveMessageWaitTimeSeconds,omitempty"`
	RedrivePolicy                 interface{} `json:"RedrivePolicy,omitempty"`
	VisibilityTimeout             interface{} `json:"VisibilityTimeout,omitempty"`
}

type SQSQueuePolicy struct {
}
