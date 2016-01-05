package asiapanbackup

import (
	"500kan/util/myinit"
//	"fmt"
)

func Add(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) {
	myinit.Myinit()
	AsiaPanBackup := new(myinit.AsiaPanBackup)
	AsiaPanBackup.ScheduleFenxiId = pan_int_info["schedule_fenxi_id"]
	AsiaPanBackup.ScheduleDate = pan_string_info["schedule_date"]
	AsiaPanBackup.ScheduleNo = pan_string_info["schedule_no"]
	AsiaPanBackup.ScheduleResultNo = pan_string_info["schedule_result_no"]
	AsiaPanBackup.ScheduleLeague = pan_string_info["schedule_league"]
	AsiaPanBackup.ScheduleHome = pan_string_info["schedule_home"]
	AsiaPanBackup.ScheduleGuest = pan_string_info["schedule_guest"]
	AsiaPanBackup.ScheduleGameDesc = pan_string_info["schedule_game_desc"]
	AsiaPanBackup.ScheduleDateDesc = pan_string_info["schedule_date_desc"]
	AsiaPanBackup.CompanyId = pan_string_info["company_id"]
	AsiaPanBackup.CompanyName = pan_string_info["company_name"]
	AsiaPanBackup.IsBigCompany = pan_int_info["is_big_company"]
	AsiaPanBackup.OpenPan = pan_float_info["open_pan"]
	AsiaPanBackup.OpenPanDesc = pan_string_info["open_pan_desc"]
	AsiaPanBackup.OpenHomeWater = pan_float_info["open_home_water"]
	AsiaPanBackup.OpenGuestWater = pan_float_info["open_guest_water"]
	AsiaPanBackup.OpenPanTime = pan_string_info["open_pan_time"]
	AsiaPanBackup.RealPan = pan_float_info["real_pan"]
	AsiaPanBackup.RealPanDesc = pan_string_info["real_pan_desc"]
	AsiaPanBackup.RealHomeWater = pan_float_info["real_home_water"]
	AsiaPanBackup.RealGuestWater = pan_float_info["real_guest_water"]
	AsiaPanBackup.PanChangeTime = pan_string_info["pan_change_time"]
	AsiaPanBackup.HomePanChangeType = pan_int_info["home_pan_change_type"]
	AsiaPanBackup.HomePanChangeTypeDesc = pan_string_info["home_pan_change_type_desc"]
	AsiaPanBackup.HomeWaterChangeType = pan_int_info["home_water_change_type"]
	AsiaPanBackup.HomeWaterChangeTypeDesc = pan_string_info["home_water_change_type_desc"]
	AsiaPanBackup.Predict1Result = pan_string_info["predict1_result"]
	AsiaPanBackup.Predict1Comment = pan_string_info["predict1_cmt"]
	AsiaPanBackup.Predict2Result = pan_string_info["predict2_result"]
	AsiaPanBackup.Predict2Comment = pan_string_info["predict2_cmt"]

	myinit.Engine.Insert(AsiaPanBackup)
	
}

func UpdateAsiaPanBackupInfo(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) (update_affected int64, update_err error) {
	schedule_fenxi_id := pan_int_info["schedule_fenxi_id"]
	company_id := pan_string_info["company_id"]
	update_lastpan := new(myinit.AsiaPanBackup)
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
//	fmt.Println(update_affected)
//	fmt.Println(update_err)
	return update_affected, update_err
}


func DeletePanByFenxiIdAndCompanyId(schedule_fenxi_id int, company_id string) {
	delete_asiapanbackup := new(myinit.AsiaPanBackup)
	myinit.Engine.Where("schedule_fenxi_id=? AND company_id=? ", schedule_fenxi_id, company_id).Delete(delete_asiapanbackup)

}