package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

type wikiresponse struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue      struct {
		Sroffset int    `json:"sroffset"`
		Continue string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Searchinfo struct {
			Totalhits int `json:"totalhits"`
		} `json:"searchinfo"`
		Search []Search `json:"search"`
	} `json:"query"`
}

//Search result struct
type Search struct {
	Lang  string
	Title string `json:"title"`
	Size  int    `json:"size"`
}

// Request returns the top 20 results based on page size from
func Request(search string, lang string) []Search {
	var toReturn []Search

	var searchURL string = fmt.Sprintf("https://%s.wikipedia.org/w/api.php", lang)
	var searchQueryString string = fmt.Sprintf("?action=query&format=json&srlimit=100&prop=&list=search&meta=&srsearch=%s&srnamespace=0&sroffset=100&srprop=size&srsort=relevance", url.QueryEscape(search))

	//fmt.Println(searchURL + searchQueryString)

	response, err := http.Get(searchURL + searchQueryString)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var wr wikiresponse
		json.Unmarshal(data, &wr)
		toReturn = addlang(wr.Query.Search, lang)
	}
	return toReturn
}

// Print search results
func Print(search []Search) {
	for i := range search {

		fmt.Printf("Size:%d Title:%s Lang:%s \n", search[i].Size, search[i].Title, search[i].Lang)
	}
}

// Combine two sets of search results and return the top 20 based on size.
func Combine(searchOne []Search, searchTwo []Search) []Search {
	var results []Search
	results = appendSearch(results, searchOne)
	results = appendSearch(results, searchTwo)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Size > results[j].Size
	})
	if len(results) < 20 {
		return results[0:len(results)]
	}
	return results[0:20]
}

func appendSearch(base []Search, added []Search) []Search {
	for i := range added {
		base = append(base, added[i])
	}
	return base
}

func addlang(search []Search, lang string) []Search {
	for i := range search {
		search[i].Lang = lang
	}
	return search
}
