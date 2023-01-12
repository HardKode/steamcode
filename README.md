# stemcode

### install go on the system
https://go.dev/dl/
In my case i used brew , you can also download the arm darwin version : go1.19.5.darwin-arm64.pkg
```
toolgocollab@Tool-Gos-iMac stemcode % go version
go version go1.19.5 darwin/arm64
```

# add modules testify and logrus
go get github.com/stretchr/testify/assert
go get "github.com/Sirupsen/logrus"

# run the tests
# basically we built go tests 
while  being at the root of this directory , run command :
APIKEY='xxxx' go test -v --tags=all

VERY IMPORTANT : The api  key is not hardcoded. it expects an environment variable called APIKEY
you can also set it : declare -x APIKEY='xxx'
or put in your bash profile, etc ...

I built two test files :
omdbapi_test.go : really basic go test with subtest
omdbapi_tabletest_test.go : this test file uses a table mechanism which is a bit more separated and can be extended with data
They all distinguished by tags 

so to run omdbapi_test.go who has tab simpletest: go test -v --tags=simpletest
so to run omdbapi_tabletest_test.go who has tab tabletest: go test -v --tags=tabletest
to run them all : go test -v --tags=all

```
go test -v --tags=tabletest
=== RUN   TestSearchTableDrivenExtended
=== RUN   TestSearchTableDrivenExtended/2-_from_search_nethod_,_get_Imdb_of_movie_titledActivision:_STEM_-_in_the_Videogame_Industry
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:16:11-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:16:11-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:16:12-04:00","totalResultsToProcess":55}
=== RUN   TestSearchTableDrivenExtended/3-_Use_get_movie_titled_The_STEM_Journals
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:16:12-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:16:12-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:16:12-04:00","totalResultsToProcess":55}
=== RUN   TestSearchTableDrivenExtended/1-_search_stem_string_confirm_more_than_30_items_,_correcy_title
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:16:12-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:16:12-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:16:12-04:00","totalResultsToProcess":55}
--- PASS: TestSearchTableDrivenExtended (1.00s)
    --- PASS: TestSearchTableDrivenExtended/2-_from_search_nethod_,_get_Imdb_of_movie_titledActivision:_STEM_-_in_the_Videogame_Industry (0.35s)
    --- PASS: TestSearchTableDrivenExtended/3-_Use_get_movie_titled_The_STEM_Journals (0.26s)
    --- PASS: TestSearchTableDrivenExtended/1-_search_stem_string_confirm_more_than_30_items_,_correcy_title (0.39s)
PASS
ok      simple_api_client       1.427s
```


```
toolgocollab@Tool-Gos-iMac stemcode % go test -v --tags=simpletest
=== RUN   TestSearch
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:14:53-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:14:57-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:14:57-04:00","totalResultsToProcess":55}
=== RUN   TestSearch/Assert_that_the_result_should_contain_at_least_30_items
=== RUN   TestSearch/Assert_that_the_result_contains:_The_STEM_Journal_,Activision:_STEM_-_in_the_Videogame_Industry_
=== RUN   TestSearch/Get_movie_by_Id(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_id_
=== RUN   TestSearch/Get_movie_by__Title__(The_STEM_Journals)_detail_using_get_by_title_
--- PASS: TestSearch (3.63s)
    --- PASS: TestSearch/Assert_that_the_result_should_contain_at_least_30_items (0.00s)
    --- PASS: TestSearch/Assert_that_the_result_contains:_The_STEM_Journal_,Activision:_STEM_-_in_the_Videogame_Industry_ (0.00s)
    --- PASS: TestSearch/Get_movie_by_Id(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_id_ (0.03s)
    --- PASS: TestSearch/Get_movie_by__Title__(The_STEM_Journals)_detail_using_get_by_title_ (0.03s)
PASS
ok      simple_api_client       3.865s

```
```
go test -v --tags=all
=== RUN   TestSearchTableDrivenExtended
=== RUN   TestSearchTableDrivenExtended/1-_search_stem_string_confirm_more_than_30_items_,_correcy_title
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:11:30-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:11:31-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:11:31-04:00","totalResultsToProcess":55}
=== RUN   TestSearchTableDrivenExtended/2-_from_search_nethod_,_get_Imdb_of_movie_titledActivision:_STEM_-_in_the_Videogame_Industry
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:11:31-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:11:31-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:11:31-04:00","totalResultsToProcess":55}
=== RUN   TestSearchTableDrivenExtended/3-_Use_get_movie_titled_The_STEM_Journals
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:11:31-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:11:31-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:11:31-04:00","totalResultsToProcess":55}
--- PASS: TestSearchTableDrivenExtended (1.01s)
    --- PASS: TestSearchTableDrivenExtended/1-_search_stem_string_confirm_more_than_30_items_,_correcy_title (0.27s)
    --- PASS: TestSearchTableDrivenExtended/2-_from_search_nethod_,_get_Imdb_of_movie_titledActivision:_STEM_-_in_the_Videogame_Industry (0.42s)
    --- PASS: TestSearchTableDrivenExtended/3-_Use_get_movie_titled_The_STEM_Journals (0.32s)
=== RUN   TestSearch
{"level":"info","msg":"http://www.omdbapi.com/?s=stem","time":"2023-01-12T00:11:31-04:00"}
{"level":"info","msg":"page count/info:","pagecount":6,"time":"2023-01-12T00:11:32-04:00","totalResults":55}
total response size: 55 , total results 55
{"level":"info","msg":"page count/info:","responsesize":55,"time":"2023-01-12T00:11:32-04:00","totalResultsToProcess":55}
=== RUN   TestSearch/Assert_that_the_result_should_contain_at_least_30_items
=== RUN   TestSearch/Assert_that_the_result_contains:_The_STEM_Journal_,Activision:_STEM_-_in_the_Videogame_Industry_
=== RUN   TestSearch/Get_movie_by_Id(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_id_
=== RUN   TestSearch/Get_movie_by__Title__(The_STEM_Journals)_detail_using_get_by_title_
--- PASS: TestSearch (0.31s)
    --- PASS: TestSearch/Assert_that_the_result_should_contain_at_least_30_items (0.00s)
    --- PASS: TestSearch/Assert_that_the_result_contains:_The_STEM_Journal_,Activision:_STEM_-_in_the_Videogame_Industry_ (0.00s)
    --- PASS: TestSearch/Get_movie_by_Id(Activision:_STEM_-_in_the_Videogame_Industry)_detail_using_get_by_id_ (0.05s)
    --- PASS: TestSearch/Get_movie_by__Title__(The_STEM_Journals)_detail_using_get_by_title_ (0.03s)
PASS
ok      simple_api_client       1.566s
```

### code

client.go : contains the client code . its a wrapper around the http library
movies.go : contains the function used to research the movie database . Uses the client
omdbapi_test.go : this has the tests methods used to exercise the client

# logging 
 just put logrus in some spots





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