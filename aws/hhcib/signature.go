package main

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Signature struct {
	Key    string
	Secret string
	Path   string
	Host   string
	Method string
	Values url.Values
	query  string
	time   time.Time
}

func (sig *Signature) Request() (*http.Request, error) {
	payload := sig.Payload()
	signature := signPayload(payload, sig.Secret)
	theUrl := "https://" + sig.Host + sig.Path + "?" + sig.query + "&Signature=" + signature
	return http.NewRequest(sig.Method, theUrl, nil)
}

func (sig *Signature) Payload() string {
	values := sig.Values
	if sig.time.IsZero() {
		sig.time = time.Now()
	}
	if len(values["Timestamp"]) == 0 {
		values.Add("Timestamp", sig.time.UTC().Format(time.RFC3339))
	}
	if len(values["AWSAccessKeyId"]) == 0 {
		values.Add("AWSAccessKeyId", sig.Key)
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

	logger.Print(sarray)

	sig.query = strings.Join(sarray, "&")
	if sig.Path == "" {
		sig.Path = "/"
	}
	return strings.Join([]string{
		sig.Method,
		sig.Host,
		sig.Path,
		sig.query,
	}, "\n")
}
