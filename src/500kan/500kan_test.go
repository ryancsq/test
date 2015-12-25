package main
import (
"fmt"
"testing"
)

func TestAnalysePanResult(t *testing.T) {
		initDb()

	open_pan := float32(-1.0)
	open_home_water := float32(0.97)
	open_guest_water := float32(0.8)
	real_pan:= float32(-1.0)
	home_real_water := float32(0.93)
	guest_real_water:= float32(0.84)
	home_pan_change_type:= " "
	schedule_game_desc:= "15世俱杯第五名"
	schedulefid:= "554737"
	cid:= "291"
	predict_result, predict_cmt := AnalysePanResult(open_pan, open_home_water, open_guest_water, real_pan, home_real_water, guest_real_water, home_pan_change_type, schedule_game_desc, schedulefid, cid)
//	if predict_result != "3" {
//		t.Error("analysePanResult failed. Got", predict_result, "Expected 3")
//	}
	fmt.Println(predict_cmt)
	fmt.Println(predict_result)
}
