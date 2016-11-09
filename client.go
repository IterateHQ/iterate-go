// Package iterate provides access to the Iterate API
package iterate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const host = "https://iteratehq.com/api/v1"

// Client manages the communication with the Iterate API.
type Client struct {
	host       string
	httpClient *http.Client
	token      string
	version    string
}

// NewClient creates a new client with an API auth token.
func New(token string) Client {
	version := "20161109"

	return Client{
		host:       host,
		httpClient: &http.Client{},
		token:      token,
		version:    version,
	}
}

func (c Client) get(path string, values url.Values) ([]byte, error) {
	r, _ := http.NewRequest("GET", c.host+path, nil)
	r.URL.RawQuery = c.withDefaultParams(values).Encode()

	return c.sendRequest(r)
}

func (c Client) post(path string, values url.Values) ([]byte, error) {
	// Configure the request
	r, _ := http.NewRequest("POST", c.host+path, bytes.NewBufferString(c.withDefaultParams(values).Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return c.sendRequest(r)
}

func (c Client) sendRequest(r *http.Request) (results []byte, err error) {
	// Send the request
	rawResp, err := c.httpClient.Do(r)
	if err != nil {
		return
	}
	defer rawResp.Body.Close()
	body, err := ioutil.ReadAll(rawResp.Body)

	// Parse the response
	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	if resp.Error != "" {
		err = errors.New(resp.Error)
		return
	}

	results, err = json.Marshal(resp.Results)
	return
}

func (c Client) withDefaultParams(values url.Values) url.Values {
	values.Add("v", c.version)
	values.Add("access_token", c.token)

	return values
}
