package parseurl

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	
	"500kan/util/myinit"
	"500kan/util/common"
	"500kan/util/schedule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func ParseBetUrl(date string,history bool){
	
	bet_url := strings.Replace(myinit.DateUrl, "TTT", date, -1)
	fmt.Println(bet_url)
	
	pan_url := myinit.PanUrl

	html_obj, _ := goquery.ParseUrl(bet_url)
	schedule_trs := html_obj.Find(".bet_table tbody tr")
	for i, _ := range schedule_trs {
		is_end := schedule_trs.Eq(i).Attr("isend")
		if is_end == "1" && history==false {
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
		schedule_string_info["schedule_bet_date"] = date
		schedule_string_info["schedule_league"] = common.ConvToGB(schedule_trs.Eq(i).Attr("lg"))
		schedule_string_info["schedule_week_day"] = common.ConvToGB(schedule_trs.Eq(i).Attr("gdate"))
		schedule_string_info["schedule_no"] = schedule_trs.Eq(i).Attr("pname")
		schedule_string_info["schedule_rq_num"] = schedule_trs.Eq(i).Attr("rq")
		week_date := schedule_string_info["schedule_no"][0:1]		
		schedule_string_info["schedule_result_no"] = myinit.WeekDesc[week_date] + schedule_string_info["schedule_no"][1:]
		
		schedule_string_info["schedule_bet_end_time"] = schedule_trs.Eq(i).Attr("pendtime")
//fmt.Println("------")
//fmt.Println(schedule_string_info["schedule_no"][0:1])
//fmt.Println(myinit.WeekDesc[week_date])
//fmt.Println(schedule_string_info["schedule_no"][1:])
fmt.Println(schedule_string_info["schedule_result_no"])
//fmt.Println("------")
//		fmt.Println(schedule_string_info["schedule_league"])
//		fmt.Println(schedule_string_info["schedule_week_day"])
//		fmt.Println(schedule_string_info["schedule_rq_num"])
//		fmt.Println(schedule_trs.Eq(i).OuterHtml())
//		fmt.Println(schedule_string_info["schedule_result_no"])
		
		today := time.Now().Format("2006-01-02")
		schedule_is_today := today == schedule_string_info["schedule_date"]
		fmt.Println("schedule_is_today:===")
		fmt.Println(schedule_is_today)
		if schedule_is_today == false && history==false {
			continue
		}
		schedule.Add(schedule_int_info, schedule_string_info)
		// end insert schedule
		

		schedule_pan_url := strings.Replace(pan_url, "TTT", strconv.Itoa(fid), -1)
		fmt.Println(schedule_pan_url)
		//		go parseOddUrl(schedule_odds_url, fid)
		res := ParsePanUrl(schedule_pan_url, fid, schedule_string_info,date )
		fmt.Println(res)
	}
}
