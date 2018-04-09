package main

import (
	"fmt"
	"kata/project/spider"
	"time"
	"math"
)

func main() {
	imgUrlCh := make(chan string, 10)
	errInfoCh := make(chan tumblr.ErrorInfo, 10000)
	for worker := 1; worker <= 10; worker++ {
		go tumblr.DownloadUrl(imgUrlCh, errInfoCh)
	}

	fmt.Println("start")
	// verall80,smile67yt
	imgUrls, err := tumblr.GetTumblrUrl("jpnhotnurse", tumblr.TumblrPhoto, 0, math.MaxInt16)
	if err != nil {
		fmt.Printf("GetTumblrUrl failed [%v]\n", err)
		return
	}

	for i := 0; i < len(imgUrls); i++ {
		imgUrlCh <- imgUrls[i]
	}

	close(imgUrlCh)
	for i := 0; i < len(imgUrls); i++ {
		select {
		case info := <-errInfoCh:
			if info.Err != nil {
				fmt.Printf("download err[%v], url[%v]\n", info.Err, info.Msg)
			}
		case <-time.After(60 * time.Second):
			fmt.Println("timeout")
		}
	}

	fmt.Println("success")
}