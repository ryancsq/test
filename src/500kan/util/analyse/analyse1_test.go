package analyse

import (
	"testing"
	"fmt"


	_ "github.com/go-sql-driver/mysql"
)

func TestAnalysePanResult1(t *testing.T) {
	pan_float_info := make(map[string] float32)
	pan_int_info := make(map[string] int)
	pan_string_info := make(map[string] string)
	pan_float_info["open_pan"] = -0.5
	pan_float_info["open_home_water"] = 0.909
	pan_float_info["open_guest_water"] = 0.909
	pan_float_info["real_pan"] = -0.5
	pan_float_info["home_real_water"] = 0.909
	pan_float_info["guest_real_water"] = 0.935

	pan_int_info["home_water_change_type"] = 0
	pan_string_info["home_water_change_desc"] = "sheng"

	predict_result, predict_cmt := AnalysePanResult1(pan_int_info, pan_float_info,pan_string_info)
	fmt.Println(predict_result)
	fmt.Println(predict_cmt)
}

