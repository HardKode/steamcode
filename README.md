# stemcode

### install go on the system
https://go.dev/dl/
In my case i used brew , you can also download the arm darwin version : go1.19.5.darwin-arm64.pkg
```
toolgocollab@Tool-Gos-iMac stemcode % go version
go version go1.19.5 darwin/arm64
```

# ad testify 
go get github.com/stretchr/testify/assert
go get "github.com/Sirupsen/logrus"

# run the tests
while  being at the root of this directory
APIKEY='xxxx' go test -v

VERY IMPORTANT : The api  key is not hardcoded. it expects an environment variable called APIKEY
you can also set it : declare -x APIKEY='xxx'
or put in your bash profile, etc ...

```
stemcode % go test -v
=== RUN   TestSearch
page count -> : 6
total response size: 55 , total results 55
=== RUN   TestSearch/Assert_that_the_result_should_contain_at_least_30_items
=== RUN   TestSearch/Assert_that_the_result_contains:_The_STEM_Journal_,Activision:_STEM_-_in_the_Videogame_Industry_
=== RUN   TestSearch/Get_movie_by_Id(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_id_
=== RUN   TestSearch/Get_movie_by__Title__(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_title_
--- PASS: TestSearch (0.53s)
    --- PASS: TestSearch/Assert_that_the_result_should_contain_at_least_30_items (0.00s)
    --- PASS: TestSearch/Assert_that_the_result_contains:_The_STEM_Journal_,Activision:_STEM_-_in_the_Videogame_Industry_ (0.00s)
    --- PASS: TestSearch/Get_movie_by_Id(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_id_ (0.08s)
    --- PASS: TestSearch/Get_movie_by__Title__(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_title_ (0.11s)
PASS
ok      simple_api_client       0.755s
```

### code

client.go : contains the client code . its a wrapper around the http library
movies.go : contains the function used to research the movie database . Uses the client
omdbapi_test.go : this has the tests methods used to exercise the client





# note about the test
Using get_by_title method, get item by title The STEM Journals and assert that the plot contains the string Science, Technology, Engineering and Math and has a runtime of 22 minutes.

this was failing because of the case sentivity : S of science,= . Not sure if that was intentional
=== CONT  TestSearch
    omdbapi_test.go:91: 
                Error Trace:    /Users/toolgocollab/stemcode/omdbapi_test.go:91
                Error:          "\"The STEM Journals\" follows hosts Brad Piccirillo as they explore new discoveries and exciting careers in the fields of science, technology, engineering and math." does not contain "Science, Technology, Engineering and Math"
                Test:           TestSearch
                Messages:       expecting the correct content in the plot

solution : fix expected string to lowercase s because its within a sentence too