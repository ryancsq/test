package analyse

import (
//	"fmt"

	"500kan/util/myinit"

	_ "github.com/go-sql-driver/mysql"
)

func AnalysePanResult1(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) (ret string, cmt string) {
	open_pan := pan_float_info["open_pan"]
	open_pan_home_water := pan_float_info["open_home_water"]
	open_pan_guest_water := pan_float_info["open_guest_water"]
	real_pan := pan_float_info["real_pan"]
	real_pan_home_water := pan_float_info["home_real_water"]
	real_pan_guest_water := pan_float_info["guest_real_water"]

	home_pan_change_type := pan_int_info["home_water_change_type"]
	schedule_game_desc := pan_string_info["schedule_game_desc"]

	schedule_fenxi_id := pan_int_info["schedule_fenxi_id"]
	company_id := pan_string_info["company_id"]
	myinit.Myinit()
	engine = myinit.Engine
//	fmt.Println("+++++++++++")
//	fmt.Println(schedule_fenxi_id)
	switch {
	case open_pan == 0:
//		fmt.Println("0 open:", open_pan)

		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜"
				if real_pan_home_water < open_pan_home_water {
					ret = "1/0"
					cmt = "主队水位	＞0.875	主队即时盘口水位小于初盘水位，多为平局"
				}
			}

			if open_pan_home_water == open_pan_guest_water {
				ret = "1"
				cmt = "主队水位=	客队水位	平"
				if real_pan_home_water < 0.875 {
					ret = "3"
					cmt = "主队水位=	客队水位	平，即时水位＜0.875队胜出"

				}
			}

			if checkPanAndWaterNotChange(schedule_fenxi_id,company_id,open_pan)==true && real_pan_home_water < 0.875 {
				ret = "3"
				cmt = "盘口、水位一直不变		即时水位＜0.875队胜出"
			}
		} else if home_pan_change_type == 1 {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜  多为平局"
			}
		} else if home_pan_change_type == -1 {
			if real_pan_guest_water > 0.875 {
				ret = "3/1"
				cmt = "客队水位	＞0.875	主队胜或平"
			} else {
				ret = "1/0"
				cmt = "客队水位	≤0.875	平或客队胜"
			}
		}

//		fmt.Println("open:", open_pan, ret, cmt)
	case open_pan == (-0.25):
//		fmt.Println("-0.25 open:", open_pan)
		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜"
			}

			if open_pan_home_water == open_pan_guest_water && real_pan_home_water > 0.875 {
				ret = "3"
				cmt = "主队水位=	客队水位	即时水位＞0.875队胜出	"
			}
			if checkPanAndWaterNotChange(schedule_fenxi_id,company_id,open_pan)==true  {
				ret = "3/0"
				cmt = "盘口、水位一直不变		双方能分胜负		德甲主队胜概率大"
			}
			if checkPanNotChange(schedule_fenxi_id,company_id,open_pan)==true && checkIsGermanyJia(schedule_game_desc) == true {
				ret = "1/0"
				cmt = "若为德甲，盘口不变而水位发生变化们一般是下盘胜出			对应结果：	1/0	"
			}
		} else if home_pan_change_type == 1 {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 && checkWaterIsDown(schedule_fenxi_id, company_id) {
					ret = "3"
					cmt = "主队水位	≤0.875	即时水位＞0.875并且水位持续下降	主队胜"
				} else if real_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "主队水位	≤0.875	即时水位≤0.875	平或客队胜"
				}
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875		平或客队胜 多为平局"
			}
		} else if home_pan_change_type == -1 {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water <= 0.875 {
					ret = "0"
					cmt = "主队水位	≤0.875	即时水位≤0.875	客队胜"
				} else {
					ret = "1"
					cmt = "主队水位	≤0.875	即时水位＞0.875	平"
				}
//			} else {
//				ret = "1/0"
//				cmt = "其余情况			平或客队胜	对应结果：	1/0"
			}

		}
	case open_pan == (-0.5):
		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜	对应结果：	3"
			} else {
				ret = "1/0"
				cmt = "主队水位	＞0.875	平或客队胜	对应结果：	1/0"
			}

			if checkPanAndAllWaterNotChange(schedule_fenxi_id,company_id,open_pan)==true {
				if open_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "盘口、水位一直不变		初盘水位	主队水位	≤0.875	平或客队胜	对应结果：	1/0"
				} else {
					ret = "3"
					cmt = "盘口、水位一直不变		初盘水位	主队水位	＞0.875	主队胜	对应结果：	3"
				}
			}

		} else if home_pan_change_type == 1 {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 {
					ret = "3"
					cmt = "主队水位	≤0.875	即时水位＞0.875	主队胜	对应结果：	3"
				} else {
					ret = "1"
					cmt = "主队水位	≤0.875	即时水位≤0.875	平	对应结果：	1"
				}
			} else {
				ret = "0"
				cmt = "主队水位	＞0.875		客队胜	对应结果：	0"
			}
		} else if home_pan_change_type == -1 {
//			fmt.Println("-0.5====")
//			fmt.Println(open_pan_home_water)
//			fmt.Println(real_pan_home_water)
			if open_pan_home_water > 0.875 {
				if real_pan_home_water <= 0.875 {
					ret = "3/1"
					cmt = "主队水位	＞0.875	即时水位≤0.875	主胜或平	对应结果：	3、1"
				} else {
					ret = "0"
					cmt = "主队水位	＞0.875	即时水位＞0.875	客队胜	对应结果：	0"
				}
//			} else {
//				ret = "1/0"
//				cmt = "其余情况			平或客队胜	对应结果：	1/0"
			}
		}
