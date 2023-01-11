package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFaces(t *testing.T) {

	httpConfiguration := &HttpConfiguration{
		BaseURL:       "http://www.omdbapi.com/",
		TimeoutMillis: 120000,
		ApiKey:        "44446a96",
	}

	searchOptions := SearchOptions{
		Page: 1,
	}
	c := NewClient(httpConfiguration)

	res, _ := c.Search("stem", &searchOptions)
	assert.NotNil(t, res, "expecting non-nil result")

	// if res != nil {
	// 	assert.Equal(t, 1, res.Count, "expecting 1 face found")
	// 	assert.Equal(t, 1, res.PagesCount, "expecting 1 PAGE found")

	// 	if res.Count > 0 {
	// 		assert.Equal(t, faceID, res.Faces[0].FaceID, "expecting correct face_id")
	// 		assert.NotEmpty(t, res.Faces[0].FaceToken, "expecting non-empty face_token")
	// 		assert.Greater(t, len(res.Faces[0].FaceImages), 0, "expecting non-empty face_images")

	// 		faceToken = res.Faces[0].FaceToken
	// 	}
	// }
}
