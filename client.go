package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpConfiguration struct {
	BaseURL       string `yaml:"BaseURL,omitempty"`
	TimeoutMillis int    `yaml:"TimeoutMillis,omitempty"`
	ApiKey        string `yaml:"ApiKey,omitempty"`
}

type HttpClient struct {
	Impl       *http.Client
	HttpConfig *HttpConfiguration
}

func NewClient(httpConfig *HttpConfiguration) *HttpClient {

	client := &http.Client{}
	client.Timeout = time.Duration(httpConfig.TimeoutMillis) * time.Millisecond
	return &HttpClient{
		Impl:       client,
		HttpConfig: httpConfig,
	}

}

// We have noticed the API returns always 2 types of response

// one when it couldnt find anything
type ErrorResponse struct {
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

// one when it has some data to return
type SuccessResponse struct {
	Search       interface{} `json:"Search"`
	TotalResults string      `json:"totalResults"`
	Response     string      `json:"Response"`
}

// good or bad , there is always a response
type MinimalResponse struct {
	Response string `json:"Response"`
}

// Client function to send the request

func (c *HttpClient) sendRequest(req *http.Request, responseInterface interface{}) error {

	//add api key
	query := req.URL.Query()
	query.Add("apikey", c.HttpConfig.ApiKey)
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())

	// Send the request
	resp, err := c.Impl.Do(req)
	if err != nil {
		return err
	}

	//Parse the response body
	parsedBody, parseErr := c.ParseBody(resp.Body)
	if parseErr != nil {
		fmt.Println("failed to parse response body")
		return errors.New(fmt.Sprintf("failed to parse response body: \n%s", parseErr))
	}
	fmt.Println("Reponse parsed successfully")

	//Keep a buffer copy since  the buffer drains after each decode
	buf_minimal_response := bytes.NewBuffer(parsedBody)
	defer resp.Body.Close()

	// probably not much to do do if you dont get a 200
	if resp.StatusCode != http.StatusOK {
		//always expect a 200
		fmt.Printf("the API did not return 200: %d\n", resp.StatusCode)
		fmt.Printf("the API did not return 200:body:%s\n", string(parsedBody))
		return errors.New(fmt.Sprintf("the API did not return 200"))
	}

	//process the response
	// for some reason , whether the search is successfull or not , we get a 200 not 404 as I would expect
	// what distinguishes the 2 cases is the existence of the Error attribute
	// we have to parse to know that unfortunately . Would have been nice to be able to just rely on error code

	var minimalResponse MinimalResponse
	if err := json.NewDecoder(buf_minimal_response).Decode(&minimalResponse); err != nil {
		fmt.Println("Unknown error from the api")
		return err
	}
	fmt.Printf("Mininal response : %v\n", minimalResponse)

	// Now we need to decode the returned entities. That is error or success . The Response field tells us which one
	if minimalResponse.Response != "True" {
		var errorResponse ErrorResponse
		buf_error_response := bytes.NewBuffer(parsedBody)
		if err = json.NewDecoder(buf_error_response).Decode(&errorResponse); err != nil {
			fmt.Println("***Unknown errResponse from the api")
			fmt.Println(err)
			fmt.Printf("Error response : %v\n", errorResponse)
			return err
		}
		return errors.New(errorResponse.Error)
	}

	successResponse := SuccessResponse{
		Search: responseInterface,
	}

	buf_full_response := bytes.NewBuffer(parsedBody)
	if err = json.NewDecoder(buf_full_response).Decode(&successResponse); err != nil {
		fmt.Println(err)
		fmt.Println("Error decoding success response from the api")
		return err
	}

	// fmt.Printf(" response : %v\n", successResponse)
	return nil

}

// This function Unmarshalls the slice of bytes into a generic struct
func Extract[T any](body []byte) (*T, error) {
	o := new(T)
	err := json.Unmarshal(body, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (c *HttpClient) ParseBody(body io.ReadCloser) ([]byte, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
