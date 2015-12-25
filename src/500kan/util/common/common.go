package common

import (
//	"fmt"
	"strconv"
	"time"

	"github.com/guotie/gogb2312"
)

func ConvToFloat32(float_string string) (float_val float32) {
//	fmt.Println(float_string)
	float_string_32, _ := strconv.ParseFloat(float_string, 32)
//		fmt.Println(float_string_32)

	return float32(float_string_32)

}

func ConvToGB(str string) (res_str string) {
	conv_str, _, _, _ := gogb2312.ConvertGB2312String(str)
	return conv_str
}

func CompareDate(time1 string, time2 string) (ret bool) {
	t1, _ := time.Parse("2006-01-02", time1)
	t2, _ := time.Parse("2006-01-02", time2)
	if t1.Before(t2) {
		return true
	}
	return false
}

func compareDateTime(time1 string, time2 string) (ret bool) {
	//先把时间字符串格式化成相同的时间类型
	t1, _ := time.Parse("2006-01-02 15:04:05", time1)
	t2, _ := time.Parse("2006-01-02 15:04:05", time2)
	if t1.Before(t2) {
		return true
	}
	return false
}
