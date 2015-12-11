package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/guotie/gogb2312"
	"github.com/opesun/goquery"
)

const AVG_WATER float32 = 0.875

type Schedule struct {
	ScheduleId       string
	ScheduleHome     string
	ScheduleGuest    string
	ScheduleDate     string
	ScheduleLeague   string
	ScheduleWeekDay  string
	ScheduleNo       string
	ScheduleFid      string
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

func convToGB(str string) (res_str string) {
	conv_str, _, _, _ := gogb2312.ConvertGB2312String(str)
	return conv_str
}

var engine *xorm.Engine

func ParseOddUrl(schedule_odds_url string, schedulefid string, pname string, schedule_date_orig string) (a string) {
	odds_html, _ := goquery.ParseUrl(schedule_odds_url)

	schedule_item := odds_html.Find(".odds_hd_cont table tbody tr td")
	home_td := schedule_item.Eq(0)
	guest_td := schedule_item.Eq(4)
	center_td := schedule_item.Eq(2)

	schedule_home_name := convToGB(home_td.Find("ul li a").Text())
	schedule_guest_name := convToGB(guest_td.Find("ul li a").Text())
	schedule_game_desc := convToGB(center_td.Find(".odds_hd_center .odds_hd_ls a").Text())
//	schedule_date := convToGB(center_td.Find(".odds_hd_center .game_time ").Text())

	schedule_date := schedule_date_orig

	odds_tr := odds_html.Find(".table_cont table tbody tr")
	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)

		company := convToGB(td_of_company.Find("p a").Attr("title"))
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

		cid := tr_item.Attr("id")

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
			delete_lastpan := new(LastPan)
			del_result, _ := engine.Where("schedule_fid=? ", schedulefid).Delete(delete_lastpan)
			delete_schedule := new(Schedule)
			del_schedule_result, _ := engine.Where("schedule_fid=? ", schedulefid).Delete(delete_schedule)
			fmt.Println(del_result, del_schedule_result)
			return "开盘>0 或者即时盘 >0"
		}
		open_pan_desc := convToGB(tds_of_opentime_pan_table.Eq(1).Text())

		//		addLastPanLog()

		predict_result, predict_cmt := analysePanResult(open_pan, open_home_water, open_guest_water, real_pan, home_real_water, guest_real_water, home_pan_change_type, schedule_game_desc, schedulefid, cid)

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

		exist_lastpan := new(LastPan)
		has, _ := engine.Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Get(exist_lastpan)
		fmt.Println(has)
		if has {
			fmt.Println(company + "pan已存在！")
			if exist_lastpan.LastPan != real_pan || exist_lastpan.LastHomeWater != home_real_water || exist_lastpan.LastChangeTime != change_time {
				fmt.Println(company + "pan有变化！")
				update_lastpan := new(LastPan)
				update_lastpan.LastPan = real_pan
				update_lastpan.LastPanDesc = real_pan_desc
				update_lastpan.LastHomeWater = home_real_water
				update_lastpan.LastGuestWater = guest_real_water
				update_lastpan.LastHomeWaterChangeType = home_water_change_type
				update_lastpan.LastHomePanChangeType = home_pan_change_type
				update_lastpan.LastChangeTime = change_time
				update_lastpan.PredictResult = predict_result
				update_lastpan.PredictComment = predict_cmt
				update_affected, update_err := engine.Cols("last_pan", "last_pan_desc", "last_hone_water", "last_guest_water", "last_change_time", "last_home_pan_change_type", "last_home_water_change_type", "predict_result", "predict_comment").Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Update(update_lastpan)
				fmt.Println(update_affected)
				fmt.Println(update_err)

				LastPanLog := new(LastPanLog)
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

				log_ins_affected, log_ins_err := engine.Insert(LastPanLog)
				fmt.Println(log_ins_affected)
				fmt.Println(log_ins_err)
			} else {

			}
		} else {
			LastPan := new(LastPan)
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

			ins_affected, ins_err := engine.Insert(LastPan)
			fmt.Println(ins_affected)
			fmt.Println(ins_err)

			LastPanLog := new(LastPanLog)
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

			log_ins_affected, log_ins_err := engine.Insert(LastPanLog)
			fmt.Println(log_ins_affected)
			fmt.Println(log_ins_err)
		}

		//		count_open_water := open_home_water +open_guest_water
		count_real_water := home_real_water + guest_real_water
		if count_real_water < 1.75 || count_real_water > 2 {
			delete_lastpan2 := new(LastPan)
			delete2, _ := engine.Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Delete(delete_lastpan2)
			fmt.Println(delete2)
		}
	}
	return "成功"
}

