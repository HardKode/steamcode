package client

import (
	"fmt"
	"net/http"
	"strconv"
)

// Could not find the data anywhere but it looked like the default page size is 10 items , top 10

const defaultPageSize int = 10

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

/*
Omdb API search function
searchstring : The string to reseaerch in the movies DB
options : this allows to select a specific page , if nil its all pages

returns : a list of SearchResItem
*/
func (c *HttpClient) Search(searchstring string, options *SearchOptions) ([]SearchResItem, error) {

	// Build base request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?s=%s", c.HttpConfig.BaseURL, searchstring), nil)
	if err != nil {
		return nil, err
	}
	//Add page if needed
	if options != nil && &options.Page != nil {
		query := req.URL.Query()
		query.Add("page", strconv.Itoa(options.Page))
		req.URL.RawQuery = query.Encode()
	}
	// fmt.Println(req.URL.String())

	response := SearchRes{}
	if err := c.sendRequest(req, &response); err != nil {
		//if the http call fails , bail out
		return nil, err
	}
	totalResultsToProcess, totalResultsToProcessErr := strconv.Atoi(response.TotalResults)
	if totalResultsToProcessErr != nil {
		fmt.Println(totalResultsToProcessErr)
		return nil, totalResultsToProcessErr
	}

	// At this stage response is considered valid , it was successfully decoded
	// var responseList []SearchResItem
	responseList := response.Search

	// if the user did NOT request a specific page , we can stop continue
	if options == nil {
		//if this is a full search , we will need to iterate
		// Even if the response is good, we have to consider pagination
		totalResults, totalResultsErr := strconv.Atoi(response.TotalResults)
		if totalResultsErr != nil {
			fmt.Println(totalResultsErr)
			return nil, totalResultsErr
		}
		//We have to consider the page size
		totalPages := totalResults / defaultPageSize
		if totalResults%defaultPageSize != 0 {
			totalPages++
		}
		fmt.Printf("page count -> : %d\n", totalPages)
		for pageIter := 2; pageIter <= totalPages; pageIter++ {
			// fmt.Printf(" page -> : %d\n", pageIter)
			//Query the page
			response := SearchRes{}
			req, err := http.NewRequest("GET", fmt.Sprintf("%s?s=%s&page=%d", c.HttpConfig.BaseURL, searchstring, pageIter), nil)
			if err != nil {
				fmt.Printf("Error building request for page: ")
				fmt.Println(err)
				return nil, err
			}

			if err := c.sendRequest(req, &response); err != nil {
				//if the http call fails , bail out
				fmt.Printf("Error sending request for page: ")
				fmt.Println(err)
				return nil, err
			}

			responseList = append(responseList, response.Search...)

		}

	}

	// fmt.Printf(" response : %+q\n", responseList)
	fmt.Printf("total response size: %d , total results %d\n", len(responseList), totalResultsToProcess)

	return responseList, nil

}
