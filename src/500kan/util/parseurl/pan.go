package parseurl

import (
	"fmt"
	"strconv"
	"strings"

	"500kan/util/analyse"
	"500kan/util/asiapan"
	"500kan/util/asiapanbackup"
	"500kan/util/asiapanlog"
	"500kan/util/common"
	"500kan/util/myinit"
	"500kan/util/schedule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func ParsePanUrl(schedule_pan_url string, schedule_fenxi_id int, schedule_string_info map[string]string, date string) (res bool) {
	useable := checkPanUseable(schedule_pan_url, schedule_fenxi_id, schedule_string_info, date)
	//	fmt.Println("userabel", useable)
	if useable == false {
		return false
	}
	do_success := doParsePanUrl(schedule_pan_url, schedule_fenxi_id, schedule_string_info, date)
	return do_success
}

func checkPanUseable(schedule_pan_url string, schedule_fenxi_id int, schedule_string_info map[string]string, date string) (res bool) {
	pan_float_info := make(map[string]float32)

	pan_html_obj, _ := goquery.ParseUrl(schedule_pan_url)

	odds_tr := pan_html_obj.Find(".table_cont table tbody tr")
	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)
		if td_of_company.Find("p a").Attr("title") == "" {
			continue
		}
		table_of_pan_detail := tr_item.Find("td .pl_table_data")

		table_of_opentime_pan := table_of_pan_detail.Eq(1)
		tds_of_opentime_pan_table := table_of_opentime_pan.Find("tbody tr td")
		open_pan_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["open_pan"] = float32(open_pan_32)

		open_home_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(0).Text(), 32)
		open_guest_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(2).Text(), 32)
		pan_float_info["open_home_water"] = float32(open_home_water_32)
		pan_float_info["open_guest_water"] = float32(open_guest_water_32)

		table_of_realtime_pan := table_of_pan_detail.Eq(0)
		tds_of_realtime_pan_table := table_of_realtime_pan.Find("tbody tr td")

		real_pan_32, _ := strconv.ParseFloat(tds_of_realtime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["real_pan"] = float32(real_pan_32)

		home_real_water_string := common.ConvToGB(tds_of_realtime_pan_table.Eq(0).Text())
		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_string := common.ConvToGB(tds_of_realtime_pan_table.Eq(2).Text())
		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		pan_float_info["real_home_water"] = float32(home_real_water_32)
		pan_float_info["real_guest_water"] = float32(guest_real_water_32)
		if pan_float_info["open_pan"] > 0 || pan_float_info["real_pan"] > 0 {
			asiapan.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			asiapanbackup.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			schedule.ClearScheduleByFenxiId(schedule_fenxi_id)
			fmt.Println("pan html:", pan_float_info["open_pan"], pan_float_info["real_pan"])
			fmt.Println("开盘>0 或者即时盘 >0")
			return false
		}

	}
	if odds_tr.Length() >= 30 {
		ajax_res := checkPanUseableFromAjax(30, schedule_fenxi_id)
		if ajax_res == false {
			return false
		}
	}
	return true
}

