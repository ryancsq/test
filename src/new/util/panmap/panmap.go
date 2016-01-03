package panmap

import (
	"500kan/util/myinit"
	"fmt"
)

func Add(pan_desc string,pan_value float32) {
	myinit.Myinit()
	has := CheckExists(pan_desc,pan_value)
	fmt.Println(has)

	if has {
		fmt.Println(pan_desc , " => " ,pan_value , "已存在！")
	} else {
		PanMap := new(myinit.PanMap)
		PanMap.PanValue = pan_value
		PanMap.PanDesc = pan_desc
		affected, _ := myinit.Engine.Insert(PanMap)
		fmt.Println(affected)
	}
}

func CheckExists(pan_desc string,pan_value float32) (has bool) {
	exist_panmap := new(myinit.PanMap)
	has, _ = myinit.Engine.Where("pan_desc=? AND pan_value=? ", pan_desc, pan_value).Get(exist_panmap)
	fmt.Println(has)
	return has
}

func GetPanValueByPanDesc(pan_desc string) (has bool,pan_value float32) {
	exist_panmap := new(myinit.PanMap)
	has, _ = myinit.Engine.Where("pan_desc=?", pan_desc).Get(exist_panmap)
	fmt.Println(has)
	
	return has,exist_panmap.PanValue
}