package mimsms

import (
	"fmt"
	"strings"
)

type contentType struct {
	val string
}

type messageType struct {
	val string
}

var (
	ContentTypeText    = contentType{"text"}
	ContentTypeUnicode = contentType{"unicode"}
)

var (
	MessageTypeTransactional = messageType{"transactional"}
	MessageTypePromotional   = messageType{"promotional"}
)

func (t *contentType) String() string {
	return t.val
}

func (t *messageType) String() string {
	return t.val
}

func (c *Client) SendMessage(sender string, recipients []string, msg string, contentType contentType, messageType messageType) (string, error) {
	sender = "88" + cleanPhoneNumber(sender)
	contacts := prepareRecipientsArray(recipients)
	query := map[string]string{
		"api_key":  c.apiKey,
		"senderid": sender,
		"type":     contentType.val,
		"label":    messageType.val,
		"contacts": contacts,
		"msg":      msg,
	}

	body, err := c.sendRequest("POST", "/smsapi", query)
	if err != nil {
		return "", err
	}

	return extractShootId(body)
}

func extractShootId(body string) (string, error) {
	parts := strings.Split(body, "- ")
	if len(parts) < 2 {
		return "", fmt.Errorf("ID Parse Failure, Original Response: %s", body)
	}
	return strings.Trim(parts[1], " "), nil
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
