package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

// Client is the couchbase client
type Client struct {
	baseURL string
	client  http.Client
}

// New creates a new couchbase client
func New(url, user, password string) Client {
	var client = Client{
		baseURL: url,
		client: http.Client{
			Transport: &AuthTransport{
				Username: user,
				Password: password,
			},
		},
	}
	go func() {
		nodes, err := client.Nodes()
		if err != nil {
			log.Warnf("couldn't verify server version compatibility: %s", err.Error())
			return
		}
		if !strings.HasPrefix(nodes.Nodes[0].Version, "5.") {
			log.Warnf("couchbase %s is not fully supported", nodes.Nodes[0].Version)
		}
	}()
	return client
}

func (c Client) url(path string) string {
	return c.baseURL + "/" + path
}

func (c Client) get(path string, v interface{}) error {
	resp, err := c.client.Get(c.url(path))
	if err != nil {
		return errors.Wrapf(err, "failed to get %s", path)
	}

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed to read response body from %s", path)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to get %s metrics: %s %d", path, string(bts), resp.StatusCode)
	}

	if err := json.Unmarshal(bts, v); err != nil {
		return errors.Wrapf(err, "failed to unmarshall %s output: %s", path, string(bts))
	}
	return nil
}

// AuthTransport is a http.RoudTripper that does the autentication
type AuthTransport struct {
	Username string
	Password string

	Transport http.RoundTripper
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// RoundTrip implements the RoundTripper interface.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	req2.SetBasicAuth(t.Username, t.Password)
	return t.transport().RoundTrip(req2)
}