func checkPanUseableFromAjax(idx int, schedule_fenxi_id int) (res bool) {
	pan_float_info := make(map[string]float32)
	odd_html := myinit.GetOddItemFromAjax(idx, schedule_fenxi_id)

	table_string := "<table>" + odd_html + "</table>"
	html_obj, _ := goquery.ParseString(table_string)

	odds_tr := html_obj.Find("table tbody tr")

	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)
		if td_of_company.Find("p a").Attr("title") == "" {
			continue
		}

		table_of_pan_detail := tr_item.Find("td .pl_table_data")

		table_of_opentime_pan := table_of_pan_detail.Eq(1)
		tds_of_opentime_pan_table := table_of_opentime_pan.Find("tbody tr td")
		open_pan_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["open_pan"] = float32(open_pan_32)

		open_home_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(0).Text(), 32)
		open_guest_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(2).Text(), 32)
		pan_float_info["open_home_water"] = float32(open_home_water_32)
		pan_float_info["open_guest_water"] = float32(open_guest_water_32)

		table_of_realtime_pan := table_of_pan_detail.Eq(0)
		tds_of_realtime_pan_table := table_of_realtime_pan.Find("tbody tr td")

		real_pan_32, _ := strconv.ParseFloat(tds_of_realtime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["real_pan"] = float32(real_pan_32)

		home_real_water_string := tds_of_realtime_pan_table.Eq(0).Text()
		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_string := tds_of_realtime_pan_table.Eq(2).Text()
		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		pan_float_info["real_home_water"] = float32(home_real_water_32)
		pan_float_info["real_guest_water"] = float32(guest_real_water_32)

		if pan_float_info["open_pan"] > 0 || pan_float_info["real_pan"] > 0 {
			asiapan.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			asiapanbackup.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			schedule.ClearScheduleByFenxiId(schedule_fenxi_id)
			fmt.Println("pan ajx:", pan_float_info["open_pan"], pan_float_info["real_pan"])

			fmt.Println("开盘>0 或者即时盘 >0")
			return false
		}

	}
	if odds_tr.Length() >= 30 {
		ajax_res := checkPanUseableFromAjax(idx+30, schedule_fenxi_id)
		if ajax_res == false {
			return false
		}
	}

	return true
}

