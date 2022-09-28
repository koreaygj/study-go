package main
import(
	"log"
	"net/http"
	"fmt"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)
var baseURL string = "https://koreaygj.github.io"
func main(){
	totalPages := getPages()
	for i := 0; i < totalPages; i++{
		getPage(i)
	}
}
func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
}
func getPages() int{
	pages := 0
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	_, err := client.Get(baseURL)
	req, err := http.NewRequest("GET", baseURL, nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	resp, err := client.Do(req)
	checkErr(err)
	checkCode(resp)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})
	return  pages
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