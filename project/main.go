package main

import (
	"fmt"
	"kata/project/spider"
	"time"
	"math"
)

func main() {
	imgUrlCh := make(chan string, 5)
	errInfoCh := make(chan tumblr.ErrorInfo, 10000)
	for worker := 1; worker <= 5; worker++ {
		go tumblr.DownloadUrl(imgUrlCh, errInfoCh)
	}

	fmt.Println("start")
	imgUrls, err := tumblr.GetTumblrUrl("***", tumblr.TumblrPhoto, 50, math.MaxInt16)
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
			} else {
				fmt.Printf("download, url[%v]\n", info.Msg)
			}
		case <-time.After(60 * time.Second):
			fmt.Println("timeout")
		}
	}

	fmt.Println("success")
}
