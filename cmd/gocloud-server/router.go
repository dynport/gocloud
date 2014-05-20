package main

import "github.com/dynport/gocloud/aws/ec2"

func init() {
	router := app.Router
	router.Get("/api/instances/:id.json", instancesShow)
	router.Get("/api/instances.json", instancesList)

	router.Get("/api/volumes.json", listVolumes)
	router.Get("/api/volumes/:id.json", showVolume)

	router.Get("/api/snapshots/:id.json", showSnapshot)
	router.Get("/api/snapshots.json", listSnapshots)

	router.Get("/api/images/self.json", listImagesFactory(&ImageFilter{Owner: "self"}))
	router.Get("/api/images/canonical.json", listImagesFactory(&ImageFilter{Owner: ec2.CANONICAL_OWNER_ID}))
	router.Get("/api/images/precise.json", listImagesFactory(&ImageFilter{Owner: ec2.CANONICAL_OWNER_ID, Name: "ubuntu/images/ebs/ubuntu-precise-12.04-amd64-server*"}))
	router.Get("/api/images/raring.json", listImagesFactory(&ImageFilter{Owner: ec2.CANONICAL_OWNER_ID, Name: "ubuntu/images/ebs/ubuntu-raring-13.04-amd64-server*"}))
	router.Get("/api/images/saucy.json", listImagesFactory(&ImageFilter{Owner: ec2.CANONICAL_OWNER_ID, Name: "ubuntu/images/ebs/ubuntu-saucy-13.10-amd64-server*"}))

	router.Get("/api/security_groups/:id.json", securityGroupsShow)
	router.Get("/api/security_groups/:id/instances.json", securityGroupsInstances)
	router.Get("/api/security_groups.json", securityGroupsList)

	router.Get("/api/vpcs/:id/subnets.json", vpcsListSubnets)
	router.Get("/api/vpcs/:id.json", vpcsShow)
	router.Get("/api/vpcs.json", vpcsList)

	router.Get("/api/stacks/:name/resources.json", stacksResources)
	router.Get("/api/stacks/:name.json", stacksShow)
	router.Get("/api/stacks.json", stacksList)

	router.Get("/kill", kill)
	router.Get("^/.*", fallback)
}