func doParsePanUrl(schedule_pan_url string, schedule_fenxi_id int, schedule_string_info map[string]string, date string) (res bool) {
	pan_int_info := make(map[string]int)
	pan_float_info := make(map[string]float32)
	pan_string_info := make(map[string]string)

	pan_html_obj, _ := goquery.ParseUrl(schedule_pan_url)

	schedule_item := pan_html_obj.Find(".odds_hd_cont table tbody tr td")
	home_td := schedule_item.Eq(0)
	guest_td := schedule_item.Eq(4)
	center_td := schedule_item.Eq(2)

	pan_int_info["schedule_fenxi_id"] = schedule_fenxi_id
	pan_string_info["schedule_date"] = schedule_string_info["schedule_date"]
	pan_string_info["schedule_no"] = schedule_string_info["schedule_no"]
	pan_string_info["schedule_result_no"] = schedule_string_info["schedule_result_no"]
	pan_string_info["schedule_league"] = schedule_string_info["schedule_league"]

	pan_string_info["schedule_home"] = common.ConvToGB(home_td.Find("ul li a").Text())
	pan_string_info["schedule_guest"] = common.ConvToGB(guest_td.Find("ul li a").Text())
	pan_string_info["schedule_game_desc"] = common.ConvToGB(center_td.Find(".odds_hd_center .odds_hd_ls a").Text())
	pan_string_info["schedule_date_desc"] = common.ConvToGB(center_td.Find(".odds_hd_center .game_time ").Text())

	odds_tr := pan_html_obj.Find(".table_cont table tbody tr")
	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)
		if td_of_company.Find("p a").Attr("title") == "" {
			continue
		}

		company_id := tr_item.Attr("id")
		pan_string_info["company_id"] = company_id
		pan_string_info["company_name"] = common.ConvToGB(td_of_company.Find("p a").Attr("title"))

		var is_big_company = 0
		if td_of_company.Find("p img").Attr("src") == "" {
			is_big_company = 0
		} else {
			is_big_company = 1
			//			fmt.Println("src:" + td_of_company.Find("p img").Attr("src"))
		}
		pan_int_info["is_big_company"] = is_big_company

		table_of_pan_detail := tr_item.Find("td .pl_table_data")

		table_of_opentime_pan := table_of_pan_detail.Eq(1)
		tds_of_opentime_pan_table := table_of_opentime_pan.Find("tbody tr td")
		open_pan_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["open_pan"] = float32(open_pan_32)
		pan_string_info["open_pan_desc"] = common.ConvToGB(tds_of_opentime_pan_table.Eq(1).Text())

		open_home_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(0).Text(), 32)
		open_guest_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(2).Text(), 32)
		pan_float_info["open_home_water"] = float32(open_home_water_32)
		pan_float_info["open_guest_water"] = float32(open_guest_water_32)

		td_of_pan_time := tr_item.Find("td time")

		pan_string_info["open_pan_time"] = td_of_pan_time.Eq(1).Text()

		table_of_realtime_pan := table_of_pan_detail.Eq(0)
		tds_of_realtime_pan_table := table_of_realtime_pan.Find("tbody tr td")

		real_pan_32, _ := strconv.ParseFloat(tds_of_realtime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["real_pan"] = float32(real_pan_32)
		pan_string_info["real_pan_desc"] = common.ConvToGB(tds_of_realtime_pan_table.Eq(1).Text())

		home_real_water_string := common.ConvToGB(tds_of_realtime_pan_table.Eq(0).Text())
		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_string := common.ConvToGB(tds_of_realtime_pan_table.Eq(2).Text())
		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		pan_float_info["real_home_water"] = float32(home_real_water_32)
		pan_float_info["real_guest_water"] = float32(guest_real_water_32)

		pan_string_info["pan_change_time"] = strings.TrimSpace(td_of_pan_time.Eq(0).Text())

		td_item_of_real_pan := tds_of_realtime_pan_table.Eq(1)
		home_pan_change_type := common.ConvToGB(td_item_of_real_pan.Find("font").Text())
		home_pan_change_type = strings.TrimSpace(home_pan_change_type)
		fmt.Println("home_pan_desc:", home_pan_change_type)

		pan_int_info["home_pan_change_type"] = 0
		pan_string_info["home_pan_change_type_desc"] = ""
		if home_pan_change_type == "升" {
			pan_int_info["home_pan_change_type"] = 1
			pan_string_info["home_pan_change_type_desc"] = home_pan_change_type
		}
		if home_pan_change_type == "降" {
			pan_int_info["home_pan_change_type"] = -1
			pan_string_info["home_pan_change_type_desc"] = home_pan_change_type
		}
		fmt.Println("home_pan_change_type_desc:", pan_string_info["home_pan_change_type_desc"])

		home_water_up_down_flag := tds_of_realtime_pan_table.Eq(0).Attr("class")
		pan_int_info["home_water_change_type"] = 0
		pan_string_info["home_water_change_type_desc"] = ""

		if home_water_up_down_flag == "ping" {
			pan_int_info["home_water_change_type"] = -1            // down
			pan_string_info["home_water_change_type_desc"] = "水位降" // down
		}
		if home_water_up_down_flag == "ying" {
			pan_int_info["home_water_change_type"] = 1             // up
			pan_string_info["home_water_change_type_desc"] = "水位升" // up
		}

		//		real_pan_string := strings.Replace(pan_string_info["real_pan_desc"], pan_string_info["home_pan_change_type_desc"], "", -1)
		//		real_pan_desc := strings.TrimSpace(real_pan_string)
		//		fmt.Println("date:", pan_string_info["schedule_date"], pan_string_info["schedule_home"])
		//		fmt.Println("open:", pan_string_info["open_pan_desc"], pan_float_info["open_pan"])
		//		fmt.Println("real:", real_pan_desc, pan_float_info["real_pan"])

		exist_asiapanlog := new(myinit.AsiaPanLog)
		has, f_err := myinit.Engine.Where("schedule_fenxi_id=? AND company_id=? AND pan_change_time=?", schedule_fenxi_id, company_id, pan_string_info["pan_change_time"]).Get(exist_asiapanlog)
		fmt.Println("has:", has, ":", schedule_fenxi_id, ":", company_id, ":", pan_string_info["pan_change_time"])
		fmt.Println("ajax:", f_err)

		if has == true {
			fmt.Println("go has:is:", pan_string_info["pan_change_time"])
			continue
		}

		if pan_float_info["open_pan"] > 0 || pan_float_info["real_pan"] > 0 {
			asiapan.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			asiapanbackup.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)

			schedule.ClearScheduleByFenxiId(schedule_fenxi_id)
			//			fmt.Println("pan html:", pan_float_info["open_pan"], pan_float_info["real_pan"])
			fmt.Println("开盘>0 或者即时盘 >0")
			return false
		}

		predict1_result, predict1_cmt := analyse.AnalysePanResult1(pan_int_info, pan_float_info, pan_string_info)
		pan_string_info["predict1_result"] = predict1_result
		pan_string_info["predict1_cmt"] = predict1_cmt
		predict2_result, predict2_cmt := analyse.AnalysePanResult2(pan_int_info, pan_float_info, pan_string_info)
		pan_string_info["predict2_result"] = predict2_result
		pan_string_info["predict2_cmt"] = predict2_cmt

		// fmt.Println("company:" + company)
		// fmt.Println("home_pan_change_type:" + home_pan_change_type)
		// fmt.Println("is big company:" + is_big_company)
		// fmt.Println("change_time:" + change_time)
		// fmt.Println("open_time:" + open_time)
		// fmt.Println("flag:" + home_water_change_type + " " + home_water_up_down_flag)
		//		fmt.Println("home_real_water:", pan_float_info["real_home_water"])
		//		fmt.Println(home_real_water_32)
		//		fmt.Println(float32(home_real_water_32))
		//		fmt.Println("home_real_water water sting:" + home_real_water_str + home_real_water_string)
		//		fmt.Println("guest_real_water:", pan_float_info["real_guest_water"])
		//		fmt.Println("guest_real_water water sting:" + guest_real_water_str + guest_real_water_string)
		//		fmt.Println("pan:", pan_float_info["real_pan"], " ", real_pan_desc)

		// fmt.Println("open_home_water water:", open_home_water)
		// fmt.Println("open_guest_water water:", open_guest_water)
		// fmt.Println("open pan:", open_pan, " ", open_pan_desc)

		addAsiaPan(schedule_fenxi_id, company_id, pan_int_info, pan_float_info, pan_string_info)
		//delete water > 2 and < 1.75
		checkWaterSum(schedule_fenxi_id, company_id, pan_float_info)
	}
	if odds_tr.Length() >= 30 {
		ParsePanUrlFromAjax(30, schedule_fenxi_id, pan_string_info)
	}

	return true
}

