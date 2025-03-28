package itest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Case struct {
	Name      string
	Before    func(t *testing.T)
	After     func(t *testing.T)
	ReqBody   any
	ExpCode   int
	ExpResult any

	Response *httptest.ResponseRecorder
}

func (c *Case) HttpTest(method, url string, body []byte, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
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