func analysePanResult(open_pan float32, open_pan_home_water float32, open_pan_guest_water float32, real_pan float32, real_pan_home_water float32, real_pan_guest_water float32, home_pan_change_type string, schedule_game_desc string, schedulefid string, cid string) (ret string, cmt string) {
	fmt.Println("+++++++++++")
	fmt.Println(schedulefid)
	fmt.Println(home_pan_change_type)

	switch {
	case open_pan == 0:
		fmt.Println("0 open:", open_pan)

		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜"
				if real_pan_home_water < open_pan_home_water {
					ret = "1/0"
					cmt = "主队水位	＞0.875	主队即时盘口水位小于初盘水位，多为平局"
				}
			}

			if open_pan_home_water == open_pan_guest_water {
				ret = "1"
				cmt = "主队水位=	客队水位	平"
				if real_pan_home_water < 0.875 {
					ret = "3"
					cmt = "主队水位=	客队水位	平，即时水位＜0.875队胜出"

				}
			}

			if checkPanAndWaterNotChange(schedulefid,cid)==true && real_pan_home_water < 0.875 {
				ret = "3"
				cmt = "盘口、水位一直不变		即时水位＜0.875队胜出"
			}
		} else if home_pan_change_type == "升" {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜  多为平局"
			}
		} else if home_pan_change_type == "降" {
			if real_pan_guest_water > 0.875 {
				ret = "3/1"
				cmt = "客队水位	＞0.875	主队胜或平"
			} else {
				ret = "1/0"
				cmt = "客队水位	≤0.875	平或客队胜"
			}
		}

		fmt.Println("open:", open_pan, ret, cmt)
	case open_pan == (-0.25):
		fmt.Println("-0.25 open:", open_pan)
		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜"
			}

			if open_pan_home_water == open_pan_guest_water && real_pan_home_water > 0.875 {
				ret = "3"
				cmt = "主队水位=	客队水位	即时水位＞0.875队胜出	"
			}
			if checkPanAndWaterNotChange(schedulefid,cid)==true  {
				ret = "3/0"
				cmt = "盘口、水位一直不变		双方能分胜负		德甲主队胜概率大"
			}
			if checkPanNotChange(schedulefid,cid)==true && checkIsGermanyJia(schedule_game_desc) == true {
				ret = "1/0"
				cmt = "若为德甲，盘口不变而水位发生变化们一般是下盘胜出			对应结果：	1/0	"
			}
		} else if home_pan_change_type == "升" {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 && checkWaterIsDown(schedulefid, cid) {
					ret = "3"
					cmt = "主队水位	≤0.875	即时水位＞0.875并且水位持续下降	主队胜"
				} else if real_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "主队水位	≤0.875	即时水位≤0.875	平或客队胜"
				}
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875		平或客队胜 多为平局"
			}
		} else if home_pan_change_type == "降" {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water <= 0.875 {
					ret = "0"
					cmt = "主队水位	≤0.875	即时水位≤0.875	客队胜"
				} else {
					ret = "1"
					cmt = "主队水位	≤0.875	即时水位＞0.875	平"
				}
//			} else {
//				ret = "1/0"
//				cmt = "其余情况			平或客队胜	对应结果：	1/0"
			}

		}
	case open_pan == (-0.5):
		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜	对应结果：	3"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜	对应结果：	1/0"
			}

			if checkPanAndWaterNotChange(schedulefid,cid)==true {
				if open_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "盘口、水位一直不变		初盘水位	主队水位	≤0.875	平或客队胜	对应结果：	1/0"
				} else {
					ret = "3"
					cmt = "盘口、水位一直不变		初盘水位	主队水位	＞0.875	主队胜	对应结果：	3"
				}
			}

		} else if home_pan_change_type == "升" {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 {
					ret = "3"
					cmt = "主队水位	≤0.875	即时水位＞0.875	主队胜	对应结果：	3"
				} else {
					ret = "1"
					cmt = "主队水位	≤0.875	即时水位≤0.875	平	对应结果：	1"
				}
			} else {
				ret = "0"
				cmt = "主队水位	＞0.875		客队胜	对应结果：	0"
			}
		} else if home_pan_change_type == "降" {
			fmt.Println("-0.5====")
			fmt.Println(open_pan_home_water)
			fmt.Println(real_pan_home_water)
			if open_pan_home_water > 0.875 {
				if real_pan_home_water <= 0.875 {
					ret = "3/1"
					cmt = "主队水位	＞0.875	即时水位≤0.875	主胜或平	对应结果：	3、1"
				} else {
					ret = "0"
					cmt = "主队水位	＞0.875	即时水位＞0.875	客队胜	对应结果：	0"
				}
//			} else {
//				ret = "1/0"
//				cmt = "其余情况			平或客队胜	对应结果：	1/0"
			}
		}
		fmt.Println("-0.5 open:", open_pan)

	case open_pan == (-0.75):
		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜	对应结果：	3"
			} else {
				if real_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "主队水位	＞0.875	即时水位≤0.875	平或客队胜	对应结果：	1/0"
				} else {
					ret = "3"
					cmt = "主队水位	＞0.875	即时水位＞0.875	主队胜	对应结果：	3"
				}
				if real_pan_home_water == open_pan_home_water {
					ret = "0"
					cmt = "主队水位	＞0.875	即时水位=初盘水位	客队胜	对应结果：	0"
				}
			}

		} else if home_pan_change_type == "升" {
		} else if home_pan_change_type == "降" {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 {
					ret = "1"
					cmt = "主队水位	≤0.875	即时水位＞0.875	平	对应结果：	1"
				} else {
					ret = "0"
					cmt = "主队水位	≤0.875	即时水位≤0.875	客队胜	对应结果：	0"
				}
			}

		}
		fmt.Println("-0.75: open:", open_pan)

	case open_pan <= -1:
		//	case -1.25:
		//	case -1.5:
		flag := false
		if(checkPanNotChange(schedulefid,cid)==true && checkWaterNotChange(schedulefid,cid)==false ){
			if open_pan_home_water > 0.875 {
				ret = "3"
				cmt = "主队水位	＞0.875	主队胜	对应结果：	3"	
				flag = true		
			
			}
		} 
		if(checkPanNotChange(schedulefid,cid)==false && open_pan==(-1) && real_pan!=(-1)){
			if open_pan_home_water > 0.875 && real_pan_home_water <= 0.875 {
				ret = "3/1"
				cmt = "主队水位	＞0.875	即时水位≤0.875"
				flag = true		
			}
		}
		
		if(open_pan<(-1.5) && checkPanNotChange(schedulefid,cid)==true){
			if(open_pan_home_water< 0.8){
				ret = "3/0"
				cmt = "初盘盘口数值＜-1.5） 初盘水位	主队水位	＜0.8	有爆冷可能	胜或负 "
				flag = true		
			}
		}
		
		if(flag==false){
				if open_pan_home_water < 0.875 {
				//其他情况
				ret = "3"
				cmt = "其余情况 初盘水位	主队水位	＜0.875	主队胜	对应结果：	3"	
				}		
		}
		fmt.Println("-1 open:", open_pan)
	default:
		fmt.Println("qita open:", open_pan)
		ret = ""
		cmt = ""

	}
	return ret, cmt
}

