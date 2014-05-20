package drive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func loadRequest(client *http.Client, req *http.Request, i interface{}) error {
	rsp, e := client.Do(req)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	buf := &bytes.Buffer{}
	r := io.TeeReader(rsp.Body, buf)
	e = json.NewDecoder(r).Decode(i)
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected 2xx got %s: %s", rsp.Status, buf.String())
	}
	if e != nil {
		log.Printf("ERROR: " + buf.String())
		return e
	}
	return nil

}
