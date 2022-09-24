package main

import(
	"fmt"
	"github.com/koreaygj/study-go/accounts"
)

func main(){
	account := accounts.NewAccount("yang")
	fmt.Println(account)
}