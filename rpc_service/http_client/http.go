package http_client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var c *http.Client

func GetHTTPClient() *http.Client {

	if c != nil {
		return c
	}

	return defaultHttpClient()
}

func defaultHttpClient() *http.Client {
	return &http.Client{
		Timeout:   3 * time.Second,
		Transport: &http.Transport{},
	}
}

func SetHttpClient(client *http.Client) {
	c = client
}

func Get(uri string, respData interface{}, header map[string]string) error {

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		return err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := GetHTTPClient().Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return jsoniter.Unmarshal(res, &respData)
}

func Post(uri string, data []byte, head map[string]string, respData interface{}) error {
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range head {
		req.Header.Set(k, v)
	}

	response, err := GetHTTPClient().Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if respData == nil {
		return nil
	}

	return jsoniter.Unmarshal(res, &respData)
}
