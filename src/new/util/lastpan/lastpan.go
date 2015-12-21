package lastpan

import (
	"500kan/util/myinit"
	"fmt"
)

func Add(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) {
	myinit.Myinit()
	LastPan := new(myinit.LastPan)
	LastPan.ScheduleNo = pan_string_info["schedule_no"]

	LastPan.ScheduleHome = pan_string_info["schedule_home_name"]
	LastPan.ScheduleGuest = pan_string_info["schedule_guest_name"]
	LastPan.ScheduleDate = pan_string_info["schedule_date"]
	LastPan.ScheduleGameDesc = pan_string_info["schedule_game_desc"]
	LastPan.CompanyCid = pan_string_info["cid"]
	LastPan.CompanyName = pan_string_info["company"]
	LastPan.ScheduleFid = pan_int_info["schedule_fenxi_id"]
	LastPan.OpenPan = pan_float_info["open_pan"]
	LastPan.OpenPanDesc = pan_string_info["open_pan_desc"]
	LastPan.OpenHomeWater = pan_float_info["open_home_water"]
	LastPan.OpenGuestWater = pan_float_info["open_guest_water"]
	LastPan.OpenPanTime = pan_string_info["open_time"]
	LastPan.LastPan = pan_float_info["real_pan"]
	LastPan.LastPanDesc = pan_string_info["real_pan_desc"]
	LastPan.LastHomeWater = pan_float_info["home_real_water"]
	LastPan.LastGuestWater = pan_float_info["guest_real_water"]
	LastPan.LastChangeTime = pan_string_info["change_time"]
	LastPan.LastHomePanChangeType = pan_string_info["home_pan_change_type"]
	LastPan.IsBigCompany = pan_int_info["is_big_company"]
	LastPan.LastHomeWaterChangeType = pan_int_info["home_water_change_type"]
	LastPan.PredictResult = pan_string_info["predict_result"]
	LastPan.PredictComment = pan_string_info["predict_cmt"]

	ins_affected, ins_err := myinit.Engine.Insert(LastPan)
	fmt.Println(ins_affected)
	fmt.Println(ins_err)
}

func updateLastPanInfo(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) (update_affected int64, update_err error) {
	schedulefid := pan_int_info["schedule_fenxi_id"]
	cid := pan_string_info["cid"]
	update_lastpan := new(myinit.LastPan)
	update_lastpan.LastPan = pan_float_info["real_pan"]
	update_lastpan.LastPanDesc = pan_string_info["real_pan_desc"]
	update_lastpan.LastHomeWater = pan_float_info["home_real_water"]
	update_lastpan.LastGuestWater = pan_float_info["guest_real_water"]
	update_lastpan.LastHomeWaterChangeType = pan_int_info["home_water_change_type"]
	update_lastpan.LastHomePanChangeType = pan_string_info["home_pan_change_type"]
	update_lastpan.LastChangeTime = pan_string_info["change_time"]
	update_lastpan.PredictResult = pan_string_info["predict_result"]
	update_lastpan.PredictComment = pan_string_info["predict_cmt"]
	update_affected, update_err = myinit.Engine.Cols("last_pan", "last_pan_desc", "last_home_water", "last_guest_water", "last_change_time", "last_home_pan_change_type", "last_home_water_change_type", "predict_result", "predict_comment").Where("schedule_fid=? AND company_cid=? ", schedulefid, cid).Update(update_lastpan)
	fmt.Println(update_affected)
	fmt.Println(update_err)
	return update_affected, update_err
}
