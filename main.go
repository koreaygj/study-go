package main

import(
	"fmt"
	"github.com/koreaygj/study-go/mydict"
)
func main(){
	dictionary := mydict.Dictionary{}
	word := "hello"
	err := dictionary.Add(word, "hil")
	if err != nil {
		fmt.Println(err)
	}
	err2 := dictionary.Update(word, "king")
	if err2 != nil {
		fmt.Println(err2)
	}
	ans, err3 := dictionary.Search(word)
	if err3 != nil{
		fmt.Println(err3)
	}
	fmt.Println("Word: ", word, "definiton: ", ans)
}