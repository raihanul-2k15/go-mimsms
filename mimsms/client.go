package mimsms

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	apiKey  string
	baseUrl string

	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	hc := &http.Client{Timeout: 30 * time.Second}
	return &Client{
		apiKey:     apiKey,
		baseUrl:    "https://isms.mimsms.com",
		httpClient: hc,
	}
}

func (c *Client) sendRequest(method string, path string, query map[string]string) (string, error) {
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
