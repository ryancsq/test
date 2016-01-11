package asiapan

import (
	"500panmap/util/myinit"
	"fmt"
)

func Add(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) {
	myinit.Myinit()
	AsiaPan := new(myinit.AsiaPan)
	AsiaPan.ScheduleFenxiId = pan_int_info["schedule_fenxi_id"]
//	AsiaPan.ScheduleBetDate = pan_string_info["schedule_bet_date"]

	AsiaPan.ScheduleDate = pan_string_info["schedule_date"]
	AsiaPan.ScheduleNo = pan_string_info["schedule_no"]
	AsiaPan.ScheduleResultNo = pan_string_info["schedule_result_no"]
	AsiaPan.ScheduleLeague = pan_string_info["schedule_league"]
	AsiaPan.ScheduleHome = pan_string_info["schedule_home"]
	AsiaPan.ScheduleGuest = pan_string_info["schedule_guest"]
	AsiaPan.ScheduleGameDesc = pan_string_info["schedule_game_desc"]
	AsiaPan.ScheduleDateDesc = pan_string_info["schedule_date_desc"]
	AsiaPan.CompanyId = pan_string_info["company_id"]
	AsiaPan.CompanyName = pan_string_info["company_name"]
	AsiaPan.IsBigCompany = pan_int_info["is_big_company"]
	AsiaPan.OpenPan = pan_float_info["open_pan"]
	AsiaPan.OpenPanDesc = pan_string_info["open_pan_desc"]
	AsiaPan.OpenHomeWater = pan_float_info["open_home_water"]
	AsiaPan.OpenGuestWater = pan_float_info["open_guest_water"]
	AsiaPan.OpenPanTime = pan_string_info["open_pan_time"]
	AsiaPan.RealPan = pan_float_info["real_pan"]
	AsiaPan.RealPanDesc = pan_string_info["real_pan_desc"]
	AsiaPan.RealHomeWater = pan_float_info["real_home_water"]
	AsiaPan.RealGuestWater = pan_float_info["real_guest_water"]
	AsiaPan.PanChangeTime = pan_string_info["pan_change_time"]
	AsiaPan.HomePanChangeType = pan_int_info["home_pan_change_type"]
	AsiaPan.HomePanChangeTypeDesc = pan_string_info["home_pan_change_type_desc"]
	AsiaPan.HomeWaterChangeType = pan_int_info["home_water_change_type"]
	AsiaPan.HomeWaterChangeTypeDesc = pan_string_info["home_water_change_type_desc"]
	AsiaPan.Predict1Result = pan_string_info["predict1_result"]
	AsiaPan.Predict1Comment = pan_string_info["predict1_cmt"]
	AsiaPan.Predict2Result = pan_string_info["predict2_result"]
	AsiaPan.Predict2Comment = pan_string_info["predict2_cmt"]

	ins_affected, ins_err := myinit.Engine.Insert(AsiaPan)
	fmt.Println(ins_affected)
	fmt.Println(ins_err)
}

func UpdateAsiaPanInfo(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) (update_affected int64, update_err error) {
	schedule_fenxi_id := pan_int_info["schedule_fenxi_id"]
	company_id := pan_string_info["company_id"]
	update_lastpan := new(myinit.AsiaPan)
	update_lastpan.RealPan = pan_float_info["real_pan"]
	update_lastpan.RealPanDesc = pan_string_info["real_pan_desc"]
	update_lastpan.RealHomeWater = pan_float_info["real_home_water"]
	update_lastpan.RealGuestWater = pan_float_info["real_guest_water"]
	update_lastpan.PanChangeTime = pan_string_info["pan_change_time"]
	update_lastpan.HomePanChangeType = pan_int_info["home_pan_change_type"]
	update_lastpan.HomePanChangeTypeDesc = pan_string_info["home_pan_change_type_desc"]
	update_lastpan.HomeWaterChangeType = pan_int_info["home_water_change_type"]
	update_lastpan.HomeWaterChangeTypeDesc = pan_string_info["home_water_change_type_desc"]
	update_lastpan.Predict1Result = pan_string_info["predict1_result"]
	update_lastpan.Predict1Comment = pan_string_info["predict1_cmt"]
	update_lastpan.Predict2Result = pan_string_info["predict2_result"]
	update_lastpan.Predict2Comment = pan_string_info["predict2_cmt"]
	update_affected, update_err =
		myinit.Engine.
			Cols("real_pan", "real_pan_desc", "real_home_water", "real_guest_water", "pan_change_time", "home_pan_change_type", "home_pan_change_type_desc", "home_water_change_type", "home_water_change_type_desc", "predict1_result", "predict1_comment", "predict2_result", "predict2_comment").
			Where("schedule_fenxi_id=? AND company_id=? ", schedule_fenxi_id, company_id).Update(update_lastpan)
	fmt.Println(update_affected)
	fmt.Println(update_err)
	return update_affected, update_err
}

func UpdateAsiaPanResult(schedule_bet_date string, pan_float_info map[string]float32, pan_string_info map[string]string) (update_affected int64, update_err error) {
	AsiaPan := new(myinit.AsiaPan)
	AsiaPan.ScheduleScore = pan_string_info["schedule_score"]
	AsiaPan.ScheduleSpfResult = pan_string_info["schedule_spf_result"]
	AsiaPan.ScheduleRqspfResult = pan_string_info["schedule_rqspf_result"]
	AsiaPan.ScheduleZjqResult = pan_string_info["schedule_zjq_result"]
	AsiaPan.ScheduleBqcResult = pan_string_info["schedule_bqc_result"]
	update_affected, update_err =
		myinit.Engine.
			Cols("schedule_score", "schedule_spf_result", "schedule_rqspf_result", "schedule_zjq_result", "schedule_bqc_result").
			Where("schedule_result_no=? AND schedule_date=? ", pan_string_info["schedule_result_no"], schedule_bet_date).Update(AsiaPan)
	fmt.Println(update_affected)
	fmt.Println(update_err)
	return update_affected, update_err
}
