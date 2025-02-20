package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

// POST 发送HTTP POST请求
func POST(url string, body []byte) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	resp, err := http.Post(url, "application/json;charset=utf-8", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}
	return respBody, nil
}

// GET 发送HTTP GET请求
func GET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}
	return respBody, nil
}

// HTTP 发送HTTP请求
func HTTP(url, method string, body []byte, headers map[string]string) ([]byte, error) {
	buf := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", string(respBody))
	}
	return respBody, nil
}

// Sign 接口签名
func Sign(params map[string]string, secret string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%s=%s&", k, params[k]))
	}
	sb.WriteString("key=" + secret)

	return MD5(sb.String())
}