//		fmt.Println("-0.5 open:", open_pan)

	case open_pan == (-0.75):
		if open_pan == real_pan {
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位	≤0.875	主队胜	对应结果：	3"
			} else {
				if real_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "主队水位	＞0.875	即时水位≤0.875	平或客队胜	对应结果：	1/0"
				} else {
					ret = "3"
					cmt = "主队水位	＞0.875	即时水位＞0.875	主队胜	对应结果：	3"
				}
				if real_pan_home_water == open_pan_home_water {
					ret = "0"
					cmt = "主队水位	＞0.875	即时水位=初盘水位	客队胜	对应结果：	0"
				}
			}

		} else if home_pan_change_type == 1 {
		} else if home_pan_change_type == -1 {
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 {
					ret = "1"
					cmt = "主队水位	≤0.875	即时水位＞0.875	平	对应结果：	1"
				} else {
					ret = "0"
					cmt = "主队水位	≤0.875	即时水位≤0.875	客队胜	对应结果：	0"
				}
			}

		}
//		fmt.Println("-0.75: open:", open_pan)

	case open_pan <= -1:
		//	case -1.25:
		//	case -1.5:
		flag := false
		if(checkPanNotLower(schedule_fenxi_id,company_id,open_pan)==true && checkWaterNotChange(schedule_fenxi_id,company_id)==false ){
			if open_pan_home_water > 0.875 {
				ret = "3"
				cmt = "主队水位	＞0.875	主队胜	对应结果：	3"	
				flag = true		
			
			}
		} 
		if(checkPanNotChange(schedule_fenxi_id,company_id,open_pan)==false && open_pan==(-1) && real_pan!=(-1)){
			if open_pan_home_water > 0.875 && real_pan_home_water <= 0.875 {
				ret = "3/1"
				cmt = "主队水位	＞0.875	即时水位≤0.875"
				flag = true		
			}
		}
		
		if(open_pan<(-1.5) && checkPanNotChange(schedule_fenxi_id,company_id,open_pan)==true){
			if(open_pan_home_water< 0.8){
				ret = "3/0"
				cmt = "初盘盘口数值＜-1.5） 初盘水位	主队水位	＜0.8	有爆冷可能	胜或负 "
				flag = true		
			}
		}
		
		if(flag==false){
				if open_pan_home_water < 0.875 {
				//其他情况
				ret = "3"
				cmt = "其余情况 初盘水位	主队水位	＜0.875	主队胜	对应结果：	3"	
				}		
		}
//		fmt.Println("-1 open:", open_pan)
	default:
//		fmt.Println("qita open:", open_pan)
		ret = ""
		cmt = ""

	}
	return ret, cmt
}

