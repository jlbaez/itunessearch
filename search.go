package itunessearch

import (
    "encoding/json"
    "io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var itunesSearchURl = "https://itunes.apple.com/search"

type SearchResponse struct {
	ResultCount int `bson:"resultCount"`
	Results []SearchResult `bson:"results"`
}

type SearchResult struct {
	ArtworkUrl30 string `json:"artworkUrl30" bson:"artworkUrl30"`
	ArtworkUrl60 string `json:"artworkUrl60" bson:"artworkUrl60"`
	ArtworkUrl100 string `json:"artworkUrl100" bson:"artworkUrl100"`
	ArtistName string `json:"artistName" bson:"artistName"`
	CollectionName string `json:"collectionName" bson:"collectionName"`
	CollectionId int `json:"collectionId" bson:"collectionId"`
	FeedUrl string `json:"feedUrl" bson:"feedUrl"`
}

func Search(query string, mediaType string, limit uint) ([]SearchResult, error)  {
	baseUrl, err := url.Parse(itunesSearchURl)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("term", query)
	params.Add("media", mediaType)
	params.Add("limit", string(limit))

	baseUrl.RawQuery = params.Encode()

	log.Printf("search url: %s", baseUrl)

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response SearchResponse

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil;
}
