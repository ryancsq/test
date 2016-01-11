package parseurl

import (
	"fmt"
	"strconv"
	"strings"

	"500panmap/util/asiapanlog"
	"500panmap/util/myinit"
	"500panmap/util/panmap"

	"github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func ParsePanChangeUrl(schedule_fenxi_id int, company_id string, pan_html_int_info map[string]int, pan_html_float_info map[string]float32, pan_html_string_info map[string]string) (res bool) {
	pan_int_info := make(map[string]int)
	pan_float_info := make(map[string]float32)
	pan_string_info := make(map[string]string)

	for k_i, v_i := range pan_html_int_info {
		pan_int_info[k_i] = v_i
	}
	for k_f, v_f := range pan_html_float_info {
		pan_float_info[k_f] = v_f
	}
	for k_s, v_s := range pan_html_string_info {
		pan_string_info[k_s] = v_s
	}

	pan_string_info["predict1_result"] = ""
	pan_string_info["predict1_cmt"] = ""
	pan_string_info["predict2_result"] = ""
	pan_string_info["predict2_cmt"] = ""

	asiapanlog.ClearOldPanLog(schedule_fenxi_id, company_id)
	body_content := myinit.GetOddsFromAjax(schedule_fenxi_id, company_id)
	body := []byte(body_content)
	body_json, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}
	tr_items, _ := body_json.Array()
	for _, tr_string := range tr_items {
		table_string := "<table>" + tr_string.(string) + "</table>"
		html_obj, _ := goquery.ParseString(table_string)
		pan_log_item := html_obj.Find("table tbody tr td")

		home_td := pan_log_item.Eq(0)
		pan_td := pan_log_item.Eq(1)
		guest_td := pan_log_item.Eq(2)
		time_td := pan_log_item.Eq(3)

		pan_string_info["real_pan_desc"] = pan_td.Text()

		home_real_water_string := home_td.Text()
		home_real_water_str := strings.Replace(home_real_water_string, "↑", "", -1)
		home_real_water_str = strings.Replace(home_real_water_str, "↓", "", -1)

		guest_real_water_string := guest_td.Text()
		guest_real_water_str := strings.Replace(guest_real_water_string, "↑", "", -1)
		guest_real_water_str = strings.Replace(guest_real_water_str, "↓", "", -1)

		home_real_water_32, _ := strconv.ParseFloat(home_real_water_str, 32)
		guest_real_water_32, _ := strconv.ParseFloat(guest_real_water_str, 32)

		pan_float_info["real_home_water"] = float32(home_real_water_32)
		pan_float_info["real_guest_water"] = float32(guest_real_water_32)

		pan_string_info["pan_change_time"] = time_td.Text()

		home_pan_change_type := pan_td.Find("font").Text()
		home_pan_change_type = strings.TrimSpace(home_pan_change_type)
		pan_int_info["home_pan_change_type"] = 0

		if home_pan_change_type == "升" {
			pan_int_info["home_pan_change_type"] = 1
			pan_string_info["home_pan_change_type_desc"] = home_pan_change_type

		}
		if home_pan_change_type == "降" {
			pan_int_info["home_pan_change_type"] = -1
			pan_string_info["home_pan_change_type_desc"] = home_pan_change_type

		}

		home_water_up_down_flag := home_td.Attr("class")
		pan_int_info["home_water_change_type"] = 0
		if home_water_up_down_flag == "tips_down" {
			pan_int_info["home_water_change_type"] = -1            // down
			pan_string_info["home_water_change_type_desc"] = "水位降" // down
		}
		if home_water_up_down_flag == "tips_up" {
			pan_int_info["home_water_change_type"] = 1             // up
			pan_string_info["home_water_change_type_desc"] = "水位升" // up
		}

		real_pan_string := strings.Replace(pan_string_info["real_pan_desc"], pan_string_info["home_pan_change_type_desc"], "", -1)
		real_pan_desc := strings.TrimSpace(real_pan_string)

		has_panmap, real_pan_value := panmap.GetPanValueByPanDesc(real_pan_desc)
		if has_panmap == false {
			fmt.Println(pan_string_info["schedule_home"])
			fmt.Println(pan_string_info["company_id"])
			fmt.Println(pan_string_info["company_name"])
			fmt.Println(real_pan_desc, "no exist")
			return false
		}
		pan_float_info["real_pan"] = real_pan_value
		asiapanlog.Add(pan_int_info, pan_float_info, pan_string_info)
	}

	return true
}
