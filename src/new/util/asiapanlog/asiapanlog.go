package asiapanlog

import (
	"500kan/util/myinit"
	"fmt"
)

func Add(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) {
	myinit.Myinit()
	AsiaPanLog := new(myinit.AsiaPanLog)
	AsiaPanLog.ScheduleFenxiId = pan_int_info["schedule_fenxi_id"]
//	AsiaPanLog.ScheduleBetDate = pan_string_info["schedule_bet_date"]
	AsiaPanLog.ScheduleDate = pan_string_info["schedule_date"]
	AsiaPanLog.ScheduleNo = pan_string_info["schedule_no"]
	AsiaPanLog.ScheduleResultNo = pan_string_info["schedule_result_no"]
	AsiaPanLog.ScheduleLeague = pan_string_info["schedule_league"]
	AsiaPanLog.ScheduleHome = pan_string_info["schedule_home"]
	AsiaPanLog.ScheduleGuest = pan_string_info["schedule_guest"]
	AsiaPanLog.ScheduleGameDesc = pan_string_info["schedule_game_desc"]
	AsiaPanLog.ScheduleDateDesc = pan_string_info["schedule_date_desc"]
	AsiaPanLog.CompanyId = pan_string_info["company_id"]
	AsiaPanLog.CompanyName = pan_string_info["company_name"]
	AsiaPanLog.IsBigCompany = pan_int_info["is_big_company"]
	AsiaPanLog.OpenPan = pan_float_info["open_pan"]
	AsiaPanLog.OpenPanDesc = pan_string_info["open_pan_desc"]
	AsiaPanLog.OpenHomeWater = pan_float_info["open_home_water"]
	AsiaPanLog.OpenGuestWater = pan_float_info["open_guest_water"]
	AsiaPanLog.OpenPanTime = pan_string_info["open_pan_time"]
	AsiaPanLog.RealPan = pan_float_info["real_pan"]
	AsiaPanLog.RealPanDesc = pan_string_info["real_pan_desc"]
	AsiaPanLog.RealHomeWater = pan_float_info["real_home_water"]
	AsiaPanLog.RealGuestWater = pan_float_info["real_guest_water"]
	AsiaPanLog.PanChangeTime = pan_string_info["pan_change_time"]
	AsiaPanLog.HomePanChangeType = pan_int_info["home_pan_change_type"]
	AsiaPanLog.HomePanChangeTypeDesc = pan_string_info["home_pan_change_type_desc"]
	AsiaPanLog.HomeWaterChangeType = pan_int_info["home_water_change_type"]
	AsiaPanLog.HomeWaterChangeTypeDesc = pan_string_info["home_water_change_type_desc"]
	AsiaPanLog.Predict1Result = pan_string_info["predict1_result"]
	AsiaPanLog.Predict1Comment = pan_string_info["predict1_cmt"]
	AsiaPanLog.Predict2Result = pan_string_info["predict2_result"]
	AsiaPanLog.Predict2Comment = pan_string_info["predict2_cmt"]

	ins_affected, ins_err := myinit.Engine.Insert(AsiaPanLog)
	fmt.Println(ins_affected)
	fmt.Println(ins_err)
}

func UpdateAsiaPanResult(schedule_bet_date string, pan_float_info map[string]float32, pan_string_info map[string]string) (update_affected int64, update_err error) {
	AsiaPanLog := new(myinit.AsiaPanLog)
	AsiaPanLog.ScheduleScore = pan_string_info["schedule_score"]
	AsiaPanLog.ScheduleSpfResult = pan_string_info["schedule_spf_result"]
	AsiaPanLog.ScheduleRqspfResult = pan_string_info["schedule_rqspf_result"]
	AsiaPanLog.ScheduleZjqResult = pan_string_info["schedule_zjq_result"]
	AsiaPanLog.ScheduleBqcResult = pan_string_info["schedule_bqc_result"]
	update_affected, update_err = 
		myinit.Engine.
		Cols("schedule_score", "schedule_spf_result",  "schedule_rqspf_result", "schedule_zjq_result", "schedule_bqc_result").
		Where("schedule_result_no=? AND schedule_date=? ", pan_string_info["schedule_result_no"], schedule_bet_date).Update(AsiaPanLog)
	fmt.Println(update_affected)
	fmt.Println(update_err)
	return update_affected, update_err
}

