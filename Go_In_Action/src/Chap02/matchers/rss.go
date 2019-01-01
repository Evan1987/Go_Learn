package matchers

import (
	"encoding/xml"
)

type (
	// 尾部tag是为了序列化时反射使用的，在序列化时key值名称首字母必须大写，如果要小写必须靠tag
	// tag不能含空格
	item struct {
		XMLName 		xml.Name	`xml:"item"`
		PubDate			string 		`xml:"pubDate"`
		Title			string		`xml:"title"`
		Description		string		`xml:"description"`
		Link 			string		`xml:"link"`
		GUID 			string		`xml:"guid"`
		GeoRssPoint		string		`xml:"georss:point"`
	}

	image struct {
		XMLName			xml.Name 	`xml:"image"`
		URL 			string		`xml:"url"`
		Title 			string 		`xml:"title"`
		Link 			string 		`xml:"link"`
	}

	chann struct {
		XMLName 		xml.Name 	`xml:"channel"`
		Title 			string		`xml:"title"`
		Description		string 		`xml:"description"`
		Link 			string  	`xml:"link"`
		PubDate 		string 		`xml:"pubDate"`
		LastBuildDate	string 		`xml:"lastBuildDate"`
		TTL				string 		`xml:"ttl"`
		Language		string 		`xml:"language"`
		ManagingEditor	string		`xml:"managingEditor"`
		WebMaster		string 		`xml:"webMaster"`
		Image			image		`xml:"image"`
		item			[]item		`xml:"item"`
	}

	rssDocument struct {
		XMLName			xml.Name	`xml:"rss"`
		Channel 		chann		`xml:"chann"`
	}
)

type rssMatcher struct {}

func init(){
	var matcher rssMatcher

}

func (m rssMatcher) retrieve(feed *search.Feed)
