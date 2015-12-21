package analyse

import (
	"fmt"
	"strings"

	"500kan/util/myinit"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func AnalysePanResult(pan_int_info map[string]int, pan_float_info map[string]float32, pan_string_info map[string]string) (ret string, cmt string) {
	open_pan := pan_float_info["open_pan"]
	open_pan_home_water := pan_float_info["open_home_water"]
	open_pan_guest_water := pan_float_info["open_guest_water"]
	real_pan := pan_float_info["real_pan"]
	real_pan_home_water := pan_float_info["home_real_water"]
	real_pan_guest_water := pan_float_info["guest_real_water"]

	home_pan_change_type := pan_int_info["home_water_change_type"]
	schedule_game_desc := pan_string_info["schedule_game_desc"]

	fid := pan_int_info["schedule_fenxi_id"]
	cid := pan_string_info["cid"]

	myinit.Myinit()
	engine = myinit.Engine
	fmt.Println("+++++++++++")
	fmt.Println(fid)
	switch {
	case open_pan == 0:
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == true {
				cmt = "一直保持平手盘（盘口数值为0）盘口、水位一直不变"
				if real_pan_home_water < real_pan_guest_water {
					ret = "3"
					cmt += "即时水位相对小的队胜出(主队)"
				} else if real_pan_home_water > real_pan_guest_water {
					ret = "0"
					cmt += "即时水位相对小的队胜出(客队)"
				} else {

				}
			} else {
				cmt = "一直保持平手盘（盘口数值为0）不变，只是水位有涨跌。"
				if open_pan_home_water < open_pan_guest_water {
					ret = "3"
					cmt += "主队水位	＜客队水位	主队胜"
				} else if open_pan_home_water > open_pan_guest_water {
					ret = "1/0"
					cmt += "主队水位	＞客队水位	平或客队胜"
					if real_pan_home_water < open_pan_home_water {
						ret = "1/0"
						cmt += "主队水位	＞客队水位	主队即时盘口水位小于初盘水位，多为平局"
					}
				} else {
					//open_pan_home_water == open_pan_guest_water
					ret = "1"
					cmt += "主队水位=	客队水位	平"

					if real_pan_home_water < real_pan_guest_water {
						ret = "3/" + ret
						cmt += "即时水位相对小的队胜出(主队)"

					} else if real_pan_home_water > real_pan_guest_water {
						ret = ret + "/0"
						cmt += "即时水位相对小的队胜出(客队)"
					} else {

					}
				}
			}

		} else if home_pan_change_type == "升" {
			cmt = "相对主队出现升盘（平手升平半、平手升半球）初盘盘口=0，即时盘口数值＜0"
			if open_pan_home_water < open_pan_guest_water {
				ret = "3"
				cmt += "主队水位	＜客队水位	主队胜"
			} else if open_pan_home_water > open_pan_guest_water {
				ret = "1/0"
				cmt += "主队水位	＞客队水位	平或客队胜  多为平局"
			}
		} else if home_pan_change_type == "降" {
			cmt = "相对主队出现降盘（初盘盘口=0，即时盘口数值＞0）"
			if real_pan_guest_water > real_pan_home_water {
				ret = "3/1"
				cmt += "客队水位	＞主队水位	主队胜或平"
			} else if real_pan_guest_water < real_pan_home_water {
				ret = "1/0"
				cmt += "客队水位	＜主队水位	平或客队胜"
			}
		}

		fmt.Println("open:", open_pan, ret, cmt)
	case open_pan == (-0.25):
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == true {
				cmt = "一直保持平半（盘口数值为-0.25）不变 盘口、水位一直不变"
				ret = "3/0"
				cmt += "双方能分胜负 德甲主队胜概率大"

			} else {
				cmt = "一直保持平半（盘口数值为-0.25）不变，只是水位有涨跌。"
				if checkIsGermanyJia(schedule_game_desc) == true {
					ret = "1/0"
					cmt += "为德甲，盘口不变而水位发生变化们一般是下盘胜出"
				} else {
					if open_pan_home_water < open_pan_guest_water {
						ret = "3"
						cmt += "主队水位	＜客队水位	主队胜"
					} else if open_pan_home_water > open_pan_guest_water {
						ret = "1/0"
						cmt += "主队水位	＞客队水位	平或客队胜"
					} else {
						//open_pan_home_water == open_pan_guest_water
						cmt += "即时水位相对大的队胜出"

						if real_pan_home_water > real_pan_guest_water {
							ret = "3"
							cmt += "即时水位相对大的队胜出(主队)"

						} else if real_pan_home_water < real_pan_guest_water {
							ret = "0"
							cmt += "即时水位相对大的队胜出(客队)"
						} else {

						}
					}
				}

			}
		} else if home_pan_change_type == "升" {
			cmt = "相对主队出现升盘（平半升半球或一球）初盘盘口=-0.25，即时盘口数值＜-0.25"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water > real_pan_guest_water && checkWaterIsDown(fid, cid) == true {
					ret = "3"
					cmt += "主队水位	<客队水位	即时水位＞客队水位并且水位持续下降	主队胜"
				} else if real_pan_home_water < real_pan_guest_water {
					ret = "1/0"
					cmt += "主队水位	<客队水位 	即时水位<客队水位	平或客队胜"
				}
			} else if open_pan_home_water > open_pan_guest_water {
				ret = "1/0"
				cmt += "主队水位	＞客队水位	平或客队胜 多为平局"
			}
		} else if home_pan_change_type == "降" {
			cmt = "相对主队出现降盘（初盘盘口=-0.25，即时盘口数值＞-0.25）"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water < real_pan_guest_water {
					ret = "0"
					cmt += "主队水位	<客队水位 	即时水位<客队水位 	客队胜"
				} else if real_pan_home_water > real_pan_guest_water {
					ret = "1"
					cmt += "主队水位	<客队水位 	即时水位＞客队水位 	平"
				}
			}

		}
	case open_pan == (-0.5):
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == true {
				cmt = "一直保持半球盘（盘口数值为-0.5）不变 盘口、水位一直不变."
				if open_pan_home_water < open_pan_guest_water {
					ret = "1/0"
					cmt += "初盘水位	主队水位	<客队水位	平或客队胜	对应结果：	1/0"

				} else if open_pan_home_water > open_pan_guest_water {
					ret = "3"
					cmt += "初盘水位	主队水位	＞客队水位	主队胜	对应结果：	3"
				}
			} else {
				cmt = "一直保持半球盘（盘口数值为-0.5）不变，只是水位有涨跌。"
				if open_pan_home_water < open_pan_guest_water {
					ret = "3"
					cmt += "主队水位	＜客队水位	主队胜"
				} else if open_pan_home_water > open_pan_guest_water {
					ret = "1/0"
					cmt += "主队水位	＞客队水位	平或客队胜"
				} else {

				}
			}
		} else if home_pan_change_type == "升" {
			cmt = "相对主队出现升盘（半球升半一或半球升一球）初盘盘口=-0.5，即时盘口数值＜-0.5"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water > real_pan_guest_water {
					ret = "3"
					cmt += "主队水位	<客队水位 	即时水位＞客队水位	主队胜	对应结果：	3"
				} else if real_pan_home_water < real_pan_guest_water {
					ret = "1"
					cmt += "主队水位	<客队水位 	即时水位<客队水位	平	对应结果：	1"
				}
			} else if open_pan_home_water > open_pan_guest_water {
				ret = "0"
				cmt += "主队水位	＞客队水位		客队胜	对应结果：	0"
			}
		} else if home_pan_change_type == "降" {
			jiang_flag := false
			cmt = "相对主队出现降盘（初盘盘口=-0.5，即时盘口数值＞-0.5）"
			if open_pan_home_water > open_pan_guest_water {
				if real_pan_home_water < real_pan_guest_water {
					ret = "3/1"
					cmt += "主队水位	＞客队水位	即时水位<客队水位	主胜或平	对应结果：	3、1"
					jiang_flag = true
				} else if real_pan_home_water > real_pan_guest_water {
					ret = "0"
					cmt += "主队水位	＞客队水位	即时水位＞客队水位	客队胜	对应结果：	0"
					jiang_flag = true

				}

			}
			if jiang_flag == false {
				ret = "1/0"
				cmt += "其余情况			平或客队胜	对应结果：	1/0"
			}
		}
	case open_pan == (-0.75):
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == false {
				cmt = "一直保持半球盘（盘口数值为-0.75）不变，只是水位有涨跌。"
				if open_pan_home_water < open_pan_guest_water {
					ret = "3"
					cmt += "主队水位	＜客队水位	主队胜"
				} else if open_pan_home_water > open_pan_guest_water {
					if real_pan_home_water < real_pan_guest_water {
						ret = "1/0"
						cmt += "主队水位	＞客队水位 即时水位＜客队水位	平或客队胜"
					} else if real_pan_home_water > real_pan_guest_water {
						ret = "3"
						cmt += "主队水位	＞客队水位 即时水位＜客队水位	主队胜"
					} else {

					}
					if real_pan_home_water == open_pan_home_water {
						ret = "0"
						cmt += "主队水位	＞客队水位 即时水位=初盘水位	客队胜"
					}

				} else {

				}
			}
		} else if home_pan_change_type == "升" {
		} else if home_pan_change_type == "降" {
			cmt = "相对主队出现降盘（初盘盘口=-0.75，即时盘口数值＞-0.75）"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water > real_pan_guest_water {
					ret = "1"
					cmt += "主队水位	<客队水位	 即时水位＞客队水位	平	对应结果：	1"
				} else if real_pan_home_water < real_pan_guest_water {
					ret = "0"
					cmt += "主队水位	<客队水位		即时水位< 客队水位	客队胜	对应结果：	0"
				}
			}

		}
		fmt.Println("-0.75: open:", open_pan)

	case open_pan <= -1:
		//	case -1.25:
		//	case -1.5:
		flag := false
		if checkPanNotChange(fid, cid, open_pan) == true && checkWaterNotChange(fid, cid) == false {
			if open_pan_home_water > open_pan_guest_water {
				ret = "3"
				cmt = "一直保持一球或以上（盘口数值为≤-1）不变，只是水位有涨跌。主队水位	＞客队水位	主队胜	对应结果：	3"
				flag = true
			}
		}
		if checkPanNotChange(fid, cid, open_pan) == false && open_pan <= (-1) && real_pan != open_pan {
			if open_pan_home_water > open_pan_guest_water && real_pan_home_water <= real_pan_guest_water {
				ret = "3/1"
				cmt = "相对主队盘口变化（初盘盘口≤-1，即时盘口数值≠初盘盘口）.主队水位	＞客队水位	即时水位<客队水位 胜或平"
				flag = true
			}
		}

		if open_pan < (-1.5) && checkPanNotChange(fid, cid, open_pan) == true {
			if open_pan_home_water < 0.8 {
				ret = "3/0"
				cmt += "初盘盘口数值＜-1.5），即时盘口数值不变 初盘水位	主队水位	＜0.8	有爆冷可能	胜或负 "
				flag = true
			}
		}

		if flag == false {
			if open_pan_home_water < open_pan_guest_water {
				//其他情况
				ret = "3"
				cmt += "其余情况 初盘水位	主队水位	＜客队水位	主队胜	对应结果：	3"
			}
		}
		fmt.Println("-1 open:", open_pan)
	default:
		fmt.Println("qita open:", open_pan)
		ret = ""
		cmt = ""

	}
	return ret, cmt
}

