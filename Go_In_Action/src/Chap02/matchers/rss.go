package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"Chap02/search"
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
		Item			[]item		`xml:"item"`
	}

	rssDocument struct {
		XMLName			xml.Name	`xml:"rss"`
		Channel 		chann		`xml:"chann"`
	}
)

type rssMatcher struct {}

func init(){
	var matcher rssMatcher
	search.Register("rss", matcher)

}

// retrieve performs a Http Get request for the rss feed and decodes the results
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error){
	if feed.URI == "" {
		return nil, errors.New("no rss feed uri provided!")
	}

	// retrieve the rss feed document from the web
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// Close the response once we return from the function
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a proper response
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type
	// We don't need to check for errors, the caller can do this
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}


func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error){
	var results []*search.Result

	log.Printf("search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// Retrieve the data to search
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}


	for _, channelItem := range document.Channel.Item {

		// Check the title for the search term
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// if we found a match save the result
		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result
		if matched {
			results = append(results, &search.Result{
				Field: "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}
