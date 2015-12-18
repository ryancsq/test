package schedule

import (
	"500kan/util/myinit"
	"fmt"
)

func Add(schedule_int_info map[string] int,schedule_string_info map[string] string) {
	myinit.Myinit()
	has := CheckExists(schedule_string_info["schedule_date"], schedule_string_info["schedule_no"])
	fmt.Println(has)

	if has {
		fmt.Println(schedule_string_info["home_team"] + " vs " + schedule_string_info["guest_team"] + "已存在！")
	} else {
		Schedule := new(myinit.Schedule)
		Schedule.ScheduleFid = schedule_int_info["fid"]
		Schedule.ScheduleHome = schedule_string_info["home_team"]
		Schedule.ScheduleGuest = schedule_string_info["guest_team"]
		Schedule.ScheduleDate = schedule_string_info["schedule_date"]
		Schedule.ScheduleLeague = schedule_string_info["lg"]
		Schedule.ScheduleWeekDay = schedule_string_info["game_date_no"]
		Schedule.ScheduleNo = schedule_string_info["schedule_no"]
		Schedule.ScheduleEndTime = schedule_string_info["end_time"]

		affected, _ := myinit.Engine.Insert(Schedule)
		fmt.Println(affected)
		fmt.Println(Schedule.ScheduleId)
	}
}

func CheckExists(schedule_date string,schedule_no string)(has bool) {
	exist_schedule := new(myinit.Schedule)
	has, _ = myinit.Engine.Where("schedule_date=? AND schedule_no=? ", schedule_date, schedule_no).Get(exist_schedule)
	fmt.Println(has)
	return has
}