func addAsiaPan(schedule_fenxi_id int, company_id string, pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) {
	exist_asiapan := new(myinit.AsiaPan)
	has, _ := myinit.Engine.Where("schedule_fenxi_id=? AND company_id=? ", schedule_fenxi_id, company_id).Get(exist_asiapan)
	if has {
		fmt.Println(pan_string_info["company_name"] + "pan已存在！")
		if exist_asiapan.PanChangeTime != pan_string_info["pan_change_time"] {
			fmt.Println(pan_string_info["company_name"] + "pan有变化！")
			asiapan.UpdateAsiaPanInfo(pan_int_info, pan_float_info, pan_string_info)
			asiapanbackup.UpdateAsiaPanBackupInfo(pan_int_info, pan_float_info, pan_string_info)
			asiapanlog.Add(pan_int_info, pan_float_info, pan_string_info)
		}
	} else {
		asiapan.Add(pan_int_info, pan_float_info, pan_string_info)
		asiapanbackup.Add(pan_int_info, pan_float_info, pan_string_info)
		asiapanlog.Add(pan_int_info, pan_float_info, pan_string_info)
	}
	//delete water > 2 and < 1.75
	checkWaterSum(schedule_fenxi_id, company_id, pan_float_info)
}

func checkWaterSum(schedule_fenxi_id int, company_id string, pan_float_info map[string]float32) (res bool) {
	sum_real_water := pan_float_info["real_home_water"] + pan_float_info["real_guest_water"]
	sum_open_water := pan_float_info["open_home_water"] + pan_float_info["open_guest_water"]

	if (sum_real_water < 1.75 || sum_real_water > 2) || (sum_open_water < 1.75 || sum_open_water > 2) {
		asiapan.DeletePanByFenxiIdAndCompanyId(schedule_fenxi_id, company_id)
		asiapanbackup.DeletePanByFenxiIdAndCompanyId(schedule_fenxi_id, company_id)
		return false
	}
	return true
}

