package myinit

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"	
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type PanMap struct {
	PanId    int
	PanDesc  string
	PanValue float32
}

type Schedule struct {
	ScheduleId         string
	ScheduleBetDate    string
	ScheduleDate       string
	ScheduleNo         string
	ScheduleResultNo   string
	ScheduleLeague     string
	ScheduleHome       string
	ScheduleGuest      string
	ScheduleWeekDay    string
	ScheduleFenxiId    int
	ScheduleBetEndTime string
	ScheduleRqNum      string

	ScheduleScore       string
	ScheduleSpfResult   string
	ScheduleSpfOdd      float32
	ScheduleRqspfResult string
	ScheduleRqspfOdd    float32
	ScheduleZjqResult   string
	ScheduleZjqOdd      float32
	ScheduleBqcResult   string
	ScheduleBqcOdd      float32
	ScheduleRate        float32
	ScheduleAlResult    string
}

type AsiaPan struct {
	AsiaId          string
	ScheduleFenxiId int
	ScheduleBetDate string

	ScheduleDate     string
	ScheduleNo       string
	ScheduleResultNo string
	ScheduleLeague   string
	ScheduleHome     string
	ScheduleGuest    string
	ScheduleGameDesc string
	ScheduleDateDesc string
	CompanyId        string
	CompanyName      string
	IsBigCompany     int

	OpenPan        float32
	OpenPanDesc    string
	OpenHomeWater  float32
	OpenGuestWater float32
	OpenPanTime    string

	RealPan        float32
	RealPanDesc    string
	RealHomeWater  float32
	RealGuestWater float32
	PanChangeTime  string

	HomePanChangeType     int
	HomePanChangeTypeDesc string

	HomeWaterChangeType     int
	HomeWaterChangeTypeDesc string
	Predict1Result          string
	Predict1Comment         string
	Predict2Result          string
	Predict2Comment         string

	ScheduleScore       string
	ScheduleSpfResult   string
	ScheduleRqspfResult string
	ScheduleZjqResult   string
	ScheduleBqcResult   string
}

type AsiaPanLog struct {
	AsiaId          string
	ScheduleFenxiId int
	ScheduleBetDate string

	ScheduleDate     string
	ScheduleNo       string
	ScheduleResultNo string
	ScheduleLeague   string
	ScheduleHome     string
	ScheduleGuest    string
	ScheduleGameDesc string
	ScheduleDateDesc string
	CompanyId        string
	CompanyName      string
	IsBigCompany     int

	OpenPan        float32
	OpenPanDesc    string
	OpenHomeWater  float32
	OpenGuestWater float32
	OpenPanTime    string

	RealPan        float32
	RealPanDesc    string
	RealHomeWater  float32
	RealGuestWater float32
	PanChangeTime  string

	HomePanChangeType     int
	HomePanChangeTypeDesc string

	HomeWaterChangeType     int
	HomeWaterChangeTypeDesc string
	Predict1Result          string
	Predict1Comment         string
	Predict2Result          string
	Predict2Comment         string

	ScheduleScore       string
	ScheduleSpfResult   string
	ScheduleRqspfResult string
	ScheduleZjqResult   string
	ScheduleBqcResult   string
}

var WeekDesc = map[string]string{
	"1": "周一",
	"2": "周二",
	"3": "周三",
	"4": "周四",
	"5": "周五",
	"6": "周六",
	"7": "周日",
}

var Engine *xorm.Engine

var DateUrl = "http://trade.500.com/jczq/?date=TTT&playtype=both"

var IndexUrl = "http://trade.500.com/jczq/"
var PanUrl = "http://odds.500.com/fenxi/yazhi-TTT.shtml"
var ResultUrl = "http://zx.500.com/jczq/kaijiang.php?d=DDD"
var PanChangeUrl = "http://odds.500.com/fenxi1/inc/yazhiajax.php?fid=FID&id=CID&r=1"

var mysql_dsn = "root:@tcp(localhost:3306)/new"

//var mysql_dsn = "root:123456@tcp(192.168.1.172:3306)/test_ha2"

//var mysql_dsn = "qichejingli:qichejingli1234QWER@tcp(rds3bhb1ed059c58i02wo.mysql.rds.aliyuncs.com:3306)/new_ha"

func Myinit() {
	initDb()
}

func initDb() {
	var err error
	Engine, err = xorm.NewEngine("mysql", mysql_dsn)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "pk_")
	Engine.SetTableMapper(tbMapper)

	Engine.ShowSQL = true   //则会在控制台打印出生成的SQL语句；
	Engine.ShowDebug = true //则会在控制台打印调试信息；
	Engine.ShowErr = true   //则会在控制台打印错误信息；
	Engine.ShowWarn = true  //则会在控制台打印警告信息；

	f, err := os.Create("sql.log")
	if err != nil {
		println("error:")
		println(err.Error())
		return
	}
	defer f.Close()
	Engine.Logger = xorm.NewSimpleLogger(f)
	fmt.Println(Engine)
}

func GetOddsFromAjax(schedule_fenxi_id int, company_id string)(body string) {
	odd_detail_url := strings.Replace(PanChangeUrl, "FID", strconv.Itoa(schedule_fenxi_id), -1)
	odd_detail_url = strings.Replace(odd_detail_url, "CID", company_id, -1)

	
	client := &http.Client{}
	request, _ := http.NewRequest("GET", odd_detail_url, nil)

	request.Header.Set("Accept", "application/json, text/javascript,*/*")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Connection", "Keep-Alive")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")

	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		fmt.Println(response.Header.Get("Content-Encoding"))
		reader, _ := gzip.NewReader(response.Body)
		for {
			buf := make([]byte, 1024)
			n, err := reader.Read(buf)

			if err != nil && err != io.EOF {
				panic(err)
			}

			if n == 0 {
				break
			}
			body += string(buf)
		}
//		fmt.Println(body)

	}
	return body
}

