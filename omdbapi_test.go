package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	httpConfiguration := &HttpConfiguration{
		BaseURL:       "http://www.omdbapi.com/",
		TimeoutMillis: 120000,
		ApiKey:        "44446a96",
	}

	c := NewClient(httpConfiguration)

	res, err := c.Search("stem", nil)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, res, "expecting non-nil result")

	if res != nil {
		assert.Greater(t, len(res), 30, "expecting more than 30 results")
	}
	stringOne := "The STEM Journals"
	stringTwo := "Activision: STEM - in the Videogame Industry"

	foundStringOne := false
	foundStringTwo := false

	for _, items := range res {
		if strings.Contains(items.Title, stringOne) {
			foundStringOne = true
		}
		if strings.Contains(items.Title, stringTwo) {
			foundStringTwo = true
		}

		if foundStringTwo == true && foundStringOne == true {
			break
		}
	}

	assert.Equal(t, true, foundStringOne, "expecting string : The STEM Journals")
	assert.Equal(t, true, foundStringTwo, "Activision: STEM - in the Videogame Industry")

}
