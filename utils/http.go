package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	httpClient *http.Client
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			DisableKeepAlives:     false,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return client
}

func HttpGet(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("respone code error")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HttpGetJson(url string, v interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("respone code error")
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// contentType application/json | application/x-www-form-urlencoded
func HttpPost(url, contentType string, params interface{}) ([]byte, error) {
	var paramsBytes []byte
	var err error
	if contentType == "application/json" {
		paramsBytes, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
	}

	if contentType == "application/x-www-form-urlencoded" {
		paramsBytes = []byte(params.(string))
	}

	res, err := http.Post(url, contentType, bytes.NewReader(paramsBytes))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("response code error")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PostFile(uri string, params map[string]string, paramName, fileName string, fileData []byte, headerParams map[string]string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	part, err := writer.CreateFormFile(paramName, fileName)
	if err != nil {
		return nil, err
	}

	if _, err := part.Write(fileData); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	for key, val := range headerParams {
		request.Header.Set(key, val)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("response.StatusCode =" + strconv.Itoa(response.StatusCode))
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PostData(uri string, data []byte, header map[string]string) ([]byte, error) {
	request, err := http.NewRequest("POST", uri, bytes.NewReader(data))

	for key, val := range header {
		request.Header.Set(key, val)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("response.StatusCode =" + strconv.Itoa(response.StatusCode))
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ErrorInfo struct {
	Err error
	Msg string
}

func DownloadUrl(url, path string, ch chan *ErrorInfo) (err error) {
	errInfo := new(ErrorInfo)
	defer func() {
		errInfo.Err = err
		errInfo.Msg = path
		ch <- errInfo
	}()
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
