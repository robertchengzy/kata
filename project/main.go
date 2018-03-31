package main

import (
	"fmt"
	"kata/project/spider"
)

func main() {
	imgUrlCh := make(chan string)
	errInfoCh := make(chan tumblr.ErrorInfo, 100)
	for worker := 1; worker <= 10; worker++ {
		go tumblr.DownloadUrl(imgUrlCh, errInfoCh)
	}

	fmt.Println("start")
	imgUrls, err := tumblr.GetTumblrUrl("anzhitinglan")
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
		}
	}

	fmt.Println("success")
}
