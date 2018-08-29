package help

import "encoding/xml"

type VideoTumblr struct {
	XMLName xml.Name   `xml:"tumblr"`
	Posts   VideoPosts `xml:"posts"`
}

type VideoPosts struct {
	XMLName xml.Name    `xml:"posts"`
	Post    []VideoPost `xml:"post"`
	Start   string      `xml:"start,attr"`
	Total   string      `xml:"total,attr"`
}

type VideoPost struct {
	XMLName     xml.Name `xml:"post"`
	VideoPlayer string   `xml:"video-player"`
}
