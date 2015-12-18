package myinit

import (
	"fmt"
	"os"
	

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"	
)

type Schedule struct {
	ScheduleId       string
	ScheduleHome     string
	ScheduleGuest    string
	ScheduleDate     string
	ScheduleLeague   string
	ScheduleWeekDay  string
	ScheduleNo       string
	ScheduleFid      int
	ScheduleEndTime  string
	ScheduleRate     float64
	ScheduleResult   string
	ScheduleAlResult string
}

type LastPan struct {
	LastId     string
	ScheduleNo string

	ScheduleHome            string
	ScheduleGuest           string
	ScheduleDate            string
	ScheduleGameDesc        string
	CompanyCid              string
	CompanyName             string
	ScheduleFid             string
	OpenPan                 float32
	OpenPanDesc             string
	OpenHomeWater           float32
	OpenGuestWater          float32
	OpenPanTime             string
	LastHomePanChangeType   string
	LastPan                 float32
	LastPanDesc             string
	LastHomeWater           float32
	LastGuestWater          float32
	LastChangeTime          string
	IsBigCompany            string
	LastHomeWaterChangeType string
	PredictResult           string
	PredictComment          string
}


type LastPanLog struct {
	LastId                  string
	ScheduleNo              string
	ScheduleHome            string
	ScheduleGuest           string
	ScheduleDate            string
	ScheduleGameDesc        string
	CompanyCid              string
	CompanyName             string
	ScheduleFid             string
	OpenPan                 float32
	OpenPanDesc             string
	OpenHomeWater           float32
	OpenGuestWater          float32
	OpenPanTime             string
	LastHomePanChangeType   string
	LastPan                 float32
	LastPanDesc             string
	LastHomeWater           float32
	LastGuestWater          float32
	LastChangeTime          string
	IsBigCompany            string
	LastHomeWaterChangeType string
	PredictResult           string
	PredictComment          string
}

var Engine *xorm.Engine

func Myinit(){
	initDb()
}

func initDb() {
	var err error
//		engine, err = xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/test_ha")
		Engine, err = xorm.NewEngine("mysql", "root:123456@tcp(192.168.1.172:3306)/test_ha2")
//	engine, err = xorm.NewEngine("mysql", "qichejingli:qichejingli1234QWER@tcp(rds3bhb1ed059c58i02wo.mysql.rds.aliyuncs.com:3306)/test_ha2")

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
		println("sql.log error:")
		println(err.Error())
		return
	}
	defer f.Close()
	Engine.Logger = xorm.NewSimpleLogger(f)
	fmt.Println(Engine)
}
