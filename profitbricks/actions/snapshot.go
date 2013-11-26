package actions

import (
	"fmt"
	"github.com/dynport/gocli"
	"github.com/dynport/gocloud/profitbricks"
)

func init() {
	ListAllSnapshots = &gocli.Action{Handler: ListAllSnapshotsHandler, Description: "List Snapshots"}

	args := gocli.NewArgs(nil)
	args.RegisterString(CLI_ROLLBACK_SNAPSHOT_STORAGE_ID, "storage_id", true, "", "Storage ID")
	args.RegisterString(CLI_ROLLBACK_SNAPSHOT_SNAPSHOT_ID, "snapshot_id", true, "", "Snapshot ID")

	RollbackSnapshot = &gocli.Action{Handler: RollbackSnapshotHandler, Description: "Rollback Snapshot", Args: args}
}

func RollbackSnapshotHandler(args *gocli.Args) error {
	req := &profitbricks.RollbackSnapshotRequest{
		StorageId:  args.MustGetString(CLI_ROLLBACK_SNAPSHOT_STORAGE_ID),
		SnapshotId: args.MustGetString(CLI_ROLLBACK_SNAPSHOT_SNAPSHOT_ID),
	}
	return profitbricks.NewFromEnv().RollbackSnapshot(req)
}

func ListAllSnapshotsHandler(args *gocli.Args) error {
	client := profitbricks.NewFromEnv()
	snapshots, e := client.GetAllSnapshots()
	if e != nil {
		return e
	}
	table := gocli.NewTable()
	table.Add("Id", "OsType", "Name", "Size", "State")
	for _, snapshot := range snapshots {
		table.Add(snapshot.SnapshotId, snapshot.OsType, snapshot.SnapshotName, snapshot.SnapshotSize, snapshot.ProvisioningState)
	}
	fmt.Println(table)
	return nil
}