func checkIsGermanyJia(str string) (ret bool) {
	return strings.Contains(str, "德甲")
}

func checkWaterIsDown(fid string, cid string) (ret bool) {
	exist_up := new(LastPanLog)
	total, _ := engine.Where("last_home_water_change_type='up' AND schedule_fid=? AND company_cid=?", fid, cid).Count(exist_up)
	if total > 0 {
		return false
	}
	return true
}

func checkPanNotChange(fid string, cid string) (ret bool) {
	exist_up := new(LastPanLog)
	total_pan_change, _ := engine.Where("open_pan!=last_pan AND schedule_fid=? AND company_cid=?", fid, cid).Count(exist_up)
	if total_pan_change > 0 {
		return false
	}	
	return true
}

func checkWaterNotChange(fid string, cid string) (ret bool) {
	exist_up := new(LastPanLog)
	total_water_change, _ := engine.Where("open_home_water!=last_home_water AND schedule_fid=? AND company_cid=?", fid, cid).Count(exist_up)
	if total_water_change > 0 {
		return false
	}	
	return true
}


func checkPanAndWaterNotChange(fid string, cid string) (ret bool) {
	if(checkPanNotChange(fid,cid)==false){
		return false;
	}
	if(checkWaterNotChange(fid,cid)==false){
		return false;
	}	
	return true
}



