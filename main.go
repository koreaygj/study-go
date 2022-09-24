package main

import(
	"fmt"

	"github.com/koreaygj/study-go/accounts"
)
func main(){
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}