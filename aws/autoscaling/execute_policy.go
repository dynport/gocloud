package autoscaling

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dynport/gocloud/aws"
)

type ExecutePolicy struct {
	AutoScalingGroupName string
	HonorCooldown        bool
	PolicyName           string
}

func (action *ExecutePolicy) Execute(client *aws.Client) (string, error) {
	ep, e := endpoint(client)
	if e != nil {
		return "", e
	}
	req, e := http.NewRequest("GET", ep+action.query(), nil)

	if e != nil {
		return "", e
	}
	client.SignAwsRequestV2(req, time.Now())
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return "", e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return "", e
	}
	if rsp.Status[0] != '2' {
		return "", fmt.Errorf("expected status 2xx, got %s. %s", rsp.Status, string(b))
	}
	return string(b), nil
}

func (action *ExecutePolicy) query() string {
	values := Values{
		"PolicyName":           action.PolicyName,
		"AutoScalingGroupName": action.AutoScalingGroupName,
		"Action":               "ExecutePolicy",
		"Version":              "2011-01-01",
	}
	if action.HonorCooldown {
		values["HonorCooldown"] = "true"
	}
	return values.query()

}
