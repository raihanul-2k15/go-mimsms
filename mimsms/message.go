package mimsms

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) SendMessage(sender string, recipients []string, msg string) (string, error) {
	sender = "88" + cleanPhoneNumber(sender)
	contacts := prepareRecipientsArray(recipients)
	query := map[string]string{
		"sendsms":  "1",
		"type":     "sms",
		"apikey":   c.apiKey,
		"apitoken": c.apiToken,
		"from":     sender,
		"to":       contacts,
		"text":     msg,
	}

	body, err := c.sendRequest("POST", "/", query)
	if err != nil {
		return "", err
	}

	return extractGroupId(body)
}

type groupIdResponse struct {
	Request string `json:"request"`
	Status  string `json:"status"`
	GroupId string `json:"group_id"`
	Date    string `json:"date"`
}

func extractGroupId(body string) (string, error) {
	var jsonResp groupIdResponse
	err := json.Unmarshal([]byte(body), &jsonResp)
	if err != nil {
		return "", err
	}

	if jsonResp.GroupId == "" {
		errMsg := jsonResp.Request + ": " + jsonResp.Status
		return "", fmt.Errorf("API Error: %s, Original Response: %s", errMsg, body)
	}

	return jsonResp.GroupId, nil
}

func prepareRecipientsArray(recipients []string) string {
	var numbers []string
	for _, number := range recipients {
		numbers = append(numbers, "88"+cleanPhoneNumber(number))
	}
	return strings.Join(numbers, "+")
}

func cleanPhoneNumber(number string) string {
	number = strings.TrimPrefix(number, "+")
	number = strings.TrimPrefix(number, "88")
	replacer := strings.NewReplacer(" ", "", "-", "", "_", "")
	number = replacer.Replace(number)
	return number
}