func initDb() {
	var err error
	//	engine, err = xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/test_ha")
//	        engine, err = xorm.NewEngine("mysql", "root:123456@tcp(192.168.1.172:3306)/jc_test")
	engine, err = xorm.NewEngine("mysql", "qichejingli:qichejingli1234QWER@tcp(rds3bhb1ed059c58i02wo.mysql.rds.aliyuncs.com:3306)/test_ha")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "pk_")
	engine.SetTableMapper(tbMapper)

	engine.ShowSQL = true   //则会在控制台打印出生成的SQL语句；
	engine.ShowDebug = true //则会在控制台打印调试信息；
	engine.ShowErr = true   //则会在控制台打印错误信息；
	engine.ShowWarn = true  //则会在控制台打印警告信息；

	f, err := os.Create("sql.log")
	if err != nil {
		println("sql.log error:")
		println(err.Error())
		return
	}
	defer f.Close()
	engine.Logger = xorm.NewSimpleLogger(f)
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
	initDb()

	x, _ := goquery.ParseUrl(index_url)
	schedule_trs := x.Find(".bet_table tbody tr")
	for i, _ := range schedule_trs {
		is_end := schedule_trs.Eq(i).Attr("isend")
		if is_end == "1" {
			fmt.Println("is_end")
			continue
		}

		// insert schedule
		fid := schedule_trs.Eq(i).Attr("fid")
		home_team := convToGB(schedule_trs.Eq(i).Attr("homesxname"))
		guest_team := convToGB(schedule_trs.Eq(i).Attr("awaysxname"))
		schedule_date := schedule_trs.Eq(i).Attr("pdate")
		lg := convToGB(schedule_trs.Eq(i).Attr("lg"))
		game_date_no := schedule_trs.Eq(i).Attr("gdate")
		schedule_no := schedule_trs.Eq(i).Attr("pname")
		end_time := schedule_trs.Eq(i).Attr("pendtime")

		now := time.Now()
		year, mon, day := now.Date()
		schedule_is_today := compareDate(end_time, time.Date(year, mon, day, 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05"))
		fmt.Println("schedule_is_today:===")
		fmt.Println(schedule_is_today)
		fmt.Println("end schedule_is_today===")
		if schedule_is_today == false {
			continue
		}
		exist_schedule := new(Schedule)
		has, _ := engine.Where("schedule_date=? AND schedule_no=? ", schedule_date, schedule_no).Get(exist_schedule)
		fmt.Println(has)
		if has {
			fmt.Println(home_team + " vs " + guest_team + "已存在！")
		} else {
			Schedule := new(Schedule)
			Schedule.ScheduleFid = fid
			Schedule.ScheduleHome = home_team
			Schedule.ScheduleGuest = guest_team
			Schedule.ScheduleDate = schedule_date
			Schedule.ScheduleLeague = lg
			Schedule.ScheduleWeekDay = game_date_no
			Schedule.ScheduleNo = schedule_no
			Schedule.ScheduleEndTime = end_time

			affected, _ := engine.Insert(Schedule)
			fmt.Println(affected)
			fmt.Println(Schedule.ScheduleId)
		}
		// end insert schedule

		schedule_odds_url := strings.Replace(odds_url, "TTT", fid, -1)
		fmt.Println(schedule_odds_url)
		//		go parseOddUrl(schedule_odds_url, fid)
		res := ParseOddUrl(schedule_odds_url, fid, schedule_no, schedule_date)
		fmt.Println(res)
	}
	moveToBackup()
}

func moveToBackup() {
	now := time.Now()
	today := now.Format("2006-01-02")
	fmt.Println(today)
	sql := "insert into `pk_last_pan_backup` select * from `pk_last_pan` where schedule_date < ?"
	ins_res, ins_err := engine.Exec(sql, today)
	fmt.Println(ins_err, ins_res)
	del_sql := "delete from `pk_last_pan` where schedule_date < ?"
	del_res, del_err := engine.Exec(del_sql, today)
	fmt.Println(del_res, del_err)

}
