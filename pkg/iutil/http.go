package iutil

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
	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewReader(body))
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
func HTTP(url, method string, body []byte, header http.Header) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = header

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
