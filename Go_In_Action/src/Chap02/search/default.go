package search

type defaultMatcher struct {}


// init registers the default matcher with the program
func init(){
	var matcher defaultMatcher
	Register("default", matcher)
}

// search implements the behavior for the default matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error){
	return nil, nil
}