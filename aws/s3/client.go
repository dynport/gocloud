package s3

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
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
	DEFAULT_ENDPOINT_HOST = "s3.amazonaws.com"
	HEADER_CONTENT_MD5    = "Content-Md5"
	HEADER_CONTENT_TYPE   = "Content-Type"
	HEADER_DATE           = "Date"
	HEADER_AUTHORIZATION  = "Authorization"
	AMZ_ACL_PUBLIC        = "public-read"
	DEFAULT_CONTENT_TYPE  = "application/octet-stream"
	HEADER_AMZ_ACL        = "X-Amz-Acl"
)

type Client struct {
	*aws.Client
	CustomEndpointHost string
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

func (client *Client) EndpointHost() string {
	if client.CustomEndpointHost != "" {
		return client.CustomEndpointHost
	}
	return DEFAULT_ENDPOINT_HOST
}

func (client *Client) Endpoint() string {
	return "https://" + client.EndpointHost()
}

type PutOptions struct {
	ContentType string
	AmzAcl      string
}

type Content struct {
	Key              string    `xml:"Key"`
	LastModified     time.Time `xml:"LastModified"`
	Etag             string    `xml:"ETag"`
	Size             int64     `xml:"Size"`
	StorageClass     string    `xml:"StorageClass"`
	OwnerID          string    `xml:"Owner>ID"`
	OwnerDisplayName string    `xml:"Owner>DisplayName"`
}

type ListBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`

	Contents []*Content `xml:"Contents"`
}

func (client *Client) Service() (r *ListAllMyBucketsResult, e error) {
	req, e := http.NewRequest("GET", client.Endpoint()+"/", nil)
	if e != nil {
		return r, e
	}
	client.SignS3Request(req, "")
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return r, e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return r, e
	}
	r = &ListAllMyBucketsResult{}
	e = xml.Unmarshal(b, r)
	if e != nil {
		return nil, e
	}
	return r, e
}

func (client *Client) Put(bucket, key string, data []byte, options *PutOptions) error {
	log.Printf("uploading %s to %s/%s", data, bucket, key)
	if options == nil {
		options = &PutOptions{ContentType: DEFAULT_CONTENT_TYPE}
	}

	buf := &bytes.Buffer{}
	buf.Write(data)
	req, e := http.NewRequest("PUT", "http://"+bucket+"."+client.EndpointHost()+"/"+key, buf)
	if e != nil {
		return e
	}

	req.Header.Add("Host", bucket+"."+client.EndpointHost())

	contentType := options.ContentType
	if contentType == "" {
		contentType = DEFAULT_CONTENT_TYPE
	}
	req.Header.Add(HEADER_CONTENT_TYPE, contentType)

	if options.AmzAcl != "" {
		req.Header.Add(HEADER_AMZ_ACL, options.AmzAcl)
	}

	b64md5, e := contentMd5(string(data))
	if e != nil {
		return e
	}
	req.Header.Add(HEADER_CONTENT_MD5, b64md5)

	client.SignS3Request(req, bucket)
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		log.Println(string(b))
		return e
	}
	if rsp.StatusCode != 200 {
		return fmt.Errorf("error uploading key: %s - %s", rsp.Status, string(b))
	}
	return nil
}

func (client *Client) SignS3Request(req *http.Request, bucket string) {
	t := time.Now().UTC()
	date := t.Format(http.TimeFormat)
	payloadParts := []string{
		req.Method,
		req.Header.Get(HEADER_CONTENT_MD5),
		req.Header.Get(HEADER_CONTENT_TYPE),
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
	req.Header.Add(HEADER_DATE, date)
	req.Header.Add(HEADER_AUTHORIZATION, "AWS "+client.Key+":"+signPayload(payload, client.newSha1Hash(client.Secret)))
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
