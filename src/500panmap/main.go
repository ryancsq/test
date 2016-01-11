package main

import (
	"time"

	"500panmap/util/parseurl"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	runPanMap()
}

func runPanMap() {
	date := ""
	for i := 30; i < 366; i++ {
		tmp := (int)(86400) * i
		ago_unix := time.Now().Unix() - int64(tmp)
		t1 := time.Unix(ago_unix, 0)
		date = t1.Format("2006-01-02")
		parseurl.ParsePanMap(date, true)
	}
}
