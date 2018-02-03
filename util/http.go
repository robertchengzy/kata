package util

import (
	. "jdk_bushu/logger"

	"io/ioutil"
	"net/http"
	"errors"
	"strconv"
	"bytes"
	"mime/multipart"
	"encoding/json"
	"os"
	"io"
)

func HttpUrlGet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		Logger.Warn("GetHttpUrl:[%v] failed:[%v] ", url, err)
		return ""
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		Logger.Warn("GetHttpUrl:[%v] resp code not ok: [%v]", url, res.StatusCode)
		return ""
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Logger.Warn("GetHttpUrl:[%v] failed:[%v] ", url, err)
		return ""
	}
	if len(data) == 0 {
		Logger.Warn("GetHttpUrl:[%v] failed length is zero ", url)
		return ""
	}
	return string(data)
}

func HttpGet(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		Logger.Warn("GetHttpUrl:[%v] failed:[%v] ", url, err)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		Logger.Warn("GetHttpUrl:[%v] resp code not ok: [%v]", url, res.StatusCode)
		return nil, errors.New("respone code error")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Logger.Warn("GetHttpUrl:[%v] failed:[%v] ", url, err)
		return nil, err
	}

	if len(data) == 0 {
		Logger.Warn("GetHttpUrl:[%v] failed length is zero ", url)
		return nil, errors.New("length is zero")
	}

	return data, nil
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

	Logger.Warn("params: [%v]", params)

	res, err := http.Post(url, contentType, bytes.NewReader(params_bytes))
	if err != nil {
		Logger.Warn("POST err:[%v]", err)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		Logger.Warn("HttpPost:[%v] resp code not ok: [%v]", url, res.StatusCode)
		return nil, errors.New("respone code error")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Logger.Warn("HttpPost:[%v] failed:[%v] ", url, err)
		return nil, err
	}

	if len(data) == 0 {
		Logger.Warn("GetHttpUrl:[%v] failed length is zero ", url)
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
		Logger.Info("POST Data: err[%v]", err)
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		Logger.Warn("response[%#v]", response)
		return nil, errors.New("response.StatusCode =" + strconv.Itoa(response.StatusCode))
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logger.Warn("resp data err [%v]", data)
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
