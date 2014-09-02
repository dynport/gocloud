package rds

import "encoding/xml"

type ModifyDBInstance struct {
	DBInstanceIdentifier string
	DBSecurityGroups     []string `xml:"DBSecurityGroups>member,omitempty"`
	VpcSecurityGroupIds  []string `xml:"VpcSecurityGroupIds>member,omitempty"`
}

type ModifyDBInstanceResponse struct {
	XMLName                xml.Name                `xml:"ModifyDBInstanceResponse"`
	ModifyDBInstanceResult *ModifyDBInstanceResult `xml:"ModifyDBInstanceResult,omitempty"`
}

type ModifyDBInstanceResult struct {
	XMLName    xml.Name    `xml:"ModifyDBInstanceResult"`
	DBInstance *DBInstance `xml:"DBInstance"`
}

func (action *ModifyDBInstance) Execute(client *Client) (res *ModifyDBInstanceResponse, e error) {
	v := newAction("ModifyDBInstance")
	if e = loadValues(v, action); e != nil {
		return nil, e
	}

	res = &ModifyDBInstanceResponse{}
	return res, client.loadResource("GET", client.Endpoint()+"?"+v.Encode(), nil, res)
}
