package mimsms

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type possibleErrorResponse struct {
	Request string `json:"request"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func isResponseError(resp string) error {
	var jsonResp possibleErrorResponse
	err := json.Unmarshal([]byte(resp), &jsonResp)
	if err != nil {
		return err
	}

	if jsonResp.Status == "error" {
		errMsg := jsonResp.Request + ": " + jsonResp.Message
		return fmt.Errorf("API Error: %s, Original Response: %s", errMsg, resp)
	}

	return nil
}

func (c *client) safeError(err error) error {
	if err == nil {
		return nil
	}

	errMsg := strings.Replace(err.Error(), c.apiKey, "********", -1)
	errMsg = strings.Replace(errMsg, c.apiToken, "********", -1)
	return errors.New(errMsg)
}
