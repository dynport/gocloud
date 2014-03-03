package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
)

const EC2_VERSION = "2013-10-15"

type AwsClient struct {
	Endpoint        string
	SecretAccessKey string
	AccessKeyId     string
}

func (client *AwsClient) loadResource(query interface{}) error {
	v := url.Values{}
	urlValues(v, query, "")
	log.Printf("sending request to %q", v.Encode())

	theUrl := client.Endpoint + "?" + v.Encode()
	request, e := http.NewRequest("GET", theUrl, nil)
	if e != nil {
		return e
	}
	logger.Print(request.URL.String())

	endpoint, e := url.Parse(client.Endpoint)
	if e != nil {
		return e
	}
	sig := &Signature{
		Host:   endpoint.Host,
		Path:   "/",
		Key:    client.AccessKeyId,
		Secret: client.SecretAccessKey,
		Values: v,
	}

	req, e := sig.Request()
	if e != nil {
		return e
	}

	log.Print(req.URL.String())

	return nil

	return nil

	rsp, e := http.DefaultClient.Do(request)
	logger.Printf("got response %#v %v", rsp, e)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	logger.Printf("status=%s", rsp.Status)
	return nil
}

var (
	b64        = base64.StdEncoding
	unreserved = make([]bool, 128)
	hex        = "0123456789ABCDEF"
)

func signPayload(payload string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum(nil))
	return string(signature)
}

func Encode(s string) string {
	encode := false
	for i := 0; i != len(s); i++ {
		c := s[i]
		if c > 127 || !unreserved[c] {
			encode = true
			break
		}
	}
	if !encode {
		return s
	}
	e := make([]byte, len(s)*3)
	ei := 0
	for i := 0; i != len(s); i++ {
		c := s[i]
		if c > 127 || !unreserved[c] {
			e[ei] = '%'
			e[ei+1] = hex[c>>4]
			e[ei+2] = hex[c&0xF]
			ei += 3
		} else {
			e[ei] = c
			ei += 1
		}
	}
	return string(e[:ei])
}
