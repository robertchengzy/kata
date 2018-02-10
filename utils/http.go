package utils

import (
	"io/ioutil"
	"net/http"
	"errors"
	"strconv"
	"bytes"
	"mime/multipart"
	"encoding/json"
	"os"
	"io"
	"github.com/xiam/resp"
)

func HttpUrlGet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return ""
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	if len(data) == 0 {
		return ""
	}
	return string(data)
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

	if len(data) == 0 {
		return nil, errors.New("length is zero")
	}

	return data, nil
}

func HttpGetJson(url string, v interface{}) error {
	rep, err := http.Get(url)

	if err != nil {
		return err
	}

	defer rep.Body.Close()
	if rep.StatusCode != http.StatusOK {
		return errors.New("respone code error")
	}

	if err := json.NewDecoder(rep.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// contentType application/json | application/x-www-form-urlencoded
func HttpPost(url, contentType string, params interface{}) ([]byte, error) {
	var params_bytes []byte
	var err error
	if contentType == "application/json" {
		params_bytes, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
	}

	if contentType == "application/x-www-form-urlencoded" {
		params_bytes = []byte(params.(string))
	}

	res, err := http.Post(url, contentType, bytes.NewReader(params_bytes))
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

	if len(data) == 0 {
		return nil, errors.New("length is zero")
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
	part.Write(fileData)

	err = writer.Close()
	if err != nil {
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
