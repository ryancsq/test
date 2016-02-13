package parseurl

import (
	"fmt"
	"time"
	"strings"

	"500kan/util/asiapan"
	"500kan/util/asiapanlog"
	"500kan/util/common"
	"500kan/util/myinit"
	"500kan/util/schedule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/opesun/goquery"
)

func ParseResultUrl(date string, history bool) {
	if date == "" {
		now := time.Now()
		date = now.Format("2006-01-02")
	}
	result_url := strings.Replace(myinit.ResultUrl, "DDD", date, -1)
	fmt.Println(result_url)
	schedule_string_info := make(map[string]string)
	schedule_float_info := make(map[string]float32)

	pan_string_info := make(map[string]string)
	pan_float_info := make(map[string]float32)
	html_obj, _ := goquery.ParseUrl(result_url)
	schedule_trs := html_obj.Find(".ld_table tbody tr")
	for i, _ := range schedule_trs {
		if i == 0 {
			continue
		}
		tr := schedule_trs.Eq(i)
		tds := tr.Find("td")
		schedule_string_info["schedule_result_no"] = common.ConvToGB(tds.Eq(0).Html())

		schedule_string_info["schedule_score"] = common.ConvToGB(tds.Eq(6).Html())
		schedule_string_info["schedule_spf_result"] = common.ConvToGB(tds.Eq(11).Html())
		schedule_float_info["schedule_spf_odd"] = common.ConvToFloat32(tds.Eq(12).Text())
		schedule_string_info["schedule_rqspf_result"] = common.ConvToGB(tds.Eq(8).Html())
		schedule_float_info["schedule_rqspf_odd"] = common.ConvToFloat32(tds.Eq(9).Text())
		schedule_string_info["schedule_zjq_result"] = common.ConvToGB(tds.Eq(14).Html())
		schedule_float_info["schedule_zjq_odd"] = common.ConvToFloat32(tds.Eq(15).Text())
		schedule_string_info["schedule_bqc_result"] = common.ConvToGB(tds.Eq(17).Html())
		schedule_float_info["schedule_bqc_odd"] = common.ConvToFloat32(tds.Eq(18).Text())

		pan_string_info["schedule_result_no"] = schedule_string_info["schedule_result_no"]
		pan_string_info["schedule_score"] = schedule_string_info["schedule_score"]
		pan_string_info["schedule_spf_result"] = schedule_string_info["schedule_spf_result"]
		pan_string_info["schedule_rqspf_result"] = schedule_string_info["schedule_rqspf_result"]
		pan_string_info["schedule_zjq_result"] = schedule_string_info["schedule_zjq_result"]
		pan_string_info["schedule_bqc_result"] = schedule_string_info["schedule_bqc_result"]
		has := schedule.CheckExistsByResultNoAndDate(schedule_string_info["schedule_result_no"], date)
		if has == false {
			continue
		}
		schedule.UpdateScheduleResult(date, schedule_float_info, schedule_string_info)
		asiapan.UpdateAsiaPanResult(date, schedule_float_info, pan_string_info)
		asiapanlog.UpdateAsiaPanResult(date, pan_float_info, pan_string_info)

	}
}
