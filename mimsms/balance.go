package mimsms

import (
	"fmt"
	"strings"
)

func (c *Client) GetBalance() (string, error) {
	body, err := c.sendRequest("GET", fmt.Sprintf("/miscapi/%s/getBalance", c.apiKey), map[string]string{})
	if err != nil {
		return "", err
	}

	parts := strings.Split(body, "BDT")
	if len(parts) != 2 {
		return "", fmt.Errorf("Balance Parse Failure, Original Response: %s", body)
	}

	balance := strings.Trim(parts[1], " ")
	balance = strings.Replace(balance, ",", "", -1)

	return balance, nil
}
