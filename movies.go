package client

import (
	"fmt"
	"net/http"
)

// Search result
type SearchRes struct {
	Search       []SearchResItem `json:"Search"`
	TotalResults string          `json:"totalResults"`
	Response     string          `json:"Response"`
}

// Search result item
type SearchResItem struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

// Search Options
type SearchOptions struct {
	Page int `json:"page"`
}

func (c *HttpClient) Search(searchstring string, options *SearchOptions) (*SearchRes, error) {
	page := 1
	if options != nil {
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?s=%s&page=%d", c.HttpConfig.BaseURL, searchstring, page), nil)
	if err != nil {
		return nil, err
	}

	response := SearchRes{}
	if err := c.sendRequest(req, response); err != nil {
		return nil, err
	}

	return &response, nil

}
