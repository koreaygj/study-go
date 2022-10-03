package main

import(
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}
func Scrape(term string){
	var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 1; i <= totalPages; i++{
		go getPage(i, baseUrl, c)
	}
	for i := 0; i <= totalPages; i++{
		extractedJobs := <- c
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
}
func writeJobs(jobs []extractedJob){
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"LInk", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)
	for _, job := range jobs{
		jobSlice := []string{"http://kr.indee.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}
func getPage(page int, url string, mainC chan<- []extractedJob)  {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++{
		job := <- c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}
func getPages(url string) int{
	pages := 0
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
    bytes, _ := ioutil.ReadAll(resp.Body)
    str := string(bytes) //바이트를 문자열로
    fmt.Println(str)
	checkCode(resp)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
	pages = s.Find("a").Length()
	})
	return pages
	}
func checkErr(err error){
	if err != nil{
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response){
	if res.StatusCode != 200{
		log.Fatalln("Request failed with Status:", res.StatusCode)	
	}
}
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func CleanString(str string) string{
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}