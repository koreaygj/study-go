package main

import(
	"fmt"
	"net/http"
)
type requestResult struct {
	url string
	status string
}
func main(){
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.aindsfadfsafdsafb.com/",
		"https://www.koreaygj.github.io",
		"https://www.netflix.com",
		"https://www.youtube.com",
	}
	for _, url := range urls{
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++{
		result := <- c
		results[result.url] = result.status
	}
	for url, status := range results{
		fmt.Println(url, status)
	}
}
func hitURL(url string, c chan<- requestResult)  {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "Ok"
	if(err != nil) || resp.StatusCode >= 400{
		status = " Requeset failed"
	}
	c <- requestResult{url: url, status: status}
}