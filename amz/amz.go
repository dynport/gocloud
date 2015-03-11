package amz

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dynport/gocloud/aws"
	"github.com/dynport/gocloud/aws/iam"
	"github.com/dynport/gocloud/aws/sts"
)

var (
	logOut = log.New(os.Stdout, "", 0)
	logErr = log.New(os.Stderr, "", 0)
)

type Config struct {
	AWSSecurityToken   string `json:"aws_security_token,omit_empty"`
	AWSAccessKeyID     string `json:"aws_access_key_id"`
	AWSSecretAccessKey string `json:"aws_secret_access_key"`
	AWSDefaultRegion   string `json:"aws_default_region"`
}

func SessionTokenFromConfigPath(path string, maxAge time.Duration) (*sts.Credentials, error) {
	c, err := loadConfig(path)
	if err != nil {
		return nil, err
	}
	return SessionTokenFromCredentials(c.AWSAccessKeyID, c.AWSSecretAccessKey, maxAge)
}

func loadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var c *Config
	return c, json.NewDecoder(f).Decode(&c)
}

func loadSessionToken(path string) (*sts.Credentials, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var c *sts.Credentials
	return c, json.NewDecoder(f).Decode(&c)
}

func SessionTokenFromCredentials(key, secret string, maxAge time.Duration) (*sts.Credentials, error) {
	if key == "" || secret == "" {
		panic("key and secret must not be blank")
	}
	sessionPath := "/tmp/wunderscale/" + key + ".json"

	creds, err := loadSessionToken(sessionPath)

	if err == nil && creds.Expiration.After(time.Now().Add(15*time.Minute)) { // at least 15 minutes valid
		return creds, nil
	}

	awsc := &aws.Client{Key: key, Secret: secret}

	sno, err := serialNumber(&iam.Client{Client: awsc})
	if err != nil {
		return nil, err
	}

	token := readToken()

	stReq := sts.GetSessionToken{
		DurationSeconds: int(maxAge.Seconds()),
		TokenCode:       token,
		SerialNumber:    sno,
	}
	stResp, err := stReq.Execute(awsc)
	if err != nil {
		return nil, err
	}

	tmpPath := sessionPath + ".tmp"
	if err := os.RemoveAll(tmpPath); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(sessionPath), 0755); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	defer os.RemoveAll(tmpPath)
	if err := json.NewEncoder(f).Encode(stResp.Credentials); err != nil {
		return nil, err
	}
	return stResp.Credentials, os.Rename(tmpPath, sessionPath)

}

func serialNumber(iamc *iam.Client) (string, error) {
	user, err := iamc.GetUser("")
	if err != nil {
		return "", err
	}

	logOut.Printf("welcome %s", user.UserName)
	devResp, err := iamc.ListMFADevices(user.UserName)
	if err != nil {
		return "", err
	}

	switch len(devResp.MFADevices) {
	case 0:
		return "", fmt.Errorf("no MFA device configured for user %q", user.UserName)
	case 1: // ignore
	default:
		return "", fmt.Errorf("multiple MFA devices configured for user %q", user.UserName)
	}
	return devResp.MFADevices[0].SerialNumber, nil
}

func readToken() string {
LOOP:
	for {
		fmt.Fprintf(os.Stdout, "Token: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		in := scanner.Text()
		if len(in) != 6 {
			logErr.Printf("token must be a 6 digit number: %d digits read", len(in))
			continue
		}

		for i := range in {
			if in[i] < '0' || in[i] > '9' {
				logErr.Printf("token must be a 6 digit number: %s is not a number", string(in[i]))
				continue LOOP
			}
		}
		return in
	}
}
