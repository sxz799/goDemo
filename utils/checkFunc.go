package utils

import (
	"errors"
	"strconv"
	"strings"
)

func IsIntNum(str string) error {
	if strings.Contains(str, ".00") {
		str = strings.ReplaceAll(str, ".00", "")
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 必须是整数")
	}
	return nil
}

func IsCorrectComeFrom(str string) error {
	var arrs = []string{"购置", "自建", "投资人投入", "接受捐赠", "盘盈", "内部销售"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return nil
	} else {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 正确内容为 " + strings.Join(arrs, ","))
	}

}

func IsCorrectManageType(str string) error {

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
		return nil
	} else {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 正确内容为 " + strings.Join(arrs, ","))
	}
}

func IsCorrectCWType(str string) error {
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
		return nil
	} else {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 正确内容为 " + strings.Join(arrs, ","))
	}

}

func IsCorrectStatus(str string) error {

	var arrs = []string{"在用",
		"在库",
		"闲置",
		"报废",
		"报损",
		"在途",
		"已售",
		"已拆分"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return nil
	} else {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 正确内容为 " + strings.Join(arrs, ","))
	}
}

func IsCorrectJiTi(str string) error {
	if str == "是" || str == "否" {
		return nil
	} else {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 只能填写 是 或者 否")
	}
}

func IsCorrectBuyDate(str string) error {
	if len(str) != 8 {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: ")
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: ")
	}
	return nil
}

func IsCorrectZJSF(str string) error {

	var arrs = []string{"平均年限法",
		"工作量法",
		"双倍余额递",
		"年数总和法",
		"新准则",
		"一次性摊销",
		"减值或变动后的平均年限法"}
	if strings.Contains(strings.Join(arrs, ","), str) {
		return nil
	} else {
		return errors.New("异常！错误值-> " + str + "  | 修改提示: 正确内容为 " + strings.Join(arrs, ","))
	}
}

func IsDoubleNum(str string) error {
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return errors.New("异常！" + str)
	}
	if strings.Contains(str, ".") {
		if len(str)-strings.Index(str, ".")-1 > 2 {
			return errors.New("异常！错误值-> " + str + "  | 修改提示: ")
		}
	}
	return nil
}

func IsCorrectMKT(str string) error {
	return nil
}

func IsCorrectDept(str string) error {
	return nil
}

func IsCorrectUser(str string) error {
	return nil
}
