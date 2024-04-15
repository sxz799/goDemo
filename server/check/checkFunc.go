package check

import (

	"gsCheck/config"
	"gsCheck/model"
	"gsCheck/utils"
	"strconv"
	"strings"
)

func checkNull(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return true, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	} else {
		return false, model.ErrInfo{}
	}
}

func IsIntNum(str string) (bool, model.ErrInfo) {

	if isNull, errInfo := checkNull(str); isNull {
		return false, errInfo
	}

	if strings.Contains(str, ".00") {
		str = strings.ReplaceAll(str, ".00", "")
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "必须是整数",
		}
	}
	if num < 0 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "不可以小于0",
		}
	}
	return true, model.ErrInfo{}
}

func IsDoubleNum(str string) (bool, model.ErrInfo) {

	if isNull, errInfo := checkNull(str); isNull {
		return false, errInfo
	}
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "请填写金额类型 如果有千分位逗号分隔,请去掉',' [tips: ctrl + f 批量替换 将 , 批量替换为空(`替换为`那里留空，不是替换成空格)] ",
		}
	}
	if strings.Contains(str, ".") {
		if len(str)-strings.Index(str, ".")-1 > 2 {
			return false, model.ErrInfo{
				ErrorMsg: "  异常！错误值->" + str,
				FixMsg:   "金额小数点后只能有两位!",
			}
		}
	}
	return true, model.ErrInfo{}
}
func IsCorrectPlace(str string) (bool, model.ErrInfo) {

	if len(str) > 100 {
		return false, model.ErrInfo{
			ErrorMsg: "  长度异常！",
			FixMsg:   "不可超过50个汉字",
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectMemo(str string) (bool, model.ErrInfo) {

	if len(str) > 200 {
		return false, model.ErrInfo{
			ErrorMsg: "  长度异常！",
			FixMsg:   "不可超过100个汉字",
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectRate(str string) (bool, model.ErrInfo) {

	if isNull, errInfo := checkNull(str); isNull {
		return false, errInfo
	}
	rate, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "请填写数字类型",
		}
	}
	if rate > 100 || rate < 0 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "值不可大于100%或小于0",
		}
	}
	// if strings.Contains(str, ".") {
	// 	if len(str)-strings.Index(str, ".")-1 > 2 {
	// 		return false, model.ErrInfo{
	// 			ErrorMsg: "  异常！错误值->" + str,
	// 			FixMsg:   "小数点后只能有两位!",
	// 		}
	// 	}
	// }
	return true, model.ErrInfo{}
}

func IsCorrectName(str string) (bool, model.ErrInfo) {
	if isNull, errInfo := checkNull(str); isNull {
		return false, errInfo
	}
	return true, model.ErrInfo{}
}

func IsCorrectComeFrom(str string) (bool, model.ErrInfo) {
	var arrs = config.CapComeFrom
	if len(str) > 0 && strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}

}

func IsCorrectManageType(str string) (bool, model.ErrInfo) {
	var arrs = config.CapManageType
	if len(str) > 0 && strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}
}

func IsCorrectStatus(str string) (bool, model.ErrInfo) {
	var arrs = config.CapStatus
	if len(str) > 0 && strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}
}

func IsCorrectJiTi(str string) (bool, model.ErrInfo) {
	if str == "是" || str == "否" {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "修改提示: 只能填写 是 或者 否",
		}
	}
}

func IsCorrectBuyDate(str string) (bool, model.ErrInfo) {
	if len(strings.Split(str, "-")) != 3 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "日期格式为2023-05-01(若修改后仍提示错误,请将日期列的单元格类型修改为文本)",
		}
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "日期格式为2023-05-01(若修改后仍提示错误,请将日期列的单元格类型修改为文本)",
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectZJSF(str string) (bool, model.ErrInfo) {
	var arrs = config.CapZJSF
	if len(str) > 0 && strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}
}

func IsCorrectMKT(str string) (bool, model.ErrInfo) {
	if isNull, errInfo := checkNull(str); isNull {
		return false, errInfo
	}

	var o model.Organization
	utils.DB.Where("mkt=?", str).First(&o)
	if o.Mkt == "" {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + str,
			FixMsg:   "没有找到该门店!(请填入提供的组织架构中的门店名称)",
		}
	}
	return true, model.ErrInfo{}

}

func IsCorrectDept(dept, mkt string) (bool, model.ErrInfo) {
	if isNull, errInfo := checkNull(dept); isNull {
		return false, errInfo
	}
	var o model.Organization
	utils.DB.Where("mkt=? and dept=?", mkt, dept).First(&o)
	if o.Dept == "" {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + dept,
			FixMsg:   "没有找到该部门!(请填入提供的组织架构中的部门名称)",
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectUser(name, mkt string) (bool, model.ErrInfo) {
	if isNull, errInfo := checkNull(name); isNull {
		return false, errInfo
	}
	var u model.User
	utils.DB.Where("mkt=? and name=?", mkt, name).First(&u)
	if u.Name == "" {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值->" + name,
			FixMsg:   "没有找到该用户!(请填入提供的组织架构中的用户姓名)",
		}
	}
	return true, model.ErrInfo{}
}

func CheckCWType(str string) (bool, model.ErrInfo, string) {
	switch {
	case len(str) > 0 && strings.Contains(strings.Join(config.CapCWCat01, ","), str):
		return true, model.ErrInfo{}, "固定资产"
	case len(str) > 0 && strings.Contains(strings.Join(config.CapCWCat02, ","), str):
		return true, model.ErrInfo{}, "低值易耗品"
	case len(str) > 0 && strings.Contains(strings.Join(config.CapCWCat03, ","), str):
		return true, model.ErrInfo{}, "无形资产"
	case len(str) > 0 && strings.Contains(strings.Join(config.CapCWCat04, ","), str):
		return true, model.ErrInfo{}, "投资性房地产"
	case len(str) > 0 && strings.Contains(strings.Join(config.CapCWCat05, ","), str):
		return true, model.ErrInfo{}, "长期待摊费用"
	case len(str) > 0 && strings.Contains(strings.Join(config.CapCWCat06, ","), str):
		return true, model.ErrInfo{}, "代管资产"
	}

	return false, model.ErrInfo{
		ErrorMsg: "  异常！错误值->" + str,
		FixMsg: "正确内容为 " +
			strings.Join(config.CapCWCat01, " , ") + " , " +
			strings.Join(config.CapCWCat02, " , ") + " , " +
			strings.Join(config.CapCWCat03, " , ") + " , " +
			strings.Join(config.CapCWCat04, " , ") + " , " +
			strings.Join(config.CapCWCat05, " , ") + " , " +
			strings.Join(config.CapCWCat06, " , "),
	}, ""

}
