package check

import (
	"gsCheck/model"
	"gsCheck/utils"
	"strconv"
	"strings"
)

func IsIntNum(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	if strings.Contains(str, ".00") {
		str = strings.ReplaceAll(str, ".00", "")
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "必须是整数",
		}
	}
	return true, model.ErrInfo{}
}

func IsDoubleNum(str string) (bool, model.ErrInfo) {

	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "请填写金额类型",
		}
	}
	if strings.Contains(str, ".") {
		if len(str)-strings.Index(str, ".")-1 > 2 {
			return false, model.ErrInfo{
				ErrorMsg: "  异常！错误值-> " + str,
				FixMsg:   "金额小数点后只能有两位!",
			}
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectRate(str string) (bool, model.ErrInfo) {

	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	rate, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "请填写数字类型",
		}
	}
	if rate > 100 || rate < 0 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "值不可大于100%或小于0",
		}
	}
	// if strings.Contains(str, ".") {
	// 	if len(str)-strings.Index(str, ".")-1 > 2 {
	// 		return false, model.ErrInfo{
	// 			ErrorMsg: "  异常！错误值-> " + str,
	// 			FixMsg:   "小数点后只能有两位!",
	// 		}
	// 	}
	// }
	return true, model.ErrInfo{}
}

func IsCorrectName(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectComeFrom(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	var arrs = []string{"购置", "自建", "投资人投入", "接受捐赠", "盘盈", "内部销售"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}

}

func IsCorrectManageType(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	var arrs = []string{"财务管理中心-税控类设备",
		"物业安全中心-保洁设备",
		"物业安全中心-安保设备",
		"物业安全中心-弱电设备",
		"物业安全中心-强电设备",
		"物业安全中心-水暖设备",
		"办公室-办公设备",
		"办公室-运输设备",
		"研发中心-信息设备",
		"超市事业部-经营设备",
		"超市事业部-营运设备",
		"招商中心-百货经营设备",
		"运营中心-百货营运设备"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}
}

func IsCorrectCWType(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	var arrs = []string{"运输工具",
		"营业设备",
		"办公设备",
		"低值易耗品",
		"工会设备",
		"电子设备",
		"机器设备",
		"房屋建筑物",
		"软件",
		"土地",
		"商标",
		"专利",
		"网络资产",
		"其他设备"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}

}

func IsCorrectStatus(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	var arrs = []string{"在用",
		"在库",
		"闲置",
		"报废",
		"报损",
		"在途",
		"已售",
		"已拆分"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}
}

func IsCorrectJiTi(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	if str == "是" || str == "否" {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "修改提示: 只能填写 是 或者 否",
		}
	}
}

func IsCorrectBuyDate(str string) (bool, model.ErrInfo) {
	if len(str) != 8 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "日期格式为20230501(若修改后仍提示错误,请将日期列的单元格类型修改为文本)",
		}
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "日期格式为20230501(若修改后仍提示错误,请将日期列的单元格类型修改为文本)",
		}
	}
	return true, model.ErrInfo{}
}

func IsCorrectZJSF(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	var arrs = []string{"平均年限法",
		"工作量法",
		"双倍余额递",
		"年数总和法",
		"新准则",
		"一次性摊销",
		"减值或变动后的平均年限法"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "正确内容为 " + strings.Join(arrs, " , "),
		}
	}
}

func IsCorrectMKT(str string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(str, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "单位名称不可为空",
		}
	}
	rows, err := utils.DB.Raw("select 1 from depts where mkt=? limit 1", str).Rows()
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "门店查询失败(SQL查询失败)!",
		}
	}
	if rows.Next() {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + str,
			FixMsg:   "没有找到该门店!(请填入提供的组织架构中的门店名称)",
		}
	}

}

func IsCorrectDept(dept, mkt string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(dept, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	rows, err := utils.DB.Raw("select 1 from depts where name=? and mkt=? limit 1", dept, mkt).Rows()
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + dept,
			FixMsg:   "部门查询失败(SQL查询失败)!",
		}
	}
	if rows.Next() {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + dept,
			FixMsg:   "没有找到此部门!(请填入提供的组织架构中的部门名称)",
		}
	}
}

func IsCorrectUser(name, mkt string) (bool, model.ErrInfo) {
	if len(strings.ReplaceAll(name, " ", "")) < 1 {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！",
			FixMsg:   "不可为空",
		}
	}
	rows, err := utils.DB.Raw("select 1 from users where name=? and mkt=? limit 1", name, mkt).Rows()
	if err != nil {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + name,
			FixMsg:   "用户查询失败(SQL查询失败)!",
		}
	}
	if rows.Next() {
		return true, model.ErrInfo{}
	} else {
		return false, model.ErrInfo{
			ErrorMsg: "  异常！错误值-> " + name,
			FixMsg:   "没有找到此用户!",
		}
	}
}
