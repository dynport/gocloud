package rds

import "encoding/xml"

type DeleteDBInstance struct {
	DBInstanceIdentifier string `xml:",omitempty"`
	// Don't create an instance final snapshot.
	SkipFinalSnapshot bool `xml:",omitempty"`
	// Name of the final instance snapshot.
	FinalDBSnapshotIdentifier string `xml:",omitempty"`
}

type DeleteDBInstanceResponse struct {
	XMLName                xml.Name                `xml:"DeleteDBInstanceResponse"`
	DeleteDBInstanceResult *DeleteDBInstanceResult `xml:"DeleDescribeDBInstancesResult"`
}

type DeleteDBInstanceResult struct {
	Instance *DBInstance `xml:"DBInstance"`
}

func (r *DeleteDBInstance) Execute(client *Client) (*DeleteDBInstanceResponse, error) {
	v := newAction("DeleteDBInstance")
	if e := loadValues(v, r); e != nil {
		return nil, e
	}

	resp := &DeleteDBInstanceResponse{}
	return resp, client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, resp)
}
