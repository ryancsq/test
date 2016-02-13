package main

import (
	"fmt"
	"time"

	"new/util/myinit"
	"new/util/parseurl"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	now := time.Now()
	t := now.Format("2006-01-02 15:04:05")
	fmt.Println(t)
	runParseUrl()
	now2 := time.Now()
	t2 := now2.Format("2006-01-02 15:04:05")
	fmt.Println(t2)
}
func runParseUrl() {
	myinit.Myinit()

	for i:=30;i<31;i++ {
		a:=86400*int64(i)
		one_ago_unix := time.Now().Unix() - a
		t1 := time.Unix(one_ago_unix, 0)
		date := t1.Format("2006-01-02")
		fmt.Println(date)
		date = "2016-01-14"
		parseurl.ParseBetUrl(date, true)
		parseurl.ParseResultUrl(date, true)
	}	

}

