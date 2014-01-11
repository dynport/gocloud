package main

import (
	"github.com/dynport/gocloud/google/actions"
)

func init() {
	router.RegisterFunc("google/zones/list", actions.ListZonesHandler, "List Zones")
	router.Register("google/machine-types/list", &actions.ListMachineTypes{}, "List Machine Types")
	router.RegisterFunc("google/images/list", actions.ListImagesHandler, "List Images")
}
