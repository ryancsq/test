package main
import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"500kan/util/analyse"
	"500kan/util/common"
	"500kan/util/lastpan"
	"500kan/util/lastpanlog"
	"500kan/util/myinit"
	"500kan/util/schedule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func ParseOddUrl(schedule_odds_url string, schedulefid string, pname string, schedule_date_orig string) (a string) {
	pan_int_info := make(map[string]int)
	pan_float_info := make(map[string]float32)
	pan_string_info := make(map[string]string)

	odds_html, _ := goquery.ParseUrl(schedule_odds_url)

	schedule_item := odds_html.Find(".odds_hd_cont table tbody tr td")
	home_td := schedule_item.Eq(0)
	guest_td := schedule_item.Eq(4)
	center_td := schedule_item.Eq(2)

	pan_int_info["schedule_fenxi_id"] = schedulefid

	pan_string_info["schedule_no"] = pname

	pan_string_info["schedule_home_name"] = common.ConvToGB(home_td.Find("ul li a").Text())
	pan_string_info["schedule_guest_name"] = common.ConvToGB(guest_td.Find("ul li a").Text())
	pan_string_info["schedule_game_desc"] = common.ConvToGB(center_td.Find(".odds_hd_center .odds_hd_ls a").Text())
	pan_string_info["schedule_date_desc"] = common.ConvToGB(center_td.Find(".odds_hd_center .game_time ").Text())
	pan_string_info["schedule_date"] = schedule_date_orig

	odds_tr := odds_html.Find(".table_cont table tbody tr")
	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)
		if td_of_company.Find("p a").Attr("title") == "" {
			continue
		}
		pan_string_info["company"] = common.ConvToGB(td_of_company.Find("p a").Attr("title"))

		var is_big_company = 0
		if td_of_company.Find("p img").Attr("src") == "" {
			is_big_company = 0
		} else {
			is_big_company = 1
			fmt.Println("src:" + td_of_company.Find("p img").Attr("src"))
		}
		pan_int_info["is_big_company"] = is_big_company

		cid := tr_item.Attr("id")
		pan_string_info["cid"] = cid

		td_of_pan_time := tr_item.Find("td time")

		pan_string_info["change_time"] = td_of_pan_time.Eq(0).Text()
		pan_string_info["open_time"] = td_of_pan_time.Eq(1).Text()

		table_of_pan_detail := tr_item.Find("td .pl_table_data")
		table_of_realtime_pan := table_of_pan_detail.Eq(0)
		tds_of_realtime_pan_table := table_of_realtime_pan.Find("tbody tr td")

		home_water_up_down_flag := tds_of_realtime_pan_table.Eq(0).Attr("class")
		pan_int_info["home_water_change_type"] = 0
		if home_water_up_down_flag == "ping" {
			pan_int_info["home_water_change_type"] = -1 // down
		}
		if home_water_up_down_flag == "ying" {
			pan_int_info["home_water_change_type"] = 1 // up
		}

		home_real_water_string := common.ConvToGB(tds_of_realtime_pan_table.Eq(0).Text())
		guest_real_water_string := common.ConvToGB(tds_of_realtime_pan_table.Eq(2).Text())

		real_pan_32, _ := strconv.ParseFloat(tds_of_realtime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["real_pan"] = float32(real_pan_32)

		pan_string_info["real_pan_desc"] = common.ConvToGB(tds_of_realtime_pan_table.Eq(1).Text())

		td_item_of_real_pan := tds_of_realtime_pan_table.Eq(1)
		home_pan_change_type := common.ConvToGB(td_item_of_real_pan.Find("font").Text())
		home_pan_change_type = strings.TrimSpace(home_pan_change_type)
		pan_string_info["home_pan_change_type"] = home_pan_change_type

		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		table_of_opentime_pan := table_of_pan_detail.Eq(1)
		tds_of_opentime_pan_table := table_of_opentime_pan.Find("tbody tr td")

		open_pan_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["open_pan"] = float32(open_pan_32)

		open_home_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(0).Text(), 32)
		open_guest_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(2).Text(), 32)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		pan_float_info["open_home_water"] = float32(open_home_water_32)
		pan_float_info["open_guest_water"] = float32(open_guest_water_32)

		pan_float_info["home_real_water"] = float32(home_real_water_32)
		pan_float_info["guest_real_water"] = float32(guest_real_water_32)

		if pan_float_info["open_pan"] > 0 || pan_float_info["real_pan"] > 0 {
			delete_lastpan := new(myinit.LastPan)
			del_result, _ := myinit.Engine.Where("schedule_fid=? ", schedulefid).Delete(delete_lastpan)
			delete_schedule := new(myinit.Schedule)
			del_schedule_result, _ := myinit.Engine.Where("schedule_fid=? ", schedulefid).Delete(delete_schedule)
			fmt.Println(del_result, del_schedule_result)
			return "开盘>0 或者即时盘 >0"
		}
		pan_string_info["open_pan_desc"] = common.ConvToGB(tds_of_opentime_pan_table.Eq(1).Text())

		//		addLastPanLog()

		// predict_result, predict_cmt := analyse.AnalysePanResult(open_pan, open_home_water, open_guest_water, real_pan, home_real_water, guest_real_water, home_pan_change_type, schedule_game_desc, schedulefid, cid)
		predict_result, predict_cmt := analyse.AnalysePanResult2(pan_int_info, pan_float_info, pan_string_info)
		pan_string_info["predict_result"] = predict_result
		pan_string_info["predict_cmt"] = predict_cmt

		// fmt.Println("float_open_pan")
		// fmt.Println(open_home_water)
		// fmt.Println("=====")

		// fmt.Println("company:" + company)
		// fmt.Println("home_pan_change_type:" + home_pan_change_type)
		// fmt.Println("is big company:" + is_big_company)
		// fmt.Println("change_time:" + change_time)
		// fmt.Println("open_time:" + open_time)
		// fmt.Println("flag:" + home_water_change_type + " " + home_water_up_down_flag)
		// fmt.Println("home_real_water:", home_real_water)
		// fmt.Println("home_real_water water sting:" + home_real_water_string)
		// fmt.Println("guest_real_water:", guest_real_water)
		// fmt.Println("guest_real_water water sting:" + guest_real_water_string)
		// fmt.Println("pan:", real_pan, " ", real_pan_desc)

		// fmt.Println("open_home_water water:", open_home_water)
		// fmt.Println("open_guest_water water:", open_guest_water)
		// fmt.Println("open pan:", open_pan, " ", open_pan_desc)

		exist_lastpan := new(myinit.LastPan)
		has, _ := myinit.Engine.Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Get(exist_lastpan)
		if has {
			fmt.Println(company + "pan已存在！")
			if exist_lastpan.LastPan != real_pan || exist_lastpan.LastHomeWater != home_real_water || exist_lastpan.LastChangeTime != change_time {
				fmt.Println(company + "pan有变化！")

				update_affected, update_err := lastpan.updateLastPanInfo(pan_int_info, pan_float_info, pan_string_info)
				fmt.Println(update_affected)
				fmt.Println(update_err)

				lastpanlog.Add(pan_int_info, pan_float_info, pan_string_info)
			} else {

			}
		} else {
			lastpan.Add(pan_int_info, pan_float_info, pan_string_info)
			lastpanlog.Add(pan_int_info, pan_float_info, pan_string_info)
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


func main() {
	myinit.Myinit()
	pan_url = myinit.PanUrl

	x, _ := goquery.ParseUrl(myinit.IndexUrl)
	schedule_trs := x.Find(".bet_table tbody tr")
	for i, _ := range schedule_trs {
		is_end := schedule_trs.Eq(i).Attr("isend")
		if is_end == "1" {
			fmt.Println("is_end")
			continue
		}

		schedule_int_info := make(map[string]int)
		schedule_string_info := make(map[string]string)

		// insert schedule
		fid, _ := strconv.ParseInt(schedule_trs.Eq(i).Attr("fid"), 10, 0)
		schedule_int_info["fid"] = int(fid)
		schedule_string_info["home_team"] = common.ConvToGB(schedule_trs.Eq(i).Attr("homesxname"))
		schedule_string_info["guest_team"] = common.ConvToGB(schedule_trs.Eq(i).Attr("awaysxname"))
		schedule_string_info["schedule_date"] = schedule_trs.Eq(i).Attr("pdate")
		schedule_string_info["schedule_league"] = common.ConvToGB(schedule_trs.Eq(i).Attr("lg"))
		schedule_string_info["game_date_no"] = schedule_trs.Eq(i).Attr("gdate")
		schedule_string_info["schedule_no"] = schedule_trs.Eq(i).Attr("pname")
		schedule_string_info["schedule_result_no"] = schedule_trs.Eq(i).Attr("pname2")
		schedule_string_info["end_time"] = schedule_trs.Eq(i).Attr("pendtime")
		
		today := time.Now().Format("2006-01-02")
		schedule_is_today := today==schedule_string_info["schedule_date"]
		fmt.Println("schedule_is_today:===")
		fmt.Println(schedule_is_today)
		fmt.Println("end schedule_is_today===")
		if(schedule_is_today == false){
			continue
		}
		schedule.Add(schedule_int_info, schedule_string_info)
		// end insert schedule

		schedule_pan_url := strings.Replace(pan_url, "TTT", fid, -1)
		fmt.Println(schedule_odds_url)
		//		go parseOddUrl(schedule_odds_url, fid)
		res := ParseOddUrl(schedule_pan_url, fid, schedule_string_info["schedule_no"], schedule_string_info["schedule_date"])
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
