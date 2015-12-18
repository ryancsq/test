package lastpanlog

import (
	"500kan/util/myinit"
	"fmt"
)

func Add(pan_int_info map[string]int, pan_string_info map[string]string) {
	myinit.Myinit()
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
