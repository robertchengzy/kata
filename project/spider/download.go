package tumblr

import (
	"time"
	"github.com/satori/go.uuid"
	"fmt"
	"os"
	"io"
	"strings"
	"path"
	"runtime"
	"net/url"
	"net/http"
)

var downloadDir string
var httpClient *http.Client
var proxy = func(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://127.0.0.1:58788")
}

func init() {
	if runtime.GOOS == "windows" {
		downloadDir = "E:/abcde/"
	} else {
		downloadDir = "/Users/cheng/robert/"
	}

	tr := &http.Transport{DisableKeepAlives: false, Proxy: proxy}
	httpClient = &http.Client{Timeout: 2 * time.Minute, Transport: tr}
}

type ErrorInfo struct {
	Err error
	Msg string
}

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
			res, err := httpClient.Get(dwUrl)
			if err != nil {
				return
			}
			ext := strings.Replace(path.Ext(dwUrl), "pnj", "png", -1)
			//uuidD, _ := uuid.NewV4()
			uuidD := uuid.NewV4()
			fileName := downloadDir + uuidD.String() + ext
			fmt.Println(fileName)
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
