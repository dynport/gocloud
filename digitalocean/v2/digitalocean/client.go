package digitalocean

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type DropletsList struct {
}

type Client struct {
	*http.Client
}

var root = "https://api.digitalocean.com"

func (c *Client) loadResponse(path string, i interface{}) error {
	rsp, e := c.Get(root + "/" + strings.TrimPrefix(path, "/"))
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}
	return json.Unmarshal(b, &i)
}

func NewFromEnv() (*Client, error) {
	cl := &transport{apiToken: os.Getenv("DIGITAL_OCEAN_API_KEY")}
	if cl.apiToken == "" {
		return nil, fmt.Errorf("DIGITAL_OCEAN_API_KEY must be set in env")
	}
	return &Client{
			Client: &http.Client{
				Transport: cl,
			},
		},
		nil
}

type transport struct {
	apiToken string
}

func (c *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer "+c.apiToken)
	return http.DefaultClient.Do(r)
}
