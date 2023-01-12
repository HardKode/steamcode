//go:build tabletest || all
// +build tabletest all

package client

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTableDrivenExtended(t *testing.T) {

	for scenario, fn := range map[string]func(t *testing.T){
		"1- search stem string confirm more than 30 items , correcy title":                             testSearch,
		"2- from search nethod , get Imdb of movie titledActivision: STEM - in the Videogame Industry": testGetById,
		"3- Use get movie titled The STEM Journals":                                                    testGetByTitle,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testSearch(t *testing.T) {

	httpConfiguration := &HttpConfiguration{
		BaseURL:       "http://www.omdbapi.com/",
		TimeoutMillis: 120000,
		ApiKey:        os.Getenv("APIKEY"),
	}

	c := NewClient(httpConfiguration)

	res, err := c.Search("stem", nil)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, res, "expecting non-nil result")

	assert.Greater(t, len(res), 30, "expecting more than 30 results")

	stringOne := "The STEM Journals"
	stringTwo := "Activision: STEM - in the Videogame Industry"

	foundStringOne := false
	foundStringTwo := false

	for _, item := range res {
		if strings.Contains(item.Title, stringOne) {
			foundStringOne = true
			// fmt.Println(item)
		}
		if strings.Contains(item.Title, stringTwo) {
			foundStringTwo = true
			// fmt.Println(item)
		}

		if foundStringTwo == true && foundStringOne == true {
			break
		}
	}

	assert.Equal(t, true, foundStringOne, "expecting string : The STEM Journals")
	assert.Equal(t, true, foundStringTwo, "Activision: STEM - in the Videogame Industry")
}

func testGetById(t *testing.T) {

	httpConfiguration := &HttpConfiguration{
		BaseURL:       "http://www.omdbapi.com/",
		TimeoutMillis: 120000,
		ApiKey:        os.Getenv("APIKEY"),
	}

	c := NewClient(httpConfiguration)

	res, err := c.Search("stem", nil)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, res, "expecting non-nil result")

	stringTwo := "Activision: STEM - in the Videogame Industry"
	stringReleaseDate := "23 Nov 2010"
	stringDirector := "Mike Feurstein"
	foundStringTwo := false

	var foundId string
	for _, item := range res {
		if strings.Contains(item.Title, stringTwo) {
			foundStringTwo = true
			foundId = item.ImdbID
			// fmt.Printf("id:%s", foundId)
		}
		if foundStringTwo == true {
			break
		}
	}
	assert.Equal(t, true, foundStringTwo, "Activision: STEM - in the Videogame Industry")
	resitem, err := c.get_by_id(foundId)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, resitem, "expecting non-nil result")

	assert.Equal(t, resitem.Released, stringReleaseDate, "expecting the correct release date")
	assert.Equal(t, resitem.Director, stringDirector, "expecting the correct director ")
}

func testGetByTitle(t *testing.T) {

	httpConfiguration := &HttpConfiguration{
		BaseURL:       "http://www.omdbapi.com/",
		TimeoutMillis: 120000,
		ApiKey:        os.Getenv("APIKEY"),
	}

	c := NewClient(httpConfiguration)

	res, err := c.Search("stem", nil)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, res, "expecting non-nil result")

	searchTitle := "The STEM Journals"
	restitle, err := c.get_by_title(searchTitle)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, restitle, "expecting non-nil result")

	// check readme for why we lower case .
	plotstring := strings.ToLower("Science, Technology, Engineering and Math")
	durationstring := "22 min"

	assert.Contains(t, restitle.Plot, plotstring, "expecting the correct content in the plot")
	assert.Equal(t, restitle.Runtime, durationstring, "expecting the correct runtime ")
}
