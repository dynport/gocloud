package main

import (
	"github.com/dynport/gocloud/profitbricks/actions"
)

func init() {
	router.Register("pb/dcs/describe", actions.DescribeDataCenter)
	router.Register("pb/dcs/list", actions.ListAllDataCenters)

	router.Register("pb/servers/start", actions.StartServer)
	router.Register("pb/servers/stop", actions.StopServer)
	router.Register("pb/servers/delete", actions.DeleteServer)
	router.Register("pb/servers/create", actions.CreateServer)
	router.Register("pb/servers/list", actions.ListAllServers)

	router.Register("pb/storages/delete", actions.DeleteStorage)
	router.Register("pb/storages/create", actions.CreateStorage)
	router.Register("pb/storages/list", actions.ListAllStorages)

	router.Register("pb/snapshots/list", actions.ListAllSnapshots)
	router.Register("pb/snapshots/rollback", actions.RollbackSnapshot)

	router.Register("pb/images/list", actions.ListAllImages)
}
