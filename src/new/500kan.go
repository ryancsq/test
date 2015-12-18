package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"500kan/util/myinit"	
	"500kan/util/analyse"	
	"500kan/util/schedule"	

	_ "github.com/go-sql-driver/mysql"
	"github.com/guotie/gogb2312"
	"github.com/opesun/goquery"
)

func convToGB(str string) (res_str string) {
	conv_str, _, _, _ := gogb2312.ConvertGB2312String(str)
	return conv_str
}

func ParseOddUrl(schedule_odds_url string, schedulefid string, pname string, schedule_date_orig string) (a string) {
	odds_html, _ := goquery.ParseUrl(schedule_odds_url)

	schedule_item := odds_html.Find(".odds_hd_cont table tbody tr td")
	home_td := schedule_item.Eq(0)
	guest_td := schedule_item.Eq(4)
	center_td := schedule_item.Eq(2)

		pan_int_info:= make(map[string]int)
		pan_string_info := make(map[string]string)
	
schedule_home_name := pan_string_info["schedule_home_name"] =  convToGB(home_td.Find("ul li a").Text())
	schedule_guest_name := pan_string_info["schedule_guest_name"] = convToGB(guest_td.Find("ul li a").Text())
pan_string_info["schedule_game_desc"] =	schedule_game_desc := convToGB(center_td.Find(".odds_hd_center .odds_hd_ls a").Text())
pan_string_info["schedule_date_desc"] =	convToGB(center_td.Find(".odds_hd_center .game_time ").Text())
	//	schedule_date := convToGB(center_td.Find(".odds_hd_center .game_time ").Text())

pan_string_info["schedule_date"] =	schedule_date := schedule_date_orig

	odds_tr := odds_html.Find(".table_cont table tbody tr")
	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)

		pan_string_info["company"] = company := convToGB(td_of_company.Find("p a").Attr("title"))
		if td_of_company.Find("p a").Attr("title") == "" {
			continue
		}

		var is_big_company = "0"
		if td_of_company.Find("p img").Attr("src") == "" {
			is_big_company = "0"
		} else {
			is_big_company = "1"
			fmt.Println("src:" + td_of_company.Find("p img").Attr("src"))
		}
