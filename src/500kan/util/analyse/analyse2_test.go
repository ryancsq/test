package analyse

import (
	"testing"
	"fmt"


	_ "github.com/go-sql-driver/mysql"
)

func TestAnalysePanResult2(t *testing.T) {
	pan_float_info := make(map[string] float32)
	pan_int_info := make(map[string] int)
	pan_string_info := make(map[string] string)
	

	
			pan_int_info["schedule_fenxi_id"] = 518963

	pan_float_info["open_pan"] = -0.25
	pan_float_info["open_home_water"] = 0.71
	pan_float_info["open_guest_water"] = 1.18
	pan_float_info["real_pan"] = -0.75
	pan_float_info["real_home_water"] = 1.2
	pan_float_info["real_guest_water"] = 0.7

	pan_int_info["home_pan_change_type"] = 1
	pan_string_info["home_pan_change_desc"] = "jiang"
	pan_string_info["company_id"] = "122" //香港马会
	
	pan_int_info["schedule_fenxi_id"] = 518963

	pan_float_info["open_pan"] = -0.25
	pan_float_info["open_home_water"] = 0.65
	pan_float_info["open_guest_water"] = 1.2
	pan_float_info["real_pan"] = -0.75
	pan_float_info["real_home_water"] = 1.1
	pan_float_info["real_guest_water"] = 0.7

	pan_int_info["home_pan_change_type"] = 1
	pan_string_info["home_pan_change_desc"] = "jiang"
	pan_string_info["company_id"] = "132" //inter
	
			pan_int_info["schedule_fenxi_id"] = 518963

	pan_float_info["open_pan"] = -0.5
	pan_float_info["open_home_water"] = 1.0
	pan_float_info["open_guest_water"] = 0.85
	pan_float_info["real_pan"] = -0.5
	pan_float_info["real_home_water"] = 0.88
	pan_float_info["real_guest_water"] = 1.25

	pan_int_info["home_pan_change_type"] = 0
	pan_string_info["home_pan_change_desc"] = "jiang"
	pan_string_info["company_id"] = "16" //10bet
	
		
	
	pan_int_info["schedule_fenxi_id"] = 524996

	pan_float_info["open_pan"] = 0.00
	pan_float_info["open_home_water"] = 0.6410
	pan_float_info["open_guest_water"] = 1.31
	pan_float_info["real_pan"] = 0.00
	pan_float_info["real_home_water"] = 0.69
	pan_float_info["real_guest_water"] = 1.25

	pan_int_info["home_pan_change_type"] = 0
	pan_string_info["home_pan_change_desc"] = "jiang"
	pan_string_info["company_id"] = "704" //Bookmaker.ag
	

	predict_result, predict_cmt := AnalysePanResult2(pan_int_info, pan_float_info,pan_string_info)
	fmt.Println(predict_result)
	fmt.Println(predict_cmt)
}

