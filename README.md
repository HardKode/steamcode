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