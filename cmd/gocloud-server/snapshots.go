package main

import (
	"net/http"

	"github.com/dynport/martini"
)

func listSnapshots(w http.ResponseWriter) (int, []byte) {
	snapshots, e := describeSnapshots(&DescribeSnapshotsParameters{Owners: []string{"self"}})
	if e != nil {
		return returnError(e)
	}
	return renderJson(w, snapshots)
}

func showSnapshot(w http.ResponseWriter, params martini.Params) (int, []byte) {
	snapshots, e := describeSnapshots(&DescribeSnapshotsParameters{Owners: []string{"self"}})
	if e != nil {
		return returnError(e)
	}
	for _, s := range snapshots {
		if s.SnapshotId == params["id"] {
			return renderJson(w, s)
		}
	}
	return 404, nil
}
