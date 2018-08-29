package help

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

var DownloadDir string
var httpClient *http.Client
var proxy = func(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://127.0.0.1:58942")
}

func init() {
	if runtime.GOOS == "windows" {
		DownloadDir = "E:/abcde/"
	} else {
		DownloadDir = "/Users/cheng/robert/"
	}

	tr := &http.Transport{DisableKeepAlives: false, Proxy: proxy}
	httpClient = &http.Client{Timeout: 2 * time.Minute, Transport: tr}
}

type ErrorInfo struct {
	Err error
	Msg string
}

var count int64 = 0

func DownloadUrl(urlCh <-chan string, errInfoCh chan<- ErrorInfo) {
	for dwUrl := range urlCh {
		func() {
			var err error
			defer func() {
				errInfo := ErrorInfo{}
				errInfo.Msg = dwUrl
				errInfo.Err = err
				errInfoCh <- errInfo
			}()
			atomic.AddInt64(&count, 1)

			fileName := strings.Replace(DownloadDir+path.Base(dwUrl), "pnj", "png", -1)
			exists, err := PathExists(fileName)
			if err != nil {
				return
			}

			if exists {
				err = errors.New("exists")
				return
			}

			fmt.Println(count, fileName)

			res, err := httpClient.Get(dwUrl)
			if err != nil {
				return
			}

			file, err := os.Create(fileName)
			if err != nil {
				return
			}

			defer file.Close()
			_, err = io.Copy(file, res.Body)
			if err != nil {
				return
			}
			return
		}()
	}
}

// golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
// 如果返回的错误为nil,说明文件或文件夹存在
// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func PathRemoveExt(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename)
	fileSuffix := path.Ext(filenameWithSuffix)
	return strings.TrimSuffix(filenameWithSuffix, fileSuffix)
}
