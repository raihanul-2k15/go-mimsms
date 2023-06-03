package mimsms

import (
	"errors"
	"fmt"
	"strings"
)

var errorMap map[string]string = map[string]string{
	"1002": "Sender Id/Masking Not Found",
	"1003": "API Not Found",
	"1004": "SPAM Detected",
	"1005": "Internal Error",
	"1006": "Internal Error",
	"1007": "Balance Insufficient",
	"1008": "Message is empty",
	"1009": "Message Type Not Set (text/unicode)",
	"1010": "Invalid User & Password",
	"1011": "Invalid User Id",
}

func isResponseError(resp string) error {
	lower := strings.ToLower(resp)
	code := ""

	if strings.Contains(lower, "error:") {
		code = strings.Trim(strings.Split(lower, "error:")[1], " ")
	} else {
		code = lower
	}

	if err, ok := errorMap[code]; ok {
		return fmt.Errorf("API Error: %s, Original Response: %s", err, resp)
	}
	return nil
}

func (c *Client) safeError(err error) error {
	if err == nil {
		return nil
	}

	return errors.New(strings.Replace(err.Error(), c.apiKey, "********", -1))
}
