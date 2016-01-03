package main

import (
	"fmt"
	"time"

	"500kan/util/myinit"
	"500kan/util/parseurl"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
//	runPanMap()
runParseUrl()
}
func runParseUrl(){
	
	myinit.Myinit()
		date := "2015-12-21"
	//	now := time.Now()
	//	date := now.Format("2006-01-02")
//	date := ""
	parseurl.ParseBetUrl(date, true)
	parseurl.ParseResultUrl(date, true)

	//	one_ago_unix := time.Now().Unix() - 86400
	//	t1 := time.Unix(one_ago_unix, 0)
	//	parseurl.ParseResultUrl(t1.Format("2006-01-02"),true)

//		moveToBackup()
}

func runPanMap() {
	date := ""
	for i := 91; i < 150; i++ {
		tmp := (int)(86400) * i
		ago_unix := time.Now().Unix() - int64(tmp)
		t1 := time.Unix(ago_unix, 0)
		date = t1.Format("2006-01-02")
		parseurl.ParsePanMap(date, true)
	}
}

func moveToBackup() {
	now := time.Now()
	today := now.Format("2006-01-02")
	fmt.Println(today)
	sql := "insert into `pk_asia_pan_backup` select * from `pk_asia_pan` where schedule_date < ?"
	ins_res, ins_err := myinit.Engine.Exec(sql, today)
	fmt.Println(ins_err, ins_res)
	del_sql := "delete from `pk_asia_pan` where schedule_date < ?"
	del_res, del_err := myinit.Engine.Exec(del_sql, today)
	fmt.Println(del_res, del_err)

	seven_ago_unix := time.Now().Unix() - 86400*7

	t1 := time.Unix(seven_ago_unix, 0)
	fmt.Println(t1.Format("2006-01-02"))

	del_backup_sql := "delete from `pk_asia_pan_backup` where schedule_date<?"
	del_backup_res, del_backup_err := myinit.Engine.Exec(del_backup_sql, t1.Format("2006-01-02"))
	fmt.Println(del_backup_res, del_backup_err)
}
