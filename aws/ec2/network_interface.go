package ec2

type CreateNetworkInterface struct {
	DeviceIndex              int
	AssociatePublicIpAddress bool
	SubnetId                 string
	SecurityGroupIds         []string
}
