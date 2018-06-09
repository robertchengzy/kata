package tumblr

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"errors"
	"net/http"
	"os"
)

type PhotoTumblr struct {
	XMLName xml.Name   `xml:"tumblr"`
	Posts   PhotoPosts `xml:"posts"`
}

type PhotoPosts struct {
	XMLName xml.Name    `xml:"posts"`
	Post    []PhotoPost `xml:"post"`
	Start   string      `xml:"start,attr"`
	Total   string      `xml:"total,attr"`
}

type PhotoPost struct {
	XMLName  xml.Name `xml:"post"`
	PhotoSet PhotoSet `xml:"photoset"`
}

type PhotoSet struct {
	XMLName xml.Name `xml:"photoset"`
	Photo   []Photo  `xml:"photo"`
}

type Photo struct {
	XMLName  xml.Name   `xml:"photo"`
	PhotoUrl []PhotoUrl `xml:"photo-url"`
}

type PhotoUrl struct {
	XMLName   xml.Name `xml:"photo-url"`
	MaxWidth  string   `xml:"max-width,attr"`
	InnerText string   `xml:",innerxml"`
}

// http://***.tumblr.com/api/read?start=0&num=50&type=photo
// http://***.tumblr.com/api/read/json?start=0&num=50&type=photo
const (
	APIUrl   = "http://%s.tumblr.com/api/read?start=%d&num=%d&type=%s"
	APIPhoto = "photo"
	APIVideo = "video"
)

func GetTumblrData(name, kind string, start, num int) (*PhotoTumblr, error) {
	reqUrl := fmt.Sprintf(APIUrl, name, start, num, kind)

	resp, err := httpClient.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("respone code error " + strconv.Itoa(resp.StatusCode))
	}

	tumblr := new(PhotoTumblr)
	if err := xml.NewDecoder(resp.Body).Decode(tumblr); err != nil {
		return nil, err
	}

	return tumblr, nil
}

func GetTumblrUrl(name, kind string, start, num int) ([]string, error) {
	if name == "" {
		return nil, errors.New("name empty")
	}

	data, err := GetTumblrData(name, kind, start, num)
	if err != nil {
		return nil, fmt.Errorf("GetTumblrData failed[%v]", err)
	}

	if data.Posts.Total == "0" {
		return nil, errors.New("tubmlr data empty")
	}

	fmt.Println("total=", data.Posts.Total)
	// TODO 根据total查找全部数据

	urls := make([]string, 0, 10)
	for i := 0; i < len(data.Posts.Post); i++ {
		for j := 0; j < len(data.Posts.Post[i].PhotoSet.Photo); j++ {
			urls = append(urls, data.Posts.Post[i].PhotoSet.Photo[j].PhotoUrl[0].InnerText)
		}
	}

	if len(urls) == 0 {
		return nil, errors.New("tubmlr data url empty")
	}

	DownloadDir = DownloadDir + name + "/"
	os.MkdirAll(DownloadDir, os.ModePerm)

	fmt.Println("urls=", len(urls))

	return urls, nil
}