pan_string_info["is_big_company"] = is_big_company

		cid := tr_item.Attr("id")
		
		pan_string_info["cid"] = cid


		td_of_pan_time := tr_item.Find("td time")

		change_time := td_of_pan_time.Eq(0).Text()
		open_time := td_of_pan_time.Eq(1).Text()

		table_of_pan_detail := tr_item.Find("td .pl_table_data")
		table_of_realtime_pan := table_of_pan_detail.Eq(0)
		tds_of_realtime_pan_table := table_of_realtime_pan.Find("tbody tr td")

		home_water_up_down_flag := tds_of_realtime_pan_table.Eq(0).Attr("class")
		var home_water_change_type = "water_unknown"
		if home_water_up_down_flag == "ping" {
			home_water_change_type = "water_down" // down
		}
		if home_water_up_down_flag == "ying" {
			home_water_change_type = "water_up" // up
		}

		home_real_water_string := convToGB(tds_of_realtime_pan_table.Eq(0).Text())
		guest_real_water_string := convToGB(tds_of_realtime_pan_table.Eq(2).Text())

		real_pan_32, _ := strconv.ParseFloat(tds_of_realtime_pan_table.Eq(1).Attr("ref"), 32)
		real_pan := float32(real_pan_32)

		real_pan_desc := convToGB(tds_of_realtime_pan_table.Eq(1).Text())

		td_item_of_real_pan := tds_of_realtime_pan_table.Eq(1)
		home_pan_change_type := convToGB(td_item_of_real_pan.Find("font").Text())
		home_pan_change_type = strings.TrimSpace(home_pan_change_type)
		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		table_of_opentime_pan := table_of_pan_detail.Eq(1)
		tds_of_opentime_pan_table := table_of_opentime_pan.Find("tbody tr td")

		open_home_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(0).Text(), 32)
		open_guest_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(2).Text(), 32)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		open_home_water := float32(open_home_water_32)
		open_guest_water := float32(open_guest_water_32)

		home_real_water := float32(home_real_water_32)
		guest_real_water := float32(guest_real_water_32)

		open_pan_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(1).Attr("ref"), 32)
		open_pan := float32(open_pan_32)

		if open_pan > 0 || real_pan > 0 {
			delete_lastpan := new(myinit.LastPan)
			del_result, _ := myinit.Engine.Where("schedule_fid=? ", schedulefid).Delete(delete_lastpan)
			delete_schedule := new(myinit.Schedule)
			del_schedule_result, _ := myinit.Engine.Where("schedule_fid=? ", schedulefid).Delete(delete_schedule)
			fmt.Println(del_result, del_schedule_result)
			return "开盘>0 或者即时盘 >0"
		}
		open_pan_desc := convToGB(tds_of_opentime_pan_table.Eq(1).Text())

		//		addLastPanLog()

		predict_result, predict_cmt := analyse.AnalysePanResult(open_pan, open_home_water, open_guest_water, real_pan, home_real_water, guest_real_water, home_pan_change_type, schedule_game_desc, schedulefid, cid)

		fmt.Println("float_open_pan")
		fmt.Println(open_home_water)
		fmt.Println("=====")

		fmt.Println("company:" + company)
		fmt.Println("home_pan_change_type:" + home_pan_change_type)
		fmt.Println("is big company:" + is_big_company)
		fmt.Println("change_time:" + change_time)
		fmt.Println("open_time:" + open_time)
		fmt.Println("flag:" + home_water_change_type + " " + home_water_up_down_flag)
		fmt.Println("home_real_water:", home_real_water)
		fmt.Println("home_real_water water sting:" + home_real_water_string)
		fmt.Println("guest_real_water:", guest_real_water)
		fmt.Println("guest_real_water water sting:" + guest_real_water_string)
		fmt.Println("pan:", real_pan, " ", real_pan_desc)

		fmt.Println("open_home_water water:", open_home_water)
		fmt.Println("open_guest_water water:", open_guest_water)
		fmt.Println("open pan:", open_pan, " ", open_pan_desc)

		exist_lastpan := new(myinit.LastPan)
		has, _ := myinit.Engine.Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Get(exist_lastpan)
		fmt.Println(has)
		if has {
			fmt.Println(company + "pan已存在！")
			if exist_lastpan.LastPan != real_pan || exist_lastpan.LastHomeWater != home_real_water || exist_lastpan.LastChangeTime != change_time {
				fmt.Println(company + "pan有变化！")
				update_lastpan := new(myinit.LastPan)
				update_lastpan.LastPan = real_pan
				update_lastpan.LastPanDesc = real_pan_desc
				update_lastpan.LastHomeWater = home_real_water
				update_lastpan.LastGuestWater = guest_real_water
				update_lastpan.LastHomeWaterChangeType = home_water_change_type
				update_lastpan.LastHomePanChangeType = home_pan_change_type
				update_lastpan.LastChangeTime = change_time
				update_lastpan.PredictResult = predict_result
				update_lastpan.PredictComment = predict_cmt
				update_affected, update_err := myinit.Engine.Cols("last_pan", "last_pan_desc", "last_home_water", "last_guest_water", "last_change_time", "last_home_pan_change_type", "last_home_water_change_type", "predict_result", "predict_comment").Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Update(update_lastpan)
				fmt.Println(update_affected)
				fmt.Println(update_err)

				LastPanLog := new(myinit.LastPanLog)
				LastPanLog.ScheduleNo = pname
				LastPanLog.ScheduleHome = schedule_home_name
				LastPanLog.ScheduleGuest = schedule_guest_name
				LastPanLog.ScheduleDate = schedule_date
				LastPanLog.ScheduleGameDesc = schedule_game_desc
				LastPanLog.CompanyCid = cid
				LastPanLog.CompanyName = company
				LastPanLog.ScheduleFid = schedulefid
				LastPanLog.OpenPan = open_pan
				LastPanLog.OpenPanDesc = open_pan_desc
				LastPanLog.OpenHomeWater = open_home_water
				LastPanLog.OpenGuestWater = open_guest_water
				LastPanLog.OpenPanTime = open_time
				LastPanLog.LastPan = real_pan
				LastPanLog.LastPanDesc = real_pan_desc
				LastPanLog.LastHomeWater = home_real_water
				LastPanLog.LastGuestWater = guest_real_water
				LastPanLog.LastChangeTime = change_time
				LastPanLog.LastHomePanChangeType = home_pan_change_type
				LastPanLog.IsBigCompany = is_big_company
				LastPanLog.LastHomeWaterChangeType = home_water_change_type
				LastPanLog.PredictResult = predict_result
				LastPanLog.PredictComment = predict_cmt

				log_ins_affected, log_ins_err := myinit.Engine.Insert(LastPanLog)
				fmt.Println(log_ins_affected)
				fmt.Println(log_ins_err)
			} else {

			}
		} else {
			LastPan := new(myinit.LastPan)
			LastPan.ScheduleNo = pname

			LastPan.ScheduleHome = schedule_home_name
			LastPan.ScheduleGuest = schedule_guest_name
			LastPan.ScheduleDate = schedule_date
			LastPan.ScheduleGameDesc = schedule_game_desc
			LastPan.CompanyCid = cid
			LastPan.CompanyName = company
			LastPan.ScheduleFid = schedulefid
			LastPan.OpenPan = open_pan
			LastPan.OpenPanDesc = open_pan_desc
			LastPan.OpenHomeWater = open_home_water
			LastPan.OpenGuestWater = open_guest_water
			LastPan.OpenPanTime = open_time
			LastPan.LastPan = real_pan
			LastPan.LastPanDesc = real_pan_desc
			LastPan.LastHomeWater = home_real_water
			LastPan.LastGuestWater = guest_real_water
			LastPan.LastChangeTime = change_time
			LastPan.LastHomePanChangeType = home_pan_change_type
			LastPan.IsBigCompany = is_big_company
			LastPan.LastHomeWaterChangeType = home_water_change_type
			LastPan.PredictResult = predict_result
			LastPan.PredictComment = predict_cmt

			ins_affected, ins_err := myinit.Engine.Insert(LastPan)
			fmt.Println(ins_affected)
			fmt.Println(ins_err)

			LastPanLog := new(myinit.LastPanLog)
			LastPanLog.ScheduleNo = pname

			LastPanLog.ScheduleHome = schedule_home_name
			LastPanLog.ScheduleGuest = schedule_guest_name
			LastPanLog.ScheduleDate = schedule_date
			LastPanLog.ScheduleGameDesc = schedule_game_desc
			LastPanLog.CompanyCid = cid
			LastPanLog.CompanyName = company
			LastPanLog.ScheduleFid = schedulefid
			LastPanLog.OpenPan = open_pan
			LastPanLog.OpenPanDesc = open_pan_desc
			LastPanLog.OpenHomeWater = open_home_water
			LastPanLog.OpenGuestWater = open_guest_water
			LastPanLog.OpenPanTime = open_time
			LastPanLog.LastPan = real_pan
			LastPanLog.LastPanDesc = real_pan_desc
			LastPanLog.LastHomeWater = home_real_water
			LastPanLog.LastGuestWater = guest_real_water
			LastPanLog.LastChangeTime = change_time
			LastPanLog.LastHomePanChangeType = home_pan_change_type
			LastPanLog.IsBigCompany = is_big_company
			LastPanLog.LastHomeWaterChangeType = home_water_change_type
			LastPanLog.PredictResult = predict_result
			LastPanLog.PredictComment = predict_cmt

			log_ins_affected, log_ins_err := myinit.Engine.Insert(LastPanLog)
			fmt.Println(log_ins_affected)
			fmt.Println(log_ins_err)
		}

		//		count_open_water := open_home_water +open_guest_water
		count_real_water := home_real_water + guest_real_water
		if count_real_water < 1.75 || count_real_water > 2 {
			delete_lastpan2 := new(myinit.LastPan)
			delete2, _ := myinit.Engine.Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Delete(delete_lastpan2)
			fmt.Println(delete2)
		}
	}
	return "成功"
}


