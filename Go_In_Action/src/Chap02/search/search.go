package search

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)
const dataFile = "/Users/lixing/code projects/Go/Go_In_Action/src/Chap02/data/data.json"

type Result struct {
	Field string
	Content string
}

type Feed struct {
	Name 	string		`json:"site"`
	URI 	string 		`json:"link"`
	Type 	string 		`json:"type"`
}

type Matcher interface {
	Search(feed *Feed, searchTerm string)([]*Result, error)
}

// read feeds json
func RetrieveFeeds() ([] *Feed, error){
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}

// matcher is launched as a go-routine for each individual feed to run
func Match(matcher Matcher, feed *Feed, searchItem string, results chan *Result) {
	searchResults, err := matcher.Search(feed, searchItem)
	if err != nil{
		log.Println(err)
		return
	}

	// write the results to the channel
	for _, result := range searchResults {
		results <- result
	}
}

// display writes results to the console window as they are received by the individual routines
func Display(results chan *Result){
	for result := range results {
		log.Printf("%s: \n%s\n\n", result.Field, result.Content)
	}
}

var matchers = make(map[string]Matcher)
// Run performs the search logic
func Run(searchTerm string){
	// Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	// create an unbuffered channel to receive match results to display
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup


	// Set the number of go-routines we need to wait for while they process the individual feeds
	waitGroup.Add(len(feeds))

	// Launch a go-routine for each feed to find the results
	for _, feed := range feeds {
		// Retieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the go-routine to perform the search
		go func(matcher Matcher, feed *Feed){
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// launch a go-routine to monitor when all work is Done
	go func(){
		// wait for everything to be processed
		waitGroup.Wait()

		// Close the channel to signal to the Display function that we can exit the program
		close(results)
	}()

	// start displaying results as they are available and return after the final result is displayed!
	Display(results)
}

func Register(feedType string, matcher Matcher){
	if _, exists := matchers[feedType]; exists{
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
