package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"net"
	"net/url"
	"os"
	"io"
	"runtime"
	"github.com/satori/go.uuid"
	"path"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}


//  http://***.tumblr.com/api/read?start=0?num=50
func main() {
	imgUrlCh := make(chan string)
	errInfoCh := make(chan ErrorInfo, 100)
	for worker := 1; worker <= 10; worker++ {
		go DownloadUrl(imgUrlCh, errInfoCh)
	}
	var count = 0

	ua := colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
	c := colly.NewCollector(ua)
	c.WithTransport(&http.Transport{
		Proxy: proxy,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.OnHTML("a img", func(e *colly.HTMLElement) {
		imgUrl := e.Attr("src")
		if strings.Contains(imgUrl, "media.tumblr.com") {
			imgUrlCh <- imgUrl
			count++
			fmt.Println("img url: ", imgUrl)
		}

		//e.Request.Visit(e.Attr("src"))
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://wanimal1983.tumblr.com/page/10")
	close(imgUrlCh)
	for i := 0; i < count; i++{
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
	return url.Parse("http://127.0.0.1:61350")
}

func init() {
	if runtime.GOOS == "windows" {
		downloadDir = "E:/abcde/"
	} else {
		downloadDir = "/data/zip/"
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

			fileName := downloadDir + uuid.NewV4().String() + ext
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