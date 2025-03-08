package ideepseek

import (
	"encoding/json"
	"fmt"

	"github.com/lcsin/glib/pkg"
)

func (c *Client) Chat(messages ...Messages) (*ResponseBody, error) {
	type payload struct {
		Model    string     `json:"model"`
		Messages []Messages `json:"messages"`
		Stream   bool       `json:"stream"`
	}

	url := fmt.Sprintf("%v/%v/%v", c.Url, "chat", "completions")
	req := payload{
		Model:    c.Model,
		Messages: messages,
		Stream:   c.Stream,
	}
	body, _ := json.Marshal(req)
	bytes, err := pkg.HTTP(url, "POST", body, map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %v", c.Apikey),
	})
	if err != nil {
		return nil, err
	}

	var resp ResponseBody
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
