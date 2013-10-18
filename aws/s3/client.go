package s3

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"github.com/dynport/gocloud/aws"
	"hash"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

var b64 = base64.StdEncoding

const (
	ENDPOINT     = "https://s3.amazonaws.com"
	CONTENT_MD5  = "Content-Md5"
	CONTENT_TYPE = "Content-Type"
)

type Client struct {
	*aws.Client
}

func NewFromEnv() *Client {
	return &Client{
		Client: aws.NewFromEnv(),
	}
}

type Bucket struct {
	Name         string    `xml:"Name"`
	CreationDate time.Time `xml:"CreationDate"`
}

type ListAllMyBucketsResult struct {
	XMLName          xml.Name `xml:"ListAllMyBucketsResult"`
	OwnerID          string   `xml:"Owner>ID"`
	OwnerDisplayName string   `xml:"Owner>DisplayName"`

	Buckets []*Bucket `xml:"Buckets>Bucket"`
}

func (client *Client) Service() (r *ListAllMyBucketsResult, e error) {
	req, e := http.NewRequest("GET", ENDPOINT+"/", nil)
	if e != nil {
		return r, e
	}
	client.SignS3Request(req, "")
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return r, e
	}
	log.Println(rsp.Status)
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return r, e
	}
	r = &ListAllMyBucketsResult{}
	e = xml.Unmarshal(b, r)
	return r, e
}

func (client *Client) SignS3Request(req *http.Request, bucket string) {
	t := time.Now().UTC()
	date := t.Format(http.TimeFormat)
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
	path := req.URL.Path
	if bucket != "" {
		path = "/" + bucket + path
	}
	payloadParts = append(payloadParts, path)
	payload := strings.Join(payloadParts, "\n")
	req.Header.Add("Date", date)
	req.Header.Add("Authorization", "AWS "+client.Key+":"+signPayload(payload, client.newSha1Hash(client.Secret)))
}

func (client *Client) newSha1Hash(secret string) hash.Hash {
	return hmac.New(sha1.New, []byte(client.Secret))
}

func signPayload(payload string, hash hash.Hash) string {
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum(nil))
	return string(signature)
}
