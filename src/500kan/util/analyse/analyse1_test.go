package analyse

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAnalysePanResult1(t *testing.T) {
	pan_float_info := make(map[string]float32)
	pan_int_info := make(map[string]int)
	pan_string_info := make(map[string]string)

	pan_int_info["schedule_fenxi_id"] = 522154

	pan_float_info["open_pan"] = -0.5
	pan_float_info["open_home_water"] = 0.78
	pan_float_info["open_guest_water"] = 1.1
	pan_float_info["real_pan"] = -0.75
	pan_float_info["home_real_water"] = 1.0
	pan_float_info["guest_real_water"] = 0.85

	pan_int_info["home_water_change_type"] = 1
	pan_string_info["home_water_change_desc"] = "升"
	pan_string_info["company_id"] = "3"

	pan_float_info["open_pan"] = -0.75
	pan_float_info["open_home_water"] = 1.04
	pan_float_info["open_guest_water"] = 0.78
	pan_float_info["real_pan"] = -0.75
	pan_float_info["home_real_water"] = 1.02
	pan_float_info["guest_real_water"] = 0.8

	pan_int_info["home_water_change_type"] = 0
	pan_string_info["home_water_change_desc"] = ""
	pan_string_info["company_id"] = "5"

	pan_float_info["open_pan"] = -1
	pan_float_info["open_home_water"] = 1.25
	pan_float_info["open_guest_water"] = 0.57
	pan_float_info["real_pan"] = -1
	pan_float_info["home_real_water"] = 1.35
	pan_float_info["guest_real_water"] = 0.53

	pan_int_info["home_water_change_type"] = 0
	pan_string_info["home_water_change_desc"] = ""
	pan_string_info["company_id"] = "108"

	pan_int_info["schedule_fenxi_id"] = 534040

	pan_float_info["open_pan"] = -0.5
	pan_float_info["open_home_water"] = 0.909
	pan_float_info["open_guest_water"] = 0.909
	pan_float_info["real_pan"] = -0.5
	pan_float_info["home_real_water"] = 0.909
	pan_float_info["guest_real_water"] = 0.935

	pan_int_info["home_water_change_type"] = 0
	pan_string_info["home_water_change_desc"] = ""
	pan_string_info["company_id"] = "51"

	pan_int_info["schedule_fenxi_id"] = 557720

	pan_float_info["open_pan"] = -0.75
	pan_float_info["open_home_water"] = 0.85
	pan_float_info["open_guest_water"] = 0.952
	pan_float_info["real_pan"] = -0.5
	pan_float_info["home_real_water"] = 0.9
	pan_float_info["guest_real_water"] = 0.91

	pan_int_info["home_water_change_type"] = -1
	pan_string_info["home_water_change_desc"] = "jiang"
	pan_string_info["company_id"] = "6"

	pan_int_info["schedule_fenxi_id"] = 534032

	pan_float_info["open_pan"] = -1.5
	pan_float_info["open_home_water"] = 1.06
	pan_float_info["open_guest_water"] = 0.86
	pan_float_info["real_pan"] = -1.5
	pan_float_info["home_real_water"] = 1.03
	pan_float_info["guest_real_water"] = 0.89

	pan_int_info["home_water_change_type"] = 0
	pan_string_info["home_water_change_desc"] = "jiang"
	pan_string_info["company_id"] = "6"

	pan_int_info["schedule_fenxi_id"] = 545803

	pan_float_info["open_pan"] = -1.25
	pan_float_info["open_home_water"] = 1.15
	pan_float_info["open_guest_water"] = 0.72
	pan_float_info["real_pan"] = -1.5
	pan_float_info["home_real_water"] = 0.85
	pan_float_info["guest_real_water"] = 1.05

	pan_int_info["home_water_change_type"] = 1
	pan_string_info["home_water_change_desc"] = "jiang"
	pan_string_info["company_id"] = "3" //bet365

	pan_int_info["schedule_fenxi_id"] = 555165

	pan_float_info["open_pan"] = -0.25
	pan_float_info["open_home_water"] = 1.05
	pan_float_info["open_guest_water"] = 0.833
	pan_float_info["real_pan"] = 0
	pan_float_info["home_real_water"] = 0.8
	pan_float_info["guest_real_water"] = 1.1

	pan_int_info["home_water_change_type"] = -1
	pan_string_info["home_water_change_desc"] = "jiang"
	pan_string_info["company_id"] = "6" //weide
	
		pan_int_info["schedule_fenxi_id"] = 518963

	pan_float_info["open_pan"] = -0.25
	pan_float_info["open_home_water"] = 0.71
	pan_float_info["open_guest_water"] = 1.18
	pan_float_info["real_pan"] = -0.75
	pan_float_info["home_real_water"] = 1.2
	pan_float_info["guest_real_water"] = 0.7

	pan_int_info["home_water_change_type"] = 1
	pan_string_info["home_water_change_desc"] = "jiang"
	pan_string_info["company_id"] = "122" //weide
	
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

	predict_result, predict_cmt := AnalysePanResult1(pan_int_info, pan_float_info, pan_string_info)
	fmt.Println(predict_result)
	fmt.Println(predict_cmt)
}
