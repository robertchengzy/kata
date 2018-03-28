package main

import (
	"fmt"
	"time"
	"net/http"
	"net/url"
	"os"
	"io"
	"runtime"
	"github.com/satori/go.uuid"
	"path"
	"strings"
	"errors"
	"encoding/json"
	"strconv"
	"encoding/xml"
)

func main() {
	imgUrlCh := make(chan string)
	errInfoCh := make(chan ErrorInfo, 100)
	for worker := 1; worker <= 10; worker++ {
		go DownloadUrl(imgUrlCh, errInfoCh)
	}
	var count = 0
	//imgUrlCh <- imgUrl
	count++

	close(imgUrlCh)
	for i := 0; i < count; i++ {
		select {
		case info := <-errInfoCh:
			if info.Err != nil {
				fmt.Printf("download err[%v], url[%v]\n", info.Err, info.Msg)
			}
		}
	}
}

var downloadDir string
var httpClient *http.Client
var proxy = func(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://127.0.0.1:58942") // 61350
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
			uuidD, _ := uuid.NewV4()
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

// http://***.tumblr.com/api/read?start=0&num=50&type=photo
const (
	TumblrUrl = "http://%s.tumblr.com/api/read?start=%d&num=%d&type=%s"
	TumblrPhoto = "photo"
)
func GetTumblrData(name, kind string, start, num int) (*Tumblr, error) {
	reqUrl := fmt.Sprintf(TumblrUrl, name, start, num, kind)

	resp, err := httpClient.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("respone code error " + strconv.Itoa(resp.StatusCode))
	}

	tumblr := new(Tumblr)
	if err := xml.NewDecoder(resp.Body).Decode(tumblr); err != nil {
		return nil, err
	}

	return tumblr, nil
}