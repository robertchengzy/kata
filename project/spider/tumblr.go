package main

import (
	"encoding/xml"
	"fmt"
)

type Tumblr struct {
	XMLName xml.Name `xml:"tumblr"`
	Posts   Posts    `xml:"posts"`
}

type Posts struct {
	XMLName xml.Name `xml:"posts"`
	Post    []Photo  `xml:"post"`
}

type Post struct {
	XMLName  xml.Name `xml:"posts"`
	PhotoSet PhotoSet  `xml:"photoset"`
}

type PhotoSet struct {
	XMLName xml.Name `xml:"photoset"`
	Post    []Photo  `xml:"post"`
}
type Photo struct {
	XMLName  xml.Name   `xml:"photoset"`
	PhotoUrl []PhotoUrl `xml:"photo-url"`
}

type PhotoUrl struct {
	XMLName   xml.Name `xml:"photo-url"`
	MaxWidth  string   `xml:"max-width,attr"`
	InnerText string   `xml:",innerxml"`
}

func main() {
	dataXml := `<?xml version="1.0" encoding="UTF-8"?>
<tumblr version="1.0"><tumblelog name="mrsimpleing" timezone="US/Eastern" title="Fxxk me">&lt;p&gt;&lt;strong&gt;nice to fuck you&lt;/strong&gt;&lt;/p&gt;</tumblelog><posts type="photo" start="0" total="465"><post id="170840626163" url="https://mrsimpleing.tumblr.com/post/170840626163" url-with-slug="https://mrsimpleing.tumblr.com/post/170840626163/joenemesis-鮮肉女神夢心玥黑色高叉泳衣原版" type="photo" date-gmt="2018-02-13 18:17:31 GMT" date="Tue, 13 Feb 2018 13:17:31" unix-timestamp="1518545851" format="html" reblog-key="5ECKFdIr" slug="joenemesis-鮮肉女神夢心玥黑色高叉泳衣原版" note-count="1654" reblogged-from-url="https://joenemesis.tumblr.com/post/163473575340/鮮肉女神夢心玥黑色高叉泳衣原版" reblogged-from-name="joenemesis" reblogged-from-title="joenemesis" reblogged-from-avatar-url-16="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_16.png" reblogged-from-avatar-url-24="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_24.png" reblogged-from-avatar-url-30="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_30.png" reblogged-from-avatar-url-40="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_40.png" reblogged-from-avatar-url-48="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_48.png" reblogged-from-avatar-url-64="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_64.png" reblogged-from-avatar-url-96="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_96.png" reblogged-from-avatar-url-128="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_128.png" reblogged-from-avatar-url-512="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_512.png" reblogged-root-url="https://joenemesis.tumblr.com/post/163473575340/鮮肉女神夢心玥黑色高叉泳衣原版" reblogged-root-name="joenemesis" reblogged-root-title="joenemesis" reblogged-root-avatar-url-16="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_16.png" reblogged-root-avatar-url-24="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_24.png" reblogged-root-avatar-url-30="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_30.png" reblogged-root-avatar-url-40="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_40.png" reblogged-root-avatar-url-48="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_48.png" reblogged-root-avatar-url-64="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_64.png" reblogged-root-avatar-url-96="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_96.png" reblogged-root-avatar-url-128="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_128.png" reblogged-root-avatar-url-512="https://78.media.tumblr.com/avatar_3a9db5e9ca0e_512.png" width="2832" height="4240"><tumblelog title="Fxxk me" name="mrsimpleing" url="https://mrsimpleing.tumblr.com/" timezone="US/Eastern" avatar-url-16="https://78.media.tumblr.com/avatar_d3390a90f0ed_16.png" avatar-url-24="https://78.media.tumblr.com/avatar_d3390a90f0ed_24.png" avatar-url-30="https://78.media.tumblr.com/avatar_d3390a90f0ed_30.png" avatar-url-40="https://78.media.tumblr.com/avatar_d3390a90f0ed_40.png" avatar-url-48="https://78.media.tumblr.com/avatar_d3390a90f0ed_48.png" avatar-url-64="https://78.media.tumblr.com/avatar_d3390a90f0ed_64.png" avatar-url-96="https://78.media.tumblr.com/avatar_d3390a90f0ed_96.png" avatar-url-128="https://78.media.tumblr.com/avatar_d3390a90f0ed_128.png" avatar-url-512="https://78.media.tumblr.com/avatar_d3390a90f0ed_512.png"/><photo-caption>&lt;p&gt;&lt;a href="https://joenemesis.tumblr.com/post/163473575340/%E9%AE%AE%E8%82%89%E5%A5%B3%E7%A5%9E%E5%A4%A2%E5%BF%83%E7%8E%A5%E9%BB%91%E8%89%B2%E9%AB%98%E5%8F%89%E6%B3%B3%E8%A1%A3%E5%8E%9F%E7%89%88" class="tumblr_blog"&gt;joenemesis&lt;/a&gt;:&lt;/p&gt;
&lt;blockquote&gt;&lt;p&gt;鮮肉女神夢心玥黑色高叉泳衣原版&lt;/p&gt;&lt;/blockquote&gt;</photo-caption><photo-url max-width="1280">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_75sq.jpg</photo-url><photoset><photo offset="o6" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/9ee423752943def8ed2f4df92033d0e1/tumblr_otqexuAAGf1um2x2oo6_75sq.jpg</photo-url></photo><photo offset="o5" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/f188bb23b702cde9576e8a5d1580aa31/tumblr_otqexuAAGf1um2x2oo5_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/f188bb23b702cde9576e8a5d1580aa31/tumblr_otqexuAAGf1um2x2oo5_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/f188bb23b702cde9576e8a5d1580aa31/tumblr_otqexuAAGf1um2x2oo5_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/f188bb23b702cde9576e8a5d1580aa31/tumblr_otqexuAAGf1um2x2oo5_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/f188bb23b702cde9576e8a5d1580aa31/tumblr_otqexuAAGf1um2x2oo5_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/f188bb23b702cde9576e8a5d1580aa31/tumblr_otqexuAAGf1um2x2oo5_75sq.jpg</photo-url></photo><photo offset="o7" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/ea897c6ff1c780d4d3cca5494a018718/tumblr_otqexuAAGf1um2x2oo7_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/ea897c6ff1c780d4d3cca5494a018718/tumblr_otqexuAAGf1um2x2oo7_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/ea897c6ff1c780d4d3cca5494a018718/tumblr_otqexuAAGf1um2x2oo7_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/ea897c6ff1c780d4d3cca5494a018718/tumblr_otqexuAAGf1um2x2oo7_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/ea897c6ff1c780d4d3cca5494a018718/tumblr_otqexuAAGf1um2x2oo7_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/ea897c6ff1c780d4d3cca5494a018718/tumblr_otqexuAAGf1um2x2oo7_75sq.jpg</photo-url></photo><photo offset="o4" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/c25727af5948585cb890a20d8cd8ca93/tumblr_otqexuAAGf1um2x2oo4_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/c25727af5948585cb890a20d8cd8ca93/tumblr_otqexuAAGf1um2x2oo4_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/c25727af5948585cb890a20d8cd8ca93/tumblr_otqexuAAGf1um2x2oo4_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/c25727af5948585cb890a20d8cd8ca93/tumblr_otqexuAAGf1um2x2oo4_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/c25727af5948585cb890a20d8cd8ca93/tumblr_otqexuAAGf1um2x2oo4_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/c25727af5948585cb890a20d8cd8ca93/tumblr_otqexuAAGf1um2x2oo4_75sq.jpg</photo-url></photo><photo offset="o2" caption="" width="4240" height="2832"><photo-url max-width="1280">https://78.media.tumblr.com/0207f7d4faeb6538ec40b68fee9d20f0/tumblr_otqexuAAGf1um2x2oo2_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/0207f7d4faeb6538ec40b68fee9d20f0/tumblr_otqexuAAGf1um2x2oo2_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/0207f7d4faeb6538ec40b68fee9d20f0/tumblr_otqexuAAGf1um2x2oo2_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/0207f7d4faeb6538ec40b68fee9d20f0/tumblr_otqexuAAGf1um2x2oo2_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/0207f7d4faeb6538ec40b68fee9d20f0/tumblr_otqexuAAGf1um2x2oo2_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/0207f7d4faeb6538ec40b68fee9d20f0/tumblr_otqexuAAGf1um2x2oo2_75sq.jpg</photo-url></photo><photo offset="o1" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/f18d851e058d3e42a98677030f1dc25e/tumblr_otqexuAAGf1um2x2oo1_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/f18d851e058d3e42a98677030f1dc25e/tumblr_otqexuAAGf1um2x2oo1_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/f18d851e058d3e42a98677030f1dc25e/tumblr_otqexuAAGf1um2x2oo1_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/f18d851e058d3e42a98677030f1dc25e/tumblr_otqexuAAGf1um2x2oo1_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/f18d851e058d3e42a98677030f1dc25e/tumblr_otqexuAAGf1um2x2oo1_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/f18d851e058d3e42a98677030f1dc25e/tumblr_otqexuAAGf1um2x2oo1_75sq.jpg</photo-url></photo><photo offset="o8" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/feb7e7ae2df22395d7e684f58bfc3272/tumblr_otqexuAAGf1um2x2oo8_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/feb7e7ae2df22395d7e684f58bfc3272/tumblr_otqexuAAGf1um2x2oo8_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/feb7e7ae2df22395d7e684f58bfc3272/tumblr_otqexuAAGf1um2x2oo8_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/feb7e7ae2df22395d7e684f58bfc3272/tumblr_otqexuAAGf1um2x2oo8_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/feb7e7ae2df22395d7e684f58bfc3272/tumblr_otqexuAAGf1um2x2oo8_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/feb7e7ae2df22395d7e684f58bfc3272/tumblr_otqexuAAGf1um2x2oo8_75sq.jpg</photo-url></photo><photo offset="o9" caption="" width="2832" height="4240"><photo-url max-width="1280">https://78.media.tumblr.com/635b18fc355e5dd77da393b49450ba92/tumblr_otqexuAAGf1um2x2oo9_1280.jpg</photo-url><photo-url max-width="500">https://78.media.tumblr.com/635b18fc355e5dd77da393b49450ba92/tumblr_otqexuAAGf1um2x2oo9_500.jpg</photo-url><photo-url max-width="400">https://78.media.tumblr.com/635b18fc355e5dd77da393b49450ba92/tumblr_otqexuAAGf1um2x2oo9_400.jpg</photo-url><photo-url max-width="250">https://78.media.tumblr.com/635b18fc355e5dd77da393b49450ba92/tumblr_otqexuAAGf1um2x2oo9_250.jpg</photo-url><photo-url max-width="100">https://78.media.tumblr.com/635b18fc355e5dd77da393b49450ba92/tumblr_otqexuAAGf1um2x2oo9_100.jpg</photo-url><photo-url max-width="75">https://78.media.tumblr.com/635b18fc355e5dd77da393b49450ba92/tumblr_otqexuAAGf1um2x2oo9_75sq.jpg</photo-url></photo></photoset></post></posts></tumblr>
`
	var tumblrData = new(Tumblr)
	if err := xml.Unmarshal([]byte(dataXml), tumblrData); err != nil {
		fmt.Printf("xml unmarshal err [%v]\n", err)
		return
	}
	data, _ := xml.Marshal(tumblrData)
	fmt.Println(string(data))
}
