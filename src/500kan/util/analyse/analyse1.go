package analyse

import (
	"fmt"

	"500kan/util/myinit"

	_ "github.com/go-sql-driver/mysql"
)

func AnalysePanResult1(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) (ret string, cmt string) {
	open_pan := pan_float_info["open_pan"]
	open_pan_home_water := pan_float_info["open_home_water"]
	open_pan_guest_water := pan_float_info["open_guest_water"]
	real_pan := pan_float_info["real_pan"]
	real_pan_home_water := pan_float_info["real_home_water"]
	real_pan_guest_water := pan_float_info["real_guest_water"]

	home_pan_change_type := pan_int_info["home_pan_change_type"]
	schedule_game_desc := pan_string_info["schedule_game_desc"]

	fid := pan_int_info["schedule_fenxi_id"]
	cid := pan_string_info["company_id"]
	myinit.Myinit()
	engine = myinit.Engine
	//	fmt.Println("+++++++++++")
	//	fmt.Println(schedule_fenxi_id)
	pre_cmt := ""
	switch {
	case open_pan == 0:
		if checkPanNotChange(fid, cid, open_pan) == true {
			pre_cmt = "一直保持平手盘（盘口数值为0）不变"
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位≤0.875,主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位＞0.875,平或客队胜"
				if real_pan_home_water < open_pan_home_water {
					ret = "1/0"
					cmt = "主队水位＞0.875,主队即时盘口水位小于初盘水位，<b>多为平局</b>"
				}
			}

			if open_pan_home_water == open_pan_guest_water {
				ret = "1"
				cmt = "主队水位=客队水位,平"
				if real_pan_home_water < 0.875 {
					ret = "3"
					cmt = "主队水位=客队水位,平，即时水位＜0.875队胜出"
				}
			}

			if checkWaterNotChange(fid, cid) == true && (real_pan_home_water < 0.875) {
				ret = "3"
				cmt = "盘口、水位一直不变,即时水位＜0.875队胜出"
			}
		} else if home_pan_change_type == 1 {
			pre_cmt = "相对主队出现升盘（平手升平半、平手升半球）初盘盘口=0，即时盘口数值＜0"
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位≤0.875,主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位＞0.875,平或客队胜,<b>多为平局</b>"
			}
		} else if home_pan_change_type == -1 {
			pre_cmt = "相对主队出现降盘（初盘盘口=0，即时盘口数值＞0）"

			if real_pan_guest_water > 0.875 {
				ret = "3/1"
				cmt = "客队水位＞0.875,主队胜或平"
			} else {
				ret = "1/0"
				cmt = "客队水位≤0.875,平或客队胜"
			}
		}
	case open_pan == (-0.25):
		if checkPanNotChange(fid, cid, open_pan) == true {
			pre_cmt = "一直保持平半（盘口数值为-0.25）不变"
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位≤0.875,主队胜"
			} else {
				ret = "1/0"
				cmt = "主队水位＞0.875,平或客队胜"
			}

			if open_pan_home_water == open_pan_guest_water && real_pan_home_water > 0.875 {
				ret = "3"
				cmt = "主队水位=客队水位,即时水位＞0.875队胜出"
			}
			if checkWaterNotChange(fid, cid) == true {
				ret = "3/0"
				cmt = "盘口、水位一直不变,双方能分胜负,<b>德甲主队胜概率大</b>"
			}
			if checkWaterNotChange(fid, cid) == false && checkIsGermanyJia(schedule_game_desc) == true {
				ret = "1/0"
				cmt = "<b>若为德甲，盘口不变而水位发生变化们一般是下盘胜出</b>,对应结果：1/0"
			}
		} else if home_pan_change_type == 1 {
			pre_cmt = "相对主队出现升盘（平半升半球或一球）初盘盘口=-0.25，即时盘口数值＜-0.25"
			if open_pan_home_water <= 0.875 {
				//				if real_pan_home_water > 0.875 && checkWaterIsDown(fid, cid) == false {
				if real_pan_home_water > 0.875 {
					ret = "3"
					cmt = "主队水位≤0.875,即时水位＞0.875,主队胜"
				} else if real_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "主队水位≤0.875,即时水位≤0.875,平或客队胜"
				}
			} else {
				ret = "1/0"
				cmt = "主队水位＞0.875,平或客队胜,<b>多为平局</b>"
			}
		} else if home_pan_change_type == -1 {
			pre_cmt = "相对主队出现降盘（初盘盘口=-0.25，即时盘口数值＞-0.25）"
			ret = "1/0"
			cmt = "其余情况:平或客队胜,对应结果：	1/0"
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water <= 0.875 {
					ret = "0"
					cmt = "主队水位≤0.875,即时水位≤0.875,客队胜"
				} else {
					ret = "1"
					cmt = "主队水位≤0.875,即时水位＞0.875,平"
				}
			}
		}
	case open_pan == (-0.5):
		if checkPanNotChange(fid, cid, open_pan) == true {
			pre_cmt = "一直保持半球盘（盘口数值为-0.5）不变"
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位≤0.875,主队胜,对应结果：3"
			} else {
				ret = "1/0"
				cmt = "主队水位＞0.875,平或客队胜,对应结果：1/0"
			}

			if checkWaterNotChange(fid, cid) == true {
				if open_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "盘口、水位一直不变,初盘水位-主队水位-≤0.875,平或客队胜,对应结果：1/0"
				} else {
					ret = "3"
					cmt = "盘口、水位一直不变,初盘水位-主队水位-＞0.875	主队胜,对应结果：3"
				}
			}
		} else if home_pan_change_type == 1 {
			pre_cmt = "相对主队出现升盘（半球升半一或半球升一球）初盘盘口=-0.5，即时盘口数值＜-0.5"
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 {
					ret = "3"
					cmt = "主队水位≤0.875	即时水位＞0.875,主队胜,对应结果：	3"
				} else {
					ret = "1"
					cmt = "主队水位≤0.875	即时水位≤0.875,平,对应结果：1"
				}
			} else {
				ret = "0"
				cmt = "主队水位＞0.875,客队胜,对应结果：0"
			}
		} else if home_pan_change_type == -1 {
			pre_cmt = "相对主队出现降盘（初盘盘口=-0.5，即时盘口数值＞-0.5）"
			ret = "1/0"
			cmt = "其余情况:平或客队胜,对应结果：	1/0"
			if open_pan_home_water > 0.875 {
				if real_pan_home_water <= 0.875 {
					ret = "3/1"
					cmt = "主队水位＞0.875,即时水位≤0.875,主胜或平,对应结果：3/1"
				} else {
					ret = "0"
					cmt = "主队水位＞0.875,即时水位＞0.875,客队胜,对应结果：0"
				}
			}
		}
	case open_pan == (-0.75):
		fmt.Println(open_pan, real_pan, open_pan_home_water, real_pan_home_water, home_pan_change_type)

		if checkPanNotChange(fid, cid, open_pan) == true {
			pre_cmt = "一直保持半球盘（盘口数值为-0.75）不变"
			if open_pan_home_water <= 0.875 {
				ret = "3"
				cmt = "主队水位≤0.875	主队胜,对应结果：3"
			} else {
				if real_pan_home_water <= 0.875 {
					ret = "1/0"
					cmt = "主队水位＞0.875即时水位≤0.875平或客队胜,对应结果：1/0"
				} else {
					ret = "3"
					cmt = "主队水位＞0.875即时水位＞0.875,主队胜,对应结果：	3"
				}
				if real_pan_home_water == open_pan_home_water {
					ret = "0"
					cmt = "主队水位＞0.875即时水位=初盘水位,客队胜,对应结果：0"
				}
			}

		} else if home_pan_change_type == 1 {
		} else if home_pan_change_type == -1 {
			pre_cmt = "相对主队出现降盘（初盘盘口=-0.75，即时盘口数值＞-0.75）"
			if open_pan_home_water <= 0.875 {
				if real_pan_home_water > 0.875 {
					ret = "1"
					cmt = "主队水位≤0.875	即时水位＞0.875,平	对应结果：1"
				} else {
					ret = "0"
					cmt = "主队水位≤0.875	即时水位≤0.875,客队胜,对应结果：0"
				}
			}
		}
	case open_pan <= -1:
		pre_cmt = "<=-1 其余情况"
		ret = "3"
		cmt = "初盘水位-主队水位＜0.875	主队胜,对应结果：3"
		if checkPanNotChange(fid, cid, open_pan) == true {
			pre_cmt = "一直保持一球或以上（盘口数值为≤-1）不变"
			if open_pan_home_water > 0.875 {
				ret = "3"
				cmt = "主队水位＞0.875主队胜,对应结果：3"
			}
		}
		if checkPanNotChange(fid, cid, open_pan) == false && open_pan <= (-1) && real_pan != open_pan {
						pre_cmt = "相对主队盘口变化（初盘盘口≤-1，即时盘口数值≠初盘盘口）"

			if open_pan_home_water > 0.875 && real_pan_home_water <= 0.875 {
				ret = "3/1"
				cmt = "主队水位＞0.875即时水位≤0.875,胜或平,对应结果：3/1"
			}
		}

		if open_pan < (-1.5) && checkPanNotChange(fid, cid, open_pan) == true {
			if open_pan_home_water <= 0.8 {
				ret = "3/0"
				cmt = "初盘盘口数值＜-1.5,初盘水位-主队水位＜=0.8,有爆冷可能	胜或负3/0"
			}
		}		
	default:
		//		fmt.Println("qita open:", open_pan)
		ret = ""
		cmt = ""
	}
	return ret, pre_cmt+":"+cmt
}
