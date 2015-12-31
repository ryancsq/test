package parseurl

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"500kan/util/common"
	"500kan/util/myinit"
	"500kan/util/schedule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func ParseBetUrl(date string, history bool) {
	bet_url := strings.Replace(myinit.DateUrl, "TTT", date, -1)
	if date == "" {
		bet_url = myinit.IndexUrl
	}
	fmt.Println("bet_url:",bet_url)

	pan_url := myinit.PanUrl

	html_obj, _ := goquery.ParseUrl(bet_url)
	schedule_trs := html_obj.Find(".bet_table tbody tr")
	for i, _ := range schedule_trs {
		is_end := schedule_trs.Eq(i).Attr("isend")
		if is_end == "1" && history == false {
			fmt.Println("is_end")
			continue
		}

		schedule_int_info := make(map[string]int)
		schedule_string_info := make(map[string]string)

		// insert schedule
		fid, _ := strconv.Atoi(schedule_trs.Eq(i).Attr("fid"))
		schedule_int_info["schedule_fenxi_id"] = int(fid)
		schedule_string_info["schedule_home"] = common.ConvToGB(schedule_trs.Eq(i).Attr("homesxname"))
		schedule_string_info["schedule_guest"] = common.ConvToGB(schedule_trs.Eq(i).Attr("awaysxname"))
		schedule_string_info["schedule_date"] = schedule_trs.Eq(i).Attr("pdate")
		schedule_string_info["schedule_league"] = common.ConvToGB(schedule_trs.Eq(i).Attr("lg"))
		schedule_string_info["schedule_week_day"] = common.ConvToGB(schedule_trs.Eq(i).Attr("gdate"))
		schedule_string_info["schedule_no"] = schedule_trs.Eq(i).Attr("pname")
		schedule_string_info["schedule_rq_num"] = schedule_trs.Eq(i).Attr("rq")
		week_date := schedule_string_info["schedule_no"][0:1]
		schedule_string_info["schedule_result_no"] = myinit.WeekDesc[week_date] + schedule_string_info["schedule_no"][1:]

		schedule_string_info["schedule_bet_end_time"] = schedule_trs.Eq(i).Attr("pendtime")

		today := time.Now().Format("2006-01-02")
		schedule_is_today := today == schedule_string_info["schedule_date"]
		fmt.Println("schedule_is_today:===",schedule_is_today)
		if schedule_is_today == false && history == false {
			continue
		}
		schedule.Add(schedule_int_info, schedule_string_info)
		// end insert schedule

		schedule_pan_url := strings.Replace(pan_url, "TTT", strconv.Itoa(fid), -1)
		fmt.Println("schedule_pan_url:",schedule_pan_url)
		//		go parseOddUrl(schedule_odds_url, fid)
		res := ParsePanUrl(schedule_pan_url, fid, schedule_string_info, date)
		fmt.Println(res)
		fmt.Println("---------")
		fmt.Println(schedule_string_info["schedule_date"])
		fmt.Println(schedule_string_info["schedule_no"])
		fmt.Println("---------")
		if res == false {
			continue
		}
		calcScheduleResult(schedule_int_info,schedule_string_info)

	}
}


func calcScheduleResult(schedule_int_info map[string]int,schedule_string_info map[string]string){
		predict1_json := make(map[string]interface{})
		predict2_json := make(map[string]interface{})

		sql1 := "select predict1_result,count(*) as predict1_cnt from `pk_asia_pan_log` where schedule_date = ? and schedule_no=? group by predict1_result"
		res_map, _ := myinit.Engine.Query(sql1, schedule_string_info["schedule_date"], schedule_string_info["schedule_no"])

		for _, row := range res_map {
			json_key := ""
			json_val := ""
			for colName, colValue := range row {

				value := common.BytesToString(colValue)
				fmt.Println("colName")
				fmt.Println(colName)
				fmt.Println("value")
				fmt.Println(value)
				if colName == "predict1_result" {
					json_key = value
				}
				if colName == "predict1_cnt" {
					json_val = value
				}
			}
			predict1_json[json_key] = json_val

		}

		sql2 := "select predict2_result,count(*) as predict2_cnt from `pk_asia_pan_log` where schedule_date = ? and schedule_no=? group by predict2_result"
		res2_map, _ := myinit.Engine.Query(sql2, schedule_string_info["schedule_date"], schedule_string_info["schedule_no"])

		for _, row2 := range res2_map {
			json_key2 := ""
			json_val2 := ""
			for colName2, colValue2 := range row2 {

				value2 := common.BytesToString(colValue2)
				if colName2 == "predict2_result" {
					json_key2 = value2
				}
				if colName2 == "predict2_cnt" {
					json_val2 = value2
				}
			}
			predict2_json[json_key2] = json_val2

		}

//		predict1_string, _ := json.Marshal(predict1_json)
//		fmt.Println(string(predict1_string))

//		predict2_string, _ := json.Marshal(predict2_json)
//		fmt.Println(string(predict2_string))

		predict_json := make(map[string]interface{})
		predict_json["predict1"] = predict1_json
		predict_json["predict2"] = predict2_json

		predict_json_string, _ := json.Marshal(predict_json)

		exist_schedule := new(myinit.Schedule)
		exist_schedule.ScheduleAlResult = string(predict_json_string)
		update_affected, update_err :=
			myinit.Engine.Cols("schedule_al_result").
				Where("schedule_fenxi_id=? AND schedule_date = ? and schedule_no=? ", schedule_int_info["schedule_fenxi_id"], schedule_string_info["schedule_date"], schedule_string_info["schedule_no"]).Update(exist_schedule)

		fmt.Println(update_affected)
		fmt.Println(update_err)
}
