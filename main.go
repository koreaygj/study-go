package main
import(
	"log"
	"net/http"
	"fmt"
	"strconv"
	"string"
	"github.com/PuerkitoBio/goquery"
)
type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}
var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
func main(){
	var jobs []extractedJob
	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++{
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
}
func getPage(page int) []extractedJobs {
	var jobs []extractedJobs
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractedJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}
func getPages() int{
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("").Length()
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
func extractJob(card *goquery.Selection){
	id, _ := cleanString(card.Attr("data-jk"))
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find("salaryText").Text())
	summary := claenString(card.Find("summary").Text())
}

func cleanString(str string) []string{
	return strings.Join(strings.Fields(strings.Trimspace(str)), " ")
}