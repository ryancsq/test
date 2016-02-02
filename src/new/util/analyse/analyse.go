package analyse

import (
	"strings"

	"new/util/myinit"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine



func checkIsGermanyJia(str string) (ret bool) {
	return strings.Contains(str, "德甲")
}

func checkWaterIsDown(fid int, cid string) (ret bool) {
	exist_up := new(myinit.AsiaPanLog)
	total, _ := engine.Where("home_water_change_type=1 AND schedule_fenxi_id=? AND company_id=?", fid, cid).Count(exist_up)
	if total > 0 {
		return false
	}
	return true
}

func checkPanNotChange(fid int, cid string, pan_value float32) (ret bool) {
	exist_up := new(myinit.AsiaPanLog)
	
	total_pan_change, _ := engine.Where("open_pan!=real_pan AND schedule_fenxi_id=? AND company_id=? AND open_pan=?", fid, cid, pan_value).Count(exist_up)

	if total_pan_change > 0 {
		return false
	}
	return true
}

func checkWaterNotChange(fid int, cid string) (ret bool) {
	exist_up := new(myinit.AsiaPanLog)
	total_water_change, _ := engine.Where("open_home_water!=real_home_water AND schedule_fenxi_id=? AND company_id=?", fid, cid).Count(exist_up)
	if total_water_change > 0 {
		return false
	}
	return true
}

func checkPanAndWaterNotChange(fid int, cid string, pan_value float32) (ret bool) {
	if checkPanNotChange(fid, cid, pan_value) == false {
		return false
	}
	if checkWaterNotChange(fid, cid) == false {
		return false
	}
	return true
}
