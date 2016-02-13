package parseurl

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"500kan/util/common"
	"500kan/util/myinit"
	"500kan/util/schedule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func getBetUrl(date string) (bet_url string) {
	bet_url = strings.Replace(myinit.DateUrl, "TTT", date, -1)
	if date == "" {
		bet_url = myinit.IndexUrl
	}
	fmt.Println("bet_url:", bet_url)

	return bet_url
}

func pareseScheduleTR(schedule_tr goquery.Nodes,date string) (schedule_int_info map[string]int, schedule_string_info map[string]string) {
	schedule_int_info = make(map[string]int)
	schedule_string_info = make(map[string]string)
	schedule_fenxi_id, _ := strconv.Atoi(schedule_tr.Attr("fid"))
	schedule_int_info["schedule_fenxi_id"] = int(schedule_fenxi_id)
	schedule_string_info["schedule_home"] = common.ConvToGB(schedule_tr.Attr("homesxname"))
	schedule_string_info["schedule_guest"] = common.ConvToGB(schedule_tr.Attr("awaysxname"))
	schedule_string_info["schedule_date"] = date
	schedule_string_info["schedule_league"] = common.ConvToGB(schedule_tr.Attr("lg"))
	schedule_string_info["schedule_week_day"] = common.ConvToGB(schedule_tr.Attr("gdate"))
	schedule_string_info["schedule_no"] = schedule_tr.Attr("pname")
	schedule_string_info["schedule_rq_num"] = schedule_tr.Attr("rq")
	week_date := schedule_string_info["schedule_no"][0:1]
	schedule_string_info["schedule_result_no"] = myinit.WeekDesc[week_date] + schedule_string_info["schedule_no"][1:]
	schedule_string_info["schedule_bet_end_time"] = schedule_tr.Attr("pendtime")
	return schedule_int_info, schedule_string_info
}

func ParseBetUrl(date string, history bool) {
	bet_url := getBetUrl(date)
	html_obj, _ := goquery.ParseUrl(bet_url)
	schedule_trs := html_obj.Find(".bet_table tbody tr")
	for i, _ := range schedule_trs {
		is_end := schedule_trs.Eq(i).Attr("isend")
		if is_end == "1" && history == false {
			continue
		}

		schedule_int_info, schedule_string_info := pareseScheduleTR(schedule_trs.Eq(i),date)
if(schedule_int_info["schedule_fenxi_id"]!=556793){
	continue
}
		//parse pan data
		res := ParsePanByScheduleFenxiId(schedule_int_info["schedule_fenxi_id"], date, schedule_string_info)
		if res == false {
			continue
		}
		schedule.Add(schedule_int_info, schedule_string_info)
		//计算预测比率
		calcScheduleResult(schedule_int_info, schedule_string_info)
//				return
		
	}
}

func ParsePanByScheduleFenxiId(schedule_fenxi_id int, date string, schedule_string_info map[string]string) (res bool) {
	pan_url := myinit.PanUrl
	schedule_pan_url := strings.Replace(pan_url, "TTT", strconv.Itoa(schedule_fenxi_id), -1)
	res = ParsePanUrl(schedule_pan_url, schedule_fenxi_id, schedule_string_info, date)

	return res
}

func getPredictResMap(predict_type int, schedule_string_info map[string]string) (res_predict_map map[string]interface{}) {
	predict_map := make(map[string]interface{})
	predict_sql := ""
	if predict_type == 1 {
		predict_sql = "select predict1_result,count(*) as predict1_cnt from `pk_asia_pan_backup` where schedule_date = ? and schedule_no=? group by predict1_result"
	} else if predict_type == 2 {
		predict_sql = "select predict2_result,count(*) as predict2_cnt from `pk_asia_pan_backup` where schedule_date = ? and schedule_no=? group by predict2_result"
	}
	predict_res_map, _ := myinit.Engine.Query(predict_sql, schedule_string_info["schedule_date"], schedule_string_info["schedule_no"])

	for _, predict_res_row := range predict_res_map {
		predict_map_key := ""
		predict_map_val := ""
		for colName, colValue := range predict_res_row {
			value := common.BytesToString(colValue)
			if value == "" {
				value = "空"
			}
			//			fmt.Println("colName")
			//			fmt.Println(colName)
			//			fmt.Println("value")
			//			fmt.Println(value)
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

func calcScheduleResult(schedule_int_info map[string]int, schedule_string_info map[string]string) {
	predict1_map := getPredictResMap(1, schedule_string_info)
	predict2_map := getPredictResMap(2, schedule_string_info)

	predict_map := make(map[string]interface{})
	predict_map["predict1"] = predict1_map
	predict_map["predict2"] = predict2_map

	predict_json, _ := json.Marshal(predict_map)
	schedule.UpdateScheduleCalcResult(predict_json, schedule_int_info, schedule_string_info)
}