func compareDate(time1 string, time2 string) (ret bool) {
	//先把时间字符串格式化成相同的时间类型
	t1, _ := time.Parse("2006-01-02 15:04:05", time1)
	t2, _ := time.Parse("2006-01-02 15:04:05", time2)
	if t1.Before(t2) {
		return true
	}
	return false
}

var index_url = "http://trade.500.com/jczq/"
var odds_url = "http://odds.500.com/fenxi/yazhi-TTT.shtml"
var log_url = "http://odds.500.com/fenxi1/inc/yazhiajax.php?fid=554298&id=526&t=1449408646779&r=1"

func main() {
	myinit.Myinit()
	

	x, _ := goquery.ParseUrl(index_url)
	schedule_trs := x.Find(".bet_table tbody tr")
	for i, _ := range schedule_trs {
		is_end := schedule_trs.Eq(i).Attr("isend")
		if is_end == "1" {
			fmt.Println("is_end")
			continue
		}

		schedule_int_info:= make(map[string]int)
		schedule_string_info := make(map[string]string)

		// insert schedule
		fid,_ := strconv.ParseInt(schedule_trs.Eq(i).Attr("fid"),10,0)
		schedule_int_info["fid"] = int(fid)
		schedule_string_info["home_team"] = convToGB(schedule_trs.Eq(i).Attr("homesxname"))
		schedule_string_info["guest_team"] = convToGB(schedule_trs.Eq(i).Attr("awaysxname"))
		schedule_string_info["schedule_date"]  = schedule_trs.Eq(i).Attr("pdate")
		schedule_string_info["lg"] = convToGB(schedule_trs.Eq(i).Attr("lg"))
		schedule_string_info["game_date_no"]  = schedule_trs.Eq(i).Attr("gdate")
		schedule_string_info["schedule_no"]  = schedule_trs.Eq(i).Attr("pname")
		schedule_string_info["end_time"]  = schedule_trs.Eq(i).Attr("pendtime")

		now := time.Now()
		year, mon, day := now.Date()
		schedule_is_today := compareDate(schedule_string_info["end_time"], time.Date(year, mon, day, 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05"))
		fmt.Println("schedule_is_today:===")
		fmt.Println(schedule_is_today)
		fmt.Println("end schedule_is_today===")
		if schedule_is_today == false {
			continue
		}
		schedule.Add(schedule_int_info,schedule_string_info)
		// end insert schedule

		schedule_odds_url := strings.Replace(odds_url, "TTT", strconv.FormatInt(fid,10), -1)
		fmt.Println(schedule_odds_url)
		//		go parseOddUrl(schedule_odds_url, fid)
		res := ParseOddUrl(schedule_odds_url, strconv.FormatInt(fid,10), schedule_string_info["schedule_no"], schedule_string_info["schedule_date"])
		fmt.Println(res)
	}
//	moveToBackup()
}

func moveToBackup() {
	now := time.Now()
	today := now.Format("2006-01-02")
	fmt.Println(today)
	sql := "insert into `pk_last_pan_backup` select * from `pk_last_pan` where schedule_date < ?"
	ins_res, ins_err := myinit.Engine.Exec(sql, today)
	fmt.Println(ins_err, ins_res)
	del_sql := "delete from `pk_last_pan` where schedule_date < ?"
	del_res, del_err := myinit.Engine.Exec(del_sql, today)
	fmt.Println(del_res, del_err)

	seven_ago_unix := time.Now().Unix() - 86400*7

	t1 := time.Unix(seven_ago_unix, 0)
	fmt.Println(t1.Format("2006-01-02"))

	del_backup_sql := "delete from `pk_last_pan_backup` where schedule_date<?"
	del_backup_res, del_backup_err := myinit.Engine.Exec(del_backup_sql, t1.Format("2006-01-02"))
	fmt.Println(del_backup_res, del_backup_err)
}
