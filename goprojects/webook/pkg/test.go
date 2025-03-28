package pkg

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Case struct {
	Name      string
	Req       any
	ExpCode   int
	ExpResult any

	Before func(t *testing.T)
	After  func(t *testing.T)

	Response *httptest.ResponseRecorder
}

func (c *Case) HttpTest(method, url string, data []byte, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	c.Response = httptest.NewRecorder()
	return request, nil
}

func (c *Case) ResponseBodyDecoder(result any) error {
	return json.NewDecoder(c.Response.Body).Decode(&result)
}
