package main

import(
	"fmt"
	"errors"
	"net/http"
)
var errRequestFailed = errors.New("Request failed")

func main(){
	var results = make(map[string]string)
	urls := []string{
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.airbnb.com/",
	}
	for _, url := range urls{
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}
func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if(err != nil) || resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}