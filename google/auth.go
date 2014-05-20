package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"

	"github.com/dynport/goauth2/oauth"
)

func LoadExchangedToken(code string, config *oauth.Config, token interface{}) error {
	u := config.TokenURL
	v := &url.Values{}
	v.Add("code", code)
	v.Add("client_id", config.ClientId)
	v.Add("client_secret", config.ClientSecret)
	v.Add("grant_type", "authorization_code")
	v.Add("redirect_uri", config.RedirectURL)
	req, e := http.NewRequest("POST", u+"?"+v.Encode(), nil)
	if e != nil {
		return e
	}
	rsp, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	defer rsp.Body.Close()
	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	fmt.Println("TOKEN: " + string(b))
	return json.Unmarshal(b, token)
}

func LoadToken(transport *oauth.Transport, port int) (e error) {
	code, e := GetAuthCode(transport.Config, port)
	if e != nil {
		return e
	}
	transport.Token, e = transport.Exchange(code)
	return e
}

type handler struct {
	codeChannel chan string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no code provided"))
		return
	}
	select {
	case h.codeChannel <- code:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("you can now close your browser and go back to the console"))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unabe to write to channel"))
	}
}

func GetAuthCode(config *oauth.Config, port int) (code string, e error) {
	addr := ""
	if port > 0 {
		addr = ":" + strconv.Itoa(port)
	}
	l, e := net.Listen("tcp", addr)
	defer l.Close()
	h := &handler{}
	h.codeChannel = make(chan string)
	go func() {
		fmt.Println("starting server")
		http.Serve(l, h)
	}()
	parts := strings.Split(l.Addr().String(), ":")
	if len(parts) > 0 {
		addr = parts[len(parts)-1]
	}
	config.RedirectURL = "http://localhost:" + addr
	log.Printf("using RedirectURL=%s", config.RedirectURL)
	url := config.AuthCodeURL("")
	cmd := exec.Command("open", url)
	e = cmd.Run()
	if e != nil {
		return "", e
	}
	fmt.Println("waiting for code")
	select {
	case code := <-h.codeChannel:
		fmt.Println("got code " + code)
		// just in case you have the terminal running already
		exec.Command("open -a Terminal").Run()
		return code, nil
		// maybe add timeout call
	}
	return
}
