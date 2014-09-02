package rds

import "encoding/xml"

type RestoreDBSnapshot struct {
	DBInstanceClass      string `xml:",omitempty"`
	DBInstanceIdentifier string `xml:",omitempty"`
	DBSnapshotIdentifier string `xml:",omitempty"`
}

type RestoreDBSnapshotResponse struct {
	XMLName                 xml.Name                 `xml:"RestoreDBInstanceFromDBSnapshotResponse"`
	RestoreDBSnapshotResult *RestoreDBSnapshotResult `xml:"RestoreDBInstanceFromDBSnapshotResult"`
}

type RestoreDBSnapshotResult struct {
	Instance *DBInstance `xml:"DBInstance"`
}

func (r *RestoreDBSnapshot) Execute(client *Client) (*RestoreDBSnapshotResponse, error) {
	v := newAction("RestoreDBInstanceFromDBSnapshot")
	if e := loadValues(v, r); e != nil {
		return nil, e
	}

	resp := &RestoreDBSnapshotResponse{}
	return resp, client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, resp)
}
