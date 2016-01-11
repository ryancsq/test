package schedule

import (
	"500panmap/util/myinit"
	"fmt"
)

func Add(schedule_int_info map[string]int, schedule_string_info map[string]string) {
	myinit.Myinit()
	has := CheckExists(schedule_string_info["schedule_date"], schedule_string_info["schedule_no"])
	fmt.Println(has)

	if has {
		fmt.Println(schedule_string_info["schedule_home"] + " vs " + schedule_string_info["schedule_guest"] + "已存在！")
	} else {
		Schedule := new(myinit.Schedule)

//		Schedule.ScheduleBetDate = schedule_string_info["schedule_bet_date"]
		Schedule.ScheduleDate = schedule_string_info["schedule_date"]
		Schedule.ScheduleNo = schedule_string_info["schedule_no"]
		Schedule.ScheduleResultNo = schedule_string_info["schedule_result_no"]
		Schedule.ScheduleLeague = schedule_string_info["schedule_league"]
		Schedule.ScheduleHome = schedule_string_info["schedule_home"]
		Schedule.ScheduleGuest = schedule_string_info["schedule_guest"]
		Schedule.ScheduleWeekDay = schedule_string_info["schedule_week_day"]
		Schedule.ScheduleFenxiId = schedule_int_info["schedule_fenxi_id"]
		Schedule.ScheduleBetEndTime = schedule_string_info["schedule_bet_end_time"]
		Schedule.ScheduleRqNum = schedule_string_info["schedule_rq_num"]

		affected, _ := myinit.Engine.Insert(Schedule)
		fmt.Println(affected)
		fmt.Println(Schedule.ScheduleId)
	}
}

func CheckExists(schedule_date string, schedule_no string) (has bool) {
	exist_schedule := new(myinit.Schedule)
	has, _ = myinit.Engine.Where("schedule_date=? AND schedule_no=? ", schedule_date, schedule_no).Get(exist_schedule)
	fmt.Println(has)
	return has
}

func CheckExistsByResultNoAndDate(schedule_result_no string, schedule_bet_date string) (has bool) {
	exist_schedule := new(myinit.Schedule)
	has, _ = myinit.Engine.Where("schedule_date=? AND schedule_result_no=? ", schedule_bet_date, schedule_result_no).Get(exist_schedule)
	fmt.Println("no:")
	fmt.Println(schedule_result_no)
	fmt.Println(schedule_bet_date)
	fmt.Println(has)
	fmt.Println("====")

	return has
}

func UpdateScheduleResult(schedule_bet_date string,schedule_float_info map[string]float32, schedule_string_info map[string]string) {
	Schedule := new(myinit.Schedule)
	Schedule.ScheduleScore = schedule_string_info["schedule_score"]
	Schedule.ScheduleSpfResult = schedule_string_info["schedule_spf_result"]
	Schedule.ScheduleSpfOdd = schedule_float_info["schedule_spf_odd"]
	Schedule.ScheduleRqspfResult = schedule_string_info["schedule_rqspf_result"]
	Schedule.ScheduleRqspfOdd = schedule_float_info["schedule_rqspf_odd"]
	Schedule.ScheduleZjqResult = schedule_string_info["schedule_zjq_result"]
	Schedule.ScheduleZjqOdd = schedule_float_info["schedule_zjq_odd"]
	Schedule.ScheduleBqcResult = schedule_string_info["schedule_bqc_result"]
	Schedule.ScheduleBqcOdd = schedule_float_info["schedule_bqc_odd"]
	update_affected, update_err := 
		myinit.Engine.
		Cols("schedule_score", "schedule_spf_result", "schedule_spf_odd", "schedule_rqspf_result", "schedule_rqspf_odd", "schedule_zjq_result", "schedule_zjq_odd", "schedule_bqc_result", "schedule_bqc_odd").
		Where("schedule_result_no=? AND schedule_date=? ", schedule_string_info["schedule_result_no"], schedule_bet_date).Update(Schedule)
	fmt.Println(update_affected)
	fmt.Println(update_err)
}
