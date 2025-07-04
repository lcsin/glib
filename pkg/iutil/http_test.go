package iutil

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPOST(t *testing.T) {
	bytes, err := POST("https://www.baidu.com", nil)
	assert.NoError(t, err)
	t.Log(string(bytes))
}

func TestGET(t *testing.T) {
	bytes, err := GET("https://www.baidu.com")
	assert.NoError(t, err)
	t.Log(string(bytes))
}

func TestHTTP(t *testing.T) {
	url := "https://api.deepseek.com/chat/completions"
	method := "POST"
	payload := strings.NewReader(`{
  "messages": [
    {
      "content": "who are you",
      "role": "user"
    }
  ],
  "model": "deepseek-chat",
  "frequency_penalty": 0,
  "max_tokens": 2048,
  "presence_penalty": 0,
  "response_format": {
    "type": "text"
  },
  "stop": null,
  "stream": false,
  "stream_options": null,
  "temperature": 1,
  "top_p": 1,
  "tools": null,
  "tool_choice": "none",
  "logprobs": false,
  "top_logprobs": null
}`)
	body, _ := io.ReadAll(payload)

	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")
	header.Set("Authorization", "Bearer <apikey>")
	bytes, err := HTTP(url, method, body, header)
	if err != nil {
		panic(err)
	}
	t.Log(string(bytes))
}

func TestSign(t *testing.T) {
	params := map[string]string{
		"uid": "1",
		"age": "18",
	}
	key := "apikey"
	sign := Sign(params, key)
	t.Log(sign)
}