func AnalysePanResult2(open_pan float32, open_pan_home_water float32, open_pan_guest_water float32, real_pan float32, real_pan_home_water float32, real_pan_guest_water float32, home_pan_change_type string, schedule_game_desc string, fid string, cid string) (ret string, cmt string) {
	myinit.Myinit()
	engine = myinit.Engine
	fmt.Println("+++++++++++")
	fmt.Println(fid)
	switch {
	case open_pan == 0:
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == true {
				cmt = "一直保持平手盘（盘口数值为0）盘口、水位一直不变"
				if real_pan_home_water < real_pan_guest_water {
					ret = "3"
					cmt += "即时水位相对小的队胜出(主队)"
				} else if real_pan_home_water > real_pan_guest_water {
					ret = "0"
					cmt += "即时水位相对小的队胜出(客队)"
				} else {

				}
			} else {
				cmt = "一直保持平手盘（盘口数值为0）不变，只是水位有涨跌。"
				if open_pan_home_water < open_pan_guest_water {
					ret = "3"
					cmt += "主队水位	＜客队水位	主队胜"
				} else if open_pan_home_water > open_pan_guest_water {
					ret = "1/0"
					cmt += "主队水位	＞客队水位	平或客队胜"
					if real_pan_home_water < open_pan_home_water {
						ret = "1/0"
						cmt += "主队水位	＞客队水位	主队即时盘口水位小于初盘水位，多为平局"
					}
				} else {
					//open_pan_home_water == open_pan_guest_water
					ret = "1"
					cmt += "主队水位=	客队水位	平"

					if real_pan_home_water < real_pan_guest_water {
						ret = "3/" + ret
						cmt += "即时水位相对小的队胜出(主队)"

					} else if real_pan_home_water > real_pan_guest_water {
						ret = ret + "/0"
						cmt += "即时水位相对小的队胜出(客队)"
					} else {

					}
				}
			}

		} else if home_pan_change_type == "升" {
			cmt = "相对主队出现升盘（平手升平半、平手升半球）初盘盘口=0，即时盘口数值＜0"
			if open_pan_home_water < open_pan_guest_water {
				ret = "3"
				cmt += "主队水位	＜客队水位	主队胜"
			} else if open_pan_home_water > open_pan_guest_water {
				ret = "1/0"
				cmt += "主队水位	＞客队水位	平或客队胜  多为平局"
			}
		} else if home_pan_change_type == "降" {
			cmt = "相对主队出现降盘（初盘盘口=0，即时盘口数值＞0）"
			if real_pan_guest_water > real_pan_home_water {
				ret = "3/1"
				cmt += "客队水位	＞主队水位	主队胜或平"
			} else if real_pan_guest_water < real_pan_home_water {
				ret = "1/0"
				cmt += "客队水位	＜主队水位	平或客队胜"
			}
		}

		fmt.Println("open:", open_pan, ret, cmt)
	case open_pan == (-0.25):
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == true {
				cmt = "一直保持平半（盘口数值为-0.25）不变 盘口、水位一直不变"
				ret = "3/0"
				cmt += "双方能分胜负 德甲主队胜概率大"

			} else {
				cmt = "一直保持平半（盘口数值为-0.25）不变，只是水位有涨跌。"
				if checkIsGermanyJia(schedule_game_desc) == true {
					ret = "1/0"
					cmt += "为德甲，盘口不变而水位发生变化们一般是下盘胜出"
				} else {
					if open_pan_home_water < open_pan_guest_water {
						ret = "3"
						cmt += "主队水位	＜客队水位	主队胜"
					} else if open_pan_home_water > open_pan_guest_water {
						ret = "1/0"
						cmt += "主队水位	＞客队水位	平或客队胜"
					} else {
						//open_pan_home_water == open_pan_guest_water
						cmt += "即时水位相对大的队胜出"

						if real_pan_home_water > real_pan_guest_water {
							ret = "3"
							cmt += "即时水位相对大的队胜出(主队)"

						} else if real_pan_home_water < real_pan_guest_water {
							ret = "0"
							cmt += "即时水位相对大的队胜出(客队)"
						} else {

						}
					}
				}

			}
		} else if home_pan_change_type == "升" {
			cmt = "相对主队出现升盘（平半升半球或一球）初盘盘口=-0.25，即时盘口数值＜-0.25"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water > real_pan_guest_water && checkWaterIsDown(fid, cid) == true {
					ret = "3"
					cmt += "主队水位	<客队水位	即时水位＞客队水位并且水位持续下降	主队胜"
				} else if real_pan_home_water < real_pan_guest_water {
					ret = "1/0"
					cmt += "主队水位	<客队水位 	即时水位<客队水位	平或客队胜"
				}
			} else if open_pan_home_water > open_pan_guest_water {
				ret = "1/0"
				cmt += "主队水位	＞客队水位	平或客队胜 多为平局"
			}
		} else if home_pan_change_type == "降" {
			cmt = "相对主队出现降盘（初盘盘口=-0.25，即时盘口数值＞-0.25）"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water < real_pan_guest_water {
					ret = "0"
					cmt += "主队水位	<客队水位 	即时水位<客队水位 	客队胜"
				} else if real_pan_home_water > real_pan_guest_water {
					ret = "1"
					cmt += "主队水位	<客队水位 	即时水位＞客队水位 	平"
				}
			}

		}
	case open_pan == (-0.5):
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == true {
				cmt = "一直保持半球盘（盘口数值为-0.5）不变 盘口、水位一直不变."
				if open_pan_home_water < open_pan_guest_water {
					ret = "1/0"
					cmt += "初盘水位	主队水位	<客队水位	平或客队胜	对应结果：	1/0"

				} else if open_pan_home_water > open_pan_guest_water {
					ret = "3"
					cmt += "初盘水位	主队水位	＞客队水位	主队胜	对应结果：	3"
				}
			} else {
				cmt = "一直保持半球盘（盘口数值为-0.5）不变，只是水位有涨跌。"
				if open_pan_home_water < open_pan_guest_water {
					ret = "3"
					cmt += "主队水位	＜客队水位	主队胜"
				} else if open_pan_home_water > open_pan_guest_water {
					ret = "1/0"
					cmt += "主队水位	＞客队水位	平或客队胜"
				} else {

				}
			}
		} else if home_pan_change_type == "升" {
			cmt = "相对主队出现升盘（半球升半一或半球升一球）初盘盘口=-0.5，即时盘口数值＜-0.5"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water > real_pan_guest_water {
					ret = "3"
					cmt += "主队水位	<客队水位 	即时水位＞客队水位	主队胜	对应结果：	3"
				} else if real_pan_home_water < real_pan_guest_water {
					ret = "1"
					cmt += "主队水位	<客队水位 	即时水位<客队水位	平	对应结果：	1"
				}
			} else if open_pan_home_water > open_pan_guest_water {
				ret = "0"
				cmt += "主队水位	＞客队水位		客队胜	对应结果：	0"
			}
		} else if home_pan_change_type == "降" {
			jiang_flag := false
			cmt = "相对主队出现降盘（初盘盘口=-0.5，即时盘口数值＞-0.5）"
			if open_pan_home_water > open_pan_guest_water {
				if real_pan_home_water < real_pan_guest_water {
					ret = "3/1"
					cmt += "主队水位	＞客队水位	即时水位<客队水位	主胜或平	对应结果：	3、1"
					jiang_flag = true
				} else if real_pan_home_water > real_pan_guest_water {
					ret = "0"
					cmt += "主队水位	＞客队水位	即时水位＞客队水位	客队胜	对应结果：	0"
					jiang_flag = true

				}

			}
			if jiang_flag == false {
				ret = "1/0"
				cmt += "其余情况			平或客队胜	对应结果：	1/0"
			}
		}
	case open_pan == (-0.75):
		if checkPanNotChange(fid, cid, open_pan) == true {
			if checkWaterNotChange(fid, cid) == false {
				cmt = "一直保持半球盘（盘口数值为-0.75）不变，只是水位有涨跌。"
				if open_pan_home_water < open_pan_guest_water {
					ret = "3"
					cmt += "主队水位	＜客队水位	主队胜"
				} else if open_pan_home_water > open_pan_guest_water {
					if real_pan_home_water < real_pan_guest_water {
						ret = "1/0"
						cmt += "主队水位	＞客队水位 即时水位＜客队水位	平或客队胜"
					} else if real_pan_home_water > real_pan_guest_water {
						ret = "3"
						cmt += "主队水位	＞客队水位 即时水位＜客队水位	主队胜"
					} else {

					}
					if real_pan_home_water == open_pan_home_water {
						ret = "0"
						cmt += "主队水位	＞客队水位 即时水位=初盘水位	客队胜"
					}

				} else {

				}
			}
		} else if home_pan_change_type == "升" {
		} else if home_pan_change_type == "降" {
			cmt = "相对主队出现降盘（初盘盘口=-0.75，即时盘口数值＞-0.75）"
			if open_pan_home_water < open_pan_guest_water {
				if real_pan_home_water > real_pan_guest_water {
					ret = "1"
					cmt += "主队水位	<客队水位	 即时水位＞客队水位	平	对应结果：	1"
				} else if real_pan_home_water < real_pan_guest_water {
					ret = "0"
					cmt += "主队水位	<客队水位		即时水位< 客队水位	客队胜	对应结果：	0"
				}
			}

		}
		fmt.Println("-0.75: open:", open_pan)

	case open_pan <= -1:
		//	case -1.25:
		//	case -1.5:
		flag := false
		if checkPanNotChange(fid, cid, open_pan) == true && checkWaterNotChange(fid, cid) == false {
			if open_pan_home_water > open_pan_guest_water {
				ret = "3"
				cmt = "一直保持一球或以上（盘口数值为≤-1）不变，只是水位有涨跌。主队水位	＞客队水位	主队胜	对应结果：	3"
				flag = true
			}
		}
		if checkPanNotChange(fid, cid, open_pan) == false && open_pan <= (-1) && real_pan != open_pan {
			if open_pan_home_water > open_pan_guest_water && real_pan_home_water <= real_pan_guest_water {
				ret = "3/1"
				cmt = "相对主队盘口变化（初盘盘口≤-1，即时盘口数值≠初盘盘口）.主队水位	＞客队水位	即时水位<客队水位 胜或平"
				flag = true
			}
		}

		if open_pan < (-1.5) && checkPanNotChange(fid, cid, open_pan) == true {
			if open_pan_home_water < 0.8 {
				ret = "3/0"
				cmt += "初盘盘口数值＜-1.5），即时盘口数值不变 初盘水位	主队水位	＜0.8	有爆冷可能	胜或负 "
				flag = true
			}
		}

		if flag == false {
			if open_pan_home_water < open_pan_guest_water {
				//其他情况
				ret = "3"
				cmt += "其余情况 初盘水位	主队水位	＜客队水位	主队胜	对应结果：	3"
			}
		}
		fmt.Println("-1 open:", open_pan)
	default:
		fmt.Println("qita open:", open_pan)
		ret = ""
		cmt = ""

	}
	return ret, cmt
}

