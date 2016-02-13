package main

import (
	"fmt"
	"time"
"encoding/json"
	

	"500kan/util/common"
	"500kan/util/myinit"
	"500kan/util/analyse"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	
	recalResult()
	
}

func recalResult() {
	myinit.Myinit()
	//	date := "2016-01-01"
	//	now := time.Now()
	//	date := now.Format("2006-01-02")
	
//	for i:=0;i<30;i++ {
//		a:=86400*int64(i)
//		one_ago_unix := time.Now().Unix() - a
//		t1 := time.Unix(one_ago_unix, 0)
//		date := t1.Format("2006-01-02")
//		recal(date)
//		recalSchedule(date)
//	}
	date := "2016-02-01"
	
	recal(date)
	recalSchedule(date)

}

func recalSchedule(date string){
	if(date==""){
		now := time.Now()
		date = now.Format("2006-01-02")
	}
	schedule := new(myinit.Schedule)
	rows, err := myinit.Engine.Where("schedule_date =?", date).Rows(schedule)
if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
	    err = rows.Scan(schedule)
		fid := schedule.ScheduleFenxiId
		date := schedule.ScheduleDate
		schedule_no := schedule.ScheduleNo
		calcScheduleResult(fid,date,schedule_no)
	}
}

func recal(date string){
	if(date==""){
		now := time.Now()
		date = now.Format("2006-01-02")
	}
	
	asia_backup := new(myinit.AsiaPanBackup)
	rows, err := myinit.Engine.Where("schedule_date =?", date).Rows(asia_backup)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
	    err = rows.Scan(asia_backup)
		pan_int_info := make(map[string]int)
		pan_float_info := make(map[string]float32)
		pan_string_info := make(map[string]string)
		
//		if(asia_backup.ScheduleFenxiId!=557720){
//			continue
//		}
//		if(asia_backup.CompanyId!="6"){
//			continue
//		}
		
	pan_int_info["schedule_fenxi_id"] = asia_backup.ScheduleFenxiId
	pan_string_info["schedule_date"] = asia_backup.ScheduleDate
	pan_string_info["schedule_league"] = asia_backup.ScheduleLeague

	pan_string_info["schedule_home"] = asia_backup.ScheduleHome
	pan_string_info["company_id"] = asia_backup.CompanyId
	pan_string_info["company_name"] = asia_backup.CompanyName

pan_float_info["open_pan"] = asia_backup.OpenPan
	pan_string_info["open_pan_desc"] = asia_backup.OpenPanDesc
pan_float_info["open_home_water"] = asia_backup.OpenHomeWater
	pan_float_info["open_guest_water"] = asia_backup.OpenGuestWater
	pan_string_info["open_pan_time"] = asia_backup.OpenPanTime
	pan_float_info["real_pan"] = asia_backup.RealPan
	pan_string_info["real_pan_desc"] = asia_backup.RealPanDesc
	pan_float_info["real_home_water"] = asia_backup.RealHomeWater
	pan_float_info["real_guest_water"] = asia_backup.RealGuestWater

	pan_string_info["pan_change_time"] = asia_backup.PanChangeTime
	pan_int_info["home_pan_change_type"] = asia_backup.HomePanChangeType
	pan_string_info["home_pan_change_type_desc"] = asia_backup.HomePanChangeTypeDesc
	pan_int_info["home_water_change_type"] = asia_backup.HomeWaterChangeType
	pan_string_info["home_water_change_type_desc"] = asia_backup.HomeWaterChangeTypeDesc

fmt.Println(asia_backup.HomePanChangeType,pan_float_info["open_pan"],"==",pan_float_info["real_pan"],"==",pan_float_info["open_home_water"],pan_float_info["real_home_water"])

	predict1_result, predict1_cmt := analyse.AnalysePanResult1(pan_int_info, pan_float_info, pan_string_info)
	pan_string_info["predict1_result"] = predict1_result
	pan_string_info["predict1_cmt"] = predict1_cmt
	predict2_result, predict2_cmt := analyse.AnalysePanResult2(pan_int_info, pan_float_info, pan_string_info)
	pan_string_info["predict2_result"] = predict2_result
	pan_string_info["predict2_cmt"] = predict2_cmt
fmt.Println(asia_backup.ScheduleHome,"==",asia_backup.CompanyName,"==",predict1_result,predict1_cmt)
		UpdateAsiaPanBackupInfo(pan_int_info,pan_float_info,pan_string_info)
	}
}

func UpdateAsiaPanBackupInfo(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) () {
	schedule_fenxi_id := pan_int_info["schedule_fenxi_id"]
	company_id := pan_string_info["company_id"]
	update_lastpan := new(myinit.AsiaPanBackup)
	update_lastpan.Predict1Result = pan_string_info["predict1_result"]
	update_lastpan.Predict1Comment = pan_string_info["predict1_cmt"]
	update_lastpan.Predict2Result = pan_string_info["predict2_result"]
	update_lastpan.Predict2Comment = pan_string_info["predict2_cmt"]
	
	upda_a,err := 	myinit.Engine.
			Cols( "predict1_result", "predict1_comment", "predict2_result", "predict2_comment").
			Where("schedule_fenxi_id=? AND company_id=? ", schedule_fenxi_id, company_id).Update(update_lastpan)
	fmt.Println(upda_a,"---",err)
	
}
func getPredictResMap(predict_type int, date string,schedule_no string) (res_predict_map map[string]interface{}) {
	predict_map := make(map[string]interface{})
	predict_sql := ""
	if predict_type == 1 {
		predict_sql = "select predict1_result,count(*) as predict1_cnt from `pk_asia_pan_backup` where schedule_date = ? and schedule_no=? group by predict1_result"
	} else if predict_type == 2 {
		predict_sql = "select predict2_result,count(*) as predict2_cnt from `pk_asia_pan_backup` where schedule_date = ? and schedule_no=? group by predict2_result"
	}
	predict_res_map, _ := myinit.Engine.Query(predict_sql, date, schedule_no)

	for _, predict_res_row := range predict_res_map {
		predict_map_key := ""
		predict_map_val := ""
		for colName, colValue := range predict_res_row {
			value := common.BytesToString(colValue)
			if value == "" {
				value = "空"
			}
			
			if predict_type == 1 {
				if colName == "predict1_result" {
					predict_map_key = value
				}
				if colName == "predict1_cnt" {
					predict_map_val = value
				}
			} else if predict_type == 2 {
				if colName == "predict2_result" {
					predict_map_key = value
				}
				if colName == "predict2_cnt" {
					predict_map_val = value
				}
			}
		}
		predict_map[predict_map_key] = predict_map_val

	}
	res_predict_map = predict_map
	return res_predict_map
}

func calcScheduleResult(fid int, date string, schedule_no string) {
	predict1_map := getPredictResMap(1, date, schedule_no)
	predict2_map := getPredictResMap(2, date, schedule_no)

	predict_map := make(map[string]interface{})
	predict_map["predict1"] = predict1_map
	predict_map["predict2"] = predict2_map

	predict_json, _ := json.Marshal(predict_map)

	exist_schedule := new(myinit.Schedule)
	exist_schedule.ScheduleAlResult = string(predict_json)

	myinit.Engine.Cols("schedule_al_result").
		Where("schedule_fenxi_id=? AND schedule_date = ? and schedule_no=? ", fid, date, schedule_no).Update(exist_schedule)

}


