package jiffybox

import (
	"encoding/json"
	"fmt"
	"github.com/dynport/gologger"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ENV_API_TOKEN = "JIFFYBOX_API_KEY"

type Client struct {
	ApiKey string
}

func abortWith(message string) {
	fmt.Println("ERROR: " + message)
	os.Exit(1)
}

func NewFromEnv() *Client {
	key := os.Getenv(ENV_API_TOKEN)
	if key == "" {
		abortWith(ENV_API_TOKEN + " must be set")
	}
	return New(key)
}

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}

func (client *Client) BaseUrl() string {
	return "https://api.jiffybox.de/" + client.ApiKey + "/v1.0"
}

type Response struct {
	Messages []string `json:"messages"`
	Result   bool     `json:"result"`
}

type HttpResponse struct {
	StatusCode int
	Content    []byte
	Response   *http.Response
}

func (client *Client) unmarshalResponse(rsp *http.Response, i interface{}) error {
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	logger.Debug(string(b))
	if e != nil {
		return e
	}
	return client.unmarshal(b, i)
}

func (client *Client) unmarshal(b []byte, i interface{}) error {
	er := &ErrorResponse{}
	e := json.Unmarshal(b, er)
	logger.Debug("result of unmarshalling err", e)
	if e == nil {
		allErrors := []string{}
		for _, message := range er.Messages {
			allErrors = append(allErrors, message.Message)
		}
		if len(allErrors) > 0 {
			return fmt.Errorf(strings.Join(allErrors, ", "))
		}
	}
	return json.Unmarshal(b, i)
}

func (client *Client) PostForm(action string, values url.Values) (rsp *HttpResponse, e error) {
	u := client.BaseUrl() + "/" + action
	logger.Infof("sending request " + u)
	httpResponse, e := http.PostForm(u, values)
	if e != nil {
		return nil, e
	}
	logger.Debugf("got status %s", httpResponse.Status)
	rsp = &HttpResponse{
		StatusCode: httpResponse.StatusCode,
	}
	rsp.Content, e = ioutil.ReadAll(httpResponse.Body)
	rsp.Response = httpResponse
	if e != nil {
		return nil, e
	}
	return rsp, e
}

var logger = gologger.NewFromEnv()

type ErrorResponse struct {
	Messages []*Message `json:"messages"`
	Result   bool       `json:"result"`
}

func (client *Client) LoadResource(action string, i interface{}) error {
	u := client.BaseUrl() + "/" + action
	logger.Debug("loading " + u)
	rsp, e := http.Get(u)
	if e != nil {
		return e
	}
	return client.unmarshalResponse(rsp, i)
}
