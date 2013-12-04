package aws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

var b64 = base64.StdEncoding

type Client struct {
	Key, Secret string
}

var Debug = os.Getenv("DEBUG") == "true"

func (client *Client) Debug(format string, i ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", i...)
	}
}

func NewFromEnv() *Client {
	client := &Client{}
	client.Key, client.Secret = os.Getenv(ENV_AWS_ACCESS_KEY), os.Getenv(ENV_AWS_SECRET_KEY)

	if client.Key == "" || client.Secret == "" {
		abortWith(fmt.Sprintf("%s and %s must be set in ENV", ENV_AWS_ACCESS_KEY, ENV_AWS_SECRET_KEY))
	}
	return client
}

type Response struct {
	StatusCode int
	Content    []byte
}

func QueryPrefix(version, action string) string {
	return "Version=" + version + "&Action=" + action
}

// list of endpoints
func (client *Client) DoSignedRequest(method string, endpoint, action string, extraAttributes map[string]string) (rsp *Response, e error) {
	request, e := http.NewRequest(method, endpoint+"?"+action, nil)
	client.SignAwsRequestV2(request, time.Now())
	raw, e := http.DefaultClient.Do(request)
	if e != nil {
		return rsp, e
	}
	defer raw.Body.Close()
	rsp = &Response{
		StatusCode: raw.StatusCode,
	}
	rsp.Content, e = ioutil.ReadAll(raw.Body)
	if e != nil {
		return rsp, e
	}
	return rsp, e
}

func (client *Client) signPayload(payload string) string {
	hash := hmac.New(sha256.New, []byte(client.Secret))
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum(nil))
	return string(signature)
}

func (client *Client) SignAwsRequest(req *http.Request) {
	date := time.Now().UTC().Format(http.TimeFormat)
	token := "AWS3-HTTPS AWSAccessKeyId=" + client.Key + ",Algorithm=HmacSHA256,Signature=" + client.signPayload(date)
	req.Header.Set("X-Amzn-Authorization", token)
	req.Header.Set("x-amz-date", date)
	return
}

func (client *Client) v2PayloadAndQuery(req *http.Request) (payload, rawQuery string) {
	values := req.URL.Query()
	if len(values["AWSAccessKeyId"]) == 0 {
		values.Add("AWSAccessKeyId", client.Key)
	}

	if len(values["SignatureVersion"]) == 0 {
		values.Add("SignatureVersion", "2")
	}

	if len(values["SignatureMethod"]) == 0 {
		values.Add("SignatureMethod", "HmacSHA256")
	}

	if len(values["Timestamp"]) == 0 {
		values.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	}

	var keys, sarray []string
	for k, _ := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sarray = append(sarray, Encode(k)+"="+Encode(values[k][0]))
	}

	joined := strings.Join(sarray, "&")
	path := "/"
	if req.URL.Path != "" {
		path = req.URL.Path
	}

	return strings.Join([]string{
		req.Method,
		req.URL.Host,
		path,
		joined,
	}, "\n"), joined
}

func (client *Client) SignAwsRequestV2(req *http.Request, t time.Time) {
	values := req.URL.Query()
	if len(values["Timestamp"]) == 0 {
		values.Add("Timestamp", t.UTC().Format(time.RFC3339))
	}
	req.URL.RawQuery = values.Encode()
	payload, query := client.v2PayloadAndQuery(req)
	query += "&Signature=" + Encode(client.signPayload(payload))
	req.URL.RawQuery = query
}

// authentication:	http://s3.amazonaws.com/doc/s3-developer-guide/RESTAuthentication.html
// upload:			http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUT.html
func (client *Client) SignS3Request(req *http.Request) {
	time := time.Now()
	date := time.Format(http.TimeFormat)
	payloadParts := []string{
		req.Method,
		req.Header.Get(CONTENT_MD5),
		req.Header.Get(CONTENT_TYPE),
		date,
	}
	amzHeaders := []string{}
	for k, v := range req.Header {
		value := strings.ToLower(k) + ":" + strings.Join(v, ",")
		if strings.HasPrefix(value, "x-amz") {
			amzHeaders = append(amzHeaders, value)
		}
	}
	sort.Strings(amzHeaders)
	payloadParts = append(payloadParts, amzHeaders...)
	payloadParts = append(payloadParts, req.URL.Path)
	payload := strings.Join(payloadParts, "\n")
	req.Header.Add("Date", date)
	req.Header.Add("Authorization", "AWS "+client.Key+":"+client.signPayload(payload))
}
