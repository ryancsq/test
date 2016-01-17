package main

import (
	"fmt"
	"time"

	"500kan/util/myinit"
	"500kan/util/parseurl"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			myinit.DeleteFile()
		}
	}()
	check := myinit.CheckFile()
	if check == false {
		fmt.Println("exist file")
		return
	}

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
	//	date := "2016-01-01"
	//	now := time.Now()
	//	date := now.Format("2006-01-02")
	date := ""
	parseurl.ParseBetUrl(date, false)
	parseurl.ParseResultUrl(date, false)

	one_ago_unix := time.Now().Unix() - 86400
	t1 := time.Unix(one_ago_unix, 0)
	parseurl.ParseResultUrl(t1.Format("2006-01-02"), false)

	moveToBackup()

	time.Sleep(30 * time.Second)
	runParseUrl()
}

func moveToBackup() {
	now := time.Now()
	today := now.Format("2006-01-02")
	del_sql := "delete from `pk_asia_pan` where schedule_date < ?"
	myinit.Engine.Exec(del_sql, today)
}
