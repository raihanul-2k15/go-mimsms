package mimsms

import (
	"encoding/json"
	"fmt"
)

func (c *client) GetBalance() (string, error) {
	body, err := c.sendRequest("GET", "/", map[string]string{
		"balance":  "1",
		"apikey":   c.apiKey,
		"apitoken": c.apiToken,
	})
	if err != nil {
		return "", err
	}

	var balanceResp struct {
		Balance string `json:"balance"`
	}
	err = json.Unmarshal([]byte(body), &balanceResp)

	if err != nil {
		return "", err
	}

	if balanceResp.Balance == "" {
		return "", fmt.Errorf("API Error: %s, Original Response: %s", "Balance is empty", body)
	}

	return balanceResp.Balance, nil
}
