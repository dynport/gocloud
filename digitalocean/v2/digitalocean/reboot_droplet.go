package digitalocean

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) RebootDroplet(id string) error {
	req, err := http.NewRequest("POST", root+"/v2/droplets/"+id+"/actions", strings.NewReader(`{"type": "reboot"}`))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	rsp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	if rsp.Status[0] != '2' {
		b, _ := ioutil.ReadAll(rsp.Body)
		return fmt.Errorf("got status %s but expected 2x. body=%s", rsp.Status, string(b))
	}
	return nil
}
