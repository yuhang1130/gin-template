package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func baseRequest(method, urlStr string, body io.Reader, headers map[string]string, timeout time.Duration) (string, error) {
	// 创建请求
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return "", err
	}
	// 添加请求头
	if len(headers) != 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}

	// 发送请求
	if timeout <= 0 {
		timeout = 15
	}
	httpClient := &http.Client{
		Timeout: timeout * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

func GetHttpRequest(urlStr string, params, headers map[string]string, timeout time.Duration) (string, error) {
	var bodyStr string
	if len(params) != 0 {
		data := url.Values{}
		for key, value := range params {
			data.Set(key, value)
		}
		bodyStr = data.Encode()
	}
	reqUrl := fmt.Sprintf("%s?%s", urlStr, bodyStr)
	return baseRequest(http.MethodGet, reqUrl, nil, headers, timeout)
}

func PostHttpRequest(urlStr string, body any, headers map[string]string, timeout time.Duration) (string, error) {
	data, err := json.Marshal(body)
	if err != nil {
		log.Printf("json Marshal Error: %s\n body: %+v\n", err, body)
		return "", err
	}
	// log.Println("PostHttpRequest body:", string(data))
	bodyReader := bytes.NewReader(data)
	return baseRequest(http.MethodPost, urlStr, bodyReader, headers, timeout)
}
