package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"srunsoft-api-sdk/configs"
	"strconv"
	"strings"
	"time"
)

// SrunResponse 北向接口响应结构
type SrunResponse struct {
	Code    int
	Message string
	Version string
	Data    interface{}
}

type httpClient struct {
	client *http.Client
}

func newHTTPClient() *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *httpClient) doRequest(method, urlPath string, data url.Values) (SrunResponse, error) {
	var (
		token string
		sr    SrunResponse
	)

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s%s", configs.Config.Scheme, configs.Config.InterfaceIP, urlPath), nil)
	if err != nil {
		configs.Log.Error(err)
		return sr, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if data != nil {
		if urlPath != GetAccessToken {
			token, err = GetToken()
			if err != nil {
				return sr, err
			}
			data.Set("access_token", token)
		}
		req.Body = io.NopCloser(strings.NewReader(data.Encode()))
		req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		configs.Log.Error(err)
		return sr, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			configs.Log.Error(err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return sr, err
	}

	return sr, nil
}

func HandlePost(urlPath string, data url.Values) (SrunResponse, error) {
	client := newHTTPClient()
	return client.doRequest(http.MethodPost, urlPath, data)
}

func HandlePut(urlPath string, data url.Values) (SrunResponse, error) {
	client := newHTTPClient()
	return client.doRequest(http.MethodPut, urlPath, data)
}

func HandleDelete(urlPath string, data url.Values) (SrunResponse, error) {
	client := newHTTPClient()
	return client.doRequest(http.MethodDelete, urlPath, data)
}

func HandleGet(urlPath string, data url.Values) (SrunResponse, error) {
	client := newHTTPClient()
	return client.doRequest(http.MethodGet, urlPath, data)
}

func Map2Values(m map[string]interface{}) url.Values {
	v := url.Values{}
	for key, value := range m {
		v.Add(key, fmt.Sprintf("%v", value))
	}
	return v
}