func checkIsGermanyJia(str string) (ret bool) {
	return strings.Contains(str, "德甲")
}

func checkWaterIsDown(fid string, cid string) (ret bool) {
	exist_up := new(myinit.LastPanLog)
	total, _ := engine.Where("last_home_water_change_type='water_up' AND schedule_fid=? AND company_cid=?", fid, cid).Count(exist_up)
	if total > 0 {
		return false
	}
	return true
}

func checkPanNotChange(fid string, cid string, pan_value float32) (ret bool) {
	exist_up := new(myinit.LastPanLog)
	//	if pan_value == 999 {
	//		total_pan_change, _ := engine.Where("open_pan!=last_pan AND schedule_fid=? AND company_cid=?", fid, cid).Count(exist_up)
	//	} else {
	total_pan_change, _ := engine.Where("open_pan!=last_pan AND schedule_fid=? AND company_cid=? AND open_pan=?", fid, cid, pan_value).Count(exist_up)
	//	}

	if total_pan_change > 0 {
		return false
	}
	return true
}

func checkWaterNotChange(fid string, cid string) (ret bool) {
	exist_up := new(myinit.LastPanLog)
	total_water_change, _ := engine.Where("open_home_water!=last_home_water AND schedule_fid=? AND company_cid=?", fid, cid).Count(exist_up)
	if total_water_change > 0 {
		return false
	}
	return true
}

func checkPanAndWaterNotChange(fid string, cid string, pan_value float32) (ret bool) {
	if checkPanNotChange(fid, cid, pan_value) == false {
		return false
	}
	if checkWaterNotChange(fid, cid) == false {
		return false
	}
	return true
}
