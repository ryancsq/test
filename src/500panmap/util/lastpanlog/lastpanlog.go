package lastpanlog

import (
	"500panmap/util/myinit"
	"fmt"
)

func Add(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) {
	myinit.Myinit()
	LastPanLog := new(myinit.LastPanLog)
	LastPanLog.ScheduleNo = pan_string_info["schedule_no"]

	LastPanLog.ScheduleHome = pan_string_info["schedule_home_name"]
	LastPanLog.ScheduleGuest = pan_string_info["schedule_guest_name"]
	LastPanLog.ScheduleDate = pan_string_info["schedule_date"]
	LastPanLog.ScheduleGameDesc = pan_string_info["schedule_game_desc"]
	LastPanLog.CompanyCid = pan_string_info["cid"]
	LastPanLog.CompanyName = pan_string_info["company"]
	LastPanLog.ScheduleFid = pan_int_info["schedule_fenxi_id"]
	LastPanLog.OpenPan = pan_float_info["open_pan"]
	LastPanLog.OpenPanDesc = pan_string_info["open_pan_desc"]
	LastPanLog.OpenHomeWater = pan_float_info["open_home_water"]
	LastPanLog.OpenGuestWater = pan_float_info["open_guest_water"]
	LastPanLog.OpenPanTime = pan_string_info["open_time"]
	LastPanLog.LastPanLog = pan_float_info["real_pan"]
	LastPanLog.LastPanDesc = pan_string_info["real_pan_desc"]
	LastPanLog.LastHomeWater = pan_float_info["home_real_water"]
	LastPanLog.LastGuestWater = pan_float_info["guest_real_water"]
	LastPanLog.LastChangeTime = pan_string_info["change_time"]
	LastPanLog.LastHomePanChangeType = pan_string_info["home_pan_change_type"]
	LastPanLog.IsBigCompany = pan_int_info["is_big_company"]
	LastPanLog.LastHomeWaterChangeType = pan_int_info["home_water_change_type"]
	LastPanLog.PredictResult = pan_string_info["predict_result"]
	LastPanLog.PredictComment = pan_string_info["predict_cmt"]

	ins_affected, ins_err := myinit.Engine.Insert(LastPanLog)
	fmt.Println(ins_affected)
	fmt.Println(ins_err)
}
