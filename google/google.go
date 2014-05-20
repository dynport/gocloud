package google

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	compute "code.google.com/p/google-api-go-client/compute/v1beta16"
	"github.com/dynport/goauth2/oauth"
)

type Client struct {
	ProjectId    string
	ClientId     string
	ClientSecret string
	service      *compute.Service
}

func (client *Client) Config() *oauth.Config {
	return &oauth.Config{
		ClientId:     client.ClientId,
		ClientSecret: client.ClientSecret,
		Scope:        compute.ComputeScope,
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
		TokenCache:   oauth.CacheFile("/tmp/google_compute_tokens_" + client.ProjectId + ".json"),
	}
}

func (client *Client) Compute() (compute *Compute, e error) {
	notSet := []string{}
	for k, v := range map[string]string{"ProjectId": client.ProjectId, "ClientId": client.ClientId, "ClientSecret": client.ClientSecret} {
		if v == "" {
			notSet = append(notSet, k)
		}
	}
	if len(notSet) > 0 {
		return nil, fmt.Errorf("%s attributes must be set", strings.Join(notSet, ","))
	}
	service, e := client.Service()
	if e != nil {
		return nil, e
	}
	return &Compute{Service: service}, nil
}

func (client *Client) Service() (service *compute.Service, e error) {
	if client.service != nil {
		return client.service, nil
	}

	config := client.Config()

	transport := &oauth.Transport{Config: config, Transport: http.DefaultTransport}
	token, e := config.TokenCache.Token()
	if token == nil {
		code, e := GetAuthCode(config, 3820)
		if e != nil {
			return nil, e
		}
		token, e = transport.Exchange(code)
		if e != nil {
			log.Fatal(e.Error())
		}
	}
	transport.Token = token
	return compute.New(transport.Client())
}

func (client *Client) Transport(config *oauth.Config) (*oauth.Transport, error) {
	transport := &oauth.Transport{Config: config, Transport: http.DefaultTransport}
	token, _ := config.TokenCache.Token()
	if token == nil {
		code, e := GetAuthCode(config, 3820)
		if e != nil {
			return nil, e
		}
		token, e = transport.Exchange(code)
		if e != nil {
			return nil, e
		}
	}
	transport.Token = token
	return transport, nil
}

type Compute struct {
	*compute.Service
}

func (c *Compute) NewZonesService() *compute.ZonesService {
	return compute.NewZonesService(c.Service)
}

func (c *Compute) NewMachineTypesService() *compute.MachineTypesService {
	return compute.NewMachineTypesService(c.Service)
}

func (c *Compute) NewImagesService() *compute.ImagesService {
	return compute.NewImagesService(c.Service)
}
