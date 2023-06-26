package mimsms

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type client struct {
	apiKey   string
	apiToken string
	baseUrl  string

	httpClient *http.Client
}

func NewClient(apiKey, apiToken string) *client {
	return &client{
		apiKey:     apiKey,
		apiToken:   apiToken,
		baseUrl:    "http://mimsms.com.bd/smsAPI",
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *client) sendRequest(method string, path string, query map[string]string) (string, error) {
	if c.apiKey == "" || c.apiToken == "" {
		return "", errors.New("apiKey or apiToken is empty")
	}

	req, err := http.NewRequest(method, c.baseUrl+path, nil)
	if err != nil {
		return "", c.safeError(err)
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", c.safeError(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", c.safeError(err)
	}

	body := string(bytes)

	err = isResponseError(body)
	if err != nil {
		return "", c.safeError(err)
	}

	return body, nil
}