func ParsePanUrlFromAjax(idx int, schedule_fenxi_id int, pan_html_string_info map[string]string) (res bool) {
	pan_int_info := make(map[string]int)
	pan_float_info := make(map[string]float32)
	pan_string_info := make(map[string]string)

	odd_html := myinit.GetOddItemFromAjax(idx, schedule_fenxi_id)

	table_string := "<table>" + odd_html + "</table>"
	html_obj, _ := goquery.ParseString(table_string)

	odds_tr := html_obj.Find("table tbody tr")

	pan_int_info["schedule_fenxi_id"] = schedule_fenxi_id
	pan_string_info["schedule_date"] = pan_html_string_info["schedule_date"]
	pan_string_info["schedule_no"] = pan_html_string_info["schedule_no"]
	pan_string_info["schedule_result_no"] = pan_html_string_info["schedule_result_no"]
	pan_string_info["schedule_league"] = pan_html_string_info["schedule_league"]

	pan_string_info["schedule_home"] = pan_html_string_info["schedule_home"]
	pan_string_info["schedule_guest"] = pan_html_string_info["schedule_guest"]
	pan_string_info["schedule_game_desc"] = pan_html_string_info["schedule_game_desc"]
	pan_string_info["schedule_date_desc"] = pan_html_string_info["schedule_date_desc"]

	for i := 0; i < odds_tr.Length(); i++ {
		tr_item := odds_tr.Eq(i)

		td_of_company := tr_item.Find("td").Eq(1)
		if td_of_company.Find("p a").Attr("title") == "" {
			continue
		}

		company_id := tr_item.Attr("id")
		pan_string_info["company_id"] = company_id
		pan_string_info["company_name"] = td_of_company.Find("p a").Attr("title")

		var is_big_company = 0
		if td_of_company.Find("p img").Attr("src") == "" {
			is_big_company = 0
		} else {
			is_big_company = 1
			//			fmt.Println("src:" + td_of_company.Find("p img").Attr("src"))
		}
		pan_int_info["is_big_company"] = is_big_company

		table_of_pan_detail := tr_item.Find("td .pl_table_data")

		table_of_opentime_pan := table_of_pan_detail.Eq(1)
		tds_of_opentime_pan_table := table_of_opentime_pan.Find("tbody tr td")
		open_pan_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["open_pan"] = float32(open_pan_32)
		pan_string_info["open_pan_desc"] = tds_of_opentime_pan_table.Eq(1).Text()

		open_home_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(0).Text(), 32)
		open_guest_water_32, _ := strconv.ParseFloat(tds_of_opentime_pan_table.Eq(2).Text(), 32)
		pan_float_info["open_home_water"] = float32(open_home_water_32)
		pan_float_info["open_guest_water"] = float32(open_guest_water_32)

		td_of_pan_time := tr_item.Find("td time")

		pan_string_info["open_pan_time"] = td_of_pan_time.Eq(1).Text()

		table_of_realtime_pan := table_of_pan_detail.Eq(0)
		tds_of_realtime_pan_table := table_of_realtime_pan.Find("tbody tr td")

		real_pan_32, _ := strconv.ParseFloat(tds_of_realtime_pan_table.Eq(1).Attr("ref"), 32)
		pan_float_info["real_pan"] = float32(real_pan_32)
		pan_string_info["real_pan_desc"] = tds_of_realtime_pan_table.Eq(1).Text()

		home_real_water_string := tds_of_realtime_pan_table.Eq(0).Text()
		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_string := tds_of_realtime_pan_table.Eq(2).Text()
		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		pan_float_info["real_home_water"] = float32(home_real_water_32)
		pan_float_info["real_guest_water"] = float32(guest_real_water_32)

		pan_string_info["pan_change_time"] = strings.TrimSpace(td_of_pan_time.Eq(0).Text())

		td_item_of_real_pan := tds_of_realtime_pan_table.Eq(1)
		home_pan_change_type := td_item_of_real_pan.Find("font").Text()
		home_pan_change_type = strings.TrimSpace(home_pan_change_type)
		fmt.Println("home_pan_desc:", home_pan_change_type)
		pan_int_info["home_pan_change_type"] = 0
		pan_string_info["home_pan_change_type_desc"] = ""

		if home_pan_change_type == "升" {
			pan_int_info["home_pan_change_type"] = 1
			pan_string_info["home_pan_change_type_desc"] = home_pan_change_type

		}
		if home_pan_change_type == "降" {
			pan_int_info["home_pan_change_type"] = -1
			pan_string_info["home_pan_change_type_desc"] = home_pan_change_type

		}
		fmt.Println("home_pan_change_type_desc:", pan_string_info["home_pan_change_type_desc"])

		home_water_up_down_flag := tds_of_realtime_pan_table.Eq(0).Attr("class")
		pan_int_info["home_water_change_type"] = 0
		pan_string_info["home_water_change_type_desc"] = ""
		if home_water_up_down_flag == "ping" {
			pan_int_info["home_water_change_type"] = -1            // down
			pan_string_info["home_water_change_type_desc"] = "水位降" // down
		}
		if home_water_up_down_flag == "ying" {
			pan_int_info["home_water_change_type"] = 1             // up
			pan_string_info["home_water_change_type_desc"] = "水位升" // up
		}

		exist_asiapanlog := new(myinit.AsiaPanLog)
		has, f_err := myinit.Engine.Where("schedule_fenxi_id=? AND company_id=? AND pan_change_time=?", schedule_fenxi_id, company_id, pan_string_info["pan_change_time"]).Get(exist_asiapanlog)
		fmt.Println("has:", has, ":", schedule_fenxi_id, ":", company_id, ":", pan_string_info["pan_change_time"])
		fmt.Println("ajax:", f_err)

		if has == true {
			fmt.Println("go has:is:", pan_string_info["pan_change_time"])
			continue
		}

		//		parse_change_data := ParsePanChangeUrl(schedule_fenxi_id, company_id, pan_int_info, pan_float_info, pan_string_info)
		//		if parse_change_data == false {
		//			continue
		//		}
		if pan_float_info["open_pan"] > 0 || pan_float_info["real_pan"] > 0 {
			asiapan.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			asiapanbackup.ClearShouPanByScheduleFenxiId(schedule_fenxi_id)
			schedule.ClearScheduleByFenxiId(schedule_fenxi_id)
			fmt.Println("pan ajx:", pan_float_info["open_pan"], pan_float_info["real_pan"])

			fmt.Println("开盘>0 或者即时盘 >0")
			return false
		}

		predict1_result, predict1_cmt := analyse.AnalysePanResult1(pan_int_info, pan_float_info, pan_string_info)
		pan_string_info["predict1_result"] = predict1_result
		pan_string_info["predict1_cmt"] = predict1_cmt
		predict2_result, predict2_cmt := analyse.AnalysePanResult2(pan_int_info, pan_float_info, pan_string_info)
		pan_string_info["predict2_result"] = predict2_result
		pan_string_info["predict2_cmt"] = predict2_cmt

		// fmt.Println("company:" + company)
		// fmt.Println("home_pan_change_type:" + home_pan_change_type)
		// fmt.Println("is big company:" + is_big_company)
		// fmt.Println("change_time:" + change_time)
		// fmt.Println("open_time:" + open_time)
		// fmt.Println("flag:" + home_water_change_type + " " + home_water_up_down_flag)
		//		fmt.Println("home_real_water:", pan_float_info["real_home_water"])
		//		fmt.Println("home_real_water water sting:" + home_real_water_string)
		// fmt.Println("guest_real_water:", guest_real_water)
		// fmt.Println("guest_real_water water sting:" + guest_real_water_string)
		// fmt.Println("pan:", real_pan, " ", real_pan_desc)

		// fmt.Println("open_home_water water:", open_home_water)
		// fmt.Println("open_guest_water water:", open_guest_water)
		// fmt.Println("open pan:", open_pan, " ", open_pan_desc)

		addAsiaPan(schedule_fenxi_id, company_id, pan_int_info, pan_float_info, pan_string_info)
		//delete water > 2 and < 1.75
		checkWaterSum(schedule_fenxi_id, company_id, pan_float_info)
	}
	if odds_tr.Length() >= 30 {
		ParsePanUrlFromAjax(idx+30, schedule_fenxi_id, pan_html_string_info)
	}

	return true
}
